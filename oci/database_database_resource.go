// Copyright (c) 2017, 2019, Oracle and/or its affiliates. All rights reserved.

package oci

import (
	"context"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/helper/validation"

	oci_database "github.com/oracle/oci-go-sdk/database"
)

func DatabaseDatabaseResource() *schema.Resource {
	return &schema.Resource{
		Timeouts: DefaultTimeout,
		Create:   createDatabaseDatabase,
		Read:     readDatabaseDatabase,
		Update:   updateDatabaseDatabase,
		Delete:   deleteDatabaseDatabase,
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
						"character_set": {
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
												"type": {
													Type:     schema.TypeString,
													Required: true,
													ForceNew: true,
												},

												// Optional
												"id": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
													ForceNew: true,
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
							DiffSuppressFunc: definedTagsDiffSuppressFunction,
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
				DiffSuppressFunc: EqualIgnoreCaseSuppressDiff,
				ValidateFunc: validation.StringInSlice([]string{
					"NONE",
				}, true),
			},

			// Optional
			"db_version": {
				Type:             schema.TypeString,
				Optional:         true,
				Computed:         true,
				ForceNew:         true,
				DiffSuppressFunc: dbVersionDiffSuppress,
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
	sync.Client = m.(*OracleClients).databaseClient

	return CreateResource(d, sync)
}

func readDatabaseDatabase(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseDatabaseResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).databaseClient

	return ReadResource(sync)
}

func updateDatabaseDatabase(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseDatabaseResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).databaseClient

	return UpdateResource(d, sync)
}

func deleteDatabaseDatabase(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseDatabaseResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).databaseClient
	sync.DisableNotFoundRetries = true

	return DeleteResource(d, sync)
}

type DatabaseDatabaseResourceCrud struct {
	BaseCrud
	Client                 *oci_database.DatabaseClient
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

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "database")

	response, err := s.Client.CreateDatabase(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Database
	return nil
}

func (s *DatabaseDatabaseResourceCrud) Get() error {
	request := oci_database.GetDatabaseRequest{}

	tmp := s.D.Id()
	request.DatabaseId = &tmp

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "database")

	response, err := s.Client.GetDatabase(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Database
	return nil
}

func (s *DatabaseDatabaseResourceCrud) Update() error {
	request := oci_database.UpdateDatabaseRequest{}

	tmp := s.D.Id()
	request.DatabaseId = &tmp

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

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "database")

	response, err := s.Client.UpdateDatabase(context.Background(), request)
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

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "database")

	_, err := s.Client.DeleteDatabase(context.Background(), request)
	return err
}

func (s *DatabaseDatabaseResourceCrud) SetData() error {
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
		s.D.Set("defined_tags", definedTagsToMap(s.Res.DefinedTags))
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	if s.Res.NcharacterSet != nil {
		s.D.Set("ncharacter_set", *s.Res.NcharacterSet)
	}

	if s.Res.PdbName != nil {
		s.D.Set("pdb_name", *s.Res.PdbName)
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
		tmp, err := mapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return result, fmt.Errorf("unable to convert defined_tags, encountered error: %v", err)
		}
		result.DefinedTags = tmp
	}

	if freeformTags, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "freeform_tags")); ok {
		result.FreeformTags = objectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if ncharacterSet, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "ncharacter_set")); ok {
		tmp := ncharacterSet.(string)
		result.NcharacterSet = &tmp
	}

	if pdbName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "pdb_name")); ok {
		tmp := pdbName.(string)
		result.PdbName = &tmp
	}

	return result, nil
}

// @CODEGEN: Method CreateDatabaseDetailsToMap is available in database_db_system_resource.go

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
		result.BackupDestinationDetails = tmp
	}

	if recoveryWindowInDays, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "recovery_window_in_days")); ok {
		tmp := recoveryWindowInDays.(int)
		result.RecoveryWindowInDays = &tmp
	}

	return result, nil
}

func (s *DatabaseDatabaseResourceCrud) mapToUpdateDbBackupConfig(fieldKeyFormat string) (oci_database.DbBackupConfig, error) {
	result := oci_database.DbBackupConfig{}

	// Service does not allow to update auto_backup_enabled and recovery_window_in_days at the same time so we must have the HasChanged check
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

// @CODEGEN: Method DbBackupConfigToMap is available in database_db_system_resource.go

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
		tmp, err := mapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return result, fmt.Errorf("unable to convert defined_tags, encountered error: %v", err)
		}
		result.DefinedTags = tmp
	}

	if freeformTags, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "freeform_tags")); ok {
		result.FreeformTags = objectMapToStringMap(freeformTags.(map[string]interface{}))
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
		request.CreateNewDatabaseDetails = details
	default:
		return fmt.Errorf("unknown source '%v' was specified", source)
	}
	return nil
}
