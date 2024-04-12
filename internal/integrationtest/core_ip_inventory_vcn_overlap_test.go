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
	CoreIpInventoryVcnOverlapDataSourceRepresentation = map[string]interface{}{
		"compartment_list": acctest.Representation{RepType: acctest.Required, Create: []string{`${var.compartment_id}`}},
		"region_list":      acctest.Representation{RepType: acctest.Required, Create: []string{`${var.region_id}`}},
		"vcn_id":           acctest.Representation{RepType: acctest.Required, Create: `${var.vcn_id}`},
	}

	CoreIpInventoryVcnOverlapResourceConfig = acctest.GenerateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", acctest.Required, acctest.Create, CoreVcnRepresentation)
)

// issue-routing-tag: core/vcnip
func TestCoreIpInventoryVcnOverlapResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestCoreIpInventoryVcnOverlapResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	vcnId := utils.GetEnvSettingWithBlankDefault("vcn_ocid")
	vncIdVariableStr := fmt.Sprintf("variable \"vcn_id\" { default = \"%s\" }\n", vcnId)

	regionId := utils.GetEnvSettingWithBlankDefault("region_id")
	regionIdVariableStr := fmt.Sprintf("variable \"region_id\" { default = \"%s\" }\n", regionId)

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_core_ip_inventory_vcn_overlaps.test_ip_inventory_vcn_overlaps"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_core_ip_inventory_vcn_overlaps", "test_ip_inventory_vcn_overlaps", acctest.Required, acctest.Create, CoreIpInventoryVcnOverlapDataSourceRepresentation) +
				compartmentIdVariableStr + vncIdVariableStr + regionIdVariableStr,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "ip_inventory_vcn_overlap_summary.0.overlapping_vcn_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "ip_inventory_vcn_overlap_summary.0.overlapping_vcn_name"),
				resource.TestCheckResourceAttr(datasourceName, "ip_inventory_vcn_overlap_summary.0.overlapping_cidr", "10.0.0.0/16"),
				resource.TestCheckResourceAttr(datasourceName, "ip_inventory_vcn_overlap_summary.0.cidr", "10.0.0.0/16"),
				resource.TestCheckResourceAttrSet(datasourceName, "last_updated_timestamp"),
				resource.TestCheckResourceAttr(datasourceName, "message", ""),
				resource.TestCheckResourceAttrSet(datasourceName, "overlap_count"),
			),
		},
	})
}
