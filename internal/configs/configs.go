package configs

import (
	"log"
	"time"

	"github.com/spf13/viper"
)

type JarchiConfig struct {
	Notifier    SectionNotifier    `mapstructure:"notifier"`
	DivarClient SectionDivarClient `mapstructure:"divar"`
	Core        SectionCore        `mapstructure:"core"`
}

type SectionNotifier struct {
	BaseUrl string `mapstructure:"baseurl"`
	Target  int    `mapstructure:"target"`
	Token   string `mapstructure:"token"`
}

type SectionDivarClient struct {
	URL      string        `mapstructure:"url"`
	MaxPrice int           `mapstructure:"max_price"`
	MinSize  int           `mapstructure:"min_size"`
	MaxDate  time.Duration `mapstructure:"max_date"`
}

type SectionCore struct {
	Interval time.Duration `mapstructure:"interval"`
}

func GetConfig(paths ...string) (JarchiConfig, error) {
	var c JarchiConfig
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
