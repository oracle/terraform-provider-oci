// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
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

func MediaServicesMediaAssetDistributionChannelAttachmentDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularMediaServicesMediaAssetDistributionChannelAttachment,
		Schema: map[string]*schema.Schema{
			"distribution_channel_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"media_asset_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"version": {
				Type:     schema.TypeString,
				Computed: true,
			},
			// Computed
			"display_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"locks": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"compartment_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"message": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"related_resource_id": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"time_created": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"type": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"is_lock_override": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"media_workflow_job_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"metadata_ref": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func readSingularMediaServicesMediaAssetDistributionChannelAttachment(d *schema.ResourceData, m interface{}) error {
	sync := &MediaServicesMediaAssetDistributionChannelAttachmentDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).MediaServicesClient()

	return tfresource.ReadResource(sync)
}

type MediaServicesMediaAssetDistributionChannelAttachmentDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_media_services.MediaServicesClient
	Res    *oci_media_services.GetMediaAssetDistributionChannelAttachmentResponse
}

func (s *MediaServicesMediaAssetDistributionChannelAttachmentDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *MediaServicesMediaAssetDistributionChannelAttachmentDataSourceCrud) Get() error {
	request := oci_media_services.GetMediaAssetDistributionChannelAttachmentRequest{}

	if distributionChannelId, ok := s.D.GetOkExists("distribution_channel_id"); ok {
		tmp := distributionChannelId.(string)
		request.DistributionChannelId = &tmp
	}

	if mediaAssetId, ok := s.D.GetOkExists("media_asset_id"); ok {
		tmp := mediaAssetId.(string)
		request.MediaAssetId = &tmp
	}

	if version, ok := s.D.GetOkExists("version"); ok {
		tmp := version.(string)
		tmpInt64, err := strconv.ParseInt(tmp, 10, 64)
		if err != nil {
			return fmt.Errorf("unable to convert version string: %s to an int64 and encountered error: %v", tmp, err)
		}
		request.Version = &tmpInt64
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "media_services")

	response, err := s.Client.GetMediaAssetDistributionChannelAttachment(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *MediaServicesMediaAssetDistributionChannelAttachmentDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("MediaServicesMediaAssetDistributionChannelAttachmentDataSource-", MediaServicesMediaAssetDistributionChannelAttachmentDataSource(), s.D))

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	locks := []interface{}{}
	for _, item := range s.Res.Locks {
		locks = append(locks, ResourceLockToMap(item))
	}
	s.D.Set("locks", locks)

	if s.Res.MediaWorkflowJobId != nil {
		s.D.Set("media_workflow_job_id", *s.Res.MediaWorkflowJobId)
	}

	if s.Res.MetadataRef != nil {
		s.D.Set("metadata_ref", *s.Res.MetadataRef)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.Version != nil {
		s.D.Set("version", strconv.FormatInt(*s.Res.Version, 10))
	}

	return nil
}

func MediaAssetDistributionChannelAttachmentSummaryToMap(obj oci_media_services.MediaAssetDistributionChannelAttachmentSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.DisplayName != nil {
		result["display_name"] = string(*obj.DisplayName)
	}

	if obj.DistributionChannelId != nil {
		result["distribution_channel_id"] = string(*obj.DistributionChannelId)
	}

	locks := []interface{}{}
	for _, item := range obj.Locks {
		locks = append(locks, ResourceLockToMap(item))
	}
	result["locks"] = locks

	if obj.MediaAssetId != nil {
		result["media_asset_id"] = string(*obj.MediaAssetId)
	}

	if obj.MediaWorkflowJobId != nil {
		result["media_workflow_job_id"] = string(*obj.MediaWorkflowJobId)
	}

	if obj.MetadataRef != nil {
		result["metadata_ref"] = string(*obj.MetadataRef)
	}

	result["state"] = string(obj.LifecycleState)

	if obj.Version != nil {
		result["version"] = strconv.FormatInt(*obj.Version, 10)
	}

	return result
}
