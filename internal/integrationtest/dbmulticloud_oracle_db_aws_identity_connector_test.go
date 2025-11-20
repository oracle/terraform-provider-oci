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
	DbmulticloudOracleDbAwsIdentityConnectorRequiredOnlyResource = DbmulticloudOracleDbAwsIdentityConnectorResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_dbmulticloud_oracle_db_aws_identity_connector", "test_oracle_db_aws_identity_connector", acctest.Required, acctest.Create, DbmulticloudOracleDbAwsIdentityConnectorRepresentation)

	DbmulticloudOracleDbAwsIdentityConnectorResourceConfig = DbmulticloudOracleDbAwsIdentityConnectorResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_dbmulticloud_oracle_db_aws_identity_connector", "test_oracle_db_aws_identity_connector", acctest.Optional, acctest.Update, DbmulticloudOracleDbAwsIdentityConnectorRepresentation)

	DbmulticloudOracleDbAwsIdentityConnectorSingularDataSourceRepresentation = map[string]interface{}{
		"oracle_db_aws_identity_connector_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_dbmulticloud_oracle_db_aws_identity_connector.test_oracle_db_aws_identity_connector.id}`},
	}

	DbmulticloudOracleDbAwsIdentityConnectorDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":   acctest.Representation{RepType: acctest.Optional, Create: `AWS_Tersi_Test`, Update: `AWS_Tersi_Test`},
		"resource_id":    acctest.Representation{RepType: acctest.Required, Create: `ocid1.cloudvmcluster.test..tersitest`},
		"filter":         acctest.RepresentationGroup{RepType: acctest.Required, Group: DbmulticloudOracleDbAwsIdentityConnectorDataSourceFilterRepresentation}}
	DbmulticloudOracleDbAwsIdentityConnectorDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_dbmulticloud_oracle_db_aws_identity_connector.test_oracle_db_aws_identity_connector.id}`}},
	}

	DbmulticloudOracleDbAwsIdentityConnectorRepresentation = map[string]interface{}{
		"aws_location":             acctest.Representation{RepType: acctest.Required, Create: `us-east1`},
		"compartment_id":           acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":             acctest.Representation{RepType: acctest.Required, Create: `AWS_Tersi_Test`, Update: `AWS_Tersi_Test`},
		"issuer_url":               acctest.Representation{RepType: acctest.Required, Create: `https://idcs-0b7b9fa060364ddf849ef39ee8001737.us-ashburn-idcs-1.identity.us-ashburn-1.oci.oraclecloud.com`},
		"oidc_scope":               acctest.Representation{RepType: acctest.Required, Create: `DBMC/aws`},
		"resource_id":              acctest.Representation{RepType: acctest.Required, Create: `ocid1.cloudvmcluster.test..awstersitest`},
		"service_role_details":     acctest.RepresentationGroup{RepType: acctest.Required, Group: DbmulticloudOracleDbAwsIdentityConnectorServiceRoleDetailsRepresentation},
		"aws_account_id":           acctest.Representation{RepType: acctest.Optional, Create: `867344470629`},
		"aws_sts_private_endpoint": acctest.Representation{RepType: acctest.Optional, Create: `https://sts.amazonaws.com`},
		"freeform_tags":            acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}},
	}
	DbmulticloudOracleDbAwsIdentityConnectorServiceRoleDetailsRepresentation = map[string]interface{}{
		"role_arn":                 acctest.Representation{RepType: acctest.Required, Create: `arn:aws:iam::867344470629:role/OracleDatabaseKMS`},
		"service_private_endpoint": acctest.Representation{RepType: acctest.Required, Create: `https://kms.us-east-1.amazonaws.com`},
		"service_type":             acctest.Representation{RepType: acctest.Required, Create: `KMS`},
	}

	DbmulticloudOracleDbAwsIdentityConnectorResourceDependencies = ""
)

// issue-routing-tag: dbmulticloud/default
func TestDbmulticloudOracleDbAwsIdentityConnectorResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDbmulticloudOracleDbAwsIdentityConnectorResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_dbmulticloud_oracle_db_aws_identity_connector.test_oracle_db_aws_identity_connector"
	datasourceName := "data.oci_dbmulticloud_oracle_db_aws_identity_connectors.test_oracle_db_aws_identity_connectors"
	singularDatasourceName := "data.oci_dbmulticloud_oracle_db_aws_identity_connector.test_oracle_db_aws_identity_connector"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+DbmulticloudOracleDbAwsIdentityConnectorResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_dbmulticloud_oracle_db_aws_identity_connector", "test_oracle_db_aws_identity_connector", acctest.Optional, acctest.Create, DbmulticloudOracleDbAwsIdentityConnectorRepresentation), "dbmulticloud", "oracleDbAwsIdentityConnector", t)

	acctest.ResourceTest(t, testAccCheckDbmulticloudOracleDbAwsIdentityConnectorDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + DbmulticloudOracleDbAwsIdentityConnectorResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_dbmulticloud_oracle_db_aws_identity_connector", "test_oracle_db_aws_identity_connector", acctest.Required, acctest.Create, DbmulticloudOracleDbAwsIdentityConnectorRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "aws_location", "us-east1"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "AWS_Tersi_Test"),
				resource.TestCheckResourceAttr(resourceName, "issuer_url", "https://idcs-0b7b9fa060364ddf849ef39ee8001737.us-ashburn-idcs-1.identity.us-ashburn-1.oci.oraclecloud.com"),
				resource.TestCheckResourceAttr(resourceName, "oidc_scope", "DBMC/aws"),
				resource.TestCheckResourceAttrSet(resourceName, "resource_id"),
				resource.TestCheckResourceAttr(resourceName, "service_role_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "service_role_details.0.role_arn", "arn:aws:iam::867344470629:role/OracleDatabaseKMS"),
				resource.TestCheckResourceAttr(resourceName, "service_role_details.0.service_private_endpoint", "https://kms.us-east-1.amazonaws.com"),
				resource.TestCheckResourceAttr(resourceName, "service_role_details.0.service_type", "KMS"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + DbmulticloudOracleDbAwsIdentityConnectorResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + DbmulticloudOracleDbAwsIdentityConnectorResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_dbmulticloud_oracle_db_aws_identity_connector", "test_oracle_db_aws_identity_connector", acctest.Optional, acctest.Create, DbmulticloudOracleDbAwsIdentityConnectorRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "aws_account_id"),
				resource.TestCheckResourceAttr(resourceName, "aws_location", "us-east1"),
				resource.TestCheckResourceAttr(resourceName, "aws_sts_private_endpoint", "https://sts.amazonaws.com"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "AWS_Tersi_Test"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "issuer_url", "https://idcs-0b7b9fa060364ddf849ef39ee8001737.us-ashburn-idcs-1.identity.us-ashburn-1.oci.oraclecloud.com"),
				resource.TestCheckResourceAttr(resourceName, "oidc_scope", "DBMC/aws"),
				resource.TestCheckResourceAttrSet(resourceName, "resource_id"),
				resource.TestCheckResourceAttr(resourceName, "service_role_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "service_role_details.0.role_arn", "arn:aws:iam::867344470629:role/OracleDatabaseKMS"),
				resource.TestCheckResourceAttr(resourceName, "service_role_details.0.service_private_endpoint", "https://kms.us-east-1.amazonaws.com"),
				resource.TestCheckResourceAttr(resourceName, "service_role_details.0.service_type", "KMS"),

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
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + DbmulticloudOracleDbAwsIdentityConnectorResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_dbmulticloud_oracle_db_aws_identity_connector", "test_oracle_db_aws_identity_connector", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(DbmulticloudOracleDbAwsIdentityConnectorRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "aws_account_id"),
				resource.TestCheckResourceAttr(resourceName, "aws_location", "us-east1"),
				resource.TestCheckResourceAttr(resourceName, "aws_sts_private_endpoint", "https://sts.amazonaws.com"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "display_name", "AWS_Tersi_Test"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "issuer_url", "https://idcs-0b7b9fa060364ddf849ef39ee8001737.us-ashburn-idcs-1.identity.us-ashburn-1.oci.oraclecloud.com"),
				resource.TestCheckResourceAttr(resourceName, "oidc_scope", "DBMC/aws"),
				resource.TestCheckResourceAttrSet(resourceName, "resource_id"),
				resource.TestCheckResourceAttr(resourceName, "service_role_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "service_role_details.0.role_arn", "arn:aws:iam::867344470629:role/OracleDatabaseKMS"),
				resource.TestCheckResourceAttr(resourceName, "service_role_details.0.service_private_endpoint", "https://kms.us-east-1.amazonaws.com"),
				resource.TestCheckResourceAttr(resourceName, "service_role_details.0.service_type", "KMS"),

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
			Config: config + compartmentIdVariableStr + DbmulticloudOracleDbAwsIdentityConnectorResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_dbmulticloud_oracle_db_aws_identity_connector", "test_oracle_db_aws_identity_connector", acctest.Optional, acctest.Update, DbmulticloudOracleDbAwsIdentityConnectorRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "aws_account_id"),
				resource.TestCheckResourceAttr(resourceName, "aws_location", "us-east1"),
				resource.TestCheckResourceAttr(resourceName, "aws_sts_private_endpoint", "https://sts.amazonaws.com"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "AWS_Tersi_Test"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "issuer_url", "https://idcs-0b7b9fa060364ddf849ef39ee8001737.us-ashburn-idcs-1.identity.us-ashburn-1.oci.oraclecloud.com"),
				resource.TestCheckResourceAttr(resourceName, "oidc_scope", "DBMC/aws"),
				resource.TestCheckResourceAttrSet(resourceName, "resource_id"),
				resource.TestCheckResourceAttr(resourceName, "service_role_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "service_role_details.0.role_arn", "arn:aws:iam::867344470629:role/OracleDatabaseKMS"),
				resource.TestCheckResourceAttr(resourceName, "service_role_details.0.service_private_endpoint", "https://kms.us-east-1.amazonaws.com"),
				resource.TestCheckResourceAttr(resourceName, "service_role_details.0.service_type", "KMS"),

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
				acctest.GenerateDataSourceFromRepresentationMap("oci_dbmulticloud_oracle_db_aws_identity_connectors", "test_oracle_db_aws_identity_connectors", acctest.Optional, acctest.Update, DbmulticloudOracleDbAwsIdentityConnectorDataSourceRepresentation) +
				compartmentIdVariableStr + DbmulticloudOracleDbAwsIdentityConnectorResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_dbmulticloud_oracle_db_aws_identity_connector", "test_oracle_db_aws_identity_connector", acctest.Optional, acctest.Update, DbmulticloudOracleDbAwsIdentityConnectorRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "AWS_Tersi_Test"),
				resource.TestCheckResourceAttrSet(datasourceName, "resource_id"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_dbmulticloud_oracle_db_aws_identity_connector", "test_oracle_db_aws_identity_connector", acctest.Required, acctest.Create, DbmulticloudOracleDbAwsIdentityConnectorSingularDataSourceRepresentation) +
				compartmentIdVariableStr + DbmulticloudOracleDbAwsIdentityConnectorResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "oracle_db_aws_identity_connector_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "aws_location", "us-east1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "aws_sts_private_endpoint", "https://sts.amazonaws.com"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "AWS_Tersi_Test"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "issuer_url", "https://idcs-0b7b9fa060364ddf849ef39ee8001737.us-ashburn-idcs-1.identity.us-ashburn-1.oci.oraclecloud.com"),
				resource.TestCheckResourceAttr(singularDatasourceName, "oidc_scope", "DBMC/aws"),
				resource.TestCheckResourceAttr(singularDatasourceName, "service_role_details.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "service_role_details.0.assume_role_status"),
				resource.TestCheckResourceAttr(singularDatasourceName, "service_role_details.0.aws_nodes.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "service_role_details.0.role_arn", "arn:aws:iam::867344470629:role/OracleDatabaseKMS"),
				resource.TestCheckResourceAttr(singularDatasourceName, "service_role_details.0.service_private_endpoint", "https://kms.us-east-1.amazonaws.com"),
				resource.TestCheckResourceAttr(singularDatasourceName, "service_role_details.0.service_type", "KMS"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
		// verify resource import
		{
			Config:                  config + DbmulticloudOracleDbAwsIdentityConnectorRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckDbmulticloudOracleDbAwsIdentityConnectorDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).DbMulticloudAwsProviderClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_dbmulticloud_oracle_db_aws_identity_connector" {
			noResourceFound = false
			request := oci_dbmulticloud.GetOracleDbAwsIdentityConnectorRequest{}

			tmp := rs.Primary.ID
			request.OracleDbAwsIdentityConnectorId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "dbmulticloud")

			response, err := client.GetOracleDbAwsIdentityConnector(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_dbmulticloud.OracleDbAwsIdentityConnectorLifecycleStateDeleted): true,
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
	if !acctest.InSweeperExcludeList("DbmulticloudOracleDbAwsIdentityConnector") {
		resource.AddTestSweepers("DbmulticloudOracleDbAwsIdentityConnector", &resource.Sweeper{
			Name:         "DbmulticloudOracleDbAwsIdentityConnector",
			Dependencies: acctest.DependencyGraph["oracleDbAwsIdentityConnector"],
			F:            sweepDbmulticloudOracleDbAwsIdentityConnectorResource,
		})
	}
}

func sweepDbmulticloudOracleDbAwsIdentityConnectorResource(compartment string) error {
	dbMulticloudAwsProviderClient := acctest.GetTestClients(&schema.ResourceData{}).DbMulticloudAwsProviderClient()
	oracleDbAwsIdentityConnectorIds, err := getDbmulticloudOracleDbAwsIdentityConnectorIds(compartment)
	if err != nil {
		return err
	}
	for _, oracleDbAwsIdentityConnectorId := range oracleDbAwsIdentityConnectorIds {
		if ok := acctest.SweeperDefaultResourceId[oracleDbAwsIdentityConnectorId]; !ok {
			deleteOracleDbAwsIdentityConnectorRequest := oci_dbmulticloud.DeleteOracleDbAwsIdentityConnectorRequest{}

			deleteOracleDbAwsIdentityConnectorRequest.OracleDbAwsIdentityConnectorId = &oracleDbAwsIdentityConnectorId

			deleteOracleDbAwsIdentityConnectorRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "dbmulticloud")
			_, error := dbMulticloudAwsProviderClient.DeleteOracleDbAwsIdentityConnector(context.Background(), deleteOracleDbAwsIdentityConnectorRequest)
			if error != nil {
				fmt.Printf("Error deleting OracleDbAwsIdentityConnector %s %s, It is possible that the resource is already deleted. Please verify manually \n", oracleDbAwsIdentityConnectorId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &oracleDbAwsIdentityConnectorId, DbmulticloudOracleDbAwsIdentityConnectorSweepWaitCondition, time.Duration(3*time.Minute),
				DbmulticloudOracleDbAwsIdentityConnectorSweepResponseFetchOperation, "dbmulticloud", true)
		}
	}
	return nil
}

func getDbmulticloudOracleDbAwsIdentityConnectorIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "OracleDbAwsIdentityConnectorId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	dbMulticloudAwsProviderClient := acctest.GetTestClients(&schema.ResourceData{}).DbMulticloudAwsProviderClient()

	listOracleDbAwsIdentityConnectorsRequest := oci_dbmulticloud.ListOracleDbAwsIdentityConnectorsRequest{}
	listOracleDbAwsIdentityConnectorsRequest.CompartmentId = &compartmentId
	listOracleDbAwsIdentityConnectorsRequest.LifecycleState = oci_dbmulticloud.OracleDbAwsIdentityConnectorLifecycleStateActive
	listOracleDbAwsIdentityConnectorsResponse, err := dbMulticloudAwsProviderClient.ListOracleDbAwsIdentityConnectors(context.Background(), listOracleDbAwsIdentityConnectorsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting OracleDbAwsIdentityConnector list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, oracleDbAwsIdentityConnector := range listOracleDbAwsIdentityConnectorsResponse.Items {
		id := *oracleDbAwsIdentityConnector.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "OracleDbAwsIdentityConnectorId", id)
	}
	return resourceIds, nil
}

func DbmulticloudOracleDbAwsIdentityConnectorSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if oracleDbAwsIdentityConnectorResponse, ok := response.Response.(oci_dbmulticloud.GetOracleDbAwsIdentityConnectorResponse); ok {
		return oracleDbAwsIdentityConnectorResponse.LifecycleState != oci_dbmulticloud.OracleDbAwsIdentityConnectorLifecycleStateDeleted
	}
	return false
}

func DbmulticloudOracleDbAwsIdentityConnectorSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.DbMulticloudAwsProviderClient().GetOracleDbAwsIdentityConnector(context.Background(), oci_dbmulticloud.GetOracleDbAwsIdentityConnectorRequest{
		OracleDbAwsIdentityConnectorId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
