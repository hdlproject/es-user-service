package helper

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/endpoints"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/kms"
)

type KMSClient struct {
	svc *kms.KMS
}

func NewKMSClient(id, secret string) (*KMSClient, error) {
	s, err := session.NewSession(
		&aws.Config{
			Region:      aws.String(endpoints.ApSoutheast2RegionID),
			Credentials: credentials.NewStaticCredentials(id, secret, ""),
		},
	)
	if err != nil {
		return nil, WrapError(err)
	}

	return &KMSClient{
		svc: kms.New(s),
	}, nil
}

func (instance *KMSClient) GenerateMac(message string) (string, error) {
	input := &kms.GenerateMacInput{
		KeyId:        aws.String("c4fed70d-6cc4-4a8a-a0d5-84eb865e1490"),
		MacAlgorithm: aws.String("HMAC_SHA_512"),
		Message:      []byte(message),
	}

	result, err := instance.svc.GenerateMac(input)
	if err != nil {
		return "", WrapError(err)
	}

	return string(result.Mac), nil
}

func (instance *KMSClient) VerifyMac(signature, message string) error {
	input := &kms.VerifyMacInput{
		KeyId:        aws.String("c4fed70d-6cc4-4a8a-a0d5-84eb865e1490"),
		MacAlgorithm: aws.String("HMAC_SHA_512"),
		Mac:          []byte(signature),
		Message:      []byte(message),
	}

	_, err := instance.svc.VerifyMac(input)
	if err != nil {
		return WrapError(err)
	}

	return nil
}

func (instance *KMSClient) Sign(message string) (string, error) {
	input := &kms.SignInput{
		KeyId:            aws.String("c720978c-a2ab-42c8-baa0-478b5462b49e"),
		SigningAlgorithm: aws.String("RSASSA_PKCS1_V1_5_SHA_512"),
		Message:          []byte(message),
		MessageType:      aws.String("RAW"),
	}

	result, err := instance.svc.Sign(input)
	if err != nil {
		return "", WrapError(err)
	}

	return string(result.Signature), nil
}

func (instance *KMSClient) Verify(signature, message string) error {
	input := &kms.VerifyInput{
		KeyId:            aws.String("c720978c-a2ab-42c8-baa0-478b5462b49e"),
		SigningAlgorithm: aws.String("RSASSA_PKCS1_V1_5_SHA_512"),
		Signature:        []byte(signature),
		Message:          []byte(message),
		MessageType:      aws.String("RAW"),
	}

	_, err := instance.svc.Verify(input)
	if err != nil {
		return WrapError(err)
	}

	return nil
}
