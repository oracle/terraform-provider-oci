// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"

	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	// before running tests, ensure to set up environment variables used below
	JmsAgentInstallerFleetId       = utils.GetEnvSettingWithBlankDefault("fleet_ocid")
	JmsAgentInstallerCompartmentId = utils.GetEnvSettingWithBlankDefault("compartment_ocid")

	JmsAgentInstallerDataSourceRepresentation = map[string]interface{}{
		"compartment_id":        acctest.Representation{RepType: acctest.Optional, Create: JmsAgentInstallerCompartmentId},
		"fleet_id":              acctest.Representation{RepType: acctest.Optional, Create: JmsAgentInstallerFleetId},
		"os_family":             acctest.Representation{RepType: acctest.Optional, Create: `LINUX`},
		"platform_architecture": acctest.Representation{RepType: acctest.Optional, Create: `X86_64`},
	}
)

// issue-routing-tag: jms/default
func TestJmsAgentInstallerResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestJmsAgentInstallerResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	datasourceName := "data.oci_jms_agent_installers.test_agent_installers"

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap(
					"oci_jms_agent_installers",
					"test_agent_installers",
					acctest.Optional,
					acctest.Create,
					JmsAgentInstallerDataSourceRepresentation,
				),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", JmsAgentInstallerCompartmentId),
				resource.TestCheckResourceAttrSet(datasourceName, "fleet_id"),
				resource.TestCheckResourceAttr(datasourceName, "os_family", "LINUX"),
				resource.TestCheckResourceAttr(datasourceName, "platform_architecture", "X86_64"),

				resource.TestCheckResourceAttrSet(datasourceName, "agent_installer_collection.#"),
				resource.TestCheckResourceAttr(datasourceName, "agent_installer_collection.0.items.#", "2"), // 2 - one for *.rpm, one for *.zip

				// check first item in the list for expected fields to be set.
				// we do not have to verify its actual value as they may change from time to time
				resource.TestCheckResourceAttrSet(datasourceName, "agent_installer_collection.0.items.0.agent_installer_description"),
				resource.TestCheckResourceAttrSet(datasourceName, "agent_installer_collection.0.items.0.agent_installer_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "agent_installer_collection.0.items.0.agent_installer_version"),
				resource.TestCheckResourceAttrSet(datasourceName, "agent_installer_collection.0.items.0.agent_version"),
				resource.TestCheckResourceAttrSet(datasourceName, "agent_installer_collection.0.items.0.approximate_file_size_in_bytes"),
				resource.TestCheckResourceAttrSet(datasourceName, "agent_installer_collection.0.items.0.java_version"),
				resource.TestCheckResourceAttrSet(datasourceName, "agent_installer_collection.0.items.0.os_family"),
				resource.TestCheckResourceAttrSet(datasourceName, "agent_installer_collection.0.items.0.package_type"),
				resource.TestCheckResourceAttrSet(datasourceName, "agent_installer_collection.0.items.0.platform_architecture"),
				resource.TestCheckResourceAttrSet(datasourceName, "agent_installer_collection.0.items.0.sha256"),
			),
		},
	})
}
