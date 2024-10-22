// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package datascience

import (
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strconv"
	"strings"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"

	"github.com/oracle/oci-go-sdk/v65/common"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"

	oci_datascience "github.com/oracle/oci-go-sdk/v65/datascience"
)

func DatascienceModelResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createDatascienceModel,
		Read:     readDatascienceModel,
		Update:   updateDatascienceModel,
		Delete:   deleteDatascienceModel,
		Schema: map[string]*schema.Schema{
			// Required
			"artifact_content_length": {
				Type:             schema.TypeString,
				Required:         true,
				ForceNew:         true,
				ValidateFunc:     tfresource.ValidateInt64TypeString,
				DiffSuppressFunc: tfresource.Int64StringDiffSuppressFunction,
			},
			"model_artifact": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"project_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"model_version_set_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"model_version_set_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"version_label": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			// Optional
			"backup_setting": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"backup_region": {
							Type:     schema.TypeString,
							Required: true,
						},
						"is_backup_enabled": {
							Type:     schema.TypeBool,
							Required: true,
						},

						// Optional
						"customer_notification_type": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},

						// Computed
					},
				},
			},
			"custom_metadata_list": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional
						"category": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"description": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"key": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"value": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"defined_metadata_list": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"category": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"description": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"key": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"value": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"artifact_content_disposition": {
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
			"input_schema": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"output_schema": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"retention_setting": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"archive_after_days": {
							Type:     schema.TypeInt,
							Required: true,
						},

						// Optional
						"customer_notification_type": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"delete_after_days": {
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"state": {
				Type:             schema.TypeString,
				Optional:         true,
				Computed:         true,
				DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
				ValidateFunc: validation.StringInSlice([]string{
					string(oci_datascience.ModelLifecycleStateActive),
					string(oci_datascience.ModelLifecycleStateInactive),
				}, true),
			},

			// Computed
			"backup_operation_details": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"backup_state": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"backup_state_details": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"time_last_backup": {
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
			"retention_operation_details": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"archive_state": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"archive_state_details": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"delete_state": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"delete_state_details": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"time_archival_scheduled": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"time_deletion_scheduled": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"artifact_content_md5": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"artifact_last_modified": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"created_by": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"empty_model": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"time_created": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func createDatascienceModel(d *schema.ResourceData, m interface{}) error {
	sync := &DatascienceModelResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataScienceClient()

	var deactivateModel = false
	if state, ok := sync.D.GetOkExists("state"); ok {
		desiredState := oci_datascience.ModelLifecycleStateEnum(strings.ToUpper(state.(string)))
		if desiredState == oci_datascience.ModelLifecycleStateInactive {
			deactivateModel = true
		}
	}

	if e := tfresource.CreateResource(d, sync); e != nil {
		return e
	}
	if deactivateModel {
		if e := sync.DeactivateModel(); e != nil {
			return e
		}
		sync.D.Set("state", oci_datascience.ModelLifecycleStateInactive)
	}
	if e := sync.CreateArtifact(); e != nil {
		return e
	}
	sync.D.Set("empty_model", false)
	return tfresource.ReadResource(sync)
}

func readDatascienceModel(d *schema.ResourceData, m interface{}) error {
	sync := &DatascienceModelResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataScienceClient()

	return tfresource.ReadResource(sync)
}

func updateDatascienceModel(d *schema.ResourceData, m interface{}) error {
	sync := &DatascienceModelResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataScienceClient()

	// Activate/Deactivate Model
	activate, deactivate := false, false

	if sync.D.HasChange("state") {
		desiredState := strings.ToUpper(sync.D.Get("state").(string))
		if oci_datascience.ModelLifecycleStateActive == oci_datascience.ModelLifecycleStateEnum(desiredState) {
			activate = true
		} else if oci_datascience.ModelLifecycleStateInactive == oci_datascience.ModelLifecycleStateEnum(desiredState) {
			deactivate = true
		}
	} else {
		currentState := strings.ToUpper(sync.D.Get("state").(string))
		if oci_datascience.ModelLifecycleStateActive == oci_datascience.ModelLifecycleStateEnum(currentState) {
			activate = true
			deactivate = true
		}
	}

	if deactivate {
		if err := sync.DeactivateModel(); err != nil {
			return err
		}
		sync.D.Set("state", oci_datascience.ModelLifecycleStateInactive)
	}
	if err := tfresource.UpdateResource(d, sync); err != nil {
		return err
	}

	if activate {
		if err := sync.ActivateModel(); err != nil {
			return err
		}
		sync.D.Set("state", oci_datascience.ModelLifecycleStateActive)
	}
	return nil
}

func deleteDatascienceModel(d *schema.ResourceData, m interface{}) error {
	sync := &DatascienceModelResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataScienceClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type HeadModelArtifact struct {
	ContentLength      *int64
	ContentDisposition *string
	ContentMd5         *string
	LastModified       *common.SDKTime
}

type DatascienceModelResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_datascience.DataScienceClient
	Res                    *oci_datascience.Model
	ArtifactHeadRes        *HeadModelArtifact
	DisableNotFoundRetries bool
}

func (s *DatascienceModelResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *DatascienceModelResourceCrud) CreatedPending() []string {
	return []string{}
}

func (s *DatascienceModelResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_datascience.ModelLifecycleStateActive),
	}
}

func (s *DatascienceModelResourceCrud) DeletedPending() []string {
	return []string{}
}

func (s *DatascienceModelResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_datascience.ModelLifecycleStateDeleted),
	}
}

func (s *DatascienceModelResourceCrud) Create() error {
	request := oci_datascience.CreateModelRequest{}

	if backupSetting, ok := s.D.GetOkExists("backup_setting"); ok {
		if tmpList := backupSetting.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "backup_setting", 0)
			tmp, err := s.mapToBackupSetting(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.BackupSetting = &tmp
		}
	}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if customMetadataList, ok := s.D.GetOkExists("custom_metadata_list"); ok {
		interfaces := customMetadataList.([]interface{})
		tmp := make([]oci_datascience.Metadata, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "custom_metadata_list", stateDataIndex)
			converted, err := s.mapToMetadata(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("custom_metadata_list") {
			request.CustomMetadataList = tmp
		}
	}

	if definedMetadataList, ok := s.D.GetOkExists("defined_metadata_list"); ok {
		interfaces := definedMetadataList.([]interface{})
		tmp := make([]oci_datascience.Metadata, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "defined_metadata_list", stateDataIndex)
			converted, err := s.mapToMetadataDefined(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("defined_metadata_list") {
			request.DefinedMetadataList = tmp
		}
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

	if inputSchema, ok := s.D.GetOkExists("input_schema"); ok {
		tmp := inputSchema.(string)
		request.InputSchema = &tmp
	}

	if outputSchema, ok := s.D.GetOkExists("output_schema"); ok {
		tmp := outputSchema.(string)
		request.OutputSchema = &tmp
	}

	if projectId, ok := s.D.GetOkExists("project_id"); ok {
		tmp := projectId.(string)
		request.ProjectId = &tmp
	}

	if retentionSetting, ok := s.D.GetOkExists("retention_setting"); ok {
		if tmpList := retentionSetting.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "retention_setting", 0)
			tmp, err := s.mapToRetentionSetting(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.RetentionSetting = &tmp
		}
	}

	if versionLabel, ok := s.D.GetOkExists("version_label"); ok {
		tmp := versionLabel.(string)
		request.VersionLabel = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "datascience")

	response, err := s.Client.CreateModel(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Model
	return nil
}

func (s *DatascienceModelResourceCrud) Get() error {
	request := oci_datascience.GetModelRequest{}

	tmp := s.D.Id()
	request.ModelId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "datascience")

	response, err := s.Client.GetModel(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Model
	if emptyModel, ok := s.D.GetOkExists("empty_model"); ok {
		tmp := emptyModel.(bool)
		if !tmp {
			err := s.GetArtifactHead()
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func (s *DatascienceModelResourceCrud) Update() error {
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(compartment)
			if err != nil {
				return err
			}
		}
	}
	request := oci_datascience.UpdateModelRequest{}

	if backupSetting, ok := s.D.GetOkExists("backup_setting"); ok {
		if tmpList := backupSetting.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "backup_setting", 0)
			tmp, err := s.mapToBackupSetting(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.BackupSetting = &tmp
		}
	}

	if customMetadataList, ok := s.D.GetOkExists("custom_metadata_list"); ok {
		interfaces := customMetadataList.([]interface{})
		tmp := make([]oci_datascience.Metadata, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "custom_metadata_list", stateDataIndex)
			converted, err := s.mapToMetadata(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("custom_metadata_list") {
			request.CustomMetadataList = tmp
		}
	}

	if definedMetadataList, ok := s.D.GetOkExists("defined_metadata_list"); ok {
		interfaces := definedMetadataList.([]interface{})
		tmp := make([]oci_datascience.Metadata, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "defined_metadata_list", stateDataIndex)
			converted, err := s.mapToMetadataDefined(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("defined_metadata_list") {
			request.DefinedMetadataList = tmp
		}
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

	tmp := s.D.Id()
	request.ModelId = &tmp

	if modelVersionSetId, ok := s.D.GetOkExists("model_version_set_id"); ok {
		tmp := modelVersionSetId.(string)
		request.ModelVersionSetId = &tmp
	}

	if retentionSetting, ok := s.D.GetOkExists("retention_setting"); ok {
		if tmpList := retentionSetting.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "retention_setting", 0)
			tmp, err := s.mapToRetentionSetting(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.RetentionSetting = &tmp
		}
	}

	if versionLabel, ok := s.D.GetOkExists("version_label"); ok {
		tmp := versionLabel.(string)
		request.VersionLabel = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "datascience")

	response, err := s.Client.UpdateModel(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Model
	return nil
}

func (s *DatascienceModelResourceCrud) Delete() error {
	request := oci_datascience.DeleteModelRequest{}

	tmp := s.D.Id()
	request.ModelId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "datascience")

	_, err := s.Client.DeleteModel(context.Background(), request)
	return err
}

func (s *DatascienceModelResourceCrud) SetData() error {
	if s.Res.BackupOperationDetails != nil {
		s.D.Set("backup_operation_details", []interface{}{BackupOperationDetailsToMap(s.Res.BackupOperationDetails)})
	} else {
		s.D.Set("backup_operation_details", nil)
	}

	if s.Res.BackupSetting != nil {
		s.D.Set("backup_setting", []interface{}{BackupSettingToMap(s.Res.BackupSetting)})
	} else {
		s.D.Set("backup_setting", nil)
	}

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.CreatedBy != nil {
		s.D.Set("created_by", *s.Res.CreatedBy)
	}

	customMetadataList := []interface{}{}
	for _, item := range s.Res.CustomMetadataList {
		customMetadataList = append(customMetadataList, MetadataToMap(item))
	}
	s.D.Set("custom_metadata_list", customMetadataList)

	definedMetadataList := []interface{}{}
	for _, item := range s.Res.DefinedMetadataList {
		definedMetadataList = append(definedMetadataList, MetadataToMap(item))
	}
	s.D.Set("defined_metadata_list", definedMetadataList)

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.Description != nil {
		s.D.Set("description", *s.Res.Description)
	}

	if s.Res.ModelVersionSetName != nil {
		s.D.Set("model_version_set_name", *s.Res.ModelVersionSetName)
	}

	if s.Res.ModelVersionSetId != nil {
		s.D.Set("model_version_set_id", *s.Res.ModelVersionSetId)
	}

	if s.Res.VersionLabel != nil {
		s.D.Set("version_label", *s.Res.VersionLabel)
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.InputSchema != nil {
		s.D.Set("input_schema", *s.Res.InputSchema)
	}

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	if s.Res.OutputSchema != nil {
		s.D.Set("output_schema", *s.Res.OutputSchema)
	}

	if s.Res.ProjectId != nil {
		s.D.Set("project_id", *s.Res.ProjectId)
	}

	if s.Res.RetentionOperationDetails != nil {
		s.D.Set("retention_operation_details", []interface{}{RetentionOperationDetailsToMap(s.Res.RetentionOperationDetails)})
	} else {
		s.D.Set("retention_operation_details", nil)
	}

	if s.Res.RetentionSetting != nil {
		s.D.Set("retention_setting", []interface{}{RetentionSettingToMap(s.Res.RetentionSetting)})
	} else {
		s.D.Set("retention_setting", nil)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	return s.SetArtifactData()
}

func BackupOperationDetailsToMap(obj *oci_datascience.BackupOperationDetails) map[string]interface{} {
	result := map[string]interface{}{}

	result["backup_state"] = string(obj.BackupState)

	if obj.BackupStateDetails != nil {
		result["backup_state_details"] = string(*obj.BackupStateDetails)
	}

	if obj.TimeLastBackup != nil {
		result["time_last_backup"] = obj.TimeLastBackup.String()
	}

	return result
}

func (s *DatascienceModelResourceCrud) mapToBackupSetting(fieldKeyFormat string) (oci_datascience.BackupSetting, error) {
	result := oci_datascience.BackupSetting{}

	if backupRegion, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "backup_region")); ok {
		tmp := backupRegion.(string)
		result.BackupRegion = &tmp
	}

	if customerNotificationType, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "customer_notification_type")); ok {
		result.CustomerNotificationType = oci_datascience.ModelSettingCustomerNotificationTypeEnum(customerNotificationType.(string))
	}

	if isBackupEnabled, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_backup_enabled")); ok {
		tmp := isBackupEnabled.(bool)
		result.IsBackupEnabled = &tmp
	}

	return result, nil
}

func BackupSettingToMap(obj *oci_datascience.BackupSetting) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.BackupRegion != nil {
		result["backup_region"] = string(*obj.BackupRegion)
	}

	result["customer_notification_type"] = string(obj.CustomerNotificationType)

	if obj.IsBackupEnabled != nil {
		result["is_backup_enabled"] = bool(*obj.IsBackupEnabled)
	}

	return result
}

func (s *DatascienceModelResourceCrud) mapToMetadata(fieldKeyFormat string) (oci_datascience.Metadata, error) {
	result := oci_datascience.Metadata{}

	if category, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "category")); ok {
		tmp := category.(string)
		result.Category = &tmp
	}

	if description, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "description")); ok {
		tmp := description.(string)
		result.Description = &tmp
	}

	if key, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "key")); ok {
		tmp := key.(string)
		result.Key = &tmp
	}

	if value, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "value")); ok {
		tmp := value.(string)
		result.Value = &tmp
	}

	return result, nil
}

/*
The New Mapping method is created for DefinedMetadataList to pass null values explicitly for 'category' and 'description'.
*/
func (s *DatascienceModelResourceCrud) mapToMetadataDefined(fieldKeyFormat string) (oci_datascience.Metadata, error) {
	result := oci_datascience.Metadata{}

	if key, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "key")); ok {
		tmp := key.(string)
		result.Key = &tmp
	}

	if value, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "value")); ok {
		tmp := value.(string)
		result.Value = &tmp
	}

	return result, nil
}

func MetadataToMap(obj oci_datascience.Metadata) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Category != nil {
		result["category"] = string(*obj.Category)
	}

	if obj.Description != nil {
		result["description"] = string(*obj.Description)
	}

	if obj.Key != nil {
		result["key"] = string(*obj.Key)
	}

	if obj.Value != nil {
		result["value"] = string(*obj.Value)
	}

	return result
}

func RetentionOperationDetailsToMap(obj *oci_datascience.RetentionOperationDetails) map[string]interface{} {
	result := map[string]interface{}{}

	result["archive_state"] = string(obj.ArchiveState)

	if obj.ArchiveStateDetails != nil {
		result["archive_state_details"] = string(*obj.ArchiveStateDetails)
	}

	result["delete_state"] = string(obj.DeleteState)

	if obj.DeleteStateDetails != nil {
		result["delete_state_details"] = string(*obj.DeleteStateDetails)
	}

	if obj.TimeArchivalScheduled != nil {
		result["time_archival_scheduled"] = obj.TimeArchivalScheduled.String()
	}

	if obj.TimeDeletionScheduled != nil {
		result["time_deletion_scheduled"] = obj.TimeDeletionScheduled.String()
	}

	return result
}

func (s *DatascienceModelResourceCrud) mapToRetentionSetting(fieldKeyFormat string) (oci_datascience.RetentionSetting, error) {
	result := oci_datascience.RetentionSetting{}

	if archiveAfterDays, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "archive_after_days")); ok {
		tmp := archiveAfterDays.(int)
		result.ArchiveAfterDays = &tmp
	}

	if customerNotificationType, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "customer_notification_type")); ok {
		result.CustomerNotificationType = oci_datascience.ModelSettingCustomerNotificationTypeEnum(customerNotificationType.(string))
	}

	if deleteAfterDays, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "delete_after_days")); ok {
		tmp := deleteAfterDays.(int)
		result.DeleteAfterDays = &tmp
	}

	return result, nil
}

func RetentionSettingToMap(obj *oci_datascience.RetentionSetting) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.ArchiveAfterDays != nil {
		result["archive_after_days"] = int(*obj.ArchiveAfterDays)
	}

	result["customer_notification_type"] = string(obj.CustomerNotificationType)

	if obj.DeleteAfterDays != nil {
		result["delete_after_days"] = int(*obj.DeleteAfterDays)
	}

	return result
}

func (s *DatascienceModelResourceCrud) updateCompartment(compartment interface{}) error {
	changeCompartmentRequest := oci_datascience.ChangeModelCompartmentRequest{}

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	idTmp := s.D.Id()
	changeCompartmentRequest.ModelId = &idTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "datascience")

	_, err := s.Client.ChangeModelCompartment(context.Background(), changeCompartmentRequest)
	if err != nil {
		return err
	}

	if waitErr := tfresource.WaitForUpdatedState(s.D, s); waitErr != nil {
		return waitErr
	}

	return nil
}

func (s *DatascienceModelResourceCrud) ActivateModel() error {
	request := oci_datascience.ActivateModelRequest{}

	tmp := s.D.Id()
	request.ModelId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "datascience")

	_, err := s.Client.ActivateModel(context.Background(), request)
	if err != nil {
		return err
	}

	retentionPolicyFunc := func() bool { return s.Res.LifecycleState == oci_datascience.ModelLifecycleStateActive }
	return tfresource.WaitForResourceCondition(s, retentionPolicyFunc, s.D.Timeout(schema.TimeoutUpdate))
}

func (s *DatascienceModelResourceCrud) DeactivateModel() error {
	request := oci_datascience.DeactivateModelRequest{}

	tmp := s.D.Id()
	request.ModelId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "datascience")

	_, err := s.Client.DeactivateModel(context.Background(), request)
	if err != nil {
		return err
	}

	retentionPolicyFunc := func() bool { return s.Res.LifecycleState == oci_datascience.ModelLifecycleStateInactive }
	return tfresource.WaitForResourceCondition(s, retentionPolicyFunc, s.D.Timeout(schema.TimeoutUpdate))
}

func (s *DatascienceModelResourceCrud) CreateArtifact() error {
	request := oci_datascience.CreateModelArtifactRequest{}

	if contentDisposition, ok := s.D.GetOkExists("artifact_content_disposition"); ok {
		tmp := contentDisposition.(string)
		request.ContentDisposition = &tmp
	}

	if contentLength, ok := s.D.GetOkExists("artifact_content_length"); ok {
		tmp := contentLength.(string)
		tmpInt64, err := strconv.ParseInt(tmp, 10, 64)
		if err != nil {
			return fmt.Errorf("unable to convert Content-Length string: %s to an int64 and encountered error: %v", tmp, err)
		}
		request.ContentLength = &tmpInt64
	}

	if modelArtifact, ok := s.D.GetOkExists("model_artifact"); ok {
		tmp := modelArtifact.(string)
		var artifactReader io.Reader
		artifactReader, err := os.Open(tmp)
		if err != nil {
			return fmt.Errorf("the specified model_artifact is not available: %q", err)
		}
		request.ModelArtifact = ioutil.NopCloser(artifactReader)
	}

	request.ModelId = s.Res.Id

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "datascience")

	_, err := s.Client.CreateModelArtifact(context.Background(), request)
	if err != nil {
		return err
	}
	return nil
}

func (s *DatascienceModelResourceCrud) GetArtifactHead() error {
	request := oci_datascience.HeadModelArtifactRequest{}

	tmp := s.D.Id()
	request.ModelId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "datascience")

	response, err := s.Client.HeadModelArtifact(context.Background(), request)
	if err != nil {
		return err
	}

	s.ArtifactHeadRes = &HeadModelArtifact{
		ContentLength:      response.ContentLength,
		ContentDisposition: response.ContentDisposition,
		ContentMd5:         response.ContentMd5,
		LastModified:       response.LastModified,
	}
	return nil
}

func (s *DatascienceModelResourceCrud) SetArtifactData() error {
	if s.ArtifactHeadRes != nil {
		if s.ArtifactHeadRes.ContentDisposition != nil {
			s.D.Set("artifact_content_disposition", *s.ArtifactHeadRes.ContentDisposition)
		}

		if s.ArtifactHeadRes.ContentLength != nil {
			s.D.Set("artifact_content_length", strconv.FormatInt(*s.ArtifactHeadRes.ContentLength, 10))
		}

		if s.ArtifactHeadRes.ContentMd5 != nil {
			s.D.Set("artifact_content_md5", *s.ArtifactHeadRes.ContentMd5)
		}

		if s.ArtifactHeadRes.LastModified != nil {
			s.D.Set("artifact_last_modified", s.ArtifactHeadRes.LastModified.String())
		}

		s.D.Set("empty_model", false)
	}

	return nil
}
