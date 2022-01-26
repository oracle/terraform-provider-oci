// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package email

import (
	"context"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"
	"github.com/terraform-providers/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	oci_email "github.com/oracle/oci-go-sdk/v56/email"
)

func EmailSenderResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createEmailSender,
		Read:     readEmailSender,
		Update:   updateEmailSender,
		Delete:   deleteEmailSender,
		Schema: map[string]*schema.Schema{
			// Required
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"email_address": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			// Optional
			"defined_tags": {
				Type:             schema.TypeMap,
				Optional:         true,
				Computed:         true,
				DiffSuppressFunc: tfresource.DefinedTagsDiffSuppressFunction,
				Elem:             schema.TypeString,
			},
			"freeform_tags": {
				Type:     schema.TypeMap,
				Optional: true,
				Computed: true,
				Elem:     schema.TypeString,
			},

			// Computed
			"email_domain_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
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
	sync.Client = m.(*client.OracleClients).EmailClient()

	return tfresource.CreateResource(d, sync)
}

func readEmailSender(d *schema.ResourceData, m interface{}) error {
	sync := &EmailSenderResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).EmailClient()

	return tfresource.ReadResource(sync)
}

func updateEmailSender(d *schema.ResourceData, m interface{}) error {
	sync := &EmailSenderResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).EmailClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteEmailSender(d *schema.ResourceData, m interface{}) error {
	sync := &EmailSenderResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).EmailClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type EmailSenderResourceCrud struct {
	tfresource.BaseCrud
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

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
		convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return err
		}
		request.DefinedTags = convertedDefinedTags
	}

	if emailAddress, ok := s.D.GetOkExists("email_address"); ok {
		tmp := emailAddress.(string)
		request.EmailAddress = &tmp
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = utils.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "email")

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

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "email")

	response, err := s.Client.GetSender(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Sender
	return nil
}

func (s *EmailSenderResourceCrud) Update() error {
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(compartment)
			if err != nil {
				return err
			}
		}
	}
	request := oci_email.UpdateSenderRequest{}

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
		convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return err
		}
		request.DefinedTags = convertedDefinedTags
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = utils.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	tmp := s.D.Id()
	request.SenderId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "email")

	response, err := s.Client.UpdateSender(context.Background(), request)
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

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "email")

	_, err := s.Client.DeleteSender(context.Background(), request)
	return err
}

func (s *EmailSenderResourceCrud) SetData() error {
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

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	return nil
}

func (s *EmailSenderResourceCrud) updateCompartment(compartment interface{}) error {
	changeCompartmentRequest := oci_email.ChangeSenderCompartmentRequest{}

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	idTmp := s.D.Id()
	changeCompartmentRequest.SenderId = &idTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "email")

	_, err := s.Client.ChangeSenderCompartment(context.Background(), changeCompartmentRequest)
	if err != nil {
		return err
	}

	if waitErr := tfresource.WaitForUpdatedState(s.D, s); waitErr != nil {
		return waitErr
	}

	return nil
}
