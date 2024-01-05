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
	DevopsRepositoryMirrorRepresentation = map[string]interface{}{
		"repository_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_devops_repository.test_repository.id}`},
	}

	DevopsRepositoryMirrorOnlyRepositoryConfigTriggerScheduleRepresentation = map[string]interface{}{
		"schedule_type": acctest.Representation{RepType: acctest.Required, Create: `NONE`},
	}

	DevopsRepositoryMirrorOnlyRepositoryConfigRepresentation = map[string]interface{}{
		"connector_id":     acctest.Representation{RepType: acctest.Required, Create: `${oci_devops_connection.test_connection.id}`},
		"repository_url":   acctest.Representation{RepType: acctest.Required, Create: `${var.mirror_repository_url}`},
		"trigger_schedule": acctest.RepresentationGroup{RepType: acctest.Required, Group: DevopsRepositoryMirrorOnlyRepositoryConfigTriggerScheduleRepresentation},
	}

	DevopsMirrorRepositoryRepresentation = map[string]interface{}{
		"name":                     acctest.Representation{RepType: acctest.Required, Create: `mirror-name`},
		"project_id":               acctest.Representation{RepType: acctest.Required, Create: `${oci_devops_project.test_project.id}`},
		"repository_type":          acctest.Representation{RepType: acctest.Required, Create: `MIRRORED`},
		"mirror_repository_config": acctest.RepresentationGroup{RepType: acctest.Required, Group: DevopsRepositoryMirrorOnlyRepositoryConfigRepresentation},
	}

	DevopsRepositoryMirrorResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_devops_project", "test_project", acctest.Required, acctest.Create, DevopsProjectRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_devops_connection", "test_connection", acctest.Required, acctest.Create, DevopsConnectionRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_devops_repository", "test_repository", acctest.Required, acctest.Update, DevopsMirrorRepositoryRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_ons_notification_topic", "test_notification_topic", acctest.Required, acctest.Create, OnsNotificationTopicRepresentation)
)

// issue-routing-tag: devops/default
func TestDevopsRepositoryMirrorResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDevopsRepositoryMirrorResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	githubAccessTokenVaultId := utils.GetEnvSettingWithBlankDefault("github_access_token_vault_id")
	githubAccessTokenVaultIdStr := fmt.Sprintf("variable \"github_access_token_vault_id\" { default = \"%s\" }\n", githubAccessTokenVaultId)

	mirrorRepositoryUrl := utils.GetEnvSettingWithBlankDefault("mirror_repository_url")
	mirrorRepositoryUrlStr := fmt.Sprintf("variable \"mirror_repository_url\" { default = \"%s\" }\n", mirrorRepositoryUrl)

	resourceName := "oci_devops_repository_mirror.test_repository_mirror"

	var resId string
	// Save TF content to create resource with only required properties. This has to be exactly the same as the config part in the create step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+DevopsRepositoryMirrorResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_devops_repository_mirror", "test_repository_mirror", acctest.Required, acctest.Create, DevopsRepositoryMirrorRepresentation), "devops", "repositoryMirror", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify create
		{
			Config: config + compartmentIdVariableStr + githubAccessTokenVaultIdStr + mirrorRepositoryUrlStr + DevopsRepositoryMirrorResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_devops_repository_mirror", "test_repository_mirror", acctest.Required, acctest.Create, DevopsRepositoryMirrorRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "repository_id"),

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
	})
}
