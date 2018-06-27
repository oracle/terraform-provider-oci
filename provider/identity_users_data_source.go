// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"context"

	"github.com/hashicorp/terraform/helper/schema"
	oci_identity "github.com/oracle/oci-go-sdk/identity"

	"github.com/oracle/terraform-provider-oci/crud"
)

func UsersDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readUsers,
		Schema: map[string]*schema.Schema{
			"filter": dataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"users": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     UserResource(),
			},
		},
	}
}

func readUsers(d *schema.ResourceData, m interface{}) error {
	sync := &UsersDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).identityClient

	return crud.ReadResource(sync)
}

type UsersDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_identity.IdentityClient
	Res    *oci_identity.ListUsersResponse
}

func (s *UsersDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *UsersDataSourceCrud) Get() error {
	request := oci_identity.ListUsersRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(false, "identity")

	response, err := s.Client.ListUsers(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListUsers(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *UsersDataSourceCrud) SetData() {
	if s.Res == nil {
		return
	}

	s.D.SetId(crud.GenerateDataSourceID())
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		user := map[string]interface{}{
			"compartment_id": *r.CompartmentId,
		}

		if r.DefinedTags != nil {
			user["defined_tags"] = definedTagsToMap(r.DefinedTags)
		}

		if r.Description != nil {
			user["description"] = *r.Description
		}

		user["freeform_tags"] = r.FreeformTags

		if r.Id != nil {
			user["id"] = *r.Id
		}

		if r.InactiveStatus != nil {
			user["inactive_state"] = *r.InactiveStatus
		}

		if r.Name != nil {
			user["name"] = *r.Name
		}

		user["state"] = r.LifecycleState

		if r.TimeCreated != nil {
			user["time_created"] = r.TimeCreated.String()
		}

		resources = append(resources, user)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = ApplyFilters(f.(*schema.Set), resources, UsersDataSource().Schema["users"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("users", resources); err != nil {
		panic(err)
	}

	return
}
