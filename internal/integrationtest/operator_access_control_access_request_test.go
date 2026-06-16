// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	OperatorAccessControlOperatorAccessControlAccessRequestSingularDataSourceRepresentation = map[string]interface{}{
		"access_request_id": acctest.Representation{RepType: acctest.Required, Create: `${data.oci_operator_access_control_access_requests.test_access_requests.access_request_collection.0.items.0.id}`},
	}

	OperatorAccessControlOperatorAccessControlAccessRequestDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"num_days":       acctest.Representation{RepType: acctest.Optional, Create: `10`},
		"resource_name":  acctest.Representation{RepType: acctest.Optional, Create: `${oci_cloud_guard_resource.test_resource.name}`},
		"resource_type":  acctest.Representation{RepType: acctest.Optional, Create: `resourceType`},
		"state":          acctest.Representation{RepType: acctest.Optional, Create: `AVAILABLE`},
		"time_end":       acctest.Representation{RepType: acctest.Optional, Create: `timeEnd`},
		"time_start":     acctest.Representation{RepType: acctest.Optional, Create: `timeStart`},
	}

	OperatorAccessControlAccessRequestResourceConfig = ""
)

// issue-routing-tag: operator_access_control/default
func TestOperatorAccessControlAccessRequestResource_basic(t *testing.T) {
	// t.Skip("Access Requests are created outside customer api. Access requests may not be available all the time")
	httpreplay.SetScenario("TestOperatorAccessControlAccessRequestResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_operator_access_control_access_requests.test_access_requests"
	singularDatasourceName := "data.oci_operator_access_control_access_request.test_access_request"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource step
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_operator_access_control_access_requests", "test_access_requests", acctest.Required, acctest.Create, OperatorAccessControlOperatorAccessControlAccessRequestDataSourceRepresentation) +
				compartmentIdVariableStr + OperatorAccessControlAccessRequestResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "access_request_collection.#", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "access_request_collection.0.items.#"),
				resource.TestCheckResourceAttrSet(datasourceName, "access_request_collection.0.items.0.id"),
				resource.TestCheckResourceAttrSet(datasourceName, "access_request_collection.0.items.0.request_id"),
				resource.TestCheckResourceAttr(datasourceName, "access_request_collection.0.items.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(datasourceName, "access_request_collection.0.items.0.resource_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "access_request_collection.0.items.0.resource_name"),
				resource.TestCheckResourceAttrSet(datasourceName, "access_request_collection.0.items.0.resource_type"),
				resource.TestCheckResourceAttrSet(datasourceName, "access_request_collection.0.items.0.state"),
				resource.TestCheckResourceAttrSet(datasourceName, "access_request_collection.0.items.0.severity"),
				resource.TestCheckResourceAttrSet(datasourceName, "access_request_collection.0.items.0.time_of_creation"),
				resource.TestCheckResourceAttrSet(datasourceName, "access_request_collection.0.items.0.time_of_modification"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_operator_access_control_access_requests", "test_access_requests", acctest.Required, acctest.Create, OperatorAccessControlOperatorAccessControlAccessRequestDataSourceRepresentation) +
				acctest.GenerateDataSourceFromRepresentationMap("oci_operator_access_control_access_request", "test_access_request", acctest.Required, acctest.Create, OperatorAccessControlOperatorAccessControlAccessRequestSingularDataSourceRepresentation) +
				compartmentIdVariableStr + OperatorAccessControlAccessRequestResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "access_reason_summary"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "compartment_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "duration"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "extend_duration"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "is_auto_approved"),
				resource.TestCheckResourceAttrPair(singularDatasourceName, "id", datasourceName, "access_request_collection.0.items.0.id"),
				resource.TestCheckResourceAttrPair(singularDatasourceName, "request_id", datasourceName, "access_request_collection.0.items.0.request_id"),
				resource.TestCheckResourceAttrPair(singularDatasourceName, "compartment_id", datasourceName, "access_request_collection.0.items.0.compartment_id"),
				resource.TestCheckResourceAttrPair(singularDatasourceName, "resource_id", datasourceName, "access_request_collection.0.items.0.resource_id"),
				resource.TestCheckResourceAttrPair(singularDatasourceName, "resource_name", datasourceName, "access_request_collection.0.items.0.resource_name"),
				resource.TestCheckResourceAttrPair(singularDatasourceName, "resource_type", datasourceName, "access_request_collection.0.items.0.resource_type"),
				resource.TestCheckResourceAttrPair(singularDatasourceName, "severity", datasourceName, "access_request_collection.0.items.0.severity"),
				resource.TestCheckResourceAttrPair(singularDatasourceName, "state", datasourceName, "access_request_collection.0.items.0.state"),
				resource.TestCheckResourceAttrPair(singularDatasourceName, "time_of_creation", datasourceName, "access_request_collection.0.items.0.time_of_creation"),
				resource.TestCheckResourceAttrPair(singularDatasourceName, "time_of_modification", datasourceName, "access_request_collection.0.items.0.time_of_modification"),
				resource.TestCheckResourceAttr(singularDatasourceName, "sub_resource_list.#", "0"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_of_user_creation"),
			),
		},
	})
}
