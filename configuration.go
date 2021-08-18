package main

type EntrezConfig struct {
	Email  string
	APIKey string
}
type Config struct {
	Host   string
	Entrez EntrezConfig
}
