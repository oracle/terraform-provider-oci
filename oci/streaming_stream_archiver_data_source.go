// Copyright (c) 2017, 2019, Oracle and/or its affiliates. All rights reserved.

package oci

import (
	"context"

	"github.com/hashicorp/terraform/helper/schema"
	oci_streaming "github.com/oracle/oci-go-sdk/streaming"
)

func init() {
	RegisterDatasource("oci_streaming_stream_archiver", StreamingStreamArchiverDataSource())
}

func StreamingStreamArchiverDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["stream_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return GetSingularDataSourceItemSchema(StreamingStreamArchiverResource(), fieldMap, readSingularStreamingStreamArchiver)
}

func readSingularStreamingStreamArchiver(d *schema.ResourceData, m interface{}) error {
	sync := &StreamingStreamArchiverDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).streamAdminClient

	return ReadResource(sync)
}

type StreamingStreamArchiverDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_streaming.StreamAdminClient
	Res    *oci_streaming.GetArchiverResponse
}

func (s *StreamingStreamArchiverDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *StreamingStreamArchiverDataSourceCrud) Get() error {
	request := oci_streaming.GetArchiverRequest{}

	if streamId, ok := s.D.GetOkExists("stream_id"); ok {
		tmp := streamId.(string)
		request.StreamId = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(false, "streaming")

	response, err := s.Client.GetArchiver(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *StreamingStreamArchiverDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(GenerateDataSourceID())

	if s.Res.BatchRolloverSizeInMBs != nil {
		s.D.Set("batch_rollover_size_in_mbs", *s.Res.BatchRolloverSizeInMBs)
	}

	if s.Res.BatchRolloverTimeInSeconds != nil {
		s.D.Set("batch_rollover_time_in_seconds", *s.Res.BatchRolloverTimeInSeconds)
	}

	if s.Res.BucketName != nil {
		s.D.Set("bucket", *s.Res.BucketName)
	}

	if s.Res.Error != nil {
		s.D.Set("error", []interface{}{ArchiverErrorToMap(s.Res.Error)})
	} else {
		s.D.Set("error", nil)
	}

	s.D.Set("start_position", s.Res.StartPosition)

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.UseExistingBucket != nil {
		s.D.Set("use_existing_bucket", *s.Res.UseExistingBucket)
	}

	return nil
}
