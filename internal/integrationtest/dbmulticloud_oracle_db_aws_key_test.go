// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
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
	DbmulticloudOracleDbAwsKeyRequiredOnlyResource = DbmulticloudOracleDbAwsKeyResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_dbmulticloud_oracle_db_aws_key", "test_oracle_db_aws_key", acctest.Required, acctest.Create, DbmulticloudOracleDbAwsKeyRepresentation)

	DbmulticloudOracleDbAwsKeyResourceConfig = DbmulticloudOracleDbAwsKeyResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_dbmulticloud_oracle_db_aws_key", "test_oracle_db_aws_key", acctest.Optional, acctest.Update, DbmulticloudOracleDbAwsKeyRepresentation)

	DbmulticloudOracleDbAwsKeySingularDataSourceRepresentation = map[string]interface{}{
		"oracle_db_aws_key_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_dbmulticloud_oracle_db_aws_key.test_oracle_db_aws_key.id}`},
	}

	DbmulticloudOracleDbAwsKeyDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":   acctest.Representation{RepType: acctest.Optional, Create: `AWS_Key_Tersi_Test`, Update: `AWS_Key_Tersi_Test`},
		"filter":         acctest.RepresentationGroup{RepType: acctest.Required, Group: DbmulticloudOracleDbAwsKeyDataSourceFilterRepresentation}}
	DbmulticloudOracleDbAwsKeyDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_dbmulticloud_oracle_db_aws_key.test_oracle_db_aws_key.id}`}},
	}

	DbmulticloudOracleDbAwsKeyRepresentation = map[string]interface{}{
		"compartment_id":         acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":           acctest.Representation{RepType: acctest.Required, Create: `AWS_Key_Tersi_Test`, Update: `AWS_Key_Tersi_Test`},
		"oracle_db_connector_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_dbmulticloud_oracle_db_aws_identity_connector.test_oracle_db_aws_identity_connector.id}`},
		"aws_account_id":         acctest.Representation{RepType: acctest.Optional, Create: `867344470629`},
		"aws_key_arn":            acctest.Representation{RepType: acctest.Required, Create: `arn:aws:iam::867344470629:role/OracleDatabaseKMS`},
		"freeform_tags":          acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}},
		"is_aws_key_enabled":     acctest.Representation{RepType: acctest.Required, Create: `false`},
	}

	DbmulticloudOracleDbAwsKeyResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_dbmulticloud_oracle_db_aws_identity_connector", "test_oracle_db_aws_identity_connector", acctest.Required, acctest.Create, DbmulticloudOracleDbAwsIdentityConnectorRepresentation) +
		DefinedTagsDependencies
)

// issue-routing-tag: dbmulticloud/default
func TestDbmulticloudOracleDbAwsKeyResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDbmulticloudOracleDbAwsKeyResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_dbmulticloud_oracle_db_aws_key.test_oracle_db_aws_key"
	datasourceName := "data.oci_dbmulticloud_oracle_db_aws_keys.test_oracle_db_aws_keys"
	singularDatasourceName := "data.oci_dbmulticloud_oracle_db_aws_key.test_oracle_db_aws_key"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+DbmulticloudOracleDbAwsKeyResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_dbmulticloud_oracle_db_aws_key", "test_oracle_db_aws_key", acctest.Optional, acctest.Create, DbmulticloudOracleDbAwsKeyRepresentation), "dbmulticloud", "oracleDbAwsKey", t)

	acctest.ResourceTest(t, testAccCheckDbmulticloudOracleDbAwsKeyDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + DbmulticloudOracleDbAwsKeyResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_dbmulticloud_oracle_db_aws_key", "test_oracle_db_aws_key", acctest.Required, acctest.Create, DbmulticloudOracleDbAwsKeyRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "AWS_Key_Tersi_Test"),
				resource.TestCheckResourceAttrSet(resourceName, "oracle_db_connector_id"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + DbmulticloudOracleDbAwsKeyResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + DbmulticloudOracleDbAwsKeyResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_dbmulticloud_oracle_db_aws_key", "test_oracle_db_aws_key", acctest.Optional, acctest.Create, DbmulticloudOracleDbAwsKeyRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "aws_account_id"),
				resource.TestCheckResourceAttr(resourceName, "aws_key_arn", "arn:aws:iam::867344470629:role/OracleDatabaseKMS"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "AWS_Key_Tersi_Test"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_aws_key_enabled", "false"),
				resource.TestCheckResourceAttrSet(resourceName, "oracle_db_connector_id"),

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
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + DbmulticloudOracleDbAwsKeyResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_dbmulticloud_oracle_db_aws_key", "test_oracle_db_aws_key", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(DbmulticloudOracleDbAwsKeyRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "aws_account_id"),
				resource.TestCheckResourceAttr(resourceName, "aws_key_arn", "arn:aws:iam::867344470629:role/OracleDatabaseKMS"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "display_name", "AWS_Key_Tersi_Test"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_aws_key_enabled", "false"),
				resource.TestCheckResourceAttrSet(resourceName, "oracle_db_connector_id"),

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
			Config: config + compartmentIdVariableStr + DbmulticloudOracleDbAwsKeyResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_dbmulticloud_oracle_db_aws_key", "test_oracle_db_aws_key", acctest.Optional, acctest.Update, DbmulticloudOracleDbAwsKeyRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "aws_account_id"),
				resource.TestCheckResourceAttr(resourceName, "aws_key_arn", "arn:aws:iam::867344470629:role/OracleDatabaseKMS"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "AWS_Key_Tersi_Test"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_aws_key_enabled", "false"),
				resource.TestCheckResourceAttrSet(resourceName, "oracle_db_connector_id"),

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
				acctest.GenerateDataSourceFromRepresentationMap("oci_dbmulticloud_oracle_db_aws_keys", "test_oracle_db_aws_keys", acctest.Optional, acctest.Update, DbmulticloudOracleDbAwsKeyDataSourceRepresentation) +
				compartmentIdVariableStr + DbmulticloudOracleDbAwsKeyResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_dbmulticloud_oracle_db_aws_key", "test_oracle_db_aws_key", acctest.Optional, acctest.Update, DbmulticloudOracleDbAwsKeyRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "AWS_Key_Tersi_Test"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_dbmulticloud_oracle_db_aws_key", "test_oracle_db_aws_key", acctest.Required, acctest.Create, DbmulticloudOracleDbAwsKeySingularDataSourceRepresentation) +
				compartmentIdVariableStr + DbmulticloudOracleDbAwsKeyResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(singularDatasourceName, "aws_key_arn", "arn:aws:iam::867344470629:role/OracleDatabaseKMS"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "AWS_Key_Tersi_Test"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "is_aws_key_enabled", "false"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
		// verify resource import
		{
			Config:                  config + DbmulticloudOracleDbAwsKeyRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckDbmulticloudOracleDbAwsKeyDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).DbMulticloudAwsProviderClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_dbmulticloud_oracle_db_aws_key" {
			noResourceFound = false
			request := oci_dbmulticloud.GetOracleDbAwsKeyRequest{}

			tmp := rs.Primary.ID
			request.OracleDbAwsKeyId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "dbmulticloud")

			response, err := client.GetOracleDbAwsKey(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_dbmulticloud.OracleDbAwsKeyLifecycleStateDeleted): true,
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
	if !acctest.InSweeperExcludeList("DbmulticloudOracleDbAwsKey") {
		resource.AddTestSweepers("DbmulticloudOracleDbAwsKey", &resource.Sweeper{
			Name:         "DbmulticloudOracleDbAwsKey",
			Dependencies: acctest.DependencyGraph["oracleDbAwsKey"],
			F:            sweepDbmulticloudOracleDbAwsKeyResource,
		})
	}
}

func sweepDbmulticloudOracleDbAwsKeyResource(compartment string) error {
	dbMulticloudAwsProviderClient := acctest.GetTestClients(&schema.ResourceData{}).DbMulticloudAwsProviderClient()
	oracleDbAwsKeyIds, err := getDbmulticloudOracleDbAwsKeyIds(compartment)
	if err != nil {
		return err
	}
	for _, oracleDbAwsKeyId := range oracleDbAwsKeyIds {
		if ok := acctest.SweeperDefaultResourceId[oracleDbAwsKeyId]; !ok {
			deleteOracleDbAwsKeyRequest := oci_dbmulticloud.DeleteOracleDbAwsKeyRequest{}

			deleteOracleDbAwsKeyRequest.OracleDbAwsKeyId = &oracleDbAwsKeyId

			deleteOracleDbAwsKeyRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "dbmulticloud")
			_, error := dbMulticloudAwsProviderClient.DeleteOracleDbAwsKey(context.Background(), deleteOracleDbAwsKeyRequest)
			if error != nil {
				fmt.Printf("Error deleting OracleDbAwsKey %s %s, It is possible that the resource is already deleted. Please verify manually \n", oracleDbAwsKeyId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &oracleDbAwsKeyId, DbmulticloudOracleDbAwsKeySweepWaitCondition, time.Duration(3*time.Minute),
				DbmulticloudOracleDbAwsKeySweepResponseFetchOperation, "dbmulticloud", true)
		}
	}
	return nil
}

func getDbmulticloudOracleDbAwsKeyIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "OracleDbAwsKeyId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	dbMulticloudAwsProviderClient := acctest.GetTestClients(&schema.ResourceData{}).DbMulticloudAwsProviderClient()

	listOracleDbAwsKeysRequest := oci_dbmulticloud.ListOracleDbAwsKeysRequest{}
	listOracleDbAwsKeysRequest.CompartmentId = &compartmentId
	listOracleDbAwsKeysRequest.LifecycleState = oci_dbmulticloud.OracleDbAwsKeyLifecycleStateActive
	listOracleDbAwsKeysResponse, err := dbMulticloudAwsProviderClient.ListOracleDbAwsKeys(context.Background(), listOracleDbAwsKeysRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting OracleDbAwsKey list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, oracleDbAwsKey := range listOracleDbAwsKeysResponse.Items {
		id := *oracleDbAwsKey.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "OracleDbAwsKeyId", id)
	}
	return resourceIds, nil
}

func DbmulticloudOracleDbAwsKeySweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if oracleDbAwsKeyResponse, ok := response.Response.(oci_dbmulticloud.GetOracleDbAwsKeyResponse); ok {
		return oracleDbAwsKeyResponse.LifecycleState != oci_dbmulticloud.OracleDbAwsKeyLifecycleStateDeleted
	}
	return false
}

func DbmulticloudOracleDbAwsKeySweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.DbMulticloudAwsProviderClient().GetOracleDbAwsKey(context.Background(), oci_dbmulticloud.GetOracleDbAwsKeyRequest{
		OracleDbAwsKeyId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
