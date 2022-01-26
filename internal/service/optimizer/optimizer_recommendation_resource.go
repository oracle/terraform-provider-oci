// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package optimizer

import (
	"context"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	oci_common "github.com/oracle/oci-go-sdk/v56/common"
	oci_optimizer "github.com/oracle/oci-go-sdk/v56/optimizer"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"
	"github.com/terraform-providers/terraform-provider-oci/internal/utils"
)

func OptimizerRecommendationResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createOptimizerRecommendation,
		Read:     readOptimizerRecommendation,
		Update:   updateOptimizerRecommendation,
		Delete:   deleteOptimizerRecommendation,
		Schema: map[string]*schema.Schema{
			// Required
			"recommendation_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"status": {
				Type:     schema.TypeString,
				Required: true,
			},

			// Optional
			"time_status_end": {
				Type:             schema.TypeString,
				Optional:         true,
				Computed:         true,
				DiffSuppressFunc: utils.TimeDiffSuppressFunction,
			},

			// Computed
			"category_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"description": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"estimated_cost_saving": {
				Type:     schema.TypeFloat,
				Computed: true,
			},
			"extended_metadata": {
				Type:     schema.TypeMap,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"importance": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"resource_counts": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"count": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"status": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"supported_levels": {
				Type:     schema.TypeList,
				Computed: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"name": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
					},
				},
			},
			"time_created": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_status_begin": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_updated": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func createOptimizerRecommendation(d *schema.ResourceData, m interface{}) error {
	sync := &OptimizerRecommendationResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).OptimizerClient()

	return tfresource.CreateResource(d, sync)
}

func readOptimizerRecommendation(d *schema.ResourceData, m interface{}) error {
	sync := &OptimizerRecommendationResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).OptimizerClient()

	return tfresource.ReadResource(sync)
}

func updateOptimizerRecommendation(d *schema.ResourceData, m interface{}) error {
	sync := &OptimizerRecommendationResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).OptimizerClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteOptimizerRecommendation(d *schema.ResourceData, m interface{}) error {
	return nil
}

type OptimizerRecommendationResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_optimizer.OptimizerClient
	Res                    *oci_optimizer.Recommendation
	DisableNotFoundRetries bool
}

func (s *OptimizerRecommendationResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *OptimizerRecommendationResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_optimizer.LifecycleStateAttaching),
		string(oci_optimizer.LifecycleStateCreating),
	}
}

func (s *OptimizerRecommendationResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_optimizer.LifecycleStateActive),
	}
}

func (s *OptimizerRecommendationResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_optimizer.LifecycleStateDetaching),
		string(oci_optimizer.LifecycleStateDeleting),
	}
}

func (s *OptimizerRecommendationResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_optimizer.LifecycleStateDeleted),
	}
}

func (s *OptimizerRecommendationResourceCrud) Create() error {
	request := oci_optimizer.UpdateRecommendationRequest{}

	if recommendationId, ok := s.D.GetOkExists("recommendation_id"); ok {
		tmp := recommendationId.(string)
		request.RecommendationId = &tmp
	}

	if status, ok := s.D.GetOkExists("status"); ok {
		request.Status = oci_optimizer.StatusEnum(status.(string))
	}

	if timeStatusEnd, ok := s.D.GetOkExists("time_status_end"); ok {
		tmp, err := time.Parse(time.RFC3339, timeStatusEnd.(string))
		if err != nil {
			return err
		}
		request.TimeStatusEnd = &oci_common.SDKTime{Time: tmp}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "optimizer")

	response, err := s.Client.UpdateRecommendation(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Recommendation
	return nil
}

func (s *OptimizerRecommendationResourceCrud) Get() error {
	request := oci_optimizer.GetRecommendationRequest{}

	tmp := s.D.Id()
	request.RecommendationId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "optimizer")

	response, err := s.Client.GetRecommendation(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Recommendation
	return nil
}

func (s *OptimizerRecommendationResourceCrud) Update() error {
	request := oci_optimizer.UpdateRecommendationRequest{}

	tmp := s.D.Id()
	request.RecommendationId = &tmp

	if status, ok := s.D.GetOkExists("status"); ok {
		request.Status = oci_optimizer.StatusEnum(status.(string))
	}

	if timeStatusEnd, ok := s.D.GetOkExists("time_status_end"); ok {
		tmp, err := time.Parse(time.RFC3339, timeStatusEnd.(string))
		if err != nil {
			return err
		}
		request.TimeStatusEnd = &oci_common.SDKTime{Time: tmp}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "optimizer")

	response, err := s.Client.UpdateRecommendation(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Recommendation
	return nil
}

func (s *OptimizerRecommendationResourceCrud) SetData() error {
	if s.Res.Id != nil {
		s.D.Set("recommendation_id", *s.Res.Id)
	}

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

func RecommendationSummaryToMap(obj oci_optimizer.RecommendationSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.CategoryId != nil {
		result["category_id"] = string(*obj.CategoryId)
	}

	if obj.CompartmentId != nil {
		result["compartment_id"] = string(*obj.CompartmentId)
	}

	if obj.Description != nil {
		result["description"] = string(*obj.Description)
	}

	if obj.EstimatedCostSaving != nil {
		result["estimated_cost_saving"] = float32(*obj.EstimatedCostSaving)
	}

	result["extended_metadata"] = obj.ExtendedMetadata

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	result["importance"] = string(obj.Importance)

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	resourceCounts := []interface{}{}
	for _, item := range obj.ResourceCounts {
		resourceCounts = append(resourceCounts, ResourceCountToMap(item))
	}
	result["resource_counts"] = resourceCounts

	result["state"] = string(obj.LifecycleState)

	result["status"] = string(obj.Status)

	if obj.SupportedLevels != nil {
		result["supported_levels"] = []interface{}{SupportedLevelsToMap(obj.SupportedLevels)}
	}

	if obj.TimeCreated != nil {
		result["time_created"] = obj.TimeCreated.String()
	}

	if obj.TimeStatusBegin != nil {
		result["time_status_begin"] = obj.TimeStatusBegin.String()
	}

	if obj.TimeStatusEnd != nil {
		result["time_status_end"] = obj.TimeStatusEnd.String()
	}

	if obj.TimeUpdated != nil {
		result["time_updated"] = obj.TimeUpdated.String()
	}

	return result
}

func SupportedLevelToMap(obj oci_optimizer.SupportedLevel) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	return result
}

func SupportedLevelsToMap(obj *oci_optimizer.SupportedLevels) map[string]interface{} {
	result := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range obj.Items {
		items = append(items, SupportedLevelToMap(item))
	}
	result["items"] = items

	return result
}
