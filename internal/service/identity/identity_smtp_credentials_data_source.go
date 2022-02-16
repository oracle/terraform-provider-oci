// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package identity

import (
	"context"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_identity "github.com/oracle/oci-go-sdk/v58/identity"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"
)

func IdentitySmtpCredentialsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readIdentitySmtpCredentials,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"user_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"smtp_credentials": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     tfresource.GetDataSourceItemSchema(IdentitySmtpCredentialResource()),
			},
		},
	}
}

func readIdentitySmtpCredentials(d *schema.ResourceData, m interface{}) error {
	sync := &IdentitySmtpCredentialsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).IdentityClient()

	return tfresource.ReadResource(sync)
}

type IdentitySmtpCredentialsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_identity.IdentityClient
	Res    *oci_identity.ListSmtpCredentialsResponse
}

func (s *IdentitySmtpCredentialsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *IdentitySmtpCredentialsDataSourceCrud) Get() error {
	request := oci_identity.ListSmtpCredentialsRequest{}

	if userId, ok := s.D.GetOkExists("user_id"); ok {
		tmp := userId.(string)
		request.UserId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "identity")

	response, err := s.Client.ListSmtpCredentials(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *IdentitySmtpCredentialsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("IdentitySmtpCredentialsDataSource-", IdentitySmtpCredentialsDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		smtpCredential := map[string]interface{}{
			"user_id": *r.UserId,
		}

		if r.Description != nil {
			smtpCredential["description"] = *r.Description
		}

		if r.Id != nil {
			smtpCredential["id"] = *r.Id
		}

		if r.InactiveStatus != nil {
			smtpCredential["inactive_state"] = strconv.FormatInt(*r.InactiveStatus, 10)
		}

		smtpCredential["state"] = r.LifecycleState

		if r.TimeCreated != nil {
			smtpCredential["time_created"] = r.TimeCreated.String()
		}

		if r.TimeExpires != nil {
			smtpCredential["time_expires"] = r.TimeExpires.String()
		}

		if r.Username != nil {
			smtpCredential["username"] = *r.Username
		}

		resources = append(resources, smtpCredential)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = tfresource.ApplyFilters(f.(*schema.Set), resources, IdentitySmtpCredentialsDataSource().Schema["smtp_credentials"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("smtp_credentials", resources); err != nil {
		return err
	}

	return nil
}
