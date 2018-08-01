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

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_load_balancer_path_route_set.test_path_route_set"
	datasourceName := "data.oci_load_balancer_path_route_sets.test_path_route_sets"

	var resId, resId2 string

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		CheckDestroy: testAccCheckLoadBalancerPathRouteSetDestroy,
		Steps: []resource.TestStep{
			// verify create
			{
				Config: config + PathRouteSetPropertyVariables + compartmentIdVariableStr + PathRouteSetResourceConfig,
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

func testAccCheckLoadBalancerPathRouteSetDestroy(s *terraform.State) error {
	noResourceFound := true
	client := testAccProvider.Meta().(*OracleClients).loadBalancerClient
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_load_balancer_path_route_set" {
			noResourceFound = false
			request := oci_load_balancer.GetPathRouteSetRequest{}

			if value, ok := rs.Primary.Attributes["load_balancer_id"]; ok {
				request.LoadBalancerId = &value
			}

			if value, ok := rs.Primary.Attributes["name"]; ok {
				request.PathRouteSetName = &value
			}

			_, err := client.GetPathRouteSet(context.Background(), request)

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
