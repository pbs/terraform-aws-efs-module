package test

import (
	"fmt"
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

func testEFS(t *testing.T, variant string) {
	t.Parallel()

	terraformDir := fmt.Sprintf("../examples/%s", variant)

	terraformOptions := &terraform.Options{
		TerraformDir: terraformDir,
		LockTimeout:  "5m",
	}

	defer terraform.Destroy(t, terraformOptions)

	terraform.InitAndApply(t, terraformOptions)

	arn := terraform.Output(t, terraformOptions, "arn")
	assert.Contains(t, arn, "arn:")

	id := terraform.Output(t, terraformOptions, "id")
	assert.Contains(t, id, "fs-")

	dns_name := terraform.Output(t, terraformOptions, "dns_name")
	assert.Contains(t, dns_name, ".amazonaws.com")
}
