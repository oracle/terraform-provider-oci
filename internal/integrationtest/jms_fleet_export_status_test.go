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
	JmsFleetExportStatusFleetId       = utils.GetEnvSettingWithBlankDefault("fleet_ocid")
	JmsFleetExportStatusCompartmentId = utils.GetEnvSettingWithBlankDefault("compartment_ocid")

	JmsFleetExportStatusSingularDataSourceRepresentation = map[string]interface{}{
		"fleet_id": acctest.Representation{RepType: acctest.Required, Create: JmsFleetExportStatusFleetId},
	}
)

// issue-routing-tag: jms/default
func TestJmsFleetExportStatusResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestJmsFleetExportStatusResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	singularDatasourceName := "data.oci_jms_fleet_export_status.test_fleet_export_status"

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap(
					"oci_jms_fleet_export_status",
					"test_fleet_export_status",
					acctest.Required,
					acctest.Create,
					JmsFleetExportStatusSingularDataSourceRepresentation,
				),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "fleet_id"),
			),
		},
	})
}

// clean up Fleet resource after test
func init() {
	if acctest.DependencyGraph == nil {
		acctest.InitDependencyGraph()
	}
	if !acctest.InSweeperExcludeList("JmsFleetExportStatus") {
		resource.AddTestSweepers("JmsFleetExportStatus", &resource.Sweeper{
			Name:         "JmsFleetExportStatus",
			Dependencies: acctest.DependencyGraph["fleet"],
			F:            sweepJmsFleetResource,
		})
	}
}
