// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/terraform-providers/terraform-provider-oci/internal/acctest"
	"github.com/terraform-providers/terraform-provider-oci/internal/resourcediscovery"
	"github.com/terraform-providers/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	compareUserAssessmentRepresentation = map[string]interface{}{
		"comparison_user_assessment_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_data_safe_user_assessment.test_user_assessment1.id}`},
		"user_assessment_id":            acctest.Representation{RepType: acctest.Required, Create: `${oci_data_safe_user_assessment.test_user_assessment2.id}`},
	}

	CompareUserAssessmentResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", acctest.Required, acctest.Create, autonomousDatabaseRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_data_safe_target_database", "test_target_database", acctest.Required, acctest.Create, targetDatabaseRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_data_safe_user_assessment", "test_user_assessment1", acctest.Required, acctest.Create, userAssessmentRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_data_safe_user_assessment", "test_user_assessment2", acctest.Required, acctest.Create, userAssessmentRepresentation)
)

func TestDataSafeCompareUserAssessmentResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDataSafeCompareUserAssessmentResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_data_safe_compare_user_assessment.test_compare_user_assessment"

	var resId string
	// Save TF content to Create resource with only required properties. This has to be exactly the same as the config part in the Create step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+CompareUserAssessmentResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_data_safe_compare_user_assessment", "test_compare_user_assessment", acctest.Required, acctest.Create, compareUserAssessmentRepresentation), "datasafe", "userAssessment", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + CompareUserAssessmentResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_data_safe_compare_user_assessment", "test_compare_user_assessment", acctest.Required, acctest.Create, compareUserAssessmentRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "comparison_user_assessment_id"),
				resource.TestCheckResourceAttrSet(resourceName, "user_assessment_id"),
				resource.TestCheckResourceAttr(resourceName, "summary.#", "0"),

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
	})
}
