// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package data_safe

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_common "github.com/oracle/oci-go-sdk/v65/common"
	oci_data_safe "github.com/oracle/oci-go-sdk/v65/datasafe"
)

func DataSafeUserAssessmentsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readDataSafeUserAssessments,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"access_level": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"compartment_id_in_subtree": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"is_baseline": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"is_schedule_assessment": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"schedule_user_assessment_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"target_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"time_created_greater_than_or_equal_to": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"time_created_less_than": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"triggered_by": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"type": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"user_assessments": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     tfresource.GetDataSourceItemSchema(DataSafeUserAssessmentResource()),
			},
			"ignored_targets": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"lifecycle_state": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"target_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"user_assessment_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func readDataSafeUserAssessments(d *schema.ResourceData, m interface{}) error {
	sync := &DataSafeUserAssessmentsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataSafeClient()

	return tfresource.ReadResource(sync)
}

type DataSafeUserAssessmentsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_data_safe.DataSafeClient
	Res    *oci_data_safe.ListUserAssessmentsResponse
}

func (s *DataSafeUserAssessmentsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DataSafeUserAssessmentsDataSourceCrud) Get() error {
	request := oci_data_safe.ListUserAssessmentsRequest{}

	if accessLevel, ok := s.D.GetOkExists("access_level"); ok {
		request.AccessLevel = oci_data_safe.ListUserAssessmentsAccessLevelEnum(accessLevel.(string))
	}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if compartmentIdInSubtree, ok := s.D.GetOkExists("compartment_id_in_subtree"); ok {
		tmp := compartmentIdInSubtree.(bool)
		request.CompartmentIdInSubtree = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if isBaseline, ok := s.D.GetOkExists("is_baseline"); ok {
		tmp := isBaseline.(bool)
		request.IsBaseline = &tmp
	}

	if isScheduleAssessment, ok := s.D.GetOkExists("is_schedule_assessment"); ok {
		tmp := isScheduleAssessment.(bool)
		request.IsScheduleAssessment = &tmp
	}

	if scheduleUserAssessmentId, ok := s.D.GetOkExists("schedule_user_assessment_id"); ok {
		tmp := scheduleUserAssessmentId.(string)
		request.ScheduleUserAssessmentId = &tmp
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_data_safe.ListUserAssessmentsLifecycleStateEnum(state.(string))
	}

	if targetId, ok := s.D.GetOkExists("target_id"); ok {
		tmp := targetId.(string)
		request.TargetId = &tmp
	}

	if timeCreatedGreaterThanOrEqualTo, ok := s.D.GetOkExists("time_created_greater_than_or_equal_to"); ok {
		tmp, err := time.Parse(time.RFC3339, timeCreatedGreaterThanOrEqualTo.(string))
		if err != nil {
			return err
		}
		request.TimeCreatedGreaterThanOrEqualTo = &oci_common.SDKTime{Time: tmp}
	}

	if timeCreatedLessThan, ok := s.D.GetOkExists("time_created_less_than"); ok {
		tmp, err := time.Parse(time.RFC3339, timeCreatedLessThan.(string))
		if err != nil {
			return err
		}
		request.TimeCreatedLessThan = &oci_common.SDKTime{Time: tmp}
	}

	if triggeredBy, ok := s.D.GetOkExists("triggered_by"); ok {
		request.TriggeredBy = oci_data_safe.ListUserAssessmentsTriggeredByEnum(triggeredBy.(string))
	}

	if type_, ok := s.D.GetOkExists("type"); ok {
		request.Type = oci_data_safe.ListUserAssessmentsTypeEnum(type_.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "data_safe")

	response, err := s.Client.ListUserAssessments(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListUserAssessments(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func IgnoredTargetsToMap(obj interface{}) interface{} {
	result := map[string]interface{}{}
	mobj, ok := obj.(map[string]interface{})
	if !ok {
		return nil
	}
	for k, v := range mobj {
		// key match, return value
		if k == "lifecycleState" {
			result["lifecycle_state"] = v
		}
		if k == "targetId" {
			result["target_id"] = v
		}
		if k == "userAssessmentId" {
			result["user_assessment_id"] = v
		}

	}
	return result
}

func (s *DataSafeUserAssessmentsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DataSafeUserAssessmentsDataSource-", DataSafeUserAssessmentsDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		userAssessment := map[string]interface{}{
			"compartment_id": *r.CompartmentId,
		}

		if r.DefinedTags != nil {
			userAssessment["defined_tags"] = tfresource.DefinedTagsToMap(r.DefinedTags)
		}

		if r.Description != nil {
			userAssessment["description"] = *r.Description
		}

		if r.DisplayName != nil {
			userAssessment["display_name"] = *r.DisplayName
		}

		userAssessment["freeform_tags"] = r.FreeformTags

		if r.Id != nil {
			userAssessment["id"] = *r.Id
		}

		ignoredAssessmentIds := []interface{}{}
		for _, item := range r.IgnoredAssessmentIds {
			ignoredAssessmentIds = append(ignoredAssessmentIds, item)
		}
		userAssessment["ignored_assessment_ids"] = ignoredAssessmentIds

		ignoredTargets := []interface{}{}
		for _, item := range r.IgnoredTargets {
			ignoredTargets = append(ignoredTargets, IgnoredTargetsToMap(item))
		}
		userAssessment["ignored_targets"] = ignoredTargets

		if r.IsAssessmentScheduled != nil {
			userAssessment["is_assessment_scheduled"] = *r.IsAssessmentScheduled
		}

		if r.IsBaseline != nil {
			userAssessment["is_baseline"] = *r.IsBaseline
		}

		if r.IsDeviatedFromBaseline != nil {
			userAssessment["is_deviated_from_baseline"] = *r.IsDeviatedFromBaseline
		}

		if r.LastComparedBaselineId != nil {
			userAssessment["last_compared_baseline_id"] = *r.LastComparedBaselineId
		}

		if r.LifecycleDetails != nil {
			userAssessment["lifecycle_details"] = *r.LifecycleDetails
		}

		if r.Schedule != nil {
			userAssessment["schedule"] = *r.Schedule
		}

		if r.ScheduleAssessmentId != nil {
			userAssessment["schedule_assessment_id"] = *r.ScheduleAssessmentId
		}

		userAssessment["state"] = r.LifecycleState

		var statsBytes, err = json.Marshal(r.Statistics)
		if err != nil {
			return fmt.Errorf("unable to Marshal Statistics, encountered error users: %v", err)
		}
		statDetails := string(statsBytes)

		userAssessment["statistics"] = statDetails

		userAssessment["target_ids"] = r.TargetIds

		if r.TimeCreated != nil {
			userAssessment["time_created"] = r.TimeCreated.String()
		}

		if r.TimeLastAssessed != nil {
			userAssessment["time_last_assessed"] = r.TimeLastAssessed.String()
		}

		if r.TimeUpdated != nil {
			userAssessment["time_updated"] = r.TimeUpdated.String()
		}

		userAssessment["triggered_by"] = r.TriggeredBy

		userAssessment["type"] = r.Type

		resources = append(resources, userAssessment)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = tfresource.ApplyFilters(f.(*schema.Set), resources, DataSafeUserAssessmentsDataSource().Schema["user_assessments"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("user_assessments", resources); err != nil {
		return err
	}

	return nil
}
