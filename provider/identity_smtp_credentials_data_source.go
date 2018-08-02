// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"context"
	"strconv"

	"github.com/hashicorp/terraform/helper/schema"
	oci_identity "github.com/oracle/oci-go-sdk/identity"
)

func SmtpCredentialsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSmtpCredentials,
		Schema: map[string]*schema.Schema{
			"filter": dataSourceFiltersSchema(),
			"user_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"smtp_credentials": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     GetDataSourceItemSchema(SmtpCredentialResource()),
			},
		},
	}
}

func readSmtpCredentials(d *schema.ResourceData, m interface{}) error {
	sync := &SmtpCredentialsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).identityClient

	return ReadResource(sync)
}

type SmtpCredentialsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_identity.IdentityClient
	Res    *oci_identity.ListSmtpCredentialsResponse
}

func (s *SmtpCredentialsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *SmtpCredentialsDataSourceCrud) Get() error {
	request := oci_identity.ListSmtpCredentialsRequest{}

	if userId, ok := s.D.GetOkExists("user_id"); ok {
		tmp := userId.(string)
		request.UserId = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(false, "identity")

	response, err := s.Client.ListSmtpCredentials(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *SmtpCredentialsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(GenerateDataSourceID())
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
		resources = ApplyFilters(f.(*schema.Set), resources, SmtpCredentialsDataSource().Schema["smtp_credentials"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("smtp_credentials", resources); err != nil {
		return err
	}

	return nil
}
