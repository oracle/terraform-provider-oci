// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	PortForwardSessionRequiredOnlyResource = PortForwardSessionResourceDependencies +
		generateResourceFromRepresentationMap("oci_bastion_session", "test_session", Required, Create, portForwardSessionRepresentation)

	PortForwardSessionResourceConfig = PortForwardSessionResourceDependencies +
		generateResourceFromRepresentationMap("oci_bastion_session", "test_session", Optional, Update, portForwardSessionRepresentation)

	portForwardSessionSingularDataSourceRepresentation = map[string]interface{}{
		"session_id": Representation{repType: Required, create: `${oci_bastion_session.test_session.id}`},
	}

	portForwardSessionDataSourceRepresentation = map[string]interface{}{
		"bastion_id":              Representation{repType: Required, create: `${oci_bastion_bastion.test_bastion.id}`},
		"display_name":            Representation{repType: Optional, create: `portForwardTest`, update: `portForwardTest2`},
		"session_id":              Representation{repType: Optional, create: `${oci_bastion_session.test_session.id}`},
		"session_lifecycle_state": Representation{repType: Optional, create: `ACTIVE`},
		"filter":                  RepresentationGroup{Required, portForwardSessionDataSourceFilterRepresentation}}
	portForwardSessionDataSourceFilterRepresentation = map[string]interface{}{
		"name":   Representation{repType: Required, create: `id`},
		"values": Representation{repType: Required, create: []string{`${oci_bastion_session.test_session.id}`}},
	}

	portForwardSessionRepresentation = map[string]interface{}{
		"bastion_id":              Representation{repType: Required, create: `${oci_bastion_bastion.test_bastion.id}`},
		"key_details":             RepresentationGroup{Required, portForwardSessionKeyDetailsRepresentation},
		"target_resource_details": RepresentationGroup{Required, portForwardSessionTargetResourceDetailsRepresentation},
		"display_name":            Representation{repType: Optional, create: `portForwardTest`, update: `portForwardTest2`},
		"key_type":                Representation{repType: Optional, create: `PUB`},
		"session_ttl_in_seconds":  Representation{repType: Optional, create: `3600`},
	}
	portForwardSessionTargetResourceDetailsRepresentation = map[string]interface{}{
		"session_type":                       Representation{repType: Required, create: `PORT_FORWARDING`},
		"target_resource_id":                 Representation{repType: Required, create: `${oci_core_instance.test_instance.id}`},
		"target_resource_port":               Representation{repType: Optional, create: `22`},
		"target_resource_private_ip_address": Representation{repType: Optional, create: `${oci_core_instance.test_instance.private_ip}`},
	}

	portForwardSessionKeyDetailsRepresentation = map[string]interface{}{
		"public_key_content": Representation{repType: Required, create: `ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQCsIMv3VTcvJeEQCz6+0Lyj4m6G02b6tGMgYwjyciXZY+GDpQhTq1RhY66dHgHP1Y6OWqpHL6lcs+OFx6iFWjNGwJmqDtR5T0/3kU7qZmPSQhe//Y4VU71wU15EA4LkbmwEX+9HIqUpPnD3XNoI+QLK7WTan3fqprK0DSc0XuRSA3H+H5meTxriFt8xbUHfo/qPBQJ0MKf9GOBQRKR+Cs6X2R0XxvF7XOmBv7NnapCpV2r7Yffkd1732m0g6i4lj2DNb1qQDDyws9TfGxH2OMlr1Rm71EYKQx3vO5Qhgszvpircxu+viADohSxTucwDii29ZaNDO8d3WZ11JSA/QH4h bastion`},
	}

	PortForwardSessionResourceDependencies = generateResourceFromRepresentationMap("oci_bastion_bastion", "test_bastion", Required, Create, bastionRepresentation) +
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
)

// issue-routing-tag: bastion/default
func TestBastionSessionResource_port_forwarding(t *testing.T) {
	httpreplay.SetScenario("TestBastionSessionResource_port_forwarding")
	defer httpreplay.SaveScenario()

	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_bastion_session.test_session"
	datasourceName := "data.oci_bastion_sessions.test_sessions"
	singularDatasourceName := "data.oci_bastion_session.test_session"

	var resId, resId2 string
	// Save TF content to create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	saveConfigContent(config+compartmentIdVariableStr+PortForwardSessionResourceDependencies+
		generateResourceFromRepresentationMap("oci_bastion_session", "test_session", Optional, Create, portForwardSessionRepresentation), "bastion", "session", t)

	ResourceTest(t, testAccCheckBastionSessionDestroy, []resource.TestStep{
		// verify create
		{
			Config: config + compartmentIdVariableStr + PortForwardSessionResourceDependencies +
				generateResourceFromRepresentationMap("oci_bastion_session", "test_session", Required, Create, portForwardSessionRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "bastion_id"),
				resource.TestCheckResourceAttr(resourceName, "key_details.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "key_details.0.public_key_content"),
				resource.TestCheckResourceAttr(resourceName, "target_resource_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "target_resource_details.0.session_type", "PORT_FORWARDING"),
				resource.TestCheckResourceAttrSet(resourceName, "target_resource_details.0.target_resource_id"),

				func(s *terraform.State) (err error) {
					resId, err = fromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next create
		{
			Config: config + compartmentIdVariableStr + PortForwardSessionResourceDependencies,
		},
		// verify create with optionals
		{
			Config: config + compartmentIdVariableStr + PortForwardSessionResourceDependencies +
				generateResourceFromRepresentationMap("oci_bastion_session", "test_session", Optional, Create, portForwardSessionRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
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
			Config: config + compartmentIdVariableStr + PortForwardSessionResourceDependencies +
				generateResourceFromRepresentationMap("oci_bastion_session", "test_session", Optional, Update, portForwardSessionRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
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
				generateDataSourceFromRepresentationMap("oci_bastion_sessions", "test_sessions", Optional, Update, portForwardSessionDataSourceRepresentation) +
				compartmentIdVariableStr + PortForwardSessionResourceDependencies +
				generateResourceFromRepresentationMap("oci_bastion_session", "test_session", Optional, Update, portForwardSessionRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
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
				generateDataSourceFromRepresentationMap("oci_bastion_session", "test_session", Required, Create, portForwardSessionSingularDataSourceRepresentation) +
				compartmentIdVariableStr + PortForwardSessionResourceConfig,
			Check: ComposeAggregateTestCheckFuncWrapper(
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
