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
	JmsFleetExportStatusCompartmentId  = utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	JmsFleetExportStatusLogGroupId     = utils.GetEnvSettingWithBlankDefault("fleet_log_group_ocid")
	JmsFleetExportStatusInventoryLogId = utils.GetEnvSettingWithBlankDefault("fleet_inventory_log_ocid")
	JmsFleetExportStatusOperationLogId = utils.GetEnvSettingWithBlankDefault("fleet_operation_log_ocid")

	JmsFleetExportStatusResourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: JmsFleetExportStatusCompartmentId},
		"display_name":   acctest.Representation{RepType: acctest.Required, Create: `Created Fleet for Export Status`},
		"description":    acctest.Representation{RepType: acctest.Optional, Create: `Created Fleet for Export Status`},
		"inventory_log": acctest.RepresentationGroup{
			RepType: acctest.Required,
			Group: map[string]interface{}{
				"log_group_id": acctest.Representation{
					RepType: acctest.Required,
					Create:  JmsFleetExportStatusLogGroupId,
					Update:  JmsFleetExportStatusLogGroupId,
				},
				"log_id": acctest.Representation{
					RepType: acctest.Required,
					Create:  JmsFleetExportStatusInventoryLogId,
					Update:  JmsFleetExportStatusInventoryLogId,
				},
			}},
		"operation_log": acctest.RepresentationGroup{
			RepType: acctest.Optional,
			Group: map[string]interface{}{
				"log_group_id": acctest.Representation{
					RepType: acctest.Required,
					Create:  JmsFleetExportStatusLogGroupId,
					Update:  JmsFleetExportStatusLogGroupId,
				},
				"log_id": acctest.Representation{
					RepType: acctest.Required,
					Create:  JmsFleetExportStatusOperationLogId,
					Update:  JmsFleetExportStatusOperationLogId,
				},
			}},
	}

	JmsFleetExportStatusSingularDataSourceRepresentation = map[string]interface{}{
		"fleet_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_jms_fleet.test_fleet.id}`},
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
				acctest.GenerateResourceFromRepresentationMap(
					"oci_jms_fleet",
					"test_fleet",
					acctest.Optional,
					acctest.Create,
					JmsFleetExportSettingResourceRepresentation,
				) +
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
