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
	DataSafeCryptoAssessmentSqlnetParameterSingularDataSourceRepresentation = map[string]interface{}{
		"crypto_assessment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.crypto_assessment_id}`},
		"parameter":            acctest.Representation{RepType: acctest.Optional, Create: `parameter`},
		"quantum_readiness":    acctest.Representation{RepType: acctest.Optional, Create: `RESISTANT`},
	}
)

// issue-routing-tag: data_safe/default
func TestDataSafeCryptoAssessmentSqlnetParameterResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDataSafeCryptoAssessmentSqlnetParameterResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)
	cryptoAssessmentId := utils.GetEnvSettingWithBlankDefault("crypto_assessment_id")
	cryptoAssessmentIdVariableStr := fmt.Sprintf("variable \"crypto_assessment_id\" { default = \"%s\" }\n", cryptoAssessmentId)

	singularDatasourceName := "data.oci_data_safe_crypto_assessment_sqlnet_parameter.test_crypto_assessment_sqlnet_parameter"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_data_safe_crypto_assessment_sqlnet_parameter", "test_crypto_assessment_sqlnet_parameter", acctest.Required, acctest.Create, DataSafeCryptoAssessmentSqlnetParameterSingularDataSourceRepresentation) +
				compartmentIdVariableStr + cryptoAssessmentIdVariableStr,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "crypto_assessment_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "parameters.#"),
				resource.TestCheckResourceAttr(singularDatasourceName, "source.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "source.0.type", "SQLNET_ORA"),
				testCheckSqlnetParameterDetailsIfPresent(singularDatasourceName),
			),
		},
	})
}

func testCheckSqlnetParameterDetailsIfPresent(datasourceName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[datasourceName]
		if !ok {
			return fmt.Errorf("not found: %s", datasourceName)
		}

		parametersCount, err := strconv.Atoi(rs.Primary.Attributes["parameters.#"])
		if err != nil {
			return err
		}

		for i := 0; i < parametersCount; i++ {
			prefix := fmt.Sprintf("parameters.%d.", i)
			for _, field := range []string{"name", "quantum_readiness", "value.#"} {
				if rs.Primary.Attributes[prefix+field] == "" {
					return fmt.Errorf("expected %s%s to be set", prefix, field)
				}
			}
		}

		return nil
	}
}
