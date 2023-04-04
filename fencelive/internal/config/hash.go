package config

type HashConfig struct {
	Salt string
}

func loadHashConfig() HashConfig {
	HashConfig := &HashConfig{}
	v := configViper("hash")
	v.BindEnv("HASH_SALT")
	err := v.ReadInConfig()
	if err != nil {
		panic(err)
	}
	v.Unmarshal(HashConfig)
	return *HashConfig
}
