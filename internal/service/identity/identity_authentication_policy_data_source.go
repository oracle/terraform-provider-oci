// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package identity

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_identity "github.com/oracle/oci-go-sdk/v56/identity"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"
)

func IdentityAuthenticationPolicyDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["compartment_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(IdentityAuthenticationPolicyResource(), fieldMap, readSingularIdentityAuthenticationPolicy)
}

func readSingularIdentityAuthenticationPolicy(d *schema.ResourceData, m interface{}) error {
	sync := &IdentityAuthenticationPolicyDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).IdentityClient()

	return tfresource.ReadResource(sync)
}

type IdentityAuthenticationPolicyDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_identity.IdentityClient
	Res    *oci_identity.GetAuthenticationPolicyResponse
}

func (s *IdentityAuthenticationPolicyDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *IdentityAuthenticationPolicyDataSourceCrud) Get() error {
	request := oci_identity.GetAuthenticationPolicyRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "identity")

	response, err := s.Client.GetAuthenticationPolicy(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *IdentityAuthenticationPolicyDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("IdentityAuthenticationPolicyDataSource-", IdentityAuthenticationPolicyDataSource(), s.D))

	if s.Res.NetworkPolicy != nil {
		s.D.Set("network_policy", []interface{}{NetworkPolicyToMap(s.Res.NetworkPolicy)})
	} else {
		s.D.Set("network_policy", nil)
	}

	if s.Res.PasswordPolicy != nil {
		s.D.Set("password_policy", []interface{}{PasswordPolicyToMap(s.Res.PasswordPolicy)})
	} else {
		s.D.Set("password_policy", nil)
	}

	return nil
}
