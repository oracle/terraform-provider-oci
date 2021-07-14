// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	tenancySingularDataSourceRepresentation = map[string]interface{}{
		"tenancy_id": Representation{repType: Required, create: `${var.tenancy_ocid}`},
	}

	TenancyResourceConfig = ""
)

func TestIdentityTenancyResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestIdentityTenancyResource_basic")
	defer httpreplay.SaveScenario()

	provider := testAccProvider
	config := testProviderConfig()

	singularDatasourceName := "data.oci_identity_tenancy.test_tenancy"

	saveConfigContent("", "", "", t)

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		Steps: []resource.TestStep{
			// verify singular datasource
			{
				Config: config +
					generateDataSourceFromRepresentationMap("oci_identity_tenancy", "test_tenancy", Required, Create, tenancySingularDataSourceRepresentation),
				Check: ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttrSet(singularDatasourceName, "tenancy_id"),

					resource.TestCheckResourceAttrSet(singularDatasourceName, "description"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "home_region_key"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "name"),
				),
			},
		},
	})
}
