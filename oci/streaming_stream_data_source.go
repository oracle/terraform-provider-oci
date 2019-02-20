// Copyright (c) 2017, 2019, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"context"

	"github.com/hashicorp/terraform/helper/schema"
	oci_streaming "github.com/oracle/oci-go-sdk/streaming"
)

func StreamingStreamDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularStreamingStream,
		Schema: map[string]*schema.Schema{
			"stream_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			// Computed
			"compartment_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"defined_tags": {
				Type:     schema.TypeMap,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"freeform_tags": {
				Type:     schema.TypeMap,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"lifecycle_state_details": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"messages_endpoint": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"partitions": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"retention_in_hours": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_created": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func readSingularStreamingStream(d *schema.ResourceData, m interface{}) error {
	sync := &StreamingStreamDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).streamAdminClient

	return ReadResource(sync)
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

	request.RequestMetadata.RetryPolicy = getRetryPolicy(false, "streaming")

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
		s.D.Set("defined_tags", definedTagsToMap(s.Res.DefinedTags))
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

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	return nil
}
