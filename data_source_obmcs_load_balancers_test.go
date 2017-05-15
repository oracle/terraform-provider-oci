// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package main

import (
	"testing"
	"time"

	"github.com/MustWin/baremetal-sdk-go"
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
		"baremetal": provider,
	}
	resourceName := "data.baremetal_load_balancers.t"
	config := `
data "baremetal_load_balancers" "t" {
  compartment_id = "${var.compartment_id}"
}
`
	config += testProviderConfig()

	compartmentID := "${var.compartment_id}"
	lb1 := baremetal.LoadBalancer{
		CompartmentID: compartmentID,
		ID:            "ocid1.loadbalancer.stub_id1",
		IPAddresses:   []baremetal.IPAddress{{IPAddress: "1.2.3.4"}},
		Shape:         "stub_shape",
		State:         baremetal.ResourceActive,
		TimeCreated:   baremetal.Time{Time: time.Now()}, // FIXME: use baremetal.Time
	}

	lb2 := lb1
	lb2.ID = "ocid1.loadbalancer.stub_id2"

	list := &baremetal.ListLoadBalancers{
		LoadBalancers: []baremetal.LoadBalancer{lb1, lb2},
	}
	client.On(
		"ListLoadBalancers",
		compartmentID,
		(*baremetal.ListOptions)(nil),
	).Return(list, nil)

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
