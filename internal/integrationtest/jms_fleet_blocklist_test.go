// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
)

var (
	JmsFleetBlocklistDataSourceRepresentation = map[string]interface{}{
		"fleet_id":  acctest.Representation{RepType: acctest.Required, Create: JmsFleetId},
		"operation": acctest.Representation{RepType: acctest.Optional, Create: `DELETE_JAVA_INSTALLATION`},
	}
)

// issue-routing-tag: jms/default
func TestJmsFleetBlocklistResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestJmsFleetBlocklistResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()
	datasourceName := "data.oci_jms_fleet_blocklists.test_fleet_blocklists"

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap(
					"oci_jms_fleet_blocklists",
					"test_fleet_blocklists",
					acctest.Optional,
					acctest.Create,
					JmsFleetBlocklistDataSourceRepresentation,
				),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "fleet_id"),
				resource.TestCheckResourceAttr(datasourceName, "operation", "DELETE_JAVA_INSTALLATION"),

				resource.TestCheckResourceAttr(datasourceName, "items.#", "0"),
			),
		},
	})
}
