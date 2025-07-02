// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_bds "github.com/oracle/oci-go-sdk/v65/bds"
	"github.com/oracle/oci-go-sdk/v65/common"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"

	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	BdsBdsInstanceNodeReplaceConfigurationRequiredOnlyResource = acctest.GenerateResourceFromRepresentationMap("oci_bds_bds_instance_node_replace_configuration", "test_bds_instance_node_replace_configuration", acctest.Required, acctest.Create, BdsBdsInstanceNodeReplaceConfigurationRepresentation)

	BdsBdsInstanceNodeReplaceConfigurationResourceConfig = BdsBdsInstanceNodeReplaceConfigurationResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_bds_bds_instance_node_replace_configuration", "test_bds_instance_node_replace_configuration", acctest.Optional, acctest.Update, BdsBdsInstanceNodeReplaceConfigurationRepresentation)

	BdsBdsInstanceNodeReplaceConfigurationSingularDataSourceRepresentation = map[string]interface{}{
		"bds_instance_id":               acctest.Representation{RepType: acctest.Required, Create: `${var.bds_instance_id}`},
		"node_replace_configuration_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_bds_bds_instance_node_replace_configuration.test_bds_instance_node_replace_configuration.id}`},
	}

	BdsBdsInstanceNodeReplaceConfigurationDataSourceRepresentation = map[string]interface{}{
		"bds_instance_id": acctest.Representation{RepType: acctest.Required, Create: `${var.bds_instance_id}`},
	}
	BdsBdsInstanceNodeReplaceConfigurationDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_bds_bds_instance_node_replace_configuration.test_bds_instance_node_replace_configuration.id}`}},
	}

	BdsBdsInstanceNodeReplaceConfigurationRepresentation = map[string]interface{}{
		"bds_instance_id":        acctest.Representation{RepType: acctest.Required, Create: `${var.bds_instance_id}`},
		"cluster_admin_password": acctest.Representation{RepType: acctest.Required, Create: `T3JhY2xlVGVhbVVTQSExMjM=`},
		"duration_in_minutes":    acctest.Representation{RepType: acctest.Required, Create: `20`, Update: `30`},
		"level_type_details":     acctest.RepresentationGroup{RepType: acctest.Required, Group: BdsBdsInstanceNodeReplaceConfigurationLevelTypeDetailsRepresentation},
		"metric_type":            acctest.Representation{RepType: acctest.Required, Create: `INSTANCE_STATUS`, Update: `INSTANCE_ACCESSIBILITY_STATUS`},
		"display_name":           acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
	}
	BdsBdsInstanceNodeReplaceConfigurationLevelTypeDetailsRepresentation = map[string]interface{}{
		"level_type": acctest.Representation{RepType: acctest.Required, Create: `NODE_TYPE_LEVEL`},
		"node_type":  acctest.Representation{RepType: acctest.Optional, Create: `MASTER`, Update: `UTILITY`},
	}

	BdsBdsInstanceNodeReplaceConfigurationResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_bds_bds_instance", "test_bds_instance", acctest.Optional, acctest.Create, bdsInstanceOdhRepresentation)
	//	acctest.GenerateResourceFromRepresentationMap("oci_core_subnet", "test_subnet", acctest.Required, acctest.Create, CoreSubnetRepresentation) +
	//	acctest.GenerateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", acctest.Required, acctest.Create, CoreVcnRepresentation) +

)

// issue-routing-tag: bds/default
func TestBdsBdsInstanceNodeReplaceConfigurationResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestBdsBdsInstanceNodeReplaceConfigurationResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	subnetId := utils.GetEnvSettingWithBlankDefault("subnet_ocid")
	subnetIdVariableStr := fmt.Sprintf("variable \"subnet_id\" { default = \"%s\" }\n", subnetId)

	bdsInstanceId := utils.GetEnvSettingWithBlankDefault("bds_instance_ocid")
	bdsInstanceIdVariableStr := fmt.Sprintf("variable \"bds_instance_id\" { default = \"%s\" }\n", bdsInstanceId)

	resourceName := "oci_bds_bds_instance_node_replace_configuration.test_bds_instance_node_replace_configuration"
	datasourceName := "data.oci_bds_bds_instance_node_replace_configurations.test_bds_instance_node_replace_configurations"
	singularDatasourceName := "data.oci_bds_bds_instance_node_replace_configuration.test_bds_instance_node_replace_configuration"

	var resId, resId2 string
	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + subnetIdVariableStr + bdsInstanceIdVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_bds_bds_instance_node_replace_configuration", "test_bds_instance_node_replace_configuration", acctest.Optional, acctest.Create, BdsBdsInstanceNodeReplaceConfigurationRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "bds_instance_id"),
				resource.TestCheckResourceAttr(resourceName, "cluster_admin_password", "T3JhY2xlVGVhbVVTQSExMjM="),
				resource.TestCheckResourceAttr(resourceName, "duration_in_minutes", "20"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "level_type_details.0.node_type", "MASTER"),
				resource.TestCheckResourceAttr(resourceName, "metric_type", "INSTANCE_STATUS"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "time_updated"),

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
			Config: config + compartmentIdVariableStr + subnetIdVariableStr + bdsInstanceIdVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_bds_bds_instance_node_replace_configuration", "test_bds_instance_node_replace_configuration", acctest.Optional, acctest.Update, BdsBdsInstanceNodeReplaceConfigurationRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "bds_instance_id"),
				resource.TestCheckResourceAttr(resourceName, "cluster_admin_password", "T3JhY2xlVGVhbVVTQSExMjM="),
				resource.TestCheckResourceAttr(resourceName, "duration_in_minutes", "30"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "level_type_details.0.node_type", "UTILITY"),
				resource.TestCheckResourceAttr(resourceName, "metric_type", "INSTANCE_ACCESSIBILITY_STATUS"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "time_updated"),

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
				acctest.GenerateDataSourceFromRepresentationMap("oci_bds_bds_instance_node_replace_configurations", "test_bds_instance_node_replace_configurations", acctest.Optional, acctest.Update, BdsBdsInstanceNodeReplaceConfigurationDataSourceRepresentation) +
				compartmentIdVariableStr + subnetIdVariableStr + bdsInstanceIdVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_bds_bds_instance_node_replace_configuration", "test_bds_instance_node_replace_configuration", acctest.Optional, acctest.Update, BdsBdsInstanceNodeReplaceConfigurationRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "bds_instance_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "node_replace_configurations.0.bds_instance_id"),
				resource.TestCheckResourceAttr(datasourceName, "node_replace_configurations.0.display_name", "displayName2"),
				resource.TestCheckResourceAttrSet(datasourceName, "node_replace_configurations.0.id"),
				resource.TestCheckResourceAttr(datasourceName, "node_replace_configurations.0.level_type_details.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "node_replace_configurations.0.level_type_details.0.level_type", "NODE_TYPE_LEVEL"),
				resource.TestCheckResourceAttr(datasourceName, "node_replace_configurations.0.level_type_details.0.node_type", "UTILITY"),
				resource.TestCheckResourceAttrSet(datasourceName, "node_replace_configurations.0.state"),
				resource.TestCheckResourceAttrSet(datasourceName, "node_replace_configurations.0.time_created"),
				resource.TestCheckResourceAttrSet(datasourceName, "node_replace_configurations.0.time_updated"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_bds_bds_instance_node_replace_configuration", "test_bds_instance_node_replace_configuration", acctest.Required, acctest.Create, BdsBdsInstanceNodeReplaceConfigurationSingularDataSourceRepresentation) +
				compartmentIdVariableStr + subnetIdVariableStr + bdsInstanceIdVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_bds_bds_instance_node_replace_configuration", "test_bds_instance_node_replace_configuration", acctest.Optional, acctest.Update, BdsBdsInstanceNodeReplaceConfigurationRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "bds_instance_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "node_replace_configuration_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "duration_in_minutes", "30"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "level_type_details.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "level_type_details.0.level_type", "NODE_TYPE_LEVEL"),
				resource.TestCheckResourceAttr(singularDatasourceName, "level_type_details.0.node_type", "UTILITY"),
				resource.TestCheckResourceAttr(singularDatasourceName, "metric_type", "INSTANCE_ACCESSIBILITY_STATUS"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
		{
			Config:            config + BdsBdsInstanceNodeReplaceConfigurationRequiredOnlyResource,
			ImportState:       true,
			ImportStateVerify: true,
			ImportStateIdFunc: getBdsReplaceConfigCompositeId(resourceName),
			ImportStateVerifyIgnore: []string{
				"cluster_admin_password",
			},
			ResourceName: resourceName,
		},
	})
}

func testAccCheckBdsBdsInstanceNodeReplaceConfigurationDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).BdsClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_bds_bds_instance_node_Replace_configuration" {
			noResourceFound = false
			request := oci_bds.GetNodeReplaceConfigurationRequest{}

			if value, ok := rs.Primary.Attributes["bds_instance_id"]; ok {
				request.BdsInstanceId = &value
			}

			if value, ok := rs.Primary.Attributes["node_Replace_configuration_id"]; ok {
				request.NodeReplaceConfigurationId = &value
			}

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "bds")

			response, err := client.GetNodeReplaceConfiguration(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_bds.NodeReplaceConfigurationLifecycleStateDeleted): true,
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
	if !acctest.InSweeperExcludeList("BdsBdsInstanceNodeReplaceConfiguration") {
		resource.AddTestSweepers("BdsBdsInstanceNodeReplaceConfiguration", &resource.Sweeper{
			Name:         "BdsBdsInstanceNodeReplaceConfiguration",
			Dependencies: acctest.DependencyGraph["bdsInstanceNodeReplaceConfiguration"],
			F:            sweepBdsBdsInstanceNodeReplaceConfigurationResource,
		})
	}
}

func sweepBdsBdsInstanceNodeReplaceConfigurationResource(compartment string) error {
	bdsClient := acctest.GetTestClients(&schema.ResourceData{}).BdsClient()
	bdsInstanceNodeReplaceConfigurationIds, err := getBdsBdsInstanceNodeReplaceConfigurationIds(compartment)
	if err != nil {
		return err
	}
	for _, bdsInstanceNodeReplaceConfigurationId := range bdsInstanceNodeReplaceConfigurationIds {
		if ok := acctest.SweeperDefaultResourceId[bdsInstanceNodeReplaceConfigurationId]; !ok {
			deleteNodeReplaceConfigurationRequest := oci_bds.RemoveNodeReplaceConfigurationRequest{}

			deleteNodeReplaceConfigurationRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "bds")
			_, error := bdsClient.RemoveNodeReplaceConfiguration(context.Background(), deleteNodeReplaceConfigurationRequest)
			if error != nil {
				fmt.Printf("Error deleting BdsInstanceNodeReplaceConfiguration %s %s, It is possible that the resource is already deleted. Please verify manually \n", bdsInstanceNodeReplaceConfigurationId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &bdsInstanceNodeReplaceConfigurationId, BdsBdsInstanceNodeReplaceConfigurationSweepWaitCondition, time.Duration(3*time.Minute),
				BdsBdsInstanceNodeReplaceConfigurationSweepResponseFetchOperation, "bds", true)
		}
	}
	return nil
}

func getBdsBdsInstanceNodeReplaceConfigurationIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "BdsInstanceNodeReplaceConfigurationId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	bdsClient := acctest.GetTestClients(&schema.ResourceData{}).BdsClient()

	listNodeReplaceConfigurationsRequest := oci_bds.ListNodeReplaceConfigurationsRequest{}

	bdsInstanceIds, error := getBdsInstanceIds(compartment)
	if error != nil {
		return resourceIds, fmt.Errorf("Error getting bdsInstanceId required for BdsInstanceNodeReplaceConfiguration resource requests \n")
	}
	for _, bdsInstanceId := range bdsInstanceIds {
		listNodeReplaceConfigurationsRequest.BdsInstanceId = &bdsInstanceId

		listNodeReplaceConfigurationsRequest.LifecycleState = oci_bds.NodeReplaceConfigurationLifecycleStateActive
		listNodeReplaceConfigurationsResponse, err := bdsClient.ListNodeReplaceConfigurations(context.Background(), listNodeReplaceConfigurationsRequest)

		if err != nil {
			return resourceIds, fmt.Errorf("Error getting BdsInstanceNodeReplaceConfiguration list for compartment id : %s , %s \n", compartmentId, err)
		}
		for _, bdsInstanceNodeReplaceConfiguration := range listNodeReplaceConfigurationsResponse.Items {
			id := *bdsInstanceNodeReplaceConfiguration.Id
			resourceIds = append(resourceIds, id)
			acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "BdsInstanceNodeReplaceConfigurationId", id)
		}

	}
	return resourceIds, nil
}

func BdsBdsInstanceNodeReplaceConfigurationSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if bdsInstanceNodeReplaceConfigurationResponse, ok := response.Response.(oci_bds.GetNodeReplaceConfigurationResponse); ok {
		return bdsInstanceNodeReplaceConfigurationResponse.LifecycleState != oci_bds.NodeReplaceConfigurationLifecycleStateDeleted
	}
	return false
}

func BdsBdsInstanceNodeReplaceConfigurationSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.BdsClient().GetNodeReplaceConfiguration(context.Background(), oci_bds.GetNodeReplaceConfigurationRequest{RequestMetadata: common.RequestMetadata{
		RetryPolicy: retryPolicy,
	},
	})
	return err
}

func getBdsReplaceConfigCompositeId(resourceName string) resource.ImportStateIdFunc {
	return func(s *terraform.State) (string, error) {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return "", fmt.Errorf("not found: %s", resourceName)
		}

		return fmt.Sprintf("bdsInstances/%s/nodeReplaceConfigurations/%s", rs.Primary.Attributes["bds_instance_id"], rs.Primary.Attributes["id"]), nil
	}
}
