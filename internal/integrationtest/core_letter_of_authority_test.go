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
	letterOfAuthoritySingularDataSourceRepresentation = map[string]interface{}{
		"cross_connect_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_core_cross_connect.test_cross_connect.id}`},
	}

	LetterOfAuthorityResourceConfig = acctest.GenerateDataSourceFromRepresentationMap("oci_core_cross_connect_locations", "test_cross_connect_locations", acctest.Required, acctest.Create, crossConnectLocationDataSourceRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_cross_connect", "test_cross_connect", acctest.Required, acctest.Create, crossConnectRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_kms_vault", "test_vault", acctest.Required, acctest.Create, vaultRepresentation) +
		acctest.GenerateDataSourceFromRepresentationMap("oci_vault_secrets", "test_secrets", acctest.Required, acctest.Create, secretDataSourceRepresentation)
)

// issue-routing-tag: core/default
func TestCoreLetterOfAuthorityResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestCoreLetterOfAuthorityResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	singularDatasourceName := "data.oci_core_letter_of_authority.test_letter_of_authority"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_core_letter_of_authority", "test_letter_of_authority", acctest.Required, acctest.Create, letterOfAuthoritySingularDataSourceRepresentation) +
				compartmentIdVariableStr + LetterOfAuthorityResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "cross_connect_id"),

				resource.TestCheckResourceAttrSet(singularDatasourceName, "authorized_entity_name"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "circuit_type"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "facility_location"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "port_name"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_expires"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_issued"),
			),
		},
	})
}
