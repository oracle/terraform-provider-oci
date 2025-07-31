// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package apm_traces

import (
	"context"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_apm_traces "github.com/oracle/oci-go-sdk/v65/apmtraces"
	oci_common "github.com/oracle/oci-go-sdk/v65/common"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func ApmTracesLogDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularApmTracesLog,
		Schema: map[string]*schema.Schema{
			"apm_domain_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"log_key": {
				Type:     schema.TypeString,
				Required: true,
			},
			"time_log_ended_less_than": {
				Type:     schema.TypeString,
				Required: true,
			},
			"time_log_started_greater_than_or_equal_to": {
				Type:     schema.TypeString,
				Required: true,
			},
			// Computed
			"attribute_metadata": {
				Type:     schema.TypeMap,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"attributes": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"attribute_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"attribute_value": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"body": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"event_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"overflow_attributes": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"severity_number": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"severity_text": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"span_key": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_created": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_observed": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"timestamp": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"trace_flags": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"trace_key": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func readSingularApmTracesLog(d *schema.ResourceData, m interface{}) error {
	sync := &ApmTracesLogDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).TraceClient()

	return tfresource.ReadResource(sync)
}

type ApmTracesLogDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_apm_traces.TraceClient
	Res    *oci_apm_traces.GetLogResponse
}

func (s *ApmTracesLogDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *ApmTracesLogDataSourceCrud) Get() error {
	request := oci_apm_traces.GetLogRequest{}

	if apmDomainId, ok := s.D.GetOkExists("apm_domain_id"); ok {
		tmp := apmDomainId.(string)
		request.ApmDomainId = &tmp
	}

	if logKey, ok := s.D.GetOkExists("log_key"); ok {
		tmp := logKey.(string)
		request.LogKey = &tmp
	}

	if timeLogEndedLessThan, ok := s.D.GetOkExists("time_log_ended_less_than"); ok {
		tmp, err := time.Parse(time.RFC3339, timeLogEndedLessThan.(string))
		if err != nil {
			return err
		}
		request.TimeLogEndedLessThan = &oci_common.SDKTime{Time: tmp}
	}

	if timeLogStartedGreaterThanOrEqualTo, ok := s.D.GetOkExists("time_log_started_greater_than_or_equal_to"); ok {
		tmp, err := time.Parse(time.RFC3339, timeLogStartedGreaterThanOrEqualTo.(string))
		if err != nil {
			return err
		}
		request.TimeLogStartedGreaterThanOrEqualTo = &oci_common.SDKTime{Time: tmp}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "apm_traces")

	response, err := s.Client.GetLog(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *ApmTracesLogDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("ApmTracesLogDataSource-", ApmTracesLogDataSource(), s.D))

	//s.D.Set("attribute_metadata", s.Res.AttributeMetadata)

	attributes := []interface{}{}
	for _, item := range s.Res.Attributes {
		attributes = append(attributes, AttributeToMap(item))
	}
	s.D.Set("attributes", attributes)

	if s.Res.Body != nil {
		s.D.Set("body", *s.Res.Body)
	}

	if s.Res.EventName != nil {
		s.D.Set("event_name", *s.Res.EventName)
	}

	if s.Res.OverflowAttributes != nil {
		s.D.Set("overflow_attributes", *s.Res.OverflowAttributes)
	}

	if s.Res.SeverityNumber != nil {
		s.D.Set("severity_number", *s.Res.SeverityNumber)
	}

	if s.Res.SeverityText != nil {
		s.D.Set("severity_text", *s.Res.SeverityText)
	}

	if s.Res.SpanKey != nil {
		s.D.Set("span_key", *s.Res.SpanKey)
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeObserved != nil {
		s.D.Set("time_observed", s.Res.TimeObserved.String())
	}

	if s.Res.Timestamp != nil {
		s.D.Set("timestamp", s.Res.Timestamp.String())
	}

	if s.Res.TraceFlags != nil {
		s.D.Set("trace_flags", *s.Res.TraceFlags)
	}

	if s.Res.TraceKey != nil {
		s.D.Set("trace_key", *s.Res.TraceKey)
	}

	return nil
}

func AttributeToMap(obj oci_apm_traces.Attribute) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.AttributeName != nil {
		result["attribute_name"] = string(*obj.AttributeName)
	}

	if obj.AttributeValue != nil {
		result["attribute_value"] = string(*obj.AttributeValue)
	}

	return result
}
