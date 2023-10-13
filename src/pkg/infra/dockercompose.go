package infra

import (
	"fmt"
	"os"
)

type DockerCompose struct {
	infraConfig   *InfraConfig
	configContent string
}

type DockerComposeConfig struct {
	FileName string
	Version  string
	Network  string
	Services []Service
}

type Service struct {
	Name          string
	Image         string
	ContainerName string
	Restart       string
	Environment   []string
	Ports         []string
}

func NewDockerCompose(ic *InfraConfig, configContent string) *DockerCompose {

	return &DockerCompose{
		infraConfig:   ic,
		configContent: configContent,
	}
}

func (dc *DockerCompose) GetContent(ss []Service) (string, error) {

	result := ""

	dcc := &DockerComposeConfig{
		Version:  "'3'",
		Network:  dc.infraConfig.Network,
		Services: ss,
	}

	result += dc.getString(0, "version:", dcc.Version)
	result += "\n"
	result += dc.getString(0, "services:", "")
	for _, item := range dcc.Services {
		result += "\n"
		result += dc.getServiceString(&item)
	}
	return result, nil
}

func (dc *DockerCompose) getMockServices(canditates []string) []Service {

	serviceMapper := NewServiceMapper()

	ss := []Service{}
	for _, item := range canditates {
		ss = append(ss, serviceMapper[item])
	}

	return ss
}

func (dc *DockerCompose) getServiceString(s *Service) (result string) {

	result += dc.getString(2, s.Name+":", "")
	result += dc.getString(4, "image:", s.Image)
	result += dc.getString(4, "container_name:", dc.infraConfig.Name+"_"+s.ContainerName)
	result += dc.getString(4, "restart:", s.Restart)

	if s.Ports != nil {
		result += dc.getString(4, "ports:", "")
		for _, port := range s.Ports {
			result += dc.getString(6, "-", port)
		}
	}

	if s.Environment != nil {
		result += dc.getString(4, "environment:", "")
		for _, env := range s.Environment {
			result += dc.getString(6, "-", env)
		}
	}

	return result
}

func (dc *DockerCompose) getString(space int8, key string, val string) (result string) {

	switch space {
	case 0:
		result += ""
	case 2:
		result += "  "
	case 4:
		result += "    "
	case 6:
		result += "      "
	}

	result += key + " " + val + "\n"

	return
}

func (d *DockerCompose) WriteDockercompose(outoutName string) error {

	services := d.getMockServices(d.infraConfig.Services)
	content, err := d.GetContent(services)
	if err != nil {
		return err
	}

	fmt.Println(content)

	fd, err := os.Create(outoutName)
	if err != nil {
		return err
	}

	_, err = fd.Write([]byte(content))
	if err != nil {
		return err
	}

	return nil
}
