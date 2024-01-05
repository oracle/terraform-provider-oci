// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"

	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	IdentityDomainsIdentityDomainsUserAttributesSettingSingularDataSourceRepresentation = map[string]interface{}{
		"idcs_endpoint":              acctest.Representation{RepType: acctest.Required, Create: `${data.oci_identity_domain.test_domain.url}`},
		"user_attributes_setting_id": acctest.Representation{RepType: acctest.Required, Create: `UserAttributesSettings`},
		"attribute_sets":             acctest.Representation{RepType: acctest.Optional, Create: []string{`all`}},
	}

	IdentityDomainsIdentityDomainsUserAttributesSettingDataSourceRepresentation = map[string]interface{}{
		"idcs_endpoint":  acctest.Representation{RepType: acctest.Required, Create: `${data.oci_identity_domain.test_domain.url}`},
		"attribute_sets": acctest.Representation{RepType: acctest.Optional, Create: []string{`all`}},
	}

	IdentityDomainsUserAttributesSettingResourceConfig = TestDomainDependencies
)

// issue-routing-tag: identity_domains/default
func TestIdentityDomainsUserAttributesSettingResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestIdentityDomainsUserAttributesSettingResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_identity_domains_user_attributes_settings.test_user_attributes_settings"
	singularDatasourceName := "data.oci_identity_domains_user_attributes_setting.test_user_attributes_setting"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		//verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_identity_domains_user_attributes_settings", "test_user_attributes_settings", acctest.Required, acctest.Create, IdentityDomainsIdentityDomainsUserAttributesSettingDataSourceRepresentation) +
				compartmentIdVariableStr + IdentityDomainsUserAttributesSettingResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "user_attributes_settings.#"),
				resource.TestCheckResourceAttr(datasourceName, "user_attributes_settings.0.schemas.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_identity_domains_user_attributes_setting", "test_user_attributes_setting", acctest.Required, acctest.Create, IdentityDomainsIdentityDomainsUserAttributesSettingSingularDataSourceRepresentation) +
				compartmentIdVariableStr + IdentityDomainsUserAttributesSettingResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "idcs_endpoint"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "user_attributes_setting_id"),

				resource.TestMatchResourceAttr(singularDatasourceName, "attribute_settings.#", regexp.MustCompile("[1-9]+")),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "ocid"),
				resource.TestCheckResourceAttr(singularDatasourceName, "schemas.#", "1"),
			),
		},
	})
}
