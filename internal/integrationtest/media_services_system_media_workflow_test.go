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
	MediaServicesMediaServicesSystemMediaWorkflowSingularDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"name":           acctest.Representation{RepType: acctest.Optional, Create: `name`},
	}

	MediaServicesSystemMediaWorkflowResourceConfig = ""
)

// issue-routing-tag: media_services/default
func TestMediaServicesSystemMediaWorkflowResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestMediaServicesSystemMediaWorkflowResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	singularDatasourceName := "data.oci_media_services_system_media_workflow.test_system_media_workflow"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_media_services_system_media_workflow", "test_system_media_workflow", acctest.Required, acctest.Create, MediaServicesMediaServicesSystemMediaWorkflowSingularDataSourceRepresentation) +
				compartmentIdVariableStr + MediaServicesSystemMediaWorkflowResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "items.#", "2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "items.0.name", "ingestWorkflow"),
				resource.TestCheckResourceAttr(singularDatasourceName, "items.0.tasks.0.type", "ingest"),
				resource.TestCheckResourceAttr(singularDatasourceName, "items.1.name", "transcodeForStreaming"),
			),
		},
	})
}
