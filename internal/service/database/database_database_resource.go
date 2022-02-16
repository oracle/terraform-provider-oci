// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package database

import (
	"context"
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"
	"github.com/terraform-providers/terraform-provider-oci/internal/utils"

	oci_work_requests "github.com/oracle/oci-go-sdk/v58/workrequests"

	oci_common "github.com/oracle/oci-go-sdk/v58/common"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"

	oci_database "github.com/oracle/oci-go-sdk/v58/database"
)

type expectedRetryDurationFn func(response oci_common.OCIOperationResponse, disableNotFoundRetries bool, service string, optionals ...interface{}) time.Duration

func DatabaseDatabaseResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: tfresource.GetTimeoutDuration("12h"),
			Update: tfresource.GetTimeoutDuration("2h"),
			Delete: tfresource.GetTimeoutDuration("2h"),
		},
		Create: createDatabaseDatabase,
		Read:   readDatabaseDatabase,
		Update: updateDatabaseDatabase,
		Delete: deleteDatabaseDatabase,
		Schema: map[string]*schema.Schema{
			// Required
			"database": {
				Type:     schema.TypeList,
				Required: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"admin_password": {
							Type:      schema.TypeString,
							Required:  true,
							ForceNew:  true,
							Sensitive: true,
						},
						"db_name": {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},

						// Optional
						"backup_id": {
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
						},
						"backup_tde_password": {
							Type:      schema.TypeString,
							Optional:  true,
							ForceNew:  true,
							Sensitive: true,
						},
						"character_set": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},
						"database_software_image_id": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},
						"db_backup_config": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							MaxItems: 1,
							MinItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional
									"auto_backup_enabled": {
										Type:     schema.TypeBool,
										Optional: true,
										Computed: true,
									},
									"auto_backup_window": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"backup_destination_details": {
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										ForceNew: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional
												"id": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
													ForceNew: true,
												},
												"type": {
													Type:     schema.TypeString,
													Optional: true,
													ForceNew: true,
												},
												"vpc_user": {
													Type:     schema.TypeString,
													Optional: true,
												},
												// Computed
											},
										},
									},
									"recovery_window_in_days": {
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
									},

									// Computed
								},
							},
						},
						"db_unique_name": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},
						"db_workload": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},
						"defined_tags": {
							Type:             schema.TypeMap,
							Optional:         true,
							Computed:         true,
							DiffSuppressFunc: tfresource.DefinedTagsDiffSuppressFunction,
							Elem:             schema.TypeString,
						},
						"freeform_tags": {
							Type:     schema.TypeMap,
							Optional: true,
							Computed: true,
							Elem:     schema.TypeString,
						},
						"ncharacter_set": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},
						"pdb_name": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},
						"sid_prefix": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},
						"tde_wallet_password": {
							Type:      schema.TypeString,
							Optional:  true,
							Computed:  true,
							Sensitive: true,
						},

						// Computed
					},
				},
			},
			"db_home_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"source": {
				Type:             schema.TypeString,
				Required:         true,
				ForceNew:         true,
				DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
				ValidateFunc: validation.StringInSlice([]string{
					"DB_BACKUP",
					"NONE",
				}, true),
			},

			// Optional
			"db_version": {
				Type:             schema.TypeString,
				Optional:         true,
				Computed:         true,
				ForceNew:         true,
				DiffSuppressFunc: tfresource.DbVersionDiffSuppress,
			},
			"kms_key_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"kms_key_version_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"kms_key_rotation": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"kms_key_migration": {
				Type:     schema.TypeBool,
				Optional: true,
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
				MaxItems: 1,
				MinItems: 1,
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
			"database_management_config": {
				Type:     schema.TypeList,
				Computed: true,
				MaxItems: 1,
				MinItems: 1,
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
				MaxItems: 1,
				MinItems: 1,
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
						"backup_destination_details": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"type": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"vpc_user": {
										Type:     schema.TypeString,
										Optional: true,
									},
								},
							},
						},
						"recovery_window_in_days": {
							Type:     schema.TypeInt,
							Computed: true,
						},
					},
				},
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
			"last_backup_timestamp": {
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
			"vm_cluster_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func createDatabaseDatabase(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseDatabaseResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()
	sync.WorkRequestClient = m.(*client.OracleClients).WorkRequestClient

	return tfresource.CreateResource(d, sync)
}

func readDatabaseDatabase(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseDatabaseResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()
	sync.WorkRequestClient = m.(*client.OracleClients).WorkRequestClient

	return tfresource.ReadResource(sync)
}

func deleteDatabaseDatabase(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseDatabaseResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()
	sync.WorkRequestClient = m.(*client.OracleClients).WorkRequestClient
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type DatabaseDatabaseResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_database.DatabaseClient
	WorkRequestClient      *oci_work_requests.WorkRequestClient
	Res                    *oci_database.Database
	DisableNotFoundRetries bool
}

func (s *DatabaseDatabaseResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *DatabaseDatabaseResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_database.DatabaseLifecycleStateProvisioning),
	}
}

func (s *DatabaseDatabaseResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_database.DatabaseLifecycleStateAvailable),
	}
}

func (s *DatabaseDatabaseResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_database.DatabaseLifecycleStateTerminating),
	}
}

func (s *DatabaseDatabaseResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_database.DatabaseLifecycleStateTerminated),
	}
}

func (s *DatabaseDatabaseResourceCrud) UpdatedPending() []string {
	return []string{
		string(oci_database.DatabaseLifecycleStateProvisioning),
		string(oci_database.DatabaseLifecycleStateUpdating),
		string(oci_database.DatabaseLifecycleStateUpgrading),
	}
}

func (s *DatabaseDatabaseResourceCrud) UpdatedTarget() []string {
	return []string{
		string(oci_database.DatabaseLifecycleStateAvailable),
	}
}

func (s *DatabaseDatabaseResourceCrud) Create() error {
	request := oci_database.CreateDatabaseRequest{}
	err := s.populateTopLevelPolymorphicCreateDatabaseRequest(&request)
	if err != nil {
		return err
	}

	createDatabaseRetryDurationFn := getdatabaseRetryDurationFunction(s.D.Timeout(schema.TimeoutCreate))
	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database", createDatabaseRetryDurationFn)

	response, err := s.Client.CreateDatabase(context.Background(), request)
	if err != nil {
		return err
	}
	workId := response.OpcWorkRequestId
	if workId != nil {
		_, err = tfresource.WaitForWorkRequestWithErrorHandling(s.WorkRequestClient, workId, "database", oci_work_requests.WorkRequestResourceActionTypeCreated, s.D.Timeout(schema.TimeoutCreate), s.DisableNotFoundRetries)
		if err != nil {
			return err
		}
	}

	s.Res = &response.Database
	return nil
}

func (s *DatabaseDatabaseResourceCrud) Get() error {
	request := oci_database.GetDatabaseRequest{}

	tmp := s.D.Id()
	request.DatabaseId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database")

	response, err := s.Client.GetDatabase(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Database
	return nil
}

func (s *DatabaseDatabaseResourceCrud) Delete() error {
	request := oci_database.DeleteDatabaseRequest{}

	tmp := s.D.Id()
	request.DatabaseId = &tmp

	if performFinalBackup, ok := s.D.GetOkExists("perform_final_backup"); ok {
		tmp := performFinalBackup.(bool)
		request.PerformFinalBackup = &tmp
	}

	deleteDatabaseRetryDurationFn := getdatabaseRetryDurationFunction(s.D.Timeout(schema.TimeoutDelete))
	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database", deleteDatabaseRetryDurationFn)

	_, err := s.Client.DeleteDatabase(context.Background(), request)
	return err
}

func (s *DatabaseDatabaseResourceCrud) SetData() error {

	s.D.Set("database", []interface{}{s.DatabaseToMap(s.Res)})

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

	if s.Res.KmsKeyId != nil {
		s.D.Set("kms_key_id", *s.Res.KmsKeyId)
	}

	if s.Res.LastBackupTimestamp != nil {
		s.D.Set("last_backup_timestamp", s.Res.LastBackupTimestamp.String())
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

	if s.Res.VmClusterId != nil {
		s.D.Set("vm_cluster_id", *s.Res.VmClusterId)
	}

	return nil
}

func (s *DatabaseDatabaseResourceCrud) mapToBackupDestinationDetails(fieldKeyFormat string) (oci_database.BackupDestinationDetails, error) {
	result := oci_database.BackupDestinationDetails{}

	if id, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "id")); ok {
		tmp := id.(string)
		result.Id = &tmp
	}

	if type_, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "type")); ok {
		result.Type = oci_database.BackupDestinationDetailsTypeEnum(type_.(string))
	}

	if vpcPassword, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "vpc_password")); ok {
		tmp := vpcPassword.(string)
		result.VpcPassword = &tmp
	}

	if vpcUser, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "vpc_user")); ok {
		tmp := vpcUser.(string)
		result.VpcUser = &tmp
	}

	return result, nil
}

func BackupDestinationDetailsToMap(obj oci_database.BackupDestinationDetails) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
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

func CloudDatabaseManagementConfigToMap(obj *oci_database.CloudDatabaseManagementConfig) map[string]interface{} {
	result := map[string]interface{}{}

	result["management_status"] = string(obj.ManagementStatus)

	result["management_type"] = string(obj.ManagementType)

	return result
}

func (s *DatabaseDatabaseResourceCrud) mapToCreateDatabaseDetails(fieldKeyFormat string) (oci_database.CreateDatabaseDetails, error) {
	result := oci_database.CreateDatabaseDetails{}

	if adminPassword, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "admin_password")); ok {
		tmp := adminPassword.(string)
		result.AdminPassword = &tmp
	}

	if characterSet, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "character_set")); ok {
		tmp := characterSet.(string)
		result.CharacterSet = &tmp
	}

	if databaseSoftwareImageId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "database_software_image_id")); ok {
		tmp := databaseSoftwareImageId.(string)
		result.DatabaseSoftwareImageId = &tmp
	}

	if dbBackupConfig, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "db_backup_config")); ok {
		if tmpList := dbBackupConfig.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "db_backup_config"), 0)
			tmp, err := s.mapToDbBackupConfig(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert db_backup_config, encountered error: %v", err)
			}
			result.DbBackupConfig = &tmp
		}
	}

	if dbName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "db_name")); ok {
		tmp := dbName.(string)
		result.DbName = &tmp
	}

	if dbUniqueName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "db_unique_name")); ok {
		tmp := dbUniqueName.(string)
		result.DbUniqueName = &tmp
	}

	if dbWorkload, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "db_workload")); ok {
		result.DbWorkload = oci_database.CreateDatabaseDetailsDbWorkloadEnum(dbWorkload.(string))
	}

	if definedTags, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "defined_tags")); ok {
		tmp, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return result, fmt.Errorf("unable to convert defined_tags, encountered error: %v", err)
		}
		result.DefinedTags = tmp
	}

	if freeformTags, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "freeform_tags")); ok {
		result.FreeformTags = utils.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if ncharacterSet, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "ncharacter_set")); ok {
		tmp := ncharacterSet.(string)
		result.NcharacterSet = &tmp
	}

	if pdbName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "pdb_name")); ok {
		tmp := pdbName.(string)
		result.PdbName = &tmp
	}

	if sidPrefix, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "sid_prefix")); ok {
		tmp := sidPrefix.(string)
		result.SidPrefix = &tmp
	}

	if tdeWalletPassword, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "tde_wallet_password")); ok {
		tmp := tdeWalletPassword.(string)
		result.TdeWalletPassword = &tmp
	}

	return result, nil
}

func (s *DatabaseDatabaseResourceCrud) mapToCreateDatabaseFromBackupDetails(fieldKeyFormat string) (oci_database.CreateDatabaseFromBackupDetails, error) {
	result := oci_database.CreateDatabaseFromBackupDetails{}

	if adminPassword, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "admin_password")); ok {
		tmp := adminPassword.(string)
		result.AdminPassword = &tmp
	}

	if backupId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "backup_id")); ok {
		tmp := backupId.(string)
		result.BackupId = &tmp
	}

	if backupTDEPassword, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "backup_tde_password")); ok {
		tmp := backupTDEPassword.(string)
		result.BackupTDEPassword = &tmp
	}

	if dbName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "db_name")); ok {
		tmp := dbName.(string)
		result.DbName = &tmp
	}

	if dbUniqueName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "db_unique_name")); ok {
		tmp := dbUniqueName.(string)
		result.DbUniqueName = &tmp
	}

	if sidPrefix, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "sid_prefix")); ok {
		tmp := sidPrefix.(string)
		result.SidPrefix = &tmp
	}

	return result, nil
}

func DatabaseConnectionStringsToMap(obj *oci_database.DatabaseConnectionStrings) map[string]interface{} {
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

func (s *DatabaseDatabaseResourceCrud) mapToDbBackupConfig(fieldKeyFormat string) (oci_database.DbBackupConfig, error) {
	result := oci_database.DbBackupConfig{}

	if autoBackupEnabled, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "auto_backup_enabled")); ok {
		tmp := autoBackupEnabled.(bool)
		result.AutoBackupEnabled = &tmp
	}

	if autoBackupWindow, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "auto_backup_window")); ok {
		result.AutoBackupWindow = oci_database.DbBackupConfigAutoBackupWindowEnum(autoBackupWindow.(string))
	}

	if backupDestinationDetails, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "backup_destination_details")); ok {
		interfaces := backupDestinationDetails.([]interface{})
		tmp := make([]oci_database.BackupDestinationDetails, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "backup_destination_details"), stateDataIndex)
			converted, err := s.mapToBackupDestinationDetails(fieldKeyFormatNextLevel)
			if err != nil {
				return result, err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "backup_destination_details")) {
			result.BackupDestinationDetails = tmp
		}
	}

	if recoveryWindowInDays, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "recovery_window_in_days")); ok {
		tmp := recoveryWindowInDays.(int)
		result.RecoveryWindowInDays = &tmp
	}

	return result, nil
}

func (s *DatabaseDatabaseResourceCrud) populateTopLevelPolymorphicCreateDatabaseRequest(request *oci_database.CreateDatabaseRequest) error {
	//discriminator
	sourceRaw, ok := s.D.GetOkExists("source")
	var source string
	if ok {
		source = sourceRaw.(string)
	} else {
		source = "NONE" // default value
	}
	switch strings.ToLower(source) {
	case strings.ToLower("DB_BACKUP"):
		details := oci_database.CreateDatabaseFromBackup{}
		if database, ok := s.D.GetOkExists("database"); ok {
			if tmpList := database.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "database", 0)
				tmp, err := s.mapToCreateDatabaseFromBackupDetails(fieldKeyFormat)
				if err != nil {
					return err
				}
				details.Database = &tmp
			}
		}
		if dbHomeId, ok := s.D.GetOkExists("db_home_id"); ok {
			tmp := dbHomeId.(string)
			details.DbHomeId = &tmp
		}
		if dbVersion, ok := s.D.GetOkExists("db_version"); ok {
			tmp := dbVersion.(string)
			details.DbVersion = &tmp
		}
		if kmsKeyId, ok := s.D.GetOkExists("kms_key_id"); ok {
			tmp := kmsKeyId.(string)
			details.KmsKeyId = &tmp
		}
		if kmsKeyVersionId, ok := s.D.GetOkExists("kms_key_version_id"); ok {
			tmp := kmsKeyVersionId.(string)
			details.KmsKeyVersionId = &tmp
		}
		request.CreateNewDatabaseDetails = details
	case strings.ToLower("NONE"):
		details := oci_database.CreateNewDatabaseDetails{}
		if database, ok := s.D.GetOkExists("database"); ok {
			if tmpList := database.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "database", 0)
				tmp, err := s.mapToCreateDatabaseDetails(fieldKeyFormat)
				if err != nil {
					return err
				}
				details.Database = &tmp
			}
		}
		if dbHomeId, ok := s.D.GetOkExists("db_home_id"); ok {
			tmp := dbHomeId.(string)
			details.DbHomeId = &tmp
		}
		if dbVersion, ok := s.D.GetOkExists("db_version"); ok {
			tmp := dbVersion.(string)
			details.DbVersion = &tmp
		}
		if kmsKeyId, ok := s.D.GetOkExists("kms_key_id"); ok {
			tmp := kmsKeyId.(string)
			details.KmsKeyId = &tmp
		}
		if kmsKeyVersionId, ok := s.D.GetOkExists("kms_key_version_id"); ok {
			tmp := kmsKeyVersionId.(string)
			details.KmsKeyVersionId = &tmp
		}
		request.CreateNewDatabaseDetails = details
	default:
		return fmt.Errorf("unknown source '%v' was specified", source)
	}
	return nil
}

func updateDatabaseDatabase(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseDatabaseResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()
	sync.WorkRequestClient = m.(*client.OracleClients).WorkRequestClient

	return tfresource.UpdateResource(d, sync)
}

func (s *DatabaseDatabaseResourceCrud) Update() error {
	request := oci_database.UpdateDatabaseRequest{}

	tmp := s.D.Id()
	request.DatabaseId = &tmp

	error := s.kmsRotation(tmp)
	if error != nil {
		return error
	}

	err := s.kmsMigration(tmp)
	if err != nil {
		return err
	}

	if database, ok := s.D.GetOkExists("database"); ok {
		if tmpList := database.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "database", 0)
			tmp, err := s.mapToUpdateDatabaseDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.UpdateDatabaseDetails = tmp
		}
	}

	updateDatabaseRetryDurationFn := getdatabaseRetryDurationFunction(s.D.Timeout(schema.TimeoutUpdate))
	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database", updateDatabaseRetryDurationFn)

	response, err := s.Client.UpdateDatabase(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	if workId != nil {
		_, err = tfresource.WaitForWorkRequestWithErrorHandling(s.WorkRequestClient, workId, "database", oci_work_requests.WorkRequestResourceActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate), s.DisableNotFoundRetries)
		if err != nil {
			return err
		}
	}

	s.Res = &response.Database
	return nil
}

func (s *DatabaseDatabaseResourceCrud) mapToUpdateDbBackupConfig(fieldKeyFormat string) (oci_database.DbBackupConfig, error) {
	result := oci_database.DbBackupConfig{}

	if autoBackupEnabled, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "auto_backup_enabled")); ok && s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "auto_backup_enabled")) {
		tmp := autoBackupEnabled.(bool)
		result.AutoBackupEnabled = &tmp
	}

	if autoBackupWindow, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "auto_backup_window")); ok {
		result.AutoBackupWindow = oci_database.DbBackupConfigAutoBackupWindowEnum(autoBackupWindow.(string))
	}

	if recoveryWindowInDays, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "recovery_window_in_days")); ok && s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "recovery_window_in_days")) {
		tmp := recoveryWindowInDays.(int)
		result.RecoveryWindowInDays = &tmp
	}

	return result, nil
}

func (s *DatabaseDatabaseResourceCrud) mapToUpdateDatabaseDetails(fieldKeyFormat string) (oci_database.UpdateDatabaseDetails, error) {
	result := oci_database.UpdateDatabaseDetails{}

	if dbBackupConfig, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "db_backup_config")); ok {
		if tmpList := dbBackupConfig.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "db_backup_config"), 0)
			tmp, err := s.mapToUpdateDbBackupConfig(fieldKeyFormatNextLevel)
			if err != nil {
				return result, err
			}
			result.DbBackupConfig = &tmp
		}
	}

	if definedTags, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "defined_tags")); ok {
		tmp, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return result, fmt.Errorf("unable to convert defined_tags, encountered error: %v", err)
		}
		result.DefinedTags = tmp
	}

	if freeformTags, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "freeform_tags")); ok {
		result.FreeformTags = utils.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if adminPassword, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "admin_password")); ok && s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "admin_password")) {
		tmp := adminPassword.(string)
		result.NewAdminPassword = &tmp
	}

	if tdeWalletPassword, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "tde_wallet_password")); ok && s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "tde_wallet_password")) {
		tmp := tdeWalletPassword.(string)
		result.NewTdeWalletPassword = &tmp
		oldTdePassword, _ := s.D.GetChange(fmt.Sprintf(fieldKeyFormat, "tde_wallet_password"))
		tmp1 := oldTdePassword.(string)
		result.OldTdeWalletPassword = &tmp1
	}

	return result, nil
}

func (s *DatabaseDatabaseResourceCrud) DatabaseToMap(obj *oci_database.Database) map[string]interface{} {
	result := map[string]interface{}{}

	if adminPassword, ok := s.D.GetOkExists("database.0.admin_password"); ok && adminPassword != nil {
		result["admin_password"] = adminPassword.(string)
	}

	if tdeWalletPassword, ok := s.D.GetOkExists("database.0.tde_wallet_password"); ok && tdeWalletPassword != nil {
		result["tde_wallet_password"] = tdeWalletPassword.(string)
	}

	if backupId, ok := s.D.GetOkExists("database.0.backup_id"); ok && backupId != nil {
		result["backup_id"] = backupId.(string)
	}

	if backupTDEPassword, ok := s.D.GetOkExists("database.0.backup_tde_password"); ok && backupTDEPassword != nil {
		result["backup_tde_password"] = backupTDEPassword.(string)
	}

	if obj.CharacterSet != nil {
		result["character_set"] = string(*obj.CharacterSet)
	}

	if obj.DatabaseManagementConfig != nil {
		result["database_management_config"] = []interface{}{CloudDatabaseManagementConfigToMap(obj.DatabaseManagementConfig)}
	}

	if obj.DbBackupConfig != nil {
		result["db_backup_config"] = []interface{}{DbBackupConfigToMap(obj.DbBackupConfig)}
	}

	if obj.DbName != nil {
		result["db_name"] = string(*obj.DbName)
	}

	if obj.DbUniqueName != nil {
		result["db_unique_name"] = string(*obj.DbUniqueName)
	}

	if obj.DbWorkload != nil {
		result["db_workload"] = string(*obj.DbWorkload)
	}

	if obj.DefinedTags != nil {
		result["defined_tags"] = tfresource.DefinedTagsToMap(obj.DefinedTags)
	}

	result["freeform_tags"] = obj.FreeformTags

	if obj.NcharacterSet != nil {
		result["ncharacter_set"] = string(*obj.NcharacterSet)
	}

	if obj.PdbName != nil {
		result["pdb_name"] = string(*obj.PdbName)
	}

	if obj.SidPrefix != nil {
		result["sid_prefix"] = string(*obj.SidPrefix)
	}

	return result
}

func (s *DatabaseDatabaseResourceCrud) kmsRotation(databaseId string) error {
	if _, ok := s.D.GetOkExists("kms_key_rotation"); ok && s.D.HasChange("kms_key_rotation") {
		rotateKeyRequest := oci_database.RotateVaultKeyRequest{}
		rotateKeyRequest.DatabaseId = &databaseId
		rotateKeyRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database")
		response, err := s.Client.RotateVaultKey(context.Background(), rotateKeyRequest)
		if err != nil {
			return err
		}
		workId := response.OpcWorkRequestId
		if workId != nil {
			_, err = tfresource.WaitForWorkRequestWithErrorHandling(s.WorkRequestClient, workId, "database", oci_work_requests.WorkRequestResourceActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate), s.DisableNotFoundRetries)
			if err != nil {
				return err
			}
		}
		return nil
	}
	return nil
}

func (s *DatabaseDatabaseResourceCrud) kmsMigration(databaseId string) error {
	migrateOperation := false
	if _, ok := s.D.GetOkExists("kms_key_migration"); ok && s.D.HasChange("kms_key_migration") && s.D.Get("kms_key_migration").(bool) {
		migrationKeyRequest := oci_database.MigrateVaultKeyRequest{}
		migrationKeyRequest.DatabaseId = &databaseId
		migrationKeyRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database")
		if kmsKeyId, ok := s.D.GetOkExists("kms_key_id"); ok && s.D.HasChange("kms_key_id") {
			oldRaw, newRaw := s.D.GetChange("kms_key_id")
			if oldRaw == "" && newRaw != "" {
				temp := kmsKeyId.(string)
				migrationKeyRequest.KmsKeyId = &temp
			}
		}

		if kmsKeyVersionId, ok := s.D.GetOkExists("kms_key_version_id"); ok && s.D.HasChange("kms_key_version_id") {
			oldRaw, newRaw := s.D.GetChange("kms_key_version_id")
			if oldRaw == "" && newRaw != "" {
				temp := kmsKeyVersionId.(string)
				migrationKeyRequest.KmsKeyVersionId = &temp
			}
		}
		response, err := s.Client.MigrateVaultKey(context.Background(), migrationKeyRequest)
		if err != nil {
			return err
		}
		workId := response.OpcWorkRequestId
		if workId != nil {
			_, err = tfresource.WaitForWorkRequestWithErrorHandling(s.WorkRequestClient, workId, "database", oci_work_requests.WorkRequestResourceActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate), s.DisableNotFoundRetries)
			if err != nil {
			}
		}
		migrateOperation = true
	}

	if _, ok := s.D.GetOkExists("kms_key_id"); ok && s.D.HasChange("kms_key_id") && !migrateOperation {
		return errors.New("kms_key_id can not be updated, please use migration or rotation")
	}

	if _, ok := s.D.GetOkExists("kms_key_version_id"); ok && s.D.HasChange("kms_key_version_id") && !migrateOperation {
		return errors.New("kms_key_version_id can not be updated, please use migration or rotation")
	}
	return nil
}

func getdatabaseRetryDurationFunction(retryTimeout time.Duration) expectedRetryDurationFn {
	return func(response oci_common.OCIOperationResponse, disableNotFoundRetries bool, service string, optionals ...interface{}) time.Duration {
		defaultRetryTime := tfresource.GetDefaultExpectedRetryDuration(response, disableNotFoundRetries)
		if response.Response == nil || response.Response.HTTPResponse() == nil {
			return defaultRetryTime
		}
		e := response.Error
		switch statusCode := response.Response.HTTPResponse().StatusCode; statusCode {
		case 409:
			if isDisable409Retry, _ := strconv.ParseBool(utils.GetEnvSettingWithDefault("disable_409_retry", "false")); isDisable409Retry {
				log.Printf("[ERROR] Resource is in conflict state due to multiple update request: %v", e.Error())
				return 0
			}
			if e := response.Error; e != nil {
				if strings.Contains(e.Error(), "IncorrectState") {
					defaultRetryTime = retryTimeout
				} else if strings.Contains(e.Error(), "InvalidatedRetryToken") {
					defaultRetryTime = 0
				} else {
					defaultRetryTime = tfresource.LongRetryTime
				}
			}
		}
		return defaultRetryTime
	}
}
