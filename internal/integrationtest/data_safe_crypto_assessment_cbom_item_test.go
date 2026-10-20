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
	DataSafeCryptoAssessmentCbomItemDataSourceRepresentation = map[string]interface{}{
		"crypto_assessment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.crypto_assessment_id}`},
	}
)

// issue-routing-tag: data_safe/default
func TestDataSafeCryptoAssessmentCbomItemResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDataSafeCryptoAssessmentCbomItemResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	cryptoAssessmentId := utils.GetEnvSettingWithBlankDefault("crypto_assessment_id")
	cryptoAssessmentIdVariableStr := fmt.Sprintf("variable \"crypto_assessment_id\" { default = \"%s\" }\n", cryptoAssessmentId)

	datasourceName := "data.oci_data_safe_crypto_assessment_cbom_items.test_crypto_assessment_cbom_items"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_data_safe_crypto_assessment_cbom_items", "test_crypto_assessment_cbom_items", acctest.Required, acctest.Create, DataSafeCryptoAssessmentCbomItemDataSourceRepresentation) +
				cryptoAssessmentIdVariableStr,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "crypto_assessment_id"),

				resource.TestCheckResourceAttrSet(datasourceName, "crypto_assessment_cbom_item_collection.#"),
			),
		},
	})
}
