// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"strconv"

	"github.com/oracle/terraform-provider-oci/internal/acctest"

	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	DatabaseAutonomousVmClusterAcdResourceUsageDataSourceRepresentation = map[string]interface{}{
		"autonomous_vm_cluster_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_database_autonomous_vm_cluster.test_autonomous_vm_cluster.id}`},
		"compartment_id":           acctest.Representation{RepType: acctest.Optional, Create: `${var.compartment_id}`},
	}
)

func getExaccDatabaseAutonomousVmClusterAcdResourceUsageDataSourceRepresentation() map[string]interface{} {
	simulateDb, _ := strconv.ParseBool(utils.GetEnvSettingWithDefault("simulate_db", "false"))
	if simulateDb {
		return DatabaseAutonomousVmClusterAcdResourceUsageDataSourceRepresentation
	} else {
		// Add the dynamic properties
		return acctest.RepresentationCopyWithNewProperties(
			DatabaseAutonomousVmClusterAcdResourceUsageDataSourceRepresentation,
			map[string]interface{}{
				"autonomous_vm_cluster_id": acctest.Representation{RepType: acctest.Required, Create: `${var.autonomous_vm_cluster_id}`},
			})
	}
}
