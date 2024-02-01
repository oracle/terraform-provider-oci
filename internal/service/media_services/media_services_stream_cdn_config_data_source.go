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

func MediaServicesStreamCdnConfigDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["stream_cdn_config_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(MediaServicesStreamCdnConfigResource(), fieldMap, readSingularMediaServicesStreamCdnConfig)
}

func readSingularMediaServicesStreamCdnConfig(d *schema.ResourceData, m interface{}) error {
	sync := &MediaServicesStreamCdnConfigDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).MediaServicesClient()

	return tfresource.ReadResource(sync)
}

type MediaServicesStreamCdnConfigDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_media_services.MediaServicesClient
	Res    *oci_media_services.GetStreamCdnConfigResponse
}

func (s *MediaServicesStreamCdnConfigDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *MediaServicesStreamCdnConfigDataSourceCrud) Get() error {
	request := oci_media_services.GetStreamCdnConfigRequest{}

	if streamCdnConfigId, ok := s.D.GetOkExists("stream_cdn_config_id"); ok {
		tmp := streamCdnConfigId.(string)
		request.StreamCdnConfigId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "media_services")

	response, err := s.Client.GetStreamCdnConfig(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *MediaServicesStreamCdnConfigDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.Config != nil {
		configArray := []interface{}{}
		if configMap := StreamCdnConfigSectionToMap(&s.Res.Config); configMap != nil {
			configArray = append(configArray, configMap)
		}
		s.D.Set("config", configArray)
	} else {
		s.D.Set("config", nil)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	if s.Res.DistributionChannelId != nil {
		s.D.Set("distribution_channel_id", *s.Res.DistributionChannelId)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.IsEnabled != nil {
		s.D.Set("is_enabled", *s.Res.IsEnabled)
	}

	if s.Res.LifecyleDetails != nil {
		s.D.Set("lifecyle_details", *s.Res.LifecyleDetails)
	}

	locks := []interface{}{}
	for _, item := range s.Res.Locks {
		locks = append(locks, ResourceLockToMap(item))
	}
	s.D.Set("locks", locks)

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

	return nil
}
