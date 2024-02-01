// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package media_services

import (
	"context"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_media_services "github.com/oracle/oci-go-sdk/v65/mediaservices"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func MediaServicesMediaAssetDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["media_asset_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(MediaServicesMediaAssetResource(), fieldMap, readSingularMediaServicesMediaAsset)
}

func readSingularMediaServicesMediaAsset(d *schema.ResourceData, m interface{}) error {
	sync := &MediaServicesMediaAssetDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).MediaServicesClient()

	return tfresource.ReadResource(sync)
}

type MediaServicesMediaAssetDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_media_services.MediaServicesClient
	Res    *oci_media_services.GetMediaAssetResponse
}

func (s *MediaServicesMediaAssetDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *MediaServicesMediaAssetDataSourceCrud) Get() error {
	request := oci_media_services.GetMediaAssetRequest{}

	if mediaAssetId, ok := s.D.GetOkExists("media_asset_id"); ok {
		tmp := mediaAssetId.(string)
		request.MediaAssetId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "media_services")

	response, err := s.Client.GetMediaAsset(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *MediaServicesMediaAssetDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

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

	locks := []interface{}{}
	for _, item := range s.Res.Locks {
		locks = append(locks, ResourceLockToMap(item))
	}
	s.D.Set("locks", locks)

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
