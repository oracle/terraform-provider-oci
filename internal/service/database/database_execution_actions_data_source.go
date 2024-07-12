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

func DatabaseExecutionActionsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readDatabaseExecutionActions,
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
			"execution_window_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"execution_actions": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     tfresource.GetDataSourceItemSchema(DatabaseExecutionActionResource()),
			},
		},
	}
}

func readDatabaseExecutionActions(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseExecutionActionsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()

	return tfresource.ReadResource(sync)
}

type DatabaseExecutionActionsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_database.DatabaseClient
	Res    *oci_database.ListExecutionActionsResponse
}

func (s *DatabaseExecutionActionsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DatabaseExecutionActionsDataSourceCrud) Get() error {
	request := oci_database.ListExecutionActionsRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if executionWindowId, ok := s.D.GetOkExists("execution_window_id"); ok {
		tmp := executionWindowId.(string)
		request.ExecutionWindowId = &tmp
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_database.ExecutionActionSummaryLifecycleStateEnum(state.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "database")

	response, err := s.Client.ListExecutionActions(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListExecutionActions(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *DatabaseExecutionActionsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DatabaseExecutionActionsDataSource-", DatabaseExecutionActionsDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		executionAction := map[string]interface{}{
			"compartment_id": *r.CompartmentId,
		}

		actionMembers := []interface{}{}
		for _, item := range r.ActionMembers {
			actionMembers = append(actionMembers, ExecutionActionMemberToMap(item))
		}
		executionAction["action_members"] = actionMembers

		executionAction["action_params"] = r.ActionParams

		executionAction["action_type"] = r.ActionType

		if r.DefinedTags != nil {
			executionAction["defined_tags"] = tfresource.DefinedTagsToMap(r.DefinedTags)
		}

		if r.Description != nil {
			executionAction["description"] = *r.Description
		}

		if r.DisplayName != nil {
			executionAction["display_name"] = *r.DisplayName
		}

		if r.EstimatedTimeInMins != nil {
			executionAction["estimated_time_in_mins"] = *r.EstimatedTimeInMins
		}

		if r.ExecutionActionOrder != nil {
			executionAction["execution_action_order"] = *r.ExecutionActionOrder
		}

		if r.ExecutionWindowId != nil {
			executionAction["execution_window_id"] = *r.ExecutionWindowId
		}

		executionAction["freeform_tags"] = r.FreeformTags

		if r.Id != nil {
			executionAction["id"] = *r.Id
		}

		if r.LifecycleDetails != nil {
			executionAction["lifecycle_details"] = *r.LifecycleDetails
		}

		executionAction["lifecycle_substate"] = r.LifecycleSubstate

		executionAction["state"] = r.LifecycleState

		if r.TimeCreated != nil {
			executionAction["time_created"] = r.TimeCreated.String()
		}

		if r.TimeUpdated != nil {
			executionAction["time_updated"] = r.TimeUpdated.String()
		}

		if r.TotalTimeTakenInMins != nil {
			executionAction["total_time_taken_in_mins"] = *r.TotalTimeTakenInMins
		}

		resources = append(resources, executionAction)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = tfresource.ApplyFilters(f.(*schema.Set), resources, DatabaseExecutionActionsDataSource().Schema["execution_actions"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("execution_actions", resources); err != nil {
		return err
	}

	return nil
}
