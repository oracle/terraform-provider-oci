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
	"github.com/oracle/oci-go-sdk/v46/common"
	oci_service_catalog "github.com/oracle/oci-go-sdk/v46/servicecatalog"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	ServiceCatalogRequiredOnlyResource = ServiceCatalogResourceDependencies +
		generateResourceFromRepresentationMap("oci_service_catalog_service_catalog", "test_service_catalog", Required, Create, serviceCatalogRepresentation)

	ServiceCatalogResourceConfig = ServiceCatalogResourceDependencies +
		generateResourceFromRepresentationMap("oci_service_catalog_service_catalog", "test_service_catalog", Optional, Update, serviceCatalogRepresentation)

	serviceCatalogSingularDataSourceRepresentation = map[string]interface{}{
		"service_catalog_id": Representation{repType: Required, create: `${oci_service_catalog_service_catalog.test_service_catalog.id}`},
	}

	serviceCatalogDataSourceRepresentation = map[string]interface{}{
		"compartment_id":     Representation{repType: Required, create: `${var.compartment_id}`},
		"display_name":       Representation{repType: Optional, create: `displayName`, update: `displayName2`},
		"service_catalog_id": Representation{repType: Optional, create: `${oci_service_catalog_service_catalog.test_service_catalog.id}`},
		"filter":             RepresentationGroup{Required, serviceCatalogDataSourceFilterRepresentation}}
	serviceCatalogDataSourceFilterRepresentation = map[string]interface{}{
		"name":   Representation{repType: Required, create: `id`},
		"values": Representation{repType: Required, create: []string{`${oci_service_catalog_service_catalog.test_service_catalog.id}`}},
	}

	serviceCatalogRepresentation = map[string]interface{}{
		"compartment_id": Representation{repType: Required, create: `${var.compartment_id}`},
		"display_name":   Representation{repType: Required, create: `displayName`, update: `displayName2`},
		"defined_tags":   Representation{repType: Optional, create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"freeform_tags":  Representation{repType: Optional, create: map[string]string{"bar-key": "value"}, update: map[string]string{"Department": "Accounting"}},
	}

	ServiceCatalogResourceDependencies = DefinedTagsDependencies
)

// issue-routing-tag: service_catalog/default
func TestServiceCatalogServiceCatalogResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestServiceCatalogServiceCatalogResource_basic")
	defer httpreplay.SaveScenario()

	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := getEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_service_catalog_service_catalog.test_service_catalog"
	datasourceName := "data.oci_service_catalog_service_catalogs.test_service_catalogs"
	singularDatasourceName := "data.oci_service_catalog_service_catalog.test_service_catalog"

	var resId, resId2 string
	// Save TF content to create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	saveConfigContent(config+compartmentIdVariableStr+ServiceCatalogResourceDependencies+
		generateResourceFromRepresentationMap("oci_service_catalog_service_catalog", "test_service_catalog", Optional, Create, serviceCatalogRepresentation), "servicecatalog", "serviceCatalog", t)

	ResourceTest(t, testAccCheckServiceCatalogServiceCatalogDestroy, []resource.TestStep{
		// verify create
		{
			Config: config + compartmentIdVariableStr + ServiceCatalogResourceDependencies +
				generateResourceFromRepresentationMap("oci_service_catalog_service_catalog", "test_service_catalog", Required, Create, serviceCatalogRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),

				func(s *terraform.State) (err error) {
					resId, err = fromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next create
		{
			Config: config + compartmentIdVariableStr + ServiceCatalogResourceDependencies,
		},
		// verify create with optionals
		{
			Config: config + compartmentIdVariableStr + ServiceCatalogResourceDependencies +
				generateResourceFromRepresentationMap("oci_service_catalog_service_catalog", "test_service_catalog", Optional, Create, serviceCatalogRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),

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
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + ServiceCatalogResourceDependencies +
				generateResourceFromRepresentationMap("oci_service_catalog_service_catalog", "test_service_catalog", Optional, Create,
					representationCopyWithNewProperties(serviceCatalogRepresentation, map[string]interface{}{
						"compartment_id": Representation{repType: Required, create: `${var.compartment_id_for_update}`},
					})),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),

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
			Config: config + compartmentIdVariableStr + ServiceCatalogResourceDependencies +
				generateResourceFromRepresentationMap("oci_service_catalog_service_catalog", "test_service_catalog", Optional, Update, serviceCatalogRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
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
			Config: config +
				generateDataSourceFromRepresentationMap("oci_service_catalog_service_catalogs", "test_service_catalogs", Optional, Update, serviceCatalogDataSourceRepresentation) +
				compartmentIdVariableStr + ServiceCatalogResourceDependencies +
				generateResourceFromRepresentationMap("oci_service_catalog_service_catalog", "test_service_catalog", Optional, Update, serviceCatalogRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
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
				generateDataSourceFromRepresentationMap("oci_service_catalog_service_catalog", "test_service_catalog", Required, Create, serviceCatalogSingularDataSourceRepresentation) +
				compartmentIdVariableStr + ServiceCatalogResourceConfig,
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "service_catalog_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "defined_tags.%", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
		// remove singular datasource from previous step so that it doesn't conflict with import tests
		{
			Config: config + compartmentIdVariableStr + ServiceCatalogResourceConfig,
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

func testAccCheckServiceCatalogServiceCatalogDestroy(s *terraform.State) error {
	noResourceFound := true
	client := testAccProvider.Meta().(*OracleClients).serviceCatalogClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_service_catalog_service_catalog" {
			noResourceFound = false
			request := oci_service_catalog.GetServiceCatalogRequest{}

			tmp := rs.Primary.ID
			request.ServiceCatalogId = &tmp

			request.RequestMetadata.RetryPolicy = getRetryPolicy(true, "service_catalog")

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
	if DependencyGraph == nil {
		initDependencyGraph()
	}
	if !inSweeperExcludeList("ServiceCatalogServiceCatalog") {
		resource.AddTestSweepers("ServiceCatalogServiceCatalog", &resource.Sweeper{
			Name:         "ServiceCatalogServiceCatalog",
			Dependencies: DependencyGraph["serviceCatalog"],
			F:            sweepServiceCatalogServiceCatalogResource,
		})
	}
}

func sweepServiceCatalogServiceCatalogResource(compartment string) error {
	serviceCatalogClient := GetTestClients(&schema.ResourceData{}).serviceCatalogClient()
	serviceCatalogIds, err := getServiceCatalogIds(compartment)
	if err != nil {
		return err
	}
	for _, serviceCatalogId := range serviceCatalogIds {
		if ok := SweeperDefaultResourceId[serviceCatalogId]; !ok {
			deleteServiceCatalogRequest := oci_service_catalog.DeleteServiceCatalogRequest{}

			deleteServiceCatalogRequest.ServiceCatalogId = &serviceCatalogId

			deleteServiceCatalogRequest.RequestMetadata.RetryPolicy = getRetryPolicy(true, "service_catalog")
			_, error := serviceCatalogClient.DeleteServiceCatalog(context.Background(), deleteServiceCatalogRequest)
			if error != nil {
				fmt.Printf("Error deleting ServiceCatalog %s %s, It is possible that the resource is already deleted. Please verify manually \n", serviceCatalogId, error)
				continue
			}
			waitTillCondition(testAccProvider, &serviceCatalogId, serviceCatalogSweepWaitCondition, time.Duration(3*time.Minute),
				serviceCatalogSweepResponseFetchOperation, "service_catalog", true)
		}
	}
	return nil
}

func getServiceCatalogIds(compartment string) ([]string, error) {
	ids := getResourceIdsToSweep(compartment, "ServiceCatalogId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	serviceCatalogClient := GetTestClients(&schema.ResourceData{}).serviceCatalogClient()

	listServiceCatalogsRequest := oci_service_catalog.ListServiceCatalogsRequest{}
	listServiceCatalogsRequest.CompartmentId = &compartmentId
	listServiceCatalogsResponse, err := serviceCatalogClient.ListServiceCatalogs(context.Background(), listServiceCatalogsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting ServiceCatalog list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, serviceCatalog := range listServiceCatalogsResponse.Items {
		id := *serviceCatalog.Id
		resourceIds = append(resourceIds, id)
		addResourceIdToSweeperResourceIdMap(compartmentId, "ServiceCatalogId", id)
	}
	return resourceIds, nil
}

func serviceCatalogSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if serviceCatalogResponse, ok := response.Response.(oci_service_catalog.GetServiceCatalogResponse); ok {
		return serviceCatalogResponse.LifecycleState != oci_service_catalog.ServiceCatalogLifecycleStateDeleted
	}
	return false
}

func serviceCatalogSweepResponseFetchOperation(client *OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.serviceCatalogClient().GetServiceCatalog(context.Background(), oci_service_catalog.GetServiceCatalogRequest{
		ServiceCatalogId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
