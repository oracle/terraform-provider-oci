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
	autonomousContainerDatabaseDataguardAssociationSingularDataSourceRepresentation = map[string]interface{}{
		"autonomous_container_database_dataguard_association_id": Representation{repType: Required, create: `${data.oci_database_autonomous_container_database_dataguard_associations.test_autonomous_container_database_dataguard_associations.autonomous_container_database_dataguard_associations[0]["id"]}`},
		"autonomous_container_database_id":                       Representation{repType: Required, create: `${oci_database_autonomous_container_database.test_autonomous_container_database.id}`},
	}

	autonomousContainerDatabaseDataguardAssociationDataSourceRepresentation = map[string]interface{}{
		"autonomous_container_database_id": Representation{repType: Required, create: `${oci_database_autonomous_container_database.test_autonomous_container_database.id}`},
	}

	AutonomousContainerDatabaseDataguardAssociationResourceConfig = generateResourceFromRepresentationMap("oci_database_autonomous_container_database", "test_autonomous_container_database", Optional, Create,
		representationCopyWithNewProperties(representationCopyWithRemovedProperties(autonomousContainerDatabaseRepresentation, []string{"vault_id", "kms_key_id"}), map[string]interface{}{
			"service_level_agreement_type":              Representation{repType: Optional, create: `AUTONOMOUS_DATAGUARD`},
			"protection_mode":                           Representation{repType: Optional, create: `MAXIMUM_AVAILABILITY`},
			"peer_autonomous_exadata_infrastructure_id": Representation{repType: Optional, create: `${oci_database_autonomous_exadata_infrastructure.peer_autonomous_exadata_infrastructure.id}`},
		})) +
		AutonomousExadataInfrastructureResourceConfig +
		generateResourceFromRepresentationMap("oci_database_autonomous_exadata_infrastructure", "peer_autonomous_exadata_infrastructure", Optional, Create, autonomousExadataInfrastructureRepresentation)
)

func TestDatabaseAutonomousContainerDatabaseDataguardAssociationResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseAutonomousContainerDatabaseDataguardAssociationResource_basic")
	defer httpreplay.SaveScenario()

	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_database_autonomous_container_database_dataguard_associations.test_autonomous_container_database_dataguard_associations"
	singularDatasourceName := "data.oci_database_autonomous_container_database_dataguard_association.test_autonomous_container_database_dataguard_association"

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		Steps: []resource.TestStep{
			// verify datasource
			{
				Config: config +
					generateDataSourceFromRepresentationMap("oci_database_autonomous_container_database_dataguard_associations", "test_autonomous_container_database_dataguard_associations", Optional, Create, autonomousContainerDatabaseDataguardAssociationDataSourceRepresentation) +
					compartmentIdVariableStr + AutonomousContainerDatabaseDataguardAssociationResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(datasourceName, "autonomous_container_database_id"),

					resource.TestCheckResourceAttrSet(datasourceName, "autonomous_container_database_dataguard_associations.#"),
					resource.TestCheckResourceAttrSet(datasourceName, "autonomous_container_database_dataguard_associations.0.autonomous_container_database_id"),
					resource.TestCheckResourceAttrSet(datasourceName, "autonomous_container_database_dataguard_associations.0.id"),
					resource.TestCheckResourceAttrSet(datasourceName, "autonomous_container_database_dataguard_associations.0.peer_autonomous_container_database_dataguard_association_id"),
					resource.TestCheckResourceAttrSet(datasourceName, "autonomous_container_database_dataguard_associations.0.peer_autonomous_container_database_id"),
					resource.TestCheckResourceAttrSet(datasourceName, "autonomous_container_database_dataguard_associations.0.peer_lifecycle_state"),
					resource.TestCheckResourceAttrSet(datasourceName, "autonomous_container_database_dataguard_associations.0.peer_role"),
					resource.TestCheckResourceAttrSet(datasourceName, "autonomous_container_database_dataguard_associations.0.protection_mode"),
					resource.TestCheckResourceAttrSet(datasourceName, "autonomous_container_database_dataguard_associations.0.role"),
					resource.TestCheckResourceAttrSet(datasourceName, "autonomous_container_database_dataguard_associations.0.state"),
					resource.TestCheckResourceAttrSet(datasourceName, "autonomous_container_database_dataguard_associations.0.time_created"),
				),
			},
			// verify singular datasource
			{
				Config: config +
					generateDataSourceFromRepresentationMap("oci_database_autonomous_container_database_dataguard_associations", "test_autonomous_container_database_dataguard_associations", Optional, Create, autonomousContainerDatabaseDataguardAssociationDataSourceRepresentation) +
					generateDataSourceFromRepresentationMap("oci_database_autonomous_container_database_dataguard_association", "test_autonomous_container_database_dataguard_association", Optional, Create, autonomousContainerDatabaseDataguardAssociationSingularDataSourceRepresentation) +
					compartmentIdVariableStr + AutonomousContainerDatabaseDataguardAssociationResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(singularDatasourceName, "autonomous_container_database_dataguard_association_id"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "autonomous_container_database_id"),

					resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "peer_autonomous_container_database_dataguard_association_id"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "peer_autonomous_container_database_id"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "peer_lifecycle_state"),
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
