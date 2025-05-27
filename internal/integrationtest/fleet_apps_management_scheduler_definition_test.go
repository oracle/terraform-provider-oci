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
	FleetAppsManagementSchedulerDefinitionRequiredOnlyResource = FleetAppsManagementSchedulerDefinitionResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_fleet_apps_management_scheduler_definition", "test_scheduler_definition", acctest.Required, acctest.Create, FleetAppsManagementSchedulerDefinitionRepresentation)

	FleetAppsManagementSchedulerDefinitionResourceConfig = FleetAppsManagementSchedulerDefinitionResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_fleet_apps_management_scheduler_definition", "test_scheduler_definition", acctest.Optional, acctest.Update, FleetAppsManagementSchedulerDefinitionRepresentation)

	FleetAppsManagementSchedulerDefinitionSingularDataSourceRepresentation = map[string]interface{}{
		"scheduler_definition_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_fleet_apps_management_scheduler_definition.test_scheduler_definition.id}`},
	}

	FleetAppsManagementSchedulerDefinitionDataSourceRepresentation = map[string]interface{}{
		"compartment_id":        acctest.Representation{RepType: acctest.Optional, Create: `${var.compartment_id}`},
		"display_name":          acctest.Representation{RepType: acctest.Optional, Create: `TerraProviderTestSchedDef`, Update: `TerraProviderTestSchedDef2`},
		"fleet_id":              acctest.Representation{RepType: acctest.Optional, Create: `${var.test_active_fleet}`},
		"maintenance_window_id": acctest.Representation{RepType: acctest.Optional, Create: `${var.maintenance_window}`},

		//"product":               acctest.Representation{RepType: acctest.Optional, Create: `product`},
		//"runbook_id":            acctest.Representation{RepType: acctest.Optional, Create: `${oci_fleet_apps_management_runbook.test_runbook.id}`},
		//"runbook_version_name":  acctest.Representation{RepType: acctest.Optional, Create: `runbookVersionName`},
		//"time_scheduled_greater_than_or_equal_to": acctest.Representation{RepType: acctest.Optional, Create: `timeScheduledGreaterThanOrEqualTo`},
		//"time_scheduled_less_than":                acctest.Representation{RepType: acctest.Optional, Create: `timeScheduledLessThan`},

		"state":  acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"filter": acctest.RepresentationGroup{RepType: acctest.Required, Group: FleetAppsManagementSchedulerDefinitionDataSourceFilterRepresentation},
	}

	FleetAppsManagementSchedulerDefinitionDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_fleet_apps_management_scheduler_definition.test_scheduler_definition.id}`}},
	}

	FleetAppsManagementSchedulerDefinitionRepresentation = map[string]interface{}{
		"action_groups":  acctest.RepresentationGroup{RepType: acctest.Required, Group: FleetAppsManagementSchedulerDefinitionActionGroupsRepresentation},
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"schedule":       acctest.RepresentationGroup{RepType: acctest.Required, Group: FleetAppsManagementSchedulerDefinitionScheduleRepresentation},
		//"activity_initiation_cut_off": acctest.Representation{RepType: acctest.Required, Create: `1`, Update: `2`},
		// "defined_tags": acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"description": acctest.Representation{RepType: acctest.Optional, Create: `description`, Update: `description2`},
		//TODO FIXAPP "display_name":  acctest.Representation{RepType: acctest.Optional, Create: `TerraProviderTestSchedDef`, Update: `TerraProviderTestSchedDef2`},
		"display_name":  acctest.Representation{RepType: acctest.Required, Create: `TerraProviderTestSchedDef`, Update: `TerraProviderTestSchedDef2`},
		"freeform_tags": acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"bar-key": "value"}},
		"run_books":     acctest.RepresentationGroup{RepType: acctest.Required, Group: FleetAppsManagementSchedulerDefinitionRunBooksRepresentation},
		//"run_books":      acctest.RepresentationGroup{RepType: acctest.Optional, Group: FleetAppsManagementSchedulerDefinitionRunBooksRepresentation},
	}

	FleetAppsManagementSchedulerDefinitionActionGroupsRepresentation = map[string]interface{}{
		// //"fleet_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_fleet_apps_management_fleet.test_fleet.id}`},
		"fleet_id":             acctest.Representation{RepType: acctest.Required, Create: `${var.test_active_fleet}`},
		"kind":                 acctest.Representation{RepType: acctest.Required, Create: `FLEET_USING_RUNBOOK`},
		"runbook_id":           acctest.Representation{RepType: acctest.Required, Create: `${var.test_runbook_ocid}`},
		"runbook_version_name": acctest.Representation{RepType: acctest.Required, Create: `1`},
		// "display_name":         acctest.Representation{RepType: acctest.Optional, Create: `TerraProviderTestSchedDef`, Update: `TerraProviderTestSchedDef2`},
		"sequence": acctest.Representation{RepType: acctest.Optional, Create: `10`, Update: `11`},
		//"lifecycle_operation": acctest.Representation{RepType: acctest.Required, Create: `PATCH`},
		//"product":             acctest.Representation{RepType: acctest.Required, Create: `Oracle Linux`},
		//"type":                acctest.Representation{RepType: acctest.Required, Create: `PRODUCT`},
	}

	FleetAppsManagementSchedulerDefinitionScheduleRepresentation = map[string]interface{}{
		"execution_startdate": acctest.Representation{RepType: acctest.Required, Create: `2025-06-01T00:00:00.000Z`},
		"type":                acctest.Representation{RepType: acctest.Required, Create: `MAINTENANCE_WINDOW`},
		// "duration":              acctest.Representation{RepType: acctest.Required, Create: `PT2H`},
		"maintenance_window_id": acctest.Representation{RepType: acctest.Required, Create: `${var.maintenance_window}`},
	}
	FleetAppsManagementSchedulerDefinitionRunBooksRepresentation = map[string]interface{}{
		"runbook_id":           acctest.Representation{RepType: acctest.Required, Create: `${var.test_runbook_ocid}`},
		"runbook_version_name": acctest.Representation{RepType: acctest.Required, Create: `1`},
		"input_parameters":     acctest.RepresentationGroup{RepType: acctest.Optional, Group: FleetAppsManagementSchedulerDefinitionRunBooksInputParametersRepresentation},
	}
	FleetAppsManagementSchedulerDefinitionRunBooksInputParametersRepresentation = map[string]interface{}{
		"step_name": acctest.Representation{RepType: acctest.Required, Create: `tesri_testing_task`, Update: `tesri_testing_task`},
		"arguments": acctest.RepresentationGroup{RepType: acctest.Optional, Group: FleetAppsManagementSchedulerDefinitionRunBooksInputParametersArgumentsRepresentation},
	}
	FleetAppsManagementSchedulerDefinitionRunBooksInputParametersArgumentsRepresentation = map[string]interface{}{
		"kind": acctest.Representation{RepType: acctest.Required, Create: `STRING`},
		"name": acctest.Representation{RepType: acctest.Required, Create: `name`},
		// "content": acctest.RepresentationGroup{RepType: acctest.Optional, Group: FleetAppsManagementSchedulerDefinitionRunBooksInputParametersArgumentsContentRepresentation},
		"value": acctest.Representation{RepType: acctest.Optional, Create: `value`, Update: `value2`},
	}
	FleetAppsManagementSchedulerDefinitionRunBooksInputParametersArgumentsContentRepresentation = map[string]interface{}{
		"bucket":      acctest.Representation{RepType: acctest.Required, Create: `bucket`, Update: `bucket2`},
		"checksum":    acctest.Representation{RepType: acctest.Required, Create: `checksum`, Update: `checksum2`},
		"namespace":   acctest.Representation{RepType: acctest.Required, Create: `namespace`, Update: `namespace2`},
		"object":      acctest.Representation{RepType: acctest.Required, Create: `object`, Update: `object2`},
		"source_type": acctest.Representation{RepType: acctest.Required, Create: `OBJECT_STORAGE_BUCKET`},
	}

	//FleetAppsManagementSchedulerDefinitionResourceDependencies = acctest.GenerateDataSourceFromRepresentationMap("oci_cloud_guard_resources", "test_resources", acctest.Required, acctest.Create, CloudGuardResourceDataSourceRepresentation) +
	//	acctest.GenerateResourceFromRepresentationMap("oci_core_subnet", "test_subnet", acctest.Required, acctest.Create, CoreSubnetRepresentation) +
	//	acctest.GenerateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", acctest.Required, acctest.Create, CoreVcnRepresentation) +
	//	acctest.GenerateResourceFromRepresentationMap("oci_datacatalog_catalog", "test_catalog", acctest.Required, acctest.Create, DatacatalogCatalogRepresentation) +
	//	acctest.GenerateResourceFromRepresentationMap("oci_dataflow_private_endpoint", "test_private_endpoint", acctest.Required, acctest.Create, DataflowPrivateEndpointRepresentation) +
	//	acctest.GenerateResourceFromRepresentationMap("oci_fleet_apps_management_fleet", "test_fleet", acctest.Required, acctest.Create, FleetAppsManagementFleetRepresentation) +
	//	acctest.GenerateResourceFromRepresentationMap("oci_fleet_apps_management_maintenance_window", "test_maintenance_window", acctest.Required, acctest.Create, FleetAppsManagementMaintenanceWindowRepresentation) +
	//	acctest.GenerateResourceFromRepresentationMap("oci_fleet_apps_management_runbook", "test_runbook", acctest.Required, acctest.Create, FleetAppsManagementRunbookRepresentation) +
	//	acctest.GenerateResourceFromRepresentationMap("oci_fleet_apps_management_task_record", "test_task_record", acctest.Required, acctest.Create, FleetAppsManagementTaskRecordRepresentation) +
	//	acctest.GenerateResourceFromRepresentationMap("oci_functions_application", "test_application", acctest.Required, acctest.Create, FunctionsApplicationRepresentation) +
	//	acctest.GenerateResourceFromRepresentationMap("oci_functions_function", "test_function", acctest.Required, acctest.Create, FunctionsFunctionRepresentation) +
	//	acctest.GenerateDataSourceFromRepresentationMap("oci_functions_pbf_listings", "test_pbf_listings", acctest.Required, acctest.Create, FunctionsPbfListingDataSourceRepresentation) +
	//	DefinedTagsDependencies +
	//	acctest.GenerateResourceFromRepresentationMap("oci_identity_group", "test_group", acctest.Required, acctest.Create, IdentityGroupRepresentation) +
	//	KeyResourceDependencyConfig +
	//	acctest.GenerateResourceFromRepresentationMap("oci_kms_vault", "test_vault", acctest.Required, acctest.Create, KmsVaultRepresentation) +
	//	acctest.GenerateResourceFromRepresentationMap("oci_objectstorage_bucket", "test_bucket", acctest.Required, acctest.Create, ObjectStorageBucketRepresentation) +
	//	GenerateDataSourceFromRepresentationMap("oci_objectstorage_namespace", "test_namespace", Required, Create, namespaceSingularDataSourceRepresentation) +
	//	acctest.GenerateResourceFromRepresentationMap("oci_ons_notification_topic", "test_notification_topic", acctest.Required, acctest.Create, OnsNotificationTopicRepresentation) +
	//	acctest.GenerateResourceFromRepresentationMap("oci_vault_secret", "test_secret", acctest.Required, acctest.Create, VaultSecretRepresentation)

	// FleetAppsManagementSchedulerDefinitionResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_fleet_apps_management_maintenance_window", "test_maintenance_window", acctest.Required, acctest.Create, FleetAppsManagementMaintenanceWindowRepresentation)
	//TODO TEMP removed: + //DefinedTagsDependencies
	FleetAppsManagementSchedulerDefinitionResourceDependencies = ""
)

// issue-routing-tag: fleet_apps_management/default
func TestFleetAppsManagementSchedulerDefinitionResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestFleetAppsManagementSchedulerDefinitionResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	// Runbooks are currently created by Oracle, and read-only. There is no Create API.
	runbookId := utils.GetEnvSettingWithBlankDefault("test_runbook_ocid")
	testRunbookStr := fmt.Sprintf("variable \"test_runbook_ocid\" { default = \"%s\" }\n", runbookId)

	maintenanceWindowId := utils.GetEnvSettingWithBlankDefault("maintenance_window")
	maintenanceWindowIdVariableStr := fmt.Sprintf("variable \"maintenance_window\" { default = \"%s\" }\n", maintenanceWindowId)

	// Fleet in ACTIVE state. Fleets require a confirmation action call not supported by Terraform to go active.
	// Thus, this needs to be created and confirmed manually.
	activeFleetId := utils.GetEnvSettingWithBlankDefault("test_active_fleet")
	activeFleetStr := fmt.Sprintf("variable \"test_active_fleet\" { default = \"%s\" }\n", activeFleetId)

	resourceName := "oci_fleet_apps_management_scheduler_definition.test_scheduler_definition"
	datasourceName := "data.oci_fleet_apps_management_scheduler_definitions.test_scheduler_definitions"
	singularDatasourceName := "data.oci_fleet_apps_management_scheduler_definition.test_scheduler_definition"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+FleetAppsManagementSchedulerDefinitionResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_fleet_apps_management_scheduler_definition", "test_scheduler_definition", acctest.Optional, acctest.Create, FleetAppsManagementSchedulerDefinitionRepresentation), "fleetappsmanagement", "schedulerDefinition", t)

	acctest.ResourceTest(t, testAccCheckFleetAppsManagementSchedulerDefinitionDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + activeFleetStr + testRunbookStr + compartmentIdVariableStr + maintenanceWindowIdVariableStr + FleetAppsManagementSchedulerDefinitionResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_fleet_apps_management_scheduler_definition", "test_scheduler_definition", acctest.Required, acctest.Create, FleetAppsManagementSchedulerDefinitionRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "action_groups.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "action_groups.0.fleet_id"),
				resource.TestCheckResourceAttr(resourceName, "action_groups.0.kind", "FLEET_USING_RUNBOOK"),
				resource.TestCheckResourceAttrSet(resourceName, "action_groups.0.runbook_id"),
				resource.TestCheckResourceAttr(resourceName, "action_groups.0.runbook_version_name", "1"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "schedule.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "schedule.0.execution_startdate"),
				resource.TestCheckResourceAttrSet(resourceName, "schedule.0.maintenance_window_id"),
				resource.TestCheckResourceAttr(resourceName, "schedule.0.type", "MAINTENANCE_WINDOW"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + activeFleetStr + testRunbookStr + compartmentIdVariableStr + FleetAppsManagementSchedulerDefinitionResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + activeFleetStr + testRunbookStr + compartmentIdVariableStr + maintenanceWindowIdVariableStr + FleetAppsManagementSchedulerDefinitionResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_fleet_apps_management_scheduler_definition", "test_scheduler_definition", acctest.Optional, acctest.Create, FleetAppsManagementSchedulerDefinitionRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "action_groups.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "action_groups.0.display_name"),
				resource.TestCheckResourceAttrSet(resourceName, "action_groups.0.fleet_id"),
				resource.TestCheckResourceAttr(resourceName, "action_groups.0.kind", "FLEET_USING_RUNBOOK"),
				resource.TestCheckResourceAttrSet(resourceName, "action_groups.0.runbook_id"),
				resource.TestCheckResourceAttr(resourceName, "action_groups.0.runbook_version_name", "1"),
				resource.TestCheckResourceAttr(resourceName, "action_groups.0.sequence", "10"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttrSet(resourceName, "display_name"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "run_books.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "run_books.0.input_parameters.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "run_books.0.input_parameters.0.arguments.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "run_books.0.input_parameters.0.arguments.0.content.#"),
				resource.TestCheckResourceAttr(resourceName, "run_books.0.input_parameters.0.arguments.0.kind", "STRING"),
				resource.TestCheckResourceAttrSet(resourceName, "run_books.0.input_parameters.0.arguments.0.name"),
				resource.TestCheckResourceAttrSet(resourceName, "run_books.0.input_parameters.0.arguments.0.value"),
				resource.TestCheckResourceAttrSet(resourceName, "run_books.0.input_parameters.0.step_name"),
				resource.TestCheckResourceAttrSet(resourceName, "run_books.0.runbook_id"),
				resource.TestCheckResourceAttr(resourceName, "run_books.0.runbook_version_name", "1"),
				resource.TestCheckResourceAttr(resourceName, "schedule.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "schedule.0.type", "MAINTENANCE_WINDOW"),
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

		// verify updates to updatable parameters
		{
			Config: config + activeFleetStr + testRunbookStr + compartmentIdVariableStr + maintenanceWindowIdVariableStr + FleetAppsManagementSchedulerDefinitionResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_fleet_apps_management_scheduler_definition", "test_scheduler_definition", acctest.Optional, acctest.Update, FleetAppsManagementSchedulerDefinitionRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				// printResourceStateToFile(resourceName, "/tmp/tersi/UpdateResourceName.json"),
				// SleepFor(10*time.Minute),
				resource.TestCheckResourceAttr(resourceName, "action_groups.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "action_groups.0.display_name"),
				resource.TestCheckResourceAttrSet(resourceName, "action_groups.0.fleet_id"),
				resource.TestCheckResourceAttr(resourceName, "action_groups.0.kind", "FLEET_USING_RUNBOOK"),
				resource.TestCheckResourceAttrSet(resourceName, "action_groups.0.runbook_id"),
				resource.TestCheckResourceAttr(resourceName, "action_groups.0.runbook_version_name", "1"),
				resource.TestCheckResourceAttr(resourceName, "action_groups.0.sequence", "11"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "description", "description2"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "TerraProviderTestSchedDef2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "run_books.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "run_books.0.input_parameters.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "run_books.0.input_parameters.0.arguments.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "run_books.0.input_parameters.0.arguments.0.content.#"),
				resource.TestCheckResourceAttr(resourceName, "run_books.0.input_parameters.0.arguments.0.kind", "STRING"),
				resource.TestCheckResourceAttr(resourceName, "run_books.0.input_parameters.0.arguments.0.name", "name"),
				resource.TestCheckResourceAttr(resourceName, "run_books.0.input_parameters.0.arguments.0.value", "value2"),
				resource.TestCheckResourceAttrSet(resourceName, "run_books.0.input_parameters.0.step_name"),
				resource.TestCheckResourceAttrSet(resourceName, "run_books.0.runbook_id"),
				resource.TestCheckResourceAttrSet(resourceName, "run_books.0.runbook_version_name"),
				resource.TestCheckResourceAttr(resourceName, "schedule.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "schedule.0.maintenance_window_id"),
				resource.TestCheckResourceAttr(resourceName, "schedule.0.type", "MAINTENANCE_WINDOW"),
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
				acctest.GenerateDataSourceFromRepresentationMap("oci_fleet_apps_management_scheduler_definitions", "test_scheduler_definitions", acctest.Optional, acctest.Update, FleetAppsManagementSchedulerDefinitionDataSourceRepresentation) +
				activeFleetStr + compartmentIdVariableStr + maintenanceWindowIdVariableStr + testRunbookStr + FleetAppsManagementSchedulerDefinitionResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_fleet_apps_management_scheduler_definition", "test_scheduler_definition", acctest.Optional, acctest.Update, FleetAppsManagementSchedulerDefinitionRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "%"),
				resource.TestCheckResourceAttrSet(datasourceName, "compartment_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "display_name"),
				resource.TestCheckResourceAttrSet(datasourceName, "fleet_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "maintenance_window_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "scheduler_definition_collection.0.items.0.products.0"),
				resource.TestCheckResourceAttrSet(datasourceName, "scheduler_definition_collection.0.items.0.run_books.#"),
				resource.TestCheckResourceAttrSet(datasourceName, "scheduler_definition_collection.0.items.0.schedule.0.execution_startdate"),
				resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),
				resource.TestCheckResourceAttrSet(datasourceName, "scheduler_definition_collection.0.items.0.schedule.0.type"),
				resource.TestCheckResourceAttrSet(datasourceName, "scheduler_definition_collection.0.items.0.state"),
				resource.TestCheckResourceAttrSet(datasourceName, "scheduler_definition_collection.#"),
				resource.TestCheckResourceAttrSet(datasourceName, "scheduler_definition_collection.0.items.#"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_fleet_apps_management_scheduler_definition", "test_scheduler_definition", acctest.Required, acctest.Create, FleetAppsManagementSchedulerDefinitionSingularDataSourceRepresentation) +
				activeFleetStr + compartmentIdVariableStr + maintenanceWindowIdVariableStr + testRunbookStr + FleetAppsManagementSchedulerDefinitionResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "scheduler_definition_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "action_groups.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "action_groups.0.display_name"),
				resource.TestCheckResourceAttr(singularDatasourceName, "action_groups.0.kind", "FLEET_USING_RUNBOOK"),
				resource.TestCheckResourceAttr(singularDatasourceName, "action_groups.0.runbook_version_name", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "action_groups.0.sequence", "11"),

				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "count_of_affected_action_groups"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "count_of_affected_resources"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "count_of_affected_targets"),
				resource.TestCheckResourceAttr(singularDatasourceName, "description", "description2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "TerraProviderTestSchedDef2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "lifecycle_operations.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "products.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "resource_region"),
				resource.TestCheckResourceAttr(singularDatasourceName, "run_books.#", "1"),

				resource.TestCheckResourceAttr(singularDatasourceName, "run_books.0.input_parameters.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "run_books.0.input_parameters.0.arguments.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "run_books.0.input_parameters.0.arguments.0.content.#"),
				// resource.TestCheckResourceAttr(singularDatasourceName, "run_books.0.input_parameters.0.arguments.0.content.0.bucket", "bucket2"),
				// resource.TestCheckResourceAttr(singularDatasourceName, "run_books.0.input_parameters.0.arguments.0.content.0.checksum", "checksum2"),
				// resource.TestCheckResourceAttr(singularDatasourceName, "run_books.0.input_parameters.0.arguments.0.content.0.namespace", "namespace2"),
				// resource.TestCheckResourceAttr(singularDatasourceName, "run_books.0.input_parameters.0.arguments.0.content.0.object", "object2"),
				// resource.TestCheckResourceAttr(singularDatasourceName, "run_books.0.input_parameters.0.arguments.0.content.0.source_type", "OBJECT_STORAGE_BUCKET"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "run_books.0.input_parameters.0.arguments.0.kind"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "run_books.0.input_parameters.0.arguments.0.name"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "run_books.0.input_parameters.0.arguments.0.value"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "run_books.0.input_parameters.0.step_name"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "run_books.0.runbook_version_name"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "schedule.#"),
				// resource.TestCheckResourceAttr(singularDatasourceName, "schedule.0.duration", "PT3H"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "schedule.0.execution_startdate"),
				resource.TestCheckResourceAttr(singularDatasourceName, "schedule.0.type", "MAINTENANCE_WINDOW"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
		// verify resource import
		{
			Config:                  config + FleetAppsManagementSchedulerDefinitionRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckFleetAppsManagementSchedulerDefinitionDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).FleetAppsManagementOperationsClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_fleet_apps_management_scheduler_definition" {
			noResourceFound = false
			request := oci_fleet_apps_management.GetSchedulerDefinitionRequest{}

			tmp := rs.Primary.ID
			request.SchedulerDefinitionId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "fleet_apps_management")

			response, err := client.GetSchedulerDefinition(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_fleet_apps_management.SchedulerDefinitionLifecycleStateDeleted): true,
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
	if !acctest.InSweeperExcludeList("FleetAppsManagementSchedulerDefinition") {
		resource.AddTestSweepers("FleetAppsManagementSchedulerDefinition", &resource.Sweeper{
			Name:         "FleetAppsManagementSchedulerDefinition",
			Dependencies: acctest.DependencyGraph["schedulerDefinition"],
			F:            sweepFleetAppsManagementSchedulerDefinitionResource,
		})
	}
}

func sweepFleetAppsManagementSchedulerDefinitionResource(compartment string) error {
	fleetAppsManagementOperationsClient := acctest.GetTestClients(&schema.ResourceData{}).FleetAppsManagementOperationsClient()
	schedulerDefinitionIds, err := getFleetAppsManagementSchedulerDefinitionIds(compartment)
	if err != nil {
		return err
	}
	for _, schedulerDefinitionId := range schedulerDefinitionIds {
		if ok := acctest.SweeperDefaultResourceId[schedulerDefinitionId]; !ok {
			deleteSchedulerDefinitionRequest := oci_fleet_apps_management.DeleteSchedulerDefinitionRequest{}

			deleteSchedulerDefinitionRequest.SchedulerDefinitionId = &schedulerDefinitionId

			deleteSchedulerDefinitionRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "fleet_apps_management")
			_, error := fleetAppsManagementOperationsClient.DeleteSchedulerDefinition(context.Background(), deleteSchedulerDefinitionRequest)
			if error != nil {
				fmt.Printf("Error deleting SchedulerDefinition %s %s, It is possible that the resource is already deleted. Please verify manually \n", schedulerDefinitionId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &schedulerDefinitionId, FleetAppsManagementSchedulerDefinitionSweepWaitCondition, time.Duration(3*time.Minute),
				FleetAppsManagementSchedulerDefinitionSweepResponseFetchOperation, "fleet_apps_management", true)
		}
	}
	return nil
}

func getFleetAppsManagementSchedulerDefinitionIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "SchedulerDefinitionId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	fleetAppsManagementOperationsClient := acctest.GetTestClients(&schema.ResourceData{}).FleetAppsManagementOperationsClient()

	listSchedulerDefinitionsRequest := oci_fleet_apps_management.ListSchedulerDefinitionsRequest{}
	listSchedulerDefinitionsRequest.CompartmentId = &compartmentId
	listSchedulerDefinitionsRequest.LifecycleState = oci_fleet_apps_management.SchedulerDefinitionLifecycleStateActive
	listSchedulerDefinitionsResponse, err := fleetAppsManagementOperationsClient.ListSchedulerDefinitions(context.Background(), listSchedulerDefinitionsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting SchedulerDefinition list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, schedulerDefinition := range listSchedulerDefinitionsResponse.Items {
		id := *schedulerDefinition.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "SchedulerDefinitionId", id)
	}
	return resourceIds, nil
}

func FleetAppsManagementSchedulerDefinitionSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if schedulerDefinitionResponse, ok := response.Response.(oci_fleet_apps_management.GetSchedulerDefinitionResponse); ok {
		return schedulerDefinitionResponse.LifecycleState != oci_fleet_apps_management.SchedulerDefinitionLifecycleStateDeleted
	}
	return false
}

func FleetAppsManagementSchedulerDefinitionSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.FleetAppsManagementOperationsClient().GetSchedulerDefinition(context.Background(), oci_fleet_apps_management.GetSchedulerDefinitionRequest{
		SchedulerDefinitionId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}

// // Custom function to print the resource state in JSON and write it to a file
// func printResourceStateToFile(dataSourceName, filePath string) resource.TestCheckFunc {
// 	fmt.Printf("Starting the function printResourceStateToFile")
// 	return func(s *terraform.State) error {
// 		// Get the data source from the Terraform state
// 		rs, ok := s.RootModule().Resources[dataSourceName]
// 		if !ok {
// 			return fmt.Errorf("data source not found: %s", dataSourceName)
// 		}

// 		// Convert the resource attributes to JSON
// 		resourceStateJSON, err := json.MarshalIndent(rs.Primary.Attributes, "", "  ")
// 		if err != nil {
// 			return fmt.Errorf("failed to marshal resource state to JSON: %s", err)
// 		}

// 		// Print the resource state in JSON format
// 		fmt.Printf("Resource state in JSON: %s\n", resourceStateJSON)

// 		// Write the JSON to the file at /tmp/debug.json
// 		if err := ioutil.WriteFile(filePath, resourceStateJSON, 0644); err != nil {
// 			return fmt.Errorf("failed to write JSON to file: %s", err)
// 		}

// 		fmt.Printf("Resource state written to file: %s\n", filePath)

// 		return nil
// 	}
// }

// func SleepFor(duration time.Duration) resource.TestCheckFunc {
// 	return func(state *terraform.State) error {
// 		print("wait for 15 minutes")
// 		time.Sleep(duration)
// 		return nil
// 	}
// }

// func dumpFullTerraformStateToFile(state *terraform.State, filePath string) error {
// 	stateJSON, err := json.MarshalIndent(state, "", "  ")
// 	if err != nil {
// 		return err
// 	}
// 	if err := os.MkdirAll(filepath.Dir(filePath), 0755); err != nil {
// 		return err
// 	}
// 	return os.WriteFile(filePath, stateJSON, 0644)
// }
