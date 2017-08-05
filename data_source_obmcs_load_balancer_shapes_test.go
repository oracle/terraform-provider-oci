// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package main

import (
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
)

func TestLoadBalancerShapesDatasource(t *testing.T) {
	client := GetTestProvider()
	providers := map[string]terraform.ResourceProvider{
		"oci": Provider(func(d *schema.ResourceData) (interface{}, error) {
			return client, nil
		}),
	}
	resourceName := "data.oci_load_balancer_shapes.t"
	config := `
data "oci_load_balancer_shapes" "t" {
  compartment_id = "${var.compartment_id}"
}
`
	config += testProviderConfig()

	resource.UnitTest(t, resource.TestCase{
		PreventPostDestroyRefresh: true,
		Providers:                 providers,
		Steps: []resource.TestStep{
			{
				Config: config,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "shapes.#"),
					resource.TestCheckResourceAttrSet(resourceName, "shapes.0.name"),
					resource.TestCheckResourceAttrSet(resourceName, "shapes.1.name"),
				),
			},
		},
	})
}
