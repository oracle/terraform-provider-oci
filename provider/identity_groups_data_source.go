// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"context"

	"github.com/hashicorp/terraform/helper/schema"
	oci_identity "github.com/oracle/oci-go-sdk/identity"

	"github.com/oracle/terraform-provider-oci/crud"
)

func GroupsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readGroups,
		Schema: map[string]*schema.Schema{
			"filter": dataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"groups": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     GroupResource(),
			},
		},
	}
}

func readGroups(d *schema.ResourceData, m interface{}) error {
	sync := &GroupsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).identityClient

	return crud.ReadResource(sync)
}

type GroupsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_identity.IdentityClient
	Res    *oci_identity.ListGroupsResponse
}

func (s *GroupsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *GroupsDataSourceCrud) Get() error {
	request := oci_identity.ListGroupsRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(false, "identity")

	response, err := s.Client.ListGroups(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListGroups(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *GroupsDataSourceCrud) SetData() {
	if s.Res == nil {
		return
	}

	s.D.SetId(crud.GenerateDataSourceID())
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		group := map[string]interface{}{
			"compartment_id": *r.CompartmentId,
		}

		if r.DefinedTags != nil {
			group["defined_tags"] = definedTagsToMap(r.DefinedTags)
		}

		if r.Description != nil {
			group["description"] = *r.Description
		}

		group["freeform_tags"] = r.FreeformTags

		if r.Id != nil {
			group["id"] = *r.Id
		}

		if r.InactiveStatus != nil {
			group["inactive_state"] = *r.InactiveStatus
		}

		if r.Name != nil {
			group["name"] = *r.Name
		}

		group["state"] = r.LifecycleState

		if r.TimeCreated != nil {
			group["time_created"] = r.TimeCreated.String()
		}

		resources = append(resources, group)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = ApplyFilters(f.(*schema.Set), resources, GroupsDataSource().Schema["groups"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("groups", resources); err != nil {
		panic(err)
	}

	return
}
