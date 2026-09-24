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
	DataSafeCryptoAssessmentWalletDataSourceRepresentation = map[string]interface{}{
		"compartment_id":            acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"assessment_id":             acctest.Representation{RepType: acctest.Required, Create: `${var.crypto_assessment_id}`},
		"assessment_type":           acctest.Representation{RepType: acctest.Required, Create: `LATEST`},
		"compartment_id_in_subtree": acctest.Representation{RepType: acctest.Required, Create: `false`},
		"target_id":                 acctest.Representation{RepType: acctest.Required, Create: `${var.target_id}`},
	}

	DataSafeCryptoAssessmentWalletResourceConfig = ""
)

// issue-routing-tag: data_safe/default
func TestDataSafeCryptoAssessmentWalletResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDataSafeCryptoAssessmentWalletResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)
	cryptoAssessmentId := utils.GetEnvSettingWithBlankDefault("crypto_assessment_id")
	cryptoAssessmentIdVariableStr := fmt.Sprintf("variable \"crypto_assessment_id\" { default = \"%s\" }\n", cryptoAssessmentId)
	targetId := utils.GetEnvSettingWithBlankDefault("target_ocid")
	targetIdVariableStr := fmt.Sprintf("variable \"target_id\" { default = \"%s\" }\n", targetId)

	datasourceName := "data.oci_data_safe_crypto_assessment_wallets.test_crypto_assessment_wallets"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_data_safe_crypto_assessment_wallets", "test_crypto_assessment_wallets", acctest.Required, acctest.Create, DataSafeCryptoAssessmentWalletDataSourceRepresentation) +
				compartmentIdVariableStr + cryptoAssessmentIdVariableStr + targetIdVariableStr + DataSafeCryptoAssessmentWalletResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(datasourceName, "crypto_assessment_wallet_collection.#"),
				resource.TestCheckResourceAttrSet(datasourceName, "crypto_assessment_wallet_collection.0.items.#"),
			),
		},
	})
}
