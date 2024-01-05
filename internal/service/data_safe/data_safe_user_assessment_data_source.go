// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package data_safe

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_data_safe "github.com/oracle/oci-go-sdk/v65/datasafe"
)

func DataSafeUserAssessmentDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["user_assessment_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(DataSafeUserAssessmentResource(), fieldMap, readSingularDataSafeUserAssessment)
}

func readSingularDataSafeUserAssessment(d *schema.ResourceData, m interface{}) error {
	sync := &DataSafeUserAssessmentDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataSafeClient()

	return tfresource.ReadResource(sync)
}

type DataSafeUserAssessmentDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_data_safe.DataSafeClient
	Res    *oci_data_safe.GetUserAssessmentResponse
}

func (s *DataSafeUserAssessmentDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DataSafeUserAssessmentDataSourceCrud) Get() error {
	request := oci_data_safe.GetUserAssessmentRequest{}

	if userAssessmentId, ok := s.D.GetOkExists("user_assessment_id"); ok {
		tmp := userAssessmentId.(string)
		request.UserAssessmentId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "data_safe")

	response, err := s.Client.GetUserAssessment(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *DataSafeUserAssessmentDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

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

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	ignoredAssessmentIds := []interface{}{}
	for _, item := range s.Res.IgnoredAssessmentIds {
		ignoredAssessmentIds = append(ignoredAssessmentIds, item)
	}
	s.D.Set("ignored_assessment_ids", ignoredAssessmentIds)

	ignoredTargets := []interface{}{}
	for _, item := range s.Res.IgnoredTargets {
		ignoredTargets = append(ignoredTargets, IgnoredTargetsToMap(item))
	}
	s.D.Set("ignored_targets", ignoredTargets)

	if s.Res.IsBaseline != nil {
		s.D.Set("is_baseline", *s.Res.IsBaseline)
	}

	if s.Res.IsDeviatedFromBaseline != nil {
		s.D.Set("is_deviated_from_baseline", *s.Res.IsDeviatedFromBaseline)
	}

	if s.Res.LastComparedBaselineId != nil {
		s.D.Set("last_compared_baseline_id", *s.Res.LastComparedBaselineId)
	}

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	if s.Res.Schedule != nil {
		s.D.Set("schedule", *s.Res.Schedule)
	}

	if s.Res.ScheduleAssessmentId != nil {
		s.D.Set("schedule_assessment_id", *s.Res.ScheduleAssessmentId)
	}

	s.D.Set("state", s.Res.LifecycleState)

	var statsBytes, err = json.Marshal(s.Res.Statistics)
	if err != nil {
		return fmt.Errorf("unable to Marshal Statistics, encountered error: %v", err)
	}
	statDetails := string(statsBytes)
	s.D.Set("statistics", statDetails)

	if s.Res.SystemTags != nil {
		s.D.Set("system_tags", tfresource.SystemTagsToMap(s.Res.SystemTags))
	}

	s.D.Set("target_ids", s.Res.TargetIds)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeLastAssessed != nil {
		s.D.Set("time_last_assessed", s.Res.TimeLastAssessed.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	s.D.Set("triggered_by", s.Res.TriggeredBy)

	s.D.Set("type", s.Res.Type)

	return nil
}
