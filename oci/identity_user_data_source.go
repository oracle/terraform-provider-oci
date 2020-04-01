// Copyright (c) 2017, 2019, Oracle and/or its affiliates. All rights reserved.

package oci

import (
	"context"
	"strconv"

	"github.com/hashicorp/terraform/helper/schema"
	oci_identity "github.com/oracle/oci-go-sdk/identity"
)

func init() {
	RegisterDatasource("oci_identity_user", IdentityUserDataSource())
}

func IdentityUserDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["user_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return GetSingularDataSourceItemSchema(IdentityUserResource(), fieldMap, readSingularIdentityUser)
}

func readSingularIdentityUser(d *schema.ResourceData, m interface{}) error {
	sync := &IdentityUserDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).identityClient

	return ReadResource(sync)
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

	request.RequestMetadata.RetryPolicy = getRetryPolicy(false, "identity")

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

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", definedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.Description != nil {
		s.D.Set("description", *s.Res.Description)
	}

	if s.Res.Email != nil {
		s.D.Set("email", *s.Res.Email)
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

	if s.Res.Name != nil {
		s.D.Set("name", *s.Res.Name)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	return nil
}
