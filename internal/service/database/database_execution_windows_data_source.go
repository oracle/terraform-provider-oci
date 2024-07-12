// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package database

import (
	"context"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_database "github.com/oracle/oci-go-sdk/v65/database"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DatabaseExecutionWindowsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readDatabaseExecutionWindows,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"execution_resource_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"execution_windows": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     tfresource.GetDataSourceItemSchema(DatabaseExecutionWindowResource()),
			},
		},
	}
}

func readDatabaseExecutionWindows(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseExecutionWindowsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()

	return tfresource.ReadResource(sync)
}

type DatabaseExecutionWindowsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_database.DatabaseClient
	Res    *oci_database.ListExecutionWindowsResponse
}

func (s *DatabaseExecutionWindowsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DatabaseExecutionWindowsDataSourceCrud) Get() error {
	request := oci_database.ListExecutionWindowsRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if executionResourceId, ok := s.D.GetOkExists("execution_resource_id"); ok {
		tmp := executionResourceId.(string)
		request.ExecutionResourceId = &tmp
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_database.ExecutionWindowSummaryLifecycleStateEnum(state.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "database")

	response, err := s.Client.ListExecutionWindows(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListExecutionWindows(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *DatabaseExecutionWindowsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DatabaseExecutionWindowsDataSource-", DatabaseExecutionWindowsDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		executionWindow := map[string]interface{}{
			"compartment_id": *r.CompartmentId,
		}

		if r.DefinedTags != nil {
			executionWindow["defined_tags"] = tfresource.DefinedTagsToMap(r.DefinedTags)
		}

		if r.Description != nil {
			executionWindow["description"] = *r.Description
		}

		if r.DisplayName != nil {
			executionWindow["display_name"] = *r.DisplayName
		}

		if r.EstimatedTimeInMins != nil {
			executionWindow["estimated_time_in_mins"] = *r.EstimatedTimeInMins
		}

		if r.ExecutionResourceId != nil {
			executionWindow["execution_resource_id"] = *r.ExecutionResourceId
		}

		executionWindow["freeform_tags"] = r.FreeformTags

		if r.Id != nil {
			executionWindow["id"] = *r.Id
		}

		if r.IsEnforcedDuration != nil {
			executionWindow["is_enforced_duration"] = *r.IsEnforcedDuration
		}

		if r.LifecycleDetails != nil {
			executionWindow["lifecycle_details"] = *r.LifecycleDetails
		}

		executionWindow["lifecycle_substate"] = r.LifecycleSubstate

		executionWindow["state"] = r.LifecycleState

		if r.TimeCreated != nil {
			executionWindow["time_created"] = r.TimeCreated.String()
		}

		if r.TimeEnded != nil {
			executionWindow["time_ended"] = r.TimeEnded.String()
		}

		if r.TimeScheduled != nil {
			executionWindow["time_scheduled"] = r.TimeScheduled.Format(time.RFC3339Nano)
		}

		if r.TimeStarted != nil {
			executionWindow["time_started"] = r.TimeStarted.String()
		}

		if r.TimeUpdated != nil {
			executionWindow["time_updated"] = r.TimeUpdated.String()
		}

		if r.TotalTimeTakenInMins != nil {
			executionWindow["total_time_taken_in_mins"] = *r.TotalTimeTakenInMins
		}

		if r.WindowDurationInMins != nil {
			executionWindow["window_duration_in_mins"] = *r.WindowDurationInMins
		}

		executionWindow["window_type"] = r.WindowType

		resources = append(resources, executionWindow)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = tfresource.ApplyFilters(f.(*schema.Set), resources, DatabaseExecutionWindowsDataSource().Schema["execution_windows"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("execution_windows", resources); err != nil {
		return err
	}

	return nil
}
