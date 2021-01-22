package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"os/user"

	"gopkg.in/yaml.v2"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func main() {

	var cfg = getConf()
	printStruct(cfg)

	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(cfg.Bucket.Region)},
	)

	if err != nil {
		exitErrorf(err.Error())
	}

	svc := s3.New(sess)

	client := &S3Client{
		svc:    svc,
		bucket: cfg.Bucket.Name,
	}

	controller := &Controller{
		s3Client: client,
	}

	println("starting server. Listenning on port 8090")
	//http.Handle("/", http.FileServer(http.Dir("./public")))

	http.HandleFunc("/", controller.handlePublic)

	http.HandleFunc("/api/objects/", controller.handleObjectsEndpoint)

	http.ListenAndServe(":8090", nil)
}

func getConf() *Config {

	usr, _ := user.Current()
	dir := usr.HomeDir
	configLocation := fmt.Sprintf("%s/.config/s3album/s3album.yaml", dir)

	fmt.Printf("Lookin for config in %s\n", configLocation)
	var c Config
	f, err := os.Open(configLocation)
	if err != nil {
		exitErrorf(err.Error())
	}
	defer f.Close()

	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&c)
	if err != nil {
		exitErrorf(err.Error())
	}

	return &c
}

func exitErrorf(msg string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, msg+"\n", args...)
	os.Exit(1)
}

type Config struct {
	Bucket struct {
		Name   string `yaml:"name"`
		Region string `yaml:"region"`
	} `yaml:"bucket"`
}

func printStruct(c *Config) {
	out, err := json.Marshal(c)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(out))
}
