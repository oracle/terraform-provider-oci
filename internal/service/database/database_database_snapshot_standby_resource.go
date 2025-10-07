// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package database

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"

	oci_database "github.com/oracle/oci-go-sdk/v65/database"
	oci_work_requests "github.com/oracle/oci-go-sdk/v65/workrequests"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DatabaseDatabaseSnapshotStandbyResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createDatabaseDatabaseSnapshotStandyby,
		Read:     readDatabaseDatabaseSnapshotStandyby,
		Delete:   deleteDatabaseDatabaseSnapshotStandyby,
		Schema: map[string]*schema.Schema{
			// Required
			"database_admin_password": {
				Type:      schema.TypeString,
				Required:  true,
				ForceNew:  true,
				Sensitive: true,
			},
			"database_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"standby_conversion_type": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
				ValidateFunc: validation.StringInSlice([]string{
					"SNAPSHOT",
					"PHYSICAL",
				}, true),
			},

			// Optional
			"snapshot_duration_in_days": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},

			// Computed
			"character_set": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"connection_strings": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"all_connection_strings": {
							Type:     schema.TypeMap,
							Computed: true,
							Elem:     schema.TypeString,
						},
						"cdb_default": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"cdb_ip_default": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"data_guard_group": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"members": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"apply_lag": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"apply_rate": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"database_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"db_system_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"is_active_data_guard_enabled": {
										Type:     schema.TypeBool,
										Computed: true,
									},
									"role": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"transport_lag": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"transport_lag_refresh": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"transport_type": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"protection_mode": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"database_management_config": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"management_status": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"management_type": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"database_software_image_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"db_backup_config": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"auto_backup_enabled": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"auto_backup_window": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"auto_full_backup_day": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"auto_full_backup_window": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"backup_deletion_policy": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"backup_destination_details": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"dbrs_policy_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"internet_proxy": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"type": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"vpc_password": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"vpc_user": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"recovery_window_in_days": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"run_immediate_full_backup": {
							Type:     schema.TypeBool,
							Computed: true,
						},
					},
				},
			},
			"db_home_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"db_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"db_system_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"db_unique_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"db_workload": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"defined_tags": {
				Type:     schema.TypeMap,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"freeform_tags": {
				Type:     schema.TypeMap,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"is_cdb": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"key_store_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"key_store_wallet_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"kms_key_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"kms_key_version_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"last_backup_duration_in_seconds": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"last_backup_timestamp": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"last_failed_backup_timestamp": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"lifecycle_details": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"ncharacter_set": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"pdb_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"sid_prefix": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"source_database_point_in_time_recovery_timestamp": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_created": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"vault_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"vm_cluster_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func createDatabaseDatabaseSnapshotStandyby(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseDatabaseSnapshotStandybyResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()
	sync.WorkRequestClient = m.(*client.OracleClients).WorkRequestClient

	return tfresource.CreateResource(d, sync)
}

func readDatabaseDatabaseSnapshotStandyby(d *schema.ResourceData, m interface{}) error {
	return nil
}

func deleteDatabaseDatabaseSnapshotStandyby(d *schema.ResourceData, m interface{}) error {
	return nil
}

type DatabaseDatabaseSnapshotStandybyResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_database.DatabaseClient
	Res                    *oci_database.Database
	DisableNotFoundRetries bool
	WorkRequestClient      *oci_work_requests.WorkRequestClient
}

func (s *DatabaseDatabaseSnapshotStandybyResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *DatabaseDatabaseSnapshotStandybyResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_database.DatabaseLifecycleStateProvisioning),
	}
}

func (s *DatabaseDatabaseSnapshotStandybyResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_database.DatabaseLifecycleStateAvailable),
	}
}

func (s *DatabaseDatabaseSnapshotStandybyResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_database.DatabaseLifecycleStateTerminating),
	}
}

func (s *DatabaseDatabaseSnapshotStandybyResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_database.DatabaseLifecycleStateTerminated),
	}
}

func (s *DatabaseDatabaseSnapshotStandybyResourceCrud) Create() error {
	request := oci_database.ConvertStandbyDatabaseTypeRequest{}

	if databaseAdminPassword, ok := s.D.GetOkExists("database_admin_password"); ok {
		tmp := databaseAdminPassword.(string)
		request.DatabaseAdminPassword = &tmp
	}

	if databaseId, ok := s.D.GetOkExists("database_id"); ok {
		tmp := databaseId.(string)
		request.DatabaseId = &tmp
	}

	if snapshotDurationInDays, ok := s.D.GetOkExists("snapshot_duration_in_days"); ok {
		tmp := snapshotDurationInDays.(int)
		request.SnapshotDurationInDays = &tmp
	}

	if standbyConversionType, ok := s.D.GetOkExists("standby_conversion_type"); ok {
		request.StandbyConversionType = oci_database.ConvertStandbyDatabaseTypeDetailsStandbyConversionTypeEnum(standbyConversionType.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database")

	response, err := s.Client.ConvertStandbyDatabaseType(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	s.Res = &response.Database

	if workId != nil {
		var identifier *string
		var err error
		identifier, err = tfresource.WaitForWorkRequestWithErrorHandling(s.WorkRequestClient, workId, "database", oci_work_requests.WorkRequestResourceActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate), s.DisableNotFoundRetries)
		if identifier != nil {
			s.D.SetId(*identifier)
		}
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *DatabaseDatabaseSnapshotStandybyResourceCrud) SetData() error {
	if s.Res.CharacterSet != nil {
		s.D.Set("character_set", *s.Res.CharacterSet)
	}

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.ConnectionStrings != nil {
		s.D.Set("connection_strings", []interface{}{DatabaseConnectionStringsToMap(s.Res.ConnectionStrings)})
	} else {
		s.D.Set("connection_strings", nil)
	}

	if s.Res.DataGuardGroup != nil {
		s.D.Set("data_guard_group", []interface{}{DataGuardGroupToMap(s.Res.DataGuardGroup)})
	} else {
		s.D.Set("data_guard_group", nil)
	}

	if s.Res.DatabaseManagementConfig != nil {
		s.D.Set("database_management_config", []interface{}{CloudDatabaseManagementConfigToMap(s.Res.DatabaseManagementConfig)})
	} else {
		s.D.Set("database_management_config", nil)
	}

	if s.Res.DatabaseSoftwareImageId != nil {
		s.D.Set("database_software_image_id", *s.Res.DatabaseSoftwareImageId)
	}

	if s.Res.DbBackupConfig != nil {
		s.D.Set("db_backup_config", []interface{}{DbBackupConfigToMap(s.Res.DbBackupConfig)})
	} else {
		s.D.Set("db_backup_config", nil)
	}

	if s.Res.DbHomeId != nil {
		s.D.Set("db_home_id", *s.Res.DbHomeId)
	}

	if s.Res.DbName != nil {
		s.D.Set("db_name", *s.Res.DbName)
	}

	if s.Res.DbSystemId != nil {
		s.D.Set("db_system_id", *s.Res.DbSystemId)
	}

	if s.Res.DbUniqueName != nil {
		s.D.Set("db_unique_name", *s.Res.DbUniqueName)
	}

	if s.Res.DbWorkload != nil {
		s.D.Set("db_workload", *s.Res.DbWorkload)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.IsCdb != nil {
		s.D.Set("is_cdb", *s.Res.IsCdb)
	}

	if s.Res.KeyStoreId != nil {
		s.D.Set("key_store_id", *s.Res.KeyStoreId)
	}

	if s.Res.KeyStoreWalletName != nil {
		s.D.Set("key_store_wallet_name", *s.Res.KeyStoreWalletName)
	}

	if s.Res.KmsKeyId != nil {
		s.D.Set("kms_key_id", *s.Res.KmsKeyId)
	}

	if s.Res.KmsKeyVersionId != nil {
		s.D.Set("kms_key_version_id", *s.Res.KmsKeyVersionId)
	}

	if s.Res.LastBackupDurationInSeconds != nil {
		s.D.Set("last_backup_duration_in_seconds", *s.Res.LastBackupDurationInSeconds)
	}

	if s.Res.LastBackupTimestamp != nil {
		s.D.Set("last_backup_timestamp", s.Res.LastBackupTimestamp.String())
	}

	if s.Res.LastFailedBackupTimestamp != nil {
		s.D.Set("last_failed_backup_timestamp", s.Res.LastFailedBackupTimestamp.String())
	}

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	if s.Res.NcharacterSet != nil {
		s.D.Set("ncharacter_set", *s.Res.NcharacterSet)
	}

	if s.Res.PdbName != nil {
		s.D.Set("pdb_name", *s.Res.PdbName)
	}

	if s.Res.SidPrefix != nil {
		s.D.Set("sid_prefix", *s.Res.SidPrefix)
	}

	if s.Res.SourceDatabasePointInTimeRecoveryTimestamp != nil {
		s.D.Set("source_database_point_in_time_recovery_timestamp", s.Res.SourceDatabasePointInTimeRecoveryTimestamp.String())
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.VaultId != nil {
		s.D.Set("vault_id", *s.Res.VaultId)
	}

	if s.Res.VmClusterId != nil {
		s.D.Set("vm_cluster_id", *s.Res.VmClusterId)
	}

	return nil
}

func (s *DatabaseDatabaseSnapshotStandybyResourceCrud) BackupDestinationDetailsToMap(obj oci_database.BackupDestinationDetails) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.DbrsPolicyId != nil {
		result["dbrs_policy_id"] = string(*obj.DbrsPolicyId)
	}

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	if obj.InternetProxy != nil {
		result["internet_proxy"] = string(*obj.InternetProxy)
	}

	result["type"] = string(obj.Type)

	if obj.VpcPassword != nil {
		result["vpc_password"] = string(*obj.VpcPassword)
	}

	if obj.VpcUser != nil {
		result["vpc_user"] = string(*obj.VpcUser)
	}

	return result
}

func (s *DatabaseDatabaseSnapshotStandybyResourceCrud) CloudDatabaseManagementConfigToMap(obj *oci_database.CloudDatabaseManagementConfig) map[string]interface{} {
	result := map[string]interface{}{}

	result["management_status"] = string(obj.ManagementStatus)

	result["management_type"] = string(obj.ManagementType)

	return result
}

func (s *DatabaseDatabaseSnapshotStandybyResourceCrud) DataGuardGroupToMap(obj *oci_database.DataGuardGroup) map[string]interface{} {
	result := map[string]interface{}{}

	members := []interface{}{}
	for _, item := range obj.Members {
		members = append(members, DataGuardGroupMemberToMap(item))
	}
	result["members"] = members

	result["protection_mode"] = string(obj.ProtectionMode)

	return result
}

func (s *DatabaseDatabaseSnapshotStandybyResourceCrud) DataGuardGroupMemberToMap(obj oci_database.DataGuardGroupMember) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.ApplyLag != nil {
		result["apply_lag"] = string(*obj.ApplyLag)
	}

	if obj.ApplyRate != nil {
		result["apply_rate"] = string(*obj.ApplyRate)
	}

	if obj.DatabaseId != nil {
		result["database_id"] = string(*obj.DatabaseId)
	}

	if obj.DbSystemId != nil {
		result["db_system_id"] = string(*obj.DbSystemId)
	}

	if obj.IsActiveDataGuardEnabled != nil {
		result["is_active_data_guard_enabled"] = bool(*obj.IsActiveDataGuardEnabled)
	}

	result["role"] = string(obj.Role)

	if obj.TransportLag != nil {
		result["transport_lag"] = string(*obj.TransportLag)
	}

	if obj.TransportLagRefresh != nil {
		result["transport_lag_refresh"] = string(*obj.TransportLagRefresh)
	}

	result["transport_type"] = string(obj.TransportType)

	return result
}

func (s *DatabaseDatabaseSnapshotStandybyResourceCrud) DatabaseConnectionStringsToMap(obj *oci_database.DatabaseConnectionStrings) map[string]interface{} {
	result := map[string]interface{}{}

	result["all_connection_strings"] = obj.AllConnectionStrings

	if obj.CdbDefault != nil {
		result["cdb_default"] = string(*obj.CdbDefault)
	}

	if obj.CdbIpDefault != nil {
		result["cdb_ip_default"] = string(*obj.CdbIpDefault)
	}

	return result
}

func (s *DatabaseDatabaseSnapshotStandybyResourceCrud) DbBackupConfigToMap(obj *oci_database.DbBackupConfig) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.AutoBackupEnabled != nil {
		result["auto_backup_enabled"] = bool(*obj.AutoBackupEnabled)
	}

	result["auto_backup_window"] = string(obj.AutoBackupWindow)

	result["auto_full_backup_day"] = string(obj.AutoFullBackupDay)

	result["auto_full_backup_window"] = string(obj.AutoFullBackupWindow)

	result["backup_deletion_policy"] = string(obj.BackupDeletionPolicy)

	backupDestinationDetails := []interface{}{}
	for _, item := range obj.BackupDestinationDetails {
		backupDestinationDetails = append(backupDestinationDetails, BackupDestinationDetailsToMap(item))
	}
	result["backup_destination_details"] = backupDestinationDetails

	if obj.RecoveryWindowInDays != nil {
		result["recovery_window_in_days"] = int(*obj.RecoveryWindowInDays)
	}

	if obj.RunImmediateFullBackup != nil {
		result["run_immediate_full_backup"] = bool(*obj.RunImmediateFullBackup)
	}

	return result
}
