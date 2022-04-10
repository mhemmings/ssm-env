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
	var internalErr error
	err = svc.GetParametersByPathPages(&input, func(out *ssm.GetParametersByPathOutput, lastPage bool) bool {
		for _, param := range out.Parameters {
			name := strings.TrimPrefix(aws.StringValue(param.Name), path)
			internalErr = os.Setenv(name, aws.StringValue(param.Value))
			if internalErr != nil {
				return false
			}
		}
		return true
	})
	if err != nil {
		return err
	}
	if internalErr != nil {
		return internalErr
	}

	return nil
}
