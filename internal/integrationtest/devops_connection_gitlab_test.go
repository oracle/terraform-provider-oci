// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/oracle/terraform-provider-oci/internal/acctest"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/oracle/terraform-provider-oci/httpreplay"
)

var (
	DevopsGitlabConnectionRequiredOnlyResource = DevopsGitlabConnectionResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_devops_connection", "test_connection", acctest.Required, acctest.Create, devopsGitlabConnectionRepresentation)

	DevopsGitlabConnectionResourceConfig = DevopsGitlabConnectionResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_devops_connection", "test_connection", acctest.Optional, acctest.Update, devopsGitlabConnectionRepresentation)

	devopsGitlabConnectionSingularDataSourceRepresentation = map[string]interface{}{
		"connection_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_devops_connection.test_connection.id}`},
	}

	devopsGitlabConnectionDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Optional, Create: `${var.compartment_id}`},
		"display_name":   acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"id":             acctest.Representation{RepType: acctest.Optional, Create: `${oci_devops_connection.test_connection.id}`},
		"project_id":     acctest.Representation{RepType: acctest.Optional, Create: `${oci_devops_project.test_project.id}`},
		"state":          acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"filter":         acctest.RepresentationGroup{RepType: acctest.Required, Group: devopsGitlabConnectionDataSourceFilterRepresentation}}
	devopsGitlabConnectionDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_devops_connection.test_connection.id}`}},
	}

	devopsGitlabConnectionRepresentation = map[string]interface{}{
		"access_token":    acctest.Representation{RepType: acctest.Required, Create: `${var.gitlab_access_token_vault_id}`},
		"connection_type": acctest.Representation{RepType: acctest.Required, Create: `GITLAB_ACCESS_TOKEN`},
		"project_id":      acctest.Representation{RepType: acctest.Required, Create: `${oci_devops_project.test_project.id}`},
		"defined_tags":    acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "UpdatedValue")}`},
		"description":     acctest.Representation{RepType: acctest.Optional, Create: `description`, Update: `description2`},
		"display_name":    acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"freeform_tags":   acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"bar-key": "value"}, Update: map[string]string{"Department": "Accounting"}},
		"lifecycle":       acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreDefinedTagsDifferencesRepresentation},
	}

	DevopsGitlabConnectionResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_devops_project", "test_project", acctest.Required, acctest.Create, DevopsProjectRepresentation) +
		DefinedTagsDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_ons_notification_topic", "test_notification_topic", acctest.Required, acctest.Create, OnsNotificationTopicRepresentation)
)

// issue-routing-tag: devops/default
func TestDevopsGitlabConnectionResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDevopsGitlabConnectionResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	gitlabAccessTokenVaultId := utils.GetEnvSettingWithBlankDefault("github_access_token_vault_id")
	gitlabAccessTokenVaultIdStr := fmt.Sprintf("variable \"gitlab_access_token_vault_id\" { default = \"%s\" }\n", gitlabAccessTokenVaultId)

	resourceName := "oci_devops_connection.test_connection"
	datasourceName := "data.oci_devops_connections.test_connections"
	singularDatasourceName := "data.oci_devops_connection.test_connection"

	var resId, resId2 string

	// Save TF content to create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+gitlabAccessTokenVaultIdStr+DevopsGitlabConnectionResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_devops_connection", "test_connection", acctest.Optional, acctest.Create, devopsGitlabConnectionRepresentation), "devops", "connection", t)

	acctest.ResourceTest(t, testAccCheckDevopsConnectionDestroy, []resource.TestStep{
		// verify create
		{
			Config: config + compartmentIdVariableStr + gitlabAccessTokenVaultIdStr + DevopsGitlabConnectionResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_devops_connection", "test_connection", acctest.Required, acctest.Create, devopsGitlabConnectionRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "access_token"),
				resource.TestCheckResourceAttr(resourceName, "connection_type", "GITLAB_ACCESS_TOKEN"),
				resource.TestCheckResourceAttrSet(resourceName, "project_id"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next create
		{
			Config: config + compartmentIdVariableStr + gitlabAccessTokenVaultIdStr + DevopsGitlabConnectionResourceDependencies,
		},
		// verify create with optionals
		{
			Config: config + compartmentIdVariableStr + gitlabAccessTokenVaultIdStr + DevopsGitlabConnectionResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_devops_connection", "test_connection", acctest.Optional, acctest.Create, devopsGitlabConnectionRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "access_token"),
				resource.TestCheckResourceAttrSet(resourceName, "compartment_id"),
				resource.TestCheckResourceAttr(resourceName, "connection_type", "GITLAB_ACCESS_TOKEN"),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "project_id"),

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
			Config: config + compartmentIdVariableStr + gitlabAccessTokenVaultIdStr + DevopsGitlabConnectionResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_devops_connection", "test_connection", acctest.Optional, acctest.Update, devopsGitlabConnectionRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "access_token"),
				resource.TestCheckResourceAttrSet(resourceName, "compartment_id"),
				resource.TestCheckResourceAttr(resourceName, "description", "description2"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "project_id"),

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
				acctest.GenerateDataSourceFromRepresentationMap("oci_devops_connections", "test_connections", acctest.Optional, acctest.Update, devopsGitlabConnectionDataSourceRepresentation) +
				compartmentIdVariableStr + gitlabAccessTokenVaultIdStr + DevopsGitlabConnectionResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_devops_connection", "test_connection", acctest.Optional, acctest.Update, devopsGitlabConnectionRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttrSet(datasourceName, "id"),
				resource.TestCheckResourceAttrSet(datasourceName, "project_id"),
				resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),

				resource.TestCheckResourceAttr(datasourceName, "connection_collection.#", "1"),
				//resource.TestCheckResourceAttr(datasourceName, "connection_collection.0.items.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_devops_connection", "test_connection", acctest.Required, acctest.Create, devopsGitlabConnectionSingularDataSourceRepresentation) +
				compartmentIdVariableStr + gitlabAccessTokenVaultIdStr + DevopsGitlabConnectionResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "connection_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "compartment_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "connection_type", "GITLAB_ACCESS_TOKEN"), //https://dyn.slack.com/archives/C021H8MJD7B/p1632873749139600
				resource.TestCheckResourceAttr(singularDatasourceName, "description", "description2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "last_connection_validation_result.#", "0"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
		// verify resource import
		{
			Config:                  config + DevopsConnectionRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}
