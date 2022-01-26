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

func OptimizerResourceActionResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createOptimizerResourceAction,
		Read:     readOptimizerResourceAction,
		Update:   updateOptimizerResourceAction,
		Delete:   deleteOptimizerResourceAction,
		Schema: map[string]*schema.Schema{
			// Required
			"resource_action_id": {
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
			"action": {
				Type:     schema.TypeList,
				Computed: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"description": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"type": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"url": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"category_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"compartment_name": {
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
			"metadata": {
				Type:     schema.TypeMap,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"recommendation_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"resource_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"resource_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"state": {
				Type:     schema.TypeString,
				Computed: true,
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

func createOptimizerResourceAction(d *schema.ResourceData, m interface{}) error {
	sync := &OptimizerResourceActionResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).OptimizerClient()

	return tfresource.CreateResource(d, sync)
}

func readOptimizerResourceAction(d *schema.ResourceData, m interface{}) error {
	sync := &OptimizerResourceActionResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).OptimizerClient()

	return tfresource.ReadResource(sync)
}

func updateOptimizerResourceAction(d *schema.ResourceData, m interface{}) error {
	sync := &OptimizerResourceActionResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).OptimizerClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteOptimizerResourceAction(d *schema.ResourceData, m interface{}) error {
	return nil
}

type OptimizerResourceActionResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_optimizer.OptimizerClient
	Res                    *oci_optimizer.ResourceAction
	DisableNotFoundRetries bool
}

func (s *OptimizerResourceActionResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *OptimizerResourceActionResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_optimizer.LifecycleStateAttaching),
		string(oci_optimizer.LifecycleStateCreating),
	}
}

func (s *OptimizerResourceActionResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_optimizer.LifecycleStateActive),
	}
}

func (s *OptimizerResourceActionResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_optimizer.LifecycleStateDetaching),
		string(oci_optimizer.LifecycleStateDeleting),
	}
}

func (s *OptimizerResourceActionResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_optimizer.LifecycleStateDeleted),
	}
}

func (s *OptimizerResourceActionResourceCrud) Create() error {
	request := oci_optimizer.UpdateResourceActionRequest{}

	if resourceActionId, ok := s.D.GetOkExists("resource_action_id"); ok {
		tmp := resourceActionId.(string)
		request.ResourceActionId = &tmp
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

	response, err := s.Client.UpdateResourceAction(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.ResourceAction
	return nil
}

func (s *OptimizerResourceActionResourceCrud) Get() error {
	request := oci_optimizer.GetResourceActionRequest{}

	tmp := s.D.Id()
	request.ResourceActionId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "optimizer")

	response, err := s.Client.GetResourceAction(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.ResourceAction
	return nil
}

func (s *OptimizerResourceActionResourceCrud) Update() error {
	request := oci_optimizer.UpdateResourceActionRequest{}

	tmp := s.D.Id()
	request.ResourceActionId = &tmp

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

	response, err := s.Client.UpdateResourceAction(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.ResourceAction
	return nil
}

func (s *OptimizerResourceActionResourceCrud) SetData() error {
	if s.Res.Id != nil {
		s.D.Set("resource_action_id", *s.Res.Id)
	}

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

	s.D.Set("extended_metadata", tfresource.GenericMapToJsonMap(s.Res.ExtendedMetadata))

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

func ResourceActionSummaryToMap(obj oci_optimizer.ResourceActionSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Action != nil {
		result["action"] = []interface{}{OptimizerActionToMap(obj.Action)}
	}

	if obj.CategoryId != nil {
		result["category_id"] = string(*obj.CategoryId)
	}

	if obj.CompartmentId != nil {
		result["compartment_id"] = string(*obj.CompartmentId)
	}

	if obj.CompartmentName != nil {
		result["compartment_name"] = string(*obj.CompartmentName)
	}

	if obj.EstimatedCostSaving != nil {
		result["estimated_cost_saving"] = float32(*obj.EstimatedCostSaving)
	}

	result["extended_metadata"] = tfresource.GenericMapToJsonMap(obj.ExtendedMetadata)

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	result["metadata"] = obj.Metadata

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	if obj.RecommendationId != nil {
		result["recommendation_id"] = string(*obj.RecommendationId)
	}

	if obj.ResourceId != nil {
		result["resource_id"] = string(*obj.ResourceId)
	}

	if obj.ResourceType != nil {
		result["resource_type"] = string(*obj.ResourceType)
	}

	result["state"] = string(obj.LifecycleState)

	result["status"] = string(obj.Status)

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
