// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	DatabaseAutonomousContainerDatabaseSnapshotStandbyRequiredOnlyResource = DatabaseAdbdAutonomousContainerDatabaseSnapshotStandbyResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_container_database_snapshot_standby", "test_autonomous_container_database_snapshot_standby", acctest.Required, acctest.Create, DatabaseAutonomousContainerDatabaseSnapshotStandbyRepresentation)

	DatabaseAutonomousContainerDatabaseSnapshotStandbyRepresentation = map[string]interface{}{
		"autonomous_container_database_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_database_autonomous_container_database_add_standby.test_autonomous_container_database_add_standby.dataguard_group_members.1.autonomous_container_database_id}`},
		"role":                             acctest.Representation{RepType: acctest.Required, Create: `SNAPSHOT_STANDBY`},
		"connection_strings_type":          acctest.Representation{RepType: acctest.Optional, Create: `SNAPSHOT_SERVICES`},
	}

	DatabaseAdbdAutonomousContainerDatabaseSnapshotStandbyResourceDependencies = DatabaseAutonomousContainerDatabaseDataguardAssociationResourceConfig +
		acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_container_database_add_standby", "test_autonomous_container_database_add_standby", acctest.Optional, acctest.Create,
			acctest.RepresentationCopyWithRemovedProperties(DatabaseAdbdAutonomousContainerDatabaseAddStandbyRepresentation, []string{"is_automatic_failover_enabled", "fast_start_fail_over_lag_limit_in_seconds"}))
)

// issue-routing-tag: database/dbaas-atp-d
func TestDatabaseAutonomousContainerDatabaseSnapshotStandbyResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseAutonomousContainerDatabaseSnapshotStandbyResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_database_autonomous_container_database_snapshot_standby.test_autonomous_container_database_snapshot_standby"

	var resId string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+DatabaseAdbdAutonomousContainerDatabaseSnapshotStandbyResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_container_database_snapshot_standby", "test_autonomous_container_database_snapshot_standby", acctest.Optional, acctest.Create, DatabaseAutonomousContainerDatabaseSnapshotStandbyRepresentation), "database", "autonomousContainerDatabaseSnapshotStandby", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + DatabaseAdbdAutonomousContainerDatabaseSnapshotStandbyResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_container_database_snapshot_standby", "test_autonomous_container_database_snapshot_standby", acctest.Required, acctest.Create, DatabaseAutonomousContainerDatabaseSnapshotStandbyRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "autonomous_container_database_id"),
				resource.TestCheckResourceAttr(resourceName, "role", "SNAPSHOT_STANDBY"),
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + DatabaseAdbdAutonomousContainerDatabaseSnapshotStandbyResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + DatabaseAdbdAutonomousContainerDatabaseSnapshotStandbyResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_container_database_snapshot_standby", "test_autonomous_container_database_snapshot_standby", acctest.Optional, acctest.Create, DatabaseAutonomousContainerDatabaseSnapshotStandbyRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "autonomous_container_database_id"),
				resource.TestCheckResourceAttr(resourceName, "connection_strings_type", "SNAPSHOT_SERVICES"),
				resource.TestCheckResourceAttr(resourceName, "role", "SNAPSHOT_STANDBY"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					if isEnableExportCompartment, _ := strconv.ParseBool(utils.GetEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
						if errExport := resourcediscovery.TestExportCompartmentWithResourceName(&resId, &compartmentId, resourceName); errExport != nil {
							return errExport
						}
					}
					return err
				},
			),
		},
	})
}
