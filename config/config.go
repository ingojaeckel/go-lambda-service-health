package config

import (
	"io/ioutil"
	"log"

	"fmt"
	yaml "gopkg.in/yaml.v2"
)

type Configuration struct {
	// Number of seconds to wait for a response from a service
	Timeout int `yaml:"timeout"`
	// AWS region of the s3 bucket that is being written to and read from.
	Region string `yaml:"region"`
	// The bucket where all data will be stored (measurements and reports)
	S3Bucket string `yaml:"s3bucket"`
	// Folder & file name within the s3 bucket where measurements are stored.
	S3KeyData string `yaml:"s3keyData"`
	// Folder & file name within the s3 bucket where the report will be stored.
	S3KeyReport string `yaml:"s3KeyReport"`

	Services []ServiceConfiguration `yaml:"services"`
}

type ServiceConfiguration struct {
	Name string `yaml:"name"`
	URL  string `yaml:"url"`
}

func (sc ServiceConfiguration) String() string {
	return fmt.Sprintf("{name:%s, url:%s}", sc.Name, sc.URL)
}

func LoadConfiguration(path string) (*Configuration, error) {
	log.Printf("Loading configuration from %s..\n", path)
	c, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var conf Configuration
	if err := yaml.Unmarshal(c, &conf); err != nil {
		log.Printf("Failed to unmarshal config: %s\n", err.Error())
		return nil, err
	}
	log.Printf("Finished loading configuration %v\n", conf)
	return &conf, err
}
