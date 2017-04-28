// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package main

import (
	"testing"

	"github.com/MustWin/baremetal-sdk-go"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"

	"github.com/oracle/terraform-provider-baremetal/client/mocks"
)

func TestLoadBalancerShapesDatasource(t *testing.T) {
	client := &mocks.BareMetalClient{}
	providers := map[string]terraform.ResourceProvider{
		"baremetal": Provider(func(d *schema.ResourceData) (interface{}, error) {
			return client, nil
		}),
	}
	resourceName := "data.baremetal_load_balancer_shapes.t"
	config := `
data "baremetal_load_balancer_shapes" "t" {
  compartment_id = "ocid1.compartment.stub_id"
}
`
	config += testProviderConfig

	compartmentID := "ocid1.compartment.stub_id"
	list := &baremetal.ListLoadBalancerShapes{
		LoadBalancerShapes: []baremetal.LoadBalancerShape{
			baremetal.LoadBalancerShape{
				Name: "stub_name1",
			},
			baremetal.LoadBalancerShape{
				Name: "stub_name2",
			},
		},
	}
	client.On(
		"ListLoadBalancerShapes",
		compartmentID,
		(*baremetal.ListLoadBalancerPolicyOptions)(nil),
	).Return(list, nil)

	resource.UnitTest(t, resource.TestCase{
		PreventPostDestroyRefresh: true,
		Providers:                 providers,
		Steps: []resource.TestStep{
			{
				Config: config,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentID),
					resource.TestCheckResourceAttr(resourceName, "shapes.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "shapes.0.name", "stub_name1"),
					resource.TestCheckResourceAttr(resourceName, "shapes.1.name", "stub_name2"),
				),
			},
		},
	})
}
