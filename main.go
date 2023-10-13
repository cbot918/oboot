package main

import (
	"log"

	"github.com/cbot918/oboot/src/pkg/infra"
)

const (
	file = "infra.o"
)

func main() {
	var err error
	infra, err := infra.NewInfra("test", file)
	if err != nil {
		log.Fatal(err)
	}
	err = infra.DockerCompose.WriteDockercompose("docker-compose.yaml")
	if err != nil {
		log.Fatal(err)
	}
}
