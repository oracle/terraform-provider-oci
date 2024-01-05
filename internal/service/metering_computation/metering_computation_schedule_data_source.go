// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package metering_computation

import (
	"context"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_metering_computation "github.com/oracle/oci-go-sdk/v65/usageapi"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func MeteringComputationScheduleDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["schedule_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(MeteringComputationScheduleResource(), fieldMap, readSingularMeteringComputationSchedule)
}

func readSingularMeteringComputationSchedule(d *schema.ResourceData, m interface{}) error {
	sync := &MeteringComputationScheduleDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).UsageapiClient()

	return tfresource.ReadResource(sync)
}

type MeteringComputationScheduleDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_metering_computation.UsageapiClient
	Res    *oci_metering_computation.GetScheduleResponse
}

func (s *MeteringComputationScheduleDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *MeteringComputationScheduleDataSourceCrud) Get() error {
	request := oci_metering_computation.GetScheduleRequest{}

	if scheduleId, ok := s.D.GetOkExists("schedule_id"); ok {
		tmp := scheduleId.(string)
		request.ScheduleId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "metering_computation")

	response, err := s.Client.GetSchedule(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *MeteringComputationScheduleDataSourceCrud) SetData() error {
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

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.Name != nil {
		s.D.Set("name", *s.Res.Name)
	}

	s.D.Set("output_file_format", s.Res.OutputFileFormat)

	if s.Res.QueryProperties != nil {
		s.D.Set("query_properties", []interface{}{QueryPropertiesToMap(s.Res.QueryProperties)})
	} else {
		s.D.Set("query_properties", nil)
	}

	if s.Res.ResultLocation != nil {
		resultLocationArray := []interface{}{}
		if resultLocationMap := ResultLocationToMap(&s.Res.ResultLocation); resultLocationMap != nil {
			resultLocationArray = append(resultLocationArray, resultLocationMap)
		}
		s.D.Set("result_location", resultLocationArray)
	} else {
		s.D.Set("result_location", nil)
	}

	if s.Res.SavedReportId != nil {
		s.D.Set("saved_report_id", *s.Res.SavedReportId)
	}

	if s.Res.ScheduleRecurrences != nil {
		s.D.Set("schedule_recurrences", *s.Res.ScheduleRecurrences)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.SystemTags != nil {
		s.D.Set("system_tags", tfresource.SystemTagsToMap(s.Res.SystemTags))
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeNextRun != nil {
		s.D.Set("time_next_run", s.Res.TimeNextRun.String())
	}

	if s.Res.TimeScheduled != nil {
		s.D.Set("time_scheduled", s.Res.TimeScheduled.Format(time.RFC3339Nano))
	}

	return nil
}
