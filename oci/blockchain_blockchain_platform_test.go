// Copyright (c) 2017, 2020, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	oci_blockchain "github.com/oracle/oci-go-sdk/v29/blockchain"
	"github.com/oracle/oci-go-sdk/v29/common"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	BlockchainPlatformRequiredOnlyResource = BlockchainPlatformResourceDependencies +
		generateResourceFromRepresentationMap("oci_blockchain_blockchain_platform", "test_blockchain_platform", Required, Create, blockchainPlatformRepresentation)

	BlockchainPlatformResourceConfig = BlockchainPlatformResourceDependencies +
		generateResourceFromRepresentationMap("oci_blockchain_blockchain_platform", "test_blockchain_platform", Optional, Update, blockchainPlatformRepresentation)

	blockchainPlatformSingularDataSourceRepresentation = map[string]interface{}{
		"blockchain_platform_id": Representation{repType: Required, create: `${oci_blockchain_blockchain_platform.test_blockchain_platform.id}`},
	}

	blockchainPlatformDataSourceRepresentation = map[string]interface{}{
		"compartment_id": Representation{repType: Required, create: `${var.compartment_id}`},
		"display_name":   Representation{repType: Optional, create: blockchainPlatformDisplayName},
		"state":          Representation{repType: Optional, create: `ACTIVE`},
		"filter":         RepresentationGroup{Required, blockchainPlatformDataSourceFilterRepresentation}}
	blockchainPlatformDataSourceFilterRepresentation = map[string]interface{}{
		"name":   Representation{repType: Required, create: `id`},
		"values": Representation{repType: Required, create: []string{`${oci_blockchain_blockchain_platform.test_blockchain_platform.id}`}},
	}

	blockchainPlatformDisplayName = randomString(10, charsetLowerCaseWithoutDigits)

	blockchainPlatformRepresentation = map[string]interface{}{
		"compartment_id":      Representation{repType: Required, create: `${var.compartment_id}`},
		"compute_shape":       Representation{repType: Required, create: `ENTERPRISE_MEDIUM`},
		"display_name":        Representation{repType: Required, create: blockchainPlatformDisplayName},
		"platform_role":       Representation{repType: Required, create: `FOUNDER`},
		"defined_tags":        Representation{repType: Optional, create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"description":         Representation{repType: Optional, create: `description`, update: `description2`},
		"federated_user_id":   Representation{repType: Optional, create: `${oci_identity_user.test_user.id}`},
		"freeform_tags":       Representation{repType: Optional, create: map[string]string{"bar-key": "value"}, update: map[string]string{"Department": "Accounting"}},
		"idcs_access_token":   Representation{repType: Required, create: `${var.idcs_access_token}`},
		"is_byol":             Representation{repType: Optional, create: `false`},
		"replicas":            RepresentationGroup{Optional, blockchainPlatformReplicasRepresentation},
		"storage_size_in_tbs": Representation{repType: Optional, create: `1.0`, update: `2.0`},
		"total_ocpu_capacity": Representation{repType: Optional, create: `4`, update: `16`},
	}
	blockchainPlatformReplicasRepresentation = map[string]interface{}{
		"ca_count":      Representation{repType: Optional, create: `3`, update: `4`},
		"console_count": Representation{repType: Optional, create: `3`, update: `3`},
		"proxy_count":   Representation{repType: Optional, create: `3`, update: `4`},
	}

	BlockchainPlatformResourceDependencies = DefinedTagsDependencies +
		generateResourceFromRepresentationMap("oci_identity_user", "test_user", Required, Create, userRepresentation)
)

func TestBlockchainBlockchainPlatformResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestBlockchainBlockchainPlatformResource_basic")
	defer httpreplay.SaveScenario()

	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := getEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	idcsAccessToken := getEnvSettingWithBlankDefault("idcs_access_token")
	idcsAccessTokenVariableStr := fmt.Sprintf("variable \"idcs_access_token\" { default = \"%s\" }\n", idcsAccessToken)

	resourceName := "oci_blockchain_blockchain_platform.test_blockchain_platform"
	datasourceName := "data.oci_blockchain_blockchain_platforms.test_blockchain_platforms"
	singularDatasourceName := "data.oci_blockchain_blockchain_platform.test_blockchain_platform"

	var resId, resId2 string

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		CheckDestroy: testAccCheckBlockchainBlockchainPlatformDestroy,
		Steps: []resource.TestStep{
			// verify create
			{
				Config: config + compartmentIdVariableStr + idcsAccessTokenVariableStr + BlockchainPlatformResourceDependencies +
					generateResourceFromRepresentationMap("oci_blockchain_blockchain_platform", "test_blockchain_platform", Required, Create, blockchainPlatformRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "compute_shape", "ENTERPRISE_MEDIUM"),
					resource.TestCheckResourceAttr(resourceName, "display_name", blockchainPlatformDisplayName),
					resource.TestCheckResourceAttr(resourceName, "platform_role", "FOUNDER"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},

			// delete before next create
			{
				Config: config + compartmentIdVariableStr + idcsAccessTokenVariableStr + BlockchainPlatformResourceDependencies,
			},
			// verify create with optionals
			{
				Config: config + compartmentIdVariableStr + idcsAccessTokenVariableStr + BlockchainPlatformResourceDependencies +
					generateResourceFromRepresentationMap("oci_blockchain_blockchain_platform", "test_blockchain_platform", Optional, Create, blockchainPlatformRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "compute_shape", "ENTERPRISE_MEDIUM"),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "description", "description"),
					resource.TestCheckResourceAttr(resourceName, "display_name", blockchainPlatformDisplayName),
					resource.TestCheckResourceAttrSet(resourceName, "federated_user_id"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "is_byol", "false"),
					resource.TestCheckResourceAttr(resourceName, "platform_role", "FOUNDER"),
					resource.TestCheckResourceAttr(resourceName, "replicas.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "replicas.0.ca_count", "3"),
					resource.TestCheckResourceAttr(resourceName, "replicas.0.console_count", "3"),
					resource.TestCheckResourceAttr(resourceName, "replicas.0.proxy_count", "3"),
					resource.TestCheckResourceAttr(resourceName, "storage_size_in_tbs", "1"),
					resource.TestCheckResourceAttr(resourceName, "total_ocpu_capacity", "4"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						if isEnableExportCompartment, _ := strconv.ParseBool(getEnvSettingWithDefault("enable_export_compartment", "false")); isEnableExportCompartment {
							if errExport := testExportCompartmentWithResourceName(&resId, &compartmentId, resourceName); errExport != nil {
								return errExport
							}
						}
						return err
					},
				),
			},

			// verify update to the compartment (the compartment will be switched back in the next step)
			{
				Config: config + compartmentIdVariableStr + idcsAccessTokenVariableStr + compartmentIdUVariableStr + BlockchainPlatformResourceDependencies +
					generateResourceFromRepresentationMap("oci_blockchain_blockchain_platform", "test_blockchain_platform", Optional, Create,
						representationCopyWithNewProperties(blockchainPlatformRepresentation, map[string]interface{}{
							"compartment_id": Representation{repType: Required, create: `${var.compartment_id_for_update}`},
						})),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
					resource.TestCheckResourceAttr(resourceName, "compute_shape", "ENTERPRISE_MEDIUM"),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "description", "description"),
					resource.TestCheckResourceAttr(resourceName, "display_name", blockchainPlatformDisplayName),
					resource.TestCheckResourceAttrSet(resourceName, "federated_user_id"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "is_byol", "false"),
					resource.TestCheckResourceAttr(resourceName, "platform_role", "FOUNDER"),
					resource.TestCheckResourceAttr(resourceName, "replicas.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "storage_size_in_tbs", "1"),
					resource.TestCheckResourceAttr(resourceName, "total_ocpu_capacity", "4"),

					func(s *terraform.State) (err error) {
						resId2, err = fromInstanceState(s, resourceName, "id")
						if resId != resId2 {
							return fmt.Errorf("resource recreated when it was supposed to be updated")
						}
						return err
					},
				),
			},

			// verify updates to updatable parameters
			{
				Config: config + compartmentIdVariableStr + idcsAccessTokenVariableStr + BlockchainPlatformResourceDependencies +
					generateResourceFromRepresentationMap("oci_blockchain_blockchain_platform", "test_blockchain_platform", Optional, Update, blockchainPlatformRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "compute_shape", "ENTERPRISE_MEDIUM"),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "description", "description2"),
					resource.TestCheckResourceAttr(resourceName, "display_name", blockchainPlatformDisplayName),
					resource.TestCheckResourceAttrSet(resourceName, "federated_user_id"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "is_byol", "false"),
					resource.TestCheckResourceAttr(resourceName, "platform_role", "FOUNDER"),
					resource.TestCheckResourceAttr(resourceName, "replicas.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "replicas.0.ca_count", "4"),
					resource.TestCheckResourceAttr(resourceName, "replicas.0.console_count", "3"),
					resource.TestCheckResourceAttr(resourceName, "replicas.0.proxy_count", "4"),
					resource.TestCheckResourceAttr(resourceName, "storage_size_in_tbs", "2"),
					resource.TestCheckResourceAttr(resourceName, "total_ocpu_capacity", "16"),

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
					generateDataSourceFromRepresentationMap("oci_blockchain_blockchain_platforms", "test_blockchain_platforms", Optional, Update, blockchainPlatformDataSourceRepresentation) +
					compartmentIdVariableStr + idcsAccessTokenVariableStr + BlockchainPlatformResourceDependencies +
					generateResourceFromRepresentationMap("oci_blockchain_blockchain_platform", "test_blockchain_platform", Optional, Update, blockchainPlatformRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
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
					generateDataSourceFromRepresentationMap("oci_blockchain_blockchain_platform", "test_blockchain_platform", Required, Create, blockchainPlatformSingularDataSourceRepresentation) +
					compartmentIdVariableStr + idcsAccessTokenVariableStr + BlockchainPlatformResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(singularDatasourceName, "blockchain_platform_id"),

					resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(singularDatasourceName, "component_details.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "compute_shape", "ENTERPRISE_MEDIUM"),
					resource.TestCheckResourceAttr(singularDatasourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(singularDatasourceName, "description", "description2"),
					resource.TestCheckResourceAttr(singularDatasourceName, "display_name", blockchainPlatformDisplayName),
					resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
					resource.TestCheckResourceAttr(singularDatasourceName, "is_byol", "false"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "is_multi_ad"),
					resource.TestCheckResourceAttr(singularDatasourceName, "platform_role", "FOUNDER"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "platform_shape_type"),
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
					resource.TestCheckResourceAttr(singularDatasourceName, "total_ocpu_capacity", "16"),
				),
			},
			// remove singular datasource from previous step so that it doesn't conflict with import tests
			{
				Config: config + compartmentIdVariableStr + idcsAccessTokenVariableStr + BlockchainPlatformResourceConfig,
			},
			// verify resource import
			{
				Config:            config,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateVerifyIgnore: []string{
					"ca_cert_archive_text",
					"federated_user_id",
					"idcs_access_token",
				},
				ResourceName: resourceName,
			},
		},
	})
}

func testAccCheckBlockchainBlockchainPlatformDestroy(s *terraform.State) error {
	noResourceFound := true
	client := testAccProvider.Meta().(*OracleClients).blockchainPlatformClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_blockchain_blockchain_platform" {
			noResourceFound = false
			request := oci_blockchain.GetBlockchainPlatformRequest{}

			tmp := rs.Primary.ID
			request.BlockchainPlatformId = &tmp

			request.RequestMetadata.RetryPolicy = getRetryPolicy(true, "blockchain")

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
	if DependencyGraph == nil {
		initDependencyGraph()
	}
	if !inSweeperExcludeList("BlockchainBlockchainPlatform") {
		resource.AddTestSweepers("BlockchainBlockchainPlatform", &resource.Sweeper{
			Name:         "BlockchainBlockchainPlatform",
			Dependencies: DependencyGraph["blockchainPlatform"],
			F:            sweepBlockchainBlockchainPlatformResource,
		})
	}
}

func sweepBlockchainBlockchainPlatformResource(compartment string) error {
	blockchainPlatformClient := GetTestClients(&schema.ResourceData{}).blockchainPlatformClient()
	blockchainPlatformIds, err := getBlockchainPlatformIds(compartment)
	if err != nil {
		return err
	}
	for _, blockchainPlatformId := range blockchainPlatformIds {
		if ok := SweeperDefaultResourceId[blockchainPlatformId]; !ok {
			deleteBlockchainPlatformRequest := oci_blockchain.DeleteBlockchainPlatformRequest{}

			deleteBlockchainPlatformRequest.BlockchainPlatformId = &blockchainPlatformId

			deleteBlockchainPlatformRequest.RequestMetadata.RetryPolicy = getRetryPolicy(true, "blockchain")
			_, error := blockchainPlatformClient.DeleteBlockchainPlatform(context.Background(), deleteBlockchainPlatformRequest)
			if error != nil {
				fmt.Printf("Error deleting BlockchainPlatform %s %s, It is possible that the resource is already deleted. Please verify manually \n", blockchainPlatformId, error)
				continue
			}
			waitTillCondition(testAccProvider, &blockchainPlatformId, blockchainPlatformSweepWaitCondition, time.Duration(3*time.Minute),
				blockchainPlatformSweepResponseFetchOperation, "blockchain", true)
		}
	}
	return nil
}

func getBlockchainPlatformIds(compartment string) ([]string, error) {
	ids := getResourceIdsToSweep(compartment, "BlockchainPlatformId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	blockchainPlatformClient := GetTestClients(&schema.ResourceData{}).blockchainPlatformClient()

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
		addResourceIdToSweeperResourceIdMap(compartmentId, "BlockchainPlatformId", id)
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

func blockchainPlatformSweepResponseFetchOperation(client *OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.blockchainPlatformClient().GetBlockchainPlatform(context.Background(), oci_blockchain.GetBlockchainPlatformRequest{
		BlockchainPlatformId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
