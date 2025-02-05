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

func DatabaseSchedulingPoliciesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readDatabaseSchedulingPolicies,
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
			"state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"scheduling_policies": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     tfresource.GetDataSourceItemSchema(DatabaseSchedulingPolicyResource()),
			},
		},
	}
}

func readDatabaseSchedulingPolicies(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseSchedulingPoliciesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()

	return tfresource.ReadResource(sync)
}

type DatabaseSchedulingPoliciesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_database.DatabaseClient
	Res    *oci_database.ListSchedulingPoliciesResponse
}

func (s *DatabaseSchedulingPoliciesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DatabaseSchedulingPoliciesDataSourceCrud) Get() error {
	request := oci_database.ListSchedulingPoliciesRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_database.SchedulingPolicySummaryLifecycleStateEnum(state.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "database")

	response, err := s.Client.ListSchedulingPolicies(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListSchedulingPolicies(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *DatabaseSchedulingPoliciesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DatabaseSchedulingPoliciesDataSource-", DatabaseSchedulingPoliciesDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		schedulingPolicy := map[string]interface{}{
			"compartment_id": *r.CompartmentId,
		}

		schedulingPolicy["cadence"] = r.Cadence

		if r.CadenceStartMonth != nil {
			schedulingPolicy["cadence_start_month"] = []interface{}{MonthToMapPolicy(r.CadenceStartMonth)}
		} else {
			schedulingPolicy["cadence_start_month"] = nil
		}

		if r.DefinedTags != nil {
			schedulingPolicy["defined_tags"] = tfresource.DefinedTagsToMap(r.DefinedTags)
		}

		if r.DisplayName != nil {
			schedulingPolicy["display_name"] = *r.DisplayName
		}

		schedulingPolicy["freeform_tags"] = r.FreeformTags

		if r.Id != nil {
			schedulingPolicy["id"] = *r.Id
		}

		if r.LifecycleDetails != nil {
			schedulingPolicy["lifecycle_details"] = *r.LifecycleDetails
		}

		schedulingPolicy["state"] = r.LifecycleState

		if r.TimeCreated != nil {
			schedulingPolicy["time_created"] = r.TimeCreated.String()
		}

		if r.TimeNextWindowStarts != nil {
			schedulingPolicy["time_next_window_starts"] = r.TimeNextWindowStarts.String()
		}

		if r.TimeUpdated != nil {
			schedulingPolicy["time_updated"] = r.TimeUpdated.String()
		}

		resources = append(resources, schedulingPolicy)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = tfresource.ApplyFilters(f.(*schema.Set), resources, DatabaseSchedulingPoliciesDataSource().Schema["scheduling_policies"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("scheduling_policies", resources); err != nil {
		return err
	}

	return nil
}
