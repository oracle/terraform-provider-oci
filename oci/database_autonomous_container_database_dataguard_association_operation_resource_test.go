// Copyright (c) 2020, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"fmt"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	AutonomousContainerDatabaseDataguardAssociationOperationSwitchOverResourceConfig = generateResourceFromRepresentationMap("oci_database_autonomous_container_database_dataguard_association_operation", "switchover", Optional, Create,
		map[string]interface{}{
			"operation":                        Representation{repType: Required, create: `switchover`},
			"autonomous_container_database_id": Representation{repType: Required, create: `${oci_database_autonomous_container_database.test_autonomous_container_database.id}`},
			"autonomous_container_database_dataguard_association_id": Representation{repType: Required, create: `${data.oci_database_autonomous_container_database_dataguard_associations.test.autonomous_container_database_dataguard_associations[0]["id"]}`},
		})

	ExaccAutonomousContainerDatabaseDataguardAssociationOperationSwitchOverResourceConfig = generateResourceFromRepresentationMap("oci_database_autonomous_container_database_dataguard_association_operation", "switchover", Optional, Create,
		map[string]interface{}{
			"operation":                        Representation{repType: Required, create: `switchover`},
			"autonomous_container_database_id": Representation{repType: Required, create: `${oci_database_autonomous_container_database.exacc_test_autonomous_container_database.id}`},
			"autonomous_container_database_dataguard_association_id": Representation{repType: Required, create: `${data.oci_database_autonomous_container_database_dataguard_associations.test.autonomous_container_database_dataguard_associations[0]["id"]}`},
		})
	AutonomousContainerDatabaseDataguardAssociationOperationFailOverResourceConfig = generateResourceFromRepresentationMap("oci_database_autonomous_container_database_dataguard_association_operation", "failover", Optional, Create,
		map[string]interface{}{
			"operation":                        Representation{repType: Required, create: `failover`},
			"autonomous_container_database_id": Representation{repType: Required, create: `${oci_database_autonomous_container_database.test_autonomous_container_database.id}`},
			"autonomous_container_database_dataguard_association_id": Representation{repType: Required, create: `${data.oci_database_autonomous_container_database_dataguard_associations.test.autonomous_container_database_dataguard_associations[0]["id"]}`},
		})

	ExaccAutonomousContainerDatabaseDataguardAssociationOperationFailOverResourceConfig = generateResourceFromRepresentationMap("oci_database_autonomous_container_database_dataguard_association_operation", "failover", Optional, Create,
		map[string]interface{}{
			"operation":                        Representation{repType: Required, create: `failover`},
			"autonomous_container_database_id": Representation{repType: Required, create: `${oci_database_autonomous_container_database.exacc_test_autonomous_container_database.id}`},
			"autonomous_container_database_dataguard_association_id": Representation{repType: Required, create: `${data.oci_database_autonomous_container_database_dataguard_associations.test.autonomous_container_database_dataguard_associations[0]["id"]}`},
		})
	AutonomousContainerDatabaseDataguardAssociationOperationReinstateResourceConfig = generateResourceFromRepresentationMap("oci_database_autonomous_container_database_dataguard_association_operation", "reinstate", Optional, Create,
		map[string]interface{}{
			"operation":                        Representation{repType: Required, create: `reinstate`},
			"autonomous_container_database_id": Representation{repType: Required, create: `${data.oci_database_autonomous_container_database_dataguard_associations.test.autonomous_container_database_dataguard_associations[0]["peer_autonomous_container_database_id"]}`},
			"autonomous_container_database_dataguard_association_id": Representation{repType: Required, create: `${data.oci_database_autonomous_container_database_dataguard_associations.test.autonomous_container_database_dataguard_associations[0]["peer_autonomous_container_database_dataguard_association_id"]}`},
		})
)

func TestDatabaseAutonomousContainerDatabaseDataguardAssociationOperationResource_basic(t *testing.T) {
	t.Skip("Skipping Test for TeamCity")
	httpreplay.SetScenario("TestDatabaseAutonomousContainerDatabaseDataguardAssociationOperationResource_basic")
	defer httpreplay.SaveScenario()

	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_database_autonomous_container_database_dataguard_associations.test"

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		Steps: []resource.TestStep{
			// verify create
			{
				Config: config +
					generateDataSourceFromRepresentationMap("oci_database_autonomous_container_database_dataguard_associations", "test", Optional, Create, autonomousContainerDatabaseDataguardAssociationDataSourceRepresentation) +
					compartmentIdVariableStr + AutonomousContainerDatabaseDataguardAssociationResourceConfig,
				Check: ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttrSet(datasourceName, "autonomous_container_database_id"),
					resource.TestCheckResourceAttrSet(datasourceName, "autonomous_container_database_dataguard_associations.#"),
					resource.TestCheckResourceAttrSet(datasourceName, "autonomous_container_database_dataguard_associations.0.autonomous_container_database_id"),
					resource.TestCheckResourceAttrSet(datasourceName, "autonomous_container_database_dataguard_associations.0.id"),
					resource.TestCheckResourceAttrSet(datasourceName, "autonomous_container_database_dataguard_associations.0.peer_autonomous_container_database_dataguard_association_id"),
					resource.TestCheckResourceAttrSet(datasourceName, "autonomous_container_database_dataguard_associations.0.peer_autonomous_container_database_id"),
					resource.TestCheckResourceAttrSet(datasourceName, "autonomous_container_database_dataguard_associations.0.peer_lifecycle_state"),
					resource.TestCheckResourceAttr(datasourceName, "autonomous_container_database_dataguard_associations.0.peer_role", "STANDBY"),
					resource.TestCheckResourceAttrSet(datasourceName, "autonomous_container_database_dataguard_associations.0.protection_mode"),
					resource.TestCheckResourceAttr(datasourceName, "autonomous_container_database_dataguard_associations.0.role", "PRIMARY"),
					resource.TestCheckResourceAttrSet(datasourceName, "autonomous_container_database_dataguard_associations.0.state"),
					resource.TestCheckResourceAttrSet(datasourceName, "autonomous_container_database_dataguard_associations.0.time_created"),
				),
			},
			// switchover
			{
				Config: config +
					generateDataSourceFromRepresentationMap("oci_database_autonomous_container_database_dataguard_associations", "test", Optional, Create, autonomousContainerDatabaseDataguardAssociationDataSourceRepresentation) +
					compartmentIdVariableStr + AutonomousContainerDatabaseDataguardAssociationResourceConfig +
					AutonomousContainerDatabaseDataguardAssociationOperationSwitchOverResourceConfig,
				Check: ComposeAggregateTestCheckFuncWrapper(
					func(s *terraform.State) (err error) {
						time.Sleep(5 * time.Minute)
						return nil
					},
				),
			},
			// verify switchover result
			{
				Config: config +
					generateDataSourceFromRepresentationMap("oci_database_autonomous_container_database_dataguard_associations", "test", Optional, Create, autonomousContainerDatabaseDataguardAssociationDataSourceRepresentation) +
					compartmentIdVariableStr + AutonomousContainerDatabaseDataguardAssociationResourceConfig,
				Check: ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttrSet(datasourceName, "autonomous_container_database_id"),

					resource.TestCheckResourceAttrSet(datasourceName, "autonomous_container_database_dataguard_associations.#"),
					resource.TestCheckResourceAttr(datasourceName, "autonomous_container_database_dataguard_associations.0.peer_role", "PRIMARY"),
					resource.TestCheckResourceAttr(datasourceName, "autonomous_container_database_dataguard_associations.0.role", "STANDBY"),
				),
			},
			// failover
			{
				Config: config +
					generateDataSourceFromRepresentationMap("oci_database_autonomous_container_database_dataguard_associations", "test", Optional, Create, autonomousContainerDatabaseDataguardAssociationDataSourceRepresentation) +
					compartmentIdVariableStr + AutonomousContainerDatabaseDataguardAssociationResourceConfig +
					AutonomousContainerDatabaseDataguardAssociationOperationFailOverResourceConfig,
				Check: ComposeAggregateTestCheckFuncWrapper(
					func(s *terraform.State) (err error) {
						time.Sleep(5 * time.Minute)
						return nil
					},
				),
			},
			// verify failover result
			{
				Config: config +
					generateDataSourceFromRepresentationMap("oci_database_autonomous_container_database_dataguard_associations", "test", Optional, Create, autonomousContainerDatabaseDataguardAssociationDataSourceRepresentation) +
					compartmentIdVariableStr + AutonomousContainerDatabaseDataguardAssociationResourceConfig,
				Check: ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttrSet(datasourceName, "autonomous_container_database_id"),

					resource.TestCheckResourceAttrSet(datasourceName, "autonomous_container_database_dataguard_associations.#"),
					resource.TestCheckResourceAttr(datasourceName, "autonomous_container_database_dataguard_associations.0.peer_role", "DISABLED_STANDBY"),
					resource.TestCheckResourceAttr(datasourceName, "autonomous_container_database_dataguard_associations.0.role", "PRIMARY"),
				),
			},
			// reinstate
			{
				Config: config +
					generateDataSourceFromRepresentationMap("oci_database_autonomous_container_database_dataguard_associations", "test", Optional, Create, autonomousContainerDatabaseDataguardAssociationDataSourceRepresentation) +
					compartmentIdVariableStr + AutonomousContainerDatabaseDataguardAssociationResourceConfig +
					AutonomousContainerDatabaseDataguardAssociationOperationReinstateResourceConfig,
				Check: ComposeAggregateTestCheckFuncWrapper(
					func(s *terraform.State) (err error) {
						time.Sleep(5 * time.Minute)
						return nil
					},
				),
			},
			// verify reinstate result
			{
				Config: config +
					generateDataSourceFromRepresentationMap("oci_database_autonomous_container_database_dataguard_associations", "test", Optional, Create, autonomousContainerDatabaseDataguardAssociationDataSourceRepresentation) +
					compartmentIdVariableStr + AutonomousContainerDatabaseDataguardAssociationResourceConfig,
				Check: ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttrSet(datasourceName, "autonomous_container_database_id"),

					resource.TestCheckResourceAttrSet(datasourceName, "autonomous_container_database_dataguard_associations.#"),
					resource.TestCheckResourceAttr(datasourceName, "autonomous_container_database_dataguard_associations.0.peer_role", "STANDBY"),
					resource.TestCheckResourceAttr(datasourceName, "autonomous_container_database_dataguard_associations.0.role", "PRIMARY"),
				),
			},
		},
	})
}

func TestDatabaseExaccAutonomousContainerDatabaseDataguardAssociationOperationResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseExaccAutonomousContainerDatabaseDataguardAssociationOperationResource_basic")
	defer httpreplay.SaveScenario()

	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_database_autonomous_container_database_dataguard_associations.test"

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		Steps: []resource.TestStep{
			// verify create
			{
				Config: config +
					generateDataSourceFromRepresentationMap("oci_database_autonomous_container_database_dataguard_associations", "test", Optional, Create, exaccAutonomousContainerDatabaseDataguardAssociationDataSourceRepresentation) +
					compartmentIdVariableStr + ExaccAutonomousContainerDatabaseDataguardAssociationResourceConfig,
				Check: ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttrSet(datasourceName, "autonomous_container_database_id"),
					resource.TestCheckResourceAttrSet(datasourceName, "autonomous_container_database_dataguard_associations.#"),
					resource.TestCheckResourceAttrSet(datasourceName, "autonomous_container_database_dataguard_associations.0.autonomous_container_database_id"),
					resource.TestCheckResourceAttrSet(datasourceName, "autonomous_container_database_dataguard_associations.0.id"),
					resource.TestCheckResourceAttrSet(datasourceName, "autonomous_container_database_dataguard_associations.0.peer_autonomous_container_database_dataguard_association_id"),
					resource.TestCheckResourceAttrSet(datasourceName, "autonomous_container_database_dataguard_associations.0.peer_autonomous_container_database_id"),
					resource.TestCheckResourceAttrSet(datasourceName, "autonomous_container_database_dataguard_associations.0.peer_lifecycle_state"),
					resource.TestCheckResourceAttr(datasourceName, "autonomous_container_database_dataguard_associations.0.peer_role", "STANDBY"),
					resource.TestCheckResourceAttrSet(datasourceName, "autonomous_container_database_dataguard_associations.0.protection_mode"),
					resource.TestCheckResourceAttr(datasourceName, "autonomous_container_database_dataguard_associations.0.role", "PRIMARY"),
					resource.TestCheckResourceAttrSet(datasourceName, "autonomous_container_database_dataguard_associations.0.state"),
					resource.TestCheckResourceAttrSet(datasourceName, "autonomous_container_database_dataguard_associations.0.time_created"),
				),
			},
			// switchover
			{
				Config: config +
					generateDataSourceFromRepresentationMap("oci_database_autonomous_container_database_dataguard_associations", "test", Optional, Create, exaccAutonomousContainerDatabaseDataguardAssociationDataSourceRepresentation) +
					compartmentIdVariableStr + ExaccAutonomousContainerDatabaseDataguardAssociationResourceConfig +
					ExaccAutonomousContainerDatabaseDataguardAssociationOperationSwitchOverResourceConfig,
				Check: ComposeAggregateTestCheckFuncWrapper(
					func(s *terraform.State) (err error) {
						time.Sleep(10 * time.Minute)
						return nil
					},
				),
			},
			// verify switchover result
			{
				Config: config +
					generateDataSourceFromRepresentationMap("oci_database_autonomous_container_database_dataguard_associations", "test", Optional, Create, exaccAutonomousContainerDatabaseDataguardAssociationDataSourceRepresentation) +
					compartmentIdVariableStr + ExaccAutonomousContainerDatabaseDataguardAssociationResourceConfig,
				Check: ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttrSet(datasourceName, "autonomous_container_database_id"),

					resource.TestCheckResourceAttrSet(datasourceName, "autonomous_container_database_dataguard_associations.#"),
					resource.TestCheckResourceAttr(datasourceName, "autonomous_container_database_dataguard_associations.0.peer_role", "PRIMARY"),
					resource.TestCheckResourceAttr(datasourceName, "autonomous_container_database_dataguard_associations.0.role", "STANDBY"),
				),
			},
			// failover
			{
				Config: config +
					generateDataSourceFromRepresentationMap("oci_database_autonomous_container_database_dataguard_associations", "test", Optional, Create, exaccAutonomousContainerDatabaseDataguardAssociationDataSourceRepresentation) +
					compartmentIdVariableStr + ExaccAutonomousContainerDatabaseDataguardAssociationResourceConfig +
					ExaccAutonomousContainerDatabaseDataguardAssociationOperationFailOverResourceConfig,
				Check: ComposeAggregateTestCheckFuncWrapper(
					func(s *terraform.State) (err error) {
						time.Sleep(5 * time.Minute)
						return nil
					},
				),
			},
			// verify failover result
			{
				Config: config +
					generateDataSourceFromRepresentationMap("oci_database_autonomous_container_database_dataguard_associations", "test", Optional, Create, exaccAutonomousContainerDatabaseDataguardAssociationDataSourceRepresentation) +
					compartmentIdVariableStr + ExaccAutonomousContainerDatabaseDataguardAssociationResourceConfig,
				Check: ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttrSet(datasourceName, "autonomous_container_database_id"),

					resource.TestCheckResourceAttrSet(datasourceName, "autonomous_container_database_dataguard_associations.#"),
					resource.TestCheckResourceAttr(datasourceName, "autonomous_container_database_dataguard_associations.0.peer_role", "DISABLED_STANDBY"),
					resource.TestCheckResourceAttr(datasourceName, "autonomous_container_database_dataguard_associations.0.role", "PRIMARY"),
				),
			},
			// reinstate
			{
				Config: config +
					generateDataSourceFromRepresentationMap("oci_database_autonomous_container_database_dataguard_associations", "test", Optional, Create, exaccAutonomousContainerDatabaseDataguardAssociationDataSourceRepresentation) +
					compartmentIdVariableStr + ExaccAutonomousContainerDatabaseDataguardAssociationResourceConfig +
					AutonomousContainerDatabaseDataguardAssociationOperationReinstateResourceConfig,
				Check: ComposeAggregateTestCheckFuncWrapper(
					func(s *terraform.State) (err error) {
						time.Sleep(5 * time.Minute)
						return nil
					},
				),
			},
			// verify reinstate result
			{
				Config: config +
					generateDataSourceFromRepresentationMap("oci_database_autonomous_container_database_dataguard_associations", "test", Optional, Create, exaccAutonomousContainerDatabaseDataguardAssociationDataSourceRepresentation) +
					compartmentIdVariableStr + ExaccAutonomousContainerDatabaseDataguardAssociationResourceConfig,
				Check: ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttrSet(datasourceName, "autonomous_container_database_id"),

					resource.TestCheckResourceAttrSet(datasourceName, "autonomous_container_database_dataguard_associations.#"),
					resource.TestCheckResourceAttr(datasourceName, "autonomous_container_database_dataguard_associations.0.peer_role", "STANDBY"),
					resource.TestCheckResourceAttr(datasourceName, "autonomous_container_database_dataguard_associations.0.role", "PRIMARY"),
				),
			},
		},
	})
}
