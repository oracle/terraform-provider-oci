// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0
package integrationtest

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	CloudGuardCloudGuardDataSourceEventSingularDataSourceRepresentation = map[string]interface{}{
		"data_source_id": acctest.Representation{RepType: acctest.Required, Create: `${var.dataSource_id}`},
		"region":         acctest.Representation{RepType: acctest.Optional, Create: `us-phoenix-1`},
	}
	CloudGuardCloudGuardDataSourceEventDataSourceRepresentation = map[string]interface{}{
		"data_source_id": acctest.Representation{RepType: acctest.Required, Create: `${var.dataSource_id}`},
		"region":         acctest.Representation{RepType: acctest.Optional, Create: `us-phoenix-1`},
	}
)

// issue-routing-tag: cloud_guard/default
func TestCloudGuardDataSourceEventResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestCloudGuardDataSourceEventResource_basic")
	defer httpreplay.SaveScenario()
	config := acctest.ProviderTestConfig()
	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	dataSourceId := utils.GetEnvSettingWithBlankDefault("dataSource_ocid")
	dataSourceIdVariableStr := fmt.Sprintf("variable \"dataSource_id\" { default = \"%s\" }\n", dataSourceId)
	datasourceName := "data.oci_cloud_guard_data_source_events.test_data_source_events"
	singularDatasourceName := "data.oci_cloud_guard_data_source_event.test_data_source_event"
	acctest.SaveConfigContent("", "", "", t)
	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_cloud_guard_data_source_events", "test_data_source_events", acctest.Required, acctest.Create, CloudGuardCloudGuardDataSourceEventDataSourceRepresentation) +
				compartmentIdVariableStr + dataSourceIdVariableStr,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "data_source_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "data_source_event_collection.#"),
				resource.TestCheckResourceAttr(datasourceName, "data_source_event_collection.0.items.#", "2"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_cloud_guard_data_source_event", "test_data_source_event", acctest.Required, acctest.Create, CloudGuardCloudGuardDataSourceEventSingularDataSourceRepresentation) +
				compartmentIdVariableStr + dataSourceIdVariableStr,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "data_source_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "items.#", "2"),
			),
		},
	})
}
