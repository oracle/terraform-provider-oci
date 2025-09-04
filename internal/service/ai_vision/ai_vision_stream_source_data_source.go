// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package ai_vision

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_ai_vision "github.com/oracle/oci-go-sdk/v65/aivision"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func AiVisionStreamSourceDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["stream_source_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(AiVisionStreamSourceResource(), fieldMap, readSingularAiVisionStreamSource)
}

func readSingularAiVisionStreamSource(d *schema.ResourceData, m interface{}) error {
	sync := &AiVisionStreamSourceDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).AiServiceVisionClient()

	return tfresource.ReadResource(sync)
}

type AiVisionStreamSourceDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_ai_vision.AIServiceVisionClient
	Res    *oci_ai_vision.GetStreamSourceResponse
}

func (s *AiVisionStreamSourceDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *AiVisionStreamSourceDataSourceCrud) Get() error {
	request := oci_ai_vision.GetStreamSourceRequest{}

	if streamSourceId, ok := s.D.GetOkExists("stream_source_id"); ok {
		tmp := streamSourceId.(string)
		request.StreamSourceId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "ai_vision")

	response, err := s.Client.GetStreamSource(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *AiVisionStreamSourceDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

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

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.StreamSourceDetails != nil {
		streamSourceDetailsArray := []interface{}{}
		if streamSourceDetailsMap := StreamSourceDetailsToMap(&s.Res.StreamSourceDetails); streamSourceDetailsMap != nil {
			streamSourceDetailsArray = append(streamSourceDetailsArray, streamSourceDetailsMap)
		}
		s.D.Set("stream_source_details", streamSourceDetailsArray)
	} else {
		s.D.Set("stream_source_details", nil)
	}

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
