// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package main

import (
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"


)

func TestLoadBalancerProtocolsDatasource(t *testing.T) {
	client := GetTestProvider()
	providers := map[string]terraform.ResourceProvider{
		"baremetal": Provider(func(d *schema.ResourceData) (interface{}, error) {
			return client, nil
		}),
	}
	resourceName := "data.baremetal_load_balancer_protocols.t"
	config := `
data "baremetal_load_balancer_protocols" "t" {
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
					resource.TestCheckResourceAttrSet(resourceName, "protocols.#"),
					resource.TestCheckResourceAttrSet(resourceName, "protocols.0.name"),
				),
			},
		},
	})
}
