// Copyright (c) 2017, 2020, Oracle and/or its affiliates. All rights reserved.
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
	dbSystemShapeDataSourceRepresentation = map[string]interface{}{
		"compartment_id":      Representation{repType: Required, create: `${var.compartment_id}`},
		"availability_domain": Representation{repType: Optional, create: `${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}`},
	}

	DbSystemShapeResourceConfig = AvailabilityDomainConfig
)

func TestDatabaseDbSystemShapeResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseDbSystemShapeResource_basic")
	defer httpreplay.SaveScenario()

	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_database_db_system_shapes.test_db_system_shapes"

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		Steps: []resource.TestStep{
			// verify datasource
			{
				Config: config +
					generateDataSourceFromRepresentationMap("oci_database_db_system_shapes", "test_db_system_shapes", Optional, Create, dbSystemShapeDataSourceRepresentation) +
					compartmentIdVariableStr + DbSystemShapeResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(datasourceName, "availability_domain"),
					resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),

					resource.TestCheckResourceAttrSet(datasourceName, "db_system_shapes.#"),
					resource.TestCheckResourceAttrSet(datasourceName, "db_system_shapes.0.available_core_count"),
					resource.TestCheckResourceAttrSet(datasourceName, "db_system_shapes.0.available_data_storage_in_tbs"),
					resource.TestCheckResourceAttrSet(datasourceName, "db_system_shapes.0.available_db_node_storage_in_gbs"),
					resource.TestCheckResourceAttrSet(datasourceName, "db_system_shapes.0.available_memory_in_gbs"),
					resource.TestCheckResourceAttrSet(datasourceName, "db_system_shapes.0.core_count_increment"),
					resource.TestCheckResourceAttrSet(datasourceName, "db_system_shapes.0.maximum_node_count"),
					resource.TestCheckResourceAttrSet(datasourceName, "db_system_shapes.0.min_core_count_per_node"),
					resource.TestCheckResourceAttrSet(datasourceName, "db_system_shapes.0.min_data_storage_in_tbs"),
					resource.TestCheckResourceAttrSet(datasourceName, "db_system_shapes.0.min_db_node_storage_per_node_in_gbs"),
					resource.TestCheckResourceAttrSet(datasourceName, "db_system_shapes.0.min_memory_per_node_in_gbs"),
					resource.TestCheckResourceAttrSet(datasourceName, "db_system_shapes.0.minimum_core_count"),
					resource.TestCheckResourceAttrSet(datasourceName, "db_system_shapes.0.minimum_node_count"),
					resource.TestCheckResourceAttrSet(datasourceName, "db_system_shapes.0.name"),
					resource.TestCheckResourceAttrSet(datasourceName, "db_system_shapes.0.shape"),
					resource.TestCheckResourceAttrSet(datasourceName, "db_system_shapes.0.shape_family"),
				),
			},
		},
	})
}
