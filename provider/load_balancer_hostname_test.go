// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.com/oracle/oci-go-sdk/common"
	oci_load_balancer "github.com/oracle/oci-go-sdk/loadbalancer"
)

const (
	HostnameResourceConfig = HostnameResourceDependencies + `
resource "oci_load_balancer_hostname" "test_hostname" {
	#Required
	hostname = "${var.hostname_hostname}"
	load_balancer_id = "${oci_load_balancer_load_balancer.test_load_balancer.id}"
	name = "${var.hostname_name}"
}
`
	HostnamePropertyVariables = `
variable "hostname_hostname" { default = "app.example.com" }
variable "hostname_name" { default = "example_hostname_001" }

`
	HostnameResourceDependencies = LoadBalancerPropertyVariables + LoadBalancerResourceConfig
)

func TestLoadBalancerHostnameResource_basic(t *testing.T) {
	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_load_balancer_hostname.test_hostname"
	datasourceName := "data.oci_load_balancer_hostnames.test_hostnames"

	var resId, resId2 string

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		CheckDestroy: testAccCheckLoadBalancerHostnameDestroy,
		Steps: []resource.TestStep{
			// verify create
			{
				Config: config + HostnamePropertyVariables + compartmentIdVariableStr + HostnameResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "hostname", "app.example.com"),
					resource.TestCheckResourceAttrSet(resourceName, "load_balancer_id"),
					resource.TestCheckResourceAttr(resourceName, "name", "example_hostname_001"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},

			// verify updates to updatable parameters
			{
				Config: config + `
variable "hostname_hostname" { default = "hostname2" }
variable "hostname_name" { default = "example_hostname_001" }

                ` + compartmentIdVariableStr + HostnameResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "hostname", "hostname2"),
					resource.TestCheckResourceAttrSet(resourceName, "load_balancer_id"),
					resource.TestCheckResourceAttr(resourceName, "name", "example_hostname_001"),

					func(s *terraform.State) (err error) {
						resId2, err = fromInstanceState(s, resourceName, "id")
						if resId != resId2 {
							return fmt.Errorf("Resource recreated when it was supposed to be updated.")
						}
						return err
					},
				),
			},
			// verify datasource
			{
				Config: config + `
variable "hostname_hostname" { default = "hostname2" }
variable "hostname_name" { default = "example_hostname_001" }

data "oci_load_balancer_hostnames" "test_hostnames" {
	#Required
	load_balancer_id = "${oci_load_balancer_load_balancer.test_load_balancer.id}"

    filter {
    	name = "name"
    	values = ["${oci_load_balancer_hostname.test_hostname.name}"]
    }
}
                ` + compartmentIdVariableStr + HostnameResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(datasourceName, "load_balancer_id"),

					resource.TestCheckResourceAttr(datasourceName, "hostnames.#", "1"),
					resource.TestCheckResourceAttr(datasourceName, "hostnames.0.hostname", "hostname2"),
					resource.TestCheckResourceAttr(datasourceName, "hostnames.0.name", "example_hostname_001"),
				),
			},
		},
	})
}

func testAccCheckLoadBalancerHostnameDestroy(s *terraform.State) error {
	noResourceFound := true
	client := testAccProvider.Meta().(*OracleClients).loadBalancerClient
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_load_balancer_hostname" {
			noResourceFound = false
			request := oci_load_balancer.GetHostnameRequest{}

			if value, ok := rs.Primary.Attributes["load_balancer_id"]; ok {
				request.LoadBalancerId = &value
			}

			if value, ok := rs.Primary.Attributes["name"]; ok {
				request.Name = &value
			}

			_, err := client.GetHostname(context.Background(), request)

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
