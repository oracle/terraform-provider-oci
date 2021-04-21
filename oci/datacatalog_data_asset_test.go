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
	"github.com/oracle/oci-go-sdk/v40/common"
	oci_datacatalog "github.com/oracle/oci-go-sdk/v40/datacatalog"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	DataAssetRequiredOnlyResource = DataAssetResourceDependencies +
		generateResourceFromRepresentationMap("oci_datacatalog_data_asset", "test_data_asset", Required, Create, dataAssetRepresentation)

	DataAssetResourceConfig = DataAssetResourceDependencies +
		generateResourceFromRepresentationMap("oci_datacatalog_data_asset", "test_data_asset", Optional, Update, dataAssetRepresentation)

	dataAssetSingularDataSourceRepresentation = map[string]interface{}{
		"catalog_id":     Representation{repType: Required, create: `${oci_datacatalog_catalog.test_catalog.id}`},
		"data_asset_key": Representation{repType: Required, create: `${oci_datacatalog_data_asset.test_data_asset.id}`},
		"fields":         Representation{repType: Optional, create: []string{`key`}},
	}
	dataAssetDataSourceRepresentation = map[string]interface{}{
		"catalog_id":            Representation{repType: Required, create: `${oci_datacatalog_catalog.test_catalog.id}`},
		"display_name":          Representation{repType: Optional, create: `displayName`, update: `displayName2`},
		"display_name_contains": Representation{repType: Optional, create: `displayNam`},
		"fields":                Representation{repType: Optional, create: []string{`key`}},
		"state":                 Representation{repType: Optional, create: `ACTIVE`},
		"type_key":              Representation{repType: Required, create: `${data.oci_datacatalog_catalog_types.test_catalog_types.type_collection.0.items.0.key}`},
		"filter":                RepresentationGroup{Required, dataAssetDataSourceFilterRepresentation},
	}
	dataAssetDataSourceFilterRepresentation = map[string]interface{}{
		"name":   Representation{repType: Required, create: `state`},
		"values": Representation{repType: Required, create: []string{`ACTIVE`}},
	}

	dataAssetRepresentation = map[string]interface{}{
		"catalog_id":   Representation{repType: Required, create: `${oci_datacatalog_catalog.test_catalog.id}`},
		"display_name": Representation{repType: Required, create: `displayName`, update: `displayName2`},
		"type_key":     Representation{repType: Required, create: `${data.oci_datacatalog_catalog_types.test_catalog_types.type_collection.0.items.0.key}`},
		"description":  Representation{repType: Optional, create: `description`, update: `description2`},
		"properties":   Representation{repType: Required, create: map[string]string{"default.host": "jbanford-host", "default.port": "1521", "default.database": "SID"}, update: map[string]string{"default.host": "jbanford-host", "default.port": "1251", "default.database": "SID"}},
	}

	DataAssetResourceDependencies = generateResourceFromRepresentationMap("oci_datacatalog_catalog", "test_catalog", Required, Create, catalogRepresentation) +
		generateDataSourceFromRepresentationMap("oci_datacatalog_catalog_types", "test_catalog_types", Optional, Create,
			representationCopyWithNewProperties(catalogTypeDataSourceRepresentation, map[string]interface{}{
				"type_category": Representation{repType: Optional, create: `dataAsset`},
				"name":          Representation{repType: Optional, create: `Oracle Database`},
			}))
)

func TestDatacatalogDataAssetResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatacatalogDataAssetResource_basic")
	defer httpreplay.SaveScenario()

	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_datacatalog_data_asset.test_data_asset"
	datasourceName := "data.oci_datacatalog_data_assets.test_data_assets"

	singularDatasourceName := "data.oci_datacatalog_data_asset.test_data_asset"

	var resId, resId2 string
	var compositeId string

	// Save TF content to create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	saveConfigContent(config+compartmentIdVariableStr+DataAssetResourceDependencies+
		generateResourceFromRepresentationMap("oci_datacatalog_data_asset", "test_data_asset", Optional, Create, dataAssetRepresentation), "datacatalog", "dataAsset", t)

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		CheckDestroy: testAccCheckDatacatalogDataAssetDestroy,
		Steps: []resource.TestStep{
			// verify create
			{
				Config: config + compartmentIdVariableStr + DataAssetResourceDependencies +
					generateResourceFromRepresentationMap("oci_datacatalog_data_asset", "test_data_asset", Required, Create, dataAssetRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "catalog_id"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
					resource.TestCheckResourceAttrSet(resourceName, "type_key"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},

			// delete before next create
			{
				Config: config + compartmentIdVariableStr + DataAssetResourceDependencies,
			},
			// verify create with optionals
			{
				Config: config + compartmentIdVariableStr + DataAssetResourceDependencies +
					generateResourceFromRepresentationMap("oci_datacatalog_data_asset", "test_data_asset", Optional, Create, dataAssetRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "catalog_id"),
					resource.TestCheckResourceAttr(resourceName, "description", "description"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
					resource.TestCheckResourceAttrSet(resourceName, "key"),
					resource.TestCheckResourceAttr(resourceName, "properties.%", "3"),
					resource.TestCheckResourceAttrSet(resourceName, "type_key"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						catalogId, _ := fromInstanceState(s, resourceName, "catalog_id")
						compositeId = getDataAssetCompositeId(catalogId, resId)
						log.Printf("[DEBUG] Composite ID to import: %s", compositeId)
						if isEnableExportCompartment, _ := strconv.ParseBool(getEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
							if errExport := testExportCompartmentWithResourceName(&compositeId, &compartmentId, resourceName); errExport != nil {
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
					generateResourceFromRepresentationMap("oci_datacatalog_data_asset", "test_data_asset", Optional, Update, dataAssetRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "catalog_id"),
					resource.TestCheckResourceAttr(resourceName, "description", "description2"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
					resource.TestCheckResourceAttrSet(resourceName, "key"),
					resource.TestCheckResourceAttr(resourceName, "properties.%", "3"),
					resource.TestCheckResourceAttrSet(resourceName, "type_key"),

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
					generateDataSourceFromRepresentationMap("oci_datacatalog_data_assets", "test_data_assets", Optional, Update, dataAssetDataSourceRepresentation) +
					compartmentIdVariableStr + DataAssetResourceDependencies +
					generateResourceFromRepresentationMap("oci_datacatalog_data_asset", "test_data_asset", Optional, Update, dataAssetRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
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
					generateDataSourceFromRepresentationMap("oci_datacatalog_data_asset", "test_data_asset", Required, Create, dataAssetSingularDataSourceRepresentation) +
					compartmentIdVariableStr + DataAssetResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(singularDatasourceName, "catalog_id"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "data_asset_key"),

					resource.TestCheckResourceAttr(singularDatasourceName, "description", "description2"),
					resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "external_key"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "key"),
					resource.TestCheckResourceAttr(singularDatasourceName, "properties.%", "3"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
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

			request.RequestMetadata.RetryPolicy = getRetryPolicy(true, "datacatalog")

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
