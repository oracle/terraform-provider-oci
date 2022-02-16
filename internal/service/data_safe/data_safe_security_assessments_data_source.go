// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package data_safe

import (
	"context"
	"time"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_common "github.com/oracle/oci-go-sdk/v58/common"
	oci_data_safe "github.com/oracle/oci-go-sdk/v58/datasafe"
)

func DataSafeSecurityAssessmentsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readDataSafeSecurityAssessments,
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
			"schedule_assessment_id": {
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
			"security_assessments": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     tfresource.GetDataSourceItemSchema(DataSafeSecurityAssessmentResource()),
			},
		},
	}
}

func readDataSafeSecurityAssessments(d *schema.ResourceData, m interface{}) error {
	sync := &DataSafeSecurityAssessmentsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataSafeClient()

	return tfresource.ReadResource(sync)
}

type DataSafeSecurityAssessmentsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_data_safe.DataSafeClient
	Res    *oci_data_safe.ListSecurityAssessmentsResponse
}

func (s *DataSafeSecurityAssessmentsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DataSafeSecurityAssessmentsDataSourceCrud) Get() error {
	request := oci_data_safe.ListSecurityAssessmentsRequest{}

	if accessLevel, ok := s.D.GetOkExists("access_level"); ok {
		request.AccessLevel = oci_data_safe.ListSecurityAssessmentsAccessLevelEnum(accessLevel.(string))
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

	if scheduleAssessmentId, ok := s.D.GetOkExists("schedule_assessment_id"); ok {
		tmp := scheduleAssessmentId.(string)
		request.ScheduleAssessmentId = &tmp
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_data_safe.ListSecurityAssessmentsLifecycleStateEnum(state.(string))
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
		request.TriggeredBy = oci_data_safe.ListSecurityAssessmentsTriggeredByEnum(triggeredBy.(string))
	}

	if type_, ok := s.D.GetOkExists("type"); ok {
		request.Type = oci_data_safe.ListSecurityAssessmentsTypeEnum(type_.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "data_safe")

	response, err := s.Client.ListSecurityAssessments(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListSecurityAssessments(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *DataSafeSecurityAssessmentsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DataSafeSecurityAssessmentsDataSource-", DataSafeSecurityAssessmentsDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		securityAssessment := map[string]interface{}{
			"compartment_id": *r.CompartmentId,
		}

		if r.DefinedTags != nil {
			securityAssessment["defined_tags"] = tfresource.DefinedTagsToMap(r.DefinedTags)
		}

		if r.Description != nil {
			securityAssessment["description"] = *r.Description
		}

		if r.DisplayName != nil {
			securityAssessment["display_name"] = *r.DisplayName
		}

		securityAssessment["freeform_tags"] = r.FreeformTags

		if r.Id != nil {
			securityAssessment["id"] = *r.Id
		}

		ignoredAssessmentIds := []interface{}{}
		for _, item := range r.IgnoredAssessmentIds {
			ignoredAssessmentIds = append(ignoredAssessmentIds, item)
		}
		securityAssessment["ignored_assessment_ids"] = ignoredAssessmentIds

		if r.IsBaseline != nil {
			securityAssessment["is_baseline"] = *r.IsBaseline
		}

		if r.IsDeviatedFromBaseline != nil {
			securityAssessment["is_deviated_from_baseline"] = *r.IsDeviatedFromBaseline
		}

		if r.LastComparedBaselineId != nil {
			securityAssessment["last_compared_baseline_id"] = *r.LastComparedBaselineId
		}

		if r.LifecycleDetails != nil {
			securityAssessment["lifecycle_details"] = *r.LifecycleDetails
		}

		if r.Link != nil {
			securityAssessment["link"] = *r.Link
		}

		if r.Schedule != nil {
			securityAssessment["schedule"] = *r.Schedule
		}

		if r.ScheduleSecurityAssessmentId != nil {
			securityAssessment["schedule_security_assessment_id"] = *r.ScheduleSecurityAssessmentId
		}

		securityAssessment["state"] = r.LifecycleState

		if r.Statistics != nil {
			securityAssessment["statistics"] = []interface{}{SecurityAssessmentStatisticsToMap(r.Statistics)}
		} else {
			securityAssessment["statistics"] = nil
		}

		securityAssessment["target_ids"] = r.TargetIds

		if r.TimeCreated != nil {
			securityAssessment["time_created"] = r.TimeCreated.String()
		}

		if r.TimeUpdated != nil {
			securityAssessment["time_updated"] = r.TimeUpdated.String()
		}

		securityAssessment["triggered_by"] = r.TriggeredBy

		securityAssessment["type"] = r.Type

		resources = append(resources, securityAssessment)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = tfresource.ApplyFilters(f.(*schema.Set), resources, DataSafeSecurityAssessmentsDataSource().Schema["security_assessments"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("security_assessments", resources); err != nil {
		return err
	}

	return nil
}
