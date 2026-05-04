// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	NetworkFirewallNetworkFirewallHealthStatusSingularDataSourceRepresentation = map[string]interface{}{
		"network_firewall_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_network_firewall_network_firewall.test_network_firewall.id}`},
	}

	NetworkFirewallNetworkFirewallHealthStatusResourceConfig = acctest.GenerateResourceFromRepresentationMap("oci_core_subnet", "test_subnet", acctest.Required, acctest.Create, CoreSubnetRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", acctest.Required, acctest.Create, CoreVcnRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_network_firewall_network_firewall_policy", "test_network_firewall_policy", acctest.Required, acctest.Create, networkFirewallPolicyRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_network_firewall_network_firewall", "test_network_firewall", acctest.Required, acctest.Create, networkFirewallRepresentation)
)

// issue-routing-tag: network_firewall/default
func TestNetworkFirewallNetworkFirewallHealthStatusResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestNetworkFirewallNetworkFirewallHealthStatusResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	singularDatasourceName := "data.oci_network_firewall_network_firewall_health_status.test_network_firewall_health_status"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_network_firewall_network_firewall_health_status", "test_network_firewall_health_status", acctest.Required, acctest.Create, NetworkFirewallNetworkFirewallHealthStatusSingularDataSourceRepresentation) +
				compartmentIdVariableStr + NetworkFirewallNetworkFirewallHealthStatusResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "network_firewall_id"),

				resource.TestCheckResourceAttrSet(singularDatasourceName, "status"),
				resource.TestMatchResourceAttr(singularDatasourceName, "status", regexp.MustCompile("^(CRITICAL|WARNING|OK|UNKNOWN)$")),
			),
		},
	})
}
