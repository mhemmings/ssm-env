package ssm

import (
	"os"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ssm"
)

// Parse sets the environment using AWS SSM. All parameters are fetched from
// SSM using the provided path, and used to set the current environment.
// It is assumed that all parameters are of type "SecureString".
func Parse(path string) error {
	if path == "" {
		return nil
	}
	if !strings.HasSuffix(path, "/") {
		path += "/"
	}
	sess, err := session.NewSession()
	if err != nil {
		return err
	}
	svc := ssm.New(sess, aws.NewConfig())
	input := ssm.GetParametersByPathInput{
		Path:           aws.String(path),
		WithDecryption: aws.Bool(true),
	}
	out, err := svc.GetParametersByPath(&input)
	if err != nil {
		return err
	}
	for _, param := range out.Parameters {
		name := strings.TrimPrefix(aws.StringValue(param.Name), path)
		err := os.Setenv(name, aws.StringValue(param.Value))
		if err != nil {
			return err
		}
	}
	return nil
}
