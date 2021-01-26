// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	oci_email "github.com/oracle/oci-go-sdk/v34/email"
)

func init() {
	RegisterResource("oci_email_suppression", EmailSuppressionResource())
}

func EmailSuppressionResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: DefaultTimeout,
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
				DiffSuppressFunc: EqualIgnoreCaseSuppressDiff,
			},

			// Optional

			// Computed
			"reason": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_created": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func createEmailSuppression(d *schema.ResourceData, m interface{}) error {
	sync := &EmailSuppressionResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).emailClient()

	return CreateResource(d, sync)
}

func readEmailSuppression(d *schema.ResourceData, m interface{}) error {
	sync := &EmailSuppressionResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).emailClient()

	return ReadResource(sync)
}

func deleteEmailSuppression(d *schema.ResourceData, m interface{}) error {
	sync := &EmailSuppressionResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).emailClient()
	sync.DisableNotFoundRetries = true

	return DeleteResource(d, sync)
}

type EmailSuppressionResourceCrud struct {
	BaseCrud
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

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "email")

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

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "email")

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

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "email")

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

	s.D.Set("reason", s.Res.Reason)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	return nil
}
