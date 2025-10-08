package validate

import (
	"fmt"
	"regexp"
)

// Validate ensures that the input string does not contain any shell injection vulnerabilities.
//
// Current protections:
// - (\(.+\)+) - Parentheses/subshells like "(command)"
// - (.*;\s*(and|or)\s+.+) - Semicolon followed by "and"/"or" like "cmd; and cmd2"
//
// SECURITY NOTE: Additional Fish shell injection vectors that are NOT currently protected:
//
// 1. Command substitution:
//   - $(command) - Command substitution
//   - `command` - Backtick command substitution
//
// 2. Pipes and redirection:
//   - cmd1 | cmd2 - Pipe output to another command
//   - cmd1 |& cmd2 - Pipe stdout and stderr
//   - cmd > file, cmd >> file, cmd < file - File redirection
//   - cmd 2> file, cmd &> file - Error redirection
//
// 3. Command chaining:
//   - cmd1 && cmd2 - Logical AND (run cmd2 if cmd1 succeeds)
//   - cmd1 || cmd2 - Logical OR (run cmd2 if cmd1 fails)
//   - cmd1; cmd2 - Sequential execution (without and/or)
//
// 4. Background processes:
//   - cmd & - Run command in background
//
// 5. Variable expansion:
//   - $var, ${var} - Variable expansion
//   - $(...) - Command substitution (different from above)
//
// 6. Glob patterns:
//   - *, ?, [...] - Wildcard expansion
//
// 7. Brace expansion:
//   - {a,b,c} - Brace expansion
//   - {1..10} - Range expansion
//
// 8. Here documents/strings:
//   - <<EOF, <<<string - Here documents
//
// 9. Process substitution:
//   - <(command), >(command) - Process substitution
//
// 10. Fish-specific metacharacters:
//   - # - Comments (rest of line ignored)
//   - \ - Escape character
//   - ' " - Quote characters that could break string parsing
func Validate(input string) error {
	var invalid = regexp.MustCompile(`(\(.+\)+)|(.*;\s*(and|or)\s+.+)`)

	if invalid.Match([]byte(input)) {
		return fmt.Errorf("input string cannot initialize subshells or run commands: %v", input)
	}

	// Additional critical injection patterns that must be blocked
	additionalPatterns := []struct {
		regex *regexp.Regexp
		desc  string
	}{
		// Command substitution
		{regexp.MustCompile(`\$\(.+\)`), "command substitution"},
		{regexp.MustCompile("`[^`]+`"), "backtick command substitution"},

		// Command chaining (most dangerous)
		{regexp.MustCompile(`.*&&.*`), "command chaining"},
		{regexp.MustCompile(`.*\|\|.*`), "command chaining"},

		// Pipes and redirection (critical for security)
		{regexp.MustCompile(`.*\|[^|&].*`), "pipe operator"},
		{regexp.MustCompile(`.*[<>]{1,2}.*`), "redirection"},

		// Background processes
		{regexp.MustCompile(`.*&\s*$`), "background process"},

		// Process substitution
		{regexp.MustCompile(`<\(.+\)`), "process substitution"},
		{regexp.MustCompile(`>\(.+\)`), "process substitution"},

		// Variable expansion (dangerous patterns)
		{regexp.MustCompile(`\$\{[^}]*\}`), "variable expansion"},

		// Comments that could hide commands
		{regexp.MustCompile(`.*;.*#.*`), "comment injection"},
	}

	for _, pattern := range additionalPatterns {
		if pattern.regex.MatchString(input) {
			return fmt.Errorf("input string cannot initialize subshells or run commands: %v", input)
		}
	}

	return nil
}
