package configs

type Config struct {
	DBUrl     string
	JWTSecret string
	Port      string
}

func Load() Config {
	return Config{
		DBUrl:     "postgres://erp:erp123@localhost:5432/erp_dev?sslmode=disable",
		JWTSecret: "troca-isso-antes-do-deploy",
		Port:      "8080",
	}
}
