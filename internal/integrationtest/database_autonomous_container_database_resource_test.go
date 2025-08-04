// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.

package integrationtest

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/oracle/terraform-provider-oci/internal/acctest"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/hashicorp/terraform-plugin-testing/terraform"

	"github.com/oracle/terraform-provider-oci/httpreplay"
)

var (
	acbDBName  = utils.RandomString(1, utils.CharsetWithoutDigits) + utils.RandomString(13, utils.Charset)
	acbDBName2 = utils.RandomString(1, utils.CharsetWithoutDigits) + utils.RandomString(13, utils.Charset)

	ExaccACDResourceConfig = ACDECPUatabaseResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_container_database", "test_autonomous_container_database", acctest.Optional, acctest.Update, ACDatabaseRepresentation)

	// This is to be able to delete the LTB after the test run
	ExaccACDResourceConfigWithoutRetentionLock = ACDECPUatabaseResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_container_database", "test_autonomous_container_database", acctest.Optional, acctest.Create, ACDatabaseRepresentation)

	ExaccACDRequiredOnlyResource = ExaccDatabaseAutonomousContainerDatabaseResourceFromAdsiDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_container_database", "test_autonomous_container_database_from_adsi", acctest.Required, acctest.Create, ExaccACDatabaseFromAdsiRepresentation)

	ACDRequiredOnlyResource = DatabaseAutonomousContainerDatabaseResourceFromAdsiDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_container_database", "test_autonomous_container_database_from_adsi", acctest.Required, acctest.Create, ACDatabaseFromAdsiRepresentation)

	ACDatabaseDataSourceRepresentation = map[string]interface{}{
		"compartment_id":           acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"autonomous_vm_cluster_id": acctest.Representation{RepType: acctest.Optional, Create: `${oci_database_autonomous_vm_cluster.test_autonomous_vm_cluster.id}`},
		"display_name":             acctest.Representation{RepType: acctest.Optional, Create: `containerdatabases2`},
		"infrastructure_type":      acctest.Representation{RepType: acctest.Optional, Create: `CLOUD_AT_CUSTOMER`},
		"state":                    acctest.Representation{RepType: acctest.Optional, Create: `AVAILABLE`},
		"filter":                   acctest.RepresentationGroup{RepType: acctest.Required, Group: DatabaseAutonomousContainerDatabaseDataSourceFilterRepresentation},
	}

	ACDatabaseRepresentation = map[string]interface{}{
		"customer_contacts":            acctest.RepresentationGroup{RepType: acctest.Optional, Group: DatabaseAutonomousContainerDatabaseCustomerContactsRepresentation},
		"db_split_threshold":           acctest.Representation{RepType: acctest.Optional, Create: `8`},
		"okv_end_point_group_name":     acctest.Representation{RepType: acctest.Optional, Create: `DUMMY_OKV_EPG_GROUP`, Update: `DUMMY_OKV_EPG_GROUP_2`},
		"distribution_affinity":        acctest.Representation{RepType: acctest.Optional, Create: `MINIMUM_DISTRIBUTION`},
		"net_services_architecture":    acctest.Representation{RepType: acctest.Optional, Create: `DEDICATED`},
		"vm_failover_reservation":      acctest.Representation{RepType: acctest.Optional, Create: `25`},
		"version_preference":           acctest.Representation{RepType: acctest.Optional, Create: `LATEST_RELEASE_UPDATE`, Update: `NEXT_RELEASE_UPDATE`},
		"display_name":                 acctest.Representation{RepType: acctest.Required, Create: `containerdatabases2`},
		"patch_model":                  acctest.Representation{RepType: acctest.Required, Create: `RELEASE_UPDATES`, Update: `RELEASE_UPDATE_REVISIONS`},
		"autonomous_vm_cluster_id":     acctest.Representation{RepType: acctest.Required, Create: `${oci_database_autonomous_vm_cluster.test_autonomous_vm_cluster.id}`},
		"backup_config":                acctest.RepresentationGroup{RepType: acctest.Required, Group: DatabaseAutonomousContainerDatabaseBackupConfigRepresentation},
		"key_store_id":                 acctest.Representation{RepType: acctest.Optional, Create: `${oci_database_key_store.test_key_store.id}`},
		"compartment_id":               acctest.Representation{RepType: acctest.Optional, Create: `${var.compartment_id}`},
		"db_unique_name":               acctest.Representation{RepType: acctest.Optional, Create: acbDBName},
		"defined_tags":                 acctest.Representation{RepType: acctest.Optional, Create: `${tomap({"${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}" = "value"})}`, Update: `${tomap({"${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}" = "updatedValue"})}`},
		"freeform_tags":                acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"maintenance_window_details":   acctest.RepresentationGroup{RepType: acctest.Optional, Group: DatabaseAutonomousContainerDatabaseMaintenanceWindowDetailsRepresentation},
		"service_level_agreement_type": acctest.Representation{RepType: acctest.Optional, Create: `STANDARD`},
		"db_name":                      acctest.Representation{RepType: acctest.Optional, Create: `DBNAME`},
		"db_version":                   acctest.Representation{RepType: acctest.Required, Create: utils.GetEnvSettingWithDefault("exacc_acd_db_version", "19.27.0.1.0")},
		"is_dst_file_update_enabled":   acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"lifecycle":                    acctest.RepresentationGroup{RepType: acctest.Required, Group: DbaasIgnoreDefinedTagsRepresentation},
	}
	ACDatabaseWithRABkpDesRepresentation = map[string]interface{}{
		"db_split_threshold":           acctest.Representation{RepType: acctest.Optional, Create: `8`},
		"distribution_affinity":        acctest.Representation{RepType: acctest.Optional, Create: `MINIMUM_DISTRIBUTION`},
		"net_services_architecture":    acctest.Representation{RepType: acctest.Optional, Create: `DEDICATED`},
		"vm_failover_reservation":      acctest.Representation{RepType: acctest.Optional, Create: `25`},
		"version_preference":           acctest.Representation{RepType: acctest.Optional, Create: `LATEST_RELEASE_UPDATE`, Update: `NEXT_RELEASE_UPDATE`},
		"display_name":                 acctest.Representation{RepType: acctest.Required, Create: `containerdatabases2`},
		"patch_model":                  acctest.Representation{RepType: acctest.Required, Create: `RELEASE_UPDATES`, Update: `RELEASE_UPDATE_REVISIONS`},
		"autonomous_vm_cluster_id":     acctest.Representation{RepType: acctest.Required, Create: `${oci_database_autonomous_vm_cluster.test_autonomous_vm_cluster.id}`},
		"backup_config":                acctest.RepresentationGroup{RepType: acctest.Required, Group: DatabaseAutonomousContainerDatabaseBackupConfigWithRAUpdateRepresentation},
		"key_store_id":                 acctest.Representation{RepType: acctest.Optional, Create: `${oci_database_key_store.test_key_store.id}`},
		"compartment_id":               acctest.Representation{RepType: acctest.Optional, Create: `${var.compartment_id}`},
		"db_unique_name":               acctest.Representation{RepType: acctest.Optional, Create: acbDBName},
		"defined_tags":                 acctest.Representation{RepType: acctest.Optional, Create: `${tomap({"${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}" = "value"})}`, Update: `${tomap({"${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}" = "updatedValue"})}`},
		"freeform_tags":                acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"maintenance_window_details":   acctest.RepresentationGroup{RepType: acctest.Optional, Group: DatabaseAutonomousContainerDatabaseMaintenanceWindowDetailsRepresentation},
		"service_level_agreement_type": acctest.Representation{RepType: acctest.Optional, Create: `STANDARD`},
		"db_name":                      acctest.Representation{RepType: acctest.Optional, Create: `DBNAME`},
		"db_version":                   acctest.Representation{RepType: acctest.Required, Create: utils.GetEnvSettingWithDefault("exacc_acd_db_version", "19.27.0.1.0")},
		"is_dst_file_update_enabled":   acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
	}

	AddStandbyACDatabaseRepresentation = map[string]interface{}{
		"version_preference":            acctest.Representation{RepType: acctest.Optional, Create: `LATEST_RELEASE_UPDATE`, Update: `NEXT_RELEASE_UPDATE`},
		"display_name":                  acctest.Representation{RepType: acctest.Required, Create: `containerDatabase2`, Update: `displayName2`},
		"patch_model":                   acctest.Representation{RepType: acctest.Required, Create: `RELEASE_UPDATES`, Update: `RELEASE_UPDATE_REVISIONS`},
		"autonomous_vm_cluster_id":      acctest.Representation{RepType: acctest.Required, Create: `${oci_database_autonomous_vm_cluster.test_autonomous_vm_cluster.id}`},
		"backup_config":                 acctest.RepresentationGroup{RepType: acctest.Required, Group: DatabaseAddStandbyAutonomousContainerDatabaseBackupConfigRepresentation},
		"is_automatic_failover_enabled": acctest.Representation{RepType: acctest.Optional, Create: `false`},
		//"key_store_id":                 acctest.Representation{RepType: acctest.Optional, Create: `${oci_database_key_store.test_key_store.id}`},
		"compartment_id":               acctest.Representation{RepType: acctest.Optional, Create: `${var.compartment_id}`},
		"db_unique_name":               acctest.Representation{RepType: acctest.Optional, Create: acbDBName},
		"defined_tags":                 acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"freeform_tags":                acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"maintenance_window_details":   acctest.RepresentationGroup{RepType: acctest.Optional, Group: DatabaseAutonomousContainerDatabaseMaintenanceWindowDetailsRepresentation},
		"service_level_agreement_type": acctest.Representation{RepType: acctest.Optional, Create: `STANDARD`},
		"db_name":                      acctest.Representation{RepType: acctest.Optional, Create: `DBNAME`},
		"db_version":                   acctest.Representation{RepType: acctest.Required, Create: utils.GetEnvSettingWithDefault("exacc_acd_db_version", "19.27.0.1.0")},
		"is_dst_file_update_enabled":   acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
	}

	ExaccACDatabaseFromAdsiRepresentation = acctest.RepresentationCopyWithNewProperties(
		acctest.GetMultipleUpdatedRepresenationCopy(
			[]string{
				"db_version", "display_name", "db_unique_name", "db_name",
			}, []interface{}{
				acctest.Representation{RepType: acctest.Optional, Create: ``},
				acctest.Representation{RepType: acctest.Required, Create: `containerdatabases3`},
				acctest.Representation{RepType: acctest.Optional, Create: acbDBName2},
				acctest.Representation{RepType: acctest.Optional, Create: `DBNAME2`},
			}, ACDatabaseRepresentation),
		map[string]interface{}{"database_software_image_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_database_autonomous_database_software_image.test_autonomous_database_software_image.id}`}})

	ACDatabaseFromAdsiRepresentation = acctest.RepresentationCopyWithNewProperties(
		map[string]interface{}{
			"version_preference":             acctest.Representation{RepType: acctest.Optional, Create: `LATEST_RELEASE_UPDATE`, Update: `NEXT_RELEASE_UPDATE`},
			"display_name":                   acctest.Representation{RepType: acctest.Required, Create: `containerDatabase3`, Update: `displayName2`},
			"patch_model":                    acctest.Representation{RepType: acctest.Required, Create: `RELEASE_UPDATES`, Update: `RELEASE_UPDATE_REVISIONS`},
			"cloud_autonomous_vm_cluster_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_database_cloud_autonomous_vm_cluster.test_cloud_autonomous_vm_cluster.id}`},
			"backup_config":                  acctest.RepresentationGroup{RepType: acctest.Optional, Group: ACDatabaseBackupConfigRepresentation},
			"compartment_id":                 acctest.Representation{RepType: acctest.Optional, Create: `${var.compartment_id}`},
			"defined_tags":                   acctest.Representation{RepType: acctest.Optional, Create: `${tomap({"${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}" = "value"})}`, Update: `${tomap({"${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}" = "updatedValue"})}`},
			"freeform_tags":                  acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
			"is_automatic_failover_enabled":  acctest.Representation{RepType: acctest.Optional, Create: `false`},
			"kms_key_id":                     acctest.Representation{RepType: acctest.Optional, Create: `${lookup(data.oci_kms_keys.test_keys_dependency.keys[0], "id")}`},
			"maintenance_window_details":     acctest.RepresentationGroup{RepType: acctest.Optional, Group: DatabaseAutonomousContainerDatabaseMaintenanceWindowDetailsRepresentation},
			"service_level_agreement_type":   acctest.Representation{RepType: acctest.Optional, Create: `STANDARD`},
			"vault_id":                       acctest.Representation{RepType: acctest.Optional, Create: `${data.oci_kms_vault.test_vault.id}`},
			"db_name":                        acctest.Representation{RepType: acctest.Optional, Create: `DBNAME3`},
			"is_dst_file_update_enabled":     acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		},
		map[string]interface{}{"database_software_image_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_database_autonomous_database_software_image.test_autonomous_database_software_image.id}`}})

	ACDatabaseBackupConfigRepresentation = map[string]interface{}{
		"backup_destination_details": acctest.RepresentationGroup{RepType: acctest.Optional, Group: ACDatabaseBackupConfigRepresentationWithImmutable},
		"recovery_window_in_days":    acctest.Representation{RepType: acctest.Optional, Create: `10`, Update: `11`},
	}

	ACDatabaseBackupConfigRepresentationWithImmutable = map[string]interface{}{
		"type":                                 acctest.Representation{RepType: acctest.Required, Create: `OBJECT_STORE`},
		"backup_retention_policy_on_terminate": acctest.Representation{RepType: acctest.Optional, Create: `RETAIN_FOR_72_HOURS`, Update: `RETAIN_PER_RETENTION_WINDOW`},
		"is_retention_lock_enabled":            acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
	}

	autonomousContainerDatabaseBackupConfigBackupDestinationDetailsRepresentation = map[string]interface{}{
		"type":                                 acctest.Representation{RepType: acctest.Required, Create: `NFS`},
		"id":                                   acctest.Representation{RepType: acctest.Optional, Create: `${oci_database_backup_destination.test_backup_destination.id}`},
		"internet_proxy":                       acctest.Representation{RepType: acctest.Optional, Create: `internetProxy`},
		"vpc_password":                         acctest.Representation{RepType: acctest.Optional, Create: `vpcPassword`, Update: `vpcPassword2`},
		"vpc_user":                             acctest.Representation{RepType: acctest.Optional, Create: `bkupUser1`},
		"backup_retention_policy_on_terminate": acctest.Representation{RepType: acctest.Optional, Create: `RETAIN_FOR_72_HOURS`, Update: `RETAIN_PER_RETENTION_WINDOW`},
		"is_retention_lock_enabled":            acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
	}

	autonomousContainerDatabaseBackupConfigBackupDestinationDetailsRepresentation1 = map[string]interface{}{
		"type":                                 acctest.Representation{RepType: acctest.Required, Create: `NFS`},
		"id":                                   acctest.Representation{RepType: acctest.Optional, Create: `${oci_database_backup_destination.test_backup_destination.id}`},
		"internet_proxy":                       acctest.Representation{RepType: acctest.Optional, Create: `internetProxy`},
		"vpc_password":                         acctest.Representation{RepType: acctest.Optional, Create: `vpcPassword`, Update: `vpcPassword2`},
		"vpc_user":                             acctest.Representation{RepType: acctest.Optional, Create: `bkupUser1`},
		"backup_retention_policy_on_terminate": acctest.Representation{RepType: acctest.Optional, Create: `RETAIN_PER_RETENTION_WINDOW`, Update: `RETAIN_FOR_72_HOURS`},
		"is_retention_lock_enabled":            acctest.Representation{RepType: acctest.Optional, Create: `true`, Update: `false`},
	}

	autonomousContainerDatabaseBackupConfigBackupDestinationDetailsRepresentationWithNoUpdate = map[string]interface{}{
		"type":           acctest.Representation{RepType: acctest.Required, Create: `NFS`},
		"id":             acctest.Representation{RepType: acctest.Optional, Create: `${oci_database_backup_destination.test_backup_destination.id}`},
		"internet_proxy": acctest.Representation{RepType: acctest.Optional, Create: `internetProxy`},
		"vpc_password":   acctest.Representation{RepType: acctest.Optional, Create: `vpcPassword`},
		"vpc_user":       acctest.Representation{RepType: acctest.Optional, Create: `bkupUser1`},
	}
	autonomousContainerDatabaseBackupConfigBackupDestinationDetailsWithRAUpdateRepresentation = map[string]interface{}{
		"type":                                 acctest.Representation{RepType: acctest.Required, Create: `NFS`, Update: `RECOVERY_APPLIANCE`},
		"id":                                   acctest.Representation{RepType: acctest.Optional, Create: `${oci_database_backup_destination.test_backup_destination.id}`, Update: `${oci_database_backup_destination.ra_backup_destination.id}`},
		"internet_proxy":                       acctest.Representation{RepType: acctest.Optional, Create: `internetProxy`},
		"vpc_password":                         acctest.Representation{RepType: acctest.Optional, Create: `vpcPassword`, Update: `vpcPassword2`},
		"backup_retention_policy_on_terminate": acctest.Representation{RepType: acctest.Optional, Create: `RETAIN_PER_RETENTION_WINDOW`},
		"is_retention_lock_enabled":            acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"vpc_user":                             acctest.Representation{RepType: acctest.Optional, Create: `bkupUser1`},
	}
	autonomousContainerDatabaseBackupConfigBackupDestinationDetailsWithRARepresentation = map[string]interface{}{
		"type":                                 acctest.Representation{RepType: acctest.Required, Create: `RECOVERY_APPLIANCE`},
		"id":                                   acctest.Representation{RepType: acctest.Optional, Create: `${oci_database_backup_destination.ra_backup_destination.id}`},
		"internet_proxy":                       acctest.Representation{RepType: acctest.Optional, Create: `internetProxy`},
		"vpc_password":                         acctest.Representation{RepType: acctest.Optional, Create: `vpcPassword`},
		"vpc_user":                             acctest.Representation{RepType: acctest.Optional, Create: `bkupUser1`},
		"backup_retention_policy_on_terminate": acctest.Representation{RepType: acctest.Optional, Create: `RETAIN_PER_RETENTION_WINDOW`},
		"is_retention_lock_enabled":            acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
	}
	acdBackupConfigLocalRepresentation = map[string]interface{}{
		"backup_destination_details": acctest.RepresentationGroup{RepType: acctest.Optional, Group: map[string]interface{}{
			"type": acctest.Representation{RepType: acctest.Required, Create: `LOCAL`}}},
		"recovery_window_in_days": acctest.Representation{RepType: acctest.Optional, Create: `7`},
	}
	ACDatabaseResourceDependencies = DatabaseAVMClusterWithSingleNetworkResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_database_backup_destination", "test_backup_destination", acctest.Optional, acctest.Create, backupDestinationNFSRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_vm_cluster", "test_autonomous_vm_cluster", acctest.Required, acctest.Create, DatabaseECPUAutonomousVmClusterRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_database_key_store", "test_key_store", acctest.Optional, acctest.Create, DatabaseKeyStoreRepresentation) +
		KmsVaultIdVariableStr + OkvSecretVariableStr

	ACDECPUatabaseResourceDependencies = DatabaseAVMClusterWithSingleNetworkResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_database_backup_destination", "test_backup_destination", acctest.Optional, acctest.Create, backupDestinationNFSRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_vm_cluster", "test_autonomous_vm_cluster", acctest.Required, acctest.Create, DatabaseECPUAutonomousVmClusterRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_database_key_store", "test_key_store", acctest.Optional, acctest.Create, DatabaseKeyStoreRepresentationWithIgnoreTagsChanges) +
		KmsVaultIdVariableStr + OkvSecretVariableStr

	dgDbUniqueName = utils.RandomString(10, utils.CharsetWithoutDigits)

	DatabaseKeyStoreRepresentationWithIgnoreTagsChanges = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":   acctest.Representation{RepType: acctest.Required, Create: `Key Store1`},
		"type_details":   acctest.RepresentationGroup{RepType: acctest.Required, Group: DatabaseKeyStoreTypeDetailsRepresentation},
		"defined_tags":   acctest.Representation{RepType: acctest.Optional, Create: `${tomap({"${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}" = "value"})}`, Update: `${tomap({"${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}" = "updatedValue"})}`},
		"freeform_tags":  acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"lifecycle":      acctest.RepresentationGroup{RepType: acctest.Optional, Group: lifecycleGroupWithTagsIgnoreChanges},
	}
	lifecycleGroupWithTagsIgnoreChanges = map[string]interface{}{
		"ignore_changes": acctest.Representation{RepType: acctest.Optional, Create: []string{`defined_tags`, `freeform_tags`}},
	}
	DatabaseAutonomousContainerExaccAutonomousContainerDatabaseDataguardAssociationResourceConfig2 = ExaccACDWithDataGuardResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_database_backup_destination", "test_backup_destination", acctest.Optional, acctest.Create, backupDestinationADBCCNFSRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_database_backup_destination", "ra_backup_destination", acctest.Optional, acctest.Create, DatabaseBackupDestinationRepresentation)

	ExaccACDWithDGUpdateBkpDesRepresentation = map[string]interface{}{
		"db_version":                   acctest.Representation{RepType: acctest.Required, Create: utils.GetEnvSettingWithDefault("acd_db_version", "19.28.0.1.0")},
		"display_name":                 acctest.Representation{RepType: acctest.Required, Create: `ACD-DG-TF-TEST`},
		"patch_model":                  acctest.Representation{RepType: acctest.Required, Create: `RELEASE_UPDATES`, Update: `RELEASE_UPDATE_REVISIONS`},
		"autonomous_vm_cluster_id":     acctest.Representation{RepType: acctest.Required, Create: `${oci_database_autonomous_vm_cluster.test_autonomous_vm_cluster.id}`},
		"backup_config":                acctest.RepresentationGroup{RepType: acctest.Optional, Group: DatabaseAutonomousContainerDatabaseBackupConfigWithRAUpdateRepresentation},
		"compartment_id":               acctest.Representation{RepType: acctest.Optional, Create: `${var.compartment_id}`},
		"db_unique_name":               acctest.Representation{RepType: acctest.Optional, Create: dgDbUniqueName},
		"freeform_tags":                acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"maintenance_window_details":   acctest.RepresentationGroup{RepType: acctest.Optional, Group: DatabaseAutonomousContainerDatabaseMaintenanceWindowDetailsRepresentation},
		"service_level_agreement_type": acctest.Representation{RepType: acctest.Optional, Create: `AUTONOMOUS_DATAGUARD`},
		"peer_autonomous_container_database_backup_config":  acctest.RepresentationGroup{RepType: acctest.Optional, Group: AddStandbyAutonomousContainerDatabaseBackupConfigWithNoUpdateRepresentation},
		"peer_autonomous_container_database_compartment_id": acctest.Representation{RepType: acctest.Optional, Create: `${var.compartment_id}`},
		"peer_autonomous_vm_cluster_id":                     acctest.Representation{RepType: acctest.Optional, Create: `${oci_database_autonomous_vm_cluster.peer_autonomous_vm_cluster.id}`},
		"peer_autonomous_container_database_display_name":   acctest.Representation{RepType: acctest.Optional, Create: `PEER-ACD-DG`},
		"protection_mode":               acctest.Representation{RepType: acctest.Optional, Create: `MAXIMUM_PERFORMANCE`},
		"is_automatic_failover_enabled": acctest.Representation{RepType: acctest.Optional, Create: `true`},
	}
	ExaCCStandbyAcdImportContainer = acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_container_database", "standby_acd", acctest.Optional, acctest.Create,
		acctest.RepresentationCopyWithNewProperties(StandbyACDRepresentation, map[string]interface{}{
			"lifecycle":     acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreDataguardChangesRep},
			"backup_config": acctest.RepresentationGroup{RepType: acctest.Required, Group: DatabaseAutonomousContainerDatabaseBackupConfigWithRARepresentation},
			//"switchover_trigger": acctest.Representation{RepType: acctest.Required, Create: `1`},
		}))
	StandbyACDRepresentation = map[string]interface{}{
		"depends_on":                 acctest.Representation{RepType: acctest.Required, Create: []string{"oci_database_autonomous_container_database.exacc_test_autonomous_container_database"}},
		"autonomous_vm_cluster_id":   acctest.Representation{RepType: acctest.Optional, Create: `${oci_database_autonomous_vm_cluster.peer_autonomous_vm_cluster.id}`},
		"compartment_id":             acctest.Representation{RepType: acctest.Optional, Create: `${var.compartment_id}`},
		"display_name":               acctest.Representation{RepType: acctest.Optional, Create: `PEER-ACD-DG`},
		"patch_model":                acctest.Representation{RepType: acctest.Required, Create: `RELEASE_UPDATES`},
		"maintenance_window_details": acctest.RepresentationGroup{RepType: acctest.Optional, Group: DatabaseAutonomousContainerDatabaseMaintenanceWindowDetailsNoPreferenceRepresentation},
	}

	ignoreDataguardChangesWithMRRep = map[string]interface{}{
		"ignore_changes": acctest.Representation{RepType: acctest.Required, Create: []string{`peer_autonomous_container_database_compartment_id`,
			`peer_autonomous_container_database_display_name`,
			`peer_autonomous_exadata_infrastructure_id`,
			`peer_autonomous_vm_cluster_id`,
			`peer_cloud_autonomous_vm_cluster_id`,
			`peer_db_unique_name`,
			`service_level_agreement_type`,
			`protection_mode`,
			`peer_autonomous_container_database_backup_config`,
			`maintenance_window_details`}},
	}
	ExaccDatabaseAutonomousContainerDatabaseResourceDependencies = DatabaseAutonomousVmClusterResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_vm_cluster", "test_autonomous_vm_cluster", acctest.Required, acctest.Create, DatabaseECPUAutonomousVmClusterRepresentation) +
		KeyResourceDependencyConfig + kmsKeyIdCreateVariableStr + kmsKeyIdUpdateVariableStr +
		acctest.GenerateResourceFromRepresentationMap("oci_database_backup_destination", "test_backup_destination", acctest.Optional, acctest.Create, backupDestinationADBCCNFSRepresentation) +
		OkvSecretVariableStr +
		acctest.GenerateResourceFromRepresentationMap("oci_database_key_store", "test_key_store", acctest.Optional, acctest.Create, DatabaseKeyStoreRepresentationWithIgnoreTagsChanges) +
		acctest.GenerateResourceFromRepresentationMap("oci_database_backup_destination", "ra_backup_destination", acctest.Optional, acctest.Create, DatabaseBackupDestinationRepresentation)

	ExaccDatabaseAutonomousContainerDatabaseResourceFromAdsiDependencies = ExaccDatabaseAutonomousDatabaseSoftwareImageResourceConfig

	DatabaseAutonomousContainerDatabaseResourceFromAdsiDependencies = DatabaseAutonomousDatabaseSoftwareImageResourceConfig
)

// issue-routing-tag: database/dbaas-atp-d
func TestDatabaseExaccAutonomousContainerDatabaseFromAdsi_basic(t *testing.T) {

	httpreplay.SetScenario("TestDatabaseExaccAutonomousContainerDatabaseFromAdsi_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_database_autonomous_container_database.test_autonomous_container_database_from_adsi"

	var resId, resId2 string

	acctest.ResourceTest(t, testAccCheckDatabaseAutonomousContainerDatabaseDestroy, []resource.TestStep{
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + ExaccDatabaseAutonomousContainerDatabaseResourceFromAdsiDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_container_database", "test_autonomous_container_database_from_adsi", acctest.Optional, acctest.Create,
					acctest.GetUpdatedRepresentationCopy("maintenance_window_details", acctest.RepresentationGroup{RepType: acctest.Optional, Group: DatabaseAutonomousContainerDatabaseMaintenanceWindowDetailsNoPreferenceRepresentation}, ExaccACDatabaseFromAdsiRepresentation)),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "autonomous_vm_cluster_id"),
				resource.TestCheckResourceAttr(resourceName, "backup_config.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "backup_config.0.recovery_window_in_days", "10"),
				resource.TestCheckResourceAttr(resourceName, "backup_config.0.backup_destination_details.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "backup_config.0.backup_destination_details.0.id"),
				resource.TestCheckResourceAttr(resourceName, "backup_config.0.backup_destination_details.0.internet_proxy", "internetProxy"),
				resource.TestCheckResourceAttr(resourceName, "backup_config.0.backup_destination_details.0.type", "NFS"),
				resource.TestCheckResourceAttr(resourceName, "backup_config.0.backup_destination_details.0.vpc_user", "bkupUser1"),
				resource.TestCheckResourceAttr(resourceName, "backup_config.0.backup_destination_details.0.vpc_password", "vpcPassword"),
				resource.TestCheckResourceAttr(resourceName, "backup_config.0.backup_destination_details.0.is_retention_lock_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "backup_config.0.backup_destination_details.0.backup_retention_policy_on_terminate", "RETAIN_FOR_72_HOURS"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "db_unique_name", acbDBName2),
				resource.TestCheckResourceAttr(resourceName, "display_name", "containerdatabases3"),
				resource.TestCheckResourceAttrSet(resourceName, "memory_per_oracle_compute_unit_in_gbs"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "key_store_id"),
				resource.TestCheckResourceAttrSet(resourceName, "key_store_wallet_name"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window.0.days_of_week.#", "0"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window.0.hours_of_day.#", "0"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window.0.months.#", "0"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window.0.preference", "NO_PREFERENCE"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window.0.weeks_of_month.#", "0"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window.0.patching_mode", "ROLLING"),
				resource.TestCheckResourceAttr(resourceName, "patch_model", "RELEASE_UPDATES"),
				resource.TestCheckResourceAttr(resourceName, "service_level_agreement_type", "STANDARD"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttr(resourceName, "version_preference", "LATEST_RELEASE_UPDATE"),
				resource.TestCheckResourceAttrSet(resourceName, "db_version"),
				resource.TestCheckResourceAttrSet(resourceName, "database_software_image_id"),
				resource.TestCheckResourceAttr(resourceName, "db_name", "DBNAME2"),
				resource.TestCheckResourceAttr(resourceName, "is_dst_file_update_enabled", "false"),
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

		// verify updates to updatable parameters
		{
			Config: config + compartmentIdVariableStr + ExaccDatabaseAutonomousContainerDatabaseResourceFromAdsiDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_container_database", "test_autonomous_container_database_from_adsi", acctest.Optional, acctest.Update, ExaccACDatabaseFromAdsiRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "autonomous_vm_cluster_id"),
				resource.TestCheckResourceAttrSet(resourceName, "autonomous_vm_cluster_id"),
				resource.TestCheckResourceAttr(resourceName, "backup_config.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "backup_config.0.recovery_window_in_days", "11"),
				resource.TestCheckResourceAttrSet(resourceName, "backup_config.0.backup_destination_details.0.id"),
				resource.TestCheckResourceAttr(resourceName, "backup_config.0.backup_destination_details.0.internet_proxy", "internetProxy"),
				resource.TestCheckResourceAttr(resourceName, "backup_config.0.backup_destination_details.0.type", "NFS"),
				resource.TestCheckResourceAttr(resourceName, "backup_config.0.backup_destination_details.0.vpc_user", "bkupUser1"),
				resource.TestCheckResourceAttr(resourceName, "backup_config.0.backup_destination_details.0.is_retention_lock_enabled", "true"),
				resource.TestCheckResourceAttr(resourceName, "backup_config.0.backup_destination_details.0.backup_retention_policy_on_terminate", "RETAIN_PER_RETENTION_WINDOW"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "db_unique_name", acbDBName2),
				resource.TestCheckResourceAttr(resourceName, "display_name", "containerdatabases3"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "key_store_id"),
				resource.TestCheckResourceAttrSet(resourceName, "key_store_wallet_name"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window.0.days_of_week.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window.0.days_of_week.0.name", "TUESDAY"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window.0.hours_of_day.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window.0.months.#", "4"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window.0.months.0.name", "FEBRUARY"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window.0.preference", "CUSTOM_PREFERENCE"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window.0.weeks_of_month.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window.0.patching_mode", "NONROLLING"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window.0.skip_ru.#", "4"),
				resource.TestCheckResourceAttr(resourceName, "patch_model", "RELEASE_UPDATE_REVISIONS"),
				resource.TestCheckResourceAttr(resourceName, "service_level_agreement_type", "STANDARD"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttr(resourceName, "version_preference", "NEXT_RELEASE_UPDATE"),
				resource.TestCheckResourceAttrSet(resourceName, "db_version"),
				resource.TestCheckResourceAttrSet(resourceName, "database_software_image_id"),
				resource.TestCheckResourceAttr(resourceName, "db_name", "DBNAME2"),
				resource.TestCheckResourceAttr(resourceName, "is_dst_file_update_enabled", "true"),
				resource.TestCheckResourceAttr(resourceName, "system_tags.%", "0"),

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
			Config:            config + ExaccACDRequiredOnlyResource + compartmentIdVariableStr,
			ImportState:       true,
			ImportStateVerify: true,
			ImportStateVerifyIgnore: []string{
				"database_software_image_id",
				"maintenance_window_details",
				"backup_config.0.backup_destination_details.0.vpc_password",
			},
			ResourceName: resourceName,
		},
	})
}

func TestDatabaseAutonomousContainerDatabaseFromAdsi_basic(t *testing.T) {

	httpreplay.SetScenario("TestDatabaseAutonomousContainerDatabaseFromAdsi_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_database_autonomous_container_database.test_autonomous_container_database_from_adsi"

	var resId, resId2 string

	acctest.ResourceTest(t, testAccCheckDatabaseAutonomousContainerDatabaseDestroy, []resource.TestStep{
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + DatabaseAutonomousContainerDatabaseResourceFromAdsiDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_container_database", "test_autonomous_container_database_from_adsi", acctest.Optional, acctest.Create,
					acctest.GetUpdatedRepresentationCopy("maintenance_window_details", acctest.RepresentationGroup{RepType: acctest.Optional, Group: DatabaseAutonomousContainerDatabaseMaintenanceWindowDetailsNoPreferenceRepresentation}, ACDatabaseFromAdsiRepresentation)),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "cloud_autonomous_vm_cluster_id"),
				resource.TestCheckResourceAttr(resourceName, "backup_config.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "backup_config.0.recovery_window_in_days", "10"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "containerDatabase3"),
				resource.TestCheckResourceAttrSet(resourceName, "memory_per_oracle_compute_unit_in_gbs"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window.0.days_of_week.#", "0"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window.0.hours_of_day.#", "0"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window.0.months.#", "0"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window.0.preference", "NO_PREFERENCE"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window.0.weeks_of_month.#", "0"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window.0.patching_mode", "ROLLING"),
				resource.TestCheckResourceAttr(resourceName, "patch_model", "RELEASE_UPDATES"),
				resource.TestCheckResourceAttr(resourceName, "service_level_agreement_type", "STANDARD"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttr(resourceName, "version_preference", "LATEST_RELEASE_UPDATE"),
				resource.TestCheckResourceAttrSet(resourceName, "db_version"),
				resource.TestCheckResourceAttrSet(resourceName, "database_software_image_id"),
				resource.TestCheckResourceAttr(resourceName, "db_name", "DBNAME3"),
				resource.TestCheckResourceAttr(resourceName, "is_dst_file_update_enabled", "false"),
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

		// verify updates to updatable parameters
		{
			Config: config + compartmentIdVariableStr + DatabaseAutonomousContainerDatabaseResourceFromAdsiDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_container_database", "test_autonomous_container_database_from_adsi", acctest.Optional, acctest.Update, ACDatabaseFromAdsiRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "cloud_autonomous_vm_cluster_id"),
				resource.TestCheckResourceAttr(resourceName, "backup_config.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "backup_config.0.recovery_window_in_days", "11"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window.0.days_of_week.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window.0.days_of_week.0.name", "TUESDAY"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window.0.hours_of_day.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window.0.months.#", "4"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window.0.months.0.name", "FEBRUARY"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window.0.preference", "CUSTOM_PREFERENCE"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window.0.weeks_of_month.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window.0.patching_mode", "NONROLLING"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window.0.skip_ru.#", "4"),
				resource.TestCheckResourceAttr(resourceName, "patch_model", "RELEASE_UPDATE_REVISIONS"),
				resource.TestCheckResourceAttr(resourceName, "service_level_agreement_type", "STANDARD"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttr(resourceName, "version_preference", "NEXT_RELEASE_UPDATE"),
				resource.TestCheckResourceAttrSet(resourceName, "db_version"),
				resource.TestCheckResourceAttrSet(resourceName, "database_software_image_id"),
				resource.TestCheckResourceAttr(resourceName, "db_name", "DBNAME3"),
				resource.TestCheckResourceAttr(resourceName, "is_dst_file_update_enabled", "true"),
				resource.TestCheckResourceAttr(resourceName, "system_tags.%", "0"),

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
			Config:            config + ACDRequiredOnlyResource + compartmentIdVariableStr,
			ImportState:       true,
			ImportStateVerify: true,
			ImportStateVerifyIgnore: []string{
				"database_software_image_id",
				"maintenance_window_details",
				"backup_config.0.backup_destination_details.0.vpc_password",
				"is_automatic_failover_enabled",
				"state",
				"time_of_last_backup",
			},
			ResourceName: resourceName,
		},
	})
}
func TestDatabaseExaccAutonomousContainerDatabase_BackupDestinationUpdate_DG(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseExaccAutonomousContainerDatabase_BackupDestinationUpdate_DG")
	defer httpreplay.SaveScenario()

	provider := acctest.TestAccProvider
	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_database_autonomous_container_database_dataguard_associations.test"
	standbyResourceName := "oci_database_autonomous_container_database.standby_acd"
	resource.Test(t, resource.TestCase{
		PreCheck: func() { acctest.PreCheck(t) },
		Providers: map[string]*schema.Provider{
			"oci": provider,
		},
		Steps: []resource.TestStep{
			//  Create dg setup
			{
				Config: config +
					acctest.GenerateDataSourceFromRepresentationMap("oci_database_autonomous_container_database_dataguard_associations", "test", acctest.Optional, acctest.Create, DatabaseExaccAutonomousContainerDatabaseDataguardAssociationDataSourceRepresentation) +
					compartmentIdVariableStr + DatabaseAutonomousContainerExaccAutonomousContainerDatabaseDataguardAssociationResourceConfig2 + acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_container_database", "exacc_test_autonomous_container_database", acctest.Optional, acctest.Create,
					acctest.GetUpdatedRepresentationCopy("maintenance_window_details", acctest.RepresentationGroup{RepType: acctest.Optional, Group: DatabaseAutonomousContainerDatabaseMaintenanceWindowDetailsNoPreferenceRepresentation}, ExaccACDWithDGUpdateBkpDesRepresentation)),
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
			// NEW STEP: Refresh state
			{
				RefreshState: true, // reload state
			},
			{ // import the standby ACD
				Config: config +
					acctest.GenerateDataSourceFromRepresentationMap("oci_database_autonomous_container_database_dataguard_associations", "test", acctest.Optional, acctest.Create, DatabaseExaccAutonomousContainerDatabaseDataguardAssociationDataSourceRepresentation) +
					compartmentIdVariableStr + DatabaseAutonomousContainerExaccAutonomousContainerDatabaseDataguardAssociationResourceConfig2 + acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_container_database", "exacc_test_autonomous_container_database", acctest.Optional, acctest.Create,
					acctest.GetUpdatedRepresentationCopy("maintenance_window_details", acctest.RepresentationGroup{RepType: acctest.Optional, Group: DatabaseAutonomousContainerDatabaseMaintenanceWindowDetailsNoPreferenceRepresentation}, ExaccACDWithDGUpdateBkpDesRepresentation)) +
					acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_container_database", "standby_acd", acctest.Optional, acctest.Create,
						acctest.RepresentationCopyWithNewProperties(StandbyACDRepresentation, map[string]interface{}{
							"lifecycle": acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreDataguardChangesWithMRRep},
						})),
				ImportState:        true,
				ImportStateIdFunc:  getStandbyAcdOcidOldDG("data.oci_database_autonomous_container_database_dataguard_associations.test"),
				ImportStatePersist: true,
				ResourceName:       standbyResourceName,
			},
			{ // update backupDestination on standby ACD to RA
				Config: config +
					acctest.GenerateDataSourceFromRepresentationMap("oci_database_autonomous_container_database_dataguard_associations", "test", acctest.Optional, acctest.Create, DatabaseExaccAutonomousContainerDatabaseDataguardAssociationDataSourceRepresentation) +
					compartmentIdVariableStr + DatabaseAutonomousContainerExaccAutonomousContainerDatabaseDataguardAssociationResourceConfig2 + acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_container_database", "exacc_test_autonomous_container_database", acctest.Optional, acctest.Create,
					acctest.GetUpdatedRepresentationCopy("maintenance_window_details", acctest.RepresentationGroup{RepType: acctest.Optional, Group: DatabaseAutonomousContainerDatabaseMaintenanceWindowDetailsNoPreferenceRepresentation}, ExaccACDWithDGUpdateBkpDesRepresentation)) +
					acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_container_database", "standby_acd", acctest.Optional, acctest.Create,
						acctest.RepresentationCopyWithNewProperties(StandbyACDRepresentation, map[string]interface{}{
							"lifecycle":     acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreDataguardChangesWithMRRep},
							"backup_config": acctest.RepresentationGroup{RepType: acctest.Required, Group: DatabaseAutonomousContainerDatabaseBackupConfigWithRARepresentation},
						})),
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(standbyResourceName, "backup_config.0.backup_destination_details.0.type", "RECOVERY_APPLIANCE"),
					resource.TestCheckResourceAttr(standbyResourceName, "display_name", "PEER-ACD-DG"),
				),
			},
		},
	})
}

func TestDatabaseExaccAutonomousContainerDatabase_BackupDestinationUpdate(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseExaccAutonomousContainerDatabase_BackupDestinationUpdate")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_database_autonomous_container_database.test_autonomous_container_database"

	var resId, resId2 string

	acctest.ResourceTest(t, testAccCheckDatabaseAutonomousContainerDatabaseDestroy, []resource.TestStep{
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + ExaccDatabaseAutonomousContainerDatabaseResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_container_database", "test_autonomous_container_database", acctest.Optional, acctest.Create,
					acctest.GetUpdatedRepresentationCopy("maintenance_window_details", acctest.RepresentationGroup{RepType: acctest.Optional, Group: DatabaseAutonomousContainerDatabaseMaintenanceWindowDetailsNoPreferenceRepresentation}, ACDatabaseWithRABkpDesRepresentation)),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "autonomous_vm_cluster_id"),
				resource.TestCheckResourceAttr(resourceName, "backup_config.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "backup_config.0.recovery_window_in_days", "10"),
				resource.TestCheckResourceAttr(resourceName, "backup_config.0.backup_destination_details.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "backup_config.0.backup_destination_details.0.id"),
				resource.TestCheckResourceAttr(resourceName, "backup_config.0.backup_destination_details.0.internet_proxy", "internetProxy"),
				resource.TestCheckResourceAttr(resourceName, "backup_config.0.backup_destination_details.0.type", "NFS"),
				resource.TestCheckResourceAttr(resourceName, "backup_config.0.backup_destination_details.0.vpc_user", "bkupUser1"),
				resource.TestCheckResourceAttr(resourceName, "backup_config.0.backup_destination_details.0.vpc_password", "vpcPassword"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "db_split_threshold", "8"),
				resource.TestCheckResourceAttr(resourceName, "db_unique_name", acbDBName),
				resource.TestCheckResourceAttr(resourceName, "distribution_affinity", "MINIMUM_DISTRIBUTION"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "containerdatabases2"),
				resource.TestCheckResourceAttrSet(resourceName, "memory_per_oracle_compute_unit_in_gbs"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "key_store_id"),
				resource.TestCheckResourceAttrSet(resourceName, "key_store_wallet_name"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window.0.days_of_week.#", "0"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window.0.hours_of_day.#", "0"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window.0.months.#", "0"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window.0.preference", "NO_PREFERENCE"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window.0.weeks_of_month.#", "0"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window.0.patching_mode", "ROLLING"),
				resource.TestCheckResourceAttr(resourceName, "net_services_architecture", "DEDICATED"),
				resource.TestCheckResourceAttr(resourceName, "patch_model", "RELEASE_UPDATES"),
				resource.TestCheckResourceAttr(resourceName, "service_level_agreement_type", "STANDARD"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttr(resourceName, "version_preference", "LATEST_RELEASE_UPDATE"),
				resource.TestCheckResourceAttr(resourceName, "vm_failover_reservation", "25"),
				resource.TestCheckResourceAttrSet(resourceName, "db_version"),
				resource.TestCheckResourceAttr(resourceName, "db_name", "DBNAME"),
				resource.TestCheckResourceAttr(resourceName, "is_dst_file_update_enabled", "false"),
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

		// verify updates to updatable parameters along
		// BackupDestination NFS to RA
		{
			Config: config + compartmentIdVariableStr + ExaccDatabaseAutonomousContainerDatabaseResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_container_database", "test_autonomous_container_database", acctest.Optional, acctest.Update, ACDatabaseWithRABkpDesRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "autonomous_vm_cluster_id"),
				resource.TestCheckResourceAttrSet(resourceName, "autonomous_vm_cluster_id"),
				resource.TestCheckResourceAttr(resourceName, "backup_config.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "backup_config.0.backup_destination_details.0.id"),
				resource.TestCheckResourceAttr(resourceName, "backup_config.0.backup_destination_details.0.internet_proxy", "internetProxy"),
				resource.TestCheckResourceAttr(resourceName, "backup_config.0.backup_destination_details.0.type", "RECOVERY_APPLIANCE"),
				resource.TestCheckResourceAttr(resourceName, "backup_config.0.backup_destination_details.0.vpc_user", "bkupUser1"),
				resource.TestCheckResourceAttr(resourceName, "associated_backup_configuration_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "associated_backup_configuration_details.0.type", "NFS"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "db_split_threshold", "8"),
				resource.TestCheckResourceAttr(resourceName, "db_unique_name", acbDBName),
				resource.TestCheckResourceAttr(resourceName, "distribution_affinity", "MINIMUM_DISTRIBUTION"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "containerdatabases2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "key_store_id"),
				resource.TestCheckResourceAttrSet(resourceName, "key_store_wallet_name"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window.0.days_of_week.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window.0.days_of_week.0.name", "TUESDAY"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window.0.hours_of_day.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window.0.months.#", "4"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window.0.months.0.name", "FEBRUARY"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window.0.preference", "CUSTOM_PREFERENCE"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window.0.weeks_of_month.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window.0.patching_mode", "NONROLLING"),
				resource.TestCheckResourceAttr(resourceName, "net_services_architecture", "DEDICATED"),
				resource.TestCheckResourceAttr(resourceName, "patch_model", "RELEASE_UPDATE_REVISIONS"),
				resource.TestCheckResourceAttr(resourceName, "service_level_agreement_type", "STANDARD"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttr(resourceName, "version_preference", "NEXT_RELEASE_UPDATE"),
				resource.TestCheckResourceAttr(resourceName, "vm_failover_reservation", "25"),
				resource.TestCheckResourceAttrSet(resourceName, "db_version"),
				resource.TestCheckResourceAttr(resourceName, "db_name", "DBNAME"),
				resource.TestCheckResourceAttr(resourceName, "is_dst_file_update_enabled", "true"),
				resource.TestCheckResourceAttr(resourceName, "system_tags.%", "0"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("Resource recreated when it was supposed to be updated.")
					}
					return err
				},
			),
		},
		{ // BackupDestination RA to NFS
			Config: config + compartmentIdVariableStr + ExaccDatabaseAutonomousContainerDatabaseResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_container_database", "test_autonomous_container_database", acctest.Optional, acctest.Create,
					acctest.GetUpdatedRepresentationCopy("maintenance_window_details", acctest.RepresentationGroup{RepType: acctest.Optional, Group: DatabaseAutonomousContainerDatabaseMaintenanceWindowDetailsNoPreferenceRepresentation}, ACDatabaseWithRABkpDesRepresentation)),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "autonomous_vm_cluster_id"),
				resource.TestCheckResourceAttr(resourceName, "backup_config.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "backup_config.0.recovery_window_in_days", "10"),
				resource.TestCheckResourceAttr(resourceName, "backup_config.0.backup_destination_details.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "backup_config.0.backup_destination_details.0.id"),
				resource.TestCheckResourceAttr(resourceName, "backup_config.0.backup_destination_details.0.internet_proxy", "internetProxy"),
				resource.TestCheckResourceAttr(resourceName, "backup_config.0.backup_destination_details.0.type", "NFS"),
				resource.TestCheckResourceAttr(resourceName, "backup_config.0.backup_destination_details.0.vpc_user", "bkupUser1"),
			),
		},
	})
}

// issue-routing-tag: database/dbaas-atp-d
func TestDatabaseExaccAutonomousContainerDatabase_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseExaccAutonomousContainerDatabase_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_database_autonomous_container_database.test_autonomous_container_database"
	datasourceName := "data.oci_database_autonomous_container_databases.test_autonomous_container_databases"
	singularDatasourceName := "data.oci_database_autonomous_container_database.test_autonomous_container_database"

	var resId, resId2 string

	acctest.ResourceTest(t, testAccCheckDatabaseAutonomousContainerDatabaseDestroy, []resource.TestStep{
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + ExaccDatabaseAutonomousContainerDatabaseResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_container_database", "test_autonomous_container_database", acctest.Optional, acctest.Create,
					acctest.GetUpdatedRepresentationCopy("maintenance_window_details", acctest.RepresentationGroup{RepType: acctest.Optional, Group: DatabaseAutonomousContainerDatabaseMaintenanceWindowDetailsNoPreferenceRepresentation}, ACDatabaseRepresentation)),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "autonomous_vm_cluster_id"),
				resource.TestCheckResourceAttr(resourceName, "backup_config.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "backup_config.0.recovery_window_in_days", "10"),
				resource.TestCheckResourceAttr(resourceName, "backup_config.0.backup_destination_details.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "backup_config.0.backup_destination_details.0.id"),
				resource.TestCheckResourceAttr(resourceName, "customer_contacts.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "customer_contacts.0.email", "test1@oracle.com"),
				resource.TestCheckResourceAttr(resourceName, "backup_config.0.backup_destination_details.0.internet_proxy", "internetProxy"),
				resource.TestCheckResourceAttr(resourceName, "backup_config.0.backup_destination_details.0.type", "NFS"),
				resource.TestCheckResourceAttr(resourceName, "backup_config.0.backup_destination_details.0.vpc_user", "bkupUser1"),
				resource.TestCheckResourceAttr(resourceName, "backup_config.0.backup_destination_details.0.vpc_password", "vpcPassword"),
				resource.TestCheckResourceAttr(resourceName, "backup_config.0.backup_destination_details.0.is_retention_lock_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "backup_config.0.backup_destination_details.0.backup_retention_policy_on_terminate", "RETAIN_FOR_72_HOURS"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "db_split_threshold", "8"),
				resource.TestCheckResourceAttr(resourceName, "db_unique_name", acbDBName),
				resource.TestCheckResourceAttr(resourceName, "distribution_affinity", "MINIMUM_DISTRIBUTION"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "containerdatabases2"),
				resource.TestCheckResourceAttrSet(resourceName, "memory_per_oracle_compute_unit_in_gbs"),
				resource.TestCheckResourceAttr(resourceName, "okv_end_point_group_name", "DUMMY_OKV_EPG_GROUP"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "key_store_id"),
				resource.TestCheckResourceAttrSet(resourceName, "key_store_wallet_name"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window.0.days_of_week.#", "0"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window.0.hours_of_day.#", "0"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window.0.months.#", "0"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window.0.preference", "NO_PREFERENCE"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window.0.weeks_of_month.#", "0"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window.0.patching_mode", "ROLLING"),
				resource.TestCheckResourceAttr(resourceName, "net_services_architecture", "DEDICATED"),
				resource.TestCheckResourceAttr(resourceName, "patch_model", "RELEASE_UPDATES"),
				resource.TestCheckResourceAttr(resourceName, "service_level_agreement_type", "STANDARD"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttr(resourceName, "version_preference", "LATEST_RELEASE_UPDATE"),
				resource.TestCheckResourceAttr(resourceName, "vm_failover_reservation", "25"),
				resource.TestCheckResourceAttrSet(resourceName, "db_version"),
				resource.TestCheckResourceAttr(resourceName, "db_name", "DBNAME"),
				resource.TestCheckResourceAttr(resourceName, "is_dst_file_update_enabled", "false"),
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

		// verify updates to updatable parameters
		{
			Config: config + compartmentIdVariableStr + ExaccDatabaseAutonomousContainerDatabaseResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_container_database", "test_autonomous_container_database", acctest.Optional, acctest.Update, ACDatabaseRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "autonomous_vm_cluster_id"),
				resource.TestCheckResourceAttrSet(resourceName, "autonomous_vm_cluster_id"),
				resource.TestCheckResourceAttr(resourceName, "backup_config.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "backup_config.0.recovery_window_in_days", "11"),
				resource.TestCheckResourceAttrSet(resourceName, "backup_config.0.backup_destination_details.0.id"),
				resource.TestCheckResourceAttr(resourceName, "customer_contacts.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "customer_contacts.0.email", "test2@oracle.com"),
				resource.TestCheckResourceAttr(resourceName, "backup_config.0.backup_destination_details.0.internet_proxy", "internetProxy"),
				resource.TestCheckResourceAttr(resourceName, "backup_config.0.backup_destination_details.0.type", "NFS"),
				resource.TestCheckResourceAttr(resourceName, "backup_config.0.backup_destination_details.0.vpc_user", "bkupUser1"),
				resource.TestCheckResourceAttr(resourceName, "backup_config.0.backup_destination_details.0.is_retention_lock_enabled", "true"),
				resource.TestCheckResourceAttr(resourceName, "backup_config.0.backup_destination_details.0.backup_retention_policy_on_terminate", "RETAIN_PER_RETENTION_WINDOW"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "okv_end_point_group_name", "DUMMY_OKV_EPG_GROUP_2"),
				resource.TestCheckResourceAttr(resourceName, "db_split_threshold", "8"),
				resource.TestCheckResourceAttr(resourceName, "db_unique_name", acbDBName),
				resource.TestCheckResourceAttr(resourceName, "distribution_affinity", "MINIMUM_DISTRIBUTION"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "containerdatabases2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "key_store_id"),
				resource.TestCheckResourceAttrSet(resourceName, "key_store_wallet_name"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window.0.days_of_week.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window.0.days_of_week.0.name", "TUESDAY"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window.0.hours_of_day.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window.0.months.#", "4"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window.0.months.0.name", "FEBRUARY"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window.0.preference", "CUSTOM_PREFERENCE"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window.0.weeks_of_month.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window.0.patching_mode", "NONROLLING"),
				resource.TestCheckResourceAttr(resourceName, "net_services_architecture", "DEDICATED"),
				resource.TestCheckResourceAttr(resourceName, "patch_model", "RELEASE_UPDATE_REVISIONS"),
				resource.TestCheckResourceAttr(resourceName, "service_level_agreement_type", "STANDARD"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttr(resourceName, "version_preference", "NEXT_RELEASE_UPDATE"),
				resource.TestCheckResourceAttr(resourceName, "vm_failover_reservation", "25"),
				resource.TestCheckResourceAttrSet(resourceName, "db_version"),
				resource.TestCheckResourceAttr(resourceName, "db_name", "DBNAME"),
				resource.TestCheckResourceAttr(resourceName, "is_dst_file_update_enabled", "true"),
				resource.TestCheckResourceAttr(resourceName, "system_tags.%", "0"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("Resource recreated when it was supposed to be updated.")
					}
					return err
				},
			),
		},
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_autonomous_container_databases", "test_autonomous_container_databases", acctest.Optional, acctest.Create, ACDatabaseDataSourceRepresentation) +
				compartmentIdVariableStr + ExaccDatabaseAutonomousContainerDatabaseResourceDependencies + ExaccAcdResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "autonomous_vm_cluster_id"),
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "containerdatabases2"),
				resource.TestCheckResourceAttr(datasourceName, "state", "AVAILABLE"),

				resource.TestCheckResourceAttr(datasourceName, "autonomous_container_databases.#", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "autonomous_container_databases.0.autonomous_vm_cluster_id"),
				resource.TestCheckResourceAttr(datasourceName, "autonomous_container_databases.0.backup_config.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "autonomous_container_databases.0.backup_config.0.recovery_window_in_days", "11"),
				resource.TestCheckResourceAttrSet(datasourceName, "autonomous_container_databases.0.backup_config.0.backup_destination_details.0.id"),
				resource.TestCheckResourceAttr(datasourceName, "autonomous_container_databases.0.backup_config.0.backup_destination_details.0.internet_proxy", "internetProxy"),
				resource.TestCheckResourceAttr(datasourceName, "autonomous_container_databases.0.backup_config.0.backup_destination_details.0.type", "NFS"),
				resource.TestCheckResourceAttr(datasourceName, "autonomous_container_databases.0.backup_config.0.backup_destination_details.0.vpc_user", "bkupUser1"),
				resource.TestCheckResourceAttr(datasourceName, "autonomous_container_databases.0.backup_config.0.backup_destination_details.0.is_retention_lock_enabled", "true"),
				resource.TestCheckResourceAttr(datasourceName, "autonomous_container_databases.0.backup_config.0.backup_destination_details.0.backup_retention_policy_on_terminate", "RETAIN_PER_RETENTION_WINDOW"),
				resource.TestCheckResourceAttr(datasourceName, "autonomous_container_databases.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "autonomous_container_databases.0.db_split_threshold", "8"),
				resource.TestCheckResourceAttr(datasourceName, "autonomous_container_databases.0.display_name", "containerdatabases2"),
				resource.TestCheckResourceAttr(datasourceName, "autonomous_container_databases.0.distribution_affinity", "MINIMUM_DISTRIBUTION"),
				resource.TestCheckResourceAttr(datasourceName, "autonomous_container_databases.0.freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "key_store_id"),
				resource.TestCheckResourceAttrSet(resourceName, "key_store_wallet_name"),
				resource.TestCheckResourceAttrSet(datasourceName, "autonomous_container_databases.0.id"),
				resource.TestCheckResourceAttr(datasourceName, "autonomous_container_databases.0.is_dst_file_update_enabled", "true"),
				resource.TestCheckResourceAttr(datasourceName, "autonomous_container_databases.0.maintenance_window.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "autonomous_container_databases.0.maintenance_window.0.days_of_week.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "autonomous_container_databases.0.maintenance_window.0.days_of_week.0.name", "TUESDAY"),
				resource.TestCheckResourceAttr(datasourceName, "autonomous_container_databases.0.maintenance_window.0.hours_of_day.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "autonomous_container_databases.0.maintenance_window.0.months.#", "4"),
				resource.TestCheckResourceAttr(datasourceName, "autonomous_container_databases.0.maintenance_window.0.months.0.name", "FEBRUARY"),
				resource.TestCheckResourceAttr(datasourceName, "autonomous_container_databases.0.maintenance_window.0.preference", "CUSTOM_PREFERENCE"),
				resource.TestCheckResourceAttr(datasourceName, "autonomous_container_databases.0.maintenance_window.0.weeks_of_month.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "autonomous_container_databases.0.net_services_architecture", "DEDICATED"),
				resource.TestCheckResourceAttr(datasourceName, "autonomous_container_databases.0.patch_model", "RELEASE_UPDATE_REVISIONS"),
				resource.TestCheckResourceAttr(datasourceName, "autonomous_container_databases.0.service_level_agreement_type", "STANDARD"),
				resource.TestCheckResourceAttrSet(datasourceName, "autonomous_container_databases.0.state"),
				resource.TestCheckResourceAttrSet(datasourceName, "autonomous_container_databases.0.time_created"),
				resource.TestCheckResourceAttrSet(datasourceName, "autonomous_container_databases.0.memory_per_oracle_compute_unit_in_gbs"),
				resource.TestCheckResourceAttr(datasourceName, "autonomous_container_databases.0.version_preference", "NEXT_RELEASE_UPDATE"),
				resource.TestCheckResourceAttr(datasourceName, "autonomous_container_databases.0.vm_failover_reservation", "25"),
				resource.TestCheckResourceAttrSet(datasourceName, "autonomous_container_databases.0.db_version"),
				resource.TestCheckResourceAttr(datasourceName, "autonomous_container_databases.0.maintenance_window.0.patching_mode", "NONROLLING"),
				resource.TestCheckResourceAttr(resourceName, "autonomous_container_databases.0.system_tags.%", "0"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_autonomous_container_database", "test_autonomous_container_database", acctest.Required, acctest.Create, DatabaseDatabaseAutonomousContainerDatabaseSingularDataSourceRepresentation) +
				compartmentIdVariableStr + ExaccDatabaseAutonomousContainerDatabaseResourceDependencies + ExaccAcdResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "autonomous_container_database_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "backup_config.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "backup_config.0.recovery_window_in_days", "11"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "backup_config.0.backup_destination_details.0.id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "backup_config.0.backup_destination_details.0.internet_proxy", "internetProxy"),
				resource.TestCheckResourceAttr(singularDatasourceName, "backup_config.0.backup_destination_details.0.type", "NFS"),
				resource.TestCheckResourceAttr(singularDatasourceName, "backup_config.0.backup_destination_details.0.vpc_user", "bkupUser1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "backup_config.0.backup_destination_details.0.is_retention_lock_enabled", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "backup_config.0.backup_destination_details.0.backup_retention_policy_on_terminate", "RETAIN_PER_RETENTION_WINDOW"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "db_split_threshold", "8"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "containerdatabases2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "distribution_affinity", "MINIMUM_DISTRIBUTION"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "is_dst_file_update_enabled", "true"),
				resource.TestCheckResourceAttrSet(resourceName, "key_store_id"),
				resource.TestCheckResourceAttrSet(resourceName, "key_store_wallet_name"),
				resource.TestCheckResourceAttr(singularDatasourceName, "maintenance_window.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "maintenance_window.0.days_of_week.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "maintenance_window.0.days_of_week.0.name", "TUESDAY"),
				resource.TestCheckResourceAttr(singularDatasourceName, "maintenance_window.0.hours_of_day.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "maintenance_window.0.months.#", "4"),
				resource.TestCheckResourceAttr(singularDatasourceName, "maintenance_window.0.months.0.name", "FEBRUARY"),
				resource.TestCheckResourceAttr(singularDatasourceName, "maintenance_window.0.preference", "CUSTOM_PREFERENCE"),
				resource.TestCheckResourceAttr(singularDatasourceName, "maintenance_window.0.weeks_of_month.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "net_services_architecture", "DEDICATED"),
				resource.TestCheckResourceAttr(singularDatasourceName, "patch_model", "RELEASE_UPDATE_REVISIONS"),
				resource.TestCheckResourceAttr(singularDatasourceName, "service_level_agreement_type", "STANDARD"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "memory_per_oracle_compute_unit_in_gbs"),
				resource.TestCheckResourceAttr(singularDatasourceName, "version_preference", "NEXT_RELEASE_UPDATE"),
				resource.TestCheckResourceAttr(singularDatasourceName, "vm_failover_reservation", "25"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "db_version"),
				resource.TestCheckResourceAttr(singularDatasourceName, "maintenance_window.0.patching_mode", "NONROLLING"),
			),
		},
		// verify resource import
		{
			Config:            config + DatabaseAutonomousContainerDatabaseRequiredOnlyResource + compartmentIdVariableStr,
			ImportState:       true,
			ImportStateVerify: true,
			ImportStateVerifyIgnore: []string{
				"maintenance_window_details",
				"backup_config.0.backup_destination_details.0.vpc_password",
			},
			ResourceName: resourceName,
		},
	})
}

// issue-routing-tag: database/dbaas-atp-d
func TestDatabaseAutonomousContainerDatabase_rotateDatabase(t *testing.T) {
	t.Skip("Skip this test as AEI and its api no longer exists.")

	httpreplay.SetScenario("TestDatabaseAutonomousContainerDatabase_rotateDatabase")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_database_autonomous_container_database.test_autonomous_container_database"
	datasourceName := "data.oci_database_autonomous_container_databases.test_autonomous_container_databases"
	singularDatasourceName := "data.oci_database_autonomous_container_database.test_autonomous_container_database"

	var resId, resId2 string

	acctest.ResourceTest(t, testAccCheckDatabaseAutonomousContainerDatabaseDestroy, []resource.TestStep{
		// verify Create with optionals and rotate key
		{
			Config: config + compartmentIdVariableStr + ATPDAutonomousContainerDatabaseResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_container_database", "test_autonomous_container_database", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(DatabaseAutonomousContainerDatabaseRepresentation, map[string]interface{}{
						"rotate_key_trigger": acctest.Representation{RepType: acctest.Optional, Create: `true`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "cloud_autonomous_vm_cluster_id"),
				resource.TestCheckResourceAttr(resourceName, "backup_config.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "backup_config.0.recovery_window_in_days", "10"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "containerdatabases2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "kms_key_id"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window.0.days_of_week.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window.0.hours_of_day.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window.0.months.#", "4"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window.0.preference", "CUSTOM_PREFERENCE"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window.0.weeks_of_month.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "patch_model", "RELEASE_UPDATES"),
				resource.TestCheckResourceAttr(resourceName, "service_level_agreement_type", "STANDARD"),
				resource.TestCheckResourceAttr(resourceName, "rotate_key_trigger", "true"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "vault_id"),

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

		// verify updates to updatable parameters
		{
			Config: config + compartmentIdVariableStr + ATPDAutonomousContainerDatabaseResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_container_database", "test_autonomous_container_database", acctest.Optional, acctest.Update,
					acctest.RepresentationCopyWithNewProperties(DatabaseAutonomousContainerDatabaseRepresentation, map[string]interface{}{
						"rotate_key_trigger": acctest.Representation{RepType: acctest.Optional, Create: `false`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "cloud_autonomous_vm_cluster_id"),
				resource.TestCheckResourceAttr(resourceName, "backup_config.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "backup_config.0.recovery_window_in_days", "11"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "kms_key_id"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window.0.days_of_week.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window.0.days_of_week.0.name", "TUESDAY"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window.0.hours_of_day.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window.0.months.#", "4"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window.0.months.0.name", "FEBRUARY"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window.0.preference", "CUSTOM_PREFERENCE"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window.0.weeks_of_month.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "patch_model", "RELEASE_UPDATE_REVISIONS"),
				resource.TestCheckResourceAttr(resourceName, "service_level_agreement_type", "STANDARD"),
				resource.TestCheckResourceAttr(resourceName, "rotate_key_trigger", "false"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "vault_id"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("Resource recreated when it was supposed to be updated.")
					}
					return err
				},
			),
		},
		// verify no rotation of key
		{
			Config: config + compartmentIdVariableStr + ATPDAutonomousContainerDatabaseResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_container_database", "test_autonomous_container_database", acctest.Optional, acctest.Update,
					acctest.RepresentationCopyWithNewProperties(DatabaseAutonomousContainerDatabaseRepresentation, map[string]interface{}{
						"rotate_key_trigger": acctest.Representation{RepType: acctest.Optional, Create: `false`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "cloud_autonomous_vm_cluster_id"),
				resource.TestCheckResourceAttr(resourceName, "backup_config.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "backup_config.0.recovery_window_in_days", "11"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "kms_key_id"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window.0.days_of_week.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window.0.days_of_week.0.name", "TUESDAY"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window.0.hours_of_day.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window.0.months.#", "4"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window.0.months.0.name", "FEBRUARY"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window.0.preference", "CUSTOM_PREFERENCE"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window.0.weeks_of_month.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "patch_model", "RELEASE_UPDATE_REVISIONS"),
				resource.TestCheckResourceAttr(resourceName, "service_level_agreement_type", "STANDARD"),
				resource.TestCheckResourceAttr(resourceName, "rotate_key_trigger", "false"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "vault_id"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("Resource recreated when it was supposed to be updated.")
					}
					return err
				},
			),
		},
		// verify rotate key
		{
			Config: config + compartmentIdVariableStr + ATPDAutonomousContainerDatabaseResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_container_database", "test_autonomous_container_database", acctest.Optional, acctest.Update,
					acctest.RepresentationCopyWithNewProperties(DatabaseAutonomousContainerDatabaseRepresentation, map[string]interface{}{
						"rotate_key_trigger": acctest.Representation{RepType: acctest.Optional, Create: `true`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "cloud_autonomous_vm_cluster_id"),
				resource.TestCheckResourceAttr(resourceName, "backup_config.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "backup_config.0.recovery_window_in_days", "11"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "kms_key_id"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window.0.days_of_week.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window.0.days_of_week.0.name", "TUESDAY"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window.0.hours_of_day.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window.0.months.#", "4"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window.0.months.0.name", "FEBRUARY"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window.0.preference", "CUSTOM_PREFERENCE"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window.0.weeks_of_month.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "patch_model", "RELEASE_UPDATE_REVISIONS"),
				resource.TestCheckResourceAttr(resourceName, "service_level_agreement_type", "STANDARD"),
				resource.TestCheckResourceAttr(resourceName, "rotate_key_trigger", "true"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "vault_id"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("Resource recreated when it was supposed to be updated.")
					}
					return err
				},
			),
		},
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_autonomous_container_databases", "test_autonomous_container_databases", acctest.Optional, acctest.Update, DatabaseDatabaseAutonomousContainerDatabaseDataSourceRepresentation) +
				compartmentIdVariableStr + ATPDAutonomousContainerDatabaseResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_container_database", "test_autonomous_container_database", acctest.Optional, acctest.Update, DatabaseAutonomousContainerDatabaseRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "cloud_autonomous_vm_cluster_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "availability_domain"),
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(datasourceName, "state", "AVAILABLE"),

				resource.TestCheckResourceAttr(datasourceName, "autonomous_container_databases.#", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "autonomous_container_databases.0.cloud_autonomous_vm_cluster_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "autonomous_container_databases.0.availability_domain"),
				resource.TestCheckResourceAttr(datasourceName, "autonomous_container_databases.0.backup_config.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "autonomous_container_databases.0.backup_config.0.recovery_window_in_days", "11"),
				resource.TestCheckResourceAttr(datasourceName, "autonomous_container_databases.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(datasourceName, "autonomous_container_databases.0.db_version"),
				resource.TestCheckResourceAttr(datasourceName, "autonomous_container_databases.0.display_name", "displayName2"),
				resource.TestCheckResourceAttr(datasourceName, "autonomous_container_databases.0.freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "autonomous_container_databases.0.id"),
				resource.TestCheckResourceAttrSet(datasourceName, "autonomous_container_databases.0.kms_key_id"),
				resource.TestCheckResourceAttr(datasourceName, "autonomous_container_databases.0.maintenance_window.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "autonomous_container_databases.0.maintenance_window.0.days_of_week.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "autonomous_container_databases.0.maintenance_window.0.days_of_week.0.name", "TUESDAY"),
				resource.TestCheckResourceAttr(datasourceName, "autonomous_container_databases.0.maintenance_window.0.hours_of_day.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "autonomous_container_databases.0.maintenance_window.0.months.#", "4"),
				resource.TestCheckResourceAttr(datasourceName, "autonomous_container_databases.0.maintenance_window.0.months.0.name", "FEBRUARY"),
				resource.TestCheckResourceAttr(datasourceName, "autonomous_container_databases.0.maintenance_window.0.preference", "CUSTOM_PREFERENCE"),
				resource.TestCheckResourceAttr(datasourceName, "autonomous_container_databases.0.maintenance_window.0.weeks_of_month.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "autonomous_container_databases.0.patch_model", "RELEASE_UPDATE_REVISIONS"),
				resource.TestCheckResourceAttr(datasourceName, "autonomous_container_databases.0.service_level_agreement_type", "STANDARD"),
				resource.TestCheckResourceAttrSet(datasourceName, "autonomous_container_databases.0.state"),
				resource.TestCheckResourceAttrSet(datasourceName, "autonomous_container_databases.0.time_created"),
				resource.TestCheckResourceAttrSet(datasourceName, "autonomous_container_databases.0.vault_id"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_autonomous_container_database", "test_autonomous_container_database", acctest.Required, acctest.Create, DatabaseDatabaseAutonomousContainerDatabaseSingularDataSourceRepresentation) +
				compartmentIdVariableStr + DatabaseAutonomousContainerDatabaseResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "autonomous_container_database_id"),

				resource.TestCheckResourceAttrSet(singularDatasourceName, "availability_domain"),
				resource.TestCheckResourceAttr(singularDatasourceName, "backup_config.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "backup_config.0.recovery_window_in_days", "11"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "db_version"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "maintenance_window.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "maintenance_window.0.days_of_week.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "maintenance_window.0.days_of_week.0.name", "TUESDAY"),
				resource.TestCheckResourceAttr(singularDatasourceName, "maintenance_window.0.hours_of_day.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "maintenance_window.0.months.#", "4"),
				resource.TestCheckResourceAttr(singularDatasourceName, "maintenance_window.0.months.0.name", "FEBRUARY"),
				resource.TestCheckResourceAttr(singularDatasourceName, "maintenance_window.0.preference", "CUSTOM_PREFERENCE"),
				resource.TestCheckResourceAttr(singularDatasourceName, "maintenance_window.0.weeks_of_month.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "patch_model", "RELEASE_UPDATE_REVISIONS"),
				resource.TestCheckResourceAttr(singularDatasourceName, "service_level_agreement_type", "STANDARD"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
			),
		},

		{
			Config: config + compartmentIdVariableStr + ATPDAutonomousContainerDatabaseResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_container_database", "test_autonomous_container_database", acctest.Optional, acctest.Update,
					acctest.GetUpdatedRepresentationCopy("maintenance_window_details", acctest.RepresentationGroup{RepType: acctest.Optional, Group: DatabaseAutonomousContainerDatabaseMaintenanceWindowDetailsNoPreferenceRepresentation}, DatabaseAutonomousContainerDatabaseRepresentation)),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "cloud_autonomous_vm_cluster_id"),
				resource.TestCheckResourceAttr(resourceName, "backup_config.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "backup_config.0.recovery_window_in_days", "11"),
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
				"rotate_key_trigger",
				"maintenance_window_details",
				"is_automatic_failover_enabled",
			},
			ResourceName: resourceName,
		},
	})
}
