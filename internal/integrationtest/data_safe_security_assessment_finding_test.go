// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"testing"

	"github.com/terraform-providers/terraform-provider-oci/internal/acctest"
	"github.com/terraform-providers/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	securityAssessmentFindingDataSourceRepresentation = map[string]interface{}{
		"security_assessment_id":    acctest.Representation{RepType: acctest.Required, Create: `${oci_data_safe_security_assessment.test_security_assessment_findings.id}`},
		"access_level":              acctest.Representation{RepType: acctest.Optional, Create: `ACCESSIBLE`},
		"compartment_id_in_subtree": acctest.Representation{RepType: acctest.Optional, Create: `true`},
		"finding_key":               acctest.Representation{RepType: acctest.Optional, Create: `findingKey`},
		"severity":                  acctest.Representation{RepType: acctest.Optional, Create: `HIGH`},
	}

	SecurityAssessmentFindingResourceConfig = acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", acctest.Required, acctest.Create, autonomousDatabaseRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_data_safe_target_database", "test_target_database", acctest.Required, acctest.Create, targetDatabaseRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_data_safe_security_assessment", "test_security_assessment_findings", acctest.Required, acctest.Create, securityAssessmentRepresentation)
)

// issue-routing-tag: data_safe/default
func TestDataSafeSecurityAssessmentFindingResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDataSafeSecurityAssessmentFindingResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_data_safe_security_assessment_findings.test_security_assessment_findings"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_data_safe_security_assessment_findings", "test_security_assessment_findings", acctest.Required, acctest.Create, securityAssessmentFindingDataSourceRepresentation) +
				compartmentIdVariableStr + SecurityAssessmentFindingResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "findings.#"),
				resource.TestCheckResourceAttrSet(datasourceName, "findings.1.key"),
				resource.TestCheckResourceAttr(datasourceName, "findings.1.references.#", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "findings.1.remarks"),
				resource.TestCheckResourceAttrSet(datasourceName, "findings.1.severity"),
				resource.TestCheckResourceAttrSet(datasourceName, "findings.1.summary"),
				resource.TestCheckResourceAttrSet(datasourceName, "findings.1.target_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "findings.1.title"),
			),
		},
	})
}
