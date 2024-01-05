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
	DataSafeuserAssessmentComparisonSingularDataSourceRepresentation = map[string]interface{}{
		"comparison_user_assessment_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_data_safe_user_assessment.test_user_assessment3.id}`},
		"user_assessment_id":            acctest.Representation{RepType: acctest.Required, Create: `${oci_data_safe_user_assessment.test_user_assessment4.id}`},
	}

	DataSafeUserAssessmentComparisonResourceConfig = acctest.GenerateResourceFromRepresentationMap("oci_data_safe_user_assessment", "test_user_assessment3", acctest.Required, acctest.Create, userAssessmentRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_data_safe_user_assessment", "test_user_assessment4", acctest.Required, acctest.Create, userAssessmentRepresentation)
)

// issue-routing-tag: data_safe/default
func TestDataSafeUserAssessmentComparisonResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDataSafeUserAssessmentComparisonResource_basic")
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
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_data_safe_user_assessment_comparison", "test_user_assessment_comparison", acctest.Required, acctest.Create, DataSafeuserAssessmentComparisonSingularDataSourceRepresentation) +
				compartmentIdVariableStr + DataSafeUserAssessmentComparisonResourceConfig + targetIdVariableStr,
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
