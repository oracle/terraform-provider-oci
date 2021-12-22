// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
	"github.com/terraform-providers/terraform-provider-oci/internal/acctest"
	"github.com/terraform-providers/terraform-provider-oci/internal/utils"
)

var (
	listingTaxDataSourceRepresentation = map[string]interface{}{
		"listing_id":     acctest.Representation{RepType: acctest.Required, Create: `${data.oci_marketplace_listings.test_listings.listings.0.id}`},
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
	}

	ListingTaxResourceConfig = acctest.GenerateDataSourceFromRepresentationMap("oci_marketplace_listings", "test_listings", acctest.Required, acctest.Create, listingDataSourceRepresentation)
)

// issue-routing-tag: marketplace/default
func TestMarketplaceListingTaxResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestMarketplaceListingTaxResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_marketplace_listing_taxes.test_listing_taxes"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_marketplace_listing_taxes", "test_listing_taxes", acctest.Required, acctest.Create, listingTaxDataSourceRepresentation) +
				compartmentIdVariableStr + ListingTaxResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(datasourceName, "listing_id"),
			),
		},
	})
}
