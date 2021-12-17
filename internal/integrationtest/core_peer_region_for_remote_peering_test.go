// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
	"github.com/terraform-providers/terraform-provider-oci/internal/acctest"
	"github.com/terraform-providers/terraform-provider-oci/internal/utils"
)

var (
	peerRegionForRemotePeeringDataSourceRepresentation = map[string]interface{}{}

	PeerRegionForRemotePeeringResourceConfig = ""
)

// issue-routing-tag: core/default
func TestCorePeerRegionForRemotePeeringResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestCorePeerRegionForRemotePeeringResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_core_peer_region_for_remote_peerings.test_peer_region_for_remote_peerings"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_core_peer_region_for_remote_peerings", "test_peer_region_for_remote_peerings", acctest.Required, acctest.Create, peerRegionForRemotePeeringDataSourceRepresentation) +
				compartmentIdVariableStr + PeerRegionForRemotePeeringResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(

				resource.TestCheckResourceAttrSet(datasourceName, "peer_region_for_remote_peerings.#"),
				resource.TestCheckResourceAttrSet(datasourceName, "peer_region_for_remote_peerings.0.name"),
			),
		},
	})
}
