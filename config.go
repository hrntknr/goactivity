package main

type Config struct {
	Server struct {
		Host string `yaml:"host"`
	} `yaml:"server"`
}
