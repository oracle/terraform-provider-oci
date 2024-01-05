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
	DataSafeAuditPolicyManagementResource = DataSafeAuditPolicyManagementResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_data_safe_audit_policy", "test_audit_policy", acctest.Required, acctest.Create, auditPolicyManagementRepresentation)

	auditPolicyManagementRepresentation = map[string]interface{}{
		"compartment_id":                        acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"target_id":                             acctest.Representation{RepType: acctest.Required, Create: `${var.target_id}`},
		"defined_tags":                          acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"description":                           acctest.Representation{RepType: acctest.Optional, Create: `Target database for HR and Payroll combined`, Update: `description2`},
		"display_name":                          acctest.Representation{RepType: acctest.Optional, Create: `AuditPolicy_HRandPayrollTarget`, Update: `displayName2`},
		"freeform_tags":                         acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"retrieve_from_target_trigger":          acctest.Representation{RepType: acctest.Optional, Create: `true`, Update: `true`},
		"is_data_safe_service_account_excluded": acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `false`},
	}

	auditConditionsRepresentation = map[string]interface{}{
		"audit_policy_name":                  acctest.Representation{RepType: acctest.Optional, Create: `Database schema changes`, Update: `Database schema changes`},
		"is_priv_users_managed_by_data_safe": acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `false`},
		"is_enabled":                         acctest.Representation{RepType: acctest.Optional, Create: `true`, Update: `true`},
	}

	DataSafeAuditPolicyManagementResourceDependencies = DefinedTagsDependencies
)

// issue-routing-tag: data_safe/default
func TestDataSafeAuditPolicyManagementResource_basic(t *testing.T) {
	t.Skip("Create operation is not available for Audit Policy Management resource")
	httpreplay.SetScenario("TestDataSafeAuditPolicyManagementResource_basic")
	defer httpreplay.SaveScenario()
	fmt.Printf("TestDataSafeAuditPolicyManagementResource_basic  ***** CALLED ***** \n")
	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	targetId := utils.GetEnvSettingWithBlankDefault("target_id")
	targetIdVariableStr := fmt.Sprintf("variable \"target_id\" { default = \"%s\" }\n", targetId)

	resourceName := "oci_data_safe_audit_policy_management.test_audit_policy_management"

	acctest.SaveConfigContent(config+compartmentIdVariableStr+targetIdVariableStr+compartmentIdUVariableStr+DataSafeAuditPolicyManagementResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_data_safe_audit_policy_management", "test_audit_policy_management", acctest.Optional, acctest.Update,
			acctest.RepresentationCopyWithNewProperties(auditPolicyManagementRepresentation, map[string]interface{}{
				"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
			})), "datasafe", "auditPolicyManagement", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify Provision, Retrieval and Update
		{
			Config: config + compartmentIdVariableStr + targetIdVariableStr + DataSafeAuditPolicyManagementResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_data_safe_audit_policy_management", "test_audit_policy_management", acctest.Optional, acctest.Update,
					acctest.RepresentationCopyWithNewProperties(auditPolicyManagementRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "description", "description2"),
			),
		},
	})
}
