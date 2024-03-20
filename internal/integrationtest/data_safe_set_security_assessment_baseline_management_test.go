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
	setSecurityAssessmentBaselineManagementRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"target_id":      acctest.Representation{RepType: acctest.Required, Create: `${var.target_id}`},
	}
	unsetSecurityAssessmentBaselineRepresentation = map[string]interface{}{
		"security_assessment_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_data_safe_set_security_assessment_baseline_management.test_set_security_assessment_baseline_management.security_assessment_id}`},
	}
)

// issue-routing-tag: data_safe/default
func TestDataSafeSetSecurityAssessmentBaselineManagementResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDataSafeSetSecurityAssessmentBaselineManagementResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	targetId := utils.GetEnvSettingWithBlankDefault("data_safe_target_ocid")
	targetIdVariableStr := fmt.Sprintf("variable \"target_id\" { default = \"%s\" }\n", targetId)

	resourceName := "oci_data_safe_set_security_assessment_baseline_management.test_set_security_assessment_baseline_management"

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// Create
		{
			Config: config + compartmentIdVariableStr + targetIdVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_data_safe_set_security_assessment_baseline_management", "test_set_security_assessment_baseline_management", acctest.Optional, acctest.Create, setSecurityAssessmentBaselineManagementRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "security_assessment_id"),
			),
		},
		// Unset baseline
		{
			Config: config + compartmentIdVariableStr + targetIdVariableStr + acctest.GenerateResourceFromRepresentationMap("oci_data_safe_set_security_assessment_baseline_management", "test_set_security_assessment_baseline_management", acctest.Optional, acctest.Create, setSecurityAssessmentBaselineManagementRepresentation) + acctest.GenerateResourceFromRepresentationMap("oci_data_safe_unset_security_assessment_baseline", "test_unset_security_assessment_baselinet", acctest.Optional, acctest.Create, unsetSecurityAssessmentBaselineRepresentation),
		},

		// Delete
		{

			Config: config + compartmentIdVariableStr + targetIdVariableStr,
		},
	})
}
