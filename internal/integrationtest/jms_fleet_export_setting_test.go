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
	JmsFleetExportSettingSingularDataSourceRepresentation = map[string]interface{}{
		"fleet_id": acctest.Representation{RepType: acctest.Required, Create: JmsFleetId},
	}
)

// issue-routing-tag: jms/default
func TestJmsFleetExportSettingResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestJmsFleetExportSettingResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	singularDatasourceName := "data.oci_jms_fleet_export_setting.test_fleet_export_setting"

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap(
					"oci_jms_fleet_export_setting",
					"test_fleet_export_setting",
					acctest.Required,
					acctest.Create,
					JmsFleetExportSettingSingularDataSourceRepresentation,
				),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "fleet_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "is_enabled", "false"),
			),
		},
	})
}
