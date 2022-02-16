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

func IdentitySwiftPasswordResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createIdentitySwiftPassword,
		Read:     readIdentitySwiftPassword,
		Update:   updateIdentitySwiftPassword,
		Delete:   deleteIdentitySwiftPassword,
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
			"expires_on": {
				Type:     schema.TypeString,
				Computed: true,
			},
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
		},
	}
}

func createIdentitySwiftPassword(d *schema.ResourceData, m interface{}) error {
	sync := &IdentitySwiftPasswordResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).IdentityClient()

	return tfresource.CreateResource(d, sync)
}

func readIdentitySwiftPassword(d *schema.ResourceData, m interface{}) error {
	sync := &IdentitySwiftPasswordResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).IdentityClient()

	return tfresource.ReadResource(sync)
}

func updateIdentitySwiftPassword(d *schema.ResourceData, m interface{}) error {
	sync := &IdentitySwiftPasswordResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).IdentityClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteIdentitySwiftPassword(d *schema.ResourceData, m interface{}) error {
	sync := &IdentitySwiftPasswordResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).IdentityClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type IdentitySwiftPasswordResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_identity.IdentityClient
	Res                    *oci_identity.SwiftPassword
	DisableNotFoundRetries bool
}

func (s *IdentitySwiftPasswordResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *IdentitySwiftPasswordResourceCrud) State() oci_identity.SwiftPasswordLifecycleStateEnum {
	return s.Res.LifecycleState
}

func (s *IdentitySwiftPasswordResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_identity.SwiftPasswordLifecycleStateCreating),
	}
}

func (s *IdentitySwiftPasswordResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_identity.SwiftPasswordLifecycleStateActive),
	}
}

func (s *IdentitySwiftPasswordResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_identity.SwiftPasswordLifecycleStateDeleting),
	}
}

func (s *IdentitySwiftPasswordResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_identity.SwiftPasswordLifecycleStateDeleted),
	}
}

func (s *IdentitySwiftPasswordResourceCrud) Create() error {
	request := oci_identity.CreateSwiftPasswordRequest{}

	if description, ok := s.D.GetOkExists("description"); ok {
		tmp := description.(string)
		request.Description = &tmp
	}

	if userId, ok := s.D.GetOkExists("user_id"); ok {
		tmp := userId.(string)
		request.UserId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "identity")

	response, err := s.Client.CreateSwiftPassword(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.SwiftPassword
	return nil
}

func (s *IdentitySwiftPasswordResourceCrud) Get() error {
	request := oci_identity.ListSwiftPasswordsRequest{}

	if userId, ok := s.D.GetOkExists("user_id"); ok {
		tmp := userId.(string)
		request.UserId = &tmp
	}

	swiftPasswordId, userId, err := parseSwiftPasswordCompositeId(s.D.Id())
	if err == nil {
		s.D.SetId(swiftPasswordId)
		request.UserId = &userId
	} else {
		log.Printf("[WARN] Get() unable to parse current ID: %s", s.D.Id())
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "identity")

	response, err := s.Client.ListSwiftPasswords(context.Background(), request)
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
	return errors.New("SwiftPassword with expected identifier not found")

}

func (s *IdentitySwiftPasswordResourceCrud) Update() error {
	request := oci_identity.UpdateSwiftPasswordRequest{}

	if description, ok := s.D.GetOkExists("description"); ok {
		tmp := description.(string)
		request.Description = &tmp
	}

	tmp := s.D.Id()
	request.SwiftPasswordId = &tmp

	if userId, ok := s.D.GetOkExists("user_id"); ok {
		tmp := userId.(string)
		request.UserId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "identity")

	response, err := s.Client.UpdateSwiftPassword(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.SwiftPassword
	return nil
}

func (s *IdentitySwiftPasswordResourceCrud) Delete() error {
	request := oci_identity.DeleteSwiftPasswordRequest{}

	tmp := s.D.Id()
	request.SwiftPasswordId = &tmp

	if userId, ok := s.D.GetOkExists("user_id"); ok {
		tmp := userId.(string)
		request.UserId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "identity")

	_, err := s.Client.DeleteSwiftPassword(context.Background(), request)
	return err
}

func (s *IdentitySwiftPasswordResourceCrud) SetData() error {
	if s.Res.Description != nil {
		s.D.Set("description", *s.Res.Description)
	}

	if s.Res.ExpiresOn != nil {
		s.D.Set("expires_on", s.Res.ExpiresOn.String())
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

	if s.Res.UserId != nil {
		s.D.Set("user_id", *s.Res.UserId)
	}

	return nil
}

func GetSwiftPasswordCompositeId(swiftPasswordId string, userId string) string {
	swiftPasswordId = url.PathEscape(swiftPasswordId)
	userId = url.PathEscape(userId)
	compositeId := "users/" + userId + "/swiftPasswords/" + swiftPasswordId
	return compositeId
}

func parseSwiftPasswordCompositeId(compositeId string) (swiftPasswordId string, userId string, err error) {
	parts := strings.Split(compositeId, "/")
	match, _ := regexp.MatchString("users/.*/swiftPasswords/.*", compositeId)
	if !match || len(parts) != 4 {
		err = fmt.Errorf("illegal compositeId %s encountered", compositeId)
		return
	}
	userId, _ = url.PathUnescape(parts[1])
	swiftPasswordId, _ = url.PathUnescape(parts[3])

	return
}
