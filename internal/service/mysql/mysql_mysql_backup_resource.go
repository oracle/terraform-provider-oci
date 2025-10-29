// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package mysql

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"

	oci_mysql "github.com/oracle/oci-go-sdk/v65/mysql"
)

func MysqlMysqlBackupResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: &tfresource.TwentyMinutes,
			Update: &tfresource.ThirtyMinutes,
			Delete: &tfresource.TwentyMinutes,
		},
		Create: createMysqlMysqlBackup,
		Read:   readMysqlMysqlBackup,
		Update: updateMysqlMysqlBackup,
		Delete: deleteMysqlMysqlBackup,
		Schema: map[string]*schema.Schema{
			// Optional
			"db_system_id": {
				Type:          schema.TypeString,
				Optional:      true,
				Computed:      true,
				ForceNew:      true,
				ConflictsWith: []string{"source_details"},
			},
			"source_details": {
				Type:          schema.TypeList,
				Optional:      true,
				ForceNew:      true,
				MaxItems:      1,
				MinItems:      1,
				ConflictsWith: []string{"db_system_id"},
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"region": {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},
						"backup_id": {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},
						"compartment_id": {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},
					},
				},
			},

			"backup_type": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"defined_tags": {
				Type:             schema.TypeMap,
				Optional:         true,
				Computed:         true,
				DiffSuppressFunc: tfresource.DefinedTagsDiffSuppressFunction,
				Elem:             schema.TypeString,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"freeform_tags": {
				Type:     schema.TypeMap,
				Optional: true,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"retention_in_days": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"soft_delete": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"validate_trigger": {
				Type:     schema.TypeInt,
				Optional: true,
			},

			// Computed
			"backup_size_in_gbs": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"backup_validation_details": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"backup_preparation_status": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"error_message": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"estimated_restore_duration": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"prepared_backup_details": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"prepared_backup_restore_reduction_in_minutes": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									"time_prepared": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"time_last_validated": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"validation_status": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"validate_backup_details": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"is_prepared_backup_required": {
							Type:     schema.TypeBool,
							Required: true,
						},
					},
				},
			},
			"creation_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"data_storage_size_in_gb": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"db_system_snapshot": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional
						"compartment_id": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},

						// Computed
						"admin_username": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"availability_domain": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"backup_policy": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"copy_policies": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional

												// Computed
												"backup_copy_retention_in_days": {
													Type:     schema.TypeInt,
													Computed: true,
												},
												"copy_to_region": {
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
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
									"is_enabled": {
										Type:     schema.TypeBool,
										Computed: true,
									},
									"pitr_policy": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional

												// Computed
												"is_enabled": {
													Type:     schema.TypeBool,
													Computed: true,
												},
											},
										},
									},
									"retention_in_days": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									"soft_delete": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"window_start_time": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"configuration_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"crash_recovery": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"data_storage": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"allocated_storage_size_in_gbs": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									"data_storage_size_in_gb": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									"data_storage_size_limit_in_gbs": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									"is_auto_expand_storage_enabled": {
										Type:     schema.TypeBool,
										Computed: true,
									},
									"max_storage_size_in_gbs": {
										Type:     schema.TypeInt,
										Computed: true,
									},
								},
							},
						},
						"data_storage_size_in_gb": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"database_management": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"defined_tags": {
							Type:     schema.TypeMap,
							Computed: true,
							Elem:     schema.TypeString,
						},
						"deletion_policy": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"automatic_backup_retention": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"final_backup": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"is_delete_protected": {
										Type:     schema.TypeBool,
										Computed: true,
									},
								},
							},
						},
						"description": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"display_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"encrypt_data": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"key_generation_type": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"key_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"endpoints": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"hostname": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"ip_address": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"modes": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"port": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									"port_x": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									"resource_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"resource_type": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"status": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"status_details": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"fault_domain": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"freeform_tags": {
							Type:     schema.TypeMap,
							Computed: true,
							Elem:     schema.TypeString,
						},
						"hostname_label": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"ip_address": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"is_highly_available": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"maintenance": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"maintenance_schedule_type": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"target_version": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"time_scheduled": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"version_preference": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"version_track_preference": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"window_start_time": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"mysql_version": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"nsg_ids": {
							Type:     schema.TypeSet,
							Computed: true,
							Set:      tfresource.LiteralTypeHashCodeForSets,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"port": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"port_x": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"read_endpoint": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"exclude_ips": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"is_enabled": {
										Type:     schema.TypeBool,
										Computed: true,
									},
									"read_endpoint_hostname_label": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"read_endpoint_ip_address": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"region": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"rest": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"configuration": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"port": {
										Type:     schema.TypeInt,
										Computed: true,
									},
								},
							},
						},
						"secure_connections": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"certificate_generation_type": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"certificate_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"security_attributes": {
							Type:     schema.TypeMap,
							Computed: true,
							Elem:     schema.TypeString,
						},
						"shape_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"subnet_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"encrypt_data": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"key_generation_type": {
							Type:     schema.TypeString,
							Required: true,
						},

						// Optional
						"key_id": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"db_system_snapshot_summary": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Computed
						"display_name": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"id": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"region": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"immediate_source_backup_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"lifecycle_details": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"mysql_version": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"original_source_backup_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"shape_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"system_tags": {
				Type:     schema.TypeMap,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"time_copy_created": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_created": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_updated": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func createMysqlMysqlBackup(d *schema.ResourceData, m interface{}) error {
	sync := &MysqlMysqlBackupResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DbBackupsClient()

	return tfresource.CreateResource(d, sync)
}

func readMysqlMysqlBackup(d *schema.ResourceData, m interface{}) error {
	sync := &MysqlMysqlBackupResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DbBackupsClient()

	return tfresource.ReadResource(sync)
}

func updateMysqlMysqlBackup(d *schema.ResourceData, m interface{}) error {
	sync := &MysqlMysqlBackupResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DbBackupsClient()

	if _, ok := sync.D.GetOkExists("validate_trigger"); ok && sync.D.HasChange("validate_trigger") {
		oldRaw, newRaw := sync.D.GetChange("validate_trigger")
		oldValue := oldRaw.(int)
		newValue := newRaw.(int)
		if oldValue < newValue {
			err := sync.ValidateBackup()

			if err != nil {
				return err
			}
		} else {
			sync.D.Set("validate_trigger", oldRaw)
			return fmt.Errorf("new value of trigger should be greater than the old value")
		}
	}

	if err := tfresource.UpdateResource(d, sync); err != nil {
		return err
	}

	return nil
}

func deleteMysqlMysqlBackup(d *schema.ResourceData, m interface{}) error {
	sync := &MysqlMysqlBackupResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DbBackupsClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type MysqlMysqlBackupResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_mysql.DbBackupsClient
	Res                    *oci_mysql.Backup
	DisableNotFoundRetries bool
}

func (s *MysqlMysqlBackupResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *MysqlMysqlBackupResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_mysql.BackupLifecycleStateCreating),
	}
}

func (s *MysqlMysqlBackupResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_mysql.BackupLifecycleStateActive),
	}
}

func (s *MysqlMysqlBackupResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_mysql.BackupLifecycleStateDeleting),
	}
}

func (s *MysqlMysqlBackupResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_mysql.BackupLifecycleStateDeleted),
	}
}

func (s *MysqlMysqlBackupResourceCrud) UpdatedPending() []string {
	return []string{
		string(oci_mysql.BackupLifecycleStateUpdating),
	}
}

func (s *MysqlMysqlBackupResourceCrud) UpdatedTarget() []string {
	return []string{
		string(oci_mysql.BackupLifecycleStateActive),
	}
}

func (s *MysqlMysqlBackupResourceCrud) Create() error {
	if s.isCopyCreate() {
		return s.createMysqlBackupCopy()
	}

	return s.createMysqlBackup()
}

func (s *MysqlMysqlBackupResourceCrud) isCopyCreate() bool {
	if sourceDetails, ok := s.D.GetOkExists("source_details"); ok {
		if tmpList := sourceDetails.([]interface{}); len(tmpList) > 0 {
			return true
		}
	}
	return false
}

func (s *MysqlMysqlBackupResourceCrud) createMysqlBackup() error {
	request := oci_mysql.CreateBackupRequest{}

	if backupType, ok := s.D.GetOkExists("backup_type"); ok {
		request.BackupType = oci_mysql.CreateBackupDetailsBackupTypeEnum(backupType.(string))
	}

	if dbSystemId, ok := s.D.GetOkExists("db_system_id"); ok {
		tmp := dbSystemId.(string)
		request.DbSystemId = &tmp
	}

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
		convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return err
		}
		request.DefinedTags = convertedDefinedTags
	}

	if description, ok := s.D.GetOkExists("description"); ok {
		tmp := description.(string)
		request.Description = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if retentionInDays, ok := s.D.GetOkExists("retention_in_days"); ok {
		tmp := retentionInDays.(int)
		request.RetentionInDays = &tmp
	}

	if softDelete, ok := s.D.GetOkExists("soft_delete"); ok {
		request.SoftDelete = oci_mysql.SoftDeleteEnum(softDelete.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "mysql")

	response, err := s.Client.CreateBackup(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Backup
	return nil
}

func (s *MysqlMysqlBackupResourceCrud) createMysqlBackupCopy() error {
	copyMysqlBackupRequest := oci_mysql.CopyBackupRequest{}

	configProvider := *s.Client.ConfigurationProvider()
	if configProvider == nil {
		return fmt.Errorf("cannot access ConfigurationProvider")
	}

	currentRegion, error := configProvider.Region()
	if error != nil {
		return fmt.Errorf("cannot access Region for the current ConfigurationProvider")
	}

	if backupCopyRetentionInDays, ok := s.D.GetOkExists("retention_in_days"); ok && backupCopyRetentionInDays != nil {
		tmp := backupCopyRetentionInDays.(int)
		copyMysqlBackupRequest.BackupCopyRetentionInDays = &tmp
	}

	if description, ok := s.D.GetOkExists("description"); ok && description != nil {
		tmp := description.(string)
		copyMysqlBackupRequest.Description = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		copyMysqlBackupRequest.DisplayName = &tmp
	}

	if sourceDetails, ok := s.D.GetOkExists("source_details"); ok && sourceDetails != nil {
		fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "source_details", 0)

		if compartmentId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "compartment_id")); ok && compartmentId != nil {
			tmp := compartmentId.(string)
			copyMysqlBackupRequest.CompartmentId = &tmp
		}

		if region, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "region")); ok {
			tmp := region.(string)
			copyMysqlBackupRequest.SourceRegion = &tmp
		}

		if sourceBackupId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "backup_id")); ok && sourceBackupId != nil {
			tmp := sourceBackupId.(string)
			copyMysqlBackupRequest.SourceBackupId = &tmp
		}
	}

	if encryptData, ok := s.D.GetOkExists("encrypt_data"); ok {
		if tmpList := encryptData.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "encrypt_data", 0)
			tmp, err := s.mapToEncryptDataDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			copyMysqlBackupRequest.EncryptData = &tmp
		}
	}

	err := s.createDbBackupClientInRegion(currentRegion)
	if err != nil {
		return err
	}

	response, err := s.Client.CopyBackup(context.Background(), copyMysqlBackupRequest)
	if err != nil {
		return err
	}

	s.Res = &response.Backup

	s.D.SetId(*s.Res.Id)
	err = tfresource.WaitForResourceCondition(s, func() bool { return s.Res.LifecycleState == oci_mysql.BackupLifecycleStateActive }, s.D.Timeout(schema.TimeoutCreate))
	if err != nil {
		return err
	}

	return nil
}

func (s *MysqlMysqlBackupResourceCrud) Get() error {
	request := oci_mysql.GetBackupRequest{}

	tmp := s.D.Id()
	request.BackupId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "mysql")

	response, err := s.Client.GetBackup(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Backup
	return nil
}

func (s *MysqlMysqlBackupResourceCrud) Update() error {
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(compartment)
			if err != nil {
				return err
			}
		}
	}
	request := oci_mysql.UpdateBackupRequest{}

	tmp := s.D.Id()
	request.BackupId = &tmp

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
		convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return err
		}
		request.DefinedTags = convertedDefinedTags
	}

	if description, ok := s.D.GetOkExists("description"); ok {
		tmp := description.(string)
		request.Description = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if retentionInDays, ok := s.D.GetOkExists("retention_in_days"); ok {
		tmp := retentionInDays.(int)
		request.RetentionInDays = &tmp
	}

	if softDelete, ok := s.D.GetOkExists("soft_delete"); ok {
		request.SoftDelete = oci_mysql.SoftDeleteEnum(softDelete.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "mysql")

	response, err := s.Client.UpdateBackup(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Backup
	return nil
}

func (s *MysqlMysqlBackupResourceCrud) Delete() error {
	request := oci_mysql.DeleteBackupRequest{}

	tmp := s.D.Id()
	request.BackupId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "mysql")

	_, err := s.Client.DeleteBackup(context.Background(), request)
	return err
}

func (s *MysqlMysqlBackupResourceCrud) isBackupCopy() bool {
	if _, ok := s.D.GetOk("immediate_source_backup_id"); ok {
		return true
	}
	return false
}

func (s *MysqlMysqlBackupResourceCrud) SetData() error {
	if s.Res.BackupSizeInGBs != nil {
		s.D.Set("backup_size_in_gbs", *s.Res.BackupSizeInGBs)
	}

	s.D.Set("backup_type", s.Res.BackupType)

	if s.Res.BackupValidationDetails != nil {
		s.D.Set("backup_validation_details", []interface{}{BackupValidationDetailsToMap(s.Res.BackupValidationDetails)})
	} else {
		s.D.Set("backup_validation_details", nil)
	}

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	s.D.Set("creation_type", s.Res.CreationType)

	if s.Res.DataStorageSizeInGBs != nil {
		s.D.Set("data_storage_size_in_gb", *s.Res.DataStorageSizeInGBs)
	}

	if s.Res.DbSystemId != nil {
		s.D.Set("db_system_id", *s.Res.DbSystemId)
	}

	if s.Res.DbSystemSnapshot != nil {
		s.D.Set("db_system_snapshot", []interface{}{DbSystemSnapshotToMap(s.Res.DbSystemSnapshot, false)})
	} else {
		s.D.Set("db_system_snapshot", nil)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.Description != nil {
		s.D.Set("description", *s.Res.Description)
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	if s.Res.EncryptData != nil {
		s.D.Set("encrypt_data", []interface{}{EncryptDataDetailsToMap(s.Res.EncryptData)})
	} else {
		s.D.Set("encrypt_data", nil)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.ImmediateSourceBackupId != nil {
		s.D.Set("immediate_source_backup_id", *s.Res.ImmediateSourceBackupId)
	}

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	if s.Res.MysqlVersion != nil {
		s.D.Set("mysql_version", *s.Res.MysqlVersion)
	}

	if s.Res.OriginalSourceBackupId != nil {
		s.D.Set("original_source_backup_id", *s.Res.OriginalSourceBackupId)
	}

	if s.Res.RetentionInDays != nil {
		s.D.Set("retention_in_days", *s.Res.RetentionInDays)
	}

	if s.Res.ShapeName != nil {
		s.D.Set("shape_name", *s.Res.ShapeName)
	}

	s.D.Set("soft_delete", s.Res.SoftDelete)

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.SystemTags != nil {
		s.D.Set("system_tags", tfresource.SystemTagsToMap(s.Res.SystemTags))
	}

	if s.Res.TimeCopyCreated != nil {
		s.D.Set("time_copy_created", s.Res.TimeCopyCreated.String())
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	return nil
}

func (s *MysqlMysqlBackupResourceCrud) ValidateBackup() error {
	request := oci_mysql.ValidateBackupRequest{}

	backupId := s.D.Id()
	request.BackupId = &backupId

	if validateBackupDetails, ok := s.D.GetOkExists("validate_backup_details"); ok && validateBackupDetails != nil {
		fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "validate_backup_details", 0)

		if isPreparedBackupRequired, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_prepared_backup_required")); ok && isPreparedBackupRequired != nil {
			tmp := isPreparedBackupRequired.(bool)
			request.IsPreparedBackupRequired = &tmp
		}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "mysql")

	response, err := s.Client.ValidateBackup(context.Background(), request)
	if err != nil {
		return err
	}

	if waitErr := tfresource.WaitForUpdatedState(s.D, s); waitErr != nil {
		return waitErr
	}

	val := s.D.Get("validate_trigger")
	s.D.Set("validate_trigger", val)

	s.Res = &response.Backup
	return nil
}

func BackupPolicyToMap(obj *oci_mysql.BackupPolicy) map[string]interface{} {
	result := map[string]interface{}{}

	copyPolicies := []interface{}{}
	for _, item := range obj.CopyPolicies {
		copyPolicies = append(copyPolicies, CopyPolicyToMap(item))
	}
	result["copy_policies"] = copyPolicies

	if obj.DefinedTags != nil {
		result["defined_tags"] = tfresource.DefinedTagsToMap(obj.DefinedTags)
	}

	result["freeform_tags"] = obj.FreeformTags

	if obj.IsEnabled != nil {
		result["is_enabled"] = bool(*obj.IsEnabled)
	}

	if obj.PitrPolicy != nil {
		result["pitr_policy"] = []interface{}{PitrPolicyToMap(obj.PitrPolicy)}
	}

	if obj.RetentionInDays != nil {
		result["retention_in_days"] = int(*obj.RetentionInDays)
	}

	result["soft_delete"] = string(obj.SoftDelete)

	if obj.WindowStartTime != nil {
		result["window_start_time"] = string(*obj.WindowStartTime)
	}

	return result
}

func BackupValidationDetailsToMap(obj *oci_mysql.BackupValidationDetails) map[string]interface{} {
	result := map[string]interface{}{}

	result["backup_preparation_status"] = string(obj.BackupPreparationStatus)

	if obj.ErrorMessage != nil {
		result["error_message"] = string(*obj.ErrorMessage)
	}

	if obj.EstimatedRestoreDuration != nil {
		result["estimated_restore_duration"] = string(*obj.EstimatedRestoreDuration)
	}

	if obj.PreparedBackupDetails != nil {
		result["prepared_backup_details"] = []interface{}{PreparedBackupDetailsToMap(obj.PreparedBackupDetails)}
	}

	if obj.TimeLastValidated != nil {
		result["time_last_validated"] = obj.TimeLastValidated.String()
	}

	result["validation_status"] = string(obj.ValidationStatus)

	return result
}

func CopyPolicyToMap(obj oci_mysql.CopyPolicy) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.BackupCopyRetentionInDays != nil {
		result["backup_copy_retention_in_days"] = int(*obj.BackupCopyRetentionInDays)
	}

	if obj.CopyToRegion != nil {
		result["copy_to_region"] = string(*obj.CopyToRegion)
	}

	return result
}

func DbSystemEndpointToMap(obj oci_mysql.DbSystemEndpoint) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Hostname != nil {
		result["hostname"] = string(*obj.Hostname)
	}

	if obj.IpAddress != nil {
		result["ip_address"] = string(*obj.IpAddress)
	}

	result["modes"] = obj.Modes

	if obj.Port != nil {
		result["port"] = int(*obj.Port)
	}

	if obj.PortX != nil {
		result["port_x"] = int(*obj.PortX)
	}

	if obj.ResourceId != nil {
		result["resource_id"] = string(*obj.ResourceId)
	}

	result["resource_type"] = string(obj.ResourceType)

	result["status"] = string(obj.Status)

	if obj.StatusDetails != nil {
		result["status_details"] = string(*obj.StatusDetails)
	}

	return result
}

func DbSystemSnapshotToMap(obj *oci_mysql.DbSystemSnapshot, datasource bool) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.AdminUsername != nil {
		result["admin_username"] = string(*obj.AdminUsername)
	}

	if obj.AvailabilityDomain != nil {
		result["availability_domain"] = string(*obj.AvailabilityDomain)
	}

	if obj.BackupPolicy != nil {
		result["backup_policy"] = []interface{}{BackupPolicyToMap(obj.BackupPolicy)}
	}

	if obj.CompartmentId != nil {
		result["compartment_id"] = string(*obj.CompartmentId)
	}

	if obj.ConfigurationId != nil {
		result["configuration_id"] = string(*obj.ConfigurationId)
	}

	result["crash_recovery"] = string(obj.CrashRecovery)

	if obj.DataStorage != nil {
		result["data_storage"] = []interface{}{DataStorageToMap(obj.DataStorage)}
	}

	if obj.DataStorageSizeInGBs != nil {
		result["data_storage_size_in_gb"] = int(*obj.DataStorageSizeInGBs)
	}

	result["database_management"] = string(obj.DatabaseManagement)

	if obj.DefinedTags != nil {
		result["defined_tags"] = tfresource.DefinedTagsToMap(obj.DefinedTags)
	}

	if obj.DeletionPolicy != nil {
		result["deletion_policy"] = []interface{}{DeletionPolicyDetailsToMap(obj.DeletionPolicy)}
	}

	if obj.Description != nil {
		result["description"] = string(*obj.Description)
	}

	if obj.DisplayName != nil {
		result["display_name"] = string(*obj.DisplayName)
	}

	if obj.EncryptData != nil {
		result["encrypt_data"] = []interface{}{EncryptDataDetailsToMap(obj.EncryptData)}
	}

	endpoints := []interface{}{}
	for _, item := range obj.Endpoints {
		endpoints = append(endpoints, DbSystemEndpointToMap(item))
	}
	result["endpoints"] = endpoints

	if obj.FaultDomain != nil {
		result["fault_domain"] = string(*obj.FaultDomain)
	}

	result["freeform_tags"] = obj.FreeformTags

	if obj.HostnameLabel != nil {
		result["hostname_label"] = string(*obj.HostnameLabel)
	}

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	if obj.IpAddress != nil {
		result["ip_address"] = string(*obj.IpAddress)
	}

	if obj.IsHighlyAvailable != nil {
		result["is_highly_available"] = bool(*obj.IsHighlyAvailable)
	}

	if obj.Maintenance != nil {
		result["maintenance"] = []interface{}{MaintenanceDetailsToMap(obj.Maintenance)}
	}

	if obj.MysqlVersion != nil {
		result["mysql_version"] = string(*obj.MysqlVersion)
	}

	nsgIds := []interface{}{}
	for _, item := range obj.NsgIds {
		nsgIds = append(nsgIds, item)
	}
	if datasource {
		result["nsg_ids"] = nsgIds
	} else {
		result["nsg_ids"] = schema.NewSet(tfresource.LiteralTypeHashCodeForSets, nsgIds)
	}

	if obj.Port != nil {
		result["port"] = int(*obj.Port)
	}

	if obj.PortX != nil {
		result["port_x"] = int(*obj.PortX)
	}

	if obj.ReadEndpoint != nil {
		result["read_endpoint"] = []interface{}{ReadEndpointDetailsToMap(obj.ReadEndpoint)}
	}

	if obj.Region != nil {
		result["region"] = string(*obj.Region)
	}

	if obj.Rest != nil {
		result["rest"] = []interface{}{RestDetailsToMap(obj.Rest)}
	}

	if obj.SecureConnections != nil {
		result["secure_connections"] = []interface{}{SecureConnectionDetailsToMap(obj.SecureConnections)}
	}

	result["security_attributes"] = obj.SecurityAttributes

	if obj.ShapeName != nil {
		result["shape_name"] = string(*obj.ShapeName)
	}

	if obj.SubnetId != nil {
		result["subnet_id"] = string(*obj.SubnetId)
	}

	return result
}

func DbSystemSnapshotSummaryToMap(obj *oci_mysql.DbSystemSnapshotSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.DisplayName != nil {
		result["display_name"] = string(*obj.DisplayName)
	}

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	if obj.Region != nil {
		result["region"] = string(*obj.Region)
	}

	return result
}

func MaintenanceDetailsToMap(obj *oci_mysql.MaintenanceDetails) map[string]interface{} {
	result := map[string]interface{}{}

	result["maintenance_schedule_type"] = string(obj.MaintenanceScheduleType)

	if obj.TargetVersion != nil {
		result["target_version"] = string(*obj.TargetVersion)
	}

	if obj.TimeScheduled != nil {
		result["time_scheduled"] = obj.TimeScheduled.String()
	}

	result["version_preference"] = string(obj.VersionPreference)

	result["version_track_preference"] = string(obj.VersionTrackPreference)

	if obj.WindowStartTime != nil {
		result["window_start_time"] = string(*obj.WindowStartTime)
	}

	return result
}

func (s *MysqlMysqlBackupResourceCrud) mapToEncryptDataDetails(fieldKeyFormat string) (oci_mysql.EncryptDataDetails, error) {
	result := oci_mysql.EncryptDataDetails{}

	if keyGenerationType, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "key_generation_type")); ok {
		result.KeyGenerationType = oci_mysql.KeyGenerationTypeEnum(keyGenerationType.(string))
	}

	if keyId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "key_id")); ok {
		tmp := keyId.(string)
		result.KeyId = &tmp
	}

	return result, nil
}

func PreparedBackupDetailsToMap(obj *oci_mysql.PreparedBackupDetails) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.PreparedBackupRestoreReductionInMinutes != nil {
		result["prepared_backup_restore_reduction_in_minutes"] = int(*obj.PreparedBackupRestoreReductionInMinutes)
	}

	if obj.TimePrepared != nil {
		result["time_prepared"] = obj.TimePrepared.String()
	}

	return result
}

func (s *MysqlMysqlBackupResourceCrud) updateCompartment(compartment interface{}) error {
	changeCompartmentRequest := oci_mysql.ChangeBackupCompartmentRequest{}

	tmp := s.D.Id()
	changeCompartmentRequest.BackupId = &tmp

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "mysql")

	_, err := s.Client.ChangeBackupCompartment(context.Background(), changeCompartmentRequest)
	if err != nil {
		return err
	}

	if waitErr := tfresource.WaitForUpdatedState(s.D, s); waitErr != nil {
		return waitErr
	}

	return nil
}
