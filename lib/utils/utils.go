package utils

import (
	"fmt"
	"github.com/aws/jsii-runtime-go"
)

type CustomCommandHooks struct{}

// Implement BeforeBundling method
func (c *CustomCommandHooks) BeforeBundling(inputDir *string, outputDir *string) *[]*string {
	return &[]*string{
		jsii.String(fmt.Sprintf("echo 'Before bundling: InputDir=%s, OutputDir=%s'", *inputDir, *outputDir)),
	}
}

// Implement BeforeInstall method
func (c *CustomCommandHooks) BeforeInstall(inputDir *string, outputDir *string) *[]*string {
	return &[]*string{
		jsii.String("echo 'Before install commands...'"),
	}
}

// Implement AfterBundling method
func (c *CustomCommandHooks) AfterBundling(inputDir *string, outputDir *string) *[]*string {
	return &[]*string{
		jsii.String("echo 'after install commands...'"),
	}
}
