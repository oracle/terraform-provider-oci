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

func TestLoadBalancerProtocolsDatasource(t *testing.T) {
	client := &mocks.BareMetalClient{}
	providers := map[string]terraform.ResourceProvider{
		"baremetal": Provider(func(d *schema.ResourceData) (interface{}, error) {
			return client, nil
		}),
	}
	resourceName := "data.baremetal_load_balancer_protocols.t"
	config := `
data "baremetal_load_balancer_protocols" "t" {
  compartment_id = "ocid1.compartment.stub_id"
}
`
	config += testProviderConfig

	compartmentID := "ocid1.compartment.stub_id"
	list := &baremetal.ListLoadBalancerProtocols{
		LoadBalancerProtocols: []baremetal.LoadBalancerProtocol{
			{Name: "stub_name1"},
			{Name: "stub_name2"},
		},
	}
	client.On(
		"ListLoadBalancerProtocols",
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
					resource.TestCheckResourceAttr(resourceName, "protocols.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "protocols.0.name", "stub_name1"),
					resource.TestCheckResourceAttr(resourceName, "protocols.1.name", "stub_name2"),
				),
			},
		},
	})
}
