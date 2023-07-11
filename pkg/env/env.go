package env

func LoadEnv() (*EnvPlain, error) {
	envPlain, err := loadEnvPlain()
	if err != nil {
		return nil, err
	}

	if err := checkEnvValues(*envPlain); err != nil {
		return nil, err
	} else {
		return envPlain, nil
	}
}
