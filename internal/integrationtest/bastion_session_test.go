// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"log"
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
	oci_bastion "github.com/oracle/oci-go-sdk/v65/bastion"
	"github.com/oracle/oci-go-sdk/v65/common"

	"github.com/oracle/terraform-provider-oci/httpreplay"
)

var (
	BastionSessionRequiredOnlyResource = BastionSessionResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_bastion_session", "test_session", acctest.Required, acctest.Create, BastionsessionRepresentation)

	BastionSessionResourceConfig = BastionSessionResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_bastion_session", "test_session", acctest.Optional, acctest.Update, BastionsessionRepresentation)

	BastionBastionsessionSingularDataSourceRepresentation = map[string]interface{}{
		"session_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_bastion_session.test_session.id}`},
	}

	sessionDataSourceRepresentation = map[string]interface{}{
		"bastion_id":              acctest.Representation{RepType: acctest.Required, Create: `${oci_bastion_bastion.test_bastion.id}`},
		"display_name":            acctest.Representation{RepType: acctest.Optional, Create: `managed_ssh`, Update: `managed_ssh2`},
		"session_id":              acctest.Representation{RepType: acctest.Optional, Create: `${oci_bastion_session.test_session.id}`},
		"session_lifecycle_state": acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"filter":                  acctest.RepresentationGroup{RepType: acctest.Required, Group: BastionsessionDataSourceFilterRepresentation}}
	BastionsessionDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_bastion_session.test_session.id}`}},
	}

	BastionsessionRepresentation = map[string]interface{}{
		"bastion_id":              acctest.Representation{RepType: acctest.Required, Create: `${oci_bastion_bastion.test_bastion.id}`},
		"key_details":             acctest.RepresentationGroup{RepType: acctest.Required, Group: BastionsessionKeyDetailsRepresentation},
		"target_resource_details": acctest.RepresentationGroup{RepType: acctest.Required, Group: BastionSessionTargetResourceDetailsRepresentation},
		"display_name":            acctest.Representation{RepType: acctest.Optional, Create: `managed_ssh`, Update: `managed_ssh2`},
		"key_type":                acctest.Representation{RepType: acctest.Optional, Create: `PUB`},
		"session_ttl_in_seconds":  acctest.Representation{RepType: acctest.Optional, Create: `1800`},
	}
	BastionSessionTargetResourceDetailsRepresentation = map[string]interface{}{
		"session_type":                               acctest.Representation{RepType: acctest.Required, Create: `MANAGED_SSH`},
		"target_resource_fqdn":                       acctest.Representation{RepType: acctest.Optional, Create: `targetResourceFqdn`},
		"target_resource_id":                         acctest.Representation{RepType: acctest.Optional, Create: `${oci_core_instance.test_instance.id}`},
		"target_resource_operating_system_user_name": acctest.Representation{RepType: acctest.Optional, Create: `opc`},
		"target_resource_port":                       acctest.Representation{RepType: acctest.Optional, Create: `22`},
		"target_resource_private_ip_address":         acctest.Representation{RepType: acctest.Optional, Create: `${oci_core_instance.test_instance.private_ip}`},
	}

	BastionsessionKeyDetailsRepresentation = map[string]interface{}{
		"public_key_content": acctest.Representation{RepType: acctest.Required, Create: `ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABgQDjk96o3uQcgHId6l/gkCwiid5J48CxKEiyk+1tPQugfhzgIIBs2Xr4xLX/rb5Xkr7MIeXuU3gdYrrMPLuMhOvthIKj6U5ROJWiZ67X00pOLq64dFyam1lQ+S/R/SaQ4W0KhKfkVskRhg7V96U07BGo8lDwYRGnvJsNb7rt3oHgnXtTFs7cy3IbzH5Sl7XBZv7yePu9sY39FrxktHw7Avz9BDZQbNYFC/cpj5eVvtPX/sMbc/D1yfrvhAIrYarhcAjEmWkjOJvkVlyKBxaSA7+mnOqFcj99hj5ZQN69h2B4TtHw2G8WEsU/nlyzBAj1iGQEvCLKnyp7Lxviy81jyKt91NQ7W6qh4tcs1mOFBsTGx/mBsNPwZhGRe4jWH15T++qnBAp6Zzw8ydPrJTgHLK+h1AMGFKMQZKYnMRV+6JYaNbnCVmLlxXoxhGsufXZMMS4qmjAQUBakZQsfiwLUxZBd0ZXmDCaZBxf6KP7HgL2x0Gb8IF38F7ryaOg9oxifqI8= chiweng@chiweng-mac`},
	}

	bastionEnabledInstanceAgentConfigRepresentation = map[string]interface{}{
		"are_all_plugins_disabled": acctest.Representation{RepType: acctest.Required, Create: `false`, Update: `false`},
		"is_management_disabled":   acctest.Representation{RepType: acctest.Required, Create: `false`, Update: `false`},
		"is_monitoring_disabled":   acctest.Representation{RepType: acctest.Required, Create: `false`, Update: `false`},
		"plugins_config":           acctest.RepresentationGroup{RepType: acctest.Required, Group: bastionEnabledInstanceAgentConfigPluginsConfigRepresentation},
	}

	bastionEnabledInstanceAgentConfigPluginsConfigRepresentation = map[string]interface{}{
		"desired_state": acctest.Representation{RepType: acctest.Required, Create: `ENABLED`},
		"name":          acctest.Representation{RepType: acctest.Required, Create: `Bastion`},
	}

	allOCIServiceGatewayServicesRepresentation = map[string]interface{}{
		"service_id": acctest.Representation{RepType: acctest.Required, Create: `${lookup(data.oci_core_services.test_services.services[1], "id")}`},
	}

	allOCIServiceRouteTableRouteRulesRepresentationWithServiceCidr = map[string]interface{}{
		"network_entity_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_core_service_gateway.test_service_gateway.id}`},
		"destination":       acctest.Representation{RepType: acctest.Required, Create: `${lookup(data.oci_core_services.test_services.services[1], "cidr_block")}`},
		"destination_type":  acctest.Representation{RepType: acctest.Required, Create: `SERVICE_CIDR_BLOCK`},
	}

	BastionSessionResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_bastion_bastion", "test_bastion", acctest.Required, acctest.Create, BastionbastionRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_subnet", "test_subnet", acctest.Required, acctest.Create, CoreSubnetRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", acctest.Required, acctest.Create, CoreVcnRepresentation) +
		AvailabilityDomainConfig +
		seestionImageInstanceDependencies +
		// Create instance as target host
		acctest.GenerateResourceFromRepresentationMap("oci_core_instance", "test_instance", acctest.Required, acctest.Create,
			acctest.RepresentationCopyWithNewProperties(CoreInstanceRepresentation, map[string]interface{}{
				"shape":        acctest.Representation{RepType: acctest.Required, Create: `VM.Standard1.1`},
				"image":        acctest.Representation{RepType: acctest.Required, Create: `${var.InstanceImageOCID[var.region]}`},
				"agent_config": acctest.RepresentationGroup{RepType: acctest.Required, Group: bastionEnabledInstanceAgentConfigRepresentation},
				"metadata":     acctest.Representation{RepType: acctest.Required, Create: map[string]string{"ssh_authorized_keys": "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABgQDjk96o3uQcgHId6l/gkCwiid5J48CxKEiyk+1tPQugfhzgIIBs2Xr4xLX/rb5Xkr7MIeXuU3gdYrrMPLuMhOvthIKj6U5ROJWiZ67X00pOLq64dFyam1lQ+S/R/SaQ4W0KhKfkVskRhg7V96U07BGo8lDwYRGnvJsNb7rt3oHgnXtTFs7cy3IbzH5Sl7XBZv7yePu9sY39FrxktHw7Avz9BDZQbNYFC/cpj5eVvtPX/sMbc/D1yfrvhAIrYarhcAjEmWkjOJvkVlyKBxaSA7+mnOqFcj99hj5ZQN69h2B4TtHw2G8WEsU/nlyzBAj1iGQEvCLKnyp7Lxviy81jyKt91NQ7W6qh4tcs1mOFBsTGx/mBsNPwZhGRe4jWH15T++qnBAp6Zzw8ydPrJTgHLK+h1AMGFKMQZKYnMRV+6JYaNbnCVmLlxXoxhGsufXZMMS4qmjAQUBakZQsfiwLUxZBd0ZXmDCaZBxf6KP7HgL2x0Gb8IF38F7ryaOg9oxifqI8= chiweng@chiweng-mac"}},
			})) +
		// Create Routable, Service Gateway, Internet Gateway for testing
		acctest.GenerateDataSourceFromRepresentationMap("oci_core_services", "test_services", acctest.Required, acctest.Create, CoreCoreServiceDataSourceRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_service_gateway", "test_service_gateway", acctest.Required, acctest.Create,
			acctest.RepresentationCopyWithNewProperties(CoreServiceGatewayRepresentation, map[string]interface{}{
				"services": acctest.RepresentationGroup{RepType: acctest.Required, Group: allOCIServiceGatewayServicesRepresentation},
			})) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_default_route_table", "default_route_table", acctest.Required, acctest.Create,
			acctest.RepresentationCopyWithNewProperties(routeTablesRepresentation, map[string]interface{}{
				"route_rules": acctest.RepresentationGroup{RepType: acctest.Required, Group: allOCIServiceRouteTableRouteRulesRepresentationWithServiceCidr},
			}))

	seestionImageInstanceDependencies = testBastionAvailableImage()
)

func testBastionAvailableImage() string {
	return `
	variable "InstanceImageOCID" {
	  type = "map"
	  default = {
		// See https://docs.us-phoenix-1.oraclecloud.com/images/
		// Oracle-provided image "Oracle-Linux-7.9-2021.04.09-0"
		us-phoenix-1 = "ocid1.image.oc1.phx.aaaaaaaanj7qmui2ux5hbiwtbtkzajuvvhuzo2y7755stim22ue6msqwv2ja"
		us-ashburn-1 = "ocid1.image.oc1.iad.aaaaaaaaw2wavtqrd3ynbrzabcnrs77pinccp55j2gqitjrrj2vf65sqj5kq"
	  }
	}`
}

// issue-routing-tag: bastion/default
func TestBastionSessionResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestBastionSessionResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_bastion_session.test_session"
	datasourceName := "data.oci_bastion_sessions.test_sessions"
	singularDatasourceName := "data.oci_bastion_session.test_session"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+BastionSessionResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_bastion_session", "test_session", acctest.Optional, acctest.Create, BastionsessionRepresentation), "bastion", "session", t)

	acctest.ResourceTest(t, testAccCheckBastionSessionDestroy, []resource.TestStep{
		// Create Dependencies
		{
			Config: config + compartmentIdVariableStr + BastionSessionResourceDependencies,
			Check: func(s *terraform.State) (err error) {
				log.Printf("[DEBUG] Wait for instance and bastion plugin to be run")
				time.Sleep(5 * time.Minute)
				return nil
			},
		},
		// verify Create
		{
			Config: config + compartmentIdVariableStr + BastionSessionResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_bastion_session", "test_session", acctest.Required, acctest.Create, BastionsessionRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "bastion_id"),
				resource.TestCheckResourceAttr(resourceName, "key_details.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "key_details.0.public_key_content"),
				resource.TestCheckResourceAttr(resourceName, "target_resource_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "target_resource_details.0.session_type", "MANAGED_SSH"),
				resource.TestCheckResourceAttrSet(resourceName, "target_resource_details.0.target_resource_id"),
				resource.TestCheckResourceAttrSet(resourceName, "target_resource_details.0.target_resource_operating_system_user_name"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + BastionSessionResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + BastionSessionResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_bastion_session", "test_session", acctest.Optional, acctest.Create, BastionsessionRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "bastion_id"),
				resource.TestCheckResourceAttrSet(resourceName, "bastion_name"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "managed_ssh"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "key_details.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "key_details.0.public_key_content"),
				resource.TestCheckResourceAttr(resourceName, "key_type", "PUB"),
				resource.TestCheckResourceAttr(resourceName, "session_ttl_in_seconds", "1800"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttr(resourceName, "target_resource_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "target_resource_details.0.session_type", "MANAGED_SSH"),
				resource.TestCheckResourceAttrSet(resourceName, "target_resource_details.0.target_resource_id"),
				resource.TestCheckResourceAttrSet(resourceName, "target_resource_details.0.target_resource_operating_system_user_name"),
				resource.TestCheckResourceAttr(resourceName, "target_resource_details.0.target_resource_port", "22"),
				resource.TestCheckResourceAttrSet(resourceName, "target_resource_details.0.target_resource_private_ip_address"),
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
			Config: config + compartmentIdVariableStr + BastionSessionResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_bastion_session", "test_session", acctest.Optional, acctest.Update, BastionsessionRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "bastion_id"),
				resource.TestCheckResourceAttrSet(resourceName, "bastion_name"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "managed_ssh2"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "key_details.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "key_details.0.public_key_content"),
				resource.TestCheckResourceAttr(resourceName, "key_type", "PUB"),
				resource.TestCheckResourceAttr(resourceName, "session_ttl_in_seconds", "1800"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttr(resourceName, "target_resource_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "target_resource_details.0.session_type", "MANAGED_SSH"),
				resource.TestCheckResourceAttrSet(resourceName, "target_resource_details.0.target_resource_id"),
				resource.TestCheckResourceAttrSet(resourceName, "target_resource_details.0.target_resource_operating_system_user_name"),
				resource.TestCheckResourceAttr(resourceName, "target_resource_details.0.target_resource_port", "22"),
				resource.TestCheckResourceAttrSet(resourceName, "target_resource_details.0.target_resource_private_ip_address"),
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
				acctest.GenerateDataSourceFromRepresentationMap("oci_bastion_sessions", "test_sessions", acctest.Optional, acctest.Update, sessionDataSourceRepresentation) +
				compartmentIdVariableStr + BastionSessionResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_bastion_session", "test_session", acctest.Optional, acctest.Update, BastionsessionRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "bastion_id"),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "managed_ssh2"),
				resource.TestCheckResourceAttrSet(datasourceName, "session_id"),
				resource.TestCheckResourceAttr(datasourceName, "session_lifecycle_state", "ACTIVE"),

				resource.TestCheckResourceAttr(datasourceName, "sessions.#", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "sessions.0.bastion_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "sessions.0.bastion_name"),
				resource.TestCheckResourceAttr(datasourceName, "sessions.0.display_name", "managed_ssh2"),
				resource.TestCheckResourceAttrSet(datasourceName, "sessions.0.id"),
				resource.TestCheckResourceAttr(datasourceName, "sessions.0.session_ttl_in_seconds", "1800"),
				resource.TestCheckResourceAttrSet(datasourceName, "sessions.0.state"),
				resource.TestCheckResourceAttr(datasourceName, "sessions.0.target_resource_details.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "sessions.0.target_resource_details.0.session_type", "MANAGED_SSH"),
				resource.TestCheckResourceAttrSet(datasourceName, "sessions.0.target_resource_details.0.target_resource_display_name"),
				resource.TestCheckResourceAttr(datasourceName, "sessions.0.target_resource_details.0.target_resource_fqdn", "targetResourceFqdn"),
				resource.TestCheckResourceAttrSet(datasourceName, "sessions.0.target_resource_details.0.target_resource_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "sessions.0.target_resource_details.0.target_resource_operating_system_user_name"),
				resource.TestCheckResourceAttr(datasourceName, "sessions.0.target_resource_details.0.target_resource_port", "22"),
				resource.TestCheckResourceAttrSet(datasourceName, "sessions.0.target_resource_details.0.target_resource_private_ip_address"),
				resource.TestCheckResourceAttrSet(datasourceName, "sessions.0.time_created"),
				resource.TestCheckResourceAttrSet(datasourceName, "sessions.0.time_updated"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_bastion_session", "test_session", acctest.Required, acctest.Create, BastionBastionsessionSingularDataSourceRepresentation) +
				compartmentIdVariableStr + BastionSessionResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "session_id"),

				resource.TestCheckResourceAttrSet(singularDatasourceName, "bastion_name"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "bastion_user_name"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "managed_ssh2"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "key_details.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "key_details.0.public_key_content"),
				resource.TestCheckResourceAttr(singularDatasourceName, "key_type", "PUB"),
				resource.TestCheckResourceAttr(singularDatasourceName, "session_ttl_in_seconds", "1800"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttr(singularDatasourceName, "target_resource_details.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "target_resource_details.0.session_type", "MANAGED_SSH"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "target_resource_details.0.target_resource_display_name"),
				resource.TestCheckResourceAttr(singularDatasourceName, "target_resource_details.0.target_resource_fqdn", "targetResourceFqdn"),
				resource.TestCheckResourceAttr(singularDatasourceName, "target_resource_details.0.target_resource_port", "22"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "target_resource_details.0.target_resource_private_ip_address"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
		// verify resource import
		{
			Config:                  config + BastionSessionRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckBastionSessionDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).BastionClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_bastion_session" {
			noResourceFound = false
			request := oci_bastion.GetSessionRequest{}

			tmp := rs.Primary.ID
			request.SessionId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "bastion")

			response, err := client.GetSession(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_bastion.SessionLifecycleStateDeleted): true,
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
	if !acctest.InSweeperExcludeList("BastionSession") {
		resource.AddTestSweepers("BastionSession", &resource.Sweeper{
			Name:         "BastionSession",
			Dependencies: acctest.DependencyGraph["session"],
			F:            sweepBastionSessionResource,
		})
	}
}

func sweepBastionSessionResource(compartment string) error {
	bastionClient := acctest.GetTestClients(&schema.ResourceData{}).BastionClient()
	sessionIds, err := getBastionSessionIds(compartment)
	if err != nil {
		return err
	}
	for _, sessionId := range sessionIds {
		if ok := acctest.SweeperDefaultResourceId[sessionId]; !ok {
			deleteSessionRequest := oci_bastion.DeleteSessionRequest{}

			deleteSessionRequest.SessionId = &sessionId

			deleteSessionRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "bastion")
			_, error := bastionClient.DeleteSession(context.Background(), deleteSessionRequest)
			if error != nil {
				fmt.Printf("Error deleting Session %s %s, It is possible that the resource is already deleted. Please verify manually \n", sessionId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &sessionId, BastionsessionsSweepWaitCondition, time.Duration(3*time.Minute),
				BastionsessionsSweepResponseFetchOperation, "bastion", true)
		}
	}
	return nil
}

func getBastionSessionIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "SessionId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	bastionClient := acctest.GetTestClients(&schema.ResourceData{}).BastionClient()

	listSessionsRequest := oci_bastion.ListSessionsRequest{}

	bastionIds, error := getBastionSessionIds(compartment)
	if error != nil {
		return resourceIds, fmt.Errorf("Error getting bastionId required for Session resource requests \n")
	}
	for _, bastionId := range bastionIds {
		listSessionsRequest.BastionId = &bastionId

		listSessionsResponse, err := bastionClient.ListSessions(context.Background(), listSessionsRequest)

		if err != nil {
			return resourceIds, fmt.Errorf("Error getting Session list for compartment id : %s , %s \n", compartmentId, err)
		}
		for _, session := range listSessionsResponse.Items {
			id := *session.Id
			resourceIds = append(resourceIds, id)
			acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "SessionId", id)
		}

	}
	return resourceIds, nil
}

func BastionsessionsSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if sessionResponse, ok := response.Response.(oci_bastion.GetSessionResponse); ok {
		return sessionResponse.LifecycleState != oci_bastion.SessionLifecycleStateDeleted
	}
	return false
}

func BastionsessionsSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.BastionClient().GetSession(context.Background(), oci_bastion.GetSessionRequest{
		SessionId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
