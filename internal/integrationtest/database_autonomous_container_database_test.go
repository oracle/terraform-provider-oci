// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/oracle/terraform-provider-oci/internal/acctest"
	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/oracle/oci-go-sdk/v65/common"
	oci_database "github.com/oracle/oci-go-sdk/v65/database"

	"github.com/oracle/terraform-provider-oci/httpreplay"
)

var (
	DatabaseAutonomousContainerDatabaseRequiredOnlyResource = DatabaseAutonomousContainerDatabaseResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_container_database", "test_autonomous_container_database", acctest.Required, acctest.Create, DatabaseAutonomousContainerDatabaseRepresentation)

	ExaccAcdResourceConfig = acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_container_database", "test_autonomous_container_database", acctest.Optional, acctest.Update, ACDatabaseRepresentation)

	DatabaseAutonomousContainerDatabaseResourceConfig = ATPDAutonomousContainerDatabaseResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_container_database", "test_autonomous_container_database", acctest.Optional, acctest.Update, DatabaseAutonomousContainerDatabaseRepresentation)

	DatabaseDatabaseAutonomousContainerDatabaseSingularDataSourceRepresentation = map[string]interface{}{
		"autonomous_container_database_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_database_autonomous_container_database.test_autonomous_container_database.id}`},
	}

	DatabaseDatabaseAutonomousContainerDatabaseDataSourceRepresentation = map[string]interface{}{
		"compartment_id":                 acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"cloud_autonomous_vm_cluster_id": acctest.Representation{RepType: acctest.Optional, Create: `${oci_database_cloud_autonomous_vm_cluster.test_cloud_autonomous_vm_cluster.id}`},
		"availability_domain":            acctest.Representation{RepType: acctest.Optional, Create: `${data.oci_identity_availability_domain.ad.name}`},
		"display_name":                   acctest.Representation{RepType: acctest.Optional, Create: `containerDatabase2`, Update: `displayName2`},
		"state":                          acctest.Representation{RepType: acctest.Optional, Create: `AVAILABLE`},
		"filter":                         acctest.RepresentationGroup{RepType: acctest.Required, Group: DatabaseAutonomousContainerDatabaseDataSourceFilterRepresentation}}

	DatabaseAutonomousContainerDatabaseDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_database_autonomous_container_database.test_autonomous_container_database.id}`}},
	}

	DatabaseAutonomousContainerDatabaseRepresentation = map[string]interface{}{
		//"autonomous_container_database_backup_id": acctest.Representation{RepType: acctest.Optional, Create: `${oci_database_autonomous_container_database_backup.test_autonomous_container_database_backup.id}`},
		"source":                         acctest.Representation{RepType: acctest.Optional, Create: `NONE`},
		"customer_contacts":              acctest.RepresentationGroup{RepType: acctest.Optional, Group: DatabaseAutonomousContainerDatabaseCustomerContactsRepresentation},
		"db_split_threshold":             acctest.Representation{RepType: acctest.Optional, Create: `12`},
		"distribution_affinity":          acctest.Representation{RepType: acctest.Optional, Create: `MINIMUM_DISTRIBUTION`},
		"net_services_architecture":      acctest.Representation{RepType: acctest.Optional, Create: `DEDICATED`},
		"vm_failover_reservation":        acctest.Representation{RepType: acctest.Optional, Create: `25`},
		"version_preference":             acctest.Representation{RepType: acctest.Optional, Create: `LATEST_RELEASE_UPDATE`, Update: `NEXT_RELEASE_UPDATE`},
		"display_name":                   acctest.Representation{RepType: acctest.Required, Create: `containerDatabase2`, Update: `displayName2`},
		"patch_model":                    acctest.Representation{RepType: acctest.Required, Create: `RELEASE_UPDATES`, Update: `RELEASE_UPDATE_REVISIONS`},
		"db_version":                     acctest.Representation{RepType: acctest.Required, Create: utils.GetEnvSettingWithDefault("acd_db_version", "19.28.0.1.0")},
		"cloud_autonomous_vm_cluster_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_database_cloud_autonomous_vm_cluster.test_cloud_autonomous_vm_cluster.id}`},
		"backup_config":                  acctest.RepresentationGroup{RepType: acctest.Optional, Group: ACDatabaseBackupConfigRepresentation},
		"compartment_id":                 acctest.Representation{RepType: acctest.Optional, Create: `${var.compartment_id}`},
		"defined_tags":                   acctest.Representation{RepType: acctest.Optional, Create: `${tomap({"${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}" = "value"})}`, Update: `${tomap({"${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}" = "updatedValue"})}`},
		"freeform_tags":                  acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"kms_key_id":                     acctest.Representation{RepType: acctest.Optional, Create: `${lookup(data.oci_kms_keys.test_keys_dependency.keys[0], "id")}`},
		"maintenance_window_details":     acctest.RepresentationGroup{RepType: acctest.Optional, Group: DatabaseAutonomousContainerDatabaseMaintenanceWindowDetailsRepresentation},
		"service_level_agreement_type":   acctest.Representation{RepType: acctest.Optional, Create: `STANDARD`},
		"vault_id":                       acctest.Representation{RepType: acctest.Optional, Create: `${data.oci_kms_vault.test_vault.id}`},
		"db_name":                        acctest.Representation{RepType: acctest.Optional, Create: `DBNAME`},
		"is_dst_file_update_enabled":     acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		//"failover_trigger":               acctest.Representation{RepType: acctest.Optional, Create: `0`, Update: `1`},
		//"reinstate_trigger":              acctest.Representation{RepType: acctest.Optional, Create: `0`, Update: `1`},
		//"switchover_trigger":             acctest.Representation{RepType: acctest.Optional, Create: `0`, Update: `1`},
	}

	DatabaseAutonomousContainerDatabaseRepresentationWithOkvEpg = map[string]interface{}{
		"okv_end_point_group_name": acctest.Representation{RepType: acctest.Optional, Create: `DUMMY_OKV_EPG_GROUP`, Update: `DUMMY_OKV_EPG_GROUP_2`},
		//"autonomous_container_database_backup_id": acctest.Representation{RepType: acctest.Optional, Create: `${oci_database_autonomous_container_database_backup.test_autonomous_container_database_backup.id}`},
		"source":                         acctest.Representation{RepType: acctest.Optional, Create: `NONE`},
		"db_split_threshold":             acctest.Representation{RepType: acctest.Optional, Create: `12`},
		"distribution_affinity":          acctest.Representation{RepType: acctest.Optional, Create: `MINIMUM_DISTRIBUTION`},
		"net_services_architecture":      acctest.Representation{RepType: acctest.Optional, Create: `DEDICATED`},
		"vm_failover_reservation":        acctest.Representation{RepType: acctest.Optional, Create: `25`},
		"version_preference":             acctest.Representation{RepType: acctest.Optional, Create: `LATEST_RELEASE_UPDATE`, Update: `NEXT_RELEASE_UPDATE`},
		"display_name":                   acctest.Representation{RepType: acctest.Required, Create: `containerDatabase2`, Update: `displayName2`},
		"patch_model":                    acctest.Representation{RepType: acctest.Required, Create: `RELEASE_UPDATES`, Update: `RELEASE_UPDATE_REVISIONS`},
		"db_version":                     acctest.Representation{RepType: acctest.Required, Create: utils.GetEnvSettingWithDefault("acd_db_version", "19.28.0.1.0")},
		"cloud_autonomous_vm_cluster_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_database_cloud_autonomous_vm_cluster.test_cloud_autonomous_vm_cluster.id}`},
		"backup_config":                  acctest.RepresentationGroup{RepType: acctest.Optional, Group: ACDatabaseBackupConfigRepresentation},
		"compartment_id":                 acctest.Representation{RepType: acctest.Optional, Create: `${var.compartment_id}`},
		"defined_tags":                   acctest.Representation{RepType: acctest.Optional, Create: `${tomap({"${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}" = "value"})}`, Update: `${tomap({"${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}" = "updatedValue"})}`},
		"freeform_tags":                  acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"maintenance_window_details":     acctest.RepresentationGroup{RepType: acctest.Optional, Group: DatabaseAutonomousContainerDatabaseMaintenanceWindowDetailsRepresentation},
		"service_level_agreement_type":   acctest.Representation{RepType: acctest.Optional, Create: `STANDARD`},
		"db_name":                        acctest.Representation{RepType: acctest.Optional, Create: `DBNAME`},
		"is_dst_file_update_enabled":     acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"key_store_id":                   acctest.Representation{RepType: acctest.Optional, Create: `${oci_database_key_store.test_key_store.id}`, Update: `${oci_database_key_store.test_key_store.id}`},
		//"failover_trigger":               acctest.Representation{RepType: acctest.Optional, Create: `0`, Update: `1`},
		//"reinstate_trigger":              acctest.Representation{RepType: acctest.Optional, Create: `0`, Update: `1`},
		//"switchover_trigger":             acctest.Representation{RepType: acctest.Optional, Create: `0`, Update: `1`},
	}

	DatabaseAutonomousContainerDatabaseBackupConfigRepresentation = map[string]interface{}{
		"backup_destination_details": acctest.RepresentationGroup{RepType: acctest.Optional, Group: autonomousContainerDatabaseBackupConfigBackupDestinationDetailsRepresentation},
		"recovery_window_in_days":    acctest.Representation{RepType: acctest.Optional, Create: `10`, Update: `11`},
	}

	DatabaseAutonomousContainerDatabaseBackupConfigRepresentationDisableRetentionOnUpdate = map[string]interface{}{
		"backup_destination_details": acctest.RepresentationGroup{RepType: acctest.Optional, Group: autonomousContainerDatabaseBackupConfigBackupDestinationDetailsRepresentation1},
		"recovery_window_in_days":    acctest.Representation{RepType: acctest.Optional, Create: `10`, Update: `11`},
	}

	DatabaseAutonomousContainerDatabaseBackupConfigWithRAUpdateRepresentation = map[string]interface{}{
		"backup_destination_details": acctest.RepresentationGroup{RepType: acctest.Optional, Group: autonomousContainerDatabaseBackupConfigBackupDestinationDetailsWithRAUpdateRepresentation},
		"recovery_window_in_days":    acctest.Representation{RepType: acctest.Optional, Create: `10`},
	}
	DatabaseAutonomousContainerDatabaseBackupConfigWithRARepresentation = map[string]interface{}{
		"backup_destination_details": acctest.RepresentationGroup{RepType: acctest.Required, Group: autonomousContainerDatabaseBackupConfigBackupDestinationDetailsWithRARepresentation},
	}

	DatabaseAutonomousContainerDatabaseCustomerContactsRepresentation = map[string]interface{}{
		"email": acctest.Representation{RepType: acctest.Optional, Create: `test1@oracle.com`, Update: `test2@oracle.com`},
	}

	DatabaseAddStandbyAutonomousContainerDatabaseBackupConfigRepresentation = map[string]interface{}{
		"backup_destination_details": acctest.RepresentationGroup{RepType: acctest.Optional, Group: autonomousContainerDatabaseBackupConfigBackupDestinationDetailsRepresentation},
		"recovery_window_in_days":    acctest.Representation{RepType: acctest.Optional, Create: `7`, Update: `7`},
	}
	AddStandbyAutonomousContainerDatabaseBackupConfigWithNoUpdateRepresentation = map[string]interface{}{
		"backup_destination_details": acctest.RepresentationGroup{RepType: acctest.Optional, Group: autonomousContainerDatabaseBackupConfigBackupDestinationDetailsRepresentationWithNoUpdate},
		"recovery_window_in_days":    acctest.Representation{RepType: acctest.Optional, Create: `10`},
	}

	DatabaseAutonomousContainerDatabaseMaintenanceWindowDetailsRepresentation = map[string]interface{}{
		"preference":                    acctest.Representation{RepType: acctest.Optional, Create: `CUSTOM_PREFERENCE`},
		"custom_action_timeout_in_mins": acctest.Representation{RepType: acctest.Optional, Create: `10`, Update: `11`},
		"days_of_week":                  acctest.RepresentationGroup{RepType: acctest.Optional, Group: DatabaseAutonomousContainerDatabaseMaintenanceWindowDetailsDaysOfWeekRepresentation},
		"hours_of_day":                  acctest.Representation{RepType: acctest.Optional, Create: []string{`4`}, Update: []string{`8`}},
		//"is_custom_action_timeout_enabled": acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		//"lead_time_in_weeks":               acctest.Representation{RepType: acctest.Optional, Create: `10`, Update: `11`},
		"months":         []acctest.RepresentationGroup{{RepType: acctest.Optional, Group: DatabaseAutonomousContainerDatabaseMaintenanceWindowDetailsMonthsRepresentation}, {RepType: acctest.Optional, Group: DatabaseAutonomousContainerDatabaseMaintenanceWindowDetailsMonthsRepresentation2}, {RepType: acctest.Optional, Group: DatabaseAutonomousContainerDatabaseMaintenanceWindowDetailsMonthsRepresentation3}, {RepType: acctest.Optional, Group: DatabaseAutonomousContainerDatabaseMaintenanceWindowDetailsMonthsRepresentation4}},
		"patching_mode":  acctest.Representation{RepType: acctest.Optional, Create: `ROLLING`, Update: `NONROLLING`},
		"skip_ru":        acctest.Representation{RepType: acctest.Optional, Create: []string{`true`, `false`, `true`, `false`}, Update: []string{`true`, `false`, `true`, `false`}},
		"weeks_of_month": acctest.Representation{RepType: acctest.Optional, Create: []string{`1`}, Update: []string{`2`}},
	}
	DatabaseAutonomousContainerDatabaseMaintenanceWindowDetailsNoPreferenceRepresentation = map[string]interface{}{
		"preference": acctest.Representation{RepType: acctest.Required, Create: `NO_PREFERENCE`},
	}

	DatabaseAutonomousContainerDatabaseMaintenanceWindowDetailsDaysOfWeekRepresentation = map[string]interface{}{
		"name": acctest.Representation{RepType: acctest.Required, Create: `MONDAY`, Update: `TUESDAY`},
	}
	DatabaseAutonomousContainerDatabaseMaintenanceWindowDetailsMonthsRepresentation = map[string]interface{}{
		"name": acctest.Representation{RepType: acctest.Required, Create: `JANUARY`, Update: `FEBRUARY`},
	}
	DatabaseAutonomousContainerDatabaseMaintenanceWindowDetailsMonthsRepresentation2 = map[string]interface{}{
		"name": acctest.Representation{RepType: acctest.Required, Create: `APRIL`, Update: `MAY`},
	}

	DatabaseAutonomousContainerDatabaseMaintenanceWindowDetailsMonthsRepresentation3 = map[string]interface{}{
		"name": acctest.Representation{RepType: acctest.Required, Create: `JULY`, Update: `AUGUST`},
	}
	DatabaseAutonomousContainerDatabaseMaintenanceWindowDetailsMonthsRepresentation4 = map[string]interface{}{
		"name": acctest.Representation{RepType: acctest.Required, Create: `OCTOBER`, Update: `NOVEMBER`},
	}

	DatabaseAutonomousContainerDatabaseResourceDependencies = DatabaseAutonomousVmClusterRequiredOnlyResource +
		acctest.GenerateResourceFromRepresentationMap("oci_database_backup_destination", "test_backup_destination", acctest.Optional, acctest.Create, DatabaseBackupDestinationRepresentation) +
		OkvSecretVariableStr +
		acctest.GenerateResourceFromRepresentationMap("oci_database_key_store", "test_key_store", acctest.Optional, acctest.Create, DatabaseKeyStoreRepresentation) + KeyResourceDependencyConfigDbaas

	DatabaseKmsKeyVersionRepresentation = map[string]interface{}{
		"key_id":              acctest.Representation{RepType: acctest.Required, Create: `${lookup(data.oci_kms_keys.test_keys_dependency.keys[0], "id")}`},
		"management_endpoint": acctest.Representation{RepType: acctest.Required, Create: `${data.oci_kms_vault.test_vault.management_endpoint}`},
		"time_of_deletion":    acctest.Representation{RepType: acctest.Required, Create: keyVersionDeletionTime.Format(time.RFC3339Nano)},
	}
	DatabaseKmsKeyVersionDataSourceRepresentation = map[string]interface{}{
		"key_id":              acctest.Representation{RepType: acctest.Required, Create: `${lookup(data.oci_kms_keys.test_keys_dependency.keys[0], "id")}`},
		"management_endpoint": acctest.Representation{RepType: acctest.Required, Create: `${data.oci_kms_vault.test_vault.management_endpoint}`},
	}
	KmsVersionResourceConfig = acctest.GenerateDataSourceFromRepresentationMap("oci_kms_key_versions", "test_key_versions", acctest.Optional, acctest.Update, DatabaseKmsKeyVersionDataSourceRepresentation)

	ATPDAutonomousContainerDatabaseResourceDependencies = DatabaseCloudAutonomousVmClusterRequiredOnlyResource + KeyResourceDependencyConfigDbaas + KmsVersionResourceConfig

	ATPDAutonomousContainerDatabaseResourceWithOkvDependencies = DatabaseCloudAutonomousVmClusterRequiredOnlyResource + KeyResourceDependencyConfigDbaas +
		acctest.GenerateResourceFromRepresentationMap("oci_database_key_store", "test_key_store", acctest.Required, acctest.Create, DatabaseKeyStoreRepresentation) + OkvSecretVariableStr

	//multi standby dg
	DatabaseAdbdAutonomousContainerDatabaseWithDGConfig = DatabaseAutonomousContainerDatabaseDataguardAssociationResourceConfig +
		acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_container_database_add_standby", "test_autonomous_container_database_add_standby", acctest.Optional, acctest.Create,
			acctest.RepresentationCopyWithRemovedProperties(DatabaseAdbdAutonomousContainerDatabaseAddStandbyRepresentation, []string{"is_automatic_failover_enabled", "fast_start_fail_over_lag_limit_in_seconds"}))

	DatabaseAdbccAutonomousContainerDatabaseWithDGConfig = AdbccDgSetupDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_container_database_add_standby", "test_autonomous_container_database_add_standby", acctest.Optional, acctest.Create,
			acctest.RepresentationCopyWithRemovedProperties(DatabaseAdbccAutonomousContainerDatabaseAddStandbyRepresentation, []string{"is_automatic_failover_enabled", "fast_start_fail_over_lag_limit_in_seconds"}))

	//multi standby dg
	DatabaseAdbdAutonomousContainerDatabaseWithDGFsfoConfig = DatabaseAutonomousContainerDatabaseDataguardAssociationResourceConfig +
		acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_container_database_add_standby", "test_autonomous_container_database_add_standby", acctest.Optional, acctest.Create,
			acctest.RepresentationCopyWithRemovedProperties(DatabaseAdbdAutonomousContainerDatabaseAddStandbyRepresentation, []string{}))

	DatabaseAdbdAutonomousContainerDatabaseWithDGUpdateConfig = DatabaseAutonomousContainerDatabaseDataguardAssociationUpdateResourceConfig +
		acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_container_database_add_standby", "test_autonomous_container_database_add_standby", acctest.Optional, acctest.Create,
			acctest.RepresentationCopyWithRemovedProperties(DatabaseAdbdAutonomousContainerDatabaseAddStandbyRepresentation, []string{}))

	DatabaseAdbdAutonomousContainerDatabaseWithDGConfigReinstate = AdbdDgDependencies + acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_container_database", "test_autonomous_container_database", acctest.Optional, acctest.Create,
		acctest.RepresentationCopyWithNewProperties(acctest.RepresentationCopyWithRemovedProperties(DatabaseAutonomousContainerDatabaseRepresentation, []string{"vault_id", "kms_key_id"}), map[string]interface{}{
			"service_level_agreement_type": acctest.Representation{RepType: acctest.Optional, Create: `STANDARD`},
			"protection_mode":              acctest.Representation{RepType: acctest.Optional, Create: `MAXIMUM_AVAILABILITY`},
			"lifecycle":                    acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreDataguardChangesRep},
			"reinstate_trigger":            acctest.Representation{RepType: acctest.Optional, Create: `1`},
		})) +
		acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_container_database_add_standby", "test_autonomous_container_database_add_standby", acctest.Optional, acctest.Create,
			acctest.RepresentationCopyWithRemovedProperties(DatabaseAdbdAutonomousContainerDatabaseAddStandbyRepresentation, []string{"is_automatic_failover_enabled", "fast_start_fail_over_lag_limit_in_seconds"}))

	DatabaseAdbdAutonomousContainerDatabaseWithDGConfigReinstateAndSwitchover = AdbdDgDependencies + acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_container_database", "test_autonomous_container_database", acctest.Optional, acctest.Create,
		acctest.RepresentationCopyWithNewProperties(acctest.RepresentationCopyWithRemovedProperties(DatabaseAutonomousContainerDatabaseRepresentation, []string{"vault_id", "kms_key_id"}), map[string]interface{}{
			"service_level_agreement_type": acctest.Representation{RepType: acctest.Optional, Create: `STANDARD`},
			"protection_mode":              acctest.Representation{RepType: acctest.Optional, Create: `MAXIMUM_AVAILABILITY`},
			"lifecycle":                    acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreDataguardChangesRep},
			"reinstate_trigger":            acctest.Representation{RepType: acctest.Optional, Create: `1`},
			"switchover_trigger":           acctest.Representation{RepType: acctest.Optional, Create: `1`},
		})) +
		acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_container_database_add_standby", "test_autonomous_container_database_add_standby", acctest.Optional, acctest.Create,
			acctest.RepresentationCopyWithRemovedProperties(DatabaseAdbdAutonomousContainerDatabaseAddStandbyRepresentation, []string{"is_automatic_failover_enabled", "fast_start_fail_over_lag_limit_in_seconds"}))

	DatabaseAdbdAutonomousContainerDatabaseWithDGConfigSwitchover = AdbdDgDependencies + acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_container_database", "test_autonomous_container_database", acctest.Optional, acctest.Create,
		acctest.RepresentationCopyWithNewProperties(acctest.RepresentationCopyWithRemovedProperties(DatabaseAutonomousContainerDatabaseRepresentation, []string{"vault_id", "kms_key_id"}), map[string]interface{}{
			"service_level_agreement_type": acctest.Representation{RepType: acctest.Optional, Create: `STANDARD`},
			"protection_mode":              acctest.Representation{RepType: acctest.Optional, Create: `MAXIMUM_AVAILABILITY`},
			"lifecycle":                    acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreDataguardChangesRep},
			"switchover_trigger":           acctest.Representation{RepType: acctest.Optional, Create: `1`},
		})) +
		acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_container_database_add_standby", "test_autonomous_container_database_add_standby", acctest.Optional, acctest.Create,
			acctest.RepresentationCopyWithRemovedProperties(DatabaseAdbdAutonomousContainerDatabaseAddStandbyRepresentation, []string{"is_automatic_failover_enabled", "fast_start_fail_over_lag_limit_in_seconds"}))

	DatabaseAdbccAutonomousContainerDatabaseWithDGConfigSwitchover = AdbccDgDependencies + acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_container_database", "test_autonomous_container_database", acctest.Optional, acctest.Create,
		acctest.RepresentationCopyWithNewProperties(acctest.GetUpdatedRepresentationCopy("maintenance_window_details", acctest.RepresentationGroup{RepType: acctest.Optional, Group: DatabaseAutonomousContainerDatabaseMaintenanceWindowDetailsNoPreferenceRepresentation}, ACDatabaseRepresentation), map[string]interface{}{
			"service_level_agreement_type": acctest.Representation{RepType: acctest.Optional, Create: `STANDARD`},
			"protection_mode":              acctest.Representation{RepType: acctest.Optional, Create: `MAXIMUM_AVAILABILITY`},
			"lifecycle":                    acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreDataguardChangesRep},
			"switchover_trigger":           acctest.Representation{RepType: acctest.Optional, Create: `1`},
		})) +
		acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_container_database_add_standby", "test_autonomous_container_database_add_standby", acctest.Optional, acctest.Create,
			acctest.RepresentationCopyWithRemovedProperties(DatabaseAdbccAutonomousContainerDatabaseAddStandbyRepresentation, []string{"is_automatic_failover_enabled", "fast_start_fail_over_lag_limit_in_seconds"}))

	AdbdStandbyACDRepresentation = map[string]interface{}{
		"depends_on":                     []string{"oci_database_autonomous_container_database.test_autonomous_container_database"},
		"compartment_id":                 acctest.Representation{RepType: acctest.Optional, Create: `${var.compartment_id}`},
		"cloud_autonomous_vm_cluster_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_database_cloud_autonomous_vm_cluster.peer_cloud_autonomous_vm_cluster.id}`},
		"display_name":                   acctest.Representation{RepType: acctest.Required, Create: `FirstStandby`},
		"patch_model":                    acctest.Representation{RepType: acctest.Required, Create: `RELEASE_UPDATES`, Update: `RELEASE_UPDATE_REVISIONS`},
		//"defined_tags":             acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value", "Oracle-Standard.Org", "HR")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		//"freeform_tags":            acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
	}

	AdbccStandbyACDRepresentation = map[string]interface{}{
		"depends_on":               []string{"oci_database_autonomous_container_database.test_autonomous_container_database"},
		"compartment_id":           acctest.Representation{RepType: acctest.Optional, Create: `${var.compartment_id}`},
		"autonomous_vm_cluster_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_database_autonomous_vm_cluster.peer_autonomous_vm_cluster.id}`},
		"display_name":             acctest.Representation{RepType: acctest.Required, Create: `FirstStandby`},
		"patch_model":              acctest.Representation{RepType: acctest.Required, Create: `RELEASE_UPDATES`, Update: `RELEASE_UPDATE_REVISIONS`},
		//"defined_tags":             acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value", "Oracle-Standard.Org", "HR")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		//"freeform_tags":            acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
	}

	AdbdStandbyAcdResourceConfig         = acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_container_database", "standby_acd", acctest.Optional, acctest.Create, AdbdStandbyACDRepresentation)
	AdbccStandbyAcdResourceConfig        = acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_container_database", "standby_acd", acctest.Optional, acctest.Create, AdbccStandbyACDRepresentation)
	AdbdStandbyAcdWithDgSwitchoverConfig = acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_container_database", "standby_acd", acctest.Optional, acctest.Create,
		acctest.RepresentationCopyWithNewProperties(AdbdStandbyACDRepresentation, map[string]interface{}{
			"lifecycle":          acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreDataguardChangesRep},
			"switchover_trigger": acctest.Representation{RepType: acctest.Required, Create: `1`},
		}))
	AdbccStandbyAcdWithDgSwitchoverConfig = acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_container_database", "standby_acd", acctest.Optional, acctest.Create,
		acctest.RepresentationCopyWithNewProperties(AdbccStandbyACDRepresentation, map[string]interface{}{
			"lifecycle":          acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreDataguardChangesRep},
			"switchover_trigger": acctest.Representation{RepType: acctest.Required, Create: `1`},
		}))
	AdbdStandbyAcdWithDgFailoverConfig = acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_container_database", "standby_acd", acctest.Optional, acctest.Create,
		acctest.RepresentationCopyWithNewProperties(AdbdStandbyACDRepresentation, map[string]interface{}{
			"lifecycle":        acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreDataguardChangesRep},
			"failover_trigger": acctest.Representation{RepType: acctest.Required, Create: `1`},
		}))
)

// multi standby test
func TestDatabaseAdbdAutonomousContainerDatabaseResource_switchover(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseAdbdAutonomousContainerDatabaseResource_switchover")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_database_autonomous_container_database.test_autonomous_container_database"
	standbyResourceName := "oci_database_autonomous_container_database.standby_acd"
	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify Create DG
		{
			Config: config + compartmentIdVariableStr + DatabaseAdbdAutonomousContainerDatabaseWithDGConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "cloud_autonomous_vm_cluster_id"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "containerDatabase2"),
				resource.TestCheckResourceAttr(resourceName, "patch_model", "RELEASE_UPDATES"),
			),
		},
		// NEW STEP: Refresh state
		{
			RefreshState: true, // reload state
		},
		{
			Config:             config + compartmentIdVariableStr + DatabaseAdbdAutonomousContainerDatabaseWithDGConfig + AdbdStandbyAcdResourceConfig,
			ImportState:        true,
			ImportStateIdFunc:  getStandbyAcdOcid(resourceName),
			ImportStatePersist: true,
			ResourceName:       standbyResourceName,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "containerDatabase2"),
				resource.TestCheckResourceAttr(standbyResourceName, "role", "STANDBY"),
			),
		},
		//switchover
		{
			Config: config + compartmentIdVariableStr + DatabaseAdbdAutonomousContainerDatabaseWithDGConfig + AdbdStandbyAcdWithDgSwitchoverConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(standbyResourceName, "id"),
				resource.TestCheckResourceAttr(standbyResourceName, "display_name", "FirstStandby"),
				resource.TestCheckResourceAttr(standbyResourceName, "role", "PRIMARY"),
			),
		},
		//switchover again
		{
			Config: config + compartmentIdVariableStr + DatabaseAdbdAutonomousContainerDatabaseWithDGConfigSwitchover + AdbdStandbyAcdWithDgSwitchoverConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "containerDatabase2"),
			),
		},
		//Delete standby
		{
			Config: config + compartmentIdVariableStr + DatabaseAdbdAutonomousContainerDatabaseWithDGConfigSwitchover,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "containerDatabase2"),
			),
		},
	})
}

func TestDatabaseAdbccAutonomousContainerDatabaseResource_switchover(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseAdbccAutonomousContainerDatabaseResource_switchover")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_database_autonomous_container_database.test_autonomous_container_database"
	standbyResourceName := "oci_database_autonomous_container_database.standby_acd"
	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify Create DG
		{
			Config: config + compartmentIdVariableStr + DatabaseAdbccAutonomousContainerDatabaseWithDGConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "autonomous_vm_cluster_id"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "containerdatabases2"),
				resource.TestCheckResourceAttr(resourceName, "patch_model", "RELEASE_UPDATES"),
			),
		},
		// NEW STEP: Refresh state
		{
			RefreshState: true, // reload state
		},
		{
			Config:             config + compartmentIdVariableStr + DatabaseAdbccAutonomousContainerDatabaseWithDGConfig + AdbccStandbyAcdResourceConfig,
			ImportState:        true,
			ImportStateIdFunc:  getStandbyAcdOcid(resourceName),
			ImportStatePersist: true,
			ResourceName:       standbyResourceName,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "containerdatabases2"),
				resource.TestCheckResourceAttr(standbyResourceName, "role", "STANDBY"),
			),
		},
		//switchover
		{
			Config: config + compartmentIdVariableStr + DatabaseAdbccAutonomousContainerDatabaseWithDGConfig + AdbccStandbyAcdWithDgSwitchoverConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(standbyResourceName, "id"),
				resource.TestCheckResourceAttr(standbyResourceName, "display_name", "FirstStandby"),
				resource.TestCheckResourceAttr(standbyResourceName, "role", "PRIMARY"),
			),
		},
		//switchover again
		{
			Config: config + compartmentIdVariableStr + DatabaseAdbccAutonomousContainerDatabaseWithDGConfigSwitchover + AdbccStandbyAcdWithDgSwitchoverConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "containerdatabases2"),
			),
		},
		//Delete standby
		{
			Config: config + compartmentIdVariableStr + DatabaseAdbccAutonomousContainerDatabaseWithDGConfigSwitchover,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "containerdatabases2"),
			),
		},
	})
}

// multi standby test
func TestDatabaseAdbdAutonomousContainerDatabaseResource_failover_reinstate(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseAdbdAutonomousContainerDatabaseResource_failover_reinstate")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_database_autonomous_container_database.test_autonomous_container_database"
	standbyResourceName := "oci_database_autonomous_container_database.standby_acd"
	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify Create DG
		{
			Config: config + compartmentIdVariableStr + DatabaseAdbdAutonomousContainerDatabaseWithDGConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "cloud_autonomous_vm_cluster_id"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "containerDatabase2"),
				resource.TestCheckResourceAttr(resourceName, "patch_model", "RELEASE_UPDATES"),
			),
		},
		// NEW STEP: Refresh state
		{
			RefreshState: true, // reload state
		},
		{
			Config:             config + compartmentIdVariableStr + DatabaseAdbdAutonomousContainerDatabaseWithDGConfig + AdbdStandbyAcdResourceConfig,
			ImportState:        true,
			ImportStateIdFunc:  getStandbyAcdOcid(resourceName),
			ImportStatePersist: true,
			ResourceName:       standbyResourceName,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "containerDatabase2"),
				resource.TestCheckResourceAttr(standbyResourceName, "role", "STANDBY"),
			),
		},
		//failover
		{
			Config: config + compartmentIdVariableStr + DatabaseAdbdAutonomousContainerDatabaseWithDGConfig + AdbdStandbyAcdWithDgFailoverConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(standbyResourceName, "id"),
				resource.TestCheckResourceAttr(standbyResourceName, "display_name", "FirstStandby"),
				resource.TestCheckResourceAttr(standbyResourceName, "role", "PRIMARY"),
			),
		},
		//reinstate
		{
			Config: config + compartmentIdVariableStr + DatabaseAdbdAutonomousContainerDatabaseWithDGConfigReinstate + AdbdStandbyAcdWithDgFailoverConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(standbyResourceName, "id"),
				resource.TestCheckResourceAttr(standbyResourceName, "display_name", "FirstStandby"),
				resource.TestCheckResourceAttr(resourceName, "role", "STANDBY"),
			),
		},
		//switchover
		{
			Config: config + compartmentIdVariableStr + DatabaseAdbdAutonomousContainerDatabaseWithDGConfigReinstateAndSwitchover + AdbdStandbyAcdWithDgFailoverConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(standbyResourceName, "id"),
				resource.TestCheckResourceAttr(standbyResourceName, "display_name", "FirstStandby"),
			),
		},

		//delete standby
		{
			Config: config + compartmentIdVariableStr + DatabaseAdbdAutonomousContainerDatabaseWithDGConfigReinstateAndSwitchover,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "containerDatabase2"),
			),
		},
	})
}

// multi standby test
func TestDatabaseAdbdAutonomousContainerDatabaseResource_updateDataGuard(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseAdbdAutonomousContainerDatabaseResource_updateDataGuard")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_database_autonomous_container_database.test_autonomous_container_database"
	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify Create DG
		{
			Config: config + compartmentIdVariableStr + DatabaseAdbdAutonomousContainerDatabaseWithDGFsfoConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "cloud_autonomous_vm_cluster_id"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "containerDatabase2"),
				resource.TestCheckResourceAttr(resourceName, "patch_model", "RELEASE_UPDATES"),
			),
		},
		// verify Update DG
		{
			Config: config + compartmentIdVariableStr + DatabaseAdbdAutonomousContainerDatabaseWithDGUpdateConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "cloud_autonomous_vm_cluster_id"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "patch_model", "RELEASE_UPDATE_REVISIONS"),
			),
		},
	})
}

// issue-routing-tag: database/dbaas-atp-d
func TestDatabaseAutonomousContainerDatabaseResource_basic(t *testing.T) {
	//t.Skip("Skip this test as AEI and its api no longer exists.")

	shouldSkipADBDtest := os.Getenv("TF_VAR_should_skip_adbd_test")

	if shouldSkipADBDtest == "true" {
		t.Skip("Skipping TestDatabaseCrossRegionDisasterRecovery_basic test.\n" + "Current TF_VAR_should_skip_adbd_test=" + shouldSkipADBDtest)
	}

	httpreplay.SetScenario("TestDatabaseAutonomousContainerDatabaseResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_database_autonomous_container_database.test_autonomous_container_database"
	datasourceName := "data.oci_database_autonomous_container_databases.test_autonomous_container_databases"
	singularDatasourceName := "data.oci_database_autonomous_container_database.test_autonomous_container_database"

	AutonomousContainerDatabaseDedicatedMaintenanceWindowDetailsRepresentation := acctest.RepresentationCopyWithRemovedProperties(
		acctest.GetUpdatedRepresentationCopy("months",
			[]acctest.RepresentationGroup{{RepType: acctest.Optional, Group: DatabaseAutonomousContainerDatabaseMaintenanceWindowDetailsMonthsRepresentation}, {RepType: acctest.Optional, Group: DatabaseAutonomousContainerDatabaseMaintenanceWindowDetailsMonthsRepresentation2}, {RepType: acctest.Optional, Group: DatabaseAutonomousContainerDatabaseMaintenanceWindowDetailsMonthsRepresentation3}, {RepType: acctest.Optional, Group: DatabaseAutonomousContainerDatabaseMaintenanceWindowDetailsMonthsRepresentation4}},
			DatabaseAutonomousContainerDatabaseMaintenanceWindowDetailsRepresentation), []string{"lead_time_in_weeks"})
	AutonomousContainerDatabaseDedicatedRepresentation := acctest.GetUpdatedRepresentationCopy("maintenance_window_details", acctest.RepresentationGroup{RepType: acctest.Optional, Group: AutonomousContainerDatabaseDedicatedMaintenanceWindowDetailsRepresentation}, DatabaseAutonomousContainerDatabaseRepresentation)

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+ATPDAutonomousContainerDatabaseResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_container_database", "test_autonomous_container_database", acctest.Optional, acctest.Create, DatabaseAutonomousContainerDatabaseRepresentation), "database", "autonomousContainerDatabase", t)

	acctest.ResourceTest(t, testAccCheckDatabaseAutonomousContainerDatabaseDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + ATPDAutonomousContainerDatabaseResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_container_database", "test_autonomous_container_database", acctest.Required, acctest.Create, DatabaseAutonomousContainerDatabaseRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				//resource.TestCheckResourceAttrSet(resourceName, "autonomous_container_database_backup_id"),
				resource.TestCheckResourceAttrSet(resourceName, "cloud_autonomous_vm_cluster_id"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "containerDatabase2"),
				resource.TestCheckResourceAttr(resourceName, "patch_model", "RELEASE_UPDATES"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next create
		{
			Config: config + compartmentIdVariableStr + ATPDAutonomousContainerDatabaseResourceDependencies,
		},
		// verify create with optionals
		{
			Config: config + compartmentIdVariableStr + ATPDAutonomousContainerDatabaseResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_container_database", "test_autonomous_container_database", acctest.Optional, acctest.Create,
					acctest.GetUpdatedRepresentationCopy("maintenance_window_details", acctest.RepresentationGroup{RepType: acctest.Optional, Group: DatabaseAutonomousContainerDatabaseMaintenanceWindowDetailsNoPreferenceRepresentation}, DatabaseAutonomousContainerDatabaseRepresentation)),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				//resource.TestCheckResourceAttrSet(resourceName, "autonomous_container_database_backup_id"),
				//resource.TestCheckResourceAttrSet(resourceName, "autonomous_exadata_infrastructure_id"),
				resource.TestCheckResourceAttr(resourceName, "backup_config.#", "1"),
				//resource.TestCheckResourceAttr(resourceName, "backup_config.0.recovery_window_in_days", "7"),
				resource.TestCheckResourceAttrSet(resourceName, "cloud_autonomous_vm_cluster_id"),
				resource.TestCheckResourceAttr(resourceName, "backup_config.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "backup_config.0.recovery_window_in_days", "10"),
				resource.TestCheckResourceAttr(resourceName, "backup_config.0.backup_destination_details.0.is_retention_lock_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "backup_config.0.backup_destination_details.0.backup_retention_policy_on_terminate", "RETAIN_FOR_72_HOURS"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "customer_contacts.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "customer_contacts.0.email", "test1@oracle.com"),
				resource.TestCheckResourceAttr(resourceName, "db_split_threshold", "12"),
				resource.TestCheckResourceAttr(resourceName, "distribution_affinity", "MINIMUM_DISTRIBUTION"),
				resource.TestCheckResourceAttrSet(resourceName, "db_version"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "containerDatabase2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				//resource.TestCheckResourceAttr(resourceName, "is_automatic_failover_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "is_dst_file_update_enabled", "false"),
				resource.TestCheckResourceAttrSet(resourceName, "kms_key_id"),
				resource.TestCheckResourceAttr(resourceName, "net_services_architecture", "DEDICATED"),
				//resource.TestCheckResourceAttrSet(resourceName, "kms_key_version_id"),
				//resource.TestCheckResourceAttr(resourceName, "maintenance_window_details.#", "1"),
				//resource.TestCheckResourceAttr(resourceName, "maintenance_window_details.0.custom_action_timeout_in_mins", "10"),
				//resource.TestCheckResourceAttr(resourceName, "maintenance_window_details.0.days_of_week.#", "1"),
				//resource.TestCheckResourceAttr(resourceName, "maintenance_window_details.0.days_of_week.0.name", "MONDAY"),
				//resource.TestCheckResourceAttr(resourceName, "maintenance_window_details.0.hours_of_day.#", "1"),
				//resource.TestCheckResourceAttr(resourceName, "maintenance_window_details.0.is_custom_action_timeout_enabled", "false"),
				//resource.TestCheckResourceAttr(resourceName, "maintenance_window_details.0.lead_time_in_weeks", "10"),
				//resource.TestCheckResourceAttr(resourceName, "maintenance_window_details.0.months.#", "1"),
				//resource.TestCheckResourceAttr(resourceName, "maintenance_window_details.0.months.0.name", "APRIL"),
				//resource.TestCheckResourceAttr(resourceName, "maintenance_window_details.0.patching_mode", "ROLLING"),
				//resource.TestCheckResourceAttr(resourceName, "maintenance_window_details.0.preference", "CUSTOM_PREFERENCE"),
				//resource.TestCheckResourceAttr(resourceName, "maintenance_window_details.0.weeks_of_month.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "patch_model", "RELEASE_UPDATES"),
				resource.TestCheckResourceAttr(resourceName, "source", "NONE"),
				// all peer related properties are not returned in GET, hence commented check on the below peer related properties
				//resource.TestCheckResourceAttr(resourceName, "peer_autonomous_container_database_backup_config.#", "1"),
				//resource.TestCheckResourceAttr(resourceName, "peer_autonomous_container_database_backup_config.0.backup_destination_details.#", "1"),
				//resource.TestCheckResourceAttr(resourceName, "peer_autonomous_container_database_backup_config.0.backup_destination_details.0.id", "id"),
				//resource.TestCheckResourceAttr(resourceName, "peer_autonomous_container_database_backup_config.0.backup_destination_details.0.internet_proxy", "internetProxy"),
				//resource.TestCheckResourceAttr(resourceName, "peer_autonomous_container_database_backup_config.0.backup_destination_details.0.type", "NFS"),
				//resource.TestCheckResourceAttr(resourceName, "peer_autonomous_container_database_backup_config.0.backup_destination_details.0.vpc_password", "vpcPassword"),
				//resource.TestCheckResourceAttr(resourceName, "peer_autonomous_container_database_backup_config.0.backup_destination_details.0.vpc_user", "vpcUser"),
				//resource.TestCheckResourceAttr(resourceName, "peer_autonomous_container_database_backup_config.0.recovery_window_in_days", "10"),
				//resource.TestCheckResourceAttrSet(resourceName, "peer_autonomous_container_database_compartment_id"),
				//resource.TestCheckResourceAttrSet(resourceName, "peer_autonomous_vm_cluster_id"),
				resource.TestCheckResourceAttr(resourceName, "service_level_agreement_type", "STANDARD"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "vault_id"),
				resource.TestCheckResourceAttr(resourceName, "version_preference", "LATEST_RELEASE_UPDATE"),
				resource.TestCheckResourceAttr(resourceName, "vm_failover_reservation", "25"),
				resource.TestCheckResourceAttr(resourceName, "db_name", "DBNAME"),
				resource.TestCheckResourceAttr(resourceName, "system_tags.%", "0"),

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

		//verify Update to the compartment (the compartment will be switched back in the next step) and maintenance_window_details
		{
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + ATPDAutonomousContainerDatabaseResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_container_database", "test_autonomous_container_database", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(AutonomousContainerDatabaseDedicatedRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: compartmentIdU},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				//resource.TestCheckResourceAttrSet(resourceName, "autonomous_container_database_backup_id"),
				//resource.TestCheckResourceAttrSet(resourceName, "autonomous_exadata_infrastructure_id"),
				resource.TestCheckResourceAttr(resourceName, "backup_config.#", "1"),
				//resource.TestCheckResourceAttr(resourceName, "backup_config.0.recovery_window_in_days", "10"),
				resource.TestCheckResourceAttrSet(resourceName, "cloud_autonomous_vm_cluster_id"),
				resource.TestCheckResourceAttr(resourceName, "backup_config.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "backup_config.0.recovery_window_in_days", "10"),
				resource.TestCheckResourceAttr(resourceName, "backup_config.0.backup_destination_details.0.is_retention_lock_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "backup_config.0.backup_destination_details.0.backup_retention_policy_on_terminate", "RETAIN_FOR_72_HOURS"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "customer_contacts.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "customer_contacts.0.email", "test1@oracle.com"),
				resource.TestCheckResourceAttr(resourceName, "db_split_threshold", "12"),
				resource.TestCheckResourceAttr(resourceName, "distribution_affinity", "MINIMUM_DISTRIBUTION"),
				resource.TestCheckResourceAttrSet(resourceName, "db_version"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "containerDatabase2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				//resource.TestCheckResourceAttr(resourceName, "is_automatic_failover_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "is_dst_file_update_enabled", "false"),
				resource.TestCheckResourceAttrSet(resourceName, "kms_key_id"),
				resource.TestCheckResourceAttr(resourceName, "net_services_architecture", "DEDICATED"),
				//resource.TestCheckResourceAttrSet(resourceName, "kms_key_version_id"),
				//resource.TestCheckResourceAttr(resourceName, "maintenance_window_details.#", "1"),
				//resource.TestCheckResourceAttr(resourceName, "maintenance_window_details.0.custom_action_timeout_in_mins", "10"),
				//resource.TestCheckResourceAttr(resourceName, "maintenance_window_details.0.days_of_week.#", "1"),
				//resource.TestCheckResourceAttr(resourceName, "maintenance_window_details.0.days_of_week.0.name", "MONDAY"),
				//resource.TestCheckResourceAttr(resourceName, "maintenance_window_details.0.hours_of_day.#", "1"),
				//resource.TestCheckResourceAttr(resourceName, "maintenance_window_details.0.is_custom_action_timeout_enabled", "false"),
				//resource.TestCheckResourceAttr(resourceName, "maintenance_window_details.0.lead_time_in_weeks", "10"),
				//resource.TestCheckResourceAttr(resourceName, "maintenance_window_details.0.months.#", "1"),
				//resource.TestCheckResourceAttr(resourceName, "maintenance_window_details.0.months.0.name", "APRIL"),
				//resource.TestCheckResourceAttr(resourceName, "maintenance_window_details.0.patching_mode", "ROLLING"),
				//resource.TestCheckResourceAttr(resourceName, "maintenance_window_details.0.preference", "CUSTOM_PREFERENCE"),
				//resource.TestCheckResourceAttr(resourceName, "maintenance_window_details.0.weeks_of_month.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "patch_model", "RELEASE_UPDATES"),
				resource.TestCheckResourceAttr(resourceName, "source", "NONE"),
				// all peer related properties are not returned in GET, hence commented check on the below peer related properties
				//resource.TestCheckResourceAttr(resourceName, "peer_autonomous_container_database_backup_config.#", "1"),
				//resource.TestCheckResourceAttr(resourceName, "peer_autonomous_container_database_backup_config.0.backup_destination_details.#", "1"),
				//resource.TestCheckResourceAttr(resourceName, "peer_autonomous_container_database_backup_config.0.backup_destination_details.0.id", "id"),
				//resource.TestCheckResourceAttr(resourceName, "peer_autonomous_container_database_backup_config.0.backup_destination_details.0.internet_proxy", "internetProxy"),
				//resource.TestCheckResourceAttr(resourceName, "peer_autonomous_container_database_backup_config.0.backup_destination_details.0.type", "NFS"),
				//resource.TestCheckResourceAttr(resourceName, "peer_autonomous_container_database_backup_config.0.backup_destination_details.0.vpc_password", "vpcPassword"),
				//resource.TestCheckResourceAttr(resourceName, "peer_autonomous_container_database_backup_config.0.backup_destination_details.0.vpc_user", "vpcUser"),
				//resource.TestCheckResourceAttr(resourceName, "peer_autonomous_container_database_backup_config.0.recovery_window_in_days", "10"),
				//resource.TestCheckResourceAttrSet(resourceName, "peer_autonomous_container_database_compartment_id"),
				//resource.TestCheckResourceAttrSet(resourceName, "peer_autonomous_vm_cluster_id"),
				resource.TestCheckResourceAttr(resourceName, "service_level_agreement_type", "STANDARD"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "vault_id"),
				resource.TestCheckResourceAttr(resourceName, "version_preference", "LATEST_RELEASE_UPDATE"),
				resource.TestCheckResourceAttr(resourceName, "vm_failover_reservation", "25"),
				resource.TestCheckResourceAttr(resourceName, "db_name", "DBNAME"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("resource recreated when it was supposed to be updated")
					}
					return err
				},
			),
		},

		// verify updates to updatable parameters
		{
			Config: config + compartmentIdVariableStr + ATPDAutonomousContainerDatabaseResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_container_database", "test_autonomous_container_database", acctest.Optional, acctest.Update, AutonomousContainerDatabaseDedicatedRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				//resource.TestCheckResourceAttrSet(resourceName, "autonomous_container_database_backup_id"),
				resource.TestCheckResourceAttrSet(resourceName, "cloud_autonomous_vm_cluster_id"),
				resource.TestCheckResourceAttr(resourceName, "backup_config.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "backup_config.0.recovery_window_in_days", "11"),
				resource.TestCheckResourceAttr(resourceName, "backup_config.0.backup_destination_details.0.is_retention_lock_enabled", "true"),
				resource.TestCheckResourceAttr(resourceName, "backup_config.0.backup_destination_details.0.backup_retention_policy_on_terminate", "RETAIN_PER_RETENTION_WINDOW"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "customer_contacts.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "customer_contacts.0.email", "test2@oracle.com"),
				resource.TestCheckResourceAttr(resourceName, "db_split_threshold", "12"),
				resource.TestCheckResourceAttr(resourceName, "distribution_affinity", "MINIMUM_DISTRIBUTION"),
				resource.TestCheckResourceAttrSet(resourceName, "db_version"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				//resource.TestCheckResourceAttr(resourceName, "is_automatic_failover_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "is_dst_file_update_enabled", "true"),
				resource.TestCheckResourceAttrSet(resourceName, "kms_key_id"),
				resource.TestCheckResourceAttr(resourceName, "net_services_architecture", "DEDICATED"),
				//resource.TestCheckResourceAttrSet(resourceName, "kms_key_version_id"),
				//resource.TestCheckResourceAttr(resourceName, "maintenance_window_details.#", "1"),
				//resource.TestCheckResourceAttr(resourceName, "maintenance_window_details.0.custom_action_timeout_in_mins", "11"),
				//resource.TestCheckResourceAttr(resourceName, "maintenance_window_details.0.days_of_week.#", "1"),
				//resource.TestCheckResourceAttr(resourceName, "maintenance_window_details.0.days_of_week.0.name", "TUESDAY"),
				//resource.TestCheckResourceAttr(resourceName, "maintenance_window_details.0.hours_of_day.#", "1"),
				//resource.TestCheckResourceAttr(resourceName, "maintenance_window_details.0.is_custom_action_timeout_enabled", "true"),
				//resource.TestCheckResourceAttr(resourceName, "maintenance_window_details.0.lead_time_in_weeks", "11"),
				//resource.TestCheckResourceAttr(resourceName, "maintenance_window_details.0.months.#", "1"),
				//resource.TestCheckResourceAttr(resourceName, "maintenance_window_details.0.months.0.name", "MAY"),
				//resource.TestCheckResourceAttr(resourceName, "maintenance_window_details.0.patching_mode", "NONROLLING"),
				//resource.TestCheckResourceAttr(resourceName, "maintenance_window_details.0.preference", "CUSTOM_PREFERENCE"),
				//resource.TestCheckResourceAttr(resourceName, "maintenance_window_details.0.weeks_of_month.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "patch_model", "RELEASE_UPDATE_REVISIONS"),
				resource.TestCheckResourceAttr(resourceName, "source", "NONE"),
				// all peer related properties are not returned in GET, hence commented check on the below peer related properties
				//resource.TestCheckResourceAttr(resourceName, "peer_autonomous_container_database_backup_config.#", "1"),
				//resource.TestCheckResourceAttr(resourceName, "peer_autonomous_container_database_backup_config.0.backup_destination_details.#", "1"),
				//resource.TestCheckResourceAttr(resourceName, "peer_autonomous_container_database_backup_config.0.backup_destination_details.0.id", "id"),
				//resource.TestCheckResourceAttr(resourceName, "peer_autonomous_container_database_backup_config.0.backup_destination_details.0.internet_proxy", "internetProxy"),
				//resource.TestCheckResourceAttr(resourceName, "peer_autonomous_container_database_backup_config.0.backup_destination_details.0.type", "NFS"),
				//resource.TestCheckResourceAttr(resourceName, "peer_autonomous_container_database_backup_config.0.backup_destination_details.0.vpc_password", "vpcPassword"),
				//resource.TestCheckResourceAttr(resourceName, "peer_autonomous_container_database_backup_config.0.backup_destination_details.0.vpc_user", "vpcUser"),
				//resource.TestCheckResourceAttr(resourceName, "peer_autonomous_container_database_backup_config.0.recovery_window_in_days", "10"),
				//resource.TestCheckResourceAttrSet(resourceName, "peer_autonomous_container_database_compartment_id"),
				//resource.TestCheckResourceAttrSet(resourceName, "peer_autonomous_vm_cluster_id"),
				resource.TestCheckResourceAttr(resourceName, "service_level_agreement_type", "STANDARD"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "vault_id"),
				resource.TestCheckResourceAttr(resourceName, "version_preference", "NEXT_RELEASE_UPDATE"),
				resource.TestCheckResourceAttr(resourceName, "vm_failover_reservation", "25"),
				resource.TestCheckResourceAttr(resourceName, "db_name", "DBNAME"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("Resource recreated when it was supposed to be updated.")
					}
					return err
				},
			),
		},
		//verify rotate key
		{
			Config: config + compartmentIdVariableStr + ATPDAutonomousContainerDatabaseResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_container_database", "test_autonomous_container_database", acctest.Optional, acctest.Update,
					acctest.RepresentationCopyWithNewProperties(AutonomousContainerDatabaseDedicatedRepresentation, map[string]interface{}{
						"rotate_key_trigger": acctest.Representation{RepType: acctest.Optional, Create: `true`},
						"key_version_id":     acctest.Representation{RepType: acctest.Optional, Create: `${lookup(data.oci_kms_key_versions.test_key_versions.key_versions[1], "id")}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "cloud_autonomous_vm_cluster_id"),
				resource.TestCheckResourceAttr(resourceName, "backup_config.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "backup_config.0.recovery_window_in_days", "11"),
				resource.TestCheckResourceAttr(resourceName, "backup_config.0.backup_destination_details.0.is_retention_lock_enabled", "true"),
				resource.TestCheckResourceAttr(resourceName, "backup_config.0.backup_destination_details.0.backup_retention_policy_on_terminate", "RETAIN_PER_RETENTION_WINDOW"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "db_split_threshold", "12"),
				resource.TestCheckResourceAttr(resourceName, "distribution_affinity", "MINIMUM_DISTRIBUTION"),
				resource.TestCheckResourceAttrSet(resourceName, "db_version"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				//resource.TestCheckResourceAttr(resourceName, "is_automatic_failover_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "is_dst_file_update_enabled", "true"),
				resource.TestCheckResourceAttrSet(resourceName, "kms_key_id"),
				resource.TestCheckResourceAttr(resourceName, "net_services_architecture", "DEDICATED"),
				resource.TestCheckResourceAttr(resourceName, "patch_model", "RELEASE_UPDATE_REVISIONS"),
				resource.TestCheckResourceAttr(resourceName, "service_level_agreement_type", "STANDARD"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "vault_id"),
				resource.TestCheckResourceAttr(resourceName, "version_preference", "NEXT_RELEASE_UPDATE"),
				resource.TestCheckResourceAttr(resourceName, "vm_failover_reservation", "25"),
				resource.TestCheckResourceAttr(resourceName, "db_name", "DBNAME"),
				//resource.TestCheckResourceAttrSet(resourceName, "key_version_id"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("Resource recreated when it was supposed to be updated.")
					}
					return err
				},
			),
		},

		//verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_autonomous_container_databases", "test_autonomous_container_databases", acctest.Optional, acctest.Update, DatabaseDatabaseAutonomousContainerDatabaseDataSourceRepresentation) +
				compartmentIdVariableStr + ATPDAutonomousContainerDatabaseResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_container_database", "test_autonomous_container_database", acctest.Optional, acctest.Update,
					acctest.RepresentationCopyWithNewProperties(AutonomousContainerDatabaseDedicatedRepresentation, map[string]interface{}{
						"rotate_key_trigger": acctest.Representation{RepType: acctest.Optional, Create: `true`},
						"key_version_id":     acctest.Representation{RepType: acctest.Optional, Create: `${lookup(data.oci_kms_key_versions.test_key_versions.key_versions[1], "id")}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "cloud_autonomous_vm_cluster_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "availability_domain"),
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(datasourceName, "state", "AVAILABLE"),
				resource.TestCheckResourceAttr(datasourceName, "autonomous_container_databases.#", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "autonomous_container_databases.0.cloud_autonomous_vm_cluster_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "autonomous_container_databases.0.availability_domain"),
				resource.TestCheckResourceAttrSet(datasourceName, "autonomous_container_databases.0.available_cpus"),
				resource.TestCheckResourceAttr(datasourceName, "autonomous_container_databases.0.backup_config.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "autonomous_container_databases.0.backup_config.0.recovery_window_in_days", "11"),
				resource.TestCheckResourceAttr(datasourceName, "autonomous_container_databases.0.backup_config.0.backup_destination_details.0.is_retention_lock_enabled", "true"),
				resource.TestCheckResourceAttr(datasourceName, "autonomous_container_databases.0.backup_config.0.backup_destination_details.0.backup_retention_policy_on_terminate", "RETAIN_PER_RETENTION_WINDOW"),
				resource.TestCheckResourceAttr(datasourceName, "autonomous_container_databases.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(datasourceName, "autonomous_container_databases.0.compute_model"),
				resource.TestCheckResourceAttr(datasourceName, "autonomous_container_databases.0.customer_contacts.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "autonomous_container_databases.0.customer_contacts.0.email", "test2@oracle.com"),
				resource.TestCheckResourceAttr(datasourceName, "autonomous_container_databases.0.dataguard.#", "0"),
				resource.TestCheckResourceAttr(datasourceName, "autonomous_container_databases.0.dataguard_group_members.#", "0"),
				resource.TestCheckResourceAttr(datasourceName, "autonomous_container_databases.0.db_split_threshold", "12"),
				resource.TestCheckResourceAttrSet(datasourceName, "autonomous_container_databases.0.db_version"),
				resource.TestCheckResourceAttr(datasourceName, "autonomous_container_databases.0.display_name", "displayName2"),
				resource.TestCheckResourceAttr(datasourceName, "autonomous_container_databases.0.distribution_affinity", "MINIMUM_DISTRIBUTION"),
				resource.TestCheckResourceAttrSet(datasourceName, "autonomous_container_databases.0.dst_file_version"),
				resource.TestCheckResourceAttr(datasourceName, "autonomous_container_databases.0.freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "autonomous_container_databases.0.id"),
				resource.TestCheckResourceAttr(datasourceName, "autonomous_container_databases.0.is_dst_file_update_enabled", "true"),
				//resource.TestCheckResourceAttr(datasourceName, "autonomous_container_databases.0.key_history_entry.#", "1"),
				//resource.TestCheckResourceAttrSet(datasourceName, "autonomous_container_databases.0.key_store_wallet_name"),
				//resource.TestCheckResourceAttrSet(datasourceName, "autonomous_container_databases.0.kms_key_id"),
				//resource.TestCheckResourceAttrSet(datasourceName, "autonomous_container_databases.0.kms_key_version_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "autonomous_container_databases.0.largest_provisionable_autonomous_database_in_cpus"),
				//resource.TestCheckResourceAttrSet(datasourceName, "autonomous_container_databases.0.last_maintenance_run_id"),
				//resource.TestCheckResourceAttrSet(datasourceName, "autonomous_container_databases.0.last_maintenance_run_id"),
				resource.TestCheckResourceAttr(datasourceName, "autonomous_container_databases.0.maintenance_window.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "autonomous_container_databases.0.net_services_architecture", "DEDICATED"),
				//resource.TestCheckResourceAttrSet(datasourceName, "autonomous_container_databases.0.patch_id"),
				resource.TestCheckResourceAttr(datasourceName, "autonomous_container_databases.0.maintenance_window.0.days_of_week.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "autonomous_container_databases.0.maintenance_window.0.days_of_week.0.name", "TUESDAY"),
				resource.TestCheckResourceAttr(datasourceName, "autonomous_container_databases.0.maintenance_window.0.hours_of_day.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "autonomous_container_databases.0.maintenance_window.0.months.#", "4"),
				resource.TestCheckResourceAttr(datasourceName, "autonomous_container_databases.0.maintenance_window.0.months.0.name", "FEBRUARY"),
				resource.TestCheckResourceAttr(datasourceName, "autonomous_container_databases.0.maintenance_window.0.preference", "CUSTOM_PREFERENCE"),
				resource.TestCheckResourceAttr(datasourceName, "autonomous_container_databases.0.maintenance_window.0.weeks_of_month.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "autonomous_container_databases.0.patch_model", "RELEASE_UPDATE_REVISIONS"),
				resource.TestCheckResourceAttrSet(datasourceName, "autonomous_container_databases.0.provisioned_cpus"),
				resource.TestCheckResourceAttrSet(datasourceName, "autonomous_container_databases.0.reclaimable_cpus"),
				resource.TestCheckResourceAttrSet(datasourceName, "autonomous_container_databases.0.reserved_cpus"),
				//resource.TestCheckResourceAttrSet(datasourceName, "autonomous_container_databases.0.role"),
				resource.TestCheckResourceAttr(datasourceName, "autonomous_container_databases.0.provisionable_cpus.#", "209"),
				resource.TestCheckResourceAttrSet(datasourceName, "autonomous_container_databases.0.reclaimable_cpus"),
				resource.TestCheckResourceAttr(datasourceName, "autonomous_container_databases.0.service_level_agreement_type", "STANDARD"),
				resource.TestCheckResourceAttrSet(datasourceName, "autonomous_container_databases.0.state"),
				resource.TestCheckResourceAttrSet(datasourceName, "autonomous_container_databases.0.time_created"),
				//resource.TestCheckResourceAttrSet(datasourceName, "autonomous_container_databases.0.time_of_last_backup"),
				//resource.TestCheckResourceAttrSet(datasourceName, "autonomous_container_databases.0.time_snapshot_standby_revert"),
				resource.TestCheckResourceAttrSet(datasourceName, "autonomous_container_databases.0.total_cpus"),
				resource.TestCheckResourceAttrSet(datasourceName, "autonomous_container_databases.0.vault_id"),
				resource.TestCheckResourceAttr(datasourceName, "autonomous_container_databases.0.version_preference", "NEXT_RELEASE_UPDATE"),
				resource.TestCheckResourceAttr(datasourceName, "autonomous_container_databases.0.vm_failover_reservation", "25"),
				resource.TestCheckResourceAttr(datasourceName, "autonomous_container_databases.0.db_name", "DBNAME"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_autonomous_container_database", "test_autonomous_container_database", acctest.Required, acctest.Create, DatabaseDatabaseAutonomousContainerDatabaseSingularDataSourceRepresentation) +
				compartmentIdVariableStr + ATPDAutonomousContainerDatabaseResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_container_database", "test_autonomous_container_database", acctest.Optional, acctest.Update, acctest.RepresentationCopyWithNewProperties(AutonomousContainerDatabaseDedicatedRepresentation, map[string]interface{}{
					"rotate_key_trigger": acctest.Representation{RepType: acctest.Optional, Create: `true`},
					"key_version_id":     acctest.Representation{RepType: acctest.Optional, Create: `${lookup(data.oci_kms_key_versions.test_key_versions.key_versions[1], "id")}`},
				})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "autonomous_container_database_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "availability_domain"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "available_cpus"),
				resource.TestCheckResourceAttr(singularDatasourceName, "backup_config.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "backup_config.0.recovery_window_in_days", "11"),
				resource.TestCheckResourceAttr(singularDatasourceName, "backup_config.0.backup_destination_details.0.is_retention_lock_enabled", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "backup_config.0.backup_destination_details.0.backup_retention_policy_on_terminate", "RETAIN_PER_RETENTION_WINDOW"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "compute_model"),
				resource.TestCheckResourceAttr(singularDatasourceName, "customer_contacts.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "customer_contacts.0.email", "test2@oracle.com"),
				resource.TestCheckResourceAttr(singularDatasourceName, "dataguard.#", "0"),
				resource.TestCheckResourceAttr(singularDatasourceName, "dataguard_group_members.#", "0"),
				resource.TestCheckResourceAttr(singularDatasourceName, "db_split_threshold", "12"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "db_version"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "distribution_affinity", "MINIMUM_DISTRIBUTION"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "dst_file_version"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "is_dst_file_update_enabled", "true"),
				//resource.TestCheckResourceAttr(singularDatasourceName, "key_history_entry.#", "1"),
				//resource.TestCheckResourceAttrSet(singularDatasourceName, "key_store_wallet_name"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "largest_provisionable_autonomous_database_in_cpus"),
				//resource.TestCheckResourceAttrSet(singularDatasourceName, "last_maintenance_run_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "maintenance_window.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "maintenance_window.0.days_of_week.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "maintenance_window.0.days_of_week.0.name", "TUESDAY"),
				resource.TestCheckResourceAttr(singularDatasourceName, "maintenance_window.0.hours_of_day.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "maintenance_window.0.months.#", "4"),
				resource.TestCheckResourceAttr(singularDatasourceName, "maintenance_window.0.months.0.name", "FEBRUARY"),
				resource.TestCheckResourceAttr(singularDatasourceName, "maintenance_window.0.preference", "CUSTOM_PREFERENCE"),
				resource.TestCheckResourceAttr(singularDatasourceName, "maintenance_window.0.weeks_of_month.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "net_services_architecture", "DEDICATED"),
				//resource.TestCheckResourceAttrSet(singularDatasourceName, "patch_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "patch_model", "RELEASE_UPDATE_REVISIONS"),
				//resource.TestCheckResourceAttrSet(singularDatasourceName, "provisioned_cpus"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "reclaimable_cpus"),
				//resource.TestCheckResourceAttrSet(singularDatasourceName, "reserved_cpus"),
				//resource.TestCheckResourceAttrSet(singularDatasourceName, "role"),
				resource.TestCheckResourceAttr(singularDatasourceName, "provisionable_cpus.#", "209"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "reclaimable_cpus"),
				resource.TestCheckResourceAttr(singularDatasourceName, "service_level_agreement_type", "STANDARD"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				//resource.TestCheckResourceAttrSet(singularDatasourceName, "time_of_last_backup"),
				//resource.TestCheckResourceAttrSet(singularDatasourceName, "time_snapshot_standby_revert"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "total_cpus"),
				resource.TestCheckResourceAttr(singularDatasourceName, "version_preference", "NEXT_RELEASE_UPDATE"),
				resource.TestCheckResourceAttr(singularDatasourceName, "vm_failover_reservation", "25"),
				resource.TestCheckResourceAttr(singularDatasourceName, "db_name", "DBNAME"),
			),
		},

		{
			Config: config + compartmentIdVariableStr + ATPDAutonomousContainerDatabaseResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_container_database", "test_autonomous_container_database", acctest.Optional, acctest.Update,
					acctest.GetUpdatedRepresentationCopy("maintenance_window_details", acctest.RepresentationGroup{RepType: acctest.Optional, Group: DatabaseAutonomousContainerDatabaseMaintenanceWindowDetailsNoPreferenceRepresentation}, acctest.RepresentationCopyWithNewProperties(AutonomousContainerDatabaseDedicatedRepresentation, map[string]interface{}{
						"rotate_key_trigger": acctest.Representation{RepType: acctest.Optional, Create: `true`},
						"key_version_id":     acctest.Representation{RepType: acctest.Optional, Create: `${lookup(data.oci_kms_key_versions.test_key_versions.key_versions[1], "id")}`},
					}))),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "cloud_autonomous_vm_cluster_id"),
				resource.TestCheckResourceAttr(resourceName, "backup_config.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "backup_config.0.recovery_window_in_days", "11"),
				resource.TestCheckResourceAttr(resourceName, "backup_config.0.backup_destination_details.0.is_retention_lock_enabled", "true"),
				resource.TestCheckResourceAttr(resourceName, "backup_config.0.backup_destination_details.0.backup_retention_policy_on_terminate", "RETAIN_PER_RETENTION_WINDOW"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window.0.preference", "NO_PREFERENCE"),
				resource.TestCheckResourceAttr(resourceName, "patch_model", "RELEASE_UPDATE_REVISIONS"),
				resource.TestCheckResourceAttr(resourceName, "service_level_agreement_type", "STANDARD"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("Resource recreated when it was supposed to be updated.")
					}
					return err
				},
			),
		},

		// verify resource import
		{
			Config:            config + DatabaseAutonomousContainerDatabaseRequiredOnlyResource + compartmentIdVariableStr,
			ImportState:       true,
			ImportStateVerify: true,
			ImportStateVerifyIgnore: []string{
				//"autonomous_container_database_backup_id",
				"database_software_image_id",
				"is_automatic_failover_enabled",
				"rotate_key_trigger",
				"key_version_id",
				"maintenance_window_details",
				"peer_autonomous_container_database_backup_config",
				"peer_autonomous_container_database_compartment_id",
				"peer_autonomous_vm_cluster_id",
				"peer_autonomous_container_database_display_name",
				"peer_autonomous_exadata_infrastructure_id",
				"peer_db_unique_name",
				"protection_mode",
				"source",
				"lifecycle_details",
				"time_of_last_backup",
			},
			ResourceName: resourceName,
		},
	})
}

func TestDatabaseAutonomousContainerDatabaseResource_OkvEpg(t *testing.T) {
	//t.Skip("Skip this test as AEI and its api no longer exists.")

	shouldSkipADBDtest := os.Getenv("TF_VAR_should_skip_adbd_test")

	if shouldSkipADBDtest == "true" {
		t.Skip("Skipping TestDatabaseAutonomousContainerDatabaseResource_OkvEpg test.\n" + "Current TF_VAR_should_skip_adbd_test=" + shouldSkipADBDtest)
	}

	httpreplay.SetScenario("TestDatabaseAutonomousContainerDatabaseResource_OkvEpg")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_database_autonomous_container_database.test_autonomous_container_database"

	AutonomousContainerDatabaseDedicatedMaintenanceWindowDetailsRepresentation := acctest.RepresentationCopyWithRemovedProperties(
		acctest.GetUpdatedRepresentationCopy("months",
			[]acctest.RepresentationGroup{{RepType: acctest.Optional, Group: DatabaseAutonomousContainerDatabaseMaintenanceWindowDetailsMonthsRepresentation}, {RepType: acctest.Optional, Group: DatabaseAutonomousContainerDatabaseMaintenanceWindowDetailsMonthsRepresentation2}, {RepType: acctest.Optional, Group: DatabaseAutonomousContainerDatabaseMaintenanceWindowDetailsMonthsRepresentation3}, {RepType: acctest.Optional, Group: DatabaseAutonomousContainerDatabaseMaintenanceWindowDetailsMonthsRepresentation4}},
			DatabaseAutonomousContainerDatabaseMaintenanceWindowDetailsRepresentation), []string{"lead_time_in_weeks"})

	AutonomousContainerDatabaseDedicatedRepresentation := acctest.GetUpdatedRepresentationCopy("maintenance_window_details", acctest.RepresentationGroup{RepType: acctest.Optional, Group: AutonomousContainerDatabaseDedicatedMaintenanceWindowDetailsRepresentation}, DatabaseAutonomousContainerDatabaseRepresentationWithOkvEpg)

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+ATPDAutonomousContainerDatabaseResourceWithOkvDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_container_database", "test_autonomous_container_database", acctest.Optional, acctest.Create, DatabaseAutonomousContainerDatabaseRepresentationWithOkvEpg), "database", "autonomousContainerDatabase", t)

	acctest.ResourceTest(t, testAccCheckDatabaseAutonomousContainerDatabaseDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + ATPDAutonomousContainerDatabaseResourceWithOkvDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_container_database", "test_autonomous_container_database", acctest.Required, acctest.Create, DatabaseAutonomousContainerDatabaseRepresentationWithOkvEpg),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				//resource.TestCheckResourceAttrSet(resourceName, "autonomous_container_database_backup_id"),
				resource.TestCheckResourceAttrSet(resourceName, "cloud_autonomous_vm_cluster_id"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "containerDatabase2"),
				resource.TestCheckResourceAttr(resourceName, "patch_model", "RELEASE_UPDATES"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next create
		{
			Config: config + compartmentIdVariableStr + ATPDAutonomousContainerDatabaseResourceWithOkvDependencies,
		},
		// verify create with optionals
		{
			Config: config + compartmentIdVariableStr + ATPDAutonomousContainerDatabaseResourceWithOkvDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_container_database", "test_autonomous_container_database", acctest.Optional, acctest.Create,
					acctest.GetUpdatedRepresentationCopy("maintenance_window_details", acctest.RepresentationGroup{RepType: acctest.Optional, Group: DatabaseAutonomousContainerDatabaseMaintenanceWindowDetailsNoPreferenceRepresentation}, DatabaseAutonomousContainerDatabaseRepresentationWithOkvEpg)),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				//resource.TestCheckResourceAttrSet(resourceName, "autonomous_container_database_backup_id"),
				//resource.TestCheckResourceAttrSet(resourceName, "autonomous_exadata_infrastructure_id"),
				resource.TestCheckResourceAttr(resourceName, "backup_config.#", "1"),
				//resource.TestCheckResourceAttr(resourceName, "backup_config.0.recovery_window_in_days", "7"),
				resource.TestCheckResourceAttrSet(resourceName, "cloud_autonomous_vm_cluster_id"),
				resource.TestCheckResourceAttr(resourceName, "backup_config.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "backup_config.0.recovery_window_in_days", "10"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "db_split_threshold", "12"),
				resource.TestCheckResourceAttr(resourceName, "distribution_affinity", "MINIMUM_DISTRIBUTION"),
				resource.TestCheckResourceAttrSet(resourceName, "db_version"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "containerDatabase2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_dst_file_update_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "net_services_architecture", "DEDICATED"),
				resource.TestCheckResourceAttr(resourceName, "okv_end_point_group_name", "DUMMY_OKV_EPG_GROUP"),
				resource.TestCheckResourceAttr(resourceName, "patch_model", "RELEASE_UPDATES"),
				resource.TestCheckResourceAttr(resourceName, "source", "NONE"),
				resource.TestCheckResourceAttr(resourceName, "service_level_agreement_type", "STANDARD"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttr(resourceName, "version_preference", "LATEST_RELEASE_UPDATE"),
				resource.TestCheckResourceAttr(resourceName, "vm_failover_reservation", "25"),
				resource.TestCheckResourceAttr(resourceName, "db_name", "DBNAME"),
				resource.TestCheckResourceAttr(resourceName, "system_tags.%", "0"),

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

		// verify updates to okv endpoint group name
		{
			Config: config + compartmentIdVariableStr + ATPDAutonomousContainerDatabaseResourceWithOkvDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_container_database", "test_autonomous_container_database", acctest.Optional, acctest.Update, AutonomousContainerDatabaseDedicatedRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "cloud_autonomous_vm_cluster_id"),
				resource.TestCheckResourceAttr(resourceName, "backup_config.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "backup_config.0.recovery_window_in_days", "11"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "db_split_threshold", "12"),
				resource.TestCheckResourceAttr(resourceName, "distribution_affinity", "MINIMUM_DISTRIBUTION"),
				resource.TestCheckResourceAttrSet(resourceName, "db_version"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_dst_file_update_enabled", "true"),
				resource.TestCheckResourceAttr(resourceName, "net_services_architecture", "DEDICATED"),
				resource.TestCheckResourceAttr(resourceName, "patch_model", "RELEASE_UPDATE_REVISIONS"),
				resource.TestCheckResourceAttr(resourceName, "source", "NONE"),
				resource.TestCheckResourceAttr(resourceName, "service_level_agreement_type", "STANDARD"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttr(resourceName, "version_preference", "NEXT_RELEASE_UPDATE"),
				resource.TestCheckResourceAttr(resourceName, "vm_failover_reservation", "25"),
				resource.TestCheckResourceAttr(resourceName, "db_name", "DBNAME"),
				resource.TestCheckResourceAttr(resourceName, "okv_end_point_group_name", "DUMMY_OKV_EPG_GROUP_2"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("Resource recreated when it was supposed to be updated.")
					}
					return err
				},
			),
		},
	})
}

func testAccCheckDatabaseAutonomousContainerDatabaseDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.GetTestClients(&schema.ResourceData{}).DatabaseClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_database_autonomous_container_database" {
			noResourceFound = false
			request := oci_database.GetAutonomousContainerDatabaseRequest{}

			tmp := rs.Primary.ID
			request.AutonomousContainerDatabaseId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "database")

			response, err := client.GetAutonomousContainerDatabase(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_database.AutonomousContainerDatabaseLifecycleStateTerminated): true,
				}
				if _, ok := deletedLifecycleStates[string(response.LifecycleState)]; !ok {
					//resource lifecycle state is not in expected deleted lifecycle states.
					return fmt.Errorf("resource lifecycle state: %s is not in expected deleted lifecycle states", response.LifecycleState)
				}
				//resource lifecycle state is in expected deleted lifecycle states. continue with next one.
				continue
			}

			//Verify that exception is for '404 not found'.
			if failure, isServiceError := common.IsServiceError(err); !isServiceError || failure.GetHTTPStatusCode() != 404 {
				return err
			}
		}
	}
	if noResourceFound {
		return fmt.Errorf("at least one resource was expected from the state file, but could not be found")
	}

	return nil
}

func TestDatabaseAutonomousContainerDatabaseResource_rotateKey(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseAutonomousContainerDatabaseResource_rotateKey")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()
	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_database_autonomous_container_database.test_autonomous_container_database"

	AutonomousContainerDatabaseDedicatedMaintenanceWindowDetailsRepresentation := acctest.RepresentationCopyWithRemovedProperties(
		acctest.GetUpdatedRepresentationCopy("months",
			[]acctest.RepresentationGroup{{RepType: acctest.Optional, Group: DatabaseAutonomousContainerDatabaseMaintenanceWindowDetailsMonthsRepresentation}, {RepType: acctest.Optional, Group: DatabaseAutonomousContainerDatabaseMaintenanceWindowDetailsMonthsRepresentation2}, {RepType: acctest.Optional, Group: DatabaseAutonomousContainerDatabaseMaintenanceWindowDetailsMonthsRepresentation3}, {RepType: acctest.Optional, Group: DatabaseAutonomousContainerDatabaseMaintenanceWindowDetailsMonthsRepresentation4}},
			DatabaseAutonomousContainerDatabaseMaintenanceWindowDetailsRepresentation), []string{"lead_time_in_weeks"})

	AutonomousContainerDatabaseDedicatedRepresentation := acctest.GetUpdatedRepresentationCopy("maintenance_window_details", acctest.RepresentationGroup{RepType: acctest.Optional, Group: AutonomousContainerDatabaseDedicatedMaintenanceWindowDetailsRepresentation}, DatabaseAutonomousContainerDatabaseRepresentation)

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+ATPDAutonomousContainerDatabaseResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_container_database", "test_autonomous_container_database", acctest.Optional, acctest.Create, DatabaseAutonomousContainerDatabaseRepresentation), "database", "autonomousContainerDatabase", t)

	acctest.ResourceTest(t, testAccCheckDatabaseAutonomousContainerDatabaseDestroy, []resource.TestStep{
		// verify create with optionals
		{
			Config: config + compartmentIdVariableStr + ATPDAutonomousContainerDatabaseResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_container_database", "test_autonomous_container_database", acctest.Optional, acctest.Create,
					acctest.GetUpdatedRepresentationCopy("maintenance_window_details", acctest.RepresentationGroup{RepType: acctest.Optional, Group: DatabaseAutonomousContainerDatabaseMaintenanceWindowDetailsNoPreferenceRepresentation}, DatabaseAutonomousContainerDatabaseRepresentation)),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(

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
		//verify rotate key with key_version_id
		{
			Config: config + compartmentIdVariableStr + ATPDAutonomousContainerDatabaseResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_container_database", "test_autonomous_container_database", acctest.Optional, acctest.Update,
					acctest.RepresentationCopyWithNewProperties(AutonomousContainerDatabaseDedicatedRepresentation, map[string]interface{}{
						"rotate_key_trigger": acctest.Representation{RepType: acctest.Optional, Create: `true`},
						"key_version_id":     acctest.Representation{RepType: acctest.Optional, Create: utils.GetEnvSettingWithBlankDefault("acd_key_version_id")},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "kms_key_version_id", utils.GetEnvSettingWithBlankDefault("acd_key_version_id")),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("Resource recreated when it was supposed to be updated.")
					}
					return err
				},
			),
		},
		//verify rotate key with key_version_id
		{
			Config: config + compartmentIdVariableStr + ATPDAutonomousContainerDatabaseResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_container_database", "test_autonomous_container_database", acctest.Optional, acctest.Update,
					acctest.RepresentationCopyWithNewProperties(AutonomousContainerDatabaseDedicatedRepresentation, map[string]interface{}{
						"rotate_key_trigger": acctest.Representation{RepType: acctest.Optional, Create: `true`},
						"key_version_id":     acctest.Representation{RepType: acctest.Optional, Create: utils.GetEnvSettingWithBlankDefault("acd_key_version_id_2")},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "kms_key_version_id", utils.GetEnvSettingWithBlankDefault("acd_key_version_id_2")),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("Resource recreated when it was supposed to be updated.")
					}
					return err
				},
			),
		},
		// unset rotate_key_trigger
		{
			Config: config + compartmentIdVariableStr + ATPDAutonomousContainerDatabaseResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_container_database", "test_autonomous_container_database", acctest.Optional, acctest.Update,
					acctest.RepresentationCopyWithNewProperties(AutonomousContainerDatabaseDedicatedRepresentation, map[string]interface{}{
						"rotate_key_trigger": acctest.Representation{RepType: acctest.Optional, Create: `false`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("Resource recreated when it was supposed to be updated.")
					}
					return err
				},
			),
		},
		//verify rotate key without key_version_id
		{
			Config: config + compartmentIdVariableStr + ATPDAutonomousContainerDatabaseResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_container_database", "test_autonomous_container_database", acctest.Optional, acctest.Update,
					acctest.RepresentationCopyWithNewProperties(AutonomousContainerDatabaseDedicatedRepresentation, map[string]interface{}{
						"rotate_key_trigger": acctest.Representation{RepType: acctest.Optional, Create: `true`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "kms_key_version_id", utils.GetEnvSettingWithBlankDefault("acd_key_version_id_2")),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("Resource recreated when it was supposed to be updated.")
					}
					return err
				},
			),
		},
	})
}

func init() {
	if acctest.DependencyGraph == nil {
		acctest.InitDependencyGraph()
	}
	if !acctest.InSweeperExcludeList("DatabaseAutonomousContainerDatabase") {
		resource.AddTestSweepers("DatabaseAutonomousContainerDatabase", &resource.Sweeper{
			Name:         "DatabaseAutonomousContainerDatabase",
			Dependencies: acctest.DependencyGraph["autonomousContainerDatabase"],
			F:            sweepDatabaseAutonomousContainerDatabaseResource,
		})
	}
}

func sweepDatabaseAutonomousContainerDatabaseResource(compartment string) error {
	databaseClient := acctest.GetTestClients(&schema.ResourceData{}).DatabaseClient()
	autonomousContainerDatabaseIds, err := getDatabaseAutonomousContainerDatabaseIds(compartment)
	if err != nil {
		return err
	}
	for _, autonomousContainerDatabaseId := range autonomousContainerDatabaseIds {
		if ok := acctest.SweeperDefaultResourceId[autonomousContainerDatabaseId]; !ok {
			terminateAutonomousContainerDatabaseRequest := oci_database.TerminateAutonomousContainerDatabaseRequest{}

			terminateAutonomousContainerDatabaseRequest.AutonomousContainerDatabaseId = &autonomousContainerDatabaseId

			terminateAutonomousContainerDatabaseRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "database")
			_, error := databaseClient.TerminateAutonomousContainerDatabase(context.Background(), terminateAutonomousContainerDatabaseRequest)
			if error != nil {
				fmt.Printf("Error deleting AutonomousContainerDatabase %s %s, It is possible that the resource is already deleted. Please verify manually \n", autonomousContainerDatabaseId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &autonomousContainerDatabaseId, DatabaseAutonomousContainerDatabaseSweepWaitCondition, time.Duration(3*time.Minute),
				DatabaseAutonomousContainerDatabaseSweepResponseFetchOperation, "database", true)
		}
	}
	return nil
}

func getDatabaseAutonomousContainerDatabaseIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "AutonomousContainerDatabaseId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	databaseClient := acctest.GetTestClients(&schema.ResourceData{}).DatabaseClient()

	listAutonomousContainerDatabasesRequest := oci_database.ListAutonomousContainerDatabasesRequest{}
	listAutonomousContainerDatabasesRequest.CompartmentId = &compartmentId
	listAutonomousContainerDatabasesRequest.LifecycleState = oci_database.AutonomousContainerDatabaseSummaryLifecycleStateAvailable
	listAutonomousContainerDatabasesResponse, err := databaseClient.ListAutonomousContainerDatabases(context.Background(), listAutonomousContainerDatabasesRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting AutonomousContainerDatabase list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, autonomousContainerDatabase := range listAutonomousContainerDatabasesResponse.Items {
		id := *autonomousContainerDatabase.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "AutonomousContainerDatabaseId", id)
	}
	return resourceIds, nil
}

func DatabaseAutonomousContainerDatabaseSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if autonomousContainerDatabaseResponse, ok := response.Response.(oci_database.GetAutonomousContainerDatabaseResponse); ok {
		return autonomousContainerDatabaseResponse.LifecycleState != oci_database.AutonomousContainerDatabaseLifecycleStateTerminated
	}
	return false
}

func DatabaseAutonomousContainerDatabaseSweepResponseFetchOperation(client *client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.DatabaseClient().GetAutonomousContainerDatabase(context.Background(), oci_database.GetAutonomousContainerDatabaseRequest{
		AutonomousContainerDatabaseId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}

func getStandbyAcdOcid(resourceName string) resource.ImportStateIdFunc {
	return func(s *terraform.State) (string, error) {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return "", fmt.Errorf("not found: %s", resourceName)
		}
		return fmt.Sprintf(rs.Primary.Attributes["dataguard_group_members.1.autonomous_container_database_id"]), nil
	}
}

func getStandbyAcdOcidOldDG(resourceName string) resource.ImportStateIdFunc {
	return func(s *terraform.State) (string, error) {
		fmt.Printf("MyState: %s", s.RootModule().Resources)
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return "", fmt.Errorf("not found: %s", resourceName)
		}
		return fmt.Sprintf(rs.Primary.Attributes["autonomous_container_database_dataguard_associations.0.peer_autonomous_container_database_id"]), nil
	}
}
