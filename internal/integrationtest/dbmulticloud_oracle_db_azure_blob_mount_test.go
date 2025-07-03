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
	DbmulticloudOracleDbAzureBlobMountRequiredOnlyResource = DbmulticloudOracleDbAzureBlobMountResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_dbmulticloud_oracle_db_azure_blob_mount", "test_oracle_db_azure_blob_mount", acctest.Required, acctest.Create, DbmulticloudOracleDbAzureBlobMountRepresentation)

	DbmulticloudOracleDbAzureBlobMountResourceConfig = DbmulticloudOracleDbAzureBlobMountResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_dbmulticloud_oracle_db_azure_blob_mount", "test_oracle_db_azure_blob_mount", acctest.Optional, acctest.Update, DbmulticloudOracleDbAzureBlobMountRepresentation)

	DbmulticloudOracleDbAzureBlobMountSingularDataSourceRepresentation = map[string]interface{}{
		"oracle_db_azure_blob_mount_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_dbmulticloud_oracle_db_azure_blob_mount.test_oracle_db_azure_blob_mount.id}`},
	}

	DbmulticloudOracleDbAzureBlobMountDataSourceRepresentation = map[string]interface{}{
		"compartment_id":                    acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":                      acctest.Representation{RepType: acctest.Required, Create: `TestOracleDBAzureBlobMountUpdated`, Update: `TestOracleDBAzureBlobMountUpdated`},
		"oracle_db_azure_blob_container_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_dbmulticloud_oracle_db_azure_blob_container.test_oracle_db_azure_blob_container.id}`},
		// "oracle_db_azure_blob_mount_id":     acctest.Representation{RepType: acctest.Optional, Create: `${oci_dbmulticloud_oracle_db_azure_blob_mount.test_oracle_db_azure_blob_mount.id}`},
		"oracle_db_azure_connector_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_dbmulticloud_oracle_db_azure_connector.test_oracle_db_azure_connector.id}`},
		// "state":                        acctest.Representation{RepType: acctest.Optional, Create: `AVAILABLE`},
		"filter": acctest.RepresentationGroup{RepType: acctest.Required, Group: DbmulticloudOracleDbAzureBlobMountDataSourceFilterRepresentation}}
	DbmulticloudOracleDbAzureBlobMountDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_dbmulticloud_oracle_db_azure_blob_mount.test_oracle_db_azure_blob_mount.id}`}},
	}

	DbmulticloudOracleDbAzureBlobMountRepresentation = map[string]interface{}{
		"compartment_id":                    acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":                      acctest.Representation{RepType: acctest.Required, Create: `TestOracleDBAzureBlobMountUpdated`, Update: `TestOracleDBAzureBlobMountUpdated`},
		"oracle_db_azure_blob_container_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_dbmulticloud_oracle_db_azure_blob_container.test_oracle_db_azure_blob_container.id}`},
		"oracle_db_azure_connector_id":      acctest.Representation{RepType: acctest.Required, Create: `${oci_dbmulticloud_oracle_db_azure_connector.test_oracle_db_azure_connector.id}`},
		// "defined_tags":                      acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		// "freeform_tags":                     acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
	}

	DbmulticloudOracleDbAzureBlobMountResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_dbmulticloud_oracle_db_azure_connector", "test_oracle_db_azure_connector", acctest.Required, acctest.Create, DbmulticloudOracleDbAzureConnectorRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_dbmulticloud_oracle_db_azure_blob_container", "test_oracle_db_azure_blob_container", acctest.Required, acctest.Create, DbmulticloudOracleDbAzureBlobContainerRepresentation) + DefinedTagsDependencies
)

// issue-routing-tag: dbmulticloud/default
func TestDbmulticloudOracleDbAzureBlobMountResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDbmulticloudOracleDbAzureBlobMountResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_dbmulticloud_oracle_db_azure_blob_mount.test_oracle_db_azure_blob_mount"
	datasourceName := "data.oci_dbmulticloud_oracle_db_azure_blob_mounts.test_oracle_db_azure_blob_mounts"
	singularDatasourceName := "data.oci_dbmulticloud_oracle_db_azure_blob_mount.test_oracle_db_azure_blob_mount"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+DbmulticloudOracleDbAzureBlobMountResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_dbmulticloud_oracle_db_azure_blob_mount", "test_oracle_db_azure_blob_mount", acctest.Optional, acctest.Create, DbmulticloudOracleDbAzureBlobMountRepresentation), "dbmulticloud", "oracleDbAzureBlobMount", t)

	acctest.ResourceTest(t, testAccCheckDbmulticloudOracleDbAzureBlobMountDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + DbmulticloudOracleDbAzureBlobMountResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_dbmulticloud_oracle_db_azure_blob_mount", "test_oracle_db_azure_blob_mount", acctest.Required, acctest.Create, DbmulticloudOracleDbAzureBlobMountRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "TestOracleDBAzureBlobMountUpdated"),
				resource.TestCheckResourceAttrSet(resourceName, "oracle_db_azure_blob_container_id"),
				resource.TestCheckResourceAttrSet(resourceName, "oracle_db_azure_connector_id"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + DbmulticloudOracleDbAzureBlobMountResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + DbmulticloudOracleDbAzureBlobMountResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_dbmulticloud_oracle_db_azure_blob_mount", "test_oracle_db_azure_blob_mount", acctest.Optional, acctest.Create, DbmulticloudOracleDbAzureBlobMountRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "TestOracleDBAzureBlobMountUpdated"),
				// resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "oracle_db_azure_blob_container_id"),
				resource.TestCheckResourceAttrSet(resourceName, "oracle_db_azure_connector_id"),

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
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + DbmulticloudOracleDbAzureBlobMountResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_dbmulticloud_oracle_db_azure_blob_mount", "test_oracle_db_azure_blob_mount", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(DbmulticloudOracleDbAzureBlobMountRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "display_name", "TestOracleDBAzureBlobMountUpdated"),
				// resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "oracle_db_azure_blob_container_id"),
				resource.TestCheckResourceAttrSet(resourceName, "oracle_db_azure_connector_id"),

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
			Config: config + compartmentIdVariableStr + DbmulticloudOracleDbAzureBlobMountResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_dbmulticloud_oracle_db_azure_blob_mount", "test_oracle_db_azure_blob_mount", acctest.Optional, acctest.Update, DbmulticloudOracleDbAzureBlobMountRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "TestOracleDBAzureBlobMountUpdated"),
				// resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "oracle_db_azure_blob_container_id"),
				resource.TestCheckResourceAttrSet(resourceName, "oracle_db_azure_connector_id"),

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
				acctest.GenerateDataSourceFromRepresentationMap("oci_dbmulticloud_oracle_db_azure_blob_mounts", "test_oracle_db_azure_blob_mounts", acctest.Optional, acctest.Update, DbmulticloudOracleDbAzureBlobMountDataSourceRepresentation) +
				compartmentIdVariableStr + DbmulticloudOracleDbAzureBlobMountResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_dbmulticloud_oracle_db_azure_blob_mount", "test_oracle_db_azure_blob_mount", acctest.Optional, acctest.Update, DbmulticloudOracleDbAzureBlobMountRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "TestOracleDBAzureBlobMountUpdated"),
				resource.TestCheckResourceAttrSet(datasourceName, "oracle_db_azure_blob_container_id"),
				// resource.TestCheckResourceAttrSet(datasourceName, "oracle_db_azure_blob_mount_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "oracle_db_azure_connector_id"),
				// resource.TestCheckResourceAttr(datasourceName, "state", "AVAILABLE"),

				resource.TestCheckResourceAttr(datasourceName, "oracle_db_azure_blob_mount_summary_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "oracle_db_azure_blob_mount_summary_collection.0.items.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_dbmulticloud_oracle_db_azure_blob_mount", "test_oracle_db_azure_blob_mount", acctest.Required, acctest.Create, DbmulticloudOracleDbAzureBlobMountSingularDataSourceRepresentation) +
				compartmentIdVariableStr + DbmulticloudOracleDbAzureBlobMountResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				// resource.TestCheckResourceAttrSet(singularDatasourceName, "oracle_db_azure_blob_mount_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "TestOracleDBAzureBlobMountUpdated"),
				// resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "last_modification"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "mount_path"),
				// resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
		// verify resource import
		{
			Config:                  config + DbmulticloudOracleDbAzureBlobMountRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckDbmulticloudOracleDbAzureBlobMountDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).OracleDBAzureBlobMountClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_dbmulticloud_oracle_db_azure_blob_mount" {
			noResourceFound = false
			request := oci_dbmulticloud.GetOracleDbAzureBlobMountRequest{}

			tmp := rs.Primary.ID
			request.OracleDbAzureBlobMountId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "dbmulticloud")

			response, err := client.GetOracleDbAzureBlobMount(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_dbmulticloud.OracleDbAzureBlobMountLifecycleStateDeleted): true,
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
	if !acctest.InSweeperExcludeList("DbmulticloudOracleDbAzureBlobMount") {
		resource.AddTestSweepers("DbmulticloudOracleDbAzureBlobMount", &resource.Sweeper{
			Name:         "DbmulticloudOracleDbAzureBlobMount",
			Dependencies: acctest.DependencyGraph["oracleDbAzureBlobMount"],
			F:            sweepDbmulticloudOracleDbAzureBlobMountResource,
		})
	}
}

func sweepDbmulticloudOracleDbAzureBlobMountResource(compartment string) error {
	oracleDBAzureBlobMountClient := acctest.GetTestClients(&schema.ResourceData{}).OracleDBAzureBlobMountClient()
	oracleDbAzureBlobMountIds, err := getDbmulticloudOracleDbAzureBlobMountIds(compartment)
	if err != nil {
		return err
	}
	for _, oracleDbAzureBlobMountId := range oracleDbAzureBlobMountIds {
		if ok := acctest.SweeperDefaultResourceId[oracleDbAzureBlobMountId]; !ok {
			deleteOracleDbAzureBlobMountRequest := oci_dbmulticloud.DeleteOracleDbAzureBlobMountRequest{}

			deleteOracleDbAzureBlobMountRequest.OracleDbAzureBlobMountId = &oracleDbAzureBlobMountId

			deleteOracleDbAzureBlobMountRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "dbmulticloud")
			_, error := oracleDBAzureBlobMountClient.DeleteOracleDbAzureBlobMount(context.Background(), deleteOracleDbAzureBlobMountRequest)
			if error != nil {
				fmt.Printf("Error deleting OracleDbAzureBlobMount %s %s, It is possible that the resource is already deleted. Please verify manually \n", oracleDbAzureBlobMountId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &oracleDbAzureBlobMountId, DbmulticloudOracleDbAzureBlobMountSweepWaitCondition, time.Duration(3*time.Minute),
				DbmulticloudOracleDbAzureBlobMountSweepResponseFetchOperation, "dbmulticloud", true)
		}
	}
	return nil
}

func getDbmulticloudOracleDbAzureBlobMountIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "OracleDbAzureBlobMountId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	oracleDBAzureBlobMountClient := acctest.GetTestClients(&schema.ResourceData{}).OracleDBAzureBlobMountClient()

	listOracleDbAzureBlobMountsRequest := oci_dbmulticloud.ListOracleDbAzureBlobMountsRequest{}
	listOracleDbAzureBlobMountsRequest.CompartmentId = &compartmentId
	listOracleDbAzureBlobMountsRequest.LifecycleState = oci_dbmulticloud.OracleDbAzureBlobMountLifecycleStateActive
	listOracleDbAzureBlobMountsResponse, err := oracleDBAzureBlobMountClient.ListOracleDbAzureBlobMounts(context.Background(), listOracleDbAzureBlobMountsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting OracleDbAzureBlobMount list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, oracleDbAzureBlobMount := range listOracleDbAzureBlobMountsResponse.Items {
		id := *oracleDbAzureBlobMount.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "OracleDbAzureBlobMountId", id)
	}
	return resourceIds, nil
}

func DbmulticloudOracleDbAzureBlobMountSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if oracleDbAzureBlobMountResponse, ok := response.Response.(oci_dbmulticloud.GetOracleDbAzureBlobMountResponse); ok {
		return oracleDbAzureBlobMountResponse.LifecycleState != oci_dbmulticloud.OracleDbAzureBlobMountLifecycleStateDeleted
	}
	return false
}

func DbmulticloudOracleDbAzureBlobMountSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.OracleDBAzureBlobMountClient().GetOracleDbAzureBlobMount(context.Background(), oci_dbmulticloud.GetOracleDbAzureBlobMountRequest{
		OracleDbAzureBlobMountId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
