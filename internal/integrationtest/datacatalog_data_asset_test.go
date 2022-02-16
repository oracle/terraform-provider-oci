// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"testing"

	tf_datacatalog "github.com/terraform-providers/terraform-provider-oci/internal/service/datacatalog"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
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
	DataAssetRequiredOnlyResource = DataAssetResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_datacatalog_data_asset", "test_data_asset", acctest.Required, acctest.Create, dataAssetRepresentation)

	DataAssetResourceConfig = DataAssetResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_datacatalog_data_asset", "test_data_asset", acctest.Optional, acctest.Update, dataAssetRepresentation)

	dataAssetSingularDataSourceRepresentation = map[string]interface{}{
		"catalog_id":     acctest.Representation{RepType: acctest.Required, Create: `${oci_datacatalog_catalog.test_catalog.id}`},
		"data_asset_key": acctest.Representation{RepType: acctest.Required, Create: `${oci_datacatalog_data_asset.test_data_asset.id}`},
		"fields":         acctest.Representation{RepType: acctest.Optional, Create: []string{`key`}},
	}

	dataAssetDataSourceRepresentation = map[string]interface{}{
		"catalog_id":            acctest.Representation{RepType: acctest.Required, Create: `${oci_datacatalog_catalog.test_catalog.id}`},
		"display_name":          acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"display_name_contains": acctest.Representation{RepType: acctest.Optional, Create: `displayNam`},
		"fields":                acctest.Representation{RepType: acctest.Optional, Create: []string{`key`}},
		"state":                 acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"type_key":              acctest.Representation{RepType: acctest.Required, Create: `${data.oci_datacatalog_catalog_types.test_catalog_types.type_collection.0.items.0.key}`},
		"filter":                acctest.RepresentationGroup{RepType: acctest.Required, Group: dataAssetDataSourceFilterRepresentation},
	}
	dataAssetDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `state`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`ACTIVE`}},
	}

	dataAssetRepresentation = map[string]interface{}{
		"catalog_id":   acctest.Representation{RepType: acctest.Required, Create: `${oci_datacatalog_catalog.test_catalog.id}`},
		"display_name": acctest.Representation{RepType: acctest.Required, Create: `displayName`, Update: `displayName2`},
		"type_key":     acctest.Representation{RepType: acctest.Required, Create: `${data.oci_datacatalog_catalog_types.test_catalog_types.type_collection.0.items.0.key}`},
		"description":  acctest.Representation{RepType: acctest.Optional, Create: `description`, Update: `description2`},
		"properties":   acctest.Representation{RepType: acctest.Required, Create: map[string]string{"default.host": "jbanford-host", "default.port": "1521", "default.database": "SID"}},
	}

	DataAssetResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_datacatalog_catalog", "test_catalog", acctest.Required, acctest.Create, catalogRepresentation) +
		acctest.GenerateDataSourceFromRepresentationMap("oci_datacatalog_catalog_types", "test_catalog_types", acctest.Optional, acctest.Create,
			acctest.RepresentationCopyWithNewProperties(catalogTypeDataSourceRepresentation, map[string]interface{}{
				"type_category": acctest.Representation{RepType: acctest.Optional, Create: `dataAsset`},
				"name":          acctest.Representation{RepType: acctest.Optional, Create: `Oracle Database`},
			}))
)

// issue-routing-tag: datacatalog/default
func TestDatacatalogDataAssetResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatacatalogDataAssetResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_datacatalog_data_asset.test_data_asset"
	datasourceName := "data.oci_datacatalog_data_assets.test_data_assets"

	singularDatasourceName := "data.oci_datacatalog_data_asset.test_data_asset"

	var resId, resId2 string
	var compositeId string

	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+DataAssetResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_datacatalog_data_asset", "test_data_asset", acctest.Optional, acctest.Create, dataAssetRepresentation), "datacatalog", "dataAsset", t)

	acctest.ResourceTest(t, testAccCheckDatacatalogDataAssetDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + DataAssetResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_datacatalog_data_asset", "test_data_asset", acctest.Required, acctest.Create, dataAssetRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "catalog_id"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttrSet(resourceName, "type_key"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + DataAssetResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + DataAssetResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_datacatalog_data_asset", "test_data_asset", acctest.Optional, acctest.Create, dataAssetRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "catalog_id"),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttrSet(resourceName, "key"),
				resource.TestCheckResourceAttr(resourceName, "properties.%", "3"),
				resource.TestCheckResourceAttrSet(resourceName, "type_key"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					catalogId, _ := acctest.FromInstanceState(s, resourceName, "catalog_id")
					compositeId = tf_datacatalog.GetDataAssetCompositeId(catalogId, resId)
					log.Printf("[DEBUG] Composite ID to import: %s", compositeId)
					if isEnableExportCompartment, _ := strconv.ParseBool(utils.GetEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
						if errExport := resourcediscovery.TestExportCompartmentWithResourceName(&compositeId, &compartmentId, resourceName); errExport != nil {
							return errExport
						}
					}
					return err
				},
			),
		},

		// verify updates to updatable parameters
		{
			Config: config + compartmentIdVariableStr + DataAssetResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_datacatalog_data_asset", "test_data_asset", acctest.Optional, acctest.Update, dataAssetRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "catalog_id"),
				resource.TestCheckResourceAttr(resourceName, "description", "description2"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttrSet(resourceName, "key"),
				resource.TestCheckResourceAttr(resourceName, "properties.%", "3"),
				resource.TestCheckResourceAttrSet(resourceName, "type_key"),

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
				acctest.GenerateDataSourceFromRepresentationMap("oci_datacatalog_data_assets", "test_data_assets", acctest.Optional, acctest.Update, dataAssetDataSourceRepresentation) +
				compartmentIdVariableStr + DataAssetResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_datacatalog_data_asset", "test_data_asset", acctest.Optional, acctest.Update, dataAssetRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "catalog_id"),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(datasourceName, "display_name_contains", "displayNam"),
				resource.TestCheckResourceAttr(datasourceName, "fields.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),
				resource.TestCheckResourceAttrSet(datasourceName, "type_key"),

				resource.TestCheckResourceAttr(datasourceName, "data_asset_collection.#", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "data_asset_collection.0.items.0.key"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_datacatalog_data_asset", "test_data_asset", acctest.Required, acctest.Create, dataAssetSingularDataSourceRepresentation) +
				compartmentIdVariableStr + DataAssetResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "catalog_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "data_asset_key"),
				resource.TestCheckResourceAttr(singularDatasourceName, "description", "description2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "external_key"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "key"),
				resource.TestCheckResourceAttr(singularDatasourceName, "properties.%", "3"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_harvested"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "type_key"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "uri"),
			),
		},
		// remove singular datasource from previous step so that it doesn't conflict with import tests
		{
			Config: config + compartmentIdVariableStr + DataAssetResourceConfig,
		},
		// verify resource import
		{
			Config:                  config,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateIdFunc:       getDataAssetImportId(resourceName),
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func getDataAssetImportId(resourceName string) resource.ImportStateIdFunc {
	return func(s *terraform.State) (string, error) {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return "", fmt.Errorf("not found: %s", resourceName)
		}
		return fmt.Sprintf("catalogs/" + rs.Primary.Attributes["catalog_id"] + "/dataAssets/" + rs.Primary.Attributes["key"]), nil
	}
}

func testAccCheckDatacatalogDataAssetDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).DataCatalogClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_datacatalog_data_asset" {
			noResourceFound = false
			request := oci_datacatalog.GetDataAssetRequest{}

			if value, ok := rs.Primary.Attributes["catalog_id"]; ok {
				request.CatalogId = &value
			}

			if value, ok := rs.Primary.Attributes["key"]; ok {
				request.DataAssetKey = &value
			}

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "datacatalog")

			response, err := client.GetDataAsset(context.Background(), request)

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
