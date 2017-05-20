// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package main

import (
	"testing"

	"github.com/MustWin/baremetal-sdk-go"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
)

func TestLoadBalancerBackendsetsDatasource(t *testing.T) {
	client := GetTestProvider()
	providers := map[string]terraform.ResourceProvider{
		"baremetal": Provider(func(d *schema.ResourceData) (interface{}, error) {
			return client, nil
		}),
	}
	resourceName := "data.baremetal_load_balancer_backendsets.t"
	config := `
data "baremetal_load_balancer_backendsets" "t" {
  load_balancer_id = "ocid1.loadbalancer.stub_id"
}
`
	config += testProviderConfig()

	loadbalancerID := "ocid1.loadbalancer.stub_id"
	list := &baremetal.ListBackendSets{
		BackendSets: []baremetal.BackendSet{
			{Name: "stub_name1"},
			{Name: "stub_name2"},
		},
	}
	client.On(
		"ListBackendSets",
		loadbalancerID,
		(*baremetal.ClientRequestOptions)(nil),
	).Return(list, nil)

	resource.UnitTest(t, resource.TestCase{
		PreventPostDestroyRefresh: true,
		Providers:                 providers,
		Steps: []resource.TestStep{
			{
				Config: config,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "load_balancer_id", loadbalancerID),
					resource.TestCheckResourceAttr(resourceName, "backendsets.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "backendsets.0.name", "stub_name1"),
					resource.TestCheckResourceAttr(resourceName, "backendsets.1.name", "stub_name2"),
				),
			},
		},
	})
}
