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

func MediaServicesStreamDistributionChannelsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readMediaServicesStreamDistributionChannels,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"stream_distribution_channel_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem:     tfresource.GetDataSourceItemSchema(MediaServicesStreamDistributionChannelResource()),
						},
					},
				},
			},
		},
	}
}

func readMediaServicesStreamDistributionChannels(d *schema.ResourceData, m interface{}) error {
	sync := &MediaServicesStreamDistributionChannelsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).MediaServicesClient()

	return tfresource.ReadResource(sync)
}

type MediaServicesStreamDistributionChannelsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_media_services.MediaServicesClient
	Res    *oci_media_services.ListStreamDistributionChannelsResponse
}

func (s *MediaServicesStreamDistributionChannelsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *MediaServicesStreamDistributionChannelsDataSourceCrud) Get() error {
	request := oci_media_services.ListStreamDistributionChannelsRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if id, ok := s.D.GetOkExists("id"); ok {
		tmp := id.(string)
		request.Id = &tmp
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_media_services.StreamDistributionChannelLifecycleStateEnum(state.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "media_services")

	response, err := s.Client.ListStreamDistributionChannels(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListStreamDistributionChannels(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *MediaServicesStreamDistributionChannelsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("MediaServicesStreamDistributionChannelsDataSource-", MediaServicesStreamDistributionChannelsDataSource(), s.D))
	resources := []map[string]interface{}{}
	streamDistributionChannel := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, StreamDistributionChannelSummaryToMap(item))
	}
	streamDistributionChannel["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, MediaServicesStreamDistributionChannelsDataSource().Schema["stream_distribution_channel_collection"].Elem.(*schema.Resource).Schema)
		streamDistributionChannel["items"] = items
	}

	resources = append(resources, streamDistributionChannel)
	if err := s.D.Set("stream_distribution_channel_collection", resources); err != nil {
		return err
	}

	return nil
}
