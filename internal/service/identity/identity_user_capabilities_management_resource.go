// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package identity

import (
	"context"
	"fmt"
	"log"
	"net/url"
	"regexp"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"

	oci_identity "github.com/oracle/oci-go-sdk/v56/identity"
)

func IdentityUserCapabilitiesManagementResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createUserCapabilitiesManagement,
		Read:     readUserCapabilitiesManagement,
		Update:   updateUserCapabilitiesManagement,
		Delete:   deleteUserCapabilitiesManagement,
		Schema: map[string]*schema.Schema{
			// Required
			"user_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			// Optional
			"can_use_api_keys": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"can_use_auth_tokens": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"can_use_console_password": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"can_use_customer_secret_keys": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"can_use_smtp_credentials": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},

			// Computed
		},
	}
}

func createUserCapabilitiesManagement(d *schema.ResourceData, m interface{}) error {
	sync := &UserCapabilitiesManagementResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).IdentityClient()
	sync.Configuration = m.(*client.OracleClients).Configuration

	return tfresource.CreateResource(d, sync)
}

func readUserCapabilitiesManagement(d *schema.ResourceData, m interface{}) error {
	sync := &UserCapabilitiesManagementResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).IdentityClient()
	return tfresource.ReadResource(sync)
}

func updateUserCapabilitiesManagement(d *schema.ResourceData, m interface{}) error {
	sync := &UserCapabilitiesManagementResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).IdentityClient()
	return tfresource.UpdateResource(d, sync)
}

func deleteUserCapabilitiesManagement(d *schema.ResourceData, m interface{}) error {
	sync := &UserCapabilitiesManagementResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).IdentityClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type UserCapabilitiesManagementResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_identity.IdentityClient
	Configuration          map[string]string
	Res                    *oci_identity.User
	DisableNotFoundRetries bool
}

func getUserCapabilitiesCompositeId(userId string) string {
	return "capabilities" + "/" + userId
}

func parseUserCapabilitiesCompositeId(compositeId string) (userId string, err error) {
	parts := strings.Split(compositeId, "/")
	match, _ := regexp.MatchString("capabilities/.*", compositeId)
	if !match || len(parts) != 2 {
		err = fmt.Errorf("illegal compositeId %s encountered", compositeId)
		return
	}
	userId, _ = url.PathUnescape(parts[1])

	return
}

func (s *UserCapabilitiesManagementResourceCrud) ID() string {
	return getUserCapabilitiesCompositeId(s.D.Get("user_id").(string))
}

func (s *UserCapabilitiesManagementResourceCrud) Create() error {
	userCapabilityRequest := oci_identity.UpdateUserCapabilitiesRequest{}

	if userId, ok := s.D.GetOkExists("user_id"); ok {
		tmp := userId.(string)
		userCapabilityRequest.UserId = &tmp
	}

	if b, ok := s.D.GetOkExists("can_use_api_keys"); ok {
		canUseApiKeys := b.(bool)
		userCapabilityRequest.CanUseApiKeys = &canUseApiKeys
	}

	if b, ok := s.D.GetOkExists("can_use_auth_tokens"); ok {
		canUseAuthTokens := b.(bool)
		userCapabilityRequest.CanUseAuthTokens = &canUseAuthTokens
	}

	if b, ok := s.D.GetOkExists("can_use_console_password"); ok {
		canUseConsolePassword := b.(bool)
		userCapabilityRequest.CanUseConsolePassword = &canUseConsolePassword
	}

	if b, ok := s.D.GetOkExists("can_use_customer_secret_keys"); ok {
		canUseCustomerSecretKeys := b.(bool)
		userCapabilityRequest.CanUseCustomerSecretKeys = &canUseCustomerSecretKeys
	}

	if b, ok := s.D.GetOkExists("can_use_smtp_credentials"); ok {
		canUseSmtpCredentials := b.(bool)
		userCapabilityRequest.CanUseSmtpCredentials = &canUseSmtpCredentials
	}

	userCapabilityResponse, err := s.Client.UpdateUserCapabilities(context.Background(), userCapabilityRequest)
	if err != nil {
		return err
	}

	s.Res = &userCapabilityResponse.User
	return nil
}

func (s *UserCapabilitiesManagementResourceCrud) Get() error {
	request := oci_identity.GetUserRequest{}

	userId, err := parseUserCapabilitiesCompositeId(s.D.Id())
	if err == nil {
		request.UserId = &userId
	} else {
		log.Printf("[WARN] Get() unable to parse current ID: %s with err %v", s.D.Id(), err)
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "identity")

	response, err := s.Client.GetUser(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.User
	return nil
}

func (s *UserCapabilitiesManagementResourceCrud) Update() error {
	return s.Create()
}

func (s *UserCapabilitiesManagementResourceCrud) Delete() error {
	// We dont know the original default values, so we pretend that the existing state is to be retained
	return nil
}

func (s *UserCapabilitiesManagementResourceCrud) SetData() error {
	if s.Res.Capabilities != nil {
		if s.Res.Capabilities.CanUseApiKeys != nil {
			s.D.Set("can_use_api_keys", *s.Res.Capabilities.CanUseApiKeys)
		}
		if s.Res.Capabilities.CanUseAuthTokens != nil {
			s.D.Set("can_use_auth_tokens", *s.Res.Capabilities.CanUseAuthTokens)
		}
		if s.Res.Capabilities.CanUseConsolePassword != nil {
			s.D.Set("can_use_console_password", *s.Res.Capabilities.CanUseConsolePassword)
		}
		if s.Res.Capabilities.CanUseCustomerSecretKeys != nil {
			s.D.Set("can_use_customer_secret_keys", *s.Res.Capabilities.CanUseCustomerSecretKeys)
		}
		if s.Res.Capabilities.CanUseSmtpCredentials != nil {
			s.D.Set("can_use_smtp_credentials", *s.Res.Capabilities.CanUseSmtpCredentials)
		}
	}
	if s.Res.Id != nil {
		s.D.Set("user_id", *s.Res.Id)
	}

	return nil
}
