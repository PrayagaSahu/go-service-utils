package consul

import (
	"fmt"

	"github.com/hashicorp/consul/api"
	"go.uber.org/zap"
)

// ServiceConfig is the input for registering a service in Consul
type ServiceConfig struct {
	ID      string
	Name    string
	Address string
	Port    int
	Tags    []string
}

// RegisterService registers a service with Consul
func RegisterService(cfg ServiceConfig) {
	consulCfg := api.DefaultConfig()
	consulCfg.Address = "127.0.0.1:8500"

	client, err := api.NewClient(consulCfg)
	if err != nil {
		zap.L().Error("❌ Consul client initialization failed", zap.Error(err))
		return
	}

	check := &api.AgentServiceCheck{
		HTTP:                           fmt.Sprintf("http://%s:%d/health", cfg.Address, cfg.Port),
		Interval:                       "10s",
		Timeout:                        "5s",
		DeregisterCriticalServiceAfter: "1m",
	}

	reg := &api.AgentServiceRegistration{
		ID:      cfg.ID,
		Name:    cfg.Name,
		Address: cfg.Address,
		Port:    cfg.Port,
		Tags:    cfg.Tags,
		Check:   check,
	}

	err = client.Agent().ServiceRegister(reg)
	if err != nil {
		zap.L().Error("❌ Failed to register service with Consul", zap.Error(err))
		return
	}

	zap.L().Info("✅ Service registered with Consul",
		zap.String("serviceID", cfg.ID),
		zap.String("serviceName", cfg.Name),
		zap.String("serviceAddress", cfg.Address),
		zap.Int("servicePort", cfg.Port),
		zap.Strings("tags", cfg.Tags),
	)
}
