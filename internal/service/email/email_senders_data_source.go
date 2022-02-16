// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package email

import (
	"context"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_email "github.com/oracle/oci-go-sdk/v58/email"
)

func EmailSendersDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readEmailSenders,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"domain": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"email_address": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"senders": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     tfresource.GetDataSourceItemSchema(EmailSenderResource()),
			},
		},
	}
}

func readEmailSenders(d *schema.ResourceData, m interface{}) error {
	sync := &EmailSendersDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).EmailClient()

	return tfresource.ReadResource(sync)
}

type EmailSendersDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_email.EmailClient
	Res    *oci_email.ListSendersResponse
}

func (s *EmailSendersDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *EmailSendersDataSourceCrud) Get() error {
	request := oci_email.ListSendersRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if domain, ok := s.D.GetOkExists("domain"); ok {
		tmp := domain.(string)
		request.Domain = &tmp
	}

	if emailAddress, ok := s.D.GetOkExists("email_address"); ok {
		tmp := emailAddress.(string)
		request.EmailAddress = &tmp
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_email.SenderLifecycleStateEnum(state.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "email")

	response, err := s.Client.ListSenders(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListSenders(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *EmailSendersDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("EmailSendersDataSource-", EmailSendersDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		sender := map[string]interface{}{
			"compartment_id": *r.CompartmentId,
		}

		if r.DefinedTags != nil {
			sender["defined_tags"] = tfresource.DefinedTagsToMap(r.DefinedTags)
		}

		if r.EmailAddress != nil {
			sender["email_address"] = *r.EmailAddress
		}

		sender["freeform_tags"] = r.FreeformTags

		if r.Id != nil {
			sender["id"] = *r.Id
		}

		sender["state"] = r.LifecycleState

		if r.TimeCreated != nil {
			sender["time_created"] = r.TimeCreated.String()
		}

		resources = append(resources, sender)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = tfresource.ApplyFilters(f.(*schema.Set), resources, EmailSendersDataSource().Schema["senders"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("senders", resources); err != nil {
		return err
	}

	return nil
}
