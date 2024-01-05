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
	DevopsDevopsRepositoryObjectSingularDataSourceRepresentation = map[string]interface{}{
		"repository_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_devops_repository.test_repository.id}`},
	}

	DevopsRepositoryObjectResourceConfig = acctest.GenerateResourceFromRepresentationMap("oci_devops_project", "test_project", acctest.Required, acctest.Create, DevopsProjectRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_devops_repository", "test_repository", acctest.Required, acctest.Create, DevopsRepositoryRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_ons_notification_topic", "test_notification_topic", acctest.Required, acctest.Create, OnsNotificationTopicRepresentation)
)

// issue-routing-tag: devops/default
func TestDevopsRepositoryObjectResource_basic(t *testing.T) {
	if !strings.Contains(utils.GetEnvSettingWithBlankDefault("enabled_tests"), "RepositoryObject") {
		t.Skip("TestDevopsRepositoryObjectResource_basic test needs a Repository resource with existing commits to test")
	}
	httpreplay.SetScenario("TestDevopsRepositoryObjectResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	singularDatasourceName := "data.oci_devops_repository_object.test_repository_object"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_devops_repository_object", "test_repository_object", acctest.Required, acctest.Create, DevopsDevopsRepositoryObjectSingularDataSourceRepresentation) +
				compartmentIdVariableStr + DevopsRepositoryObjectResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(singularDatasourceName, "file_path", "filePath"),
				resource.TestCheckResourceAttr(singularDatasourceName, "ref_name", "refName"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "repository_id"),

				resource.TestCheckResourceAttrSet(singularDatasourceName, "is_binary"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "sha"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "size_in_bytes"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "type"),
			),
		},
	})
}
