// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"

	"github.com/oracle/terraform-provider-oci/internal/acctest"

	"github.com/oracle/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/oracle/terraform-provider-oci/httpreplay"
)

var (
	MaskingPoliciesApplyDifferenceToMaskingColumnsRepresentation = map[string]interface{}{
		"masking_policy_id":                acctest.Representation{RepType: acctest.Required, Create: `${oci_data_safe_masking_policy.test_masking_policy.id}`},
		"sdm_masking_policy_difference_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_data_safe_sdm_masking_policy_difference.test_sdm_masking_policy_difference.id}`},
	}

	MaskingPoliciesApplyDifferenceToMaskingColumnsDependencies = acctest.GenerateResourceFromRepresentationMap("oci_data_safe_masking_policy", "test_masking_policy", acctest.Required, acctest.Create, maskingPolicyRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_data_safe_sdm_masking_policy_difference", "test_sdm_masking_policy_difference", acctest.Required, acctest.Create, DataSafeSdmMaskingPolicyDifferenceRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_data_safe_sensitive_data_model", "test_sensitive_data_model1", acctest.Required, acctest.Create, sensitiveDataModelRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_data_safe_sensitive_data_model", "test_sensitive_data_model2", acctest.Required, acctest.Create, sensitiveDataModelRepresentation)
)

// issue-routing-tag: data_safe/default
func TestDataSafeMaskingPoliciesApplyDifferenceToMaskingColumnsResource_basic(t *testing.T) {
	httpreplay.SetScenario("MaskingPoliciesApplyDifferenceToMaskingColumnsResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	targetId := utils.GetEnvSettingWithBlankDefault("data_safe_target_ocid")
	targetIdVariableStr := fmt.Sprintf("variable \"target_id\" { default = \"%s\" }\n", targetId)

	resourceName := "oci_data_safe_masking_policies_apply_difference_to_masking_columns.test_masking_policies_apply_difference_to_masking_columns"

	var resId string
	// Save TF content to Create resource with only required properties. This has to be exactly the same as the config part in the Create step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+MaskingPoliciesApplyDifferenceToMaskingColumnsDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_data_safe_masking_policies_apply_difference_to_masking_columns", "test_masking_policies_apply_difference_to_masking_columns", acctest.Required, acctest.Create, MaskingPoliciesApplyDifferenceToMaskingColumnsRepresentation), "datasafe", "maskingPolicy", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + MaskingPoliciesApplyDifferenceToMaskingColumnsDependencies + targetIdVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_data_safe_masking_policies_apply_difference_to_masking_columns", "test_masking_policies_apply_difference_to_masking_columns", acctest.Required, acctest.Create, MaskingPoliciesApplyDifferenceToMaskingColumnsRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(

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
