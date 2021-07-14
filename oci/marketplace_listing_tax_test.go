// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	listingTaxDataSourceRepresentation = map[string]interface{}{
		"listing_id":     Representation{repType: Required, create: `${data.oci_marketplace_listings.test_listings.listings.0.id}`},
		"compartment_id": Representation{repType: Required, create: `${var.compartment_id}`},
	}

	ListingTaxResourceConfig = generateDataSourceFromRepresentationMap("oci_marketplace_listings", "test_listings", Required, Create, listingDataSourceRepresentation)
)

func TestMarketplaceListingTaxResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestMarketplaceListingTaxResource_basic")
	defer httpreplay.SaveScenario()

	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_marketplace_listing_taxes.test_listing_taxes"

	saveConfigContent("", "", "", t)

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		Steps: []resource.TestStep{
			// verify datasource
			{
				Config: config +
					generateDataSourceFromRepresentationMap("oci_marketplace_listing_taxes", "test_listing_taxes", Required, Create, listingTaxDataSourceRepresentation) +
					compartmentIdVariableStr + ListingTaxResourceConfig,
				Check: ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttrSet(datasourceName, "listing_id"),
				),
			},
		},
	})
}
