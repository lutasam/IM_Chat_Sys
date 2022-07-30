package utils

import (
	"github.com/asaskevich/govalidator"
)

func IsValidIP(ip string) bool {
	return govalidator.IsIP(ip)
}

func IsValidURL(url string) bool {
	return govalidator.IsURL(url)
}

func IsValidPort(port int) bool {
	return port >= 1024 && port < 65536
}
