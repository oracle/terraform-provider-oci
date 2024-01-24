// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"

	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/oracle/oci-go-sdk/v65/common"
	oci_datascience "github.com/oracle/oci-go-sdk/v65/datascience"

	"github.com/oracle/terraform-provider-oci/httpreplay"
)

var (
	DatascienceNotebookSessionRequiredOnlyResource = DatascienceNotebookSessionResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_datascience_notebook_session", "test_notebook_session", acctest.Required, acctest.Create, DatascienceNotebookSessionRepresentation)

	NotebookSessionConfigurationDetailsRequiredOnlyResource = NotebookSessionConfigurationDetailsResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_datascience_notebook_session", "test_notebook_session", acctest.Required, acctest.Create, notebookSessionConfigurationDetailsRepresentation)

	DatascienceNotebookSessionResourceConfig = DatascienceNotebookSessionResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_datascience_notebook_session", "test_notebook_session", acctest.Optional, acctest.Update, DatascienceNotebookSessionRepresentation)

	NotebookSessionConfigurationDetailsResourceConfig = NotebookSessionConfigurationDetailsResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_datascience_notebook_session", "test_notebook_session", acctest.Optional, acctest.Update, notebookSessionConfigurationDetailsRepresentation)

	notebookSessionConfigDetailsSingularDataSourceRepresentation = map[string]interface{}{
		"notebook_session_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_datascience_notebook_session.test_notebook_session.id}`},
	}

	notebookSessionConfigurationDetailsSingularDataSourceRepresentation = map[string]interface{}{
		"notebook_session_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_datascience_notebook_session.test_notebook_session.id}`},
	}

	DatascienceDatascienceNotebookSessionDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":   acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"id":             acctest.Representation{RepType: acctest.Optional, Create: `${oci_datascience_notebook_session.test_notebook_session.id}`},
		"project_id":     acctest.Representation{RepType: acctest.Optional, Create: `${oci_datascience_project.test_project.id}`},
		"state":          acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"filter":         acctest.RepresentationGroup{RepType: acctest.Required, Group: DatascienceNotebookSessionDataSourceFilterRepresentation},
	}

	notebookSessionConfigurationDetailsDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":   acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"id":             acctest.Representation{RepType: acctest.Optional, Create: `${oci_datascience_notebook_session.test_notebook_session.id}`},
		"project_id":     acctest.Representation{RepType: acctest.Optional, Create: `${oci_datascience_project.test_project.id}`},
		"state":          acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"filter":         acctest.RepresentationGroup{RepType: acctest.Required, Group: notebookSessionConfigurationDetailsDataSourceFilterRepresentation},
	}

	DatascienceNotebookSessionDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_datascience_notebook_session.test_notebook_session.id}`}},
	}

	notebookSessionConfigurationDetailsDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_datascience_notebook_session.test_notebook_session.id}`}},
	}

	DatascienceNotebookSessionRepresentation = map[string]interface{}{
		"compartment_id":                  acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"project_id":                      acctest.Representation{RepType: acctest.Required, Create: `${oci_datascience_project.test_project.id}`},
		"defined_tags":                    acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"display_name":                    acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"freeform_tags":                   acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"notebook_session_config_details": acctest.RepresentationGroup{RepType: acctest.Required, Group: notebookSessionConfigDetailsRepresentation},
		"lifecycle":                       acctest.RepresentationGroup{RepType: acctest.Optional, Group: definedTagsIgnoreRepresentation},
	}

	notebookSessionConfigDetailsRepresentation = map[string]interface{}{
		"shape":                                 acctest.Representation{RepType: acctest.Required, Create: `VM.Standard2.1`},
		"block_storage_size_in_gbs":             acctest.Representation{RepType: acctest.Optional, Create: `100`},
		"notebook_session_shape_config_details": acctest.RepresentationGroup{RepType: acctest.Optional, Group: notebookSessionConfigDetailsShapeConfigDetailsRepresentation},
		"private_endpoint_id":                   acctest.Representation{RepType: acctest.Required, Create: `${oci_datascience_private_endpoint.test_data_science_private_endpoint.id}`},
		"subnet_id":                             acctest.Representation{RepType: acctest.Required, Create: `${oci_core_subnet.test_subnet.id}`},
	}

	notebookSessionConfigurationDetailsRepresentation = map[string]interface{}{
		"compartment_id":                         acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"notebook_session_configuration_details": acctest.RepresentationGroup{RepType: acctest.Required, Group: DatascienceNotebookSessionConfigurationDetailsRepresentation},
		"project_id":                             acctest.Representation{RepType: acctest.Required, Create: `${oci_datascience_project.test_project.id}`},
		"defined_tags":                           acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"display_name":                           acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"freeform_tags":                          acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"lifecycle":                              acctest.RepresentationGroup{RepType: acctest.Optional, Group: definedTagsIgnoreRepresentation},
	}

	DatascienceNotebookSessionConfigurationDetailsRepresentation = map[string]interface{}{
		"shape":                     acctest.Representation{RepType: acctest.Required, Create: `VM.Standard2.1`},
		"subnet_id":                 acctest.Representation{RepType: acctest.Required, Create: `${oci_core_subnet.test_subnet.id}`},
		"block_storage_size_in_gbs": acctest.Representation{RepType: acctest.Optional, Create: `100`},
		"private_endpoint_id":       acctest.Representation{RepType: acctest.Required, Create: `${oci_datascience_private_endpoint.test_data_science_private_endpoint.id}`},
	}

	notebookSessionConfigDetailsShapeConfigDetailsRepresentation = map[string]interface{}{
		"memory_in_gbs": acctest.Representation{RepType: acctest.Optional, Create: `1.0`},
		"ocpus":         acctest.Representation{RepType: acctest.Optional, Create: `1.0`},
	}

	definedTagsIgnoreRepresentation = map[string]interface{}{
		"ignore_changes": acctest.Representation{RepType: acctest.Required, Create: []string{`defined_tags`, `compartment_id`, `lifecycle_details`, `notebook_session_configuration_details[0].notebook_session_shape_config_details`, `notebook_session_config_details[0].notebook_session_shape_config_details`, `notebook_session_runtime_config_details`}},
	}

	DatascienceNotebookSessionResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_datascience_private_endpoint", "test_data_science_private_endpoint", acctest.Required, acctest.Create, DataSciencePrivateEndpointRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_subnet", "test_subnet", acctest.Required, acctest.Create, CoreSubnetRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", acctest.Required, acctest.Create, CoreVcnRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_datascience_project", "test_project", acctest.Required, acctest.Create, DatascienceProjectRepresentation) +
		DefinedTagsDependencies

	NotebookSessionConfigurationDetailsResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_datascience_private_endpoint", "test_data_science_private_endpoint", acctest.Required, acctest.Create, DataSciencePrivateEndpointRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_subnet", "test_subnet", acctest.Required, acctest.Create, CoreSubnetRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", acctest.Required, acctest.Create, CoreVcnRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_datascience_project", "test_project", acctest.Required, acctest.Create, DatascienceProjectRepresentation) +
		AvailabilityDomainConfig +
		DefinedTagsDependencies
)

// issue-routing-tag: datascience/default
func TestDatascienceNotebookSessionWithConfigDetailsResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatascienceNotebookSessionWithConfigDetailsResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_datascience_notebook_session.test_notebook_session"
	singularDatasourceName := "data.oci_datascience_notebook_session.test_notebook_session"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+DatascienceNotebookSessionResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_datascience_notebook_session", "test_notebook_session", acctest.Optional, acctest.Create, DatascienceNotebookSessionRepresentation), "datascience", "notebookSession", t)

	acctest.ResourceTest(t, testAccCheckDatascienceNotebookSessionDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + DatascienceNotebookSessionResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_datascience_notebook_session", "test_notebook_session", acctest.Required, acctest.Create, DatascienceNotebookSessionRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName, "project_id"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + DatascienceNotebookSessionResourceDependencies,
		},
		// verify Create with NotebookSessionConfigDetails with optionals
		{
			Config: config + compartmentIdVariableStr + DatascienceNotebookSessionResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_datascience_notebook_session", "test_notebook_session", acctest.Optional, acctest.Create, DatascienceNotebookSessionRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName, "created_by"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "notebook_session_config_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "notebook_session_config_details.0.block_storage_size_in_gbs", "100"),
				resource.TestCheckResourceAttrSet(resourceName, "notebook_session_config_details.0.private_endpoint_id"),
				resource.TestCheckResourceAttr(resourceName, "notebook_session_config_details.0.shape", "VM.Standard2.1"),
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
		// verify Create with NotebookSessionConfigurationDetails with optionals
		{
			Config: config + compartmentIdVariableStr + DatascienceNotebookSessionResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_datascience_notebook_session", "test_notebook_session", acctest.Optional, acctest.Create, DatascienceNotebookSessionRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "created_by"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
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
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + DatascienceNotebookSessionResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_datascience_notebook_session", "test_notebook_session", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(DatascienceNotebookSessionRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName, "created_by"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "notebook_session_config_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "notebook_session_config_details.0.block_storage_size_in_gbs", "100"),
				resource.TestCheckResourceAttr(resourceName, "notebook_session_config_details.0.notebook_session_shape_config_details.#", "0"),
				resource.TestCheckResourceAttrSet(resourceName, "notebook_session_config_details.0.private_endpoint_id"),
				resource.TestCheckResourceAttr(resourceName, "notebook_session_config_details.0.shape", "VM.Standard2.1"),
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
			Config: config + compartmentIdVariableStr + DatascienceNotebookSessionResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_datascience_notebook_session", "test_notebook_session", acctest.Optional, acctest.Update, DatascienceNotebookSessionRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName, "created_by"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "notebook_session_config_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "notebook_session_config_details.0.block_storage_size_in_gbs", "100"),
				resource.TestCheckResourceAttr(resourceName, "notebook_session_config_details.0.notebook_session_shape_config_details.#", "0"),
				resource.TestCheckResourceAttrSet(resourceName, "notebook_session_config_details.0.private_endpoint_id"),
				resource.TestCheckResourceAttr(resourceName, "notebook_session_config_details.0.shape", "VM.Standard2.1"),
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
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_datascience_notebook_session", "test_notebook_session", acctest.Required, acctest.Create, notebookSessionConfigDetailsSingularDataSourceRepresentation) +
				compartmentIdVariableStr + DatascienceNotebookSessionResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "notebook_session_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "created_by"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "notebook_session_config_details.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "notebook_session_config_details.0.block_storage_size_in_gbs", "100"),
				resource.TestCheckResourceAttr(singularDatasourceName, "notebook_session_config_details.0.notebook_session_shape_config_details.#", "0"),
				resource.TestCheckResourceAttr(singularDatasourceName, "notebook_session_config_details.0.shape", "VM.Standard2.1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "notebook_session_url"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
			),
		},
		// verify resource import
		{
			Config:                  config + DatascienceNotebookSessionRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

// issue-routing-tag: datascience/default
func TestDatascienceNotebookSessionWithConfigurationDetailsResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatascienceNotebookSessionWithConfigurationDetailsResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_datascience_notebook_session.test_notebook_session"
	datasourceName := "data.oci_datascience_notebook_sessions.test_notebook_session"
	singularDatasourceName := "data.oci_datascience_notebook_session.test_notebook_session"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+NotebookSessionConfigurationDetailsResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_datascience_notebook_session", "test_notebook_session", acctest.Optional, acctest.Create, notebookSessionConfigurationDetailsRepresentation), "datascience", "notebookSession", t)

	acctest.ResourceTest(t, testAccCheckDatascienceNotebookSessionDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + NotebookSessionConfigurationDetailsResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_datascience_notebook_session", "test_notebook_session", acctest.Required, acctest.Create, notebookSessionConfigurationDetailsRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName, "project_id"),

				func(s *terraform.State) (err error) {
					_, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},
		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + NotebookSessionConfigurationDetailsResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + NotebookSessionConfigurationDetailsResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_datascience_notebook_session", "test_notebook_session", acctest.Optional, acctest.Create, notebookSessionConfigurationDetailsRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName, "created_by"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
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
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + NotebookSessionConfigurationDetailsResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_datascience_notebook_session", "test_notebook_session", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(notebookSessionConfigurationDetailsRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName, "created_by"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
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
			Config: config + compartmentIdVariableStr + NotebookSessionConfigurationDetailsResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_datascience_notebook_session", "test_notebook_session", acctest.Optional, acctest.Update, notebookSessionConfigurationDetailsRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName, "created_by"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
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
				acctest.GenerateDataSourceFromRepresentationMap("oci_datascience_notebook_sessions", "test_notebook_session", acctest.Optional, acctest.Update, notebookSessionConfigurationDetailsDataSourceRepresentation) +
				compartmentIdVariableStr + NotebookSessionConfigurationDetailsResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_datascience_notebook_session", "test_notebook_session", acctest.Optional, acctest.Update, notebookSessionConfigurationDetailsRepresentation),
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
				resource.TestCheckResourceAttr(datasourceName, "notebook_sessions.0.notebook_session_config_details.0.block_storage_size_in_gbs", "100"),
				resource.TestCheckResourceAttr(datasourceName, "notebook_sessions.0.notebook_session_config_details.0.shape", "VM.Standard2.1"),
				resource.TestCheckResourceAttrSet(datasourceName, "notebook_sessions.0.notebook_session_url"),
				resource.TestCheckResourceAttrSet(datasourceName, "notebook_sessions.0.project_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "notebook_sessions.0.state"),
				resource.TestCheckResourceAttrSet(datasourceName, "notebook_sessions.0.time_created"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_datascience_notebook_session", "test_notebook_session", acctest.Required, acctest.Create, notebookSessionConfigurationDetailsSingularDataSourceRepresentation) +
				compartmentIdVariableStr + NotebookSessionConfigurationDetailsResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "notebook_session_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "created_by"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "notebook_session_url"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
			),
		},
		// verify resource import
		{
			Config:                  config + DatascienceNotebookSessionRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckDatascienceNotebookSessionDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).DataScienceClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_datascience_notebook_session" {
			noResourceFound = false
			request := oci_datascience.GetNotebookSessionRequest{}

			tmp := rs.Primary.ID
			request.NotebookSessionId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "datascience")

			response, err := client.GetNotebookSession(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_datascience.NotebookSessionLifecycleStateDeleted): true,
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
	if !acctest.InSweeperExcludeList("DatascienceNotebookSession") {
		resource.AddTestSweepers("DatascienceNotebookSession", &resource.Sweeper{
			Name:         "DatascienceNotebookSession",
			Dependencies: acctest.DependencyGraph["notebookSession"],
			F:            sweepDatascienceNotebookSessionResource,
		})
	}
}

func sweepDatascienceNotebookSessionResource(compartment string) error {
	dataScienceClient := acctest.GetTestClients(&schema.ResourceData{}).DataScienceClient()
	notebookSessionIds, err := getDatascienceNotebookSessionIds(compartment)
	if err != nil {
		return err
	}
	for _, notebookSessionId := range notebookSessionIds {
		if ok := acctest.SweeperDefaultResourceId[notebookSessionId]; !ok {
			deleteNotebookSessionRequest := oci_datascience.DeleteNotebookSessionRequest{}

			deleteNotebookSessionRequest.NotebookSessionId = &notebookSessionId

			deleteNotebookSessionRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "datascience")
			_, err = dataScienceClient.DeleteNotebookSession(context.Background(), deleteNotebookSessionRequest)
			if err != nil {
				fmt.Printf("Error deleting NotebookSession %s %s, It is possible that the resource is already deleted. Please verify manually \n", notebookSessionId, err)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &notebookSessionId, DatascienceNotebookSessionSweepWaitCondition, time.Duration(3*time.Minute),
				DatascienceNotebookSessionSweepResponseFetchOperation, "datascience", true)
		}
	}
	return nil
}

func getDatascienceNotebookSessionIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "NotebookSessionId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	dataScienceClient := acctest.GetTestClients(&schema.ResourceData{}).DataScienceClient()

	listNotebookSessionsRequest := oci_datascience.ListNotebookSessionsRequest{}
	listNotebookSessionsRequest.CompartmentId = &compartmentId
	listNotebookSessionsRequest.LifecycleState = oci_datascience.ListNotebookSessionsLifecycleStateActive
	listNotebookSessionsResponse, err := dataScienceClient.ListNotebookSessions(context.Background(), listNotebookSessionsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting NotebookSession list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, notebookSession := range listNotebookSessionsResponse.Items {
		id := *notebookSession.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "NotebookSessionId", id)
	}
	return resourceIds, nil
}

func DatascienceNotebookSessionSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if notebookSessionResponse, ok := response.Response.(oci_datascience.GetNotebookSessionResponse); ok {
		return notebookSessionResponse.LifecycleState != oci_datascience.NotebookSessionLifecycleStateDeleted
	}
	return false
}

func DatascienceNotebookSessionSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.DataScienceClient().GetNotebookSession(context.Background(), oci_datascience.GetNotebookSessionRequest{
		NotebookSessionId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
