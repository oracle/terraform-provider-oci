// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_identity "github.com/oracle/oci-go-sdk/v45/identity"
)

func init() {
	RegisterDatasource("oci_identity_tag", IdentityTagDataSource())
}

func IdentityTagDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["tag_name"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	fieldMap["tag_namespace_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return GetSingularDataSourceItemSchema(IdentityTagResource(), fieldMap, readSingularIdentityTag)
}

func readSingularIdentityTag(d *schema.ResourceData, m interface{}) error {
	sync := &IdentityTagDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).identityClient()

	return ReadResource(sync)
}

type IdentityTagDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_identity.IdentityClient
	Res    *oci_identity.GetTagResponse
}

func (s *IdentityTagDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *IdentityTagDataSourceCrud) Get() error {
	request := oci_identity.GetTagRequest{}

	if tagName, ok := s.D.GetOkExists("tag_name"); ok {
		tmp := tagName.(string)
		request.TagName = &tmp
	}

	if tagNamespaceId, ok := s.D.GetOkExists("tag_namespace_id"); ok {
		tmp := tagNamespaceId.(string)
		request.TagNamespaceId = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(false, "identity")

	response, err := s.Client.GetTag(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *IdentityTagDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", definedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.Description != nil {
		s.D.Set("description", *s.Res.Description)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.IsCostTracking != nil {
		s.D.Set("is_cost_tracking", *s.Res.IsCostTracking)
	}

	if s.Res.IsRetired != nil {
		s.D.Set("is_retired", *s.Res.IsRetired)
	}

	if s.Res.Name != nil {
		s.D.Set("name", *s.Res.Name)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.Validator != nil {
		validatorArray := []interface{}{}
		if validatorMap := BaseTagDefinitionValidatorToMap(&s.Res.Validator); validatorMap != nil {
			validatorArray = append(validatorArray, validatorMap)
		}
		s.D.Set("validator", validatorArray)
	} else {
		s.D.Set("validator", nil)
	}

	return nil
}
