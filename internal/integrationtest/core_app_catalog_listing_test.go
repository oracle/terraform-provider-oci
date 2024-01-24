// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
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
	CoreCoreAppCatalogListingSingularDataSourceRepresentation = map[string]interface{}{
		"listing_id": acctest.Representation{RepType: acctest.Required, Create: `${lookup(data.oci_core_app_catalog_listings.test_app_catalog_listings.app_catalog_listings[0],"listing_id")}`},
	}

	CoreCoreAppCatalogListingDataSourceRepresentation = map[string]interface{}{}

	CoreAppCatalogListingResourceConfig = ""
)

// issue-routing-tag: core/computeImaging
func TestCoreAppCatalogListingResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestCoreAppCatalogListingResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_core_app_catalog_listings.test_app_catalog_listings"
	singularDatasourceName := "data.oci_core_app_catalog_listing.test_app_catalog_listing"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_core_app_catalog_listings", "test_app_catalog_listings", acctest.Required, acctest.Create, CoreCoreAppCatalogListingDataSourceRepresentation) +
				compartmentIdVariableStr + CoreAppCatalogListingResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
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
				acctest.GenerateDataSourceFromRepresentationMap("oci_core_app_catalog_listings", "test_app_catalog_listings", acctest.Required, acctest.Create, CoreCoreAppCatalogListingDataSourceRepresentation) +
				acctest.GenerateDataSourceFromRepresentationMap("oci_core_app_catalog_listing", "test_app_catalog_listing", acctest.Required, acctest.Create, CoreCoreAppCatalogListingSingularDataSourceRepresentation) +
				compartmentIdVariableStr + CoreAppCatalogListingResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
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
	})
}
