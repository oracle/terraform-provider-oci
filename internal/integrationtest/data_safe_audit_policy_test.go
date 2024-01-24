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
	DataSafeAuditPolicyRequiredOnlyResource = DataSafeAuditPolicyResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_data_safe_audit_policy", "test_audit_policy", acctest.Required, acctest.Create, auditPolicyRepresentation)

	DataSafeAuditPolicyResourceConfig = DataSafeAuditPolicyResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_data_safe_audit_policy", "test_audit_policy", acctest.Optional, acctest.Update, auditPolicyRepresentation)

	DataSafeauditPolicySingularDataSourceRepresentation = map[string]interface{}{
		"audit_policy_id": acctest.Representation{RepType: acctest.Required, Create: `${var.policy_id}`},
	}

	DataSafeauditPolicyDataSourceRepresentation = map[string]interface{}{
		"audit_policy_id": acctest.Representation{RepType: acctest.Optional, Create: `${var.policy_id}`},
	}

	auditPolicyRepresentation = map[string]interface{}{
		"audit_policy_id":              acctest.Representation{RepType: acctest.Required, Create: `${var.policy_id}`},
		"defined_tags":                 acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"description":                  acctest.Representation{RepType: acctest.Optional, Create: `Target database for HR and Payroll combined`, Update: `description2`},
		"display_name":                 acctest.Representation{RepType: acctest.Optional, Create: `AuditPolicy_HRandPayrollTarget`, Update: `displayName2`},
		"freeform_tags":                acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"provision_trigger":            acctest.Representation{RepType: acctest.Optional, Create: `0`, Update: `1`},
		"retrieve_from_target_trigger": acctest.Representation{RepType: acctest.Optional, Create: `0`, Update: `1`},
	}

	DataSafeAuditPolicyResourceDependencies = DefinedTagsDependencies
)

// issue-routing-tag: data_safe/default
func TestDataSafeAuditPolicyResource_basic(t *testing.T) {
	t.Skip("Create operation is not available for Audit Policy resource")
	httpreplay.SetScenario("TestDataSafeAuditPolicyResource_basic")
	defer httpreplay.SaveScenario()
	fmt.Printf("TestDataSafeAuditPolicyResource_basic  ***** CALLED ***** \n")
	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	policyId := utils.GetEnvSettingWithBlankDefault("policy_ocid")
	policyIdVariableStr := fmt.Sprintf("variable \"policy_id\" { default = \"%s\" }\n", policyId)

	resourceName := "oci_data_safe_audit_policy.test_audit_policy"
	datasourceName := "data.oci_data_safe_audit_policy.test_audit_policy"
	singularDatasourceName := "data.oci_data_safe_audit_policy.test_audit_policy"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+policyIdVariableStr+compartmentIdUVariableStr+DataSafeAuditPolicyResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_data_safe_audit_policy", "test_audit_policy", acctest.Optional, acctest.Create,
			acctest.RepresentationCopyWithNewProperties(auditPolicyRepresentation, map[string]interface{}{
				"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
			})), "datasafe", "auditPolicy", t)

	fmt.Printf("TestDataSafeAuditPolicyResource_basic  ***** CREATED ***** \n")

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + policyIdVariableStr + DataSafeAuditPolicyResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_data_safe_audit_policy", "test_audit_policy", acctest.Required, acctest.Create, auditPolicyRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "audit_policy_id"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + policyIdVariableStr + DataSafeAuditPolicyResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + policyIdVariableStr + compartmentIdUVariableStr + DataSafeAuditPolicyResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_data_safe_audit_policy", "test_audit_policy", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(auditPolicyRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "audit_policy_id"),
				resource.TestCheckResourceAttr(resourceName, "description", "Target database for HR and Payroll combined"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "AuditPolicy_HRandPayrollTarget"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "is_data_safe_service_account_excluded"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					if isEnableExportCompartment, _ := strconv.ParseBool(utils.GetEnvSettingWithDefault("enable_export_compartment", "false")); isEnableExportCompartment {
						if errExport := resourcediscovery.TestExportCompartmentWithResourceName(&resId, &compartmentId, resourceName); errExport != nil {
							return errExport
						}
					}
					return err
				},
			),
		},

		// verify updates to updatable parameters
		{
			Config: config + compartmentIdVariableStr + policyIdVariableStr + DataSafeAuditPolicyResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_data_safe_audit_policy", "test_audit_policy", acctest.Optional, acctest.Update,
					acctest.RepresentationCopyWithNewProperties(auditPolicyRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "audit_policy_id"),
				resource.TestCheckResourceAttr(resourceName, "description", "description2"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "is_data_safe_service_account_excluded"),
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
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_data_safe_audit_policy", "test_audit_policy", acctest.Optional, acctest.Update, DataSafeauditPolicyDataSourceRepresentation) +
				compartmentIdVariableStr + policyIdVariableStr + DataSafeAuditPolicyResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_data_safe_audit_policy", "test_audit_policy", acctest.Optional, acctest.Update, auditPolicyRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "audit_policy_id"),
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_data_safe_audit_policy", "test_audit_policy", acctest.Required, acctest.Create, DataSafeauditPolicySingularDataSourceRepresentation) +
				compartmentIdVariableStr + policyIdVariableStr + DataSafeAuditPolicyResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "audit_policy_id"),

				//resource.TestCheckResourceAttr(singularDatasourceName, "audit_conditions.#", "19"),
				//resource.TestCheckResourceAttr(singularDatasourceName, "audit_specifications.#", "19"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "description", "description2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "is_data_safe_service_account_excluded"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				//resource.TestCheckResourceAttrSet(singularDatasourceName, "time_last_provisioned"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_last_retrieved"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
		// remove singular datasource from previous step so that it doesn't conflict with import tests
		{
			Config: config + compartmentIdVariableStr + policyIdVariableStr + DataSafeAuditPolicyResourceConfig,
		},
		// verify resource import
		{
			Config:                  config,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{`audit_policy_id`, `provision_trigger`, `retrieve_from_target_trigger`},
			ResourceName:            resourceName,
		},
	})
}
