// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"terraform-provider-oci/httpreplay"
	"terraform-provider-oci/internal/acctest"
	"terraform-provider-oci/internal/utils"
)

var (
	fleetBlocklistDataSourceRepresentation = map[string]interface{}{
		"fleet_id":  acctest.Representation{RepType: acctest.Required, Create: `${oci_jms_fleet.test_fleet.id}`},
		"operation": acctest.Representation{RepType: acctest.Optional, Create: `DELETE_JAVA_INSTALLATION`},
	}

	FleetBlocklistResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_jms_fleet", "test_fleet", acctest.Required, acctest.Create, fleetRepresentation)
)

// issue-routing-tag: jms/default
func TestJmsFleetBlocklistResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestJmsFleetBlocklistResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_jms_fleet_blocklists.test_fleet_blocklists"

	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_jms_fleet_blocklists", "test_fleet_blocklists", acctest.Optional, acctest.Update, fleetBlocklistDataSourceRepresentation) +
				compartmentIdVariableStr + FleetBlocklistResourceDependencies,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "fleet_id"),
				resource.TestCheckResourceAttr(datasourceName, "operation", "DELETE_JAVA_INSTALLATION"),

				resource.TestCheckResourceAttr(datasourceName, "items.#", "0"),
			),
		},
	})
}

func init() {
	if acctest.DependencyGraph == nil {
		acctest.InitDependencyGraph()
	}
	if !acctest.InSweeperExcludeList("JmsFleetBlocklist") {
		resource.AddTestSweepers("JmsFleetBlocklist", &resource.Sweeper{
			Name:         "JmsFleetBlocklist",
			Dependencies: acctest.DependencyGraph["fleet"],
			F:            sweepJmsFleetResource,
		})
	}
}
