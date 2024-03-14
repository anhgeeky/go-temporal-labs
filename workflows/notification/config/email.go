package config

type EmailConfig struct {
	SmtpHost     string `mapstructure:"SMTP_HOST"`
	SmtpPort     int    `mapstructure:"SMTP_PORT"`
	SmtpAccount  string `mapstructure:"SMTP_ACCOUNT"`
	SmtpPassword string `mapstructure:"SMTP_PASSWORD"`
	EmailFrom    string `mapstructure:"EMAIL_FROM"`
	EmailTo      string `mapstructure:"EMAIL_TO"`
}
