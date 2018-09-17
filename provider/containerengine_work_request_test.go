// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

const (
	WorkRequestResourceConfig = WorkRequestResourceDependencies + `

`
	WorkRequestPropertyVariables = `
variable "work_request_resource_type" { default = "CLUSTER" }
variable "work_request_status" { default = [] }

`
	WorkRequestResourceDependencies = ClusterPropertyVariables + ClusterRequiredOnlyResource
)

func TestContainerengineWorkRequestResource_basic(t *testing.T) {
	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_containerengine_work_requests.test_work_requests"

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		Steps: []resource.TestStep{
			// verify datasource
			{
				Config: config + `

data "oci_containerengine_work_requests" "test_work_requests" {
	#Required
	compartment_id = "${var.compartment_id}"

	#Optional
	cluster_id = "${oci_containerengine_cluster.test_cluster.id}"
	resource_id = "${oci_containerengine_cluster.test_cluster.id}"
	resource_type = "${var.work_request_resource_type}"
	status = "${var.work_request_status}"
}
                ` + compartmentIdVariableStr + WorkRequestPropertyVariables + WorkRequestResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(datasourceName, "cluster_id"),
					resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),

					resource.TestCheckResourceAttrSet(datasourceName, "work_requests.#"),
					resource.TestCheckResourceAttrSet(datasourceName, "work_requests.0.id"),
				),
			},
		},
	})
}
