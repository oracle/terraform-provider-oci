// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package optimizer

import (
	"context"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_optimizer "github.com/oracle/oci-go-sdk/v65/optimizer"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func OptimizerRecommendationDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["recommendation_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(OptimizerRecommendationResource(), fieldMap, readSingularOptimizerRecommendation)
}

func readSingularOptimizerRecommendation(d *schema.ResourceData, m interface{}) error {
	sync := &OptimizerRecommendationDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).OptimizerClient()

	return tfresource.ReadResource(sync)
}

type OptimizerRecommendationDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_optimizer.OptimizerClient
	Res    *oci_optimizer.GetRecommendationResponse
}

func (s *OptimizerRecommendationDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *OptimizerRecommendationDataSourceCrud) Get() error {
	request := oci_optimizer.GetRecommendationRequest{}

	if recommendationId, ok := s.D.GetOkExists("recommendation_id"); ok {
		tmp := recommendationId.(string)
		request.RecommendationId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "optimizer")

	response, err := s.Client.GetRecommendation(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *OptimizerRecommendationDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.CategoryId != nil {
		s.D.Set("category_id", *s.Res.CategoryId)
	}

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.Description != nil {
		s.D.Set("description", *s.Res.Description)
	}

	if s.Res.EstimatedCostSaving != nil {
		s.D.Set("estimated_cost_saving", *s.Res.EstimatedCostSaving)
	}

	s.D.Set("extended_metadata", s.Res.ExtendedMetadata)

	s.D.Set("importance", s.Res.Importance)

	if s.Res.Name != nil {
		s.D.Set("name", *s.Res.Name)
	}

	resourceCounts := []interface{}{}
	for _, item := range s.Res.ResourceCounts {
		resourceCounts = append(resourceCounts, ResourceCountToMap(item))
	}
	s.D.Set("resource_counts", resourceCounts)

	s.D.Set("state", s.Res.LifecycleState)

	s.D.Set("status", s.Res.Status)

	if s.Res.SupportedLevels != nil {
		s.D.Set("supported_levels", []interface{}{SupportedLevelsToMap(s.Res.SupportedLevels)})
	} else {
		s.D.Set("supported_levels", nil)
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeStatusBegin != nil {
		s.D.Set("time_status_begin", s.Res.TimeStatusBegin.String())
	}

	if s.Res.TimeStatusEnd != nil {
		s.D.Set("time_status_end", s.Res.TimeStatusEnd.Format(time.RFC3339Nano))
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	return nil
}
