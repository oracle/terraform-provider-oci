// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package identity

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/url"
	"regexp"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"

	oci_identity "github.com/oracle/oci-go-sdk/v58/identity"
)

func IdentitySmtpCredentialResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createIdentitySmtpCredential,
		Read:     readIdentitySmtpCredential,
		Update:   updateIdentitySmtpCredential,
		Delete:   deleteIdentitySmtpCredential,
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
			"time_expires": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"username": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func createIdentitySmtpCredential(d *schema.ResourceData, m interface{}) error {
	sync := &IdentitySmtpCredentialResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).IdentityClient()

	return tfresource.CreateResource(d, sync)
}

func readIdentitySmtpCredential(d *schema.ResourceData, m interface{}) error {
	sync := &IdentitySmtpCredentialResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).IdentityClient()

	return tfresource.ReadResource(sync)
}

func updateIdentitySmtpCredential(d *schema.ResourceData, m interface{}) error {
	sync := &IdentitySmtpCredentialResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).IdentityClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteIdentitySmtpCredential(d *schema.ResourceData, m interface{}) error {
	sync := &IdentitySmtpCredentialResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).IdentityClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type IdentitySmtpCredentialResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_identity.IdentityClient
	Res                    *oci_identity.SmtpCredential
	DisableNotFoundRetries bool
}

func (s *IdentitySmtpCredentialResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *IdentitySmtpCredentialResourceCrud) State() oci_identity.SmtpCredentialLifecycleStateEnum {
	return s.Res.LifecycleState
}

func (s *IdentitySmtpCredentialResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_identity.SmtpCredentialLifecycleStateCreating),
	}
}

func (s *IdentitySmtpCredentialResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_identity.SmtpCredentialLifecycleStateActive),
	}
}

func (s *IdentitySmtpCredentialResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_identity.SmtpCredentialLifecycleStateDeleting),
	}
}

func (s *IdentitySmtpCredentialResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_identity.SmtpCredentialLifecycleStateDeleted),
	}
}

func (s *IdentitySmtpCredentialResourceCrud) Create() error {
	request := oci_identity.CreateSmtpCredentialRequest{}

	if description, ok := s.D.GetOkExists("description"); ok {
		tmp := description.(string)
		request.Description = &tmp
	}

	if userId, ok := s.D.GetOkExists("user_id"); ok {
		tmp := userId.(string)
		request.UserId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "identity")

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

func (s *IdentitySmtpCredentialResourceCrud) Get() error {
	request := oci_identity.ListSmtpCredentialsRequest{}

	if userId, ok := s.D.GetOkExists("user_id"); ok {
		tmp := userId.(string)
		request.UserId = &tmp
	}

	smtpCredentialId, userId, err := parseSmtpCredentialCompositeId(s.D.Id())
	if err == nil {
		s.D.SetId(smtpCredentialId)
		request.UserId = &userId
	} else {
		log.Printf("[WARN] Get() unable to parse current ID: %s", s.D.Id())
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "identity")

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
	return errors.New("SmtpCredential with expected identifier not found")

}

func (s *IdentitySmtpCredentialResourceCrud) Update() error {
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

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "identity")

	response, err := s.Client.UpdateSmtpCredential(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = fromSmtpCredentialSummary(response.SmtpCredentialSummary)
	return nil
}

func (s *IdentitySmtpCredentialResourceCrud) Delete() error {
	request := oci_identity.DeleteSmtpCredentialRequest{}

	tmp := s.D.Id()
	request.SmtpCredentialId = &tmp

	if userId, ok := s.D.GetOkExists("user_id"); ok {
		tmp := userId.(string)
		request.UserId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "identity")

	_, err := s.Client.DeleteSmtpCredential(context.Background(), request)
	return err
}

func (s *IdentitySmtpCredentialResourceCrud) SetData() error {
	if s.Res.Description != nil {
		s.D.Set("description", *s.Res.Description)
	}

	if s.Res.InactiveStatus != nil {
		s.D.Set("inactive_state", strconv.FormatInt(*s.Res.InactiveStatus, 10))
	}

	if s.Res.Password != nil {
		s.D.Set("password", *s.Res.Password)
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

	return nil
}

func GetSmtpCredentialCompositeId(smtpCredentialId string, userId string) string {
	smtpCredentialId = url.PathEscape(smtpCredentialId)
	userId = url.PathEscape(userId)
	compositeId := "users/" + userId + "/smtpCredentials/" + smtpCredentialId
	return compositeId
}

func parseSmtpCredentialCompositeId(compositeId string) (smtpCredentialId string, userId string, err error) {
	parts := strings.Split(compositeId, "/")
	match, _ := regexp.MatchString("users/.*/smtpCredentials/.*", compositeId)
	if !match || len(parts) != 4 {
		err = fmt.Errorf("illegal compositeId %s encountered", compositeId)
		return
	}
	userId, _ = url.PathUnescape(parts[1])
	smtpCredentialId, _ = url.PathUnescape(parts[3])

	return
}
