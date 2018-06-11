// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"context"
	"errors"
	"regexp"

	"github.com/hashicorp/terraform/helper/schema"

	"github.com/oracle/terraform-provider-oci/crud"

	oci_identity "github.com/oracle/oci-go-sdk/identity"
)

func ApiKeyResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: crud.DefaultTimeout,
		Create:   createApiKey,
		Read:     readApiKey,
		Delete:   deleteApiKey,
		Schema: map[string]*schema.Schema{
			// Required
			"user_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
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

			// Optional

			// Computed
			"fingerprint": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"inactive_status": {
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
		},
	}
}

func createApiKey(d *schema.ResourceData, m interface{}) error {
	sync := &ApiKeyResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).identityClient

	return crud.CreateResource(d, sync)
}

func readApiKey(d *schema.ResourceData, m interface{}) error {
	sync := &ApiKeyResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).identityClient

	return crud.ReadResource(sync)
}

func deleteApiKey(d *schema.ResourceData, m interface{}) error {
	sync := &ApiKeyResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).identityClient
	sync.DisableNotFoundRetries = true

	return crud.DeleteResource(d, sync)
}

type ApiKeyResourceCrud struct {
	crud.BaseCrud
	Client                 *oci_identity.IdentityClient
	Res                    *oci_identity.ApiKey
	DisableNotFoundRetries bool
}

func (s *ApiKeyResourceCrud) ID() string {
	return *s.Res.KeyId
}

func (s *ApiKeyResourceCrud) State() oci_identity.ApiKeyLifecycleStateEnum {
	return s.Res.LifecycleState
}

func (s *ApiKeyResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_identity.ApiKeyLifecycleStateCreating),
	}
}

func (s *ApiKeyResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_identity.ApiKeyLifecycleStateActive),
	}
}

func (s *ApiKeyResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_identity.ApiKeyLifecycleStateDeleting),
	}
}

func (s *ApiKeyResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_identity.ApiKeyLifecycleStateDeleted),
	}
}

func (s *ApiKeyResourceCrud) Create() error {
	request := oci_identity.UploadApiKeyRequest{}

	if key, ok := s.D.GetOkExists("key_value"); ok {
		tmp := key.(string)
		request.Key = &tmp
	}

	if userId, ok := s.D.GetOkExists("user_id"); ok {
		tmp := userId.(string)
		request.UserId = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "identity")

	response, err := s.Client.UploadApiKey(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.ApiKey
	return nil
}

func (s *ApiKeyResourceCrud) Get() error {
	request := oci_identity.ListApiKeysRequest{}

	if userId, ok := s.D.GetOkExists("user_id"); ok {
		tmp := userId.(string)
		request.UserId = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "identity")

	response, err := s.Client.ListApiKeys(context.Background(), request)
	if err != nil {
		return err
	}

	fingerprint := s.D.Get("fingerprint").(string)
	for _, item := range response.Items {
		if *item.Fingerprint == fingerprint {
			s.Res = &item
			return nil
		}
	}

	return errors.New("API key not found")
}

func (s *ApiKeyResourceCrud) Delete() error {
	request := oci_identity.DeleteApiKeyRequest{}

	if fingerprint, ok := s.D.GetOkExists("fingerprint"); ok {
		tmp := fingerprint.(string)
		request.Fingerprint = &tmp
	}

	if userId, ok := s.D.GetOkExists("user_id"); ok {
		tmp := userId.(string)
		request.UserId = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "identity")

	_, err := s.Client.DeleteApiKey(context.Background(), request)
	return err
}

func (s *ApiKeyResourceCrud) SetData() {
	if s.Res.Fingerprint != nil {
		s.D.Set("fingerprint", *s.Res.Fingerprint)
	}

	if s.Res.KeyId != nil {
		s.D.Set("id", *s.Res.KeyId)
	}

	if s.Res.InactiveStatus != nil {
		s.D.Set("inactive_status", *s.Res.InactiveStatus)
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

}
