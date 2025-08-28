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
	DataSafeSecurityAssessmentCheckDataSourceRepresentation = map[string]interface{}{
		"security_assessment_id":    acctest.Representation{RepType: acctest.Required, Create: `${var.security_assessment_id}`},
		"access_level":              acctest.Representation{RepType: acctest.Optional, Create: `ACCESSIBLE`},
		"compartment_id_in_subtree": acctest.Representation{RepType: acctest.Optional, Create: `true`},
		"key":                       acctest.Representation{RepType: acctest.Optional, Create: `${var.key}`},
	}

	DataSafeSecurityAssessmentCheckRepresentation = map[string]interface{}{
		"security_assessment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.security_assessment_id}`},
	}
)

// issue-routing-tag: data_safe/default
func TestDataSafeSecurityAssessmentCheckResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDataSafeSecurityAssessmentCheckResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	assessmentId := utils.GetEnvSettingWithBlankDefault("data_safe_security_assessment_id")
	securityAssessmentIdVariableStr := fmt.Sprintf("variable \"security_assessment_id\" { default = \"%s\" }\n", assessmentId)

	findingKey := utils.GetEnvSettingWithBlankDefault("data_safe_key")
	findingKeyVariableStr := fmt.Sprintf("variable \"key\" { default = \"%s\" }\n", findingKey)

	datasourceName := "data.oci_data_safe_security_assessment_checks.test_security_assessment_checks"

	// Save TF content to Create resource with only required properties. This has to be exactly the same as the config part in the create step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+
		acctest.GenerateResourceFromRepresentationMap("oci_data_safe_security_assessment_check", "test_security_assessment_check", acctest.Required, acctest.Create, DataSafeSecurityAssessmentCheckRepresentation), "datasafe", "securityAssessmentCheck", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_data_safe_security_assessment_checks", "test_security_assessment_checks", acctest.Optional, acctest.Update, DataSafeSecurityAssessmentCheckDataSourceRepresentation) +
				compartmentIdVariableStr + securityAssessmentIdVariableStr + findingKeyVariableStr,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "access_level", "ACCESSIBLE"),
				resource.TestCheckResourceAttr(datasourceName, "compartment_id_in_subtree", "true"),
				resource.TestCheckResourceAttr(datasourceName, "key", findingKey),

				resource.TestCheckResourceAttr(datasourceName, "checks.#", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "checks.0.category"),
				resource.TestCheckResourceAttrSet(datasourceName, "checks.0.key"),
				resource.TestCheckResourceAttrSet(datasourceName, "checks.0.oneline"),
				resource.TestCheckResourceAttr(datasourceName, "checks.0.references.#", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "checks.0.remarks"),
				resource.TestCheckResourceAttrSet(datasourceName, "checks.0.suggested_severity"),
				resource.TestCheckResourceAttrSet(datasourceName, "checks.0.title"),
			),
		},
	})
}
