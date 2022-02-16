// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/terraform-providers/terraform-provider-oci/internal/acctest"
	tf_client "github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/resourcediscovery"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"
	"github.com/terraform-providers/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	oci_bds "github.com/oracle/oci-go-sdk/v58/bds"
	"github.com/oracle/oci-go-sdk/v58/common"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	BdsInstanceMetastoreConfigRequiredOnlyResource = BdsInstanceMetastoreConfigResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_bds_bds_instance_metastore_config", "test_bds_instance_metastore_config", acctest.Required, acctest.Create, bdsInstanceMetastoreConfigRepresentation)

	BdsInstanceMetastoreConfigResourceConfig = BdsInstanceMetastoreConfigResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_bds_bds_instance_metastore_config", "test_bds_instance_metastore_config", acctest.Optional, acctest.Update, bdsInstanceMetastoreConfigRepresentation)

	bdsInstanceMetastoreConfigSingularDataSourceRepresentation = map[string]interface{}{
		"bds_instance_id":     acctest.Representation{RepType: acctest.Required, Create: `${oci_bds_bds_instance.test_bds_instance.id}`},
		"metastore_config_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_bds_bds_instance_metastore_config.test_bds_instance_metastore_config.id}`},
	}

	bdsInstanceMetastoreConfigDataSourceRepresentation = map[string]interface{}{
		"bds_instance_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_bds_bds_instance.test_bds_instance.id}`},
		"state":           acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"filter":          acctest.RepresentationGroup{RepType: acctest.Required, Group: bdsInstanceMetastoreConfigDataSourceFilterRepresentation}}
	bdsInstanceMetastoreConfigDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_bds_bds_instance_metastore_config.test_bds_instance_metastore_config.id}`}},
	}

	bdsInstanceMetastoreConfigRepresentation = map[string]interface{}{
		"bds_api_key_id":         acctest.Representation{RepType: acctest.Required, Create: `${oci_bds_bds_instance_api_key.test_bds_instance_api_key.id}`},
		"bds_api_key_passphrase": acctest.Representation{RepType: acctest.Required, Create: `V2VsY29tZTE=`, Update: `V2VsY29tZTE=`},
		"bds_instance_id":        acctest.Representation{RepType: acctest.Required, Create: `${oci_bds_bds_instance.test_bds_instance.id}`},
		"cluster_admin_password": acctest.Representation{RepType: acctest.Required, Create: `V2VsY29tZTE=`, Update: `V2VsY29tZTE=`},
		"metastore_id":           acctest.Representation{RepType: acctest.Required, Create: `${var.metastore_id}`},
		"display_name":           acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
	}

	BdsInstanceMetastoreConfigResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_bds_bds_instance", "test_bds_instance", acctest.Required, acctest.Create, bdsInstanceOdhWithNatGatewayRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_subnet", "test_subnet", acctest.Required, acctest.Create, subnetRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", acctest.Required, acctest.Create, vcnRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_bds_bds_instance_api_key", "test_bds_instance_api_key", acctest.Required, acctest.Create, bdsInstanceApiKeyRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_identity_user", "test_user", acctest.Required, acctest.Create, userRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_identity_dynamic_group", "test_dynamic_group", acctest.Required, acctest.Create, bdsDcatDynamicGroupRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_identity_group", "test_group", acctest.Required, acctest.Create, bdsApiKeyUserGroupRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_identity_policy", "test_policy", acctest.Required, acctest.Create, bdsMetastoreConfigTestPolicyRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_identity_user_group_membership", "test_user_group_membership", acctest.Required, acctest.Create, bdsMetastoreConfigTestUserGroupMembershipRepresentation)
)

// issue-routing-tag: bds/default
func TestBdsBdsInstanceMetastoreConfigResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestBdsBdsInstanceMetastoreConfigResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	metastoreId := utils.GetEnvSettingWithBlankDefault("metastore_id")
	metastoreIdVariableStr := fmt.Sprintf("variable \"metastore_id\" { default = \"%s\" }\n", metastoreId)

	resourceName := "oci_bds_bds_instance_metastore_config.test_bds_instance_metastore_config"
	datasourceName := "data.oci_bds_bds_instance_metastore_configs.test_bds_instance_metastore_configs"
	singularDatasourceName := "data.oci_bds_bds_instance_metastore_config.test_bds_instance_metastore_config"

	var resId, resId2, bdsId string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+BdsInstanceMetastoreConfigResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_bds_bds_instance_metastore_config", "test_bds_instance_metastore_config", acctest.Optional, acctest.Create, bdsInstanceMetastoreConfigRepresentation), "bds", "bdsInstanceMetastoreConfig", t)

	acctest.ResourceTest(t, testAccCheckBdsBdsInstanceMetastoreConfigDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + BdsInstanceMetastoreConfigResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_bds_bds_instance_metastore_config", "test_bds_instance_metastore_config", acctest.Required, acctest.Create, bdsInstanceMetastoreConfigRepresentation) + metastoreIdVariableStr,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "bds_api_key_id"),
				resource.TestCheckResourceAttr(resourceName, "bds_api_key_passphrase", "V2VsY29tZTE="),
				resource.TestCheckResourceAttrSet(resourceName, "bds_instance_id"),
				resource.TestCheckResourceAttr(resourceName, "cluster_admin_password", "V2VsY29tZTE="),
				resource.TestCheckResourceAttrSet(resourceName, "metastore_id"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					bdsId, _ = acctest.FromInstanceState(s, resourceName, "bds_instance_id")
					return err
				},
			),
		},

		// delete before next Create
		{
			PreConfig: func() {
				activateLocalMetastore(bdsId)
			},
			Config: config + compartmentIdVariableStr + BdsInstanceMetastoreConfigResourceDependencies + metastoreIdVariableStr,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + BdsInstanceMetastoreConfigResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_bds_bds_instance_metastore_config", "test_bds_instance_metastore_config", acctest.Optional, acctest.Create, bdsInstanceMetastoreConfigRepresentation) + metastoreIdVariableStr,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "bds_api_key_id"),
				resource.TestCheckResourceAttr(resourceName, "bds_api_key_passphrase", "V2VsY29tZTE="),
				resource.TestCheckResourceAttrSet(resourceName, "bds_instance_id"),
				resource.TestCheckResourceAttr(resourceName, "cluster_admin_password", "V2VsY29tZTE="),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "metastore_id"),
				resource.TestCheckResourceAttrSet(resourceName, "metastore_type"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),

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
			Config: config + compartmentIdVariableStr + BdsInstanceMetastoreConfigResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_bds_bds_instance_metastore_config", "test_bds_instance_metastore_config", acctest.Optional, acctest.Update, bdsInstanceMetastoreConfigRepresentation) + metastoreIdVariableStr,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "bds_api_key_id"),
				resource.TestCheckResourceAttr(resourceName, "bds_api_key_passphrase", "V2VsY29tZTE="),
				resource.TestCheckResourceAttrSet(resourceName, "bds_instance_id"),
				resource.TestCheckResourceAttr(resourceName, "cluster_admin_password", "V2VsY29tZTE="),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "metastore_id"),
				resource.TestCheckResourceAttrSet(resourceName, "metastore_type"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),

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
				acctest.GenerateDataSourceFromRepresentationMap("oci_bds_bds_instance_metastore_configs", "test_bds_instance_metastore_configs", acctest.Optional, acctest.Update, bdsInstanceMetastoreConfigDataSourceRepresentation) +
				compartmentIdVariableStr + BdsInstanceMetastoreConfigResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_bds_bds_instance_metastore_config", "test_bds_instance_metastore_config", acctest.Optional, acctest.Update, bdsInstanceMetastoreConfigRepresentation) + metastoreIdVariableStr,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "bds_instance_id"),
				resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),

				resource.TestCheckResourceAttr(datasourceName, "bds_metastore_configurations.#", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "bds_metastore_configurations.0.bds_api_key_id"),
				resource.TestCheckResourceAttr(datasourceName, "bds_metastore_configurations.0.display_name", "displayName2"),
				resource.TestCheckResourceAttrSet(datasourceName, "bds_metastore_configurations.0.id"),
				resource.TestCheckResourceAttrSet(datasourceName, "bds_metastore_configurations.0.metastore_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "bds_metastore_configurations.0.metastore_type"),
				resource.TestCheckResourceAttrSet(datasourceName, "bds_metastore_configurations.0.state"),
				resource.TestCheckResourceAttrSet(datasourceName, "bds_metastore_configurations.0.time_created"),
				resource.TestCheckResourceAttrSet(datasourceName, "bds_metastore_configurations.0.time_updated"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_bds_bds_instance_metastore_config", "test_bds_instance_metastore_config", acctest.Required, acctest.Create, bdsInstanceMetastoreConfigSingularDataSourceRepresentation) +
				compartmentIdVariableStr + BdsInstanceMetastoreConfigResourceConfig + metastoreIdVariableStr,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "bds_instance_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "metastore_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "metastore_type"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
		// remove singular datasource from previous step so that it doesn't conflict with import tests
		{
			PreConfig: func() {
				activateLocalMetastore(bdsId)
			},
			Config: config + compartmentIdVariableStr + BdsInstanceMetastoreConfigResourceConfig + metastoreIdVariableStr,
		},
		// verify resource import
		{
			Config:            config,
			ImportState:       true,
			ImportStateVerify: true,
			ImportStateIdFunc: getBdsMetastoreConfigCompositeId(resourceName),
			ImportStateVerifyIgnore: []string{
				"bds_api_key_passphrase",
				"cluster_admin_password",
			},
			ResourceName: resourceName,
		},
	})
}

func testAccCheckBdsBdsInstanceMetastoreConfigDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).BdsClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_bds_bds_instance_metastore_config" {
			noResourceFound = false
			request := oci_bds.GetBdsMetastoreConfigurationRequest{}

			if value, ok := rs.Primary.Attributes["bds_instance_id"]; ok {
				request.BdsInstanceId = &value
			}

			if value, ok := rs.Primary.Attributes["id"]; ok {
				request.MetastoreConfigId = &value
			}

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "bds")

			response, err := client.GetBdsMetastoreConfiguration(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_bds.BdsMetastoreConfigurationLifecycleStateDeleted): true,
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
	if !acctest.InSweeperExcludeList("BdsBdsInstanceMetastoreConfig") {
		resource.AddTestSweepers("BdsBdsInstanceMetastoreConfig", &resource.Sweeper{
			Name:         "BdsBdsInstanceMetastoreConfig",
			Dependencies: acctest.DependencyGraph["bdsInstanceMetastoreConfig"],
			F:            sweepBdsBdsInstanceMetastoreConfigResource,
		})
	}
}

func sweepBdsBdsInstanceMetastoreConfigResource(compartment string) error {
	bdsClient := acctest.GetTestClients(&schema.ResourceData{}).BdsClient()
	bdsInstanceMetastoreConfigIds, err := getBdsInstanceMetastoreConfigIds(compartment)
	if err != nil {
		return err
	}
	for _, bdsInstanceMetastoreConfigId := range bdsInstanceMetastoreConfigIds {
		if ok := acctest.SweeperDefaultResourceId[bdsInstanceMetastoreConfigId]; !ok {
			deleteBdsMetastoreConfigurationRequest := oci_bds.DeleteBdsMetastoreConfigurationRequest{}

			deleteBdsMetastoreConfigurationRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "bds")
			_, error := bdsClient.DeleteBdsMetastoreConfiguration(context.Background(), deleteBdsMetastoreConfigurationRequest)
			if error != nil {
				fmt.Printf("Error deleting BdsInstanceMetastoreConfig %s %s, It is possible that the resource is already deleted. Please verify manually \n", bdsInstanceMetastoreConfigId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &bdsInstanceMetastoreConfigId, bdsInstanceMetastoreConfigSweepWaitCondition, time.Duration(3*time.Minute),
				bdsInstanceMetastoreConfigSweepResponseFetchOperation, "bds", true)
		}
	}
	return nil
}

func getBdsInstanceMetastoreConfigIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "BdsInstanceMetastoreConfigId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	bdsClient := acctest.GetTestClients(&schema.ResourceData{}).BdsClient()

	listBdsMetastoreConfigurationsRequest := oci_bds.ListBdsMetastoreConfigurationsRequest{}

	bdsInstanceIds, error := getBdsInstanceIds(compartment)
	if error != nil {
		return resourceIds, fmt.Errorf("Error getting bdsInstanceId required for BdsInstanceMetastoreConfig resource requests \n")
	}
	for _, bdsInstanceId := range bdsInstanceIds {
		listBdsMetastoreConfigurationsRequest.BdsInstanceId = &bdsInstanceId

		listBdsMetastoreConfigurationsRequest.LifecycleState = oci_bds.BdsMetastoreConfigurationLifecycleStateActive
		listBdsMetastoreConfigurationsResponse, err := bdsClient.ListBdsMetastoreConfigurations(context.Background(), listBdsMetastoreConfigurationsRequest)

		if err != nil {
			return resourceIds, fmt.Errorf("Error getting BdsInstanceMetastoreConfig list for compartment id : %s , %s \n", compartmentId, err)
		}
		for _, bdsInstanceMetastoreConfig := range listBdsMetastoreConfigurationsResponse.Items {
			id := *bdsInstanceMetastoreConfig.Id
			resourceIds = append(resourceIds, id)
			acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "BdsInstanceMetastoreConfigId", id)
		}

	}
	return resourceIds, nil
}

func bdsInstanceMetastoreConfigSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if bdsInstanceMetastoreConfigResponse, ok := response.Response.(oci_bds.GetBdsMetastoreConfigurationResponse); ok {
		return bdsInstanceMetastoreConfigResponse.LifecycleState != oci_bds.BdsMetastoreConfigurationLifecycleStateDeleted
	}
	return false
}

func bdsInstanceMetastoreConfigSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.BdsClient().GetBdsMetastoreConfiguration(context.Background(), oci_bds.GetBdsMetastoreConfigurationRequest{RequestMetadata: common.RequestMetadata{
		RetryPolicy: retryPolicy,
	},
	})
	return err
}

func getBdsMetastoreConfigCompositeId(resourceName string) resource.ImportStateIdFunc {
	return func(s *terraform.State) (string, error) {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return "", fmt.Errorf("not found: %s", resourceName)
		}

		return fmt.Sprintf("bdsInstances/%s/metastoreConfigs/%s", rs.Primary.Attributes["bds_instance_id"], rs.Primary.Attributes["id"]), nil
	}
}
