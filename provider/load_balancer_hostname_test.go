// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
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

	compartmentId := getRequiredEnvSetting("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_load_balancer_hostname.test_hostname"
	datasourceName := "data.oci_load_balancer_hostnames.test_hostnames"

	var resId, resId2 string

	resource.Test(t, resource.TestCase{
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		Steps: []resource.TestStep{
			// verify create
			{
				ImportState:       true,
				ImportStateVerify: true,
				Config:            config + HostnamePropertyVariables + compartmentIdVariableStr + HostnameResourceConfig,
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
