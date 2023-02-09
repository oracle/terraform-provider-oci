// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	MarketplaceMarketplaceListingPackageAgreementDataSourceRepresentation = map[string]interface{}{
		"agreement_id":    acctest.Representation{RepType: acctest.Required, Create: `${data.oci_marketplace_listing_package_agreements.test_listing_package_agreements.agreements.0.id}`},
		"listing_id":      acctest.Representation{RepType: acctest.Required, Create: `${data.oci_marketplace_listing.test_listing.id}`},
		"package_version": acctest.Representation{RepType: acctest.Required, Create: `${data.oci_marketplace_listing.test_listing.default_package_version}`},
		"compartment_id":  acctest.Representation{RepType: acctest.Optional, Create: `${var.compartment_id}`},
	}

	MarketplaceMarketplaceListingPackageAgreementsDataSourceRepresentation = map[string]interface{}{
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

	listDatasourceName := "data.oci_marketplace_listing_package_agreements.test_listing_package_agreements"
	getDatasourceName := "data.oci_marketplace_listing_package_agreement.test_listing_package_agreement"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify get datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_marketplace_listing_package_agreement", "test_listing_package_agreement", acctest.Required, acctest.Create, MarketplaceMarketplaceListingPackageAgreementDataSourceRepresentation) +
				acctest.GenerateDataSourceFromRepresentationMap("oci_marketplace_listing_package_agreements", "test_listing_package_agreements", acctest.Required, acctest.Create, MarketplaceMarketplaceListingPackageAgreementsDataSourceRepresentation) +
				compartmentIdVariableStr + MarketplaceListingPackageAgreementResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(getDatasourceName, "agreement_id"),
				resource.TestCheckResourceAttrSet(getDatasourceName, "listing_id"),
				resource.TestCheckResourceAttrSet(getDatasourceName, "package_version"),

				resource.TestCheckResourceAttrSet(getDatasourceName, "content_url"),
				resource.TestCheckResourceAttrSet(getDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(getDatasourceName, "prompt"),
				resource.TestCheckResourceAttrSet(getDatasourceName, "signature"),
			),
		},
		// verify list datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_marketplace_listing_package_agreements", "test_listing_package_agreements", acctest.Required, acctest.Create, MarketplaceMarketplaceListingPackageAgreementsDataSourceRepresentation) +
				compartmentIdVariableStr + MarketplaceListingPackageAgreementResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(listDatasourceName, "listing_id"),

				resource.TestCheckResourceAttrSet(listDatasourceName, "agreements.#"),
				resource.TestCheckResourceAttrSet(listDatasourceName, "agreements.0.author"),
				resource.TestCheckResourceAttrSet(listDatasourceName, "agreements.0.content_url"),
				resource.TestCheckResourceAttrSet(listDatasourceName, "agreements.0.id"),
				resource.TestCheckResourceAttrSet(listDatasourceName, "agreements.0.prompt"),
			),
		},
	})
}
