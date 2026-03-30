// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package generative_ai

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_generative_ai "github.com/oracle/oci-go-sdk/v65/generativeai"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func GenerativeAiSemanticStoreDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["semantic_store_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchemaWithContext(GenerativeAiSemanticStoreResource(), fieldMap, readSingularGenerativeAiSemanticStoreWithContext)
}

func readSingularGenerativeAiSemanticStoreWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &GenerativeAiSemanticStoreDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).GenerativeAiClient()

	return tfresource.HandleDiagError(m, tfresource.ReadResourceWithContext(ctx, sync))
}

type GenerativeAiSemanticStoreDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_generative_ai.GenerativeAiClient
	Res    *oci_generative_ai.GetSemanticStoreResponse
}

func (s *GenerativeAiSemanticStoreDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *GenerativeAiSemanticStoreDataSourceCrud) GetWithContext(ctx context.Context) error {
	request := oci_generative_ai.GetSemanticStoreRequest{}

	if semanticStoreId, ok := s.D.GetOkExists("semantic_store_id"); ok {
		tmp := semanticStoreId.(string)
		request.SemanticStoreId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "generative_ai")

	response, err := s.Client.GetSemanticStore(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *GenerativeAiSemanticStoreDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DataSource != nil {
		dataSourceArray := []interface{}{}
		if dataSourceMap := DataSourceDetailsToMap(&s.Res.DataSource); dataSourceMap != nil {
			dataSourceArray = append(dataSourceArray, dataSourceMap)
		}
		s.D.Set("data_source", dataSourceArray)
	} else {
		s.D.Set("data_source", nil)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	} else {
		s.D.Set("defined_tags", nil)
	}

	if s.Res.Description != nil {
		s.D.Set("description", *s.Res.Description)
	} else {
		s.D.Set("description", nil)
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	if s.Res.RefreshSchedule != nil {
		refreshScheduleArray := []interface{}{}
		if refreshScheduleMap := RefreshScheduleDetailsToMap(&s.Res.RefreshSchedule); refreshScheduleMap != nil {
			refreshScheduleArray = append(refreshScheduleArray, refreshScheduleMap)
		}
		s.D.Set("refresh_schedule", refreshScheduleArray)
	} else {
		s.D.Set("refresh_schedule", nil)
	}

	if s.Res.Schemas != nil {
		schemasArray := []interface{}{}
		if schemasMap := SchemasDetailsToMap(&s.Res.Schemas); schemasMap != nil {
			schemasArray = append(schemasArray, schemasMap)
		}
		s.D.Set("schemas", schemasArray)
	} else {
		s.D.Set("schemas", nil)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.SystemTags != nil {
		s.D.Set("system_tags", tfresource.SystemTagsToMap(s.Res.SystemTags))
	} else {
		s.D.Set("system_tags", nil)
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	return nil
}
