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
	DataSafeAuditProfileRequiredOnlyResource = DataSafeAuditProfileResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_data_safe_audit_profile", "test_audit_profile", acctest.Required, acctest.Create, auditProfileRepresentation)

	DataSafeAuditProfileResourceConfig = DataSafeAuditProfileResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_data_safe_audit_profile", "test_audit_profile", acctest.Optional, acctest.Update, auditProfileRepresentation)

	DataSafeauditProfileSingularDataSourceRepresentation = map[string]interface{}{
		"audit_profile_id": acctest.Representation{RepType: acctest.Required, Create: `${var.profile_id}`},
	}

	DataSafeauditProfileDataSourceRepresentation = map[string]interface{}{
		"audit_profile_id": acctest.Representation{RepType: acctest.Optional, Create: `${var.profile_id}`},
	}

	auditProfileRepresentation = map[string]interface{}{
		"audit_profile_id":         acctest.Representation{RepType: acctest.Required, Create: `${var.profile_id}`},
		"defined_tags":             acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"description":              acctest.Representation{RepType: acctest.Optional, Create: `updated-description`, Update: `description2`},
		"freeform_tags":            acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"change_retention_trigger": acctest.Representation{RepType: acctest.Optional, Create: `0`, Update: `1`},
	}

	DataSafeAuditProfileResourceDependencies = DefinedTagsDependencies
)

// issue-routing-tag: data_safe/default
func TestDataSafeAuditProfileResource_basic(t *testing.T) {
	t.Skip("Create/Delete operation is not available for Audit Profile resource")
	httpreplay.SetScenario("TestDataSafeAuditProfileResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	profileId := utils.GetEnvSettingWithBlankDefault("profile_ocid")
	profileIdVariableStr := fmt.Sprintf("variable \"profile_id\" { default = \"%s\" }\n", profileId)

	resourceName := "oci_data_safe_audit_profile.test_audit_profile"
	datasourceName := "data.oci_data_safe_audit_profile.test_audit_profile"
	singularDatasourceName := "data.oci_data_safe_audit_profile.test_audit_profile"
	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+profileIdVariableStr+compartmentIdUVariableStr+DataSafeAuditProfileResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_data_safe_audit_profile", "test_audit_profile", acctest.Optional, acctest.Create,
			acctest.RepresentationCopyWithNewProperties(auditProfileRepresentation, map[string]interface{}{
				"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
			})), "datasafe", "auditProfile", t)
	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + profileIdVariableStr + DataSafeAuditProfileRequiredOnlyResource,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "audit_profile_id"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + profileIdVariableStr + DataSafeAuditProfileResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + profileIdVariableStr + compartmentIdUVariableStr + DataSafeAuditProfileResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_data_safe_audit_profile", "test_audit_profile", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(auditProfileRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "audit_profile_id"),
				resource.TestCheckResourceAttr(resourceName, "description", "updated-description"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "offline_months"),
				resource.TestCheckResourceAttrSet(resourceName, "online_months"),
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

		// verify updates to updatable parameters
		{
			Config: config + compartmentIdVariableStr + profileIdVariableStr + DataSafeAuditProfileResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_data_safe_audit_profile", "test_audit_profile", acctest.Optional, acctest.Update,
					acctest.RepresentationCopyWithNewProperties(auditProfileRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "audit_profile_id"),
				resource.TestCheckResourceAttr(resourceName, "description", "description2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "offline_months"),
				resource.TestCheckResourceAttrSet(resourceName, "online_months"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "time_updated"),

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
				acctest.GenerateDataSourceFromRepresentationMap("oci_data_safe_audit_profile", "test_audit_profile", acctest.Optional, acctest.Update, DataSafeauditProfileDataSourceRepresentation) +
				compartmentIdVariableStr + profileIdVariableStr + DataSafeAuditProfileResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_data_safe_audit_profile", "test_audit_profile", acctest.Optional, acctest.Update, auditProfileRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "audit_profile_id"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_data_safe_audit_profile", "test_audit_profile", acctest.Required, acctest.Create, DataSafeauditProfileSingularDataSourceRepresentation) +
				compartmentIdVariableStr + profileIdVariableStr + DataSafeAuditProfileResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "audit_profile_id"),

				resource.TestCheckResourceAttrSet(singularDatasourceName, "audit_collected_volume"),
				resource.TestCheckResourceAttr(singularDatasourceName, "audit_trails.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "description", "description2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "is_override_global_retention_setting"),
				resource.TestCheckResourceAttr(singularDatasourceName, "is_paid_usage_enabled", "true"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "offline_months"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "online_months"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
		// remove singular datasource from previous step so that it doesn't conflict with import tests
		{
			Config: config + compartmentIdVariableStr + profileIdVariableStr + DataSafeAuditProfileResourceConfig,
		},
		// verify resource import
		{
			Config:                  config,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{"audit_profile_id", "change_retention_trigger"},
			ResourceName:            resourceName,
		},
	})
}
