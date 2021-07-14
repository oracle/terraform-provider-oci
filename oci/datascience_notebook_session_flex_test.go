// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	NotebookSessionResourceFlexConfig = NotebookSessionResourceDependencies +
		generateResourceFromRepresentationMap("oci_datascience_notebook_session", "test_notebook_session", Optional, Update, notebookSessionFlexRepresentation)

	notebookSessionNotebookSessionConfigurationDetailsNotebookSessionShapeConfigDetailsRepresentation = map[string]interface{}{
		"memory_in_gbs": Representation{repType: Required, create: `16.0`, update: `64.0`},
		"ocpus":         Representation{repType: Required, create: `1.0`, update: `4.0`},
	}
	notebookSessionFlexRepresentation = getUpdatedRepresentationCopy("notebook_session_configuration_details", RepresentationGroup{Required,
		representationCopyWithNewProperties(notebookSessionNotebookSessionConfigurationDetailsRepresentation, map[string]interface{}{
			"shape":                                 Representation{repType: Required, create: `VM.Standard.E3.Flex`},
			"subnet_id":                             Representation{repType: Required, create: `${oci_core_subnet.test_subnet.id}`},
			"block_storage_size_in_gbs":             Representation{repType: Optional, create: `50`, update: `51`},
			"notebook_session_shape_config_details": RepresentationGroup{Required, notebookSessionNotebookSessionConfigurationDetailsNotebookSessionShapeConfigDetailsRepresentation}})},
		notebookSessionRepresentation)
)

func TestDatascienceNotebookSessionResource_flex(t *testing.T) {
	httpreplay.SetScenario("TestDatascienceNotebookSessionResource_flex")
	defer httpreplay.SaveScenario()

	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := getEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_datascience_notebook_session.test_notebook_session"
	datasourceName := "data.oci_datascience_notebook_sessions.test_notebook_sessions"
	singularDatasourceName := "data.oci_datascience_notebook_session.test_notebook_session"

	var resId, resId2 string
	// Save TF content to create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	saveConfigContent(config+compartmentIdVariableStr+NotebookSessionResourceDependencies+
		generateResourceFromRepresentationMap("oci_datascience_notebook_session", "test_notebook_session", Optional, Create, notebookSessionFlexRepresentation), "datascience", "notebookSession", t)

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		CheckDestroy: testAccCheckDatascienceNotebookSessionDestroy,
		Steps: []resource.TestStep{
			// verify create
			{
				Config: config + compartmentIdVariableStr + NotebookSessionResourceDependencies +
					generateResourceFromRepresentationMap("oci_datascience_notebook_session", "test_notebook_session", Required, Create, notebookSessionFlexRepresentation),
				Check: ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "notebook_session_configuration_details.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "notebook_session_configuration_details.0.shape", "VM.Standard.E3.Flex"),
					resource.TestCheckResourceAttr(resourceName, "notebook_session_configuration_details.0.notebook_session_shape_config_details.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "notebook_session_configuration_details.0.notebook_session_shape_config_details.0.memory_in_gbs", "16"),
					resource.TestCheckResourceAttr(resourceName, "notebook_session_configuration_details.0.notebook_session_shape_config_details.0.ocpus", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "notebook_session_configuration_details.0.subnet_id"),
					resource.TestCheckResourceAttrSet(resourceName, "project_id"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},

			// delete before next create
			{
				Config: config + compartmentIdVariableStr + NotebookSessionResourceDependencies,
			},
			// verify create with optionals
			{
				Config: config + compartmentIdVariableStr + NotebookSessionResourceDependencies +
					generateResourceFromRepresentationMap("oci_datascience_notebook_session", "test_notebook_session", Optional, Create, notebookSessionFlexRepresentation),
				Check: ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttrSet(resourceName, "created_by"),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "notebook_session_configuration_details.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "notebook_session_configuration_details.0.block_storage_size_in_gbs", "50"),
					resource.TestCheckResourceAttr(resourceName, "notebook_session_configuration_details.0.notebook_session_shape_config_details.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "notebook_session_configuration_details.0.notebook_session_shape_config_details.0.memory_in_gbs", "16"),
					resource.TestCheckResourceAttr(resourceName, "notebook_session_configuration_details.0.notebook_session_shape_config_details.0.ocpus", "1"),
					resource.TestCheckResourceAttr(resourceName, "notebook_session_configuration_details.0.shape", "VM.Standard.E3.Flex"),
					resource.TestCheckResourceAttrSet(resourceName, "notebook_session_configuration_details.0.subnet_id"),
					resource.TestCheckResourceAttrSet(resourceName, "project_id"),
					resource.TestCheckResourceAttrSet(resourceName, "state"),
					resource.TestCheckResourceAttrSet(resourceName, "time_created"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						if isEnableExportCompartment, _ := strconv.ParseBool(getEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
							if errExport := testExportCompartmentWithResourceName(&resId, &compartmentId, resourceName); errExport != nil {
								return errExport
							}
						}
						return err
					},
				),
			},

			// verify update to the compartment (the compartment will be switched back in the next step)
			{
				Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + NotebookSessionResourceDependencies +
					generateResourceFromRepresentationMap("oci_datascience_notebook_session", "test_notebook_session", Optional, Create,
						representationCopyWithNewProperties(notebookSessionFlexRepresentation, map[string]interface{}{
							"compartment_id": Representation{repType: Required, create: `${var.compartment_id_for_update}`},
						})),
				Check: ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
					resource.TestCheckResourceAttrSet(resourceName, "created_by"),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "notebook_session_configuration_details.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "notebook_session_configuration_details.0.block_storage_size_in_gbs", "50"),
					resource.TestCheckResourceAttr(resourceName, "notebook_session_configuration_details.0.notebook_session_shape_config_details.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "notebook_session_configuration_details.0.notebook_session_shape_config_details.0.memory_in_gbs", "16"),
					resource.TestCheckResourceAttr(resourceName, "notebook_session_configuration_details.0.notebook_session_shape_config_details.0.ocpus", "1"),
					resource.TestCheckResourceAttr(resourceName, "notebook_session_configuration_details.0.shape", "VM.Standard.E3.Flex"),
					resource.TestCheckResourceAttrSet(resourceName, "notebook_session_configuration_details.0.subnet_id"),
					resource.TestCheckResourceAttrSet(resourceName, "project_id"),
					resource.TestCheckResourceAttrSet(resourceName, "state"),
					resource.TestCheckResourceAttrSet(resourceName, "time_created"),

					func(s *terraform.State) (err error) {
						resId2, err = fromInstanceState(s, resourceName, "id")
						if resId != resId2 {
							return fmt.Errorf("resource recreated when it was supposed to be updated")
						}
						return err
					},
				),
			},

			// verify updates to updatable parameters
			{
				Config: config + compartmentIdVariableStr + NotebookSessionResourceDependencies +
					generateResourceFromRepresentationMap("oci_datascience_notebook_session", "test_notebook_session", Optional, Update, notebookSessionFlexRepresentation),
				Check: ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttrSet(resourceName, "created_by"),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "notebook_session_configuration_details.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "notebook_session_configuration_details.0.block_storage_size_in_gbs", "51"),
					resource.TestCheckResourceAttr(resourceName, "notebook_session_configuration_details.0.notebook_session_shape_config_details.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "notebook_session_configuration_details.0.notebook_session_shape_config_details.0.memory_in_gbs", "64"),
					resource.TestCheckResourceAttr(resourceName, "notebook_session_configuration_details.0.notebook_session_shape_config_details.0.ocpus", "4"),
					resource.TestCheckResourceAttr(resourceName, "notebook_session_configuration_details.0.shape", "VM.Standard.E3.Flex"),
					resource.TestCheckResourceAttrSet(resourceName, "notebook_session_configuration_details.0.subnet_id"),
					resource.TestCheckResourceAttrSet(resourceName, "project_id"),
					resource.TestCheckResourceAttrSet(resourceName, "state"),
					resource.TestCheckResourceAttrSet(resourceName, "time_created"),

					func(s *terraform.State) (err error) {
						resId2, err = fromInstanceState(s, resourceName, "id")
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
					generateDataSourceFromRepresentationMap("oci_datascience_notebook_sessions", "test_notebook_sessions", Optional, Update, notebookSessionDataSourceRepresentation) +
					compartmentIdVariableStr + NotebookSessionResourceDependencies +
					generateResourceFromRepresentationMap("oci_datascience_notebook_session", "test_notebook_session", Optional, Update, notebookSessionFlexRepresentation),
				Check: ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
					resource.TestCheckResourceAttrSet(datasourceName, "id"),
					resource.TestCheckResourceAttrSet(datasourceName, "project_id"),
					resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),
					resource.TestCheckResourceAttr(datasourceName, "notebook_sessions.#", "1"),
					resource.TestCheckResourceAttr(datasourceName, "notebook_sessions.0.compartment_id", compartmentId),
					resource.TestCheckResourceAttrSet(datasourceName, "notebook_sessions.0.created_by"),
					resource.TestCheckResourceAttr(datasourceName, "notebook_sessions.0.defined_tags.%", "1"),
					resource.TestCheckResourceAttr(datasourceName, "notebook_sessions.0.display_name", "displayName2"),
					resource.TestCheckResourceAttr(datasourceName, "notebook_sessions.0.freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(datasourceName, "notebook_sessions.0.id"),
					resource.TestCheckResourceAttr(datasourceName, "notebook_sessions.0.notebook_session_configuration_details.#", "1"),
					resource.TestCheckResourceAttr(datasourceName, "notebook_sessions.0.notebook_session_configuration_details.0.block_storage_size_in_gbs", "51"),
					resource.TestCheckResourceAttr(datasourceName, "notebook_sessions.0.notebook_session_configuration_details.0.notebook_session_shape_config_details.#", "1"),
					resource.TestCheckResourceAttr(datasourceName, "notebook_sessions.0.notebook_session_configuration_details.0.notebook_session_shape_config_details.0.memory_in_gbs", "64"),
					resource.TestCheckResourceAttr(datasourceName, "notebook_sessions.0.notebook_session_configuration_details.0.notebook_session_shape_config_details.0.ocpus", "4"),
					resource.TestCheckResourceAttr(datasourceName, "notebook_sessions.0.notebook_session_configuration_details.0.shape", "VM.Standard.E3.Flex"),
					resource.TestCheckResourceAttrSet(datasourceName, "notebook_sessions.0.notebook_session_configuration_details.0.subnet_id"),
					resource.TestCheckResourceAttrSet(datasourceName, "notebook_sessions.0.notebook_session_url"),
					resource.TestCheckResourceAttrSet(datasourceName, "notebook_sessions.0.project_id"),
					resource.TestCheckResourceAttrSet(datasourceName, "notebook_sessions.0.state"),
					resource.TestCheckResourceAttrSet(datasourceName, "notebook_sessions.0.time_created"),
				),
			},
			// verify singular datasource
			{
				Config: config +
					generateDataSourceFromRepresentationMap("oci_datascience_notebook_session", "test_notebook_session", Required, Create, notebookSessionSingularDataSourceRepresentation) +
					compartmentIdVariableStr + NotebookSessionResourceFlexConfig,
				Check: ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttrSet(singularDatasourceName, "notebook_session_id"),
					resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "created_by"),
					resource.TestCheckResourceAttr(singularDatasourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
					resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
					resource.TestCheckResourceAttr(singularDatasourceName, "notebook_session_configuration_details.#", "1"),
					resource.TestCheckResourceAttr(singularDatasourceName, "notebook_session_configuration_details.0.block_storage_size_in_gbs", "51"),
					resource.TestCheckResourceAttr(singularDatasourceName, "notebook_session_configuration_details.0.notebook_session_shape_config_details.#", "1"),
					resource.TestCheckResourceAttr(singularDatasourceName, "notebook_session_configuration_details.0.notebook_session_shape_config_details.0.memory_in_gbs", "64"),
					resource.TestCheckResourceAttr(singularDatasourceName, "notebook_session_configuration_details.0.notebook_session_shape_config_details.0.ocpus", "4"),
					resource.TestCheckResourceAttr(singularDatasourceName, "notebook_session_configuration_details.0.shape", "VM.Standard.E3.Flex"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "notebook_session_url"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				),
			},
			// remove singular datasource from previous step so that it doesn't conflict with import tests
			{
				Config: config + compartmentIdVariableStr + NotebookSessionResourceFlexConfig,
			},
			// verify resource import
			{
				Config:                  config,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
				ResourceName:            resourceName,
			},
		},
	})
}

func init() {
	if DependencyGraph == nil {
		initDependencyGraph()
	}
	if !inSweeperExcludeList("DatascienceFlexNotebookSession") {
		resource.AddTestSweepers("DatascienceFlexNotebookSession", &resource.Sweeper{
			Name:         "DatascienceFlexNotebookSession",
			Dependencies: DependencyGraph["notebookSession"],
			F:            sweepDatascienceNotebookSessionResource,
		})
	}
}
