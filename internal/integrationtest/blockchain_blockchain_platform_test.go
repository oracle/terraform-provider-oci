// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	oci_blockchain "github.com/oracle/oci-go-sdk/v65/blockchain"
	"github.com/oracle/oci-go-sdk/v65/common"

	"github.com/oracle/terraform-provider-oci/httpreplay"
)

var (
	BlockchainBlockchainPlatformRequiredOnlyResource = BlockchainBlockchainPlatformResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_blockchain_blockchain_platform", "test_blockchain_platform", acctest.Required, acctest.Create, blockchainPlatformRepresentation)

	BlockchainBlockchainPlatformResourceConfig = BlockchainBlockchainPlatformResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_blockchain_blockchain_platform", "test_blockchain_platform", acctest.Optional, acctest.Update, blockchainPlatformRepresentation)

	BlockchainblockchainPlatformSingularDataSourceRepresentation = map[string]interface{}{
		"blockchain_platform_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_blockchain_blockchain_platform.test_blockchain_platform.id}`},
	}

	BlockchainblockchainPlatformDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":   acctest.Representation{RepType: acctest.Optional, Create: blockchainPlatformDisplayName},
		"state":          acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"filter":         acctest.RepresentationGroup{RepType: acctest.Required, Group: blockchainPlatformDataSourceFilterRepresentation}}
	blockchainPlatformDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_blockchain_blockchain_platform.test_blockchain_platform.id}`}},
	}

	blockchainPlatformDisplayName = utils.RandomString(10, utils.CharsetLowerCaseWithoutDigits)

	blockchainPlatformRepresentation = map[string]interface{}{
		"compartment_id":      acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"compute_shape":       acctest.Representation{RepType: acctest.Required, Create: `ENTERPRISE_MEDIUM`},
		"display_name":        acctest.Representation{RepType: acctest.Required, Create: blockchainPlatformDisplayName},
		"idcs_access_token":   acctest.Representation{RepType: acctest.Required, Create: `${var.idcs_access_token}`},
		"platform_role":       acctest.Representation{RepType: acctest.Required, Create: `FOUNDER`},
		"defined_tags":        acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"description":         acctest.Representation{RepType: acctest.Optional, Create: `description`, Update: `description2`},
		"federated_user_id":   acctest.Representation{RepType: acctest.Optional, Create: `${oci_identity_user.test_user.id}`},
		"freeform_tags":       acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"bar-key": "value"}, Update: map[string]string{"Department": "Accounting"}},
		"is_byol":             acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"replicas":            acctest.RepresentationGroup{RepType: acctest.Optional, Group: blockchainPlatformReplicasRepresentation},
		"storage_size_in_tbs": acctest.Representation{RepType: acctest.Optional, Create: `1.0`, Update: `2.0`},
		"total_ocpu_capacity": acctest.Representation{RepType: acctest.Optional, Create: `16`, Update: `32`},
		"platform_version":    acctest.Representation{RepType: acctest.Optional, Create: `Hyperledger Fabric v1.4.7`},
		"load_balancer_shape": acctest.Representation{RepType: acctest.Optional, Create: `LB_100_MBPS`, Update: `LB_400_MBPS`},
	}
	blockchainPlatformReplicasRepresentation = map[string]interface{}{
		"ca_count":      acctest.Representation{RepType: acctest.Optional, Create: `3`, Update: `4`},
		"console_count": acctest.Representation{RepType: acctest.Optional, Create: `3`, Update: `3`},
		"proxy_count":   acctest.Representation{RepType: acctest.Optional, Create: `3`, Update: `4`},
	}

	BlockchainBlockchainPlatformResourceDependencies = DefinedTagsDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_identity_user", "test_user", acctest.Required, acctest.Create, IdentityUserRepresentation)
)

// issue-routing-tag: blockchain/default
func TestBlockchainBlockchainPlatformResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestBlockchainBlockchainPlatformResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	idcsAccessToken := utils.GetEnvSettingWithBlankDefault("idcs_access_token")
	idcsAccessTokenVariableStr := fmt.Sprintf("variable \"idcs_access_token\" { default = \"%s\" }\n", idcsAccessToken)

	resourceName := "oci_blockchain_blockchain_platform.test_blockchain_platform"
	datasourceName := "data.oci_blockchain_blockchain_platforms.test_blockchain_platforms"
	singularDatasourceName := "data.oci_blockchain_blockchain_platform.test_blockchain_platform"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+BlockchainBlockchainPlatformResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_blockchain_blockchain_platform", "test_blockchain_platform", acctest.Optional, acctest.Create, blockchainPlatformRepresentation), "blockchain", "blockchainPlatform", t)

	acctest.ResourceTest(t, testAccCheckBlockchainBlockchainPlatformDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + idcsAccessTokenVariableStr + BlockchainBlockchainPlatformResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_blockchain_blockchain_platform", "test_blockchain_platform", acctest.Required, acctest.Create, blockchainPlatformRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "compute_shape", "ENTERPRISE_MEDIUM"),
				resource.TestCheckResourceAttr(resourceName, "display_name", blockchainPlatformDisplayName),
				resource.TestCheckResourceAttrSet(resourceName, "idcs_access_token"),
				resource.TestCheckResourceAttr(resourceName, "platform_role", "FOUNDER"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + idcsAccessTokenVariableStr + BlockchainBlockchainPlatformResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + idcsAccessTokenVariableStr + BlockchainBlockchainPlatformResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_blockchain_blockchain_platform", "test_blockchain_platform", acctest.Optional, acctest.Create, blockchainPlatformRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "compute_shape", "ENTERPRISE_MEDIUM"),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", blockchainPlatformDisplayName),
				resource.TestCheckResourceAttrSet(resourceName, "federated_user_id"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_byol", "false"),
				resource.TestCheckResourceAttr(resourceName, "platform_role", "FOUNDER"),
				resource.TestCheckResourceAttr(resourceName, "platform_version", "Hyperledger Fabric v1.4.7"),
				resource.TestCheckResourceAttr(resourceName, "replicas.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "replicas.0.ca_count", "3"),
				resource.TestCheckResourceAttr(resourceName, "replicas.0.console_count", "3"),
				resource.TestCheckResourceAttr(resourceName, "replicas.0.proxy_count", "3"),
				resource.TestCheckResourceAttr(resourceName, "storage_size_in_tbs", "1"),
				resource.TestCheckResourceAttr(resourceName, "total_ocpu_capacity", "16"),
				resource.TestCheckResourceAttr(resourceName, "load_balancer_shape", "LB_100_MBPS"),

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

		// verify Update to the compartment (the compartment will be switched back in the next step)
		{
			Config: config + compartmentIdVariableStr + idcsAccessTokenVariableStr + compartmentIdUVariableStr + BlockchainBlockchainPlatformResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_blockchain_blockchain_platform", "test_blockchain_platform", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(blockchainPlatformRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "compute_shape", "ENTERPRISE_MEDIUM"),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", blockchainPlatformDisplayName),
				resource.TestCheckResourceAttrSet(resourceName, "federated_user_id"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_byol", "false"),
				resource.TestCheckResourceAttr(resourceName, "platform_role", "FOUNDER"),
				resource.TestCheckResourceAttr(resourceName, "platform_version", "Hyperledger Fabric v1.4.7"),
				resource.TestCheckResourceAttr(resourceName, "replicas.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "storage_size_in_tbs", "1"),
				resource.TestCheckResourceAttr(resourceName, "total_ocpu_capacity", "16"),
				resource.TestCheckResourceAttr(resourceName, "load_balancer_shape", "LB_100_MBPS"),

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
			Config: config + compartmentIdVariableStr + idcsAccessTokenVariableStr + BlockchainBlockchainPlatformResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_blockchain_blockchain_platform", "test_blockchain_platform", acctest.Optional, acctest.Update, blockchainPlatformRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "compute_shape", "ENTERPRISE_MEDIUM"),
				resource.TestCheckResourceAttr(resourceName, "description", "description2"),
				resource.TestCheckResourceAttr(resourceName, "display_name", blockchainPlatformDisplayName),
				resource.TestCheckResourceAttrSet(resourceName, "federated_user_id"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_byol", "false"),
				resource.TestCheckResourceAttr(resourceName, "platform_role", "FOUNDER"),
				resource.TestCheckResourceAttr(resourceName, "platform_version", "Hyperledger Fabric v1.4.7"),
				resource.TestCheckResourceAttr(resourceName, "replicas.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "replicas.0.ca_count", "4"),
				resource.TestCheckResourceAttr(resourceName, "replicas.0.console_count", "3"),
				resource.TestCheckResourceAttr(resourceName, "replicas.0.proxy_count", "4"),
				resource.TestCheckResourceAttr(resourceName, "storage_size_in_tbs", "2"),
				resource.TestCheckResourceAttr(resourceName, "total_ocpu_capacity", "32"),
				resource.TestCheckResourceAttr(resourceName, "load_balancer_shape", "LB_400_MBPS"),

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
				acctest.GenerateDataSourceFromRepresentationMap("oci_blockchain_blockchain_platforms", "test_blockchain_platforms", acctest.Optional, acctest.Update, BlockchainblockchainPlatformDataSourceRepresentation) +
				compartmentIdVariableStr + idcsAccessTokenVariableStr + BlockchainBlockchainPlatformResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_blockchain_blockchain_platform", "test_blockchain_platform", acctest.Optional, acctest.Update, blockchainPlatformRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", blockchainPlatformDisplayName),
				resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),

				resource.TestCheckResourceAttr(datasourceName, "blockchain_platform_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "blockchain_platform_collection.0.items.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_blockchain_blockchain_platform", "test_blockchain_platform", acctest.Required, acctest.Create, BlockchainblockchainPlatformSingularDataSourceRepresentation) +
				compartmentIdVariableStr + idcsAccessTokenVariableStr + BlockchainBlockchainPlatformResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "blockchain_platform_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "component_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "compute_shape", "ENTERPRISE_MEDIUM"),
				resource.TestCheckResourceAttr(singularDatasourceName, "description", "description2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", blockchainPlatformDisplayName),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "is_byol", "false"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "is_multi_ad"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "load_balancer_shape"),
				resource.TestCheckResourceAttr(singularDatasourceName, "platform_role", "FOUNDER"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "platform_shape_type"),
				resource.TestCheckResourceAttr(singularDatasourceName, "platform_version", "Hyperledger Fabric v1.4.7"),
				resource.TestCheckResourceAttr(singularDatasourceName, "replicas.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "replicas.0.ca_count"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "replicas.0.console_count"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "replicas.0.proxy_count"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "service_endpoint"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "service_version"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttr(singularDatasourceName, "storage_size_in_tbs", "2"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "storage_used_in_tbs"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
				resource.TestCheckResourceAttr(singularDatasourceName, "total_ocpu_capacity", "32"),
				resource.TestCheckResourceAttr(singularDatasourceName, "load_balancer_shape", "LB_400_MBPS"),
			),
		},
		// verify resource import
		{
			Config:            config + BlockchainBlockchainPlatformRequiredOnlyResource,
			ImportState:       true,
			ImportStateVerify: true,
			ImportStateVerifyIgnore: []string{
				"ca_cert_archive_text",
				"federated_user_id",
				"idcs_access_token",
			},
			ResourceName: resourceName,
		},
	})
}

func testAccCheckBlockchainBlockchainPlatformDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).BlockchainPlatformClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_blockchain_blockchain_platform" {
			noResourceFound = false
			request := oci_blockchain.GetBlockchainPlatformRequest{}

			tmp := rs.Primary.ID
			request.BlockchainPlatformId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "blockchain")

			response, err := client.GetBlockchainPlatform(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_blockchain.BlockchainPlatformLifecycleStateDeleted): true,
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
	if !acctest.InSweeperExcludeList("BlockchainBlockchainPlatform") {
		resource.AddTestSweepers("BlockchainBlockchainPlatform", &resource.Sweeper{
			Name:         "BlockchainBlockchainPlatform",
			Dependencies: acctest.DependencyGraph["blockchainPlatform"],
			F:            sweepBlockchainBlockchainPlatformResource,
		})
	}
}

func sweepBlockchainBlockchainPlatformResource(compartment string) error {
	blockchainPlatformClient := acctest.GetTestClients(&schema.ResourceData{}).BlockchainPlatformClient()
	blockchainPlatformIds, err := getBlockchainPlatformIds(compartment)
	if err != nil {
		return err
	}
	for _, blockchainPlatformId := range blockchainPlatformIds {
		if ok := acctest.SweeperDefaultResourceId[blockchainPlatformId]; !ok {
			deleteBlockchainPlatformRequest := oci_blockchain.DeleteBlockchainPlatformRequest{}

			deleteBlockchainPlatformRequest.BlockchainPlatformId = &blockchainPlatformId

			deleteBlockchainPlatformRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "blockchain")
			_, error := blockchainPlatformClient.DeleteBlockchainPlatform(context.Background(), deleteBlockchainPlatformRequest)
			if error != nil {
				fmt.Printf("Error deleting BlockchainPlatform %s %s, It is possible that the resource is already deleted. Please verify manually \n", blockchainPlatformId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &blockchainPlatformId, blockchainPlatformSweepWaitCondition, time.Duration(3*time.Minute),
				blockchainPlatformSweepResponseFetchOperation, "blockchain", true)
		}
	}
	return nil
}

func getBlockchainPlatformIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "BlockchainPlatformId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	blockchainPlatformClient := acctest.GetTestClients(&schema.ResourceData{}).BlockchainPlatformClient()

	listBlockchainPlatformsRequest := oci_blockchain.ListBlockchainPlatformsRequest{}
	listBlockchainPlatformsRequest.CompartmentId = &compartmentId
	listBlockchainPlatformsRequest.LifecycleState = oci_blockchain.BlockchainPlatformLifecycleStateActive
	listBlockchainPlatformsResponse, err := blockchainPlatformClient.ListBlockchainPlatforms(context.Background(), listBlockchainPlatformsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting BlockchainPlatform list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, blockchainPlatform := range listBlockchainPlatformsResponse.Items {
		id := *blockchainPlatform.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "BlockchainPlatformId", id)
	}
	return resourceIds, nil
}

func blockchainPlatformSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if blockchainPlatformResponse, ok := response.Response.(oci_blockchain.GetBlockchainPlatformResponse); ok {
		return blockchainPlatformResponse.LifecycleState != oci_blockchain.BlockchainPlatformLifecycleStateDeleted
	}
	return false
}

func blockchainPlatformSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.BlockchainPlatformClient().GetBlockchainPlatform(context.Background(), oci_blockchain.GetBlockchainPlatformRequest{
		BlockchainPlatformId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
