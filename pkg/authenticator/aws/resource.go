package aws

import (
	"fmt"
	"strings"

	"github.com/AndrewCopeland/conjur-authn-iam-client/pkg/utils/aws"
	"github.com/AndrewCopeland/conjur-authn-iam-client/pkg/utils/aws/ec2"
	"github.com/AndrewCopeland/conjur-authn-iam-client/pkg/utils/aws/ecs"
	"github.com/AndrewCopeland/conjur-authn-iam-client/pkg/utils/aws/lambda"
)

func getAwsResources() []aws.AwsResource {
	resources := []aws.AwsResource{}
	resources = append(resources, ec2.New())
	resources = append(resources, lambda.New())
	resources = append(resources, ecs.New())
	return resources
}

// GetAwsResource will return an interface that has the ability to retrieve IAM AWS credentials from the desired metadata endpoint
func GetAwsResource(name string) (aws.AwsResource, error) {
	resources := getAwsResources()
	for _, r := range resources {
		if strings.ToLower(name) == r.Name() {
			return r, nil
		}
	}

	return nil, fmt.Errorf("Failed to retrieve AWS resource with type '%s'", name)
}
