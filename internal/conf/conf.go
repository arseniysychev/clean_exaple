package conf

import (
	"fmt"
	"github.com/caarlos0/env/v6"
	"os"
	"sync"
)

var once sync.Once
var config Config

type Config struct {
	DB   DB
	LOG  LOG
	GRPC GRPC
}

func New() Config {
	once.Do(func() {
		println("Config")
		if err := env.Parse(&config); err != nil {
			fmt.Println("parsing configuration:", err)
			os.Exit(-1)
		}
		//config = Config{
		//	DB: DB{
		//		Name:     "test",
		//		Host:     "0.0.0.0",
		//		Port:     26257,
		//		User:     "root",
		//		Password: "",
		//	},
		//	GRPC: GRPC{
		//		Host: "0.0.0.0",
		//		Port: "50051",
		//	},
		//}
	})
	return config
}

// DB describes database config
type DB struct {
	Name     string `env:"DB_NAME,required"`
	Host     string `env:"DB_HOST,required"`
	Port     int    `env:"DB_PORT,required"`
	User     string `env:"DB_USER,required"`
	Password string `env:"DB_PASSWORD,required"`
}

func (d *DB) DSN() string {
	fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s timezone=%s", d.Host, d.Port, d.User, d.Password, d.Name, d.SSL, d.TimeZone)
	return "postgresql://root@localhost:26257?sslmode=disable"
}

type GRPC struct {
	Host string `env:"GRPC_HOST" envDefault:"0.0.0.0"`
	Port string `env:"GRPC_PORT" envDefault:"10000"`
}

func (grpc *GRPC) Address() string {
	return grpc.Host + ":" + grpc.Port
}

type LOG struct {
	Level      string `env:"LOG_LEVEL" envDefault:"debug"`
	Json       bool   `env:"LOG_JSON" envDefault:"false"`
	TimeFormat string `env:"LOG_TIME_FORMAT" envDefault:"2006-01-02T15:04:05Z"`
}
