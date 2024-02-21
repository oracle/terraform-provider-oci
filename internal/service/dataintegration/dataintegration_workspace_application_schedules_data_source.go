// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package dataintegration

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_dataintegration "github.com/oracle/oci-go-sdk/v65/dataintegration"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DataintegrationWorkspaceApplicationSchedulesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readDataintegrationWorkspaceApplicationSchedules,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"application_key": {
				Type:     schema.TypeString,
				Required: true,
			},
			"identifier": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"key": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"type": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"workspace_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"schedule_summary_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem:     DataintegrationWorkspaceApplicationScheduleResource(),
						},
					},
				},
			},
		},
	}
}

func readDataintegrationWorkspaceApplicationSchedules(d *schema.ResourceData, m interface{}) error {
	sync := &DataintegrationWorkspaceApplicationSchedulesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataIntegrationClient()

	return tfresource.ReadResource(sync)
}

type DataintegrationWorkspaceApplicationSchedulesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_dataintegration.DataIntegrationClient
	Res    *oci_dataintegration.ListSchedulesResponse
}

func (s *DataintegrationWorkspaceApplicationSchedulesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DataintegrationWorkspaceApplicationSchedulesDataSourceCrud) Get() error {
	request := oci_dataintegration.ListSchedulesRequest{}

	if applicationKey, ok := s.D.GetOkExists("application_key"); ok {
		tmp := applicationKey.(string)
		request.ApplicationKey = &tmp
	}

	if identifier, ok := s.D.GetOkExists("identifier"); ok {
		interfaces := identifier.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("identifier") {
			request.Identifier = tmp
		}
	}

	if name, ok := s.D.GetOkExists("name"); ok {
		tmp := name.(string)
		request.Name = &tmp
	}

	if type_, ok := s.D.GetOkExists("type"); ok {
		interfaces := type_.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("type") {
			request.Type = tmp
		}
	}

	if workspaceId, ok := s.D.GetOkExists("workspace_id"); ok {
		tmp := workspaceId.(string)
		request.WorkspaceId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "dataintegration")

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

func (s *DataintegrationWorkspaceApplicationSchedulesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DataintegrationWorkspaceApplicationSchedulesDataSource-", DataintegrationWorkspaceApplicationSchedulesDataSource(), s.D))
	resources := []map[string]interface{}{}
	workspaceApplicationSchedule := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, ScheduleSummaryToMap(item))
	}
	workspaceApplicationSchedule["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, DataintegrationWorkspaceApplicationSchedulesDataSource().Schema["schedule_summary_collection"].Elem.(*schema.Resource).Schema)
		workspaceApplicationSchedule["items"] = items
	}

	resources = append(resources, workspaceApplicationSchedule)
	if err := s.D.Set("schedule_summary_collection", resources); err != nil {
		return err
	}

	return nil
}
