package utils

import (
	"context"
	"encoding/base64"
	"github.com/docker/docker/client"
	"github.com/golang-jwt/jwt/v4"
)

var tokenKey string = "VTON2023COMPETITION-govton"

type VitonClaim struct {
	jwt.StandardClaims
}

func CheckContainerRunning(cli *client.Client, id string) bool {
	inspect, err := cli.ContainerInspect(context.Background(), id)
	if err != nil {
		return false
	}
	return inspect.State.Running
}

func ConvertImage(s string) []byte {
	data, _ := base64.StdEncoding.DecodeString(s)
	return data
}

func GenerateToken() (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, VitonClaim{})
	tokenString, err := token.SignedString([]byte(tokenKey))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
