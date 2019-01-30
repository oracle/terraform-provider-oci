// Copyright (c) 2017, 2019, Oracle and/or its affiliates. All rights reserved.

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

var (
	pathRouteSetDataSourceRepresentation = map[string]interface{}{
		"load_balancer_id": Representation{repType: Required, create: `${oci_load_balancer_load_balancer.test_load_balancer.id}`},
		"filter":           RepresentationGroup{Required, pathRouteSetDataSourceFilterRepresentation}}
	pathRouteSetDataSourceFilterRepresentation = map[string]interface{}{
		"name":   Representation{repType: Required, create: `name`},
		"values": Representation{repType: Required, create: []string{`${oci_load_balancer_path_route_set.test_path_route_set.name}`}},
	}

	pathRouteSetRepresentation = map[string]interface{}{
		"load_balancer_id": Representation{repType: Required, create: `${oci_load_balancer_load_balancer.test_load_balancer.id}`},
		"name":             Representation{repType: Required, create: `example_path_route_set`},
		"path_routes":      RepresentationGroup{Required, pathRouteSetPathRoutesRepresentation},
	}
	pathRouteSetPathRoutesRepresentation = map[string]interface{}{
		"backend_set_name": Representation{repType: Required, create: `${oci_load_balancer_backend_set.test_backend_set.name}`},
		"path":             Representation{repType: Required, create: `/example/video/123`, update: `path2`},
		"path_match_type":  RepresentationGroup{Required, pathRouteSetPathRoutesPathMatchTypeRepresentation},
	}
	pathRouteSetPathRoutesPathMatchTypeRepresentation = map[string]interface{}{
		"match_type": Representation{repType: Required, create: `EXACT_MATCH`},
	}

	PathRouteSetResourceDependencies = BackendSetRequiredOnlyResource
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
				Config: config + compartmentIdVariableStr + PathRouteSetResourceDependencies +
					generateResourceFromRepresentationMap("oci_load_balancer_path_route_set", "test_path_route_set", Required, Create, pathRouteSetRepresentation),
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
				Config: config + compartmentIdVariableStr + PathRouteSetResourceDependencies +
					generateResourceFromRepresentationMap("oci_load_balancer_path_route_set", "test_path_route_set", Optional, Update, pathRouteSetRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "load_balancer_id"),
					resource.TestCheckResourceAttr(resourceName, "name", "example_path_route_set"),
					resource.TestCheckResourceAttr(resourceName, "path_routes.#", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "path_routes.0.backend_set_name"),
					resource.TestCheckResourceAttr(resourceName, "path_routes.0.path", "path2"),
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
				Config: config +
					generateDataSourceFromRepresentationMap("oci_load_balancer_path_route_sets", "test_path_route_sets", Optional, Update, pathRouteSetDataSourceRepresentation) +
					compartmentIdVariableStr + PathRouteSetResourceDependencies +
					generateResourceFromRepresentationMap("oci_load_balancer_path_route_set", "test_path_route_set", Optional, Update, pathRouteSetRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(datasourceName, "load_balancer_id"),

					resource.TestCheckResourceAttr(datasourceName, "path_route_sets.#", "1"),
					resource.TestCheckResourceAttr(datasourceName, "path_route_sets.0.name", "example_path_route_set"),
					resource.TestCheckResourceAttr(datasourceName, "path_route_sets.0.path_routes.#", "1"),
					resource.TestCheckResourceAttrSet(datasourceName, "path_route_sets.0.path_routes.0.backend_set_name"),
					resource.TestCheckResourceAttr(datasourceName, "path_route_sets.0.path_routes.0.path", "path2"),
					resource.TestCheckResourceAttr(datasourceName, "path_route_sets.0.path_routes.0.path_match_type.#", "1"),
					resource.TestCheckResourceAttr(datasourceName, "path_route_sets.0.path_routes.0.path_match_type.0.match_type", "EXACT_MATCH"),
				),
			},
			// verify resource import
			{
				Config:            config,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateVerifyIgnore: []string{
					"state",
				},
				ResourceName: resourceName,
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
