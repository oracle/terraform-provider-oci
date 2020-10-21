// Copyright (c) 2017, 2020, Oracle and/or its affiliates. All rights reserved.
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
	crossConnectStatusSingularDataSourceRepresentation = map[string]interface{}{
		"cross_connect_id": Representation{repType: Required, create: `${oci_core_cross_connect.test_cross_connect.id}`},
	}

	CrossConnectStatusResourceConfig = generateDataSourceFromRepresentationMap("oci_core_cross_connect_locations", "test_cross_connect_locations", Required, Create, crossConnectLocationDataSourceRepresentation) +
		generateResourceFromRepresentationMap("oci_core_cross_connect", "test_cross_connect", Required, Create, crossConnectRepresentation)
)

func TestCoreCrossConnectStatusResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestCoreCrossConnectStatusResource_basic")
	defer httpreplay.SaveScenario()

	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	singularDatasourceName := "data.oci_core_cross_connect_status.test_cross_connect_status"

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		Steps: []resource.TestStep{
			// verify singular datasource
			{
				Config: config +
					generateDataSourceFromRepresentationMap("oci_core_cross_connect_status", "test_cross_connect_status", Required, Create, crossConnectStatusSingularDataSourceRepresentation) +
					compartmentIdVariableStr + CrossConnectStatusResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(singularDatasourceName, "cross_connect_id"),

					resource.TestCheckResourceAttrSet(singularDatasourceName, "interface_state"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "light_level_ind_bm"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "light_level_indicator"),
				),
			},
		},
	})
}
