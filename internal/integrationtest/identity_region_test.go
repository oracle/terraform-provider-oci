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
	regionDataSourceRepresentation = map[string]interface{}{
		"filter": acctest.RepresentationGroup{RepType: acctest.Required, Group: regionDataSourceFilterRepresentation}}

	regionDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `name`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${var.region}`}},
	}

	RegionResourceConfig = ""
)

// issue-routing-tag: identity/default
func TestIdentityRegionResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestIdentityRegionResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_identity_regions.test_regions"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_identity_regions", "test_regions", acctest.Required, acctest.Create, regionDataSourceRepresentation) +
				compartmentIdVariableStr + RegionResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(

				resource.TestCheckResourceAttrSet(datasourceName, "regions.#"),
				resource.TestCheckResourceAttrSet(datasourceName, "regions.0.key"),
				resource.TestCheckResourceAttrSet(datasourceName, "regions.0.name"),
			),
		},
	})
}
