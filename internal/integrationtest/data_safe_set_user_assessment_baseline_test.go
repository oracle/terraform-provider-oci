// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/oracle/terraform-provider-oci/internal/acctest"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/oracle/terraform-provider-oci/httpreplay"
)

var (
	setUserAssessmentBaselineRepresentation = map[string]interface{}{
		"user_assessment_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_data_safe_user_assessment.test_user_assessment1.id}`},
		"assessment_ids":     acctest.Representation{RepType: acctest.Optional, Create: []string{`${oci_data_safe_user_assessment.test_user_assessment2.id}`}},
	}

	unsetAssessmentBaselineRepresentation = map[string]interface{}{
		"user_assessment_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_data_safe_user_assessment.test_user_assessment1.id}`},
	}

	DataSafeSetUserAssessmentBaselineResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_data_safe_user_assessment", "test_user_assessment1", acctest.Required, acctest.Create, userAssessmentRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_data_safe_user_assessment", "test_user_assessment2", acctest.Required, acctest.Create, userAssessmentRepresentation)
)

// issue-routing-tag: data_safe/default
func TestDataSafeSetUserAssessmentBaselineResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDataSafeSetUserAssessmentBaselineResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	targetId := utils.GetEnvSettingWithBlankDefault("data_safe_target_ocid")
	targetIdVariableStr := fmt.Sprintf("variable \"target_id\" { default = \"%s\" }\n", targetId)

	resourceName := "oci_data_safe_set_user_assessment_baseline.test_set_user_assessment_baseline"

	var resId string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+DataSafeSetUserAssessmentBaselineResourceDependencies+targetIdVariableStr+
		acctest.GenerateResourceFromRepresentationMap("oci_data_safe_set_user_assessment_baseline", "test_set_user_assessment_baseline", acctest.Optional, acctest.Create, setUserAssessmentBaselineRepresentation), "datasafe", "setUserAssessmentBaseline", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + DataSafeSetUserAssessmentBaselineResourceDependencies + targetIdVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_data_safe_set_user_assessment_baseline", "test_set_user_assessment_baseline", acctest.Optional, acctest.Create, setUserAssessmentBaselineRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "user_assessment_id"),
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + DataSafeSetUserAssessmentBaselineResourceDependencies + targetIdVariableStr,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + DataSafeSetUserAssessmentBaselineResourceDependencies + targetIdVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_data_safe_set_user_assessment_baseline", "test_set_user_assessment_baseline", acctest.Optional, acctest.Create, setUserAssessmentBaselineRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "assessment_ids.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "user_assessment_id"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					if isEnableExportCompartment, _ := strconv.ParseBool(utils.GetEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
						if errExport := resourcediscovery.TestExportCompartmentWithResourceName(&resId, &compartmentId, resourceName); errExport != nil {
							return errExport
						}
					}
					return err
				},
			),
		},
		{
			Config: config + compartmentIdVariableStr + DataSafeSetUserAssessmentBaselineResourceDependencies + targetIdVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_data_safe_unset_user_assessment_baseline", "test_unset_user_assessment_baseline", acctest.Required, acctest.Create, unsetAssessmentBaselineRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				// resource.TestCheckResourceAttrSet(resourceNameAssessment, "user_assessment_id"),
				func(s *terraform.State) (err error) {
					return err
				},
			),
		},
	})
}
