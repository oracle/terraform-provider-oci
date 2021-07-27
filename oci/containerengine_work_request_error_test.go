// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	workRequestErrorDataSourceRepresentation = map[string]interface{}{
		"compartment_id":  Representation{repType: Required, create: `${var.compartment_id}`},
		"work_request_id": Representation{repType: Required, create: `${lookup(data.oci_containerengine_work_requests.test_work_requests.work_requests[0], "id")}`},
	}

	WorkRequestErrorResourceConfig = WorkRequestResourceConfig +
		generateDataSourceFromRepresentationMap("oci_containerengine_work_requests", "test_work_requests", Optional, Create, workRequestDataSourceRepresentation)
)

// issue-routing-tag: containerengine/default
func TestContainerengineWorkRequestErrorResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestContainerengineWorkRequestErrorResource_basic")
	defer httpreplay.SaveScenario()

	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_containerengine_work_request_errors.test_work_request_errors"

	saveConfigContent("", "", "", t)

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		Steps: []resource.TestStep{
			// verify datasource
			{
				Config: config +
					generateDataSourceFromRepresentationMap("oci_containerengine_work_request_errors", "test_work_request_errors", Required, Create, workRequestErrorDataSourceRepresentation) +
					compartmentIdVariableStr + WorkRequestErrorResourceConfig,
				Check: ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttrSet(datasourceName, "work_request_id"),

					resource.TestCheckResourceAttrSet(datasourceName, "work_request_errors.#"),
				),
			},
		},
	})
}
