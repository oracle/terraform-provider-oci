// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	peerRegionForRemotePeeringDataSourceRepresentation = map[string]interface{}{}

	PeerRegionForRemotePeeringResourceConfig = ""
)

// issue-routing-tag: core/default
func TestCorePeerRegionForRemotePeeringResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestCorePeerRegionForRemotePeeringResource_basic")
	defer httpreplay.SaveScenario()

	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_core_peer_region_for_remote_peerings.test_peer_region_for_remote_peerings"

	saveConfigContent("", "", "", t)

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		Steps: []resource.TestStep{
			// verify datasource
			{
				Config: config +
					generateDataSourceFromRepresentationMap("oci_core_peer_region_for_remote_peerings", "test_peer_region_for_remote_peerings", Required, Create, peerRegionForRemotePeeringDataSourceRepresentation) +
					compartmentIdVariableStr + PeerRegionForRemotePeeringResourceConfig,
				Check: ComposeAggregateTestCheckFuncWrapper(

					resource.TestCheckResourceAttrSet(datasourceName, "peer_region_for_remote_peerings.#"),
					resource.TestCheckResourceAttrSet(datasourceName, "peer_region_for_remote_peerings.0.name"),
				),
			},
		},
	})
}
