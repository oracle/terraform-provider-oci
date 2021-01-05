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
	autonomousExadataInfrastructureShapeDataSourceRepresentation = map[string]interface{}{
		"availability_domain": Representation{repType: Required, create: `${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}`},
		"compartment_id":      Representation{repType: Required, create: `${var.compartment_id}`},
	}

	AutonomousExadataInfrastructureShapeResourceConfig = AvailabilityDomainConfig
)

func TestDatabaseAutonomousExadataInfrastructureShapeResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseAutonomousExadataInfrastructureShapeResource_basic")
	defer httpreplay.SaveScenario()

	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_database_autonomous_exadata_infrastructure_shapes.test_autonomous_exadata_infrastructure_shapes"

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		Steps: []resource.TestStep{
			// verify datasource
			{
				Config: config +
					generateDataSourceFromRepresentationMap("oci_database_autonomous_exadata_infrastructure_shapes", "test_autonomous_exadata_infrastructure_shapes", Required, Create, autonomousExadataInfrastructureShapeDataSourceRepresentation) +
					compartmentIdVariableStr + AutonomousExadataInfrastructureShapeResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(datasourceName, "availability_domain"),
					resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),

					resource.TestCheckResourceAttrSet(datasourceName, "autonomous_exadata_infrastructure_shapes.#"),
					resource.TestCheckResourceAttrSet(datasourceName, "autonomous_exadata_infrastructure_shapes.0.available_core_count"),
					resource.TestCheckResourceAttrSet(datasourceName, "autonomous_exadata_infrastructure_shapes.0.core_count_increment"),
					resource.TestCheckResourceAttrSet(datasourceName, "autonomous_exadata_infrastructure_shapes.0.maximum_node_count"),
					resource.TestCheckResourceAttrSet(datasourceName, "autonomous_exadata_infrastructure_shapes.0.minimum_core_count"),
					resource.TestCheckResourceAttrSet(datasourceName, "autonomous_exadata_infrastructure_shapes.0.minimum_node_count"),
					resource.TestCheckResourceAttrSet(datasourceName, "autonomous_exadata_infrastructure_shapes.0.name"),
				),
			},
		},
	})
}
