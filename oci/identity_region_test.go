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
	regionDataSourceRepresentation = map[string]interface{}{
		"filter": RepresentationGroup{Required, regionDataSourceFilterRepresentation}}

	regionDataSourceFilterRepresentation = map[string]interface{}{
		"name":   Representation{repType: Required, create: `name`},
		"values": Representation{repType: Required, create: []string{`${var.region}`}},
	}

	RegionResourceConfig = ""
)

// issue-routing-tag: identity/default
func TestIdentityRegionResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestIdentityRegionResource_basic")
	defer httpreplay.SaveScenario()

	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_identity_regions.test_regions"

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
					generateDataSourceFromRepresentationMap("oci_identity_regions", "test_regions", Required, Create, regionDataSourceRepresentation) +
					compartmentIdVariableStr + RegionResourceConfig,
				Check: ComposeAggregateTestCheckFuncWrapper(

					resource.TestCheckResourceAttrSet(datasourceName, "regions.#"),
					resource.TestCheckResourceAttrSet(datasourceName, "regions.0.key"),
					resource.TestCheckResourceAttrSet(datasourceName, "regions.0.name"),
				),
			},
		},
	})
}
