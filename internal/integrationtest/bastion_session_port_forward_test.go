// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/terraform-providers/terraform-provider-oci/internal/acctest"
	"github.com/terraform-providers/terraform-provider-oci/internal/resourcediscovery"
	"github.com/terraform-providers/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	PortForwardSessionRequiredOnlyResource = PortForwardSessionResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_bastion_session", "test_session", acctest.Required, acctest.Create, portForwardSessionRepresentation)

	PortForwardSessionResourceConfig = PortForwardSessionResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_bastion_session", "test_session", acctest.Optional, acctest.Update, portForwardSessionRepresentation)

	portForwardSessionSingularDataSourceRepresentation = map[string]interface{}{
		"session_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_bastion_session.test_session.id}`},
	}

	portForwardSessionDataSourceRepresentation = map[string]interface{}{
		"bastion_id":              acctest.Representation{RepType: acctest.Required, Create: `${oci_bastion_bastion.test_bastion.id}`},
		"display_name":            acctest.Representation{RepType: acctest.Optional, Create: `portForwardTest`, Update: `portForwardTest2`},
		"session_id":              acctest.Representation{RepType: acctest.Optional, Create: `${oci_bastion_session.test_session.id}`},
		"session_lifecycle_state": acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"filter":                  acctest.RepresentationGroup{RepType: acctest.Required, Group: portForwardSessionDataSourceFilterRepresentation}}
	portForwardSessionDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_bastion_session.test_session.id}`}},
	}

	portForwardSessionRepresentation = map[string]interface{}{
		"bastion_id":              acctest.Representation{RepType: acctest.Required, Create: `${oci_bastion_bastion.test_bastion.id}`},
		"key_details":             acctest.RepresentationGroup{RepType: acctest.Required, Group: portForwardSessionKeyDetailsRepresentation},
		"target_resource_details": acctest.RepresentationGroup{RepType: acctest.Required, Group: portForwardSessionTargetResourceDetailsRepresentation},
		"display_name":            acctest.Representation{RepType: acctest.Optional, Create: `portForwardTest`, Update: `portForwardTest2`},
		"key_type":                acctest.Representation{RepType: acctest.Optional, Create: `PUB`},
		"session_ttl_in_seconds":  acctest.Representation{RepType: acctest.Optional, Create: `3600`},
	}
	portForwardSessionTargetResourceDetailsRepresentation = map[string]interface{}{
		"session_type":                       acctest.Representation{RepType: acctest.Required, Create: `PORT_FORWARDING`},
		"target_resource_id":                 acctest.Representation{RepType: acctest.Required, Create: `${oci_core_instance.test_instance.id}`},
		"target_resource_port":               acctest.Representation{RepType: acctest.Optional, Create: `22`},
		"target_resource_private_ip_address": acctest.Representation{RepType: acctest.Optional, Create: `${oci_core_instance.test_instance.private_ip}`},
	}

	portForwardSessionKeyDetailsRepresentation = map[string]interface{}{
		"public_key_content": acctest.Representation{RepType: acctest.Required, Create: `ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQCsIMv3VTcvJeEQCz6+0Lyj4m6G02b6tGMgYwjyciXZY+GDpQhTq1RhY66dHgHP1Y6OWqpHL6lcs+OFx6iFWjNGwJmqDtR5T0/3kU7qZmPSQhe//Y4VU71wU15EA4LkbmwEX+9HIqUpPnD3XNoI+QLK7WTan3fqprK0DSc0XuRSA3H+H5meTxriFt8xbUHfo/qPBQJ0MKf9GOBQRKR+Cs6X2R0XxvF7XOmBv7NnapCpV2r7Yffkd1732m0g6i4lj2DNb1qQDDyws9TfGxH2OMlr1Rm71EYKQx3vO5Qhgszvpircxu+viADohSxTucwDii29ZaNDO8d3WZ11JSA/QH4h bastion`},
	}

	PortForwardSessionResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_bastion_bastion", "test_bastion", acctest.Required, acctest.Create, bastionRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_subnet", "test_subnet", acctest.Required, acctest.Create, subnetRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", acctest.Required, acctest.Create, vcnRepresentation) +
		AvailabilityDomainConfig +
		seestionImageInstanceDependencies +
		// Create instance as target host
		acctest.GenerateResourceFromRepresentationMap("oci_core_instance", "test_instance", acctest.Required, acctest.Create,
			acctest.RepresentationCopyWithNewProperties(instanceRepresentation, map[string]interface{}{
				"shape":        acctest.Representation{RepType: acctest.Required, Create: `VM.Standard1.1`},
				"image":        acctest.Representation{RepType: acctest.Required, Create: `${var.InstanceImageOCID[var.region]}`},
				"agent_config": acctest.RepresentationGroup{RepType: acctest.Required, Group: bastionEnabledInstanceAgentConfigRepresentation},
				"metadata":     acctest.Representation{RepType: acctest.Required, Create: map[string]string{"ssh_authorized_keys": "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABgQDjk96o3uQcgHId6l/gkCwiid5J48CxKEiyk+1tPQugfhzgIIBs2Xr4xLX/rb5Xkr7MIeXuU3gdYrrMPLuMhOvthIKj6U5ROJWiZ67X00pOLq64dFyam1lQ+S/R/SaQ4W0KhKfkVskRhg7V96U07BGo8lDwYRGnvJsNb7rt3oHgnXtTFs7cy3IbzH5Sl7XBZv7yePu9sY39FrxktHw7Avz9BDZQbNYFC/cpj5eVvtPX/sMbc/D1yfrvhAIrYarhcAjEmWkjOJvkVlyKBxaSA7+mnOqFcj99hj5ZQN69h2B4TtHw2G8WEsU/nlyzBAj1iGQEvCLKnyp7Lxviy81jyKt91NQ7W6qh4tcs1mOFBsTGx/mBsNPwZhGRe4jWH15T++qnBAp6Zzw8ydPrJTgHLK+h1AMGFKMQZKYnMRV+6JYaNbnCVmLlxXoxhGsufXZMMS4qmjAQUBakZQsfiwLUxZBd0ZXmDCaZBxf6KP7HgL2x0Gb8IF38F7ryaOg9oxifqI8= chiweng@chiweng-mac"}},
			})) +
		// Create Routable, Service Gateway, Internet Gateway for testing
		acctest.GenerateDataSourceFromRepresentationMap("oci_core_services", "test_services", acctest.Required, acctest.Create, serviceDataSourceRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_service_gateway", "test_service_gateway", acctest.Required, acctest.Create,
			acctest.RepresentationCopyWithNewProperties(serviceGatewayRepresentation, map[string]interface{}{
				"services": acctest.RepresentationGroup{RepType: acctest.Required, Group: allOCIServiceGatewayServicesRepresentation},
			})) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_default_route_table", "default_route_table", acctest.Required, acctest.Create,
			acctest.RepresentationCopyWithNewProperties(routeTablesRepresentation, map[string]interface{}{
				"route_rules": acctest.RepresentationGroup{RepType: acctest.Required, Group: allOCIServiceRouteTableRouteRulesRepresentationWithServiceCidr},
			}))
)

// issue-routing-tag: bastion/default
func TestBastionSessionResource_port_forwarding(t *testing.T) {
	httpreplay.SetScenario("TestBastionSessionResource_port_forwarding")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_bastion_session.test_session"
	datasourceName := "data.oci_bastion_sessions.test_sessions"
	singularDatasourceName := "data.oci_bastion_session.test_session"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+PortForwardSessionResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_bastion_session", "test_session", acctest.Optional, acctest.Create, portForwardSessionRepresentation), "bastion", "session", t)

	acctest.ResourceTest(t, testAccCheckBastionSessionDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + PortForwardSessionResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_bastion_session", "test_session", acctest.Required, acctest.Create, portForwardSessionRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "bastion_id"),
				resource.TestCheckResourceAttr(resourceName, "key_details.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "key_details.0.public_key_content"),
				resource.TestCheckResourceAttr(resourceName, "target_resource_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "target_resource_details.0.session_type", "PORT_FORWARDING"),
				resource.TestCheckResourceAttrSet(resourceName, "target_resource_details.0.target_resource_id"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + PortForwardSessionResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + PortForwardSessionResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_bastion_session", "test_session", acctest.Optional, acctest.Create, portForwardSessionRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "bastion_id"),
				resource.TestCheckResourceAttrSet(resourceName, "bastion_name"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "portForwardTest"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "key_details.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "key_details.0.public_key_content"),
				resource.TestCheckResourceAttr(resourceName, "key_type", "PUB"),
				resource.TestCheckResourceAttr(resourceName, "session_ttl_in_seconds", "3600"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttr(resourceName, "target_resource_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "target_resource_details.0.session_type", "PORT_FORWARDING"),
				resource.TestCheckResourceAttrSet(resourceName, "target_resource_details.0.target_resource_id"),
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
			Config: config + compartmentIdVariableStr + PortForwardSessionResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_bastion_session", "test_session", acctest.Optional, acctest.Update, portForwardSessionRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "bastion_id"),
				resource.TestCheckResourceAttrSet(resourceName, "bastion_name"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "portForwardTest2"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "key_details.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "key_details.0.public_key_content"),
				resource.TestCheckResourceAttr(resourceName, "key_type", "PUB"),
				resource.TestCheckResourceAttr(resourceName, "session_ttl_in_seconds", "3600"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttr(resourceName, "target_resource_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "target_resource_details.0.session_type", "PORT_FORWARDING"),
				resource.TestCheckResourceAttrSet(resourceName, "target_resource_details.0.target_resource_id"),
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
				acctest.GenerateDataSourceFromRepresentationMap("oci_bastion_sessions", "test_sessions", acctest.Optional, acctest.Update, portForwardSessionDataSourceRepresentation) +
				compartmentIdVariableStr + PortForwardSessionResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_bastion_session", "test_session", acctest.Optional, acctest.Update, portForwardSessionRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "bastion_id"),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "portForwardTest2"),
				resource.TestCheckResourceAttrSet(datasourceName, "session_id"),
				resource.TestCheckResourceAttr(datasourceName, "session_lifecycle_state", "ACTIVE"),

				resource.TestCheckResourceAttr(datasourceName, "sessions.#", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "sessions.0.bastion_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "sessions.0.bastion_name"),
				resource.TestCheckResourceAttr(datasourceName, "sessions.0.display_name", "portForwardTest2"),
				resource.TestCheckResourceAttrSet(datasourceName, "sessions.0.id"),
				resource.TestCheckResourceAttr(datasourceName, "sessions.0.session_ttl_in_seconds", "3600"),
				resource.TestCheckResourceAttrSet(datasourceName, "sessions.0.state"),
				resource.TestCheckResourceAttr(datasourceName, "sessions.0.target_resource_details.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "sessions.0.target_resource_details.0.session_type", "PORT_FORWARDING"),
				resource.TestCheckResourceAttrSet(datasourceName, "sessions.0.target_resource_details.0.target_resource_id"),
				resource.TestCheckResourceAttr(datasourceName, "sessions.0.target_resource_details.0.target_resource_port", "22"),
				resource.TestCheckResourceAttrSet(resourceName, "target_resource_details.0.target_resource_private_ip_address"),
				resource.TestCheckResourceAttrSet(datasourceName, "sessions.0.time_created"),
				resource.TestCheckResourceAttrSet(datasourceName, "sessions.0.time_updated"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_bastion_session", "test_session", acctest.Required, acctest.Create, portForwardSessionSingularDataSourceRepresentation) +
				compartmentIdVariableStr + PortForwardSessionResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "session_id"),

				resource.TestCheckResourceAttrSet(singularDatasourceName, "bastion_name"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "portForwardTest2"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "key_details.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "key_details.0.public_key_content"),
				resource.TestCheckResourceAttr(singularDatasourceName, "key_type", "PUB"),
				resource.TestCheckResourceAttr(singularDatasourceName, "session_ttl_in_seconds", "3600"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttr(singularDatasourceName, "target_resource_details.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "target_resource_details.0.session_type", "PORT_FORWARDING"),
				resource.TestCheckResourceAttr(singularDatasourceName, "target_resource_details.0.target_resource_port", "22"),
				resource.TestCheckResourceAttrSet(resourceName, "target_resource_details.0.target_resource_private_ip_address"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
		// remove singular datasource from previous step so that it doesn't conflict with import tests
		{
			Config: config + compartmentIdVariableStr + PortForwardSessionResourceConfig,
		},
		// verify resource import
		{
			Config:                  config,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}
