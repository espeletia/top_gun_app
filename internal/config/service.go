package config

type ServiceConfig struct {
	ID          string
	Environment string
	Name        string
}

func loadServiceConfig() ServiceConfig {
	serviceConfig := &ServiceConfig{}
	v := configViper("service")
	v.BindEnv("Envinroment")
	v.BindEnv("ID")
	err := v.ReadInConfig()
	if err != nil {
		panic(err)
	}
	v.Unmarshal(serviceConfig)
	return *serviceConfig
}
