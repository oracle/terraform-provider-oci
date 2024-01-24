// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"testing"

	"github.com/oracle/terraform-provider-oci/internal/acctest"
	"github.com/oracle/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"github.com/oracle/terraform-provider-oci/httpreplay"
)

var (
	DevopsDevopsRepositoryArchiveContentSingularDataSourceRepresentation = map[string]interface{}{
		"repository_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_devops_repository.test_repository.id}`},
		"format":        acctest.Representation{RepType: acctest.Optional, Create: `zip`},
		"ref_name":      acctest.Representation{RepType: acctest.Optional, Create: `main`},
	}

	DevopsRepositoryArchiveContentResourceConfig = acctest.GenerateResourceFromRepresentationMap("oci_devops_project", "test_project", acctest.Required, acctest.Create, DevopsProjectRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_devops_repository", "test_repository", acctest.Required, acctest.Create, DevopsRepositoryRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_ons_notification_topic", "test_notification_topic", acctest.Required, acctest.Create, OnsNotificationTopicRepresentation)
)

// issue-routing-tag: devops/default
func TestDevopsRepositoryArchiveContentResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDevopsRepositoryArchiveContentResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	singularDatasourceName := "data.oci_devops_repository_archive_content.test_repository_archive_content"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_devops_repository_archive_content", "test_repository_archive_content", acctest.Required, acctest.Create, DevopsDevopsRepositoryArchiveContentSingularDataSourceRepresentation) +
				compartmentIdVariableStr + DevopsRepositoryArchiveContentResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(singularDatasourceName, "format", "zip"),
				resource.TestCheckResourceAttr(singularDatasourceName, "ref_name", "main"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "repository_id"),
			),
		},
	})
}
