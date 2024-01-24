// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	PathAnalysiRequiredOnlyResource = PathAnalysiResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_vn_monitoring_path_analysi", "test_path_analysi", acctest.Required, acctest.Create, pathAnalysiRepresentation)

	pathAnalysiRepresentation = map[string]interface{}{
		"type":                  acctest.Representation{RepType: acctest.Required, Create: `ADHOC_QUERY`},
		"cache_control":         acctest.Representation{RepType: acctest.Optional, Create: `cacheControl`},
		"compartment_id":        acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"destination_endpoint":  acctest.RepresentationGroup{RepType: acctest.Required, Group: pathAnalysiDestinationEndpointRepresentation},
		"path_analyzer_test_id": acctest.Representation{RepType: acctest.Optional, Create: `${oci_vn_monitoring_path_analyzer_test.test_path_analyzer_test.id}`},
		"protocol":              acctest.Representation{RepType: acctest.Required, Create: `1`},
		"protocol_parameters":   acctest.RepresentationGroup{RepType: acctest.Required, Group: pathAnalysiProtocolParametersRepresentation},
		"query_options":         acctest.RepresentationGroup{RepType: acctest.Optional, Group: pathAnalysiQueryOptionsRepresentation},
		"source_endpoint":       acctest.RepresentationGroup{RepType: acctest.Required, Group: pathAnalysiSourceEndpointRepresentation},
	}
	pathAnalysiDestinationEndpointRepresentation = map[string]interface{}{
		"type":                     acctest.Representation{RepType: acctest.Required, Create: `IP_ADDRESS`},
		"address":                  acctest.Representation{RepType: acctest.Required, Create: `${oci_core_instance.test_instance.private_ip}`},
		"instance_id":              acctest.Representation{RepType: acctest.Optional, Create: `${oci_core_instance.test_instance.id}`},
		"listener_id":              acctest.Representation{RepType: acctest.Optional, Create: `${oci_load_balancer_listener.test_listener.id}`},
		"load_balancer_id":         acctest.Representation{RepType: acctest.Optional, Create: `${oci_load_balancer_load_balancer.test_load_balancer.id}`},
		"network_load_balancer_id": acctest.Representation{RepType: acctest.Optional, Create: `${oci_network_load_balancer_network_load_balancer.test_network_load_balancer.id}`},
		"subnet_id":                acctest.Representation{RepType: acctest.Optional, Create: `${oci_core_subnet.test_subnet.id}`},
		"vlan_id":                  acctest.Representation{RepType: acctest.Optional, Create: `${oci_core_vlan.test_vlan.id}`},
		"vnic_id":                  acctest.Representation{RepType: acctest.Optional, Create: `${oci_core_vnic_attachment.test_vnic_attachment.vnic_id}`},
	}
	pathAnalysiProtocolParametersRepresentation = map[string]interface{}{
		"type":             acctest.Representation{RepType: acctest.Required, Create: `ICMP`},
		"destination_port": acctest.Representation{RepType: acctest.Required, Create: `10`},
		"icmp_code":        acctest.Representation{RepType: acctest.Required, Create: `4`},
		"icmp_type":        acctest.Representation{RepType: acctest.Required, Create: `3`},
		"source_port":      acctest.Representation{RepType: acctest.Required, Create: `10`},
	}
	pathAnalysiQueryOptionsRepresentation = map[string]interface{}{
		"is_bi_directional_analysis": acctest.Representation{RepType: acctest.Optional, Create: `false`},
	}
	pathAnalysiSourceEndpointRepresentation = map[string]interface{}{
		"type":                     acctest.Representation{RepType: acctest.Required, Create: `IP_ADDRESS`},
		"address":                  acctest.Representation{RepType: acctest.Required, Create: `${oci_core_instance.test_instance2.private_ip}`},
		"instance_id":              acctest.Representation{RepType: acctest.Optional, Create: `${oci_core_instance.test_instance2.id}`},
		"listener_id":              acctest.Representation{RepType: acctest.Optional, Create: `${oci_load_balancer_listener.test_listener.id}`},
		"load_balancer_id":         acctest.Representation{RepType: acctest.Optional, Create: `${oci_load_balancer_load_balancer.test_load_balancer.id}`},
		"network_load_balancer_id": acctest.Representation{RepType: acctest.Optional, Create: `${oci_network_load_balancer_network_load_balancer.test_network_load_balancer.id}`},
		"subnet_id":                acctest.Representation{RepType: acctest.Optional, Create: `${oci_core_subnet.test_subnet.id}`},
		"vlan_id":                  acctest.Representation{RepType: acctest.Optional, Create: `${oci_core_vlan.test_vlan.id}`},
		"vnic_id":                  acctest.Representation{RepType: acctest.Optional, Create: `${oci_core_vnic_attachment.test_vnic_attachment2.vnic_id}`},
	}
	listenerRepresentation = map[string]interface{}{
		"default_backend_set_name": acctest.Representation{RepType: acctest.Required, Create: `${oci_load_balancer_backend_set.test_backend_set.name}`},
		"load_balancer_id":         acctest.Representation{RepType: acctest.Required, Create: `${oci_load_balancer_load_balancer.test_load_balancer.id}`},
		"name":                     acctest.Representation{RepType: acctest.Required, Create: `myListener1`},
		"port":                     acctest.Representation{RepType: acctest.Required, Create: `10`, Update: `11`},
		"protocol":                 acctest.Representation{RepType: acctest.Required, Create: `HTTP`},
		"connection_configuration": acctest.RepresentationGroup{RepType: acctest.Optional, Group: listenerConnectionConfigurationRepresentation},
		"hostname_names":           acctest.Representation{RepType: acctest.Optional, Create: []string{`${oci_load_balancer_hostname.test_hostname.name}`}},
		"path_route_set_name":      acctest.Representation{RepType: acctest.Optional, Create: `${oci_load_balancer_path_route_set.test_path_route_set.name}`},
		"routing_policy_name":      acctest.Representation{RepType: acctest.Optional, Create: `${oci_load_balancer_load_balancer_routing_policy.test_load_balancer_routing_policy.name}`},
		"rule_set_names":           acctest.Representation{RepType: acctest.Optional, Create: []string{`${oci_load_balancer_rule_set.test_rule_set.name}`}},
		"ssl_configuration":        acctest.RepresentationGroup{RepType: acctest.Optional, Group: listenerSslConfigurationRepresentationOciCerts},
	}
	vlanRepresentationVnMonitoring = map[string]interface{}{
		"cidr_block":          acctest.Representation{RepType: acctest.Required, Create: `10.0.1.0/24`, Update: "10.0.0.0/16"},
		"compartment_id":      acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"vcn_id":              acctest.Representation{RepType: acctest.Required, Create: `${oci_core_vcn.test_vcn.id}`},
		"availability_domain": acctest.Representation{RepType: acctest.Optional, Create: `${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}`},
		"defined_tags":        acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"display_name":        acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"freeform_tags":       acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"nsg_ids":             acctest.Representation{RepType: acctest.Optional, Create: []string{`${oci_core_network_security_group.test_network_security_group.id}`}, Update: []string{}},
		"route_table_id":      acctest.Representation{RepType: acctest.Optional, Create: `${oci_core_route_table.test_route_table.id}`},
		"vlan_tag":            acctest.Representation{RepType: acctest.Optional, Create: `10`},
		"lifecycle":           acctest.RepresentationGroup{RepType: acctest.Required, Group: CoreNetworkSecurityIgnoreChangesNsgRepresentation},
	}
	PathAnalysiResourceDependencies = utils.OciImageIdsVariable +
		acctest.GenerateResourceFromRepresentationMap("oci_core_instance", "test_instance", acctest.Required, acctest.Create, CoreInstanceRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_instance", "test_instance2", acctest.Required, acctest.Create, instance2Representation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_subnet", "test_subnet", acctest.Required, acctest.Create, CoreSubnetRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", acctest.Required, acctest.Create, CoreVcnRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_vlan", "test_vlan", acctest.Required, acctest.Create, vlanRepresentationVnMonitoring) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_vnic_attachment", "test_vnic_attachment", acctest.Required, acctest.Create, CoreVnicAttachmentRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_vnic_attachment", "test_vnic_attachment2", acctest.Required, acctest.Create, vnicAttachment2Representation) +
		AvailabilityDomainConfig +
		acctest.GenerateResourceFromRepresentationMap("oci_load_balancer_backend_set", "test_backend_set", acctest.Required, acctest.Create, backendSetRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_load_balancer_listener", "test_listener", acctest.Required, acctest.Create, listenerRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_load_balancer_load_balancer", "test_load_balancer", acctest.Required, acctest.Create, loadBalancerRepresentation) +
		LoadBalancerSubnetDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_network_load_balancer_network_load_balancer", "test_network_load_balancer", acctest.Required, acctest.Create, NetworkLoadBalancerNetworkLoadBalancerRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_vn_monitoring_path_analyzer_test", "test_path_analyzer_test", acctest.Required, acctest.Create, pathAnalyzerTestRepresentation)
)

// issue-routing-tag: vn_monitoring/default
func TestVnMonitoringPathAnalysiResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestVnMonitoringPathAnalysiResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_vn_monitoring_path_analysi.test_path_analysi"

	var resId string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+PathAnalysiResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_vn_monitoring_path_analysi", "test_path_analysi", acctest.Optional, acctest.Create, pathAnalysiRepresentation), "vnmonitoring", "pathAnalysi", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + PathAnalysiResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_vn_monitoring_path_analysi", "test_path_analysi", acctest.Required, acctest.Create, pathAnalysiRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "destination_endpoint.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "destination_endpoint.0.address"),
				resource.TestCheckResourceAttr(resourceName, "destination_endpoint.0.type", "IP_ADDRESS"),
				resource.TestCheckResourceAttr(resourceName, "protocol", "1"),
				resource.TestCheckResourceAttr(resourceName, "source_endpoint.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "source_endpoint.0.address"),
				resource.TestCheckResourceAttr(resourceName, "source_endpoint.0.type", "IP_ADDRESS"),
				resource.TestCheckResourceAttr(resourceName, "type", "ADHOC_QUERY"),
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + PathAnalysiResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + PathAnalysiResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_vn_monitoring_path_analysi", "test_path_analysi", acctest.Optional, acctest.Create, pathAnalysiRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "destination_endpoint.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "destination_endpoint.0.address"),
				resource.TestCheckResourceAttr(resourceName, "destination_endpoint.0.type", "IP_ADDRESS"),
				resource.TestCheckResourceAttrSet(resourceName, "path_analyzer_test_id"),
				resource.TestCheckResourceAttr(resourceName, "protocol", "1"),
				resource.TestCheckResourceAttr(resourceName, "protocol_parameters.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "protocol_parameters.0.icmp_code", "4"),
				resource.TestCheckResourceAttr(resourceName, "protocol_parameters.0.icmp_type", "3"),
				resource.TestCheckResourceAttr(resourceName, "protocol_parameters.0.type", "ICMP"),
				resource.TestCheckResourceAttr(resourceName, "query_options.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "query_options.0.is_bi_directional_analysis", "false"),
				resource.TestCheckResourceAttr(resourceName, "source_endpoint.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "source_endpoint.0.address"),
				resource.TestCheckResourceAttrSet(resourceName, "source_endpoint.0.instance_id"),
				resource.TestCheckResourceAttr(resourceName, "source_endpoint.0.type", "IP_ADDRESS"),
				resource.TestCheckResourceAttrSet(resourceName, "source_endpoint.0.vnic_id"),
				resource.TestCheckResourceAttr(resourceName, "type", "ADHOC_QUERY"),

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
	})
}
