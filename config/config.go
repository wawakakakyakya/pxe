package config

type Config struct {
	Dhcp struct {
		Listen   string `yaml:"listen"`
		Enabled  bool   `yaml:"enabled"`
		Port     int    `yaml:"port"`
		Range    string `yaml:"range"`
		LeaseSec int    `yaml:"lease_sec"`
		Hosts    []struct {
			Name string `yaml:"name"`
			Addr string `yaml:"addr"`
			Mac  string `yaml:"mac"`
		} `yaml:"hosts"`
		Options struct {
			Subnet   string `yaml:"subnet"`
			Gateway  string `yaml:"gateway"`
			Dns      string `yaml:"dns"`
			TftpSrv  string `yaml:"tftp_srv"`
			BootFile string `yaml:"boot_file"`
		} `yaml:"options"`
	} `yaml:"dhcp"`
	Http struct {
		Enabled   bool   `yaml:"enabled"`
		Port      int    `yaml:"port"`
		UseSsl    bool   `yaml:"use_ssl"`
		ImagePath string `yaml:"image_path"`
	} `yaml:"http"`
	Tftp struct {
		Enabled   bool   `yaml:"enabled"`
		Port      port   `yaml:"port"`
		ImagePath string `yaml:"root_dir"`
	} `yaml:"tftp"`
}

func init() {
	setArgs()
}

func NewConfig() (*Config, error) {
	yamlC, err := LoadYamlConfig(configPath)
	if err != nil {
		return nil, err
	}
	return yamlC, err
}
