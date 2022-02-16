// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package identity

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_identity "github.com/oracle/oci-go-sdk/v58/identity"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"
)

func IdentityTagDefaultDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["tag_default_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(IdentityTagDefaultResource(), fieldMap, readSingularIdentityTagDefault)
}

func readSingularIdentityTagDefault(d *schema.ResourceData, m interface{}) error {
	sync := &IdentityTagDefaultDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).IdentityClient()

	return tfresource.ReadResource(sync)
}

type IdentityTagDefaultDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_identity.IdentityClient
	Res    *oci_identity.GetTagDefaultResponse
}

func (s *IdentityTagDefaultDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *IdentityTagDefaultDataSourceCrud) Get() error {
	request := oci_identity.GetTagDefaultRequest{}

	if tagDefaultId, ok := s.D.GetOkExists("tag_default_id"); ok {
		tmp := tagDefaultId.(string)
		request.TagDefaultId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "identity")

	response, err := s.Client.GetTagDefault(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *IdentityTagDefaultDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.IsRequired != nil {
		s.D.Set("is_required", *s.Res.IsRequired)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TagDefinitionId != nil {
		s.D.Set("tag_definition_id", *s.Res.TagDefinitionId)
	}

	if s.Res.TagDefinitionName != nil {
		s.D.Set("tag_definition_name", *s.Res.TagDefinitionName)
	}

	if s.Res.TagNamespaceId != nil {
		s.D.Set("tag_namespace_id", *s.Res.TagNamespaceId)
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.Value != nil {
		s.D.Set("value", *s.Res.Value)
	}

	return nil
}
