// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package datascience

import (
	"context"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_datascience "github.com/oracle/oci-go-sdk/v58/datascience"
)

func DatascienceProjectsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readDatascienceProjects,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"created_by": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"projects": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     tfresource.GetDataSourceItemSchema(DatascienceProjectResource()),
			},
		},
	}
}

func readDatascienceProjects(d *schema.ResourceData, m interface{}) error {
	sync := &DatascienceProjectsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataScienceClient()

	return tfresource.ReadResource(sync)
}

type DatascienceProjectsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_datascience.DataScienceClient
	Res    *oci_datascience.ListProjectsResponse
}

func (s *DatascienceProjectsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DatascienceProjectsDataSourceCrud) Get() error {
	request := oci_datascience.ListProjectsRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if createdBy, ok := s.D.GetOkExists("created_by"); ok {
		tmp := createdBy.(string)
		request.CreatedBy = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if id, ok := s.D.GetOkExists("id"); ok {
		tmp := id.(string)
		request.Id = &tmp
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_datascience.ListProjectsLifecycleStateEnum(state.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "datascience")

	response, err := s.Client.ListProjects(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListProjects(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *DatascienceProjectsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DatascienceProjectsDataSource-", DatascienceProjectsDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		project := map[string]interface{}{
			"compartment_id": *r.CompartmentId,
		}

		if r.CreatedBy != nil {
			project["created_by"] = *r.CreatedBy
		}

		if r.DefinedTags != nil {
			project["defined_tags"] = tfresource.DefinedTagsToMap(r.DefinedTags)
		}

		if r.Description != nil {
			project["description"] = *r.Description
		}

		if r.DisplayName != nil {
			project["display_name"] = *r.DisplayName
		}

		project["freeform_tags"] = r.FreeformTags

		if r.Id != nil {
			project["id"] = *r.Id
		}

		project["state"] = r.LifecycleState

		if r.TimeCreated != nil {
			project["time_created"] = r.TimeCreated.String()
		}

		resources = append(resources, project)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = tfresource.ApplyFilters(f.(*schema.Set), resources, DatascienceProjectsDataSource().Schema["projects"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("projects", resources); err != nil {
		return err
	}

	return nil
}
