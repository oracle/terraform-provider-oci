// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"

	"github.com/oracle/terraform-provider-oci/internal/utils"
)

/*

Example dbsystem representation

resource "oci_psql_db_system" "test_flexdb_system" {
  #Required
  db_version          = "15"
  display_name = "crr-tf-replicate_2_DO_NOT_DELETE"
  network_details {
    subnet_id = var.subnet_id
  }
  shape = "PostgreSQL.VM.Standard.E5.Flex"
  storage_details {
    is_regionally_durable = false
    availability_domain = "gXfg:AP-MUMBAI-1-AD-1"
    system_type = "OCI_OPTIMIZED_STORAGE"
    iops = "75000"
  }
  #apply_change_mode_to_stand_alone = "REPLAY_PENDING_UPDATES"
  source {
    source_type = "DB_SYSTEM"
    primary_db_system_id = var.primary_db_system_id
  }
  replication_config {
    is_rpo_enforced = true
    rpo_in_seconds = "300"
  }
  compartment_id      = var.compartment_id
  instance_count = "1"
  instance_ocpu_count = "2"
  instance_memory_size_in_gbs = "32"
  system_type = "OCI_OPTIMIZED_STORAGE"
  config_id = var.config_id
  management_policy {
    backup_policy {
      kind              = "NONE"
    }
    maintenance_window_start = "THU 17:00"
  }
  # Add freeform tags
  freeform_tags = {
    Environment = "Dev"
    Owner       = "Sivaram2"
  }
  timeouts {
    create = "60m"
    update = "60m"
  }
}
*/

var (
	PsqlDbSystemRepresentationReplica = map[string]interface{}{
		"compartment_id":                   acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"db_version":                       acctest.Representation{RepType: acctest.Required, Create: `15`},
		"display_name":                     acctest.Representation{RepType: acctest.Required, Create: `crr-test-terraform`, Update: `crr-test-terraform-2`},
		"network_details":                  acctest.RepresentationGroup{RepType: acctest.Required, Group: PsqlReplicaDbSystemIpNetworkDetailsRepresentation},
		"shape":                            acctest.Representation{RepType: acctest.Required, Create: `PostgreSQL.VM.Standard.E5.Flex`},
		"storage_details":                  acctest.RepresentationGroup{RepType: acctest.Required, Group: PsqlDbSystemRegionalStorageDetailsReplicaRepresentation},
		"instance_count":                   acctest.Representation{RepType: acctest.Required, Create: `1`},
		"instance_memory_size_in_gbs":      acctest.Representation{RepType: acctest.Optional, Create: `32`},
		"instance_ocpu_count":              acctest.Representation{RepType: acctest.Optional, Create: `2`},
		"management_policy":                acctest.RepresentationGroup{RepType: acctest.Optional, Group: PsqlDbSystemNoneManagementPolicyRepresentation},
		"config_id":                        acctest.Representation{RepType: acctest.Optional, Create: `${var.default_config_id}`},
		"source":                           acctest.RepresentationGroup{RepType: acctest.Optional, Group: PsqlDbSystemReplicaSourceRepresentation},
		"replication_config":               acctest.RepresentationGroup{RepType: acctest.Optional, Group: PsqlDbSystemReplicaReplicationConfigRepresentation},
		"apply_change_mode_to_stand_alone": acctest.Representation{RepType: acctest.Optional, Update: `REPLAY_PENDING_UPDATES`},
	}

	PsqlDbSystemRepresentationUpdateReplica = map[string]interface{}{
		"compartment_id":                   acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"db_version":                       acctest.Representation{RepType: acctest.Required, Create: `15`},
		"display_name":                     acctest.Representation{RepType: acctest.Required, Create: `crr-test-terraform`, Update: `crr-test-terraform-2`},
		"network_details":                  acctest.RepresentationGroup{RepType: acctest.Required, Group: PsqlReplicaDbSystemIpNetworkDetailsRepresentation},
		"shape":                            acctest.Representation{RepType: acctest.Required, Create: `PostgreSQL.VM.Standard.E5.Flex`},
		"storage_details":                  acctest.RepresentationGroup{RepType: acctest.Required, Group: PsqlDbSystemRegionalStorageDetailsReplicaRepresentation},
		"instance_count":                   acctest.Representation{RepType: acctest.Required, Create: `1`},
		"instance_memory_size_in_gbs":      acctest.Representation{RepType: acctest.Optional, Create: `32`},
		"instance_ocpu_count":              acctest.Representation{RepType: acctest.Optional, Create: `2`},
		"management_policy":                acctest.RepresentationGroup{RepType: acctest.Optional, Group: PsqlDbSystemNoneManagementPolicyRepresentation},
		"config_id":                        acctest.Representation{RepType: acctest.Optional, Create: `${var.default_config_id}`},
		"source":                           acctest.RepresentationGroup{RepType: acctest.Optional, Group: PsqlDbSystemUpdateReplicaSourceRepresentation},
		"replication_config":               acctest.RepresentationGroup{RepType: acctest.Optional, Group: PsqlDbSystemReplicaReplicationConfigRepresentation},
		"apply_change_mode_to_stand_alone": acctest.Representation{RepType: acctest.Optional, Update: `REPLAY_PENDING_UPDATES`},
	}

	PsqlReplicaDbSystemIpNetworkDetailsRepresentation = map[string]interface{}{
		"subnet_id": acctest.Representation{RepType: acctest.Required, Create: `${var.subnet_id}`},
	}

	PsqlDbSystemRegionalStorageDetailsReplicaRepresentation = map[string]interface{}{
		"is_regionally_durable": acctest.Representation{RepType: acctest.Required, Create: `false`},
		"system_type":           acctest.Representation{RepType: acctest.Required, Create: `OCI_OPTIMIZED_STORAGE`},
		"availability_domain":   acctest.Representation{RepType: acctest.Required, Create: `gXfg:AP-MUMBAI-1-AD-1`},
		"iops":                  acctest.Representation{RepType: acctest.Optional, Create: `75000`},
	}

	PsqlDbSystemNoneManagementPolicyRepresentation = map[string]interface{}{
		"backup_policy":            acctest.RepresentationGroup{RepType: acctest.Optional, Group: PsqlDbSystemNoneManagementPolicyBackupPolicyRepresentation},
		"maintenance_window_start": acctest.Representation{RepType: acctest.Optional, Create: `SUN 12:00`},
	}
	PsqlDbSystemNoneManagementPolicyBackupPolicyRepresentation = map[string]interface{}{
		"kind": acctest.Representation{RepType: acctest.Optional, Create: `NONE`},
	}

	PsqlDbSystemReplicaSourceRepresentation = map[string]interface{}{
		"source_type":          acctest.Representation{RepType: acctest.Optional, Create: `DB_SYSTEM`, Update: nil},
		"primary_db_system_id": acctest.Representation{RepType: acctest.Optional, Create: `${var.primary_db_system_id}`, Update: nil},
	}

	PsqlDbSystemUpdateReplicaSourceRepresentation = map[string]interface{}{}

	PsqlDbSystemReplicaReplicationConfigRepresentation = map[string]interface{}{
		"is_rpo_enforced": acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"rpo_in_seconds":  acctest.Representation{RepType: acctest.Optional, Create: `300`, Update: `300`},
	}

	PsqlDbSystemReplicaDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Optional, Create: `${var.compartment_id}`},
		"filter":         acctest.RepresentationGroup{RepType: acctest.Optional, Group: PsqlReplicaDbSystemDataSourceFilterRepresentation},
	}

	PsqlReplicaDbSystemDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_psql_db_system.test_replica_db_system.id}`}},
	}

	PsqlDbSystemReplicaResourceConfig = acctest.GenerateResourceFromRepresentationMap("oci_psql_db_system", "test_replica_db_system", acctest.Optional, acctest.Create, PsqlDbSystemRepresentationReplica)
)

// issue-routing-tag: psql/default
func TestPsqlDbSystemReplicaResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestPsqlDbSystemReplicaResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	defaultConfigId := utils.GetEnvSettingWithBlankDefault("default_config_id")
	defaultConfigIdVariableStr := fmt.Sprintf("variable \"default_config_id\" { default = \"%s\" }\n", defaultConfigId)

	subnetId := utils.GetEnvSettingWithBlankDefault("subnet_id")
	subnetIdVariableStr = fmt.Sprintf("variable \"subnet_id\" { default = \"%s\" }\n", subnetId)

	primaryDbSystemId := utils.GetEnvSettingWithBlankDefault("primary_db_system_id")
	primaryDbSystemIdVariableStr := fmt.Sprintf("variable \"primary_db_system_id\" { default = \"%s\" }\n", primaryDbSystemId)

	//datasourceName := "data.oci_psql_db_system_replicas.test_db_system_replicas"

	acctest.SaveConfigContent(config+compartmentIdVariableStr+defaultConfigIdVariableStr+subnetIdVariableStr+primaryDbSystemIdVariableStr+
		acctest.GenerateResourceFromRepresentationMap("oci_psql_db_system", "test_replica_db_system", acctest.Optional, acctest.Create, PsqlDbSystemRepresentationReplica), "psql", "dbSystem", t)

	acctest.ResourceTest(t, testAccCheckPsqlDbSystemDestroy, []resource.TestStep{
		// Replica test

		{
			Config: config + compartmentIdVariableStr + defaultConfigIdVariableStr + subnetIdVariableStr + primaryDbSystemIdVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_psql_db_system", "test_replica_db_system", acctest.Optional, acctest.Create, PsqlDbSystemRepresentationReplica),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr("oci_psql_db_system.test_replica_db_system", "compartment_id", compartmentId),
				resource.TestCheckResourceAttr("oci_psql_db_system.test_replica_db_system", "config_id", defaultConfigId),
				resource.TestCheckResourceAttr("oci_psql_db_system.test_replica_db_system", "network_details.0.subnet_id", subnetId),
				resource.TestCheckResourceAttr("oci_psql_db_system.test_replica_db_system", "source.0.primary_db_system_id", primaryDbSystemId),
				resource.TestCheckResourceAttr("oci_psql_db_system.test_replica_db_system", "source.0.source_type", "DB_SYSTEM"),
				resource.TestCheckResourceAttr("oci_psql_db_system.test_replica_db_system", "system_role", "WARM_STANDBY_DB_SYSTEM"),
				resource.TestCheckResourceAttr("oci_psql_db_system.test_replica_db_system", "replication_config.0.is_rpo_enforced", "false"),
				resource.TestCheckResourceAttr("oci_psql_db_system.test_replica_db_system", "replication_config.0.rpo_in_seconds", "300"),
			),
		},

		{
			Config: config + compartmentIdVariableStr + defaultConfigIdVariableStr + subnetIdVariableStr + primaryDbSystemIdVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_psql_db_system", "test_replica_db_system", acctest.Optional, acctest.Update, PsqlDbSystemRepresentationUpdateReplica),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr("oci_psql_db_system.test_replica_db_system", "compartment_id", compartmentId),
				resource.TestCheckResourceAttr("oci_psql_db_system.test_replica_db_system", "config_id", defaultConfigId),
				resource.TestCheckResourceAttr("oci_psql_db_system.test_replica_db_system", "network_details.0.subnet_id", subnetId),
				resource.TestCheckResourceAttr("oci_psql_db_system.test_replica_db_system", "source.0.primary_db_system_id", ""),
				resource.TestCheckResourceAttr("oci_psql_db_system.test_replica_db_system", "source.0.source_type", "NONE"),
				resource.TestCheckResourceAttr("oci_psql_db_system.test_replica_db_system", "system_role", "STANDALONE_DB_SYSTEM"),
				resource.TestCheckResourceAttr("oci_psql_db_system.test_replica_db_system", "replication_config.0.is_rpo_enforced", "true"),
				resource.TestCheckResourceAttr("oci_psql_db_system.test_replica_db_system", "replication_config.0.rpo_in_seconds", "300")),
		},

		{
			Config: config + compartmentIdVariableStr + defaultConfigIdVariableStr + subnetIdVariableStr + primaryDbSystemIdVariableStr,
		},

		//verify singular data source
		{
			Config: config + compartmentIdVariableStr + defaultConfigIdVariableStr + subnetIdVariableStr + primaryDbSystemIdVariableStr + PsqlDbSystemReplicaResourceConfig +
				acctest.GenerateDataSourceFromRepresentationMap("oci_psql_db_systems", "list_db_systems", acctest.Optional, acctest.Create, PsqlDbSystemReplicaDataSourceRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet("data.oci_psql_db_systems.list_db_systems", "compartment_id"),
				resource.TestCheckResourceAttrSet("data.oci_psql_db_systems.list_db_systems", "id"),
			),
		},
	})
}
