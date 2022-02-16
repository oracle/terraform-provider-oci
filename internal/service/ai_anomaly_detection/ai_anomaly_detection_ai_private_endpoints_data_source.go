// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package ai_anomaly_detection

import (
	"context"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_ai_anomaly_detection "github.com/oracle/oci-go-sdk/v58/aianomalydetection"
)

func AiAnomalyDetectionAiPrivateEndpointsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readAiAnomalyDetectionAiPrivateEndpoints,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"ai_private_endpoint_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem:     tfresource.GetDataSourceItemSchema(AiAnomalyDetectionAiPrivateEndpointResource()),
						},
					},
				},
			},
		},
	}
}

func readAiAnomalyDetectionAiPrivateEndpoints(d *schema.ResourceData, m interface{}) error {
	sync := &AiAnomalyDetectionAiPrivateEndpointsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).AnomalyDetectionClient()

	return tfresource.ReadResource(sync)
}

type AiAnomalyDetectionAiPrivateEndpointsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_ai_anomaly_detection.AnomalyDetectionClient
	Res    *oci_ai_anomaly_detection.ListAiPrivateEndpointsResponse
}

func (s *AiAnomalyDetectionAiPrivateEndpointsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *AiAnomalyDetectionAiPrivateEndpointsDataSourceCrud) Get() error {
	request := oci_ai_anomaly_detection.ListAiPrivateEndpointsRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if id, ok := s.D.GetOkExists("id"); ok {
		tmp := id.(string)
		request.Id = &tmp
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_ai_anomaly_detection.AiPrivateEndpointLifecycleStateEnum(state.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "ai_anomaly_detection")

	response, err := s.Client.ListAiPrivateEndpoints(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListAiPrivateEndpoints(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *AiAnomalyDetectionAiPrivateEndpointsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("AiAnomalyDetectionAiPrivateEndpointsDataSource-", AiAnomalyDetectionAiPrivateEndpointsDataSource(), s.D))
	resources := []map[string]interface{}{}
	aiPrivateEndpoint := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, AiPrivateEndpointSummaryToMap(item))
	}
	aiPrivateEndpoint["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, AiAnomalyDetectionAiPrivateEndpointsDataSource().Schema["ai_private_endpoint_collection"].Elem.(*schema.Resource).Schema)
		aiPrivateEndpoint["items"] = items
	}

	resources = append(resources, aiPrivateEndpoint)
	if err := s.D.Set("ai_private_endpoint_collection", resources); err != nil {
		return err
	}

	return nil
}
