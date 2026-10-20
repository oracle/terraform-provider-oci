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
	DataSafeCryptoAssessmentFindingDataSourceRepresentation = map[string]interface{}{
		"crypto_assessment_id":       acctest.Representation{RepType: acctest.Required, Create: `${var.crypto_assessment_id}`},
		"category":                   acctest.Representation{RepType: acctest.Required, Create: `CERTIFICATES_AND_KEY_MANAGEMENT`},
		"finding_key":                acctest.Representation{RepType: acctest.Required, Create: `CONF.ALLOWED_WEAK_CERT_ALGORITHMS`},
		"is_quantum_readiness_check": acctest.Representation{RepType: acctest.Required, Create: `false`},
		"status":                     acctest.Representation{RepType: acctest.Optional, Create: `PASS`},
		"title":                      acctest.Representation{RepType: acctest.Optional, Create: `title`},
	}
)

// issue-routing-tag: data_safe/default
func TestDataSafeCryptoAssessmentFindingResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDataSafeCryptoAssessmentFindingResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)
	cryptoAssessmentId := utils.GetEnvSettingWithBlankDefault("crypto_assessment_id")
	cryptoAssessmentIdVariableStr := fmt.Sprintf("variable \"crypto_assessment_id\" { default = \"%s\" }\n", cryptoAssessmentId)

	datasourceName := "data.oci_data_safe_crypto_assessment_findings.test_crypto_assessment_findings"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_data_safe_crypto_assessment_findings", "test_crypto_assessment_findings", acctest.Required, acctest.Create, DataSafeCryptoAssessmentFindingDataSourceRepresentation) +
				compartmentIdVariableStr + cryptoAssessmentIdVariableStr,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "crypto_assessment_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "crypto_assessment_finding_collection.#"),
				resource.TestCheckResourceAttr(datasourceName, "crypto_assessment_finding_collection.0.assessment_type", "LATEST"),
				resource.TestCheckResourceAttrSet(datasourceName, "crypto_assessment_finding_collection.0.database_version"),
				testCheckCryptoAssessmentFindingDetailsIfPresent(datasourceName),
			),
		},
	})
}

func testCheckCryptoAssessmentFindingDetailsIfPresent(datasourceName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[datasourceName]
		if !ok {
			return fmt.Errorf("not found: %s", datasourceName)
		}

		itemsCountStr := rs.Primary.Attributes["crypto_assessment_finding_collection.0.items.#"]
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

		if rs.Primary.Attributes["crypto_assessment_finding_collection.0.items.0.category"] == "" {
			return fmt.Errorf("expected category to be set")
		}
		if rs.Primary.Attributes["crypto_assessment_finding_collection.0.items.0.finding_key"] == "" {
			return fmt.Errorf("expected finding_key to be set")
		}
		if rs.Primary.Attributes["crypto_assessment_finding_collection.0.items.0.is_quantum_readiness_check"] == "" {
			return fmt.Errorf("expected is_quantum_readiness_check to be set")
		}
		if rs.Primary.Attributes["crypto_assessment_finding_collection.0.items.0.status"] == "" {
			return fmt.Errorf("expected status to be set")
		}
		if rs.Primary.Attributes["crypto_assessment_finding_collection.0.items.0.title"] == "" {
			return fmt.Errorf("expected title to be set")
		}
		if rs.Primary.Attributes["crypto_assessment_finding_collection.0.items.0.compliance"] == "" {
			return fmt.Errorf("expected compliance to be set")
		}
		if rs.Primary.Attributes["crypto_assessment_finding_collection.0.items.0.expected_value"] == "" {
			return fmt.Errorf("expected expected_value to be set")
		}
		if rs.Primary.Attributes["crypto_assessment_finding_collection.0.items.0.recommended_value"] == "" {
			return fmt.Errorf("expected recommended_value to be set")
		}
		if rs.Primary.Attributes["crypto_assessment_finding_collection.0.items.0.remediation"] == "" {
			return fmt.Errorf("expected remediation to be set")
		}
		if rs.Primary.Attributes["crypto_assessment_finding_collection.0.items.0.summary"] == "" {
			return fmt.Errorf("expected summary to be set")
		}
		if rs.Primary.Attributes["crypto_assessment_finding_collection.0.items.0.url"] == "" {
			return fmt.Errorf("expected url to be set")
		}
		if rs.Primary.Attributes["crypto_assessment_finding_collection.0.target_id"] == "" {
			return fmt.Errorf("expected target_id to be set")
		}

		return nil
	}
}
