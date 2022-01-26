// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package dataintegration

import (
	"context"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_dataintegration "github.com/oracle/oci-go-sdk/v56/dataintegration"
)

func DataintegrationWorkspacesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readDataintegrationWorkspaces,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"workspaces": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     tfresource.GetDataSourceItemSchema(DataintegrationWorkspaceResource()),
			},
		},
	}
}

func readDataintegrationWorkspaces(d *schema.ResourceData, m interface{}) error {
	sync := &DataintegrationWorkspacesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataIntegrationClient()

	return tfresource.ReadResource(sync)
}

type DataintegrationWorkspacesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_dataintegration.DataIntegrationClient
	Res    *oci_dataintegration.ListWorkspacesResponse
}

func (s *DataintegrationWorkspacesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DataintegrationWorkspacesDataSourceCrud) Get() error {
	request := oci_dataintegration.ListWorkspacesRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if name, ok := s.D.GetOkExists("name"); ok {
		tmp := name.(string)
		request.Name = &tmp
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_dataintegration.WorkspaceLifecycleStateEnum(state.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "dataintegration")

	response, err := s.Client.ListWorkspaces(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListWorkspaces(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *DataintegrationWorkspacesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DataintegrationWorkspacesDataSource-", DataintegrationWorkspacesDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		workspace := map[string]interface{}{
			"compartment_id": *r.CompartmentId,
		}

		if r.DefinedTags != nil {
			workspace["defined_tags"] = tfresource.DefinedTagsToMap(r.DefinedTags)
		}

		if r.Description != nil {
			workspace["description"] = *r.Description
		}

		if r.DisplayName != nil {
			workspace["display_name"] = *r.DisplayName
		}

		workspace["freeform_tags"] = r.FreeformTags

		if r.Id != nil {
			workspace["id"] = *r.Id
		}

		workspace["state"] = r.LifecycleState

		if r.StateMessage != nil {
			workspace["state_message"] = *r.StateMessage
		}

		if r.TimeCreated != nil {
			workspace["time_created"] = r.TimeCreated.String()
		}

		if r.TimeUpdated != nil {
			workspace["time_updated"] = r.TimeUpdated.String()
		}

		resources = append(resources, workspace)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = tfresource.ApplyFilters(f.(*schema.Set), resources, DataintegrationWorkspacesDataSource().Schema["workspaces"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("workspaces", resources); err != nil {
		return err
	}

	return nil
}
