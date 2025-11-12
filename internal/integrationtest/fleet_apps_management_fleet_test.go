// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/oracle/oci-go-sdk/v65/common"
	oci_fleet_apps_management "github.com/oracle/oci-go-sdk/v65/fleetappsmanagement"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	FleetAppsManagementFleetRequiredOnlyResource = FleetAppsManagementFleetResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_fleet_apps_management_fleet", "test_fleet", acctest.Required, acctest.Create, FleetAppsManagementFleetRepresentation)

	FleetAppsManagementFleetResourceConfig = FleetAppsManagementFleetResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_fleet_apps_management_fleet", "test_fleet", acctest.Optional, acctest.Update, FleetAppsManagementFleetRepresentation)

	FleetAppsManagementFleetSingularDataSourceRepresentation = map[string]interface{}{
		"fleet_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_fleet_apps_management_fleet.test_fleet.id}`},
	}

	FleetAppsManagementFleetDataSourceRepresentation = map[string]interface{}{
		"application_type": acctest.Representation{RepType: acctest.Optional, Create: `applicationType`},
		"compartment_id":   acctest.Representation{RepType: acctest.Optional, Create: `${var.compartment_id}`},
		"display_name":     acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"environment_type": acctest.Representation{RepType: acctest.Optional, Create: `environmentType`, Update: `environmentType2`},
		"fleet_type":       acctest.Representation{RepType: acctest.Optional, Create: `GENERIC`},
		"id":               acctest.Representation{RepType: acctest.Optional, Create: `${oci_fleet_apps_management_fleet.test_fleet.id}`},
		"state":            acctest.Representation{RepType: acctest.Optional, Create: `NEEDS_ATTENTION`},
		"filter":           acctest.RepresentationGroup{RepType: acctest.Required, Group: FleetAppsManagementFleetDataSourceFilterRepresentation}}
	FleetAppsManagementFleetDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_fleet_apps_management_fleet.test_fleet.id}`}},
	}

	FleetAppsManagementFleetRepresentation = map[string]interface{}{
		"compartment_id":     acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":       acctest.Representation{RepType: acctest.Required, Create: `displayName`, Update: `displayName2`},
		"resource_selection": acctest.RepresentationGroup{RepType: acctest.Required, Group: FleetAppsManagementFleetResourceSelectionRepresentation},
		"credentials":        acctest.RepresentationGroup{RepType: acctest.Optional, Group: FleetAppsManagementFleetCredentialsRepresentation},
		// "defined_tags":       acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"defined_tags":             acctest.Representation{RepType: acctest.Optional, Create: `${map("Oracle-Tags.CreatedBy", "value")}`, Update: `${map("Oracle-Tags.CreatedBy", "updatedValue")}`},
		"description":              acctest.Representation{RepType: acctest.Optional, Create: `description`, Update: `description2`},
		"details":                  acctest.RepresentationGroup{RepType: acctest.Required, Group: FleetAppsManagementFleetDetailsRepresentation},
		"environment_type":         acctest.Representation{RepType: acctest.Optional, Create: `Stage`, Update: `Stage`},
		"freeform_tags":            acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"bar-key": "value"}, Update: map[string]string{"Department": "Accounting"}},
		"is_target_auto_confirm":   acctest.Representation{RepType: acctest.Required, Create: `false`, Update: `true`},
		"notification_preferences": acctest.RepresentationGroup{RepType: acctest.Optional, Group: FleetAppsManagementFleetNotificationPreferencesRepresentation},
	}
	FleetAppsManagementFleetResourceSelectionRepresentation = map[string]interface{}{
		"resource_selection_type": acctest.Representation{RepType: acctest.Required, Create: `DYNAMIC`, Update: `DYNAMIC`},
		"rule_selection_criteria": acctest.RepresentationGroup{RepType: acctest.Required, Group: FleetAppsManagementFleetResourceSelectionRuleSelectionCriteriaRepresentation},
	}

	FleetAppsManagementFleetCredentialsRepresentation = map[string]interface{}{
		"compartment_id":   acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":     acctest.Representation{RepType: acctest.Required, Create: `tersi-testing-credential`},
		"entity_specifics": acctest.RepresentationGroup{RepType: acctest.Required, Group: FleetAppsManagementFleetCredentialsEntitySpecificsRepresentation},
		"password":         acctest.RepresentationGroup{RepType: acctest.Required, Group: FleetAppsManagementFleetCredentialsPasswordRepresentation},
		"user":             acctest.RepresentationGroup{RepType: acctest.Required, Group: FleetAppsManagementFleetCredentialsUserRepresentation},
	}
	FleetAppsManagementFleetDetailsRepresentation = map[string]interface{}{
		"fleet_type": acctest.Representation{RepType: acctest.Required, Create: `GENERIC`},
	}
	FleetAppsManagementFleetNotificationPreferencesRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"topic_id":       acctest.Representation{RepType: acctest.Required, Create: `${var.oci_ons_notification_topic}`},
		"preferences":    acctest.RepresentationGroup{RepType: acctest.Optional, Group: FleetAppsManagementFleetNotificationPreferencesPreferencesRepresentation},
	}
	FleetAppsManagementFleetPropertiesRepresentation = map[string]interface{}{
		"compartment_id":      acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"fleet_property_type": acctest.Representation{RepType: acctest.Required, Create: `STRING`},
		"display_name":        acctest.Representation{RepType: acctest.Optional, Create: `displayName`},
		"is_required":         acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"value":               acctest.Representation{RepType: acctest.Optional, Create: `value`},
	}
	FleetAppsManagementFleetResourcesRepresentation = map[string]interface{}{
		"compartment_id":      acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"resource_id":         acctest.Representation{RepType: acctest.Required, Create: `${var.test_instance_id}`},
		"tenancy_id":          acctest.Representation{RepType: acctest.Required, Create: `${var.tenancy_id}`},
		"fleet_resource_type": acctest.Representation{RepType: acctest.Optional, Create: `Instance`},
	}
	FleetAppsManagementFleetResourceSelectionRuleSelectionCriteriaRepresentation = map[string]interface{}{
		"match_condition": acctest.Representation{RepType: acctest.Required, Create: `ANY`, Update: `MATCH_ALL`},
		"rules":           acctest.RepresentationGroup{RepType: acctest.Required, Group: FleetAppsManagementFleetResourceSelectionRuleSelectionCriteriaRulesRepresentation},
	}

	//FleetAppsManagementFleetRuleSelectionCriteriaRepresentation = map[string]interface{}{
	//	"match_condition": acctest.Representation{RepType: acctest.Required, Create: `MATCH_ALL`, Update: `ANY`},
	//	"rules":           acctest.RepresentationGroup{RepType: acctest.Required, Group: FleetAppsManagementFleetRuleSelectionCriteriaRulesRepresentation},
	//}

	FleetAppsManagementFleetCredentialsEntitySpecificsRepresentation = map[string]interface{}{
		"credential_level": acctest.Representation{RepType: acctest.Required, Create: `FLEET`},
		"resource_id":      acctest.Representation{RepType: acctest.Optional, Create: `${var.test_instance_id}`},
		"target":           acctest.Representation{RepType: acctest.Optional, Create: `target`},
		"variables":        acctest.RepresentationGroup{RepType: acctest.Optional, Group: FleetAppsManagementFleetCredentialsEntitySpecificsVariablesRepresentation},
	}
	FleetAppsManagementFleetCredentialsPasswordRepresentation = map[string]interface{}{
		"credential_type": acctest.Representation{RepType: acctest.Required, Create: `PLAIN_TEXT`},
		"key_id":          acctest.Representation{RepType: acctest.Optional, Create: `${var.key_id}`},
		"key_version":     acctest.Representation{RepType: acctest.Optional, Create: `keyVersion`},
		"secret_id":       acctest.Representation{RepType: acctest.Optional, Create: `${var.fams_user_password}`},
		"secret_version":  acctest.Representation{RepType: acctest.Optional, Create: `secretVersion`},
		"value":           acctest.Representation{RepType: acctest.Optional, Create: `value`},
		"vault_id":        acctest.Representation{RepType: acctest.Optional, Create: `${var.vault_id}`},
	}
	FleetAppsManagementFleetCredentialsUserRepresentation = map[string]interface{}{
		"credential_type": acctest.Representation{RepType: acctest.Required, Create: `PLAIN_TEXT`},
		"key_id":          acctest.Representation{RepType: acctest.Optional, Create: `${var.key_id}`},
		"key_version":     acctest.Representation{RepType: acctest.Optional, Create: `keyVersion`},
		"secret_id":       acctest.Representation{RepType: acctest.Optional, Create: `${var.fams_user_id}`},
		"secret_version":  acctest.Representation{RepType: acctest.Optional, Create: `secretVersion`},
		"value":           acctest.Representation{RepType: acctest.Optional, Create: `value`},
		"vault_id":        acctest.Representation{RepType: acctest.Optional, Create: `${var.vault_id}`},
	}
	FleetAppsManagementFleetNotificationPreferencesPreferencesRepresentation = map[string]interface{}{
		"on_job_canceled":            acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"on_job_failure":             acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"on_job_schedule_change":     acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"on_job_start":               acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"on_job_success":             acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"on_resource_non_compliance": acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"on_runbook_newer_version":   acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"on_task_failure":            acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"on_task_pause":              acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"on_task_success":            acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"on_topology_modification":   acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"upcoming_schedule":          acctest.RepresentationGroup{RepType: acctest.Optional, Group: FleetAppsManagementFleetNotificationPreferencesPreferencesUpcomingScheduleRepresentation},
	}
	FleetAppsManagementFleetResourceSelectionRuleSelectionCriteriaRulesRepresentation = map[string]interface{}{
		"basis":                     acctest.Representation{RepType: acctest.Required, Create: `inventoryProperties`, Update: `inventoryProperties`},
		"compartment_id":            acctest.Representation{RepType: acctest.Required, Create: `${var.tenancy_id}`},
		"compartment_id_in_subtree": acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"conditions":                acctest.RepresentationGroup{RepType: acctest.Required, Group: FleetAppsManagementFleetResourceSelectionRuleSelectionCriteriaRulesConditionsRepresentation},
		"match_condition":           acctest.Representation{RepType: acctest.Optional, Create: `MATCH_ALL`, Update: `ANY`},
		// "resource_compartment_id":   acctest.Representation{RepType: acctest.Optional, Create: `${oci_identity_compartment.test_compartment.id}`},
		"resource_compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		///// EARLIER CHANGES BEGIN /////
		// "basis":                   acctest.Representation{RepType: acctest.Required, Create: `inventoryProperties`},
		// "compartment_id":          acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		// "conditions":              acctest.RepresentationGroup{RepType: acctest.Required, Group: FleetAppsManagementFleetResourceSelectionRuleSelectionCriteriaRulesConditionsRepresentation},
		// "resource_compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		// //"resource_compartment_id": acctest.Representation{RepType: acctest.Optional, Create: `${oci_identity_compartment.test_compartment.id}`},
		///// EARLIER CHANGES END /////
	}
	FleetAppsManagementFleetCredentialsEntitySpecificsVariablesRepresentation = map[string]interface{}{
		"name":  acctest.Representation{RepType: acctest.Optional, Create: `name`},
		"value": acctest.Representation{RepType: acctest.Optional, Create: `value`},
	}
	FleetAppsManagementFleetNotificationPreferencesPreferencesUpcomingScheduleRepresentation = map[string]interface{}{
		"notify_before":        acctest.Representation{RepType: acctest.Optional, Create: `notifyBefore`, Update: `notifyBefore2`},
		"on_upcoming_schedule": acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
	}
	FleetAppsManagementFleetResourceSelectionRuleSelectionCriteriaRulesConditionsRepresentation = map[string]interface{}{
		"attr_group": acctest.Representation{RepType: acctest.Required, Create: `Instance`},
		"attr_key":   acctest.Representation{RepType: acctest.Required, Create: `displayName`, Update: `shape`},
		"attr_value": acctest.Representation{RepType: acctest.Required, Create: `attrValue1`, Update: `VM.Standard.E4.Flex`},
	}

	FleetAppsManagementFleetResourceDependencies = ""
)

// issue-routing-tag: fleet_apps_management/default
func TestFleetAppsManagementFleetResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestFleetAppsManagementFleetResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	tenancyId := utils.GetEnvSettingWithBlankDefault("tenancy_ocid")
	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	instanceId := utils.GetEnvSettingWithBlankDefault("self_hosted_instance_id")
	keyId := utils.GetEnvSettingWithBlankDefault("key_id")
	vaultId := utils.GetEnvSettingWithBlankDefault("vault_id")
	onsNotificationTopicId := utils.GetEnvSettingWithBlankDefault("test_ons_topic")
	kmsVaultId := utils.GetEnvSettingWithBlankDefault("kms_vault_id")
	famsUserId := utils.GetEnvSettingWithBlankDefault("fams_user_id")
	famsUserPassword := utils.GetEnvSettingWithBlankDefault("fams_user_password")
	compatibleProduct := utils.GetEnvSettingWithBlankDefault("compatible_product")

	tenancyIdVariableStr := fmt.Sprintf("variable \"tenancy_id\" { default = \"%s\" }\n", tenancyId)
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	//TODO: FIX APP
	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_create", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)
	instanceIdVariableStr := fmt.Sprintf("variable \"test_instance_id\" { default = \"%s\" }\n", instanceId)
	keyIdVariableStr := fmt.Sprintf("variable \"key_id\" { default = \"%s\" }\n", keyId)
	vaultIdVariableStr := fmt.Sprintf("variable \"vault_id\" { default = \"%s\" }\n", vaultId)
	onsNotificationTopicIdVariableStr := fmt.Sprintf("variable \"oci_ons_notification_topic\" { default = \"%s\" }\n", onsNotificationTopicId)
	kmsVaultIdVariableStr := fmt.Sprintf("variable \"oci_kms_vault\" { default = \"%s\" }\n", kmsVaultId)
	famsUserIdVariableStr := fmt.Sprintf("variable \"fams_user_id\" { default = \"%s\" }\n", famsUserId)
	famsUserPasswordVariableStr := fmt.Sprintf("variable \"fams_user_password\" { default = \"%s\" }\n", famsUserPassword)
	compatibleProductVariableStr := fmt.Sprintf("variable \"compatible_product\" { default = \"%s\" }\n", compatibleProduct)

	resourceName := "oci_fleet_apps_management_fleet.test_fleet"
	datasourceName := "data.oci_fleet_apps_management_fleets.test_fleets"
	singularDatasourceName := "data.oci_fleet_apps_management_fleet.test_fleet"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+FleetAppsManagementFleetResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_fleet_apps_management_fleet", "test_fleet", acctest.Optional, acctest.Create, FleetAppsManagementFleetRepresentation), "fleetappsmanagement", "fleet", t)

	acctest.ResourceTest(t, testAccCheckFleetAppsManagementFleetDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + tenancyIdVariableStr + instanceIdVariableStr + keyIdVariableStr + vaultIdVariableStr + FleetAppsManagementFleetResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_fleet_apps_management_fleet", "test_fleet", acctest.Required, acctest.Create, FleetAppsManagementFleetRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttrSet(resourceName, "details.0.fleet_type"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "is_target_auto_confirm"),
				resource.TestCheckResourceAttrSet(resourceName, "resource_selection.0.resource_selection_type"),
				resource.TestCheckResourceAttr(resourceName, "state", "ACTIVE"),
				resource.TestCheckResourceAttr(resourceName, "resource_selection.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "resource_selection.0.rule_selection_criteria.#"),
				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},
		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + instanceIdVariableStr + keyIdVariableStr + vaultIdVariableStr + FleetAppsManagementFleetResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + tenancyIdVariableStr + instanceIdVariableStr + compatibleProductVariableStr + keyIdVariableStr + vaultIdVariableStr + onsNotificationTopicIdVariableStr + kmsVaultIdVariableStr + famsUserIdVariableStr + famsUserPasswordVariableStr + FleetAppsManagementFleetResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_fleet_apps_management_fleet", "test_fleet", acctest.Optional, acctest.Create, FleetAppsManagementFleetRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName, "credentials.#"),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "details.0.fleet_type", "GENERIC"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "environment_type", "Stage"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_target_auto_confirm", "false"),
				resource.TestCheckResourceAttr(resourceName, "notification_preferences.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "notification_preferences.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "notification_preferences.0.preferences.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "notification_preferences.0.preferences.0.on_job_canceled", "false"),
				resource.TestCheckResourceAttr(resourceName, "notification_preferences.0.preferences.0.on_job_failure", "false"),
				resource.TestCheckResourceAttr(resourceName, "notification_preferences.0.preferences.0.on_job_schedule_change", "false"),
				resource.TestCheckResourceAttr(resourceName, "notification_preferences.0.preferences.0.on_job_start", "false"),
				resource.TestCheckResourceAttr(resourceName, "notification_preferences.0.preferences.0.on_job_success", "false"),
				resource.TestCheckResourceAttr(resourceName, "notification_preferences.0.preferences.0.on_resource_non_compliance", "false"),
				resource.TestCheckResourceAttr(resourceName, "notification_preferences.0.preferences.0.on_runbook_newer_version", "false"),
				resource.TestCheckResourceAttr(resourceName, "notification_preferences.0.preferences.0.on_task_failure", "false"),
				resource.TestCheckResourceAttr(resourceName, "notification_preferences.0.preferences.0.on_task_pause", "false"),
				resource.TestCheckResourceAttr(resourceName, "notification_preferences.0.preferences.0.on_task_success", "false"),
				resource.TestCheckResourceAttr(resourceName, "notification_preferences.0.preferences.0.on_topology_modification", "false"),
				resource.TestCheckResourceAttr(resourceName, "notification_preferences.0.preferences.0.upcoming_schedule.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "notification_preferences.0.preferences.0.upcoming_schedule.0.notify_before", "notifyBefore"),
				resource.TestCheckResourceAttr(resourceName, "notification_preferences.0.preferences.0.upcoming_schedule.0.on_upcoming_schedule", "false"),
				resource.TestCheckResourceAttrSet(resourceName, "notification_preferences.0.topic_id"),
				resource.TestCheckResourceAttr(resourceName, "resource_selection.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "resource_selection.0.resource_selection_type", "DYNAMIC"),
				resource.TestCheckResourceAttr(resourceName, "resource_selection.0.rule_selection_criteria.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "resource_selection.0.rule_selection_criteria.0.match_condition", "ANY"),
				resource.TestCheckResourceAttr(resourceName, "resource_selection.0.rule_selection_criteria.0.rules.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "resource_selection.0.rule_selection_criteria.0.rules.0.basis", "inventoryProperties"),
				resource.TestCheckResourceAttrSet(resourceName, "resource_selection.0.rule_selection_criteria.0.rules.0.compartment_id"),
				resource.TestCheckResourceAttr(resourceName, "resource_selection.0.rule_selection_criteria.0.rules.0.compartment_id_in_subtree", "false"),
				resource.TestCheckResourceAttr(resourceName, "resource_selection.0.rule_selection_criteria.0.rules.0.conditions.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "resource_selection.0.rule_selection_criteria.0.rules.0.conditions.0.attr_group", "Instance"),
				resource.TestCheckResourceAttr(resourceName, "resource_selection.0.rule_selection_criteria.0.rules.0.conditions.0.attr_key", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "resource_selection.0.rule_selection_criteria.0.rules.0.conditions.0.attr_value", "attrValue1"),
				resource.TestCheckResourceAttr(resourceName, "resource_selection.0.rule_selection_criteria.0.rules.0.match_condition", "MATCH_ALL"),
				resource.TestCheckResourceAttrSet(resourceName, "resource_selection.0.rule_selection_criteria.0.rules.0.resource_compartment_id"),
				resource.TestCheckResourceAttrSet(resourceName, "resources.#"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),

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
		// verify Update to the compartment (the compartment will be switched back in the next step)
		{
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + tenancyIdVariableStr + instanceIdVariableStr + compatibleProductVariableStr + keyIdVariableStr + vaultIdVariableStr + onsNotificationTopicIdVariableStr + kmsVaultIdVariableStr + famsUserIdVariableStr + famsUserPasswordVariableStr + FleetAppsManagementFleetResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_fleet_apps_management_fleet", "test_fleet", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(FleetAppsManagementFleetRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName, "credentials.#"),
				// Credentials are not part of API Response. They are async calls
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "details.0.fleet_type", "GENERIC"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "environment_type", "Stage"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_target_auto_confirm", "false"),
				resource.TestCheckResourceAttr(resourceName, "notification_preferences.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "notification_preferences.0.compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "notification_preferences.0.preferences.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "notification_preferences.0.preferences.0.on_job_canceled", "false"),
				resource.TestCheckResourceAttr(resourceName, "notification_preferences.0.preferences.0.on_job_failure", "false"),
				resource.TestCheckResourceAttr(resourceName, "notification_preferences.0.preferences.0.on_job_schedule_change", "false"),
				resource.TestCheckResourceAttr(resourceName, "notification_preferences.0.preferences.0.on_job_start", "false"),
				resource.TestCheckResourceAttr(resourceName, "notification_preferences.0.preferences.0.on_job_success", "false"),
				resource.TestCheckResourceAttr(resourceName, "notification_preferences.0.preferences.0.on_resource_non_compliance", "false"),
				resource.TestCheckResourceAttr(resourceName, "notification_preferences.0.preferences.0.on_runbook_newer_version", "false"),
				resource.TestCheckResourceAttr(resourceName, "notification_preferences.0.preferences.0.on_task_failure", "false"),
				resource.TestCheckResourceAttr(resourceName, "notification_preferences.0.preferences.0.on_task_pause", "false"),
				resource.TestCheckResourceAttr(resourceName, "notification_preferences.0.preferences.0.on_task_success", "false"),
				resource.TestCheckResourceAttr(resourceName, "notification_preferences.0.preferences.0.on_topology_modification", "false"),
				resource.TestCheckResourceAttr(resourceName, "notification_preferences.0.preferences.0.upcoming_schedule.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "notification_preferences.0.preferences.0.upcoming_schedule.0.notify_before", "notifyBefore"),
				resource.TestCheckResourceAttr(resourceName, "notification_preferences.0.preferences.0.upcoming_schedule.0.on_upcoming_schedule", "false"),
				resource.TestCheckResourceAttrSet(resourceName, "notification_preferences.0.topic_id"),
				resource.TestCheckResourceAttrSet(resourceName, "products.#"),
				resource.TestCheckResourceAttrSet(resourceName, "properties.#"),
				resource.TestCheckResourceAttr(resourceName, "resource_selection.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "resource_selection.0.resource_selection_type", "DYNAMIC"),
				resource.TestCheckResourceAttr(resourceName, "resource_selection.0.rule_selection_criteria.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "resource_selection.0.rule_selection_criteria.0.rules.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "resource_selection.0.rule_selection_criteria.0.rules.0.basis", "inventoryProperties"),
				resource.TestCheckResourceAttr(resourceName, "resource_selection.0.rule_selection_criteria.0.rules.0.compartment_id_in_subtree", "false"),
				resource.TestCheckResourceAttr(resourceName, "resource_selection.0.rule_selection_criteria.0.rules.0.conditions.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "resource_selection.0.rule_selection_criteria.0.rules.0.conditions.0.attr_group", "Instance"),
				resource.TestCheckResourceAttr(resourceName, "resource_selection.0.rule_selection_criteria.0.rules.0.conditions.0.attr_key", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "resource_selection.0.rule_selection_criteria.0.rules.0.conditions.0.attr_value", "attrValue1"),
				resource.TestCheckResourceAttr(resourceName, "resource_selection.0.rule_selection_criteria.0.rules.0.match_condition", "MATCH_ALL"),
				resource.TestCheckResourceAttrSet(resourceName, "resource_selection.0.rule_selection_criteria.0.rules.0.resource_compartment_id"),
				resource.TestCheckResourceAttrSet(resourceName, "resources.#"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("resource recreated when it was supposed to be updated")
					}
					return err
				},
			),
		},
		// verify updates to updatable parameters
		{
			Config: config + compartmentIdVariableStr + tenancyIdVariableStr + instanceIdVariableStr + keyIdVariableStr + vaultIdVariableStr + onsNotificationTopicIdVariableStr + kmsVaultIdVariableStr + famsUserIdVariableStr + famsUserPasswordVariableStr + FleetAppsManagementFleetResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_fleet_apps_management_fleet", "test_fleet", acctest.Optional, acctest.Update, FleetAppsManagementFleetRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName, "credentials.#"),
				resource.TestCheckResourceAttr(resourceName, "description", "description2"),
				resource.TestCheckResourceAttr(resourceName, "details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "details.0.fleet_type", "GENERIC"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttrSet(resourceName, "environment_type"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_target_auto_confirm", "true"),
				resource.TestCheckResourceAttr(resourceName, "notification_preferences.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "notification_preferences.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "notification_preferences.0.preferences.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "notification_preferences.0.preferences.0.on_job_canceled", "true"),
				resource.TestCheckResourceAttr(resourceName, "notification_preferences.0.preferences.0.on_job_failure", "true"),
				resource.TestCheckResourceAttr(resourceName, "notification_preferences.0.preferences.0.on_job_schedule_change", "true"),
				resource.TestCheckResourceAttr(resourceName, "notification_preferences.0.preferences.0.on_job_start", "true"),
				resource.TestCheckResourceAttr(resourceName, "notification_preferences.0.preferences.0.on_job_success", "true"),
				resource.TestCheckResourceAttr(resourceName, "notification_preferences.0.preferences.0.on_resource_non_compliance", "true"),
				resource.TestCheckResourceAttr(resourceName, "notification_preferences.0.preferences.0.on_runbook_newer_version", "true"),
				resource.TestCheckResourceAttr(resourceName, "notification_preferences.0.preferences.0.on_task_failure", "true"),
				resource.TestCheckResourceAttr(resourceName, "notification_preferences.0.preferences.0.on_task_pause", "true"),
				resource.TestCheckResourceAttr(resourceName, "notification_preferences.0.preferences.0.on_task_success", "true"),
				resource.TestCheckResourceAttr(resourceName, "notification_preferences.0.preferences.0.on_topology_modification", "true"),
				resource.TestCheckResourceAttr(resourceName, "notification_preferences.0.preferences.0.upcoming_schedule.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "notification_preferences.0.preferences.0.upcoming_schedule.0.notify_before", "notifyBefore2"),
				resource.TestCheckResourceAttr(resourceName, "notification_preferences.0.preferences.0.upcoming_schedule.0.on_upcoming_schedule", "true"),
				resource.TestCheckResourceAttrSet(resourceName, "notification_preferences.0.topic_id"),
				resource.TestCheckResourceAttrSet(resourceName, "products.#"),
				resource.TestCheckResourceAttrSet(resourceName, "properties.#"),
				resource.TestCheckResourceAttr(resourceName, "resource_selection.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "resource_selection.0.resource_selection_type", "DYNAMIC"),
				resource.TestCheckResourceAttr(resourceName, "resource_selection.0.rule_selection_criteria.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "resource_selection.0.rule_selection_criteria.0.rules.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "resource_selection.0.rule_selection_criteria.0.rules.0.basis", "inventoryProperties"),
				resource.TestCheckResourceAttr(resourceName, "resource_selection.0.rule_selection_criteria.0.rules.0.compartment_id_in_subtree", "true"),
				resource.TestCheckResourceAttr(resourceName, "resource_selection.0.rule_selection_criteria.0.rules.0.conditions.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "resource_selection.0.rule_selection_criteria.0.rules.0.conditions.0.attr_group", "Instance"),
				resource.TestCheckResourceAttr(resourceName, "resource_selection.0.rule_selection_criteria.0.rules.0.conditions.0.attr_key", "shape"),
				resource.TestCheckResourceAttr(resourceName, "resource_selection.0.rule_selection_criteria.0.rules.0.conditions.0.attr_value", "VM.Standard.E4.Flex"),
				resource.TestCheckResourceAttr(resourceName, "resource_selection.0.rule_selection_criteria.0.rules.0.match_condition", "ANY"),
				resource.TestCheckResourceAttrSet(resourceName, "resource_selection.0.rule_selection_criteria.0.rules.0.resource_compartment_id"),
				resource.TestCheckResourceAttrSet(resourceName, "resources.#"),
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
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_fleet_apps_management_fleets", "test_fleets", acctest.Optional, acctest.Update, FleetAppsManagementFleetDataSourceRepresentation) +
				compartmentIdVariableStr + tenancyIdVariableStr + instanceIdVariableStr + keyIdVariableStr + vaultIdVariableStr + onsNotificationTopicIdVariableStr + kmsVaultIdVariableStr + famsUserIdVariableStr + famsUserPasswordVariableStr + FleetAppsManagementFleetResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_fleet_apps_management_fleet", "test_fleet", acctest.Optional, acctest.Update, FleetAppsManagementFleetRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "application_type", "applicationType"),
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttrSet(datasourceName, "environment_type"),
				resource.TestCheckResourceAttr(datasourceName, "fleet_type", "GENERIC"),
				resource.TestCheckResourceAttrSet(datasourceName, "id"),
				resource.TestCheckResourceAttr(datasourceName, "state", "NEEDS_ATTENTION"),
				resource.TestCheckResourceAttr(datasourceName, "fleet_collection.#", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "fleet_collection.0.items.#"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_fleet_apps_management_fleet", "test_fleet", acctest.Required, acctest.Create, FleetAppsManagementFleetSingularDataSourceRepresentation) +
				compartmentIdVariableStr + tenancyIdVariableStr + instanceIdVariableStr + keyIdVariableStr + vaultIdVariableStr + onsNotificationTopicIdVariableStr + kmsVaultIdVariableStr + famsUserIdVariableStr + famsUserPasswordVariableStr + FleetAppsManagementFleetResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "fleet_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "description", "description2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "details.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "details.0.fleet_type", "GENERIC"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "environment_type"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "is_target_auto_confirm", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "notification_preferences.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "notification_preferences.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "notification_preferences.0.preferences.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "notification_preferences.0.preferences.0.on_job_canceled", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "notification_preferences.0.preferences.0.on_job_failure", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "notification_preferences.0.preferences.0.on_job_schedule_change", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "notification_preferences.0.preferences.0.on_job_start", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "notification_preferences.0.preferences.0.on_job_success", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "notification_preferences.0.preferences.0.on_resource_non_compliance", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "notification_preferences.0.preferences.0.on_runbook_newer_version", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "notification_preferences.0.preferences.0.on_task_failure", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "notification_preferences.0.preferences.0.on_task_pause", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "notification_preferences.0.preferences.0.on_task_success", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "notification_preferences.0.preferences.0.on_topology_modification", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "notification_preferences.0.preferences.0.upcoming_schedule.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "notification_preferences.0.preferences.0.upcoming_schedule.0.notify_before", "notifyBefore2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "notification_preferences.0.preferences.0.upcoming_schedule.0.on_upcoming_schedule", "true"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "products.#"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "resource_region"),
				resource.TestCheckResourceAttr(singularDatasourceName, "resource_selection.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "resource_selection.0.resource_selection_type", "DYNAMIC"),
				resource.TestCheckResourceAttr(singularDatasourceName, "resource_selection.0.rule_selection_criteria.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "resource_selection.0.rule_selection_criteria.0.match_condition", "MATCH_ALL"),
				resource.TestCheckResourceAttr(singularDatasourceName, "resource_selection.0.rule_selection_criteria.0.rules.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "resource_selection.0.rule_selection_criteria.0.rules.0.basis", "inventoryProperties"),
				resource.TestCheckResourceAttr(singularDatasourceName, "resource_selection.0.rule_selection_criteria.0.rules.0.compartment_id_in_subtree", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "resource_selection.0.rule_selection_criteria.0.rules.0.conditions.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "resource_selection.0.rule_selection_criteria.0.rules.0.conditions.0.attr_group", "Instance"),
				resource.TestCheckResourceAttr(singularDatasourceName, "resource_selection.0.rule_selection_criteria.0.rules.0.conditions.0.attr_key", "shape"),
				resource.TestCheckResourceAttr(singularDatasourceName, "resource_selection.0.rule_selection_criteria.0.rules.0.conditions.0.attr_value", "VM.Standard.E4.Flex"),
				resource.TestCheckResourceAttr(singularDatasourceName, "resource_selection.0.rule_selection_criteria.0.rules.0.match_condition", "ANY"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "resources.#"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
		// verify resource import
		{
			Config:                  config + FleetAppsManagementFleetRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       false,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckFleetAppsManagementFleetDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).FleetAppsManagementClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_fleet_apps_management_fleet" {
			noResourceFound = false
			request := oci_fleet_apps_management.GetFleetRequest{}

			tmp := rs.Primary.ID
			request.FleetId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "fleet_apps_management")

			response, err := client.GetFleet(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_fleet_apps_management.FleetLifecycleStateDeleted): true,
				}
				if _, ok := deletedLifecycleStates[string(response.LifecycleState)]; !ok {
					//resource lifecycle state is not in expected deleted lifecycle states.
					return fmt.Errorf("resource lifecycle state: %s is not in expected deleted lifecycle states", response.LifecycleState)
				}
				//resource lifecycle state is in expected deleted lifecycle states. continue with next one.
				continue
			}

			//Verify that exception is for '404 not found'.
			if failure, isServiceError := common.IsServiceError(err); !isServiceError || failure.GetHTTPStatusCode() != 404 {
				return err
			}
		}
	}
	if noResourceFound {
		return fmt.Errorf("at least one resource was expected from the state file, but could not be found")
	}

	return nil
}

func init() {
	if acctest.DependencyGraph == nil {
		acctest.InitDependencyGraph()
	}
	if !acctest.InSweeperExcludeList("FleetAppsManagementFleet") {
		resource.AddTestSweepers("FleetAppsManagementFleet", &resource.Sweeper{
			Name:         "FleetAppsManagementFleet",
			Dependencies: acctest.DependencyGraph["fleet"],
			F:            sweepFleetAppsManagementFleetResource,
		})
	}
}

func sweepFleetAppsManagementFleetResource(compartment string) error {
	fleetAppsManagementClient := acctest.GetTestClients(&schema.ResourceData{}).FleetAppsManagementClient()
	fleetIds, err := getFleetAppsManagementFleetIds(compartment)
	if err != nil {
		return err
	}
	for _, fleetId := range fleetIds {
		if ok := acctest.SweeperDefaultResourceId[fleetId]; !ok {
			deleteFleetRequest := oci_fleet_apps_management.DeleteFleetRequest{}

			deleteFleetRequest.FleetId = &fleetId

			deleteFleetRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "fleet_apps_management")
			_, error := fleetAppsManagementClient.DeleteFleet(context.Background(), deleteFleetRequest)
			if error != nil {
				fmt.Printf("Error deleting Fleet %s %s, It is possible that the resource is already deleted. Please verify manually \n", fleetId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &fleetId, FleetAppsManagementFleetSweepWaitCondition, time.Duration(3*time.Minute),
				FleetAppsManagementFleetSweepResponseFetchOperation, "fleet_apps_management", true)
		}
	}
	return nil
}

func getFleetAppsManagementFleetIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "FleetId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	fleetAppsManagementClient := acctest.GetTestClients(&schema.ResourceData{}).FleetAppsManagementClient()

	listFleetsRequest := oci_fleet_apps_management.ListFleetsRequest{}
	listFleetsRequest.CompartmentId = &compartmentId
	listFleetsRequest.LifecycleState = oci_fleet_apps_management.FleetLifecycleStateActive
	listFleetsResponse, err := fleetAppsManagementClient.ListFleets(context.Background(), listFleetsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting Fleet list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, fleet := range listFleetsResponse.Items {
		id := *fleet.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "FleetId", id)
	}
	return resourceIds, nil
}

func FleetAppsManagementFleetSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if fleetResponse, ok := response.Response.(oci_fleet_apps_management.GetFleetResponse); ok {
		return fleetResponse.LifecycleState != oci_fleet_apps_management.FleetLifecycleStateDeleted
	}
	return false
}

func FleetAppsManagementFleetSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.FleetAppsManagementClient().GetFleet(context.Background(), oci_fleet_apps_management.GetFleetRequest{
		FleetId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
