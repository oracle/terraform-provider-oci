// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package apm_synthetics

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_apm_synthetics "github.com/oracle/oci-go-sdk/v58/apmsynthetics"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"
)

func ApmSyntheticsResultDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularApmSyntheticsResult,
		Schema: map[string]*schema.Schema{
			"apm_domain_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"execution_time": {
				Type:     schema.TypeString,
				Required: true,
			},
			"monitor_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"result_content_type": {
				Type:     schema.TypeString,
				Required: true,
			},
			"result_type": {
				Type:     schema.TypeString,
				Required: true,
			},
			"vantage_point": {
				Type:     schema.TypeString,
				Required: true,
			},
			// Computed
			"result_data_set": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"byte_content": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"string_content": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"timestamp": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func readSingularApmSyntheticsResult(d *schema.ResourceData, m interface{}) error {
	sync := &ApmSyntheticsResultDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ApmSyntheticClient()

	return tfresource.ReadResource(sync)
}

type ApmSyntheticsResultDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_apm_synthetics.ApmSyntheticClient
	Res    *oci_apm_synthetics.GetMonitorResultResponse
}

func (s *ApmSyntheticsResultDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *ApmSyntheticsResultDataSourceCrud) Get() error {
	request := oci_apm_synthetics.GetMonitorResultRequest{}

	if apmDomainId, ok := s.D.GetOkExists("apm_domain_id"); ok {
		tmp := apmDomainId.(string)
		request.ApmDomainId = &tmp
	}

	if executionTime, ok := s.D.GetOkExists("execution_time"); ok {
		tmp := executionTime.(string)
		request.ExecutionTime = &tmp
	}

	if monitorId, ok := s.D.GetOkExists("monitor_id"); ok {
		tmp := monitorId.(string)
		request.MonitorId = &tmp
	}

	if resultContentType, ok := s.D.GetOkExists("result_content_type"); ok {
		tmp := resultContentType.(string)
		request.ResultContentType = &tmp
	}

	if resultType, ok := s.D.GetOkExists("result_type"); ok {
		tmp := resultType.(string)
		request.ResultType = &tmp
	}

	if vantagePoint, ok := s.D.GetOkExists("vantage_point"); ok {
		tmp := vantagePoint.(string)
		request.VantagePoint = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "apm_synthetics")

	response, err := s.Client.GetMonitorResult(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *ApmSyntheticsResultDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("ApmSyntheticsResultDataSource-", ApmSyntheticsResultDataSource(), s.D))

	resultDataSet := []interface{}{}
	for _, item := range s.Res.ResultDataSet {
		resultDataSet = append(resultDataSet, MonitorResultDataToMap(item))
	}
	s.D.Set("result_data_set", resultDataSet)

	return nil
}

func MonitorResultDataToMap(obj oci_apm_synthetics.MonitorResultData) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.ByteContent != nil {
		result["byte_content"] = string(obj.ByteContent)
	}

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	if obj.StringContent != nil {
		result["string_content"] = string(*obj.StringContent)
	}

	if obj.Timestamp != nil {
		result["timestamp"] = obj.Timestamp.String()
	}

	return result
}
