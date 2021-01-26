// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"
	"fmt"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/oracle/oci-go-sdk/v34/common"
	oci_load_balancer "github.com/oracle/oci-go-sdk/v34/loadbalancer"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
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

	PathRouteSetResourceDependencies = generateResourceFromRepresentationMap("oci_load_balancer_backend_set", "test_backend_set", Required, Create, backendSetRepresentation) +
		generateResourceFromRepresentationMap("oci_load_balancer_certificate", "test_certificate", Required, Create, certificateRepresentation) +
		generateResourceFromRepresentationMap("oci_load_balancer_load_balancer", "test_load_balancer", Required, Create, loadBalancerRepresentation) +
		LoadBalancerSubnetDependencies
)

func TestLoadBalancerPathRouteSetResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestLoadBalancerPathRouteSetResource_basic")
	defer httpreplay.SaveScenario()

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
						if isEnableExportCompartment, _ := strconv.ParseBool(getEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
							if errExport := testExportCompartmentWithResourceName(&resId, &compartmentId, resourceName); errExport != nil {
								return errExport
							}
						}
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
	client := testAccProvider.Meta().(*OracleClients).loadBalancerClient()
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

			request.RequestMetadata.RetryPolicy = getRetryPolicy(true, "load_balancer")

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

func init() {
	if DependencyGraph == nil {
		initDependencyGraph()
	}
	if !inSweeperExcludeList("LoadBalancerPathRouteSet") {
		resource.AddTestSweepers("LoadBalancerPathRouteSet", &resource.Sweeper{
			Name:         "LoadBalancerPathRouteSet",
			Dependencies: DependencyGraph["pathRouteSet"],
			F:            sweepLoadBalancerPathRouteSetResource,
		})
	}
}

func sweepLoadBalancerPathRouteSetResource(compartment string) error {
	loadBalancerClient := GetTestClients(&schema.ResourceData{}).loadBalancerClient()
	pathRouteSetIds, err := getPathRouteSetIds(compartment)
	if err != nil {
		return err
	}
	for _, pathRouteSetId := range pathRouteSetIds {
		if ok := SweeperDefaultResourceId[pathRouteSetId]; !ok {
			deletePathRouteSetRequest := oci_load_balancer.DeletePathRouteSetRequest{}

			deletePathRouteSetRequest.RequestMetadata.RetryPolicy = getRetryPolicy(true, "load_balancer")
			_, error := loadBalancerClient.DeletePathRouteSet(context.Background(), deletePathRouteSetRequest)
			if error != nil {
				fmt.Printf("Error deleting PathRouteSet %s %s, It is possible that the resource is already deleted. Please verify manually \n", pathRouteSetId, error)
				continue
			}
		}
	}
	return nil
}

func getPathRouteSetIds(compartment string) ([]string, error) {
	ids := getResourceIdsToSweep(compartment, "PathRouteSetId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	loadBalancerClient := GetTestClients(&schema.ResourceData{}).loadBalancerClient()

	listPathRouteSetsRequest := oci_load_balancer.ListPathRouteSetsRequest{}

	loadBalancerIds, error := getLoadBalancerIds(compartment)
	if error != nil {
		return resourceIds, fmt.Errorf("Error getting loadBalancerId required for PathRouteSet resource requests \n")
	}
	for _, loadBalancerId := range loadBalancerIds {
		listPathRouteSetsRequest.LoadBalancerId = &loadBalancerId

		listPathRouteSetsResponse, err := loadBalancerClient.ListPathRouteSets(context.Background(), listPathRouteSetsRequest)

		if err != nil {
			return resourceIds, fmt.Errorf("Error getting PathRouteSet list for compartment id : %s , %s \n", compartmentId, err)
		}
		for _, pathRouteSet := range listPathRouteSetsResponse.Items {
			id := *pathRouteSet.Name
			resourceIds = append(resourceIds, id)
			addResourceIdToSweeperResourceIdMap(compartmentId, "PathRouteSetId", id)
		}

	}
	return resourceIds, nil
}
