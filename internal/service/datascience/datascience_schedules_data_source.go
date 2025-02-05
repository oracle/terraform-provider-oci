// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package datascience

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_datascience "github.com/oracle/oci-go-sdk/v65/datascience"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DatascienceSchedulesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readDatascienceSchedules,
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
			"id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"project_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"schedules": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     tfresource.GetDataSourceItemSchema(DatascienceScheduleResource()),
			},
		},
	}
}

func readDatascienceSchedules(d *schema.ResourceData, m interface{}) error {
	sync := &DatascienceSchedulesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataScienceClient()

	return tfresource.ReadResource(sync)
}

type DatascienceSchedulesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_datascience.DataScienceClient
	Res    *oci_datascience.ListSchedulesResponse
}

func (s *DatascienceSchedulesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DatascienceSchedulesDataSourceCrud) Get() error {
	request := oci_datascience.ListSchedulesRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if id, ok := s.D.GetOkExists("id"); ok {
		tmp := id.(string)
		request.Id = &tmp
	}

	if projectId, ok := s.D.GetOkExists("project_id"); ok {
		tmp := projectId.(string)
		request.ProjectId = &tmp
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_datascience.ListSchedulesLifecycleStateEnum(state.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "datascience")

	response, err := s.Client.ListSchedules(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListSchedules(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *DatascienceSchedulesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DatascienceSchedulesDataSource-", DatascienceSchedulesDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		schedule := map[string]interface{}{
			"compartment_id": *r.CompartmentId,
		}

		if r.CreatedBy != nil {
			schedule["created_by"] = *r.CreatedBy
		}

		if r.DefinedTags != nil {
			schedule["defined_tags"] = tfresource.DefinedTagsToMap(r.DefinedTags)
		}

		if r.DisplayName != nil {
			schedule["display_name"] = *r.DisplayName
		}

		schedule["freeform_tags"] = r.FreeformTags

		if r.Id != nil {
			schedule["id"] = *r.Id
		}

		if r.ProjectId != nil {
			schedule["project_id"] = *r.ProjectId
		}

		schedule["state"] = r.LifecycleState

		if r.SystemTags != nil {
			schedule["system_tags"] = tfresource.SystemTagsToMap(r.SystemTags)
		}

		if r.TimeCreated != nil {
			schedule["time_created"] = r.TimeCreated.String()
		}

		if r.TimeUpdated != nil {
			schedule["time_updated"] = r.TimeUpdated.String()
		}

		if r.Trigger != nil {
			triggerArray := []interface{}{}
			if triggerMap := ScheduleTriggerToMap(&r.Trigger); triggerMap != nil {
				triggerArray = append(triggerArray, triggerMap)
			}
			schedule["trigger"] = triggerArray
		} else {
			schedule["trigger"] = nil
		}

		resources = append(resources, schedule)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = tfresource.ApplyFilters(f.(*schema.Set), resources, DatascienceSchedulesDataSource().Schema["schedules"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("schedules", resources); err != nil {
		return err
	}

	return nil
}
