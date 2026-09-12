package config

func defaultConfig() *Config {
	return &Config{
		App: AppConfig{
			Name:    "AnagataSentinel",
			Version: "0.1.0",
			Debug:   false,
		},
		Window: WindowConfig{
			Width:     1280,
			Height:    800,
			MinWidth:  1024,
			MinHeight: 640,
		},
		Database: DatabaseConfig{
			Path:     "data/sentinel.db",
			Password: "0123456789",
		},
		Logging: LoggingConfig{
			Level:      "info",
			File:       "logs/app.log",
			MaxSizeMB:  50,
			MaxBackups: 3,
			Compress:   true,
		},
		Splash: SplashConfig{
			Enabled:  true,
			Duration: 10000,
		},
		Report: ReportConfig{
			Enabled: true,
			DSN:     "https://241ee37658fca3cae3255e5435902c54@o4512071942406144.ingest.us.sentry.io/4512071947190272",
		},
	}
}
