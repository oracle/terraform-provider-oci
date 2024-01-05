// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"testing"

	tf_datacatalog "github.com/oracle/terraform-provider-oci/internal/service/datacatalog"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
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
	DatacatalogConnectionRequiredOnlyResource = DatacatalogConnectionResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_datacatalog_connection", "test_connection", acctest.Required, acctest.Create, DatacatalogConnectionRepresentation)

	DatacatalogConnectionResourceConfig = DatacatalogConnectionResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_datacatalog_connection", "test_connection", acctest.Optional, acctest.Update, DatacatalogConnectionRepresentation)

	DatacatalogDatacatalogConnectionSingularDataSourceRepresentation = map[string]interface{}{
		"catalog_id":     acctest.Representation{RepType: acctest.Required, Create: `${oci_datacatalog_catalog.test_catalog.id}`},
		"connection_key": acctest.Representation{RepType: acctest.Required, Create: `${oci_datacatalog_connection.test_connection.id}`},
		"data_asset_key": acctest.Representation{RepType: acctest.Required, Create: `${oci_datacatalog_data_asset.test_data_asset.id}`},
		"fields":         acctest.Representation{RepType: acctest.Optional, Create: []string{}},
	}

	DatacatalogDatacatalogConnectionDataSourceRepresentation = map[string]interface{}{
		"catalog_id":            acctest.Representation{RepType: acctest.Required, Create: `${oci_datacatalog_catalog.test_catalog.id}`},
		"data_asset_key":        acctest.Representation{RepType: acctest.Required, Create: `${oci_datacatalog_data_asset.test_data_asset.id}`},
		"display_name":          acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"display_name_contains": acctest.Representation{RepType: acctest.Optional, Create: `displayName`},
		"fields":                acctest.Representation{RepType: acctest.Optional, Create: []string{}},
		"is_default":            acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"state":                 acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"filter":                acctest.RepresentationGroup{RepType: acctest.Required, Group: DatacatalogConnectionDataSourceFilterRepresentation}}
	DatacatalogConnectionDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `is_default`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`true`}},
	}

	DatacatalogConnectionRepresentation = map[string]interface{}{
		"catalog_id":     acctest.Representation{RepType: acctest.Required, Create: `${oci_datacatalog_catalog.test_catalog.id}`},
		"data_asset_key": acctest.Representation{RepType: acctest.Required, Create: `${oci_datacatalog_data_asset.test_data_asset.id}`},
		"display_name":   acctest.Representation{RepType: acctest.Required, Create: `displayName`, Update: `displayName2`},
		"properties":     acctest.Representation{RepType: acctest.Required, Create: map[string]string{"default.username": "scott", "default.passwordAndSecrets": "passwordField"}, Update: map[string]string{"default.username": "wardon", "default.passwordAndSecrets": "passwordField"}},
		"type_key":       acctest.Representation{RepType: acctest.Required, Create: `${data.oci_datacatalog_catalog_types.test_catalog_types_connection.type_collection.0.items.0.key}`},
		"description":    acctest.Representation{RepType: acctest.Optional, Create: `description`, Update: `description2`},
		"enc_properties": acctest.Representation{RepType: acctest.Required, Create: map[string]string{"default.password": "tiger"}, Update: map[string]string{"default.password": "lion"}},
		"is_default":     acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
	}

	DatacatalogConnectionResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_datacatalog_catalog", "test_catalog", acctest.Required, acctest.Create, DatacatalogCatalogRepresentation) +
		acctest.GenerateDataSourceFromRepresentationMap("oci_datacatalog_catalog_types", "test_catalog_types_dataAssset", acctest.Optional, acctest.Create,
			acctest.RepresentationCopyWithNewProperties(DatacatalogDatacatalogCatalogTypeDataSourceRepresentation, map[string]interface{}{
				"type_category": acctest.Representation{RepType: acctest.Optional, Create: `dataAsset`},
				"name":          acctest.Representation{RepType: acctest.Optional, Create: `Oracle Database`}})) +
		acctest.GenerateDataSourceFromRepresentationMap("oci_datacatalog_catalog_types", "test_catalog_types_connection", acctest.Optional, acctest.Create,
			acctest.RepresentationCopyWithNewProperties(DatacatalogDatacatalogCatalogTypeDataSourceRepresentation, map[string]interface{}{
				"type_category": acctest.Representation{RepType: acctest.Optional, Create: `connection`},
				"name":          acctest.Representation{RepType: acctest.Optional, Create: `JDBC`},
			})) +
		acctest.GenerateResourceFromRepresentationMap("oci_datacatalog_data_asset", "test_data_asset", acctest.Required, acctest.Create,
			acctest.RepresentationCopyWithNewProperties(DatacatalogDataAssetRepresentation, map[string]interface{}{
				"type_key": acctest.Representation{RepType: acctest.Required, Create: `${data.oci_datacatalog_catalog_types.test_catalog_types_dataAssset.type_collection.0.items.0.key}`}}))
)

// issue-routing-tag: datacatalog/default
func TestDatacatalogConnectionResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatacatalogConnectionResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_datacatalog_connection.test_connection"
	datasourceName := "data.oci_datacatalog_connections.test_connections"
	singularDatasourceName := "data.oci_datacatalog_connection.test_connection"

	var resId, resId2 string
	var compositeId string

	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+DatacatalogConnectionResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_datacatalog_connection", "test_connection", acctest.Optional, acctest.Create, DatacatalogConnectionRepresentation), "datacatalog", "connection", t)

	acctest.ResourceTest(t, testAccCheckDatacatalogConnectionDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + DatacatalogConnectionResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_datacatalog_connection", "test_connection", acctest.Required, acctest.Create, DatacatalogConnectionRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "catalog_id"),
				resource.TestCheckResourceAttrSet(resourceName, "data_asset_key"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "properties.%", "2"),
				resource.TestCheckResourceAttrSet(resourceName, "type_key"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + DatacatalogConnectionResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + DatacatalogConnectionResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_datacatalog_connection", "test_connection", acctest.Optional, acctest.Create, DatacatalogConnectionRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "catalog_id"),
				resource.TestCheckResourceAttrSet(resourceName, "data_asset_key"),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "enc_properties.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "is_default", "false"),
				resource.TestCheckResourceAttrSet(resourceName, "key"),
				resource.TestCheckResourceAttr(resourceName, "properties.%", "2"),
				resource.TestCheckResourceAttrSet(resourceName, "type_key"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					dataAssetKey, _ := acctest.FromInstanceState(s, resourceName, "data_asset_key")
					catalogId, _ := acctest.FromInstanceState(s, resourceName, "catalog_id")
					compositeId = tf_datacatalog.GetConnectionCompositeId(catalogId, resId, dataAssetKey)
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
			Config: config + compartmentIdVariableStr + DatacatalogConnectionResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_datacatalog_connection", "test_connection", acctest.Optional, acctest.Update, DatacatalogConnectionRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "catalog_id"),
				resource.TestCheckResourceAttrSet(resourceName, "data_asset_key"),
				resource.TestCheckResourceAttr(resourceName, "description", "description2"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "enc_properties.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "is_default", "true"),
				resource.TestCheckResourceAttrSet(resourceName, "key"),
				resource.TestCheckResourceAttr(resourceName, "properties.%", "2"),
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
				acctest.GenerateDataSourceFromRepresentationMap("oci_datacatalog_connections", "test_connections", acctest.Optional, acctest.Update, DatacatalogDatacatalogConnectionDataSourceRepresentation) +
				compartmentIdVariableStr + DatacatalogConnectionResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_datacatalog_connection", "test_connection", acctest.Optional, acctest.Update, DatacatalogConnectionRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
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
				acctest.GenerateDataSourceFromRepresentationMap("oci_datacatalog_connection", "test_connection", acctest.Required, acctest.Create, DatacatalogDatacatalogConnectionSingularDataSourceRepresentation) +
				compartmentIdVariableStr + DatacatalogConnectionResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "catalog_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "connection_key"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "data_asset_key"),

				resource.TestCheckResourceAttrSet(singularDatasourceName, "data_asset_key"),
				resource.TestCheckResourceAttr(singularDatasourceName, "description", "description2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "external_key"),
				resource.TestCheckResourceAttr(singularDatasourceName, "is_default", "true"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "key"),
				resource.TestCheckResourceAttr(singularDatasourceName, "properties.%", "2"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "type_key"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "uri"),
			),
		},
		// verify resource import
		{
			Config:            config + DatacatalogConnectionRequiredOnlyResource,
			ImportState:       true,
			ImportStateVerify: true,
			ImportStateIdFunc: getDataAssetConnectionImportId(resourceName),
			ImportStateVerifyIgnore: []string{
				"enc_properties",
			},
			ResourceName: resourceName,
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
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).DataCatalogClient()
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

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "datacatalog")

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
