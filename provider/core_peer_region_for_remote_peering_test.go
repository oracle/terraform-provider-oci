// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

const (
	PeerRegionForRemotePeeringResourceConfig = PeerRegionForRemotePeeringResourceDependencies + `

`
	PeerRegionForRemotePeeringPropertyVariables = `

`
	PeerRegionForRemotePeeringResourceDependencies = ""
)

func TestCorePeerRegionForRemotePeeringResource_basic(t *testing.T) {
	provider := testAccProvider
	config := testProviderConfig()

	compartmentId2 := getRequiredEnvSetting("compartment_id_for_update")
	compartmentIdVariableStr2 := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId2)

	datasourceName := "data.oci_core_peer_region_for_remote_peerings.test_peer_region_for_remote_peerings"

	resource.Test(t, resource.TestCase{
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		Steps: []resource.TestStep{
			// verify datasource
			{
				Config: config + `

data "oci_core_peer_region_for_remote_peerings" "test_peer_region_for_remote_peerings" {
}
                ` + compartmentIdVariableStr2 + PeerRegionForRemotePeeringResourceConfig,
				Check: resource.ComposeTestCheckFunc(

					resource.TestCheckResourceAttrSet(datasourceName, "peer_region_for_remote_peerings.#"),
					resource.TestCheckResourceAttrSet(datasourceName, "peer_region_for_remote_peerings.0.name"),
				),
			},
		},
	})
}
