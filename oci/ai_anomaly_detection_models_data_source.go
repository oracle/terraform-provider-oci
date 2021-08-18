// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_ai_anomaly_detection "github.com/oracle/oci-go-sdk/v46/aianomalydetection"
)

func init() {
	RegisterDatasource("oci_ai_anomaly_detection_models", AiAnomalyDetectionModelsDataSource())
}

func AiAnomalyDetectionModelsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readAiAnomalyDetectionModels,
		Schema: map[string]*schema.Schema{
			"filter": dataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"display_name": {
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
			"model_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem:     GetDataSourceItemSchema(AiAnomalyDetectionModelResource()),
						},
					},
				},
			},
		},
	}
}

func readAiAnomalyDetectionModels(d *schema.ResourceData, m interface{}) error {
	sync := &AiAnomalyDetectionModelsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).anomalyDetectionClient()

	return ReadResource(sync)
}

type AiAnomalyDetectionModelsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_ai_anomaly_detection.AnomalyDetectionClient
	Res    *oci_ai_anomaly_detection.ListModelsResponse
}

func (s *AiAnomalyDetectionModelsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *AiAnomalyDetectionModelsDataSourceCrud) Get() error {
	request := oci_ai_anomaly_detection.ListModelsRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if projectId, ok := s.D.GetOkExists("project_id"); ok {
		tmp := projectId.(string)
		request.ProjectId = &tmp
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_ai_anomaly_detection.ModelLifecycleStateEnum(state.(string))
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(false, "ai_anomaly_detection")

	response, err := s.Client.ListModels(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListModels(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *AiAnomalyDetectionModelsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(GenerateDataSourceHashID("AiAnomalyDetectionModelsDataSource-", AiAnomalyDetectionModelsDataSource(), s.D))
	resources := []map[string]interface{}{}
	model := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, ModelSummaryToMap(item))
	}
	model["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = ApplyFiltersInCollection(f.(*schema.Set), items, AiAnomalyDetectionModelsDataSource().Schema["model_collection"].Elem.(*schema.Resource).Schema)
		model["items"] = items
	}

	resources = append(resources, model)
	if err := s.D.Set("model_collection", resources); err != nil {
		return err
	}

	return nil
}
