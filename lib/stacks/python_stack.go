package stacks

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/aws-cdk-go/awscdklambdapythonalpha/v2"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

type CdktestGoStackProps struct {
	awscdk.StackProps
}

func NewCdkPythonTestAppGoStack(scope constructs.Construct, id string, props *CdktestGoStackProps) awscdk.Stack {
	var sprops awscdk.StackProps
	if props != nil {
		sprops = props.StackProps
	}
	stack := awscdk.NewStack(scope, &id, &sprops)

	awscdklambdapythonalpha.NewPythonFunction(stack, jsii.String("MyFunction"), &awscdklambdapythonalpha.PythonFunctionProps{
		Entry:   jsii.String("./lib/lambda/python"),
		Runtime: awslambda.Runtime_PYTHON_3_8(),
		Index:   jsii.String("testlambda.py"),
		Handler: jsii.String("lambda_handler"),
	})
	return stack
}
