// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package testing

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	identityProviderGroupDataSourceRepresentation = map[string]interface{}{
		"identity_provider_id": Representation{RepType: Required, Create: `${oci_identity_identity_provider.test_identity_provider.id}`},
		"name":                 Representation{RepType: Optional, Create: `test-idp-saml2-adfs`},
		"state":                Representation{RepType: Optional, Create: `ACTIVE`},
	}

	IdentityProviderGroupResourceConfig = acctest.GenerateResourceFromRepresentationMap("oci_identity_identity_provider", "test_identity_provider", Required, Create, identityProviderRepresentation) +
		IdentityProviderPropertyVariables
)

// issue-routing-tag: identity/default
func TestIdentityIdentityProviderGroupResource_basic(t *testing.T) {
	metadataFile := utils.GetEnvSettingWithBlankDefault("identity_provider_metadata_file")
	if metadataFile == "" {
		t.Skip("Skipping generated test for now as it has a dependency on federation metadata file")
	}

	httpreplay.SetScenario("TestIdentityIdentityProviderGroupResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_identity_identity_provider_groups.test_identity_provider_groups"

	acctest.SaveConfigContent("", "", "", t)

	_, tokenFn := acctest.TokenizeWithHttpReplay("identity_group_resource")
	IdentityProviderGroupResourceConfig = tokenFn(IdentityProviderGroupResourceConfig, map[string]string{"metadata_file": metadataFile})

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_identity_identity_provider_groups", "test_identity_provider_groups", Optional, Create, identityProviderGroupDataSourceRepresentation) +
				compartmentIdVariableStr + IdentityProviderGroupResourceConfig,
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "identity_provider_id"),
				resource.TestCheckResourceAttr(datasourceName, "name", "test-idp-saml2-adfs"),
				resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),

				resource.TestCheckResourceAttrSet(datasourceName, "identity_provider_groups.#"),
			),
		},
	})
}
