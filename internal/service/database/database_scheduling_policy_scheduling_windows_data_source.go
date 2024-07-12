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

func DatabaseSchedulingPolicySchedulingWindowsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readDatabaseSchedulingPolicySchedulingWindows,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"scheduling_policy_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"scheduling_windows": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     tfresource.GetDataSourceItemSchema(DatabaseSchedulingPolicySchedulingWindowResource()),
			},
		},
	}
}

func readDatabaseSchedulingPolicySchedulingWindows(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseSchedulingPolicySchedulingWindowsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()

	return tfresource.ReadResource(sync)
}

type DatabaseSchedulingPolicySchedulingWindowsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_database.DatabaseClient
	Res    *oci_database.ListSchedulingWindowsResponse
}

func (s *DatabaseSchedulingPolicySchedulingWindowsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DatabaseSchedulingPolicySchedulingWindowsDataSourceCrud) Get() error {
	request := oci_database.ListSchedulingWindowsRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if schedulingPolicyId, ok := s.D.GetOkExists("scheduling_policy_id"); ok {
		tmp := schedulingPolicyId.(string)
		request.SchedulingPolicyId = &tmp
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_database.SchedulingWindowSummaryLifecycleStateEnum(state.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "database")

	response, err := s.Client.ListSchedulingWindows(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListSchedulingWindows(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *DatabaseSchedulingPolicySchedulingWindowsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DatabaseSchedulingPolicySchedulingWindowsDataSource-", DatabaseSchedulingPolicySchedulingWindowsDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		schedulingPolicySchedulingWindow := map[string]interface{}{
			"scheduling_policy_id": *r.SchedulingPolicyId,
		}

		if r.CompartmentId != nil {
			schedulingPolicySchedulingWindow["compartment_id"] = *r.CompartmentId
		}

		if r.DefinedTags != nil {
			schedulingPolicySchedulingWindow["defined_tags"] = tfresource.DefinedTagsToMap(r.DefinedTags)
		}

		if r.DisplayName != nil {
			schedulingPolicySchedulingWindow["display_name"] = *r.DisplayName
		}

		schedulingPolicySchedulingWindow["freeform_tags"] = r.FreeformTags

		if r.Id != nil {
			schedulingPolicySchedulingWindow["id"] = *r.Id
		}

		if r.LifecycleDetails != nil {
			schedulingPolicySchedulingWindow["lifecycle_details"] = *r.LifecycleDetails
		}

		schedulingPolicySchedulingWindow["state"] = r.LifecycleState

		if r.TimeCreated != nil {
			schedulingPolicySchedulingWindow["time_created"] = r.TimeCreated.String()
		}

		if r.TimeNextSchedulingWindowStarts != nil {
			schedulingPolicySchedulingWindow["time_next_scheduling_window_starts"] = r.TimeNextSchedulingWindowStarts.String()
		}

		if r.TimeUpdated != nil {
			schedulingPolicySchedulingWindow["time_updated"] = r.TimeUpdated.String()
		}

		if r.WindowPreference != nil {
			schedulingPolicySchedulingWindow["window_preference"] = []interface{}{WindowPreferenceDetailToMap(r.WindowPreference)}
		} else {
			schedulingPolicySchedulingWindow["window_preference"] = nil
		}

		resources = append(resources, schedulingPolicySchedulingWindow)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = tfresource.ApplyFilters(f.(*schema.Set), resources, DatabaseSchedulingPolicySchedulingWindowsDataSource().Schema["scheduling_windows"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("scheduling_windows", resources); err != nil {
		return err
	}

	return nil
}
