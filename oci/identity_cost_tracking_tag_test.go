// Copyright (c) 2017, 2019, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	costTrackingTagDataSourceRepresentation = map[string]interface{}{
		"compartment_id": Representation{repType: Required, create: `${var.compartment_id}`},
	}

	CostTrackingTagResourceConfig = ""
)

func TestIdentityCostTrackingTagResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestIdentityCostTrackingTagResource_basic")
	defer httpreplay.SaveScenario()

	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_identity_cost_tracking_tags.test_cost_tracking_tags"

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		Steps: []resource.TestStep{
			// verify datasource
			{
				Config: config +
					generateDataSourceFromRepresentationMap("oci_identity_cost_tracking_tags", "test_cost_tracking_tags", Required, Create, costTrackingTagDataSourceRepresentation) +
					compartmentIdVariableStr + CostTrackingTagResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),

					resource.TestCheckResourceAttrSet(datasourceName, "tags.#"),
					resource.TestCheckResourceAttrSet(datasourceName, "tags.0.compartment_id"),
					resource.TestCheckResourceAttrSet(datasourceName, "tags.0.description"),
					resource.TestCheckResourceAttrSet(datasourceName, "tags.0.id"),
					resource.TestCheckResourceAttrSet(datasourceName, "tags.0.is_cost_tracking"),
					resource.TestCheckResourceAttrSet(datasourceName, "tags.0.is_retired"),
					resource.TestCheckResourceAttrSet(datasourceName, "tags.0.name"),
					resource.TestCheckResourceAttrSet(datasourceName, "tags.0.tag_namespace_id"),
					resource.TestCheckResourceAttrSet(datasourceName, "tags.0.tag_namespace_name"),
					resource.TestCheckResourceAttrSet(datasourceName, "tags.0.time_created"),
				),
			},
		},
	})
}
