// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"testing"

	"github.com/oracle/terraform-provider-oci/internal/acctest"
	"github.com/oracle/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	oci_common "github.com/oracle/oci-go-sdk/v65/common"

	"github.com/oracle/terraform-provider-oci/httpreplay"
)

var (
	DataSafesecurityAssessmentComparisonSingularDataSourceRepresentation = map[string]interface{}{
		"comparison_security_assessment_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_data_safe_security_assessment.test_security_assessment3.id}`},
		"security_assessment_id":            acctest.Representation{RepType: acctest.Required, Create: `${oci_data_safe_security_assessment.test_security_assessment4.id}`},
	}

	DataSafeSecurityAssessmentComparisonResourceConfig = acctest.GenerateResourceFromRepresentationMap("oci_data_safe_security_assessment", "test_security_assessment3", acctest.Required, acctest.Create, securityAssessmentRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_data_safe_security_assessment", "test_security_assessment4", acctest.Required, acctest.Create, securityAssessmentRepresentation)
)

// issue-routing-tag: data_safe/default
func TestDataSafeSecurityAssessmentComparisonResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDataSafeSecurityAssessmentComparisonResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	targetId := utils.GetEnvSettingWithBlankDefault("data_safe_target_ocid")
	targetIdVariableStr := fmt.Sprintf("variable \"target_id\" { default = \"%s\" }\n", targetId)

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify singular datasource
		{
			Config: config + DataSafeSecurityAssessmentComparisonResourceConfig + compartmentIdVariableStr + targetIdVariableStr +
				acctest.GenerateDataSourceFromRepresentationMap("oci_data_safe_security_assessment_comparison", "test_security_assessment_comparison", acctest.Required, acctest.Create, DataSafesecurityAssessmentComparisonSingularDataSourceRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
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
