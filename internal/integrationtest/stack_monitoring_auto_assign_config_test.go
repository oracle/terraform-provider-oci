// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/oracle/oci-go-sdk/v65/common"
	oci_stack_monitoring "github.com/oracle/oci-go-sdk/v65/stackmonitoring"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	StackMonitoringAutoAssignConfigRequiredOnlyResource = acctest.GenerateResourceFromRepresentationMap("oci_stack_monitoring_config", "test_auto_assign_config", acctest.Required, acctest.Create, StackMonitoringAutoAssignConfigRepresentation)

	StackMonitoringAutoAssignConfigResourceConfig = acctest.GenerateResourceFromRepresentationMap("oci_stack_monitoring_config", "test_auto_assign_config", acctest.Optional, acctest.Update, StackMonitoringAutoAssignConfigRepresentation)

	StackMonitoringAutoAssignConfigSingularDataSourceRepresentation = map[string]interface{}{
		"config_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_stack_monitoring_config.test_auto_assign_config.id}`},
	}

	StackMonitoringAutoAssignConfigDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":   acctest.Representation{RepType: acctest.Optional, Create: `TerraformAutoAssignConfig`, Update: `TerraformAutoAssignConfigUpdated`},
		"state":          acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"type":           acctest.Representation{RepType: acctest.Optional, Create: `LICENSE_AUTO_ASSIGN`},
		"filter":         acctest.RepresentationGroup{RepType: acctest.Required, Group: StackMonitoringAutoAssignConfigDataSourceFilterRepresentation}}
	StackMonitoringAutoAssignConfigDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_stack_monitoring_config.test_auto_assign_config.id}`}},
	}

	StackMonitoringAutoAssignConfigRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"config_type":    acctest.Representation{RepType: acctest.Required, Create: `LICENSE_AUTO_ASSIGN`},
		"license":        acctest.Representation{RepType: acctest.Required, Create: `STANDARD_EDITION`, Update: `ENTERPRISE_EDITION`},
		"display_name":   acctest.Representation{RepType: acctest.Optional, Create: `TerraformAutoAssignConfig`, Update: `TerraformAutoAssignConfigUpdated`},
		"freeform_tags":  acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"bar-key": "value"}, Update: map[string]string{"Department": "Accounting"}},
	}
)

// issue-routing-tag: stack_monitoring/default
func TestStackMonitoringAutoAssignConfigResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestStackMonitoringAutoAssignConfigResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_stack_monitoring_config.test_auto_assign_config"
	datasourceName := "data.oci_stack_monitoring_configs.test_auto_assign_configs"
	singularDatasourceName := "data.oci_stack_monitoring_config.test_auto_assign_config"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+
		acctest.GenerateResourceFromRepresentationMap("oci_stack_monitoring_config", "test_auto_assign_config", acctest.Optional, acctest.Create, StackMonitoringAutoAssignConfigRepresentation), "stackmonitoring", "config", t)

	acctest.ResourceTest(t, testAccCheckStackMonitoringAutoAssignConfigDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_stack_monitoring_config", "test_auto_assign_config", acctest.Required, acctest.Create, StackMonitoringAutoAssignConfigRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "config_type", "LICENSE_AUTO_ASSIGN"),
				resource.TestCheckResourceAttr(resourceName, "license", "STANDARD_EDITION"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_stack_monitoring_config", "test_auto_assign_config", acctest.Optional, acctest.Create, StackMonitoringAutoAssignConfigRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "config_type", "LICENSE_AUTO_ASSIGN"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "TerraformAutoAssignConfig"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "license", "STANDARD_EDITION"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),

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
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_stack_monitoring_config", "test_auto_assign_config", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(StackMonitoringAutoAssignConfigRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "config_type", "LICENSE_AUTO_ASSIGN"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "TerraformAutoAssignConfig"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "license", "STANDARD_EDITION"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),

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
			Config: config + compartmentIdVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_stack_monitoring_config", "test_auto_assign_config", acctest.Optional, acctest.Update, StackMonitoringAutoAssignConfigRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "config_type", "LICENSE_AUTO_ASSIGN"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "TerraformAutoAssignConfigUpdated"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "license", "ENTERPRISE_EDITION"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),

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
				acctest.GenerateDataSourceFromRepresentationMap("oci_stack_monitoring_configs", "test_auto_assign_configs", acctest.Optional, acctest.Update, StackMonitoringAutoAssignConfigDataSourceRepresentation) +
				compartmentIdVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_stack_monitoring_config", "test_auto_assign_config", acctest.Optional, acctest.Update, StackMonitoringAutoAssignConfigRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "TerraformAutoAssignConfigUpdated"),
				resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),
				resource.TestCheckResourceAttr(datasourceName, "type", "LICENSE_AUTO_ASSIGN"),

				resource.TestCheckResourceAttr(datasourceName, "config_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "config_collection.0.items.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_stack_monitoring_config", "test_auto_assign_config", acctest.Required, acctest.Create, StackMonitoringAutoAssignConfigSingularDataSourceRepresentation) +
				compartmentIdVariableStr + StackMonitoringAutoAssignConfigResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "config_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "config_type", "LICENSE_AUTO_ASSIGN"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "TerraformAutoAssignConfigUpdated"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "license", "ENTERPRISE_EDITION"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
		// verify resource import
		{
			Config:                  config + StackMonitoringAutoAssignConfigRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckStackMonitoringAutoAssignConfigDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).StackMonitoringClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_stack_monitoring_config" {
			noResourceFound = false
			request := oci_stack_monitoring.GetConfigRequest{}

			tmp := rs.Primary.ID
			request.ConfigId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "stack_monitoring")

			response, err := client.GetConfig(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_stack_monitoring.ConfigLifecycleStateDeleted): true,
				}
				if _, ok := deletedLifecycleStates[string(response.GetLifecycleState())]; !ok {
					//resource lifecycle state is not in expected deleted lifecycle states.
					return fmt.Errorf("resource lifecycle state: %s is not in expected deleted lifecycle states", response.GetLifecycleState())
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
	if !acctest.InSweeperExcludeList("StackMonitoringAutoAssignConfig") {
		resource.AddTestSweepers("StackMonitoringAutoAssignConfig", &resource.Sweeper{
			Name:         "StackMonitoringAutoAssignConfig",
			Dependencies: acctest.DependencyGraph["config"],
			F:            sweepStackMonitoringAutoAssignConfigResource,
		})
	}
}

func sweepStackMonitoringAutoAssignConfigResource(compartment string) error {
	stackMonitoringClient := acctest.GetTestClients(&schema.ResourceData{}).StackMonitoringClient()
	configIds, err := getStackMonitoringAutoAssignConfigIds(compartment)
	if err != nil {
		return err
	}
	for _, configId := range configIds {
		if ok := acctest.SweeperDefaultResourceId[configId]; !ok {
			deleteConfigRequest := oci_stack_monitoring.DeleteConfigRequest{}

			deleteConfigRequest.ConfigId = &configId

			deleteConfigRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "stack_monitoring")
			_, error := stackMonitoringClient.DeleteConfig(context.Background(), deleteConfigRequest)
			if error != nil {
				fmt.Printf("Error deleting Config %s %s, It is possible that the resource is already deleted. Please verify manually \n", configId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &configId, StackMonitoringAutoAssignConfigSweepWaitCondition, time.Duration(3*time.Minute),
				StackMonitoringAutoAssignConfigSweepResponseFetchOperation, "stack_monitoring", true)
		}
	}
	return nil
}

func getStackMonitoringAutoAssignConfigIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "ConfigId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	stackMonitoringClient := acctest.GetTestClients(&schema.ResourceData{}).StackMonitoringClient()

	listConfigsRequest := oci_stack_monitoring.ListConfigsRequest{}
	listConfigsRequest.CompartmentId = &compartmentId
	listConfigsRequest.LifecycleState = oci_stack_monitoring.ConfigLifecycleStateActive
	listConfigsResponse, err := stackMonitoringClient.ListConfigs(context.Background(), listConfigsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting Config list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, config := range listConfigsResponse.Items {
		id := *config.GetId()
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "ConfigId", id)
	}
	return resourceIds, nil
}

func StackMonitoringAutoAssignConfigSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if configResponse, ok := response.Response.(oci_stack_monitoring.GetConfigResponse); ok {
		return configResponse.GetLifecycleState() != oci_stack_monitoring.ConfigLifecycleStateDeleted
	}
	return false
}

func StackMonitoringAutoAssignConfigSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.StackMonitoringClient().GetConfig(context.Background(), oci_stack_monitoring.GetConfigRequest{
		ConfigId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
