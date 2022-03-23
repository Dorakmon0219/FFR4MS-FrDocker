package types

type ServiceActuatorInfo struct {
	Leaf int
	Port string
}

type ServiceActuatorHealth struct {
	Status string
}

type GatewayActuatorInfo struct {
	Getway string
	Port   string
}

type EurekaConfig struct {
	ArrayIpPort []string
	ArrayGetWay []string
	ArrayGroup  []string
}

type GateWayReplaceService struct {
	ServiceName      string
	DownInstanceHost string
	DownInstancePort string
}

type AddContainerDTO struct {
	ServiceGroup string `json:"serviceGroup"`
	ServiceIP    string `json:"serviceIP"`
	ServicePort  string `json:"servicePort"`
}
