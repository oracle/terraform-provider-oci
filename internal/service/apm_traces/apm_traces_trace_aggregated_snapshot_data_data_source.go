// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package apm_traces

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_apm_traces "github.com/oracle/oci-go-sdk/v65/apmtraces"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func ApmTracesTraceAggregatedSnapshotDataDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularApmTracesTraceAggregatedSnapshotData,
		Schema: map[string]*schema.Schema{
			"apm_domain_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"trace_key": {
				Type:     schema.TypeString,
				Required: true,
			},
			// Computed
			"details": {
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

func readSingularApmTracesTraceAggregatedSnapshotData(d *schema.ResourceData, m interface{}) error {
	sync := &ApmTracesTraceAggregatedSnapshotDataDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).TraceClient()

	return tfresource.ReadResource(sync)
}

type ApmTracesTraceAggregatedSnapshotDataDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_apm_traces.TraceClient
	Res    *oci_apm_traces.GetAggregatedSnapshotResponse
}

func (s *ApmTracesTraceAggregatedSnapshotDataDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *ApmTracesTraceAggregatedSnapshotDataDataSourceCrud) Get() error {
	request := oci_apm_traces.GetAggregatedSnapshotRequest{}

	if apmDomainId, ok := s.D.GetOkExists("apm_domain_id"); ok {
		tmp := apmDomainId.(string)
		request.ApmDomainId = &tmp
	}

	if traceKey, ok := s.D.GetOkExists("trace_key"); ok {
		tmp := traceKey.(string)
		request.TraceKey = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "apm_traces")

	response, err := s.Client.GetAggregatedSnapshot(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *ApmTracesTraceAggregatedSnapshotDataDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("ApmTracesTraceAggregatedSnapshotDataDataSource-", ApmTracesTraceAggregatedSnapshotDataDataSource(), s.D))

	details := []interface{}{}
	for _, item := range s.Res.Details {
		details = append(details, SnapshotDetailToMap(item))
	}
	s.D.Set("details", details)

	return nil
}
