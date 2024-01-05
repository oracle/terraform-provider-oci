// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package fusion_apps

import (
	"context"
	"log"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_common "github.com/oracle/oci-go-sdk/v65/common"
	oci_fusion_apps "github.com/oracle/oci-go-sdk/v65/fusionapps"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func FusionAppsFusionEnvironmentScheduledActivitiesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readFusionAppsFusionEnvironmentScheduledActivities,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"fusion_environment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"run_cycle": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"time_expected_finish_less_than_or_equal_to": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"time_scheduled_start_greater_than_or_equal_to": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"scheduled_activity_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"actions": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional

												// Computed
												"action_type": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"artifact": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"category": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"description": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"mode": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"qualifier": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"reference_key": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"state": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"version": {
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},
									"delay_in_hours": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									"display_name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"fusion_environment_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"lifecycle_details": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"run_cycle": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"service_availability": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"state": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"time_accepted": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"time_expected_finish": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"time_finished": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"time_scheduled_start": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"time_updated": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"freeform_tags": {
										Type:     schema.TypeMap,
										Computed: true,
										Elem:     schema.TypeString,
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

func readFusionAppsFusionEnvironmentScheduledActivities(d *schema.ResourceData, m interface{}) error {
	sync := &FusionAppsFusionEnvironmentScheduledActivitiesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).FusionApplicationsClient()

	return tfresource.ReadResource(sync)
}

type FusionAppsFusionEnvironmentScheduledActivitiesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_fusion_apps.FusionApplicationsClient
	Res    *oci_fusion_apps.ListScheduledActivitiesResponse
}

func (s *FusionAppsFusionEnvironmentScheduledActivitiesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *FusionAppsFusionEnvironmentScheduledActivitiesDataSourceCrud) Get() error {
	request := oci_fusion_apps.ListScheduledActivitiesRequest{}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if fusionEnvironmentId, ok := s.D.GetOkExists("fusion_environment_id"); ok {
		tmp := fusionEnvironmentId.(string)
		request.FusionEnvironmentId = &tmp
	}

	if runCycle, ok := s.D.GetOkExists("run_cycle"); ok {
		request.RunCycle = oci_fusion_apps.ScheduledActivityRunCycleEnum(runCycle.(string))
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_fusion_apps.ScheduledActivityLifecycleStateEnum(state.(string))
	}

	if timeExpectedFinishLessThanOrEqualTo, ok := s.D.GetOkExists("time_expected_finish_less_than_or_equal_to"); ok {
		tmp, err := time.Parse(time.RFC3339, timeExpectedFinishLessThanOrEqualTo.(string))
		if err != nil {
			return err
		}
		request.TimeExpectedFinishLessThanOrEqualTo = &oci_common.SDKTime{Time: tmp}
	}

	if timeScheduledStartGreaterThanOrEqualTo, ok := s.D.GetOkExists("time_scheduled_start_greater_than_or_equal_to"); ok {
		tmp, err := time.Parse(time.RFC3339, timeScheduledStartGreaterThanOrEqualTo.(string))
		if err != nil {
			return err
		}
		request.TimeScheduledStartGreaterThanOrEqualTo = &oci_common.SDKTime{Time: tmp}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "fusion_apps")

	response, err := s.Client.ListScheduledActivities(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListScheduledActivities(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *FusionAppsFusionEnvironmentScheduledActivitiesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("FusionAppsFusionEnvironmentScheduledActivitiesDataSource-", FusionAppsFusionEnvironmentScheduledActivitiesDataSource(), s.D))
	resources := []map[string]interface{}{}
	fusionEnvironmentScheduledActivity := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, ScheduledActivitySummaryToMap(item))
	}
	fusionEnvironmentScheduledActivity["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, FusionAppsFusionEnvironmentScheduledActivitiesDataSource().Schema["scheduled_activity_collection"].Elem.(*schema.Resource).Schema)
		fusionEnvironmentScheduledActivity["items"] = items
	}

	resources = append(resources, fusionEnvironmentScheduledActivity)
	if err := s.D.Set("scheduled_activity_collection", resources); err != nil {
		return err
	}

	return nil
}

func ActionToMap(obj oci_fusion_apps.Action) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (obj).(type) {
	case oci_fusion_apps.PatchAction:
		result["action_type"] = "PATCH"

		if v.Artifact != nil {
			result["artifact"] = string(*v.Artifact)
		}

		result["category"] = string(v.Category)

		result["mode"] = string(v.Mode)
	case oci_fusion_apps.UpgradeAction:
		result["action_type"] = "QUARTERLY_UPGRADE"

		if v.Qualifier != nil {
			result["qualifier"] = string(*v.Qualifier)
		}

		if v.Version != nil {
			result["version"] = string(*v.Version)
		}
	case oci_fusion_apps.VertexAction:
		result["action_type"] = "VERTEX"

		if v.Artifact != nil {
			result["artifact"] = string(*v.Artifact)
		}
	default:
		log.Printf("[WARN] Received 'action_type' of unknown type %v", obj)
		return nil
	}

	return result
}

func ScheduledActivitySummaryToMap(obj oci_fusion_apps.ScheduledActivitySummary) map[string]interface{} {
	result := map[string]interface{}{}

	actions := []interface{}{}
	for _, item := range obj.Actions {
		actions = append(actions, ActionToMap(item))
	}
	result["actions"] = actions

	if obj.DefinedTags != nil {
		result["defined_tags"] = tfresource.DefinedTagsToMap(obj.DefinedTags)
	}

	if obj.DelayInHours != nil {
		result["delay_in_hours"] = int(*obj.DelayInHours)
	}

	if obj.DisplayName != nil {
		result["display_name"] = string(*obj.DisplayName)
	}

	result["freeform_tags"] = obj.FreeformTags

	if obj.FusionEnvironmentId != nil {
		result["fusion_environment_id"] = string(*obj.FusionEnvironmentId)
	}

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	if obj.LifecycleDetails != nil {
		result["lifecycle_details"] = string(*obj.LifecycleDetails)
	}

	result["run_cycle"] = string(obj.RunCycle)

	result["service_availability"] = string(obj.ServiceAvailability)

	result["state"] = string(obj.LifecycleState)

	if obj.TimeAccepted != nil {
		result["time_accepted"] = obj.TimeAccepted.String()
	}

	if obj.TimeExpectedFinish != nil {
		result["time_expected_finish"] = obj.TimeExpectedFinish.String()
	}

	if obj.TimeFinished != nil {
		result["time_finished"] = obj.TimeFinished.String()
	}

	if obj.TimeScheduledStart != nil {
		result["time_scheduled_start"] = obj.TimeScheduledStart.String()
	}

	if obj.TimeUpdated != nil {
		result["time_updated"] = obj.TimeUpdated.String()
	}

	return result
}
