package stacks

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/aws-cdk-go/awscdklambdagoalpha/v2"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	Config "market-data/lib/lambda/golang/config"
	"market-data/lib/utils"
)

type CdkGotestGoStackProps struct {
	awscdk.StackProps
}

func NewCdkGoTestAppGoStack(scope constructs.Construct, id string, props *CdkGotestGoStackProps) awscdk.Stack {
	var sprops awscdk.StackProps
	if props != nil {
		sprops = props.StackProps
	}
	stack := awscdk.NewStack(scope, &id, &sprops)

	Config.Init()

	awscdklambdagoalpha.NewGoFunction(stack, jsii.String("MyFunction"), &awscdklambdagoalpha.GoFunctionProps{
		Entry:   jsii.String("./lib/lambda/golang"),
		Runtime: awslambda.Runtime_PROVIDED_AL2(),
		Bundling: &awscdklambdagoalpha.BundlingOptions{
			CommandHooks: &utils.CustomCommandHooks{},
		},
	})

	return stack
}
