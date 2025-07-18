package model

type ServiceConfig struct {
	ID      string
	Name    string
	Address string
	Port    int
	Tags    []string
}
