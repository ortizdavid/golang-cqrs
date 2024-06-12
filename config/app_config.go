package config

func ListenAddr() string {
	return GetEnv("APP_HOST")+":"+GetEnv("APP_PORT")
}

