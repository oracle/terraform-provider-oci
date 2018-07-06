// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"fmt"
	"testing"

	"strings"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

const (
	CrossConnectStatusResourceConfig = CrossConnectStatusResourceDependencies + `

`
	CrossConnectStatusPropertyVariables = `

`
	CrossConnectStatusResourceDependencies = CrossConnectPropertyVariables + CrossConnectRequiredOnlyResource
)

func TestCoreCrossConnectStatusResource_basic(t *testing.T) {
	region := getRequiredEnvSetting("region")
	if !strings.EqualFold("r1", region) {
		t.Skip("FastConnect tests are not yet supported in production regions")
	}

	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getRequiredEnvSetting("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	singularDatasourceName := "data.oci_core_cross_connect_status.test_cross_connect_status"

	resource.Test(t, resource.TestCase{
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		Steps: []resource.TestStep{
			// verify singular datasource
			{
				Config: config + `

data "oci_core_cross_connect_status" "test_cross_connect_status" {
	#Required
	cross_connect_id = "${oci_core_cross_connect.test_cross_connect.id}"
}
                ` + compartmentIdVariableStr + CrossConnectStatusResourceConfig,
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
