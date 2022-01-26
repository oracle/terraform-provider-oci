// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package optimizer

import (
	"context"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_optimizer "github.com/oracle/oci-go-sdk/v56/optimizer"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"
)

func OptimizerResourceActionDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["resource_action_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(OptimizerResourceActionResource(), fieldMap, readSingularOptimizerResourceAction)
}

func readSingularOptimizerResourceAction(d *schema.ResourceData, m interface{}) error {
	sync := &OptimizerResourceActionDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).OptimizerClient()

	return tfresource.ReadResource(sync)
}

type OptimizerResourceActionDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_optimizer.OptimizerClient
	Res    *oci_optimizer.GetResourceActionResponse
}

func (s *OptimizerResourceActionDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *OptimizerResourceActionDataSourceCrud) Get() error {
	request := oci_optimizer.GetResourceActionRequest{}

	if resourceActionId, ok := s.D.GetOkExists("resource_action_id"); ok {
		tmp := resourceActionId.(string)
		request.ResourceActionId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "optimizer")

	response, err := s.Client.GetResourceAction(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *OptimizerResourceActionDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.Action != nil {
		s.D.Set("action", []interface{}{OptimizerActionToMap(s.Res.Action)})
	} else {
		s.D.Set("action", nil)
	}

	if s.Res.CategoryId != nil {
		s.D.Set("category_id", *s.Res.CategoryId)
	}

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.CompartmentName != nil {
		s.D.Set("compartment_name", *s.Res.CompartmentName)
	}

	if s.Res.EstimatedCostSaving != nil {
		s.D.Set("estimated_cost_saving", *s.Res.EstimatedCostSaving)
	}

	s.D.Set("extended_metadata", s.Res.ExtendedMetadata)

	s.D.Set("metadata", s.Res.Metadata)

	if s.Res.Name != nil {
		s.D.Set("name", *s.Res.Name)
	}

	if s.Res.RecommendationId != nil {
		s.D.Set("recommendation_id", *s.Res.RecommendationId)
	}

	if s.Res.ResourceId != nil {
		s.D.Set("resource_id", *s.Res.ResourceId)
	}

	if s.Res.ResourceType != nil {
		s.D.Set("resource_type", *s.Res.ResourceType)
	}

	s.D.Set("state", s.Res.LifecycleState)

	s.D.Set("status", s.Res.Status)

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
