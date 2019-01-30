// Copyright (c) 2017, 2019, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

var (
	appCatalogListingResourceVersionSingularDataSourceRepresentation = map[string]interface{}{
		"listing_id":       Representation{repType: Required, create: `${lookup(data.oci_core_app_catalog_listing_resource_versions.test_app_catalog_listing_resource_versions.app_catalog_listing_resource_versions[0],"listing_id")}`},
		"resource_version": Representation{repType: Required, create: `${lookup(data.oci_core_app_catalog_listing_resource_versions.test_app_catalog_listing_resource_versions.app_catalog_listing_resource_versions[0],"listing_resource_version")}`},
	}

	appCatalogListingResourceVersionDataSourceRepresentation = map[string]interface{}{
		"listing_id": Representation{repType: Required, create: `${lookup(data.oci_core_app_catalog_listings.test_app_catalog_listings.app_catalog_listings[0],"listing_id")}`},
	}

	AppCatalogListingResourceVersionResourceConfig = `
	
	data "oci_core_app_catalog_listings" "test_app_catalog_listings" {}

	`
)

func TestCoreAppCatalogListingResourceVersionResource_basic(t *testing.T) {
	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_core_app_catalog_listing_resource_versions.test_app_catalog_listing_resource_versions"
	singularDatasourceName := "data.oci_core_app_catalog_listing_resource_version.test_app_catalog_listing_resource_version"

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		Steps: []resource.TestStep{
			// verify datasource
			{
				Config: config +
					generateDataSourceFromRepresentationMap("oci_core_app_catalog_listing_resource_versions", "test_app_catalog_listing_resource_versions", Required, Create, appCatalogListingResourceVersionDataSourceRepresentation) +
					compartmentIdVariableStr + AppCatalogListingResourceVersionResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(datasourceName, "listing_id"),

					resource.TestCheckResourceAttrSet(datasourceName, "app_catalog_listing_resource_versions.#"),
					resource.TestCheckResourceAttrSet(datasourceName, "app_catalog_listing_resource_versions.0.listing_id"),
				),
			},
			// verify singular datasource
			{
				Config: config +
					generateDataSourceFromRepresentationMap("oci_core_app_catalog_listing_resource_versions", "test_app_catalog_listing_resource_versions", Required, Create, appCatalogListingResourceVersionDataSourceRepresentation) +
					generateDataSourceFromRepresentationMap("oci_core_app_catalog_listing_resource_version", "test_app_catalog_listing_resource_version", Required, Create, appCatalogListingResourceVersionSingularDataSourceRepresentation) +
					compartmentIdVariableStr + AppCatalogListingResourceVersionResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(singularDatasourceName, "listing_id"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "listing_resource_version"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "resource_version"),
				),
			},
		},
	})
}
