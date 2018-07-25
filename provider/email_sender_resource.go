// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"context"

	"github.com/hashicorp/terraform/helper/schema"

	"github.com/oracle/terraform-provider-oci/crud"

	oci_email "github.com/oracle/oci-go-sdk/email"
)

func SenderResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: crud.DefaultTimeout,
		Create:   createSender,
		Read:     readSender,
		Delete:   deleteSender,
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

func createSender(d *schema.ResourceData, m interface{}) error {
	sync := &SenderResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).emailClient

	return crud.CreateResource(d, sync)
}

func readSender(d *schema.ResourceData, m interface{}) error {
	sync := &SenderResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).emailClient

	return crud.ReadResource(sync)
}

func deleteSender(d *schema.ResourceData, m interface{}) error {
	sync := &SenderResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).emailClient
	sync.DisableNotFoundRetries = true

	return crud.DeleteResource(d, sync)
}

type SenderResourceCrud struct {
	crud.BaseCrud
	Client                 *oci_email.EmailClient
	Res                    *oci_email.Sender
	DisableNotFoundRetries bool
}

func (s *SenderResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *SenderResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_email.SenderLifecycleStateCreating),
	}
}

func (s *SenderResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_email.SenderLifecycleStateActive),
	}
}

func (s *SenderResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_email.SenderLifecycleStateDeleting),
	}
}

func (s *SenderResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_email.SenderLifecycleStateDeleted),
	}
}

func (s *SenderResourceCrud) Create() error {
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

func (s *SenderResourceCrud) Get() error {
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

func (s *SenderResourceCrud) Delete() error {
	request := oci_email.DeleteSenderRequest{}

	tmp := s.D.Id()
	request.SenderId = &tmp

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "email")

	_, err := s.Client.DeleteSender(context.Background(), request)
	return err
}

func (s *SenderResourceCrud) SetData() error {
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
