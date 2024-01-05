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
	oci_vn_monitoring "github.com/oracle/oci-go-sdk/v65/vnmonitoring"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"

	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
)

var (
	PathAnalyzerTestRequiredOnlyResource = PathAnalyzerTestResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_vn_monitoring_path_analyzer_test", "test_path_analyzer_test", acctest.Required, acctest.Create, pathAnalyzerTestRepresentation)

	PathAnalyzerTestResourceConfig = PathAnalyzerTestResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_vn_monitoring_path_analyzer_test", "test_path_analyzer_test", acctest.Optional, acctest.Update, pathAnalyzerTestRepresentation)

	pathAnalyzerTestSingularDataSourceRepresentation = map[string]interface{}{
		"path_analyzer_test_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_vn_monitoring_path_analyzer_test.test_path_analyzer_test.id}`},
	}

	pathAnalyzerTestDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":   acctest.Representation{RepType: acctest.Optional, Create: `Path Analyzer Test`, Update: `displayName2`},
		"state":          acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"filter":         acctest.RepresentationGroup{RepType: acctest.Required, Group: pathAnalyzerTestDataSourceFilterRepresentation}}
	pathAnalyzerTestDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_vn_monitoring_path_analyzer_test.test_path_analyzer_test.id}`}},
	}

	pathAnalyzerTestRepresentation = map[string]interface{}{
		"compartment_id":       acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"destination_endpoint": acctest.RepresentationGroup{RepType: acctest.Required, Group: pathAnalyzerTestDestinationEndpointRepresentation},
		"protocol":             acctest.Representation{RepType: acctest.Required, Create: `1`, Update: `17`},
		"source_endpoint":      acctest.RepresentationGroup{RepType: acctest.Required, Group: pathAnalyzerTestSourceEndpointRepresentation},
		"defined_tags":         acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"display_name":         acctest.Representation{RepType: acctest.Optional, Create: `Path Analyzer Test`, Update: `displayName2`},
		"freeform_tags":        acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"bar-key": "value"}, Update: map[string]string{"Department": "Accounting"}},
		"protocol_parameters":  acctest.RepresentationGroup{RepType: acctest.Required, Group: pathAnalyzerTestProtocolParametersRepresentation},
		"query_options":        acctest.RepresentationGroup{RepType: acctest.Optional, Group: pathAnalyzerTestQueryOptionsRepresentation},
		"lifecycle":            acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreVnMonitoringChangesRep},
	}
	ignoreVnMonitoringChangesRep = map[string]interface{}{
		"ignore_changes": acctest.Representation{RepType: acctest.Required, Create: []string{`defined_tags`, `destination_endpoint`, `source_endpoint`}},
	}
	pathAnalyzerTestDestinationEndpointRepresentation = map[string]interface{}{
		"type":                     acctest.Representation{RepType: acctest.Required, Create: `COMPUTE_INSTANCE`},
		"address":                  acctest.Representation{RepType: acctest.Required, Create: `${oci_core_instance.test_instance.private_ip}`},
		"instance_id":              acctest.Representation{RepType: acctest.Required, Create: `${oci_core_instance.test_instance.id}`},
		"listener_id":              acctest.Representation{RepType: acctest.Optional, Create: `${oci_load_balancer_listener.test_listener.id}`},
		"load_balancer_id":         acctest.Representation{RepType: acctest.Optional, Create: `${oci_load_balancer_load_balancer.test_load_balancer.id}`},
		"network_load_balancer_id": acctest.Representation{RepType: acctest.Optional, Create: `${oci_network_load_balancer_network_load_balancer.test_network_load_balancer.id}`},
		"subnet_id":                acctest.Representation{RepType: acctest.Optional, Create: `${oci_core_subnet.test_subnet.id}`},
		"vlan_id":                  acctest.Representation{RepType: acctest.Optional, Create: `${oci_core_vlan.test_vlan.id}`},
		"vnic_id":                  acctest.Representation{RepType: acctest.Required, Create: `${oci_core_vnic_attachment.test_vnic_attachment.vnic_id}`},
	}
	pathAnalyzerTestSourceEndpointRepresentation = map[string]interface{}{
		"type":                     acctest.Representation{RepType: acctest.Required, Create: `COMPUTE_INSTANCE`},
		"address":                  acctest.Representation{RepType: acctest.Required, Create: `${oci_core_instance.test_instance2.private_ip}`},
		"instance_id":              acctest.Representation{RepType: acctest.Required, Create: `${oci_core_instance.test_instance2.id}`},
		"listener_id":              acctest.Representation{RepType: acctest.Optional, Create: `${oci_load_balancer_listener.test_listener.id}`},
		"load_balancer_id":         acctest.Representation{RepType: acctest.Optional, Create: `${oci_load_balancer_load_balancer.test_load_balancer.id}`},
		"network_load_balancer_id": acctest.Representation{RepType: acctest.Optional, Create: `${oci_network_load_balancer_network_load_balancer.test_network_load_balancer.id}`},
		"subnet_id":                acctest.Representation{RepType: acctest.Optional, Create: `${oci_core_subnet.test_subnet.id}`},
		"vlan_id":                  acctest.Representation{RepType: acctest.Optional, Create: `${oci_core_vlan.test_vlan.id}`},
		"vnic_id":                  acctest.Representation{RepType: acctest.Required, Create: `${oci_core_vnic_attachment.test_vnic_attachment2.vnic_id}`},
	}
	pathAnalyzerTestProtocolParametersRepresentation = map[string]interface{}{
		"type":             acctest.Representation{RepType: acctest.Required, Create: `ICMP`, Update: `UDP`},
		"destination_port": acctest.Representation{RepType: acctest.Required, Create: `00`, Update: `11`},
		"icmp_code":        acctest.Representation{RepType: acctest.Required, Create: `10`, Update: `00`},
		"icmp_type":        acctest.Representation{RepType: acctest.Required, Create: `10`, Update: `00`},
		"source_port":      acctest.Representation{RepType: acctest.Required, Create: `00`, Update: `11`},
	}
	pathAnalyzerTestQueryOptionsRepresentation = map[string]interface{}{
		"is_bi_directional_analysis": acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
	}
	instance2Representation = map[string]interface{}{
		"availability_domain":                 acctest.Representation{RepType: acctest.Required, Create: `${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}`},
		"compartment_id":                      acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"shape":                               acctest.Representation{RepType: acctest.Required, Create: `VM.Standard2.1`},
		"agent_config":                        acctest.RepresentationGroup{RepType: acctest.Optional, Group: CoreInstanceAgentConfigRepresentation},
		"availability_config":                 acctest.RepresentationGroup{RepType: acctest.Optional, Group: CoreInstanceAvailabilityConfigRepresentation},
		"create_vnic_details":                 acctest.RepresentationGroup{RepType: acctest.Optional, Group: CoreInstanceCreateVnicDetailsRepresentation},
		"dedicated_vm_host_id":                acctest.Representation{RepType: acctest.Optional, Create: `${oci_core_dedicated_vm_host.test_dedicated_vm_host.id}`},
		"defined_tags":                        acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"display_name":                        acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"fault_domain":                        acctest.Representation{RepType: acctest.Optional, Create: `FAULT-DOMAIN-3`},
		"freeform_tags":                       acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"hostname_label":                      acctest.Representation{RepType: acctest.Optional, Create: `hostnamelabel`},
		"instance_options":                    acctest.RepresentationGroup{RepType: acctest.Optional, Group: CoreInstanceInstanceOptionsRepresentation},
		"image":                               acctest.Representation{RepType: acctest.Required, Create: `${var.InstanceImageOCID[var.region]}`},
		"ipxe_script":                         acctest.Representation{RepType: acctest.Optional, Create: `ipxeScript`},
		"is_pv_encryption_in_transit_enabled": acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"launch_options":                      acctest.RepresentationGroup{RepType: acctest.Optional, Group: CoreInstanceLaunchOptionsRepresentation},
		"metadata":                            acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"user_data": "abcd"}, Update: map[string]string{"user_data": "abcd", "volatile_data": "stringE"}},
		"shape_config":                        acctest.RepresentationGroup{RepType: acctest.Optional, Group: CoreInstanceShapeConfigRepresentation},
		"source_details":                      acctest.RepresentationGroup{RepType: acctest.Optional, Group: CoreInstanceSourceDetailsRepresentation},
		"subnet_id":                           acctest.Representation{RepType: acctest.Required, Create: `${oci_core_subnet.test_subnet.id}`},
		"state":                               acctest.Representation{RepType: acctest.Optional, Create: `STOPPED`, Update: `RUNNING`},
	}
	vnicAttachment2Representation = map[string]interface{}{
		"create_vnic_details": acctest.RepresentationGroup{RepType: acctest.Required, Group: CoreVnicAttachmentCreateVnicDetailsRepresentation},
		"instance_id":         acctest.Representation{RepType: acctest.Required, Create: `${oci_core_instance.test_instance2.id}`},
		"display_name":        acctest.Representation{RepType: acctest.Optional, Create: `displayName`},
		"nic_index":           acctest.Representation{RepType: acctest.Optional, Create: `0`},
	}

	PathAnalyzerTestResourceDependencies = utils.OciImageIdsVariable +
		acctest.GenerateResourceFromRepresentationMap("oci_core_instance", "test_instance", acctest.Required, acctest.Create, CoreInstanceRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_instance", "test_instance2", acctest.Required, acctest.Create, instance2Representation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_subnet", "test_subnet", acctest.Required, acctest.Create, CoreSubnetRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", acctest.Required, acctest.Create, CoreVcnRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_vlan", "test_vlan", acctest.Required, acctest.Create, vlanRepresentationVnMonitoring) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_vnic_attachment", "test_vnic_attachment", acctest.Required, acctest.Create, CoreVnicAttachmentRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_vnic_attachment", "test_vnic_attachment2", acctest.Required, acctest.Create, vnicAttachment2Representation) +
		AvailabilityDomainConfig +
		DefinedTagsDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_load_balancer_backend_set", "test_backend_set", acctest.Required, acctest.Create, backendSetRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_load_balancer_listener", "test_listener", acctest.Required, acctest.Create, listenerRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_load_balancer_load_balancer", "test_load_balancer", acctest.Required, acctest.Create, loadBalancerRepresentation) +
		LoadBalancerSubnetDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_network_load_balancer_network_load_balancer", "test_network_load_balancer", acctest.Required, acctest.Create, NetworkLoadBalancerNetworkLoadBalancerRepresentation)
)

// issue-routing-tag: vn_monitoring/default
func TestVnMonitoringPathAnalyzerTestResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestVnMonitoringPathAnalyzerTestResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_vn_monitoring_path_analyzer_test.test_path_analyzer_test"
	datasourceName := "data.oci_vn_monitoring_path_analyzer_tests.test_path_analyzer_tests"
	singularDatasourceName := "data.oci_vn_monitoring_path_analyzer_test.test_path_analyzer_test"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+PathAnalyzerTestResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_vn_monitoring_path_analyzer_test", "test_path_analyzer_test", acctest.Optional, acctest.Create, pathAnalyzerTestRepresentation), "vnmonitoring", "pathAnalyzerTest", t)

	acctest.ResourceTest(t, testAccCheckVnMonitoringPathAnalyzerTestDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + PathAnalyzerTestResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_vn_monitoring_path_analyzer_test", "test_path_analyzer_test", acctest.Required, acctest.Create, pathAnalyzerTestRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "destination_endpoint.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "destination_endpoint.0.address"),
				resource.TestCheckResourceAttrSet(resourceName, "destination_endpoint.0.instance_id"),
				resource.TestCheckResourceAttr(resourceName, "destination_endpoint.0.type", "COMPUTE_INSTANCE"),
				resource.TestCheckResourceAttrSet(resourceName, "destination_endpoint.0.vnic_id"),
				resource.TestCheckResourceAttr(resourceName, "protocol", "1"),
				resource.TestCheckResourceAttr(resourceName, "source_endpoint.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "source_endpoint.0.address"),
				resource.TestCheckResourceAttrSet(resourceName, "source_endpoint.0.instance_id"),
				resource.TestCheckResourceAttr(resourceName, "source_endpoint.0.type", "COMPUTE_INSTANCE"),
				resource.TestCheckResourceAttrSet(resourceName, "source_endpoint.0.vnic_id"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + PathAnalyzerTestResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + PathAnalyzerTestResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_vn_monitoring_path_analyzer_test", "test_path_analyzer_test", acctest.Optional, acctest.Create, pathAnalyzerTestRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "destination_endpoint.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "destination_endpoint.0.address"),
				resource.TestCheckResourceAttrSet(resourceName, "destination_endpoint.0.instance_id"),
				resource.TestCheckResourceAttr(resourceName, "destination_endpoint.0.type", "COMPUTE_INSTANCE"),
				resource.TestCheckResourceAttrSet(resourceName, "destination_endpoint.0.vnic_id"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "Path Analyzer Test"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "protocol", "1"),
				resource.TestCheckResourceAttr(resourceName, "protocol_parameters.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "protocol_parameters.0.icmp_code", "10"),
				resource.TestCheckResourceAttr(resourceName, "protocol_parameters.0.icmp_type", "10"),
				resource.TestCheckResourceAttr(resourceName, "protocol_parameters.0.type", "ICMP"),
				resource.TestCheckResourceAttr(resourceName, "query_options.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "query_options.0.is_bi_directional_analysis", "false"),
				resource.TestCheckResourceAttr(resourceName, "source_endpoint.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "source_endpoint.0.address"),
				resource.TestCheckResourceAttrSet(resourceName, "source_endpoint.0.instance_id"),
				resource.TestCheckResourceAttr(resourceName, "source_endpoint.0.type", "COMPUTE_INSTANCE"),
				resource.TestCheckResourceAttrSet(resourceName, "source_endpoint.0.vnic_id"),
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

		// verify Update to the compartment (the compartment will be switched back in the next step)
		{
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + PathAnalyzerTestResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_vn_monitoring_path_analyzer_test", "test_path_analyzer_test", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(pathAnalyzerTestRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "destination_endpoint.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "destination_endpoint.0.address"),
				resource.TestCheckResourceAttrSet(resourceName, "destination_endpoint.0.instance_id"),
				resource.TestCheckResourceAttr(resourceName, "destination_endpoint.0.type", "COMPUTE_INSTANCE"),
				resource.TestCheckResourceAttrSet(resourceName, "destination_endpoint.0.vnic_id"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "Path Analyzer Test"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "protocol", "1"),
				resource.TestCheckResourceAttr(resourceName, "protocol_parameters.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "protocol_parameters.0.icmp_code", "10"),
				resource.TestCheckResourceAttr(resourceName, "protocol_parameters.0.icmp_type", "10"),
				resource.TestCheckResourceAttr(resourceName, "protocol_parameters.0.type", "ICMP"),
				resource.TestCheckResourceAttr(resourceName, "query_options.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "query_options.0.is_bi_directional_analysis", "false"),
				resource.TestCheckResourceAttr(resourceName, "source_endpoint.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "source_endpoint.0.address"),
				resource.TestCheckResourceAttrSet(resourceName, "source_endpoint.0.instance_id"),
				resource.TestCheckResourceAttr(resourceName, "source_endpoint.0.type", "COMPUTE_INSTANCE"),
				resource.TestCheckResourceAttrSet(resourceName, "source_endpoint.0.vnic_id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "time_updated"),

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
			Config: config + compartmentIdVariableStr + PathAnalyzerTestResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_vn_monitoring_path_analyzer_test", "test_path_analyzer_test", acctest.Optional, acctest.Update, pathAnalyzerTestRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "destination_endpoint.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "destination_endpoint.0.address"),
				resource.TestCheckResourceAttrSet(resourceName, "destination_endpoint.0.instance_id"),
				resource.TestCheckResourceAttr(resourceName, "destination_endpoint.0.type", "COMPUTE_INSTANCE"),
				resource.TestCheckResourceAttrSet(resourceName, "destination_endpoint.0.vnic_id"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "protocol", "17"),
				resource.TestCheckResourceAttr(resourceName, "protocol_parameters.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "protocol_parameters.0.destination_port", "11"),
				resource.TestCheckResourceAttr(resourceName, "protocol_parameters.0.source_port", "11"),
				resource.TestCheckResourceAttr(resourceName, "protocol_parameters.0.type", "UDP"),
				resource.TestCheckResourceAttr(resourceName, "query_options.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "query_options.0.is_bi_directional_analysis", "true"),
				resource.TestCheckResourceAttr(resourceName, "source_endpoint.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "source_endpoint.0.address"),
				resource.TestCheckResourceAttrSet(resourceName, "source_endpoint.0.instance_id"),
				resource.TestCheckResourceAttr(resourceName, "source_endpoint.0.type", "COMPUTE_INSTANCE"),
				resource.TestCheckResourceAttrSet(resourceName, "source_endpoint.0.vnic_id"),
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
				acctest.GenerateDataSourceFromRepresentationMap("oci_vn_monitoring_path_analyzer_tests", "test_path_analyzer_tests", acctest.Optional, acctest.Update, pathAnalyzerTestDataSourceRepresentation) +
				compartmentIdVariableStr + PathAnalyzerTestResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_vn_monitoring_path_analyzer_test", "test_path_analyzer_test", acctest.Optional, acctest.Update, pathAnalyzerTestRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),

				resource.TestCheckResourceAttr(datasourceName, "path_analyzer_test_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "path_analyzer_test_collection.0.items.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_vn_monitoring_path_analyzer_test", "test_path_analyzer_test", acctest.Required, acctest.Create, pathAnalyzerTestSingularDataSourceRepresentation) +
				compartmentIdVariableStr + PathAnalyzerTestResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "path_analyzer_test_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "destination_endpoint.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "destination_endpoint.0.address"),
				resource.TestCheckResourceAttr(singularDatasourceName, "destination_endpoint.0.type", "COMPUTE_INSTANCE"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "protocol", "17"),
				resource.TestCheckResourceAttr(singularDatasourceName, "protocol_parameters.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "protocol_parameters.0.destination_port", "11"),
				resource.TestCheckResourceAttr(singularDatasourceName, "protocol_parameters.0.source_port", "11"),
				resource.TestCheckResourceAttr(singularDatasourceName, "protocol_parameters.0.type", "UDP"),
				resource.TestCheckResourceAttr(singularDatasourceName, "query_options.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "query_options.0.is_bi_directional_analysis", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "source_endpoint.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "source_endpoint.0.address"),
				resource.TestCheckResourceAttr(singularDatasourceName, "source_endpoint.0.type", "COMPUTE_INSTANCE"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
		// verify resource import
		{
			Config:                  config + PathAnalyzerTestRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckVnMonitoringPathAnalyzerTestDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).VnMonitoringClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_vn_monitoring_path_analyzer_test" {
			noResourceFound = false
			request := oci_vn_monitoring.GetPathAnalyzerTestRequest{}

			tmp := rs.Primary.ID
			request.PathAnalyzerTestId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "vn_monitoring")

			response, err := client.GetPathAnalyzerTest(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_vn_monitoring.PathAnalyzerTestLifecycleStateDeleted): true,
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
	if !acctest.InSweeperExcludeList("VnMonitoringPathAnalyzerTest") {
		resource.AddTestSweepers("VnMonitoringPathAnalyzerTest", &resource.Sweeper{
			Name:         "VnMonitoringPathAnalyzerTest",
			Dependencies: acctest.DependencyGraph["pathAnalyzerTest"],
			F:            sweepVnMonitoringPathAnalyzerTestResource,
		})
	}
}

func sweepVnMonitoringPathAnalyzerTestResource(compartment string) error {
	vnMonitoringClient := acctest.GetTestClients(&schema.ResourceData{}).VnMonitoringClient()
	pathAnalyzerTestIds, err := getPathAnalyzerTestIds(compartment)
	if err != nil {
		return err
	}
	for _, pathAnalyzerTestId := range pathAnalyzerTestIds {
		if ok := acctest.SweeperDefaultResourceId[pathAnalyzerTestId]; !ok {
			deletePathAnalyzerTestRequest := oci_vn_monitoring.DeletePathAnalyzerTestRequest{}

			deletePathAnalyzerTestRequest.PathAnalyzerTestId = &pathAnalyzerTestId

			deletePathAnalyzerTestRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "vn_monitoring")
			_, error := vnMonitoringClient.DeletePathAnalyzerTest(context.Background(), deletePathAnalyzerTestRequest)
			if error != nil {
				fmt.Printf("Error deleting PathAnalyzerTest %s %s, It is possible that the resource is already deleted. Please verify manually \n", pathAnalyzerTestId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &pathAnalyzerTestId, pathAnalyzerTestSweepWaitCondition, time.Duration(3*time.Minute),
				pathAnalyzerTestSweepResponseFetchOperation, "vn_monitoring", true)
		}
	}
	return nil
}

func getPathAnalyzerTestIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "PathAnalyzerTestId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	vnMonitoringClient := acctest.GetTestClients(&schema.ResourceData{}).VnMonitoringClient()

	listPathAnalyzerTestsRequest := oci_vn_monitoring.ListPathAnalyzerTestsRequest{}
	listPathAnalyzerTestsRequest.CompartmentId = &compartmentId
	listPathAnalyzerTestsRequest.LifecycleState = oci_vn_monitoring.PathAnalyzerTestLifecycleStateActive
	listPathAnalyzerTestsResponse, err := vnMonitoringClient.ListPathAnalyzerTests(context.Background(), listPathAnalyzerTestsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting PathAnalyzerTest list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, pathAnalyzerTest := range listPathAnalyzerTestsResponse.Items {
		id := *pathAnalyzerTest.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "PathAnalyzerTestId", id)
	}
	return resourceIds, nil
}

func pathAnalyzerTestSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if pathAnalyzerTestResponse, ok := response.Response.(oci_vn_monitoring.GetPathAnalyzerTestResponse); ok {
		return pathAnalyzerTestResponse.LifecycleState != oci_vn_monitoring.PathAnalyzerTestLifecycleStateDeleted
	}
	return false
}

func pathAnalyzerTestSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.VnMonitoringClient().GetPathAnalyzerTest(context.Background(), oci_vn_monitoring.GetPathAnalyzerTestRequest{
		PathAnalyzerTestId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
