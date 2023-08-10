package config

type Server struct {
	Zap    Zap    `mapstructure:"zap" json:"zap" yaml:"zap"`
	System System `mapstructure:"system" json:"system" yaml:"system"`
	// gorm
	Mysql      Mysql      `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
	Timer      Timer      `mapstructure:"timer" json:"timer" yaml:"timer"`
	Kubernetes Kubernetes `mapstructure:"kubernetes" json:"kubernetes" yaml:"kubernetes"`
}
