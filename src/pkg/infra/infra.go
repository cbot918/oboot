package infra

import (
	"os"
	"regexp"
	"strings"
)

const (
	configFileName = "infra.o"
)

type Infra struct {
	InfraConfig   *InfraConfig
	DockerCompose *DockerCompose
}

type InfraConfig struct {
	Name     string
	Network  string
	Services []string
}

func NewInfra(infraName string, config string) (*Infra, error) {

	content, err := os.ReadFile(config)
	if err != nil {
		return nil, err
	}
	network := regexp.MustCompile(`network: (.*)`).FindStringSubmatch(string(content))[1]

	service := regexp.MustCompile(`service: (.*)`).FindStringSubmatch(string(content))[1]

	service = strings.Replace(service, " ", "", -1)
	service = strings.Replace(service, "[", "", -1)
	service = strings.Replace(service, "]", "", -1)
	serviceArray := strings.Split(service, ",")

	ic := &InfraConfig{
		Name:     infraName,
		Network:  network,
		Services: serviceArray,
	}

	return &Infra{
		InfraConfig: ic,
		DockerCompose: NewDockerCompose(
			ic,
			string(content)),
	}, err
}
