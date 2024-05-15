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
	CoreIpInventorySubnetCidrSingularDataSourceRepresentation = map[string]interface{}{
		"subnet_id": acctest.Representation{RepType: acctest.Required, Create: `${var.subnet_id}`},
	}

	CoreIpInventorySubnetCidrResourceConfig = acctest.GenerateResourceFromRepresentationMap("oci_core_subnet", "test_subnet", acctest.Required, acctest.Create, CoreSubnetRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", acctest.Required, acctest.Create, CoreVcnRepresentation)
)

// issue-routing-tag: core/vcnip
func TestCoreIpInventorySubnetCidrResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestCoreIpInventorySubnetCidrResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	subnetId := utils.GetEnvSettingWithBlankDefault("subnet_ocid")
	subnetIdVariableStr := fmt.Sprintf("variable \"subnet_id\" { default = \"%s\" }\n", subnetId)

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	singularDatasourceName := "data.oci_core_ip_inventory_subnet_cidr.test_ip_inventory_subnet_cidr"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_core_ip_inventory_subnet_cidr", "test_ip_inventory_subnet_cidr", acctest.Required, acctest.Create, CoreIpInventorySubnetCidrSingularDataSourceRepresentation) +
				compartmentIdVariableStr + subnetIdVariableStr,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "compartment_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "ip_inventory_cidr_utilization_summary.#", "2"),
				/* Below - Pass result to check as env variable not working */
				resource.TestCheckResourceAttr(singularDatasourceName, "ip_inventory_subnet_cidr_count", "2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "ip_inventory_cidr_utilization_summary.0.address_type", "Private_IPv4"),
				resource.TestCheckResourceAttr(singularDatasourceName, "ip_inventory_cidr_utilization_summary.0.cidr", "10.0.0.0/24"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "ip_inventory_cidr_utilization_summary.0.utilization"),
				resource.TestCheckResourceAttr(singularDatasourceName, "ip_inventory_cidr_utilization_summary.1.address_type", "ULA_IPv6"),
				resource.TestCheckResourceAttr(singularDatasourceName, "ip_inventory_cidr_utilization_summary.1.cidr", "fc00:0000:0000:0000:0000:0000:0000:0000/64"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "ip_inventory_cidr_utilization_summary.1.utilization"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "last_updated_timestamp"),
				resource.TestCheckResourceAttr(singularDatasourceName, "message", ""),
			),
		},
	})
}
