package config

type Config struct {
	Http struct {
		Port int `mapstructure:"port"`
	} `mapstructure:"http"`

	Postgres struct {
		Host     string `mapstructure:"host"`
		Port     int    `mapstructure:"port"`
		Database string `mapstructure:"database"`
		User     string `mapstructure:"user"`
		Passwrod string `mapstructure:"password"`
	} `mapstructure:"postgres"`

	Database struct {
		Migration bool `mapstructure:"host"`
	} `mapstructure:"database"`

	Kafka struct {
		Brokers string `mapstructure:"brokers"`
	} `mapstructure:"kafka"`

	Aliyun struct {
		AccessKeyId      string `mapstructure:"accessKeyId"`
		AccessKeySecret  string `mapstructure:"accessKeySecret"`
		CalledShowNumber string `mapstructure:"calledShowNumber"`
	} `mapstructure:"aliyun"`

	Cmccopen struct {
		ApiKey     string `mapstructure:"apiKey"`
		SecretKey  string `mapstructure:"secretKey"`
		BaseUrl    string `mapstructure:"baseUrl"`
		DisplayNbr string `mapstructure:"displayNbr"`
		RespUrl    string `mapstructure:"respUrl"`
	} `mapstructure:"cmccopen"`
}
