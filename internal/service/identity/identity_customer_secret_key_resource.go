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

func IdentityCustomerSecretKeyResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createIdentityCustomerSecretKey,
		Read:     readIdentityCustomerSecretKey,
		Update:   updateIdentityCustomerSecretKey,
		Delete:   deleteIdentityCustomerSecretKey,
		Schema: map[string]*schema.Schema{
			// Required
			"display_name": {
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
			"key": {
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

func createIdentityCustomerSecretKey(d *schema.ResourceData, m interface{}) error {
	sync := &IdentityCustomerSecretKeyResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).IdentityClient()

	return tfresource.CreateResource(d, sync)
}

func readIdentityCustomerSecretKey(d *schema.ResourceData, m interface{}) error {
	sync := &IdentityCustomerSecretKeyResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).IdentityClient()

	return tfresource.ReadResource(sync)
}

func updateIdentityCustomerSecretKey(d *schema.ResourceData, m interface{}) error {
	sync := &IdentityCustomerSecretKeyResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).IdentityClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteIdentityCustomerSecretKey(d *schema.ResourceData, m interface{}) error {
	sync := &IdentityCustomerSecretKeyResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).IdentityClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type IdentityCustomerSecretKeyResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_identity.IdentityClient
	Res                    *oci_identity.CustomerSecretKey
	DisableNotFoundRetries bool
}

func (s *IdentityCustomerSecretKeyResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *IdentityCustomerSecretKeyResourceCrud) State() oci_identity.CustomerSecretKeyLifecycleStateEnum {
	return s.Res.LifecycleState
}

func (s *IdentityCustomerSecretKeyResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_identity.CustomerSecretKeyLifecycleStateCreating),
	}
}

func (s *IdentityCustomerSecretKeyResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_identity.CustomerSecretKeyLifecycleStateActive),
	}
}

func (s *IdentityCustomerSecretKeyResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_identity.CustomerSecretKeyLifecycleStateDeleting),
	}
}

func (s *IdentityCustomerSecretKeyResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_identity.CustomerSecretKeyLifecycleStateDeleted),
	}
}

func (s *IdentityCustomerSecretKeyResourceCrud) Create() error {
	request := oci_identity.CreateCustomerSecretKeyRequest{}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if userId, ok := s.D.GetOkExists("user_id"); ok {
		tmp := userId.(string)
		request.UserId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "identity")

	response, err := s.Client.CreateCustomerSecretKey(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.CustomerSecretKey
	return nil
}

func (s *IdentityCustomerSecretKeyResourceCrud) Get() error {
	request := oci_identity.ListCustomerSecretKeysRequest{}

	if userId, ok := s.D.GetOkExists("user_id"); ok {
		tmp := userId.(string)
		request.UserId = &tmp
	}

	customerSecretKeyId, userId, err := parseCustomerSecretKeyCompositeId(s.D.Id())
	if err == nil {
		s.D.SetId(customerSecretKeyId)
		request.UserId = &userId
	} else {
		log.Printf("[WARN] Get() unable to parse current ID: %s", s.D.Id())
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "identity")

	response, err := s.Client.ListCustomerSecretKeys(context.Background(), request)
	if err != nil {
		return err
	}

	id := s.D.Id()
	for _, item := range response.Items {
		if *item.Id == id {
			s.Res = fromCustomerSecretKeySummary(item)
			return nil
		}
	}
	return errors.New("CustomerSecretKey with expected identifier not found")

}

func (s *IdentityCustomerSecretKeyResourceCrud) Update() error {
	request := oci_identity.UpdateCustomerSecretKeyRequest{}

	tmp := s.D.Id()
	request.CustomerSecretKeyId = &tmp

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if userId, ok := s.D.GetOkExists("user_id"); ok {
		tmp := userId.(string)
		request.UserId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "identity")

	response, err := s.Client.UpdateCustomerSecretKey(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = fromCustomerSecretKeySummary(response.CustomerSecretKeySummary)
	return nil
}

func (s *IdentityCustomerSecretKeyResourceCrud) Delete() error {
	request := oci_identity.DeleteCustomerSecretKeyRequest{}

	tmp := s.D.Id()
	request.CustomerSecretKeyId = &tmp

	if userId, ok := s.D.GetOkExists("user_id"); ok {
		tmp := userId.(string)
		request.UserId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "identity")

	_, err := s.Client.DeleteCustomerSecretKey(context.Background(), request)
	return err
}

func (s *IdentityCustomerSecretKeyResourceCrud) SetData() error {
	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	if s.Res.InactiveStatus != nil {
		s.D.Set("inactive_state", strconv.FormatInt(*s.Res.InactiveStatus, 10))
	}

	if s.Res.Key != nil {
		s.D.Set("key", *s.Res.Key)
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

func GetCustomerSecretKeyCompositeId(customerSecretKeyId string, userId string) string {
	customerSecretKeyId = url.PathEscape(customerSecretKeyId)
	userId = url.PathEscape(userId)
	compositeId := "users/" + userId + "/customerSecretKeys/" + customerSecretKeyId
	return compositeId
}

func parseCustomerSecretKeyCompositeId(compositeId string) (customerSecretKeyId string, userId string, err error) {
	parts := strings.Split(compositeId, "/")
	match, _ := regexp.MatchString("users/.*/customerSecretKeys/.*", compositeId)
	if !match || len(parts) != 4 {
		err = fmt.Errorf("illegal compositeId %s encountered", compositeId)
		return
	}
	userId, _ = url.PathUnescape(parts[1])
	customerSecretKeyId, _ = url.PathUnescape(parts[3])

	return
}

func fromCustomerSecretKeySummary(summary oci_identity.CustomerSecretKeySummary) *oci_identity.CustomerSecretKey {
	s := &oci_identity.CustomerSecretKey{}
	s.Id = summary.Id
	s.DisplayName = summary.DisplayName
	s.UserId = summary.UserId
	s.TimeExpires = summary.TimeExpires
	s.TimeCreated = summary.TimeCreated
	s.LifecycleState = oci_identity.CustomerSecretKeyLifecycleStateEnum(summary.LifecycleState)
	s.InactiveStatus = summary.InactiveStatus
	return s
}
