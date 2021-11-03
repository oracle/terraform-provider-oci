// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	listingTaxDataSourceRepresentation = map[string]interface{}{
		"listing_id":     Representation{RepType: Required, Create: `${data.oci_marketplace_listings.test_listings.listings.0.id}`},
		"compartment_id": Representation{RepType: Required, Create: `${var.compartment_id}`},
	}

	ListingTaxResourceConfig = GenerateDataSourceFromRepresentationMap("oci_marketplace_listings", "test_listings", Required, Create, listingDataSourceRepresentation)
)

// issue-routing-tag: marketplace/default
func TestMarketplaceListingTaxResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestMarketplaceListingTaxResource_basic")
	defer httpreplay.SaveScenario()

	config := ProviderTestConfig()

	compartmentId := GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_marketplace_listing_taxes.test_listing_taxes"

	SaveConfigContent("", "", "", t)

	ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				GenerateDataSourceFromRepresentationMap("oci_marketplace_listing_taxes", "test_listing_taxes", Required, Create, listingTaxDataSourceRepresentation) +
				compartmentIdVariableStr + ListingTaxResourceConfig,
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(datasourceName, "listing_id"),
			),
		},
	})
}
