// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"

	// "strconv"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/oracle/oci-go-sdk/v65/common"
	oci_dbmulticloud "github.com/oracle/oci-go-sdk/v65/dbmulticloud"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"

	// "github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	DbmulticloudOracleDbAzureBlobContainerRequiredOnlyResource = DbmulticloudOracleDbAzureBlobContainerResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_dbmulticloud_oracle_db_azure_blob_container", "test_oracle_db_azure_blob_container", acctest.Required, acctest.Create, DbmulticloudOracleDbAzureBlobContainerRepresentation)

	DbmulticloudOracleDbAzureBlobContainerResourceConfig = DbmulticloudOracleDbAzureBlobContainerResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_dbmulticloud_oracle_db_azure_blob_container", "test_oracle_db_azure_blob_container", acctest.Optional, acctest.Update, DbmulticloudOracleDbAzureBlobContainerRepresentation)

	DbmulticloudOracleDbAzureBlobContainerSingularDataSourceRepresentation = map[string]interface{}{
		"oracle_db_azure_blob_container_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_dbmulticloud_oracle_db_azure_blob_container.test_oracle_db_azure_blob_container.id}`},
	}

	DbmulticloudOracleDbAzureBlobContainerDataSourceRepresentation = map[string]interface{}{
		"compartment_id":               acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"azure_storage_account_name":   acctest.Representation{RepType: acctest.Optional, Create: `ADBDAppStorageAccounts`, Update: `ADBDAppStorageAccounts`},
		"azure_storage_container_name": acctest.Representation{RepType: acctest.Optional, Create: `ADBDContainers`, Update: `ADBDContainers`},
		"display_name":                 acctest.Representation{RepType: acctest.Optional, Create: `TestDBAzureBlobContainerUpdate`, Update: `displayName2`},
		// "oracle_db_azure_blob_container_id": acctest.Representation{RepType: acctest.Optional, Create: `${oci_dbmulticloud_oracle_db_azure_blob_container.test_oracle_db_azure_blob_container.id}`},
		"state":  acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"filter": acctest.RepresentationGroup{RepType: acctest.Required, Group: DbmulticloudOracleDbAzureBlobContainerDataSourceFilterRepresentation}}
	DbmulticloudOracleDbAzureBlobContainerDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_dbmulticloud_oracle_db_azure_blob_container.test_oracle_db_azure_blob_container.id}`}},
	}

	DbmulticloudOracleDbAzureBlobContainerRepresentation = map[string]interface{}{
		"azure_storage_account_name":   acctest.Representation{RepType: acctest.Required, Create: `ADBDAppStorageAccounts`, Update: `ADBDAppStorageAccounts`},
		"azure_storage_container_name": acctest.Representation{RepType: acctest.Required, Create: `ADBDContainers`, Update: `ADBDContainers`},
		"compartment_id":               acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":                 acctest.Representation{RepType: acctest.Required, Create: `TestDBAzureBlobContainerUpdate`, Update: `displayName2`},
		// "defined_tags":                 acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		// "freeform_tags":                acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		// "private_endpoint_dns_alias":   acctest.Representation{RepType: acctest.Optional, Create: `test.oracle.com`, Update: `privateEndpointDnsAlias2`},
		// "private_endpoint_ip_address":  acctest.Representation{RepType: acctest.Optional, Create: `196.168.0.1`, Update: `privateEndpointIpAddress2`},
	}

	DbmulticloudOracleDbAzureBlobContainerResourceDependencies = ""
	//  acctest.GenerateResourceFromRepresentationMap("oci_dbmulticloud_oracle_db_azure_blob_container", "test_oracle_db_azure_blob_container", acctest.Required, acctest.Create, DbmulticloudOracleDbAzureBlobContainerRepresentation) +
	// DefinedTagsDependencies
)

// issue-routing-tag: dbmulticloud/default
func TestDbmulticloudOracleDbAzureBlobContainerResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDbmulticloudOracleDbAzureBlobContainerResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_dbmulticloud_oracle_db_azure_blob_container.test_oracle_db_azure_blob_container"
	datasourceName := "data.oci_dbmulticloud_oracle_db_azure_blob_containers.test_oracle_db_azure_blob_containers"
	singularDatasourceName := "data.oci_dbmulticloud_oracle_db_azure_blob_container.test_oracle_db_azure_blob_container"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+DbmulticloudOracleDbAzureBlobContainerResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_dbmulticloud_oracle_db_azure_blob_container", "test_oracle_db_azure_blob_container", acctest.Optional, acctest.Create, DbmulticloudOracleDbAzureBlobContainerRepresentation), "dbmulticloud", "oracleDbAzureBlobContainer", t)

	acctest.ResourceTest(t, testAccCheckDbmulticloudOracleDbAzureBlobContainerDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + DbmulticloudOracleDbAzureBlobContainerResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_dbmulticloud_oracle_db_azure_blob_container", "test_oracle_db_azure_blob_container", acctest.Required, acctest.Create, DbmulticloudOracleDbAzureBlobContainerRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "azure_storage_account_name", "ADBDAppStorageAccounts"),
				resource.TestCheckResourceAttr(resourceName, "azure_storage_container_name", "ADBDContainers"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "TestDBAzureBlobContainerUpdate"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + DbmulticloudOracleDbAzureBlobContainerResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + DbmulticloudOracleDbAzureBlobContainerResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_dbmulticloud_oracle_db_azure_blob_container", "test_oracle_db_azure_blob_container", acctest.Optional, acctest.Create, DbmulticloudOracleDbAzureBlobContainerRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "azure_storage_account_name", "ADBDAppStorageAccounts"),
				resource.TestCheckResourceAttr(resourceName, "azure_storage_container_name", "ADBDContainers"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "TestDBAzureBlobContainerUpdate"),
				// resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				// resource.TestCheckResourceAttr(resourceName, "private_endpoint_dns_alias", "test.oracle.com"),
				// resource.TestCheckResourceAttr(resourceName, "private_endpoint_ip_address", "196.168.0.1"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					// if isEnableExportCompartment, _ := strconv.ParseBool(utils.GetEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
					// 	if errExport := resourcediscovery.TestExportCompartmentWithResourceName(&resId, &compartmentId, resourceName); errExport != nil {
					// 		return errExport
					// 	}
					// }
					return err
				},
			),
		},

		// verify Update to the compartment (the compartment will be switched back in the next step)
		{
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + DbmulticloudOracleDbAzureBlobContainerResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_dbmulticloud_oracle_db_azure_blob_container", "test_oracle_db_azure_blob_container", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(DbmulticloudOracleDbAzureBlobContainerRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "azure_storage_account_name", "ADBDAppStorageAccounts"),
				resource.TestCheckResourceAttr(resourceName, "azure_storage_container_name", "ADBDContainers"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "display_name", "TestDBAzureBlobContainerUpdate"),
				// resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				// resource.TestCheckResourceAttr(resourceName, "private_endpoint_dns_alias", "test.oracle.com"),
				// resource.TestCheckResourceAttr(resourceName, "private_endpoint_ip_address", "196.168.0.1"),

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
			Config: config + compartmentIdVariableStr + DbmulticloudOracleDbAzureBlobContainerResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_dbmulticloud_oracle_db_azure_blob_container", "test_oracle_db_azure_blob_container", acctest.Optional, acctest.Update, DbmulticloudOracleDbAzureBlobContainerRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "azure_storage_account_name", "ADBDAppStorageAccounts"),
				resource.TestCheckResourceAttr(resourceName, "azure_storage_container_name", "ADBDContainers"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				// resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				// resource.TestCheckResourceAttr(resourceName, "private_endpoint_dns_alias", "privateEndpointDnsAlias2"),
				// resource.TestCheckResourceAttr(resourceName, "private_endpoint_ip_address", "privateEndpointIpAddress2"),

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
				acctest.GenerateDataSourceFromRepresentationMap("oci_dbmulticloud_oracle_db_azure_blob_containers", "test_oracle_db_azure_blob_containers", acctest.Optional, acctest.Update, DbmulticloudOracleDbAzureBlobContainerDataSourceRepresentation) +
				compartmentIdVariableStr + DbmulticloudOracleDbAzureBlobContainerResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_dbmulticloud_oracle_db_azure_blob_container", "test_oracle_db_azure_blob_container", acctest.Optional, acctest.Update, DbmulticloudOracleDbAzureBlobContainerRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "azure_storage_account_name", "ADBDAppStorageAccounts"),
				resource.TestCheckResourceAttr(datasourceName, "azure_storage_container_name", "ADBDContainers"),
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
				// resource.TestCheckResourceAttrSet(datasourceName, "oracle_db_azure_blob_container_id"),
				resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),

				// resource.TestCheckResourceAttr(datasourceName, "oracle_db_azure_blob_container_summary_collection.#", "1"),
				// resource.TestCheckResourceAttr(datasourceName, "oracle_db_azure_blob_container_summary_collection.0.items.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_dbmulticloud_oracle_db_azure_blob_container", "test_oracle_db_azure_blob_container", acctest.Required, acctest.Create, DbmulticloudOracleDbAzureBlobContainerSingularDataSourceRepresentation) +
				compartmentIdVariableStr + DbmulticloudOracleDbAzureBlobContainerResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				// resource.TestCheckResourceAttrSet(singularDatasourceName, "oracle_db_azure_blob_container_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "azure_storage_account_name", "ADBDAppStorageAccounts"),
				resource.TestCheckResourceAttr(singularDatasourceName, "azure_storage_container_name", "ADBDContainers"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				// resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				// resource.TestCheckResourceAttrSet(singularDatasourceName, "last_modification"),
				// resource.TestCheckResourceAttr(singularDatasourceName, "private_endpoint_dns_alias", "privateEndpointDnsAlias2"),
				// resource.TestCheckResourceAttr(singularDatasourceName, "private_endpoint_ip_address", "privateEndpointIpAddress2"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
		// verify resource import
		{
			Config:                  config + DbmulticloudOracleDbAzureBlobContainerRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckDbmulticloudOracleDbAzureBlobContainerDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).OracleDBAzureBlobContainerClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_dbmulticloud_oracle_db_azure_blob_container" {
			noResourceFound = false
			request := oci_dbmulticloud.GetOracleDbAzureBlobContainerRequest{}

			tmp := rs.Primary.ID
			request.OracleDbAzureBlobContainerId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "dbmulticloud")

			response, err := client.GetOracleDbAzureBlobContainer(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_dbmulticloud.OracleDbAzureBlobContainerLifecycleStateDeleted): true,
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
	if !acctest.InSweeperExcludeList("DbmulticloudOracleDbAzureBlobContainer") {
		resource.AddTestSweepers("DbmulticloudOracleDbAzureBlobContainer", &resource.Sweeper{
			Name:         "DbmulticloudOracleDbAzureBlobContainer",
			Dependencies: acctest.DependencyGraph["oracleDbAzureBlobContainer"],
			F:            sweepDbmulticloudOracleDbAzureBlobContainerResource,
		})
	}
}

func sweepDbmulticloudOracleDbAzureBlobContainerResource(compartment string) error {
	oracleDBAzureBlobContainerClient := acctest.GetTestClients(&schema.ResourceData{}).OracleDBAzureBlobContainerClient()
	oracleDbAzureBlobContainerIds, err := getDbmulticloudOracleDbAzureBlobContainerIds(compartment)
	if err != nil {
		return err
	}
	for _, oracleDbAzureBlobContainerId := range oracleDbAzureBlobContainerIds {
		if ok := acctest.SweeperDefaultResourceId[oracleDbAzureBlobContainerId]; !ok {
			deleteOracleDbAzureBlobContainerRequest := oci_dbmulticloud.DeleteOracleDbAzureBlobContainerRequest{}

			deleteOracleDbAzureBlobContainerRequest.OracleDbAzureBlobContainerId = &oracleDbAzureBlobContainerId

			deleteOracleDbAzureBlobContainerRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "dbmulticloud")
			_, error := oracleDBAzureBlobContainerClient.DeleteOracleDbAzureBlobContainer(context.Background(), deleteOracleDbAzureBlobContainerRequest)
			if error != nil {
				fmt.Printf("Error deleting OracleDbAzureBlobContainer %s %s, It is possible that the resource is already deleted. Please verify manually \n", oracleDbAzureBlobContainerId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &oracleDbAzureBlobContainerId, DbmulticloudOracleDbAzureBlobContainerSweepWaitCondition, time.Duration(3*time.Minute),
				DbmulticloudOracleDbAzureBlobContainerSweepResponseFetchOperation, "dbmulticloud", true)
		}
	}
	return nil
}

func getDbmulticloudOracleDbAzureBlobContainerIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "OracleDbAzureBlobContainerId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	oracleDBAzureBlobContainerClient := acctest.GetTestClients(&schema.ResourceData{}).OracleDBAzureBlobContainerClient()

	listOracleDbAzureBlobContainersRequest := oci_dbmulticloud.ListOracleDbAzureBlobContainersRequest{}
	listOracleDbAzureBlobContainersRequest.CompartmentId = &compartmentId
	listOracleDbAzureBlobContainersRequest.LifecycleState = oci_dbmulticloud.OracleDbAzureBlobContainerLifecycleStateActive
	listOracleDbAzureBlobContainersResponse, err := oracleDBAzureBlobContainerClient.ListOracleDbAzureBlobContainers(context.Background(), listOracleDbAzureBlobContainersRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting OracleDbAzureBlobContainer list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, oracleDbAzureBlobContainer := range listOracleDbAzureBlobContainersResponse.Items {
		id := *oracleDbAzureBlobContainer.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "OracleDbAzureBlobContainerId", id)
	}
	return resourceIds, nil
}

func DbmulticloudOracleDbAzureBlobContainerSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is ACTIVE beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if oracleDbAzureBlobContainerResponse, ok := response.Response.(oci_dbmulticloud.GetOracleDbAzureBlobContainerResponse); ok {
		return oracleDbAzureBlobContainerResponse.LifecycleState != oci_dbmulticloud.OracleDbAzureBlobContainerLifecycleStateDeleted
	}
	return false
}

func DbmulticloudOracleDbAzureBlobContainerSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.OracleDBAzureBlobContainerClient().GetOracleDbAzureBlobContainer(context.Background(), oci_dbmulticloud.GetOracleDbAzureBlobContainerRequest{
		OracleDbAzureBlobContainerId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
