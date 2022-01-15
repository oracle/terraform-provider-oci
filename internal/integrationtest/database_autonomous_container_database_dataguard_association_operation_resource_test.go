// Copyright (c) 2020, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"testing"
	"time"

	"github.com/terraform-providers/terraform-provider-oci/internal/acctest"
	"github.com/terraform-providers/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	AutonomousContainerDatabaseDataguardAssociationOperationSwitchOverResourceConfig = acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_container_database_dataguard_association_operation", "switchover", acctest.Optional, acctest.Create,
		map[string]interface{}{
			"operation":                        acctest.Representation{RepType: acctest.Required, Create: `switchover`},
			"autonomous_container_database_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_database_autonomous_container_database.test_autonomous_container_database.id}`},
			"autonomous_container_database_dataguard_association_id": acctest.Representation{RepType: acctest.Required, Create: `${data.oci_database_autonomous_container_database_dataguard_associations.test.autonomous_container_database_dataguard_associations[0]["id"]}`},
		})

	ExaccAutonomousContainerDatabaseDataguardAssociationOperationSwitchOverResourceConfig = acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_container_database_dataguard_association_operation", "switchover", acctest.Optional, acctest.Create,
		map[string]interface{}{
			"operation":                        acctest.Representation{RepType: acctest.Required, Create: `switchover`},
			"autonomous_container_database_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_database_autonomous_container_database.exacc_test_autonomous_container_database.id}`},
			"autonomous_container_database_dataguard_association_id": acctest.Representation{RepType: acctest.Required, Create: `${data.oci_database_autonomous_container_database_dataguard_associations.test.autonomous_container_database_dataguard_associations[0]["id"]}`},
		})
	AutonomousContainerDatabaseDataguardAssociationOperationFailOverResourceConfig = acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_container_database_dataguard_association_operation", "failover", acctest.Optional, acctest.Create,
		map[string]interface{}{
			"operation":                        acctest.Representation{RepType: acctest.Required, Create: `failover`},
			"autonomous_container_database_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_database_autonomous_container_database.test_autonomous_container_database.id}`},
			"autonomous_container_database_dataguard_association_id": acctest.Representation{RepType: acctest.Required, Create: `${data.oci_database_autonomous_container_database_dataguard_associations.test.autonomous_container_database_dataguard_associations[0]["id"]}`},
		})

	ExaccAutonomousContainerDatabaseDataguardAssociationOperationFailOverResourceConfig = acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_container_database_dataguard_association_operation", "failover", acctest.Optional, acctest.Create,
		map[string]interface{}{
			"operation":                        acctest.Representation{RepType: acctest.Required, Create: `failover`},
			"autonomous_container_database_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_database_autonomous_container_database.exacc_test_autonomous_container_database.id}`},
			"autonomous_container_database_dataguard_association_id": acctest.Representation{RepType: acctest.Required, Create: `${data.oci_database_autonomous_container_database_dataguard_associations.test.autonomous_container_database_dataguard_associations[0]["id"]}`},
		})
	AutonomousContainerDatabaseDataguardAssociationOperationReinstateResourceConfig = acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_container_database_dataguard_association_operation", "reinstate", acctest.Optional, acctest.Create,
		map[string]interface{}{
			"operation":                        acctest.Representation{RepType: acctest.Required, Create: `reinstate`},
			"autonomous_container_database_id": acctest.Representation{RepType: acctest.Required, Create: `${data.oci_database_autonomous_container_database_dataguard_associations.test.autonomous_container_database_dataguard_associations[0]["peer_autonomous_container_database_id"]}`},
			"autonomous_container_database_dataguard_association_id": acctest.Representation{RepType: acctest.Required, Create: `${data.oci_database_autonomous_container_database_dataguard_associations.test.autonomous_container_database_dataguard_associations[0]["peer_autonomous_container_database_dataguard_association_id"]}`},
		})
)

// issue-routing-tag: database/dbaas-atp-d
func TestDatabaseAutonomousContainerDatabaseDataguardAssociationOperationResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseAutonomousContainerDatabaseDataguardAssociationOperationResource_basic")
	defer httpreplay.SaveScenario()

	provider := acctest.TestAccProvider
	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_database_autonomous_container_database_dataguard_associations.test"

	resource.Test(t, resource.TestCase{
		PreCheck: func() { acctest.PreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		Steps: []resource.TestStep{
			// verify Create
			{
				Config: config +
					acctest.GenerateDataSourceFromRepresentationMap("oci_database_autonomous_container_database_dataguard_associations", "test", acctest.Optional, acctest.Create, autonomousContainerDatabaseDataguardAssociationDataSourceRepresentation) +
					compartmentIdVariableStr + AutonomousContainerDatabaseDataguardAssociationResourceConfig,
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
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
					acctest.GenerateDataSourceFromRepresentationMap("oci_database_autonomous_container_database_dataguard_associations", "test", acctest.Optional, acctest.Create, autonomousContainerDatabaseDataguardAssociationDataSourceRepresentation) +
					compartmentIdVariableStr + AutonomousContainerDatabaseDataguardAssociationResourceConfig +
					AutonomousContainerDatabaseDataguardAssociationOperationSwitchOverResourceConfig,
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					func(s *terraform.State) (err error) {
						time.Sleep(5 * time.Minute)
						return nil
					},
				),
			},
			// verify switchover result
			{
				Config: config +
					acctest.GenerateDataSourceFromRepresentationMap("oci_database_autonomous_container_database_dataguard_associations", "test", acctest.Optional, acctest.Create, autonomousContainerDatabaseDataguardAssociationDataSourceRepresentation) +
					compartmentIdVariableStr + AutonomousContainerDatabaseDataguardAssociationResourceConfig,
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttrSet(datasourceName, "autonomous_container_database_id"),

					resource.TestCheckResourceAttrSet(datasourceName, "autonomous_container_database_dataguard_associations.#"),
					resource.TestCheckResourceAttr(datasourceName, "autonomous_container_database_dataguard_associations.0.peer_role", "PRIMARY"),
					resource.TestCheckResourceAttr(datasourceName, "autonomous_container_database_dataguard_associations.0.role", "STANDBY"),
				),
			},
			// failover
			{
				Config: config +
					acctest.GenerateDataSourceFromRepresentationMap("oci_database_autonomous_container_database_dataguard_associations", "test", acctest.Optional, acctest.Create, autonomousContainerDatabaseDataguardAssociationDataSourceRepresentation) +
					compartmentIdVariableStr + AutonomousContainerDatabaseDataguardAssociationResourceConfig +
					AutonomousContainerDatabaseDataguardAssociationOperationFailOverResourceConfig,
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					func(s *terraform.State) (err error) {
						time.Sleep(5 * time.Minute)
						return nil
					},
				),
			},
			// verify failover result
			{
				Config: config +
					acctest.GenerateDataSourceFromRepresentationMap("oci_database_autonomous_container_database_dataguard_associations", "test", acctest.Optional, acctest.Create, autonomousContainerDatabaseDataguardAssociationDataSourceRepresentation) +
					compartmentIdVariableStr + AutonomousContainerDatabaseDataguardAssociationResourceConfig,
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttrSet(datasourceName, "autonomous_container_database_id"),

					resource.TestCheckResourceAttrSet(datasourceName, "autonomous_container_database_dataguard_associations.#"),
					resource.TestCheckResourceAttr(datasourceName, "autonomous_container_database_dataguard_associations.0.peer_role", "DISABLED_STANDBY"),
					resource.TestCheckResourceAttr(datasourceName, "autonomous_container_database_dataguard_associations.0.role", "PRIMARY"),
				),
			},
			// reinstate
			{
				Config: config +
					acctest.GenerateDataSourceFromRepresentationMap("oci_database_autonomous_container_database_dataguard_associations", "test", acctest.Optional, acctest.Create, autonomousContainerDatabaseDataguardAssociationDataSourceRepresentation) +
					compartmentIdVariableStr + AutonomousContainerDatabaseDataguardAssociationResourceConfig +
					AutonomousContainerDatabaseDataguardAssociationOperationReinstateResourceConfig,
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					func(s *terraform.State) (err error) {
						time.Sleep(5 * time.Minute)
						return nil
					},
				),
			},
			// verify reinstate result
			{
				Config: config +
					acctest.GenerateDataSourceFromRepresentationMap("oci_database_autonomous_container_database_dataguard_associations", "test", acctest.Optional, acctest.Create, autonomousContainerDatabaseDataguardAssociationDataSourceRepresentation) +
					compartmentIdVariableStr + AutonomousContainerDatabaseDataguardAssociationResourceConfig,
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttrSet(datasourceName, "autonomous_container_database_id"),

					resource.TestCheckResourceAttrSet(datasourceName, "autonomous_container_database_dataguard_associations.#"),
					resource.TestCheckResourceAttr(datasourceName, "autonomous_container_database_dataguard_associations.0.peer_role", "STANDBY"),
					resource.TestCheckResourceAttr(datasourceName, "autonomous_container_database_dataguard_associations.0.role", "PRIMARY"),
				),
			},
		},
	})
}

// issue-routing-tag: database/dbaas-atp-d
func TestDatabaseExaccAutonomousContainerDatabaseDataguardAssociationOperationResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseExaccAutonomousContainerDatabaseDataguardAssociationOperationResource_basic")
	defer httpreplay.SaveScenario()

	provider := acctest.TestAccProvider
	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_database_autonomous_container_database_dataguard_associations.test"

	resource.Test(t, resource.TestCase{
		PreCheck: func() { acctest.PreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		Steps: []resource.TestStep{
			// verify Create
			{
				Config: config +
					acctest.GenerateDataSourceFromRepresentationMap("oci_database_autonomous_container_database_dataguard_associations", "test", acctest.Optional, acctest.Create, exaccAutonomousContainerDatabaseDataguardAssociationDataSourceRepresentation) +
					compartmentIdVariableStr + ExaccAutonomousContainerDatabaseDataguardAssociationResourceConfig,
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
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
					acctest.GenerateDataSourceFromRepresentationMap("oci_database_autonomous_container_database_dataguard_associations", "test", acctest.Optional, acctest.Create, exaccAutonomousContainerDatabaseDataguardAssociationDataSourceRepresentation) +
					compartmentIdVariableStr + ExaccAutonomousContainerDatabaseDataguardAssociationResourceConfig +
					ExaccAutonomousContainerDatabaseDataguardAssociationOperationSwitchOverResourceConfig,
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					func(s *terraform.State) (err error) {
						time.Sleep(10 * time.Minute)
						return nil
					},
				),
			},
			// verify switchover result
			{
				Config: config +
					acctest.GenerateDataSourceFromRepresentationMap("oci_database_autonomous_container_database_dataguard_associations", "test", acctest.Optional, acctest.Create, exaccAutonomousContainerDatabaseDataguardAssociationDataSourceRepresentation) +
					compartmentIdVariableStr + ExaccAutonomousContainerDatabaseDataguardAssociationResourceConfig,
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttrSet(datasourceName, "autonomous_container_database_id"),

					resource.TestCheckResourceAttrSet(datasourceName, "autonomous_container_database_dataguard_associations.#"),
					resource.TestCheckResourceAttr(datasourceName, "autonomous_container_database_dataguard_associations.0.peer_role", "PRIMARY"),
					resource.TestCheckResourceAttr(datasourceName, "autonomous_container_database_dataguard_associations.0.role", "STANDBY"),
				),
			},
			// failover
			{
				Config: config +
					acctest.GenerateDataSourceFromRepresentationMap("oci_database_autonomous_container_database_dataguard_associations", "test", acctest.Optional, acctest.Create, exaccAutonomousContainerDatabaseDataguardAssociationDataSourceRepresentation) +
					compartmentIdVariableStr + ExaccAutonomousContainerDatabaseDataguardAssociationResourceConfig +
					ExaccAutonomousContainerDatabaseDataguardAssociationOperationFailOverResourceConfig,
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					func(s *terraform.State) (err error) {
						time.Sleep(5 * time.Minute)
						return nil
					},
				),
			},
			// verify failover result
			{
				Config: config +
					acctest.GenerateDataSourceFromRepresentationMap("oci_database_autonomous_container_database_dataguard_associations", "test", acctest.Optional, acctest.Create, exaccAutonomousContainerDatabaseDataguardAssociationDataSourceRepresentation) +
					compartmentIdVariableStr + ExaccAutonomousContainerDatabaseDataguardAssociationResourceConfig,
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttrSet(datasourceName, "autonomous_container_database_id"),

					resource.TestCheckResourceAttrSet(datasourceName, "autonomous_container_database_dataguard_associations.#"),
					resource.TestCheckResourceAttr(datasourceName, "autonomous_container_database_dataguard_associations.0.peer_role", "DISABLED_STANDBY"),
					resource.TestCheckResourceAttr(datasourceName, "autonomous_container_database_dataguard_associations.0.role", "PRIMARY"),
				),
			},
			// reinstate
			{
				Config: config +
					acctest.GenerateDataSourceFromRepresentationMap("oci_database_autonomous_container_database_dataguard_associations", "test", acctest.Optional, acctest.Create, exaccAutonomousContainerDatabaseDataguardAssociationDataSourceRepresentation) +
					compartmentIdVariableStr + ExaccAutonomousContainerDatabaseDataguardAssociationResourceConfig +
					AutonomousContainerDatabaseDataguardAssociationOperationReinstateResourceConfig,
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					func(s *terraform.State) (err error) {
						time.Sleep(5 * time.Minute)
						return nil
					},
				),
			},
			// verify reinstate result
			{
				Config: config +
					acctest.GenerateDataSourceFromRepresentationMap("oci_database_autonomous_container_database_dataguard_associations", "test", acctest.Optional, acctest.Create, exaccAutonomousContainerDatabaseDataguardAssociationDataSourceRepresentation) +
					compartmentIdVariableStr + ExaccAutonomousContainerDatabaseDataguardAssociationResourceConfig,
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttrSet(datasourceName, "autonomous_container_database_id"),

					resource.TestCheckResourceAttrSet(datasourceName, "autonomous_container_database_dataguard_associations.#"),
					resource.TestCheckResourceAttr(datasourceName, "autonomous_container_database_dataguard_associations.0.peer_role", "STANDBY"),
					resource.TestCheckResourceAttr(datasourceName, "autonomous_container_database_dataguard_associations.0.role", "PRIMARY"),
				),
			},
		},
	})
}
