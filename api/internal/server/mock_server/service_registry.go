package mock_server

import (
	"fmt"
	"reflect"
)

type ServiceRegistry struct {
	services     map[reflect.Type]interface{}
	serviceTypes []reflect.Type
}

func NewServiceRegistry() *ServiceRegistry {
	return &ServiceRegistry{
		services: make(map[reflect.Type]interface{}),
	}
}

func (s *ServiceRegistry) Register(service interface{}) error {
	serviceType := reflect.TypeOf(service)
	if serviceType.Kind() != reflect.Ptr {
		return fmt.Errorf("service must be a pointer, received non-pointer type: %T", service)
	}

	if _, exists := s.services[serviceType]; exists {
		return fmt.Errorf("service already exists: %v", serviceType)
	}
	s.services[serviceType] = service
	s.serviceTypes = append(s.serviceTypes, serviceType)
	return nil
}

func (s *ServiceRegistry) Fetch(service interface{}) error {
	if reflect.TypeOf(service).Kind() != reflect.Ptr || reflect.ValueOf(service).Elem().Kind() != reflect.Ptr {
		return fmt.Errorf("input must be a pointer to a pointer, received: %T", service)
	}

	serviceType := reflect.ValueOf(service).Elem().Type()
	if runningService, ok := s.services[serviceType]; ok {
		reflect.ValueOf(service).Elem().Set(reflect.ValueOf(runningService))
		return nil
	}
	return fmt.Errorf("unknown service: %T", service)
}
