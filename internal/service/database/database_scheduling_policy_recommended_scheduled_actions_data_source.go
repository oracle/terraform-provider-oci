// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package database

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_database "github.com/oracle/oci-go-sdk/v65/database"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DatabaseSchedulingPolicyRecommendedScheduledActionsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readDatabaseSchedulingPolicyRecommendedScheduledActions,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"plan_intent": {
				Type:     schema.TypeString,
				Required: true,
			},
			"scheduling_policy_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"scheduling_policy_target_resource_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"recommended_scheduled_actions_collection": {
				Type:     schema.TypeList,
				Computed: true,
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
									"action_members": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional

												// Computed
												"estimated_time_in_mins": {
													Type:     schema.TypeInt,
													Computed: true,
												},
												"member_id": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"member_order": {
													Type:     schema.TypeInt,
													Computed: true,
												},
											},
										},
									},
									"action_order": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									"action_params": {
										Type:     schema.TypeMap,
										Computed: true,
										Elem:     schema.TypeString,
									},
									"action_type": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"display_name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"estimated_time_in_mins": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									"scheduling_window_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func readDatabaseSchedulingPolicyRecommendedScheduledActions(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseSchedulingPolicyRecommendedScheduledActionsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()

	return tfresource.ReadResource(sync)
}

type DatabaseSchedulingPolicyRecommendedScheduledActionsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_database.DatabaseClient
	Res    *oci_database.ListRecommendedScheduledActionsResponse
}

func (s *DatabaseSchedulingPolicyRecommendedScheduledActionsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DatabaseSchedulingPolicyRecommendedScheduledActionsDataSourceCrud) Get() error {
	request := oci_database.ListRecommendedScheduledActionsRequest{}

	if planIntent, ok := s.D.GetOkExists("plan_intent"); ok {
		request.PlanIntent = oci_database.ListRecommendedScheduledActionsPlanIntentEnum(planIntent.(string))
	}

	if schedulingPolicyId, ok := s.D.GetOkExists("scheduling_policy_id"); ok {
		tmp := schedulingPolicyId.(string)
		request.SchedulingPolicyId = &tmp
	}

	if schedulingPolicyTargetResourceId, ok := s.D.GetOkExists("scheduling_policy_target_resource_id"); ok {
		tmp := schedulingPolicyTargetResourceId.(string)
		request.SchedulingPolicyTargetResourceId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "database")

	response, err := s.Client.ListRecommendedScheduledActions(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListRecommendedScheduledActions(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *DatabaseSchedulingPolicyRecommendedScheduledActionsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DatabaseSchedulingPolicyRecommendedScheduledActionsDataSource-", DatabaseSchedulingPolicyRecommendedScheduledActionsDataSource(), s.D))
	resources := []map[string]interface{}{}
	schedulingPolicyRecommendedScheduledAction := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, RecommendedScheduledActionSummaryToMap(item))
	}
	schedulingPolicyRecommendedScheduledAction["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, DatabaseSchedulingPolicyRecommendedScheduledActionsDataSource().Schema["recommended_scheduled_actions_collection"].Elem.(*schema.Resource).Schema)
		schedulingPolicyRecommendedScheduledAction["items"] = items
	}

	resources = append(resources, schedulingPolicyRecommendedScheduledAction)
	if err := s.D.Set("recommended_scheduled_actions_collection", resources); err != nil {
		return err
	}

	return nil
}

func RecommendedScheduledActionSummaryToMap(obj oci_database.RecommendedScheduledActionSummary) map[string]interface{} {
	result := map[string]interface{}{}

	actionMembers := []interface{}{}
	for _, item := range obj.ActionMembers {
		actionMembers = append(actionMembers, ActionMemberToMap(item))
	}
	result["action_members"] = actionMembers

	if obj.ActionOrder != nil {
		result["action_order"] = int(*obj.ActionOrder)
	}

	result["action_params"] = obj.ActionParams

	result["action_type"] = string(obj.ActionType)

	if obj.DisplayName != nil {
		result["display_name"] = string(*obj.DisplayName)
	}

	if obj.EstimatedTimeInMins != nil {
		result["estimated_time_in_mins"] = int(*obj.EstimatedTimeInMins)
	}

	if obj.SchedulingWindowId != nil {
		result["scheduling_window_id"] = string(*obj.SchedulingWindowId)
	}

	return result
}
