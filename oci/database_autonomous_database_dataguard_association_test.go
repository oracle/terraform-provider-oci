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
	autonomousDatabaseDataguardAssociationSingularDataSourceRepresentation = map[string]interface{}{
		"autonomous_database_dataguard_association_id": Representation{repType: Required, create: `${data.oci_database_autonomous_database_dataguard_associations.test_autonomous_database_dataguard_associations.autonomous_database_dataguard_associations[0]["id"]}`},
		"autonomous_database_id":                       Representation{repType: Required, create: `${oci_database_autonomous_database.test_autonomous_database.id}`},
	}

	autonomousDatabaseDataguardAssociationDataSourceRepresentation = map[string]interface{}{
		"autonomous_database_id": Representation{repType: Required, create: `${oci_database_autonomous_database.test_autonomous_database.id}`},
	}

	autonomousDatabaseDedicatedRepresentationForDataguard = representationCopyWithNewProperties(
		representationCopyWithRemovedProperties(getUpdatedRepresentationCopy("db_name", Representation{repType: Required, create: adbDedicatedName}, autonomousDatabaseRepresentation), []string{"license_model", "whitelisted_ips", "db_version", "is_auto_scaling_enabled"}),
		map[string]interface{}{
			"autonomous_container_database_id": Representation{repType: Optional, create: `${oci_database_autonomous_container_database.test_autonomous_container_database.id}`},
			"is_dedicated":                     Representation{repType: Optional, create: `true`},
			"display_name":                     Representation{repType: Optional, create: adDedicatedName},
			"data_safe_status":                 Representation{repType: Optional, create: `REGISTERED`},
		})
)

func TestDatabaseAutonomousDatabaseDataguardAssociationResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseAutonomousDatabaseDataguardAssociationResource_basic")
	defer httpreplay.SaveScenario()

	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_database_autonomous_database_dataguard_associations.test_autonomous_database_dataguard_associations"
	singularDatasourceName := "data.oci_database_autonomous_database_dataguard_association.test_autonomous_database_dataguard_association"

	saveConfigContent("", "", "", t)

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		Steps: []resource.TestStep{
			// verify datasource
			{
				Config: config +
					generateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", Optional, Create, autonomousDatabaseDedicatedRepresentationForDataguard) +
					generateDataSourceFromRepresentationMap("oci_database_autonomous_database_dataguard_associations", "test_autonomous_database_dataguard_associations", Required, Create, autonomousDatabaseDataguardAssociationDataSourceRepresentation) +
					compartmentIdVariableStr + AutonomousContainerDatabaseDataguardAssociationResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(datasourceName, "autonomous_database_id"),

					resource.TestCheckResourceAttrSet(datasourceName, "autonomous_database_dataguard_associations.#"),
					resource.TestCheckResourceAttrSet(datasourceName, "autonomous_database_dataguard_associations.0.autonomous_database_id"),
					resource.TestCheckResourceAttrSet(datasourceName, "autonomous_database_dataguard_associations.0.id"),
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
					generateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", Optional, Create, autonomousDatabaseDedicatedRepresentationForDataguard) +
					generateDataSourceFromRepresentationMap("oci_database_autonomous_database_dataguard_associations", "test_autonomous_database_dataguard_associations", Required, Create, autonomousDatabaseDataguardAssociationDataSourceRepresentation) +
					generateDataSourceFromRepresentationMap("oci_database_autonomous_database_dataguard_association", "test_autonomous_database_dataguard_association", Required, Create, autonomousDatabaseDataguardAssociationSingularDataSourceRepresentation) +
					compartmentIdVariableStr + AutonomousContainerDatabaseDataguardAssociationResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
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
		},
	})
}
