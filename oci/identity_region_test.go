// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	regionDataSourceRepresentation = map[string]interface{}{
		"filter": RepresentationGroup{Required, regionDataSourceFilterRepresentation}}

	regionDataSourceFilterRepresentation = map[string]interface{}{
		"name":   Representation{RepType: Required, Create: `name`},
		"values": Representation{RepType: Required, Create: []string{`${var.region}`}},
	}

	RegionResourceConfig = ""
)

// issue-routing-tag: identity/default
func TestIdentityRegionResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestIdentityRegionResource_basic")
	defer httpreplay.SaveScenario()

	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_identity_regions.test_regions"

	SaveConfigContent("", "", "", t)

	ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				GenerateDataSourceFromRepresentationMap("oci_identity_regions", "test_regions", Required, Create, regionDataSourceRepresentation) +
				compartmentIdVariableStr + RegionResourceConfig,
			Check: ComposeAggregateTestCheckFuncWrapper(

				resource.TestCheckResourceAttrSet(datasourceName, "regions.#"),
				resource.TestCheckResourceAttrSet(datasourceName, "regions.0.key"),
				resource.TestCheckResourceAttrSet(datasourceName, "regions.0.name"),
			),
		},
	})
}
