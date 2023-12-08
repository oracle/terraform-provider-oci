// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
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
	oci_datacatalog "github.com/oracle/oci-go-sdk/v65/datacatalog"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	DatacatalogCatalogRequiredOnlyResource = DatacatalogCatalogResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_datacatalog_catalog", "test_catalog", acctest.Required, acctest.Create, DatacatalogCatalogRepresentation)

	DatacatalogCatalogResourceConfig = DatacatalogCatalogResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_datacatalog_catalog", "test_catalog", acctest.Optional, acctest.Update, DatacatalogCatalogRepresentation)

	DatacatalogDatacatalogCatalogSingularDataSourceRepresentation = map[string]interface{}{
		"catalog_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_datacatalog_catalog.test_catalog.id}`},
	}

	DatacatalogDatacatalogCatalogDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":   acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"state":          acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"filter":         acctest.RepresentationGroup{RepType: acctest.Required, Group: DatacatalogCatalogDataSourceFilterRepresentation}}
	DatacatalogCatalogDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_datacatalog_catalog.test_catalog.id}`}},
	}

	DatacatalogCatalogRepresentation = map[string]interface{}{
		"compartment_id":                     acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"defined_tags":                       acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"display_name":                       acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"freeform_tags":                      acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"attached_catalog_private_endpoints": acctest.Representation{RepType: acctest.Optional, Create: []string{`${oci_datacatalog_catalog_private_endpoint.test_catalog_private_endpoint.id}`}},
		"lifecycle":                          acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreDatacatalogDefinednSystemTagsChangesRepresentation},
	}

	// need to ignore the defined tags created by the OCI service tenancy
	ignoreDatacatalogDefinednSystemTagsChangesRepresentation = map[string]interface{}{
		"ignore_changes": acctest.Representation{RepType: acctest.Required, Create: []string{`defined_tags`, `system_tags`}},
	}

	DatacatalogCatalogResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_datacatalog_catalog_private_endpoint", "test_catalog_private_endpoint", acctest.Required, acctest.Create, catalogPrivateEndpointRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_subnet", "test_subnet", acctest.Required, acctest.Create, CoreSubnetRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", acctest.Required, acctest.Create, CoreVcnRepresentation) +
		DefinedTagsDependencies
)

// issue-routing-tag: datacatalog/default
func TestDatacatalogCatalogResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatacatalogCatalogResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_datacatalog_catalog.test_catalog"
	datasourceName := "data.oci_datacatalog_catalogs.test_catalogs"
	singularDatasourceName := "data.oci_datacatalog_catalog.test_catalog"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+DatacatalogCatalogResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_datacatalog_catalog", "test_catalog", acctest.Optional, acctest.Create, DatacatalogCatalogRepresentation), "datacatalog", "catalog", t)

	acctest.ResourceTest(t, testAccCheckDatacatalogCatalogDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + DatacatalogCatalogResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_datacatalog_catalog", "test_catalog", acctest.Required, acctest.Create, DatacatalogCatalogRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + DatacatalogCatalogResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + DatacatalogCatalogResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_datacatalog_catalog", "test_catalog", acctest.Optional, acctest.Create, DatacatalogCatalogRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "attached_catalog_private_endpoints.#", "1"),

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
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + DatacatalogCatalogResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_datacatalog_catalog", "test_catalog", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(DatacatalogCatalogRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "attached_catalog_private_endpoints.#", "1"),

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
			Config: config + compartmentIdVariableStr + DatacatalogCatalogResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_datacatalog_catalog", "test_catalog", acctest.Optional, acctest.Update, DatacatalogCatalogRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "attached_catalog_private_endpoints.#", "1"),

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
				acctest.GenerateDataSourceFromRepresentationMap("oci_datacatalog_catalogs", "test_catalogs", acctest.Optional, acctest.Update, DatacatalogDatacatalogCatalogDataSourceRepresentation) +
				compartmentIdVariableStr + DatacatalogCatalogResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_datacatalog_catalog", "test_catalog", acctest.Optional, acctest.Update, DatacatalogCatalogRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),

				resource.TestCheckResourceAttr(datasourceName, "catalogs.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "catalogs.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "catalogs.0.attached_catalog_private_endpoints.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "catalogs.0.display_name", "displayName2"),
				resource.TestCheckResourceAttr(datasourceName, "catalogs.0.freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "catalogs.0.id"),
				resource.TestCheckResourceAttr(datasourceName, "catalogs.0.locks.#", "0"),
				resource.TestCheckResourceAttrSet(datasourceName, "catalogs.0.number_of_objects"),
				resource.TestCheckResourceAttrSet(datasourceName, "catalogs.0.state"),
				resource.TestCheckResourceAttrSet(datasourceName, "catalogs.0.time_created"),
				resource.TestCheckResourceAttrSet(datasourceName, "catalogs.0.time_updated"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_datacatalog_catalog", "test_catalog", acctest.Required, acctest.Create, DatacatalogDatacatalogCatalogSingularDataSourceRepresentation) +
				compartmentIdVariableStr + DatacatalogCatalogResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "catalog_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "attached_catalog_private_endpoints.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "locks.#", "0"),
				resource.TestCheckResourceAttr(singularDatasourceName, "system_tags.%", "0"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "number_of_objects"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "service_api_url"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
		// verify resource import
		{
			Config:                  config + DatacatalogCatalogRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckDatacatalogCatalogDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).DataCatalogClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_datacatalog_catalog" {
			noResourceFound = false
			request := oci_datacatalog.GetCatalogRequest{}

			tmp := rs.Primary.ID
			request.CatalogId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "datacatalog")

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
	if acctest.DependencyGraph == nil {
		acctest.InitDependencyGraph()
	}
	if !acctest.InSweeperExcludeList("DatacatalogCatalog") {
		resource.AddTestSweepers("DatacatalogCatalog", &resource.Sweeper{
			Name:         "DatacatalogCatalog",
			Dependencies: acctest.DependencyGraph["catalog"],
			F:            sweepDatacatalogCatalogResource,
		})
	}
}

func sweepDatacatalogCatalogResource(compartment string) error {
	dataCatalogClient := acctest.GetTestClients(&schema.ResourceData{}).DataCatalogClient()
	catalogIds, err := getDatacatalogCatalogIds(compartment)
	if err != nil {
		return err
	}
	for _, catalogId := range catalogIds {
		if ok := acctest.SweeperDefaultResourceId[catalogId]; !ok {
			deleteCatalogRequest := oci_datacatalog.DeleteCatalogRequest{}

			deleteCatalogRequest.CatalogId = &catalogId

			deleteCatalogRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "datacatalog")
			_, error := dataCatalogClient.DeleteCatalog(context.Background(), deleteCatalogRequest)
			if error != nil {
				fmt.Printf("Error deleting Catalog %s %s, It is possible that the resource is already deleted. Please verify manually \n", catalogId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &catalogId, DatacatalogCatalogSweepWaitCondition, time.Duration(3*time.Minute),
				DatacatalogCatalogSweepResponseFetchOperation, "datacatalog", true)
		}
	}
	return nil
}

func getDatacatalogCatalogIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "CatalogId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	dataCatalogClient := acctest.GetTestClients(&schema.ResourceData{}).DataCatalogClient()

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
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "CatalogId", id)
	}
	return resourceIds, nil
}

func DatacatalogCatalogSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if catalogResponse, ok := response.Response.(oci_datacatalog.GetCatalogResponse); ok {
		return catalogResponse.LifecycleState != oci_datacatalog.LifecycleStateDeleted
	}
	return false
}

func DatacatalogCatalogSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.DataCatalogClient().GetCatalog(context.Background(), oci_datacatalog.GetCatalogRequest{
		CatalogId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
