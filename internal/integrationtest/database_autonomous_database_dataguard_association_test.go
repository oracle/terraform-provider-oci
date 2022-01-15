// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"testing"

	"github.com/terraform-providers/terraform-provider-oci/internal/acctest"
	"github.com/terraform-providers/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	autonomousDatabaseDataguardAssociationSingularDataSourceRepresentation = map[string]interface{}{
		"autonomous_database_dataguard_association_id": acctest.Representation{RepType: acctest.Required, Create: `${data.oci_database_autonomous_database_dataguard_associations.test_autonomous_database_dataguard_associations.autonomous_database_dataguard_associations[0]["id"]}`},
		"autonomous_database_id":                       acctest.Representation{RepType: acctest.Required, Create: `${oci_database_autonomous_database.test_autonomous_database.id}`},
	}

	autonomousDatabaseDataguardAssociationDataSourceRepresentation = map[string]interface{}{
		"autonomous_database_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_database_autonomous_database.test_autonomous_database.id}`},
	}

	autonomousDatabaseDedicatedRepresentationForDataguard = acctest.RepresentationCopyWithNewProperties(
		acctest.RepresentationCopyWithRemovedProperties(acctest.GetUpdatedRepresentationCopy("db_name", acctest.Representation{RepType: acctest.Required, Create: adbDedicatedName}, autonomousDatabaseRepresentation), []string{"license_model", "whitelisted_ips", "db_version", "is_auto_scaling_enabled", "customer_contacts", "kms_key_id", "vault_id", "autonomous_maintenance_schedule_type", "scheduled_operations"}),
		map[string]interface{}{
			"autonomous_container_database_id": acctest.Representation{RepType: acctest.Optional, Create: `${oci_database_autonomous_container_database.test_autonomous_container_database.id}`},
			"is_dedicated":                     acctest.Representation{RepType: acctest.Optional, Create: `true`},
			"display_name":                     acctest.Representation{RepType: acctest.Optional, Create: adDedicatedName},
			"data_safe_status":                 acctest.Representation{RepType: acctest.Optional, Create: `REGISTERED`},
		})
)

// issue-routing-tag: database/dbaas-atp-d
func TestDatabaseAutonomousDatabaseDataguardAssociationResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseAutonomousDatabaseDataguardAssociationResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_database_autonomous_database_dataguard_associations.test_autonomous_database_dataguard_associations"
	singularDatasourceName := "data.oci_database_autonomous_database_dataguard_association.test_autonomous_database_dataguard_association"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", acctest.Optional, acctest.Create, autonomousDatabaseDedicatedRepresentationForDataguard) +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_autonomous_database_dataguard_associations", "test_autonomous_database_dataguard_associations", acctest.Required, acctest.Create, autonomousDatabaseDataguardAssociationDataSourceRepresentation) +
				compartmentIdVariableStr + AutonomousContainerDatabaseDataguardAssociationResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "autonomous_database_id"),

				resource.TestCheckResourceAttrSet(datasourceName, "autonomous_database_dataguard_associations.#"),
				resource.TestCheckResourceAttrSet(datasourceName, "autonomous_database_dataguard_associations.0.autonomous_database_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "autonomous_database_dataguard_associations.0.id"),
				resource.TestCheckResourceAttrSet(datasourceName, "autonomous_database_dataguard_associations.0.is_automatic_failover_enabled"),
				resource.TestCheckResourceAttrSet(datasourceName, "autonomous_database_dataguard_associations.0.peer_autonomous_database_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "autonomous_database_dataguard_associations.0.peer_autonomous_database_life_cycle_state"),
				resource.TestCheckResourceAttrSet(datasourceName, "autonomous_database_dataguard_associations.0.peer_role"),
				resource.TestCheckResourceAttrSet(datasourceName, "autonomous_database_dataguard_associations.0.protection_mode"),
				resource.TestCheckResourceAttrSet(datasourceName, "autonomous_database_dataguard_associations.0.role"),
				resource.TestCheckResourceAttrSet(datasourceName, "autonomous_database_dataguard_associations.0.state"),
				resource.TestCheckResourceAttrSet(datasourceName, "autonomous_database_dataguard_associations.0.time_created"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", acctest.Optional, acctest.Create, autonomousDatabaseDedicatedRepresentationForDataguard) +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_autonomous_database_dataguard_associations", "test_autonomous_database_dataguard_associations", acctest.Required, acctest.Create, autonomousDatabaseDataguardAssociationDataSourceRepresentation) +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_autonomous_database_dataguard_association", "test_autonomous_database_dataguard_association", acctest.Required, acctest.Create, autonomousDatabaseDataguardAssociationSingularDataSourceRepresentation) +
				compartmentIdVariableStr + AutonomousContainerDatabaseDataguardAssociationResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "autonomous_database_dataguard_association_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "autonomous_database_id"),

				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "peer_autonomous_database_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "peer_autonomous_database_life_cycle_state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "peer_role"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "protection_mode"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "role"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
			),
		},
	})
}
