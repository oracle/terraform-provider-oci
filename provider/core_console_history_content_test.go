// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

const (
	ConsoleHistoryContentResourceConfig = ConsoleHistoryContentResourceDependencies + `

`
	ConsoleHistoryContentPropertyVariables = `
variable "console_history_content_length" { default = 10 }
variable "console_history_content_offset" { default = 10 }

`
	ConsoleHistoryContentResourceDependencies = ConsoleHistoryPropertyVariables + ConsoleHistoryResourceConfig
)

func TestCoreConsoleHistoryContentResource_basic(t *testing.T) {
	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getRequiredEnvSetting("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	singularDatasourceName := "data.oci_core_console_history_data.test_console_history_content"

	resource.Test(t, resource.TestCase{
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		Steps: []resource.TestStep{
			// verify singular datasource
			{
				Config: config + `
variable "console_history_content_length" { default = 10240 }
variable "console_history_content_offset" { default = 0 }

data "oci_core_console_history_data" "test_console_history_content" {
	#Required
	console_history_id = "${oci_core_console_history.test_console_history.id}"

	#Optional
	length = "${var.console_history_content_length}"
	offset = "${var.console_history_content_offset}"
}
                ` + compartmentIdVariableStr + ConsoleHistoryContentResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(singularDatasourceName, "console_history_id"),
					resource.TestCheckResourceAttr(singularDatasourceName, "length", "10240"),
					resource.TestCheckResourceAttr(singularDatasourceName, "offset", "0"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "data"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				),
			},
		},
	})
}
