// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"testing"

	"github.com/oracle/terraform-provider-oci/internal/acctest"
	"github.com/oracle/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"github.com/oracle/terraform-provider-oci/httpreplay"
)

var (
	DatabaseDatabaseAutonomousExadataInfrastructureShapeDataSourceRepresentation = map[string]interface{}{
		"availability_domain": acctest.Representation{RepType: acctest.Required, Create: `${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}`},
		"compartment_id":      acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
	}

	DatabaseAutonomousExadataInfrastructureShapeResourceConfig = AvailabilityDomainConfig
)

// issue-routing-tag: database/dbaas-atp-d
func TestDatabaseAutonomousExadataInfrastructureShapeResource_basic(t *testing.T) {
	t.Skip("Skip this test as AEI and its api no longer exists.")

	httpreplay.SetScenario("TestDatabaseAutonomousExadataInfrastructureShapeResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_database_autonomous_exadata_infrastructure_shapes.test_autonomous_exadata_infrastructure_shapes"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_autonomous_exadata_infrastructure_shapes", "test_autonomous_exadata_infrastructure_shapes", acctest.Required, acctest.Create, DatabaseDatabaseAutonomousExadataInfrastructureShapeDataSourceRepresentation) +
				compartmentIdVariableStr + DatabaseAutonomousExadataInfrastructureShapeResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
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
