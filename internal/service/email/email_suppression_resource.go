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

func EmailSuppressionResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createEmailSuppression,
		Read:     readEmailSuppression,
		Delete:   deleteEmailSuppression,
		Schema: map[string]*schema.Schema{
			// Required
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"email_address": {
				Type:             schema.TypeString,
				Required:         true,
				ForceNew:         true,
				DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
			},

			// Optional

			// Computed
			"error_detail": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"error_source": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"message_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"reason": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_created": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_last_suppressed": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func createEmailSuppression(d *schema.ResourceData, m interface{}) error {
	sync := &EmailSuppressionResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).EmailClient()

	return tfresource.CreateResource(d, sync)
}

func readEmailSuppression(d *schema.ResourceData, m interface{}) error {
	sync := &EmailSuppressionResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).EmailClient()

	return tfresource.ReadResource(sync)
}

func deleteEmailSuppression(d *schema.ResourceData, m interface{}) error {
	sync := &EmailSuppressionResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).EmailClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type EmailSuppressionResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_email.EmailClient
	Res                    *oci_email.Suppression
	DisableNotFoundRetries bool
}

func (s *EmailSuppressionResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *EmailSuppressionResourceCrud) Create() error {
	request := oci_email.CreateSuppressionRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if emailAddress, ok := s.D.GetOkExists("email_address"); ok {
		tmp := emailAddress.(string)
		request.EmailAddress = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "email")

	response, err := s.Client.CreateSuppression(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Suppression
	return nil
}

func (s *EmailSuppressionResourceCrud) Get() error {
	request := oci_email.GetSuppressionRequest{}

	tmp := s.D.Id()
	request.SuppressionId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "email")

	response, err := s.Client.GetSuppression(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Suppression
	return nil
}

func (s *EmailSuppressionResourceCrud) Delete() error {
	request := oci_email.DeleteSuppressionRequest{}

	tmp := s.D.Id()
	request.SuppressionId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "email")

	_, err := s.Client.DeleteSuppression(context.Background(), request)
	return err
}

func (s *EmailSuppressionResourceCrud) SetData() error {
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
