// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"context"

	"github.com/hashicorp/terraform/helper/schema"

	"github.com/oracle/terraform-provider-oci/crud"

	oci_identity "github.com/oracle/oci-go-sdk/identity"
)

func UiPasswordResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: crud.DefaultTimeout,
		Create:   createUiPassword,
		Read:     readUiPassword,
		Delete:   deleteUiPassword,
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
				Type:     schema.TypeInt,
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

func createUiPassword(d *schema.ResourceData, m interface{}) error {
	sync := &UiPasswordResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).identityClient

	return crud.CreateResource(d, sync)
}

func readUiPassword(d *schema.ResourceData, m interface{}) error {
	return nil
}

func deleteUiPassword(d *schema.ResourceData, m interface{}) error {
	return nil
}

type UiPasswordResourceCrud struct {
	crud.BaseCrud
	Client                 *oci_identity.IdentityClient
	Res                    *oci_identity.UiPassword
	DisableNotFoundRetries bool
}

func (s *UiPasswordResourceCrud) ID() string {
	return *s.Res.UserId
}

func (s *UiPasswordResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_identity.UiPasswordLifecycleStateCreating),
	}
}

func (s *UiPasswordResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_identity.UiPasswordLifecycleStateActive),
	}
}

func (s *UiPasswordResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_identity.UiPasswordLifecycleStateDeleting),
	}
}

func (s *UiPasswordResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_identity.UiPasswordLifecycleStateDeleted),
	}
}

func (s *UiPasswordResourceCrud) Create() error {
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

func (s *UiPasswordResourceCrud) SetData() {
	if s.Res.InactiveStatus != nil {
		s.D.Set("inactive_status", *s.Res.InactiveStatus)
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

}
