// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/terraform"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	DataSafeSecurityPolicyManagementResourceConfig = DataSafeSecurityPolicyManagementResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_data_safe_security_policy_management", "test_security_policy_management", acctest.Optional, acctest.Create, DataSafeSecurityPolicyManagementRepresentation)

	DataSafeSecurityPolicyManagementRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"target_id":      acctest.Representation{RepType: acctest.Optional, Create: `${var.target_id}`},
		"description":    acctest.Representation{RepType: acctest.Optional, Create: `description2`, Update: `updatedDescription`},
		"display_name":   acctest.Representation{RepType: acctest.Optional, Create: `displayName2`, Update: `updatedDisplayName`},
		"freeform_tags":  acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"lifecycle":      acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreSecurityPolicyManagementTagsChangesRep},
	}

	DataSafeSecurityPolicyCreateRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"description":    acctest.Representation{RepType: acctest.Optional, Create: `description1`, Update: `updatedDescription`},
		"display_name":   acctest.Representation{RepType: acctest.Optional, Create: `displayName1`, Update: `updatedDisplayName`},
		"freeform_tags":  acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"lifecycle":      acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreSecurityPolicyManagementTagsChangesRep},
	}

	ignoreSecurityPolicyManagementTagsChangesRep = map[string]interface{}{
		"ignore_changes": acctest.Representation{RepType: acctest.Required, Create: []string{`system_tags`, `defined_tags`, `freeform_tags`}},
	}

	DataSafeSecurityPolicyManagementResourceDependencies = ""
)

// issue-routing-tag: data_safe/default
func TestDataSafeSecurityPolicyManagementResource_basic(t *testing.T) {
	t.Skip("Skipping this test as the target ocid is hardcoded and may not exist when the test runs")
	httpreplay.SetScenario("TestDataSafeSecurityPolicyManagementResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	targetId := utils.GetEnvSettingWithBlankDefault("target_ocid")
	targetIdVariableStr := fmt.Sprintf("variable \"target_id\" { default = \"%s\" }\n", targetId)

	resourceName := "oci_data_safe_security_policy_management.test_security_policy_management"

	var resId, resId2, resId3, resId4 string

	acctest.SaveConfigContent(config+compartmentIdVariableStr+compartmentIdUVariableStr+targetIdVariableStr+DataSafeSecurityPolicyManagementResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_data_safe_security_policy_management", "test_security_policy_management", acctest.Optional, acctest.Create,
			acctest.RepresentationCopyWithNewProperties(DataSafeSecurityPolicyManagementRepresentation, map[string]interface{}{
				"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
			})), "datasafe", "securityPolicyManagement", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{

		//verify updates to updatable parameters for autocreated Security Policy if targetId is present
		{
			Config: config + compartmentIdVariableStr + targetIdVariableStr + DataSafeSecurityPolicyManagementResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "description", "description2"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		//verify update to the compartment (the compartment will be switched back in the next step)
		{
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + targetIdVariableStr + DataSafeSecurityPolicyManagementResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_data_safe_security_policy_management", "test_security_policy_management", acctest.Optional, acctest.Update,
					acctest.RepresentationCopyWithNewProperties(DataSafeSecurityPolicyManagementRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Update: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "description", "updatedDescription"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "updatedDisplayName"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("Resource recreated when it was supposed to be updated.")
					}
					return err
				},
			),
		},

		//revert back the compartment change
		{
			Config: config + compartmentIdVariableStr + targetIdVariableStr + DataSafeSecurityPolicyManagementResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("Resource recreated when it was supposed to be updated.")
					}
					return err
				},
			),
		},

		// Delete before next Create
		{
			Config: config + compartmentIdVariableStr,
		},

		// Create new Security Policy if target_id is not present
		{
			Config: config + compartmentIdVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_data_safe_security_policy_management", "test_security_policy_management", acctest.Optional, acctest.Create, DataSafeSecurityPolicyCreateRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "description", "description1"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),

				func(s *terraform.State) (err error) {
					resId3, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId == resId3 {
						return fmt.Errorf("New resource did not get created when target_id is not present")
					}
					return err
				},
			),
		},

		//verify updates to updatable parameters
		{
			Config: config + compartmentIdVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_data_safe_security_policy_management", "test_security_policy_management", acctest.Optional, acctest.Update, DataSafeSecurityPolicyCreateRepresentation),

			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "description", "updatedDescription"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "updatedDisplayName"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),

				func(s *terraform.State) (err error) {
					resId4, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId3 != resId4 {
						return fmt.Errorf("Resource recreated when it was supposed to be updated.")
					}
					return err
				},
			),
		},

		//verify update to the compartment (the compartment will be switched back in the next step)
		{
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_data_safe_security_policy_management", "test_security_policy_management", acctest.Optional, acctest.Update,
					acctest.RepresentationCopyWithNewProperties(DataSafeSecurityPolicyCreateRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Update: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "description", "updatedDescription"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "updatedDisplayName"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),

				func(s *terraform.State) (err error) {
					resId4, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId3 != resId4 {
						return fmt.Errorf("Resource recreated when it was supposed to be updated.")
					}
					return err
				},
			),
		},

		//revert back the compartment change
		{
			Config: config + compartmentIdVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_data_safe_security_policy_management", "test_security_policy_management", acctest.Optional, acctest.Update, DataSafeSecurityPolicyCreateRepresentation),

			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				func(s *terraform.State) (err error) {
					resId4, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId3 != resId4 {
						return fmt.Errorf("Resource recreated when it was supposed to be updated.")
					}
					return err
				},
			),
		},

		// Delete resource
		{
			Config: config + compartmentIdVariableStr,
		},
	})
}
