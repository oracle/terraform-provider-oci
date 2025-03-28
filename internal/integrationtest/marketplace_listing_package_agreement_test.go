// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	MarketplaceListingPackageAgreementManagementRepresentation = map[string]interface{}{
		"agreement_id":    acctest.Representation{RepType: acctest.Required, Create: `${data.oci_marketplace_listing_package_agreements.test_listing_package_agreements.agreements.0.id}`},
		"listing_id":      acctest.Representation{RepType: acctest.Required, Create: `${data.oci_marketplace_listing.test_listing.id}`},
		"package_version": acctest.Representation{RepType: acctest.Required, Create: `${data.oci_marketplace_listing.test_listing.default_package_version}`},
		"compartment_id":  acctest.Representation{RepType: acctest.Optional, Create: `${var.compartment_id}`},
	}

	MarketplaceMarketplaceListingPackageAgreementDataSourceRepresentation = map[string]interface{}{
		"listing_id":      acctest.Representation{RepType: acctest.Required, Create: `${data.oci_marketplace_listing.test_listing.id}`},
		"package_version": acctest.Representation{RepType: acctest.Required, Create: `${data.oci_marketplace_listing.test_listing.default_package_version}`},
		"compartment_id":  acctest.Representation{RepType: acctest.Optional, Create: `${var.compartment_id}`},
	}

	MarketplaceListingPackageAgreementResourceConfig = acctest.GenerateDataSourceFromRepresentationMap("oci_marketplace_listing", "test_listing", acctest.Required, acctest.Create, MarketplaceMarketplaceListingSingularDataSourceRepresentation) +
		acctest.GenerateDataSourceFromRepresentationMap("oci_marketplace_listings", "test_listings", acctest.Required, acctest.Create, MarketplaceMarketplaceListingDataSourceRepresentation)
)

// issue-routing-tag: marketplace/default
func TestMarketplaceListingPackageAgreementResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestMarketplaceListingPackageAgreementResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_marketplace_listing_package_agreements.test_listing_package_agreements"
	resourceName := "oci_marketplace_listing_package_agreement.test_listing_package_agreement"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify resource
		{
			Config: config +
				acctest.GenerateResourceFromRepresentationMap("oci_marketplace_listing_package_agreement", "test_listing_package_agreement", acctest.Required, acctest.Create, MarketplaceListingPackageAgreementManagementRepresentation) +
				acctest.GenerateDataSourceFromRepresentationMap("oci_marketplace_listing_package_agreements", "test_listing_package_agreements", acctest.Required, acctest.Create, MarketplaceMarketplaceListingPackageAgreementDataSourceRepresentation) +
				compartmentIdVariableStr + MarketplaceListingPackageAgreementResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "agreement_id"),
				resource.TestCheckResourceAttrSet(resourceName, "listing_id"),
				resource.TestCheckResourceAttrSet(resourceName, "package_version"),

				resource.TestCheckResourceAttrSet(resourceName, "content_url"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "prompt"),
				resource.TestCheckResourceAttrSet(resourceName, "signature"),
			),
		},
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_marketplace_listing_package_agreements", "test_listing_package_agreements", acctest.Required, acctest.Create, MarketplaceMarketplaceListingPackageAgreementDataSourceRepresentation) +
				compartmentIdVariableStr + MarketplaceListingPackageAgreementResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "listing_id"),

				resource.TestCheckResourceAttrSet(datasourceName, "agreements.#"),
				resource.TestCheckResourceAttrSet(datasourceName, "agreements.0.author"),
				resource.TestCheckResourceAttrSet(datasourceName, "agreements.0.content_url"),
				resource.TestCheckResourceAttrSet(datasourceName, "agreements.0.id"),
				resource.TestCheckResourceAttrSet(datasourceName, "agreements.0.prompt"),
			),
		},
	})
}
