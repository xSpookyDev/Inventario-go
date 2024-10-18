package settings

import (
	_ "embed"

	"github.com/rs/zerolog/log"
	"gopkg.in/yaml.v3"
)

//go:embed settings.yaml
var settingsFile []byte

type (
	DatabaseConfig struct {
		Host     string `yaml:"host"`
		Port     int    `yaml:"port"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		Name     string `yaml:"name"`
	}

	Settings struct {
		Port string         `yaml:"port"`
		DB   DatabaseConfig `yaml:"database"`
	}
)

func New() (*Settings, error) {
	var s Settings
	log.Info().Msg("Cargando Configuraciones de Settings")
	err := yaml.Unmarshal(settingsFile, &s)
	if err != nil {
		log.Err(err).Msg("Error al cargar el archivo")
		return nil, err
	}

	log.Info().Msg("Configuraciones cargadas con exito")
	return &s, nil
}
