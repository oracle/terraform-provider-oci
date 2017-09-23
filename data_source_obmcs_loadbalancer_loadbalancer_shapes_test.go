// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package main

import (
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
)

func TestAccDatasourceLoadBalancerShapes_basic(t *testing.T) {
	providers := testAccProviders
	config := testProviderConfig() + `
	data "oci_load_balancer_shapes" "t" {
		compartment_id = "${var.compartment_id}"
	}`

	resourceName := "data.oci_load_balancer_shapes.t"

	resource.Test(t, resource.TestCase{
		PreventPostDestroyRefresh: true,
		Providers:                 providers,
		Steps: []resource.TestStep{
			{
				ImportState:       true,
				ImportStateVerify: true,
				Config:            config,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "shapes.#"),
					resource.TestCheckResourceAttrSet(resourceName, "shapes.0.name"),
				),
			},
		},
	})
}
