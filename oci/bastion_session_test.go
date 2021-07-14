// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
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
	oci_bastion "github.com/oracle/oci-go-sdk/v44/bastion"
	"github.com/oracle/oci-go-sdk/v44/common"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	SessionRequiredOnlyResource = SessionResourceDependencies +
		generateResourceFromRepresentationMap("oci_bastion_session", "test_session", Required, Create, sessionRepresentation)

	SessionResourceConfig = SessionResourceDependencies +
		generateResourceFromRepresentationMap("oci_bastion_session", "test_session", Optional, Update, sessionRepresentation)

	sessionSingularDataSourceRepresentation = map[string]interface{}{
		"session_id": Representation{repType: Required, create: `${oci_bastion_session.test_session.id}`},
	}

	sessionDataSourceRepresentation = map[string]interface{}{
		"bastion_id":              Representation{repType: Required, create: `${oci_bastion_bastion.test_bastion.id}`},
		"display_name":            Representation{repType: Optional, create: `managed_ssh`, update: `managed_ssh2`},
		"session_id":              Representation{repType: Optional, create: `${oci_bastion_session.test_session.id}`},
		"session_lifecycle_state": Representation{repType: Optional, create: `ACTIVE`},
		"filter":                  RepresentationGroup{Required, sessionDataSourceFilterRepresentation}}
	sessionDataSourceFilterRepresentation = map[string]interface{}{
		"name":   Representation{repType: Required, create: `id`},
		"values": Representation{repType: Required, create: []string{`${oci_bastion_session.test_session.id}`}},
	}

	sessionRepresentation = map[string]interface{}{
		"bastion_id":              Representation{repType: Required, create: `${oci_bastion_bastion.test_bastion.id}`},
		"key_details":             RepresentationGroup{Required, sessionKeyDetailsRepresentation},
		"target_resource_details": RepresentationGroup{Required, sessionTargetResourceDetailsRepresentation},
		"display_name":            Representation{repType: Optional, create: `managed_ssh`, update: `managed_ssh2`},
		"key_type":                Representation{repType: Optional, create: `PUB`},
		"session_ttl_in_seconds":  Representation{repType: Optional, create: `1800`},
	}
	sessionTargetResourceDetailsRepresentation = map[string]interface{}{
		"session_type":       Representation{repType: Required, create: `MANAGED_SSH`},
		"target_resource_id": Representation{repType: Required, create: `${oci_core_instance.test_instance.id}`},
		"target_resource_operating_system_user_name": Representation{repType: Required, create: `opc`},
		"target_resource_port":                       Representation{repType: Optional, create: `22`},
		"target_resource_private_ip_address":         Representation{repType: Optional, create: `${oci_core_instance.test_instance.private_ip}`},
	}

	sessionKeyDetailsRepresentation = map[string]interface{}{
		"public_key_content": Representation{repType: Required, create: `ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABgQDjk96o3uQcgHId6l/gkCwiid5J48CxKEiyk+1tPQugfhzgIIBs2Xr4xLX/rb5Xkr7MIeXuU3gdYrrMPLuMhOvthIKj6U5ROJWiZ67X00pOLq64dFyam1lQ+S/R/SaQ4W0KhKfkVskRhg7V96U07BGo8lDwYRGnvJsNb7rt3oHgnXtTFs7cy3IbzH5Sl7XBZv7yePu9sY39FrxktHw7Avz9BDZQbNYFC/cpj5eVvtPX/sMbc/D1yfrvhAIrYarhcAjEmWkjOJvkVlyKBxaSA7+mnOqFcj99hj5ZQN69h2B4TtHw2G8WEsU/nlyzBAj1iGQEvCLKnyp7Lxviy81jyKt91NQ7W6qh4tcs1mOFBsTGx/mBsNPwZhGRe4jWH15T++qnBAp6Zzw8ydPrJTgHLK+h1AMGFKMQZKYnMRV+6JYaNbnCVmLlxXoxhGsufXZMMS4qmjAQUBakZQsfiwLUxZBd0ZXmDCaZBxf6KP7HgL2x0Gb8IF38F7ryaOg9oxifqI8= chiweng@chiweng-mac`},
	}

	bastionEnabledInstanceAgentConfigRepresentation = map[string]interface{}{
		"are_all_plugins_disabled": Representation{repType: Required, create: `false`, update: `false`},
		"is_management_disabled":   Representation{repType: Required, create: `false`, update: `false`},
		"is_monitoring_disabled":   Representation{repType: Required, create: `false`, update: `false`},
		"plugins_config":           RepresentationGroup{Required, bastionEnabledInstanceAgentConfigPluginsConfigRepresentation},
	}

	bastionEnabledInstanceAgentConfigPluginsConfigRepresentation = map[string]interface{}{
		"desired_state": Representation{repType: Required, create: `ENABLED`},
		"name":          Representation{repType: Required, create: `Bastion`},
	}

	allOCIServiceGatewayServicesRepresentation = map[string]interface{}{
		"service_id": Representation{repType: Required, create: `${lookup(data.oci_core_services.test_services.services[1], "id")}`},
	}

	allOCIServiceRouteTableRouteRulesRepresentationWithServiceCidr = map[string]interface{}{
		"network_entity_id": Representation{repType: Required, create: `${oci_core_service_gateway.test_service_gateway.id}`},
		"destination":       Representation{repType: Required, create: `${lookup(data.oci_core_services.test_services.services[1], "cidr_block")}`},
		"destination_type":  Representation{repType: Required, create: `SERVICE_CIDR_BLOCK`},
	}

	SessionResourceDependencies = generateResourceFromRepresentationMap("oci_bastion_bastion", "test_bastion", Required, Create, bastionRepresentation) +
		generateResourceFromRepresentationMap("oci_core_subnet", "test_subnet", Required, Create, subnetRepresentation) +
		generateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", Required, Create, vcnRepresentation) +
		AvailabilityDomainConfig +
		seestionImageInstanceDependencies +
		// Create instance as target host
		generateResourceFromRepresentationMap("oci_core_instance", "test_instance", Required, Create,
			representationCopyWithNewProperties(instanceRepresentation, map[string]interface{}{
				"shape":        Representation{repType: Required, create: `VM.Standard1.1`},
				"image":        Representation{repType: Required, create: `${var.InstanceImageOCID[var.region]}`},
				"agent_config": RepresentationGroup{Required, bastionEnabledInstanceAgentConfigRepresentation},
				"metadata":     Representation{repType: Required, create: map[string]string{"ssh_authorized_keys": "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABgQDjk96o3uQcgHId6l/gkCwiid5J48CxKEiyk+1tPQugfhzgIIBs2Xr4xLX/rb5Xkr7MIeXuU3gdYrrMPLuMhOvthIKj6U5ROJWiZ67X00pOLq64dFyam1lQ+S/R/SaQ4W0KhKfkVskRhg7V96U07BGo8lDwYRGnvJsNb7rt3oHgnXtTFs7cy3IbzH5Sl7XBZv7yePu9sY39FrxktHw7Avz9BDZQbNYFC/cpj5eVvtPX/sMbc/D1yfrvhAIrYarhcAjEmWkjOJvkVlyKBxaSA7+mnOqFcj99hj5ZQN69h2B4TtHw2G8WEsU/nlyzBAj1iGQEvCLKnyp7Lxviy81jyKt91NQ7W6qh4tcs1mOFBsTGx/mBsNPwZhGRe4jWH15T++qnBAp6Zzw8ydPrJTgHLK+h1AMGFKMQZKYnMRV+6JYaNbnCVmLlxXoxhGsufXZMMS4qmjAQUBakZQsfiwLUxZBd0ZXmDCaZBxf6KP7HgL2x0Gb8IF38F7ryaOg9oxifqI8= chiweng@chiweng-mac"}},
			})) +
		// Create Routable, Service Gateway, Internet Gateway for testing
		generateDataSourceFromRepresentationMap("oci_core_services", "test_services", Required, Create, serviceDataSourceRepresentation) +
		generateResourceFromRepresentationMap("oci_core_service_gateway", "test_service_gateway", Required, Create,
			representationCopyWithNewProperties(serviceGatewayRepresentation, map[string]interface{}{
				"services": RepresentationGroup{Required, allOCIServiceGatewayServicesRepresentation},
			})) +
		generateResourceFromRepresentationMap("oci_core_default_route_table", "default_route_table", Required, Create,
			representationCopyWithNewProperties(routeTablesRepresentation, map[string]interface{}{
				"route_rules": RepresentationGroup{Required, allOCIServiceRouteTableRouteRulesRepresentationWithServiceCidr},
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

func TestBastionSessionResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestBastionSessionResource_basic")
	defer httpreplay.SaveScenario()

	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_bastion_session.test_session"
	datasourceName := "data.oci_bastion_sessions.test_sessions"
	singularDatasourceName := "data.oci_bastion_session.test_session"

	var resId, resId2 string
	// Save TF content to create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	saveConfigContent(config+compartmentIdVariableStr+SessionResourceDependencies+
		generateResourceFromRepresentationMap("oci_bastion_session", "test_session", Optional, Create, sessionRepresentation), "bastion", "session", t)

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		CheckDestroy: testAccCheckBastionSessionDestroy,
		Steps: []resource.TestStep{
			// verify create
			{
				Config: config + compartmentIdVariableStr + SessionResourceDependencies +
					generateResourceFromRepresentationMap("oci_bastion_session", "test_session", Required, Create, sessionRepresentation),
				Check: ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttrSet(resourceName, "bastion_id"),
					resource.TestCheckResourceAttr(resourceName, "key_details.#", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "key_details.0.public_key_content"),
					resource.TestCheckResourceAttr(resourceName, "target_resource_details.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "target_resource_details.0.session_type", "MANAGED_SSH"),
					resource.TestCheckResourceAttrSet(resourceName, "target_resource_details.0.target_resource_id"),
					resource.TestCheckResourceAttrSet(resourceName, "target_resource_details.0.target_resource_operating_system_user_name"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},

			// delete before next create
			{
				Config: config + compartmentIdVariableStr + SessionResourceDependencies,
			},
			// verify create with optionals
			{
				Config: config + compartmentIdVariableStr + SessionResourceDependencies +
					generateResourceFromRepresentationMap("oci_bastion_session", "test_session", Optional, Create, sessionRepresentation),
				Check: ComposeAggregateTestCheckFuncWrapper(
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
						resId, err = fromInstanceState(s, resourceName, "id")
						if isEnableExportCompartment, _ := strconv.ParseBool(getEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
							if errExport := testExportCompartmentWithResourceName(&resId, &compartmentId, resourceName); errExport != nil {
								return errExport
							}
						}
						return err
					},
				),
			},
			// verify updates to updatable parameters
			{
				Config: config + compartmentIdVariableStr + SessionResourceDependencies +
					generateResourceFromRepresentationMap("oci_bastion_session", "test_session", Optional, Update, sessionRepresentation),
				Check: ComposeAggregateTestCheckFuncWrapper(
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
					generateDataSourceFromRepresentationMap("oci_bastion_sessions", "test_sessions", Optional, Update, sessionDataSourceRepresentation) +
					compartmentIdVariableStr + SessionResourceDependencies +
					generateResourceFromRepresentationMap("oci_bastion_session", "test_session", Optional, Update, sessionRepresentation),
				Check: ComposeAggregateTestCheckFuncWrapper(
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
					generateDataSourceFromRepresentationMap("oci_bastion_session", "test_session", Required, Create, sessionSingularDataSourceRepresentation) +
					compartmentIdVariableStr + SessionResourceConfig,
				Check: ComposeAggregateTestCheckFuncWrapper(
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
					resource.TestCheckResourceAttr(singularDatasourceName, "target_resource_details.0.target_resource_port", "22"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "target_resource_details.0.target_resource_private_ip_address"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
				),
			},
			// remove singular datasource from previous step so that it doesn't conflict with import tests
			{
				Config: config + compartmentIdVariableStr + SessionResourceConfig,
			},
			// verify resource import
			{
				Config:                  config,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
				ResourceName:            resourceName,
			},
		},
	})
}

func testAccCheckBastionSessionDestroy(s *terraform.State) error {
	noResourceFound := true
	client := testAccProvider.Meta().(*OracleClients).bastionClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_bastion_session" {
			noResourceFound = false
			request := oci_bastion.GetSessionRequest{}

			tmp := rs.Primary.ID
			request.SessionId = &tmp

			request.RequestMetadata.RetryPolicy = getRetryPolicy(true, "bastion")

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
	if DependencyGraph == nil {
		initDependencyGraph()
	}
	if !inSweeperExcludeList("BastionSession") {
		resource.AddTestSweepers("BastionSession", &resource.Sweeper{
			Name:         "BastionSession",
			Dependencies: DependencyGraph["session"],
			F:            sweepBastionSessionResource,
		})
	}
}

func sweepBastionSessionResource(compartment string) error {
	bastionClient := GetTestClients(&schema.ResourceData{}).bastionClient()
	sessionIds, err := getSessionIds(compartment)
	if err != nil {
		return err
	}
	for _, sessionId := range sessionIds {
		if ok := SweeperDefaultResourceId[sessionId]; !ok {
			deleteSessionRequest := oci_bastion.DeleteSessionRequest{}

			deleteSessionRequest.SessionId = &sessionId

			deleteSessionRequest.RequestMetadata.RetryPolicy = getRetryPolicy(true, "bastion")
			_, error := bastionClient.DeleteSession(context.Background(), deleteSessionRequest)
			if error != nil {
				fmt.Printf("Error deleting Session %s %s, It is possible that the resource is already deleted. Please verify manually \n", sessionId, error)
				continue
			}
			waitTillCondition(testAccProvider, &sessionId, sessionSweepWaitCondition, time.Duration(3*time.Minute),
				sessionSweepResponseFetchOperation, "bastion", true)
		}
	}
	return nil
}

func getSessionIds(compartment string) ([]string, error) {
	ids := getResourceIdsToSweep(compartment, "SessionId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	bastionClient := GetTestClients(&schema.ResourceData{}).bastionClient()

	listSessionsRequest := oci_bastion.ListSessionsRequest{}

	bastionIds, error := getBastionIds(compartment)
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
			addResourceIdToSweeperResourceIdMap(compartmentId, "SessionId", id)
		}

	}
	return resourceIds, nil
}

func sessionSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if sessionResponse, ok := response.Response.(oci_bastion.GetSessionResponse); ok {
		return sessionResponse.LifecycleState != oci_bastion.SessionLifecycleStateDeleted
	}
	return false
}

func sessionSweepResponseFetchOperation(client *OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.bastionClient().GetSession(context.Background(), oci_bastion.GetSessionRequest{
		SessionId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
