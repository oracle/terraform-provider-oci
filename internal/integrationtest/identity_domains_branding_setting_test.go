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
	IdentityDomainsBrandingSettingSingularDataSourceRepresentation = map[string]interface{}{
		"branding_setting_id": acctest.Representation{RepType: acctest.Required, Create: `BrandingSettings`},
		"idcs_endpoint":       acctest.Representation{RepType: acctest.Required, Create: `${data.oci_identity_domain.test_domain.url}`},
		"attribute_sets":      acctest.Representation{RepType: acctest.Optional, Create: []string{`all`}},
	}

	IdentityDomainsBrandingSettingDataSourceRepresentation = map[string]interface{}{
		"idcs_endpoint":  acctest.Representation{RepType: acctest.Required, Create: `${data.oci_identity_domain.test_domain.url}`},
		"attribute_sets": acctest.Representation{RepType: acctest.Optional, Create: []string{`all`}},
	}

	// BrandingSettings API is GET only. Adding dependency resource of Settings here so that validation on attributes such as "custom_branding" etc. can be validated in the singular test case below.
	IdentityDomainsBrandingSettingResourceConfig = TestDomainDependencies + acctest.GenerateResourceFromRepresentationMap("oci_identity_domains_setting", "test_setting", acctest.Optional, acctest.Create, IdentityDomainsSettingRepresentation)
)

// issue-routing-tag: identity_domains/default
func TestIdentityDomainsBrandingSettingResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestIdentityDomainsBrandingSettingResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_identity_domains_branding_settings.test_branding_settings"
	singularDatasourceName := "data.oci_identity_domains_branding_setting.test_branding_setting"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_identity_domains_branding_settings", "test_branding_settings", acctest.Optional, acctest.Create, IdentityDomainsBrandingSettingDataSourceRepresentation) +
				compartmentIdVariableStr + IdentityDomainsBrandingSettingResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "attribute_sets.#", "1"),

				resource.TestCheckResourceAttrSet(datasourceName, "branding_settings.#"),
				resource.TestCheckResourceAttr(datasourceName, "branding_settings.0.schemas.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_identity_domains_branding_setting", "test_branding_setting", acctest.Optional, acctest.Create, IdentityDomainsBrandingSettingSingularDataSourceRepresentation) +
				compartmentIdVariableStr + IdentityDomainsBrandingSettingResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(singularDatasourceName, "attribute_sets.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "branding_setting_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "idcs_endpoint"),

				resource.TestCheckResourceAttr(singularDatasourceName, "company_names.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "custom_branding"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "custom_css_location"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "custom_html_location"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "custom_translation"),
				// Numbers of items in default_* properties are fixed by IDCS.
				resource.TestCheckResourceAttrSet(singularDatasourceName, "default_company_names.#"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "default_images.#"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "default_login_texts.#"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "enable_terms_of_use"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "external_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "images.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "is_hosted_page"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "locale"),
				resource.TestCheckResourceAttr(singularDatasourceName, "login_texts.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "preferred_language"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "privacy_policy_url"),
				resource.TestCheckResourceAttr(singularDatasourceName, "schemas.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "terms_of_use_url"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "timezone"),
			),
		},
	})
}
