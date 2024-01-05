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
	JmsFleetDiagnosesCompartmentId  = utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	JmsFleetDiagnosesLogGroupId     = utils.GetEnvSettingWithBlankDefault("fleet_log_group_ocid")
	JmsFleetDiagnosesInventoryLogId = utils.GetEnvSettingWithBlankDefault("fleet_inventory_log_ocid")
	JmsFleetDiagnosesOperationLogId = utils.GetEnvSettingWithBlankDefault("fleet_operation_log_ocid")

	JmsFleetDiagnosesResourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: JmsFleetDiagnosesCompartmentId},
		"display_name":   acctest.Representation{RepType: acctest.Required, Create: `Created Fleet for Diagnoses`},
		"description":    acctest.Representation{RepType: acctest.Optional, Create: `Created Fleet for Diagnoses`},
		"inventory_log": acctest.RepresentationGroup{RepType: acctest.Required, Group: map[string]interface{}{
			"log_group_id": acctest.Representation{
				RepType: acctest.Required,
				Create:  JmsFleetDiagnosesLogGroupId,
				Update:  JmsFleetDiagnosesLogGroupId,
			},
			"log_id": acctest.Representation{
				RepType: acctest.Required,
				Create:  JmsFleetDiagnosesInventoryLogId,
				Update:  JmsFleetDiagnosesInventoryLogId,
			},
		}},
		"operation_log": acctest.RepresentationGroup{RepType: acctest.Optional, Group: map[string]interface{}{
			"log_group_id": acctest.Representation{
				RepType: acctest.Required,
				Create:  JmsFleetDiagnosesLogGroupId,
				Update:  JmsFleetDiagnosesLogGroupId,
			},
			"log_id": acctest.Representation{
				RepType: acctest.Required,
				Create:  JmsFleetDiagnosesOperationLogId,
				Update:  JmsFleetDiagnosesOperationLogId,
			},
		}},
	}

	JmsFleetDiagnosesDataSourceRepresentation = map[string]interface{}{
		"fleet_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_jms_fleet.test_fleet.id}`},
	}
)

// issue-routing-tag: jms/default
func TestJmsFleetDiagnosesResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestJmsFleetDiagnosesResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	datasourceName := "data.oci_jms_fleet_diagnoses.test_fleet_diagnoses"

	acctest.ResourceTest(t, nil, []resource.TestStep{

		// verify datasource
		{
			Config: config +
				acctest.GenerateResourceFromRepresentationMap(
					"oci_jms_fleet",
					"test_fleet",
					acctest.Optional,
					acctest.Create,
					JmsFleetDiagnosesResourceRepresentation,
				) +
				acctest.GenerateDataSourceFromRepresentationMap(
					"oci_jms_fleet_diagnoses",
					"test_fleet_diagnoses",
					acctest.Optional,
					acctest.Create,
					JmsFleetDiagnosesDataSourceRepresentation,
				),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "fleet_id"),

				resource.TestCheckResourceAttrSet(datasourceName, "fleet_diagnosis_collection.#"),
				// we expect the diagnosis to return zero record because
				// the fleet uses existing inventory log and operation logs,
				// and it also does ot have any advanced feature enabled.
				resource.TestCheckResourceAttr(datasourceName, "fleet_diagnosis_collection.0.items.#", "0"),
			),
		},
	})
}

// clean up Fleet resource after test
func init() {
	if acctest.DependencyGraph == nil {
		acctest.InitDependencyGraph()
	}
	if !acctest.InSweeperExcludeList("JmsFleetDiagnoses") {
		resource.AddTestSweepers("JmsFleetDiagnoses", &resource.Sweeper{
			Name:         "JmsFleetDiagnoses",
			Dependencies: acctest.DependencyGraph["fleet"],
			F:            sweepJmsFleetResource,
		})
	}
}
