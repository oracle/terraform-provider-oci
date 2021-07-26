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
	UnsetSecurityAssessmentBaselineRepresentation = map[string]interface{}{
		"security_assessment_id": Representation{repType: Required, create: `${oci_data_safe_security_assessment.test_security_assessment1.id}`},
	}

	UnsetSecurityAssessmentBaselineResourceDependencies = SetSecurityAssessmentBaselineResourceDependencies +
		generateResourceFromRepresentationMap("oci_data_safe_set_security_assessment_baseline", "test_set_security_assessment_baseline", Required, Create, setSecurityAssessmentBaselineRepresentation)
)

// issue-routing-tag: data_safe/default
func TestDataSafeUnsetSecurityAssessmentBaselineResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDataSafeUnsetSecurityAssessmentBaselineResource_basic")
	defer httpreplay.SaveScenario()

	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	// Save TF content to create resource with only required properties. This has to be exactly the same as the config part in the create step in the test.
	saveConfigContent(config+compartmentIdVariableStr+UnsetSecurityAssessmentBaselineResourceDependencies+
		generateResourceFromRepresentationMap("oci_data_safe_unset_security_assessment_baseline", "test_unset_security_assessment_baseline", Required, Create, UnsetSecurityAssessmentBaselineRepresentation), "datasafe", "unsetSecurityAssessmentBaseline", t)

	ResourceTest(t, nil, []resource.TestStep{
		// verify create
		{
			Config: config + compartmentIdVariableStr + UnsetSecurityAssessmentBaselineResourceDependencies +
				generateResourceFromRepresentationMap("oci_data_safe_unset_security_assessment_baseline", "test_unset_security_assessment_baseline", Required, Create, UnsetSecurityAssessmentBaselineRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				func(s *terraform.State) (err error) {
					return nil
				},
			),
		},
	})
}
