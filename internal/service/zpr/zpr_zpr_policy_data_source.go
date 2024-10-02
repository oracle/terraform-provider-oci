// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package zpr

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_zpr "github.com/oracle/oci-go-sdk/v65/zpr"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func ZprZprPolicyDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["zpr_policy_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(ZprZprPolicyResource(), fieldMap, readSingularZprZprPolicy)
}

func readSingularZprZprPolicy(d *schema.ResourceData, m interface{}) error {
	sync := &ZprZprPolicyDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ZprClient()

	return tfresource.ReadResource(sync)
}

type ZprZprPolicyDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_zpr.ZprClient
	Res    *oci_zpr.GetZprPolicyResponse
}

func (s *ZprZprPolicyDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *ZprZprPolicyDataSourceCrud) Get() error {
	request := oci_zpr.GetZprPolicyRequest{}

	if zprPolicyId, ok := s.D.GetOkExists("zpr_policy_id"); ok {
		tmp := zprPolicyId.(string)
		request.ZprPolicyId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "zpr")

	response, err := s.Client.GetZprPolicy(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *ZprZprPolicyDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.Description != nil {
		s.D.Set("description", *s.Res.Description)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	if s.Res.Name != nil {
		s.D.Set("name", *s.Res.Name)
	}

	s.D.Set("state", s.Res.LifecycleState)

	s.D.Set("statements", s.Res.Statements)

	if s.Res.SystemTags != nil {
		s.D.Set("system_tags", tfresource.SystemTagsToMap(s.Res.SystemTags))
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	return nil
}
