// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	oci_bds "github.com/oracle/oci-go-sdk/v65/bds"
	"github.com/oracle/oci-go-sdk/v65/common"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	BdsBdsInstanceNodeBackupConfigurationRequiredOnlyResource = BdsBdsInstanceNodeBackupConfigurationResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_bds_bds_instance_node_backup_configuration", "test_bds_instance_node_backup_configuration", acctest.Required, acctest.Create, BdsBdsInstanceNodeBackupConfigurationRepresentation)

	BdsBdsInstanceNodeBackupConfigurationResourceConfig = BdsBdsInstanceNodeBackupConfigurationResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_bds_bds_instance_node_backup_configuration", "test_bds_instance_node_backup_configuration", acctest.Optional, acctest.Update, BdsBdsInstanceNodeBackupConfigurationRepresentation)

	BdsBdsInstanceNodeBackupConfigurationSingularDataSourceRepresentation = map[string]interface{}{
		"bds_instance_id":              acctest.Representation{RepType: acctest.Required, Create: `${oci_bds_bds_instance.test_bds_instance.id}`},
		"node_backup_configuration_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_bds_bds_instance_node_backup_configuration.test_bds_instance_node_backup_configuration.id}`},
	}

	BdsBdsInstanceNodeBackupConfigurationDataSourceRepresentation = map[string]interface{}{
		"bds_instance_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_bds_bds_instance.test_bds_instance.id}`},
	}

	BdsBdsInstanceNodeBackupConfigurationDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_bds_bds_instance_node_backup_configuration.test_bds_instance_node_backup_configuration.id}`}},
	}

	BdsBdsInstanceNodeBackupConfigurationRepresentation = map[string]interface{}{
		"bds_instance_id":             acctest.Representation{RepType: acctest.Required, Create: `${oci_bds_bds_instance.test_bds_instance.id}`},
		"level_type_details":          acctest.RepresentationGroup{RepType: acctest.Required, Group: BdsBdsInstanceNodeBackupConfigurationLevelTypeDetailsRepresentation},
		"schedule":                    acctest.Representation{RepType: acctest.Required, Create: `FREQ=WEEKLY;BYDAY=MO;BYHOUR=9`, Update: `FREQ=WEEKLY;BYDAY=FR;BYHOUR=18`},
		"backup_type":                 acctest.Representation{RepType: acctest.Optional, Create: `FULL`, Update: `INCREMENTAL`},
		"number_of_backups_to_retain": acctest.Representation{RepType: acctest.Optional, Create: `1`, Update: `2`},
	}
	BdsBdsInstanceNodeBackupConfigurationLevelTypeDetailsRepresentation = map[string]interface{}{
		"level_type": acctest.Representation{RepType: acctest.Required, Create: `NODE_TYPE_LEVEL`},
		"node_type":  acctest.Representation{RepType: acctest.Optional, Create: `MASTER`, Update: `UTILITY`},
	}

	BdsBdsInstanceNodeBackupConfigurationResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_bds_bds_instance", "test_bds_instance", acctest.Optional, acctest.Create, bdsInstanceOdhRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_subnet", "test_subnet", acctest.Required, acctest.Create, CoreSubnetRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", acctest.Required, acctest.Create, CoreVcnRepresentation)
)

// issue-routing-tag: bds/default
func TestBdsBdsInstanceNodeBackupConfigurationResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestBdsBdsInstanceNodeBackupConfigurationResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	subnetId := utils.GetEnvSettingWithBlankDefault("subnet_ocid")
	subnetIdVariableStr := fmt.Sprintf("variable \"subnet_id\" { default = \"%s\" }\n", subnetId)

	resourceName := "oci_bds_bds_instance_node_backup_configuration.test_bds_instance_node_backup_configuration"
	datasourceName := "data.oci_bds_bds_instance_node_backup_configurations.test_bds_instance_node_backup_configurations"
	singularDatasourceName := "data.oci_bds_bds_instance_node_backup_configuration.test_bds_instance_node_backup_configuration"

	var resId, resId2 string
	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + subnetIdVariableStr + BdsBdsInstanceNodeBackupConfigurationResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_bds_bds_instance_node_backup_configuration", "test_bds_instance_node_backup_configuration", acctest.Optional, acctest.Create, BdsBdsInstanceNodeBackupConfigurationRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "backup_type", "FULL"),
				resource.TestCheckResourceAttrSet(resourceName, "bds_instance_id"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "display_name"),
				resource.TestCheckResourceAttr(resourceName, "level_type_details.0.level_type", "NODE_TYPE_LEVEL"),
				resource.TestCheckResourceAttr(resourceName, "level_type_details.0.node_type", "MASTER"),
				resource.TestCheckResourceAttr(resourceName, "schedule", "FREQ=WEEKLY;BYDAY=MO;BYHOUR=9"),
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
			Config: config + compartmentIdVariableStr + subnetIdVariableStr + BdsBdsInstanceNodeBackupConfigurationResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_bds_bds_instance_node_backup_configuration", "test_bds_instance_node_backup_configuration", acctest.Optional, acctest.Update, BdsBdsInstanceNodeBackupConfigurationRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "backup_type", "INCREMENTAL"),
				resource.TestCheckResourceAttrSet(resourceName, "bds_instance_id"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "display_name"),
				resource.TestCheckResourceAttr(resourceName, "level_type_details.0.node_type", "UTILITY"),
				resource.TestCheckResourceAttr(resourceName, "schedule", "FREQ=WEEKLY;BYDAY=FR;BYHOUR=18"),
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
			Config: config + acctest.GenerateResourceFromRepresentationMap("oci_bds_bds_instance_node_backup_configuration", "test_bds_instance_node_backup_configuration", acctest.Optional, acctest.Update, BdsBdsInstanceNodeBackupConfigurationRepresentation) +
				acctest.GenerateDataSourceFromRepresentationMap("oci_bds_bds_instance_node_backup_configurations", "test_bds_instance_node_backup_configurations", acctest.Optional, acctest.Update, BdsBdsInstanceNodeBackupConfigurationDataSourceRepresentation) +
				compartmentIdVariableStr + subnetIdVariableStr + BdsBdsInstanceNodeBackupConfigurationResourceDependencies,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "node_backup_configurations.0.bds_instance_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "node_backup_configurations.0.id"),
				resource.TestCheckResourceAttr(datasourceName, "node_backup_configurations.0.level_type_details.0.level_type", "NODE_TYPE_LEVEL"),
				resource.TestCheckResourceAttr(datasourceName, "node_backup_configurations.0.level_type_details.0.node_type", "UTILITY"),
				resource.TestCheckResourceAttrSet(datasourceName, "node_backup_configurations.0.time_created"),
				resource.TestCheckResourceAttrSet(datasourceName, "node_backup_configurations.0.time_updated"),
			),
		},
		// verify singular datasource
		{
			Config: config + acctest.GenerateResourceFromRepresentationMap("oci_bds_bds_instance_node_backup_configuration", "test_bds_instance_node_backup_configuration", acctest.Optional, acctest.Update, BdsBdsInstanceNodeBackupConfigurationRepresentation) +
				acctest.GenerateDataSourceFromRepresentationMap("oci_bds_bds_instance_node_backup_configuration", "test_bds_instance_node_backup_configuration", acctest.Required, acctest.Create, BdsBdsInstanceNodeBackupConfigurationSingularDataSourceRepresentation) +
				compartmentIdVariableStr + subnetIdVariableStr + BdsBdsInstanceNodeBackupConfigurationResourceDependencies,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "bds_instance_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "node_backup_configuration_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "backup_type", "INCREMENTAL"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "level_type_details.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "level_type_details.0.level_type", "NODE_TYPE_LEVEL"),
				resource.TestCheckResourceAttr(singularDatasourceName, "level_type_details.0.node_type", "UTILITY"),
				resource.TestCheckResourceAttr(singularDatasourceName, "number_of_backups_to_retain", "2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "schedule", "FREQ=WEEKLY;BYDAY=FR;BYHOUR=18"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
				resource.TestCheckResourceAttr(singularDatasourceName, "timezone", "Etc/UTC"),
			),
		},
		// verify resource import
		{
			Config:            config + BdsBdsInstanceNodeBackupConfigurationRequiredOnlyResource,
			ImportState:       true,
			ImportStateVerify: true,
			ImportStateIdFunc: getBdsBackupConfigCompositeId(resourceName),
			ImportStateVerifyIgnore: []string{
				"cluster_admin_password",
			},
			ResourceName: resourceName,
		},
	})
}

func testAccCheckBdsBdsInstanceNodeBackupConfigurationDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).BdsClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_bds_bds_instance_node_backup_configuration" {
			noResourceFound = false
			request := oci_bds.GetNodeBackupConfigurationRequest{}

			if value, ok := rs.Primary.Attributes["bds_instance_id"]; ok {
				request.BdsInstanceId = &value
			}

			if value, ok := rs.Primary.Attributes["node_backup_configuration_id"]; ok {
				request.NodeBackupConfigurationId = &value
			}

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "bds")

			response, err := client.GetNodeBackupConfiguration(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_bds.NodeBackupConfigurationLifecycleStateDeleted): true,
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
	if !acctest.InSweeperExcludeList("BdsBdsInstanceNodeBackupConfiguration") {
		resource.AddTestSweepers("BdsBdsInstanceNodeBackupConfiguration", &resource.Sweeper{
			Name:         "BdsBdsInstanceNodeBackupConfiguration",
			Dependencies: acctest.DependencyGraph["bdsInstanceNodeBackupConfiguration"],
			F:            sweepBdsBdsInstanceNodeBackupConfigurationResource,
		})
	}
}

func sweepBdsBdsInstanceNodeBackupConfigurationResource(compartment string) error {
	bdsClient := acctest.GetTestClients(&schema.ResourceData{}).BdsClient()
	bdsInstanceNodeBackupConfigurationIds, err := getBdsBdsInstanceNodeBackupConfigurationIds(compartment)
	if err != nil {
		return err
	}
	for _, bdsInstanceNodeBackupConfigurationId := range bdsInstanceNodeBackupConfigurationIds {
		if ok := acctest.SweeperDefaultResourceId[bdsInstanceNodeBackupConfigurationId]; !ok {
			deleteNodeBackupConfigurationRequest := oci_bds.DeleteNodeBackupConfigurationRequest{}

			deleteNodeBackupConfigurationRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "bds")
			_, error := bdsClient.DeleteNodeBackupConfiguration(context.Background(), deleteNodeBackupConfigurationRequest)
			if error != nil {
				fmt.Printf("Error deleting BdsInstanceNodeBackupConfiguration %s %s, It is possible that the resource is already deleted. Please verify manually \n", bdsInstanceNodeBackupConfigurationId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &bdsInstanceNodeBackupConfigurationId, BdsBdsInstanceNodeBackupConfigurationSweepWaitCondition, time.Duration(3*time.Minute),
				BdsBdsInstanceNodeBackupConfigurationSweepResponseFetchOperation, "bds", true)
		}
	}
	return nil
}

func getBdsBdsInstanceNodeBackupConfigurationIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "BdsInstanceNodeBackupConfigurationId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	bdsClient := acctest.GetTestClients(&schema.ResourceData{}).BdsClient()

	listNodeBackupConfigurationsRequest := oci_bds.ListNodeBackupConfigurationsRequest{}
	//listNodeBackupConfigurationsRequest.CompartmentId = &compartmentId

	bdsInstanceIds, error := getBdsInstanceIds(compartment)
	if error != nil {
		return resourceIds, fmt.Errorf("Error getting bdsInstanceId required for BdsInstanceNodeBackupConfiguration resource requests \n")
	}
	for _, bdsInstanceId := range bdsInstanceIds {
		listNodeBackupConfigurationsRequest.BdsInstanceId = &bdsInstanceId

		listNodeBackupConfigurationsRequest.LifecycleState = oci_bds.NodeBackupConfigurationLifecycleStateActive
		listNodeBackupConfigurationsResponse, err := bdsClient.ListNodeBackupConfigurations(context.Background(), listNodeBackupConfigurationsRequest)

		if err != nil {
			return resourceIds, fmt.Errorf("Error getting BdsInstanceNodeBackupConfiguration list for compartment id : %s , %s \n", compartmentId, err)
		}
		for _, bdsInstanceNodeBackupConfiguration := range listNodeBackupConfigurationsResponse.Items {
			id := *bdsInstanceNodeBackupConfiguration.Id
			resourceIds = append(resourceIds, id)
			acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "BdsInstanceNodeBackupConfigurationId", id)
		}

	}
	return resourceIds, nil
}

func BdsBdsInstanceNodeBackupConfigurationSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if bdsInstanceNodeBackupConfigurationResponse, ok := response.Response.(oci_bds.GetNodeBackupConfigurationResponse); ok {
		return bdsInstanceNodeBackupConfigurationResponse.LifecycleState != oci_bds.NodeBackupConfigurationLifecycleStateDeleted
	}
	return false
}

func BdsBdsInstanceNodeBackupConfigurationSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.BdsClient().GetNodeBackupConfiguration(context.Background(), oci_bds.GetNodeBackupConfigurationRequest{RequestMetadata: common.RequestMetadata{
		RetryPolicy: retryPolicy,
	},
	})
	return err
}

func getBdsBackupConfigCompositeId(resourceName string) resource.ImportStateIdFunc {
	return func(s *terraform.State) (string, error) {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return "", fmt.Errorf("not found: %s", resourceName)
		}

		return fmt.Sprintf("bdsInstances/%s/nodeBackupConfigurations/%s", rs.Primary.Attributes["bds_instance_id"], rs.Primary.Attributes["id"]), nil
	}
}
