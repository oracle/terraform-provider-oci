// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"context"

	"github.com/hashicorp/terraform/helper/schema"
	oci_identity "github.com/oracle/oci-go-sdk/identity"

	"github.com/oracle/terraform-provider-oci/crud"
)

func UserGroupMembershipsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readUserGroupMemberships,
		Schema: map[string]*schema.Schema{
			"filter": dataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"group_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"user_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"memberships": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     UserGroupMembershipResource(),
			},
		},
	}
}

func readUserGroupMemberships(d *schema.ResourceData, m interface{}) error {
	sync := &UserGroupMembershipsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).identityClient

	return crud.ReadResource(sync)
}

type UserGroupMembershipsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_identity.IdentityClient
	Res    *oci_identity.ListUserGroupMembershipsResponse
}

func (s *UserGroupMembershipsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *UserGroupMembershipsDataSourceCrud) Get() error {
	request := oci_identity.ListUserGroupMembershipsRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if groupId, ok := s.D.GetOkExists("group_id"); ok {
		tmp := groupId.(string)
		request.GroupId = &tmp
	}

	if userId, ok := s.D.GetOkExists("user_id"); ok {
		tmp := userId.(string)
		request.UserId = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(false, "identity")

	response, err := s.Client.ListUserGroupMemberships(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListUserGroupMemberships(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *UserGroupMembershipsDataSourceCrud) SetData() {
	if s.Res == nil {
		return
	}

	s.D.SetId(crud.GenerateDataSourceID())
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		userGroupMembership := map[string]interface{}{
			"compartment_id": *r.CompartmentId,
		}

		if r.GroupId != nil {
			userGroupMembership["group_id"] = *r.GroupId
		}

		if r.Id != nil {
			userGroupMembership["id"] = *r.Id
		}

		if r.InactiveStatus != nil {
			userGroupMembership["inactive_state"] = *r.InactiveStatus
		}

		userGroupMembership["state"] = r.LifecycleState

		if r.TimeCreated != nil {
			userGroupMembership["time_created"] = r.TimeCreated.String()
		}

		if r.UserId != nil {
			userGroupMembership["user_id"] = *r.UserId
		}

		resources = append(resources, userGroupMembership)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = ApplyFilters(f.(*schema.Set), resources, UserGroupMembershipsDataSource().Schema["memberships"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("memberships", resources); err != nil {
		panic(err)
	}

	return
}
