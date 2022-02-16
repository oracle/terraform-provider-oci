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

func IdentityDbCredentialResource() *schema.Resource {
	return &schema.Resource{
		Timeouts: tfresource.DefaultTimeout,
		Create:   createIdentityDbCredential,
		Read:     readIdentityDbCredential,
		Delete:   deleteIdentityDbCredential,
		Schema: map[string]*schema.Schema{
			// Required
			"description": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"password": {
				Type:      schema.TypeString,
				Required:  true,
				ForceNew:  true,
				Sensitive: true,
			},
			"user_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			// Optional

			// Computed
			"lifecycle_details": {
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
		},
	}
}

func createIdentityDbCredential(d *schema.ResourceData, m interface{}) error {
	sync := &IdentityDbCredentialResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).IdentityClient()

	return tfresource.CreateResource(d, sync)
}

func readIdentityDbCredential(d *schema.ResourceData, m interface{}) error {
	sync := &IdentityDbCredentialResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).IdentityClient()

	return tfresource.ReadResource(sync)
}

func deleteIdentityDbCredential(d *schema.ResourceData, m interface{}) error {
	sync := &IdentityDbCredentialResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).IdentityClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type IdentityDbCredentialResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_identity.IdentityClient
	Res                    *oci_identity.DbCredential
	DisableNotFoundRetries bool
}

func (s *IdentityDbCredentialResourceCrud) ID() string {
	return GetDbCredentialCompositeId(*s.Res.Id, *s.Res.UserId)
}

func (s *IdentityDbCredentialResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_identity.DbCredentialLifecycleStateCreating),
	}
}

func (s *IdentityDbCredentialResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_identity.DbCredentialLifecycleStateActive),
	}
}

func (s *IdentityDbCredentialResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_identity.DbCredentialLifecycleStateDeleting),
	}
}

func (s *IdentityDbCredentialResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_identity.DbCredentialLifecycleStateDeleted),
	}
}

func (s *IdentityDbCredentialResourceCrud) Create() error {
	request := oci_identity.CreateDbCredentialRequest{}

	if description, ok := s.D.GetOkExists("description"); ok {
		tmp := description.(string)
		request.Description = &tmp
	}

	if password, ok := s.D.GetOkExists("password"); ok {
		tmp := password.(string)
		request.Password = &tmp
	}

	if userId, ok := s.D.GetOkExists("user_id"); ok {
		tmp := userId.(string)
		request.UserId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "identity")

	response, err := s.Client.CreateDbCredential(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.DbCredential
	return nil
}

func (s *IdentityDbCredentialResourceCrud) Get() error {
	request := oci_identity.ListDbCredentialsRequest{}

	if userId, ok := s.D.GetOkExists("user_id"); ok {
		tmp := userId.(string)
		request.UserId = &tmp
	}

	dbCredentialId, userId, err := parseDbCredentialCompositeId(s.D.Id())
	if err == nil {
		s.D.SetId(dbCredentialId)
		request.UserId = &userId
	} else {
		log.Printf("[WARN] Get() unable to parse current ID: %s", s.D.Id())
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "identity")

	response, err := s.Client.ListDbCredentials(context.Background(), request)
	if err != nil {
		return err
	}

	id := s.D.Id()
	for _, item := range response.Items {
		if *item.Id == id {
			s.Res = fromDbCredentialSummary(item)
			return nil
		}
	}
	return errors.New("DbCredential with expected identifier not found")

}

func fromDbCredentialSummary(c oci_identity.DbCredentialSummary) *oci_identity.DbCredential {
	return &oci_identity.DbCredential{Id: c.Id,
		UserId:         c.UserId,
		TimeCreated:    c.TimeCreated,
		TimeExpires:    c.TimeExpires,
		LifecycleState: oci_identity.DbCredentialLifecycleStateEnum(c.LifecycleState)}
}

func (s *IdentityDbCredentialResourceCrud) Delete() error {
	request := oci_identity.DeleteDbCredentialRequest{}

	tmp := s.D.Id()
	request.DbCredentialId = &tmp

	if userId, ok := s.D.GetOkExists("user_id"); ok {
		tmp := userId.(string)
		request.UserId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "identity")

	_, err := s.Client.DeleteDbCredential(context.Background(), request)
	return err
}

func (s *IdentityDbCredentialResourceCrud) SetData() error {

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", strconv.FormatInt(*s.Res.LifecycleDetails, 10))
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

	return nil
}

func GetDbCredentialCompositeId(dbCredentialId string, userId string) string {
	dbCredentialId = url.PathEscape(dbCredentialId)
	userId = url.PathEscape(userId)
	compositeId := "users/" + userId + "/dbCredentials/" + dbCredentialId
	return compositeId
}

func parseDbCredentialCompositeId(compositeId string) (dbCredentialId string, userId string, err error) {
	parts := strings.Split(compositeId, "/")
	match, _ := regexp.MatchString("users/.*/dbCredentials/.*", compositeId)
	if !match || len(parts) != 4 {
		err = fmt.Errorf("illegal compositeId %s encountered", compositeId)
		return
	}
	userId, _ = url.PathUnescape(parts[1])
	dbCredentialId, _ = url.PathUnescape(parts[3])

	return
}
