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
	managementAgentImageDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"install_type":   acctest.Representation{RepType: acctest.Optional, Create: `AGENT`},
		"state":          acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
	}

	ManagementAgentImageResourceConfig = ""
)

// issue-routing-tag: management_agent/default
func TestManagementAgentManagementAgentImageResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestManagementAgentManagementAgentImageResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_management_agent_management_agent_images.test_management_agent_images"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_management_agent_management_agent_images", "test_management_agent_images", acctest.Required, acctest.Create, managementAgentImageDataSourceRepresentation) +
				compartmentIdVariableStr + ManagementAgentImageResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(datasourceName, "management_agent_images.#"),
				resource.TestCheckResourceAttrSet(datasourceName, "management_agent_images.0.checksum"),
				resource.TestCheckResourceAttrSet(datasourceName, "management_agent_images.0.id"),
				resource.TestCheckResourceAttr(datasourceName, "management_agent_images.0.image_object_storage_details.#", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "management_agent_images.0.object_url"),
				resource.TestCheckResourceAttrSet(datasourceName, "management_agent_images.0.package_architecture_type"),
				resource.TestCheckResourceAttrSet(datasourceName, "management_agent_images.0.package_type"),
				resource.TestCheckResourceAttrSet(datasourceName, "management_agent_images.0.platform_name"),
				resource.TestCheckResourceAttrSet(datasourceName, "management_agent_images.0.platform_type"),
				resource.TestCheckResourceAttrSet(datasourceName, "management_agent_images.0.size"),
				resource.TestCheckResourceAttrSet(datasourceName, "management_agent_images.0.state"),
				resource.TestCheckResourceAttrSet(datasourceName, "management_agent_images.0.version"),
				resource.TestCheckResourceAttr(datasourceName, "management_agent_images.0.state", "ACTIVE"),
			),
		},
	})
}
