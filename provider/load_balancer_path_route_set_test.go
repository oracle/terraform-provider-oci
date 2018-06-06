// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

const (
	PathRouteSetResourceConfig = PathRouteSetResourceDependencies + `
resource "oci_load_balancer_path_route_set" "test_path_route_set" {
	#Required
	load_balancer_id = "${oci_load_balancer_load_balancer.test_load_balancer.id}"
	name = "${var.path_route_set_name}"
	path_routes {
		#Required
		backend_set_name = "${oci_load_balancer_backend_set.test_backend_set.name}"
		path = "${var.path_route_set_path_routes_path}"
		path_match_type {
			#Required
			match_type = "${var.path_route_set_path_routes_path_match_type_match_type}"
		}
	}
}
`
	PathRouteSetPropertyVariables = `
variable "path_route_set_name" { default = "example_path_route_set" }
variable "path_route_set_path_routes_path" { default = "/example/video/123" }
variable "path_route_set_path_routes_path_match_type_match_type" { default = "EXACT_MATCH" }

`
	PathRouteSetResourceDependencies = BackendSetRequiredOnlyResource + BackendSetPropertyVariables
)

func TestLoadBalancerPathRouteSetResource_basic(t *testing.T) {
	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getRequiredEnvSetting("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_load_balancer_path_route_set.test_path_route_set"
	datasourceName := "data.oci_load_balancer_path_route_sets.test_path_route_sets"

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
				Config:            config + PathRouteSetPropertyVariables + compartmentIdVariableStr + PathRouteSetResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "load_balancer_id"),
					resource.TestCheckResourceAttr(resourceName, "name", "example_path_route_set"),
					resource.TestCheckResourceAttr(resourceName, "path_routes.#", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "path_routes.0.backend_set_name"),
					resource.TestCheckResourceAttr(resourceName, "path_routes.0.path", "/example/video/123"),
					resource.TestCheckResourceAttr(resourceName, "path_routes.0.path_match_type.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "path_routes.0.path_match_type.0.match_type", "EXACT_MATCH"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},

			// verify updates to updatable parameters
			{
				Config: config + `
variable "path_route_set_name" { default = "example_path_route_set" }
variable "path_route_set_path_routes_path" { default = "/example/video/123" }
variable "path_route_set_path_routes_path_match_type_match_type" { default = "EXACT_MATCH" }

                ` + compartmentIdVariableStr + PathRouteSetResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "load_balancer_id"),
					resource.TestCheckResourceAttr(resourceName, "name", "example_path_route_set"),
					resource.TestCheckResourceAttr(resourceName, "path_routes.#", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "path_routes.0.backend_set_name"),
					resource.TestCheckResourceAttr(resourceName, "path_routes.0.path", "/example/video/123"),
					resource.TestCheckResourceAttr(resourceName, "path_routes.0.path_match_type.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "path_routes.0.path_match_type.0.match_type", "EXACT_MATCH"),

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
variable "path_route_set_name" { default = "example_path_route_set" }
variable "path_route_set_path_routes_path" { default = "/example/video/123" }
variable "path_route_set_path_routes_path_match_type_match_type" { default = "EXACT_MATCH" }

data "oci_load_balancer_path_route_sets" "test_path_route_sets" {
	#Required
	load_balancer_id = "${oci_load_balancer_load_balancer.test_load_balancer.id}"

    filter {
    	name = "name"
    	values = ["${oci_load_balancer_path_route_set.test_path_route_set.name}"]
    }
}
                ` + compartmentIdVariableStr + PathRouteSetResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(datasourceName, "load_balancer_id"),

					resource.TestCheckResourceAttr(datasourceName, "path_route_sets.#", "1"),
					resource.TestCheckResourceAttr(datasourceName, "path_route_sets.0.name", "example_path_route_set"),
					resource.TestCheckResourceAttr(datasourceName, "path_route_sets.0.path_routes.#", "1"),
					resource.TestCheckResourceAttrSet(datasourceName, "path_route_sets.0.path_routes.0.backend_set_name"),
					resource.TestCheckResourceAttr(datasourceName, "path_route_sets.0.path_routes.0.path", "/example/video/123"),
					resource.TestCheckResourceAttr(datasourceName, "path_route_sets.0.path_routes.0.path_match_type.#", "1"),
					resource.TestCheckResourceAttr(datasourceName, "path_route_sets.0.path_routes.0.path_match_type.0.match_type", "EXACT_MATCH"),
				),
			},
		},
	})
}
