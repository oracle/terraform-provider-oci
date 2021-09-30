// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	workRequestErrorDataSourceRepresentation = map[string]interface{}{
		"compartment_id":  Representation{RepType: Required, Create: `${var.compartment_id}`},
		"work_request_id": Representation{RepType: Required, Create: `${lookup(data.oci_containerengine_work_requests.test_work_requests.work_requests[0], "id")}`},
	}

	WorkRequestErrorResourceConfig = WorkRequestResourceConfig +
		GenerateDataSourceFromRepresentationMap("oci_containerengine_work_requests", "test_work_requests", Optional, Create, workRequestDataSourceRepresentation)
)

// issue-routing-tag: containerengine/default
func TestContainerengineWorkRequestErrorResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestContainerengineWorkRequestErrorResource_basic")
	defer httpreplay.SaveScenario()

	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_containerengine_work_request_errors.test_work_request_errors"

	SaveConfigContent("", "", "", t)

	ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				GenerateDataSourceFromRepresentationMap("oci_containerengine_work_request_errors", "test_work_request_errors", Required, Create, workRequestErrorDataSourceRepresentation) +
				compartmentIdVariableStr + WorkRequestErrorResourceConfig,
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(datasourceName, "work_request_id"),

				resource.TestCheckResourceAttrSet(datasourceName, "work_request_errors.#"),
			),
		},
	})
}
