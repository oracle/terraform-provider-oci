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

func IdentityDbCredentialsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readIdentityDbCredentials,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"user_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"db_credentials": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     tfresource.GetDataSourceItemSchema(IdentityDbCredentialResource()),
			},
		},
	}
}

func readIdentityDbCredentials(d *schema.ResourceData, m interface{}) error {
	sync := &IdentityDbCredentialsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).IdentityClient()

	return tfresource.ReadResource(sync)
}

type IdentityDbCredentialsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_identity.IdentityClient
	Res    *oci_identity.ListDbCredentialsResponse
}

func (s *IdentityDbCredentialsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *IdentityDbCredentialsDataSourceCrud) Get() error {
	request := oci_identity.ListDbCredentialsRequest{}

	if name, ok := s.D.GetOkExists("name"); ok {
		tmp := name.(string)
		request.Name = &tmp
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_identity.DbCredentialLifecycleStateEnum(state.(string))
	}

	if userId, ok := s.D.GetOkExists("user_id"); ok {
		tmp := userId.(string)
		request.UserId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "identity")

	response, err := s.Client.ListDbCredentials(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListDbCredentials(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *IdentityDbCredentialsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("IdentityDbCredentialsDataSource-", IdentityDbCredentialsDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		dbCredential := map[string]interface{}{
			"user_id": *r.UserId,
		}

		if r.Description != nil {
			dbCredential["description"] = *r.Description
		}

		if r.Id != nil {
			dbCredential["id"] = *r.Id
		}

		dbCredential["state"] = r.LifecycleState

		if r.TimeCreated != nil {
			dbCredential["time_created"] = r.TimeCreated.String()
		}

		if r.TimeExpires != nil {
			dbCredential["time_expires"] = r.TimeExpires.String()
		}

		resources = append(resources, dbCredential)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = tfresource.ApplyFilters(f.(*schema.Set), resources, IdentityDbCredentialsDataSource().Schema["db_credentials"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("db_credentials", resources); err != nil {
		return err
	}

	return nil
}
