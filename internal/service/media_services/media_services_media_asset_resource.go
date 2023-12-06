// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package media_services

import (
	"context"
	"fmt"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	oci_media_services "github.com/oracle/oci-go-sdk/v65/mediaservices"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func MediaServicesMediaAssetResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createMediaServicesMediaAsset,
		Read:     readMediaServicesMediaAsset,
		Update:   updateMediaServicesMediaAsset,
		Delete:   deleteMediaServicesMediaAsset,
		Schema: map[string]*schema.Schema{
			// Required
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"type": {
				Type:     schema.TypeString,
				Required: true,
			},

			// Optional
			"bucket": {
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
			"master_media_asset_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"media_asset_tags": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"value": {
							Type:     schema.TypeString,
							Required: true,
						},

						// Optional
						"type": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},

						// Computed
					},
				},
			},
			"media_workflow_job_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"metadata": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"metadata": {
							Type:     schema.TypeString,
							Required: true,
						},

						// Optional

						// Computed
					},
				},
			},
			"namespace": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"object": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"object_etag": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"parent_media_asset_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"segment_range_end_index": {
				Type:             schema.TypeString,
				Optional:         true,
				Computed:         true,
				ForceNew:         true,
				ValidateFunc:     tfresource.ValidateInt64TypeString,
				DiffSuppressFunc: tfresource.Int64StringDiffSuppressFunction,
			},
			"segment_range_start_index": {
				Type:             schema.TypeString,
				Optional:         true,
				Computed:         true,
				ForceNew:         true,
				ValidateFunc:     tfresource.ValidateInt64TypeString,
				DiffSuppressFunc: tfresource.Int64StringDiffSuppressFunction,
			},
			"source_media_workflow_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"source_media_workflow_version": {
				Type:             schema.TypeString,
				Optional:         true,
				Computed:         true,
				ForceNew:         true,
				ValidateFunc:     tfresource.ValidateInt64TypeString,
				DiffSuppressFunc: tfresource.Int64StringDiffSuppressFunction,
			},

			// Computed
			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"system_tags": {
				Type:     schema.TypeMap,
				Computed: true,
				Elem:     schema.TypeString,
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

func createMediaServicesMediaAsset(d *schema.ResourceData, m interface{}) error {
	sync := &MediaServicesMediaAssetResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).MediaServicesClient()

	return tfresource.CreateResource(d, sync)
}

func readMediaServicesMediaAsset(d *schema.ResourceData, m interface{}) error {
	sync := &MediaServicesMediaAssetResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).MediaServicesClient()

	return tfresource.ReadResource(sync)
}

func updateMediaServicesMediaAsset(d *schema.ResourceData, m interface{}) error {
	sync := &MediaServicesMediaAssetResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).MediaServicesClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteMediaServicesMediaAsset(d *schema.ResourceData, m interface{}) error {
	sync := &MediaServicesMediaAssetResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).MediaServicesClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type MediaServicesMediaAssetResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_media_services.MediaServicesClient
	Res                    *oci_media_services.MediaAsset
	DisableNotFoundRetries bool
}

func (s *MediaServicesMediaAssetResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *MediaServicesMediaAssetResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_media_services.LifecycleStateCreating),
	}
}

func (s *MediaServicesMediaAssetResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_media_services.LifecycleStateActive),
	}
}

func (s *MediaServicesMediaAssetResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_media_services.LifecycleStateDeleting),
	}
}

func (s *MediaServicesMediaAssetResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_media_services.LifecycleStateDeleted),
	}
}

func (s *MediaServicesMediaAssetResourceCrud) Create() error {
	request := oci_media_services.CreateMediaAssetRequest{}

	if bucket, ok := s.D.GetOkExists("bucket"); ok {
		tmp := bucket.(string)
		request.BucketName = &tmp
	}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
		convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return err
		}
		request.DefinedTags = convertedDefinedTags
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if masterMediaAssetId, ok := s.D.GetOkExists("master_media_asset_id"); ok {
		tmp := masterMediaAssetId.(string)
		request.MasterMediaAssetId = &tmp
	}

	if mediaAssetTags, ok := s.D.GetOkExists("media_asset_tags"); ok {
		interfaces := mediaAssetTags.([]interface{})
		tmp := make([]oci_media_services.MediaAssetTag, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "media_asset_tags", stateDataIndex)
			converted, err := s.mapToMediaAssetTag(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("media_asset_tags") {
			request.MediaAssetTags = tmp
		}
	}

	if mediaWorkflowJobId, ok := s.D.GetOkExists("media_workflow_job_id"); ok {
		tmp := mediaWorkflowJobId.(string)
		request.MediaWorkflowJobId = &tmp
	}

	if metadata, ok := s.D.GetOkExists("metadata"); ok {
		interfaces := metadata.([]interface{})
		tmp := make([]oci_media_services.Metadata, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "metadata", stateDataIndex)
			converted, err := s.mapToMetadata(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("metadata") {
			request.Metadata = tmp
		}
	}

	if namespace, ok := s.D.GetOkExists("namespace"); ok {
		tmp := namespace.(string)
		request.NamespaceName = &tmp
	}

	if object, ok := s.D.GetOkExists("object"); ok {
		tmp := object.(string)
		request.ObjectName = &tmp
	}

	if objectEtag, ok := s.D.GetOkExists("object_etag"); ok {
		tmp := objectEtag.(string)
		request.ObjectEtag = &tmp
	}

	if parentMediaAssetId, ok := s.D.GetOkExists("parent_media_asset_id"); ok {
		tmp := parentMediaAssetId.(string)
		request.ParentMediaAssetId = &tmp
	}

	if segmentRangeEndIndex, ok := s.D.GetOkExists("segment_range_end_index"); ok {
		tmp := segmentRangeEndIndex.(string)
		tmpInt64, err := strconv.ParseInt(tmp, 10, 64)
		if err != nil {
			return fmt.Errorf("unable to convert segmentRangeEndIndex string: %s to an int64 and encountered error: %v", tmp, err)
		}
		request.SegmentRangeEndIndex = &tmpInt64
	}

	if segmentRangeStartIndex, ok := s.D.GetOkExists("segment_range_start_index"); ok {
		tmp := segmentRangeStartIndex.(string)
		tmpInt64, err := strconv.ParseInt(tmp, 10, 64)
		if err != nil {
			return fmt.Errorf("unable to convert segmentRangeStartIndex string: %s to an int64 and encountered error: %v", tmp, err)
		}
		request.SegmentRangeStartIndex = &tmpInt64
	}

	if sourceMediaWorkflowId, ok := s.D.GetOkExists("source_media_workflow_id"); ok {
		tmp := sourceMediaWorkflowId.(string)
		request.SourceMediaWorkflowId = &tmp
	}

	if sourceMediaWorkflowVersion, ok := s.D.GetOkExists("source_media_workflow_version"); ok {
		tmp := sourceMediaWorkflowVersion.(string)
		tmpInt64, err := strconv.ParseInt(tmp, 10, 64)
		if err != nil {
			return fmt.Errorf("unable to convert sourceMediaWorkflowVersion string: %s to an int64 and encountered error: %v", tmp, err)
		}
		request.SourceMediaWorkflowVersion = &tmpInt64
	}

	if type_, ok := s.D.GetOkExists("type"); ok {
		request.Type = oci_media_services.AssetTypeEnum(type_.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "media_services")

	response, err := s.Client.CreateMediaAsset(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.MediaAsset
	return nil
}

func (s *MediaServicesMediaAssetResourceCrud) Get() error {
	request := oci_media_services.GetMediaAssetRequest{}

	tmp := s.D.Id()
	request.MediaAssetId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "media_services")

	response, err := s.Client.GetMediaAsset(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.MediaAsset
	return nil
}

func (s *MediaServicesMediaAssetResourceCrud) Update() error {
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(compartment)
			if err != nil {
				return err
			}
		}
	}
	request := oci_media_services.UpdateMediaAssetRequest{}

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
		convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return err
		}
		request.DefinedTags = convertedDefinedTags
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if masterMediaAssetId, ok := s.D.GetOkExists("master_media_asset_id"); ok {
		tmp := masterMediaAssetId.(string)
		request.MasterMediaAssetId = &tmp
	}

	tmp := s.D.Id()
	request.MediaAssetId = &tmp

	if mediaAssetTags, ok := s.D.GetOkExists("media_asset_tags"); ok {
		interfaces := mediaAssetTags.([]interface{})
		tmp := make([]oci_media_services.MediaAssetTag, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "media_asset_tags", stateDataIndex)
			converted, err := s.mapToMediaAssetTag(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("media_asset_tags") {
			request.MediaAssetTags = tmp
		}
	}

	if metadata, ok := s.D.GetOkExists("metadata"); ok {
		interfaces := metadata.([]interface{})
		tmp := make([]oci_media_services.Metadata, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "metadata", stateDataIndex)
			converted, err := s.mapToMetadata(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("metadata") {
			request.Metadata = tmp
		}
	}

	if parentMediaAssetId, ok := s.D.GetOkExists("parent_media_asset_id"); ok {
		tmp := parentMediaAssetId.(string)
		request.ParentMediaAssetId = &tmp
	}

	if type_, ok := s.D.GetOkExists("type"); ok {
		request.Type = oci_media_services.AssetTypeEnum(type_.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "media_services")

	response, err := s.Client.UpdateMediaAsset(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.MediaAsset
	return nil
}

func (s *MediaServicesMediaAssetResourceCrud) Delete() error {
	request := oci_media_services.DeleteMediaAssetRequest{}

	if deleteMode, ok := s.D.GetOkExists("delete_mode"); ok {
		request.DeleteMode = oci_media_services.DeleteMediaAssetDeleteModeEnum(deleteMode.(string))
	}

	tmp := s.D.Id()
	request.MediaAssetId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "media_services")

	_, err := s.Client.DeleteMediaAsset(context.Background(), request)
	return err
}

func (s *MediaServicesMediaAssetResourceCrud) SetData() error {
	if s.Res.BucketName != nil {
		s.D.Set("bucket", *s.Res.BucketName)
	}

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)
	if s.Res.MasterMediaAssetId != nil {
		s.D.Set("master_media_asset_id", *s.Res.MasterMediaAssetId)
	}

	mediaAssetTags := []interface{}{}
	for _, item := range s.Res.MediaAssetTags {
		mediaAssetTags = append(mediaAssetTags, MediaAssetTagToMap(item))
	}
	s.D.Set("media_asset_tags", mediaAssetTags)

	if s.Res.MediaWorkflowJobId != nil {
		s.D.Set("media_workflow_job_id", *s.Res.MediaWorkflowJobId)
	}

	metadata := []interface{}{}
	for _, item := range s.Res.Metadata {
		metadata = append(metadata, MetadataToMap(item))
	}
	s.D.Set("metadata", metadata)

	if s.Res.NamespaceName != nil {
		s.D.Set("namespace", *s.Res.NamespaceName)
	}

	if s.Res.ObjectName != nil {
		s.D.Set("object", *s.Res.ObjectName)
	}

	if s.Res.ObjectEtag != nil {
		s.D.Set("object_etag", *s.Res.ObjectEtag)
	}

	if s.Res.ParentMediaAssetId != nil {
		s.D.Set("parent_media_asset_id", *s.Res.ParentMediaAssetId)
	}

	if s.Res.SegmentRangeEndIndex != nil {
		s.D.Set("segment_range_end_index", strconv.FormatInt(*s.Res.SegmentRangeEndIndex, 10))
	}

	if s.Res.SegmentRangeStartIndex != nil {
		s.D.Set("segment_range_start_index", strconv.FormatInt(*s.Res.SegmentRangeStartIndex, 10))
	}

	if s.Res.SourceMediaWorkflowId != nil {
		s.D.Set("source_media_workflow_id", *s.Res.SourceMediaWorkflowId)
	}

	if s.Res.SourceMediaWorkflowVersion != nil {
		s.D.Set("source_media_workflow_version", strconv.FormatInt(*s.Res.SourceMediaWorkflowVersion, 10))
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.SystemTags != nil {
		s.D.Set("system_tags", tfresource.SystemTagsToMap(s.Res.SystemTags))
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	s.D.Set("type", s.Res.Type)

	return nil
}

func MediaAssetSummaryToMap(obj oci_media_services.MediaAssetSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.CompartmentId != nil {
		result["compartment_id"] = string(*obj.CompartmentId)
	}

	if obj.DefinedTags != nil {
		result["defined_tags"] = tfresource.DefinedTagsToMap(obj.DefinedTags)
	}

	if obj.DisplayName != nil {
		result["display_name"] = string(*obj.DisplayName)
	}

	result["freeform_tags"] = obj.FreeformTags

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	if obj.MasterMediaAssetId != nil {
		result["master_media_asset_id"] = string(*obj.MasterMediaAssetId)
	}

	if obj.ParentMediaAssetId != nil {
		result["parent_media_asset_id"] = string(*obj.ParentMediaAssetId)
	}

	result["state"] = string(obj.LifecycleState)

	if obj.SystemTags != nil {
		result["system_tags"] = tfresource.SystemTagsToMap(obj.SystemTags)
	}

	if obj.TimeCreated != nil {
		result["time_created"] = obj.TimeCreated.String()
	}

	if obj.TimeUpdated != nil {
		result["time_updated"] = obj.TimeUpdated.String()
	}

	result["type"] = string(obj.Type)

	return result
}

func (s *MediaServicesMediaAssetResourceCrud) mapToMediaAssetTag(fieldKeyFormat string) (oci_media_services.MediaAssetTag, error) {
	result := oci_media_services.MediaAssetTag{}

	if type_, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "type")); ok {
		result.Type = oci_media_services.MediaAssetTagTypeEnum(type_.(string))
	}

	if value, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "value")); ok {
		tmp := value.(string)
		result.Value = &tmp
	}

	return result, nil
}

func MediaAssetTagToMap(obj oci_media_services.MediaAssetTag) map[string]interface{} {
	result := map[string]interface{}{}

	result["type"] = string(obj.Type)

	if obj.Value != nil {
		result["value"] = string(*obj.Value)
	}

	return result
}

func (s *MediaServicesMediaAssetResourceCrud) mapToMetadata(fieldKeyFormat string) (oci_media_services.Metadata, error) {
	result := oci_media_services.Metadata{}

	if metadata, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "metadata")); ok {
		tmp := metadata.(string)
		result.Metadata = &tmp
	}

	return result, nil
}

func MetadataToMap(obj oci_media_services.Metadata) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Metadata != nil {
		result["metadata"] = string(*obj.Metadata)
	}

	return result
}

func (s *MediaServicesMediaAssetResourceCrud) updateCompartment(compartment interface{}) error {
	changeCompartmentRequest := oci_media_services.ChangeMediaAssetCompartmentRequest{}

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	idTmp := s.D.Id()
	changeCompartmentRequest.MediaAssetId = &idTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "media_services")

	_, err := s.Client.ChangeMediaAssetCompartment(context.Background(), changeCompartmentRequest)
	if err != nil {
		return err
	}

	if waitErr := tfresource.WaitForUpdatedState(s.D, s); waitErr != nil {
		return waitErr
	}

	return nil
}
