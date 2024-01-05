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

	"github.com/oracle/terraform-provider-oci/httpreplay"
)

var (
	unsetUserAssessmentBaselineRepresentation = map[string]interface{}{
		"user_assessment_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_data_safe_user_assessment.test_user_assessment.id}`},
	}

	DataSafeUnsetUserAssessmentBaselineResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_data_safe_user_assessment", "test_user_assessment", acctest.Required, acctest.Create, userAssessmentRepresentation)
)

// issue-routing-tag: data_safe/default
func TestDataSafeUnsetUserAssessmentBaselineResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDataSafeUnsetUserAssessmentBaselineResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	targetId := utils.GetEnvSettingWithBlankDefault("data_safe_target_ocid")
	targetIdVariableStr := fmt.Sprintf("variable \"target_id\" { default = \"%s\" }\n", targetId)

	// Save TF content to Create resource with only required properties. This has to be exactly the same as the config part in the Create step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+DataSafeUnsetUserAssessmentBaselineResourceDependencies+targetIdVariableStr+
		acctest.GenerateResourceFromRepresentationMap("oci_data_safe_unset_user_assessment_baseline", "test_unset_user_assessment_baseline", acctest.Required, acctest.Create, unsetUserAssessmentBaselineRepresentation), "datasafe", "unsetUserAssessmentBaseline", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + DataSafeUnsetUserAssessmentBaselineResourceDependencies + targetIdVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_data_safe_unset_user_assessment_baseline", "test_unset_user_assessment_baseline", acctest.Required, acctest.Create, unsetUserAssessmentBaselineRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				func(s *terraform.State) (err error) {
					return nil
				},
			),
		},
	})
}
