// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package identity

import (
	"context"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_identity "github.com/oracle/oci-go-sdk/v58/identity"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"
)

func IdentityUserGroupMembershipsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readIdentityUserGroupMemberships,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
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
				Elem:     tfresource.GetDataSourceItemSchema(IdentityUserGroupMembershipResource()),
			},
		},
	}
}

func readIdentityUserGroupMemberships(d *schema.ResourceData, m interface{}) error {
	sync := &IdentityUserGroupMembershipsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).IdentityClient()

	return tfresource.ReadResource(sync)
}

type IdentityUserGroupMembershipsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_identity.IdentityClient
	Res    *oci_identity.ListUserGroupMembershipsResponse
}

func (s *IdentityUserGroupMembershipsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *IdentityUserGroupMembershipsDataSourceCrud) Get() error {
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

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "identity")

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

func (s *IdentityUserGroupMembershipsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("IdentityUserGroupMembershipsDataSource-", IdentityUserGroupMembershipsDataSource(), s.D))
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
			userGroupMembership["inactive_state"] = strconv.FormatInt(*r.InactiveStatus, 10)
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
		resources = tfresource.ApplyFilters(f.(*schema.Set), resources, IdentityUserGroupMembershipsDataSource().Schema["memberships"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("memberships", resources); err != nil {
		return err
	}

	return nil
}
