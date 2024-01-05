// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package media_services

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_media_services "github.com/oracle/oci-go-sdk/v65/mediaservices"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func MediaServicesStreamPackagingConfigsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readMediaServicesStreamPackagingConfigs,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"distribution_channel_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"stream_packaging_config_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"stream_packaging_config_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem:     tfresource.GetDataSourceItemSchema(MediaServicesStreamPackagingConfigResource()),
						},
					},
				},
			},
		},
	}
}

func readMediaServicesStreamPackagingConfigs(d *schema.ResourceData, m interface{}) error {
	sync := &MediaServicesStreamPackagingConfigsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).MediaServicesClient()

	return tfresource.ReadResource(sync)
}

type MediaServicesStreamPackagingConfigsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_media_services.MediaServicesClient
	Res    *oci_media_services.ListStreamPackagingConfigsResponse
}

func (s *MediaServicesStreamPackagingConfigsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *MediaServicesStreamPackagingConfigsDataSourceCrud) Get() error {
	request := oci_media_services.ListStreamPackagingConfigsRequest{}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if distributionChannelId, ok := s.D.GetOkExists("distribution_channel_id"); ok {
		tmp := distributionChannelId.(string)
		request.DistributionChannelId = &tmp
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_media_services.StreamPackagingConfigLifecycleStateEnum(state.(string))
	}

	if streamPackagingConfigId, ok := s.D.GetOkExists("id"); ok {
		tmp := streamPackagingConfigId.(string)
		request.StreamPackagingConfigId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "media_services")

	response, err := s.Client.ListStreamPackagingConfigs(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListStreamPackagingConfigs(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *MediaServicesStreamPackagingConfigsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("MediaServicesStreamPackagingConfigsDataSource-", MediaServicesStreamPackagingConfigsDataSource(), s.D))
	resources := []map[string]interface{}{}
	streamPackagingConfig := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, StreamPackagingConfigSummaryToMap(item))
	}
	streamPackagingConfig["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, MediaServicesStreamPackagingConfigsDataSource().Schema["stream_packaging_config_collection"].Elem.(*schema.Resource).Schema)
		streamPackagingConfig["items"] = items
	}

	resources = append(resources, streamPackagingConfig)
	if err := s.D.Set("stream_packaging_config_collection", resources); err != nil {
		return err
	}

	return nil
}
