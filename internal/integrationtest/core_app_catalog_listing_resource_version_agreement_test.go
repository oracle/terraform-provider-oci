// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
	"github.com/terraform-providers/terraform-provider-oci/internal/acctest"
	"github.com/terraform-providers/terraform-provider-oci/internal/utils"
)

var (
	AppCatalogListingResourceVersionAgreementResourceConfig = AppCatalogListingResourceVersionAgreementResourceDependencies + `

resource "oci_core_app_catalog_listing_resource_version_agreement" "test_app_catalog_listing_resource_version_agreement" {
	#Required
	listing_id = "${lookup(data.oci_core_app_catalog_listing_resource_versions.test_app_catalog_listing_resource_versions.app_catalog_listing_resource_versions[0], "listing_id")}"
	listing_resource_version = "${lookup(data.oci_core_app_catalog_listing_resource_versions.test_app_catalog_listing_resource_versions.app_catalog_listing_resource_versions[0], "listing_resource_version")}"
}
`
	AppCatalogListingResourceVersionAgreementResourceDependencies = AppCatalogListingResourceVersionResourceConfig + `
	data oci_core_app_catalog_listing_resource_versions test_app_catalog_listing_resource_versions {
		listing_id = "${lookup(data.oci_core_app_catalog_listings.test_app_catalog_listings.app_catalog_listings[0],"listing_id")}"
	}
	`
)

// issue-routing-tag: core/computeImaging
func TestResourceAppCatalogListingResourceVersionAgreement_basic(t *testing.T) {
	httpreplay.SetScenario("TestCoreAppCatalogListingResourceVersionAgreementResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()
	RCF3339NanoReg := regexp.MustCompile(`^\d{4}-\d{2}-\d{2}(T|t)\d{2}:\d{2}:\d{2}\.(\d|\d{2}|\d{3})Z$`)
	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_core_app_catalog_listing_resource_version_agreement.test_app_catalog_listing_resource_version_agreement"

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// Create resource
		{
			Config: config + compartmentIdVariableStr + AppCatalogListingResourceVersionAgreementResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "listing_id"),
				resource.TestCheckResourceAttrSet(resourceName, "listing_resource_version"),
				resource.TestCheckResourceAttrSet(resourceName, "oracle_terms_of_use_link"),
				resource.TestCheckResourceAttrSet(resourceName, "signature"),
				resource.TestMatchResourceAttr(resourceName, "time_retrieved", RCF3339NanoReg),
			),
		},
	})
}
