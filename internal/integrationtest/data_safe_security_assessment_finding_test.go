// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"testing"

	"github.com/oracle/terraform-provider-oci/internal/acctest"
	"github.com/oracle/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"github.com/oracle/terraform-provider-oci/httpreplay"
)

var (
	DataSafesecurityAssessmentFindingDataSourceRepresentation = map[string]interface{}{
		"compartment_id":            acctest.Representation{RepType: acctest.Optional, Create: `${var.compartment_id}`},
		"security_assessment_id":    acctest.Representation{RepType: acctest.Required, Create: `${var.security_assessment_id}`},
		"access_level":              acctest.Representation{RepType: acctest.Optional, Create: `ACCESSIBLE`},
		"compartment_id_in_subtree": acctest.Representation{RepType: acctest.Optional, Create: `true`},
		"finding_key":               acctest.Representation{RepType: acctest.Optional, Create: `findingKey`},
		"severity":                  acctest.Representation{RepType: acctest.Optional, Create: `HIGH`},
		"state":                     acctest.Representation{RepType: acctest.Optional, Create: `AVAILABLE`},
		"target_id":                 acctest.Representation{RepType: acctest.Optional, Create: `${var.target_id}`},
		"is_top_finding":            acctest.Representation{RepType: acctest.Optional, Create: `false`},
	}
	DataSafeSecurityAssessmentFindingReferencesDataSourceRepresentation = DefinedTagsDependencies
)

// issue-routing-tag: data_safe/default
func TestDataSafeSecurityAssessmentFindingResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDataSafeSecurityAssessmentFindingResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	targetId := utils.GetEnvSettingWithBlankDefault("data_safe_target_ocid")
	targetIdVariableStr := fmt.Sprintf("variable \"target_id\" { default = \"%s\" }\n", targetId)

	assessmentId := utils.GetEnvSettingWithBlankDefault("data_safe_security_assessment_id")
	securityAssessmentIdVariableStr := fmt.Sprintf("variable \"security_assessment_id\" { default = \"%s\" }\n", assessmentId)
	datasourceName := "data.oci_data_safe_security_assessment_findings.test_security_assessment_findings"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_data_safe_security_assessment_findings", "test_security_assessment_findings", acctest.Required, acctest.Create, DataSafesecurityAssessmentFindingDataSourceRepresentation) +
				compartmentIdVariableStr + targetIdVariableStr + securityAssessmentIdVariableStr,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "security_assessment_id"),
				resource.TestCheckResourceAttr(datasourceName, "severity", "HIGH"),
				resource.TestCheckResourceAttrSet(datasourceName, "target_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "findings.#"),
				resource.TestCheckResourceAttrSet(datasourceName, "findings.0.has_target_db_risk_level_changed"),
				resource.TestCheckResourceAttrSet(datasourceName, "findings.0.is_risk_modified"),
				resource.TestCheckResourceAttrSet(datasourceName, "findings.0.key"),
				//resource.TestCheckResourceAttrSet(datasourceName, "findings.0.oneline"),
				resource.TestCheckResourceAttrSet(datasourceName, "findings.0.oracle_defined_severity"),
				resource.TestCheckResourceAttr(datasourceName, "findings.0.references.#", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "findings.0.is_top_finding"),
				resource.TestCheckResourceAttrSet(datasourceName, "findings.0.key"),
				resource.TestCheckResourceAttrSet(datasourceName, "findings.0.remarks"),
				resource.TestCheckResourceAttrSet(datasourceName, "findings.0.severity"),
				resource.TestCheckResourceAttrSet(datasourceName, "findings.0.summary"),
				resource.TestCheckResourceAttrSet(datasourceName, "findings.0.target_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "findings.0.title"),
			),
		},
	})
}
