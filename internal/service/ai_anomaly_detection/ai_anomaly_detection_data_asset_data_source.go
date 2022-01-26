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

func AiAnomalyDetectionDataAssetDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["data_asset_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(AiAnomalyDetectionDataAssetResource(), fieldMap, readSingularAiAnomalyDetectionDataAsset)
}

func readSingularAiAnomalyDetectionDataAsset(d *schema.ResourceData, m interface{}) error {
	sync := &AiAnomalyDetectionDataAssetDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).AnomalyDetectionClient()

	return tfresource.ReadResource(sync)
}

type AiAnomalyDetectionDataAssetDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_ai_anomaly_detection.AnomalyDetectionClient
	Res    *oci_ai_anomaly_detection.GetDataAssetResponse
}

func (s *AiAnomalyDetectionDataAssetDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *AiAnomalyDetectionDataAssetDataSourceCrud) Get() error {
	request := oci_ai_anomaly_detection.GetDataAssetRequest{}

	if dataAssetId, ok := s.D.GetOkExists("data_asset_id"); ok {
		tmp := dataAssetId.(string)
		request.DataAssetId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "ai_anomaly_detection")

	response, err := s.Client.GetDataAsset(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *AiAnomalyDetectionDataAssetDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DataSourceDetails != nil {
		dataSourceDetailsArray := []interface{}{}
		if dataSourceDetailsMap := DataSourceDetailsToMap(&s.Res.DataSourceDetails); dataSourceDetailsMap != nil {
			dataSourceDetailsArray = append(dataSourceDetailsArray, dataSourceDetailsMap)
		}
		s.D.Set("data_source_details", dataSourceDetailsArray)
	} else {
		s.D.Set("data_source_details", nil)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.Description != nil {
		s.D.Set("description", *s.Res.Description)
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.PrivateEndpointId != nil {
		s.D.Set("private_endpoint_id", *s.Res.PrivateEndpointId)
	}

	if s.Res.ProjectId != nil {
		s.D.Set("project_id", *s.Res.ProjectId)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.SystemTags != nil {
		s.D.Set("system_tags", tfresource.SystemTagsToMap(s.Res.SystemTags))
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	return nil
}
