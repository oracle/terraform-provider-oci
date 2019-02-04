// Copyright (c) 2017, 2019, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"context"

	"github.com/hashicorp/terraform/helper/schema"

	oci_email "github.com/oracle/oci-go-sdk/email"
)

func EmailSenderResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: DefaultTimeout,
		Create:   createEmailSender,
		Read:     readEmailSender,
		Delete:   deleteEmailSender,
		Schema: map[string]*schema.Schema{
			// Required
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"email_address": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			// Optional

			// Computed
			"is_spf": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"state": {
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

func createEmailSender(d *schema.ResourceData, m interface{}) error {
	sync := &EmailSenderResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).emailClient

	return CreateResource(d, sync)
}

func readEmailSender(d *schema.ResourceData, m interface{}) error {
	sync := &EmailSenderResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).emailClient

	return ReadResource(sync)
}

func deleteEmailSender(d *schema.ResourceData, m interface{}) error {
	sync := &EmailSenderResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).emailClient
	sync.DisableNotFoundRetries = true

	return DeleteResource(d, sync)
}

type EmailSenderResourceCrud struct {
	BaseCrud
	Client                 *oci_email.EmailClient
	Res                    *oci_email.Sender
	DisableNotFoundRetries bool
}

func (s *EmailSenderResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *EmailSenderResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_email.SenderLifecycleStateCreating),
	}
}

func (s *EmailSenderResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_email.SenderLifecycleStateActive),
	}
}

func (s *EmailSenderResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_email.SenderLifecycleStateDeleting),
	}
}

func (s *EmailSenderResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_email.SenderLifecycleStateDeleted),
	}
}

func (s *EmailSenderResourceCrud) Create() error {
	request := oci_email.CreateSenderRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if emailAddress, ok := s.D.GetOkExists("email_address"); ok {
		tmp := emailAddress.(string)
		request.EmailAddress = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "email")

	response, err := s.Client.CreateSender(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Sender
	return nil
}

func (s *EmailSenderResourceCrud) Get() error {
	request := oci_email.GetSenderRequest{}

	tmp := s.D.Id()
	request.SenderId = &tmp

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "email")

	response, err := s.Client.GetSender(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Sender
	return nil
}

func (s *EmailSenderResourceCrud) Delete() error {
	request := oci_email.DeleteSenderRequest{}

	tmp := s.D.Id()
	request.SenderId = &tmp

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "email")

	_, err := s.Client.DeleteSender(context.Background(), request)
	return err
}

func (s *EmailSenderResourceCrud) SetData() error {
	if s.Res.EmailAddress != nil {
		s.D.Set("email_address", *s.Res.EmailAddress)
	}

	if s.Res.IsSpf != nil {
		s.D.Set("is_spf", *s.Res.IsSpf)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	return nil
}
