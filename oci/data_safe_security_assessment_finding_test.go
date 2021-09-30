// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	securityAssessmentFindingDataSourceRepresentation = map[string]interface{}{
		"security_assessment_id":    Representation{RepType: Required, Create: `${oci_data_safe_security_assessment.test_security_assessment_findings.id}`},
		"access_level":              Representation{RepType: Optional, Create: `ACCESSIBLE`},
		"compartment_id_in_subtree": Representation{RepType: Optional, Create: `true`},
		"finding_key":               Representation{RepType: Optional, Create: `findingKey`},
		"severity":                  Representation{RepType: Optional, Create: `HIGH`},
	}

	SecurityAssessmentFindingResourceConfig = GenerateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", Required, Create, autonomousDatabaseRepresentation) +
		GenerateResourceFromRepresentationMap("oci_data_safe_target_database", "test_target_database", Required, Create, targetDatabaseRepresentation) +
		GenerateResourceFromRepresentationMap("oci_data_safe_security_assessment", "test_security_assessment_findings", Required, Create, securityAssessmentRepresentation)
)

// issue-routing-tag: data_safe/default
func TestDataSafeSecurityAssessmentFindingResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDataSafeSecurityAssessmentFindingResource_basic")
	defer httpreplay.SaveScenario()

	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_data_safe_security_assessment_findings.test_security_assessment_findings"

	SaveConfigContent("", "", "", t)

	ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				GenerateDataSourceFromRepresentationMap("oci_data_safe_security_assessment_findings", "test_security_assessment_findings", Required, Create, securityAssessmentFindingDataSourceRepresentation) +
				compartmentIdVariableStr + SecurityAssessmentFindingResourceConfig,
			Check: ComposeAggregateTestCheckFuncWrapper(
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
