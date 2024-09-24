// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"

	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	FleetAppsManagementInventoryResourceDataSourceRepresentation = map[string]interface{}{
		"compartment_id":          acctest.Representation{RepType: acctest.Required, Create: `${var.tenancy_ocid}`},
		"resource_compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.tenancy_ocid}`},
		"inventory_properties":    acctest.Representation{RepType: acctest.Optional, Create: []string{`Instance.id=${oci_core_instance.test_instance.id}`}},
		"matching_criteria":       acctest.Representation{RepType: acctest.Required, Create: `ANY`},
		"resource_region":         acctest.Representation{RepType: acctest.Required, Create: `${var.region}`},
		"state":                   acctest.Representation{RepType: acctest.Optional, Create: `RUNNING`},
	}

	FleetAppsManagementInventoryResourceResourceConfig = acctest.GenerateResourceFromRepresentationMap("oci_core_subnet", "test_subnet", acctest.Required, acctest.Create, CoreSubnetRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", acctest.Required, acctest.Create, CoreVcnRepresentation) +
		utils.OciImageIdsVariable +
		acctest.GenerateResourceFromRepresentationMap("oci_core_instance", "test_instance", acctest.Required, acctest.Create, CoreInstanceRepresentation) +
		AvailabilityDomainConfig
)

// issue-routing-tag: fleet_apps_management/default
func TestFleetAppsManagementInventoryResourceResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestFleetAppsManagementInventoryResourceResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("tenancy_ocid")
	region := utils.GetEnvSettingWithBlankDefault("region")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_fleet_apps_management_inventory_resources.test_inventory_resources"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_fleet_apps_management_inventory_resources", "test_inventory_resources", acctest.Required, acctest.Create, FleetAppsManagementInventoryResourceDataSourceRepresentation) +
				compartmentIdVariableStr + FleetAppsManagementInventoryResourceResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "matching_criteria", "ANY"),
				resource.TestCheckResourceAttrSet(datasourceName, "resource_compartment_id"),
				resource.TestCheckResourceAttr(datasourceName, "resource_region", region),

				resource.TestCheckResourceAttrSet(datasourceName, "inventory_resource_collection.#"),
				resource.TestMatchResourceAttr(datasourceName, "inventory_resource_collection.0.items.#", regexp.MustCompile("[1-9][0-9]*")),
			),
		},
	})
}
