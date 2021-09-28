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
	CatalogRequiredOnlyResource = CatalogResourceDependencies +
		generateResourceFromRepresentationMap("oci_datacatalog_catalog", "test_catalog", Required, Create, catalogRepresentation)

	CatalogResourceConfig = CatalogResourceDependencies +
		generateResourceFromRepresentationMap("oci_datacatalog_catalog", "test_catalog", Optional, Update, catalogRepresentation)

	catalogSingularDataSourceRepresentation = map[string]interface{}{
		"catalog_id": Representation{repType: Required, create: `${oci_datacatalog_catalog.test_catalog.id}`},
	}

	catalogDataSourceRepresentation = map[string]interface{}{
		"compartment_id": Representation{repType: Required, create: `${var.compartment_id}`},
		"display_name":   Representation{repType: Optional, create: `displayName`, update: `displayName2`},
		"state":          Representation{repType: Optional, create: `ACTIVE`},
		"filter":         RepresentationGroup{Required, catalogDataSourceFilterRepresentation}}
	catalogDataSourceFilterRepresentation = map[string]interface{}{
		"name":   Representation{repType: Required, create: `id`},
		"values": Representation{repType: Required, create: []string{`${oci_datacatalog_catalog.test_catalog.id}`}},
	}

	catalogRepresentation = map[string]interface{}{
		"compartment_id":                     Representation{repType: Required, create: `${var.compartment_id}`},
		"defined_tags":                       Representation{repType: Optional, create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"display_name":                       Representation{repType: Optional, create: `displayName`, update: `displayName2`},
		"freeform_tags":                      Representation{repType: Optional, create: map[string]string{"Department": "Finance"}, update: map[string]string{"Department": "Accounting"}},
		"attached_catalog_private_endpoints": Representation{repType: Optional, create: []string{`${oci_datacatalog_catalog_private_endpoint.test_catalog_private_endpoint.id}`}},
	}

	CatalogResourceDependencies = generateResourceFromRepresentationMap("oci_datacatalog_catalog_private_endpoint", "test_catalog_private_endpoint", Required, Create, catalogPrivateEndpointRepresentation) +
		generateResourceFromRepresentationMap("oci_core_subnet", "test_subnet", Required, Create, subnetRepresentation) +
		generateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", Required, Create, vcnRepresentation) +
		DefinedTagsDependencies
)

// issue-routing-tag: datacatalog/default
func TestDatacatalogCatalogResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatacatalogCatalogResource_basic")
	defer httpreplay.SaveScenario()

	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := getEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_datacatalog_catalog.test_catalog"
	datasourceName := "data.oci_datacatalog_catalogs.test_catalogs"
	singularDatasourceName := "data.oci_datacatalog_catalog.test_catalog"

	var resId, resId2 string
	// Save TF content to create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	saveConfigContent(config+compartmentIdVariableStr+CatalogResourceDependencies+
		generateResourceFromRepresentationMap("oci_datacatalog_catalog", "test_catalog", Optional, Create, catalogRepresentation), "datacatalog", "catalog", t)

	ResourceTest(t, testAccCheckDatacatalogCatalogDestroy, []resource.TestStep{
		// verify create
		{
			Config: config + compartmentIdVariableStr + CatalogResourceDependencies +
				generateResourceFromRepresentationMap("oci_datacatalog_catalog", "test_catalog", Required, Create, catalogRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),

				func(s *terraform.State) (err error) {
					resId, err = fromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next create
		{
			Config: config + compartmentIdVariableStr + CatalogResourceDependencies,
		},
		// verify create with optionals
		{
			Config: config + compartmentIdVariableStr + CatalogResourceDependencies +
				generateResourceFromRepresentationMap("oci_datacatalog_catalog", "test_catalog", Optional, Create, catalogRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "attached_catalog_private_endpoints.#", "1"),

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
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + CatalogResourceDependencies +
				generateResourceFromRepresentationMap("oci_datacatalog_catalog", "test_catalog", Optional, Create,
					representationCopyWithNewProperties(catalogRepresentation, map[string]interface{}{
						"compartment_id": Representation{repType: Required, create: `${var.compartment_id_for_update}`},
					})),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "attached_catalog_private_endpoints.#", "1"),

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
			Config: config + compartmentIdVariableStr + CatalogResourceDependencies +
				generateResourceFromRepresentationMap("oci_datacatalog_catalog", "test_catalog", Optional, Update, catalogRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "attached_catalog_private_endpoints.#", "1"),

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
				generateDataSourceFromRepresentationMap("oci_datacatalog_catalogs", "test_catalogs", Optional, Update, catalogDataSourceRepresentation) +
				compartmentIdVariableStr + CatalogResourceDependencies +
				generateResourceFromRepresentationMap("oci_datacatalog_catalog", "test_catalog", Optional, Update, catalogRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),

				resource.TestCheckResourceAttr(datasourceName, "catalogs.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "catalogs.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "catalogs.0.attached_catalog_private_endpoints.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "catalogs.0.defined_tags.%", "1"),
				resource.TestCheckResourceAttr(datasourceName, "catalogs.0.display_name", "displayName2"),
				resource.TestCheckResourceAttr(datasourceName, "catalogs.0.freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "catalogs.0.id"),
				resource.TestCheckResourceAttrSet(datasourceName, "catalogs.0.number_of_objects"),
				resource.TestCheckResourceAttrSet(datasourceName, "catalogs.0.state"),
				resource.TestCheckResourceAttrSet(datasourceName, "catalogs.0.time_created"),
				resource.TestCheckResourceAttrSet(datasourceName, "catalogs.0.time_updated"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				generateDataSourceFromRepresentationMap("oci_datacatalog_catalog", "test_catalog", Required, Create, catalogSingularDataSourceRepresentation) +
				compartmentIdVariableStr + CatalogResourceConfig,
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "catalog_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "attached_catalog_private_endpoints.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "defined_tags.%", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "number_of_objects"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "service_api_url"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
		// remove singular datasource from previous step so that it doesn't conflict with import tests
		{
			Config: config + compartmentIdVariableStr + CatalogResourceConfig,
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

func testAccCheckDatacatalogCatalogDestroy(s *terraform.State) error {
	noResourceFound := true
	client := testAccProvider.Meta().(*OracleClients).dataCatalogClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_datacatalog_catalog" {
			noResourceFound = false
			request := oci_datacatalog.GetCatalogRequest{}

			tmp := rs.Primary.ID
			request.CatalogId = &tmp

			request.RequestMetadata.RetryPolicy = getRetryPolicy(true, "datacatalog")

			response, err := client.GetCatalog(context.Background(), request)

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
	if !inSweeperExcludeList("DatacatalogCatalog") {
		resource.AddTestSweepers("DatacatalogCatalog", &resource.Sweeper{
			Name:         "DatacatalogCatalog",
			Dependencies: DependencyGraph["catalog"],
			F:            sweepDatacatalogCatalogResource,
		})
	}
}

func sweepDatacatalogCatalogResource(compartment string) error {
	dataCatalogClient := GetTestClients(&schema.ResourceData{}).dataCatalogClient()
	catalogIds, err := getCatalogIds(compartment)
	if err != nil {
		return err
	}
	for _, catalogId := range catalogIds {
		if ok := SweeperDefaultResourceId[catalogId]; !ok {
			deleteCatalogRequest := oci_datacatalog.DeleteCatalogRequest{}

			deleteCatalogRequest.CatalogId = &catalogId

			deleteCatalogRequest.RequestMetadata.RetryPolicy = getRetryPolicy(true, "datacatalog")
			_, error := dataCatalogClient.DeleteCatalog(context.Background(), deleteCatalogRequest)
			if error != nil {
				fmt.Printf("Error deleting Catalog %s %s, It is possible that the resource is already deleted. Please verify manually \n", catalogId, error)
				continue
			}
			waitTillCondition(testAccProvider, &catalogId, catalogSweepWaitCondition, time.Duration(3*time.Minute),
				catalogSweepResponseFetchOperation, "datacatalog", true)
		}
	}
	return nil
}

func getCatalogIds(compartment string) ([]string, error) {
	ids := getResourceIdsToSweep(compartment, "CatalogId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	dataCatalogClient := GetTestClients(&schema.ResourceData{}).dataCatalogClient()

	listCatalogsRequest := oci_datacatalog.ListCatalogsRequest{}
	listCatalogsRequest.CompartmentId = &compartmentId
	listCatalogsRequest.LifecycleState = oci_datacatalog.ListCatalogsLifecycleStateActive
	listCatalogsResponse, err := dataCatalogClient.ListCatalogs(context.Background(), listCatalogsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting Catalog list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, catalog := range listCatalogsResponse.Items {
		id := *catalog.Id
		resourceIds = append(resourceIds, id)
		addResourceIdToSweeperResourceIdMap(compartmentId, "CatalogId", id)
	}
	return resourceIds, nil
}

func catalogSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if catalogResponse, ok := response.Response.(oci_datacatalog.GetCatalogResponse); ok {
		return catalogResponse.LifecycleState != oci_datacatalog.LifecycleStateDeleted
	}
	return false
}

func catalogSweepResponseFetchOperation(client *OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.dataCatalogClient().GetCatalog(context.Background(), oci_datacatalog.GetCatalogRequest{
		CatalogId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
