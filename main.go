package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ssm"
)

var parameters = map[string]string{
	"/zservices/menu-service/dev/MENU_SERVICE_SERVER_PORT": "80",
	"/zservices/menu-service/dev/MENU_SERVICE_LOG_LEVEL":   "debug",
	"/zservices/menu-service/dev/MENU_SERVICE_LOG_FORMAT":  "text",
}

func main() {
	// aws login
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("ap-south-1"),
	})
	if err != nil {
		panic(err)
	}

	ssmsvc := ssm.New(sess, aws.NewConfig().WithRegion("ap-south-1"))
	for key, value := range parameters {
		_, err := ssmsvc.PutParameter(&ssm.PutParameterInput{
			Type:  aws.String("String"),
			Name:  aws.String(key),
			Value: aws.String(value),
		})
		if err != nil {
			fmt.Printf("Failed to add key %v \n", key)
			fmt.Println(err)
		}
	}
}
