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
	JmsPluginErrorCompartmentId     = utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	JmsPluginErrorManagedInstanceId = utils.GetEnvSettingWithBlankDefault("managed_instance_ocid")

	JmsPluginErrorDataSourceRepresentation = map[string]interface{}{
		"compartment_id":            acctest.Representation{RepType: acctest.Optional, Create: JmsPluginErrorCompartmentId},
		"compartment_id_in_subtree": acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"managed_instance_id":       acctest.Representation{RepType: acctest.Optional, Create: JmsPluginErrorManagedInstanceId},
	}

	JmsPluginErrorResourceConfig = acctest.GenerateResourceFromRepresentationMap("oci_os_management_hub_managed_instance", "test_managed_instance", acctest.Required, acctest.Create, OsManagementHubManagedInstanceRepresentation)
)

// issue-routing-tag: jms/default
func TestJmsPluginErrorResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestJmsPluginErrorResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	datasourceName := "data.oci_jms_plugin_errors.test_plugin_errors"

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap(
					"oci_jms_plugin_errors",
					"test_plugin_errors",
					acctest.Optional,
					acctest.Create,
					JmsPluginErrorDataSourceRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", JmsPluginErrorCompartmentId),
				resource.TestCheckResourceAttr(datasourceName, "managed_instance_id", JmsPluginErrorManagedInstanceId),
				resource.TestCheckResourceAttr(datasourceName, "compartment_id_in_subtree", `false`),

				resource.TestCheckResourceAttrSet(datasourceName, "plugin_error_collection.#"),
				// we can only verify that response contain zero items because we are using dummy test data values
				// we cannot use actual values because it requires create API.
				resource.TestCheckResourceAttr(datasourceName, "plugin_error_collection.0.items.#", "0"),
			),
		},
		// verify singular datasource
		// note: we cannot write test to verify singular data source because
		// crypto analysis processing requires create API.
	})
}
