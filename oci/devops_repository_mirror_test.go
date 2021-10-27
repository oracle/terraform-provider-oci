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
	repositoryMirrorRepresentation = map[string]interface{}{
		"repository_id": Representation{RepType: Required, Create: `${oci_devops_repository.test_repository.id}`},
	}

	repositoryMirrorOnlyRepositoryConfigTriggerScheduleRepresentation = map[string]interface{}{
		"schedule_type": Representation{RepType: Required, Create: `NONE`},
	}

	repositoryMirrorOnlyRepositoryConfigRepresentation = map[string]interface{}{
		"connector_id":     Representation{RepType: Required, Create: `${oci_devops_connection.test_connection.id}`},
		"repository_url":   Representation{RepType: Required, Create: `${var.mirror_repository_url}`},
		"trigger_schedule": RepresentationGroup{Required, repositoryMirrorOnlyRepositoryConfigTriggerScheduleRepresentation},
	}

	devopsMirrorRepositoryRepresentation = map[string]interface{}{
		"name":                     Representation{RepType: Required, Create: `mirror-name`},
		"project_id":               Representation{RepType: Required, Create: `${oci_devops_project.test_project.id}`},
		"repository_type":          Representation{RepType: Required, Create: `MIRRORED`},
		"mirror_repository_config": RepresentationGroup{Required, repositoryMirrorOnlyRepositoryConfigRepresentation},
	}

	RepositoryMirrorResourceDependencies = GenerateResourceFromRepresentationMap("oci_devops_project", "test_project", Required, Create, devopsProjectRepresentation) +
		GenerateResourceFromRepresentationMap("oci_devops_connection", "test_connection", Required, Create, devopsConnectionRepresentation) +
		GenerateResourceFromRepresentationMap("oci_devops_repository", "test_repository", Required, Update, devopsMirrorRepositoryRepresentation) +
		GenerateResourceFromRepresentationMap("oci_ons_notification_topic", "test_notification_topic", Required, Create, notificationTopicRepresentation)
)

// issue-routing-tag: devops/default
func TestDevopsRepositoryMirrorResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDevopsRepositoryMirrorResource_basic")
	defer httpreplay.SaveScenario()

	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	githubAccessTokenVaultId := getEnvSettingWithBlankDefault("github_access_token_vault_id")
	githubAccessTokenVaultIdStr := fmt.Sprintf("variable \"github_access_token_vault_id\" { default = \"%s\" }\n", githubAccessTokenVaultId)

	mirrorRepositoryUrl := getEnvSettingWithBlankDefault("mirror_repository_url")
	mirrorRepositoryUrlStr := fmt.Sprintf("variable \"mirror_repository_url\" { default = \"%s\" }\n", mirrorRepositoryUrl)

	resourceName := "oci_devops_repository_mirror.test_repository_mirror"

	var resId string
	// Save TF content to create resource with only required properties. This has to be exactly the same as the config part in the create step in the test.
	SaveConfigContent(config+compartmentIdVariableStr+RepositoryMirrorResourceDependencies+
		GenerateResourceFromRepresentationMap("oci_devops_repository_mirror", "test_repository_mirror", Required, Create, repositoryMirrorRepresentation), "devops", "repositoryMirror", t)

	ResourceTest(t, nil, []resource.TestStep{
		// verify create
		{
			Config: config + compartmentIdVariableStr + githubAccessTokenVaultIdStr + mirrorRepositoryUrlStr + RepositoryMirrorResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_devops_repository_mirror", "test_repository_mirror", Required, Create, repositoryMirrorRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "repository_id"),

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
	})
}
