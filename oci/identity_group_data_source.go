// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_identity "github.com/oracle/oci-go-sdk/v35/identity"
)

func init() {
	RegisterDatasource("oci_identity_group", IdentityGroupDataSource())
}

func IdentityGroupDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["group_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return GetSingularDataSourceItemSchema(IdentityGroupResource(), fieldMap, readSingularIdentityGroup)
}

func readSingularIdentityGroup(d *schema.ResourceData, m interface{}) error {
	sync := &IdentityGroupDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).identityClient()

	return ReadResource(sync)
}

type IdentityGroupDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_identity.IdentityClient
	Res    *oci_identity.GetGroupResponse
}

func (s *IdentityGroupDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *IdentityGroupDataSourceCrud) Get() error {
	request := oci_identity.GetGroupRequest{}

	if groupId, ok := s.D.GetOkExists("group_id"); ok {
		tmp := groupId.(string)
		request.GroupId = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(false, "identity")

	response, err := s.Client.GetGroup(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *IdentityGroupDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", definedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.Description != nil {
		s.D.Set("description", *s.Res.Description)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

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
