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
	appCatalogListingSingularDataSourceRepresentation = map[string]interface{}{
		"listing_id": Representation{repType: Required, create: `${lookup(data.oci_core_app_catalog_listings.test_app_catalog_listings.app_catalog_listings[0],"listing_id")}`},
	}

	appCatalogListingDataSourceRepresentation = map[string]interface{}{}

	AppCatalogListingResourceConfig = ""
)

func TestCoreAppCatalogListingResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestCoreAppCatalogListingResource_basic")
	defer httpreplay.SaveScenario()

	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_core_app_catalog_listings.test_app_catalog_listings"
	singularDatasourceName := "data.oci_core_app_catalog_listing.test_app_catalog_listing"

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
					generateDataSourceFromRepresentationMap("oci_core_app_catalog_listings", "test_app_catalog_listings", Required, Create, appCatalogListingDataSourceRepresentation) +
					compartmentIdVariableStr + AppCatalogListingResourceConfig,
				Check: ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttrSet(datasourceName, "app_catalog_listings.#"),
					resource.TestCheckResourceAttrSet(datasourceName, "app_catalog_listings.0.display_name"),
					resource.TestCheckResourceAttrSet(datasourceName, "app_catalog_listings.0.listing_id"),
					resource.TestCheckResourceAttrSet(datasourceName, "app_catalog_listings.0.publisher_name"),
					resource.TestCheckResourceAttrSet(datasourceName, "app_catalog_listings.0.summary"),
				),
			},
			// verify singular datasource
			{
				Config: config +
					generateDataSourceFromRepresentationMap("oci_core_app_catalog_listings", "test_app_catalog_listings", Required, Create, appCatalogListingDataSourceRepresentation) +
					generateDataSourceFromRepresentationMap("oci_core_app_catalog_listing", "test_app_catalog_listing", Required, Create, appCatalogListingSingularDataSourceRepresentation) +
					compartmentIdVariableStr + AppCatalogListingResourceConfig,
				Check: ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttrSet(singularDatasourceName, "listing_id"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "contact_url"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "description"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "display_name"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "listing_id"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "publisher_logo_url"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "publisher_name"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "summary"),
				),
			},
		},
	})
}
