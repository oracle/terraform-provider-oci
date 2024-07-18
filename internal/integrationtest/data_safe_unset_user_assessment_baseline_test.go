// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/oracle/terraform-provider-oci/internal/acctest"
	"github.com/oracle/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/oracle/terraform-provider-oci/httpreplay"
)

var (
	DataSafeUnsetUserAssessmentBaselineRepresentation = map[string]interface{}{
		"user_assessment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.user_assessment_id}`},
		"target_ids":         acctest.Representation{RepType: acctest.Optional, Create: []string{`${var.target_id}`}},
	}

	DataSafeUnsetUserAssessmentBaselineResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_data_safe_user_assessment", "test_user_assessment", acctest.Required, acctest.Create, DataSafeUnsetUserAssessmentBaselineRepresentation)
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

	userAssessmentId := utils.GetEnvSettingWithBlankDefault("data_safe_user_assessment_id")
	userAssessmentIdVariableStr := fmt.Sprintf("variable \"user_assessment_id\" { default = \"%s\" }\n", userAssessmentId) //var resId string
	resourceName := "oci_data_safe_unset_user_assessment_baseline.test_unset_user_assessment_baseline"
	acctest.SaveConfigContent(config+compartmentIdVariableStr+DataSafeUnsetUserAssessmentBaselineResourceDependencies+targetIdVariableStr+
		acctest.GenerateResourceFromRepresentationMap("oci_data_safe_unset_user_assessment_baseline", "test_unset_user_assessment_baseline", acctest.Required, acctest.Create, DataSafeUnsetUserAssessmentBaselineRepresentation), "datasafe", "unsetUserAssessmentBaseline", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + targetIdVariableStr + userAssessmentIdVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_data_safe_unset_user_assessment_baseline", "test_unset_user_assessment_baseline", acctest.Optional, acctest.Create, DataSafeUnsetUserAssessmentBaselineRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "target_ids.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "user_assessment_id"),

				func(s *terraform.State) (err error) {
					return nil
				},
			),
		},
	})
}
