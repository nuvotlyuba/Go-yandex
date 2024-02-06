package config

import (
	"errors"
	"strconv"
	"strings"
)

const (
	Port = 8080
	Host = "localhost"

)

type NetAddress struct {
	Host string
	Port int
}

type BaseUrl string

func (a NetAddress) String() string {
    return a.Host + ":" + strconv.Itoa(a.Port)
}

func (a *NetAddress) Set(s string) error {
    hp := strings.Split(s, ":")
    if len(hp) != 2 {
        return errors.New("Need address in a form host:port")
    }
    port, err := strconv.Atoi(hp[1])
    if err != nil{
        return err
    }
    a.Host = hp[0]
    a.Port = port
    return nil
}

func GetDefaultBaseUrl() string {
	return Host + ":" + strconv.Itoa(Port)
}
