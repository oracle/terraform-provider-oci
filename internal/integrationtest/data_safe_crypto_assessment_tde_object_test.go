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
	DataSafeCryptoAssessmentTdeObjectDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"object_type":    acctest.Representation{RepType: acctest.Required, Create: `TABLESPACE`},
		"assessment_id":  acctest.Representation{RepType: acctest.Optional, Create: `${var.crypto_assessment_id}`},
	}

	DataSafeCryptoAssessmentTdeObjectResourceConfig = ""
)

func testCheckTdeObjectDetailsIfPresent(datasourceName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[datasourceName]
		if !ok {
			return fmt.Errorf("not found: %s", datasourceName)
		}

		itemsCountStr := rs.Primary.Attributes["crypto_assessment_tde_object_collection.0.items.#"]
		if itemsCountStr == "" {
			return fmt.Errorf("missing item count for %s", datasourceName)
		}

		itemsCount, err := strconv.Atoi(itemsCountStr)
		if err != nil {
			return err
		}

		if itemsCount == 0 {
			return nil
		}

		if rs.Primary.Attributes["crypto_assessment_tde_object_collection.0.items.0.assessment_id"] == "" {
			return fmt.Errorf("expected assessment_id to be set")
		}
		if rs.Primary.Attributes["crypto_assessment_tde_object_collection.0.items.0.tablespace_name"] == "" {
			return fmt.Errorf("expected tablespace_name to be set")
		}
		if rs.Primary.Attributes["crypto_assessment_tde_object_collection.0.items.0.target_id"] == "" {
			return fmt.Errorf("expected target_id to be set")
		}
		if rs.Primary.Attributes["crypto_assessment_tde_object_collection.0.items.0.time_last_assessed"] == "" {
			return fmt.Errorf("expected time_last_assessed to be set")
		}

		return nil
	}
}

// issue-routing-tag: data_safe/default
func TestDataSafeCryptoAssessmentTdeObjectResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDataSafeCryptoAssessmentTdeObjectResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)
	cryptoAssessmentId := utils.GetEnvSettingWithBlankDefault("crypto_assessment_id")
	cryptoAssessmentIdVariableStr := fmt.Sprintf("variable \"crypto_assessment_id\" { default = \"%s\" }\n", cryptoAssessmentId)

	datasourceName := "data.oci_data_safe_crypto_assessment_tde_objects.test_crypto_assessment_tde_objects"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_data_safe_crypto_assessment_tde_objects", "test_crypto_assessment_tde_objects", acctest.Required, acctest.Create, DataSafeCryptoAssessmentTdeObjectDataSourceRepresentation) +
				compartmentIdVariableStr + cryptoAssessmentIdVariableStr + DataSafeCryptoAssessmentTdeObjectResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "object_type", "TABLESPACE"),

				resource.TestCheckResourceAttrSet(datasourceName, "crypto_assessment_tde_object_collection.#"),
				testCheckTdeObjectDetailsIfPresent(datasourceName),
			),
		},
	})
}
