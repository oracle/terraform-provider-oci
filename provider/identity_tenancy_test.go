// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

func TestIdentityTenancyResource_basic(t *testing.T) {
	provider := testAccProvider
	config := testProviderConfig()

	datasourceName := "data.oci_identity_tenancy.test_tenancy"

	resource.Test(t, resource.TestCase{
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		Steps: []resource.TestStep{
			// verify datasource
			{
				Config: config + `

data "oci_identity_tenancy" "test_tenancy" {
	tenancy_id = "${var.tenancy_ocid}"
}
                `,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(datasourceName, "tenancy_id"),
					resource.TestCheckResourceAttrSet(datasourceName, "description"),
					resource.TestCheckResourceAttrSet(datasourceName, "home_region_key"),
					resource.TestCheckResourceAttrSet(datasourceName, "name"),
				),
			},
		},
	})
}
