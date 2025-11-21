// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package jms

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_jms "github.com/oracle/oci-go-sdk/v65/jms"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func JmsFleetCryptoAnalysisResultDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularJmsFleetCryptoAnalysisResult,
		Schema: map[string]*schema.Schema{
			"crypto_analysis_result_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"fleet_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			// Computed
			"aggregation_mode": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"bucket": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"crypto_roadmap_version": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"finding_count": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"host_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"managed_instance_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"namespace": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"non_compliant_finding_count": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"object": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"summarized_event_count": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"time_created": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_finished": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_first_event": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_last_event": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_started": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"total_event_count": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"work_request_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func readSingularJmsFleetCryptoAnalysisResult(d *schema.ResourceData, m interface{}) error {
	sync := &JmsFleetCryptoAnalysisResultDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).JavaManagementServiceClient()

	return tfresource.ReadResource(sync)
}

type JmsFleetCryptoAnalysisResultDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_jms.JavaManagementServiceClient
	Res    *oci_jms.GetCryptoAnalysisResultResponse
}

func (s *JmsFleetCryptoAnalysisResultDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *JmsFleetCryptoAnalysisResultDataSourceCrud) Get() error {
	request := oci_jms.GetCryptoAnalysisResultRequest{}

	if cryptoAnalysisResultId, ok := s.D.GetOkExists("crypto_analysis_result_id"); ok {
		tmp := cryptoAnalysisResultId.(string)
		request.CryptoAnalysisResultId = &tmp
	}

	if fleetId, ok := s.D.GetOkExists("fleet_id"); ok {
		tmp := fleetId.(string)
		request.FleetId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "jms")

	response, err := s.Client.GetCryptoAnalysisResult(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *JmsFleetCryptoAnalysisResultDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	s.D.Set("aggregation_mode", s.Res.AggregationMode)

	if s.Res.BucketName != nil {
		s.D.Set("bucket", *s.Res.BucketName)
	}

	if s.Res.CryptoRoadmapVersion != nil {
		s.D.Set("crypto_roadmap_version", *s.Res.CryptoRoadmapVersion)
	}

	if s.Res.FindingCount != nil {
		s.D.Set("finding_count", *s.Res.FindingCount)
	}

	if s.Res.HostName != nil {
		s.D.Set("host_name", *s.Res.HostName)
	}

	if s.Res.ManagedInstanceId != nil {
		s.D.Set("managed_instance_id", *s.Res.ManagedInstanceId)
	}

	if s.Res.Namespace != nil {
		s.D.Set("namespace", *s.Res.Namespace)
	}

	if s.Res.NonCompliantFindingCount != nil {
		s.D.Set("non_compliant_finding_count", *s.Res.NonCompliantFindingCount)
	}

	if s.Res.ObjectName != nil {
		s.D.Set("object", *s.Res.ObjectName)
	}

	if s.Res.SummarizedEventCount != nil {
		s.D.Set("summarized_event_count", *s.Res.SummarizedEventCount)
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeFinished != nil {
		s.D.Set("time_finished", s.Res.TimeFinished.String())
	}

	if s.Res.TimeFirstEvent != nil {
		s.D.Set("time_first_event", s.Res.TimeFirstEvent.String())
	}

	if s.Res.TimeLastEvent != nil {
		s.D.Set("time_last_event", s.Res.TimeLastEvent.String())
	}

	if s.Res.TimeStarted != nil {
		s.D.Set("time_started", s.Res.TimeStarted.String())
	}

	if s.Res.TotalEventCount != nil {
		s.D.Set("total_event_count", *s.Res.TotalEventCount)
	}

	if s.Res.WorkRequestId != nil {
		s.D.Set("work_request_id", *s.Res.WorkRequestId)
	}

	return nil
}
