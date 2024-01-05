// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"

	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	MediaServicesMediaServicesMediaWorkflowTaskDeclarationSingularDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"is_current":     acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"name":           acctest.Representation{RepType: acctest.Optional, Create: `name`},
		"version":        acctest.Representation{RepType: acctest.Optional, Create: `10`},
	}

	MediaServicesMediaWorkflowTaskDeclarationResourceConfig = ""
)

// issue-routing-tag: media_services/default
func TestMediaServicesMediaWorkflowTaskDeclarationResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestMediaServicesMediaWorkflowTaskDeclarationResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	singularDatasourceName := "data.oci_media_services_media_workflow_task_declaration.test_media_workflow_task_declaration"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_media_services_media_workflow_task_declaration", "test_media_workflow_task_declaration", acctest.Required, acctest.Create, MediaServicesMediaServicesMediaWorkflowTaskDeclarationSingularDataSourceRepresentation) +
				compartmentIdVariableStr + MediaServicesMediaWorkflowTaskDeclarationResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "items.#", "7"),
				resource.TestCheckResourceAttr(singularDatasourceName, "items.0.name", "getFiles"),
				resource.TestCheckResourceAttr(singularDatasourceName, "items.0.version", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "items.1.name", "ingest"),
			),
		},
	})
}
