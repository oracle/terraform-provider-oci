// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package apm_traces

import (
	"context"
	"encoding/json"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_apm_traces "github.com/oracle/oci-go-sdk/v65/apmtraces"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func ApmTracesTraceSnapshotDataDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularApmTracesTraceSnapshotData,
		Schema: map[string]*schema.Schema{
			"apm_domain_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"is_summarized": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"snapshot_time": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"thread_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"trace_key": {
				Type:     schema.TypeString,
				Required: true,
			},
			// Computed
			"key": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_ended": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_started": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"trace_snapshot_details": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"key": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"value": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func readSingularApmTracesTraceSnapshotData(d *schema.ResourceData, m interface{}) error {
	sync := &ApmTracesTraceSnapshotDataDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).TraceClient()

	return tfresource.ReadResource(sync)
}

type ApmTracesTraceSnapshotDataDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_apm_traces.TraceClient
	Res    *oci_apm_traces.GetTraceSnapshotResponse
}

func (s *ApmTracesTraceSnapshotDataDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *ApmTracesTraceSnapshotDataDataSourceCrud) Get() error {
	request := oci_apm_traces.GetTraceSnapshotRequest{}

	if apmDomainId, ok := s.D.GetOkExists("apm_domain_id"); ok {
		tmp := apmDomainId.(string)
		request.ApmDomainId = &tmp
	}

	if isSummarized, ok := s.D.GetOkExists("is_summarized"); ok {
		tmp := isSummarized.(bool)
		request.IsSummarized = &tmp
	}

	if snapshotTime, ok := s.D.GetOkExists("snapshot_time"); ok {
		tmp := snapshotTime.(string)
		request.SnapshotTime = &tmp
	}

	if threadId, ok := s.D.GetOkExists("thread_id"); ok {
		tmp := threadId.(string)
		request.ThreadId = &tmp
	}

	if traceKey, ok := s.D.GetOkExists("trace_key"); ok {
		tmp := traceKey.(string)
		request.TraceKey = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "apm_traces")

	response, err := s.Client.GetTraceSnapshot(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *ApmTracesTraceSnapshotDataDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("ApmTracesTraceSnapshotDataDataSource-", ApmTracesTraceSnapshotDataDataSource(), s.D))

	if s.Res.Key != nil {
		s.D.Set("key", *s.Res.Key)
	}

	if s.Res.TimeEnded != nil {
		s.D.Set("time_ended", s.Res.TimeEnded.String())
	}

	if s.Res.TimeStarted != nil {
		s.D.Set("time_started", s.Res.TimeStarted.String())
	}

	traceSnapshotDetails := []interface{}{}
	for _, item := range s.Res.TraceSnapshotDetails {
		traceSnapshotDetails = append(traceSnapshotDetails, SnapshotDetailToMap(item))
	}
	s.D.Set("trace_snapshot_details", traceSnapshotDetails)

	return nil
}

func SnapshotDetailToMap(obj oci_apm_traces.SnapshotDetail) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Key != nil {
		result["key"] = string(*obj.Key)
	}

	if obj.Value != nil {
		tmp, _ := json.Marshal(obj.Value)
		result["value"] = string(tmp)
	}

	return result
}
