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
	listingSingularDataSourceRepresentation = map[string]interface{}{
		"listing_id":     acctest.Representation{RepType: acctest.Required, Create: `${data.oci_marketplace_listings.test_listings.listings.0.id}`},
		"compartment_id": acctest.Representation{RepType: acctest.Optional, Create: `${var.compartment_id}`},
	}

	listingDataSourceRepresentation = map[string]interface{}{
		"category":          acctest.Representation{RepType: acctest.Optional, Create: []string{`category`}},
		"compartment_id":    acctest.Representation{RepType: acctest.Optional, Create: `${var.compartment_id}`},
		"is_featured":       acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"listing_id":        acctest.Representation{RepType: acctest.Optional, Create: `${oci_marketplace_listing.test_listing.id}`},
		"listing_types":     acctest.Representation{RepType: acctest.Optional, Create: []string{`listingTypes`}},
		"name":              acctest.Representation{RepType: acctest.Optional, Create: []string{`name`}},
		"operating_systems": acctest.Representation{RepType: acctest.Optional, Create: []string{`operatingSystems`}},
		"package_type":      acctest.Representation{RepType: acctest.Optional, Create: `packageType`},
		"pricing":           acctest.Representation{RepType: acctest.Optional, Create: []string{`pricing`}},
		"publisher_id":      acctest.Representation{RepType: acctest.Optional, Create: `${oci_marketplace_publisher.test_publisher.id}`},
	}

	ListingResourceConfig = ``
)

// issue-routing-tag: marketplace/default
func TestMarketplaceListingResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestMarketplaceListingResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_marketplace_listings.test_listings"
	singularDatasourceName := "data.oci_marketplace_listing.test_listing"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_marketplace_listings", "test_listings", acctest.Required, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(listingDataSourceRepresentation, map[string]interface{}{
						"name": acctest.Representation{RepType: acctest.Required, Create: []string{`FortiGate Next-Gen Firewall (2 cores)`}},
					})) +
				compartmentIdVariableStr + ListingResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(

				resource.TestCheckResourceAttrSet(datasourceName, "listings.#"),
				resource.TestCheckResourceAttrSet(datasourceName, "listings.0.categories.#"),
				resource.TestCheckResourceAttrSet(datasourceName, "listings.0.icon.#"),
				resource.TestCheckResourceAttrSet(datasourceName, "listings.0.id"),
				resource.TestCheckResourceAttrSet(datasourceName, "listings.0.is_featured"),
				resource.TestCheckResourceAttrSet(datasourceName, "listings.0.name"),
				resource.TestCheckResourceAttrSet(datasourceName, "listings.0.regions.#"),
				resource.TestCheckResourceAttrSet(datasourceName, "listings.0.package_type"),
				resource.TestCheckResourceAttrSet(datasourceName, "listings.0.publisher.#"),
				resource.TestCheckResourceAttrSet(datasourceName, "listings.0.short_description"),
				resource.TestCheckResourceAttrSet(datasourceName, "listings.0.supported_operating_systems.#"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_marketplace_listings", "test_listings", acctest.Required, acctest.Create, listingDataSourceRepresentation) +
				acctest.GenerateDataSourceFromRepresentationMap("oci_marketplace_listing", "test_listing", acctest.Required, acctest.Create, listingSingularDataSourceRepresentation) +
				compartmentIdVariableStr + ListingResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "listing_id"),

				resource.TestCheckResourceAttrSet(singularDatasourceName, "categories.#"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "default_package_version"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "documentation_links.#"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "icon.#"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "is_featured"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "keywords"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "languages.#"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "links.#"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "long_description"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "name"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "package_type"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "publisher.#"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "regions.#"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "screenshots.#"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "short_description"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "supported_operating_systems.#"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "support_contacts.#"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "support_links.#"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "system_requirements"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "tagline"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "usage_information"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "videos.#"),
			),
		},
	})
}
