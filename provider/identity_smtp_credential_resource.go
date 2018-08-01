// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"context"

	"github.com/hashicorp/terraform/helper/schema"

	oci_identity "github.com/oracle/oci-go-sdk/identity"
)

func SmtpCredentialResource() *schema.Resource {
	return &schema.Resource{
		Timeouts: DefaultTimeout,
		Create:   createSmtpCredential,
		Read:     readSmtpCredential,
		Update:   updateSmtpCredential,
		Delete:   deleteSmtpCredential,
		Schema: map[string]*schema.Schema{
			// Required
			"description": {
				Type:     schema.TypeString,
				Required: true,
			},
			"user_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			// Optional

			// Computed
			"inactive_state": {
				Type:     schema.TypeInt,
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
			"time_expires": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"username": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"password": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func createSmtpCredential(d *schema.ResourceData, m interface{}) error {
	sync := &SmtpCredentialResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).identityClient

	return CreateResource(d, sync)
}

func readSmtpCredential(d *schema.ResourceData, m interface{}) error {
	sync := &SmtpCredentialResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).identityClient

	return ReadResource(sync)
}

func updateSmtpCredential(d *schema.ResourceData, m interface{}) error {
	sync := &SmtpCredentialResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).identityClient

	return UpdateResource(d, sync)
}

func deleteSmtpCredential(d *schema.ResourceData, m interface{}) error {
	sync := &SmtpCredentialResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).identityClient
	sync.DisableNotFoundRetries = true

	return DeleteResource(d, sync)
}

type SmtpCredentialResourceCrud struct {
	BaseCrud
	Client                 *oci_identity.IdentityClient
	Res                    *oci_identity.SmtpCredential
	DisableNotFoundRetries bool
}

func (s *SmtpCredentialResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *SmtpCredentialResourceCrud) State() oci_identity.SmtpCredentialLifecycleStateEnum {
	return s.Res.LifecycleState
}

func (s *SmtpCredentialResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_identity.SmtpCredentialLifecycleStateCreating),
	}
}

func (s *SmtpCredentialResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_identity.SmtpCredentialLifecycleStateActive),
	}
}

func (s *SmtpCredentialResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_identity.SmtpCredentialLifecycleStateDeleting),
	}
}

func (s *SmtpCredentialResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_identity.SmtpCredentialLifecycleStateDeleted),
	}
}

func (s *SmtpCredentialResourceCrud) Create() error {
	request := oci_identity.CreateSmtpCredentialRequest{}

	if description, ok := s.D.GetOkExists("description"); ok {
		tmp := description.(string)
		request.Description = &tmp
	}

	if userId, ok := s.D.GetOkExists("user_id"); ok {
		tmp := userId.(string)
		request.UserId = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "identity")

	response, err := s.Client.CreateSmtpCredential(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.SmtpCredential
	return nil
}

func fromSmtpCredentialSummary(c oci_identity.SmtpCredentialSummary) *oci_identity.SmtpCredential {
	return &oci_identity.SmtpCredential{Username: c.Username,
		Id:             c.Id,
		UserId:         c.UserId,
		Description:    c.Description,
		TimeCreated:    c.TimeCreated,
		TimeExpires:    c.TimeExpires,
		LifecycleState: oci_identity.SmtpCredentialLifecycleStateEnum(c.LifecycleState),
		InactiveStatus: c.InactiveStatus}
}

func (s *SmtpCredentialResourceCrud) Get() error {
	request := oci_identity.ListSmtpCredentialsRequest{}

	if userId, ok := s.D.GetOkExists("user_id"); ok {
		tmp := userId.(string)
		request.UserId = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "identity")

	response, err := s.Client.ListSmtpCredentials(context.Background(), request)
	if err != nil {
		return err
	}

	id := s.D.Id()
	for _, item := range response.Items {
		if *item.Id == id {
			s.Res = fromSmtpCredentialSummary(item)
			return nil
		}
	}

	return nil
}

func (s *SmtpCredentialResourceCrud) Update() error {
	request := oci_identity.UpdateSmtpCredentialRequest{}

	if description, ok := s.D.GetOkExists("description"); ok {
		tmp := description.(string)
		request.Description = &tmp
	}

	tmp := s.D.Id()
	request.SmtpCredentialId = &tmp

	if userId, ok := s.D.GetOkExists("user_id"); ok {
		tmp := userId.(string)
		request.UserId = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "identity")

	response, err := s.Client.UpdateSmtpCredential(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = fromSmtpCredentialSummary(response.SmtpCredentialSummary)
	return nil
}

func (s *SmtpCredentialResourceCrud) Delete() error {
	request := oci_identity.DeleteSmtpCredentialRequest{}

	tmp := s.D.Id()
	request.SmtpCredentialId = &tmp

	if userId, ok := s.D.GetOkExists("user_id"); ok {
		tmp := userId.(string)
		request.UserId = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "identity")

	_, err := s.Client.DeleteSmtpCredential(context.Background(), request)
	return err
}

func (s *SmtpCredentialResourceCrud) SetData() error {
	if s.Res.Description != nil {
		s.D.Set("description", *s.Res.Description)
	}

	if s.Res.InactiveStatus != nil {
		s.D.Set("inactive_state", *s.Res.InactiveStatus)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeExpires != nil {
		s.D.Set("time_expires", s.Res.TimeExpires.String())
	}

	if s.Res.UserId != nil {
		s.D.Set("user_id", *s.Res.UserId)
	}

	if s.Res.Username != nil {
		s.D.Set("username", *s.Res.Username)
	}

	if s.Res.Password != nil {
		s.D.Set("password", *s.Res.Password)
	}

	return nil
}
