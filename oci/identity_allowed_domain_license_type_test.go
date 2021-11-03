// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	allowedDomainLicenseTypeDataSourceRepresentation = map[string]interface{}{
		"current_license_type_name": Representation{RepType: Optional, Create: `free`},
	}

	AllowedDomainLicenseTypeResourceConfig = ""
)

// issue-routing-tag: identity/default
func TestIdentityAllowedDomainLicenseTypeResource_basic(t *testing.T) {
	t.Skip("Skip this test because henosis tenancy is needed")
	httpreplay.SetScenario("TestIdentityAllowedDomainLicenseTypeResource_basic")
	defer httpreplay.SaveScenario()

	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_identity_allowed_domain_license_types.test_allowed_domain_license_types"

	SaveConfigContent("", "", "", t)

	ResourceTest(t, nil, []resource.TestStep{
		// verify datasource without current_license_type_name param
		{
			Config: config +
				GenerateDataSourceFromRepresentationMap("oci_identity_allowed_domain_license_types", "test_allowed_domain_license_types", Required, Create, allowedDomainLicenseTypeDataSourceRepresentation) +
				compartmentIdVariableStr + AllowedDomainLicenseTypeResourceConfig,
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "allowed_domain_license_types.#"),
				resource.TestCheckResourceAttrSet(datasourceName, "allowed_domain_license_types.0.description"),
				resource.TestCheckResourceAttrSet(datasourceName, "allowed_domain_license_types.0.license_type"),
				resource.TestCheckResourceAttrSet(datasourceName, "allowed_domain_license_types.0.name"),
			),
		},
		// verify datasource with current_license_type_name param
		{
			Config: config +
				GenerateDataSourceFromRepresentationMap("oci_identity_allowed_domain_license_types", "test_allowed_domain_license_types", Optional, Create, allowedDomainLicenseTypeDataSourceRepresentation) +
				compartmentIdVariableStr + AllowedDomainLicenseTypeResourceConfig,
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "current_license_type_name", "free"),

				resource.TestCheckResourceAttrSet(datasourceName, "allowed_domain_license_types.#"),
				resource.TestCheckResourceAttrSet(datasourceName, "allowed_domain_license_types.0.description"),
				resource.TestCheckResourceAttrSet(datasourceName, "allowed_domain_license_types.0.license_type"),
				resource.TestCheckResourceAttrSet(datasourceName, "allowed_domain_license_types.0.name"),
			),
		},
	})
}
