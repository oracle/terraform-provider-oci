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

func IdentityUserDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["user_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(IdentityUserResource(), fieldMap, readSingularIdentityUser)
}

func readSingularIdentityUser(d *schema.ResourceData, m interface{}) error {
	sync := &IdentityUserDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).IdentityClient()

	return tfresource.ReadResource(sync)
}

type IdentityUserDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_identity.IdentityClient
	Res    *oci_identity.GetUserResponse
}

func (s *IdentityUserDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *IdentityUserDataSourceCrud) Get() error {
	request := oci_identity.GetUserRequest{}

	if userId, ok := s.D.GetOkExists("user_id"); ok {
		tmp := userId.(string)
		request.UserId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "identity")

	response, err := s.Client.GetUser(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *IdentityUserDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.Capabilities != nil {
		s.D.Set("capabilities", []interface{}{UserCapabilitiesToMap(s.Res.Capabilities)})
	} else {
		s.D.Set("capabilities", nil)
	}

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DbUserName != nil {
		s.D.Set("db_user_name", *s.Res.DbUserName)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.Description != nil {
		s.D.Set("description", *s.Res.Description)
	}

	if s.Res.Email != nil {
		s.D.Set("email", *s.Res.Email)
	}

	if s.Res.EmailVerified != nil {
		s.D.Set("email_verified", *s.Res.EmailVerified)
	}

	if s.Res.ExternalIdentifier != nil {
		s.D.Set("external_identifier", *s.Res.ExternalIdentifier)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.IdentityProviderId != nil {
		s.D.Set("identity_provider_id", *s.Res.IdentityProviderId)
	}

	if s.Res.InactiveStatus != nil {
		s.D.Set("inactive_state", strconv.FormatInt(*s.Res.InactiveStatus, 10))
	}

	if s.Res.LastSuccessfulLoginTime != nil {
		s.D.Set("last_successful_login_time", s.Res.LastSuccessfulLoginTime.String())
	}

	if s.Res.Name != nil {
		s.D.Set("name", *s.Res.Name)
	}

	if s.Res.PreviousSuccessfulLoginTime != nil {
		s.D.Set("previous_successful_login_time", s.Res.PreviousSuccessfulLoginTime.String())
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	return nil
}
