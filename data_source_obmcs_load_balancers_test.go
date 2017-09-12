// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package main

import (
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
)

func TestLoadBalancersDatasource(t *testing.T) {
	client := GetTestProvider()
	provider := Provider(func(d *schema.ResourceData) (interface{}, error) {
		return client, nil
	})
	providers := map[string]terraform.ResourceProvider{
		"oci": provider,
	}
	resourceName := "data.oci_load_balancers.t"
	config := `
data "oci_load_balancers" "t" {
  compartment_id = "${var.compartment_id}"
}
`
	config += testProviderConfig()

	compartmentID := "${var.compartment_id}"

	resource.UnitTest(t, resource.TestCase{
		PreventPostDestroyRefresh: true,
		Providers:                 providers,
		Steps: []resource.TestStep{
			{
				Config: config,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentID),
					resource.TestCheckResourceAttr(resourceName, "load_balancers.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "load_balancers.0.id", "ocid1.loadbalancer.stub_id1"),
					resource.TestCheckResourceAttr(resourceName, "load_balancers.1.id", "ocid1.loadbalancer.stub_id2"),
				),
			},
		},
	})
}
