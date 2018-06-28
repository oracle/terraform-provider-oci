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

	singularDatasourceName := "data.oci_identity_tenancy.test_tenancy"

	resource.Test(t, resource.TestCase{
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		Steps: []resource.TestStep{
			// verify singular datasource
			{
				Config: config + `

data "oci_identity_tenancy" "test_tenancy" {
	#Required
	tenancy_id = "${var.tenancy_ocid}"
}
                `,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(singularDatasourceName, "tenancy_id"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "description"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "home_region_key"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "name"),
				),
			},
		},
	})
}
