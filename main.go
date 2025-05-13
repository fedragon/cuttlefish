package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/fedragon/cuttlefish/git"
	"github.com/fedragon/cuttlefish/globals"
	"github.com/fedragon/cuttlefish/paths"
	"github.com/fedragon/cuttlefish/ssh"
	"github.com/knadh/koanf"
	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/file"
)

var k = koanf.New(".")

// ConfigFileName is the name of the config file used by this script
const ConfigFileName = ".cuttlefish.yaml"

var sysTmpDir = os.Getenv("TMPDIR")

// TmpFile is the tmp file where the script stores the path to the previous config, if any
var TmpFile = filepath.Join(sysTmpDir, "cuttlefish-lastvisited")

func main() {
	config, err := find(os.Getenv("PWD"))
	if err != nil {
		// signal that there's nothing to source
		os.Exit(1)
	}

	previousConfig, err := os.ReadFile(TmpFile)
	if err != nil {
		previousConfig = []byte("")
	}

	if string(previousConfig) == config {
		// nothing to do
		os.Exit(1)
	}

	if err := k.Load(file.Provider(config), yaml.Parser()); err != nil {
		log.Fatal(err)
	}

	var cmds []string

	if strings.TrimSpace(string(previousConfig)) == "" {
		cmds = withoutPreviousConfig()
	} else {
		cmds, err = withPreviousConfig(string(previousConfig))
		if err != nil {
			log.Fatal(err)
		}
	}

	if err = os.WriteFile(TmpFile, []byte(config), 0644); err != nil {
		log.Fatal(err)
	}

	fmt.Println(strings.Join(cmds, "\n"))
}

func find(pwd string) (string, error) {
	absPwd, err := filepath.Abs(pwd)
	if err != nil {
		return "", err
	}

	if !strings.HasPrefix(absPwd, os.Getenv("HOME")) {
		return "", errors.New("only checking in home dir")
	}

	fullPath := filepath.Join(absPwd, ConfigFileName)

	_, err = os.Stat(fullPath)
	if os.IsNotExist(err) {
		if os.Getenv("HOME") == absPwd {
			return "", fmt.Errorf("no %v found in any parent (stopped at %v)", ConfigFileName, os.Getenv("HOME"))
		}

		return find(filepath.Dir(absPwd))
	}

	return fullPath, nil
}

func withoutPreviousConfig() []string {
	cmds := paths.Set(k.Strings("user_paths"), nil)
	cmds = append(cmds, ssh.Set(k.Strings("ssh_identities"), nil)...)
	cmds = append(cmds, globals.Set(k.StringMap("global_variables"), nil)...)
	return append(cmds, git.SetName(k.String("git_config.name")), git.SetEmail(k.String("git_config.email")))
}

func withPreviousConfig(previous string) ([]string, error) {
	var pk = koanf.New(".")
	if err := pk.Load(file.Provider(previous), yaml.Parser()); err != nil {
		return nil, err
	}

	cmds := paths.Set(k.Strings("user_paths"), pk.Strings("user_paths"))
	cmds = append(cmds, git.SetName(k.String("git_config.name")), git.SetEmail(k.String("git_config.email")))
	cmds = append(cmds, ssh.Set(k.Strings("ssh_identities"), pk.Strings("ssh_identities"))...)

	vars := pk.StringMap("global_variables")
	keys := make([]string, 0, len(vars))

	for key := range vars {
		keys = append(keys, key)
	}

	return append(cmds, globals.Set(k.StringMap("global_variables"), keys)...), nil
}
