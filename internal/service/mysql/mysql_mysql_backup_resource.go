// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package mysql

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"

	oci_mysql "github.com/oracle/oci-go-sdk/v65/mysql"
)

func MysqlMysqlBackupResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createMysqlMysqlBackup,
		Read:     readMysqlMysqlBackup,
		Update:   updateMysqlMysqlBackup,
		Delete:   deleteMysqlMysqlBackup,
		Schema: map[string]*schema.Schema{
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

			// Optional
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

			// Computed
			"backup_size_in_gbs": {
				Type:     schema.TypeInt,
				Computed: true,
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
						"port": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"port_x": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"region": {
							Type:     schema.TypeString,
							Computed: true,
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

	return tfresource.UpdateResource(d, sync)
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
	DestRegionClient       *oci_mysql.DbBackupsClient
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
		err := s.createMysqlBackupCopy()
		if err != nil {
			return err
		}
		s.D.SetId(*s.Res.Id)
		err = tfresource.WaitForResourceCondition(s, func() bool { return s.Res.LifecycleState == oci_mysql.BackupLifecycleStateActive }, s.D.Timeout(schema.TimeoutCreate))
		if err != nil {
			return err
		}
		// Update for some fields that can't be created by copy
		return s.Update()
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

	err := s.createDbBackupClientInRegion(currentRegion)
	if err != nil {
		return err
	}

	response, err := s.DestRegionClient.CopyBackup(context.Background(), copyMysqlBackupRequest)
	if err != nil {
		return err
	}

	s.Res = &response.Backup

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

	if s.isBackupCopy() {
		destinationRegion := utils.GetEnvSettingWithBlankDefault("destination_region")
		s.Client.SetRegion(destinationRegion)
	}

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
		s.D.Set("db_system_snapshot", []interface{}{DbSystemSnapshotToMap(s.Res.DbSystemSnapshot)})
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

	s.D.Set("state", s.Res.LifecycleState)

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

func BackupPolicyToMap(obj *oci_mysql.BackupPolicy) map[string]interface{} {
	result := map[string]interface{}{}

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

	if obj.WindowStartTime != nil {
		result["window_start_time"] = string(*obj.WindowStartTime)
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

func DbSystemSnapshotToMap(obj *oci_mysql.DbSystemSnapshot) map[string]interface{} {
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

	if obj.Port != nil {
		result["port"] = int(*obj.Port)
	}

	if obj.PortX != nil {
		result["port_x"] = int(*obj.PortX)
	}

	if obj.Region != nil {
		result["region"] = string(*obj.Region)
	}

	if obj.SecureConnections != nil {
		result["secure_connections"] = []interface{}{SecureConnectionDetailsToMap(obj.SecureConnections)}
	}

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

	if obj.WindowStartTime != nil {
		result["window_start_time"] = string(*obj.WindowStartTime)
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
