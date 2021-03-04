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
	listingPackageSingularDataSourceRepresentation = map[string]interface{}{
		"listing_id":      Representation{repType: Required, create: `${data.oci_marketplace_listing.test_listing.id}`},
		"compartment_id":  Representation{repType: Optional, create: `${var.compartment_id}`},
		"package_version": Representation{repType: Required, create: `${data.oci_marketplace_listing.test_listing.default_package_version}`},
	}

	listingPackageDataSourceRepresentation = map[string]interface{}{
		"listing_id":      Representation{repType: Required, create: `${data.oci_marketplace_listings.test_listings.listings.0.id}`},
		"compartment_id":  Representation{repType: Optional, create: `${var.compartment_id}`},
		"package_type":    Representation{repType: Optional, create: `packageType`},
		"package_version": Representation{repType: Optional, create: `packageVersion`},
	}

	ListingPackageResourceConfig = generateDataSourceFromRepresentationMap("oci_marketplace_listings", "test_listings", Required, Create, listingDataSourceRepresentation) +
		generateDataSourceFromRepresentationMap("oci_marketplace_listing", "test_listing", Required, Create, listingSingularDataSourceRepresentation)
)

func TestMarketplaceListingPackageResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestMarketplaceListingPackageResource_basic")
	defer httpreplay.SaveScenario()

	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_marketplace_listing_packages.test_listing_packages"
	singularDatasourceName := "data.oci_marketplace_listing_package.test_listing_package"

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
					generateDataSourceFromRepresentationMap("oci_marketplace_listing_packages", "test_listing_packages", Required, Create, listingPackageDataSourceRepresentation) +
					compartmentIdVariableStr + ListingPackageResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(datasourceName, "listing_id"),

					resource.TestCheckResourceAttrSet(datasourceName, "listing_packages.#"),
					resource.TestCheckResourceAttrSet(datasourceName, "listing_packages.0.listing_id"),
					resource.TestCheckResourceAttrSet(datasourceName, "listing_packages.0.resource_id"),
					resource.TestCheckResourceAttrSet(datasourceName, "listing_packages.0.time_created"),
					resource.TestCheckResourceAttrSet(datasourceName, "listing_packages.0.package_type"),
					resource.TestCheckResourceAttrSet(datasourceName, "listing_packages.0.package_version"),
				),
			},
			// verify singular datasource
			{
				Config: config +
					generateDataSourceFromRepresentationMap("oci_marketplace_listing_package", "test_listing_package", Required, Create, listingPackageSingularDataSourceRepresentation) +
					compartmentIdVariableStr + ListingPackageResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(singularDatasourceName, "listing_id"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "package_version"),

					resource.TestCheckResourceAttrSet(singularDatasourceName, "app_catalog_listing_id"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "app_catalog_listing_resource_version"),
					resource.TestCheckResourceAttr(singularDatasourceName, "operating_system.#", "0"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "image_id"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "package_type"),
					resource.TestCheckResourceAttr(singularDatasourceName, "pricing.#", "1"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "pricing.0.rate"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "pricing.0.type"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "resource_id"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "version"),
				),
			},
		},
	})
}
