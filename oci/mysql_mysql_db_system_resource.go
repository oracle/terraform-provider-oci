// Copyright (c) 2017, 2020, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"

	oci_mysql "github.com/oracle/oci-go-sdk/v25/mysql"
)

func init() {
	RegisterResource("oci_mysql_mysql_db_system", MysqlMysqlDbSystemResource())
}

func MysqlMysqlDbSystemResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: getTimeoutDuration("1h"),
			Update: getTimeoutDuration("1h"),
			Delete: getTimeoutDuration("1h"),
		},
		Create: createMysqlMysqlDbSystem,
		Read:   readMysqlMysqlDbSystem,
		Update: updateMysqlMysqlDbSystem,
		Delete: deleteMysqlMysqlDbSystem,
		Schema: map[string]*schema.Schema{
			// Required
			"admin_password": {
				Type:      schema.TypeString,
				Required:  true,
				ForceNew:  true,
				Sensitive: true,
			},
			"admin_username": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"availability_domain": {
				Type:             schema.TypeString,
				Required:         true,
				ForceNew:         true,
				DiffSuppressFunc: EqualIgnoreCaseSuppressDiff,
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"configuration_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"shape_name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"subnet_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			// Optional
			"backup_policy": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional
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
						"is_enabled": {
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
						},
						"retention_in_days": {
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"window_start_time": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},

						// Computed
					},
				},
			},
			"data_storage_size_in_gb": {
				Type:     schema.TypeInt,
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
			"fault_domain": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"freeform_tags": {
				Type:     schema.TypeMap,
				Optional: true,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"hostname_label": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"ip_address": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"maintenance": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"window_start_time": {
							Type:     schema.TypeString,
							Required: true,
						},

						// Optional

						// Computed
					},
				},
			},
			"mysql_version": {
				Type:             schema.TypeString,
				Optional:         true,
				Computed:         true,
				ForceNew:         true,
				DiffSuppressFunc: mySqlVersionDiffSuppress,
			},
			"port": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"port_x": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"source": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				ForceNew: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"source_type": {
							Type:             schema.TypeString,
							Required:         true,
							ForceNew:         true,
							DiffSuppressFunc: EqualIgnoreCaseSuppressDiff,
							ValidateFunc: validation.StringInSlice([]string{
								"BACKUP",
							}, true),
						},

						// Optional
						"backup_id": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},

						// Computed
					},
				},
			},
			"state": {
				Type:             schema.TypeString,
				Computed:         true,
				Optional:         true,
				DiffSuppressFunc: EqualIgnoreCaseSuppressDiff,
				ValidateFunc: validation.StringInSlice([]string{
					string(oci_mysql.DbSystemLifecycleStateInactive),
					string(oci_mysql.DbSystemLifecycleStateActive),
				}, true),
			},
			"shutdown_type": {
				Type:             schema.TypeString,
				Optional:         true,
				DiffSuppressFunc: EqualIgnoreCaseSuppressDiff,
				ValidateFunc: validation.StringInSlice([]string{
					string(oci_mysql.InnoDbShutdownModeFast),
					string(oci_mysql.InnoDbShutdownModeImmediate),
					string(oci_mysql.InnoDbShutdownModeSlow),
				}, true),
			},

			// Computed
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
			"lifecycle_details": {
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

func createMysqlMysqlDbSystem(d *schema.ResourceData, m interface{}) error {
	sync := &MysqlMysqlDbSystemResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).dbSystemClient()

	var powerOff = false
	if powerState, ok := sync.D.GetOkExists("state"); ok {
		wantedPowerState := oci_mysql.DbSystemLifecycleStateEnum(strings.ToUpper(powerState.(string)))
		if wantedPowerState == oci_mysql.DbSystemLifecycleStateInactive {
			powerOff = true
		}
	}

	if e := CreateResource(d, sync); e != nil {
		return e
	}

	if powerOff {
		if err := sync.StopMysqlDbInstance(); err != nil {
			return err
		}
		sync.D.Set("state", oci_mysql.DbSystemLifecycleStateInactive)
	}

	return nil
}

func readMysqlMysqlDbSystem(d *schema.ResourceData, m interface{}) error {
	sync := &MysqlMysqlDbSystemResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).dbSystemClient()

	return ReadResource(sync)
}

func updateMysqlMysqlDbSystem(d *schema.ResourceData, m interface{}) error {
	sync := &MysqlMysqlDbSystemResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).dbSystemClient()

	// switch to power on
	powerOn, powerOff := false, false

	if sync.D.HasChange("state") {
		wantedState := strings.ToUpper(sync.D.Get("state").(string))
		if oci_mysql.DbSystemLifecycleStateActive == oci_mysql.DbSystemLifecycleStateEnum(wantedState) {
			powerOn = true
		} else if oci_mysql.DbSystemLifecycleStateInactive == oci_mysql.DbSystemLifecycleStateEnum(wantedState) {
			powerOff = true
		}
	}

	if powerOn {
		if err := sync.StartMysqlDbInstance(); err != nil {
			return err
		}
		sync.D.Set("state", oci_mysql.DbSystemLifecycleStateActive)
	}

	if err := UpdateResource(d, sync); err != nil {
		return err
	}

	// switch to power off
	if powerOff {
		if err := sync.StopMysqlDbInstance(); err != nil {
			return err
		}
		sync.D.Set("state", oci_mysql.DbSystemLifecycleStateInactive)
	}

	return nil
}

func deleteMysqlMysqlDbSystem(d *schema.ResourceData, m interface{}) error {
	sync := &MysqlMysqlDbSystemResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).dbSystemClient()
	sync.DisableNotFoundRetries = true

	return DeleteResource(d, sync)
}

type MysqlMysqlDbSystemResourceCrud struct {
	BaseCrud
	Client                 *oci_mysql.DbSystemClient
	Res                    *oci_mysql.DbSystem
	DisableNotFoundRetries bool
}

func (s *MysqlMysqlDbSystemResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *MysqlMysqlDbSystemResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_mysql.DbSystemLifecycleStateCreating),
	}
}

func (s *MysqlMysqlDbSystemResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_mysql.DbSystemLifecycleStateActive),
	}
}

func (s *MysqlMysqlDbSystemResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_mysql.DbSystemLifecycleStateDeleting),
	}
}

func (s *MysqlMysqlDbSystemResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_mysql.DbSystemLifecycleStateDeleted),
	}
}

func (s *MysqlMysqlDbSystemResourceCrud) UpdatedPending() []string {
	return []string{
		string(oci_mysql.DbSystemLifecycleStateUpdating),
	}
}

func (s *MysqlMysqlDbSystemResourceCrud) UpdatedTarget() []string {
	return []string{
		string(oci_mysql.DbSystemLifecycleStateActive),
	}
}

func (s *MysqlMysqlDbSystemResourceCrud) Create() error {
	request := oci_mysql.CreateDbSystemRequest{}

	if adminPassword, ok := s.D.GetOkExists("admin_password"); ok {
		tmp := adminPassword.(string)
		request.AdminPassword = &tmp
	}

	if adminUsername, ok := s.D.GetOkExists("admin_username"); ok {
		tmp := adminUsername.(string)
		request.AdminUsername = &tmp
	}

	if availabilityDomain, ok := s.D.GetOkExists("availability_domain"); ok {
		tmp := availabilityDomain.(string)
		request.AvailabilityDomain = &tmp
	}

	if backupPolicy, ok := s.D.GetOkExists("backup_policy"); ok {
		if tmpList := backupPolicy.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "backup_policy", 0)
			tmp, err := s.mapToCreateBackupPolicyDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.BackupPolicy = &tmp
		}
	}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if configurationId, ok := s.D.GetOkExists("configuration_id"); ok {
		tmp := configurationId.(string)
		request.ConfigurationId = &tmp
	}

	if dataStorageSizeInGB, ok := s.D.GetOkExists("data_storage_size_in_gb"); ok {
		tmp := dataStorageSizeInGB.(int)
		request.DataStorageSizeInGBs = &tmp
	}

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
		convertedDefinedTags, err := mapToDefinedTags(definedTags.(map[string]interface{}))
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

	if faultDomain, ok := s.D.GetOkExists("fault_domain"); ok {
		tmp := faultDomain.(string)
		request.FaultDomain = &tmp
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = objectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if hostnameLabel, ok := s.D.GetOkExists("hostname_label"); ok {
		tmp := hostnameLabel.(string)
		request.HostnameLabel = &tmp
	}

	if ipAddress, ok := s.D.GetOkExists("ip_address"); ok {
		tmp := ipAddress.(string)
		request.IpAddress = &tmp
	}

	if maintenance, ok := s.D.GetOkExists("maintenance"); ok {
		if tmpList := maintenance.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "maintenance", 0)
			tmp, err := s.mapToCreateMaintenanceDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.Maintenance = &tmp
		}
	}

	if mysqlVersion, ok := s.D.GetOkExists("mysql_version"); ok {
		tmp := mysqlVersion.(string)
		request.MysqlVersion = &tmp
	}

	if port, ok := s.D.GetOkExists("port"); ok {
		tmp := port.(int)
		request.Port = &tmp
	}

	if portX, ok := s.D.GetOkExists("port_x"); ok {
		tmp := portX.(int)
		request.PortX = &tmp
	}

	if shapeName, ok := s.D.GetOkExists("shape_name"); ok {
		tmp := shapeName.(string)
		request.ShapeName = &tmp
	}

	if source, ok := s.D.GetOkExists("source"); ok {
		if tmpList := source.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "source", 0)
			tmp, err := s.mapToCreateDbSystemSourceDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.Source = tmp
		}
	}

	if subnetId, ok := s.D.GetOkExists("subnet_id"); ok {
		tmp := subnetId.(string)
		request.SubnetId = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "mysql")

	response, err := s.Client.CreateDbSystem(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.DbSystem
	return nil
}

func (s *MysqlMysqlDbSystemResourceCrud) Get() error {
	request := oci_mysql.GetDbSystemRequest{}

	tmp := s.D.Id()
	request.DbSystemId = &tmp

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "mysql")

	response, err := s.Client.GetDbSystem(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.DbSystem
	return nil
}

func (s *MysqlMysqlDbSystemResourceCrud) Update() error {
	request := oci_mysql.UpdateDbSystemRequest{}

	if backupPolicy, ok := s.D.GetOkExists("backup_policy"); ok && s.D.HasChange("backup_policy") {
		if tmpList := backupPolicy.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "backup_policy", 0)
			tmp, err := s.mapToUpdateBackupPolicyDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.BackupPolicy = &tmp
		}
	}

	tmp := s.D.Id()
	request.DbSystemId = &tmp

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok && s.D.HasChange("defined_tags") {
		convertedDefinedTags, err := mapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return err
		}
		request.DefinedTags = convertedDefinedTags
	}

	if description, ok := s.D.GetOkExists("description"); ok && s.D.HasChange("description") {
		tmp := description.(string)
		request.Description = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok && s.D.HasChange("display_name") {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok && s.D.HasChange("freeform_tags") {
		request.FreeformTags = objectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if maintenance, ok := s.D.GetOkExists("maintenance"); ok && s.D.HasChange("maintenance") {
		if tmpList := maintenance.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "maintenance", 0)
			tmp, err := s.mapToUpdateMaintenanceDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.Maintenance = &tmp
		}
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "mysql")

	_, err := s.Client.UpdateDbSystem(context.Background(), request)
	if err != nil {
		return err
	}

	return s.Get()
}

func (s *MysqlMysqlDbSystemResourceCrud) Delete() error {
	request := oci_mysql.DeleteDbSystemRequest{}

	tmp := s.D.Id()
	request.DbSystemId = &tmp

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "mysql")

	_, err := s.Client.DeleteDbSystem(context.Background(), request)
	return err
}

func (s *MysqlMysqlDbSystemResourceCrud) SetData() error {
	if s.Res.AvailabilityDomain != nil {
		s.D.Set("availability_domain", *s.Res.AvailabilityDomain)
	}

	if s.Res.BackupPolicy != nil {
		s.D.Set("backup_policy", []interface{}{BackupPolicyToMap(s.Res.BackupPolicy)})
	} else {
		s.D.Set("backup_policy", nil)
	}

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.ConfigurationId != nil {
		s.D.Set("configuration_id", *s.Res.ConfigurationId)
	}

	if s.Res.DataStorageSizeInGBs != nil {
		s.D.Set("data_storage_size_in_gb", *s.Res.DataStorageSizeInGBs)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", definedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.Description != nil {
		s.D.Set("description", *s.Res.Description)
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	endpoints := []interface{}{}
	for _, item := range s.Res.Endpoints {
		endpoints = append(endpoints, DbSystemEndpointToMap(item))
	}
	s.D.Set("endpoints", endpoints)

	if s.Res.FaultDomain != nil {
		s.D.Set("fault_domain", *s.Res.FaultDomain)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.HostnameLabel != nil {
		s.D.Set("hostname_label", *s.Res.HostnameLabel)
	}

	if s.Res.IpAddress != nil {
		s.D.Set("ip_address", *s.Res.IpAddress)
	}

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	if s.Res.Maintenance != nil {
		s.D.Set("maintenance", []interface{}{MaintenanceDetailsToMap(s.Res.Maintenance)})
	} else {
		s.D.Set("maintenance", nil)
	}

	if s.Res.MysqlVersion != nil {
		s.D.Set("mysql_version", *s.Res.MysqlVersion)
	}

	if s.Res.Port != nil {
		s.D.Set("port", *s.Res.Port)
	}

	if s.Res.PortX != nil {
		s.D.Set("port_x", *s.Res.PortX)
	}

	if s.Res.ShapeName != nil {
		s.D.Set("shape_name", *s.Res.ShapeName)
	}

	if s.Res.Source != nil {
		sourceArray := []interface{}{}
		if sourceMap := DbSystemSourceToMap(&s.Res.Source); sourceMap != nil {
			sourceArray = append(sourceArray, sourceMap)
		}
		s.D.Set("source", sourceArray)
	} else {
		result := map[string]interface{}{}
		result["source_type"] = "NONE"
		sourceArray := []interface{}{}
		sourceArray = append(sourceArray, result)
		s.D.Set("source", sourceArray)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.SubnetId != nil {
		s.D.Set("subnet_id", *s.Res.SubnetId)
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	return nil
}

func (s *MysqlMysqlDbSystemResourceCrud) mapToCreateBackupPolicyDetails(fieldKeyFormat string) (oci_mysql.CreateBackupPolicyDetails, error) {
	result := oci_mysql.CreateBackupPolicyDetails{}

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

	if isEnabled, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_enabled")); ok {
		tmp := isEnabled.(bool)
		result.IsEnabled = &tmp
	}

	if retentionInDays, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "retention_in_days")); ok {
		tmp := retentionInDays.(int)
		result.RetentionInDays = &tmp
	}

	if windowStartTime, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "window_start_time")); ok {
		tmp := windowStartTime.(string)
		result.WindowStartTime = &tmp
	}

	return result, nil
}

func BackupPolicyToMap(obj *oci_mysql.BackupPolicy) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.DefinedTags != nil {
		result["defined_tags"] = definedTagsToMap(obj.DefinedTags)
	}

	result["freeform_tags"] = obj.FreeformTags

	if obj.IsEnabled != nil {
		result["is_enabled"] = bool(*obj.IsEnabled)
	}

	if obj.RetentionInDays != nil {
		result["retention_in_days"] = int(*obj.RetentionInDays)
	}

	if obj.WindowStartTime != nil {
		result["window_start_time"] = string(*obj.WindowStartTime)
	}

	return result
}

func (s *MysqlMysqlDbSystemResourceCrud) mapToCreateDbSystemSourceDetails(fieldKeyFormat string) (oci_mysql.CreateDbSystemSourceDetails, error) {
	var baseObject oci_mysql.CreateDbSystemSourceDetails
	//discriminator
	sourceTypeRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "source_type"))
	var sourceType string
	if ok {
		sourceType = sourceTypeRaw.(string)
	} else {
		sourceType = "NONE" // default value
	}
	switch strings.ToLower(sourceType) {
	case strings.ToLower("BACKUP"):
		details := oci_mysql.CreateDbSystemSourceFromBackupDetails{}
		if backupId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "backup_id")); ok {
			tmp := backupId.(string)
			details.BackupId = &tmp
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown source_type '%v' was specified", sourceType)
	}
	return baseObject, nil
}

func DbSystemSourceToMap(obj *oci_mysql.DbSystemSource) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (*obj).(type) {
	case oci_mysql.DbSystemSourceFromBackup:
		result["source_type"] = "BACKUP"

		if v.BackupId != nil {
			result["backup_id"] = string(*v.BackupId)
		}
	default:
		log.Printf("[WARN] Received 'source_type' of unknown type %v", *obj)
		return nil
	}

	return result
}

func (s *MysqlMysqlDbSystemResourceCrud) mapToCreateMaintenanceDetails(fieldKeyFormat string) (oci_mysql.CreateMaintenanceDetails, error) {
	result := oci_mysql.CreateMaintenanceDetails{}

	if windowStartTime, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "window_start_time")); ok {
		tmp := windowStartTime.(string)
		result.WindowStartTime = &tmp
	}

	return result, nil
}

func MaintenanceDetailsToMap(obj *oci_mysql.MaintenanceDetails) map[string]interface{} {
	result := map[string]interface{}{}

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

	result["status"] = string(obj.Status)

	if obj.StatusDetails != nil {
		result["status_details"] = string(*obj.StatusDetails)
	}

	return result
}

func (s *MysqlMysqlDbSystemResourceCrud) mapToUpdateBackupPolicyDetails(fieldKeyFormat string) (oci_mysql.UpdateBackupPolicyDetails, error) {
	result := oci_mysql.UpdateBackupPolicyDetails{}

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

	if isEnabled, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_enabled")); ok {
		tmp := isEnabled.(bool)
		result.IsEnabled = &tmp
	}

	if retentionInDays, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "retention_in_days")); ok {
		tmp := retentionInDays.(int)
		result.RetentionInDays = &tmp
	}

	if windowStartTime, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "window_start_time")); ok {
		tmp := windowStartTime.(string)
		result.WindowStartTime = &tmp
	}

	return result, nil
}

func (s *MysqlMysqlDbSystemResourceCrud) mapToUpdateMaintenanceDetails(fieldKeyFormat string) (oci_mysql.UpdateMaintenanceDetails, error) {
	result := oci_mysql.UpdateMaintenanceDetails{}

	if windowStartTime, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "window_start_time")); ok {
		tmp := windowStartTime.(string)
		result.WindowStartTime = &tmp
	}

	return result, nil
}

func (s *MysqlMysqlDbSystemResourceCrud) StartMysqlDbInstance() error {
	request := oci_mysql.StartDbSystemRequest{}

	tmp := s.D.Id()
	request.DbSystemId = &tmp

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "mysql")

	_, err := s.Client.StartDbSystem(context.Background(), request)
	if err != nil {
		return err
	}

	retentionPolicyFunc := func() bool { return s.Res.LifecycleState == oci_mysql.DbSystemLifecycleStateActive }
	return WaitForResourceCondition(s, retentionPolicyFunc, s.D.Timeout(schema.TimeoutUpdate))
}

func (s *MysqlMysqlDbSystemResourceCrud) StopMysqlDbInstance() error {
	request := oci_mysql.StopDbSystemRequest{}

	tmp := s.D.Id()
	request.DbSystemId = &tmp

	shutdownTypeRaw, ok := s.D.GetOkExists("shutdown_type")

	var shutdown_type string
	if ok {
		shutdown_type = shutdownTypeRaw.(string)
	} else {
		shutdown_type = "FAST"
	}

	switch strings.ToLower(shutdown_type) {
	case strings.ToLower("SLOW"):
		request.ShutdownType = oci_mysql.InnoDbShutdownModeSlow
	case strings.ToLower("FAST"):
		request.ShutdownType = oci_mysql.InnoDbShutdownModeFast
	case strings.ToLower("IMMEDIATE"):
		request.ShutdownType = oci_mysql.InnoDbShutdownModeImmediate
	default:
		return fmt.Errorf("unknown shutdown_type '%v' was specified", shutdown_type)
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "mysql")

	_, err := s.Client.StopDbSystem(context.Background(), request)
	if err != nil {
		return err
	}

	retentionPolicyFunc := func() bool { return s.Res.LifecycleState == oci_mysql.DbSystemLifecycleStateInactive }
	return WaitForResourceCondition(s, retentionPolicyFunc, s.D.Timeout(schema.TimeoutUpdate))
}
