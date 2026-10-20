// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"

	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	DataSafeCryptoAssessmentKeyDataSourceRepresentation = map[string]interface{}{
		"compartment_id":            acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"access_level":              acctest.Representation{RepType: acctest.Required, Create: `RESTRICTED`},
		"assessment_id":             acctest.Representation{RepType: acctest.Required, Create: `${var.crypto_assessment_id}`},
		"assessment_type":           acctest.Representation{RepType: acctest.Required, Create: `LATEST`},
		"compartment_id_in_subtree": acctest.Representation{RepType: acctest.Required, Create: `false`},
		"feature":                   acctest.Representation{RepType: acctest.Required, Create: `TDE`},
	}

	DataSafeCryptoAssessmentKeyResourceConfig = ""
)

// issue-routing-tag: data_safe/default
func TestDataSafeCryptoAssessmentKeyResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDataSafeCryptoAssessmentKeyResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)
	cryptoAssessmentId := utils.GetEnvSettingWithBlankDefault("crypto_assessment_id")
	cryptoAssessmentIdVariableStr := fmt.Sprintf("variable \"crypto_assessment_id\" { default = \"%s\" }\n", cryptoAssessmentId)

	datasourceName := "data.oci_data_safe_crypto_assessment_keys.test_crypto_assessment_keys"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_data_safe_crypto_assessment_keys", "test_crypto_assessment_keys", acctest.Required, acctest.Create, DataSafeCryptoAssessmentKeyDataSourceRepresentation) +
				compartmentIdVariableStr + cryptoAssessmentIdVariableStr + DataSafeCryptoAssessmentKeyResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "access_level", "RESTRICTED"),
				resource.TestCheckResourceAttrSet(datasourceName, "assessment_id"),
				resource.TestCheckResourceAttr(datasourceName, "assessment_type", "LATEST"),
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "compartment_id_in_subtree", "false"),
				resource.TestCheckResourceAttr(datasourceName, "feature", "TDE"),
				resource.TestCheckResourceAttrSet(datasourceName, "crypto_assessment_key_collection.#"),
				testCheckCryptoAssessmentKeyDetailsIfPresent(datasourceName),
			),
		},
	})
}

func testCheckCryptoAssessmentKeyDetailsIfPresent(datasourceName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[datasourceName]
		if !ok {
			return fmt.Errorf("not found: %s", datasourceName)
		}

		itemsCount, err := strconv.Atoi(rs.Primary.Attributes["crypto_assessment_key_collection.0.items.#"])
		if err != nil {
			return err
		}

		for i := 0; i < itemsCount; i++ {
			prefix := fmt.Sprintf("crypto_assessment_key_collection.0.items.%d.", i)
			for _, field := range []string{"assessment_id", "target_id", "key_id", "feature", "algorithm", "time_created"} {
				if rs.Primary.Attributes[prefix+field] == "" {
					return fmt.Errorf("expected %s%s to be set", prefix, field)
				}
			}
		}

		return nil
	}
}
