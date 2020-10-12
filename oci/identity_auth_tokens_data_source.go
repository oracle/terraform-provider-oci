// Copyright (c) 2017, 2020, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"
	"strconv"

	"github.com/hashicorp/terraform/helper/schema"
	oci_identity "github.com/oracle/oci-go-sdk/v27/identity"
)

func init() {
	RegisterDatasource("oci_identity_auth_tokens", IdentityAuthTokensDataSource())
}

func IdentityAuthTokensDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readIdentityAuthTokens,
		Schema: map[string]*schema.Schema{
			"filter": dataSourceFiltersSchema(),
			"user_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"tokens": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     GetDataSourceItemSchema(IdentityAuthTokenResource()),
			},
		},
	}
}

func readIdentityAuthTokens(d *schema.ResourceData, m interface{}) error {
	sync := &IdentityAuthTokensDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).identityClient()

	return ReadResource(sync)
}

type IdentityAuthTokensDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_identity.IdentityClient
	Res    *oci_identity.ListAuthTokensResponse
}

func (s *IdentityAuthTokensDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *IdentityAuthTokensDataSourceCrud) Get() error {
	request := oci_identity.ListAuthTokensRequest{}

	if userId, ok := s.D.GetOkExists("user_id"); ok {
		tmp := userId.(string)
		request.UserId = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(false, "identity")

	response, err := s.Client.ListAuthTokens(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *IdentityAuthTokensDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(GenerateDataSourceHashID("IdentityAuthTokensDataSource-", IdentityAuthTokensDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		authToken := map[string]interface{}{
			"user_id": *r.UserId,
		}

		if r.Description != nil {
			authToken["description"] = *r.Description
		}

		if r.Id != nil {
			authToken["id"] = *r.Id
		}

		if r.InactiveStatus != nil {
			authToken["inactive_state"] = strconv.FormatInt(*r.InactiveStatus, 10)
		}

		authToken["state"] = r.LifecycleState

		if r.TimeCreated != nil {
			authToken["time_created"] = r.TimeCreated.String()
		}

		if r.TimeExpires != nil {
			authToken["time_expires"] = r.TimeExpires.String()
		}

		if r.Token != nil {
			authToken["token"] = *r.Token
		}

		resources = append(resources, authToken)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = ApplyFilters(f.(*schema.Set), resources, IdentityAuthTokensDataSource().Schema["tokens"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("tokens", resources); err != nil {
		return err
	}

	return nil
}
