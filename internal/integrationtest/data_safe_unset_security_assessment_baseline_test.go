// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"testing"

	"github.com/terraform-providers/terraform-provider-oci/internal/acctest"
	"github.com/terraform-providers/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	UnsetSecurityAssessmentBaselineRepresentation = map[string]interface{}{
		"security_assessment_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_data_safe_security_assessment.test_security_assessment1.id}`},
	}

	UnsetSecurityAssessmentBaselineResourceDependencies = SetSecurityAssessmentBaselineResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_data_safe_set_security_assessment_baseline", "test_set_security_assessment_baseline", acctest.Required, acctest.Create, setSecurityAssessmentBaselineRepresentation)
)

// issue-routing-tag: data_safe/default
func TestDataSafeUnsetSecurityAssessmentBaselineResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDataSafeUnsetSecurityAssessmentBaselineResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	// Save TF content to Create resource with only required properties. This has to be exactly the same as the config part in the Create step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+UnsetSecurityAssessmentBaselineResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_data_safe_unset_security_assessment_baseline", "test_unset_security_assessment_baseline", acctest.Required, acctest.Create, UnsetSecurityAssessmentBaselineRepresentation), "datasafe", "unsetSecurityAssessmentBaseline", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + UnsetSecurityAssessmentBaselineResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_data_safe_unset_security_assessment_baseline", "test_unset_security_assessment_baseline", acctest.Required, acctest.Create, UnsetSecurityAssessmentBaselineRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				func(s *terraform.State) (err error) {
					return nil
				},
			),
		},
	})
}
