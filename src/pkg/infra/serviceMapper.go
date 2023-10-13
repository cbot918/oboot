package infra

type serviceMapper map[string]Service

func NewServiceMapper() serviceMapper {

	sm := make(serviceMapper)

	sm["redis"] = Service{
		Name:          "redis",
		Image:         "redis:latest",
		ContainerName: "redis",
		Restart:       "always",
		Ports:         []string{"6377:6379"},
	}

	sm["rabbitmq"] = Service{
		Name:          "rabbitmq",
		Image:         "rabbitmq:3-management",
		ContainerName: "rabbitmq",
		Restart:       "always",
		Environment:   []string{"RABBITMQ_DEFAULT_USER=yale918", "RABBITMQ_DEFAULT_PASS=12345"},
		Ports:         []string{"5672:5672", "5673:15672"},
	}

	sm["nginx"] = Service{
		Name:          "nginx",
		Image:         "nginx:latest",
		ContainerName: "nginx",
		Restart:       "always",
		Ports:         []string{"80:80"},
	}

	return sm

}
