// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"

	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	MarketplaceMarketplaceMetadataPublicKeyDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Optional, Create: `${var.compartment_id}`},
	}

	MarketplaceMarketplaceMetadataPublicKeyResourceConfig = ""
)

// issue-routing-tag: marketplace/default
func TestMarketplaceMarketplaceMetadataPublicKeyResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestMarketplaceMarketplaceMetadataPublicKeyResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_marketplace_marketplace_metadata_public_keys.test_marketplace_metadata_public_keys"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_marketplace_marketplace_metadata_public_keys", "test_marketplace_metadata_public_keys", acctest.Required, acctest.Create, MarketplaceMarketplaceMetadataPublicKeyDataSourceRepresentation) +
				compartmentIdVariableStr + MarketplaceMarketplaceMetadataPublicKeyResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(

				resource.TestCheckResourceAttrSet(datasourceName, "marketplace_metadata_public_keys.#"),
				resource.TestCheckResourceAttr(datasourceName, "marketplace_metadata_public_keys.0.certificate_chain.#", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "marketplace_metadata_public_keys.0.certificate_thumbprint"),
				resource.TestCheckResourceAttrSet(datasourceName, "marketplace_metadata_public_keys.0.exponent"),
				resource.TestCheckResourceAttrSet(datasourceName, "marketplace_metadata_public_keys.0.key_algorithm"),
				resource.TestCheckResourceAttrSet(datasourceName, "marketplace_metadata_public_keys.0.key_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "marketplace_metadata_public_keys.0.key_type"),
				resource.TestCheckResourceAttrSet(datasourceName, "marketplace_metadata_public_keys.0.key_use"),
				resource.TestCheckResourceAttrSet(datasourceName, "marketplace_metadata_public_keys.0.modulus"),
			),
		},
	})
}
