// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package streaming

import (
	"context"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_streaming "github.com/oracle/oci-go-sdk/v56/streaming"
)

func StreamingStreamsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readStreamingStreams,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"stream_pool_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"streams": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     tfresource.GetDataSourceItemSchema(StreamingStreamResource()),
			},
		},
	}
}

func readStreamingStreams(d *schema.ResourceData, m interface{}) error {
	sync := &StreamingStreamsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).StreamAdminClient()

	return tfresource.ReadResource(sync)
}

type StreamingStreamsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_streaming.StreamAdminClient
	Res    *oci_streaming.ListStreamsResponse
}

func (s *StreamingStreamsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *StreamingStreamsDataSourceCrud) Get() error {
	request := oci_streaming.ListStreamsRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if id, ok := s.D.GetOkExists("id"); ok {
		tmp := id.(string)
		request.Id = &tmp
	}

	if name, ok := s.D.GetOkExists("name"); ok {
		tmp := name.(string)
		request.Name = &tmp
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_streaming.StreamLifecycleStateEnum(state.(string))
	}

	if streamPoolId, ok := s.D.GetOkExists("stream_pool_id"); ok {
		tmp := streamPoolId.(string)
		request.StreamPoolId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "streaming")

	response, err := s.Client.ListStreams(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListStreams(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *StreamingStreamsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("StreamingStreamsDataSource-", StreamingStreamsDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		stream := map[string]interface{}{}

		if r.CompartmentId != nil {
			stream["compartment_id"] = *r.CompartmentId
		}

		if r.DefinedTags != nil {
			stream["defined_tags"] = tfresource.DefinedTagsToMap(r.DefinedTags)
		}

		stream["freeform_tags"] = r.FreeformTags

		if r.Id != nil {
			stream["id"] = *r.Id
		}

		if r.MessagesEndpoint != nil {
			stream["messages_endpoint"] = *r.MessagesEndpoint
		}

		if r.Name != nil {
			stream["name"] = *r.Name
		}

		if r.Partitions != nil {
			stream["partitions"] = *r.Partitions
		}

		stream["state"] = r.LifecycleState

		if r.StreamPoolId != nil {
			stream["stream_pool_id"] = *r.StreamPoolId
		}

		if r.TimeCreated != nil {
			stream["time_created"] = r.TimeCreated.String()
		}

		resources = append(resources, stream)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = tfresource.ApplyFilters(f.(*schema.Set), resources, StreamingStreamsDataSource().Schema["streams"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("streams", resources); err != nil {
		return err
	}

	return nil
}
