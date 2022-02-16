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

func EmailSuppressionDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["suppression_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(EmailSuppressionResource(), fieldMap, readSingularEmailSuppression)
}

func readSingularEmailSuppression(d *schema.ResourceData, m interface{}) error {
	sync := &EmailSuppressionDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).EmailClient()

	return tfresource.ReadResource(sync)
}

type EmailSuppressionDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_email.EmailClient
	Res    *oci_email.GetSuppressionResponse
}

func (s *EmailSuppressionDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *EmailSuppressionDataSourceCrud) Get() error {
	request := oci_email.GetSuppressionRequest{}

	if suppressionId, ok := s.D.GetOkExists("suppression_id"); ok {
		tmp := suppressionId.(string)
		request.SuppressionId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "email")

	response, err := s.Client.GetSuppression(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *EmailSuppressionDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.EmailAddress != nil {
		s.D.Set("email_address", *s.Res.EmailAddress)
	}

	if s.Res.ErrorDetail != nil {
		s.D.Set("error_detail", *s.Res.ErrorDetail)
	}

	if s.Res.ErrorSource != nil {
		s.D.Set("error_source", *s.Res.ErrorSource)
	}

	if s.Res.MessageId != nil {
		s.D.Set("message_id", *s.Res.MessageId)
	}

	s.D.Set("reason", s.Res.Reason)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeLastSuppressed != nil {
		s.D.Set("time_last_suppressed", s.Res.TimeLastSuppressed.String())
	}

	return nil
}
