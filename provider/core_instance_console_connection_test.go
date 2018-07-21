// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.com/oracle/oci-go-sdk/common"
	oci_core "github.com/oracle/oci-go-sdk/core"
)

const (
	InstanceConsoleConnectionRequiredOnlyResource = InstanceConsoleConnectionResourceDependencies + `
resource "oci_core_instance_console_connection" "test_instance_console_connection" {
	#Required
	instance_id = "${oci_core_instance.test_instance.id}"
	public_key = "${var.instance_console_connection_public_key}"
}
`

	InstanceConsoleConnectionResourceConfig = InstanceConsoleConnectionResourceDependencies + `
resource "oci_core_instance_console_connection" "test_instance_console_connection" {
	#Required
	instance_id = "${oci_core_instance.test_instance.id}"
	public_key = "${var.instance_console_connection_public_key}"

	#Optional
	defined_tags = "${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "${var.instance_console_connection_defined_tags_value}")}"
	freeform_tags = "${var.instance_console_connection_freeform_tags}"
}
`
	InstanceConsoleConnectionPropertyVariables = `
variable "instance_console_connection_defined_tags_value" { default = "value" }
variable "instance_console_connection_freeform_tags" { default = {"Department"= "Finance"} }
variable "instance_console_connection_public_key" { default = "ssh-rsa KKKLK3NzaC1yc2EAAAADAQABAAABAQC+UC9MFNA55NIVtKPIBCNw7++ACXhD0hx+Zyj25JfHykjz/QU3Q5FAU3DxDbVXyubgXfb/GJnrKRY8O4QDdvnZZRvQFFEOaApThAmCAM5MuFUIHdFvlqP+0W+ZQnmtDhwVe2NCfcmOrMuaPEgOKO3DOW6I/qOOdO691Xe2S9NgT9HhN0ZfFtEODVgvYulgXuCCXsJs+NUqcHAOxxFUmwkbPvYi0P0e2DT8JKeiOOC8VKUEgvVx+GKmqasm+Y6zHFW7vv3g2GstE1aRs3mttHRoC/JPM86PRyIxeWXEMzyG5wHqUu4XZpDbnWNxi6ugxnAGiL3CrIFdCgRNgHz5qS1l MustWin" }

`
	InstanceConsoleConnectionResourceDependencies = DefinedTagsDependencies + InstancePropertyVariables + InstanceResourceAsDependencyConfig
)

func TestCoreInstanceConsoleConnectionResource_basic(t *testing.T) {
	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getRequiredEnvSetting("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_core_instance_console_connection.test_instance_console_connection"
	datasourceName := "data.oci_core_instance_console_connections.test_instance_console_connections"

	resource.Test(t, resource.TestCase{
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		CheckDestroy: testAccCheckCoreInstanceConsoleConnectionDestroy,
		Steps: []resource.TestStep{
			// verify create
			{
				Config: config + InstanceConsoleConnectionPropertyVariables + compartmentIdVariableStr + InstanceConsoleConnectionRequiredOnlyResource,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "instance_id"),
					resource.TestCheckResourceAttr(resourceName, "public_key", "ssh-rsa KKKLK3NzaC1yc2EAAAADAQABAAABAQC+UC9MFNA55NIVtKPIBCNw7++ACXhD0hx+Zyj25JfHykjz/QU3Q5FAU3DxDbVXyubgXfb/GJnrKRY8O4QDdvnZZRvQFFEOaApThAmCAM5MuFUIHdFvlqP+0W+ZQnmtDhwVe2NCfcmOrMuaPEgOKO3DOW6I/qOOdO691Xe2S9NgT9HhN0ZfFtEODVgvYulgXuCCXsJs+NUqcHAOxxFUmwkbPvYi0P0e2DT8JKeiOOC8VKUEgvVx+GKmqasm+Y6zHFW7vv3g2GstE1aRs3mttHRoC/JPM86PRyIxeWXEMzyG5wHqUu4XZpDbnWNxi6ugxnAGiL3CrIFdCgRNgHz5qS1l MustWin"),
				),
			},

			// delete before next create
			{
				Config: config + compartmentIdVariableStr + InstanceConsoleConnectionResourceDependencies,
			},
			// verify create with optionals
			{
				Config: config + InstanceConsoleConnectionPropertyVariables + compartmentIdVariableStr + InstanceConsoleConnectionResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "instance_id"),
					resource.TestCheckResourceAttr(resourceName, "public_key", "ssh-rsa KKKLK3NzaC1yc2EAAAADAQABAAABAQC+UC9MFNA55NIVtKPIBCNw7++ACXhD0hx+Zyj25JfHykjz/QU3Q5FAU3DxDbVXyubgXfb/GJnrKRY8O4QDdvnZZRvQFFEOaApThAmCAM5MuFUIHdFvlqP+0W+ZQnmtDhwVe2NCfcmOrMuaPEgOKO3DOW6I/qOOdO691Xe2S9NgT9HhN0ZfFtEODVgvYulgXuCCXsJs+NUqcHAOxxFUmwkbPvYi0P0e2DT8JKeiOOC8VKUEgvVx+GKmqasm+Y6zHFW7vv3g2GstE1aRs3mttHRoC/JPM86PRyIxeWXEMzyG5wHqUu4XZpDbnWNxi6ugxnAGiL3CrIFdCgRNgHz5qS1l MustWin"),
				),
			},

			// verify datasource
			{
				Config: config + `
variable "instance_console_connection_defined_tags_value" { default = "value" }
variable "instance_console_connection_freeform_tags" { default = {"Department"= "Finance"} }
variable "instance_console_connection_public_key" { default = "ssh-rsa KKKLK3NzaC1yc2EAAAADAQABAAABAQC+UC9MFNA55NIVtKPIBCNw7++ACXhD0hx+Zyj25JfHykjz/QU3Q5FAU3DxDbVXyubgXfb/GJnrKRY8O4QDdvnZZRvQFFEOaApThAmCAM5MuFUIHdFvlqP+0W+ZQnmtDhwVe2NCfcmOrMuaPEgOKO3DOW6I/qOOdO691Xe2S9NgT9HhN0ZfFtEODVgvYulgXuCCXsJs+NUqcHAOxxFUmwkbPvYi0P0e2DT8JKeiOOC8VKUEgvVx+GKmqasm+Y6zHFW7vv3g2GstE1aRs3mttHRoC/JPM86PRyIxeWXEMzyG5wHqUu4XZpDbnWNxi6ugxnAGiL3CrIFdCgRNgHz5qS1l MustWin" }

data "oci_core_instance_console_connections" "test_instance_console_connections" {
	#Required
	compartment_id = "${var.compartment_id}"

	#Optional
	instance_id = "${oci_core_instance.test_instance.id}"

    filter {
    	name = "id"
    	values = ["${oci_core_instance_console_connection.test_instance_console_connection.id}"]
    }
}
                ` + compartmentIdVariableStr + InstanceConsoleConnectionResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttrSet(datasourceName, "instance_id"),

					resource.TestCheckResourceAttr(datasourceName, "instance_console_connections.#", "1"),
					resource.TestCheckResourceAttr(datasourceName, "instance_console_connections.0.defined_tags.%", "1"),
					resource.TestCheckResourceAttr(datasourceName, "instance_console_connections.0.freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(datasourceName, "instance_console_connections.0.instance_id"),
				),
			},
			// verify resource import
			{
				Config:            config,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateVerifyIgnore: []string{
					"public_key",
				},
				ResourceName: resourceName,
			},
		},
	})
}

func testAccCheckCoreInstanceConsoleConnectionDestroy(s *terraform.State) error {
	noResourceFound := true
	client := testAccProvider.Meta().(*OracleClients).computeClient
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_core_instance_console_connection" {
			noResourceFound = false
			request := oci_core.GetInstanceConsoleConnectionRequest{}

			tmp := rs.Primary.ID
			request.InstanceConsoleConnectionId = &tmp

			_, err := client.GetInstanceConsoleConnection(context.Background(), request)

			if err == nil {
				return fmt.Errorf("resource still exists")
			}
			//Verify that exception is for '404 not found'.
			if failure, isServiceError := common.IsServiceError(err); !isServiceError || failure.GetHTTPStatusCode() != 404 {
				return err
			}
		}
	}
	if noResourceFound {
		return fmt.Errorf("at least one resource was expected from the state file, but could not be found")
	}

	return nil
}
