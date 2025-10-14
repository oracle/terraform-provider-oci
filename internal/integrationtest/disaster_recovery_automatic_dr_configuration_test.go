// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"regexp"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/oracle/oci-go-sdk/v65/common"
	oci_disaster_recovery "github.com/oracle/oci-go-sdk/v65/disasterrecovery"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	DisasterRecoveryAutomaticDrConfigurationRequiredOnlyResource = DisasterRecoveryAutomaticDrConfigurationResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_disaster_recovery_automatic_dr_configuration", "test_automatic_dr_configuration", acctest.Required, acctest.Create, DisasterRecoveryAutomaticDrConfigurationRepresentation)

	DisasterRecoveryAutomaticDrConfigurationResourceConfig = DisasterRecoveryAutomaticDrConfigurationResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_disaster_recovery_automatic_dr_configuration", "test_automatic_dr_configuration", acctest.Optional, acctest.Create, DisasterRecoveryAutomaticDrConfigurationRepresentation)

	DisasterRecoveryDisasterRecoveryAutomaticDrConfigurationSingularDataSourceRepresentation = map[string]interface{}{
		"automatic_dr_configuration_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_disaster_recovery_automatic_dr_configuration.test_automatic_dr_configuration.id}`},
	}

	//Automatic DR Configuration Data source
	DisasterRecoveryDisasterRecoveryAutomaticDrConfigurationDataSourceRepresentation = map[string]interface{}{
		"dr_protection_group_id": acctest.Representation{RepType: acctest.Required, Create: `${data.oci_disaster_recovery_dr_protection_groups.test_drpg.dr_protection_group_collection.0.items.0.id}`},
		"display_name":           acctest.Representation{RepType: acctest.Optional, Create: `My Automatic DR Configuration`},
		"state":                  acctest.Representation{RepType: acctest.Optional, Create: `FAILED`},
		"filter":                 acctest.RepresentationGroup{RepType: acctest.Required, Group: DisasterRecoveryAutomaticDrConfigurationDataSourceFilterRepresentation}}
	DisasterRecoveryAutomaticDrConfigurationDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_disaster_recovery_automatic_dr_configuration.test_automatic_dr_configuration.id}`}},
	}

	DisasterRecoveryAutomaticDrConfigurationRepresentation = map[string]interface{}{
		"display_name":                  acctest.Representation{RepType: acctest.Required, Create: `My Automatic DR Configuration`},
		"dr_protection_group_id":        acctest.Representation{RepType: acctest.Required, Create: `${data.oci_disaster_recovery_dr_protection_groups.test_drpg.dr_protection_group_collection.0.items.0.id}`},
		"members":                       acctest.RepresentationGroup{RepType: acctest.Required, Group: DisasterRecoveryAutomaticDrConfigurationMembersRepresentation},
		"default_switchover_dr_plan_id": acctest.Representation{RepType: acctest.Required, Create: `${data.oci_disaster_recovery_dr_plans.test_switchover_plan.dr_plan_collection.0.items.0.id}`},
		"default_failover_dr_plan_id":   acctest.Representation{RepType: acctest.Required, Create: `${data.oci_disaster_recovery_dr_plans.test_failover_plan.dr_plan_collection.0.items.0.id}`},
		"freeform_tags":                 acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}},
	}

	DisasterRecoveryAutomaticDrConfigurationMembersRepresentation = map[string]interface{}{
		"member_id":                  acctest.Representation{RepType: acctest.Required, Create: `${data.oci_database_autonomous_databases.test_adb.autonomous_databases.0.id}`},
		"member_type":                acctest.Representation{RepType: acctest.Required, Create: `AUTONOMOUS_DATABASE`},
		"is_auto_failover_enabled":   acctest.Representation{RepType: acctest.Required, Create: `true`},
		"is_auto_switchover_enabled": acctest.Representation{RepType: acctest.Required, Create: `true`},
	}

	// Data source configurations for pre-existing resources
	DrProtectionGroupConfiguration = `
	data "oci_disaster_recovery_dr_protection_groups" "test_drpg" {
		compartment_id = var.compartment_id
		display_name   = "terraform-standby-drpg-iad"
	}
	`

	DefaultSwitchoverPlanConfiguration = `
	data "oci_disaster_recovery_dr_plans" "test_switchover_plan" {
		dr_protection_group_id = data.oci_disaster_recovery_dr_protection_groups.test_drpg.dr_protection_group_collection.0.items.0.id
		display_name           = "terraform-default-switchover"
	}
	`

	DefaultFailoverPlanConfiguration = `
	data "oci_disaster_recovery_dr_plans" "test_failover_plan" {
		dr_protection_group_id = data.oci_disaster_recovery_dr_protection_groups.test_drpg.dr_protection_group_collection.0.items.0.id
		display_name           = "terraform-default-failover"
	}
	`

	AutonomousDatabaseConfiguration = `
	data "oci_database_autonomous_databases" "test_adb" {
		compartment_id = var.compartment_id
		display_name   = "fsdradbs01"
	}
	`

	DisasterRecoveryAutomaticDrConfigurationResourceDependencies = DrProtectionGroupConfiguration +
		DefaultSwitchoverPlanConfiguration +
		DefaultFailoverPlanConfiguration +
		AutonomousDatabaseConfiguration
)

// issue-routing-tag: disaster_recovery/default
func TestDisasterRecoveryAutomaticDrConfigurationResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDisasterRecoveryAutomaticDrConfigurationResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", utils.GetEnvSettingWithBlankDefault("compartment_ocid"))

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// Load pre-existing resources via data sources
		{
			Config: config + compartmentIdVariableStr + DisasterRecoveryAutomaticDrConfigurationResourceDependencies,
		},
		{
			Config: config + compartmentIdVariableStr + DisasterRecoveryAutomaticDrConfigurationResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_disaster_recovery_automatic_dr_configuration", "test_automatic_dr_configuration", acctest.Optional, acctest.Create, DisasterRecoveryAutomaticDrConfigurationRepresentation),
			ExpectError: regexp.MustCompile("Work Request error"),
		},
	})
}

func testAccCheckDisasterRecoveryAutomaticDrConfigurationDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).DisasterRecoveryClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_disaster_recovery_automatic_dr_configuration" {
			noResourceFound = false
			request := oci_disaster_recovery.GetAutomaticDrConfigurationRequest{}

			tmp := rs.Primary.ID
			request.AutomaticDrConfigurationId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "disaster_recovery")

			response, err := client.GetAutomaticDrConfiguration(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_disaster_recovery.AutomaticDrConfigurationLifecycleStateDeleted):  true,
					string(oci_disaster_recovery.AutomaticDrConfigurationLifecycleStateDeleting): true,
					string(oci_disaster_recovery.AutomaticDrConfigurationLifecycleStateFailed):   true,
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
	if !acctest.InSweeperExcludeList("DisasterRecoveryAutomaticDrConfiguration") {
		resource.AddTestSweepers("DisasterRecoveryAutomaticDrConfiguration", &resource.Sweeper{
			Name:         "DisasterRecoveryAutomaticDrConfiguration",
			Dependencies: acctest.DependencyGraph["automaticDrConfiguration"],
			F:            sweepDisasterRecoveryAutomaticDrConfigurationResource,
		})
	}
}

func sweepDisasterRecoveryAutomaticDrConfigurationResource(compartment string) error {
	disasterRecoveryClient := acctest.GetTestClients(&schema.ResourceData{}).DisasterRecoveryClient()
	automaticDrConfigurationIds, err := getDisasterRecoveryAutomaticDrConfigurationIds(compartment)
	if err != nil {
		return err
	}
	for _, automaticDrConfigurationId := range automaticDrConfigurationIds {
		if ok := acctest.SweeperDefaultResourceId[automaticDrConfigurationId]; !ok {
			deleteAutomaticDrConfigurationRequest := oci_disaster_recovery.DeleteAutomaticDrConfigurationRequest{}

			deleteAutomaticDrConfigurationRequest.AutomaticDrConfigurationId = &automaticDrConfigurationId

			deleteAutomaticDrConfigurationRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "disaster_recovery")
			_, error := disasterRecoveryClient.DeleteAutomaticDrConfiguration(context.Background(), deleteAutomaticDrConfigurationRequest)
			if error != nil {
				fmt.Printf("Error deleting AutomaticDrConfiguration %s %s, It is possible that the resource is already deleted. Please verify manually \n", automaticDrConfigurationId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &automaticDrConfigurationId, DisasterRecoveryAutomaticDrConfigurationSweepWaitCondition, time.Duration(3*time.Minute),
				DisasterRecoveryAutomaticDrConfigurationSweepResponseFetchOperation, "disaster_recovery", true)
		}
	}
	return nil
}

func getDisasterRecoveryAutomaticDrConfigurationIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "AutomaticDrConfigurationId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	disasterRecoveryClient := acctest.GetTestClients(&schema.ResourceData{}).DisasterRecoveryClient()

	listAutomaticDrConfigurationsRequest := oci_disaster_recovery.ListAutomaticDrConfigurationsRequest{}
	//listAutomaticDrConfigurationsRequest.CompartmentId = &compartmentId

	drProtectionGroupIds, error := getDisasterRecoveryDrProtectionGroupIds(compartment)
	if error != nil {
		return resourceIds, fmt.Errorf("Error getting drProtectionGroupId required for AutomaticDrConfiguration resource requests \n")
	}
	for _, drProtectionGroupId := range drProtectionGroupIds {
		listAutomaticDrConfigurationsRequest.DrProtectionGroupId = &drProtectionGroupId

		listAutomaticDrConfigurationsRequest.LifecycleState = oci_disaster_recovery.ListAutomaticDrConfigurationsLifecycleStateActive
		listAutomaticDrConfigurationsResponse, err := disasterRecoveryClient.ListAutomaticDrConfigurations(context.Background(), listAutomaticDrConfigurationsRequest)

		if err != nil {
			return resourceIds, fmt.Errorf("Error getting AutomaticDrConfiguration list for compartment id : %s , %s \n", compartmentId, err)
		}
		for _, automaticDrConfiguration := range listAutomaticDrConfigurationsResponse.Items {
			id := *automaticDrConfiguration.Id
			resourceIds = append(resourceIds, id)
			acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "AutomaticDrConfigurationId", id)
			acctest.SweeperDefaultResourceId[*automaticDrConfiguration.DefaultFailoverDrPlanId] = true
			acctest.SweeperDefaultResourceId[*automaticDrConfiguration.DefaultSwitchoverDrPlanId] = true

		}

	}
	return resourceIds, nil
}

func DisasterRecoveryAutomaticDrConfigurationSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if automaticDrConfigurationResponse, ok := response.Response.(oci_disaster_recovery.GetAutomaticDrConfigurationResponse); ok {
		return automaticDrConfigurationResponse.LifecycleState != oci_disaster_recovery.AutomaticDrConfigurationLifecycleStateDeleted
	}
	return false
}

func DisasterRecoveryAutomaticDrConfigurationSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.DisasterRecoveryClient().GetAutomaticDrConfiguration(context.Background(), oci_disaster_recovery.GetAutomaticDrConfigurationRequest{
		AutomaticDrConfigurationId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
