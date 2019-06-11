// Copyright (c) 2017, 2019, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"testing"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"

	"github.com/hashicorp/terraform/helper/resource"
)

func TestAccDatasourceLoadBalancerPolicies_basic(t *testing.T) {
	httpreplay.SetScenario("TestAccDatasourceLoadBalancerPolicies_basic")
	defer httpreplay.SaveScenario()
	providers := testAccProviders
	config := legacyTestProviderConfig() + `
	data "oci_load_balancer_policies" "t" {
		compartment_id = "${var.compartment_id}"
	}`

	resourceName := "data.oci_load_balancer_policies.t"

	resource.Test(t, resource.TestCase{
		PreCheck:                  func() { testAccPreCheck(t) },
		PreventPostDestroyRefresh: true,
		Providers:                 providers,
		Steps: []resource.TestStep{
			{
				Config: config,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "compartment_id"),
					resource.TestCheckResourceAttrSet(resourceName, "policies.#"),
					resource.TestCheckResourceAttrSet(resourceName, "policies.0.name"),
				),
			},
		},
	})
}
