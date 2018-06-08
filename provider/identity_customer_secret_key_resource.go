// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"context"

	"github.com/hashicorp/terraform/helper/schema"

	"github.com/oracle/terraform-provider-oci/crud"

	oci_identity "github.com/oracle/oci-go-sdk/identity"
)

func CustomerSecretKeyResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: crud.DefaultTimeout,
		Create:   createCustomerSecretKey,
		Read:     readCustomerSecretKey,
		Update:   updateCustomerSecretKey,
		Delete:   deleteCustomerSecretKey,
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
			"id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"inactive_state": {
				Type:     schema.TypeInt,
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

func createCustomerSecretKey(d *schema.ResourceData, m interface{}) error {
	sync := &CustomerSecretKeyResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).identityClient

	return crud.CreateResource(d, sync)
}

func readCustomerSecretKey(d *schema.ResourceData, m interface{}) error {
	sync := &CustomerSecretKeyResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).identityClient

	return crud.ReadResource(sync)
}

func updateCustomerSecretKey(d *schema.ResourceData, m interface{}) error {
	sync := &CustomerSecretKeyResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).identityClient

	return crud.UpdateResource(d, sync)
}

func deleteCustomerSecretKey(d *schema.ResourceData, m interface{}) error {
	sync := &CustomerSecretKeyResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).identityClient
	sync.DisableNotFoundRetries = true

	return crud.DeleteResource(d, sync)
}

type CustomerSecretKeyResourceCrud struct {
	crud.BaseCrud
	Client                 *oci_identity.IdentityClient
	Res                    *oci_identity.CustomerSecretKey
	DisableNotFoundRetries bool
}

func (s *CustomerSecretKeyResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *CustomerSecretKeyResourceCrud) State() oci_identity.CustomerSecretKeyLifecycleStateEnum {
	return s.Res.LifecycleState
}

func (s *CustomerSecretKeyResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_identity.CustomerSecretKeyLifecycleStateCreating),
	}
}

func (s *CustomerSecretKeyResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_identity.CustomerSecretKeyLifecycleStateActive),
	}
}

func (s *CustomerSecretKeyResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_identity.CustomerSecretKeyLifecycleStateDeleting),
	}
}

func (s *CustomerSecretKeyResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_identity.CustomerSecretKeyLifecycleStateDeleted),
	}
}

func (s *CustomerSecretKeyResourceCrud) Create() error {
	request := oci_identity.CreateCustomerSecretKeyRequest{}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if userId, ok := s.D.GetOkExists("user_id"); ok {
		tmp := userId.(string)
		request.UserId = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "identity")

	response, err := s.Client.CreateCustomerSecretKey(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.CustomerSecretKey
	return nil
}

func (s *CustomerSecretKeyResourceCrud) Get() error {
	request := oci_identity.ListCustomerSecretKeysRequest{}

	if userId, ok := s.D.GetOkExists("user_id"); ok {
		tmp := userId.(string)
		request.UserId = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "identity")

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

	return nil
}

func (s *CustomerSecretKeyResourceCrud) Update() error {
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

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "identity")

	response, err := s.Client.UpdateCustomerSecretKey(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = fromCustomerSecretKeySummary(response.CustomerSecretKeySummary)
	return nil
}

func (s *CustomerSecretKeyResourceCrud) Delete() error {
	request := oci_identity.DeleteCustomerSecretKeyRequest{}

	tmp := s.D.Id()
	request.CustomerSecretKeyId = &tmp

	if userId, ok := s.D.GetOkExists("user_id"); ok {
		tmp := userId.(string)
		request.UserId = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "identity")

	_, err := s.Client.DeleteCustomerSecretKey(context.Background(), request)
	return err
}

func (s *CustomerSecretKeyResourceCrud) SetData() {
	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	if s.Res.Id != nil {
		s.D.Set("id", *s.Res.Id)
	}

	if s.Res.InactiveStatus != nil {
		s.D.Set("inactive_state", *s.Res.InactiveStatus)
	}

	if s.Res.Key != nil && *s.Res.Key != "" {
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

}

func fromCustomerSecretKeySummary(summary oci_identity.CustomerSecretKeySummary) *oci_identity.CustomerSecretKey {
	s := &oci_identity.CustomerSecretKey{}
	s.Id = summary.Id
	s.DisplayName = summary.DisplayName
	s.TimeExpires = summary.TimeExpires
	s.TimeCreated = summary.TimeCreated
	s.LifecycleState = oci_identity.CustomerSecretKeyLifecycleStateEnum(summary.LifecycleState)
	s.InactiveStatus = summary.InactiveStatus
	return s
}
