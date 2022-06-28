// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/oracle/oci-go-sdk/v65/common"
	oci_data_connectivity "github.com/oracle/oci-go-sdk/v65/dataconnectivity"

	"terraform-provider-oci/httpreplay"
	"terraform-provider-oci/internal/acctest"
	tf_client "terraform-provider-oci/internal/client"
	"terraform-provider-oci/internal/resourcediscovery"
	"terraform-provider-oci/internal/tfresource"
	"terraform-provider-oci/internal/utils"
)

var (
	DataConnectivityRegistryConnectionRequiredOnlyResource = DataConnectivityRegistryConnectionResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_data_connectivity_registry_connection", "test_registry_connection", acctest.Required, acctest.Create, DataConnectivityRegistryConnectionRepresentation)

	DataConnectivityRegistryConnectionResourceConfig = DataConnectivityRegistryConnectionResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_data_connectivity_registry_connection", "test_registry_connection", acctest.Optional, acctest.Update, DataConnectivityRegistryConnectionRepresentation)

	DataConnectivityDataConnectivityRegistryConnectionSingularDataSourceRepresentation = map[string]interface{}{
		"connection_key": acctest.Representation{RepType: acctest.Required, Create: `${oci_data_connectivity_registry_connection.test_registry_connection.key}`},
		"registry_id":    acctest.Representation{RepType: acctest.Required, Create: `${oci_data_connectivity_registry.test_registry.id}`},
	}

	DataConnectivityDataConnectivityRegistryConnectionDataSourceRepresentation = map[string]interface{}{
		"data_asset_key": acctest.Representation{RepType: acctest.Required, Create: `${oci_data_connectivity_registry_data_asset.test_registry_data_asset.key}`},
		"registry_id":    acctest.Representation{RepType: acctest.Required, Create: `${oci_data_connectivity_registry.test_registry.id}`}}
	DataConnectivityRegistryConnectionDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `name`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_data_connectivity_registry_connection.test_registry_connection.name}`}},
	}

	DataConnectivityRegistryConnectionRepresentation = map[string]interface{}{
		"identifier":            acctest.Representation{RepType: acctest.Required, Create: `CONNECTIONIDENTIFIER`, Update: `CONNECTIONIDENTIFIER2`},
		"name":                  acctest.Representation{RepType: acctest.Required, Create: `connectionname`, Update: `connectionname2`},
		"properties":            acctest.Representation{RepType: acctest.Required, Create: map[string]string{"username": "username", "password": "password"}, Update: map[string]string{"username": "username", "password": "password"}},
		"registry_id":           acctest.Representation{RepType: acctest.Required, Create: `${oci_data_connectivity_registry.test_registry.id}`},
		"type":                  acctest.Representation{RepType: acctest.Required, Create: `GENERIC_JDBC_CONNECTION`},
		"connection_properties": acctest.RepresentationGroup{RepType: acctest.Optional, Group: DataConnectivityRegistryConnectionConnectionPropertiesRepresentation},
		"description":           acctest.Representation{RepType: acctest.Optional, Create: `description`, Update: `description2`},
		"is_default":            acctest.Representation{RepType: acctest.Required, Create: `false`, Update: `false`},
		"model_type":            acctest.Representation{RepType: acctest.Required, Create: `GENERIC_CONNECTION`},
		"registry_metadata":     acctest.RepresentationGroup{RepType: acctest.Required, Group: DataConnectivityRegistryConnectionRegistryMetadataRepresentation},
		"lifecycle":             acctest.RepresentationGroup{RepType: acctest.Required, Group: dcmsConnectionIgnoreChangesRepresentation},
	}
	dcmsConnectionIgnoreChangesRepresentation = map[string]interface{}{
		"ignore_changes": acctest.Representation{RepType: acctest.Required, Create: []string{`properties`}},
	}
	DataConnectivityRegistryConnectionConnectionPropertiesRepresentation = map[string]interface{}{
		"name":  acctest.Representation{RepType: acctest.Optional, Create: `name`, Update: `name2`},
		"value": acctest.Representation{RepType: acctest.Optional, Create: `value`, Update: `value2`},
	}
	DataConnectivityRegistryConnectionMetadataRepresentation = map[string]interface{}{
		"aggregator":       acctest.RepresentationGroup{RepType: acctest.Optional, Group: DataConnectivityRegistryConnectionMetadataAggregatorRepresentation},
		"aggregator_key":   acctest.Representation{RepType: acctest.Optional, Create: `aggregatorKey`, Update: `aggregatorKey2`},
		"created_by":       acctest.Representation{RepType: acctest.Optional, Create: `createdBy`, Update: `createdBy2`},
		"created_by_name":  acctest.Representation{RepType: acctest.Optional, Create: `createdByName`, Update: `createdByName2`},
		"identifier_path":  acctest.Representation{RepType: acctest.Optional, Create: `identifierPath`, Update: `identifierPath2`},
		"info_fields":      acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"infoFields": "infoFields"}, Update: map[string]string{"infoFields2": "infoFields2"}},
		"is_favorite":      acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"labels":           acctest.Representation{RepType: acctest.Optional, Create: []string{`labels`}, Update: []string{`labels2`}},
		"registry_version": acctest.Representation{RepType: acctest.Optional, Create: `10`, Update: `11`},
		"time_created":     acctest.Representation{RepType: acctest.Optional, Create: `timeCreated`, Update: `timeCreated2`},
		"time_updated":     acctest.Representation{RepType: acctest.Optional, Create: `timeUpdated`, Update: `timeUpdated2`},
		"updated_by":       acctest.Representation{RepType: acctest.Optional, Create: `updatedBy`, Update: `updatedBy2`},
		"updated_by_name":  acctest.Representation{RepType: acctest.Optional, Create: `updatedByName`, Update: `updatedByName2`},
	}
	DataConnectivityRegistryConnectionPrimarySchemaRepresentation = map[string]interface{}{
		"identifier":         acctest.Representation{RepType: acctest.Required, Create: `identifier`, Update: `identifier2`},
		"key":                acctest.Representation{RepType: acctest.Required, Create: `key`, Update: `key2`},
		"model_type":         acctest.Representation{RepType: acctest.Required, Create: `modelType`, Update: `modelType2`},
		"name":               acctest.Representation{RepType: acctest.Required, Create: `name`, Update: `name2`},
		"default_connection": acctest.Representation{RepType: acctest.Optional, Create: `defaultConnection`, Update: `defaultConnection2`},
		"description":        acctest.Representation{RepType: acctest.Optional, Create: `description`, Update: `description2`},
		"external_key":       acctest.Representation{RepType: acctest.Optional, Create: `externalKey`, Update: `externalKey2`},
		"is_has_containers":  acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"metadata":           acctest.RepresentationGroup{RepType: acctest.Optional, Group: DataConnectivityRegistryConnectionPrimarySchemaMetadataRepresentation},
		"model_version":      acctest.Representation{RepType: acctest.Optional, Create: `modelVersion`, Update: `modelVersion2`},
		"object_status":      acctest.Representation{RepType: acctest.Optional, Create: `10`, Update: `11`},
		"object_version":     acctest.Representation{RepType: acctest.Optional, Create: `10`, Update: `11`},
		"parent_ref":         acctest.RepresentationGroup{RepType: acctest.Optional, Group: DataConnectivityRegistryConnectionPrimarySchemaParentRefRepresentation},
		"resource_name":      acctest.Representation{RepType: acctest.Optional, Create: `resourceName`, Update: `resourceName2`},
	}
	DataConnectivityRegistryConnectionRegistryMetadataRepresentation = map[string]interface{}{
		"aggregator_key": acctest.Representation{RepType: acctest.Required, Create: `${oci_data_connectivity_registry_data_asset.test_registry_data_asset.key}`, Update: `${oci_data_connectivity_registry_data_asset.test_registry_data_asset.key}`},
	}
	DataConnectivityRegistryConnectionMetadataAggregatorRepresentation = map[string]interface{}{
		"description": acctest.Representation{RepType: acctest.Optional, Create: `description`, Update: `description2`},
		"identifier":  acctest.Representation{RepType: acctest.Optional, Create: `identifier`, Update: `identifier2`},
		"key":         acctest.Representation{RepType: acctest.Optional, Create: `key`, Update: `key2`},
		"name":        acctest.Representation{RepType: acctest.Optional, Create: `name`, Update: `name2`},
		"type":        acctest.Representation{RepType: acctest.Optional, Create: `type`, Update: `type2`},
	}
	DataConnectivityRegistryConnectionPrimarySchemaMetadataRepresentation = map[string]interface{}{
		"aggregator":       acctest.RepresentationGroup{RepType: acctest.Optional, Group: DataConnectivityRegistryConnectionPrimarySchemaMetadataAggregatorRepresentation},
		"aggregator_key":   acctest.Representation{RepType: acctest.Optional, Create: `aggregatorKey`, Update: `aggregatorKey2`},
		"created_by":       acctest.Representation{RepType: acctest.Optional, Create: `createdBy`, Update: `createdBy2`},
		"created_by_name":  acctest.Representation{RepType: acctest.Optional, Create: `createdByName`, Update: `createdByName2`},
		"identifier_path":  acctest.Representation{RepType: acctest.Optional, Create: `identifierPath`, Update: `identifierPath2`},
		"info_fields":      acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"infoFields": "infoFields"}, Update: map[string]string{"infoFields2": "infoFields2"}},
		"is_favorite":      acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"labels":           acctest.Representation{RepType: acctest.Optional, Create: []string{`labels`}, Update: []string{`labels2`}},
		"registry_version": acctest.Representation{RepType: acctest.Optional, Create: `10`, Update: `11`},
		"time_created":     acctest.Representation{RepType: acctest.Optional, Create: `timeCreated`, Update: `timeCreated2`},
		"time_updated":     acctest.Representation{RepType: acctest.Optional, Create: `timeUpdated`, Update: `timeUpdated2`},
		"updated_by":       acctest.Representation{RepType: acctest.Optional, Create: `updatedBy`, Update: `updatedBy2`},
		"updated_by_name":  acctest.Representation{RepType: acctest.Optional, Create: `updatedByName`, Update: `updatedByName2`},
	}
	DataConnectivityRegistryConnectionPrimarySchemaParentRefRepresentation = map[string]interface{}{
		"parent": acctest.Representation{RepType: acctest.Optional, Create: `parent`, Update: `parent2`},
	}
	DataConnectivityRegistryConnectionPrimarySchemaMetadataAggregatorRepresentation = map[string]interface{}{
		"description": acctest.Representation{RepType: acctest.Optional, Create: `description`, Update: `description2`},
		"identifier":  acctest.Representation{RepType: acctest.Optional, Create: `identifier`, Update: `identifier2`},
		"key":         acctest.Representation{RepType: acctest.Optional, Create: `key`, Update: `key2`},
		"name":        acctest.Representation{RepType: acctest.Optional, Create: `name`, Update: `name2`},
		"type":        acctest.Representation{RepType: acctest.Optional, Create: `type`, Update: `type2`},
	}

	DataConnectivityRegistryConnectionResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_data_connectivity_registry", "test_registry", acctest.Required, acctest.Create, DataConnectivityRegistryRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_identity_user", "test_user", acctest.Required, acctest.Create, IdentityUserRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_data_connectivity_registry_data_asset", "test_registry_data_asset", acctest.Required, acctest.Create, DataConnectivityRegistryDataAssetRepresentation)
)

// issue-routing-tag: data_connectivity/default
func TestDataConnectivityRegistryConnectionResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDataConnectivityRegistryConnectionResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_data_connectivity_registry_connection.test_registry_connection"
	datasourceName := "data.oci_data_connectivity_registry_connections.test_registry_connections"
	singularDatasourceName := "data.oci_data_connectivity_registry_connection.test_registry_connection"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+DataConnectivityRegistryConnectionResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_data_connectivity_registry_connection", "test_registry_connection", acctest.Optional, acctest.Create, DataConnectivityRegistryConnectionRepresentation), "dataconnectivity", "registryConnection", t)

	acctest.ResourceTest(t, testAccCheckDataConnectivityRegistryConnectionDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + DataConnectivityRegistryConnectionResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_data_connectivity_registry_connection", "test_registry_connection", acctest.Required, acctest.Create, DataConnectivityRegistryConnectionRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "identifier", "CONNECTIONIDENTIFIER"),
				resource.TestCheckResourceAttr(resourceName, "name", "connectionname"),
				resource.TestCheckResourceAttr(resourceName, "properties.%", "3"),
				resource.TestCheckResourceAttrSet(resourceName, "registry_id"),
				resource.TestCheckResourceAttr(resourceName, "type", "GENERIC_JDBC_CONNECTION"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + DataConnectivityRegistryConnectionResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + DataConnectivityRegistryConnectionResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_data_connectivity_registry_connection", "test_registry_connection", acctest.Optional, acctest.Create, DataConnectivityRegistryConnectionRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "connection_properties.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "connection_properties.0.name", "name"),
				resource.TestCheckResourceAttr(resourceName, "connection_properties.0.value", "value"),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "identifier", "CONNECTIONIDENTIFIER"),
				resource.TestCheckResourceAttr(resourceName, "is_default", "false"),
				resource.TestCheckResourceAttr(resourceName, "model_type", "GENERIC_CONNECTION"),
				resource.TestCheckResourceAttr(resourceName, "name", "connectionname"),
				resource.TestCheckResourceAttr(resourceName, "object_status", "8"),
				resource.TestCheckResourceAttr(resourceName, "object_version", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "key"),
				resource.TestCheckResourceAttrSet(resourceName, "registry_id"),

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

		// verify updates to updatable parameters
		{
			Config: config + compartmentIdVariableStr + DataConnectivityRegistryConnectionResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_data_connectivity_registry_connection", "test_registry_connection", acctest.Optional, acctest.Update, DataConnectivityRegistryConnectionRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "connection_properties.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "connection_properties.0.name", "name2"),
				resource.TestCheckResourceAttr(resourceName, "connection_properties.0.value", "value2"),
				resource.TestCheckResourceAttr(resourceName, "description", "description2"),
				resource.TestCheckResourceAttr(resourceName, "identifier", "CONNECTIONIDENTIFIER2"),
				resource.TestCheckResourceAttr(resourceName, "is_default", "false"),
				resource.TestCheckResourceAttr(resourceName, "name", "connectionname2"),
				resource.TestCheckResourceAttrSet(resourceName, "key"),
				resource.TestCheckResourceAttrSet(resourceName, "registry_id"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId2 != resId2 {
						return fmt.Errorf("Resource recreated when it was supposed to be updated.")
					}
					return err
				},
			),
		},
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_data_connectivity_registry_connections", "test_registry_connections", acctest.Optional, acctest.Update, DataConnectivityDataConnectivityRegistryConnectionDataSourceRepresentation) +
				compartmentIdVariableStr + DataConnectivityRegistryConnectionResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_data_connectivity_registry_connection", "test_registry_connection", acctest.Optional, acctest.Update, DataConnectivityRegistryConnectionRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "connection_summary_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "connection_summary_collection.0.items.#", "2"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_data_connectivity_registry_connection", "test_registry_connection", acctest.Required, acctest.Create, DataConnectivityDataConnectivityRegistryConnectionSingularDataSourceRepresentation) +
				compartmentIdVariableStr + DataConnectivityRegistryConnectionResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "connection_key"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "registry_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "connection_properties.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "connection_properties.0.name", "name2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "connection_properties.0.value", "value2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "description", "description2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "identifier", "CONNECTIONIDENTIFIER2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "is_default", "false"),
			),
		},
		// verify resource import
		{
			Config:                  config + DataConnectivityRegistryConnectionRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{"registry_metadata"},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckDataConnectivityRegistryConnectionDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).DataConnectivityManagementClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_data_connectivity_registry_connection" {
			noResourceFound = false
			request := oci_data_connectivity.GetConnectionRequest{}

			if value, ok := rs.Primary.Attributes["key"]; ok {
				request.ConnectionKey = &value
			}

			if value, ok := rs.Primary.Attributes["registry_id"]; ok {
				request.RegistryId = &value
			}

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "data_connectivity")

			_, err := client.GetConnection(context.Background(), request)

			if err == nil {
				return fmt.Errorf("resource still exists")
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
	if !acctest.InSweeperExcludeList("DataConnectivityRegistryConnection") {
		resource.AddTestSweepers("DataConnectivityRegistryConnection", &resource.Sweeper{
			Name:         "DataConnectivityRegistryConnection",
			Dependencies: acctest.DependencyGraph["registryConnection"],
			F:            sweepDataConnectivityRegistryConnectionResource,
		})
	}
}

func sweepDataConnectivityRegistryConnectionResource(compartment string) error {
	dataConnectivityManagementClient := acctest.GetTestClients(&schema.ResourceData{}).DataConnectivityManagementClient()
	registryConnectionIds, err := getDataConnectivityRegistryConnectionIds(compartment)
	if err != nil {
		return err
	}
	for _, registryConnectionId := range registryConnectionIds {
		if ok := acctest.SweeperDefaultResourceId[registryConnectionId]; !ok {
			deleteConnectionRequest := oci_data_connectivity.DeleteConnectionRequest{}

			deleteConnectionRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "data_connectivity")
			_, error := dataConnectivityManagementClient.DeleteConnection(context.Background(), deleteConnectionRequest)
			if error != nil {
				fmt.Printf("Error deleting RegistryConnection %s %s, It is possible that the resource is already deleted. Please verify manually \n", registryConnectionId, error)
				continue
			}
		}
	}
	return nil
}

func getDataConnectivityRegistryConnectionIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "RegistryConnectionId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	dataConnectivityManagementClient := acctest.GetTestClients(&schema.ResourceData{}).DataConnectivityManagementClient()

	listConnectionsRequest := oci_data_connectivity.ListConnectionsRequest{}

	dataAssetKeys, error := getDataConnectivityRegistryDataAssetIds(compartment)
	if error != nil {
		return resourceIds, fmt.Errorf("Error getting dataAssetKey required for RegistryConnection resource requests \n")
	}
	for _, dataAssetKey := range dataAssetKeys {
		listConnectionsRequest.DataAssetKey = &dataAssetKey

		registryIds, error := getDataConnectivityRegistryIds(compartment)
		if error != nil {
			return resourceIds, fmt.Errorf("Error getting registryId required for RegistryConnection resource requests \n")
		}
		for _, registryId := range registryIds {
			listConnectionsRequest.RegistryId = &registryId

			listConnectionsResponse, err := dataConnectivityManagementClient.ListConnections(context.Background(), listConnectionsRequest)

			if err != nil {
				return resourceIds, fmt.Errorf("Error getting RegistryConnection list for compartment id : %s , %s \n", compartmentId, err)
			}
			for _, registryConnection := range listConnectionsResponse.Items {
				id := *registryConnection.Key
				resourceIds = append(resourceIds, id)
				acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "RegistryConnectionId", id)
			}

		}
	}
	return resourceIds, nil
}
