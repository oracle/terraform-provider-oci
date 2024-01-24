// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package ai_anomaly_detection

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_ai_anomaly_detection "github.com/oracle/oci-go-sdk/v65/aianomalydetection"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func AiAnomalyDetectionDetectAnomalyJobsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readAiAnomalyDetectionDetectAnomalyJobs,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"detect_anomaly_job_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"model_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"project_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"detect_anomaly_job_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem:     tfresource.GetDataSourceItemSchema(AiAnomalyDetectionDetectAnomalyJobResource()),
						},
					},
				},
			},
		},
	}
}

func readAiAnomalyDetectionDetectAnomalyJobs(d *schema.ResourceData, m interface{}) error {
	sync := &AiAnomalyDetectionDetectAnomalyJobsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).AnomalyDetectionClient()

	return tfresource.ReadResource(sync)
}

type AiAnomalyDetectionDetectAnomalyJobsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_ai_anomaly_detection.AnomalyDetectionClient
	Res    *oci_ai_anomaly_detection.ListDetectAnomalyJobsResponse
}

func (s *AiAnomalyDetectionDetectAnomalyJobsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *AiAnomalyDetectionDetectAnomalyJobsDataSourceCrud) Get() error {
	request := oci_ai_anomaly_detection.ListDetectAnomalyJobsRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if detectAnomalyJobId, ok := s.D.GetOkExists("id"); ok {
		tmp := detectAnomalyJobId.(string)
		request.DetectAnomalyJobId = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if modelId, ok := s.D.GetOkExists("model_id"); ok {
		tmp := modelId.(string)
		request.ModelId = &tmp
	}

	if projectId, ok := s.D.GetOkExists("project_id"); ok {
		tmp := projectId.(string)
		request.ProjectId = &tmp
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_ai_anomaly_detection.DetectAnomalyJobLifecycleStateEnum(state.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "ai_anomaly_detection")

	response, err := s.Client.ListDetectAnomalyJobs(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListDetectAnomalyJobs(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *AiAnomalyDetectionDetectAnomalyJobsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("AiAnomalyDetectionDetectAnomalyJobsDataSource-", AiAnomalyDetectionDetectAnomalyJobsDataSource(), s.D))
	resources := []map[string]interface{}{}
	detectAnomalyJob := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, DetectAnomalyJobSummaryToMap(item))
	}
	detectAnomalyJob["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, AiAnomalyDetectionDetectAnomalyJobsDataSource().Schema["detect_anomaly_job_collection"].Elem.(*schema.Resource).Schema)
		detectAnomalyJob["items"] = items
	}

	resources = append(resources, detectAnomalyJob)
	if err := s.D.Set("detect_anomaly_job_collection", resources); err != nil {
		return err
	}

	return nil
}
