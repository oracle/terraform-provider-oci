// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package main

import (
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
)

func TestLoadBalancerBackendsetsDatasource(t *testing.T) {
	client := GetTestProvider()
	providers := map[string]terraform.ResourceProvider{
		"oci": Provider(func(d *schema.ResourceData) (interface{}, error) {
			return client, nil
		}),
	}
	resourceName := "data.oci_load_balancer_backendsets.t"
	config := `
data "oci_load_balancer_backendsets" "t" {
  load_balancer_id = "ocid1.loadbalancer.stub_id"
}
`
	config += testProviderConfig()

	loadbalancerID := "ocid1.loadbalancer.stub_id"

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
