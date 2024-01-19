// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"

	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/oracle/oci-go-sdk/v65/common"
	oci_database_migration "github.com/oracle/oci-go-sdk/v65/databasemigration"

	"github.com/oracle/terraform-provider-oci/httpreplay"
)

var (
	DatabaseMigrationMigrationRequiredOnlyResource = acctest.GenerateResourceFromRepresentationMap("oci_database_migration_migration", "test_migration", acctest.Required, acctest.Create, DatabaseMigrationMigrationRDSRepresentation)

	DatabaseMigrationMigrationResourceConfig = acctest.GenerateResourceFromRepresentationMap("oci_database_migration_migration", "test_migration", acctest.Optional, acctest.Update, DatabaseMigrationMigrationRepresentation)

	DatabaseMigrationmigrationSingularDataSourceRepresentation = map[string]interface{}{
		"migration_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_database_migration_migration.test_migration.id}`},
	}

	DatabaseMigrationmigrationDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":   acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"state":          acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"filter":         acctest.RepresentationGroup{RepType: acctest.Required, Group: migrationDataSourceFilterRepresentation}}
	migrationDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_database_migration_migration.test_migration.id}`}},
	}

	DatabaseMigrationMigrationRepresentation = map[string]interface{}{
		"compartment_id":                          acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"type":                                    acctest.Representation{RepType: acctest.Required, Create: `ONLINE`},
		"data_transfer_medium_details":            acctest.RepresentationGroup{RepType: acctest.Optional, Group: migrationDataTransferMediumDetailsRepresentation},
		"data_transfer_medium_details_v2":         acctest.RepresentationGroup{RepType: acctest.Optional, Group: DatabaseMigrationMigrationDataTransferMediumDetailsV2Representation},
		"datapump_settings":                       acctest.RepresentationGroup{RepType: acctest.Optional, Group: DatabaseMigrationMigrationDatapumpSettingsRepresentation},
		"dump_transfer_details":                   acctest.RepresentationGroup{RepType: acctest.Optional, Group: DatabaseMigrationMigrationDumpTransferDetailsRepresentation},
		"source_database_connection_id":           acctest.Representation{RepType: acctest.Required, Create: `${var.source_connection_id}`},
		"target_database_connection_id":           acctest.Representation{RepType: acctest.Required, Create: `${var.target_connection_no_autonomous_id}`},
		"advisor_settings":                        acctest.RepresentationGroup{RepType: acctest.Optional, Group: DatabaseMigrationMigrationAdvisorSettingsRepresentation},
		"exclude_objects":                         acctest.RepresentationGroup{RepType: acctest.Optional, Group: migrationExcludeObjectsRepresentation},
		"display_name":                            acctest.Representation{RepType: acctest.Optional, Create: `TFtestOnline1`, Update: `TFtestOnline2`},
		"freeform_tags":                           acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"bar-key": "value"}, Update: map[string]string{"Department": "Accounting"}},
		"golden_gate_details":                     acctest.RepresentationGroup{RepType: acctest.Optional, Group: migrationGoldenGateDetailsRepresentation},
		"golden_gate_service_details":             acctest.RepresentationGroup{RepType: acctest.Optional, Group: DatabaseMigrationMigrationGoldenGateServiceDetailsRepresentation},
		"source_container_database_connection_id": acctest.Representation{RepType: acctest.Optional, Create: `${var.source_connection_container_id}`},
		"vault_details":                           acctest.RepresentationGroup{RepType: acctest.Required, Group: DatabaseMigrationMigrationVaultDetailsRepresentation},
	}

	DatabaseMigrationMigrationRDSRepresentation = map[string]interface{}{
		"compartment_id":                  acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"type":                            acctest.Representation{RepType: acctest.Required, Create: `ONLINE`},
		"data_transfer_medium_details_v2": acctest.RepresentationGroup{RepType: acctest.Optional, Group: DatabaseMigrationMigrationDataTransferMediumDetailsV2Representation},
		"datapump_settings":               acctest.RepresentationGroup{RepType: acctest.Optional, Group: DatabaseMigrationMigrationDatapumpSettingsRepresentation},
		"dump_transfer_details":           acctest.RepresentationGroup{RepType: acctest.Optional, Group: DatabaseMigrationMigrationDumpTransferDetailsRepresentation},
		"source_database_connection_id":   acctest.Representation{RepType: acctest.Required, Create: `${var.target_connection_rds_id}`},
		"target_database_connection_id":   acctest.Representation{RepType: acctest.Required, Create: `${var.target_connection_no_autonomous_id}`},
		"advisor_settings":                acctest.RepresentationGroup{RepType: acctest.Optional, Group: DatabaseMigrationMigrationAdvisorSettingsRepresentation},
		"exclude_objects":                 acctest.RepresentationGroup{RepType: acctest.Optional, Group: migrationExcludeObjectsRepresentation},
		"display_name":                    acctest.Representation{RepType: acctest.Optional, Create: `TFtestOnline1`, Update: `TFtestOnline2`},
		"freeform_tags":                   acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"bar-key": "value"}, Update: map[string]string{"Department": "Accounting"}},
		"golden_gate_service_details":     acctest.RepresentationGroup{RepType: acctest.Optional, Group: DatabaseMigrationMigrationGoldenGateServiceDetailsRepresentation},
		"vault_details":                   acctest.RepresentationGroup{RepType: acctest.Required, Group: DatabaseMigrationMigrationVaultDetailsRepresentation},
	}

	migrationRepresentationMig = map[string]interface{}{
		"compartment_id":                          acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"source_database_connection_id":           acctest.Representation{RepType: acctest.Required, Create: `${var.source_connection_id}`},
		"target_database_connection_id":           acctest.Representation{RepType: acctest.Required, Create: `${var.target_connection_id}`},
		"type":                                    acctest.Representation{RepType: acctest.Required, Create: `ONLINE`},
		"csv_text":                                acctest.Representation{RepType: acctest.Optional, Create: `MY_BIZZ,SRC_CITY,TABLE,EXCLUDE`},
		"data_transfer_medium_details":            acctest.RepresentationGroup{RepType: acctest.Required, Group: migrationDataTransferMediumDetailsRepresentation},
		"datapump_settings":                       acctest.RepresentationGroup{RepType: acctest.Required, Group: migrationDatapumpSettingsRepresentation},
		"display_name":                            acctest.Representation{RepType: acctest.Optional, Create: `TF_ONLINE_MIG`, Update: `TF_ONLINE_MIG`},
		"golden_gate_details":                     acctest.RepresentationGroup{RepType: acctest.Required, Group: migrationGoldenGateDetailsRepresentation},
		"golden_gate_service_details":             acctest.RepresentationGroup{RepType: acctest.Optional, Group: DatabaseMigrationMigrationGoldenGateServiceDetailsRepresentation},
		"source_container_database_connection_id": acctest.Representation{RepType: acctest.Required, Create: `${var.source_connection_container_id}`},
		"vault_details":                           acctest.RepresentationGroup{RepType: acctest.Required, Group: DatabaseMigrationMigrationVaultDetailsRepresentation},
	}

	migrationRepresentationNoSshMig = map[string]interface{}{
		"compartment_id":                acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"source_database_connection_id": acctest.Representation{RepType: acctest.Required, Create: `${var.source_connection_container}`},
		"target_database_connection_id": acctest.Representation{RepType: acctest.Required, Create: `${var.target_connection_nossh_id}`},
		"type":                          acctest.Representation{RepType: acctest.Required, Create: `OFFLINE`},
		"exclude_objects":               acctest.RepresentationGroup{RepType: acctest.Optional, Group: migrationExcludeObjectsRepresentation},
		"data_transfer_medium_details":  acctest.RepresentationGroup{RepType: acctest.Required, Group: migrationDataTransferMediumDetailsRepresentation},
		"datapump_settings":             acctest.RepresentationGroup{RepType: acctest.Required, Group: migrationDatapumpSettingsNoSshRepresentation},
		"display_name":                  acctest.Representation{RepType: acctest.Optional, Create: `TF_OFFLINE_MIG`, Update: `TF_OFFLINE_MIG2`},
		"vault_details":                 acctest.RepresentationGroup{RepType: acctest.Required, Group: DatabaseMigrationMigrationVaultDetailsRepresentation},
		"dump_transfer_details":         acctest.RepresentationGroup{RepType: acctest.Required, Group: migrationDumpTransferRepresentationNoSsh},
	}

	migrationDataTransferMediumDetailsRepresentation = map[string]interface{}{
		"object_storage_details": acctest.RepresentationGroup{RepType: acctest.Required, Group: migrationDataTransferMediumDetailsObjectStorageDetailsRepresentation},
	}
	DatabaseMigrationMigrationAdvisorSettingsRepresentation = map[string]interface{}{
		"is_ignore_errors": acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"is_skip_advisor":  acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `false`},
	}
	DatabaseMigrationMigrationDataTransferMediumDetailsRepresentation = map[string]interface{}{
		"database_link_details":  acctest.RepresentationGroup{RepType: acctest.Optional, Group: migrationDataTransferMediumDetailsDatabaseLinkDetailsRepresentation},
		"object_storage_details": acctest.RepresentationGroup{RepType: acctest.Optional, Group: migrationDataTransferMediumDetailsObjectStorageDetailsRepresentation},
	}
	DatabaseMigrationMigrationDataTransferMediumDetailsV2Representation = map[string]interface{}{
		"type": acctest.Representation{RepType: acctest.Required, Create: `OBJECT_STORAGE`, Update: `OBJECT_STORAGE`},
	}
	DatabaseMigrationMigrationDatapumpSettingsRepresentation = map[string]interface{}{
		"data_pump_parameters":    acctest.RepresentationGroup{RepType: acctest.Optional, Group: DatabaseMigrationMigrationDatapumpSettingsDataPumpParametersRepresentation},
		"export_directory_object": acctest.RepresentationGroup{RepType: acctest.Optional, Group: migrationDatapumpSettingsExportDirectoryObjectRepresentation},
		//"import_directory_object": acctest.RepresentationGroup{RepType: acctest.Optional, Group: migrationDatapumpSettingsImportDirectoryObjectRepresentation}, // Import Directory Object is not supported for Migrations using an Autonomous database as target connection.
		"job_mode":        acctest.Representation{RepType: acctest.Optional, Create: `SCHEMA`, Update: `SCHEMA`},
		"metadata_remaps": acctest.RepresentationGroup{RepType: acctest.Optional, Group: migrationDatapumpSettingsMetadataRemapsRepresentation},
	}
	DatabaseMigrationMigrationDumpTransferDetailsRepresentation = map[string]interface{}{
		"source": acctest.RepresentationGroup{RepType: acctest.Optional, Group: migrationSourceHostDumpTransferDetailsRepresentation},
		"target": acctest.RepresentationGroup{RepType: acctest.Optional, Group: migrationTargetHostDumpTransferDetailsRepresentation},
	}
	migrationDataTransferMediumDetailsRepresentationBeforeCPAT = map[string]interface{}{
		"database_link_details":  acctest.RepresentationGroup{RepType: acctest.Optional, Group: migrationDataTransferMediumDetailsDatabaseLinkDetailsRepresentation},
		"object_storage_details": acctest.RepresentationGroup{RepType: acctest.Required, Group: migrationDataTransferMediumDetailsObjectStorageDetailsRepresentation},
	}
	migrationDatapumpSettingsRepresentation = map[string]interface{}{
		"export_directory_object": acctest.RepresentationGroup{RepType: acctest.Required, Group: migrationDatapumpSettingsExportDirectoryObjectRepresentation},
		"metadata_remaps":         acctest.RepresentationGroup{RepType: acctest.Required, Group: migrationDatapumpSettingsMetadataRemapsRepresentation},
	}
	migrationDatapumpSettingsNoSshRepresentation = map[string]interface{}{
		"export_directory_object": acctest.RepresentationGroup{RepType: acctest.Required, Group: DatabaseMigrationMigrationDatapumpSettingsExportDirectoryObjectRepresentation},
		"import_directory_object": acctest.RepresentationGroup{RepType: acctest.Optional, Group: migrationDatapumpSettingsImportDirectoryObjectRepresentation},
		"metadata_remaps":         acctest.RepresentationGroup{RepType: acctest.Optional, Group: migrationDatapumpSettingsMetadataRemapsRepresentation},
	}
	migrationExcludeObjectsRepresentation = map[string]interface{}{
		"object": acctest.Representation{RepType: acctest.Required, Create: `.*`, Update: `object2`},
		"owner":  acctest.Representation{RepType: acctest.Required, Create: `owner`, Update: `owner2`},
		"is_omit_excluded_table_from_replication": acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"type": acctest.Representation{RepType: acctest.Optional, Create: `ALL`, Update: `TABLE`},
	}

	migrationDumpTransferRepresentationNoSsh = map[string]interface{}{
		"source": acctest.RepresentationGroup{RepType: acctest.Optional, Group: migrationSourceHostDumpTransferDetailsRepresentationNoSsh},
		"target": acctest.RepresentationGroup{RepType: acctest.Optional, Group: migrationTargetHostDumpTransferDetailsRepresentationNoSsh},
	}
	migrationSourceHostDumpTransferDetailsRepresentationNoSsh = map[string]interface{}{
		"kind":            acctest.Representation{RepType: acctest.Required, Create: `OCI_CLI`, Update: `OCI_CLI`},
		"oci_home":        acctest.Representation{RepType: acctest.Optional, Create: `ociHome`, Update: `ociHome2`},
		"wallet_location": acctest.Representation{RepType: acctest.Optional, Create: `walletLocation`, Update: `walletLocation2`},
	}
	migrationTargetHostDumpTransferDetailsRepresentationNoSsh = map[string]interface{}{
		"kind":            acctest.Representation{RepType: acctest.Required, Create: `OCI_CLI`, Update: `OCI_CLI`},
		"oci_home":        acctest.Representation{RepType: acctest.Optional, Create: `ociHome`, Update: `ociHome2`},
		"wallet_location": acctest.Representation{RepType: acctest.Optional, Create: `walletLocation`, Update: `walletLocation2`},
	}
	migrationDumpTransferRepresentation = map[string]interface{}{
		"source": acctest.RepresentationGroup{RepType: acctest.Optional, Group: migrationSourceHostDumpTransferDetailsRepresentation},
		"target": acctest.RepresentationGroup{RepType: acctest.Optional, Group: migrationTargetHostDumpTransferDetailsRepresentation},
	}
	migrationSourceHostDumpTransferDetailsRepresentation = map[string]interface{}{
		"kind":     acctest.Representation{RepType: acctest.Required, Create: `OCI_CLI`, Update: `OCI_CLI`},
		"oci_home": acctest.Representation{RepType: acctest.Optional, Create: `/path/to/ociCli`, Update: `/path/to/ociCli`},
	}
	migrationTargetHostDumpTransferDetailsRepresentation = map[string]interface{}{
		"kind":     acctest.Representation{RepType: acctest.Required, Create: `OCI_CLI`, Update: `OCI_CLI`},
		"oci_home": acctest.Representation{RepType: acctest.Optional, Create: `ociHome`, Update: `ociHome2`},
	}
	migrationGoldenGateDetailsRepresentation = map[string]interface{}{
		"hub":      acctest.RepresentationGroup{RepType: acctest.Required, Group: migrationGoldenGateDetailsHubRepresentation},
		"settings": acctest.RepresentationGroup{RepType: acctest.Optional, Group: migrationGoldenGateDetailsSettingsRepresentation},
	}
	DatabaseMigrationMigrationIncludeObjectsRepresentation = map[string]interface{}{
		"object": acctest.Representation{RepType: acctest.Required, Create: `object`, Update: `object2`},
		"owner":  acctest.Representation{RepType: acctest.Required, Create: `owner`, Update: `owner2`},
		"is_omit_excluded_table_from_replication": acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"type": acctest.Representation{RepType: acctest.Optional, Create: `type`, Update: `type2`},
	}
	DatabaseMigrationMigrationGoldenGateServiceDetailsRepresentation = map[string]interface{}{
		"settings": acctest.RepresentationGroup{RepType: acctest.Required, Group: DatabaseMigrationMigrationGoldenGateServiceDetailsSettingsRepresentation},
	}
	DatabaseMigrationMigrationVaultDetailsRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"key_id":         acctest.Representation{RepType: acctest.Required, Create: `${var.kms_key_id}`},
		"vault_id":       acctest.Representation{RepType: acctest.Required, Create: `${var.kms_vault_id}`},
	}
	migrationDataTransferMediumDetailsDatabaseLinkDetailsRepresentation = map[string]interface{}{
		"name": acctest.Representation{RepType: acctest.Required, Create: `name`, Update: `name2`},
	}
	migrationDataTransferMediumDetailsObjectStorageDetailsRepresentation = map[string]interface{}{
		"bucket":    acctest.Representation{RepType: acctest.Required, Create: `${var.bucket_id}`, Update: `${var.bucket_id}`},
		"namespace": acctest.Representation{RepType: acctest.Required, Create: `ax5cpn0vohdh`, Update: `ax5cpn0vohdh`},
	}
	DatabaseMigrationMigrationDataTransferMediumDetailsV2ObjectStorageBucketRepresentation = map[string]interface{}{
		"bucket":    acctest.Representation{RepType: acctest.Optional, Create: `${var.bucket_id}`, Update: `${var.bucket_id}`},
		"namespace": acctest.Representation{RepType: acctest.Optional, Create: `ax5cpn0vohdh`, Update: `ax5cpn0vohdh`},
	}
	DatabaseMigrationMigrationDatapumpSettingsDataPumpParametersRepresentation = map[string]interface{}{
		"estimate":                  acctest.Representation{RepType: acctest.Optional, Create: `BLOCKS`, Update: `STATISTICS`},
		"exclude_parameters":        acctest.Representation{RepType: acctest.Optional, Create: [3]string{`INDEX`, `MATERIALIZED_VIEW`, `MATERIALIZED_VIEW_LOG`}, Update: [1]string{`excludeParameters2`}},
		"export_parallelism_degree": acctest.Representation{RepType: acctest.Optional, Create: `10`, Update: `11`},
		"import_parallelism_degree": acctest.Representation{RepType: acctest.Optional, Create: `10`, Update: `11`},
		"is_cluster":                acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"table_exists_action":       acctest.Representation{RepType: acctest.Optional, Create: `TRUNCATE`, Update: `REPLACE`},
	}
	migrationDatapumpSettingsExportDirectoryObjectRepresentation = map[string]interface{}{
		"name": acctest.Representation{RepType: acctest.Required, Create: `test_export_dir`, Update: `test_export_dir`},
		"path": acctest.Representation{RepType: acctest.Required, Create: `/u01/app/oracle/product/19.0.0.0/dbhome_1/rdbms/log`, Update: `/u01/app/oracle/product/19.0.0.0/dbhome_1/rdbms/log`},
	}
	DatabaseMigrationMigrationDatapumpSettingsExportDirectoryObjectRepresentation = map[string]interface{}{
		"name": acctest.Representation{RepType: acctest.Required, Create: `test_export_dir`, Update: `test_export_dir`},
		"path": acctest.Representation{RepType: acctest.Required, Create: `/u01/app/oracle/product/19.0.0.0/dbhome_1/rdbms/log`, Update: `/u01/app/oracle/product/19.0.0.0/dbhome_1/rdbms/log`},
	}
	migrationDatapumpSettingsImportDirectoryObjectRepresentation = map[string]interface{}{
		"name": acctest.Representation{RepType: acctest.Required, Create: `dumpdir`, Update: `dumpdir`},
		"path": acctest.Representation{RepType: acctest.Optional, Create: `/u01/app/oracle/dumpdir`, Update: `/u01/app/oracle/dumpdir`},
	}
	migrationDatapumpSettingsMetadataRemapsRepresentation = map[string]interface{}{
		"new_value": acctest.Representation{RepType: acctest.Required, Create: `DATA`, Update: `DATA`},
		"old_value": acctest.Representation{RepType: acctest.Required, Create: `USERS`, Update: `USERS`},
		"type":      acctest.Representation{RepType: acctest.Required, Create: `TABLESPACE`, Update: `TABLESPACE`},
	}
	migrationGoldenGateDetailsHubRepresentation = map[string]interface{}{
		"rest_admin_credentials": acctest.RepresentationGroup{RepType: acctest.Required, Group: migrationGoldenGateDetailsHubRestAdminCredentialsRepresentation},

		"source_container_db_admin_credentials": acctest.RepresentationGroup{RepType: acctest.Required, Group: migrationGoldenGateDetailsHubSourceContainerDbAdminCredentialsRepresentation},
		"source_db_admin_credentials":           acctest.RepresentationGroup{RepType: acctest.Required, Group: migrationGoldenGateDetailsHubSourceDbAdminCredentialsRepresentation},
		"target_db_admin_credentials":           acctest.RepresentationGroup{RepType: acctest.Required, Group: migrationGoldenGateDetailsHubTargetDbAdminCredentialsRepresentation},
		"url":                                   acctest.Representation{RepType: acctest.Required, Create: `https://138.1.36.139`, Update: `https://138.1.36.139`},
		"source_microservices_deployment_name":  acctest.Representation{RepType: acctest.Required, Create: `Marketplace`},
		"target_microservices_deployment_name":  acctest.Representation{RepType: acctest.Required, Create: `Marketplace`},
	}
	migrationGoldenGateDetailsSettingsRepresentation = map[string]interface{}{
		"acceptable_lag": acctest.Representation{RepType: acctest.Optional, Create: `10`, Update: `11`},
		"extract":        acctest.RepresentationGroup{RepType: acctest.Optional, Group: migrationGoldenGateDetailsSettingsExtractRepresentation},
	}
	DatabaseMigrationMigrationGoldenGateServiceDetailsSettingsRepresentation = map[string]interface{}{
		"acceptable_lag": acctest.Representation{RepType: acctest.Optional, Create: `10`, Update: `11`},
		"extract":        acctest.RepresentationGroup{RepType: acctest.Optional, Group: DatabaseMigrationMigrationGoldenGateServiceDetailsSettingsExtractRepresentation},
	}
	DatabaseMigrationMigrationGoldenGateServiceDetailsSourceContainerDbCredentialsRepresentation = map[string]interface{}{
		"password": acctest.Representation{RepType: acctest.Required, Create: `BEstrO0ng_#11`, Update: `BEstrO0ng_#12`},
		"username": acctest.Representation{RepType: acctest.Required, Create: `username`, Update: `username2`},
	}
	DatabaseMigrationMigrationGoldenGateServiceDetailsSourceDbCredentialsRepresentation = map[string]interface{}{
		"password": acctest.Representation{RepType: acctest.Required, Create: `BEstrO0ng_#11`, Update: `BEstrO0ng_#12`},
		"username": acctest.Representation{RepType: acctest.Required, Create: `username`, Update: `username2`},
	}
	DatabaseMigrationMigrationGoldenGateServiceDetailsTargetDbCredentialsRepresentation = map[string]interface{}{
		"password": acctest.Representation{RepType: acctest.Required, Create: `BEstrO0ng_#11`, Update: `BEstrO0ng_#12`},
		"username": acctest.Representation{RepType: acctest.Required, Create: `username`, Update: `username2`},
	}
	DatabaseMigrationMigrationDataTransferMediumDetailsDatabaseLinkDetailsWalletBucketRepresentation = map[string]interface{}{
		"bucket":    acctest.Representation{RepType: acctest.Required, Create: `bucket`, Update: `bucket2`},
		"namespace": acctest.Representation{RepType: acctest.Required, Create: `namespace`, Update: `namespace2`},
	}
	migrationGoldenGateDetailsHubRestAdminCredentialsRepresentation = map[string]interface{}{
		"password": acctest.Representation{RepType: acctest.Required, Create: `ytW8z_.vZMwdBk0y`, Update: `ytW8z_.vZMwdBk0y`},
		"username": acctest.Representation{RepType: acctest.Required, Create: `oggadmin`, Update: `oggadmin`},
	}
	migrationGoldenGateDetailsHubSourceDbAdminCredentialsRepresentation = map[string]interface{}{
		"password": acctest.Representation{RepType: acctest.Required, Create: `DMS-pswd-2023#`, Update: `DMS-pswd-2023#`},
		"username": acctest.Representation{RepType: acctest.Required, Create: `ggadmin`, Update: `ggadmin`},
	}
	DatabaseMigrationMigrationGoldenGateDetailsHubSourceContainerDbAdminCredentialsRepresentation = map[string]interface{}{
		"password": acctest.Representation{RepType: acctest.Required, Create: `BEstrO0ng_#11`, Update: `BEstrO0ng_#12`},
		"username": acctest.Representation{RepType: acctest.Required, Create: `username`, Update: `username2`},
	}
	DatabaseMigrationMigrationGoldenGateDetailsHubSourceDbAdminCredentialsRepresentation = map[string]interface{}{
		"password": acctest.Representation{RepType: acctest.Required, Create: `BEstrO0ng_#11`, Update: `BEstrO0ng_#12`},
		"username": acctest.Representation{RepType: acctest.Required, Create: `username`, Update: `username2`},
	}
	DatabaseMigrationMigrationGoldenGateDetailsHubTargetDbAdminCredentialsRepresentation = map[string]interface{}{
		"password": acctest.Representation{RepType: acctest.Required, Create: `BEstrO0ng_#11`, Update: `BEstrO0ng_#12`},
		"username": acctest.Representation{RepType: acctest.Required, Create: `username`, Update: `username2`},
	}
	DatabaseMigrationMigrationGoldenGateDetailsSettingsExtractRepresentation = map[string]interface{}{
		"long_trans_duration": acctest.Representation{RepType: acctest.Optional, Create: `10`, Update: `11`},
		"performance_profile": acctest.Representation{RepType: acctest.Optional, Create: `LOW`, Update: `MEDIUM`},
	}
	DatabaseMigrationMigrationGoldenGateDetailsSettingsReplicatRepresentation = map[string]interface{}{
		"map_parallelism":       acctest.Representation{RepType: acctest.Optional, Create: `10`, Update: `11`},
		"max_apply_parallelism": acctest.Representation{RepType: acctest.Optional, Create: `10`, Update: `11`},
		"min_apply_parallelism": acctest.Representation{RepType: acctest.Optional, Create: `10`, Update: `11`},
	}
	DatabaseMigrationMigrationGoldenGateServiceDetailsSettingsExtractRepresentation = map[string]interface{}{
		"long_trans_duration": acctest.Representation{RepType: acctest.Optional, Create: `10`, Update: `11`},
		"performance_profile": acctest.Representation{RepType: acctest.Optional, Create: `LOW`, Update: `MEDIUM`},
	}
	DatabaseMigrationMigrationGoldenGateServiceDetailsSettingsReplicatRepresentation = map[string]interface{}{
		"map_parallelism":       acctest.Representation{RepType: acctest.Required, Create: `10`, Update: `11`},
		"max_apply_parallelism": acctest.Representation{RepType: acctest.Required, Create: `10`, Update: `11`},
		"min_apply_parallelism": acctest.Representation{RepType: acctest.Required, Create: `10`, Update: `11`},
	}
	migrationGoldenGateDetailsHubTargetDbAdminCredentialsRepresentation = map[string]interface{}{
		"password": acctest.Representation{RepType: acctest.Required, Create: `DMS-pswd-2023#`, Update: `DMS-pswd-2023#`},
		"username": acctest.Representation{RepType: acctest.Required, Create: `ggadmin`, Update: `ggadmin`},
	}
	migrationGoldenGateDetailsHubSourceContainerDbAdminCredentialsRepresentation = map[string]interface{}{
		"password": acctest.Representation{RepType: acctest.Required, Create: `DMS-pswd-2023#`, Update: `DMS-pswd-2023#`},
		"username": acctest.Representation{RepType: acctest.Required, Create: `c##ggadmin`, Update: `c##ggadmin`},
	}
	migrationGoldenGateDetailsSettingsExtractRepresentation = map[string]interface{}{
		"long_trans_duration": acctest.Representation{RepType: acctest.Optional, Create: `10`, Update: `11`},
		"performance_profile": acctest.Representation{RepType: acctest.Optional, Create: `LOW`, Update: `MEDIUM`},
	}
	migrationGoldenGateDetailsSettingsReplicatRepresentation = map[string]interface{}{
		"map_parallelism":       acctest.Representation{RepType: acctest.Required, Create: `1`, Update: `0`},
		"min_apply_parallelism": acctest.Representation{RepType: acctest.Required, Create: `1`, Update: `0`},
		"max_apply_parallelism": acctest.Representation{RepType: acctest.Required, Create: `1`, Update: `0`},
		"performance_profile":   acctest.Representation{RepType: acctest.Optional, Create: `LOW`, Update: `HIGH`},
	}

	DatabaseMigrationMigrationResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_database_migration_connection", "test_connection", acctest.Required, acctest.Create, connectionRepresentationTarget) +
		acctest.GenerateResourceFromRepresentationMap("oci_database_migration_connection", "test_connection_source", acctest.Required, acctest.Create, connectionRepresentationSource)

	DatabaseMigrationMigrationNoSshResourceDependencies = //acctest.GenerateResourceFromRepresentationMap("oci_database_migration_connection", "test_connection_usr_managed_oci", acctest.Required, acctest.Create, connectionRepresentationUserManagedOciTarget) +
	acctest.GenerateResourceFromRepresentationMap("oci_database_migration_connection", "test_connection_nossh_source", acctest.Required, acctest.Create, connectionRepresentationNoSshSource)
)

// issue-routing-tag: database_migration/default
func TestDatabaseMigrationMigrationResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseMigrationMigrationResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	//Asigning values to variables used on representations. Environment variables are taken by tf config file and used to create resources.

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithBlankDefault("compartment_id_for_update")
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	databaseSourceContainerId := utils.GetEnvSettingWithBlankDefault("source_connection_container_id")
	databaseSourceContainerStr := fmt.Sprintf("variable \"source_connection_container_id\" { default = \"%s\" }\n", databaseSourceContainerId)

	databaseSourcePDBId := utils.GetEnvSettingWithBlankDefault("source_connection_id")
	databaseSourcePDBStr := fmt.Sprintf("variable \"source_connection_id\" { default = \"%s\" }\n", databaseSourcePDBId)

	databaseTargetId := utils.GetEnvSettingWithBlankDefault("target_connection_id")
	databaseTargetStr := fmt.Sprintf("variable \"target_connection_id\" { default = \"%s\" }\n", databaseTargetId)
	databaseSourceId := utils.GetEnvSettingWithBlankDefault("database_source")
	databaseSourceStr := fmt.Sprintf("variable \"database_source\" { default = \"%s\" }\n", databaseSourceId)

	databaseSourceFDB := utils.GetEnvSettingWithBlankDefault("source_connection_plus_id")
	databaseSourceFDBStr := fmt.Sprintf("variable \"source_connection_plus_id\" { default = \"%s\" }\n", databaseSourceFDB)

	databaseSourceNoSSHDBId := utils.GetEnvSettingWithBlankDefault("database_container_source_id")
	databaseSourceNoSSHDBStr := fmt.Sprintf("variable \"database_container_source_id\" { default = \"%s\" }\n", databaseSourceNoSSHDBId)

	databaseSourceNoSsh := utils.GetEnvSettingWithBlankDefault("source_connection_container")
	databaseSourceNoSshStr := fmt.Sprintf("variable \"source_connection_container\" { default = \"%s\" }\n", databaseSourceNoSsh)

	databaseSourceContainerFDB := utils.GetEnvSettingWithBlankDefault("source_connection_container_plus_id")
	databaseSourceContainerFDBStr := fmt.Sprintf("variable \"source_connection_container_plus_id\" { default = \"%s\" }\n", databaseSourceContainerFDB)

	databaseTargetNoSshId := utils.GetEnvSettingWithBlankDefault("target_connection_nossh_id")
	databaseTargetNoSshStr := fmt.Sprintf("variable \"target_connection_nossh_id\" { default = \"%s\" }\n", databaseTargetNoSshId)

	databaseTargetNoAutId := utils.GetEnvSettingWithBlankDefault("target_connection_no_autonomous_id")
	databaseTargetNoAutStr := fmt.Sprintf("variable \"target_connection_no_autonomous_id\" { default = \"%s\" }\n", databaseTargetNoAutId)

	databaseTargetRDSId := utils.GetEnvSettingWithBlankDefault("target_connection_rds_id")
	databaseTargetRDSStr := fmt.Sprintf("variable \"target_connection_rds_id\" { default = \"%s\" }\n", databaseTargetRDSId)

	fileStorageMountId := utils.GetEnvSettingWithBlankDefault("file_storage_mount")
	fileStorageMountStr := fmt.Sprintf("variable \"file_storage_mount\" { default = \"%s\" }\n", fileStorageMountId)

	sourceDBId := utils.GetEnvSettingWithBlankDefault("database_source_id")
	sourceDBStr := fmt.Sprintf("variable \"database_source_id\" { default = \"%s\" }\n", sourceDBId)

	targetDBId := utils.GetEnvSettingWithBlankDefault("database_id")
	targetDBStr := fmt.Sprintf("variable \"database_id\" { default = \"%s\" }\n", targetDBId)

	kmsVaultId := utils.GetEnvSettingWithBlankDefault("kms_vault_id")
	kmsVaultIdVariableStr := fmt.Sprintf("\nvariable \"kms_vault_id\" { default = \"%s\" }\n", kmsVaultId)

	vcnId := utils.GetEnvSettingWithBlankDefault("vcn_id")
	vcnIdVariableStr := fmt.Sprintf("variable \"vcn_id\" { default = \"%s\" }\n", vcnId)

	subnetId := utils.GetEnvSettingWithBlankDefault("subnet_id")
	subnetIdStr := fmt.Sprintf("variable \"subnet_id\" { default = \"%s\" }\n", subnetId)

	kmsKeyId := utils.GetEnvSettingWithBlankDefault("kms_key_id")
	KmsKeyIdVariableStr := fmt.Sprintf("\nvariable \"kms_key_id\" { default = \"%s\" }\n", kmsKeyId)

	bucketId := utils.GetEnvSettingWithBlankDefault("bucket_id")
	bucketIdVariableStr := fmt.Sprintf("\nvariable \"bucket_id\" { default = \"%s\" }\n", bucketId)

	sshKey := utils.GetEnvSettingWithBlankDefault("ssh_key")
	sshKeyStr := fmt.Sprintf("variable \"ssh_key\" {\n type = \"string\"\n default = <<EOF\n%s\nEOF \n}\n", sshKey)

	resourceName := "oci_database_migration_migration.test_migration"
	resourceNoSshName := "oci_database_migration_migration.test_nossh_migration"
	datasourceName := "data.oci_database_migration_migrations.test_migrations"
	singularDatasourceName := "data.oci_database_migration_migration.test_migration"
	var resId, resId2, resId3 string

	//Initial configuration to build the test set
	acctest.SaveConfigContent(config+compartmentIdVariableStr+databaseSourceStr+databaseTargetStr+kmsVaultIdVariableStr+vcnIdVariableStr+subnetIdStr+KmsKeyIdVariableStr+databaseSourceContainerStr+bucketIdVariableStr+sshKeyStr+databaseSourceFDBStr+databaseSourceContainerFDBStr+sourceDBStr+targetDBStr+databaseTargetNoSshStr+databaseSourceNoSshStr+databaseSourceNoSSHDBStr+databaseTargetRDSStr+ //DatabaseMigrationMigrationResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_database_migration_migration", "test_migration", acctest.Optional, acctest.Create, DatabaseMigrationMigrationRepresentation)+
		DatabaseMigrationMigrationNoSshResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_database_migration_migration", "test_nossh_migration", acctest.Optional, acctest.Create, migrationRepresentationNoSshMig), "databasemigration", "migration", t)
	acctest.ResourceTest(t, testAccCheckDatabaseMigrationMigrationDestroy, []resource.TestStep{
		// verify Create
		{

			Config: config + compartmentIdVariableStr + databaseSourceStr + databaseTargetStr + kmsVaultIdVariableStr + vcnIdVariableStr + subnetIdStr + KmsKeyIdVariableStr + bucketIdVariableStr + sshKeyStr + databaseSourceContainerStr + databaseTargetNoSshStr + databaseSourceFDBStr + databaseSourceContainerFDBStr + sourceDBStr + targetDBStr + databaseSourcePDBStr + databaseTargetRDSStr + //DatabaseMigrationMigrationResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_migration_migration", "test_migration", acctest.Required, acctest.Create, migrationRepresentationMig),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName, "source_database_connection_id"),
				resource.TestCheckResourceAttrSet(resourceName, "target_database_connection_id"),
				resource.TestCheckResourceAttr(resourceName, "type", "ONLINE"),
				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + databaseSourceStr + databaseTargetStr + kmsVaultIdVariableStr + vcnIdVariableStr + subnetIdStr + KmsKeyIdVariableStr + bucketIdVariableStr + sshKeyStr + databaseSourceContainerStr + databaseTargetRDSStr, //+ DatabaseMigrationMigrationResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + databaseSourceStr + databaseTargetStr + databaseTargetNoAutStr + kmsVaultIdVariableStr + vcnIdVariableStr + subnetIdStr + KmsKeyIdVariableStr + bucketIdVariableStr + sshKeyStr + databaseSourceContainerStr + fileStorageMountStr + databaseTargetNoSshStr + databaseSourceFDBStr + databaseSourceContainerFDBStr + sourceDBStr + targetDBStr + databaseSourcePDBStr + databaseSourceNoSshStr + databaseSourceNoSSHDBStr + databaseTargetRDSStr + //DatabaseMigrationMigrationResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_migration_migration", "test_migration", acctest.Optional, acctest.Create, DatabaseMigrationMigrationRDSRepresentation) +
				acctest.GenerateResourceFromRepresentationMap("oci_database_migration_migration", "test_nossh_migration", acctest.Optional, acctest.Create, migrationRepresentationNoSshMig),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceNoSshName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceNoSshName, "display_name", "TF_OFFLINE_MIG"),
				resource.TestCheckResourceAttr(resourceNoSshName, "type", "OFFLINE"),
				resource.TestCheckResourceAttrSet(resourceNoSshName, "id"),
				resource.TestCheckResourceAttrSet(resourceNoSshName, "source_database_connection_id"),
				resource.TestCheckResourceAttrSet(resourceNoSshName, "state"),
				resource.TestCheckResourceAttrSet(resourceNoSshName, "target_database_connection_id"),
				resource.TestCheckResourceAttrSet(resourceNoSshName, "time_created"),
				resource.TestCheckResourceAttr(resourceNoSshName, "dump_transfer_details.#", "1"),
				resource.TestCheckResourceAttr(resourceNoSshName, "dump_transfer_details.0.source.#", "1"),
				resource.TestCheckResourceAttr(resourceNoSshName, "dump_transfer_details.0.source.0.kind", "OCI_CLI"),
				resource.TestCheckResourceAttr(resourceNoSshName, "dump_transfer_details.0.source.0.oci_home", "ociHome"),
				resource.TestCheckResourceAttr(resourceNoSshName, "dump_transfer_details.0.source.0.wallet_location", "walletLocation"),
				resource.TestCheckResourceAttr(resourceNoSshName, "dump_transfer_details.0.target.#", "1"),
				resource.TestCheckResourceAttr(resourceNoSshName, "dump_transfer_details.0.target.0.kind", "OCI_CLI"),
				resource.TestCheckResourceAttr(resourceNoSshName, "dump_transfer_details.0.target.0.oci_home", "ociHome"),
				resource.TestCheckResourceAttr(resourceNoSshName, "dump_transfer_details.0.target.0.wallet_location", "walletLocation"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "data_transfer_medium_details_v2.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "data_transfer_medium_details_v2.0.type", "OBJECT_STORAGE"),
				resource.TestCheckResourceAttr(resourceName, "datapump_settings.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "datapump_settings.0.export_directory_object.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "datapump_settings.0.export_directory_object.0.name", "test_export_dir"),
				resource.TestCheckResourceAttr(resourceName, "datapump_settings.0.export_directory_object.0.path", "/u01/app/oracle/product/19.0.0.0/dbhome_1/rdbms/log"),
				resource.TestCheckResourceAttr(resourceName, "datapump_settings.0.metadata_remaps.#", "1"),

				resource.TestCheckResourceAttr(resourceName, "display_name", "TFtestOnline1"),

				resource.TestCheckResourceAttr(resourceName, "dump_transfer_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "dump_transfer_details.0.source.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "dump_transfer_details.0.source.0.kind", "OCI_CLI"),
				resource.TestCheckResourceAttr(resourceName, "dump_transfer_details.0.source.0.oci_home", "/path/to/ociCli"),
				resource.TestCheckResourceAttr(resourceName, "dump_transfer_details.0.target.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "dump_transfer_details.0.target.0.kind", "OCI_CLI"),
				resource.TestCheckResourceAttr(resourceName, "dump_transfer_details.0.target.0.oci_home", "ociHome"),
				resource.TestCheckResourceAttr(resourceName, "exclude_objects.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "exclude_objects.0.is_omit_excluded_table_from_replication", "false"),
				resource.TestCheckResourceAttr(resourceName, "exclude_objects.0.owner", "owner"),
				resource.TestCheckResourceAttr(resourceName, "exclude_objects.0.object", ".*"),
				resource.TestCheckResourceAttr(resourceName, "exclude_objects.0.type", "ALL"),
				acctest.CheckResourceSetContainsElementWithProperties(resourceName, "datapump_settings.0.metadata_remaps", map[string]string{
					"new_value": "DATA",
					"old_value": "USERS",
					"type":      "TABLESPACE",
				},
					[]string{}),
				resource.TestCheckResourceAttr(resourceName, "display_name", "TFtestOnline1"),
				resource.TestCheckResourceAttr(resourceName, "golden_gate_service_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "golden_gate_service_details.0.settings.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "golden_gate_service_details.0.settings.0.acceptable_lag", "10"),
				resource.TestCheckResourceAttr(resourceName, "golden_gate_service_details.0.settings.0.extract.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "golden_gate_service_details.0.settings.0.extract.0.long_trans_duration", "10"),
				resource.TestCheckResourceAttr(resourceName, "golden_gate_service_details.0.settings.0.extract.0.performance_profile", "LOW"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),

				resource.TestCheckResourceAttrSet(resourceName, "source_database_connection_id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "target_database_connection_id"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttr(resourceName, "type", "ONLINE"),
				resource.TestCheckResourceAttr(resourceName, "vault_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "vault_details.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName, "vault_details.0.key_id"),
				resource.TestCheckResourceAttrSet(resourceName, "vault_details.0.vault_id"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					if isEnableExportCompartment, _ := strconv.ParseBool(utils.GetEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
						if errExport := resourcediscovery.TestExportCompartmentWithResourceName(&resId, &compartmentId, resourceName); errExport != nil {
							return errExport
						}
					}
					resId2, err = acctest.FromInstanceState(s, resourceNoSshName, "id")
					if isEnableExportCompartment, _ := strconv.ParseBool(utils.GetEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
						if errExport := resourcediscovery.TestExportCompartmentWithResourceName(&resId2, &compartmentId, resourceNoSshName); errExport != nil {
							return errExport
						}
					}
					return err
				},
			),
		},

		// verify Update to the compartment (the compartment will be switched back in the next step)
		{
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + databaseSourceStr + fileStorageMountStr + databaseTargetStr + databaseTargetNoAutStr + kmsVaultIdVariableStr + vcnIdVariableStr + subnetIdStr + KmsKeyIdVariableStr + bucketIdVariableStr + sshKeyStr + databaseSourceContainerStr + databaseSourceFDBStr + databaseSourceContainerFDBStr + sourceDBStr + targetDBStr + databaseTargetNoSshStr + databaseSourcePDBStr + databaseSourceNoSshStr + databaseSourceNoSSHDBStr + databaseTargetRDSStr + //DatabaseMigrationMigrationResourceDependencies + //DatabaseMigrationMigrationResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_migration_migration", "test_migration", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(DatabaseMigrationMigrationRDSRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})) +
				acctest.GenerateResourceFromRepresentationMap("oci_database_migration_migration", "test_nossh_migration", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(migrationRepresentationNoSshMig, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "data_transfer_medium_details_v2.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "data_transfer_medium_details_v2.0.type", "OBJECT_STORAGE"),
				resource.TestCheckResourceAttr(resourceName, "datapump_settings.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "datapump_settings.0.export_directory_object.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "datapump_settings.0.export_directory_object.0.name", "test_export_dir"),
				resource.TestCheckResourceAttr(resourceName, "datapump_settings.0.export_directory_object.0.path", "/u01/app/oracle/product/19.0.0.0/dbhome_1/rdbms/log"),
				resource.TestCheckResourceAttr(resourceName, "datapump_settings.0.metadata_remaps.#", "1"),

				resource.TestCheckResourceAttr(resourceName, "display_name", "TFtestOnline1"),

				resource.TestCheckResourceAttr(resourceName, "dump_transfer_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "dump_transfer_details.0.source.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "dump_transfer_details.0.source.0.kind", "OCI_CLI"),
				resource.TestCheckResourceAttr(resourceName, "dump_transfer_details.0.source.0.oci_home", "/path/to/ociCli"),
				resource.TestCheckResourceAttr(resourceName, "dump_transfer_details.0.target.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "dump_transfer_details.0.target.0.kind", "OCI_CLI"),
				resource.TestCheckResourceAttr(resourceName, "dump_transfer_details.0.target.0.oci_home", "ociHome"),
				resource.TestCheckResourceAttr(resourceName, "exclude_objects.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "exclude_objects.0.is_omit_excluded_table_from_replication", "false"),
				resource.TestCheckResourceAttr(resourceName, "exclude_objects.0.owner", "owner"),
				resource.TestCheckResourceAttr(resourceName, "exclude_objects.0.object", ".*"),
				resource.TestCheckResourceAttr(resourceName, "exclude_objects.0.type", "ALL"),
				acctest.CheckResourceSetContainsElementWithProperties(resourceName, "datapump_settings.0.metadata_remaps", map[string]string{
					"new_value": "DATA",
					"old_value": "USERS",
					"type":      "TABLESPACE",
				},
					[]string{}),
				resource.TestCheckResourceAttr(resourceName, "display_name", "TFtestOnline1"),
				resource.TestCheckResourceAttr(resourceName, "golden_gate_service_details.#", "1"),

				resource.TestCheckResourceAttr(resourceName, "golden_gate_service_details.0.settings.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "golden_gate_service_details.0.settings.0.acceptable_lag", "10"),
				resource.TestCheckResourceAttr(resourceName, "golden_gate_service_details.0.settings.0.extract.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "golden_gate_service_details.0.settings.0.extract.0.long_trans_duration", "10"),
				resource.TestCheckResourceAttr(resourceName, "golden_gate_service_details.0.settings.0.extract.0.performance_profile", "LOW"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),

				resource.TestCheckResourceAttrSet(resourceName, "source_database_connection_id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "target_database_connection_id"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttr(resourceName, "type", "ONLINE"),
				resource.TestCheckResourceAttr(resourceName, "vault_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "vault_details.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName, "vault_details.0.key_id"),
				resource.TestCheckResourceAttrSet(resourceName, "vault_details.0.vault_id"),

				resource.TestCheckResourceAttr(resourceNoSshName, "compartment_id", compartmentIdU),

				func(s *terraform.State) (err error) {
					resId3, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId3 {
						return fmt.Errorf("update to the compartment: resource %s recreated when it was supposed to be updated", resourceName)
					}
					resId3, err = acctest.FromInstanceState(s, resourceNoSshName, "id")
					if resId2 != resId3 {
						return fmt.Errorf("update to the compartment: resource %s recreated when it was supposed to be updated", resourceNoSshName)
					}
					return err
				},
			),
		},

		// verify updates to updatable parameters
		{

			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + databaseSourceStr + databaseTargetStr + databaseTargetNoAutStr + kmsVaultIdVariableStr + vcnIdVariableStr + subnetIdStr + KmsKeyIdVariableStr + bucketIdVariableStr + sshKeyStr + databaseSourceContainerStr + databaseSourceFDBStr + databaseSourceContainerFDBStr + sourceDBStr + targetDBStr + databaseTargetNoSshStr + databaseSourcePDBStr + databaseSourceNoSshStr + databaseSourceNoSSHDBStr + databaseTargetRDSStr + //DatabaseMigrationMigrationResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_migration_migration", "test_migration", acctest.Optional, acctest.Update, DatabaseMigrationMigrationRDSRepresentation) +
				acctest.GenerateResourceFromRepresentationMap("oci_database_migration_migration", "test_nossh_migration", acctest.Optional, acctest.Update, migrationRepresentationNoSshMig),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceNoSshName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceNoSshName, "display_name", "TF_OFFLINE_MIG2"),
				resource.TestCheckResourceAttr(resourceNoSshName, "type", "OFFLINE"),
				resource.TestCheckResourceAttrSet(resourceNoSshName, "id"),
				resource.TestCheckResourceAttrSet(resourceNoSshName, "source_database_connection_id"),
				resource.TestCheckResourceAttrSet(resourceNoSshName, "state"),
				resource.TestCheckResourceAttrSet(resourceNoSshName, "target_database_connection_id"),
				resource.TestCheckResourceAttrSet(resourceNoSshName, "time_created"),
				resource.TestCheckResourceAttr(resourceNoSshName, "dump_transfer_details.#", "1"),
				resource.TestCheckResourceAttr(resourceNoSshName, "dump_transfer_details.0.source.#", "1"),
				resource.TestCheckResourceAttr(resourceNoSshName, "dump_transfer_details.0.source.0.kind", "OCI_CLI"),
				resource.TestCheckResourceAttr(resourceNoSshName, "dump_transfer_details.0.source.0.oci_home", "ociHome2"),
				resource.TestCheckResourceAttr(resourceNoSshName, "dump_transfer_details.0.source.0.wallet_location", "walletLocation2"),
				resource.TestCheckResourceAttr(resourceNoSshName, "dump_transfer_details.0.target.#", "1"),
				resource.TestCheckResourceAttr(resourceNoSshName, "dump_transfer_details.0.target.0.kind", "OCI_CLI"),
				resource.TestCheckResourceAttr(resourceNoSshName, "dump_transfer_details.0.target.0.oci_home", "ociHome2"),
				resource.TestCheckResourceAttr(resourceNoSshName, "dump_transfer_details.0.target.0.wallet_location", "walletLocation2"),

				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "data_transfer_medium_details_v2.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "data_transfer_medium_details_v2.0.type", "OBJECT_STORAGE"),
				resource.TestCheckResourceAttr(resourceName, "datapump_settings.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "datapump_settings.0.export_directory_object.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "datapump_settings.0.export_directory_object.0.name", "test_export_dir"),
				resource.TestCheckResourceAttr(resourceName, "datapump_settings.0.export_directory_object.0.path", "/u01/app/oracle/product/19.0.0.0/dbhome_1/rdbms/log"),
				resource.TestCheckResourceAttr(resourceName, "datapump_settings.0.metadata_remaps.#", "1"),

				resource.TestCheckResourceAttr(resourceName, "display_name", "TFtestOnline2"),

				resource.TestCheckResourceAttr(resourceName, "dump_transfer_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "dump_transfer_details.0.source.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "dump_transfer_details.0.source.0.kind", "OCI_CLI"),
				resource.TestCheckResourceAttr(resourceName, "dump_transfer_details.0.source.0.oci_home", "/path/to/ociCli"),
				resource.TestCheckResourceAttr(resourceName, "dump_transfer_details.0.target.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "dump_transfer_details.0.target.0.kind", "OCI_CLI"),
				resource.TestCheckResourceAttr(resourceName, "dump_transfer_details.0.target.0.oci_home", "ociHome2"),
				resource.TestCheckResourceAttr(resourceName, "exclude_objects.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "exclude_objects.0.is_omit_excluded_table_from_replication", "true"),
				resource.TestCheckResourceAttr(resourceName, "exclude_objects.0.owner", "owner2"),
				resource.TestCheckResourceAttr(resourceName, "exclude_objects.0.object", "object2"),
				resource.TestCheckResourceAttr(resourceName, "exclude_objects.0.type", "TABLE"),
				acctest.CheckResourceSetContainsElementWithProperties(resourceName, "datapump_settings.0.metadata_remaps", map[string]string{
					"new_value": "DATA",
					"old_value": "USERS",
					"type":      "TABLESPACE",
				},
					[]string{}),
				resource.TestCheckResourceAttr(resourceName, "golden_gate_service_details.#", "1"),

				resource.TestCheckResourceAttr(resourceName, "golden_gate_service_details.0.settings.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "golden_gate_service_details.0.settings.0.acceptable_lag", "11"),
				resource.TestCheckResourceAttr(resourceName, "golden_gate_service_details.0.settings.0.extract.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "golden_gate_service_details.0.settings.0.extract.0.long_trans_duration", "11"),
				resource.TestCheckResourceAttr(resourceName, "golden_gate_service_details.0.settings.0.extract.0.performance_profile", "MEDIUM"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),

				resource.TestCheckResourceAttrSet(resourceName, "source_database_connection_id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "target_database_connection_id"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttr(resourceName, "type", "ONLINE"),
				resource.TestCheckResourceAttr(resourceName, "vault_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "vault_details.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName, "vault_details.0.key_id"),
				resource.TestCheckResourceAttrSet(resourceName, "vault_details.0.vault_id"),

				func(s *terraform.State) (err error) {
					resId3, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId3 {
						return fmt.Errorf("updates to updatable parameters: resource %s recreated when it was supposed to be updated", resourceName)
					}
					resId3, err = acctest.FromInstanceState(s, resourceNoSshName, "id")
					if resId2 != resId3 {
						return fmt.Errorf("updates to updatable parameters: resource %s recreated when it was supposed to be updated", resourceNoSshName)
					}
					return err
				},
			),
		},
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_migration_migrations", "test_migrations", acctest.Optional, acctest.Update, DatabaseMigrationmigrationDataSourceRepresentation) +
				compartmentIdVariableStr + compartmentIdUVariableStr + databaseSourceStr + databaseTargetStr + databaseTargetNoAutStr + kmsVaultIdVariableStr + vcnIdVariableStr + subnetIdStr + KmsKeyIdVariableStr + bucketIdVariableStr + sshKeyStr + databaseSourceContainerStr + databaseSourceFDBStr + databaseSourceContainerFDBStr + sourceDBStr + targetDBStr + databaseTargetNoSshStr + databaseSourcePDBStr + databaseSourceNoSshStr + databaseSourceNoSSHDBStr + // DatabaseMigrationMigrationResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_migration_migration", "test_migration", acctest.Optional, acctest.Update, DatabaseMigrationMigrationRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),
				resource.TestCheckResourceAttr(datasourceName, "migration_collection.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_migration_migration", "test_migration", acctest.Required, acctest.Create, DatabaseMigrationmigrationSingularDataSourceRepresentation) +
				compartmentIdVariableStr + compartmentIdUVariableStr + databaseSourceStr + databaseTargetStr + databaseTargetNoAutStr + kmsVaultIdVariableStr + vcnIdVariableStr + subnetIdStr + KmsKeyIdVariableStr + bucketIdVariableStr + sshKeyStr + databaseSourceContainerStr + databaseSourceFDBStr + databaseSourceContainerFDBStr + sourceDBStr + targetDBStr + databaseTargetNoSshStr + databaseSourcePDBStr + databaseSourceNoSshStr + databaseSourceNoSSHDBStr + DatabaseMigrationMigrationResourceConfig, // DatabaseMigrationMigrationResourceDependencies
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "migration_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "credentials_secret_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "data_transfer_medium_details.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "data_transfer_medium_details.0.object_storage_details.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "data_transfer_medium_details.0.object_storage_details.0.bucket"),
				resource.TestCheckResourceAttr(singularDatasourceName, "data_transfer_medium_details.0.object_storage_details.0.namespace", "ax5cpn0vohdh"),
				resource.TestCheckResourceAttr(singularDatasourceName, "data_transfer_medium_details_v2.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "data_transfer_medium_details_v2.0.type", "OBJECT_STORAGE"),
				resource.TestCheckResourceAttr(singularDatasourceName, "datapump_settings.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "datapump_settings.0.data_pump_parameters.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "datapump_settings.0.data_pump_parameters.0.estimate", "STATISTICS"),
				resource.TestCheckResourceAttr(singularDatasourceName, "datapump_settings.0.data_pump_parameters.0.export_parallelism_degree", "11"),
				resource.TestCheckResourceAttr(singularDatasourceName, "datapump_settings.0.data_pump_parameters.0.import_parallelism_degree", "11"),
				resource.TestCheckResourceAttr(singularDatasourceName, "datapump_settings.0.data_pump_parameters.0.is_cluster", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "datapump_settings.0.data_pump_parameters.0.table_exists_action", "REPLACE"),
				resource.TestCheckResourceAttr(singularDatasourceName, "datapump_settings.0.export_directory_object.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "datapump_settings.0.export_directory_object.0.name", "test_export_dir"),
				resource.TestCheckResourceAttr(singularDatasourceName, "datapump_settings.0.export_directory_object.0.path", "/u01/app/oracle/product/19.0.0.0/dbhome_1/rdbms/log"),
				resource.TestCheckResourceAttr(singularDatasourceName, "datapump_settings.0.metadata_remaps.#", "1"),
				acctest.CheckResourceSetContainsElementWithProperties(resourceName, "datapump_settings.0.metadata_remaps", map[string]string{
					"new_value": "DATA",
					"old_value": "USERS",
					"type":      "TABLESPACE",
				},
					[]string{}),

				resource.TestCheckResourceAttr(singularDatasourceName, "dump_transfer_details.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "dump_transfer_details.0.source.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "dump_transfer_details.0.source.0.kind", "OCI_CLI"),
				resource.TestCheckResourceAttr(singularDatasourceName, "dump_transfer_details.0.source.0.oci_home", "/path/to/ociCli"),
				resource.TestCheckResourceAttr(singularDatasourceName, "dump_transfer_details.0.target.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "dump_transfer_details.0.target.0.kind", "OCI_CLI"),
				resource.TestCheckResourceAttr(singularDatasourceName, "dump_transfer_details.0.target.0.oci_home", "ociHome2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "exclude_objects.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "exclude_objects.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "exclude_objects.0.is_omit_excluded_table_from_replication", "true"),
				resource.TestCheckResourceAttr(resourceName, "exclude_objects.0.owner", "owner2"),
				resource.TestCheckResourceAttr(resourceName, "exclude_objects.0.object", "object2"),
				resource.TestCheckResourceAttr(resourceName, "exclude_objects.0.type", "TABLE"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "TFtestOnline2"),

				resource.TestCheckResourceAttr(singularDatasourceName, "golden_gate_details.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "golden_gate_details.0.hub.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "golden_gate_details.0.hub.0.rest_admin_credentials.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "golden_gate_details.0.hub.0.rest_admin_credentials.0.username", "oggadmin"),
				resource.TestCheckResourceAttr(singularDatasourceName, "golden_gate_details.0.hub.0.source_db_admin_credentials.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "golden_gate_details.0.hub.0.source_db_admin_credentials.0.username", "ggadmin"),
				resource.TestCheckResourceAttr(singularDatasourceName, "golden_gate_details.0.hub.0.target_db_admin_credentials.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "golden_gate_details.0.hub.0.target_db_admin_credentials.0.username", "ggadmin"),
				resource.TestCheckResourceAttr(singularDatasourceName, "golden_gate_details.0.hub.0.url", "https://138.1.36.139"),
				resource.TestCheckResourceAttr(singularDatasourceName, "golden_gate_details.0.settings.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "golden_gate_details.0.settings.0.acceptable_lag", "11"),
				resource.TestCheckResourceAttr(singularDatasourceName, "golden_gate_details.0.settings.0.extract.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "golden_gate_details.0.settings.0.extract.0.long_trans_duration", "11"),
				resource.TestCheckResourceAttr(singularDatasourceName, "golden_gate_details.0.settings.0.extract.0.performance_profile", "MEDIUM"),

				resource.TestCheckResourceAttr(singularDatasourceName, "golden_gate_service_details.0.settings.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "golden_gate_service_details.0.settings.0.acceptable_lag", "11"),
				resource.TestCheckResourceAttr(singularDatasourceName, "golden_gate_service_details.0.settings.0.extract.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "golden_gate_service_details.0.settings.0.extract.0.long_trans_duration", "11"),
				resource.TestCheckResourceAttr(singularDatasourceName, "golden_gate_service_details.0.settings.0.extract.0.performance_profile", "MEDIUM"),

				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
				resource.TestCheckResourceAttr(singularDatasourceName, "type", "ONLINE"),
				resource.TestCheckResourceAttr(singularDatasourceName, "vault_details.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "vault_details.0.compartment_id", compartmentId),
			),
		},
		// verify resource import
		{
			Config:            config + DatabaseMigrationMigrationRequiredOnlyResource,
			ImportState:       true,
			ImportStateVerify: true,
			ImportStateVerifyIgnore: []string{
				"csv_text",
				"golden_gate_details.0.hub.0.rest_admin_credentials.0.password",
				"golden_gate_details.0.hub.0.source_container_db_admin_credentials.0.password",
				"golden_gate_details.0.hub.0.source_db_admin_credentials.0.password",
				"golden_gate_details.0.hub.0.target_db_admin_credentials.0.password",
			},
			ResourceName: resourceName,
		},
	})
}

func testAccCheckDatabaseMigrationMigrationDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).DatabaseMigrationClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_database_migration_migration" {
			noResourceFound = false
			request := oci_database_migration.GetMigrationRequest{}

			tmp := rs.Primary.ID
			request.MigrationId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "database_migration")

			response, err := client.GetMigration(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_database_migration.LifecycleStatesDeleted): true,
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

func init() {
	if acctest.DependencyGraph == nil {
		acctest.InitDependencyGraph()
	}
	if !acctest.InSweeperExcludeList("DatabaseMigrationMigration") {
		resource.AddTestSweepers("DatabaseMigrationMigration", &resource.Sweeper{
			Name:         "DatabaseMigrationMigration",
			Dependencies: acctest.DependencyGraph["migration"],
			F:            sweepDatabaseMigrationMigrationResource,
		})
	}
}

func sweepDatabaseMigrationMigrationResource(compartment string) error {
	databaseMigrationClient := acctest.GetTestClients(&schema.ResourceData{}).DatabaseMigrationClient()
	migrationIds, err := getDatabaseMigrationMigrationIds(compartment)
	if err != nil {
		return err
	}
	for _, migrationId := range migrationIds {
		if ok := acctest.SweeperDefaultResourceId[migrationId]; !ok {
			deleteMigrationRequest := oci_database_migration.DeleteMigrationRequest{}

			deleteMigrationRequest.MigrationId = &migrationId

			deleteMigrationRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "database_migration")
			_, error := databaseMigrationClient.DeleteMigration(context.Background(), deleteMigrationRequest)
			if error != nil {
				fmt.Printf("Error deleting Migration %s %s, It is possible that the resource is already deleted. Please verify manually \n", migrationId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &migrationId, DatabaseMigrationmigrationsSweepWaitCondition, time.Duration(3*time.Minute),
				DatabaseMigrationmigrationsSweepResponseFetchOperation, "database_migration", true)
		}
	}
	return nil
}

func getDatabaseMigrationMigrationIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "MigrationId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	databaseMigrationClient := acctest.GetTestClients(&schema.ResourceData{}).DatabaseMigrationClient()

	listMigrationsRequest := oci_database_migration.ListMigrationsRequest{}
	listMigrationsRequest.CompartmentId = &compartmentId
	listMigrationsRequest.LifecycleState = oci_database_migration.ListMigrationsLifecycleStateActive
	listMigrationsResponse, err := databaseMigrationClient.ListMigrations(context.Background(), listMigrationsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting Migration list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, migration := range listMigrationsResponse.Items {
		id := *migration.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "MigrationId", id)
	}
	return resourceIds, nil
}

func DatabaseMigrationmigrationsSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	/*if migrationResponse, ok := response.Response.(oci_database_migration.GetMigrationResponse); ok {
		return migrationResponse.LifecycleState != oci_database_migration.LifecycleStatesDeleted
	}*/
	return false
}

func DatabaseMigrationmigrationsSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.DatabaseMigrationClient().GetMigration(context.Background(), oci_database_migration.GetMigrationRequest{
		MigrationId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
