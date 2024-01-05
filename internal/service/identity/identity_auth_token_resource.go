// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
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

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	oci_identity "github.com/oracle/oci-go-sdk/v65/identity"
)

func IdentityAuthTokenResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createIdentityAuthToken,
		Read:     readIdentityAuthToken,
		Update:   updateIdentityAuthToken,
		Delete:   deleteIdentityAuthToken,
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

func createIdentityAuthToken(d *schema.ResourceData, m interface{}) error {
	sync := &IdentityAuthTokenResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).IdentityClient()

	return tfresource.CreateResource(d, sync)
}

func readIdentityAuthToken(d *schema.ResourceData, m interface{}) error {
	sync := &IdentityAuthTokenResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).IdentityClient()

	return tfresource.ReadResource(sync)
}

func updateIdentityAuthToken(d *schema.ResourceData, m interface{}) error {
	sync := &IdentityAuthTokenResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).IdentityClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteIdentityAuthToken(d *schema.ResourceData, m interface{}) error {
	sync := &IdentityAuthTokenResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).IdentityClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type IdentityAuthTokenResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_identity.IdentityClient
	Res                    *oci_identity.AuthToken
	DisableNotFoundRetries bool
}

func (s *IdentityAuthTokenResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *IdentityAuthTokenResourceCrud) State() oci_identity.AuthTokenLifecycleStateEnum {
	return s.Res.LifecycleState
}

func (s *IdentityAuthTokenResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_identity.AuthTokenLifecycleStateCreating),
	}
}

func (s *IdentityAuthTokenResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_identity.AuthTokenLifecycleStateActive),
	}
}

func (s *IdentityAuthTokenResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_identity.AuthTokenLifecycleStateDeleting),
	}
}

func (s *IdentityAuthTokenResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_identity.AuthTokenLifecycleStateDeleted),
	}
}

func (s *IdentityAuthTokenResourceCrud) Create() error {
	request := oci_identity.CreateAuthTokenRequest{}

	if description, ok := s.D.GetOkExists("description"); ok {
		tmp := description.(string)
		request.Description = &tmp
	}

	if userId, ok := s.D.GetOkExists("user_id"); ok {
		tmp := userId.(string)
		request.UserId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "identity")

	response, err := s.Client.CreateAuthToken(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.AuthToken
	return nil
}

func (s *IdentityAuthTokenResourceCrud) Get() error {
	request := oci_identity.ListAuthTokensRequest{}

	if userId, ok := s.D.GetOkExists("user_id"); ok {
		tmp := userId.(string)
		request.UserId = &tmp
	}

	authTokenId, userId, err := parseAuthTokenCompositeId(s.D.Id())
	if err == nil {
		s.D.SetId(authTokenId)
		request.UserId = &userId
	} else {
		log.Printf("[WARN] Get() unable to parse current ID: %s", s.D.Id())
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "identity")

	response, err := s.Client.ListAuthTokens(context.Background(), request)
	if err != nil {
		return err
	}

	id := s.D.Id()
	for _, item := range response.Items {
		if *item.Id == id {
			s.Res = &item
			return nil
		}
	}
	return errors.New("AuthToken with expected identifier not found")

}

func (s *IdentityAuthTokenResourceCrud) Update() error {
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

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "identity")

	response, err := s.Client.UpdateAuthToken(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.AuthToken
	return nil
}

func (s *IdentityAuthTokenResourceCrud) Delete() error {
	request := oci_identity.DeleteAuthTokenRequest{}

	tmp := s.D.Id()
	request.AuthTokenId = &tmp

	if userId, ok := s.D.GetOkExists("user_id"); ok {
		tmp := userId.(string)
		request.UserId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "identity")

	_, err := s.Client.DeleteAuthToken(context.Background(), request)
	return err
}

func (s *IdentityAuthTokenResourceCrud) SetData() error {
	if s.Res.Description != nil {
		s.D.Set("description", *s.Res.Description)
	}

	if s.Res.InactiveStatus != nil {
		s.D.Set("inactive_state", strconv.FormatInt(*s.Res.InactiveStatus, 10))
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

	return nil
}

func GetAuthTokenCompositeId(authTokenId string, userId string) string {
	authTokenId = url.PathEscape(authTokenId)
	userId = url.PathEscape(userId)
	compositeId := "users/" + userId + "/authTokens/" + authTokenId
	return compositeId
}

func parseAuthTokenCompositeId(compositeId string) (authTokenId string, userId string, err error) {
	parts := strings.Split(compositeId, "/")
	match, _ := regexp.MatchString("users/.*/authTokens/.*", compositeId)
	if !match || len(parts) != 4 {
		err = fmt.Errorf("illegal compositeId %s encountered", compositeId)
		return
	}
	userId, _ = url.PathUnescape(parts[1])
	authTokenId, _ = url.PathUnescape(parts[3])

	return
}
