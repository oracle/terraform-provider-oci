// Copyright (c) 2017, 2019, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	workRequestDataSourceRepresentation = map[string]interface{}{
		"compartment_id": Representation{repType: Required, create: `${var.compartment_id}`},
		"cluster_id":     Representation{repType: Optional, create: `${oci_containerengine_cluster.test_cluster.id}`},
		"resource_id":    Representation{repType: Optional, create: `${oci_containerengine_cluster.test_cluster.id}`},
		"resource_type":  Representation{repType: Optional, create: `CLUSTER`},
		"status":         Representation{repType: Optional, create: []string{}},
	}

	WorkRequestResourceConfig = ClusterRequiredOnlyResource
)

func TestContainerengineWorkRequestResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestContainerengineWorkRequestResource_basic")
	defer httpreplay.SaveScenario()

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
				Config: config +
					generateDataSourceFromRepresentationMap("oci_containerengine_work_requests", "test_work_requests", Optional, Create, workRequestDataSourceRepresentation) +
					compartmentIdVariableStr + WorkRequestResourceConfig,
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
