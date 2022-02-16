// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package streaming

import (
	"context"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_streaming "github.com/oracle/oci-go-sdk/v58/streaming"
)

func StreamingStreamDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["stream_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(StreamingStreamResource(), fieldMap, readSingularStreamingStream)
}

func readSingularStreamingStream(d *schema.ResourceData, m interface{}) error {
	sync := &StreamingStreamDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).StreamAdminClient()

	return tfresource.ReadResource(sync)
}

type StreamingStreamDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_streaming.StreamAdminClient
	Res    *oci_streaming.GetStreamResponse
}

func (s *StreamingStreamDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *StreamingStreamDataSourceCrud) Get() error {
	request := oci_streaming.GetStreamRequest{}

	if streamId, ok := s.D.GetOkExists("stream_id"); ok {
		tmp := streamId.(string)
		request.StreamId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "streaming")

	response, err := s.Client.GetStream(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *StreamingStreamDataSourceCrud) SetData() error {
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

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.LifecycleStateDetails != nil {
		s.D.Set("lifecycle_state_details", *s.Res.LifecycleStateDetails)
	}

	if s.Res.MessagesEndpoint != nil {
		s.D.Set("messages_endpoint", *s.Res.MessagesEndpoint)
	}

	if s.Res.Name != nil {
		s.D.Set("name", *s.Res.Name)
	}

	if s.Res.Partitions != nil {
		s.D.Set("partitions", *s.Res.Partitions)
	}

	if s.Res.RetentionInHours != nil {
		s.D.Set("retention_in_hours", *s.Res.RetentionInHours)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.StreamPoolId != nil {
		s.D.Set("stream_pool_id", *s.Res.StreamPoolId)
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	return nil
}
