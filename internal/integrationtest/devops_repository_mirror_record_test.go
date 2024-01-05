// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"strings"
	"testing"

	"github.com/oracle/terraform-provider-oci/internal/acctest"
	"github.com/oracle/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"github.com/oracle/terraform-provider-oci/httpreplay"
)

var (
	DevopsDevopsRepositoryMirrorRecordSingularDataSourceRepresentation = map[string]interface{}{
		"mirror_record_type": acctest.Representation{RepType: acctest.Required, Create: `lastSuccessful`},
		"repository_id":      acctest.Representation{RepType: acctest.Required, Create: `${oci_devops_repository.test_repository.id}`},
	}

	DevopsDevopsRepositoryMirrorRecordDataSourceRepresentation = map[string]interface{}{
		"repository_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_devops_repository.test_repository.id}`},
	}

	DevopsMirrorRecordRepositoryConfigTriggerScheduleRepresentation = map[string]interface{}{
		"schedule_type": acctest.Representation{RepType: acctest.Required, Create: `NONE`},
	}

	DevopsMirrorRecordRepositoryRepresentation = map[string]interface{}{
		"name":                     acctest.Representation{RepType: acctest.Required, Create: `name`},
		"project_id":               acctest.Representation{RepType: acctest.Required, Create: `${oci_devops_project.test_project.id}`},
		"repository_type":          acctest.Representation{RepType: acctest.Required, Create: `MIRRORED`},
		"mirror_repository_config": acctest.RepresentationGroup{RepType: acctest.Required, Group: DevopsRepositoryMirrorRecordRepositoryConfigRepresentation},
	}

	DevopsRepositoryMirrorRecordRepositoryConfigRepresentation = map[string]interface{}{
		"connector_id":     acctest.Representation{RepType: acctest.Required, Create: `${oci_devops_connection.test_connection.id}`},
		"repository_url":   acctest.Representation{RepType: acctest.Required, Create: `${var.mirror_repository_url}`},
		"trigger_schedule": acctest.RepresentationGroup{RepType: acctest.Required, Group: DevopsMirrorRecordRepositoryConfigTriggerScheduleRepresentation},
	}

	DevopsConnectionRepresentationMirrorRecord = map[string]interface{}{
		"access_token":    acctest.Representation{RepType: acctest.Required, Create: `${var.github_access_token_vault_id}`},
		"connection_type": acctest.Representation{RepType: acctest.Required, Create: `GITHUB_ACCESS_TOKEN`},
		"project_id":      acctest.Representation{RepType: acctest.Required, Create: `${oci_devops_project.test_project.id}`},
		"defined_tags":    acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"description":     acctest.Representation{RepType: acctest.Optional, Create: `description`, Update: `description2`},
		"display_name":    acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"freeform_tags":   acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"bar-key": "value"}, Update: map[string]string{"Department": "Accounting"}},
	}

	DevopsRepositoryMirrorRecordResourceConfig = acctest.GenerateResourceFromRepresentationMap("oci_devops_project", "test_project", acctest.Required, acctest.Create, DevopsProjectRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_devops_repository", "test_repository", acctest.Required, acctest.Create, DevopsMirrorRecordRepositoryRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_devops_connection", "test_connection", acctest.Required, acctest.Create, DevopsConnectionRepresentationMirrorRecord) +
		acctest.GenerateResourceFromRepresentationMap("oci_ons_notification_topic", "test_notification_topic", acctest.Required, acctest.Create, OnsNotificationTopicRepresentation)
)

// issue-routing-tag: devops/default
func TestDevopsRepositoryMirrorRecordResource_basic(t *testing.T) {
	if !strings.Contains(utils.GetEnvSettingWithBlankDefault("enabled_tests"), "RepositoryMirrorRecord") {
		t.Skip("TestDevopsRepositoryMirrorRecordResource_basic test needs a Repository resource with existing commits to test")
	}
	httpreplay.SetScenario("TestDevopsRepositoryMirrorRecordResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	githubAccessTokenVaultId := utils.GetEnvSettingWithBlankDefault("github_access_token_vault_id")
	githubAccessTokenVaultIdStr := fmt.Sprintf("variable \"github_access_token_vault_id\" { default = \"%s\" }\n", githubAccessTokenVaultId)

	mirrorRepositoryUrl := utils.GetEnvSettingWithBlankDefault("mirror_repository_url")
	mirrorRepositoryUrlStr := fmt.Sprintf("variable \"mirror_repository_url\" { default = \"%s\" }\n", mirrorRepositoryUrl)

	datasourceName := "data.oci_devops_repository_mirror_records.test_repository_mirror_records"
	singularDatasourceName := "data.oci_devops_repository_mirror_record.test_repository_mirror_record"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config + compartmentIdVariableStr + githubAccessTokenVaultIdStr + mirrorRepositoryUrlStr + DevopsRepositoryMirrorRecordResourceConfig +
				acctest.GenerateDataSourceFromRepresentationMap("oci_devops_repository_mirror_records", "test_repository_mirror_records", acctest.Required, acctest.Create, DevopsDevopsRepositoryMirrorRecordDataSourceRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "repository_id"),

				resource.TestCheckResourceAttrSet(datasourceName, "repository_mirror_record_collection.#"),
			),
		},
		// verify singular datasource
		{
			Config: config + compartmentIdVariableStr + githubAccessTokenVaultIdStr + mirrorRepositoryUrlStr + DevopsRepositoryMirrorRecordResourceConfig +
				acctest.GenerateDataSourceFromRepresentationMap("oci_devops_repository_mirror_record", "test_repository_mirror_record", acctest.Required, acctest.Create, DevopsDevopsRepositoryMirrorRecordSingularDataSourceRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(singularDatasourceName, "mirror_record_type", "current"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "repository_id"),

				resource.TestCheckResourceAttrSet(singularDatasourceName, "mirror_status"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_ended"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_enqueued"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_started"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "work_request_id"),
			),
		},
	})
}
