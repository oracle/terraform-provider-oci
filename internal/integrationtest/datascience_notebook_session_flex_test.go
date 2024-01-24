// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"testing"

	"github.com/oracle/terraform-provider-oci/internal/acctest"
	"github.com/oracle/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/oracle/terraform-provider-oci/httpreplay"
)

var (
	DatascienceNotebookSessionFlexRequiredOnlyResource = acctest.GenerateResourceFromRepresentationMap("oci_datascience_notebook_session", "test_notebook_session", acctest.Required, acctest.Create, notebookSessionFlexRepresentation)

	NotebookSessionResourceFlexConfig = DatascienceNotebookSessionResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_datascience_notebook_session", "test_notebook_session", acctest.Optional, acctest.Update, notebookSessionFlexRepresentation)

	notebookSessionFlexRepresentation = map[string]interface{}{
		"compartment_id":                  acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"project_id":                      acctest.Representation{RepType: acctest.Required, Create: `${oci_datascience_project.test_project.id}`},
		"defined_tags":                    acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"display_name":                    acctest.Representation{RepType: acctest.Optional, Create: `displayName`},
		"freeform_tags":                   acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"notebook_session_config_details": acctest.RepresentationGroup{RepType: acctest.Required, Group: notebookSessionFlexConfigDetailsRepresentation},
		"lifecycle":                       acctest.RepresentationGroup{RepType: acctest.Optional, Group: definedTagsIgnoreRepresentation},
	}

	notebookSessionFlexConfigDetailsRepresentation = map[string]interface{}{
		"shape":                     acctest.Representation{RepType: acctest.Required, Create: `VM.Standard2.1`},
		"block_storage_size_in_gbs": acctest.Representation{RepType: acctest.Optional, Create: `50`},
	}

	notebookSessionShapeConfigDetailsRepresentation = map[string]interface{}{
		"memory_in_gbs": acctest.Representation{RepType: acctest.Optional, Create: `16.0`},
		"ocpus":         acctest.Representation{RepType: acctest.Optional, Create: `1.0`},
	}
)

// issue-routing-tag: datascience/default
func TestDatascienceNotebookSessionResource_flex(t *testing.T) {
	httpreplay.SetScenario("TestDatascienceNotebookSessionResource_flex")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_datascience_notebook_session.test_notebook_session"
	singularDatasourceName := "data.oci_datascience_notebook_session.test_notebook_session"

	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+DatascienceNotebookSessionResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_datascience_notebook_session", "test_notebook_session", acctest.Optional, acctest.Create, notebookSessionFlexRepresentation), "datascience", "notebookSession", t)

	acctest.ResourceTest(t, testAccCheckDatascienceNotebookSessionDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + DatascienceNotebookSessionResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_datascience_notebook_session", "test_notebook_session", acctest.Required, acctest.Create, notebookSessionFlexRepresentation),
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
			Config: config + compartmentIdVariableStr + DatascienceNotebookSessionResourceDependencies,
		},

		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_datascience_notebook_session", "test_notebook_session", acctest.Required, acctest.Create, notebookSessionConfigDetailsSingularDataSourceRepresentation) +
				compartmentIdVariableStr + NotebookSessionResourceFlexConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "notebook_session_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "created_by"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "notebook_session_url"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
			),
		},
		// verify resource import
		{
			Config:                  config + DatascienceNotebookSessionFlexRequiredOnlyResource,
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
