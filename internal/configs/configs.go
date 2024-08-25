package configs

import (
	"log"
	"time"

	"github.com/spf13/viper"
)

type ServiceConfig struct {
	Notifier    SectionNotifier    `yaml:"notifier"`
	DivarClient SectionDivarClient `yaml:"divar"`
	Core        SectionCore        `yaml:"core"`
}

type SectionNotifier struct {
	BaseUrl string `yaml:"baseurl"`
	Target  int    `yaml:"target"`
	Token   string `yaml:"token"`
}

type SectionDivarClient struct {
	URL string `yaml:"url"`
}

type SectionCore struct {
	Interval time.Duration `yaml:"interval"`
}

func GetConfig(paths ...string) (ServiceConfig, error) {
	var c ServiceConfig
	viper.SetConfigType("yaml")
	viper.SetConfigName("divar-notifier")
	for _, p := range paths {
		viper.AddConfigPath(p)
	}
	err := viper.ReadInConfig()
	if err != nil {
		return c, err
	}
	err = viper.Unmarshal(&c)
	log.Printf("pathes are %v and config is %v\n", paths, c)
	return c, err
}
