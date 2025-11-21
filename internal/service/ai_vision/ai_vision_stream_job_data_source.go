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

func AiVisionStreamJobDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["stream_job_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(AiVisionStreamJobResource(), fieldMap, readSingularAiVisionStreamJob)
}

func readSingularAiVisionStreamJob(d *schema.ResourceData, m interface{}) error {
	sync := &AiVisionStreamJobDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).AiServiceVisionClient()

	return tfresource.ReadResource(sync)
}

type AiVisionStreamJobDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_ai_vision.AIServiceVisionClient
	Res    *oci_ai_vision.GetStreamJobResponse
}

func (s *AiVisionStreamJobDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *AiVisionStreamJobDataSourceCrud) Get() error {
	request := oci_ai_vision.GetStreamJobRequest{}

	if streamJobId, ok := s.D.GetOkExists("stream_job_id"); ok {
		tmp := streamJobId.(string)
		request.StreamJobId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "ai_vision")

	response, err := s.Client.GetStreamJob(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *AiVisionStreamJobDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.AgentParticipantId != nil {
		s.D.Set("agent_participant_id", *s.Res.AgentParticipantId)
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

	features := []interface{}{}
	for _, item := range s.Res.Features {
		features = append(features, VideoStreamFeatureToMap(item))
	}
	s.D.Set("features", features)

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.StreamOutputLocation != nil {
		streamOutputLocationArray := []interface{}{}
		if streamOutputLocationMap := StreamOutputLocationToMap(&s.Res.StreamOutputLocation); streamOutputLocationMap != nil {
			streamOutputLocationArray = append(streamOutputLocationArray, streamOutputLocationMap)
		}
		s.D.Set("stream_output_location", streamOutputLocationArray)
	} else {
		s.D.Set("stream_output_location", nil)
	}

	if s.Res.StreamSourceId != nil {
		s.D.Set("stream_source_id", *s.Res.StreamSourceId)
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
