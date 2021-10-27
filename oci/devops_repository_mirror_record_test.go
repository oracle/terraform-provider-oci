// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	repositoryMirrorRecordSingularDataSourceRepresentation = map[string]interface{}{
		"mirror_record_type": Representation{RepType: Required, Create: `lastSuccessful`},
		"repository_id":      Representation{RepType: Required, Create: `${oci_devops_repository.test_repository.id}`},
	}

	repositoryMirrorRecordDataSourceRepresentation = map[string]interface{}{
		"repository_id": Representation{RepType: Required, Create: `${oci_devops_repository.test_repository.id}`},
	}

	MirrorRecordRepositoryConfigTriggerScheduleRepresentation = map[string]interface{}{
		"schedule_type": Representation{RepType: Required, Create: `NONE`},
	}

	devopsMirrorRecordRepositoryRepresentation = map[string]interface{}{
		"name":                     Representation{RepType: Required, Create: `name`},
		"project_id":               Representation{RepType: Required, Create: `${oci_devops_project.test_project.id}`},
		"repository_type":          Representation{RepType: Required, Create: `MIRRORED`},
		"mirror_repository_config": RepresentationGroup{Required, repositoryMirrorRecordRepositoryConfigRepresentation},
	}

	repositoryMirrorRecordRepositoryConfigRepresentation = map[string]interface{}{
		"connector_id":     Representation{RepType: Required, Create: `${oci_devops_connection.test_connection.id}`},
		"repository_url":   Representation{RepType: Required, Create: `${var.mirror_repository_url}`},
		"trigger_schedule": RepresentationGroup{Required, MirrorRecordRepositoryConfigTriggerScheduleRepresentation},
	}

	devopsConnectionRepresentationMirrorRecord = map[string]interface{}{
		"access_token":    Representation{RepType: Required, Create: `${var.github_access_token_vault_id}`},
		"connection_type": Representation{RepType: Required, Create: `GITHUB_ACCESS_TOKEN`},
		"project_id":      Representation{RepType: Required, Create: `${oci_devops_project.test_project.id}`},
		"defined_tags":    Representation{RepType: Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"description":     Representation{RepType: Optional, Create: `description`, Update: `description2`},
		"display_name":    Representation{RepType: Optional, Create: `displayName`, Update: `displayName2`},
		"freeform_tags":   Representation{RepType: Optional, Create: map[string]string{"bar-key": "value"}, Update: map[string]string{"Department": "Accounting"}},
	}

	RepositoryMirrorRecordResourceConfig = GenerateResourceFromRepresentationMap("oci_devops_project", "test_project", Required, Create, devopsProjectRepresentation) +
		GenerateResourceFromRepresentationMap("oci_devops_repository", "test_repository", Required, Create, devopsMirrorRecordRepositoryRepresentation) +
		GenerateResourceFromRepresentationMap("oci_devops_connection", "test_connection", Required, Create, devopsConnectionRepresentationMirrorRecord) +
		GenerateResourceFromRepresentationMap("oci_ons_notification_topic", "test_notification_topic", Required, Create, notificationTopicRepresentation)
)

// issue-routing-tag: devops/default
func TestDevopsRepositoryMirrorRecordResource_basic(t *testing.T) {
	if !strings.Contains(getEnvSettingWithBlankDefault("enabled_tests"), "RepositoryMirrorRecord") {
		t.Skip("TestDevopsRepositoryMirrorRecordResource_basic test needs a Repository resource with existing commits to test")
	}
	httpreplay.SetScenario("TestDevopsRepositoryMirrorRecordResource_basic")
	defer httpreplay.SaveScenario()

	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	githubAccessTokenVaultId := getEnvSettingWithBlankDefault("github_access_token_vault_id")
	githubAccessTokenVaultIdStr := fmt.Sprintf("variable \"github_access_token_vault_id\" { default = \"%s\" }\n", githubAccessTokenVaultId)

	mirrorRepositoryUrl := getEnvSettingWithBlankDefault("mirror_repository_url")
	mirrorRepositoryUrlStr := fmt.Sprintf("variable \"mirror_repository_url\" { default = \"%s\" }\n", mirrorRepositoryUrl)

	datasourceName := "data.oci_devops_repository_mirror_records.test_repository_mirror_records"
	singularDatasourceName := "data.oci_devops_repository_mirror_record.test_repository_mirror_record"

	SaveConfigContent("", "", "", t)

	ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config + compartmentIdVariableStr + githubAccessTokenVaultIdStr + mirrorRepositoryUrlStr + RepositoryMirrorRecordResourceConfig +
				GenerateDataSourceFromRepresentationMap("oci_devops_repository_mirror_records", "test_repository_mirror_records", Required, Create, repositoryMirrorRecordDataSourceRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "repository_id"),

				resource.TestCheckResourceAttrSet(datasourceName, "repository_mirror_record_collection.#"),
			),
		},
		// verify singular datasource
		{
			Config: config + compartmentIdVariableStr + githubAccessTokenVaultIdStr + mirrorRepositoryUrlStr + RepositoryMirrorRecordResourceConfig +
				GenerateDataSourceFromRepresentationMap("oci_devops_repository_mirror_record", "test_repository_mirror_record", Required, Create, repositoryMirrorRecordSingularDataSourceRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
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
