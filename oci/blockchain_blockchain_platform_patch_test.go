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
	blockchainPlatformPatchDataSourceRepresentation = map[string]interface{}{
		"blockchain_platform_id": Representation{RepType: Required, Create: `${oci_blockchain_blockchain_platform.test_blockchain_platform.id}`},
	}

	BlockchainPlatformPatchResourceConfig = GenerateResourceFromRepresentationMap("oci_blockchain_blockchain_platform", "test_blockchain_platform", Required, Create, blockchainPlatformRepresentation)
)

// issue-routing-tag: blockchain/default
func TestBlockchainBlockchainPlatformPatchResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestBlockchainBlockchainPlatformPatchResource_basic")
	defer httpreplay.SaveScenario()

	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	idcsAccessToken := getEnvSettingWithBlankDefault("idcs_access_token")
	idcsAccessTokenVariableStr := fmt.Sprintf("variable \"idcs_access_token\" { default = \"%s\" }\n", idcsAccessToken)

	datasourceName := "data.oci_blockchain_blockchain_platform_patches.test_blockchain_platform_patches"

	SaveConfigContent("", "", "", t)

	ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config + compartmentIdVariableStr + idcsAccessTokenVariableStr + BlockchainPlatformPatchResourceConfig +
				GenerateDataSourceFromRepresentationMap("oci_blockchain_blockchain_platform_patches", "test_blockchain_platform_patches", Required, Create, blockchainPlatformPatchDataSourceRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "blockchain_platform_patch_collection.#"),
				resource.TestCheckResourceAttr(datasourceName, "blockchain_platform_patch_collection.0.items.#", "0"),
			),
		},
	})
}
