// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	autonomousExadataInfrastructureShapeDataSourceRepresentation = map[string]interface{}{
		"availability_domain": Representation{RepType: Required, Create: `${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}`},
		"compartment_id":      Representation{RepType: Required, Create: `${var.compartment_id}`},
	}

	AutonomousExadataInfrastructureShapeResourceConfig = AvailabilityDomainConfig
)

// issue-routing-tag: database/dbaas-atp-d
func TestDatabaseAutonomousExadataInfrastructureShapeResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseAutonomousExadataInfrastructureShapeResource_basic")
	defer httpreplay.SaveScenario()

	config := ProviderTestConfig()

	compartmentId := GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_database_autonomous_exadata_infrastructure_shapes.test_autonomous_exadata_infrastructure_shapes"

	SaveConfigContent("", "", "", t)

	ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				GenerateDataSourceFromRepresentationMap("oci_database_autonomous_exadata_infrastructure_shapes", "test_autonomous_exadata_infrastructure_shapes", Required, Create, autonomousExadataInfrastructureShapeDataSourceRepresentation) +
				compartmentIdVariableStr + AutonomousExadataInfrastructureShapeResourceConfig,
			Check: ComposeAggregateTestCheckFuncWrapper(
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
	})
}
