package vars 

type ConfigSchema struct {
	JWT_SECRET string `required:"true"`

	DB_USER string `required:"true"`
	DB_PASSWORD string `required:"true"`
	DB_DATABASE string `required:"true"`
	DB_HOST string `required:"true"`
}

var Config = ConfigSchema{}