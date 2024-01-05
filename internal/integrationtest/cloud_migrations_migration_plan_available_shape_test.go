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
	CloudMigrationsCloudMigrationsMigrationPlanAvailableShapeSingularDataSourceRepresentation = map[string]interface{}{
		"migration_plan_id":    acctest.Representation{RepType: acctest.Required, Create: `${oci_cloud_migrations_migration_plan.test_migration_plan.id}`},
		"availability_domain":  acctest.Representation{RepType: acctest.Optional, Create: CloudMigrationsAvailabilityDomain},
		"compartment_id":       acctest.Representation{RepType: acctest.Optional, Create: `${var.compartment_id}`},
		"reserved_capacity_id": acctest.Representation{RepType: acctest.Optional, Create: `${oci_cloud_migrations_reserved_capacity.test_reserved_capacity.id}`},
	}

	CloudMigrationsCloudMigrationsMigrationPlanAvailableShapeDataSourceRepresentation = map[string]interface{}{
		"migration_plan_id":    acctest.Representation{RepType: acctest.Required, Create: `${oci_cloud_migrations_migration_plan.test_migration_plan.id}`},
		"availability_domain":  acctest.Representation{RepType: acctest.Optional, Create: CloudMigrationsAvailabilityDomain},
		"compartment_id":       acctest.Representation{RepType: acctest.Optional, Create: `${var.compartment_id}`},
		"reserved_capacity_id": acctest.Representation{RepType: acctest.Optional, Create: `${oci_cloud_migrations_reserved_capacity.test_reserved_capacity.id}`},
	}

	CloudMigrationsMigrationPlanAvailableShapeResourceConfig = acctest.GenerateResourceFromRepresentationMap("oci_cloud_migrations_migration_plan", "test_migration_plan", acctest.Required, acctest.Create, CloudMigrationsMigrationPlanRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_cloud_migrations_migration", "test_migration", acctest.Required, acctest.Create, CloudMigrationsMigrationRepresentation)
)

// issue-routing-tag: cloud_migrations/default
func TestCloudMigrationsMigrationPlanAvailableShapeResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestCloudMigrationsMigrationPlanAvailableShapeResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_cloud_migrations_migration_plan_available_shapes.test_migration_plan_available_shapes"
	singularDatasourceName := "data.oci_cloud_migrations_migration_plan_available_shape.test_migration_plan_available_shape"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_cloud_migrations_migration_plan_available_shapes", "test_migration_plan_available_shapes", acctest.Required, acctest.Create, CloudMigrationsCloudMigrationsMigrationPlanAvailableShapeDataSourceRepresentation) +
				compartmentIdVariableStr + CloudMigrationsMigrationPlanAvailableShapeResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "available_shapes_collection.#"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_cloud_migrations_migration_plan_available_shape", "test_migration_plan_available_shape", acctest.Required, acctest.Create, CloudMigrationsCloudMigrationsMigrationPlanAvailableShapeSingularDataSourceRepresentation) +
				compartmentIdVariableStr + CloudMigrationsMigrationPlanAvailableShapeResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "migration_plan_id"),
			),
		},
	})
}
