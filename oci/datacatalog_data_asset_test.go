// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/oracle/oci-go-sdk/v55/common"
	oci_datacatalog "github.com/oracle/oci-go-sdk/v55/datacatalog"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	DataAssetRequiredOnlyResource = DataAssetResourceDependencies +
		GenerateResourceFromRepresentationMap("oci_datacatalog_data_asset", "test_data_asset", Required, Create, dataAssetRepresentation)

	DataAssetResourceConfig = DataAssetResourceDependencies +
		GenerateResourceFromRepresentationMap("oci_datacatalog_data_asset", "test_data_asset", Optional, Update, dataAssetRepresentation)

	dataAssetSingularDataSourceRepresentation = map[string]interface{}{
		"catalog_id":     Representation{RepType: Required, Create: `${oci_datacatalog_catalog.test_catalog.id}`},
		"data_asset_key": Representation{RepType: Required, Create: `${oci_datacatalog_data_asset.test_data_asset.id}`},
		"fields":         Representation{RepType: Optional, Create: []string{`key`}},
	}

	dataAssetDataSourceRepresentation = map[string]interface{}{
		"catalog_id":            Representation{RepType: Required, Create: `${oci_datacatalog_catalog.test_catalog.id}`},
		"display_name":          Representation{RepType: Optional, Create: `displayName`, Update: `displayName2`},
		"display_name_contains": Representation{RepType: Optional, Create: `displayNam`},
		"fields":                Representation{RepType: Optional, Create: []string{`key`}},
		"state":                 Representation{RepType: Optional, Create: `ACTIVE`},
		"type_key":              Representation{RepType: Required, Create: `${data.oci_datacatalog_catalog_types.test_catalog_types.type_collection.0.items.0.key}`},
		"filter":                RepresentationGroup{Required, dataAssetDataSourceFilterRepresentation},
	}
	dataAssetDataSourceFilterRepresentation = map[string]interface{}{
		"name":   Representation{RepType: Required, Create: `state`},
		"values": Representation{RepType: Required, Create: []string{`ACTIVE`}},
	}

	dataAssetRepresentation = map[string]interface{}{
		"catalog_id":   Representation{RepType: Required, Create: `${oci_datacatalog_catalog.test_catalog.id}`},
		"display_name": Representation{RepType: Required, Create: `displayName`, Update: `displayName2`},
		"type_key":     Representation{RepType: Required, Create: `${data.oci_datacatalog_catalog_types.test_catalog_types.type_collection.0.items.0.key}`},
		"description":  Representation{RepType: Optional, Create: `description`, Update: `description2`},
		"properties":   Representation{RepType: Required, Create: map[string]string{"default.host": "jbanford-host", "default.port": "1521", "default.database": "SID"}},
	}

	DataAssetResourceDependencies = GenerateResourceFromRepresentationMap("oci_datacatalog_catalog", "test_catalog", Required, Create, catalogRepresentation) +
		GenerateDataSourceFromRepresentationMap("oci_datacatalog_catalog_types", "test_catalog_types", Optional, Create,
			RepresentationCopyWithNewProperties(catalogTypeDataSourceRepresentation, map[string]interface{}{
				"type_category": Representation{RepType: Optional, Create: `dataAsset`},
				"name":          Representation{RepType: Optional, Create: `Oracle Database`},
			}))
)

// issue-routing-tag: datacatalog/default
func TestDatacatalogDataAssetResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatacatalogDataAssetResource_basic")
	defer httpreplay.SaveScenario()

	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_datacatalog_data_asset.test_data_asset"
	datasourceName := "data.oci_datacatalog_data_assets.test_data_assets"

	singularDatasourceName := "data.oci_datacatalog_data_asset.test_data_asset"

	var resId, resId2 string
	var compositeId string

	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	SaveConfigContent(config+compartmentIdVariableStr+DataAssetResourceDependencies+
		GenerateResourceFromRepresentationMap("oci_datacatalog_data_asset", "test_data_asset", Optional, Create, dataAssetRepresentation), "datacatalog", "dataAsset", t)

	ResourceTest(t, testAccCheckDatacatalogDataAssetDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + DataAssetResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_datacatalog_data_asset", "test_data_asset", Required, Create, dataAssetRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "catalog_id"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttrSet(resourceName, "type_key"),

				func(s *terraform.State) (err error) {
					resId, err = FromInstanceState(s, resourceName, "id")
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
				GenerateResourceFromRepresentationMap("oci_datacatalog_data_asset", "test_data_asset", Optional, Create, dataAssetRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "catalog_id"),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttrSet(resourceName, "key"),
				resource.TestCheckResourceAttr(resourceName, "properties.%", "3"),
				resource.TestCheckResourceAttrSet(resourceName, "type_key"),

				func(s *terraform.State) (err error) {
					resId, err = FromInstanceState(s, resourceName, "id")
					catalogId, _ := FromInstanceState(s, resourceName, "catalog_id")
					compositeId = getDataAssetCompositeId(catalogId, resId)
					log.Printf("[DEBUG] Composite ID to import: %s", compositeId)
					if isEnableExportCompartment, _ := strconv.ParseBool(getEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
						if errExport := TestExportCompartmentWithResourceName(&compositeId, &compartmentId, resourceName); errExport != nil {
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
				GenerateResourceFromRepresentationMap("oci_datacatalog_data_asset", "test_data_asset", Optional, Update, dataAssetRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "catalog_id"),
				resource.TestCheckResourceAttr(resourceName, "description", "description2"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttrSet(resourceName, "key"),
				resource.TestCheckResourceAttr(resourceName, "properties.%", "3"),
				resource.TestCheckResourceAttrSet(resourceName, "type_key"),

				func(s *terraform.State) (err error) {
					resId2, err = FromInstanceState(s, resourceName, "id")
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
				GenerateDataSourceFromRepresentationMap("oci_datacatalog_data_assets", "test_data_assets", Optional, Update, dataAssetDataSourceRepresentation) +
				compartmentIdVariableStr + DataAssetResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_datacatalog_data_asset", "test_data_asset", Optional, Update, dataAssetRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
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
				GenerateDataSourceFromRepresentationMap("oci_datacatalog_data_asset", "test_data_asset", Required, Create, dataAssetSingularDataSourceRepresentation) +
				compartmentIdVariableStr + DataAssetResourceConfig,
			Check: ComposeAggregateTestCheckFuncWrapper(
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
	client := testAccProvider.Meta().(*OracleClients).dataCatalogClient()
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

			request.RequestMetadata.RetryPolicy = GetRetryPolicy(true, "datacatalog")

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
