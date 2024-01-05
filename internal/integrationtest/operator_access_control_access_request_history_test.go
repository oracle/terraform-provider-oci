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
	accessReqId = utils.GetEnvSettingWithBlankDefault("test_access_req_id")

	OperatorAccessControlOperatorAccessControlAccessRequestHistoryDataSourceRepresentation = map[string]interface{}{
		"access_request_id": acctest.Representation{RepType: acctest.Required, Create: accessReqId},
	}
	OperatorAccessControlOperatorAccessControlAccessRequestHistorySingularDataSourceRepresentation = map[string]interface{}{
		"access_request_id": acctest.Representation{RepType: acctest.Required, Create: accessReqId},
	}

	OperatorAccessControlAccessRequestHistoryResourceConfig = acctest.GenerateDataSourceFromRepresentationMap("oci_operator_access_control_access_requests", "test_access_requests", acctest.Required, acctest.Create, OperatorAccessControlOperatorAccessControlAccessRequestDataSourceRepresentation)
)

// issue-routing-tag: operator_access_control/default
func TestOperatorAccessControlAccessRequestHistoryResource_basic(t *testing.T) {
	t.Skip("Access Requests are created outside customer api. Access requests may not be available all the time")
	httpreplay.SetScenario("TestOperatorAccessControlAccessRequestHistoryResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_operator_access_control_access_request_history.test_access_request_history"
	singularDatasourceName := "data.oci_operator_access_control_access_request_history.test_access_request_history"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_operator_access_control_access_request_history", "test_access_request_history", acctest.Required, acctest.Create, OperatorAccessControlOperatorAccessControlAccessRequestHistoryDataSourceRepresentation) +
				compartmentIdVariableStr + OperatorAccessControlAccessRequestHistoryResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "access_request_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "items.#"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_operator_access_control_access_request_history", "test_access_request_history", acctest.Required, acctest.Create, OperatorAccessControlOperatorAccessControlAccessRequestHistorySingularDataSourceRepresentation) +
				compartmentIdVariableStr + OperatorAccessControlAccessRequestHistoryResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "access_request_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "items.#"),
			),
		},
	})
}
