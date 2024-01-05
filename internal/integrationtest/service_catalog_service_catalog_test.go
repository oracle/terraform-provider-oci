// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/oracle/oci-go-sdk/v65/common"
	oci_service_catalog "github.com/oracle/oci-go-sdk/v65/servicecatalog"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	ServiceCatalogServiceCatalogRequiredOnlyResource = ServiceCatalogServiceCatalogResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_service_catalog_service_catalog", "test_service_catalog", acctest.Required, acctest.Create, ServiceCatalogServiceCatalogRepresentation)

	ServiceCatalogServiceCatalogResourceConfig = ServiceCatalogServiceCatalogResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_service_catalog_service_catalog", "test_service_catalog", acctest.Optional, acctest.Update, ServiceCatalogServiceCatalogRepresentation)

	ServiceCatalogServiceCatalogServiceCatalogSingularDataSourceRepresentation = map[string]interface{}{
		"service_catalog_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_service_catalog_service_catalog.test_service_catalog.id}`},
	}

	ServiceCatalogServiceCatalogServiceCatalogDataSourceRepresentation = map[string]interface{}{
		"compartment_id":     acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":       acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"service_catalog_id": acctest.Representation{RepType: acctest.Optional, Create: `${oci_service_catalog_service_catalog.test_service_catalog.id}`},
		"filter":             acctest.RepresentationGroup{RepType: acctest.Required, Group: ServiceCatalogServiceCatalogDataSourceFilterRepresentation}}
	ServiceCatalogServiceCatalogDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_service_catalog_service_catalog.test_service_catalog.id}`}},
	}

	ServiceCatalogServiceCatalogRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":   acctest.Representation{RepType: acctest.Required, Create: `displayName`, Update: `displayName2`},
		"defined_tags":   acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"freeform_tags":  acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"bar-key": "value"}, Update: map[string]string{"Department": "Accounting"}},
	}

	ServiceCatalogServiceCatalogResourceDependencies = DefinedTagsDependencies
)

// issue-routing-tag: service_catalog/default
func TestServiceCatalogServiceCatalogResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestServiceCatalogServiceCatalogResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_service_catalog_service_catalog.test_service_catalog"
	datasourceName := "data.oci_service_catalog_service_catalogs.test_service_catalogs"
	singularDatasourceName := "data.oci_service_catalog_service_catalog.test_service_catalog"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+ServiceCatalogServiceCatalogResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_service_catalog_service_catalog", "test_service_catalog", acctest.Optional, acctest.Create, ServiceCatalogServiceCatalogRepresentation), "servicecatalog", "serviceCatalog", t)

	acctest.ResourceTest(t, testAccCheckServiceCatalogServiceCatalogDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + ServiceCatalogServiceCatalogResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_service_catalog_service_catalog", "test_service_catalog", acctest.Required, acctest.Create, ServiceCatalogServiceCatalogRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + ServiceCatalogServiceCatalogResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + ServiceCatalogServiceCatalogResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_service_catalog_service_catalog", "test_service_catalog", acctest.Optional, acctest.Create, ServiceCatalogServiceCatalogRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),

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
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + ServiceCatalogServiceCatalogResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_service_catalog_service_catalog", "test_service_catalog", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(ServiceCatalogServiceCatalogRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),

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
			Config: config + compartmentIdVariableStr + ServiceCatalogServiceCatalogResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_service_catalog_service_catalog", "test_service_catalog", acctest.Optional, acctest.Update, ServiceCatalogServiceCatalogRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),

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
				acctest.GenerateDataSourceFromRepresentationMap("oci_service_catalog_service_catalogs", "test_service_catalogs", acctest.Optional, acctest.Update, ServiceCatalogServiceCatalogServiceCatalogDataSourceRepresentation) +
				compartmentIdVariableStr + ServiceCatalogServiceCatalogResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_service_catalog_service_catalog", "test_service_catalog", acctest.Optional, acctest.Update, ServiceCatalogServiceCatalogRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttrSet(datasourceName, "service_catalog_id"),

				resource.TestCheckResourceAttr(datasourceName, "service_catalog_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "service_catalog_collection.0.items.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_service_catalog_service_catalog", "test_service_catalog", acctest.Required, acctest.Create, ServiceCatalogServiceCatalogServiceCatalogSingularDataSourceRepresentation) +
				compartmentIdVariableStr + ServiceCatalogServiceCatalogResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "service_catalog_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
		// verify resource import
		{
			Config:                  config + ServiceCatalogServiceCatalogRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckServiceCatalogServiceCatalogDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).ServiceCatalogClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_service_catalog_service_catalog" {
			noResourceFound = false
			request := oci_service_catalog.GetServiceCatalogRequest{}

			tmp := rs.Primary.ID
			request.ServiceCatalogId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "service_catalog")

			response, err := client.GetServiceCatalog(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_service_catalog.ServiceCatalogLifecycleStateDeleted): true,
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
	if !acctest.InSweeperExcludeList("ServiceCatalogServiceCatalog") {
		resource.AddTestSweepers("ServiceCatalogServiceCatalog", &resource.Sweeper{
			Name:         "ServiceCatalogServiceCatalog",
			Dependencies: acctest.DependencyGraph["serviceCatalog"],
			F:            sweepServiceCatalogServiceCatalogResource,
		})
	}
}

func sweepServiceCatalogServiceCatalogResource(compartment string) error {
	serviceCatalogClient := acctest.GetTestClients(&schema.ResourceData{}).ServiceCatalogClient()
	serviceCatalogIds, err := getServiceCatalogServiceCatalogIds(compartment)
	if err != nil {
		return err
	}
	for _, serviceCatalogId := range serviceCatalogIds {
		if ok := acctest.SweeperDefaultResourceId[serviceCatalogId]; !ok {
			deleteServiceCatalogRequest := oci_service_catalog.DeleteServiceCatalogRequest{}

			deleteServiceCatalogRequest.ServiceCatalogId = &serviceCatalogId

			deleteServiceCatalogRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "service_catalog")
			_, error := serviceCatalogClient.DeleteServiceCatalog(context.Background(), deleteServiceCatalogRequest)
			if error != nil {
				fmt.Printf("Error deleting ServiceCatalog %s %s, It is possible that the resource is already deleted. Please verify manually \n", serviceCatalogId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &serviceCatalogId, ServiceCatalogServiceCatalogSweepWaitCondition, time.Duration(3*time.Minute),
				ServiceCatalogServiceCatalogSweepResponseFetchOperation, "service_catalog", true)
		}
	}
	return nil
}

func getServiceCatalogServiceCatalogIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "ServiceCatalogId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	serviceCatalogClient := acctest.GetTestClients(&schema.ResourceData{}).ServiceCatalogClient()

	listServiceCatalogsRequest := oci_service_catalog.ListServiceCatalogsRequest{}
	listServiceCatalogsRequest.CompartmentId = &compartmentId
	listServiceCatalogsResponse, err := serviceCatalogClient.ListServiceCatalogs(context.Background(), listServiceCatalogsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting ServiceCatalog list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, serviceCatalog := range listServiceCatalogsResponse.Items {
		id := *serviceCatalog.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "ServiceCatalogId", id)
	}
	return resourceIds, nil
}

func ServiceCatalogServiceCatalogSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if serviceCatalogResponse, ok := response.Response.(oci_service_catalog.GetServiceCatalogResponse); ok {
		return serviceCatalogResponse.LifecycleState != oci_service_catalog.ServiceCatalogLifecycleStateDeleted
	}
	return false
}

func ServiceCatalogServiceCatalogSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.ServiceCatalogClient().GetServiceCatalog(context.Background(), oci_service_catalog.GetServiceCatalogRequest{
		ServiceCatalogId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
