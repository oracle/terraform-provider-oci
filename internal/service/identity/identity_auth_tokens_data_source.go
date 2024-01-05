// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package identity

import (
	"context"
	"strconv"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_identity "github.com/oracle/oci-go-sdk/v65/identity"
)

func IdentityAuthTokensDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readIdentityAuthTokens,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"user_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"tokens": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     tfresource.GetDataSourceItemSchema(IdentityAuthTokenResource()),
			},
		},
	}
}

func readIdentityAuthTokens(d *schema.ResourceData, m interface{}) error {
	sync := &IdentityAuthTokensDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).IdentityClient()

	return tfresource.ReadResource(sync)
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

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "identity")

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

	s.D.SetId(tfresource.GenerateDataSourceHashID("IdentityAuthTokensDataSource-", IdentityAuthTokensDataSource(), s.D))
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
		resources = tfresource.ApplyFilters(f.(*schema.Set), resources, IdentityAuthTokensDataSource().Schema["tokens"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("tokens", resources); err != nil {
		return err
	}

	return nil
}
