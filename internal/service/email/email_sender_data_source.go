// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package email

import (
	"context"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_email "github.com/oracle/oci-go-sdk/v65/email"
)

func EmailSenderDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["sender_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(EmailSenderResource(), fieldMap, readSingularEmailSender)
}

func readSingularEmailSender(d *schema.ResourceData, m interface{}) error {
	sync := &EmailSenderDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).EmailClient()

	return tfresource.ReadResource(sync)
}

type EmailSenderDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_email.EmailClient
	Res    *oci_email.GetSenderResponse
}

func (s *EmailSenderDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *EmailSenderDataSourceCrud) Get() error {
	request := oci_email.GetSenderRequest{}

	if senderId, ok := s.D.GetOkExists("sender_id"); ok {
		tmp := senderId.(string)
		request.SenderId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "email")

	response, err := s.Client.GetSender(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *EmailSenderDataSourceCrud) SetData() error {
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

	if s.Res.EmailAddress != nil {
		s.D.Set("email_address", *s.Res.EmailAddress)
	}

	if s.Res.EmailDomainId != nil {
		s.D.Set("email_domain_id", *s.Res.EmailDomainId)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.IsSpf != nil {
		s.D.Set("is_spf", *s.Res.IsSpf)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.SystemTags != nil {
		s.D.Set("system_tags", tfresource.SystemTagsToMap(s.Res.SystemTags))
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	return nil
}
