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

func StreamingStreamPoolsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readStreamingStreamPools,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
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
			"stream_pools": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     tfresource.GetDataSourceItemSchema(StreamingStreamPoolResource()),
			},
		},
	}
}

func readStreamingStreamPools(d *schema.ResourceData, m interface{}) error {
	sync := &StreamingStreamPoolsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).StreamAdminClient()

	return tfresource.ReadResource(sync)
}

type StreamingStreamPoolsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_streaming.StreamAdminClient
	Res    *oci_streaming.ListStreamPoolsResponse
}

func (s *StreamingStreamPoolsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *StreamingStreamPoolsDataSourceCrud) Get() error {
	request := oci_streaming.ListStreamPoolsRequest{}

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
		request.LifecycleState = oci_streaming.StreamPoolSummaryLifecycleStateEnum(state.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "streaming")

	response, err := s.Client.ListStreamPools(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListStreamPools(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *StreamingStreamPoolsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("StreamingStreamPoolsDataSource-", StreamingStreamPoolsDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		streamPool := map[string]interface{}{
			"compartment_id": *r.CompartmentId,
		}

		if r.DefinedTags != nil {
			streamPool["defined_tags"] = tfresource.DefinedTagsToMap(r.DefinedTags)
		}

		streamPool["freeform_tags"] = r.FreeformTags

		if r.Id != nil {
			streamPool["id"] = *r.Id
		}

		if r.IsPrivate != nil {
			streamPool["is_private"] = *r.IsPrivate
		}

		if r.Name != nil {
			streamPool["name"] = *r.Name
		}

		streamPool["state"] = r.LifecycleState

		if r.TimeCreated != nil {
			streamPool["time_created"] = r.TimeCreated.String()
		}

		resources = append(resources, streamPool)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = tfresource.ApplyFilters(f.(*schema.Set), resources, StreamingStreamPoolsDataSource().Schema["stream_pools"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("stream_pools", resources); err != nil {
		return err
	}

	return nil
}
