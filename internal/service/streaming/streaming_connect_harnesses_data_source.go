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

func StreamingConnectHarnessesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readStreamingConnectHarnesses,
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
			"connect_harness": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     tfresource.GetDataSourceItemSchema(StreamingConnectHarnessResource()),
			},
		},
	}
}

func readStreamingConnectHarnesses(d *schema.ResourceData, m interface{}) error {
	sync := &StreamingConnectHarnessesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).StreamAdminClient()

	return tfresource.ReadResource(sync)
}

type StreamingConnectHarnessesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_streaming.StreamAdminClient
	Res    *oci_streaming.ListConnectHarnessesResponse
}

func (s *StreamingConnectHarnessesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *StreamingConnectHarnessesDataSourceCrud) Get() error {
	request := oci_streaming.ListConnectHarnessesRequest{}

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
		request.LifecycleState = oci_streaming.ConnectHarnessSummaryLifecycleStateEnum(state.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "streaming")

	response, err := s.Client.ListConnectHarnesses(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListConnectHarnesses(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *StreamingConnectHarnessesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("StreamingConnectHarnessesDataSource-", StreamingConnectHarnessesDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		connectHarness := map[string]interface{}{
			"compartment_id": *r.CompartmentId,
		}

		if r.DefinedTags != nil {
			connectHarness["defined_tags"] = tfresource.DefinedTagsToMap(r.DefinedTags)
		}

		connectHarness["freeform_tags"] = r.FreeformTags

		if r.Id != nil {
			connectHarness["id"] = *r.Id
		}

		if r.Name != nil {
			connectHarness["name"] = *r.Name
		}

		connectHarness["state"] = r.LifecycleState

		if r.TimeCreated != nil {
			connectHarness["time_created"] = r.TimeCreated.String()
		}

		resources = append(resources, connectHarness)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = tfresource.ApplyFilters(f.(*schema.Set), resources, StreamingConnectHarnessesDataSource().Schema["connect_harness"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("connect_harness", resources); err != nil {
		return err
	}

	return nil
}
