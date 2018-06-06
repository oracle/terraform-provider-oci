// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

func TestIdentityRegionResource_basic(t *testing.T) {
	provider := testAccProvider
	config := testProviderConfig()

	datasourceName := "data.oci_identity_regions.test_regions"

	resource.Test(t, resource.TestCase{
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		Steps: []resource.TestStep{
			// verify datasource
			{
				Config: config + `
data "oci_identity_regions" "test_regions" {
	filter {
		name = "name"
		values = ["${var.region}"]
	}
}
                `,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(datasourceName, "regions.#", "1"),
				),
			},
		},
	})
}
