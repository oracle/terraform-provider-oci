// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	// before running tests, ensure to set up environment variables used below.
	JmsPluginFleetId       = utils.GetEnvSettingWithBlankDefault("fleet_advanced_feature_ocid")
	JmsPluginCompartmentId = utils.GetEnvSettingWithBlankDefault("compartment_ocid")

	JmsPluginDataSourceRepresentation = map[string]interface{}{
		"fleet_id":                  acctest.Representation{RepType: acctest.Optional, Create: JmsPluginFleetId},
		"compartment_id":            acctest.Representation{RepType: acctest.Optional, Create: JmsPluginCompartmentId},
		"compartment_id_in_subtree": acctest.Representation{RepType: acctest.Optional, Create: `false`},
	}
)

// issue-routing-tag: jms/default
func TestJmsPluginResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestJmsPluginResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	datasourceName := "data.oci_jms_jms_plugins.test_jms_plugins"

	acctest.ResourceTest(t, nil, []resource.TestStep{

		// verify Create
		// note: we cannot write test for this case because
		// it requires setup of fleet -> compute instance -> management agent -> jms plugin.

		// verify update
		// note: we cannot write test for this case because
		// it requires setup of fleet -> compute instance -> management agent -> jms plugin.

		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap(
					"oci_jms_jms_plugins",
					"test_jms_plugins",
					acctest.Optional,
					acctest.Create,
					JmsPluginDataSourceRepresentation,
				),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "fleet_id"),
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", JmsPluginCompartmentId),
				resource.TestCheckResourceAttr(datasourceName, "compartment_id_in_subtree", "false"),

				// we can only verify that response contain zero items because
				// it requires setup of fleet -> compute instance -> management agent -> jms plugin.
				resource.TestCheckResourceAttrSet(datasourceName, "jms_plugin_collection.#"),
				resource.TestCheckResourceAttr(datasourceName, "jms_plugin_collection.0.items.#", "0"),
			),
		},

		// verify singular datasource
		// note: we cannot write test for this case because
		// it requires setup of fleet -> compute instance -> management agent -> jms plugin.

		// verify resource import
		// note: we cannot write test for this case because
		// it requires setup of fleet -> compute instance -> management agent -> jms plugin.
	})
}
