// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package media_services

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_media_services "github.com/oracle/oci-go-sdk/v65/mediaservices"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func MediaServicesStreamPackagingConfigDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["stream_packaging_config_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(MediaServicesStreamPackagingConfigResource(), fieldMap, readSingularMediaServicesStreamPackagingConfig)
}

func readSingularMediaServicesStreamPackagingConfig(d *schema.ResourceData, m interface{}) error {
	sync := &MediaServicesStreamPackagingConfigDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).MediaServicesClient()

	return tfresource.ReadResource(sync)
}

type MediaServicesStreamPackagingConfigDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_media_services.MediaServicesClient
	Res    *oci_media_services.GetStreamPackagingConfigResponse
}

func (s *MediaServicesStreamPackagingConfigDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *MediaServicesStreamPackagingConfigDataSourceCrud) Get() error {
	request := oci_media_services.GetStreamPackagingConfigRequest{}

	if streamPackagingConfigId, ok := s.D.GetOkExists("stream_packaging_config_id"); ok {
		tmp := streamPackagingConfigId.(string)
		request.StreamPackagingConfigId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "media_services")

	response, err := s.Client.GetStreamPackagingConfig(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *MediaServicesStreamPackagingConfigDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.GetId())

	if s.Res.GetTimeCreated() != nil {
		s.D.Set("time_created", s.Res.GetTimeCreated().String())
	}

	if s.Res.GetTimeUpdated() != nil {
		s.D.Set("time_updated", s.Res.GetTimeUpdated().String())
	}

	if s.Res.GetCompartmentId() != nil {
		s.D.Set("compartment_id", *s.Res.GetCompartmentId())
	}

	switch v := (s.Res.StreamPackagingConfig).(type) {
	case oci_media_services.DashStreamPackagingConfig:
		s.D.Set("stream_packaging_format", "DASH")

		if v.DefinedTags != nil {
			s.D.Set("defined_tags", tfresource.DefinedTagsToMap(v.DefinedTags))
		}

		if v.DisplayName != nil {
			s.D.Set("display_name", *v.DisplayName)
		}

		if v.DistributionChannelId != nil {
			s.D.Set("distribution_channel_id", *v.DistributionChannelId)
		}

		if v.Encryption != nil {
			encryptionArray := []interface{}{}
			if encryptionMap := StreamPackagingConfigEncryptionToMap(&v.Encryption); encryptionMap != nil {
				encryptionArray = append(encryptionArray, encryptionMap)
			}
			s.D.Set("encryption", encryptionArray)
		} else {
			s.D.Set("encryption", nil)
		}

		s.D.Set("freeform_tags", v.FreeformTags)

		if v.SegmentTimeInSeconds != nil {
			s.D.Set("segment_time_in_seconds", *v.SegmentTimeInSeconds)
		}

		s.D.Set("state", v.LifecycleState)
	case oci_media_services.HlsStreamPackagingConfig:
		s.D.Set("stream_packaging_format", "HLS")

		if v.DefinedTags != nil {
			s.D.Set("defined_tags", tfresource.DefinedTagsToMap(v.DefinedTags))
		}

		if v.DisplayName != nil {
			s.D.Set("display_name", *v.DisplayName)
		}

		if v.DistributionChannelId != nil {
			s.D.Set("distribution_channel_id", *v.DistributionChannelId)
		}

		if v.Encryption != nil {
			encryptionArray := []interface{}{}
			if encryptionMap := StreamPackagingConfigEncryptionToMap(&v.Encryption); encryptionMap != nil {
				encryptionArray = append(encryptionArray, encryptionMap)
			}
			s.D.Set("encryption", encryptionArray)
		} else {
			s.D.Set("encryption", nil)
		}

		s.D.Set("freeform_tags", v.FreeformTags)

		if v.SegmentTimeInSeconds != nil {
			s.D.Set("segment_time_in_seconds", *v.SegmentTimeInSeconds)
		}

		s.D.Set("state", v.LifecycleState)
	default:
		log.Printf("[WARN] Received 'stream_packaging_format' of unknown type %v", s.Res.StreamPackagingConfig)
		return nil
	}

	return nil
}
