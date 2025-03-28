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
	JmsFleetErrorCompartmentId            = utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	JmsFleetErrorFleetId                  = utils.GetEnvSettingWithBlankDefault("fleet_advanced_feature_ocid")
	JmsFleetErrorDataSourceRepresentation = map[string]interface{}{
		"compartment_id":            acctest.Representation{RepType: acctest.Optional, Create: JmsFleetErrorCompartmentId},
		"compartment_id_in_subtree": acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"fleet_id":                  acctest.Representation{RepType: acctest.Optional, Create: JmsFleetErrorFleetId},
	}

	JmsFleetErrorResourceConfig = acctest.GenerateResourceFromRepresentationMap("oci_jms_fleet", "test_fleet", acctest.Required, acctest.Create, JmsFleetRepresentation)
)

// issue-routing-tag: jms/default
func TestJmsFleetErrorResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestJmsFleetErrorResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	datasourceName := "data.oci_jms_fleet_errors.test_fleet_errors"

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify create
		// note: we cannot write test for this case because
		// we don't have create API.

		// verify update
		// note: we cannot write test for this case because
		// we don't have update API.

		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap(
					"oci_jms_fleet_errors",
					"test_fleet_errors",
					acctest.Optional,
					acctest.Create,
					JmsFleetErrorDataSourceRepresentation) +
				JmsFleetErrorResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", JmsFleetErrorCompartmentId),
				resource.TestCheckResourceAttr(datasourceName, "compartment_id_in_subtree", "false"),
				resource.TestCheckResourceAttr(datasourceName, "fleet_id", JmsFleetErrorFleetId),

				resource.TestCheckResourceAttrSet(datasourceName, "fleet_error_collection.#"),
				// we can only verify that response contain zero items because we are using dummy test data values
				// we cannot use actual values because it requires create API.
				resource.TestCheckResourceAttr(datasourceName, "fleet_error_collection.0.items.#", "0"),
			),
		},
		// verify singular datasource
		// note: we cannot write test to verify singular data source because
		// crypto analysis processing requires create API.
	})
}
