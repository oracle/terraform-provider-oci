// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"testing"

	"github.com/oracle/terraform-provider-oci/internal/acctest"
	"github.com/oracle/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"github.com/oracle/terraform-provider-oci/httpreplay"
)

var (
	BlockchainblockchainPlatformPatchDataSourceRepresentation = map[string]interface{}{
		"blockchain_platform_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_blockchain_blockchain_platform.test_blockchain_platform.id}`},
	}

	BlockchainBlockchainPlatformPatchResourceConfig = acctest.GenerateResourceFromRepresentationMap("oci_blockchain_blockchain_platform", "test_blockchain_platform", acctest.Required, acctest.Create, blockchainPlatformRepresentation)
)

// issue-routing-tag: blockchain/default
func TestBlockchainBlockchainPlatformPatchResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestBlockchainBlockchainPlatformPatchResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	idcsAccessToken := utils.GetEnvSettingWithBlankDefault("idcs_access_token")
	idcsAccessTokenVariableStr := fmt.Sprintf("variable \"idcs_access_token\" { default = \"%s\" }\n", idcsAccessToken)

	datasourceName := "data.oci_blockchain_blockchain_platform_patches.test_blockchain_platform_patches"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config + compartmentIdVariableStr + idcsAccessTokenVariableStr + BlockchainBlockchainPlatformPatchResourceConfig +
				acctest.GenerateDataSourceFromRepresentationMap("oci_blockchain_blockchain_platform_patches", "test_blockchain_platform_patches", acctest.Required, acctest.Create, BlockchainblockchainPlatformPatchDataSourceRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "blockchain_platform_patch_collection.#"),
				resource.TestCheckResourceAttr(datasourceName, "blockchain_platform_patch_collection.0.items.#", "0"),
			),
		},
	})
}
