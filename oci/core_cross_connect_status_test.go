// Copyright (c) 2017, 2019, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

var (
	crossConnectStatusSingularDataSourceRepresentation = map[string]interface{}{
		"cross_connect_id": Representation{repType: Required, create: `${oci_core_cross_connect.test_cross_connect.id}`},
	}

	CrossConnectStatusResourceConfig = CrossConnectResourceConfig
)

func TestCoreCrossConnectStatusResource_basic(t *testing.T) {
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

					resource.TestCheckResourceAttr(singularDatasourceName, "interface_state", "UP"),
					resource.TestCheckResourceAttr(singularDatasourceName, "light_level_ind_bm", "-2.740000009536743"),
					resource.TestCheckResourceAttr(singularDatasourceName, "light_level_indicator", "GOOD"),
				),
			},
		},
	})
}
