package service

type ServiceGroup struct {
	JwtService jwtService
}

var ServiceGroupApp = new(ServiceGroup)
