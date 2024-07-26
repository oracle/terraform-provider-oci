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

var (
	FleetSoftwareUpdateFsuCollectionDBRequiredOnlyResource = FleetSoftwareUpdateFsuCollectionResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_fleet_software_update_fsu_collection", "test_fsu_collection", acctest.Required, acctest.Create, FleetSoftwareUpdateFsuCollectionDBRepresentation)

	FleetSoftwareUpdateFsuCollectionGIRequiredOnlyResource = FleetSoftwareUpdateFsuCollectionResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_fleet_software_update_fsu_collection", "test_fsu_collection", acctest.Required, acctest.Create, FleetSoftwareUpdateFsuCollectionGIRepresentation)

	FleetSoftwareUpdateFsuCollectionDBResourceConfig = FleetSoftwareUpdateFsuCollectionResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_fleet_software_update_fsu_collection", "test_fsu_collection", acctest.Optional, acctest.Update, FleetSoftwareUpdateFsuCollectionDBRepresentation)

	FleetSoftwareUpdateFsuCollectionGIResourceConfig = FleetSoftwareUpdateFsuCollectionResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_fleet_software_update_fsu_collection", "test_fsu_collection", acctest.Optional, acctest.Update, FleetSoftwareUpdateFsuCollectionGIRepresentation)

	FleetSoftwareUpdateFsuCollectionSingularDataSourceRepresentation = map[string]interface{}{
		"fsu_collection_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_fleet_software_update_fsu_collection.test_fsu_collection.id}`},
	}

	FleetSoftwareUpdateFsuCollectionDBDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":   acctest.Representation{RepType: acctest.Optional, Create: `TF_TEST_Collection`, Update: `TF_TEST_Collection_Updated`},
		"state":          acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"type":           acctest.Representation{RepType: acctest.Optional, Create: `DB`},
		"filter":         acctest.RepresentationGroup{RepType: acctest.Required, Group: FleetSoftwareUpdateFsuCollectionDataSourceFilterRepresentation}}

	FleetSoftwareUpdateFsuCollectionGIDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":   acctest.Representation{RepType: acctest.Optional, Create: `TF_TEST_Collection`, Update: `TF_TEST_Collection_Updated`},
		"state":          acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"type":           acctest.Representation{RepType: acctest.Optional, Create: `GI`},
		"filter":         acctest.RepresentationGroup{RepType: acctest.Required, Group: FleetSoftwareUpdateFsuCollectionDataSourceFilterRepresentation}}

	FleetSoftwareUpdateFsuCollectionDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_fleet_software_update_fsu_collection.test_fsu_collection.id}`}},
	}

	ignoreFsuCollectionDefinedTagsChangesRepresentation = map[string]interface{}{
		"ignore_changes": acctest.Representation{RepType: acctest.Required, Create: []string{`defined_tags`, `system_tags`, `freeform_tags`}},
	}

	// https://docs.oracle.com/en-us/iaas/api/#/en/edsfu/20220528/datatypes/CreateDbFsuCollectionDetails
	FleetSoftwareUpdateFsuCollectionDBRepresentation = map[string]interface{}{
		"compartment_id":       acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"service_type":         acctest.Representation{RepType: acctest.Required, Create: `EXACS`},
		"source_major_version": acctest.Representation{RepType: acctest.Required, Create: `DB_19`},
		"type":                 acctest.Representation{RepType: acctest.Required, Create: `DB`},
		"display_name":         acctest.Representation{RepType: acctest.Optional, Create: `TF_TEST_Collection`, Update: `TF_TEST_Collection_Updated`},
		"fleet_discovery":      acctest.RepresentationGroup{RepType: acctest.Required, Group: fsuCollectionFleetDiscoveryDBRepresentation},
		"lifecycle":            acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreFsuCollectionDefinedTagsChangesRepresentation},
	}

	// https://docs.oracle.com/en-us/iaas/api/#/en/edsfu/20220528/datatypes/CreateGiFsuCollectionDetails
	FleetSoftwareUpdateFsuCollectionGIRepresentation = map[string]interface{}{
		"compartment_id":       acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"service_type":         acctest.Representation{RepType: acctest.Required, Create: `EXACS`},
		"source_major_version": acctest.Representation{RepType: acctest.Required, Create: `GI_19`},
		"type":                 acctest.Representation{RepType: acctest.Required, Create: `GI`},
		"display_name":         acctest.Representation{RepType: acctest.Optional, Create: `TF_TEST_Collection`, Update: `TF_TEST_Collection_Updated`},
		"fleet_discovery":      acctest.RepresentationGroup{RepType: acctest.Required, Group: fsuCollectionFleetDiscoveryGIRepresentation},
		"lifecycle":            acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreFsuCollectionDefinedTagsChangesRepresentation},
	}

	// https://docs.oracle.com/en-us/iaas/api/#/en/edsfu/20220528/datatypes/DbFleetDiscoveryDetails
	fsuCollectionFleetDiscoveryDBRepresentation = map[string]interface{}{
		"strategy": acctest.Representation{RepType: acctest.Required, Create: `TARGET_LIST`},
		"targets":  acctest.Representation{RepType: acctest.Required, Create: []string{`${var.db_target_1}`}},
	}

	// https://docs.oracle.com/en-us/iaas/api/#/en/edsfu/20220528/datatypes/GiFleetDiscoveryDetails
	fsuCollectionFleetDiscoveryGIRepresentation = map[string]interface{}{
		"strategy": acctest.Representation{RepType: acctest.Required, Create: `TARGET_LIST`},
		"targets":  acctest.Representation{RepType: acctest.Required, Create: []string{`${var.gi_target_1}`}},
	}

	FleetSoftwareUpdateFsuCollectionResourceDependencies = DefinedTagsDependencies
)

// issue-routing-tag: fleet_software_update/default
func TestFleetSoftwareUpdateFsuCollectionResource_DB_basic(t *testing.T) {
	httpreplay.SetScenario("TestFleetSoftwareUpdateFsuCollectionResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()
	//compartment_id
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

	resourceName := "oci_fleet_software_update_fsu_collection.test_fsu_collection"
	datasourceName := "data.oci_fleet_software_update_fsu_collections.test_fsu_collections"
	singularDatasourceName := "data.oci_fleet_software_update_fsu_collection.test_fsu_collection"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	var testConfig = config + variablesStr + FleetSoftwareUpdateFsuCollectionResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_fleet_software_update_fsu_collection", "test_fsu_collection", acctest.Optional,
			acctest.Create, FleetSoftwareUpdateFsuCollectionDBRepresentation)
	acctest.SaveConfigContent(testConfig, "fleetsoftwareupdate", "fsuCollection", t)
	fmt.Printf("FSU_TEST_LOG CONF:\n%s\n", testConfig)

	acctest.ResourceTest(t, testAccCheckFleetSoftwareUpdateFsuCollectionDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + variablesStr + FleetSoftwareUpdateFsuCollectionResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_fleet_software_update_fsu_collection", "test_fsu_collection", acctest.Required, acctest.Create, FleetSoftwareUpdateFsuCollectionDBRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "service_type", "EXACS"),
				resource.TestCheckResourceAttr(resourceName, "source_major_version", "DB_19"),
				resource.TestCheckResourceAttr(resourceName, "type", "DB"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + variablesStr + FleetSoftwareUpdateFsuCollectionResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + variablesStr + FleetSoftwareUpdateFsuCollectionResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_fleet_software_update_fsu_collection", "test_fsu_collection", acctest.Optional, acctest.Create, FleetSoftwareUpdateFsuCollectionDBRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "TF_TEST_Collection"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "service_type", "EXACS"),
				resource.TestCheckResourceAttr(resourceName, "source_major_version", "DB_19"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttr(resourceName, "type", "DB"),

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
			Config: config + variablesStr + compartmentIdUVariableStr + FleetSoftwareUpdateFsuCollectionResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_fleet_software_update_fsu_collection", "test_fsu_collection", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(FleetSoftwareUpdateFsuCollectionDBRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "display_name", "TF_TEST_Collection"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "service_type", "EXACS"),
				resource.TestCheckResourceAttr(resourceName, "source_major_version", "DB_19"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttr(resourceName, "type", "DB"),

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
			Config: config + variablesStr + FleetSoftwareUpdateFsuCollectionResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_fleet_software_update_fsu_collection", "test_fsu_collection", acctest.Optional, acctest.Update, FleetSoftwareUpdateFsuCollectionDBRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "TF_TEST_Collection_Updated"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "service_type", "EXACS"),
				resource.TestCheckResourceAttr(resourceName, "source_major_version", "DB_19"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttr(resourceName, "type", "DB"),

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
				acctest.GenerateDataSourceFromRepresentationMap("oci_fleet_software_update_fsu_collections", "test_fsu_collections", acctest.Optional, acctest.Update, FleetSoftwareUpdateFsuCollectionDBDataSourceRepresentation) +
				variablesStr + FleetSoftwareUpdateFsuCollectionResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_fleet_software_update_fsu_collection", "test_fsu_collection", acctest.Optional, acctest.Update, FleetSoftwareUpdateFsuCollectionDBRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "TF_TEST_Collection_Updated"),
				resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),
				resource.TestCheckResourceAttr(datasourceName, "type", "DB"),
				resource.TestCheckResourceAttr(datasourceName, "fsu_collection_summary_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "fsu_collection_summary_collection.0.items.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_fleet_software_update_fsu_collection", "test_fsu_collection", acctest.Required, acctest.Create, FleetSoftwareUpdateFsuCollectionSingularDataSourceRepresentation) +
				variablesStr + FleetSoftwareUpdateFsuCollectionDBResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "fsu_collection_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "active_fsu_cycle.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "TF_TEST_Collection_Updated"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "service_type", "EXACS"),
				resource.TestCheckResourceAttr(singularDatasourceName, "source_major_version", "DB_19"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "target_count"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
				resource.TestCheckResourceAttr(singularDatasourceName, "type", "DB"),
			),
		},
		// verify resource import
		{
			Config:                  config + FleetSoftwareUpdateFsuCollectionDBRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func TestFleetSoftwareUpdateFsuCollectionResource_GI_basic(t *testing.T) {
	httpreplay.SetScenario("TestFleetSoftwareUpdateFsuCollectionGIResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()
	//compartment_id
	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	giTargetId1 := utils.GetEnvSettingWithBlankDefault("fsu_gi_target_1")
	giTargetId1VariableStr := fmt.Sprintf("variable \"gi_target_1\" { default = \"%s\" }\n", giTargetId1)

	giTargetId2 := utils.GetEnvSettingWithBlankDefault("fsu_gi_target_2")
	giTargetId2VariableStr := fmt.Sprintf("variable \"gi_target_2\" { default = \"%s\" }\n", giTargetId2)

	var variablesStr = compartmentIdVariableStr + giTargetId1VariableStr + giTargetId2VariableStr

	resourceName := "oci_fleet_software_update_fsu_collection.test_fsu_collection"
	datasourceName := "data.oci_fleet_software_update_fsu_collections.test_fsu_collections"
	singularDatasourceName := "data.oci_fleet_software_update_fsu_collection.test_fsu_collection"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	var testConfig = config + variablesStr + FleetSoftwareUpdateFsuCollectionResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_fleet_software_update_fsu_collection", "test_fsu_collection",
			acctest.Optional, acctest.Create, FleetSoftwareUpdateFsuCollectionGIRepresentation)
	acctest.SaveConfigContent(testConfig, "fleetsoftwareupdate", "fsuCollection", t)
	fmt.Printf("FSU_TEST_LOG CONF:\n%s\n", testConfig)

	acctest.ResourceTest(t, testAccCheckFleetSoftwareUpdateFsuCollectionDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + variablesStr + FleetSoftwareUpdateFsuCollectionResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_fleet_software_update_fsu_collection", "test_fsu_collection",
					acctest.Required, acctest.Create, FleetSoftwareUpdateFsuCollectionGIRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "service_type", "EXACS"),
				resource.TestCheckResourceAttr(resourceName, "source_major_version", "GI_19"),
				resource.TestCheckResourceAttr(resourceName, "type", "GI"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + variablesStr + FleetSoftwareUpdateFsuCollectionResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + variablesStr + FleetSoftwareUpdateFsuCollectionResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_fleet_software_update_fsu_collection", "test_fsu_collection", acctest.Optional, acctest.Create, FleetSoftwareUpdateFsuCollectionGIRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "TF_TEST_Collection"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "service_type", "EXACS"),
				resource.TestCheckResourceAttr(resourceName, "source_major_version", "GI_19"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttr(resourceName, "type", "GI"),

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
			Config: config + variablesStr + compartmentIdUVariableStr + FleetSoftwareUpdateFsuCollectionResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_fleet_software_update_fsu_collection", "test_fsu_collection", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(FleetSoftwareUpdateFsuCollectionGIRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "display_name", "TF_TEST_Collection"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "service_type", "EXACS"),
				resource.TestCheckResourceAttr(resourceName, "source_major_version", "GI_19"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttr(resourceName, "type", "GI"),

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
			Config: config + variablesStr + FleetSoftwareUpdateFsuCollectionResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_fleet_software_update_fsu_collection", "test_fsu_collection", acctest.Optional, acctest.Update, FleetSoftwareUpdateFsuCollectionGIRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "TF_TEST_Collection_Updated"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "service_type", "EXACS"),
				resource.TestCheckResourceAttr(resourceName, "source_major_version", "GI_19"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttr(resourceName, "type", "GI"),

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
				acctest.GenerateDataSourceFromRepresentationMap("oci_fleet_software_update_fsu_collections", "test_fsu_collections", acctest.Optional, acctest.Update, FleetSoftwareUpdateFsuCollectionGIDataSourceRepresentation) +
				variablesStr + FleetSoftwareUpdateFsuCollectionResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_fleet_software_update_fsu_collection", "test_fsu_collection", acctest.Optional, acctest.Update, FleetSoftwareUpdateFsuCollectionGIRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "TF_TEST_Collection_Updated"),
				resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),
				resource.TestCheckResourceAttr(datasourceName, "type", "GI"),
				resource.TestCheckResourceAttr(datasourceName, "fsu_collection_summary_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "fsu_collection_summary_collection.0.items.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_fleet_software_update_fsu_collection", "test_fsu_collection", acctest.Required, acctest.Create, FleetSoftwareUpdateFsuCollectionSingularDataSourceRepresentation) +
				variablesStr + FleetSoftwareUpdateFsuCollectionGIResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "fsu_collection_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "active_fsu_cycle.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "TF_TEST_Collection_Updated"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "service_type", "EXACS"),
				resource.TestCheckResourceAttr(singularDatasourceName, "source_major_version", "GI_19"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "target_count"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
				resource.TestCheckResourceAttr(singularDatasourceName, "type", "GI"),
			),
		},
		// verify resource import
		{
			Config:                  config + FleetSoftwareUpdateFsuCollectionGIRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckFleetSoftwareUpdateFsuCollectionDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).FleetSoftwareUpdateClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_fleet_software_update_fsu_collection" {
			noResourceFound = false
			request := oci_fleet_software_update.GetFsuCollectionRequest{}

			tmp := rs.Primary.ID
			request.FsuCollectionId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "fleet_software_update")

			response, err := client.GetFsuCollection(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_fleet_software_update.CollectionLifecycleStatesDeleted): true,
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
	if !acctest.InSweeperExcludeList("FleetSoftwareUpdateFsuCollection") {
		resource.AddTestSweepers("FleetSoftwareUpdateFsuCollection", &resource.Sweeper{
			Name:         "FleetSoftwareUpdateFsuCollection",
			Dependencies: acctest.DependencyGraph["fsuCollection"],
			F:            sweepFleetSoftwareUpdateFsuCollectionResource,
		})
	}
}

func sweepFleetSoftwareUpdateFsuCollectionResource(compartment string) error {
	fleetSoftwareUpdateClient := acctest.GetTestClients(&schema.ResourceData{}).FleetSoftwareUpdateClient()
	fsuCollectionIds, err := getFleetSoftwareUpdateFsuCollectionIds(compartment)
	if err != nil {
		return err
	}
	for _, fsuCollectionId := range fsuCollectionIds {
		if ok := acctest.SweeperDefaultResourceId[fsuCollectionId]; !ok {
			deleteFsuCollectionRequest := oci_fleet_software_update.DeleteFsuCollectionRequest{}

			deleteFsuCollectionRequest.FsuCollectionId = &fsuCollectionId

			deleteFsuCollectionRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "fleet_software_update")
			_, error := fleetSoftwareUpdateClient.DeleteFsuCollection(context.Background(), deleteFsuCollectionRequest)
			if error != nil {
				fmt.Printf("Error deleting FsuCollection %s %s, It is possible that the resource is already deleted. Please verify manually \n", fsuCollectionId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &fsuCollectionId, FleetSoftwareUpdateFsuCollectionSweepWaitCondition, time.Duration(3*time.Minute),
				FleetSoftwareUpdateFsuCollectionSweepResponseFetchOperation, "fleet_software_update", true)
		}
	}
	return nil
}

func getFleetSoftwareUpdateFsuCollectionIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "FsuCollectionId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	fleetSoftwareUpdateClient := acctest.GetTestClients(&schema.ResourceData{}).FleetSoftwareUpdateClient()

	listFsuCollectionsRequest := oci_fleet_software_update.ListFsuCollectionsRequest{}
	listFsuCollectionsRequest.CompartmentId = &compartmentId
	listFsuCollectionsRequest.LifecycleState = oci_fleet_software_update.ListFsuCollectionsLifecycleStateNeedsAttention
	listFsuCollectionsResponse, err := fleetSoftwareUpdateClient.ListFsuCollections(context.Background(), listFsuCollectionsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting FsuCollection list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, fsuCollection := range listFsuCollectionsResponse.Items {
		id := *fsuCollection.GetId()
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "FsuCollectionId", id)
	}
	return resourceIds, nil
}

func FleetSoftwareUpdateFsuCollectionSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if fsuCollectionResponse, ok := response.Response.(oci_fleet_software_update.GetFsuCollectionResponse); ok {
		return fsuCollectionResponse.GetLifecycleState() != oci_fleet_software_update.CollectionLifecycleStatesDeleted
	}
	return false
}

func FleetSoftwareUpdateFsuCollectionSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.FleetSoftwareUpdateClient().GetFsuCollection(context.Background(), oci_fleet_software_update.GetFsuCollectionRequest{
		FsuCollectionId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
