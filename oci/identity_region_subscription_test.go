// Copyright (c) 2017, 2020, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	regionSubscriptionDataSourceRepresentation = map[string]interface{}{
		"tenancy_id": Representation{repType: Required, create: `${var.tenancy_ocid}`},
		"filter":     RepresentationGroup{Required, regionSubscriptionDataSourceFilterRepresentation},
	}

	regionSubscriptionDataSourceFilterRepresentation = map[string]interface{}{
		"name":   Representation{repType: Required, create: `is_home_region`},
		"values": Representation{repType: Required, create: []string{`true`}},
	}

	RegionSubscriptionResourceConfig = ""
)

func TestIdentityRegionSubscriptionResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestIdentityRegionSubscriptionResource_basic")
	defer httpreplay.SaveScenario()

	provider := testAccProvider
	config := testProviderConfig()

	datasourceName := "data.oci_identity_region_subscriptions.test_region_subscriptions"

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		Steps: []resource.TestStep{
			// verify datasource
			{
				Config: config +
					generateDataSourceFromRepresentationMap("oci_identity_region_subscriptions", "test_region_subscriptions", Required, Create, regionSubscriptionDataSourceRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(datasourceName, "tenancy_id"),
					resource.TestCheckResourceAttr(datasourceName, "region_subscriptions.#", "1"),
					resource.TestCheckResourceAttr(datasourceName, "region_subscriptions.0.is_home_region", "true"),
					resource.TestCheckResourceAttrSet(datasourceName, "region_subscriptions.0.region_key"),
					resource.TestCheckResourceAttr(datasourceName, "region_subscriptions.0.region_name", getRequiredEnvSetting("region")),
					resource.TestCheckResourceAttrSet(datasourceName, "region_subscriptions.0.state"),
				),
			},
		},
	})
}
