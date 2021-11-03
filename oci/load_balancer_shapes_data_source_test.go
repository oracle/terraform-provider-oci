// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"testing"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

// issue-routing-tag: load_balancer/default
func TestAccDatasourceLoadBalancerShapes_basic(t *testing.T) {
	httpreplay.SetScenario("TestAccDatasourceLoadBalancerShapes_basic")
	defer httpreplay.SaveScenario()
	providers := TestAccProviders
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
		PreCheck:                  func() { PreCheck() },
		PreventPostDestroyRefresh: true,
		Providers:                 providers,
		Steps: []resource.TestStep{
			{
				Config: config,
				Check: ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(resourceName, "shapes.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "shapes.0.name", "100Mbps"),
				),
			},
		},
	})
}
