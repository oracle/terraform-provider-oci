// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package ai_anomaly_detection

import (
	"context"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_ai_anomaly_detection "github.com/oracle/oci-go-sdk/v56/aianomalydetection"
)

func AiAnomalyDetectionDataAssetsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readAiAnomalyDetectionDataAssets,
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
			"project_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"data_asset_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem:     tfresource.GetDataSourceItemSchema(AiAnomalyDetectionDataAssetResource()),
						},
					},
				},
			},
		},
	}
}

func readAiAnomalyDetectionDataAssets(d *schema.ResourceData, m interface{}) error {
	sync := &AiAnomalyDetectionDataAssetsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).AnomalyDetectionClient()

	return tfresource.ReadResource(sync)
}

type AiAnomalyDetectionDataAssetsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_ai_anomaly_detection.AnomalyDetectionClient
	Res    *oci_ai_anomaly_detection.ListDataAssetsResponse
}

func (s *AiAnomalyDetectionDataAssetsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *AiAnomalyDetectionDataAssetsDataSourceCrud) Get() error {
	request := oci_ai_anomaly_detection.ListDataAssetsRequest{}

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
		request.LifecycleState = oci_ai_anomaly_detection.DataAssetLifecycleStateEnum(state.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "ai_anomaly_detection")

	response, err := s.Client.ListDataAssets(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListDataAssets(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *AiAnomalyDetectionDataAssetsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("AiAnomalyDetectionDataAssetsDataSource-", AiAnomalyDetectionDataAssetsDataSource(), s.D))
	resources := []map[string]interface{}{}
	dataAsset := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, AiDataAssetSummaryToMap(item))
	}
	dataAsset["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, AiAnomalyDetectionDataAssetsDataSource().Schema["data_asset_collection"].Elem.(*schema.Resource).Schema)
		dataAsset["items"] = items
	}

	resources = append(resources, dataAsset)
	if err := s.D.Set("data_asset_collection", resources); err != nil {
		return err
	}

	return nil
}
