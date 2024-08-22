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
	unsetUserAssessmentBaselineManagementRepresentationDependencies = acctest.GenerateResourceFromRepresentationMap("oci_data_safe_set_user_assessment_baseline_management", "test_set_user_assessment_baseline_management", acctest.Optional, acctest.Create, setUserAssessmentBaselineManagementRepresentation)
	unsetUserAssessmentBaselineManagementRepresentation             = map[string]interface{}{
		"user_assessment_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_data_safe_set_user_assessment_baseline_management.test_set_user_assessment_baseline_management.user_assessment_id}`},
		"compartment_id":     acctest.Representation{RepType: acctest.Required, Create: "ocid1.tenancy.oc1..aaaaaaaaiu6b5x762v2un36ubxhzhoqyd5a6tfu4ny3su6g5xjtvnlomlyrq"},
	}
)

// issue-routing-tag: data_safe/default
func TestDataSafeUnsetUserAssessmentBaselineManagementResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDataSafeUnsetUserAssessmentBaselineManagementResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	targetId := utils.GetEnvSettingWithBlankDefault("data_safe_target_ocid")
	targetIdVariableStr := fmt.Sprintf("variable \"target_id\" { default = \"%s\" }\n", targetId)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// Create
		{
			Config: config + compartmentIdVariableStr + targetIdVariableStr + unsetUserAssessmentBaselineManagementRepresentationDependencies + acctest.GenerateResourceFromRepresentationMap("oci_data_safe_unset_user_assessment_baseline_management", "test_unset_user_assessment_baseline_management", acctest.Optional, acctest.Create, unsetUserAssessmentBaselineManagementRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				func(s *terraform.State) (err error) {
					return nil
				},
			),
		},
		// Delete
		{

			Config: config + compartmentIdVariableStr + targetIdVariableStr,
		},
	})
}
