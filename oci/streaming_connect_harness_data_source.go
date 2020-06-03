// Copyright (c) 2017, 2020, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"

	"github.com/hashicorp/terraform/helper/schema"
	oci_streaming "github.com/oracle/oci-go-sdk/streaming"
)

func init() {
	RegisterDatasource("oci_streaming_connect_harness", StreamingConnectHarnessDataSource())
}

func StreamingConnectHarnessDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["connect_harness_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return GetSingularDataSourceItemSchema(StreamingConnectHarnessResource(), fieldMap, readSingularStreamingConnectHarness)
}

func readSingularStreamingConnectHarness(d *schema.ResourceData, m interface{}) error {
	sync := &StreamingConnectHarnessDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).streamAdminClient()

	return ReadResource(sync)
}

type StreamingConnectHarnessDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_streaming.StreamAdminClient
	Res    *oci_streaming.GetConnectHarnessResponse
}

func (s *StreamingConnectHarnessDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *StreamingConnectHarnessDataSourceCrud) Get() error {
	request := oci_streaming.GetConnectHarnessRequest{}

	if connectHarnessId, ok := s.D.GetOkExists("connect_harness_id"); ok {
		tmp := connectHarnessId.(string)
		request.ConnectHarnessId = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(false, "streaming")

	response, err := s.Client.GetConnectHarness(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *StreamingConnectHarnessDataSourceCrud) SetData() error {
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

	if s.Res.Name != nil {
		s.D.Set("name", *s.Res.Name)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	return nil
}
