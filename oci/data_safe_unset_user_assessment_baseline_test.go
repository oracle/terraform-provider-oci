// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	unsetUserAssessmentBaselineRepresentation = map[string]interface{}{
		"user_assessment_id": Representation{RepType: Required, Create: `${oci_data_safe_user_assessment.test_user_assessment.id}`},
	}

	UnsetUserAssessmentBaselineResourceDependencies = GenerateResourceFromRepresentationMap("oci_data_safe_user_assessment", "test_user_assessment", Required, Create, userAssessmentRepresentation)
)

func TestDataSafeUnsetUserAssessmentBaselineResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDataSafeUnsetUserAssessmentBaselineResource_basic")
	defer httpreplay.SaveScenario()

	config := ProviderTestConfig()

	compartmentId := GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	// Save TF content to Create resource with only required properties. This has to be exactly the same as the config part in the Create step in the test.
	SaveConfigContent(config+compartmentIdVariableStr+UnsetUserAssessmentBaselineResourceDependencies+
		GenerateResourceFromRepresentationMap("oci_data_safe_unset_user_assessment_baseline", "test_unset_user_assessment_baseline", Required, Create, unsetUserAssessmentBaselineRepresentation), "datasafe", "unsetUserAssessmentBaseline", t)

	ResourceTest(t, nil, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + UnsetUserAssessmentBaselineResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_data_safe_unset_user_assessment_baseline", "test_unset_user_assessment_baseline", Required, Create, unsetUserAssessmentBaselineRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				func(s *terraform.State) (err error) {
					return nil
				},
			),
		},
	})
}
