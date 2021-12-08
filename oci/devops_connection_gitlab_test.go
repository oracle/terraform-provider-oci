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
	DevopsGitlabConnectionRequiredOnlyResource = DevopsGitlabConnectionResourceDependencies +
		GenerateResourceFromRepresentationMap("oci_devops_connection", "test_connection", Required, Create, devopsGitlabConnectionRepresentation)

	DevopsGitlabConnectionResourceConfig = DevopsGitlabConnectionResourceDependencies +
		GenerateResourceFromRepresentationMap("oci_devops_connection", "test_connection", Optional, Update, devopsGitlabConnectionRepresentation)

	devopsGitlabConnectionSingularDataSourceRepresentation = map[string]interface{}{
		"connection_id": Representation{RepType: Required, Create: `${oci_devops_connection.test_connection.id}`},
	}

	devopsGitlabConnectionDataSourceRepresentation = map[string]interface{}{
		"compartment_id":  Representation{RepType: Optional, Create: `${var.compartment_id}`},
		"connection_type": Representation{RepType: Optional, Create: `GitlabAccessToken`},
		"display_name":    Representation{RepType: Optional, Create: `displayName`, Update: `displayName2`},
		"id":              Representation{RepType: Optional, Create: `${oci_devops_connection.test_connection.id}`},
		"project_id":      Representation{RepType: Optional, Create: `${oci_devops_project.test_project.id}`},
		"state":           Representation{RepType: Optional, Create: `ACTIVE`},
		"filter":          RepresentationGroup{Required, devopsGitlabConnectionDataSourceFilterRepresentation}}
	devopsGitlabConnectionDataSourceFilterRepresentation = map[string]interface{}{
		"name":   Representation{RepType: Required, Create: `id`},
		"values": Representation{RepType: Required, Create: []string{`${oci_devops_connection.test_connection.id}`}},
	}

	devopsGitlabConnectionRepresentation = map[string]interface{}{
		"access_token":    Representation{RepType: Required, Create: `${var.gitlab_access_token_vault_id}`},
		"connection_type": Representation{RepType: Required, Create: `GITLAB_ACCESS_TOKEN`},
		"project_id":      Representation{RepType: Required, Create: `${oci_devops_project.test_project.id}`},
		"defined_tags":    Representation{RepType: Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "UpdatedValue")}`},
		"description":     Representation{RepType: Optional, Create: `description`, Update: `description2`},
		"display_name":    Representation{RepType: Optional, Create: `displayName`, Update: `displayName2`},
		"freeform_tags":   Representation{RepType: Optional, Create: map[string]string{"bar-key": "value"}, Update: map[string]string{"Department": "Accounting"}},
	}

	DevopsGitlabConnectionResourceDependencies = GenerateResourceFromRepresentationMap("oci_devops_project", "test_project", Required, Create, devopsProjectRepresentation) +
		DefinedTagsDependencies +
		GenerateResourceFromRepresentationMap("oci_ons_notification_topic", "test_notification_topic", Required, Create, notificationTopicRepresentation)
)

// issue-routing-tag: devops/default
func TestDevopsGitlabConnectionResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDevopsGitlabConnectionResource_basic")
	defer httpreplay.SaveScenario()

	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	gitlabAccessTokenVaultId := getEnvSettingWithBlankDefault("github_access_token_vault_id")
	gitlabAccessTokenVaultIdStr := fmt.Sprintf("variable \"gitlab_access_token_vault_id\" { default = \"%s\" }\n", gitlabAccessTokenVaultId)

	resourceName := "oci_devops_connection.test_connection"
	datasourceName := "data.oci_devops_connections.test_connections"
	singularDatasourceName := "data.oci_devops_connection.test_connection"

	var resId, resId2 string

	// Save TF content to create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	SaveConfigContent(config+compartmentIdVariableStr+gitlabAccessTokenVaultIdStr+DevopsGitlabConnectionResourceDependencies+
		GenerateResourceFromRepresentationMap("oci_devops_connection", "test_connection", Optional, Create, devopsGitlabConnectionRepresentation), "devops", "connection", t)

	ResourceTest(t, testAccCheckDevopsConnectionDestroy, []resource.TestStep{
		// verify create
		{
			Config: config + compartmentIdVariableStr + gitlabAccessTokenVaultIdStr + DevopsGitlabConnectionResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_devops_connection", "test_connection", Required, Create, devopsGitlabConnectionRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "access_token"),
				resource.TestCheckResourceAttr(resourceName, "connection_type", "GITLAB_ACCESS_TOKEN"),
				resource.TestCheckResourceAttrSet(resourceName, "project_id"),

				func(s *terraform.State) (err error) {
					resId, err = FromInstanceState(s, resourceName, "id")
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
				GenerateResourceFromRepresentationMap("oci_devops_connection", "test_connection", Optional, Create, devopsGitlabConnectionRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "access_token"),
				resource.TestCheckResourceAttrSet(resourceName, "compartment_id"),
				resource.TestCheckResourceAttr(resourceName, "connection_type", "GITLAB_ACCESS_TOKEN"),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "project_id"),

				func(s *terraform.State) (err error) {
					resId, err = FromInstanceState(s, resourceName, "id")
					if isEnableExportCompartment, _ := strconv.ParseBool(getEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
						if errExport := TestExportCompartmentWithResourceName(&resId, &compartmentId, resourceName); errExport != nil {
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
				GenerateResourceFromRepresentationMap("oci_devops_connection", "test_connection", Optional, Update, devopsGitlabConnectionRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "access_token"),
				resource.TestCheckResourceAttrSet(resourceName, "compartment_id"),
				resource.TestCheckResourceAttr(resourceName, "description", "description2"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "project_id"),

				func(s *terraform.State) (err error) {
					resId2, err = FromInstanceState(s, resourceName, "id")
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
				GenerateDataSourceFromRepresentationMap("oci_devops_connections", "test_connections", Optional, Update, devopsGitlabConnectionDataSourceRepresentation) +
				compartmentIdVariableStr + gitlabAccessTokenVaultIdStr + DevopsGitlabConnectionResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_devops_connection", "test_connection", Optional, Update, devopsGitlabConnectionRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "connection_type", "GitlabAccessToken"), //TODO: Needs to accept `GITLAB_ACCESS_TOKEN`
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
				GenerateDataSourceFromRepresentationMap("oci_devops_connection", "test_connection", Required, Create, devopsGitlabConnectionSingularDataSourceRepresentation) +
				compartmentIdVariableStr + gitlabAccessTokenVaultIdStr + DevopsGitlabConnectionResourceConfig,
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "connection_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "compartment_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "connection_type", "GITLAB_ACCESS_TOKEN"), //https://dyn.slack.com/archives/C021H8MJD7B/p1632873749139600
				resource.TestCheckResourceAttr(singularDatasourceName, "description", "description2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
		// remove singular datasource from previous step so that it doesn't conflict with import tests
		{
			Config: config + compartmentIdVariableStr + gitlabAccessTokenVaultIdStr + DevopsGitlabConnectionResourceConfig,
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
