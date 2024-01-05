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
	IdentityIdentityAllowedDomainLicenseTypeDataSourceRepresentation = map[string]interface{}{
		"current_license_type_name": acctest.Representation{RepType: acctest.Optional, Create: `free`},
	}

	IdentityAllowedDomainLicenseTypeResourceConfig = ""
)

// issue-routing-tag: identity/default
func TestIdentityAllowedDomainLicenseTypeResource_basic(t *testing.T) {
	t.Skip("Skip this test because henosis tenancy is needed")
	httpreplay.SetScenario("TestIdentityAllowedDomainLicenseTypeResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_identity_allowed_domain_license_types.test_allowed_domain_license_types"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource without current_license_type_name param
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_identity_allowed_domain_license_types", "test_allowed_domain_license_types", acctest.Required, acctest.Create, IdentityIdentityAllowedDomainLicenseTypeDataSourceRepresentation) +
				compartmentIdVariableStr + IdentityAllowedDomainLicenseTypeResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "allowed_domain_license_types.#"),
				resource.TestCheckResourceAttrSet(datasourceName, "allowed_domain_license_types.0.description"),
				resource.TestCheckResourceAttrSet(datasourceName, "allowed_domain_license_types.0.license_type"),
				resource.TestCheckResourceAttrSet(datasourceName, "allowed_domain_license_types.0.name"),
			),
		},
		// verify datasource with current_license_type_name param
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_identity_allowed_domain_license_types", "test_allowed_domain_license_types", acctest.Optional, acctest.Create, IdentityIdentityAllowedDomainLicenseTypeDataSourceRepresentation) +
				compartmentIdVariableStr + IdentityAllowedDomainLicenseTypeResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "current_license_type_name", "free"),

				resource.TestCheckResourceAttrSet(datasourceName, "allowed_domain_license_types.#"),
				resource.TestCheckResourceAttrSet(datasourceName, "allowed_domain_license_types.0.description"),
				resource.TestCheckResourceAttrSet(datasourceName, "allowed_domain_license_types.0.license_type"),
				resource.TestCheckResourceAttrSet(datasourceName, "allowed_domain_license_types.0.name"),
			),
		},
	})
}
