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
	CoreIpInventorySubnetSingularDataSourceRepresentation = map[string]interface{}{
		"subnet_id": acctest.Representation{RepType: acctest.Required, Create: `${var.subnet_id}`},
	}

	CoreIpInventorySubnetResourceConfig = acctest.GenerateResourceFromRepresentationMap("oci_core_subnet", "test_subnet", acctest.Required, acctest.Create, CoreSubnetRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", acctest.Required, acctest.Create, CoreVcnRepresentation)
)

// issue-routing-tag: core/vcnip
func TestCoreIpInventorySubnetResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestCoreIpInventorySubnetResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	subnetId := utils.GetEnvSettingWithBlankDefault("subnet_ocid")
	subnetIdVariableStr := fmt.Sprintf("variable \"subnet_id\" { default = \"%s\" }\n", subnetId)

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	singularDatasourceName := "data.oci_core_ip_inventory_subnet.test_ip_inventory_subnet"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_core_ip_inventory_subnet", "test_ip_inventory_subnet", acctest.Required, acctest.Create, CoreIpInventorySubnetSingularDataSourceRepresentation) +
				compartmentIdVariableStr + subnetIdVariableStr,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "subnet_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "compartment_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "ip_inventory_subnet_resource_summary.#"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "ip_inventory_subnet_resource_summary.0.ip_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "ip_inventory_subnet_resource_summary.0.ip_address"),
				resource.TestCheckResourceAttr(singularDatasourceName, "ip_inventory_subnet_resource_summary.0.ip_address_lifetime", "Ephemeral"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "ip_inventory_subnet_resource_summary.0.associated_public_ip"),
				resource.TestCheckResourceAttr(singularDatasourceName, "ip_inventory_subnet_resource_summary.0.public_ip_lifetime", "Ephemeral"),
				resource.TestCheckResourceAttr(singularDatasourceName, "ip_inventory_subnet_resource_summary.0.associated_public_ip_pool", "ORACLE"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "ip_inventory_subnet_resource_summary.0.dns_host_name"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "ip_inventory_subnet_resource_summary.0.assigned_resource_name"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "ip_inventory_subnet_resource_summary.0.assigned_resource_type"),
				resource.TestCheckResourceAttr(singularDatasourceName, "ip_inventory_subnet_resource_summary.0.address_type", "Private_IPv4"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "ip_inventory_subnet_resource_summary.0.assigned_time"),

				resource.TestCheckResourceAttr(singularDatasourceName, "ip_inventory_subnet_resource_summary.1.address_type", "ULA_IPv6"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "ip_inventory_subnet_resource_summary.1.assigned_time"),

				resource.TestCheckResourceAttrSet(singularDatasourceName, "ip_inventory_subnet_count"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "last_updated_timestamp"),
				resource.TestCheckResourceAttr(singularDatasourceName, "message", ""),
			),
		},
	})
}
