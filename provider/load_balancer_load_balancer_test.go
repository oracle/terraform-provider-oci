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
	LoadBalancerRequiredOnlyResource = LoadBalancerResourceDependencies + `
resource "oci_load_balancer_load_balancer" "test_load_balancer" {
	#Required
	compartment_id = "${var.compartment_id}"
	display_name = "${var.load_balancer_display_name}"
	shape = "${var.load_balancer_shape}"
	subnet_ids = ["${oci_core_subnet.lb_test_subnet_1.id}", "${oci_core_subnet.lb_test_subnet_2.id}"]
}
`

	LoadBalancerResourceConfig = LoadBalancerResourceDependencies + `
resource "oci_load_balancer_load_balancer" "test_load_balancer" {
	#Required
	compartment_id = "${var.compartment_id}"
	display_name = "${var.load_balancer_display_name}"
	shape = "${var.load_balancer_shape}"
	subnet_ids = ["${oci_core_subnet.lb_test_subnet_1.id}", "${oci_core_subnet.lb_test_subnet_2.id}"]

	#Optional
	is_private = "${var.load_balancer_is_private}"
}
`
	LoadBalancerPropertyVariables = `
variable "load_balancer_detail" { default = "detail" }
variable "load_balancer_display_name" { default = "example_load_balancer" }
variable "load_balancer_is_private" { default = false }
variable "load_balancer_shape" { default = "100Mbps" }
variable "load_balancer_state" { default = "ACTIVE" }

`

	LoadBalancerSubnetDependencies = `
	data "oci_load_balancer_shapes" "t" {
		compartment_id = "${var.compartment_id}"
	}

	data "oci_identity_availability_domains" "ADs" {
		compartment_id = "${var.compartment_id}"
	}

	resource "oci_core_subnet" "lb_test_subnet_1" {
		#Required
		availability_domain = "${lookup(data.oci_identity_availability_domains.ADs.availability_domains[0],"name")}"
		cidr_block = "10.0.0.0/24"
		compartment_id = "${var.compartment_id}"
		vcn_id = "${oci_core_vcn.test_vcn.id}"
		display_name        = "lbTestSubnet"
		security_list_ids = ["${oci_core_vcn.test_vcn.default_security_list_id}"]
	}
	
	resource "oci_core_subnet" "lb_test_subnet_2" {
		#Required
		availability_domain = "${lookup(data.oci_identity_availability_domains.ADs.availability_domains[1],"name")}"
		cidr_block = "10.0.1.0/24"
		compartment_id = "${var.compartment_id}"
		vcn_id = "${oci_core_vcn.test_vcn.id}"
		display_name        = "lbTestSubnet2"
		security_list_ids = ["${oci_core_vcn.test_vcn.default_security_list_id}"]
	}
`

	LoadBalancerResourceDependencies = VcnPropertyVariables + VcnResourceConfig + LoadBalancerSubnetDependencies
)

func TestLoadBalancerLoadBalancerResource_basic(t *testing.T) {
	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_load_balancer_load_balancer.test_load_balancer"
	datasourceName := "data.oci_load_balancer_load_balancers.test_load_balancers"

	var resId, resId2 string

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		CheckDestroy: testAccCheckLoadBalancerLoadBalancerDestroy,
		Steps: []resource.TestStep{
			// verify create
			{
				Config: config + LoadBalancerPropertyVariables + compartmentIdVariableStr + LoadBalancerRequiredOnlyResource,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "display_name", "example_load_balancer"),
					resource.TestCheckResourceAttr(resourceName, "shape", "100Mbps"),
					resource.TestCheckResourceAttr(resourceName, "subnet_ids.#", "2"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},

			// delete before next create
			{
				Config: config + compartmentIdVariableStr + LoadBalancerResourceDependencies,
			},
			// verify create with optionals
			{
				Config: config + LoadBalancerPropertyVariables + compartmentIdVariableStr + LoadBalancerResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "display_name", "example_load_balancer"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "is_private", "false"),
					resource.TestCheckResourceAttr(resourceName, "shape", "100Mbps"),
					resource.TestCheckResourceAttrSet(resourceName, "state"),
					resource.TestCheckResourceAttr(resourceName, "subnet_ids.#", "2"),
					resource.TestCheckResourceAttrSet(resourceName, "time_created"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},

			// verify updates to updatable parameters
			{
				Config: config + `
variable "load_balancer_detail" { default = "detail" }
variable "load_balancer_display_name" { default = "example_load_balancer" }
variable "load_balancer_is_private" { default = false }
variable "load_balancer_shape" { default = "100Mbps" }
variable "load_balancer_state" { default = "ACTIVE" }

                ` + compartmentIdVariableStr + LoadBalancerResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "display_name", "example_load_balancer"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "is_private", "false"),
					resource.TestCheckResourceAttr(resourceName, "shape", "100Mbps"),
					resource.TestCheckResourceAttrSet(resourceName, "state"),
					resource.TestCheckResourceAttr(resourceName, "subnet_ids.#", "2"),
					resource.TestCheckResourceAttrSet(resourceName, "time_created"),

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
variable "load_balancer_detail" { default = "detail" }
variable "load_balancer_display_name" { default = "example_load_balancer" }
variable "load_balancer_is_private" { default = false }
variable "load_balancer_shape" { default = "100Mbps" }
variable "load_balancer_state" { default = "ACTIVE" }

data "oci_load_balancer_load_balancers" "test_load_balancers" {
	#Required
	compartment_id = "${var.compartment_id}"

	#Optional
	detail = "${var.load_balancer_detail}"
	display_name = "${var.load_balancer_display_name}"
	state = "${var.load_balancer_state}"

    filter {
    	name = "id"
    	values = ["${oci_load_balancer_load_balancer.test_load_balancer.id}"]
    }
}
                ` + compartmentIdVariableStr + LoadBalancerResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(datasourceName, "detail", "detail"),
					resource.TestCheckResourceAttr(datasourceName, "display_name", "example_load_balancer"),
					resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),

					resource.TestCheckResourceAttr(datasourceName, "load_balancers.#", "1"),
					resource.TestCheckResourceAttr(datasourceName, "load_balancers.0.compartment_id", compartmentId),
					resource.TestCheckResourceAttr(datasourceName, "load_balancers.0.display_name", "example_load_balancer"),
					resource.TestCheckResourceAttrSet(datasourceName, "load_balancers.0.id"),
					resource.TestCheckResourceAttr(datasourceName, "load_balancers.0.is_private", "false"),
					resource.TestCheckResourceAttr(datasourceName, "load_balancers.0.shape", "100Mbps"),
					resource.TestCheckResourceAttrSet(datasourceName, "load_balancers.0.state"),
					resource.TestCheckResourceAttr(datasourceName, "load_balancers.0.subnet_ids.#", "2"),
					resource.TestCheckResourceAttrSet(datasourceName, "load_balancers.0.time_created"),
				),
			},
			// verify resource import
			{
				Config:                  config,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
				ResourceName:            resourceName,
			},
		},
	})
}

func testAccCheckLoadBalancerLoadBalancerDestroy(s *terraform.State) error {
	noResourceFound := true
	client := testAccProvider.Meta().(*OracleClients).loadBalancerClient
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_load_balancer_load_balancer" {
			noResourceFound = false
			request := oci_load_balancer.GetLoadBalancerRequest{}

			tmp := rs.Primary.ID
			request.LoadBalancerId = &tmp

			response, err := client.GetLoadBalancer(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_load_balancer.LoadBalancerLifecycleStateDeleted): true,
				}
				if _, ok := deletedLifecycleStates[string(response.LifecycleState)]; !ok {
					//resource lifecycle state is not in expected deleted lifecycle states.
					return fmt.Errorf("resource lifecycle state: %s is not in expected deleted lifecycle states", response.LifecycleState)
				}
				//resource lifecycle state is in expected deleted lifecycle states. continue with next one.
				continue
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
