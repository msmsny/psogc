package psogc

import (
	"bytes"
	"embed"
	"fmt"
	"io"
	"io/fs"
	"strings"

	"gopkg.in/yaml.v2"
)

type CharacterConfig struct {
	Characters []*Character `yaml:"characters,flow"`
}

type Character struct {
	Name     string    `yaml:"name"`
	Statuses []*Status `yaml:"statuses,flow"`
}

type Status struct {
	Level        int     `yaml:"lv"`
	HP           int     `yaml:"hp"`
	TP           int     `yaml:"tp"`
	Attack       int     `yaml:"atp"`
	Defense      int     `yaml:"dfp"`
	MindStrength int     `yaml:"mst"`
	Accuracy     float32 `yaml:"ata"`
	Evasion      int     `yaml:"evp"`
}

//go:embed config/*
var configFS embed.FS

func LoadConfig(wd string) (*CharacterConfig, error) {
	config := &CharacterConfig{}
	return config, fs.WalkDir(configFS, "config", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return fmt.Errorf("fs.WalkDir: %s", err)
		}
		if d.IsDir() || !strings.HasPrefix(path, "config/character") {
			return nil
		}

		rawFile, err := configFS.ReadFile(path)
		if err != nil {
			return fmt.Errorf("config.ReadFile: %s", err)
		}
		configByCharacter := &CharacterConfig{}
		if err := yaml.NewDecoder(bytes.NewReader(rawFile)).Decode(configByCharacter); err != nil {
			if err == io.EOF {
				return nil
			}
			return fmt.Errorf("yaml decode failed: %s", err)
		}
		config.Characters = append(config.Characters, configByCharacter.Characters...)

		return nil
	})
}
