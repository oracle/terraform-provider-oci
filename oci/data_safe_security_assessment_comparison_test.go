// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	oci_common "github.com/oracle/oci-go-sdk/v48/common"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	securityAssessmentComparisonSingularDataSourceRepresentation = map[string]interface{}{
		"comparison_security_assessment_id": Representation{repType: Required, create: `${oci_data_safe_security_assessment.test_security_assessment3.id}`},
		"security_assessment_id":            Representation{repType: Required, create: `${oci_data_safe_security_assessment.test_security_assessment4.id}`},
	}

	SecurityAssessmentComparisonResourceConfig = generateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", Required, Create, autonomousDatabaseRepresentation) +
		generateResourceFromRepresentationMap("oci_data_safe_target_database", "test_target_database", Required, Create, targetDatabaseRepresentation) +
		generateResourceFromRepresentationMap("oci_data_safe_security_assessment", "test_security_assessment3", Required, Create, securityAssessmentRepresentation) +
		generateResourceFromRepresentationMap("oci_data_safe_security_assessment", "test_security_assessment4", Required, Create, securityAssessmentRepresentation)
)

// issue-routing-tag: data_safe/default
func TestDataSafeSecurityAssessmentComparisonResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDataSafeSecurityAssessmentComparisonResource_basic")
	defer httpreplay.SaveScenario()

	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	saveConfigContent("", "", "", t)

	ResourceTest(t, nil, []resource.TestStep{
		// verify singular datasource
		{
			Config: config + SecurityAssessmentComparisonResourceConfig + compartmentIdVariableStr +
				generateDataSourceFromRepresentationMap("oci_data_safe_security_assessment_comparison", "test_security_assessment_comparison", Required, Create, securityAssessmentComparisonSingularDataSourceRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				func(s *terraform.State) (err error) {
					if failure, isServiceError := oci_common.IsServiceError(err); !isServiceError || failure.GetHTTPStatusCode() != 404 {
						return err
					}
					return nil
				},
			),
		},
	})
}
