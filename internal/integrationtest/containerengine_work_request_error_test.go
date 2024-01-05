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
	ContainerengineContainerengineWorkRequestErrorDataSourceRepresentation = map[string]interface{}{
		"compartment_id":  acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"work_request_id": acctest.Representation{RepType: acctest.Required, Create: `${lookup(data.oci_containerengine_work_requests.test_work_requests.work_requests[0], "id")}`},
	}

	ContainerengineWorkRequestErrorResourceConfig = ContainerengineWorkRequestResourceConfig +
		acctest.GenerateDataSourceFromRepresentationMap("oci_containerengine_work_requests", "test_work_requests", acctest.Optional, acctest.Create, ContainerengineContainerengineWorkRequestDataSourceRepresentation)
)

// issue-routing-tag: containerengine/default
func TestContainerengineWorkRequestErrorResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestContainerengineWorkRequestErrorResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_containerengine_work_request_errors.test_work_request_errors"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_containerengine_work_request_errors", "test_work_request_errors", acctest.Required, acctest.Create, ContainerengineContainerengineWorkRequestErrorDataSourceRepresentation) +
				compartmentIdVariableStr + ContainerengineWorkRequestErrorResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(datasourceName, "work_request_id"),

				resource.TestCheckResourceAttrSet(datasourceName, "work_request_errors.#"),
			),
		},
	})
}
