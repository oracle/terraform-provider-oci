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

func IdentityApiKeyResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createIdentityApiKey,
		Read:     readIdentityApiKey,
		Delete:   deleteIdentityApiKey,
		Schema: map[string]*schema.Schema{
			// Required
			"key_value": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
				DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
					r := regexp.MustCompile("\\s")
					strippedOld := r.ReplaceAllString(old, "")
					strippedNew := r.ReplaceAllString(new, "")
					return (strippedOld == strippedNew)
				},
			},
			"user_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			// Optional

			// Computed
			"fingerprint": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"inactive_status": {
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

func createIdentityApiKey(d *schema.ResourceData, m interface{}) error {
	sync := &IdentityApiKeyResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).IdentityClient()

	return tfresource.CreateResource(d, sync)
}

func readIdentityApiKey(d *schema.ResourceData, m interface{}) error {
	sync := &IdentityApiKeyResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).IdentityClient()

	return tfresource.ReadResource(sync)
}

func deleteIdentityApiKey(d *schema.ResourceData, m interface{}) error {
	sync := &IdentityApiKeyResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).IdentityClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type IdentityApiKeyResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_identity.IdentityClient
	Res                    *oci_identity.ApiKey
	DisableNotFoundRetries bool
}

func (s *IdentityApiKeyResourceCrud) ID() string {
	return *s.Res.KeyId
}

func (s *IdentityApiKeyResourceCrud) State() oci_identity.ApiKeyLifecycleStateEnum {
	return s.Res.LifecycleState
}

func (s *IdentityApiKeyResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_identity.ApiKeyLifecycleStateCreating),
	}
}

func (s *IdentityApiKeyResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_identity.ApiKeyLifecycleStateActive),
	}
}

func (s *IdentityApiKeyResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_identity.ApiKeyLifecycleStateDeleting),
	}
}

func (s *IdentityApiKeyResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_identity.ApiKeyLifecycleStateDeleted),
	}
}

func (s *IdentityApiKeyResourceCrud) Create() error {
	request := oci_identity.UploadApiKeyRequest{}

	if keyValue, ok := s.D.GetOkExists("key_value"); ok {
		tmp := keyValue.(string)
		request.Key = &tmp
	}

	if userId, ok := s.D.GetOkExists("user_id"); ok {
		tmp := userId.(string)
		request.UserId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "identity")

	response, err := s.Client.UploadApiKey(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.ApiKey
	return nil
}

func (s *IdentityApiKeyResourceCrud) Get() error {
	request := oci_identity.ListApiKeysRequest{}

	if userId, ok := s.D.GetOkExists("user_id"); ok {
		tmp := userId.(string)
		request.UserId = &tmp
	}

	fingerprintFromCompositeId, userId, err := parseApiKeyCompositeId(s.D.Id())
	if err == nil {
		s.D.Set("fingerprint", fingerprintFromCompositeId)
		request.UserId = &userId
	} else {
		log.Printf("[WARN] Get() unable to parse current ID: %s", s.D.Id())
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "identity")

	response, err := s.Client.ListApiKeys(context.Background(), request)
	if err != nil {
		return err
	}

	fingerprint := s.D.Get("fingerprint").(string)
	for _, item := range response.Items {
		if *item.Fingerprint == fingerprint {
			s.Res = &item
			s.D.SetId(*s.Res.KeyId)
			return nil
		}
	}
	return errors.New("ApiKey with expected identifier not found")

}

func (s *IdentityApiKeyResourceCrud) Delete() error {
	request := oci_identity.DeleteApiKeyRequest{}

	if fingerprint, ok := s.D.GetOkExists("fingerprint"); ok {
		tmp := fingerprint.(string)
		request.Fingerprint = &tmp
	}

	if userId, ok := s.D.GetOkExists("user_id"); ok {
		tmp := userId.(string)
		request.UserId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "identity")

	_, err := s.Client.DeleteApiKey(context.Background(), request)
	return err
}

func (s *IdentityApiKeyResourceCrud) SetData() error {
	if s.Res.Fingerprint != nil {
		s.D.Set("fingerprint", *s.Res.Fingerprint)
	}

	if s.Res.InactiveStatus != nil {
		s.D.Set("inactive_status", strconv.FormatInt(*s.Res.InactiveStatus, 10))
	}

	if s.Res.KeyValue != nil {
		s.D.Set("key_value", *s.Res.KeyValue)
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

func GetApiKeyCompositeId(fingerprint string, userId string) string {
	fingerprint = url.PathEscape(fingerprint)
	userId = url.PathEscape(userId)
	compositeId := "users/" + userId + "/apiKeys/" + fingerprint
	return compositeId
}

func parseApiKeyCompositeId(compositeId string) (fingerprint string, userId string, err error) {
	parts := strings.Split(compositeId, "/")
	match, _ := regexp.MatchString("users/.*/apiKeys/.*", compositeId)
	if !match || len(parts) != 4 {
		err = fmt.Errorf("illegal compositeId %s encountered", compositeId)
		return
	}
	userId, _ = url.PathUnescape(parts[1])
	fingerprint, _ = url.PathUnescape(parts[3])

	return
}
