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

func IdentityUsersDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readIdentityUsers,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"external_identifier": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"identity_provider_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"users": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     tfresource.GetDataSourceItemSchema(IdentityUserResource()),
			},
		},
	}
}

func readIdentityUsers(d *schema.ResourceData, m interface{}) error {
	sync := &IdentityUsersDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).IdentityClient()

	return tfresource.ReadResource(sync)
}

type IdentityUsersDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_identity.IdentityClient
	Res    *oci_identity.ListUsersResponse
}

func (s *IdentityUsersDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *IdentityUsersDataSourceCrud) Get() error {
	request := oci_identity.ListUsersRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if externalIdentifier, ok := s.D.GetOkExists("external_identifier"); ok {
		tmp := externalIdentifier.(string)
		request.ExternalIdentifier = &tmp
	}

	if identityProviderId, ok := s.D.GetOkExists("identity_provider_id"); ok {
		tmp := identityProviderId.(string)
		request.IdentityProviderId = &tmp
	}

	if name, ok := s.D.GetOkExists("name"); ok {
		tmp := name.(string)
		request.Name = &tmp
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_identity.UserLifecycleStateEnum(state.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "identity")

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

func (s *IdentityUsersDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("IdentityUsersDataSource-", IdentityUsersDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		user := map[string]interface{}{
			"compartment_id": *r.CompartmentId,
		}

		if r.Capabilities != nil {
			user["capabilities"] = []interface{}{UserCapabilitiesToMap(r.Capabilities)}
		} else {
			user["capabilities"] = nil
		}

		if r.DbUserName != nil {
			user["db_user_name"] = *r.DbUserName
		}

		if r.DefinedTags != nil {
			user["defined_tags"] = tfresource.DefinedTagsToMap(r.DefinedTags)
		}

		if r.Description != nil {
			user["description"] = *r.Description
		}

		if r.Email != nil {
			user["email"] = *r.Email
		}

		if r.EmailVerified != nil {
			user["email_verified"] = *r.EmailVerified
		}

		if r.ExternalIdentifier != nil {
			user["external_identifier"] = *r.ExternalIdentifier
		}

		user["freeform_tags"] = r.FreeformTags

		if r.Id != nil {
			user["id"] = *r.Id
		}

		if r.IdentityProviderId != nil {
			user["identity_provider_id"] = *r.IdentityProviderId
		}

		if r.InactiveStatus != nil {
			user["inactive_state"] = strconv.FormatInt(*r.InactiveStatus, 10)
		}

		if r.LastSuccessfulLoginTime != nil {
			user["last_successful_login_time"] = r.LastSuccessfulLoginTime.String()
		}

		if r.Name != nil {
			user["name"] = *r.Name
		}

		if r.PreviousSuccessfulLoginTime != nil {
			user["previous_successful_login_time"] = r.PreviousSuccessfulLoginTime.String()
		}

		user["state"] = r.LifecycleState

		if r.TimeCreated != nil {
			user["time_created"] = r.TimeCreated.String()
		}

		resources = append(resources, user)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = tfresource.ApplyFilters(f.(*schema.Set), resources, IdentityUsersDataSource().Schema["users"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("users", resources); err != nil {
		return err
	}

	return nil
}
