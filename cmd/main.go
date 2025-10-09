package main

import (
	"errors"
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/fedragon/cuttlefish/internal/git"
	"github.com/fedragon/cuttlefish/internal/globals"
	"github.com/fedragon/cuttlefish/internal/paths"
	"github.com/fedragon/cuttlefish/internal/ssh"
	"github.com/knadh/koanf"
	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/file"
)

// configFileName is the name of the config file used by this script
const configFileName = ".cuttlefish.yaml"

// lastSeenFileName is the name of the file where the script stores the path to the last-seen config, if any
const lastSeenFileName = "cuttlefish-lastseen"

func main() {
	homeDir, ok := os.LookupEnv("HOME")
	if !ok {
		log.Fatal("HOME environment variable not set")
	}
	tmpDir, ok := os.LookupEnv("TMPDIR")
	if !ok {
		log.Fatal("TMPDIR environment variable not set")
	}
	pwdDir, ok := os.LookupEnv("PWD")
	if !ok {
		log.Fatal("PWD environment variable not set")
	}

	currentConfigPath, err := find(homeDir, pwdDir)
	if err != nil {
		// signal that there's nothing to source
		os.Exit(1)
	}

	tmpFile := filepath.Join(tmpDir, lastSeenFileName)

	stat, err := os.Stat(tmpFile)
	if err != nil {
		if os.IsNotExist(err) {
			// ignore error: the last seen config will be nil
		} else {
			log.Fatal(err)
		}
	} else if stat.Mode().Perm() != 0644 {
		log.Fatal(fmt.Errorf("invalid permissions for %v: they need to be 0644", tmpFile))
	}

	lastSeenConfigBytes, err := os.ReadFile(tmpFile)
	if err != nil {
		// ignore error: the last seen config will be nil
	}
	lastSeenConfigPath := string(lastSeenConfigBytes)

	if lastSeenConfigPath == currentConfigPath {
		// nothing to do
		os.Exit(0)
	}

	currentConfig := koanf.New(".")
	if err := currentConfig.Load(file.Provider(currentConfigPath), yaml.Parser()); err != nil {
		log.Fatal(err)
	}

	var commands []string
	if len(strings.TrimSpace(lastSeenConfigPath)) == 0 {
		commands = fromCurrentConfig(currentConfig)
	} else {
		lastSeenConfig := koanf.New(".")
		if err := lastSeenConfig.Load(file.Provider(lastSeenConfigPath), yaml.Parser()); err != nil {
			log.Fatal(err)
		}

		commands, err = fromCurrentAndLastSeenConfig(currentConfig, lastSeenConfig)
		if err != nil {
			log.Fatal(err)
		}
	}

	if err := os.WriteFile(tmpFile, []byte(currentConfigPath), 0644); err != nil {
		log.Fatal(err)
	}

	fmt.Println(strings.Join(commands, "\n"))
}

func find(homeDir, pwd string) (string, error) {
	absPwd, err := filepath.Abs(pwd)
	if err != nil {
		return "", err
	}

	if !strings.HasPrefix(absPwd, homeDir) {
		return "", errors.New("cannot check outside home directory")
	}

	fullPath := filepath.Join(absPwd, configFileName)
	stat, err := os.Stat(fullPath)
	if errors.Is(err, fs.ErrNotExist) {
		if homeDir == absPwd {
			return "", fmt.Errorf("no %v found in any parent (stopped at %v)", configFileName, homeDir)
		}

		return find(homeDir, filepath.Dir(absPwd))
	}

	if stat.Mode().Perm() != 0644 {
		return "", fmt.Errorf("invalid permissions for %v: they need to be 0644", fullPath)
	}

	return fullPath, nil
}

func fromCurrentConfig(config *koanf.Koanf) []string {
	cmds := paths.Set(config.Strings("user_paths"), nil)
	cmds = append(cmds, ssh.Set(config.Strings("ssh_identities"), nil)...)
	cmds = append(cmds, globals.Set(config.StringMap("global_variables"), nil)...)

	if name, ok := git.SetName(config.String("git_config.name")); ok {
		cmds = append(cmds, name)
	}
	if email, ok := git.SetEmail(config.String("git_config.email")); ok {
		cmds = append(cmds, email)
	}

	return cmds
}

func fromCurrentAndLastSeenConfig(current *koanf.Koanf, lastSeen *koanf.Koanf) ([]string, error) {
	cmds := paths.Set(current.Strings("user_paths"), lastSeen.Strings("user_paths"))
	cmds = append(cmds, ssh.Set(current.Strings("ssh_identities"), lastSeen.Strings("ssh_identities"))...)

	if name, ok := git.SetName(current.String("git_config.name")); ok {
		cmds = append(cmds, name)
	}
	if email, ok := git.SetEmail(current.String("git_config.email")); ok {
		cmds = append(cmds, email)
	}

	vars := lastSeen.StringMap("global_variables")
	keys := make([]string, 0, len(vars))
	for key := range vars {
		keys = append(keys, key)
	}

	return append(cmds, globals.Set(current.StringMap("global_variables"), keys)...), nil
}
