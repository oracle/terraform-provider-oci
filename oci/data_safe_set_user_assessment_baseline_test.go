// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	setUserAssessmentBaselineRepresentation = map[string]interface{}{
		"user_assessment_id": Representation{repType: Required, create: `${oci_data_safe_user_assessment.test_user_assessment1.id}`},
		"assessment_ids":     Representation{repType: Optional, create: []string{`${oci_data_safe_user_assessment.test_user_assessment2.id}`}},
	}

	unsetAssessmentBaselineRepresentation = map[string]interface{}{
		"user_assessment_id": Representation{repType: Required, create: `${oci_data_safe_user_assessment.test_user_assessment1.id}`},
	}

	SetUserAssessmentBaselineResourceDependencies = generateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", Required, Create, autonomousDatabaseRepresentation) +
		generateResourceFromRepresentationMap("oci_data_safe_target_database", "test_target_database", Required, Create, targetDatabaseRepresentation) +
		generateResourceFromRepresentationMap("oci_data_safe_user_assessment", "test_user_assessment1", Required, Create, userAssessmentRepresentation) +
		generateResourceFromRepresentationMap("oci_data_safe_user_assessment", "test_user_assessment2", Required, Create, userAssessmentRepresentation)
)

func TestDataSafeSetUserAssessmentBaselineResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDataSafeSetUserAssessmentBaselineResource_basic")
	defer httpreplay.SaveScenario()

	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_data_safe_set_user_assessment_baseline.test_set_user_assessment_baseline"

	var resId string
	// Save TF content to create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	saveConfigContent(config+compartmentIdVariableStr+SetUserAssessmentBaselineResourceDependencies+
		generateResourceFromRepresentationMap("oci_data_safe_set_user_assessment_baseline", "test_set_user_assessment_baseline", Optional, Create, setUserAssessmentBaselineRepresentation), "datasafe", "setUserAssessmentBaseline", t)

	ResourceTest(t, nil, []resource.TestStep{
		// verify create
		{
			Config: config + compartmentIdVariableStr + SetUserAssessmentBaselineResourceDependencies +
				generateResourceFromRepresentationMap("oci_data_safe_set_user_assessment_baseline", "test_set_user_assessment_baseline", Optional, Create, setUserAssessmentBaselineRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "user_assessment_id"),
			),
		},

		// delete before next create
		{
			Config: config + compartmentIdVariableStr + SetUserAssessmentBaselineResourceDependencies,
		},
		// verify create with optionals
		{
			Config: config + compartmentIdVariableStr + SetUserAssessmentBaselineResourceDependencies +
				generateResourceFromRepresentationMap("oci_data_safe_set_user_assessment_baseline", "test_set_user_assessment_baseline", Optional, Create, setUserAssessmentBaselineRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "assessment_ids.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "user_assessment_id"),

				func(s *terraform.State) (err error) {
					resId, err = fromInstanceState(s, resourceName, "id")
					if isEnableExportCompartment, _ := strconv.ParseBool(getEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
						if errExport := testExportCompartmentWithResourceName(&resId, &compartmentId, resourceName); errExport != nil {
							return errExport
						}
					}
					return err
				},
			),
		},
		{
			Config: config + compartmentIdVariableStr + SetUserAssessmentBaselineResourceDependencies +
				generateResourceFromRepresentationMap("oci_data_safe_unset_user_assessment_baseline", "test_unset_user_assessment_baseline", Required, Create, unsetAssessmentBaselineRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				// resource.TestCheckResourceAttrSet(resourceNameAssessment, "user_assessment_id"),
				func(s *terraform.State) (err error) {
					return err
				},
			),
		},
	})
}
