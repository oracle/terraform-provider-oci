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
	DataSafeCryptoAssessmentFindingTargetDataSourceRepresentation = map[string]interface{}{
		"compartment_id":             acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"finding_key":                acctest.Representation{RepType: acctest.Required, Create: []string{`CONF.ALLOWED_WEAK_CERT_ALGORITHMS`}},
		"access_level":               acctest.Representation{RepType: acctest.Required, Create: `ACCESSIBLE`},
		"assessment_type":            acctest.Representation{RepType: acctest.Required, Create: `LATEST`},
		"compartment_id_in_subtree":  acctest.Representation{RepType: acctest.Required, Create: `true`},
		"is_quantum_readiness_check": acctest.Representation{RepType: acctest.Required, Create: `false`},
		"target_id":                  acctest.Representation{RepType: acctest.Required, Create: `${var.target_id}`},
	}
)

// issue-routing-tag: data_safe/default
func TestDataSafeCryptoAssessmentFindingTargetResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDataSafeCryptoAssessmentFindingTargetResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)
	targetId := utils.GetEnvSettingWithBlankDefault("target_ocid")
	targetIdVariableStr := fmt.Sprintf("variable \"target_id\" { default = \"%s\" }\n", targetId)

	datasourceName := "data.oci_data_safe_crypto_assessment_finding_targets.test_crypto_assessment_finding_targets"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_data_safe_crypto_assessment_finding_targets", "test_crypto_assessment_finding_targets", acctest.Required, acctest.Create, DataSafeCryptoAssessmentFindingTargetDataSourceRepresentation) +
				compartmentIdVariableStr + targetIdVariableStr,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "access_level", "ACCESSIBLE"),
				resource.TestCheckResourceAttr(datasourceName, "assessment_type", "LATEST"),
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "compartment_id_in_subtree", "true"),
				resource.TestCheckResourceAttr(datasourceName, "finding_key.0", "CONF.ALLOWED_WEAK_CERT_ALGORITHMS"),
				resource.TestCheckResourceAttr(datasourceName, "is_quantum_readiness_check", "false"),
				resource.TestCheckResourceAttr(datasourceName, "target_id", targetId),
				resource.TestCheckResourceAttrSet(datasourceName, "crypto_assessment_finding_target_collection.#"),
				testCheckFindingTargetDetailsIfPresent(datasourceName),
			),
		},
	})
}

func testCheckFindingTargetDetailsIfPresent(datasourceName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[datasourceName]
		if !ok {
			return fmt.Errorf("not found: %s", datasourceName)
		}

		itemsCount, err := strconv.Atoi(rs.Primary.Attributes["crypto_assessment_finding_target_collection.0.items.#"])
		if err != nil {
			return err
		}

		for i := 0; i < itemsCount; i++ {
			prefix := fmt.Sprintf("crypto_assessment_finding_target_collection.0.items.%d.", i)
			for _, field := range []string{"assessment_id", "database_version", "finding_key", "observed_value", "target_id"} {
				if rs.Primary.Attributes[prefix+field] == "" {
					return fmt.Errorf("expected %s%s to be set", prefix, field)
				}
			}
		}

		return nil
	}
}
