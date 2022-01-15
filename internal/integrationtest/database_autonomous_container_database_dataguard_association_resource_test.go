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
	atpdAutonomousContainerDatabaseDataguardAssociationResourceRepresentation = map[string]interface{}{
		"autonomous_container_database_dataguard_association_id": acctest.Representation{RepType: acctest.Required, Create: `${data.oci_database_autonomous_container_database_dataguard_associations.test_autonomous_container_database_dataguard_associations.autonomous_container_database_dataguard_associations[0]["id"]}`},
		"autonomous_container_database_id":                       acctest.Representation{RepType: acctest.Required, Create: `${oci_database_autonomous_container_database.test_autonomous_container_database.id}`},
		"is_automatic_failover_enabled":                          acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
	}

	exaccAutonomousContainerDatabaseDataguardAssociationResourceRepresentation = map[string]interface{}{
		"autonomous_container_database_dataguard_association_id": acctest.Representation{RepType: acctest.Required, Create: `${data.oci_database_autonomous_container_database_dataguard_associations.exacc_test_autonomous_container_database_dataguard_associations.autonomous_container_database_dataguard_associations[0]["id"]}`},
		"autonomous_container_database_id":                       acctest.Representation{RepType: acctest.Required, Create: `${oci_database_autonomous_container_database.exacc_test_autonomous_container_database.id}`},
		"is_automatic_failover_enabled":                          acctest.Representation{RepType: acctest.Optional, Create: `true`, Update: `false`},
	}
)

// issue-routing-tag: database/dbaas-atp-d
func TestDatabaseAutonomousContainerDatabaseDataguardAssociationResource_update(t *testing.T) {

	httpreplay.SetScenario("TestDatabaseAutonomousContainerDatabaseDataguardAssociationResource_update")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_database_autonomous_container_database_dataguard_associations.test_autonomous_container_database_dataguard_associations"
	singularDatasourceName := "data.oci_database_autonomous_container_database_dataguard_association.test_autonomous_container_database_dataguard_association"
	resourceName := "oci_database_autonomous_container_database_dataguard_association.test_update_autonomous_container_database_dataguard_association"
	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{

		//create datasource
		{
			Config: config +

				acctest.GenerateDataSourceFromRepresentationMap("oci_database_autonomous_container_database_dataguard_associations", "test_autonomous_container_database_dataguard_associations", acctest.Optional, acctest.Create, autonomousContainerDatabaseDataguardAssociationDataSourceRepresentation) +
				compartmentIdVariableStr + AutonomousContainerDatabaseDataguardAssociationResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "autonomous_container_database_id"),

				resource.TestCheckResourceAttrSet(datasourceName, "autonomous_container_database_dataguard_associations.#"),
				resource.TestCheckResourceAttrSet(datasourceName, "autonomous_container_database_dataguard_associations.0.autonomous_container_database_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "autonomous_container_database_dataguard_associations.0.id"),
				resource.TestCheckResourceAttr(datasourceName, "autonomous_container_database_dataguard_associations.0.is_automatic_failover_enabled", "false"),
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
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_autonomous_container_database_dataguard_associations", "test_autonomous_container_database_dataguard_associations", acctest.Optional, acctest.Create, autonomousContainerDatabaseDataguardAssociationDataSourceRepresentation) +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_autonomous_container_database_dataguard_association", "test_autonomous_container_database_dataguard_association", acctest.Optional, acctest.Create, atpdAutonomousContainerDatabaseDataguardAssociationSingularDataSourceRepresentation) +
				compartmentIdVariableStr + AutonomousContainerDatabaseDataguardAssociationResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "autonomous_container_database_dataguard_association_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "autonomous_container_database_id"),

				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "is_automatic_failover_enabled", "false"),
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

		// verify create with optional parameters
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_autonomous_container_database_dataguard_associations", "test_autonomous_container_database_dataguard_associations", acctest.Optional, acctest.Create, autonomousContainerDatabaseDataguardAssociationDataSourceRepresentation) +
				acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_container_database_dataguard_association", "test_update_autonomous_container_database_dataguard_association", acctest.Optional, acctest.Create, atpdAutonomousContainerDatabaseDataguardAssociationResourceRepresentation) +
				compartmentIdVariableStr + AutonomousContainerDatabaseDataguardAssociationResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "autonomous_container_database_dataguard_association_id"),
				resource.TestCheckResourceAttrSet(resourceName, "autonomous_container_database_id"),

				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_automatic_failover_enabled", "false"),
				resource.TestCheckResourceAttrSet(resourceName, "peer_autonomous_container_database_dataguard_association_id"),
				resource.TestCheckResourceAttrSet(resourceName, "peer_autonomous_container_database_id"),
				resource.TestCheckResourceAttrSet(resourceName, "peer_lifecycle_state"),
				resource.TestCheckResourceAttrSet(resourceName, "peer_role"),
				resource.TestCheckResourceAttrSet(resourceName, "protection_mode"),
				resource.TestCheckResourceAttrSet(resourceName, "role"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
			),
		},

		// verify updates with optional parameters
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_autonomous_container_database_dataguard_associations", "test_autonomous_container_database_dataguard_associations", acctest.Optional, acctest.Create, autonomousContainerDatabaseDataguardAssociationDataSourceRepresentation) +
				acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_container_database_dataguard_association", "test_update_autonomous_container_database_dataguard_association", acctest.Optional, acctest.Update, atpdAutonomousContainerDatabaseDataguardAssociationResourceRepresentation) +
				compartmentIdVariableStr + AutonomousContainerDatabaseDataguardAssociationResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "autonomous_container_database_dataguard_association_id"),
				resource.TestCheckResourceAttrSet(resourceName, "autonomous_container_database_id"),

				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_automatic_failover_enabled", "true"),
				resource.TestCheckResourceAttrSet(resourceName, "peer_autonomous_container_database_dataguard_association_id"),
				resource.TestCheckResourceAttrSet(resourceName, "peer_autonomous_container_database_id"),
				resource.TestCheckResourceAttrSet(resourceName, "peer_lifecycle_state"),
				resource.TestCheckResourceAttrSet(resourceName, "peer_role"),
				resource.TestCheckResourceAttrSet(resourceName, "protection_mode"),
				resource.TestCheckResourceAttrSet(resourceName, "role"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
			),
		},
	})
}

func TestDatabaseExaccAutonomousContainerDatabaseDataguardAssociationResource_update(t *testing.T) {

	httpreplay.SetScenario("TestDatabaseExaccAutonomousContainerDatabaseDataguardAssociationResource_update")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_database_autonomous_container_database_dataguard_associations.exacc_test_autonomous_container_database_dataguard_associations"
	singularDatasourceName := "data.oci_database_autonomous_container_database_dataguard_association.exacc_test_autonomous_container_database_dataguard_association"
	resourceName := "oci_database_autonomous_container_database_dataguard_association.exacc_test_autonomous_container_database_dataguard_association"
	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_autonomous_container_database_dataguard_associations", "exacc_test_autonomous_container_database_dataguard_associations", acctest.Optional, acctest.Create, exaccAutonomousContainerDatabaseDataguardAssociationDataSourceRepresentation) +
				compartmentIdVariableStr + ExaccACDFSFOResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "autonomous_container_database_id"),

				resource.TestCheckResourceAttrSet(datasourceName, "autonomous_container_database_dataguard_associations.#"),
				resource.TestCheckResourceAttrSet(datasourceName, "autonomous_container_database_dataguard_associations.0.autonomous_container_database_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "autonomous_container_database_dataguard_associations.0.id"),
				resource.TestCheckResourceAttr(datasourceName, "autonomous_container_database_dataguard_associations.0.is_automatic_failover_enabled", "true"),
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
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_autonomous_container_database_dataguard_associations", "exacc_test_autonomous_container_database_dataguard_associations", acctest.Optional, acctest.Create, exaccAutonomousContainerDatabaseDataguardAssociationDataSourceRepresentation) +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_autonomous_container_database_dataguard_association", "exacc_test_autonomous_container_database_dataguard_association", acctest.Optional, acctest.Create, exaccAutonomousContainerDatabaseDataguardAssociationSingularDataSourceRepresentation) +
				compartmentIdVariableStr + ExaccACDFSFOResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "autonomous_container_database_dataguard_association_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "autonomous_container_database_id"),
				resource.TestCheckResourceAttr(datasourceName, "autonomous_container_database_dataguard_associations.0.is_automatic_failover_enabled", "true"),
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
		// verify create with optional parameters
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_autonomous_container_database_dataguard_associations", "exacc_test_autonomous_container_database_dataguard_associations", acctest.Optional, acctest.Create, exaccAutonomousContainerDatabaseDataguardAssociationDataSourceRepresentation) +
				acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_container_database_dataguard_association", "exacc_test_autonomous_container_database_dataguard_association", acctest.Optional, acctest.Create, exaccAutonomousContainerDatabaseDataguardAssociationResourceRepresentation) +
				compartmentIdVariableStr + ExaccACDFSFOResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "autonomous_container_database_dataguard_association_id"),
				resource.TestCheckResourceAttrSet(resourceName, "autonomous_container_database_id"),

				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_automatic_failover_enabled", "true"),
				resource.TestCheckResourceAttrSet(resourceName, "peer_autonomous_container_database_dataguard_association_id"),
				resource.TestCheckResourceAttrSet(resourceName, "peer_autonomous_container_database_id"),
				resource.TestCheckResourceAttrSet(resourceName, "peer_lifecycle_state"),
				resource.TestCheckResourceAttrSet(resourceName, "peer_role"),
				resource.TestCheckResourceAttrSet(resourceName, "protection_mode"),
				resource.TestCheckResourceAttrSet(resourceName, "role"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
			),
		},
		//verify update
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_autonomous_container_database_dataguard_associations", "exacc_test_autonomous_container_database_dataguard_associations", acctest.Optional, acctest.Create, exaccAutonomousContainerDatabaseDataguardAssociationDataSourceRepresentation) +
				acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_container_database_dataguard_association", "exacc_test_autonomous_container_database_dataguard_association", acctest.Optional, acctest.Update, exaccAutonomousContainerDatabaseDataguardAssociationResourceRepresentation) +
				compartmentIdVariableStr + ExaccACDFSFOResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "autonomous_container_database_dataguard_association_id"),
				resource.TestCheckResourceAttrSet(resourceName, "autonomous_container_database_id"),

				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_automatic_failover_enabled", "false"),
				resource.TestCheckResourceAttrSet(resourceName, "peer_autonomous_container_database_dataguard_association_id"),
				resource.TestCheckResourceAttrSet(resourceName, "peer_autonomous_container_database_id"),
				resource.TestCheckResourceAttrSet(resourceName, "peer_lifecycle_state"),
				resource.TestCheckResourceAttrSet(resourceName, "peer_role"),
				resource.TestCheckResourceAttrSet(resourceName, "protection_mode"),
				resource.TestCheckResourceAttrSet(resourceName, "role"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
			),
		},
	})
}
