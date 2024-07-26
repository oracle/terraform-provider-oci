// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package generative_ai

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_generative_ai "github.com/oracle/oci-go-sdk/v65/generativeai"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func GenerativeAiModelDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["model_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(GenerativeAiModelResource(), fieldMap, readSingularGenerativeAiModel)
}

func readSingularGenerativeAiModel(d *schema.ResourceData, m interface{}) error {
	sync := &GenerativeAiModelDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).GenerativeAiClient()

	return tfresource.ReadResource(sync)
}

type GenerativeAiModelDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_generative_ai.GenerativeAiClient
	Res    *oci_generative_ai.GetModelResponse
}

func (s *GenerativeAiModelDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *GenerativeAiModelDataSourceCrud) Get() error {
	request := oci_generative_ai.GetModelRequest{}

	if modelId, ok := s.D.GetOkExists("model_id"); ok {
		tmp := modelId.(string)
		request.ModelId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "generative_ai")

	response, err := s.Client.GetModel(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *GenerativeAiModelDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.BaseModelId != nil {
		s.D.Set("base_model_id", *s.Res.BaseModelId)
	}

	s.D.Set("capabilities", s.Res.Capabilities)

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
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

	if s.Res.FineTuneDetails != nil {
		s.D.Set("fine_tune_details", []interface{}{FineTuneDetailsToMap(s.Res.FineTuneDetails)})
	} else {
		s.D.Set("fine_tune_details", nil)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.IsLongTermSupported != nil {
		s.D.Set("is_long_term_supported", *s.Res.IsLongTermSupported)
	}

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	if s.Res.ModelMetrics != nil {
		modelMetricsArray := []interface{}{}
		if modelMetricsMap := ModelMetricsToMap(&s.Res.ModelMetrics); modelMetricsMap != nil {
			modelMetricsArray = append(modelMetricsArray, modelMetricsMap)
		}
		s.D.Set("model_metrics", modelMetricsArray)
	} else {
		s.D.Set("model_metrics", nil)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.SystemTags != nil {
		s.D.Set("system_tags", tfresource.SystemTagsToMap(s.Res.SystemTags))
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeDeprecated != nil {
		s.D.Set("time_deprecated", s.Res.TimeDeprecated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	s.D.Set("type", s.Res.Type)

	if s.Res.Vendor != nil {
		s.D.Set("vendor", *s.Res.Vendor)
	}

	if s.Res.Version != nil {
		s.D.Set("version", *s.Res.Version)
	}

	return nil
}
