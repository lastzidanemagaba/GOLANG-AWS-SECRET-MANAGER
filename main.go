package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/secretsmanager"
)

func main() {
	// Replace these values with your AWS credentials and secret name
	accessKeyID := ""
	secretAccessKey := ""
	region := ""
	secretName := ""

	// create AWS session
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String(region),
		Credentials: credentials.NewStaticCredentials(accessKeyID, secretAccessKey, ""),
	})
	if err != nil {
		fmt.Println(err)
		return
	}

	// create SecretsManager client
	svc := secretsmanager.New(sess)

	// retrieve secret value
	input := &secretsmanager.GetSecretValueInput{
		SecretId: aws.String(secretName),
	}
	result, err := svc.GetSecretValue(input)
	if err != nil {
		fmt.Println(err)
		return
	}
	secretValue := aws.StringValue(result.SecretString)

	// use secret value in your application
	fmt.Println(secretValue)
}
