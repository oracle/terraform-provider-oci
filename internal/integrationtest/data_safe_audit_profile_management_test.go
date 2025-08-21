// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	DataSafeAuditProfileManagementResourceRepresentation = map[string]interface{}{
		"compartment_id":                acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"target_id":                     acctest.Representation{RepType: acctest.Required, Create: `${var.target_id}`},
		"target_type":                   acctest.Representation{RepType: acctest.Required, Create: `TARGET_DATABASE`},
		"display_name":                  acctest.Representation{RepType: acctest.Optional, Create: `Audit_1`, Update: `displayName2`},
		"description":                   acctest.Representation{RepType: acctest.Optional, Create: `Description`, Update: `description2`},
		"freeform_tags":                 acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"is_override_global_paid_usage": acctest.Representation{RepType: acctest.Optional, Create: `true`},
		"is_paid_usage_enabled":         acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
	}
	DataSafeAuditProfileManagementResourceDependencies = DefinedTagsDependencies
)

// issue-routing-tag: data_safe/default
func TestDataSafeAuditProfileManagementResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDataSafeAuditProfileResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	targetId := utils.GetEnvSettingWithBlankDefault("target_ocid")
	targetIdVariableStr := fmt.Sprintf("variable \"target_id\" { default = \"%s\" }\n", targetId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_data_safe_audit_profile_management.test_audit_profile_management"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+
		acctest.GenerateResourceFromRepresentationMap("oci_data_safe_audit_profile_management", "test_audit_profile_management", acctest.Optional, acctest.Create, DataSafeAuditProfileManagementResourceRepresentation), "datasafe", "auditProfileManagement", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify Create
		// Audit profile for target database
		{
			Config: config + compartmentIdVariableStr + targetIdVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_data_safe_audit_profile_management", "test_audit_profile_management", acctest.Required, acctest.Create, DataSafeAuditProfileManagementResourceRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName, "target_id"),
				resource.TestCheckResourceAttr(resourceName, "target_type", "TARGET_DATABASE"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					fmt.Print(resId)
					return err
				},
			),
		},

		// verify Update
		// Audit profile for target database
		{
			Config: config + compartmentIdVariableStr + targetIdVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_data_safe_audit_profile_management", "test_audit_profile_management", acctest.Optional, acctest.Update, DataSafeAuditProfileManagementResourceRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName, "target_id"),
				resource.TestCheckResourceAttr(resourceName, "target_type", "TARGET_DATABASE"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + targetIdVariableStr,
		},
		// Create Audit profile for target group with optional
		{
			Config: config + compartmentIdVariableStr + DataSafeAuditProfileResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_data_safe_audit_profile_management", "test_audit_profile_management", acctest.Optional, acctest.Create, DataSafeAuditProfileRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "description", "Description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "Audit_1"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_override_global_paid_usage", "true"),
				resource.TestCheckResourceAttrSet(resourceName, "is_override_global_retention_setting"),
				resource.TestCheckResourceAttr(resourceName, "is_paid_usage_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "offline_months", "10"),
				resource.TestCheckResourceAttr(resourceName, "online_months", "10"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "target_id"),
				resource.TestCheckResourceAttr(resourceName, "target_type", "TARGET_DATABASE_GROUP"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "time_updated"),

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
		// Verify update
		{
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + DataSafeAuditProfileResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_data_safe_audit_profile_management", "test_audit_profile_management", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(DataSafeAuditProfileRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "description", "Description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "Audit_1"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_override_global_paid_usage", "true"),
				resource.TestCheckResourceAttrSet(resourceName, "is_override_global_retention_setting"),
				resource.TestCheckResourceAttr(resourceName, "is_paid_usage_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "offline_months", "10"),
				resource.TestCheckResourceAttr(resourceName, "online_months", "10"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "target_id"),
				resource.TestCheckResourceAttr(resourceName, "target_type", "TARGET_DATABASE_GROUP"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "time_updated"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("resource recreated when it was supposed to be updated")
					}
					return err
				},
			),
		},
		// Delete Audit profile for Target group
		{
			Config: config + compartmentIdVariableStr,
		},
	})
}
