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
	// before running tests, ensure to set up environment variables used below
	JmsFleetBlocklistWithAdvancedFeature = utils.GetEnvSettingWithBlankDefault("fleet_advanced_feature_ocid")

	JmsFleetBlocklistDataSourceRepresentation = map[string]interface{}{
		"fleet_id":  acctest.Representation{RepType: acctest.Required, Create: JmsFleetBlocklistWithAdvancedFeature},
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
