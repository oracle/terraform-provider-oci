// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
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
	oci_fleet_apps_management "github.com/oracle/oci-go-sdk/v65/fleetappsmanagement"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	FleetAppsManagementFleetCredentialResourceConfig = FleetAppsManagementFleetCredentialResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_fleet_apps_management_fleet_credential", "test_fleet_credential", acctest.Optional, acctest.Update, FleetAppsManagementFleetCredentialRepresentation)

	FleetAppsManagementFleetCredentialSingularDataSourceRepresentation = map[string]interface{}{
		"fleet_credential_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_fleet_apps_management_fleet_credential.test_fleet_credential.id}`},
		"fleet_id":            acctest.Representation{RepType: acctest.Required, Create: `${var.test_active_fleet}`},
	}

	FleetAppsManagementFleetCredentialDataSourceRepresentation = map[string]interface{}{
		"fleet_id":         acctest.Representation{RepType: acctest.Required, Create: `${oci_fleet_apps_management_fleet.test_fleet.id}`},
		"compartment_id":   acctest.Representation{RepType: acctest.Optional, Create: `${var.tenancy_ocid}`},
		"credential_level": acctest.Representation{RepType: acctest.Optional, Create: `TARGET`},
		"display_name":     acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"state":            acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"filter":           acctest.RepresentationGroup{RepType: acctest.Required, Group: FleetAppsManagementFleetCredentialDataSourceFilterRepresentation}}
	FleetAppsManagementFleetCredentialDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_fleet_apps_management_fleet_credential.test_fleet_credential.id}`}},
	}

	FleetAppsManagementFleetCredentialRepresentation = map[string]interface{}{
		"lifecycle":        acctest.RepresentationGroup{RepType: acctest.Required, Group: fleetCredentialIgnoreChangesRecipeRepresentation},
		"depends_on":       acctest.Representation{RepType: acctest.Required, Create: []string{`oci_fleet_apps_management_fleet_resource.test_fleet_resource`}},
		"compartment_id":   acctest.Representation{RepType: acctest.Required, Create: `${var.tenancy_ocid}`},
		"display_name":     acctest.Representation{RepType: acctest.Required, Create: `displayName`, Update: `displayName2`},
		"entity_specifics": acctest.RepresentationGroup{RepType: acctest.Required, Group: FleetAppsManagementFleetCredentialEntitySpecificsRepresentation},
		"fleet_id":         acctest.Representation{RepType: acctest.Required, Create: `${oci_fleet_apps_management_fleet.test_fleet.id}`},
		"password":         acctest.RepresentationGroup{RepType: acctest.Required, Group: FleetAppsManagementFleetCredentialPasswordRepresentation},
		"user":             acctest.RepresentationGroup{RepType: acctest.Required, Group: FleetAppsManagementFleetCredentialUserRepresentation},
	}
	FleetAppsManagementFleetCredentialEntitySpecificsRepresentation = map[string]interface{}{
		"credential_level": acctest.Representation{RepType: acctest.Required, Create: `TARGET`},
		"resource_id":      acctest.Representation{RepType: acctest.Required, Create: `${oci_core_instance.test_instance.id}`},
		"target":           acctest.Representation{RepType: acctest.Required, Create: `/home/oracle/Oracle/Middleware/Oracle_Home/wlserver`, Update: `/home/oracle/Oracle/Middleware/Oracle_Home/otherserver`},
	}
	FleetAppsManagementFleetCredentialPasswordRepresentation = map[string]interface{}{
		"secret_id":       acctest.Representation{RepType: acctest.Required, Create: `somePasswordSecretId`},
		"secret_version":  acctest.Representation{RepType: acctest.Required, Create: `1`, Update: `2`},
		"credential_type": acctest.Representation{RepType: acctest.Required, Create: `VAULT_SECRET`},
	}
	FleetAppsManagementFleetCredentialUserRepresentation = map[string]interface{}{
		"secret_id":       acctest.Representation{RepType: acctest.Required, Create: `someUserSecretId`},
		"secret_version":  acctest.Representation{RepType: acctest.Required, Create: `1`, Update: `2`},
		"credential_type": acctest.Representation{RepType: acctest.Required, Create: `VAULT_SECRET`},
	}

	FleetAppsManagementFleetCredentialResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_core_subnet", "test_subnet", acctest.Required, acctest.Create, CoreSubnetRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", acctest.Required, acctest.Create, CoreVcnRepresentation) +
		utils.OciImageIdsVariable +
		acctest.GenerateResourceFromRepresentationMap("oci_core_instance", "test_instance", acctest.Required, acctest.Create, CoreInstanceRepresentation) +
		AvailabilityDomainConfig +
		acctest.GenerateResourceFromRepresentationMap("oci_fleet_apps_management_fleet", "test_fleet", acctest.Required, acctest.Create, FleetAppsManagementFleetRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_fleet_apps_management_fleet_resource", "test_fleet_resource", acctest.Required, acctest.Create, FleetAppsManagementFleetResourceRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_ons_notification_topic", "test_notification_topic", acctest.Required, acctest.Create, OnsNotificationTopicRepresentation)

	fleetCredentialIgnoreChangesRecipeRepresentation = map[string]interface{}{
		"ignore_changes": acctest.Representation{RepType: acctest.Required, Create: []string{`system_tags`}},
	}
)

// issue-routing-tag: fleet_apps_management/default
func TestFleetAppsManagementFleetCredentialResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestFleetAppsManagementFleetCredentialResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("tenancy_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	// Fleet in ACTIVE state. Fleets require a confirmation action call not supported by Terraform to go active.
	// Thus, this needs to be created and confirmed manually.
	activeFleetId := utils.GetEnvSettingWithBlankDefault("test_active_fleet")
	activeFleetStr := fmt.Sprintf("variable \"test_active_fleet\" { default = \"%s\" }\n", activeFleetId)

	resourceName := "oci_fleet_apps_management_fleet_credential.test_fleet_credential"
	datasourceName := "data.oci_fleet_apps_management_fleet_credentials.test_fleet_credentials"
	singularDatasourceName := "data.oci_fleet_apps_management_fleet_credential.test_fleet_credential"

	var resId, resId2 string
	// Save TF content to Create resource with only required properties. This has to be exactly the same as the config part in the create step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+FleetAppsManagementFleetCredentialResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_fleet_apps_management_fleet_credential", "test_fleet_credential", acctest.Required, acctest.Create, FleetAppsManagementFleetCredentialRepresentation), "fleetappsmanagement", "fleetCredential", t)

	acctest.ResourceTest(t, testAccCheckFleetAppsManagementFleetCredentialDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + activeFleetStr + compartmentIdVariableStr + FleetAppsManagementFleetCredentialResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_fleet_apps_management_fleet_credential", "test_fleet_credential", acctest.Required, acctest.Create, FleetAppsManagementFleetCredentialRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "entity_specifics.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "entity_specifics.0.credential_level", "TARGET"),
				resource.TestCheckResourceAttrSet(resourceName, "entity_specifics.0.resource_id"),
				resource.TestCheckResourceAttr(resourceName, "entity_specifics.0.target", "/home/oracle/Oracle/Middleware/Oracle_Home/wlserver"),
				resource.TestCheckResourceAttrSet(resourceName, "fleet_id"),
				resource.TestCheckResourceAttr(resourceName, "password.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "password.0.credential_type", "VAULT_SECRET"),
				resource.TestCheckResourceAttr(resourceName, "user.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "user.0.credential_type", "VAULT_SECRET"),

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
			Config: config + activeFleetStr + compartmentIdVariableStr + FleetAppsManagementFleetCredentialResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_fleet_apps_management_fleet_credential", "test_fleet_credential", acctest.Optional, acctest.Update, FleetAppsManagementFleetCredentialRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "entity_specifics.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "entity_specifics.0.credential_level", "TARGET"),
				resource.TestCheckResourceAttrSet(resourceName, "entity_specifics.0.resource_id"),
				resource.TestCheckResourceAttr(resourceName, "entity_specifics.0.target", "/home/oracle/Oracle/Middleware/Oracle_Home/otherserver"),
				resource.TestCheckResourceAttrSet(resourceName, "fleet_id"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "password.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "password.0.credential_type", "VAULT_SECRET"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttr(resourceName, "user.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "user.0.credential_type", "VAULT_SECRET"),

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
				acctest.GenerateDataSourceFromRepresentationMap("oci_fleet_apps_management_fleet_credentials", "test_fleet_credentials", acctest.Optional, acctest.Update, FleetAppsManagementFleetCredentialDataSourceRepresentation) +
				activeFleetStr + compartmentIdVariableStr + FleetAppsManagementFleetCredentialResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_fleet_apps_management_fleet_credential", "test_fleet_credential", acctest.Optional, acctest.Update, FleetAppsManagementFleetCredentialRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "credential_level", "TARGET"),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttrSet(datasourceName, "fleet_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "id"),
				resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),

				resource.TestCheckResourceAttr(datasourceName, "fleet_credential_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "fleet_credential_collection.0.items.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_fleet_apps_management_fleet_credential", "test_fleet_credential", acctest.Required, acctest.Create, FleetAppsManagementFleetCredentialSingularDataSourceRepresentation) +
				activeFleetStr + compartmentIdVariableStr + FleetAppsManagementFleetCredentialResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "fleet_credential_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "fleet_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "entity_specifics.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "entity_specifics.0.credential_level", "TARGET"),
				resource.TestCheckResourceAttr(singularDatasourceName, "entity_specifics.0.target", "/home/oracle/Oracle/Middleware/Oracle_Home/otherserver"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "password.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "password.0.credential_type", "VAULT_SECRET"),
				resource.TestCheckResourceAttr(singularDatasourceName, "password.0.secret_version", "2"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
				resource.TestCheckResourceAttr(singularDatasourceName, "user.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "user.0.credential_type", "VAULT_SECRET"),
				resource.TestCheckResourceAttr(singularDatasourceName, "user.0.secret_version", "2"),
			),
		},
	})
}

func testAccCheckFleetAppsManagementFleetCredentialDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).FleetAppsManagementClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_fleet_apps_management_fleet_credential" {
			noResourceFound = false
			request := oci_fleet_apps_management.GetFleetCredentialRequest{}

			tmp := rs.Primary.ID
			request.FleetCredentialId = &tmp

			if value, ok := rs.Primary.Attributes["fleet_id"]; ok {
				request.FleetId = &value
			}

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "fleet_apps_management")

			response, err := client.GetFleetCredential(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_fleet_apps_management.FleetCredentialLifecycleStateDeleted): true,
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
	if !acctest.InSweeperExcludeList("FleetAppsManagementFleetCredential") {
		resource.AddTestSweepers("FleetAppsManagementFleetCredential", &resource.Sweeper{
			Name:         "FleetAppsManagementFleetCredential",
			Dependencies: acctest.DependencyGraph["fleetCredential"],
			F:            sweepFleetAppsManagementFleetCredentialResource,
		})
	}
}

func sweepFleetAppsManagementFleetCredentialResource(compartment string) error {
	fleetAppsManagementClient := acctest.GetTestClients(&schema.ResourceData{}).FleetAppsManagementClient()
	fleetCredentialIds, err := getFleetAppsManagementFleetCredentialIds(compartment)
	if err != nil {
		return err
	}
	for _, fleetCredentialId := range fleetCredentialIds {
		if ok := acctest.SweeperDefaultResourceId[fleetCredentialId]; !ok {
			deleteFleetCredentialRequest := oci_fleet_apps_management.DeleteFleetCredentialRequest{}

			deleteFleetCredentialRequest.FleetCredentialId = &fleetCredentialId

			deleteFleetCredentialRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "fleet_apps_management")
			_, error := fleetAppsManagementClient.DeleteFleetCredential(context.Background(), deleteFleetCredentialRequest)
			if error != nil {
				fmt.Printf("Error deleting FleetCredential %s %s, It is possible that the resource is already deleted. Please verify manually \n", fleetCredentialId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &fleetCredentialId, FleetAppsManagementFleetCredentialSweepWaitCondition, time.Duration(3*time.Minute),
				FleetAppsManagementFleetCredentialSweepResponseFetchOperation, "fleet_apps_management", true)
		}
	}
	return nil
}

func getFleetAppsManagementFleetCredentialIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "FleetCredentialId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	fleetAppsManagementClient := acctest.GetTestClients(&schema.ResourceData{}).FleetAppsManagementClient()

	listFleetCredentialsRequest := oci_fleet_apps_management.ListFleetCredentialsRequest{}
	listFleetCredentialsRequest.CompartmentId = &compartmentId

	fleetIds, error := getFleetAppsManagementFleetIds(compartment)
	if error != nil {
		return resourceIds, fmt.Errorf("Error getting fleetId required for FleetCredential resource requests \n")
	}
	for _, fleetId := range fleetIds {
		listFleetCredentialsRequest.FleetId = &fleetId

		listFleetCredentialsRequest.LifecycleState = oci_fleet_apps_management.FleetCredentialLifecycleStateActive
		listFleetCredentialsResponse, err := fleetAppsManagementClient.ListFleetCredentials(context.Background(), listFleetCredentialsRequest)

		if err != nil {
			return resourceIds, fmt.Errorf("Error getting FleetCredential list for compartment id : %s , %s \n", compartmentId, err)
		}
		for _, fleetCredential := range listFleetCredentialsResponse.Items {
			id := *fleetCredential.Id
			resourceIds = append(resourceIds, id)
			acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "FleetCredentialId", id)
		}

	}
	return resourceIds, nil
}

func FleetAppsManagementFleetCredentialSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if fleetCredentialResponse, ok := response.Response.(oci_fleet_apps_management.GetFleetCredentialResponse); ok {
		return fleetCredentialResponse.LifecycleState != oci_fleet_apps_management.FleetCredentialLifecycleStateDeleted
	}
	return false
}

func FleetAppsManagementFleetCredentialSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.FleetAppsManagementClient().GetFleetCredential(context.Background(), oci_fleet_apps_management.GetFleetCredentialRequest{
		FleetCredentialId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
