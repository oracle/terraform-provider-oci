// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	CoreIpSecConnectionTunnelRequiredOnlyResource = CoreIpSecConnectionTunnelResourceConfig +
		acctest.GenerateResourceFromRepresentationMap("oci_core_ipsec", "test_ip_sec_connection", acctest.Required, acctest.Create, CoreIpSecConnectionRepresentation)

	CoreIpSecConnectionTunnelSingularDataSourceRepresentation = map[string]interface{}{
		"ipsec_id":  acctest.Representation{RepType: acctest.Required, Create: `${oci_core_ipsec.test_ip_sec_connection.id}`},
		"tunnel_id": acctest.Representation{RepType: acctest.Required, Create: `${data.oci_core_ipsec_connection_tunnels.test_ip_sec_connection_tunnels.ip_sec_connection_tunnels.0.id}`},
	}

	CoreIpSecConnectionTunnelGroupDataSourceRepresentation = map[string]interface{}{
		"ipsec_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_core_ipsec.test_ip_sec_connection.id}`},
	}

	// First Tunnel
	//// 1. create with defaults - static
	//// 2. create with optional - static
	//// 3. update - bgp
	CoreIpSecConnectionTunnelRepresentation1 = map[string]interface{}{
		"ipsec_id":                acctest.Representation{RepType: acctest.Required, Create: `${oci_core_ipsec.test_ip_sec_connection.id}`},
		"tunnel_id":               acctest.Representation{RepType: acctest.Required, Create: `${data.oci_core_ipsec_connection_tunnels.test_ip_sec_connection_tunnels.ip_sec_connection_tunnels.0.id}`},
		"routing":                 acctest.Representation{RepType: acctest.Optional, Create: `STATIC`, Update: `BGP`},
		"ike_version":             acctest.Representation{RepType: acctest.Optional, Create: `V1`, Update: `V2`},
		"display_name":            acctest.Representation{RepType: acctest.Optional, Create: `MyIPSecConnectionTunnel1`, Update: `MyIPSecConnectionTunnel1 - updated`},
		"bgp_session_info":        acctest.RepresentationGroup{RepType: acctest.Optional, Group: CoreIpSecConnectionTunnelConfigurationBgpSessionInfoRepresentation1},
		"shared_secret":           acctest.Representation{RepType: acctest.Optional, Create: `sharedsecret1`, Update: `sharedsecret2`},
		"dpd_config":              acctest.RepresentationGroup{RepType: acctest.Optional, Group: ipSecConnectionTunnelConfigurationDpdConfigRepresentation1},
		"nat_translation_enabled": acctest.Representation{RepType: acctest.Optional, Create: `ENABLED`, Update: `DISABLED`},
		"oracle_can_initiate":     acctest.Representation{RepType: acctest.Optional, Create: "RESPONDER_ONLY", Update: "INITIATOR_OR_RESPONDER"},
		"phase_one_details":       acctest.RepresentationGroup{RepType: acctest.Optional, Group: ipSecConnectionTunnelConfigurationPhaseOneDetailsRepresentation1},
		"phase_two_details":       acctest.RepresentationGroup{RepType: acctest.Optional, Group: ipSecConnectionTunnelConfigurationPhaseTwoDetailsRepresentation1},
	}

	CoreIpSecConnectionTunnelConfigurationBgpSessionInfoRepresentation1 = map[string]interface{}{
		"customer_bgp_asn":      acctest.Representation{RepType: acctest.Optional, Update: `158`},
		"customer_interface_ip": acctest.Representation{RepType: acctest.Optional, Update: `10.0.0.16/31`},
		"oracle_interface_ip":   acctest.Representation{RepType: acctest.Optional, Update: `10.0.0.17/31`},
	}

	ipSecConnectionTunnelConfigurationDpdConfigRepresentation1 = map[string]interface{}{
		"dpd_mode":           acctest.Representation{RepType: acctest.Optional, Create: `INITIATE_AND_RESPOND`, Update: `RESPOND_ONLY`},
		"dpd_timeout_in_sec": acctest.Representation{RepType: acctest.Optional, Create: `50`, Update: `51`},
	}

	ipSecConnectionTunnelConfigurationPhaseOneDetailsRepresentation1 = map[string]interface{}{
		"custom_authentication_algorithm": acctest.Representation{RepType: acctest.Optional, Create: `SHA2_384`, Update: `SHA2_256`},
		"custom_encryption_algorithm":     acctest.Representation{RepType: acctest.Optional, Create: `AES_192_CBC`, Update: `AES_256_CBC`},
		"custom_dh_group":                 acctest.Representation{RepType: acctest.Optional, Create: `GROUP2`, Update: `GROUP19`},
		"is_custom_phase_one_config":      acctest.Representation{RepType: acctest.Optional, Create: `true`, Update: `true`},
		"lifetime":                        acctest.Representation{RepType: acctest.Optional, Create: `6`, Update: `5`},
	}

	ipSecConnectionTunnelConfigurationPhaseTwoDetailsRepresentation1 = map[string]interface{}{
		"custom_authentication_algorithm": acctest.Representation{RepType: acctest.Optional, Update: `HMAC_SHA2_256_128`},
		"custom_encryption_algorithm":     acctest.Representation{RepType: acctest.Optional, Create: `AES_128_GCM`, Update: `AES_192_CBC`},
		"dh_group":                        acctest.Representation{RepType: acctest.Optional, Update: `GROUP20`},
		"is_custom_phase_two_config":      acctest.Representation{RepType: acctest.Optional, Create: `true`, Update: `true`},  // if false then can't select custom encryption alg
		"is_pfs_enabled":                  acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`}, // if false then can't select dh group
		"lifetime":                        acctest.Representation{RepType: acctest.Optional, Create: `32`, Update: `31`},
	}

	// Second Tunnel
	//// 1. create optional - do static with bgp including ipv6
	//// 2. update - do policy based,  encryption_domain_config field

	CoreIpSecConnectionTunnelRepresentation2 = map[string]interface{}{
		"ipsec_id":                 acctest.Representation{RepType: acctest.Required, Create: `${oci_core_ipsec.test_ip_sec_connection.id}`},
		"tunnel_id":                acctest.Representation{RepType: acctest.Required, Create: `${data.oci_core_ipsec_connection_tunnels.test_ip_sec_connection_tunnels.ip_sec_connection_tunnels.1.id}`},
		"routing":                  acctest.Representation{RepType: acctest.Optional, Update: `POLICY`},
		"ike_version":              acctest.Representation{RepType: acctest.Optional, Create: `V2`, Update: `V1`},
		"display_name":             acctest.Representation{RepType: acctest.Optional, Create: `MyIPSecConnectionTunnel2`, Update: `MyIPSecConnectionTunnel2 - updated`},
		"bgp_session_info":         acctest.RepresentationGroup{RepType: acctest.Optional, Group: CoreIpSecConnectionTunnelConfigurationBgpSessionInfoRepresentation2},
		"shared_secret":            acctest.Representation{RepType: acctest.Optional, Create: `sharedsecret3`, Update: `sharedsecret4`},
		"dpd_config":               acctest.RepresentationGroup{RepType: acctest.Optional, Group: ipSecConnectionTunnelConfigurationDpdConfigRepresentation2},
		"nat_translation_enabled":  acctest.Representation{RepType: acctest.Optional, Create: `DISABLED`, Update: `ENABLED`},
		"oracle_can_initiate":      acctest.Representation{RepType: acctest.Optional, Create: "INITIATOR_OR_RESPONDER", Update: "RESPONDER_ONLY"},
		"phase_one_details":        acctest.RepresentationGroup{RepType: acctest.Optional, Group: ipSecConnectionTunnelConfigurationPhaseOneDetailsRepresentation2},
		"phase_two_details":        acctest.RepresentationGroup{RepType: acctest.Optional, Group: ipSecConnectionTunnelConfigurationPhaseTwoDetailsRepresentation2},
		"encryption_domain_config": acctest.RepresentationGroup{RepType: acctest.Optional, Group: ipSecConnectionTunnelConfigurationEncryptionDomainConfigRepresentation},
	}

	CoreIpSecConnectionTunnelConfigurationBgpSessionInfoRepresentation2 = map[string]interface{}{
		"customer_interface_ip":   acctest.Representation{RepType: acctest.Optional, Create: `10.0.0.16/31`, Update: `10.0.0.16/31`},
		"oracle_interface_ip":     acctest.Representation{RepType: acctest.Optional, Create: `10.0.0.17/31`, Update: `10.0.0.17/31`},
		"customer_interface_ipv6": acctest.Representation{RepType: acctest.Optional, Create: `2002:db2::6/64`},
		"oracle_interface_ipv6":   acctest.Representation{RepType: acctest.Optional, Create: `2002:db2::7/64`},
	}

	ipSecConnectionTunnelConfigurationDpdConfigRepresentation2 = map[string]interface{}{
		"dpd_mode":           acctest.Representation{RepType: acctest.Optional, Create: `RESPOND_ONLY`, Update: `INITIATE_AND_RESPOND`},
		"dpd_timeout_in_sec": acctest.Representation{RepType: acctest.Optional, Create: `35`, Update: `36`},
	}

	ipSecConnectionTunnelConfigurationPhaseOneDetailsRepresentation2 = map[string]interface{}{
		"is_custom_phase_one_config": acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `false`},
		"lifetime":                   acctest.Representation{RepType: acctest.Optional, Create: `28700`, Update: `28600`},
	}

	ipSecConnectionTunnelConfigurationPhaseTwoDetailsRepresentation2 = map[string]interface{}{
		"dh_group":                   acctest.Representation{RepType: acctest.Optional, Update: `GROUP20`},
		"is_custom_phase_two_config": acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `false`}, // if false then can't select custom encryption & authentication alg
		"is_pfs_enabled":             acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},  // if false then can't select dh group
		"lifetime":                   acctest.Representation{RepType: acctest.Optional, Create: `3601`, Update: `3602`},
	}

	ipSecConnectionTunnelConfigurationEncryptionDomainConfigRepresentation = map[string]interface{}{
		"cpe_traffic_selector":    acctest.Representation{RepType: acctest.Optional, Update: []string{`10.0.0.16/31`, `11.0.0.16/31`}},
		"oracle_traffic_selector": acctest.Representation{RepType: acctest.Optional, Update: []string{`12.0.0.16/31`}},
	}

	CoreIpSecConnectionTunnelResourceConfig = IpSecConnectionOptionalResource
)

// issue-routing-tag: core/default
func TestCoreIpSecConnectionTunnelResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestCoreIpSecConnectionTunnelResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid") // Test in an NgVpn tenancy
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_core_ipsec_connection_tunnel_management.test_ip_sec_connection_tunnel_management_tunnel"
	resourceName2 := "oci_core_ipsec_connection_tunnel_management.test_ip_sec_connection_tunnel_management_tunnel_2"
	datasourceName := "data.oci_core_ipsec_connection_tunnels.test_ip_sec_connection_tunnels"
	singularDatasourceName := "data.oci_core_ipsec_connection_tunnel.test_ip_sec_connection_tunnel"

	var resId, resId2 string

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify Create with required - Tunnel 1
		{
			Config: config + compartmentIdVariableStr + CoreIpSecConnectionTunnelResourceConfig +
				acctest.GenerateDataSourceFromRepresentationMap("oci_core_ipsec_connection_tunnels", "test_ip_sec_connection_tunnels", acctest.Required, acctest.Create, CoreIpSecConnectionTunnelGroupDataSourceRepresentation) +
				acctest.GenerateResourceFromRepresentationMap("oci_core_ipsec_connection_tunnel_management", "test_ip_sec_connection_tunnel_management_tunnel", acctest.Required, acctest.Create, CoreIpSecConnectionTunnelRepresentation1),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "cpe_ip", `203.0.113.6`),
				resource.TestCheckResourceAttrSet(resourceName, "shared_secret"),
				resource.TestCheckResourceAttrSet(resourceName, "display_name"),
				resource.TestCheckResourceAttrSet(resourceName, "ipsec_id"),
				resource.TestCheckResourceAttrSet(resourceName, "tunnel_id"),
				resource.TestCheckResourceAttrSet(resourceName, "vpn_ip"),
				resource.TestCheckResourceAttrSet(resourceName, "status"),
				resource.TestCheckResourceAttr(resourceName, "routing", "STATIC"), // Static is default
				resource.TestCheckResourceAttr(resourceName, "ike_version", "V1"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "time_status_updated"),

				resource.TestCheckNoResourceAttr(resourceName, "bgp_session_info"),
				resource.TestCheckNoResourceAttr(resourceName, "encryption_domain_config"),

				resource.TestCheckResourceAttr(resourceName, "nat_translation_enabled", "AUTO"),
				resource.TestCheckResourceAttr(resourceName, "oracle_can_initiate", "INITIATOR_OR_RESPONDER"),
				resource.TestCheckResourceAttr(resourceName, "dpd_mode", "INITIATE_AND_RESPOND"),
				resource.TestCheckResourceAttr(resourceName, "dpd_timeout_in_sec", "20"),

				resource.TestCheckResourceAttr(resourceName, "phase_one_details.0.custom_authentication_algorithm", ""), // Null is interpreted as empty string, which looks like known issue for getokexists
				resource.TestCheckResourceAttr(resourceName, "phase_one_details.0.custom_encryption_algorithm", ""),
				resource.TestCheckResourceAttr(resourceName, "phase_one_details.0.custom_dh_group", ""),
				resource.TestCheckResourceAttr(resourceName, "phase_one_details.0.is_custom_phase_one_config", "false"),
				resource.TestCheckResourceAttr(resourceName, "phase_one_details.0.lifetime", `28800`),
				resource.TestCheckResourceAttr(resourceName, "phase_one_details.0.remaining_lifetime", ""),
				resource.TestCheckResourceAttr(resourceName, "phase_one_details.0.negotiated_authentication_algorithm", ""),
				resource.TestCheckResourceAttr(resourceName, "phase_one_details.0.negotiated_encryption_algorithm", ""),
				resource.TestCheckResourceAttr(resourceName, "phase_one_details.0.negotiated_dh_group", ""),
				resource.TestCheckResourceAttr(resourceName, "phase_one_details.0.is_ike_established", "false"),
				resource.TestCheckResourceAttrSet(resourceName, "phase_one_details.0.remaining_lifetime_last_retrieved"),

				resource.TestCheckResourceAttr(resourceName, "phase_two_details.0.dh_group", "GROUP5"),
				resource.TestCheckResourceAttr(resourceName, "phase_two_details.0.custom_encryption_algorithm", ""),
				resource.TestCheckResourceAttr(resourceName, "phase_two_details.0.custom_authentication_algorithm", ""),
				resource.TestCheckResourceAttr(resourceName, "phase_two_details.0.lifetime", "3600"),
				resource.TestCheckResourceAttr(resourceName, "phase_two_details.0.is_custom_phase_two_config", "false"),
				resource.TestCheckResourceAttr(resourceName, "phase_two_details.0.negotiated_authentication_algorithm", ""),
				resource.TestCheckResourceAttr(resourceName, "phase_two_details.0.negotiated_encryption_algorithm", ""),
				resource.TestCheckResourceAttr(resourceName, "phase_two_details.0.negotiated_dh_group", ""),
				resource.TestCheckResourceAttr(resourceName, "phase_two_details.0.is_pfs_enabled", "true"),
				resource.TestCheckResourceAttr(resourceName, "phase_two_details.0.is_esp_established", "false"),
				resource.TestCheckResourceAttrSet(resourceName, "phase_two_details.0.remaining_lifetime_last_retrieved"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// verify Create with optionals - Tunnel 1 & 2
		{
			Config: config + compartmentIdVariableStr + CoreIpSecConnectionTunnelResourceConfig +
				acctest.GenerateDataSourceFromRepresentationMap("oci_core_ipsec_connection_tunnels", "test_ip_sec_connection_tunnels", acctest.Required, acctest.Create, CoreIpSecConnectionTunnelGroupDataSourceRepresentation) +
				acctest.GenerateResourceFromRepresentationMap("oci_core_ipsec_connection_tunnel_management", "test_ip_sec_connection_tunnel_management_tunnel", acctest.Optional, acctest.Create, CoreIpSecConnectionTunnelRepresentation1) +
				acctest.GenerateResourceFromRepresentationMap("oci_core_ipsec_connection_tunnel_management", "test_ip_sec_connection_tunnel_management_tunnel_2", acctest.Optional, acctest.Create, CoreIpSecConnectionTunnelRepresentation2),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				// check tunnel 1
				resource.TestCheckResourceAttrSet(resourceName, "ipsec_id"),
				resource.TestCheckResourceAttrSet(resourceName, "tunnel_id"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "cpe_ip", `203.0.113.6`),
				resource.TestCheckResourceAttr(resourceName, "ike_version", "V1"),
				resource.TestCheckResourceAttr(resourceName, "shared_secret", "sharedsecret1"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "MyIPSecConnectionTunnel1"),
				resource.TestCheckResourceAttrSet(resourceName, "vpn_ip"),
				resource.TestCheckResourceAttrSet(resourceName, "status"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttr(resourceName, "routing", "STATIC"),
				resource.TestCheckResourceAttr(resourceName, "nat_translation_enabled", "ENABLED"),
				resource.TestCheckResourceAttr(resourceName, "oracle_can_initiate", "RESPONDER_ONLY"),
				resource.TestCheckResourceAttr(resourceName, "dpd_mode", "INITIATE_AND_RESPOND"),
				resource.TestCheckResourceAttr(resourceName, "dpd_timeout_in_sec", "50"),

				resource.TestCheckNoResourceAttr(resourceName, "bgp_session_info"),
				resource.TestCheckNoResourceAttr(resourceName, "encryption_domain_config"),

				resource.TestCheckResourceAttr(resourceName, "phase_one_details.0.custom_authentication_algorithm", "SHA2_384"),
				resource.TestCheckResourceAttr(resourceName, "phase_one_details.0.custom_encryption_algorithm", "AES_192_CBC"),
				resource.TestCheckResourceAttr(resourceName, "phase_one_details.0.custom_dh_group", "GROUP2"),
				resource.TestCheckResourceAttr(resourceName, "phase_one_details.0.is_custom_phase_one_config", "true"),
				resource.TestCheckResourceAttr(resourceName, "phase_one_details.0.lifetime", `6`),
				resource.TestCheckResourceAttr(resourceName, "phase_one_details.0.remaining_lifetime", ""),
				resource.TestCheckResourceAttr(resourceName, "phase_one_details.0.negotiated_authentication_algorithm", ""),
				resource.TestCheckResourceAttr(resourceName, "phase_one_details.0.negotiated_encryption_algorithm", ""),
				resource.TestCheckResourceAttr(resourceName, "phase_one_details.0.negotiated_dh_group", ""),
				resource.TestCheckResourceAttr(resourceName, "phase_one_details.0.is_ike_established", "false"),
				resource.TestCheckResourceAttrSet(resourceName, "phase_one_details.0.remaining_lifetime_last_retrieved"),

				resource.TestCheckResourceAttr(resourceName, "phase_two_details.0.dh_group", ""), // no dh group when pfs is disabled
				resource.TestCheckResourceAttr(resourceName, "phase_two_details.0.custom_encryption_algorithm", "AES_128_GCM"),
				resource.TestCheckResourceAttr(resourceName, "phase_two_details.0.custom_authentication_algorithm", ""),
				resource.TestCheckResourceAttr(resourceName, "phase_two_details.0.lifetime", "32"),
				resource.TestCheckResourceAttr(resourceName, "phase_two_details.0.is_custom_phase_two_config", "true"),
				resource.TestCheckResourceAttr(resourceName, "phase_two_details.0.negotiated_authentication_algorithm", ""),
				resource.TestCheckResourceAttr(resourceName, "phase_two_details.0.negotiated_encryption_algorithm", ""),
				resource.TestCheckResourceAttr(resourceName, "phase_two_details.0.negotiated_dh_group", ""),
				resource.TestCheckResourceAttr(resourceName, "phase_two_details.0.is_pfs_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "phase_two_details.0.is_esp_established", "false"),
				resource.TestCheckResourceAttrSet(resourceName, "phase_two_details.0.remaining_lifetime_last_retrieved"),

				// check tunnel 2
				resource.TestCheckResourceAttrSet(resourceName2, "ipsec_id"),
				resource.TestCheckResourceAttrSet(resourceName2, "tunnel_id"),
				resource.TestCheckResourceAttr(resourceName2, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName2, "cpe_ip", `203.0.113.6`),
				resource.TestCheckResourceAttrSet(resourceName2, "vpn_ip"),
				resource.TestCheckResourceAttrSet(resourceName2, "status"),
				resource.TestCheckResourceAttrSet(resourceName2, "state"),
				resource.TestCheckResourceAttr(resourceName2, "shared_secret", "sharedsecret3"),
				resource.TestCheckResourceAttr(resourceName2, "display_name", "MyIPSecConnectionTunnel2"),
				resource.TestCheckResourceAttr(resourceName2, "routing", "STATIC"),
				resource.TestCheckResourceAttr(resourceName2, "ike_version", "V2"),
				resource.TestCheckResourceAttr(resourceName2, "nat_translation_enabled", "DISABLED"),
				resource.TestCheckResourceAttr(resourceName2, "oracle_can_initiate", "INITIATOR_OR_RESPONDER"),
				resource.TestCheckResourceAttr(resourceName2, "dpd_mode", "RESPOND_ONLY"),
				resource.TestCheckResourceAttr(resourceName2, "dpd_timeout_in_sec", "35"),

				resource.TestCheckResourceAttr(resourceName2, "bgp_session_info.0.customer_bgp_asn", ""),
				resource.TestCheckResourceAttr(resourceName2, "bgp_session_info.0.oracle_bgp_asn", ""),
				resource.TestCheckResourceAttr(resourceName2, "bgp_session_info.0.customer_interface_ip", "10.0.0.16/31"),
				resource.TestCheckResourceAttr(resourceName2, "bgp_session_info.0.oracle_interface_ip", "10.0.0.17/31"),
				resource.TestCheckResourceAttr(resourceName2, "bgp_session_info.0.customer_interface_ipv6", "2002:db2::6/64"),
				resource.TestCheckResourceAttr(resourceName2, "bgp_session_info.0.oracle_interface_ipv6", "2002:db2::7/64"),

				resource.TestCheckNoResourceAttr(resourceName2, "encryption_domain_config"),

				resource.TestCheckResourceAttr(resourceName2, "phase_one_details.0.custom_authentication_algorithm", ""),
				resource.TestCheckResourceAttr(resourceName2, "phase_one_details.0.custom_encryption_algorithm", ""),
				resource.TestCheckResourceAttr(resourceName2, "phase_one_details.0.custom_dh_group", ""),
				resource.TestCheckResourceAttr(resourceName2, "phase_one_details.0.is_custom_phase_one_config", "false"),
				resource.TestCheckResourceAttr(resourceName2, "phase_one_details.0.lifetime", `28700`),
				resource.TestCheckResourceAttr(resourceName2, "phase_one_details.0.remaining_lifetime", ""),
				resource.TestCheckResourceAttr(resourceName2, "phase_one_details.0.negotiated_authentication_algorithm", ""),
				resource.TestCheckResourceAttr(resourceName2, "phase_one_details.0.negotiated_encryption_algorithm", ""),
				resource.TestCheckResourceAttr(resourceName2, "phase_one_details.0.negotiated_dh_group", ""),
				resource.TestCheckResourceAttr(resourceName2, "phase_one_details.0.is_ike_established", "false"),
				resource.TestCheckResourceAttrSet(resourceName2, "phase_one_details.0.remaining_lifetime_last_retrieved"),

				resource.TestCheckResourceAttr(resourceName2, "phase_two_details.0.dh_group", ""),
				resource.TestCheckResourceAttr(resourceName2, "phase_two_details.0.custom_encryption_algorithm", ""),
				resource.TestCheckResourceAttr(resourceName2, "phase_two_details.0.custom_authentication_algorithm", ""),
				resource.TestCheckResourceAttr(resourceName2, "phase_two_details.0.lifetime", "3601"),
				resource.TestCheckResourceAttr(resourceName2, "phase_two_details.0.is_custom_phase_two_config", "false"),
				resource.TestCheckResourceAttr(resourceName2, "phase_two_details.0.negotiated_authentication_algorithm", ""),
				resource.TestCheckResourceAttr(resourceName2, "phase_two_details.0.negotiated_encryption_algorithm", ""),
				resource.TestCheckResourceAttr(resourceName2, "phase_two_details.0.negotiated_dh_group", ""),
				resource.TestCheckResourceAttr(resourceName2, "phase_two_details.0.is_pfs_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName2, "phase_two_details.0.is_esp_established", "false"),
				resource.TestCheckResourceAttrSet(resourceName2, "phase_two_details.0.remaining_lifetime_last_retrieved"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},
		// verify updates to updatable parameters - Tunnel 1 & 2
		{
			Config: config + compartmentIdVariableStr + CoreIpSecConnectionTunnelResourceConfig +
				acctest.GenerateDataSourceFromRepresentationMap("oci_core_ipsec_connection_tunnels", "test_ip_sec_connection_tunnels", acctest.Required, acctest.Create, CoreIpSecConnectionTunnelGroupDataSourceRepresentation) +
				acctest.GenerateResourceFromRepresentationMap("oci_core_ipsec_connection_tunnel_management", "test_ip_sec_connection_tunnel_management_tunnel", acctest.Optional, acctest.Update, CoreIpSecConnectionTunnelRepresentation1) +
				acctest.GenerateResourceFromRepresentationMap("oci_core_ipsec_connection_tunnel_management", "test_ip_sec_connection_tunnel_management_tunnel_2", acctest.Optional, acctest.Update, CoreIpSecConnectionTunnelRepresentation2),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				// Tunnel 1 checks
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName, "cpe_ip"),
				resource.TestCheckResourceAttr(resourceName, "ike_version", "V2"),
				resource.TestCheckResourceAttr(resourceName, "shared_secret", "sharedsecret2"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "MyIPSecConnectionTunnel1 - updated"),
				resource.TestCheckResourceAttrSet(resourceName, "vpn_ip"),
				resource.TestCheckResourceAttrSet(resourceName, "status"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttr(resourceName, "routing", "BGP"),
				resource.TestCheckResourceAttr(resourceName, "nat_translation_enabled", "DISABLED"),
				resource.TestCheckResourceAttr(resourceName, "oracle_can_initiate", "INITIATOR_OR_RESPONDER"),
				resource.TestCheckResourceAttr(resourceName, "dpd_mode", "RESPOND_ONLY"),
				resource.TestCheckResourceAttr(resourceName, "dpd_timeout_in_sec", "51"),

				resource.TestCheckResourceAttr(resourceName, "phase_one_details.0.custom_authentication_algorithm", "SHA2_256"),
				resource.TestCheckResourceAttr(resourceName, "phase_one_details.0.custom_encryption_algorithm", "AES_256_CBC"),
				resource.TestCheckResourceAttr(resourceName, "phase_one_details.0.custom_dh_group", "GROUP19"),
				resource.TestCheckResourceAttr(resourceName, "phase_one_details.0.is_custom_phase_one_config", "true"),
				resource.TestCheckResourceAttr(resourceName, "phase_one_details.0.lifetime", "5"),

				resource.TestCheckResourceAttr(resourceName, "phase_one_details.0.remaining_lifetime", ""),
				resource.TestCheckResourceAttr(resourceName, "phase_one_details.0.remaining_lifetime_int", "0"),
				resource.TestCheckResourceAttr(resourceName, "phase_one_details.0.negotiated_authentication_algorithm", ""),
				resource.TestCheckResourceAttr(resourceName, "phase_one_details.0.negotiated_encryption_algorithm", ""),
				resource.TestCheckResourceAttr(resourceName, "phase_one_details.0.negotiated_dh_group", ""),
				resource.TestCheckResourceAttr(resourceName, "phase_one_details.0.is_ike_established", "false"),
				resource.TestCheckResourceAttrSet(resourceName, "phase_one_details.0.remaining_lifetime_last_retrieved"),

				resource.TestCheckResourceAttr(resourceName, "phase_two_details.0.custom_encryption_algorithm", "AES_192_CBC"),
				resource.TestCheckResourceAttr(resourceName, "phase_two_details.0.dh_group", "GROUP20"),
				resource.TestCheckResourceAttr(resourceName, "phase_two_details.0.is_custom_phase_two_config", "true"),
				resource.TestCheckResourceAttr(resourceName, "phase_two_details.0.is_pfs_enabled", "true"),
				resource.TestCheckResourceAttr(resourceName, "phase_two_details.0.lifetime", "31"),

				resource.TestCheckResourceAttr(resourceName, "phase_two_details.0.negotiated_authentication_algorithm", ""),
				resource.TestCheckResourceAttr(resourceName, "phase_two_details.0.negotiated_encryption_algorithm", ""),
				resource.TestCheckResourceAttr(resourceName, "phase_two_details.0.negotiated_dh_group", ""),
				resource.TestCheckResourceAttr(resourceName, "phase_two_details.0.is_esp_established", "false"),
				resource.TestCheckResourceAttrSet(resourceName, "phase_two_details.0.remaining_lifetime_last_retrieved"),
				resource.TestCheckResourceAttr(resourceName, "phase_two_details.0.remaining_lifetime", ""),
				resource.TestCheckResourceAttr(resourceName, "phase_two_details.0.remaining_lifetime_int", "0"),

				resource.TestCheckResourceAttr(resourceName, "bgp_session_info.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "bgp_session_info.0.customer_bgp_asn", "158"),
				resource.TestCheckResourceAttrSet(resourceName, "bgp_session_info.0.oracle_bgp_asn"),
				resource.TestCheckResourceAttr(resourceName, "bgp_session_info.0.customer_interface_ip", "10.0.0.16/31"),
				resource.TestCheckResourceAttr(resourceName, "bgp_session_info.0.oracle_interface_ip", "10.0.0.17/31"),
				resource.TestCheckResourceAttr(resourceName, "dpd_config.0.dpd_mode", "RESPOND_ONLY"),
				resource.TestCheckResourceAttr(resourceName, "dpd_config.0.dpd_timeout_in_sec", "51"),
				resource.TestCheckResourceAttr(resourceName, "bgp_session_info.0.customer_interface_ipv6", ""),
				resource.TestCheckResourceAttr(resourceName, "bgp_session_info.0.oracle_interface_ipv6", ""),
				resource.TestCheckResourceAttr(resourceName, "bgp_session_info.0.bgp_state", "DOWN"),
				resource.TestCheckResourceAttr(resourceName, "bgp_session_info.0.bgp_ipv6state", "DOWN"),
				resource.TestCheckResourceAttr(resourceName, "bgp_session_info.0.bgp_ipv6_state", "DOWN"),

				resource.TestCheckNoResourceAttr(resourceName, "encryption_domain_config"),

				// Tunnel 2 checks
				resource.TestCheckResourceAttr(resourceName2, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName2, "cpe_ip"),
				resource.TestCheckResourceAttrSet(resourceName2, "vpn_ip"),
				resource.TestCheckResourceAttrSet(resourceName2, "status"),
				resource.TestCheckResourceAttrSet(resourceName2, "state"),
				resource.TestCheckResourceAttr(resourceName2, "shared_secret", "sharedsecret4"),
				resource.TestCheckResourceAttr(resourceName2, "display_name", "MyIPSecConnectionTunnel2 - updated"),
				resource.TestCheckResourceAttr(resourceName2, "routing", "POLICY"),
				resource.TestCheckResourceAttr(resourceName2, "ike_version", "V1"),
				resource.TestCheckResourceAttr(resourceName2, "nat_translation_enabled", "ENABLED"),
				resource.TestCheckResourceAttr(resourceName2, "oracle_can_initiate", "RESPONDER_ONLY"),
				resource.TestCheckResourceAttr(resourceName2, "dpd_mode", "INITIATE_AND_RESPOND"),
				resource.TestCheckResourceAttr(resourceName2, "dpd_timeout_in_sec", "36"),

				resource.TestCheckResourceAttr(resourceName2, "phase_one_details.0.custom_authentication_algorithm", ""),
				resource.TestCheckResourceAttr(resourceName2, "phase_one_details.0.custom_encryption_algorithm", ""),
				resource.TestCheckResourceAttr(resourceName2, "phase_one_details.0.custom_dh_group", ""),
				resource.TestCheckResourceAttr(resourceName2, "phase_one_details.0.is_custom_phase_one_config", "false"),
				resource.TestCheckResourceAttr(resourceName2, "phase_one_details.0.lifetime", "28600"),

				resource.TestCheckResourceAttr(resourceName2, "phase_one_details.0.remaining_lifetime", ""),
				resource.TestCheckResourceAttr(resourceName2, "phase_one_details.0.negotiated_authentication_algorithm", ""),
				resource.TestCheckResourceAttr(resourceName2, "phase_one_details.0.negotiated_encryption_algorithm", ""),
				resource.TestCheckResourceAttr(resourceName2, "phase_one_details.0.negotiated_dh_group", ""),
				resource.TestCheckResourceAttr(resourceName2, "phase_one_details.0.is_ike_established", "false"),
				resource.TestCheckResourceAttrSet(resourceName2, "phase_one_details.0.remaining_lifetime_last_retrieved"),

				resource.TestCheckResourceAttr(resourceName2, "phase_two_details.0.custom_encryption_algorithm", ""),
				resource.TestCheckResourceAttr(resourceName2, "phase_two_details.0.dh_group", "GROUP20"),
				resource.TestCheckResourceAttr(resourceName2, "phase_two_details.0.is_custom_phase_two_config", "false"),
				resource.TestCheckResourceAttr(resourceName2, "phase_two_details.0.is_pfs_enabled", "true"),
				resource.TestCheckResourceAttr(resourceName2, "phase_two_details.0.lifetime", "3602"),

				resource.TestCheckResourceAttr(resourceName2, "phase_two_details.0.negotiated_authentication_algorithm", ""),
				resource.TestCheckResourceAttr(resourceName2, "phase_two_details.0.negotiated_encryption_algorithm", ""),
				resource.TestCheckResourceAttr(resourceName2, "phase_two_details.0.negotiated_dh_group", ""),
				resource.TestCheckResourceAttr(resourceName2, "phase_two_details.0.is_esp_established", "false"),
				resource.TestCheckResourceAttrSet(resourceName2, "phase_two_details.0.remaining_lifetime_last_retrieved"),

				resource.TestCheckResourceAttr(resourceName2, "bgp_session_info.0.customer_bgp_asn", ""),
				resource.TestCheckResourceAttrSet(resourceName2, "bgp_session_info.0.oracle_bgp_asn"), // Policy tunnel sets oracle asn
				resource.TestCheckResourceAttr(resourceName2, "bgp_session_info.0.customer_interface_ip", "10.0.0.16/31"),
				resource.TestCheckResourceAttr(resourceName2, "bgp_session_info.0.oracle_interface_ip", "10.0.0.17/31"),
				resource.TestCheckResourceAttr(resourceName2, "bgp_session_info.0.customer_interface_ipv6", "2002:db2::6/64"),
				resource.TestCheckResourceAttr(resourceName2, "bgp_session_info.0.oracle_interface_ipv6", "2002:db2::7/64"),
				resource.TestCheckResourceAttr(resourceName2, "bgp_session_info.0.bgp_state", ""),
				resource.TestCheckResourceAttr(resourceName2, "bgp_session_info.0.bgp_ipv6state", ""),
				resource.TestCheckResourceAttr(resourceName2, "bgp_session_info.0.bgp_ipv6_state", ""),

				resource.TestCheckResourceAttr(resourceName2, "encryption_domain_config.0.oracle_traffic_selector.#", "1"),
				resource.TestCheckResourceAttr(resourceName2, "encryption_domain_config.0.oracle_traffic_selector.0", "12.0.0.16/31"),
				resource.TestCheckResourceAttr(resourceName2, "encryption_domain_config.0.cpe_traffic_selector.#", "2"),
				resource.TestCheckResourceAttr(resourceName2, "encryption_domain_config.0.cpe_traffic_selector.0", "10.0.0.16/31"),
				resource.TestCheckResourceAttr(resourceName2, "encryption_domain_config.0.cpe_traffic_selector.1", "11.0.0.16/31"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("Resource recreated when it was supposed to be updated.")
					}
					return err
				},
			),
		},
		// verify datasource - both tunnels
		{
			Config: config + compartmentIdVariableStr + CoreIpSecConnectionTunnelResourceConfig +
				acctest.GenerateDataSourceFromRepresentationMap("oci_core_ipsec_connection_tunnels", "test_ip_sec_connection_tunnels", acctest.Required, acctest.Update, CoreIpSecConnectionTunnelGroupDataSourceRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "ipsec_id"),

				resource.TestCheckResourceAttr(datasourceName, "ip_sec_connection_tunnels.#", "2"),

				// Tunnel 1
				resource.TestCheckResourceAttrSet(datasourceName, "ip_sec_connection_tunnels.0.bgp_session_info.0.bgp_ipv6_state"),
				resource.TestCheckResourceAttrSet(datasourceName, "ip_sec_connection_tunnels.0.bgp_session_info.0.bgp_ipv6state"),
				resource.TestCheckResourceAttrSet(datasourceName, "ip_sec_connection_tunnels.0.bgp_session_info.0.bgp_state"),
				resource.TestCheckResourceAttr(datasourceName, "ip_sec_connection_tunnels.0.associated_virtual_circuits.#", "0"),
				resource.TestCheckResourceAttr(datasourceName, "ip_sec_connection_tunnels.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(datasourceName, "ip_sec_connection_tunnels.0.cpe_ip"),
				resource.TestCheckResourceAttr(datasourceName, "ip_sec_connection_tunnels.0.display_name", "MyIPSecConnectionTunnel1 - updated"),
				resource.TestCheckResourceAttr(datasourceName, "ip_sec_connection_tunnels.0.dpd_mode", "RESPOND_ONLY"),
				resource.TestCheckResourceAttr(datasourceName, "ip_sec_connection_tunnels.0.dpd_timeout_in_sec", "51"),
				resource.TestCheckResourceAttrSet(datasourceName, "ip_sec_connection_tunnels.0.id"),
				resource.TestCheckResourceAttr(datasourceName, "ip_sec_connection_tunnels.0.ike_version", "V2"),
				resource.TestCheckResourceAttr(datasourceName, "ip_sec_connection_tunnels.0.nat_translation_enabled", "DISABLED"),
				resource.TestCheckResourceAttr(datasourceName, "ip_sec_connection_tunnels.0.oracle_can_initiate", "INITIATOR_OR_RESPONDER"),
				resource.TestCheckResourceAttr(datasourceName, "ip_sec_connection_tunnels.0.phase_one_details.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "ip_sec_connection_tunnels.0.phase_two_details.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "ip_sec_connection_tunnels.0.routing", "BGP"),
				resource.TestCheckResourceAttrSet(datasourceName, "ip_sec_connection_tunnels.0.state"),
				resource.TestCheckResourceAttrSet(datasourceName, "ip_sec_connection_tunnels.0.status"),
				resource.TestCheckResourceAttrSet(datasourceName, "ip_sec_connection_tunnels.0.time_created"),
				resource.TestCheckResourceAttrSet(datasourceName, "ip_sec_connection_tunnels.0.time_status_updated"),
				resource.TestCheckResourceAttrSet(datasourceName, "ip_sec_connection_tunnels.0.vpn_ip"),

				resource.TestCheckResourceAttr(datasourceName, "ip_sec_connection_tunnels.0.phase_one_details.0.custom_authentication_algorithm", "SHA2_256"),
				resource.TestCheckResourceAttr(datasourceName, "ip_sec_connection_tunnels.0.phase_one_details.0.custom_encryption_algorithm", "AES_256_CBC"),
				resource.TestCheckResourceAttr(datasourceName, "ip_sec_connection_tunnels.0.phase_one_details.0.custom_dh_group", "GROUP19"),
				resource.TestCheckResourceAttr(datasourceName, "ip_sec_connection_tunnels.0.phase_one_details.0.is_custom_phase_one_config", "true"),
				resource.TestCheckResourceAttr(datasourceName, "ip_sec_connection_tunnels.0.phase_one_details.0.lifetime", "5"),
				resource.TestCheckResourceAttr(datasourceName, "ip_sec_connection_tunnels.0.phase_one_details.0.remaining_lifetime", ""),
				resource.TestCheckResourceAttr(datasourceName, "ip_sec_connection_tunnels.0.phase_one_details.0.negotiated_authentication_algorithm", ""),
				resource.TestCheckResourceAttr(datasourceName, "ip_sec_connection_tunnels.0.phase_one_details.0.negotiated_encryption_algorithm", ""),
				resource.TestCheckResourceAttr(datasourceName, "ip_sec_connection_tunnels.0.phase_one_details.0.negotiated_dh_group", ""),
				resource.TestCheckResourceAttr(datasourceName, "ip_sec_connection_tunnels.0.phase_one_details.0.is_ike_established", "false"),
				resource.TestCheckResourceAttrSet(datasourceName, "ip_sec_connection_tunnels.0.phase_one_details.0.remaining_lifetime_last_retrieved"),

				resource.TestCheckResourceAttr(datasourceName, "ip_sec_connection_tunnels.0.phase_two_details.0.custom_encryption_algorithm", "AES_192_CBC"),
				resource.TestCheckResourceAttr(datasourceName, "ip_sec_connection_tunnels.0.phase_two_details.0.dh_group", "GROUP20"),
				resource.TestCheckResourceAttr(datasourceName, "ip_sec_connection_tunnels.0.phase_two_details.0.is_custom_phase_two_config", "true"),
				resource.TestCheckResourceAttr(datasourceName, "ip_sec_connection_tunnels.0.phase_two_details.0.is_pfs_enabled", "true"),
				resource.TestCheckResourceAttr(datasourceName, "ip_sec_connection_tunnels.0.phase_two_details.0.lifetime", "31"),
				resource.TestCheckResourceAttr(datasourceName, "ip_sec_connection_tunnels.0.phase_two_details.0.negotiated_authentication_algorithm", ""),
				resource.TestCheckResourceAttr(datasourceName, "ip_sec_connection_tunnels.0.phase_two_details.0.negotiated_encryption_algorithm", ""),
				resource.TestCheckResourceAttr(datasourceName, "ip_sec_connection_tunnels.0.phase_two_details.0.negotiated_dh_group", ""),
				resource.TestCheckResourceAttr(datasourceName, "ip_sec_connection_tunnels.0.phase_two_details.0.is_esp_established", "false"),
				resource.TestCheckResourceAttrSet(datasourceName, "ip_sec_connection_tunnels.0.phase_two_details.0.remaining_lifetime_last_retrieved"),

				resource.TestCheckResourceAttr(datasourceName, "ip_sec_connection_tunnels.0.bgp_session_info.0.customer_bgp_asn", "158"),
				resource.TestCheckResourceAttrSet(datasourceName, "ip_sec_connection_tunnels.0.bgp_session_info.0.oracle_bgp_asn"),
				resource.TestCheckResourceAttr(datasourceName, "ip_sec_connection_tunnels.0.bgp_session_info.0.customer_interface_ip", "10.0.0.16/31"),
				resource.TestCheckResourceAttr(datasourceName, "ip_sec_connection_tunnels.0.bgp_session_info.0.oracle_interface_ip", "10.0.0.17/31"),
				resource.TestCheckResourceAttr(datasourceName, "ip_sec_connection_tunnels.0.bgp_session_info.0.customer_interface_ipv6", ""),
				resource.TestCheckResourceAttr(datasourceName, "ip_sec_connection_tunnels.0.bgp_session_info.0.oracle_interface_ipv6", ""),

				resource.TestCheckNoResourceAttr(datasourceName, "ip_sec_connection_tunnels.0.encryption_domain_config"),

				// Tunnel 2
				resource.TestCheckResourceAttrSet(datasourceName, "ip_sec_connection_tunnels.1.id"),
				resource.TestCheckResourceAttrSet(datasourceName, "ip_sec_connection_tunnels.1.time_created"),
				resource.TestCheckResourceAttrSet(datasourceName, "ip_sec_connection_tunnels.1.time_status_updated"),
				resource.TestCheckResourceAttr(datasourceName, "ip_sec_connection_tunnels.1.compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(datasourceName, "ip_sec_connection_tunnels.1.cpe_ip"),
				resource.TestCheckResourceAttrSet(datasourceName, "ip_sec_connection_tunnels.1.vpn_ip"),
				resource.TestCheckResourceAttrSet(datasourceName, "ip_sec_connection_tunnels.1.status"),
				resource.TestCheckResourceAttrSet(datasourceName, "ip_sec_connection_tunnels.1.state"),
				resource.TestCheckResourceAttr(datasourceName, "ip_sec_connection_tunnels.1.display_name", "MyIPSecConnectionTunnel2 - updated"),
				resource.TestCheckResourceAttr(datasourceName, "ip_sec_connection_tunnels.1.routing", "POLICY"),
				resource.TestCheckResourceAttr(datasourceName, "ip_sec_connection_tunnels.1.ike_version", "V1"),
				resource.TestCheckResourceAttr(datasourceName, "ip_sec_connection_tunnels.1.nat_translation_enabled", "ENABLED"),
				resource.TestCheckResourceAttr(datasourceName, "ip_sec_connection_tunnels.1.oracle_can_initiate", "RESPONDER_ONLY"),
				resource.TestCheckResourceAttr(datasourceName, "ip_sec_connection_tunnels.1.dpd_mode", "INITIATE_AND_RESPOND"),
				resource.TestCheckResourceAttr(datasourceName, "ip_sec_connection_tunnels.1.dpd_timeout_in_sec", "36"),

				resource.TestCheckResourceAttr(datasourceName, "ip_sec_connection_tunnels.1.phase_one_details.0.custom_authentication_algorithm", ""),
				resource.TestCheckResourceAttr(datasourceName, "ip_sec_connection_tunnels.1.phase_one_details.0.custom_encryption_algorithm", ""),
				resource.TestCheckResourceAttr(datasourceName, "ip_sec_connection_tunnels.1.phase_one_details.0.custom_dh_group", ""),
				resource.TestCheckResourceAttr(datasourceName, "ip_sec_connection_tunnels.1.phase_one_details.0.is_custom_phase_one_config", "false"),
				resource.TestCheckResourceAttr(datasourceName, "ip_sec_connection_tunnels.1.phase_one_details.0.lifetime", "28600"),
				resource.TestCheckResourceAttr(datasourceName, "ip_sec_connection_tunnels.1.phase_one_details.0.remaining_lifetime", ""),
				resource.TestCheckResourceAttr(datasourceName, "ip_sec_connection_tunnels.1.phase_one_details.0.negotiated_authentication_algorithm", ""),
				resource.TestCheckResourceAttr(datasourceName, "ip_sec_connection_tunnels.1.phase_one_details.0.negotiated_encryption_algorithm", ""),
				resource.TestCheckResourceAttr(datasourceName, "ip_sec_connection_tunnels.1.phase_one_details.0.negotiated_dh_group", ""),
				resource.TestCheckResourceAttr(datasourceName, "ip_sec_connection_tunnels.1.phase_one_details.0.is_ike_established", "false"),
				resource.TestCheckResourceAttrSet(datasourceName, "ip_sec_connection_tunnels.1.phase_one_details.0.remaining_lifetime_last_retrieved"),

				resource.TestCheckResourceAttr(datasourceName, "ip_sec_connection_tunnels.1.phase_two_details.0.custom_encryption_algorithm", ""),
				resource.TestCheckResourceAttr(datasourceName, "ip_sec_connection_tunnels.1.phase_two_details.0.dh_group", "GROUP20"),
				resource.TestCheckResourceAttr(datasourceName, "ip_sec_connection_tunnels.1.phase_two_details.0.is_custom_phase_two_config", "false"),
				resource.TestCheckResourceAttr(datasourceName, "ip_sec_connection_tunnels.1.phase_two_details.0.is_pfs_enabled", "true"),
				resource.TestCheckResourceAttr(datasourceName, "ip_sec_connection_tunnels.1.phase_two_details.0.lifetime", "3602"),
				resource.TestCheckResourceAttr(datasourceName, "ip_sec_connection_tunnels.1.phase_two_details.0.negotiated_authentication_algorithm", ""),
				resource.TestCheckResourceAttr(datasourceName, "ip_sec_connection_tunnels.1.phase_two_details.0.negotiated_encryption_algorithm", ""),
				resource.TestCheckResourceAttr(datasourceName, "ip_sec_connection_tunnels.1.phase_two_details.0.negotiated_dh_group", ""),
				resource.TestCheckResourceAttr(datasourceName, "ip_sec_connection_tunnels.1.phase_two_details.0.is_esp_established", "false"),
				resource.TestCheckResourceAttrSet(datasourceName, "ip_sec_connection_tunnels.1.phase_two_details.0.remaining_lifetime_last_retrieved"),

				resource.TestCheckResourceAttr(datasourceName, "ip_sec_connection_tunnels.1.bgp_session_info.0.customer_bgp_asn", ""),
				resource.TestCheckResourceAttrSet(datasourceName, "ip_sec_connection_tunnels.1.bgp_session_info.0.oracle_bgp_asn"), // policy tunnel sets oracle asn
				resource.TestCheckResourceAttr(datasourceName, "ip_sec_connection_tunnels.1.bgp_session_info.0.bgp_state", ""),     // Policy tunnel does not have bgp state set
				resource.TestCheckResourceAttr(datasourceName, "ip_sec_connection_tunnels.1.bgp_session_info.0.bgp_ipv6state", ""),
				resource.TestCheckResourceAttr(datasourceName, "ip_sec_connection_tunnels.1.bgp_session_info.0.bgp_ipv6_state", ""),
				resource.TestCheckResourceAttr(datasourceName, "ip_sec_connection_tunnels.1.bgp_session_info.0.customer_interface_ip", "10.0.0.16/31"),
				resource.TestCheckResourceAttr(datasourceName, "ip_sec_connection_tunnels.1.bgp_session_info.0.oracle_interface_ip", "10.0.0.17/31"),
				resource.TestCheckResourceAttr(datasourceName, "ip_sec_connection_tunnels.1.bgp_session_info.0.customer_interface_ipv6", "2002:db2::6/64"),
				resource.TestCheckResourceAttr(datasourceName, "ip_sec_connection_tunnels.1.bgp_session_info.0.oracle_interface_ipv6", "2002:db2::7/64"),

				resource.TestCheckResourceAttr(datasourceName, "ip_sec_connection_tunnels.1.encryption_domain_config.0.oracle_traffic_selector.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "ip_sec_connection_tunnels.1.encryption_domain_config.0.oracle_traffic_selector.0", "12.0.0.16/31"),
				resource.TestCheckResourceAttr(datasourceName, "ip_sec_connection_tunnels.1.encryption_domain_config.0.cpe_traffic_selector.#", "2"),
				resource.TestCheckResourceAttr(datasourceName, "ip_sec_connection_tunnels.1.encryption_domain_config.0.cpe_traffic_selector.0", "10.0.0.16/31"),
				resource.TestCheckResourceAttr(datasourceName, "ip_sec_connection_tunnels.1.encryption_domain_config.0.cpe_traffic_selector.1", "11.0.0.16/31"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_core_ipsec_connection_tunnel", "test_ip_sec_connection_tunnel", acctest.Required, acctest.Create, CoreIpSecConnectionTunnelSingularDataSourceRepresentation) +
				acctest.GenerateDataSourceFromRepresentationMap("oci_core_ipsec_connection_tunnels", "test_ip_sec_connection_tunnels", acctest.Required, acctest.Create, CoreIpSecConnectionTunnelGroupDataSourceRepresentation) +
				compartmentIdVariableStr + CoreIpSecConnectionTunnelResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "ipsec_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "tunnel_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "cpe_ip"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "MyIPSecConnectionTunnel1 - updated"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "ike_version", "V2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "routing", "BGP"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "status"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_status_updated"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "vpn_ip"),

				resource.TestCheckResourceAttr(singularDatasourceName, "nat_translation_enabled", "DISABLED"),
				resource.TestCheckResourceAttr(singularDatasourceName, "oracle_can_initiate", "INITIATOR_OR_RESPONDER"),
				resource.TestCheckResourceAttr(singularDatasourceName, "dpd_mode", "RESPOND_ONLY"),
				resource.TestCheckResourceAttr(singularDatasourceName, "dpd_timeout_in_sec", "51"),

				resource.TestCheckResourceAttr(singularDatasourceName, "phase_one_details.0.custom_authentication_algorithm", "SHA2_256"),
				resource.TestCheckResourceAttr(singularDatasourceName, "phase_one_details.0.custom_encryption_algorithm", "AES_256_CBC"),
				resource.TestCheckResourceAttr(singularDatasourceName, "phase_one_details.0.custom_dh_group", "GROUP19"),
				resource.TestCheckResourceAttr(singularDatasourceName, "phase_one_details.0.is_custom_phase_one_config", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "phase_one_details.0.lifetime", "5"),
				resource.TestCheckResourceAttr(singularDatasourceName, "phase_one_details.0.remaining_lifetime", ""),
				resource.TestCheckResourceAttr(singularDatasourceName, "phase_one_details.0.remaining_lifetime_int", "0"),
				resource.TestCheckResourceAttr(singularDatasourceName, "phase_one_details.0.negotiated_authentication_algorithm", ""),
				resource.TestCheckResourceAttr(singularDatasourceName, "phase_one_details.0.negotiated_encryption_algorithm", ""),
				resource.TestCheckResourceAttr(singularDatasourceName, "phase_one_details.0.negotiated_dh_group", ""),
				resource.TestCheckResourceAttr(singularDatasourceName, "phase_one_details.0.is_ike_established", "false"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "phase_one_details.0.remaining_lifetime_last_retrieved"),

				resource.TestCheckResourceAttr(singularDatasourceName, "phase_two_details.0.custom_encryption_algorithm", "AES_192_CBC"),
				resource.TestCheckResourceAttr(singularDatasourceName, "phase_two_details.0.dh_group", "GROUP20"),
				resource.TestCheckResourceAttr(singularDatasourceName, "phase_two_details.0.is_custom_phase_two_config", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "phase_two_details.0.is_pfs_enabled", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "phase_two_details.0.lifetime", "31"),
				resource.TestCheckResourceAttr(singularDatasourceName, "phase_two_details.0.remaining_lifetime", ""),
				resource.TestCheckResourceAttr(singularDatasourceName, "phase_two_details.0.remaining_lifetime_int", "0"),
				resource.TestCheckResourceAttr(singularDatasourceName, "phase_two_details.0.negotiated_authentication_algorithm", ""),
				resource.TestCheckResourceAttr(singularDatasourceName, "phase_two_details.0.negotiated_encryption_algorithm", ""),
				resource.TestCheckResourceAttr(singularDatasourceName, "phase_two_details.0.negotiated_dh_group", ""),
				resource.TestCheckResourceAttr(singularDatasourceName, "phase_two_details.0.is_esp_established", "false"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "phase_two_details.0.remaining_lifetime_last_retrieved"),

				resource.TestCheckResourceAttr(singularDatasourceName, "bgp_session_info.0.customer_bgp_asn", "158"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "bgp_session_info.0.oracle_bgp_asn"),

				resource.TestCheckResourceAttr(singularDatasourceName, "bgp_session_info.0.customer_interface_ip", "10.0.0.16/31"),
				resource.TestCheckResourceAttr(singularDatasourceName, "bgp_session_info.0.oracle_interface_ip", "10.0.0.17/31"),
				resource.TestCheckResourceAttr(singularDatasourceName, "bgp_session_info.0.customer_interface_ipv6", ""),
				resource.TestCheckResourceAttr(singularDatasourceName, "bgp_session_info.0.oracle_interface_ipv6", ""),

				resource.TestCheckNoResourceAttr(singularDatasourceName, "encryption_domain_config"),
			),
		},
	})
}
