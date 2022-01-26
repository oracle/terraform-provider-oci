// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package osmanagement

import (
	"context"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_osmanagement "github.com/oracle/oci-go-sdk/v56/osmanagement"
)

func OsmanagementSoftwareSourcesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readOsmanagementSoftwareSources,
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
			"software_sources": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     tfresource.GetDataSourceItemSchema(OsmanagementSoftwareSourceResource()),
			},
		},
	}
}

func readOsmanagementSoftwareSources(d *schema.ResourceData, m interface{}) error {
	sync := &OsmanagementSoftwareSourcesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).OsManagementClient()

	return tfresource.ReadResource(sync)
}

type OsmanagementSoftwareSourcesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_osmanagement.OsManagementClient
	Res    *oci_osmanagement.ListSoftwareSourcesResponse
}

func (s *OsmanagementSoftwareSourcesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *OsmanagementSoftwareSourcesDataSourceCrud) Get() error {
	request := oci_osmanagement.ListSoftwareSourcesRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_osmanagement.ListSoftwareSourcesLifecycleStateEnum(state.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "osmanagement")

	response, err := s.Client.ListSoftwareSources(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListSoftwareSources(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *OsmanagementSoftwareSourcesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("OsmanagementSoftwareSourcesDataSource-", OsmanagementSoftwareSourcesDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		softwareSource := map[string]interface{}{
			"compartment_id": *r.CompartmentId,
		}

		if r.DefinedTags != nil {
			softwareSource["defined_tags"] = tfresource.DefinedTagsToMap(r.DefinedTags)
		}

		if r.Description != nil {
			softwareSource["description"] = *r.Description
		}

		if r.DisplayName != nil {
			softwareSource["display_name"] = *r.DisplayName
		}

		softwareSource["freeform_tags"] = r.FreeformTags

		if r.Id != nil {
			softwareSource["id"] = *r.Id
		}

		if r.Packages != nil {
			softwareSource["packages"] = *r.Packages
		}

		if r.ParentId != nil {
			softwareSource["parent_id"] = *r.ParentId
		}

		if r.ParentName != nil {
			softwareSource["parent_name"] = *r.ParentName
		}

		if r.RepoType != nil {
			softwareSource["repo_type"] = *r.RepoType
		}

		softwareSource["state"] = r.LifecycleState

		softwareSource["status"] = r.Status

		resources = append(resources, softwareSource)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = tfresource.ApplyFilters(f.(*schema.Set), resources, OsmanagementSoftwareSourcesDataSource().Schema["software_sources"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("software_sources", resources); err != nil {
		return err
	}

	return nil
}
