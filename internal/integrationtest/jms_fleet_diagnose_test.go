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
	JmsFleetDiagnosesDataSourceRepresentation = map[string]interface{}{
		"fleet_id": acctest.Representation{RepType: acctest.Required, Create: JmsFleetId},
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
