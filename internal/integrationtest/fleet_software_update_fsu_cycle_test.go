// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"

	"strconv"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/oracle/oci-go-sdk/v65/common"
	oci_fleet_software_update "github.com/oracle/oci-go-sdk/v65/fleetsoftwareupdate"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var now time.Time = time.Now()
var scheduledStageTime string = now.AddDate(0, 0, 10).In(time.UTC).Format(time.RFC3339)
var scheduledStageTimeUpdate string = now.AddDate(0, 0, 15).In(time.UTC).Format(time.RFC3339)
var scheduledApplyTime string = now.AddDate(0, 0, 20).In(time.UTC).Format(time.RFC3339)
var scheduledApplyTimeUpdate string = now.AddDate(0, 0, 25).In(time.UTC).Format(time.RFC3339)

var (
	FleetSoftwareUpdateFsuCycleRequiredOnlyResource_DB_VersionType = FleetSoftwareUpdateFsuCycleDBResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_fleet_software_update_fsu_cycle", "test_fsu_cycle", acctest.Required, acctest.Create, FleetSoftwareUpdateFsuCycleRepresentation_DB_VersionType)
	FleetSoftwareUpdateFsuCycleResourceConfig_DB_VersionType = FleetSoftwareUpdateFsuCycleDBResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_fleet_software_update_fsu_cycle", "test_fsu_cycle", acctest.Optional, acctest.Update, FleetSoftwareUpdateFsuCycleRepresentation_DB_VersionType)

	FleetSoftwareUpdateFsuCycleRequiredOnlyResource_DB_ImageType = FleetSoftwareUpdateFsuCycleDBResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_fleet_software_update_fsu_cycle", "test_fsu_cycle", acctest.Required, acctest.Create, FleetSoftwareUpdateFsuCycleRepresentation_DB_ImageType)
	FleetSoftwareUpdateFsuCycleResourceConfig_DB_ImageType = FleetSoftwareUpdateFsuCycleDBResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_fleet_software_update_fsu_cycle", "test_fsu_cycle", acctest.Optional, acctest.Update, FleetSoftwareUpdateFsuCycleRepresentation_DB_ImageType)

	FleetSoftwareUpdateFsuCycleRequiredOnlyResource_GI_ImageType = FleetSoftwareUpdateFsuCycleGIResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_fleet_software_update_fsu_cycle", "test_fsu_cycle", acctest.Required, acctest.Create, FleetSoftwareUpdateFsuCycleRepresentation_GI_ImageType)
	FleetSoftwareUpdateFsuCycleResourceConfig_GI_ImageType = FleetSoftwareUpdateFsuCycleGIResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_fleet_software_update_fsu_cycle", "test_fsu_cycle", acctest.Optional, acctest.Update, FleetSoftwareUpdateFsuCycleRepresentation_GI_ImageType)

	FleetSoftwareUpdateFsuCycleSingularDataSourceRepresentation = map[string]interface{}{
		"fsu_cycle_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_fleet_software_update_fsu_cycle.test_fsu_cycle.id}`},
	}

	FleetSoftwareUpdateFsuCycle_DB_DataSourceRepresentation = map[string]interface{}{
		"compartment_id":    acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"collection_type":   acctest.Representation{RepType: acctest.Optional, Create: `DB`},
		"display_name":      acctest.Representation{RepType: acctest.Optional, Create: `TF_TEST_Cycle`, Update: `TF_TEST_Cycle_Updated`},
		"fsu_collection_id": acctest.Representation{RepType: acctest.Optional, Create: `${oci_fleet_software_update_fsu_collection.test_fsu_collection.id}`},
		"state":             acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"target_version":    acctest.Representation{RepType: acctest.Optional, Create: `targetVersion`},
		"filter":            acctest.RepresentationGroup{RepType: acctest.Required, Group: FleetSoftwareUpdateFsuCycleDataSourceFilterRepresentation},
	}

	FleetSoftwareUpdateFsuCycle_GI_DataSourceRepresentation = map[string]interface{}{
		"compartment_id":    acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"collection_type":   acctest.Representation{RepType: acctest.Optional, Create: `GI`},
		"display_name":      acctest.Representation{RepType: acctest.Optional, Create: `TF_TEST_Cycle`, Update: `TF_TEST_Cycle_Updated`},
		"fsu_collection_id": acctest.Representation{RepType: acctest.Optional, Create: `${oci_fleet_software_update_fsu_collection.test_fsu_collection.id}`},
		"state":             acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"target_version":    acctest.Representation{RepType: acctest.Optional, Create: `targetVersion`},
		"filter":            acctest.RepresentationGroup{RepType: acctest.Required, Group: FleetSoftwareUpdateFsuCycleDataSourceFilterRepresentation},
	}

	FleetSoftwareUpdateFsuCycleDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_fleet_software_update_fsu_cycle.test_fsu_cycle.id}`}},
	}

	ignoreFsuCycleDefinedTagsChangesRepresentation = map[string]interface{}{
		"ignore_changes": acctest.Representation{RepType: acctest.Required, Create: []string{`defined_tags`, `system_tags`, `freeform_tags`}},
	}

	FleetSoftwareUpdateFsuCycleRepresentation_DB_VersionType = map[string]interface{}{
		"compartment_id":               acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"fsu_collection_id":            acctest.Representation{RepType: acctest.Required, Create: `${oci_fleet_software_update_fsu_collection.test_fsu_collection.id}`},
		"goal_version_details":         acctest.RepresentationGroup{RepType: acctest.Required, Group: FleetSoftwareUpdateFsuCycleGoalVersionDetailsRepresentation_DB_VersionType},
		"type":                         acctest.Representation{RepType: acctest.Required, Create: `PATCH`},
		"batching_strategy":            acctest.RepresentationGroup{RepType: acctest.Optional, Group: FleetSoftwareUpdateFsuCycleBatchingStrategyRepresentation},
		"defined_tags":                 acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"display_name":                 acctest.Representation{RepType: acctest.Optional, Create: `TF_TEST_Cycle`, Update: `TF_TEST_Cycle_Updated`},
		"freeform_tags":                acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"bar-key": "value"}, Update: map[string]string{"Department": "Accounting"}},
		"is_ignore_missing_patches":    acctest.Representation{RepType: acctest.Optional, Create: []string{`isIgnoreMissingPatches`}, Update: []string{`isIgnoreMissingPatches2`}},
		"is_ignore_patches":            acctest.Representation{RepType: acctest.Optional, Create: `true`, Update: `false`},
		"is_keep_placement":            acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"max_drain_timeout_in_seconds": acctest.Representation{RepType: acctest.Optional, Create: `10`, Update: `11`},
		"lifecycle":                    acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreFsuCycleDefinedTagsChangesRepresentation},
		// UDX-22040-OPT-IN
		"diagnostics_collection": acctest.RepresentationGroup{RepType: acctest.Optional, Group: DataCollectionModesRepresentation},
	}

	FleetSoftwareUpdateFsuCycleRepresentation_DB_ImageType = map[string]interface{}{
		"compartment_id":               acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"fsu_collection_id":            acctest.Representation{RepType: acctest.Required, Create: `${oci_fleet_software_update_fsu_collection.test_fsu_collection.id}`},
		"goal_version_details":         acctest.RepresentationGroup{RepType: acctest.Required, Group: FleetSoftwareUpdateFsuCycleGoalVersionDetailsRepresentation_DB_ImageType},
		"type":                         acctest.Representation{RepType: acctest.Required, Create: `PATCH`},
		"batching_strategy":            acctest.RepresentationGroup{RepType: acctest.Optional, Group: FleetSoftwareUpdateFsuCycleBatchingStrategyRepresentation},
		"defined_tags":                 acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"display_name":                 acctest.Representation{RepType: acctest.Optional, Create: `TF_TEST_Cycle`, Update: `TF_TEST_Cycle_Updated`},
		"freeform_tags":                acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"bar-key": "value"}, Update: map[string]string{"Department": "Accounting"}},
		"is_ignore_missing_patches":    acctest.Representation{RepType: acctest.Optional, Create: []string{`isIgnoreMissingPatches`}, Update: []string{`isIgnoreMissingPatches2`}},
		"is_ignore_patches":            acctest.Representation{RepType: acctest.Optional, Create: `true`, Update: `false`},
		"is_keep_placement":            acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"max_drain_timeout_in_seconds": acctest.Representation{RepType: acctest.Optional, Create: `10`, Update: `11`},
		"lifecycle":                    acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreFsuCycleDefinedTagsChangesRepresentation},
		// UDX-22040-OPT-IN
		"diagnostics_collection": acctest.RepresentationGroup{RepType: acctest.Optional, Group: DataCollectionModesRepresentation},
	}

	FleetSoftwareUpdateFsuCycleRepresentation_GI_ImageType = map[string]interface{}{
		"compartment_id":    acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"fsu_collection_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_fleet_software_update_fsu_collection.test_fsu_collection.id}`},
		// UDX-21995-GI-CUSTOM-IMAGE
		"goal_version_details":         acctest.RepresentationGroup{RepType: acctest.Required, Group: FleetSoftwareUpdateFsuCycleGoalVersionDetailsRepresentation_GI_ImageType},
		"type":                         acctest.Representation{RepType: acctest.Required, Create: `PATCH`},
		"batching_strategy":            acctest.RepresentationGroup{RepType: acctest.Optional, Group: FleetSoftwareUpdateFsuCycleBatchingStrategyRepresentation},
		"defined_tags":                 acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"display_name":                 acctest.Representation{RepType: acctest.Optional, Create: `TF_TEST_Cycle`, Update: `TF_TEST_Cycle_Updated`},
		"freeform_tags":                acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"bar-key": "value"}, Update: map[string]string{"Department": "Accounting"}},
		"is_ignore_missing_patches":    acctest.Representation{RepType: acctest.Optional, Create: []string{`isIgnoreMissingPatches`}, Update: []string{`isIgnoreMissingPatches2`}},
		"is_ignore_patches":            acctest.Representation{RepType: acctest.Optional, Create: `true`, Update: `false`},
		"is_keep_placement":            acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"max_drain_timeout_in_seconds": acctest.Representation{RepType: acctest.Optional, Create: `10`, Update: `11`},
		"lifecycle":                    acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreFsuCycleDefinedTagsChangesRepresentation},
		// UDX-22040-OPT-IN
		"diagnostics_collection": acctest.RepresentationGroup{RepType: acctest.Optional, Group: DataCollectionModesRepresentation},
	}

	FleetSoftwareUpdateFsuCycleGoalVersionDetailsRepresentation_DB_VersionType = map[string]interface{}{
		"type":            acctest.Representation{RepType: acctest.Required, Create: `VERSION`, Update: `VERSION`},
		"version":         acctest.Representation{RepType: acctest.Required, Create: `19.17.0.0.0`, Update: `19.19.0.0.0`},
		"home_policy":     acctest.Representation{RepType: acctest.Optional, Create: `CREATE_NEW`, Update: `USE_EXISTING`},
		"new_home_prefix": acctest.Representation{RepType: acctest.Optional, Create: `newHomePrefix`, Update: `newHomePrefix2`},
	}

	FleetSoftwareUpdateFsuCycleGoalVersionDetailsRepresentation_DB_ImageType = map[string]interface{}{
		"type":              acctest.Representation{RepType: acctest.Required, Create: `IMAGE_ID`},
		"software_image_id": acctest.Representation{RepType: acctest.Required, Create: `${var.db_software_image_1}`},
		"home_policy":       acctest.Representation{RepType: acctest.Optional, Create: `CREATE_NEW`, Update: `USE_EXISTING`},
		"new_home_prefix":   acctest.Representation{RepType: acctest.Optional, Create: `newHomePrefix`, Update: `newHomePrefix2`},
	}

	// UDX-21995-GI-CUSTOM-IMAGE
	FleetSoftwareUpdateFsuCycleGoalVersionDetailsRepresentation_GI_ImageType = map[string]interface{}{
		"type":              acctest.Representation{RepType: acctest.Required, Create: `IMAGE_ID`},
		"software_image_id": acctest.Representation{RepType: acctest.Required, Create: `${var.db_grid_software_image_1}`},
	}

	FleetSoftwareUpdateFsuCycleBatchingStrategyRepresentation = map[string]interface{}{
		"type":                     acctest.Representation{RepType: acctest.Required, Create: `SEQUENTIAL`, Update: `FIFTY_FIFTY`},
		"is_force_rolling":         acctest.Representation{RepType: acctest.Required, Create: `true`, Update: `true`},
		"is_wait_for_batch_resume": acctest.Representation{RepType: acctest.Optional, Create: nil, Update: `false`},
	}

	FleetSoftwareUpdateFsuCycleStageActionScheduleRepresentation = map[string]interface{}{
		"type":          acctest.Representation{RepType: acctest.Required, Create: `START_TIME`, Update: `START_TIME`},
		"time_to_start": acctest.Representation{RepType: acctest.Required, Create: scheduledStageTime, Update: scheduledStageTimeUpdate},
	}

	FleetSoftwareUpdateFsuCycleApplyActionScheduleRepresentation = map[string]interface{}{
		"type":          acctest.Representation{RepType: acctest.Required, Create: `START_TIME`, Update: `START_TIME`},
		"time_to_start": acctest.Representation{RepType: acctest.Required, Create: scheduledApplyTime, Update: scheduledApplyTimeUpdate},
	}

	// UDX-22040-OPT-IN
	DataCollectionModesRepresentation = map[string]interface{}{
		"log_collection_mode": acctest.Representation{RepType: acctest.Optional, Create: `ENABLE`, Update: `NO_CHANGE`},
	}

	FleetSoftwareUpdateFsuCycleDBResourceDependencies = utils.OciImageIdsVariable +
		acctest.GenerateResourceFromRepresentationMap("oci_fleet_software_update_fsu_collection", "test_fsu_collection", acctest.Required, acctest.Create, FleetSoftwareUpdateFsuCollectionDBRepresentation) +
		DefinedTagsDependencies

	FleetSoftwareUpdateFsuCycleGIResourceDependencies = utils.OciImageIdsVariable +
		acctest.GenerateResourceFromRepresentationMap("oci_fleet_software_update_fsu_collection", "test_fsu_collection", acctest.Required, acctest.Create, FleetSoftwareUpdateFsuCollectionGIRepresentation) +
		DefinedTagsDependencies
)

// issue-routing-tag: fleet_software_update/default
func TestFleetSoftwareUpdateFsuCycleResource_DB_VersionDetails(t *testing.T) {
	httpreplay.SetScenario("TestFleetSoftwareUpdateFsuCycleResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	dbTargetId1 := utils.GetEnvSettingWithBlankDefault("fsu_db_target_1")
	dbTargetId1VariableStr := fmt.Sprintf("variable \"db_target_1\" { default = \"%s\" }\n", dbTargetId1)

	dbTargetId2 := utils.GetEnvSettingWithBlankDefault("fsu_db_target_2")
	dbTargetId2VariableStr := fmt.Sprintf("variable \"db_target_2\" { default = \"%s\" }\n", dbTargetId2)

	var variablesStr = compartmentIdVariableStr + dbTargetId1VariableStr + dbTargetId2VariableStr

	resourceName := "oci_fleet_software_update_fsu_cycle.test_fsu_cycle"
	datasourceName := "data.oci_fleet_software_update_fsu_cycles.test_fsu_cycles"
	singularDatasourceName := "data.oci_fleet_software_update_fsu_cycle.test_fsu_cycle"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+variablesStr+FleetSoftwareUpdateFsuCycleDBResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_fleet_software_update_fsu_cycle", "test_fsu_cycle", acctest.Optional, acctest.Create, FleetSoftwareUpdateFsuCycleRepresentation_DB_VersionType), "fleetsoftwareupdate", "fsuCycle", t)

	acctest.ResourceTest(t, testAccCheckFleetSoftwareUpdateFsuCycleDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + variablesStr + FleetSoftwareUpdateFsuCycleDBResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_fleet_software_update_fsu_cycle", "test_fsu_cycle", acctest.Required, acctest.Create, FleetSoftwareUpdateFsuCycleRepresentation_DB_VersionType),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName, "fsu_collection_id"),
				resource.TestCheckResourceAttr(resourceName, "goal_version_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "goal_version_details.0.type", "VERSION"),
				resource.TestCheckResourceAttr(resourceName, "goal_version_details.0.version", "19.17.0.0.0"),
				resource.TestCheckResourceAttr(resourceName, "type", "PATCH"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + variablesStr + FleetSoftwareUpdateFsuCycleDBResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + variablesStr + FleetSoftwareUpdateFsuCycleDBResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_fleet_software_update_fsu_cycle", "test_fsu_cycle", acctest.Optional, acctest.Create, FleetSoftwareUpdateFsuCycleRepresentation_DB_VersionType),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "batching_strategy.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "batching_strategy.0.is_force_rolling", "true"),
				resource.TestCheckResourceAttr(resourceName, "batching_strategy.0.type", "SEQUENTIAL"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "TF_TEST_Cycle"),
				resource.TestCheckResourceAttrSet(resourceName, "fsu_collection_id"),
				resource.TestCheckResourceAttr(resourceName, "goal_version_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "goal_version_details.0.home_policy", "CREATE_NEW"),
				resource.TestCheckResourceAttr(resourceName, "goal_version_details.0.new_home_prefix", "newHomePrefix"),
				resource.TestCheckResourceAttr(resourceName, "goal_version_details.0.type", "VERSION"),
				resource.TestCheckResourceAttr(resourceName, "goal_version_details.0.version", "19.17.0.0.0"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_ignore_missing_patches.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "is_ignore_patches", "true"),
				resource.TestCheckResourceAttr(resourceName, "is_keep_placement", "false"),
				resource.TestCheckResourceAttr(resourceName, "max_drain_timeout_in_seconds", "10"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttr(resourceName, "type", "PATCH"),
				// UDX-22040-OPT-IN
				resource.TestCheckResourceAttr(resourceName, "diagnostics_collection.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "diagnostics_collection.0.log_collection_mode", "ENABLE"),

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
			Config: config + variablesStr + compartmentIdUVariableStr + FleetSoftwareUpdateFsuCycleDBResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_fleet_software_update_fsu_cycle", "test_fsu_cycle", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(FleetSoftwareUpdateFsuCycleRepresentation_DB_VersionType, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "batching_strategy.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "batching_strategy.0.is_force_rolling", "true"),
				resource.TestCheckResourceAttr(resourceName, "batching_strategy.0.type", "SEQUENTIAL"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "display_name", "TF_TEST_Cycle"),
				resource.TestCheckResourceAttrSet(resourceName, "fsu_collection_id"),
				resource.TestCheckResourceAttr(resourceName, "goal_version_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "goal_version_details.0.home_policy", "CREATE_NEW"),
				resource.TestCheckResourceAttr(resourceName, "goal_version_details.0.new_home_prefix", "newHomePrefix"),
				resource.TestCheckResourceAttr(resourceName, "goal_version_details.0.type", "VERSION"),
				resource.TestCheckResourceAttr(resourceName, "goal_version_details.0.version", "19.17.0.0.0"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_ignore_missing_patches.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "is_ignore_patches", "true"),
				resource.TestCheckResourceAttr(resourceName, "is_keep_placement", "false"),
				resource.TestCheckResourceAttr(resourceName, "max_drain_timeout_in_seconds", "10"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttr(resourceName, "type", "PATCH"),
				// UDX-22040-OPT-IN
				resource.TestCheckResourceAttr(resourceName, "diagnostics_collection.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "diagnostics_collection.0.log_collection_mode", "ENABLE"),

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
			Config: config + variablesStr + FleetSoftwareUpdateFsuCycleDBResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_fleet_software_update_fsu_cycle", "test_fsu_cycle", acctest.Optional, acctest.Update, FleetSoftwareUpdateFsuCycleRepresentation_DB_VersionType),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "batching_strategy.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "batching_strategy.0.is_force_rolling", "true"),
				resource.TestCheckResourceAttr(resourceName, "batching_strategy.0.is_wait_for_batch_resume", "false"),
				resource.TestCheckResourceAttr(resourceName, "batching_strategy.0.type", "FIFTY_FIFTY"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "TF_TEST_Cycle_Updated"),
				resource.TestCheckResourceAttrSet(resourceName, "fsu_collection_id"),
				resource.TestCheckResourceAttr(resourceName, "goal_version_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "goal_version_details.0.home_policy", "USE_EXISTING"),
				resource.TestCheckResourceAttr(resourceName, "goal_version_details.0.new_home_prefix", "newHomePrefix2"),
				resource.TestCheckResourceAttr(resourceName, "goal_version_details.0.type", "VERSION"),
				resource.TestCheckResourceAttr(resourceName, "goal_version_details.0.version", "19.19.0.0.0"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_ignore_missing_patches.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "is_ignore_patches", "false"),
				resource.TestCheckResourceAttr(resourceName, "is_keep_placement", "true"),
				resource.TestCheckResourceAttr(resourceName, "max_drain_timeout_in_seconds", "11"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttr(resourceName, "type", "PATCH"),
				// UDX-22040-OPT-IN
				resource.TestCheckResourceAttr(resourceName, "diagnostics_collection.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "diagnostics_collection.0.log_collection_mode", "NO_CHANGE"),

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
				acctest.GenerateDataSourceFromRepresentationMap("oci_fleet_software_update_fsu_cycles", "test_fsu_cycles", acctest.Optional, acctest.Update, FleetSoftwareUpdateFsuCycle_DB_DataSourceRepresentation) +
				variablesStr + FleetSoftwareUpdateFsuCycleDBResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_fleet_software_update_fsu_cycle", "test_fsu_cycle", acctest.Optional, acctest.Update, FleetSoftwareUpdateFsuCycleRepresentation_DB_VersionType),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "collection_type", "DB"),
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "TF_TEST_Cycle_Updated"),
				resource.TestCheckResourceAttrSet(datasourceName, "fsu_collection_id"),
				resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),
				resource.TestCheckResourceAttr(datasourceName, "target_version", "targetVersion"),
				resource.TestCheckResourceAttr(datasourceName, "fsu_cycle_summary_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "fsu_cycle_summary_collection.0.items.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_fleet_software_update_fsu_cycle", "test_fsu_cycle", acctest.Required, acctest.Create, FleetSoftwareUpdateFsuCycleSingularDataSourceRepresentation) +
				variablesStr + FleetSoftwareUpdateFsuCycleResourceConfig_DB_VersionType,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "fsu_cycle_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "batching_strategy.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "batching_strategy.0.is_force_rolling", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "batching_strategy.0.is_wait_for_batch_resume", "false"),
				resource.TestCheckResourceAttr(singularDatasourceName, "batching_strategy.0.type", "FIFTY_FIFTY"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "collection_type"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "TF_TEST_Cycle_Updated"),
				resource.TestCheckResourceAttr(singularDatasourceName, "goal_version_details.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "goal_version_details.0.home_policy", "USE_EXISTING"),
				resource.TestCheckResourceAttr(singularDatasourceName, "goal_version_details.0.new_home_prefix", "newHomePrefix2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "goal_version_details.0.type", "VERSION"),
				resource.TestCheckResourceAttr(singularDatasourceName, "goal_version_details.0.version", "19.19.0.0.0"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "is_ignore_missing_patches.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "is_ignore_patches", "false"),
				resource.TestCheckResourceAttr(singularDatasourceName, "is_keep_placement", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "max_drain_timeout_in_seconds", "11"),
				resource.TestCheckResourceAttr(singularDatasourceName, "next_action_to_execute.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
				resource.TestCheckResourceAttr(singularDatasourceName, "type", "PATCH"),
			),
		},
		// verify resource import
		{
			Config:                  config + FleetSoftwareUpdateFsuCycleRequiredOnlyResource_DB_VersionType,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{"identity_domain", "defined_tags", "system_tags", "freeform_tags"},
			ResourceName:            resourceName,
		},
	})
}

func TestFleetSoftwareUpdateFsuCycleResource_DB_ImageIdDetails(t *testing.T) {
	httpreplay.SetScenario("TestFleetSoftwareUpdateFsuCycleResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	dbTargetId1 := utils.GetEnvSettingWithBlankDefault("fsu_db_target_1")
	dbTargetId1VariableStr := fmt.Sprintf("variable \"db_target_1\" { default = \"%s\" }\n", dbTargetId1)

	dbTargetId2 := utils.GetEnvSettingWithBlankDefault("fsu_db_target_2")
	dbTargetId2VariableStr := fmt.Sprintf("variable \"db_target_2\" { default = \"%s\" }\n", dbTargetId2)

	dbSwImage1 := utils.GetEnvSettingWithBlankDefault("fsu_db_software_image_1")
	dbSwImage1VariableStr := fmt.Sprintf("variable \"db_software_image_1\" { default = \"%s\" }\n", dbSwImage1)

	var variablesStr = compartmentIdVariableStr + dbTargetId1VariableStr + dbTargetId2VariableStr + dbSwImage1VariableStr

	resourceName := "oci_fleet_software_update_fsu_cycle.test_fsu_cycle"
	datasourceName := "data.oci_fleet_software_update_fsu_cycles.test_fsu_cycles"
	singularDatasourceName := "data.oci_fleet_software_update_fsu_cycle.test_fsu_cycle"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+variablesStr+FleetSoftwareUpdateFsuCycleDBResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_fleet_software_update_fsu_cycle", "test_fsu_cycle", acctest.Optional, acctest.Create, FleetSoftwareUpdateFsuCycleRepresentation_DB_ImageType), "fleetsoftwareupdate", "fsuCycle", t)

	acctest.ResourceTest(t, testAccCheckFleetSoftwareUpdateFsuCycleDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + variablesStr + FleetSoftwareUpdateFsuCycleDBResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_fleet_software_update_fsu_cycle", "test_fsu_cycle", acctest.Required, acctest.Create, FleetSoftwareUpdateFsuCycleRepresentation_DB_ImageType),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName, "fsu_collection_id"),
				resource.TestCheckResourceAttr(resourceName, "goal_version_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "goal_version_details.0.type", "IMAGE_ID"),
				resource.TestCheckResourceAttrSet(resourceName, "goal_version_details.0.software_image_id"),
				resource.TestCheckResourceAttr(resourceName, "type", "PATCH"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + variablesStr + FleetSoftwareUpdateFsuCycleDBResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + variablesStr + FleetSoftwareUpdateFsuCycleDBResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_fleet_software_update_fsu_cycle", "test_fsu_cycle", acctest.Optional, acctest.Create, FleetSoftwareUpdateFsuCycleRepresentation_DB_ImageType),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "batching_strategy.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "batching_strategy.0.is_force_rolling", "true"),
				resource.TestCheckResourceAttr(resourceName, "batching_strategy.0.type", "SEQUENTIAL"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "TF_TEST_Cycle"),
				resource.TestCheckResourceAttrSet(resourceName, "fsu_collection_id"),
				resource.TestCheckResourceAttr(resourceName, "goal_version_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "goal_version_details.0.home_policy", "CREATE_NEW"),
				resource.TestCheckResourceAttr(resourceName, "goal_version_details.0.new_home_prefix", "newHomePrefix"),
				resource.TestCheckResourceAttrSet(resourceName, "goal_version_details.0.software_image_id"),
				resource.TestCheckResourceAttr(resourceName, "goal_version_details.0.type", "IMAGE_ID"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_ignore_missing_patches.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "is_ignore_patches", "true"),
				resource.TestCheckResourceAttr(resourceName, "is_keep_placement", "false"),
				resource.TestCheckResourceAttr(resourceName, "max_drain_timeout_in_seconds", "10"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttr(resourceName, "type", "PATCH"),
				// UDX-22040-OPT-IN
				resource.TestCheckResourceAttr(resourceName, "diagnostics_collection.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "diagnostics_collection.0.log_collection_mode", "ENABLE"),

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
			Config: config + variablesStr + compartmentIdUVariableStr + FleetSoftwareUpdateFsuCycleDBResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_fleet_software_update_fsu_cycle", "test_fsu_cycle", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(FleetSoftwareUpdateFsuCycleRepresentation_DB_ImageType, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "batching_strategy.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "batching_strategy.0.is_force_rolling", "true"),
				resource.TestCheckResourceAttr(resourceName, "batching_strategy.0.type", "SEQUENTIAL"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "display_name", "TF_TEST_Cycle"),
				resource.TestCheckResourceAttrSet(resourceName, "fsu_collection_id"),
				resource.TestCheckResourceAttr(resourceName, "goal_version_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "goal_version_details.0.home_policy", "CREATE_NEW"),
				resource.TestCheckResourceAttr(resourceName, "goal_version_details.0.new_home_prefix", "newHomePrefix"),
				resource.TestCheckResourceAttrSet(resourceName, "goal_version_details.0.software_image_id"),
				resource.TestCheckResourceAttr(resourceName, "goal_version_details.0.type", "IMAGE_ID"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_ignore_missing_patches.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "is_ignore_patches", "true"),
				resource.TestCheckResourceAttr(resourceName, "is_keep_placement", "false"),
				resource.TestCheckResourceAttr(resourceName, "max_drain_timeout_in_seconds", "10"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttr(resourceName, "type", "PATCH"),
				// UDX-22040-OPT-IN
				resource.TestCheckResourceAttr(resourceName, "diagnostics_collection.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "diagnostics_collection.0.log_collection_mode", "ENABLE"),

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
			Config: config + variablesStr + FleetSoftwareUpdateFsuCycleDBResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_fleet_software_update_fsu_cycle", "test_fsu_cycle", acctest.Optional, acctest.Update, FleetSoftwareUpdateFsuCycleRepresentation_DB_ImageType),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "batching_strategy.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "batching_strategy.0.is_force_rolling", "true"),
				resource.TestCheckResourceAttr(resourceName, "batching_strategy.0.is_wait_for_batch_resume", "false"),
				resource.TestCheckResourceAttr(resourceName, "batching_strategy.0.type", "FIFTY_FIFTY"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "TF_TEST_Cycle_Updated"),
				resource.TestCheckResourceAttrSet(resourceName, "fsu_collection_id"),
				resource.TestCheckResourceAttr(resourceName, "goal_version_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "goal_version_details.0.home_policy", "USE_EXISTING"),
				resource.TestCheckResourceAttr(resourceName, "goal_version_details.0.new_home_prefix", "newHomePrefix2"),
				resource.TestCheckResourceAttrSet(resourceName, "goal_version_details.0.software_image_id"),
				resource.TestCheckResourceAttr(resourceName, "goal_version_details.0.type", "IMAGE_ID"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_ignore_missing_patches.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "is_ignore_patches", "false"),
				resource.TestCheckResourceAttr(resourceName, "is_keep_placement", "true"),
				resource.TestCheckResourceAttr(resourceName, "max_drain_timeout_in_seconds", "11"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttr(resourceName, "type", "PATCH"),
				// UDX-22040-OPT-IN
				resource.TestCheckResourceAttr(resourceName, "diagnostics_collection.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "diagnostics_collection.0.log_collection_mode", "NO_CHANGE"),

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
				acctest.GenerateDataSourceFromRepresentationMap("oci_fleet_software_update_fsu_cycles", "test_fsu_cycles", acctest.Optional, acctest.Update, FleetSoftwareUpdateFsuCycle_DB_DataSourceRepresentation) +
				variablesStr + FleetSoftwareUpdateFsuCycleDBResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_fleet_software_update_fsu_cycle", "test_fsu_cycle", acctest.Optional, acctest.Update, FleetSoftwareUpdateFsuCycleRepresentation_DB_ImageType),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "collection_type", "DB"),
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "TF_TEST_Cycle_Updated"),
				resource.TestCheckResourceAttrSet(datasourceName, "fsu_collection_id"),
				resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),
				resource.TestCheckResourceAttr(datasourceName, "target_version", "targetVersion"),
				resource.TestCheckResourceAttr(datasourceName, "fsu_cycle_summary_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "fsu_cycle_summary_collection.0.items.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_fleet_software_update_fsu_cycle", "test_fsu_cycle", acctest.Required, acctest.Create, FleetSoftwareUpdateFsuCycleSingularDataSourceRepresentation) +
				variablesStr + FleetSoftwareUpdateFsuCycleResourceConfig_DB_ImageType,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "fsu_cycle_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "batching_strategy.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "batching_strategy.0.is_force_rolling", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "batching_strategy.0.is_wait_for_batch_resume", "false"),
				resource.TestCheckResourceAttr(singularDatasourceName, "batching_strategy.0.type", "FIFTY_FIFTY"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "collection_type"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "TF_TEST_Cycle_Updated"),
				resource.TestCheckResourceAttr(singularDatasourceName, "goal_version_details.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "goal_version_details.0.home_policy", "USE_EXISTING"),
				resource.TestCheckResourceAttr(singularDatasourceName, "goal_version_details.0.new_home_prefix", "newHomePrefix2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "goal_version_details.0.type", "IMAGE_ID"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "is_ignore_missing_patches.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "is_ignore_patches", "false"),
				resource.TestCheckResourceAttr(singularDatasourceName, "is_keep_placement", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "max_drain_timeout_in_seconds", "11"),
				resource.TestCheckResourceAttr(singularDatasourceName, "next_action_to_execute.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
				resource.TestCheckResourceAttr(singularDatasourceName, "type", "PATCH"),
			),
		},
		// verify resource import
		{
			Config:                  config + FleetSoftwareUpdateFsuCycleRequiredOnlyResource_DB_ImageType,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{"identity_domain", "defined_tags", "system_tags", "freeform_tags"},
			ResourceName:            resourceName,
		},
	})
}

func TestFleetSoftwareUpdateFsuCycleResource_GI_ImageIdDetails(t *testing.T) {
	httpreplay.SetScenario("TestFleetSoftwareUpdateFsuCycleResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	giTargetId1 := utils.GetEnvSettingWithBlankDefault("fsu_gi_target_1")
	giTargetId1VariableStr := fmt.Sprintf("variable \"gi_target_1\" { default = \"%s\" }\n", giTargetId1)

	giTargetId2 := utils.GetEnvSettingWithBlankDefault("fsu_gi_target_2")
	giTargetId2VariableStr := fmt.Sprintf("variable \"gi_target_2\" { default = \"%s\" }\n", giTargetId2)

	dbGridSwImage1 := utils.GetEnvSettingWithBlankDefault("fsu_db_grid_software_image_1")
	dbGridSwImage1VariableStr := fmt.Sprintf("variable \"db_grid_software_image_1\" { default = \"%s\" }\n", dbGridSwImage1)

	var variablesStr = compartmentIdVariableStr + giTargetId1VariableStr + giTargetId2VariableStr + dbGridSwImage1VariableStr

	resourceName := "oci_fleet_software_update_fsu_cycle.test_fsu_cycle"
	datasourceName := "data.oci_fleet_software_update_fsu_cycles.test_fsu_cycles"
	singularDatasourceName := "data.oci_fleet_software_update_fsu_cycle.test_fsu_cycle"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+variablesStr+FleetSoftwareUpdateFsuCycleGIResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_fleet_software_update_fsu_cycle", "test_fsu_cycle", acctest.Optional, acctest.Create, FleetSoftwareUpdateFsuCycleRepresentation_GI_ImageType), "fleetsoftwareupdate", "fsuCycle", t)

	acctest.ResourceTest(t, testAccCheckFleetSoftwareUpdateFsuCycleDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + variablesStr + FleetSoftwareUpdateFsuCycleGIResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_fleet_software_update_fsu_cycle", "test_fsu_cycle", acctest.Required, acctest.Create, FleetSoftwareUpdateFsuCycleRepresentation_GI_ImageType),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName, "fsu_collection_id"),
				resource.TestCheckResourceAttr(resourceName, "goal_version_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "type", "PATCH"),
				// UDX-21995-GI-CUSTOM-IMAGE
				resource.TestCheckResourceAttr(resourceName, "goal_version_details.0.type", "IMAGE_ID"),
				resource.TestCheckResourceAttrSet(resourceName, "goal_version_details.0.software_image_id"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + variablesStr + FleetSoftwareUpdateFsuCycleGIResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + variablesStr + FleetSoftwareUpdateFsuCycleGIResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_fleet_software_update_fsu_cycle", "test_fsu_cycle", acctest.Optional, acctest.Create, FleetSoftwareUpdateFsuCycleRepresentation_GI_ImageType),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "batching_strategy.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "batching_strategy.0.is_force_rolling", "true"),
				resource.TestCheckResourceAttr(resourceName, "batching_strategy.0.type", "SEQUENTIAL"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "TF_TEST_Cycle"),
				resource.TestCheckResourceAttrSet(resourceName, "fsu_collection_id"),
				resource.TestCheckResourceAttr(resourceName, "goal_version_details.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_ignore_missing_patches.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "is_ignore_patches", "true"),
				resource.TestCheckResourceAttr(resourceName, "is_keep_placement", "false"),
				resource.TestCheckResourceAttr(resourceName, "max_drain_timeout_in_seconds", "10"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttr(resourceName, "type", "PATCH"),
				// UDX-22040-OPT-IN
				resource.TestCheckResourceAttr(resourceName, "diagnostics_collection.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "diagnostics_collection.0.log_collection_mode", "ENABLE"),
				// UDX-21995-GI-CUSTOM-IMAGE
				resource.TestCheckResourceAttrSet(resourceName, "goal_version_details.0.software_image_id"),
				resource.TestCheckResourceAttr(resourceName, "goal_version_details.0.type", "IMAGE_ID"),

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
			Config: config + variablesStr + compartmentIdUVariableStr + FleetSoftwareUpdateFsuCycleGIResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_fleet_software_update_fsu_cycle", "test_fsu_cycle", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(FleetSoftwareUpdateFsuCycleRepresentation_GI_ImageType, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "batching_strategy.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "batching_strategy.0.is_force_rolling", "true"),
				resource.TestCheckResourceAttr(resourceName, "batching_strategy.0.type", "SEQUENTIAL"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "display_name", "TF_TEST_Cycle"),
				resource.TestCheckResourceAttrSet(resourceName, "fsu_collection_id"),
				resource.TestCheckResourceAttr(resourceName, "goal_version_details.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_ignore_missing_patches.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "is_ignore_patches", "true"),
				resource.TestCheckResourceAttr(resourceName, "is_keep_placement", "false"),
				resource.TestCheckResourceAttr(resourceName, "max_drain_timeout_in_seconds", "10"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttr(resourceName, "type", "PATCH"),
				// UDX-22040-OPT-IN
				resource.TestCheckResourceAttr(resourceName, "diagnostics_collection.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "diagnostics_collection.0.log_collection_mode", "ENABLE"),
				// UDX-21995-GI-CUSTOM-IMAGE
				resource.TestCheckResourceAttrSet(resourceName, "goal_version_details.0.software_image_id"),
				resource.TestCheckResourceAttr(resourceName, "goal_version_details.0.type", "IMAGE_ID"),

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
			Config: config + variablesStr + FleetSoftwareUpdateFsuCycleGIResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_fleet_software_update_fsu_cycle", "test_fsu_cycle", acctest.Optional, acctest.Update, FleetSoftwareUpdateFsuCycleRepresentation_GI_ImageType),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "batching_strategy.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "batching_strategy.0.is_force_rolling", "true"),
				resource.TestCheckResourceAttr(resourceName, "batching_strategy.0.is_wait_for_batch_resume", "false"),
				resource.TestCheckResourceAttr(resourceName, "batching_strategy.0.type", "FIFTY_FIFTY"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "TF_TEST_Cycle_Updated"),
				resource.TestCheckResourceAttrSet(resourceName, "fsu_collection_id"),
				resource.TestCheckResourceAttr(resourceName, "goal_version_details.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_ignore_missing_patches.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "is_ignore_patches", "false"),
				resource.TestCheckResourceAttr(resourceName, "is_keep_placement", "true"),
				resource.TestCheckResourceAttr(resourceName, "max_drain_timeout_in_seconds", "11"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttr(resourceName, "type", "PATCH"),
				// UDX-22040-OPT-IN
				resource.TestCheckResourceAttr(resourceName, "diagnostics_collection.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "diagnostics_collection.0.log_collection_mode", "NO_CHANGE"),
				// UDX-21995-GI-CUSTOM-IMAGE
				resource.TestCheckResourceAttrSet(resourceName, "goal_version_details.0.software_image_id"),
				resource.TestCheckResourceAttr(resourceName, "goal_version_details.0.type", "IMAGE_ID"),

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
				acctest.GenerateDataSourceFromRepresentationMap("oci_fleet_software_update_fsu_cycles", "test_fsu_cycles", acctest.Optional, acctest.Update, FleetSoftwareUpdateFsuCycle_GI_DataSourceRepresentation) +
				variablesStr + FleetSoftwareUpdateFsuCycleGIResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_fleet_software_update_fsu_cycle", "test_fsu_cycle", acctest.Optional, acctest.Update, FleetSoftwareUpdateFsuCycleRepresentation_GI_ImageType),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "collection_type", "GI"),
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "TF_TEST_Cycle_Updated"),
				resource.TestCheckResourceAttrSet(datasourceName, "fsu_collection_id"),
				resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),
				resource.TestCheckResourceAttr(datasourceName, "target_version", "targetVersion"),
				resource.TestCheckResourceAttr(datasourceName, "fsu_cycle_summary_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "fsu_cycle_summary_collection.0.items.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_fleet_software_update_fsu_cycle", "test_fsu_cycle", acctest.Required, acctest.Create, FleetSoftwareUpdateFsuCycleSingularDataSourceRepresentation) +
				variablesStr + FleetSoftwareUpdateFsuCycleResourceConfig_GI_ImageType,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "fsu_cycle_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "batching_strategy.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "batching_strategy.0.is_force_rolling", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "batching_strategy.0.is_wait_for_batch_resume", "false"),
				resource.TestCheckResourceAttr(singularDatasourceName, "batching_strategy.0.type", "FIFTY_FIFTY"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "collection_type"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "TF_TEST_Cycle_Updated"),
				resource.TestCheckResourceAttr(singularDatasourceName, "goal_version_details.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "is_ignore_missing_patches.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "is_ignore_patches", "false"),
				resource.TestCheckResourceAttr(singularDatasourceName, "is_keep_placement", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "max_drain_timeout_in_seconds", "11"),
				resource.TestCheckResourceAttr(singularDatasourceName, "next_action_to_execute.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
				resource.TestCheckResourceAttr(singularDatasourceName, "type", "PATCH"),
				// UDX-21995-GI-CUSTOM-IMAGE
				resource.TestCheckResourceAttrSet(resourceName, "goal_version_details.0.software_image_id"),
				resource.TestCheckResourceAttr(resourceName, "goal_version_details.0.type", "IMAGE_ID"),
			),
		},
		// verify resource import
		{
			Config:                  config + FleetSoftwareUpdateFsuCycleRequiredOnlyResource_GI_ImageType,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{"identity_domain", "defined_tags", "system_tags", "freeform_tags"},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckFleetSoftwareUpdateFsuCycleDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).FleetSoftwareUpdateClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_fleet_software_update_fsu_cycle" {
			noResourceFound = false
			request := oci_fleet_software_update.GetFsuCycleRequest{}

			tmp := rs.Primary.ID
			request.FsuCycleId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "fleet_software_update")

			response, err := client.GetFsuCycle(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_fleet_software_update.CycleLifecycleStatesDeleted): true,
				}
				if _, ok := deletedLifecycleStates[string(response.GetLifecycleState())]; !ok {
					//resource lifecycle state is not in expected deleted lifecycle states.
					return fmt.Errorf("resource lifecycle state: %s is not in expected deleted lifecycle states", response.GetLifecycleState())
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
	if !acctest.InSweeperExcludeList("FleetSoftwareUpdateFsuCycle") {
		resource.AddTestSweepers("FleetSoftwareUpdateFsuCycle", &resource.Sweeper{
			Name:         "FleetSoftwareUpdateFsuCycle",
			Dependencies: acctest.DependencyGraph["fsuCycle"],
			F:            sweepFleetSoftwareUpdateFsuCycleResource,
		})
	}
}

func sweepFleetSoftwareUpdateFsuCycleResource(compartment string) error {
	fleetSoftwareUpdateClient := acctest.GetTestClients(&schema.ResourceData{}).FleetSoftwareUpdateClient()
	fsuCycleIds, err := getFleetSoftwareUpdateFsuCycleIds(compartment)
	if err != nil {
		return err
	}
	for _, fsuCycleId := range fsuCycleIds {
		if ok := acctest.SweeperDefaultResourceId[fsuCycleId]; !ok {
			deleteFsuCycleRequest := oci_fleet_software_update.DeleteFsuCycleRequest{}

			deleteFsuCycleRequest.FsuCycleId = &fsuCycleId

			deleteFsuCycleRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "fleet_software_update")
			_, error := fleetSoftwareUpdateClient.DeleteFsuCycle(context.Background(), deleteFsuCycleRequest)
			if error != nil {
				fmt.Printf("Error deleting FsuCycle %s %s, It is possible that the resource is already deleted. Please verify manually \n", fsuCycleId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &fsuCycleId, FleetSoftwareUpdateFsuCycleSweepWaitCondition, time.Duration(3*time.Minute),
				FleetSoftwareUpdateFsuCycleSweepResponseFetchOperation, "fleet_software_update", true)
		}
	}
	return nil
}

func getFleetSoftwareUpdateFsuCycleIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "FsuCycleId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	fleetSoftwareUpdateClient := acctest.GetTestClients(&schema.ResourceData{}).FleetSoftwareUpdateClient()

	// FPPCS-1864: get all cycles in state Active, NeedsAttention and Succeeded
	var wantedCycleStatesString = [3]string{"ACTIVE", "NEEDS_ATTENTION", "SUCCEEDED"}

	for _, cycleStateString := range wantedCycleStatesString {
		listFsuCyclesRequest := oci_fleet_software_update.ListFsuCyclesRequest{}
		listFsuCyclesRequest.CompartmentId = &compartmentId
		lifeCycleState, _ := oci_fleet_software_update.GetMappingListFsuCyclesLifecycleStateEnum(cycleStateString)
		listFsuCyclesRequest.LifecycleState = lifeCycleState
		listFsuCyclesResponse, err := fleetSoftwareUpdateClient.ListFsuCycles(context.Background(), listFsuCyclesRequest)
		if err != nil {
			return resourceIds, fmt.Errorf("Error getting FsuCycle list for compartment id : %s , %s \n", compartmentId, err)
		}
		for _, fsuCycle := range listFsuCyclesResponse.Items {
			id := *fsuCycle.Id
			resourceIds = append(resourceIds, id)
			acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "FsuCycleId", id)
		}
	}

	return resourceIds, nil
}

func FleetSoftwareUpdateFsuCycleSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if fsuCycleResponse, ok := response.Response.(oci_fleet_software_update.GetFsuCycleResponse); ok {
		return fsuCycleResponse.GetLifecycleState() != oci_fleet_software_update.CycleLifecycleStatesDeleted
	}
	return false
}

func FleetSoftwareUpdateFsuCycleSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.FleetSoftwareUpdateClient().GetFsuCycle(context.Background(), oci_fleet_software_update.GetFsuCycleRequest{
		FsuCycleId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
