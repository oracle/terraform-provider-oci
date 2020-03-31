// Copyright (c) 2017, 2019, Oracle and/or its affiliates. All rights reserved.

package oci

import (
	"context"
	"strconv"

	"github.com/hashicorp/terraform/helper/schema"

	oci_identity "github.com/oracle/oci-go-sdk/identity"
)

func init() {
	RegisterResource("oci_identity_ui_password", IdentityUiPasswordResource())
}

func IdentityUiPasswordResource() *schema.Resource {
	return &schema.Resource{
		Timeouts: DefaultTimeout,
		Create:   createIdentityUiPassword,
		Read:     readIdentityUiPassword,
		Delete:   deleteIdentityUiPassword,
		Schema: map[string]*schema.Schema{
			// Required
			"user_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			// Optional

			// Computed
			"inactive_status": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"password": {
				Type:     schema.TypeString,
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

func createIdentityUiPassword(d *schema.ResourceData, m interface{}) error {
	sync := &IdentityUiPasswordResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).identityClient()

	return CreateResource(d, sync)
}

func readIdentityUiPassword(d *schema.ResourceData, m interface{}) error {
	return nil
}

func deleteIdentityUiPassword(d *schema.ResourceData, m interface{}) error {
	return nil
}

type IdentityUiPasswordResourceCrud struct {
	BaseCrud
	Client                 *oci_identity.IdentityClient
	Res                    *oci_identity.UiPassword
	DisableNotFoundRetries bool
}

func (s *IdentityUiPasswordResourceCrud) ID() string {
	return *s.Res.UserId
}

func (s *IdentityUiPasswordResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_identity.UiPasswordLifecycleStateCreating),
	}
}

func (s *IdentityUiPasswordResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_identity.UiPasswordLifecycleStateActive),
	}
}

func (s *IdentityUiPasswordResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_identity.UiPasswordLifecycleStateDeleting),
	}
}

func (s *IdentityUiPasswordResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_identity.UiPasswordLifecycleStateDeleted),
	}
}

func (s *IdentityUiPasswordResourceCrud) Create() error {
	request := oci_identity.CreateOrResetUIPasswordRequest{}

	if userId, ok := s.D.GetOkExists("user_id"); ok {
		tmp := userId.(string)
		request.UserId = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "identity")

	response, err := s.Client.CreateOrResetUIPassword(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.UiPassword
	return nil
}

func (s *IdentityUiPasswordResourceCrud) SetData() error {
	if s.Res.InactiveStatus != nil {
		s.D.Set("inactive_status", strconv.FormatInt(*s.Res.InactiveStatus, 10))
	}

	if s.Res.Password != nil {
		s.D.Set("password", *s.Res.Password)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.UserId != nil {
		s.D.Set("user_id", *s.Res.UserId)
	}

	return nil
}
