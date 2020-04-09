package main

import (
	"fmt"
	"github.com/oklog/ulid/v2"
	"gopkg.in/yaml.v2"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"time"
)

var cfg Config

type Config struct {
	Server struct {
		Port string `yaml:"port"`
		Host string `yaml:"host"`
	} `yaml:"server"`
}

func main() {

	configFile := "./config/config.yml"
	if len(os.Args) > 1 {
		configFile = os.Args[1]
	}
	log.Println(configFile)

	readConfig(&cfg, configFile)

	http.HandleFunc("/", GenerateHandler)
	http.HandleFunc("/timestamp", TimestampHandler)
	http.HandleFunc("/datetime", DatetimeHandler)

	serve := fmt.Sprintf("%s:%s", cfg.Server.Host, cfg.Server.Port)
	if err := http.ListenAndServe(serve, nil); err != nil {
		log.Fatal(err)
	}
}

func readConfig(cfg *Config, configFile string) {
	f, err := os.Open(configFile)
	if err != nil {
		processError(err)
	}
	defer f.Close()

	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(cfg)
	if err != nil {
		processError(err)
	}
}

func GenerateHandler(w http.ResponseWriter, r *http.Request) {

	t := time.Now().UTC()
	entropy := rand.New(rand.NewSource(t.UnixNano()))
	id := ulid.MustNew(ulid.Timestamp(t), entropy)
	fmt.Println(r.RemoteAddr, id.String())
	w.Write([]byte (id.String()))

}

func TimestampHandler(w http.ResponseWriter, r *http.Request) {

	source := r.URL.Query().Get("id")
	id := ulid.MustParse(source)
	w.Write([]byte(strconv.FormatUint(id.Time(), 10)))

}

func DatetimeHandler(w http.ResponseWriter, r *http.Request) {

	source := r.URL.Query().Get("id")
	timestamp := ulid.MustParse(source).Time()
	datetime := time.Unix(int64(timestamp)/1000,0)
	w.Write([]byte(datetime.String()))
}

func processError(err error) {
	fmt.Println(err)
	os.Exit(2)
}
