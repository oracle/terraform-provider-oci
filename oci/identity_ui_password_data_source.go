// Copyright (c) 2017, 2019, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"context"

	"github.com/hashicorp/terraform/helper/schema"
	oci_identity "github.com/oracle/oci-go-sdk/identity"
)

func IdentityUiPasswordDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["user_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return GetSingularDataSourceItemSchema(IdentityUiPasswordResource(), fieldMap, readSingularIdentityUiPassword)
}

func readSingularIdentityUiPassword(d *schema.ResourceData, m interface{}) error {
	sync := &IdentityUiPasswordDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).identityClient

	return ReadResource(sync)
}

type IdentityUiPasswordDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_identity.IdentityClient
	Res    *oci_identity.GetUserUIPasswordInformationResponse
}

func (s *IdentityUiPasswordDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *IdentityUiPasswordDataSourceCrud) Get() error {
	request := oci_identity.GetUserUIPasswordInformationRequest{}

	if userId, ok := s.D.GetOkExists("user_id"); ok {
		tmp := userId.(string)
		request.UserId = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(false, "identity")

	response, err := s.Client.GetUserUIPasswordInformation(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *IdentityUiPasswordDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(GenerateDataSourceID())

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	return nil
}
