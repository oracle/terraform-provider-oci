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
	"github.com/oracle/oci-go-sdk/v42/common"
	oci_datacatalog "github.com/oracle/oci-go-sdk/v42/datacatalog"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	ConnectionRequiredOnlyResource = ConnectionResourceDependencies +
		generateResourceFromRepresentationMap("oci_datacatalog_connection", "test_connection", Required, Create, connectionRepresentation)

	ConnectionResourceConfig = ConnectionResourceDependencies +
		generateResourceFromRepresentationMap("oci_datacatalog_connection", "test_connection", Optional, Update, connectionRepresentation)

	connectionSingularDataSourceRepresentation = map[string]interface{}{
		"catalog_id":     Representation{repType: Required, create: `${oci_datacatalog_catalog.test_catalog.id}`},
		"connection_key": Representation{repType: Required, create: `${oci_datacatalog_connection.test_connection.id}`},
		"data_asset_key": Representation{repType: Required, create: `${oci_datacatalog_data_asset.test_data_asset.id}`},
		"fields":         Representation{repType: Optional, create: []string{}},
	}

	connectionDataSourceRepresentation = map[string]interface{}{
		"catalog_id":            Representation{repType: Required, create: `${oci_datacatalog_catalog.test_catalog.id}`},
		"data_asset_key":        Representation{repType: Required, create: `${oci_datacatalog_data_asset.test_data_asset.id}`},
		"display_name":          Representation{repType: Optional, create: `displayName`, update: `displayName2`},
		"display_name_contains": Representation{repType: Optional, create: `displayName`},
		"fields":                Representation{repType: Optional, create: []string{}},
		"is_default":            Representation{repType: Optional, create: `false`, update: `true`},
		"state":                 Representation{repType: Optional, create: `ACTIVE`},
		"filter":                RepresentationGroup{Required, connectionDataSourceFilterRepresentation}}
	connectionDataSourceFilterRepresentation = map[string]interface{}{
		"name":   Representation{repType: Required, create: `is_default`},
		"values": Representation{repType: Required, create: []string{`true`}},
	}

	connectionRepresentation = map[string]interface{}{
		"catalog_id":     Representation{repType: Required, create: `${oci_datacatalog_catalog.test_catalog.id}`},
		"data_asset_key": Representation{repType: Required, create: `${oci_datacatalog_data_asset.test_data_asset.id}`},
		"display_name":   Representation{repType: Required, create: `displayName`, update: `displayName2`},
		"properties":     Representation{repType: Required, create: map[string]string{"default.username": "scott"}, update: map[string]string{"default.username": "wardon"}},
		"type_key":       Representation{repType: Required, create: `${data.oci_datacatalog_catalog_types.test_catalog_types_connection.type_collection.0.items.0.key}`},
		"description":    Representation{repType: Optional, create: `description`, update: `description2`},
		"enc_properties": Representation{repType: Required, create: map[string]string{"default.password": "tiger"}, update: map[string]string{"default.password": "lion"}},
		"is_default":     Representation{repType: Optional, create: `false`, update: `true`},
	}

	ConnectionResourceDependencies = generateResourceFromRepresentationMap("oci_datacatalog_catalog", "test_catalog", Required, Create, catalogRepresentation) +
		generateDataSourceFromRepresentationMap("oci_datacatalog_catalog_types", "test_catalog_types_dataAssset", Optional, Create,
			representationCopyWithNewProperties(catalogTypeDataSourceRepresentation, map[string]interface{}{
				"type_category": Representation{repType: Optional, create: `dataAsset`},
				"name":          Representation{repType: Optional, create: `Oracle Database`}})) +
		generateDataSourceFromRepresentationMap("oci_datacatalog_catalog_types", "test_catalog_types_connection", Optional, Create,
			representationCopyWithNewProperties(catalogTypeDataSourceRepresentation, map[string]interface{}{
				"type_category": Representation{repType: Optional, create: `connection`},
				"name":          Representation{repType: Optional, create: `JDBC`},
			})) +
		generateResourceFromRepresentationMap("oci_datacatalog_data_asset", "test_data_asset", Required, Create,
			representationCopyWithNewProperties(dataAssetRepresentation, map[string]interface{}{
				"type_key": Representation{repType: Required, create: `${data.oci_datacatalog_catalog_types.test_catalog_types_dataAssset.type_collection.0.items.0.key}`}}))
)

func TestDatacatalogConnectionResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatacatalogConnectionResource_basic")
	defer httpreplay.SaveScenario()

	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_datacatalog_connection.test_connection"
	datasourceName := "data.oci_datacatalog_connections.test_connections"
	singularDatasourceName := "data.oci_datacatalog_connection.test_connection"

	var resId, resId2 string
	var compositeId string

	// Save TF content to create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	saveConfigContent(config+compartmentIdVariableStr+ConnectionResourceDependencies+
		generateResourceFromRepresentationMap("oci_datacatalog_connection", "test_connection", Optional, Create, connectionRepresentation), "datacatalog", "connection", t)

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		CheckDestroy: testAccCheckDatacatalogConnectionDestroy,
		Steps: []resource.TestStep{
			// verify create
			{
				Config: config + compartmentIdVariableStr + ConnectionResourceDependencies +
					generateResourceFromRepresentationMap("oci_datacatalog_connection", "test_connection", Required, Create, connectionRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "catalog_id"),
					resource.TestCheckResourceAttrSet(resourceName, "data_asset_key"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
					resource.TestCheckResourceAttr(resourceName, "properties.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "type_key"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},

			// delete before next create
			{
				Config: config + compartmentIdVariableStr + ConnectionResourceDependencies,
			},
			// verify create with optionals
			{
				Config: config + compartmentIdVariableStr + ConnectionResourceDependencies +
					generateResourceFromRepresentationMap("oci_datacatalog_connection", "test_connection", Optional, Create, connectionRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "catalog_id"),
					resource.TestCheckResourceAttrSet(resourceName, "data_asset_key"),
					resource.TestCheckResourceAttr(resourceName, "description", "description"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
					resource.TestCheckResourceAttr(resourceName, "enc_properties.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "is_default", "false"),
					resource.TestCheckResourceAttrSet(resourceName, "key"),
					resource.TestCheckResourceAttr(resourceName, "properties.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "type_key"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						dataAssetKey, _ := fromInstanceState(s, resourceName, "data_asset_key")
						catalogId, _ := fromInstanceState(s, resourceName, "catalog_id")
						compositeId = getConnectionCompositeId(catalogId, resId, dataAssetKey)
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
				Config: config + compartmentIdVariableStr + ConnectionResourceDependencies +
					generateResourceFromRepresentationMap("oci_datacatalog_connection", "test_connection", Optional, Update, connectionRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "catalog_id"),
					resource.TestCheckResourceAttrSet(resourceName, "data_asset_key"),
					resource.TestCheckResourceAttr(resourceName, "description", "description2"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
					resource.TestCheckResourceAttr(resourceName, "enc_properties.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "is_default", "true"),
					resource.TestCheckResourceAttrSet(resourceName, "key"),
					resource.TestCheckResourceAttr(resourceName, "properties.%", "1"),
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
					generateDataSourceFromRepresentationMap("oci_datacatalog_connections", "test_connections", Optional, Update, connectionDataSourceRepresentation) +
					compartmentIdVariableStr + ConnectionResourceDependencies +
					generateResourceFromRepresentationMap("oci_datacatalog_connection", "test_connection", Optional, Update, connectionRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(datasourceName, "catalog_id"),
					resource.TestCheckResourceAttrSet(datasourceName, "data_asset_key"),
					resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
					resource.TestCheckResourceAttr(datasourceName, "display_name_contains", "displayName"),
					resource.TestCheckResourceAttr(datasourceName, "is_default", "true"),
					resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),
					resource.TestCheckResourceAttr(datasourceName, "connection_collection.#", "1"),
				),
			},
			// verify singular datasource
			{
				Config: config +
					generateDataSourceFromRepresentationMap("oci_datacatalog_connection", "test_connection", Required, Create, connectionSingularDataSourceRepresentation) +
					compartmentIdVariableStr + ConnectionResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(singularDatasourceName, "catalog_id"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "connection_key"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "data_asset_key"),

					resource.TestCheckResourceAttrSet(singularDatasourceName, "data_asset_key"),
					resource.TestCheckResourceAttr(singularDatasourceName, "description", "description2"),
					resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "external_key"),
					resource.TestCheckResourceAttr(singularDatasourceName, "is_default", "true"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "key"),
					resource.TestCheckResourceAttr(singularDatasourceName, "properties.%", "1"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "type_key"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "uri"),
				),
			},
			// remove singular datasource from previous step so that it doesn't conflict with import tests
			{
				Config: config + compartmentIdVariableStr + ConnectionResourceConfig,
			},
			// verify resource import
			{
				Config:            config,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateIdFunc: getDataAssetConnectionImportId(resourceName),
				ImportStateVerifyIgnore: []string{
					"enc_properties",
				},
				ResourceName: resourceName,
			},
		},
	})
}

func getDataAssetConnectionImportId(resourceName string) resource.ImportStateIdFunc {
	return func(s *terraform.State) (string, error) {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return "", fmt.Errorf("not found: %s", resourceName)
		}
		return fmt.Sprintf("catalogs/" + rs.Primary.Attributes["catalog_id"] + "/dataAssets/" + rs.Primary.Attributes["data_asset_key"] + "/connections/" + rs.Primary.Attributes["key"]), nil
	}
}

func testAccCheckDatacatalogConnectionDestroy(s *terraform.State) error {
	noResourceFound := true
	client := testAccProvider.Meta().(*OracleClients).dataCatalogClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_datacatalog_connection" {
			noResourceFound = false
			request := oci_datacatalog.GetConnectionRequest{}

			if value, ok := rs.Primary.Attributes["catalog_id"]; ok {
				request.CatalogId = &value
			}

			tmp := rs.Primary.ID
			request.ConnectionKey = &tmp

			if value, ok := rs.Primary.Attributes["data_asset_key"]; ok {
				request.DataAssetKey = &value
			}

			request.RequestMetadata.RetryPolicy = getRetryPolicy(true, "datacatalog")

			response, err := client.GetConnection(context.Background(), request)

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
