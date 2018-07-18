// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

const (
	WorkRequestLogEntryResourceConfig = WorkRequestLogEntryResourceDependencies + `

data "oci_containerengine_work_requests" "test_work_requests" {
	#Required
	compartment_id = "${var.compartment_id}"

	#Optional
	cluster_id = "${oci_containerengine_cluster.test_cluster.id}"
	resource_id = "${oci_containerengine_cluster.test_cluster.id}"
	resource_type = "${var.work_request_resource_type}"
	status = "${var.work_request_status}"
}
`
	WorkRequestLogEntryPropertyVariables = `

`
	WorkRequestLogEntryResourceDependencies = WorkRequestPropertyVariables + WorkRequestResourceConfig
)

func TestContainerengineWorkRequestLogEntryResource_basic(t *testing.T) {
	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getRequiredEnvSetting("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_containerengine_work_request_log_entries.test_work_request_log_entries"

	resource.Test(t, resource.TestCase{
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		Steps: []resource.TestStep{
			// verify datasource
			{
				Config: config + `

data "oci_containerengine_work_request_log_entries" "test_work_request_log_entries" {
	#Required
	compartment_id = "${var.compartment_id}"
	work_request_id = "${lookup(data.oci_containerengine_work_requests.test_work_requests.work_requests[0], "id")}"
}
                ` + compartmentIdVariableStr + WorkRequestLogEntryResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttrSet(datasourceName, "work_request_id"),

					resource.TestCheckResourceAttrSet(datasourceName, "work_request_log_entries.#"),
					resource.TestCheckResourceAttrSet(datasourceName, "work_request_log_entries.0.message"),
					resource.TestCheckResourceAttrSet(datasourceName, "work_request_log_entries.0.timestamp"),
				),
			},
		},
	})
}
