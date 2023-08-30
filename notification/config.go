package notification

import (
	"io/ioutil"
	"os"

	"github.com/rs/zerolog/log"
	"gopkg.in/yaml.v3"
)

type ConfigSlackYaml struct {
	ClusterName string `yaml:"clusterName"`
	SlackClient struct {
		WebhookUrl string `yaml:"webhookUrl"`
		UserName   string `yaml:"userName"`
		Channel    string `yaml:"channel"`
	} `yaml:"slackClient"`
}

var ConfigurationSlack ConfigSlackYaml

func LoadConfSlack(confPath string) {
	yamlFile, err := ioutil.ReadFile(confPath)
	if err != nil {
		log.Fatal().Err(err).Msg("Can't read notification config file")
		os.Exit(1)
	}
	err = yaml.Unmarshal(yamlFile, &ConfigurationSlack)
	if err != nil {
		log.Fatal().Err(err).Msg("Unmarshal error")
		os.Exit(1)
	}
}
