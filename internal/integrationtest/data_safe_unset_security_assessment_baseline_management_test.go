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
	unsetSecurityAssessmentBaselineManagementRepresentationDependencies = acctest.GenerateResourceFromRepresentationMap("oci_data_safe_set_security_assessment_baseline_management", "test_set_security_assessment_baseline_management", acctest.Optional, acctest.Create, setSecurityAssessmentBaselineManagementRepresentation)
	unsetSecurityAssessmentBaselineManagementRepresentation             = map[string]interface{}{
		"security_assessment_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_data_safe_set_security_assessment_baseline_management.test_set_security_assessment_baseline_management.security_assessment_id}`},
		"compartment_id":         acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
	}
)

// issue-routing-tag: data_safe/default
func TestDataSafeUnsetSecurityAssessmentBaselineManagementResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDataSafeUnsetSecurityAssessmentBaselineManagementResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	targetId := utils.GetEnvSettingWithBlankDefault("data_safe_target_ocid")
	targetIdVariableStr := fmt.Sprintf("variable \"target_id\" { default = \"%s\" }\n", targetId)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// Create
		{
			Config: config + compartmentIdVariableStr + targetIdVariableStr + unsetSecurityAssessmentBaselineManagementRepresentationDependencies + acctest.GenerateResourceFromRepresentationMap("oci_data_safe_unset_security_assessment_baseline_management", "test_unset_security_assessment_baseline_management", acctest.Optional, acctest.Create, unsetSecurityAssessmentBaselineManagementRepresentation),
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
