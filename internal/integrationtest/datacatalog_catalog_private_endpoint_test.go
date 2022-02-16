// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/oracle/oci-go-sdk/v58/common"
	oci_datacatalog "github.com/oracle/oci-go-sdk/v58/datacatalog"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
	"github.com/terraform-providers/terraform-provider-oci/internal/acctest"
	tf_client "github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/resourcediscovery"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"
	"github.com/terraform-providers/terraform-provider-oci/internal/utils"
)

var (
	CatalogPrivateEndpointRequiredOnlyResource = CatalogPrivateEndpointResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_datacatalog_catalog_private_endpoint", "test_catalog_private_endpoint", acctest.Required, acctest.Create, catalogPrivateEndpointRepresentation)

	CatalogPrivateEndpointResourceConfig = CatalogPrivateEndpointResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_datacatalog_catalog_private_endpoint", "test_catalog_private_endpoint", acctest.Optional, acctest.Update, catalogPrivateEndpointRepresentation)

	catalogPrivateEndpointSingularDataSourceRepresentation = map[string]interface{}{
		"catalog_private_endpoint_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_datacatalog_catalog_private_endpoint.test_catalog_private_endpoint.id}`},
	}

	catalogPrivateEndpointDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":   acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"state":          acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"filter":         acctest.RepresentationGroup{RepType: acctest.Required, Group: catalogPrivateEndpointDataSourceFilterRepresentation}}
	catalogPrivateEndpointDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_datacatalog_catalog_private_endpoint.test_catalog_private_endpoint.id}`}},
	}

	catalogPrivateEndpointRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"dns_zones":      acctest.Representation{RepType: acctest.Required, Create: []string{`custpvtsubnet.oraclevcn.com`}, Update: []string{`db.custpvtsubnet.oraclevcn.com`}},
		"subnet_id":      acctest.Representation{RepType: acctest.Required, Create: `${oci_core_subnet.test_subnet.id}`},
		"defined_tags":   acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"display_name":   acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"freeform_tags":  acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"bar-key": "value"}, Update: map[string]string{"Department": "Accounting"}},
	}

	CatalogPrivateEndpointResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_core_subnet", "test_subnet", acctest.Required, acctest.Create, subnetRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", acctest.Required, acctest.Create, vcnRepresentation) +
		DefinedTagsDependencies
)

// issue-routing-tag: datacatalog/default
func TestDatacatalogCatalogPrivateEndpointResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatacatalogCatalogPrivateEndpointResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_datacatalog_catalog_private_endpoint.test_catalog_private_endpoint"
	datasourceName := "data.oci_datacatalog_catalog_private_endpoints.test_catalog_private_endpoints"
	singularDatasourceName := "data.oci_datacatalog_catalog_private_endpoint.test_catalog_private_endpoint"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+CatalogPrivateEndpointResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_datacatalog_catalog_private_endpoint", "test_catalog_private_endpoint", acctest.Optional, acctest.Create, catalogPrivateEndpointRepresentation), "datacatalog", "catalogPrivateEndpoint", t)

	acctest.ResourceTest(t, testAccCheckDatacatalogCatalogPrivateEndpointDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + CatalogPrivateEndpointResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_datacatalog_catalog_private_endpoint", "test_catalog_private_endpoint", acctest.Required, acctest.Create, catalogPrivateEndpointRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "dns_zones.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "subnet_id"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + CatalogPrivateEndpointResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + CatalogPrivateEndpointResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_datacatalog_catalog_private_endpoint", "test_catalog_private_endpoint", acctest.Optional, acctest.Create, catalogPrivateEndpointRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "dns_zones.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "subnet_id"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					if isEnableExportCompartment, _ := strconv.ParseBool(utils.GetEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
						if errExport := resourcediscovery.TestExportCompartmentWithResourceName(&resId, &compartmentId, resourceName); errExport != nil {
							return errExport
						}
					}
					return err
				},
			),
		},

		// verify Update to the compartment (the compartment will be switched back in the next step)
		{
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + CatalogPrivateEndpointResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_datacatalog_catalog_private_endpoint", "test_catalog_private_endpoint", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(catalogPrivateEndpointRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "dns_zones.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "subnet_id"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
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
				acctest.GenerateResourceFromRepresentationMap("oci_datacatalog_catalog_private_endpoint", "test_catalog_private_endpoint", acctest.Optional, acctest.Update, catalogPrivateEndpointRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "dns_zones.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "subnet_id"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
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
				acctest.GenerateDataSourceFromRepresentationMap("oci_datacatalog_catalog_private_endpoints", "test_catalog_private_endpoints", acctest.Optional, acctest.Update, catalogPrivateEndpointDataSourceRepresentation) +
				compartmentIdVariableStr + CatalogPrivateEndpointResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_datacatalog_catalog_private_endpoint", "test_catalog_private_endpoint", acctest.Optional, acctest.Update, catalogPrivateEndpointRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),

				resource.TestCheckResourceAttr(datasourceName, "catalog_private_endpoints.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "catalog_private_endpoints.0.compartment_id", compartmentId),
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
				acctest.GenerateDataSourceFromRepresentationMap("oci_datacatalog_catalog_private_endpoint", "test_catalog_private_endpoint", acctest.Required, acctest.Create, catalogPrivateEndpointSingularDataSourceRepresentation) +
				compartmentIdVariableStr + CatalogPrivateEndpointResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "catalog_private_endpoint_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
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
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).DataCatalogClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_datacatalog_catalog_private_endpoint" {
			noResourceFound = false
			request := oci_datacatalog.GetCatalogPrivateEndpointRequest{}

			tmp := rs.Primary.ID
			request.CatalogPrivateEndpointId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "datacatalog")

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
	if acctest.DependencyGraph == nil {
		acctest.InitDependencyGraph()
	}
	if !acctest.InSweeperExcludeList("DatacatalogCatalogPrivateEndpoint") {
		resource.AddTestSweepers("DatacatalogCatalogPrivateEndpoint", &resource.Sweeper{
			Name:         "DatacatalogCatalogPrivateEndpoint",
			Dependencies: acctest.DependencyGraph["catalogPrivateEndpoint"],
			F:            sweepDatacatalogCatalogPrivateEndpointResource,
		})
	}
}

func sweepDatacatalogCatalogPrivateEndpointResource(compartment string) error {
	dataCatalogClient := acctest.GetTestClients(&schema.ResourceData{}).DataCatalogClient()
	catalogPrivateEndpointIds, err := getCatalogPrivateEndpointIds(compartment)
	if err != nil {
		return err
	}
	for _, catalogPrivateEndpointId := range catalogPrivateEndpointIds {
		if ok := acctest.SweeperDefaultResourceId[catalogPrivateEndpointId]; !ok {
			deleteCatalogPrivateEndpointRequest := oci_datacatalog.DeleteCatalogPrivateEndpointRequest{}

			deleteCatalogPrivateEndpointRequest.CatalogPrivateEndpointId = &catalogPrivateEndpointId

			deleteCatalogPrivateEndpointRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "datacatalog")
			_, error := dataCatalogClient.DeleteCatalogPrivateEndpoint(context.Background(), deleteCatalogPrivateEndpointRequest)
			if error != nil {
				fmt.Printf("Error deleting CatalogPrivateEndpoint %s %s, It is possible that the resource is already deleted. Please verify manually \n", catalogPrivateEndpointId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &catalogPrivateEndpointId, catalogPrivateEndpointSweepWaitCondition, time.Duration(3*time.Minute),
				catalogPrivateEndpointSweepResponseFetchOperation, "datacatalog", true)
		}
	}
	return nil
}

func getCatalogPrivateEndpointIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "CatalogPrivateEndpointId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	dataCatalogClient := acctest.GetTestClients(&schema.ResourceData{}).DataCatalogClient()

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
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "CatalogPrivateEndpointId", id)
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

func catalogPrivateEndpointSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.DataCatalogClient().GetCatalogPrivateEndpoint(context.Background(), oci_datacatalog.GetCatalogPrivateEndpointRequest{
		CatalogPrivateEndpointId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
