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

func MediaServicesMediaAssetsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readMediaServicesMediaAssets,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"bucket": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"distribution_channel_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"master_media_asset_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"media_workflow_job_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"object": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"parent_media_asset_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"source_media_workflow_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"source_media_workflow_version": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"type": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"media_asset_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem:     tfresource.GetDataSourceItemSchema(MediaServicesMediaAssetResource()),
						},
					},
				},
			},
		},
	}
}

func readMediaServicesMediaAssets(d *schema.ResourceData, m interface{}) error {
	sync := &MediaServicesMediaAssetsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).MediaServicesClient()

	return tfresource.ReadResource(sync)
}

type MediaServicesMediaAssetsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_media_services.MediaServicesClient
	Res    *oci_media_services.ListMediaAssetsResponse
}

func (s *MediaServicesMediaAssetsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *MediaServicesMediaAssetsDataSourceCrud) Get() error {
	request := oci_media_services.ListMediaAssetsRequest{}

	if bucket, ok := s.D.GetOkExists("bucket"); ok {
		tmp := bucket.(string)
		request.BucketName = &tmp
	}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if distributionChannelId, ok := s.D.GetOkExists("distribution_channel_id"); ok {
		tmp := distributionChannelId.(string)
		request.DistributionChannelId = &tmp
	}

	if masterMediaAssetId, ok := s.D.GetOkExists("master_media_asset_id"); ok {
		tmp := masterMediaAssetId.(string)
		request.MasterMediaAssetId = &tmp
	}

	if mediaWorkflowJobId, ok := s.D.GetOkExists("media_workflow_job_id"); ok {
		tmp := mediaWorkflowJobId.(string)
		request.MediaWorkflowJobId = &tmp
	}

	if object, ok := s.D.GetOkExists("object"); ok {
		tmp := object.(string)
		request.ObjectName = &tmp
	}

	if parentMediaAssetId, ok := s.D.GetOkExists("parent_media_asset_id"); ok {
		tmp := parentMediaAssetId.(string)
		request.ParentMediaAssetId = &tmp
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

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_media_services.ListMediaAssetsLifecycleStateEnum(state.(string))
	}

	if type_, ok := s.D.GetOkExists("type"); ok {
		request.Type = oci_media_services.ListMediaAssetsTypeEnum(type_.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "media_services")

	response, err := s.Client.ListMediaAssets(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListMediaAssets(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *MediaServicesMediaAssetsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("MediaServicesMediaAssetsDataSource-", MediaServicesMediaAssetsDataSource(), s.D))
	resources := []map[string]interface{}{}
	mediaAsset := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, MediaAssetSummaryToMap(item))
	}
	mediaAsset["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, MediaServicesMediaAssetsDataSource().Schema["media_asset_collection"].Elem.(*schema.Resource).Schema)
		mediaAsset["items"] = items
	}

	resources = append(resources, mediaAsset)
	if err := s.D.Set("media_asset_collection", resources); err != nil {
		return err
	}

	return nil
}
