// Copyright (c) 2017, 2019, Oracle and/or its affiliates. All rights reserved.

package oci

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
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

func TestCoreAppCatalogListingResourceVersionAgreementResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestCoreAppCatalogListingResourceVersionAgreementResource_basic")
	defer httpreplay.SaveScenario()

	provider := testAccProvider
	config := testProviderConfig()
	RCF3339NanoReg := regexp.MustCompile(`^\d{4}-\d{2}-\d{2}(T|t)\d{2}:\d{2}:\d{2}\.(\d|\d{2}|\d{3})Z$`)
	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_core_app_catalog_listing_resource_version_agreement.test_app_catalog_listing_resource_version_agreement"

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		Steps: []resource.TestStep{
			// create resource
			{
				Config: config + compartmentIdVariableStr + AppCatalogListingResourceVersionAgreementResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "listing_id"),
					resource.TestCheckResourceAttrSet(resourceName, "listing_resource_version"),
					resource.TestCheckResourceAttrSet(resourceName, "oracle_terms_of_use_link"),
					resource.TestCheckResourceAttrSet(resourceName, "signature"),
					resource.TestMatchResourceAttr(resourceName, "time_retrieved", RCF3339NanoReg),
				),
			},
		},
	})
}
