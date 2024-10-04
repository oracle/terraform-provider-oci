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
	oci_core "github.com/oracle/oci-go-sdk/v65/core"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	CoreVirtualCircuitRequiredOnlyResource = CoreVirtualCircuitResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_core_virtual_circuit", "test_virtual_circuit", acctest.Required, acctest.Create, CoreVirtualCircuitRequiredOnlyRepresentation)

	CoreVirtualCircuitRequiredOnlyResourceNoMacSec = VirtualCircuitResourceDependenciesCopyForVC +
		acctest.GenerateResourceFromRepresentationMap("oci_core_virtual_circuit", "test_virtual_circuit", acctest.Required, acctest.Create, CoreVirtualCircuitRequiredOnlyRepresentation)

	CoreCoreVirtualCircuitSingularDataSourceRepresentation = map[string]interface{}{
		"virtual_circuit_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_core_virtual_circuit.test_virtual_circuit.id}`},
	}

	CoreCoreVirtualCircuitDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":   acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"state":          acctest.Representation{RepType: acctest.Optional, Create: `PROVISIONED`},
		"filter":         acctest.RepresentationGroup{RepType: acctest.Required, Group: CoreVirtualCircuitDataSourceFilterRepresentation}}
	CoreVirtualCircuitDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_core_virtual_circuit.test_virtual_circuit.id}`}},
	}

	CoreVirtualCircuitPublicRequiredOnlyRepresentation = map[string]interface{}{
		"compartment_id":         acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"type":                   acctest.Representation{RepType: acctest.Required, Create: `${var.virtual_circuit_type}`},
		"cross_connect_mappings": acctest.RepresentationGroup{RepType: acctest.Required, Group: CoreVirtualCircuitCrossConnectMappingsPublicRequiredOnlyRepresentation},
		"customer_bgp_asn":       acctest.Representation{RepType: acctest.Required, Create: `10`, Update: `11`},
		"public_prefixes":        acctest.RepresentationGroup{RepType: acctest.Required, Group: CoreVirtualCircuitPublicPrefixesRepresentation},
	}
	CoreVirtualCircuitPublicRequiredOnlyWithoutDeprecatedRepresentation = acctest.RepresentationCopyWithNewProperties(
		acctest.RepresentationCopyWithRemovedProperties(CoreVirtualCircuitPublicRequiredOnlyRepresentation, []string{"customer_bgp_asn"}), map[string]interface{}{
			"customer_asn": acctest.Representation{RepType: acctest.Required, Create: `10`, Update: `11`},
		})

	CoreVirtualCircuitRequiredOnlyRepresentation = map[string]interface{}{
		"compartment_id":         acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"type":                   acctest.Representation{RepType: acctest.Required, Create: `${var.virtual_circuit_type}`},
		"cross_connect_mappings": acctest.RepresentationGroup{RepType: acctest.Required, Group: CoreVirtualCircuitCrossConnectMappingsRequiredOnlyRepresentation},
		"customer_asn":           acctest.Representation{RepType: acctest.Required, Create: `10`, Update: `11`},
		"gateway_id":             acctest.Representation{RepType: acctest.Required, Create: `${oci_core_drg.test_drg.id}`},
	}
	CoreVirtualCircuitRepresentation = map[string]interface{}{
		"ip_mtu":                 acctest.Representation{RepType: acctest.Optional, Create: `MTU_1500`, Update: `MTU_9000`},
		"compartment_id":         acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"type":                   acctest.Representation{RepType: acctest.Required, Create: `${var.virtual_circuit_type}`},
		"bandwidth_shape_name":   acctest.Representation{RepType: acctest.Optional, Create: `10 Gbps`, Update: `20 Gbps`},
		"cross_connect_mappings": acctest.RepresentationGroup{RepType: acctest.Required, Group: CoreVirtualCircuitCrossConnectMappingsRequiredOnlyRepresentation},
		"customer_asn":           acctest.Representation{RepType: acctest.Required, Create: `10`, Update: `11`},
		"defined_tags":           acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"display_name":           acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"freeform_tags":          acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"gateway_id":             acctest.Representation{RepType: acctest.Optional, Create: `${oci_core_drg.test_drg.id}`},
		"region":                 acctest.Representation{RepType: acctest.Optional, Create: `us-phoenix-1`},
		"routing_policy":         acctest.Representation{RepType: acctest.Optional, Create: []string{`REGIONAL`}, Update: []string{`GLOBAL`}},
		"bgp_admin_state":        acctest.Representation{RepType: acctest.Optional, Create: `ENABLED`, Update: `DISABLED`},
		"is_bfd_enabled":         acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"is_transport_mode":      acctest.Representation{RepType: acctest.Optional, Create: `false`},
	}
	CoreVirtualCircuitRepresentationTransportMode = map[string]interface{}{
		"ip_mtu":                 acctest.Representation{RepType: acctest.Optional, Create: `MTU_1500`, Update: `MTU_9000`},
		"compartment_id":         acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"type":                   acctest.Representation{RepType: acctest.Required, Create: `${var.virtual_circuit_type}`},
		"bandwidth_shape_name":   acctest.Representation{RepType: acctest.Optional, Create: `10 Gbps`, Update: `20 Gbps`},
		"cross_connect_mappings": acctest.RepresentationGroup{RepType: acctest.Required, Group: CoreVirtualCircuitCrossConnectMappingsRequiredOnlyRepresentation},
		"customer_asn":           acctest.Representation{RepType: acctest.Required, Create: `10`, Update: `11`},
		"defined_tags":           acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"display_name":           acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"freeform_tags":          acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"gateway_id":             acctest.Representation{RepType: acctest.Optional, Create: `${var.drg_ocid_transport_mode}`},
		"region":                 acctest.Representation{RepType: acctest.Optional, Create: `us-phoenix-1`},
		"routing_policy":         acctest.Representation{RepType: acctest.Optional, Create: []string{`REGIONAL`}, Update: []string{`GLOBAL`}},
		"bgp_admin_state":        acctest.Representation{RepType: acctest.Optional, Create: `ENABLED`, Update: `DISABLED`},
		"is_bfd_enabled":         acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"is_transport_mode":      acctest.Representation{RepType: acctest.Optional, Create: `true`},
	}

	CoreVirtualCircuitWithProviderRepresentation = map[string]interface{}{
		"compartment_id":         acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"type":                   acctest.Representation{RepType: acctest.Required, Create: `${var.virtual_circuit_type}`},
		"bandwidth_shape_name":   acctest.Representation{RepType: acctest.Optional, Create: "${data.oci_core_virtual_circuit_bandwidth_shapes.test_virtual_circuit_bandwidth_shapes_layer2.virtual_circuit_bandwidth_shapes.0.name}"},
		"cross_connect_mappings": acctest.RepresentationGroup{RepType: acctest.Required, Group: CoreVirtualCircuitCrossConnectMappingsRepresentation},
		"customer_asn":           acctest.Representation{RepType: acctest.Required, Create: `10`, Update: `11`},
		"display_name":           acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"gateway_id":             acctest.Representation{RepType: acctest.Optional, Create: `${oci_core_drg.test_drg.id}`},
		"provider_service_id":    acctest.Representation{RepType: acctest.Optional, Create: `${data.oci_core_fast_connect_provider_services.test_fast_connect_provider_services_layer2.fast_connect_provider_services.0.id}`},
		// provider_service_key_name can only be updated by a Fast Connect Service Provider
		// "provider_service_key_name": acctest.Representation{RepType: acctest.Optional, Create: `d8f7a443-28c2-4dcf-996c-286351908c58`},
		"region": acctest.Representation{RepType: acctest.Optional, Create: `us-phoenix-1`},
	}
	CorePrivateVirtualCircuitUpdateRepresentation = acctest.RepresentationCopyWithNewProperties(
		acctest.RepresentationCopyWithRemovedProperties(CoreVirtualCircuitRequiredOnlyRepresentation, []string{"cross_connect_mappings"}), map[string]interface{}{
			"cross_connect_mappings": acctest.RepresentationGroup{RepType: acctest.Required, Group: CoreVirtualCircuitCrossConnectMappingsRepresentation},
		})

	CoreVirtualCircuitWithProviderLayer3Representation = acctest.RepresentationCopyWithNewProperties(
		acctest.RepresentationCopyWithRemovedProperties(CoreVirtualCircuitRequiredOnlyRepresentation, []string{"cross_connect_mappings", "bandwidth_shape_name", "provider_service_id", "customer_asn"}), map[string]interface{}{
			"bandwidth_shape_name": acctest.Representation{RepType: acctest.Optional, Create: "${data.oci_core_virtual_circuit_bandwidth_shapes.test_virtual_circuit_bandwidth_shapes_layer3.virtual_circuit_bandwidth_shapes.0.name}"},
			"provider_service_id":  acctest.Representation{RepType: acctest.Optional, Create: `${data.oci_core_fast_connect_provider_services.test_fast_connect_provider_services_layer3.fast_connect_provider_services.0.id}`},
		})

	CoreVirtualCircuitCrossConnectMappingsPublicRequiredOnlyRepresentation = map[string]interface{}{
		"cross_connect_or_cross_connect_group_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_core_cross_connect.test_cross_connect.cross_connect_group_id}`},
		"vlan": acctest.Representation{RepType: acctest.Required, Create: `200`, Update: `300`},
	}
	CoreVirtualCircuitCrossConnectMappingsRequiredOnlyRepresentation = map[string]interface{}{
		"cross_connect_or_cross_connect_group_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_core_cross_connect.test_cross_connect.cross_connect_group_id}`},
		"customer_bgp_peering_ip":                 acctest.Representation{RepType: acctest.Required, Create: `10.0.0.18/31`, Update: `10.0.0.20/31`},
		"oracle_bgp_peering_ip":                   acctest.Representation{RepType: acctest.Required, Create: `10.0.0.19/31`, Update: `10.0.0.21/31`},
		"vlan":                                    acctest.Representation{RepType: acctest.Required, Create: `200`, Update: `300`},
	}
	CoreVirtualCircuitCrossConnectMappingsRepresentation = map[string]interface{}{
		"customer_bgp_peering_ip": acctest.Representation{RepType: acctest.Required, Create: `10.0.0.18/31`, Update: `10.0.0.20/31`},
		"oracle_bgp_peering_ip":   acctest.Representation{RepType: acctest.Required, Create: `10.0.0.19/31`, Update: `10.0.0.21/31`},
	}
	CoreVirtualCircuitPublicPrefixesRepresentation = map[string]interface{}{
		"cidr_block": acctest.Representation{RepType: acctest.Required, Create: `11.0.0.0/24`, Update: `11.0.1.0/24`},
	}

	VirtualCircuitWithProviderResourceConfigFilter = `
data "oci_core_fast_connect_provider_services" "test_fast_connect_provider_services_layer2" {
	#Required
	compartment_id = "${var.compartment_id}"

	filter {
		name = "type"
		values = [ "LAYER2" ]
	}

	filter {
		name = "provider_name"
		values = ["OracleL2IntegDeployment"]
	}

	filter {
		name = "private_peering_bgp_management"
		values = [ "CUSTOMER_MANAGED" ]
	}

	filter {
		name = "supported_virtual_circuit_types"
		values = [ "${var.virtual_circuit_type}" ]
	}

	filter {
		name = "public_peering_bgp_management"
		values = [ "ORACLE_MANAGED" ]
	}

	filter {
		name = "provider_service_key_management"
		values = ["PROVIDER_MANAGED"]
	}
}

data "oci_core_fast_connect_provider_services" "test_fast_connect_provider_services_layer3" {
	#Required
	compartment_id = "${var.compartment_id}"

	filter {
		name = "type"
		values = [ "LAYER3" ]
	}

	filter {
		name = "provider_name"
		values = ["OracleL3IntegDeployment"]
	}

	filter {
		name = "supported_virtual_circuit_types"
		values = [ "${var.virtual_circuit_type}" ]
	}

	filter {
		name = "public_peering_bgp_management"
		values = [ "ORACLE_MANAGED" ]
	}

	filter {
		name = "provider_service_key_management"
		values = ["PROVIDER_MANAGED"]
	}
}

data "oci_core_virtual_circuit_bandwidth_shapes" "test_virtual_circuit_bandwidth_shapes_layer2" {
  #Required
  provider_service_id = "${data.oci_core_fast_connect_provider_services.test_fast_connect_provider_services_layer2.fast_connect_provider_services.0.id}"
}

data "oci_core_virtual_circuit_bandwidth_shapes" "test_virtual_circuit_bandwidth_shapes_layer3" {
  #Required
  provider_service_id = "${data.oci_core_fast_connect_provider_services.test_fast_connect_provider_services_layer3.fast_connect_provider_services.0.id}"
}
`

	VirtualCircuitPrivatePropertyVariables = `
variable "virtual_circuit_type" { default = "PRIVATE" }

`

	VirtualCircuitPublicPropertyVariables = `
variable "virtual_circuit_type" { default = "PUBLIC" }

`
	// VC dependencies with Macsec capabilities
	CoreVirtualCircuitResourceDependencies = CoreDrgRequiredOnlyResource + CrossConnectWithGroupResourceConfig
	// VC dependencies without Macsec capabilities
	VirtualCircuitResourceDependenciesCopyForVC = CoreDrgRequiredOnlyResource + CrossConnectWithGroupResourceConfigCopyForVC
)

// issue-routing-tag: core/default
func TestCoreVirtualCircuitResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestCoreVirtualCircuitResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	secretIdCKN := utils.GetEnvSettingWithBlankDefault("secret_ocid_ckn")
	secretIdVariableStrCKN := fmt.Sprintf("variable \"secret_ocid_ckn\" { default = \"%s\" }\n", secretIdCKN)

	secretIdCAK := utils.GetEnvSettingWithBlankDefault("secret_ocid_cak")
	secretIdVariableStrCAK := fmt.Sprintf("variable \"secret_ocid_cak\" { default = \"%s\" }\n", secretIdCAK)

	secretVersionCAK := utils.GetEnvSettingWithBlankDefault("secret_version_cak")
	secretVersionStrCAK := fmt.Sprintf("variable \"secret_version_cak\" { default = \"%s\" }\n", secretVersionCAK)

	secretVersionCKN := utils.GetEnvSettingWithBlankDefault("secret_version_ckn")
	secretVersionStrCKN := fmt.Sprintf("variable \"secret_version_ckn\" { default = \"%s\" }\n", secretVersionCKN)

	secretVersionCAKU := utils.GetEnvSettingWithBlankDefault("secret_version_cak_for_update")
	secretVersionStrCAKU := fmt.Sprintf("variable \"secret_version_cak_for_update\" { default = \"%s\" }\n", secretVersionCAKU)

	drgTransportMode := utils.GetEnvSettingWithBlankDefault("drg_ocid_transport_mode")
	drgTransportModeStr := fmt.Sprintf("variable \"drg_ocid_transport_mode\" { default = \"%s\" }\n", drgTransportMode)

	resourceName := "oci_core_virtual_circuit.test_virtual_circuit"
	datasourceName := "data.oci_core_virtual_circuits.test_virtual_circuits"
	singularDatasourceName := "data.oci_core_virtual_circuit.test_virtual_circuit"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+CoreVirtualCircuitResourceDependencies+VirtualCircuitPublicPropertyVariables+secretIdVariableStrCKN+secretIdVariableStrCAK+secretVersionStrCAK+secretVersionStrCKN+secretVersionStrCAKU+
		acctest.GenerateResourceFromRepresentationMap("oci_core_virtual_circuit", "test_virtual_circuit", acctest.Required, acctest.Create, CoreVirtualCircuitPublicRequiredOnlyRepresentation), "core", "virtualCircuit", t)

	acctest.ResourceTest(t, testAccCheckCoreVirtualCircuitDestroy, []resource.TestStep{
		// verify Create - PUBLIC Virtual Circuit
		{
			Config: config + compartmentIdVariableStr + VirtualCircuitResourceDependenciesCopyForVC + VirtualCircuitPublicPropertyVariables + secretIdVariableStrCKN + secretIdVariableStrCAK + secretVersionStrCAK + secretVersionStrCKN + secretVersionStrCAKU +
				acctest.GenerateResourceFromRepresentationMap("oci_core_virtual_circuit", "test_virtual_circuit", acctest.Required, acctest.Create, CoreVirtualCircuitPublicRequiredOnlyRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "cross_connect_mappings.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "cross_connect_mappings.0.cross_connect_or_cross_connect_group_id"),
				resource.TestCheckResourceAttr(resourceName, "cross_connect_mappings.0.vlan", "200"),
				resource.TestCheckResourceAttr(resourceName, "customer_bgp_asn", "10"),
				resource.TestCheckResourceAttr(resourceName, "public_prefixes.#", "1"),
				acctest.CheckResourceSetContainsElementWithProperties(resourceName, "public_prefixes", map[string]string{
					"cidr_block": "11.0.0.0/24",
				},
					[]string{}),
				resource.TestCheckResourceAttr(resourceName, "type", "PUBLIC"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},
		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + VirtualCircuitResourceDependenciesCopyForVC + secretIdVariableStrCKN + secretIdVariableStrCAK + secretVersionStrCAK + secretVersionStrCKN,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + VirtualCircuitResourceDependenciesCopyForVC + VirtualCircuitPublicPropertyVariables + secretIdVariableStrCKN + secretIdVariableStrCAK + secretVersionStrCAK + secretVersionStrCKN +
				acctest.GenerateResourceFromRepresentationMap("oci_core_virtual_circuit", "test_virtual_circuit", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(acctest.RepresentationCopyWithRemovedProperties(CoreVirtualCircuitRepresentation, []string{"gateway_id", "cross_connect_mappings", "customer_asn"}),
						map[string]interface{}{
							"cross_connect_mappings": acctest.RepresentationGroup{RepType: acctest.Required, Group: CoreVirtualCircuitCrossConnectMappingsPublicRequiredOnlyRepresentation},
							"customer_bgp_asn":       acctest.Representation{RepType: acctest.Required, Create: `10`, Update: `11`},
						})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "cross_connect_mappings.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "customer_bgp_asn", "10"),
				resource.TestCheckResourceAttr(resourceName, "type", "PUBLIC"),
				resource.TestCheckResourceAttr(resourceName, "routing_policy.#", "1"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},
		// verify Update from customer_bgp_asn to customer_asn
		{
			Config: config + compartmentIdVariableStr + VirtualCircuitResourceDependenciesCopyForVC + VirtualCircuitPublicPropertyVariables + secretIdVariableStrCKN + secretIdVariableStrCAK + secretVersionStrCAK + secretVersionStrCKN +
				acctest.GenerateResourceFromRepresentationMap("oci_core_virtual_circuit", "test_virtual_circuit", acctest.Required, acctest.Create, CoreVirtualCircuitPublicRequiredOnlyWithoutDeprecatedRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "cross_connect_mappings.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "cross_connect_mappings.0.cross_connect_or_cross_connect_group_id"),
				resource.TestCheckResourceAttr(resourceName, "cross_connect_mappings.0.vlan", "200"),
				resource.TestCheckResourceAttr(resourceName, "customer_asn", "10"),
				resource.TestCheckResourceAttr(resourceName, "public_prefixes.#", "1"),
				acctest.CheckResourceSetContainsElementWithProperties(resourceName, "public_prefixes", map[string]string{
					"cidr_block": "11.0.0.0/24",
				},
					[]string{}),
				resource.TestCheckResourceAttr(resourceName, "type", "PUBLIC"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("Resource recreated when it was supposed to be updated.")
					}
					return err
				},
			),
		},
		// verify Update - PUBLIC Virtual Circuit
		{
			Config: config + compartmentIdVariableStr + VirtualCircuitResourceDependenciesCopyForVC + VirtualCircuitPublicPropertyVariables + secretIdVariableStrCKN + secretIdVariableStrCAK + secretVersionStrCAK + secretVersionStrCKN +
				acctest.GenerateResourceFromRepresentationMap("oci_core_virtual_circuit", "test_virtual_circuit", acctest.Required, acctest.Update, CoreVirtualCircuitPublicRequiredOnlyWithoutDeprecatedRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "cross_connect_mappings.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "cross_connect_mappings.0.cross_connect_or_cross_connect_group_id"),
				resource.TestCheckResourceAttr(resourceName, "cross_connect_mappings.0.vlan", "300"),
				resource.TestCheckResourceAttr(resourceName, "customer_asn", "11"),
				resource.TestCheckResourceAttr(resourceName, "public_prefixes.#", "1"),
				acctest.CheckResourceSetContainsElementWithProperties(resourceName, "public_prefixes", map[string]string{
					"cidr_block": "11.0.1.0/24",
				},
					[]string{}),
				resource.TestCheckResourceAttr(resourceName, "type", "PUBLIC"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("Resource recreated when it was supposed to be updated.")
					}
					return err
				},
			),
		},
		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + VirtualCircuitResourceDependenciesCopyForVC + secretIdVariableStrCKN + secretIdVariableStrCAK + secretVersionStrCAK + secretVersionStrCKN,
		},
		// verify Create - PRIVATE Virtual Circuit with Provider Layer 2
		{
			Config: config + compartmentIdVariableStr + VirtualCircuitResourceDependenciesCopyForVC + VirtualCircuitPrivatePropertyVariables + VirtualCircuitWithProviderResourceConfigFilter +
				acctest.GenerateResourceFromRepresentationMap("oci_core_virtual_circuit", "test_virtual_circuit", acctest.Optional, acctest.Create,
					CoreVirtualCircuitWithProviderRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "cross_connect_mappings.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "cross_connect_mappings.0.customer_bgp_peering_ip", "10.0.0.18/31"),
				resource.TestCheckResourceAttr(resourceName, "cross_connect_mappings.0.oracle_bgp_peering_ip", "10.0.0.19/31"),
				resource.TestCheckResourceAttr(resourceName, "customer_asn", "10"),
				resource.TestCheckResourceAttrSet(resourceName, "gateway_id"),
				resource.TestCheckResourceAttrSet(resourceName, "provider_service_id"),
				resource.TestCheckResourceAttr(resourceName, "provider_state", "INACTIVE"),
				resource.TestCheckResourceAttr(resourceName, "type", "PRIVATE"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},
		// verify Update - PRIVATE Virtual Circuit with Provider Layer 2
		{
			Config: config + compartmentIdVariableStr + VirtualCircuitResourceDependenciesCopyForVC + VirtualCircuitPrivatePropertyVariables + VirtualCircuitWithProviderResourceConfigFilter + secretIdVariableStrCKN + secretIdVariableStrCAK + secretVersionStrCAK + secretVersionStrCKN +
				acctest.GenerateResourceFromRepresentationMap("oci_core_virtual_circuit", "test_virtual_circuit", acctest.Optional, acctest.Update, CoreVirtualCircuitWithProviderRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "cross_connect_mappings.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "cross_connect_mappings.0.customer_bgp_peering_ip", "10.0.0.20/31"),
				resource.TestCheckResourceAttr(resourceName, "cross_connect_mappings.0.oracle_bgp_peering_ip", "10.0.0.21/31"),
				resource.TestCheckResourceAttr(resourceName, "customer_asn", "11"),
				resource.TestCheckResourceAttrSet(resourceName, "gateway_id"),
				resource.TestCheckResourceAttrSet(resourceName, "provider_service_id"),
				resource.TestCheckResourceAttr(resourceName, "provider_state", "INACTIVE"),
				resource.TestCheckResourceAttr(resourceName, "type", "PRIVATE"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("Resource recreated when it was supposed to be updated.")
					}
					return err
				},
			),
		},
		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + VirtualCircuitResourceDependenciesCopyForVC + secretIdVariableStrCKN + secretIdVariableStrCAK + secretVersionStrCAK + secretVersionStrCKN,
		},
		// verify Create - PRIVATE Virtual Circuit with Provider Layer 3
		{
			Config: config + compartmentIdVariableStr + VirtualCircuitResourceDependenciesCopyForVC + VirtualCircuitPrivatePropertyVariables + VirtualCircuitWithProviderResourceConfigFilter +
				acctest.GenerateResourceFromRepresentationMap("oci_core_virtual_circuit", "test_virtual_circuit", acctest.Optional, acctest.Create,
					CoreVirtualCircuitWithProviderLayer3Representation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "cross_connect_mappings.#", "0"),
				resource.TestCheckResourceAttrSet(resourceName, "gateway_id"),
				resource.TestCheckResourceAttrSet(resourceName, "provider_service_id"),
				resource.TestCheckResourceAttr(resourceName, "provider_state", "INACTIVE"),
				resource.TestCheckResourceAttr(resourceName, "type", "PRIVATE"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},
		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + VirtualCircuitResourceDependenciesCopyForVC + secretIdVariableStrCKN + secretIdVariableStrCAK + secretVersionStrCAK + secretVersionStrCKN,
		},

		// verify Create - PRIVATE Virtual Circuit
		{
			Config: config + VirtualCircuitPrivatePropertyVariables + compartmentIdVariableStr + CoreVirtualCircuitRequiredOnlyResourceNoMacSec + secretIdVariableStrCKN + secretIdVariableStrCAK + secretVersionStrCAK + secretVersionStrCKN + secretVersionStrCAKU,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "cross_connect_mappings.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "cross_connect_mappings.0.cross_connect_or_cross_connect_group_id"),
				resource.TestCheckResourceAttr(resourceName, "cross_connect_mappings.0.customer_bgp_peering_ip", "10.0.0.18/31"),
				resource.TestCheckResourceAttr(resourceName, "cross_connect_mappings.0.oracle_bgp_peering_ip", "10.0.0.19/31"),
				resource.TestCheckResourceAttr(resourceName, "cross_connect_mappings.0.vlan", "200"),
				resource.TestCheckResourceAttr(resourceName, "customer_asn", "10"),
				resource.TestCheckResourceAttrSet(resourceName, "gateway_id"),
				resource.TestCheckResourceAttr(resourceName, "type", "PRIVATE"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},
		// verify Update - PRIVATE Virtual Circuit
		{
			Config: config + VirtualCircuitPrivatePropertyVariables + compartmentIdVariableStr + secretIdVariableStrCKN + secretIdVariableStrCAK + secretVersionStrCAK + secretVersionStrCKN + secretVersionStrCAKU +
				VirtualCircuitResourceDependenciesCopyForVC +
				acctest.GenerateResourceFromRepresentationMap("oci_core_virtual_circuit", "test_virtual_circuit", acctest.Required, acctest.Update, CorePrivateVirtualCircuitUpdateRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "cross_connect_mappings.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "cross_connect_mappings.0.cross_connect_or_cross_connect_group_id"),
				resource.TestCheckResourceAttr(resourceName, "cross_connect_mappings.0.customer_bgp_peering_ip", "10.0.0.20/31"),
				resource.TestCheckResourceAttr(resourceName, "cross_connect_mappings.0.oracle_bgp_peering_ip", "10.0.0.21/31"),
				resource.TestCheckResourceAttr(resourceName, "cross_connect_mappings.0.vlan", "200"),
				resource.TestCheckResourceAttr(resourceName, "customer_asn", "11"),
				resource.TestCheckResourceAttrSet(resourceName, "gateway_id"),
				resource.TestCheckResourceAttr(resourceName, "type", "PRIVATE"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + VirtualCircuitResourceDependenciesCopyForVC + secretIdVariableStrCKN + secretIdVariableStrCAK + secretVersionStrCAK + secretVersionStrCKN,
		},
		// verify is_transport_mode flag set to true
		{
			Config: config + compartmentIdVariableStr + VirtualCircuitResourceDependenciesCopyForVC + VirtualCircuitPrivatePropertyVariables + secretIdVariableStrCKN + secretIdVariableStrCAK + secretVersionStrCAK + secretVersionStrCKN + drgTransportModeStr +
				acctest.GenerateResourceFromRepresentationMap("oci_core_virtual_circuit", "test_virtual_circuit", acctest.Optional, acctest.Create, CoreVirtualCircuitRepresentationTransportMode),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "is_transport_mode", "true"),
				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},
		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + VirtualCircuitResourceDependenciesCopyForVC + secretIdVariableStrCKN + secretIdVariableStrCAK + secretVersionStrCAK + secretVersionStrCKN,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + VirtualCircuitResourceDependenciesCopyForVC + VirtualCircuitPrivatePropertyVariables + secretIdVariableStrCKN + secretIdVariableStrCAK + secretVersionStrCAK + secretVersionStrCKN + VirtualCircuitWithProviderResourceConfigFilter +
				acctest.GenerateResourceFromRepresentationMap("oci_core_virtual_circuit", "test_virtual_circuit", acctest.Optional, acctest.Create, CoreVirtualCircuitRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "bandwidth_shape_name", "10 Gbps"),
				resource.TestCheckResourceAttr(resourceName, "bgp_admin_state", "ENABLED"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "cross_connect_mappings.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "cross_connect_mappings.0.cross_connect_or_cross_connect_group_id"),
				resource.TestCheckResourceAttr(resourceName, "cross_connect_mappings.0.customer_bgp_peering_ip", "10.0.0.18/31"),
				resource.TestCheckResourceAttr(resourceName, "cross_connect_mappings.0.oracle_bgp_peering_ip", "10.0.0.19/31"),
				resource.TestCheckResourceAttr(resourceName, "cross_connect_mappings.0.vlan", "200"),
				resource.TestCheckResourceAttr(resourceName, "customer_asn", "10"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "gateway_id"),
				resource.TestCheckResourceAttr(resourceName, "ip_mtu", "MTU_1500"),
				resource.TestCheckResourceAttr(resourceName, "is_bfd_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "is_transport_mode", "false"),
				resource.TestCheckResourceAttr(resourceName, "region", "us-phoenix-1"),
				resource.TestCheckResourceAttr(resourceName, "routing_policy.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "type", "PRIVATE"),

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
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + VirtualCircuitResourceDependenciesCopyForVC + VirtualCircuitPrivatePropertyVariables + secretIdVariableStrCKN + secretIdVariableStrCAK + secretVersionStrCAK + secretVersionStrCKN +
				acctest.GenerateResourceFromRepresentationMap("oci_core_virtual_circuit", "test_virtual_circuit", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(CoreVirtualCircuitRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "bandwidth_shape_name", "10 Gbps"),
				resource.TestCheckResourceAttr(resourceName, "bgp_admin_state", "ENABLED"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "cross_connect_mappings.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "cross_connect_mappings.0.cross_connect_or_cross_connect_group_id"),
				resource.TestCheckResourceAttr(resourceName, "cross_connect_mappings.0.customer_bgp_peering_ip", "10.0.0.18/31"),
				resource.TestCheckResourceAttr(resourceName, "cross_connect_mappings.0.oracle_bgp_peering_ip", "10.0.0.19/31"),
				resource.TestCheckResourceAttr(resourceName, "cross_connect_mappings.0.vlan", "200"),
				resource.TestCheckResourceAttr(resourceName, "customer_asn", "10"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttrSet(resourceName, "gateway_id"),
				resource.TestCheckResourceAttr(resourceName, "ip_mtu", "MTU_1500"),
				resource.TestCheckResourceAttr(resourceName, "is_bfd_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "is_transport_mode", "false"),
				resource.TestCheckResourceAttr(resourceName, "region", "us-phoenix-1"),
				resource.TestCheckResourceAttr(resourceName, "routing_policy.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "type", "PRIVATE"),

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
			Config: config + compartmentIdVariableStr + VirtualCircuitResourceDependenciesCopyForVC + VirtualCircuitPrivatePropertyVariables + secretIdVariableStrCKN + secretIdVariableStrCAK + secretVersionStrCAK + secretVersionStrCKN +
				acctest.GenerateResourceFromRepresentationMap("oci_core_virtual_circuit", "test_virtual_circuit", acctest.Optional, acctest.Update, CoreVirtualCircuitRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "bandwidth_shape_name", "20 Gbps"),
				resource.TestCheckResourceAttr(resourceName, "bgp_admin_state", "DISABLED"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "cross_connect_mappings.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "cross_connect_mappings.0.cross_connect_or_cross_connect_group_id"),
				resource.TestCheckResourceAttr(resourceName, "cross_connect_mappings.0.customer_bgp_peering_ip", "10.0.0.20/31"),
				resource.TestCheckResourceAttr(resourceName, "cross_connect_mappings.0.oracle_bgp_peering_ip", "10.0.0.21/31"),
				resource.TestCheckResourceAttr(resourceName, "cross_connect_mappings.0.vlan", "300"),
				resource.TestCheckResourceAttr(resourceName, "customer_asn", "11"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "gateway_id"),
				resource.TestCheckResourceAttr(resourceName, "ip_mtu", "MTU_9000"),
				resource.TestCheckResourceAttr(resourceName, "is_bfd_enabled", "true"),
				resource.TestCheckResourceAttr(resourceName, "region", "us-phoenix-1"),
				resource.TestCheckResourceAttr(resourceName, "routing_policy.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "type", "PRIVATE"),

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
			Config: config + acctest.GenerateDataSourceFromRepresentationMap("oci_core_virtual_circuits", "test_virtual_circuits", acctest.Optional, acctest.Update, CoreCoreVirtualCircuitDataSourceRepresentation) +
				compartmentIdVariableStr + VirtualCircuitResourceDependenciesCopyForVC + VirtualCircuitPrivatePropertyVariables + secretIdVariableStrCKN + secretIdVariableStrCAK + secretVersionStrCAK + secretVersionStrCKN +
				acctest.GenerateResourceFromRepresentationMap("oci_core_virtual_circuit", "test_virtual_circuit", acctest.Optional, acctest.Update, CoreVirtualCircuitRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),

				resource.TestCheckResourceAttr(datasourceName, "virtual_circuits.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "virtual_circuits.0.bandwidth_shape_name", "20 Gbps"),
				resource.TestCheckResourceAttr(datasourceName, "virtual_circuits.0.bgp_admin_state", "DISABLED"),
				resource.TestCheckResourceAttrSet(datasourceName, "virtual_circuits.0.bgp_ipv6session_state"),
				resource.TestCheckResourceAttrSet(datasourceName, "virtual_circuits.0.bgp_management"),
				resource.TestCheckResourceAttrSet(datasourceName, "virtual_circuits.0.bgp_session_state"),
				resource.TestCheckResourceAttr(datasourceName, "virtual_circuits.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "virtual_circuits.0.cross_connect_mappings.#", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "virtual_circuits.0.cross_connect_mappings.0.cross_connect_or_cross_connect_group_id"),
				resource.TestCheckResourceAttr(datasourceName, "virtual_circuits.0.cross_connect_mappings.0.customer_bgp_peering_ip", "10.0.0.20/31"),
				resource.TestCheckResourceAttr(datasourceName, "virtual_circuits.0.cross_connect_mappings.0.oracle_bgp_peering_ip", "10.0.0.21/31"),
				resource.TestCheckResourceAttr(datasourceName, "virtual_circuits.0.cross_connect_mappings.0.vlan", "300"),
				resource.TestCheckResourceAttr(datasourceName, "virtual_circuits.0.customer_asn", "11"),
				resource.TestCheckResourceAttr(datasourceName, "virtual_circuits.0.display_name", "displayName2"),
				resource.TestCheckResourceAttr(datasourceName, "virtual_circuits.0.freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "virtual_circuits.0.gateway_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "virtual_circuits.0.id"),
				resource.TestCheckResourceAttr(datasourceName, "virtual_circuits.0.ip_mtu", "MTU_9000"),
				resource.TestCheckResourceAttr(datasourceName, "virtual_circuits.0.is_bfd_enabled", "true"),
				resource.TestCheckResourceAttrSet(datasourceName, "virtual_circuits.0.oracle_bgp_asn"),
				resource.TestCheckResourceAttr(datasourceName, "virtual_circuits.0.routing_policy.#", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "virtual_circuits.0.service_type"),
				resource.TestCheckResourceAttrSet(datasourceName, "virtual_circuits.0.state"),
				resource.TestCheckResourceAttrSet(datasourceName, "virtual_circuits.0.time_created"),
				resource.TestCheckResourceAttr(datasourceName, "virtual_circuits.0.state", "PROVISIONED"),
				resource.TestCheckResourceAttr(datasourceName, "virtual_circuits.0.type", "PRIVATE"),
				resource.TestCheckResourceAttr(datasourceName, "virtual_circuits.0.virtual_circuit_redundancy_metadata.#", "0"),
			),
		},
		// verify singular datasource
		{
			Config: config + acctest.GenerateDataSourceFromRepresentationMap("oci_core_virtual_circuit", "test_virtual_circuit", acctest.Required, acctest.Create, CoreCoreVirtualCircuitSingularDataSourceRepresentation) +
				compartmentIdVariableStr + VirtualCircuitResourceDependenciesCopyForVC + VirtualCircuitPrivatePropertyVariables + secretIdVariableStrCKN + secretIdVariableStrCAK + secretVersionStrCAK + secretVersionStrCKN +
				acctest.GenerateResourceFromRepresentationMap("oci_core_virtual_circuit", "test_virtual_circuit", acctest.Optional, acctest.Update, CoreVirtualCircuitRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "gateway_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "virtual_circuit_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "bandwidth_shape_name", "20 Gbps"),
				resource.TestCheckResourceAttr(singularDatasourceName, "bgp_admin_state", "DISABLED"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "bgp_ipv6session_state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "bgp_management"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "bgp_session_state"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "cross_connect_mappings.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "cross_connect_mappings.0.customer_bgp_peering_ip", "10.0.0.20/31"),
				resource.TestCheckResourceAttr(singularDatasourceName, "cross_connect_mappings.0.oracle_bgp_peering_ip", "10.0.0.21/31"),
				resource.TestCheckResourceAttr(singularDatasourceName, "cross_connect_mappings.0.vlan", "300"),
				resource.TestCheckResourceAttr(singularDatasourceName, "customer_asn", "11"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "ip_mtu", "MTU_9000"),
				resource.TestCheckResourceAttr(singularDatasourceName, "is_bfd_enabled", "true"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "oracle_bgp_asn"),
				resource.TestCheckResourceAttr(singularDatasourceName, "routing_policy.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "service_type"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttr(singularDatasourceName, "oracle_bgp_asn", "31898"),
				resource.TestCheckResourceAttr(singularDatasourceName, "service_type", "COLOCATED"),
				resource.TestCheckResourceAttr(singularDatasourceName, "state", "PROVISIONED"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttr(singularDatasourceName, "type", "PRIVATE"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "virtual_circuit_redundancy_metadata.#"),
			),
		},
		// verify resource import
		{
			Config:            config + CoreVirtualCircuitRequiredOnlyResourceNoMacSec,
			ImportState:       true,
			ImportStateVerify: true,
			ImportStateVerifyIgnore: []string{
				"region",
			},
			ResourceName: resourceName,
		},
	})
}

func testAccCheckCoreVirtualCircuitDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).VirtualNetworkClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_core_virtual_circuit" {
			noResourceFound = false
			request := oci_core.GetVirtualCircuitRequest{}

			tmp := rs.Primary.ID
			request.VirtualCircuitId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "core")

			response, err := client.GetVirtualCircuit(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_core.VirtualCircuitLifecycleStateTerminated): true,
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
	if !acctest.InSweeperExcludeList("CoreVirtualCircuit") {
		resource.AddTestSweepers("CoreVirtualCircuit", &resource.Sweeper{
			Name:         "CoreVirtualCircuit",
			Dependencies: acctest.DependencyGraph["virtualCircuit"],
			F:            sweepCoreVirtualCircuitResource,
		})
	}
}

func sweepCoreVirtualCircuitResource(compartment string) error {
	virtualNetworkClient := acctest.GetTestClients(&schema.ResourceData{}).VirtualNetworkClient()
	virtualCircuitIds, err := getCoreVirtualCircuitIds(compartment)
	if err != nil {
		return err
	}
	for _, virtualCircuitId := range virtualCircuitIds {
		if ok := acctest.SweeperDefaultResourceId[virtualCircuitId]; !ok {
			deleteVirtualCircuitRequest := oci_core.DeleteVirtualCircuitRequest{}

			deleteVirtualCircuitRequest.VirtualCircuitId = &virtualCircuitId

			deleteVirtualCircuitRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "core")
			_, error := virtualNetworkClient.DeleteVirtualCircuit(context.Background(), deleteVirtualCircuitRequest)
			if error != nil {
				fmt.Printf("Error deleting VirtualCircuit %s %s, It is possible that the resource is already deleted. Please verify manually \n", virtualCircuitId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &virtualCircuitId, CoreVirtualCircuitSweepWaitCondition, time.Duration(3*time.Minute),
				CoreVirtualCircuitSweepResponseFetchOperation, "core", true)
		}
	}
	return nil
}

func getCoreVirtualCircuitIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "VirtualCircuitId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	virtualNetworkClient := acctest.GetTestClients(&schema.ResourceData{}).VirtualNetworkClient()

	listVirtualCircuitsRequest := oci_core.ListVirtualCircuitsRequest{}
	listVirtualCircuitsRequest.CompartmentId = &compartmentId
	listVirtualCircuitsRequest.LifecycleState = oci_core.VirtualCircuitLifecycleStateProvisioned
	listVirtualCircuitsResponse, err := virtualNetworkClient.ListVirtualCircuits(context.Background(), listVirtualCircuitsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting VirtualCircuit list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, virtualCircuit := range listVirtualCircuitsResponse.Items {
		id := *virtualCircuit.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "VirtualCircuitId", id)
	}
	return resourceIds, nil
}

func CoreVirtualCircuitSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if virtualCircuitResponse, ok := response.Response.(oci_core.GetVirtualCircuitResponse); ok {
		return virtualCircuitResponse.LifecycleState != oci_core.VirtualCircuitLifecycleStateTerminated
	}
	return false
}

func CoreVirtualCircuitSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.VirtualNetworkClient().GetVirtualCircuit(context.Background(), oci_core.GetVirtualCircuitRequest{
		VirtualCircuitId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
