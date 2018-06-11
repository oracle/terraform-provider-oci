// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"context"

	"github.com/hashicorp/terraform/helper/schema"

	"github.com/oracle/terraform-provider-oci/crud"

	oci_identity "github.com/oracle/oci-go-sdk/identity"
)

func AuthTokenResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: crud.DefaultTimeout,
		Create:   createAuthToken,
		Read:     readAuthToken,
		Update:   updateAuthToken,
		Delete:   deleteAuthToken,
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
			"id": {
				Type:     schema.TypeString,
				Computed: true,
			},
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
			"token": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func createAuthToken(d *schema.ResourceData, m interface{}) error {
	sync := &AuthTokenResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).identityClient

	return crud.CreateResource(d, sync)
}

func readAuthToken(d *schema.ResourceData, m interface{}) error {
	sync := &AuthTokenResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).identityClient

	return crud.ReadResource(sync)
}

func updateAuthToken(d *schema.ResourceData, m interface{}) error {
	sync := &AuthTokenResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).identityClient

	return crud.UpdateResource(d, sync)
}

func deleteAuthToken(d *schema.ResourceData, m interface{}) error {
	sync := &AuthTokenResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).identityClient
	sync.DisableNotFoundRetries = true

	return crud.DeleteResource(d, sync)
}

type AuthTokenResourceCrud struct {
	crud.BaseCrud
	Client                 *oci_identity.IdentityClient
	Res                    *oci_identity.AuthToken
	DisableNotFoundRetries bool
}

func (s *AuthTokenResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *AuthTokenResourceCrud) State() oci_identity.AuthTokenLifecycleStateEnum {
	return s.Res.LifecycleState
}

func (s *AuthTokenResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_identity.AuthTokenLifecycleStateCreating),
	}
}

func (s *AuthTokenResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_identity.AuthTokenLifecycleStateActive),
	}
}

func (s *AuthTokenResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_identity.AuthTokenLifecycleStateDeleting),
	}
}

func (s *AuthTokenResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_identity.AuthTokenLifecycleStateDeleted),
	}
}

func (s *AuthTokenResourceCrud) Create() error {
	request := oci_identity.CreateAuthTokenRequest{}

	if description, ok := s.D.GetOkExists("description"); ok {
		tmp := description.(string)
		request.Description = &tmp
	}

	if userId, ok := s.D.GetOkExists("user_id"); ok {
		tmp := userId.(string)
		request.UserId = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "identity")

	response, err := s.Client.CreateAuthToken(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.AuthToken
	return nil
}

func (s *AuthTokenResourceCrud) Get() error {
	request := oci_identity.ListAuthTokensRequest{}

	if userId, ok := s.D.GetOkExists("user_id"); ok {
		tmp := userId.(string)
		request.UserId = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "identity")

	response, err := s.Client.ListAuthTokens(context.Background(), request)
	if err != nil {
		return err
	}

	id := s.D.Get("id").(string)
	for _, item := range response.Items {
		if *item.Id == id {
			s.Res = &item
			return nil
		}
	}

	return nil
}

func (s *AuthTokenResourceCrud) Update() error {
	request := oci_identity.UpdateAuthTokenRequest{}

	tmp := s.D.Id()
	request.AuthTokenId = &tmp

	if description, ok := s.D.GetOkExists("description"); ok {
		tmp := description.(string)
		request.Description = &tmp
	}

	if userId, ok := s.D.GetOkExists("user_id"); ok {
		tmp := userId.(string)
		request.UserId = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "identity")

	response, err := s.Client.UpdateAuthToken(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.AuthToken
	return nil
}

func (s *AuthTokenResourceCrud) Delete() error {
	request := oci_identity.DeleteAuthTokenRequest{}

	tmp := s.D.Id()
	request.AuthTokenId = &tmp

	if userId, ok := s.D.GetOkExists("user_id"); ok {
		tmp := userId.(string)
		request.UserId = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "identity")

	_, err := s.Client.DeleteAuthToken(context.Background(), request)
	return err
}

func (s *AuthTokenResourceCrud) SetData() {
	if s.Res.Description != nil {
		s.D.Set("description", *s.Res.Description)
	}

	if s.Res.Id != nil {
		s.D.Set("id", *s.Res.Id)
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

	if s.Res.Token != nil && *s.Res.Token != "" {
		s.D.Set("token", *s.Res.Token)
	}

	if s.Res.UserId != nil {
		s.D.Set("user_id", *s.Res.UserId)
	}

}
