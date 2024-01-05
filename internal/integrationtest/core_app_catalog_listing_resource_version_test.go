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
	CoreCoreAppCatalogListingResourceVersionSingularDataSourceRepresentation = map[string]interface{}{
		"listing_id":       acctest.Representation{RepType: acctest.Required, Create: `${lookup(data.oci_core_app_catalog_listing_resource_versions.test_app_catalog_listing_resource_versions.app_catalog_listing_resource_versions[0],"listing_id")}`},
		"resource_version": acctest.Representation{RepType: acctest.Required, Create: `${lookup(data.oci_core_app_catalog_listing_resource_versions.test_app_catalog_listing_resource_versions.app_catalog_listing_resource_versions[0],"listing_resource_version")}`},
	}

	CoreCoreAppCatalogListingResourceVersionDataSourceRepresentation = map[string]interface{}{
		"listing_id": acctest.Representation{RepType: acctest.Required, Create: `${lookup(data.oci_core_app_catalog_listings.test_app_catalog_listings.app_catalog_listings[0],"listing_id")}`},
	}

	CoreAppCatalogListingResourceVersionResourceConfig = `
	
	data "oci_core_app_catalog_listings" "test_app_catalog_listings" {}

	`
)

// issue-routing-tag: core/computeImaging
func TestCoreAppCatalogListingResourceVersionResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestCoreAppCatalogListingResourceVersionResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_core_app_catalog_listing_resource_versions.test_app_catalog_listing_resource_versions"
	singularDatasourceName := "data.oci_core_app_catalog_listing_resource_version.test_app_catalog_listing_resource_version"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_core_app_catalog_listing_resource_versions", "test_app_catalog_listing_resource_versions", acctest.Required, acctest.Create, CoreCoreAppCatalogListingResourceVersionDataSourceRepresentation) +
				compartmentIdVariableStr + CoreAppCatalogListingResourceVersionResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "listing_id"),

				resource.TestCheckResourceAttrSet(datasourceName, "app_catalog_listing_resource_versions.#"),
				resource.TestCheckResourceAttrSet(datasourceName, "app_catalog_listing_resource_versions.0.listing_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "app_catalog_listing_resource_versions.0.listing_resource_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "app_catalog_listing_resource_versions.0.listing_resource_version"),
				resource.TestCheckResourceAttrSet(datasourceName, "app_catalog_listing_resource_versions.0.time_published"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_core_app_catalog_listing_resource_versions", "test_app_catalog_listing_resource_versions", acctest.Required, acctest.Create, CoreCoreAppCatalogListingResourceVersionDataSourceRepresentation) +
				acctest.GenerateDataSourceFromRepresentationMap("oci_core_app_catalog_listing_resource_version", "test_app_catalog_listing_resource_version", acctest.Required, acctest.Create, CoreCoreAppCatalogListingResourceVersionSingularDataSourceRepresentation) +
				compartmentIdVariableStr + CoreAppCatalogListingResourceVersionResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "listing_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "resource_version"),

				resource.TestCheckResourceAttrSet(singularDatasourceName, "listing_resource_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "listing_resource_version"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_published"),
			),
		},
	})
}
