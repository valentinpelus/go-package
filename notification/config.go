package notification

import (
	"os"

	"github.com/rs/zerolog/log"
	"gopkg.in/yaml.v2"
)

type Config struct {
	SlackClient []SlackClient `yaml:"slackClient"`
}

type SlackClient struct {
	webhookUrl string `yaml:"webhookUrl"`
	userName   string `yaml:"userName"`
	channel    string `yaml:"channel"`
}

var Conf Config

func LoadConfSlack(confPath string) {
	yamlFile, err := os.ReadFile(confPath)
	if err != nil {
		log.Fatal().Err(err).Msg("Can't read notification config file")
		os.Exit(1)
	}
	err = yaml.Unmarshal(yamlFile, &Conf)
	if err != nil {
		log.Fatal().Err(err).Msg("Unmarshal error")
		os.Exit(1)
	}
}
