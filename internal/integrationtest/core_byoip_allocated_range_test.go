// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"log"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	CoreCoreByoipAllocatedRangeDataSourceRepresentation = map[string]interface{}{
		"byoip_range_id": acctest.Representation{RepType: acctest.Required, Create: `${var.byoip_range_id}`},
	}

	CoreByoipAllocatedRangeResourceConfig = PublicIpPoolAddCapacityResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_core_public_ip_pool_capacity", "test_public_ip_pool_capacity", acctest.Required, acctest.Create, publicIpPoolCapacityRepresentation)
)

// issue-routing-tag: core/vcnip
func TestCoreByoipAllocatedRangeResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestCoreByoipAllocatedRangeResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_core_byoip_allocated_ranges.test_byoip_allocated_ranges"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// Create dependencies
		{
			Config: config + compartmentIdVariableStr + CoreByoipAllocatedRangeResourceConfig,
			Check: func(s *terraform.State) (err error) {
				log.Printf("[DEBUG] Wait for oci_core_public_ip_pool and oci_core_public_ip_pool_capacity resource to get created first")
				time.Sleep(2 * time.Minute)
				return nil
			},
		},
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_core_byoip_allocated_ranges", "test_byoip_allocated_ranges", acctest.Required, acctest.Create, CoreCoreByoipAllocatedRangeDataSourceRepresentation) +
				compartmentIdVariableStr + CoreByoipAllocatedRangeResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "byoip_range_id"),

				resource.TestCheckResourceAttrSet(datasourceName, "byoip_allocated_range_collection.#"),
				resource.TestCheckResourceAttr(datasourceName, "byoip_allocated_range_collection.0.items.#", "1"),
			),
		},
	})
}
