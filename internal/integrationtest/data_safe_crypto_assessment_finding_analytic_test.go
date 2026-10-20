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
	DataSafeCryptoAssessmentFindingAnalyticDataSourceRepresentation = map[string]interface{}{
		"compartment_id":             acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"access_level":               acctest.Representation{RepType: acctest.Optional, Create: `ACCESSIBLE`},
		"category":                   acctest.Representation{RepType: acctest.Optional, Create: `NETWORK_ENCRYPTION`},
		"compartment_id_in_subtree":  acctest.Representation{RepType: acctest.Optional, Create: `true`},
		"finding_key":                acctest.Representation{RepType: acctest.Optional, Create: `findingKey`},
		"is_quantum_readiness_check": acctest.Representation{RepType: acctest.Optional, Create: `false`},
	}

	DataSafeCryptoAssessmentFindingAnalyticResourceConfig = ""
)

// issue-routing-tag: data_safe/default
func TestDataSafeCryptoAssessmentFindingAnalyticResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDataSafeCryptoAssessmentFindingAnalyticResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_data_safe_crypto_assessment_finding_analytics.test_crypto_assessment_finding_analytics"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_data_safe_crypto_assessment_finding_analytics", "test_crypto_assessment_finding_analytics", acctest.Required, acctest.Create, DataSafeCryptoAssessmentFindingAnalyticDataSourceRepresentation) +
				compartmentIdVariableStr + DataSafeCryptoAssessmentFindingAnalyticResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "crypto_assessment_finding_analytics_collection.#"),
				testCheckCryptoAssessmentFindingAnalyticsDetailsIfPresent(datasourceName),
			),
		},
	})
}

func testCheckCryptoAssessmentFindingAnalyticsDetailsIfPresent(datasourceName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[datasourceName]
		if !ok {
			return fmt.Errorf("not found: %s", datasourceName)
		}

		itemsCountStr := rs.Primary.Attributes["crypto_assessment_finding_analytics_collection.0.items.#"]
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

		if rs.Primary.Attributes["crypto_assessment_finding_analytics_collection.0.items.0.category"] == "" {
			return fmt.Errorf("expected category to be set")
		}
		if rs.Primary.Attributes["crypto_assessment_finding_analytics_collection.0.items.0.finding_key"] == "" {
			return fmt.Errorf("expected finding_key to be set")
		}
		if rs.Primary.Attributes["crypto_assessment_finding_analytics_collection.0.items.0.target_count"] == "" {
			return fmt.Errorf("expected target_count to be set")
		}
		if rs.Primary.Attributes["crypto_assessment_finding_analytics_collection.0.items.0.title"] == "" {
			return fmt.Errorf("expected title to be set")
		}

		return nil
	}
}
