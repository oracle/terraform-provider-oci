// Copyright (c) 2017, 2019, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"testing"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"

	"github.com/hashicorp/terraform/helper/resource"
)

func TestAccDatasourceLoadBalancerShapes_basic(t *testing.T) {
	httpreplay.SetScenario("TestAccDatasourceLoadBalancerShapes_basic")
	defer httpreplay.SaveScenario()
	providers := testAccProviders
	config := legacyTestProviderConfig() + `
	data "oci_load_balancer_shapes" "t" {
		compartment_id = "${var.compartment_id}"
		filter {
			name = "name"
			values = ["100Mbps"]
		}
	}`

	resourceName := "data.oci_load_balancer_shapes.t"

	resource.Test(t, resource.TestCase{
		PreCheck:                  func() { testAccPreCheck(t) },
		PreventPostDestroyRefresh: true,
		Providers:                 providers,
		Steps: []resource.TestStep{
			{
				Config: config,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "shapes.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "shapes.0.name", "100Mbps"),
				),
			},
		},
	})
}
