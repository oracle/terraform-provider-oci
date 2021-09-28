// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/oracle/oci-go-sdk/v48/common"
	oci_datacatalog "github.com/oracle/oci-go-sdk/v48/datacatalog"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	CatalogPrivateEndpointRequiredOnlyResource = CatalogPrivateEndpointResourceDependencies +
		generateResourceFromRepresentationMap("oci_datacatalog_catalog_private_endpoint", "test_catalog_private_endpoint", Required, Create, catalogPrivateEndpointRepresentation)

	CatalogPrivateEndpointResourceConfig = CatalogPrivateEndpointResourceDependencies +
		generateResourceFromRepresentationMap("oci_datacatalog_catalog_private_endpoint", "test_catalog_private_endpoint", Optional, Update, catalogPrivateEndpointRepresentation)

	catalogPrivateEndpointSingularDataSourceRepresentation = map[string]interface{}{
		"catalog_private_endpoint_id": Representation{repType: Required, create: `${oci_datacatalog_catalog_private_endpoint.test_catalog_private_endpoint.id}`},
	}

	catalogPrivateEndpointDataSourceRepresentation = map[string]interface{}{
		"compartment_id": Representation{repType: Required, create: `${var.compartment_id}`},
		"display_name":   Representation{repType: Optional, create: `displayName`, update: `displayName2`},
		"state":          Representation{repType: Optional, create: `ACTIVE`},
		"filter":         RepresentationGroup{Required, catalogPrivateEndpointDataSourceFilterRepresentation}}
	catalogPrivateEndpointDataSourceFilterRepresentation = map[string]interface{}{
		"name":   Representation{repType: Required, create: `id`},
		"values": Representation{repType: Required, create: []string{`${oci_datacatalog_catalog_private_endpoint.test_catalog_private_endpoint.id}`}},
	}

	catalogPrivateEndpointRepresentation = map[string]interface{}{
		"compartment_id": Representation{repType: Required, create: `${var.compartment_id}`},
		"dns_zones":      Representation{repType: Required, create: []string{`custpvtsubnet.oraclevcn.com`}, update: []string{`db.custpvtsubnet.oraclevcn.com`}},
		"subnet_id":      Representation{repType: Required, create: `${oci_core_subnet.test_subnet.id}`},
		"defined_tags":   Representation{repType: Optional, create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"display_name":   Representation{repType: Optional, create: `displayName`, update: `displayName2`},
		"freeform_tags":  Representation{repType: Optional, create: map[string]string{"bar-key": "value"}, update: map[string]string{"Department": "Accounting"}},
	}

	CatalogPrivateEndpointResourceDependencies = generateResourceFromRepresentationMap("oci_core_subnet", "test_subnet", Required, Create, subnetRepresentation) +
		generateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", Required, Create, vcnRepresentation) +
		DefinedTagsDependencies
)

// issue-routing-tag: datacatalog/default
func TestDatacatalogCatalogPrivateEndpointResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatacatalogCatalogPrivateEndpointResource_basic")
	defer httpreplay.SaveScenario()

	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := getEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_datacatalog_catalog_private_endpoint.test_catalog_private_endpoint"
	datasourceName := "data.oci_datacatalog_catalog_private_endpoints.test_catalog_private_endpoints"
	singularDatasourceName := "data.oci_datacatalog_catalog_private_endpoint.test_catalog_private_endpoint"

	var resId, resId2 string
	// Save TF content to create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	saveConfigContent(config+compartmentIdVariableStr+CatalogPrivateEndpointResourceDependencies+
		generateResourceFromRepresentationMap("oci_datacatalog_catalog_private_endpoint", "test_catalog_private_endpoint", Optional, Create, catalogPrivateEndpointRepresentation), "datacatalog", "catalogPrivateEndpoint", t)

	ResourceTest(t, testAccCheckDatacatalogCatalogPrivateEndpointDestroy, []resource.TestStep{
		// verify create
		{
			Config: config + compartmentIdVariableStr + CatalogPrivateEndpointResourceDependencies +
				generateResourceFromRepresentationMap("oci_datacatalog_catalog_private_endpoint", "test_catalog_private_endpoint", Required, Create, catalogPrivateEndpointRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "dns_zones.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "subnet_id"),

				func(s *terraform.State) (err error) {
					resId, err = fromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next create
		{
			Config: config + compartmentIdVariableStr + CatalogPrivateEndpointResourceDependencies,
		},
		// verify create with optionals
		{
			Config: config + compartmentIdVariableStr + CatalogPrivateEndpointResourceDependencies +
				generateResourceFromRepresentationMap("oci_datacatalog_catalog_private_endpoint", "test_catalog_private_endpoint", Optional, Create, catalogPrivateEndpointRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "dns_zones.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "subnet_id"),

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

		// verify update to the compartment (the compartment will be switched back in the next step)
		{
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + CatalogPrivateEndpointResourceDependencies +
				generateResourceFromRepresentationMap("oci_datacatalog_catalog_private_endpoint", "test_catalog_private_endpoint", Optional, Create,
					representationCopyWithNewProperties(catalogPrivateEndpointRepresentation, map[string]interface{}{
						"compartment_id": Representation{repType: Required, create: `${var.compartment_id_for_update}`},
					})),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "dns_zones.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "subnet_id"),

				func(s *terraform.State) (err error) {
					resId2, err = fromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("resource recreated when it was supposed to be updated")
					}
					return err
				},
			),
		},

		// verify updates to updatable parameters
		{
			Config: config + compartmentIdVariableStr + CatalogPrivateEndpointResourceDependencies +
				generateResourceFromRepresentationMap("oci_datacatalog_catalog_private_endpoint", "test_catalog_private_endpoint", Optional, Update, catalogPrivateEndpointRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "dns_zones.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "subnet_id"),

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
				generateDataSourceFromRepresentationMap("oci_datacatalog_catalog_private_endpoints", "test_catalog_private_endpoints", Optional, Update, catalogPrivateEndpointDataSourceRepresentation) +
				compartmentIdVariableStr + CatalogPrivateEndpointResourceDependencies +
				generateResourceFromRepresentationMap("oci_datacatalog_catalog_private_endpoint", "test_catalog_private_endpoint", Optional, Update, catalogPrivateEndpointRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),

				resource.TestCheckResourceAttr(datasourceName, "catalog_private_endpoints.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "catalog_private_endpoints.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "catalog_private_endpoints.0.defined_tags.%", "1"),
				resource.TestCheckResourceAttr(datasourceName, "catalog_private_endpoints.0.display_name", "displayName2"),
				resource.TestCheckResourceAttr(datasourceName, "catalog_private_endpoints.0.dns_zones.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "catalog_private_endpoints.0.freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "catalog_private_endpoints.0.id"),
				resource.TestCheckResourceAttrSet(datasourceName, "catalog_private_endpoints.0.state"),
				resource.TestCheckResourceAttrSet(datasourceName, "catalog_private_endpoints.0.subnet_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "catalog_private_endpoints.0.time_created"),
				resource.TestCheckResourceAttrSet(datasourceName, "catalog_private_endpoints.0.time_updated"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				generateDataSourceFromRepresentationMap("oci_datacatalog_catalog_private_endpoint", "test_catalog_private_endpoint", Required, Create, catalogPrivateEndpointSingularDataSourceRepresentation) +
				compartmentIdVariableStr + CatalogPrivateEndpointResourceConfig,
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "catalog_private_endpoint_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "defined_tags.%", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "dns_zones.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
		// remove singular datasource from previous step so that it doesn't conflict with import tests
		{
			Config: config + compartmentIdVariableStr + CatalogPrivateEndpointResourceConfig,
		},
		// verify resource import
		{
			Config:                  config,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckDatacatalogCatalogPrivateEndpointDestroy(s *terraform.State) error {
	noResourceFound := true
	client := testAccProvider.Meta().(*OracleClients).dataCatalogClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_datacatalog_catalog_private_endpoint" {
			noResourceFound = false
			request := oci_datacatalog.GetCatalogPrivateEndpointRequest{}

			tmp := rs.Primary.ID
			request.CatalogPrivateEndpointId = &tmp

			request.RequestMetadata.RetryPolicy = getRetryPolicy(true, "datacatalog")

			response, err := client.GetCatalogPrivateEndpoint(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_datacatalog.LifecycleStateDeleted): true,
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

func init() {
	if DependencyGraph == nil {
		initDependencyGraph()
	}
	if !inSweeperExcludeList("DatacatalogCatalogPrivateEndpoint") {
		resource.AddTestSweepers("DatacatalogCatalogPrivateEndpoint", &resource.Sweeper{
			Name:         "DatacatalogCatalogPrivateEndpoint",
			Dependencies: DependencyGraph["catalogPrivateEndpoint"],
			F:            sweepDatacatalogCatalogPrivateEndpointResource,
		})
	}
}

func sweepDatacatalogCatalogPrivateEndpointResource(compartment string) error {
	dataCatalogClient := GetTestClients(&schema.ResourceData{}).dataCatalogClient()
	catalogPrivateEndpointIds, err := getCatalogPrivateEndpointIds(compartment)
	if err != nil {
		return err
	}
	for _, catalogPrivateEndpointId := range catalogPrivateEndpointIds {
		if ok := SweeperDefaultResourceId[catalogPrivateEndpointId]; !ok {
			deleteCatalogPrivateEndpointRequest := oci_datacatalog.DeleteCatalogPrivateEndpointRequest{}

			deleteCatalogPrivateEndpointRequest.CatalogPrivateEndpointId = &catalogPrivateEndpointId

			deleteCatalogPrivateEndpointRequest.RequestMetadata.RetryPolicy = getRetryPolicy(true, "datacatalog")
			_, error := dataCatalogClient.DeleteCatalogPrivateEndpoint(context.Background(), deleteCatalogPrivateEndpointRequest)
			if error != nil {
				fmt.Printf("Error deleting CatalogPrivateEndpoint %s %s, It is possible that the resource is already deleted. Please verify manually \n", catalogPrivateEndpointId, error)
				continue
			}
			waitTillCondition(testAccProvider, &catalogPrivateEndpointId, catalogPrivateEndpointSweepWaitCondition, time.Duration(3*time.Minute),
				catalogPrivateEndpointSweepResponseFetchOperation, "datacatalog", true)
		}
	}
	return nil
}

func getCatalogPrivateEndpointIds(compartment string) ([]string, error) {
	ids := getResourceIdsToSweep(compartment, "CatalogPrivateEndpointId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	dataCatalogClient := GetTestClients(&schema.ResourceData{}).dataCatalogClient()

	listCatalogPrivateEndpointsRequest := oci_datacatalog.ListCatalogPrivateEndpointsRequest{}
	listCatalogPrivateEndpointsRequest.CompartmentId = &compartmentId
	listCatalogPrivateEndpointsRequest.LifecycleState = oci_datacatalog.ListCatalogPrivateEndpointsLifecycleStateActive
	listCatalogPrivateEndpointsResponse, err := dataCatalogClient.ListCatalogPrivateEndpoints(context.Background(), listCatalogPrivateEndpointsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting CatalogPrivateEndpoint list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, catalogPrivateEndpoint := range listCatalogPrivateEndpointsResponse.Items {
		id := *catalogPrivateEndpoint.Id
		resourceIds = append(resourceIds, id)
		addResourceIdToSweeperResourceIdMap(compartmentId, "CatalogPrivateEndpointId", id)
	}
	return resourceIds, nil
}

func catalogPrivateEndpointSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if catalogPrivateEndpointResponse, ok := response.Response.(oci_datacatalog.GetCatalogPrivateEndpointResponse); ok {
		return catalogPrivateEndpointResponse.LifecycleState != oci_datacatalog.LifecycleStateDeleted
	}
	return false
}

func catalogPrivateEndpointSweepResponseFetchOperation(client *OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.dataCatalogClient().GetCatalogPrivateEndpoint(context.Background(), oci_datacatalog.GetCatalogPrivateEndpointRequest{
		CatalogPrivateEndpointId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
