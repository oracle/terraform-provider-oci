// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/terraform-providers/terraform-provider-oci/internal/acctest"
	"github.com/terraform-providers/terraform-provider-oci/internal/resourcediscovery"
	"github.com/terraform-providers/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	NotebookSessionResourceFlexConfig = NotebookSessionResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_datascience_notebook_session", "test_notebook_session", acctest.Optional, acctest.Update, notebookSessionFlexRepresentation)

	notebookSessionNotebookSessionConfigurationDetailsNotebookSessionShapeConfigDetailsRepresentation = map[string]interface{}{
		"memory_in_gbs": acctest.Representation{RepType: acctest.Required, Create: `16.0`, Update: `64.0`},
		"ocpus":         acctest.Representation{RepType: acctest.Required, Create: `1.0`, Update: `4.0`},
	}
	notebookSessionFlexRepresentation = acctest.GetUpdatedRepresentationCopy("notebook_session_configuration_details", acctest.RepresentationGroup{RepType: acctest.Required, Group: acctest.RepresentationCopyWithNewProperties(notebookSessionNotebookSessionConfigurationDetailsRepresentation, map[string]interface{}{
		"shape":                                 acctest.Representation{RepType: acctest.Required, Create: `VM.Standard.E3.Flex`},
		"subnet_id":                             acctest.Representation{RepType: acctest.Required, Create: `${oci_core_subnet.test_subnet.id}`},
		"block_storage_size_in_gbs":             acctest.Representation{RepType: acctest.Optional, Create: `50`, Update: `51`},
		"notebook_session_shape_config_details": acctest.RepresentationGroup{RepType: acctest.Required, Group: notebookSessionNotebookSessionConfigurationDetailsNotebookSessionShapeConfigDetailsRepresentation}})},
		notebookSessionRepresentation)
)

// issue-routing-tag: datascience/default
func TestDatascienceNotebookSessionResource_flex(t *testing.T) {
	httpreplay.SetScenario("TestDatascienceNotebookSessionResource_flex")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_datascience_notebook_session.test_notebook_session"
	datasourceName := "data.oci_datascience_notebook_sessions.test_notebook_sessions"
	singularDatasourceName := "data.oci_datascience_notebook_session.test_notebook_session"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+NotebookSessionResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_datascience_notebook_session", "test_notebook_session", acctest.Optional, acctest.Create, notebookSessionFlexRepresentation), "datascience", "notebookSession", t)

	acctest.ResourceTest(t, testAccCheckDatascienceNotebookSessionDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + NotebookSessionResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_datascience_notebook_session", "test_notebook_session", acctest.Required, acctest.Create, notebookSessionFlexRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "notebook_session_configuration_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "notebook_session_configuration_details.0.shape", "VM.Standard.E3.Flex"),
				resource.TestCheckResourceAttr(resourceName, "notebook_session_configuration_details.0.notebook_session_shape_config_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "notebook_session_configuration_details.0.notebook_session_shape_config_details.0.memory_in_gbs", "16"),
				resource.TestCheckResourceAttr(resourceName, "notebook_session_configuration_details.0.notebook_session_shape_config_details.0.ocpus", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "notebook_session_configuration_details.0.subnet_id"),
				resource.TestCheckResourceAttrSet(resourceName, "project_id"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + NotebookSessionResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + NotebookSessionResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_datascience_notebook_session", "test_notebook_session", acctest.Optional, acctest.Create, notebookSessionFlexRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName, "created_by"),
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
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + NotebookSessionResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_datascience_notebook_session", "test_notebook_session", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(notebookSessionFlexRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttrSet(resourceName, "created_by"),
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
			Config: config + compartmentIdVariableStr + NotebookSessionResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_datascience_notebook_session", "test_notebook_session", acctest.Optional, acctest.Update, notebookSessionFlexRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName, "created_by"),
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
				acctest.GenerateDataSourceFromRepresentationMap("oci_datascience_notebook_sessions", "test_notebook_sessions", acctest.Optional, acctest.Update, notebookSessionDataSourceRepresentation) +
				compartmentIdVariableStr + NotebookSessionResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_datascience_notebook_session", "test_notebook_session", acctest.Optional, acctest.Update, notebookSessionFlexRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttrSet(datasourceName, "id"),
				resource.TestCheckResourceAttrSet(datasourceName, "project_id"),
				resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),
				resource.TestCheckResourceAttr(datasourceName, "notebook_sessions.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "notebook_sessions.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(datasourceName, "notebook_sessions.0.created_by"),
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
				acctest.GenerateDataSourceFromRepresentationMap("oci_datascience_notebook_session", "test_notebook_session", acctest.Required, acctest.Create, notebookSessionSingularDataSourceRepresentation) +
				compartmentIdVariableStr + NotebookSessionResourceFlexConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "notebook_session_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "created_by"),
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
	})
}

func init() {
	if acctest.DependencyGraph == nil {
		acctest.InitDependencyGraph()
	}
	if !acctest.InSweeperExcludeList("DatascienceFlexNotebookSession") {
		resource.AddTestSweepers("DatascienceFlexNotebookSession", &resource.Sweeper{
			Name:         "DatascienceFlexNotebookSession",
			Dependencies: acctest.DependencyGraph["notebookSession"],
			F:            sweepDatascienceNotebookSessionResource,
		})
	}
}
