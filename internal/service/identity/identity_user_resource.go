// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package identity

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/oracle/terraform-provider-oci/internal/globalvar"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"

	oci_identity "github.com/oracle/oci-go-sdk/v65/identity"
)

func IdentityUserResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createIdentityUser,
		Read:     readIdentityUser,
		Update:   updateIdentityUser,
		Delete:   deleteIdentityUser,
		Schema: map[string]*schema.Schema{
			// The legacy provider exposed this as read-only/computed. The API requires this param. For legacy users who are
			// not supplying a value, make it optional, behind the scenes it will use the tenancy ocid if not supplied.
			// If a user supplies the value, then changes it, it requires forcing new.
			"compartment_id": {
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
				ForceNew: true,
			},
			// Required
			"description": {
				Type:     schema.TypeString,
				Required: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			// Optional
			"defined_tags": {
				Type:             schema.TypeMap,
				Optional:         true,
				Computed:         true,
				DiffSuppressFunc: tfresource.DefinedTagsDiffSuppressFunction,
				Elem:             schema.TypeString,
			},
			"email": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"freeform_tags": {
				Type:     schema.TypeMap,
				Optional: true,
				Computed: true,
				Elem:     schema.TypeString,
			},

			// Computed
			"capabilities": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"can_use_api_keys": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"can_use_auth_tokens": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"can_use_console_password": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"can_use_customer_secret_keys": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"can_use_db_credentials": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"can_use_oauth2client_credentials": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"can_use_smtp_credentials": {
							Type:     schema.TypeBool,
							Computed: true,
						},
					},
				},
			},
			"db_user_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"email_verified": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"external_identifier": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"identity_provider_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"inactive_state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"last_successful_login_time": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"previous_successful_login_time": {
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

func createIdentityUser(d *schema.ResourceData, m interface{}) error {
	sync := &IdentityUserResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).IdentityClient()
	sync.Configuration = m.(*client.OracleClients).Configuration

	return tfresource.CreateResource(d, sync)
}

func readIdentityUser(d *schema.ResourceData, m interface{}) error {
	sync := &IdentityUserResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).IdentityClient()

	return tfresource.ReadResource(sync)
}

func updateIdentityUser(d *schema.ResourceData, m interface{}) error {
	sync := &IdentityUserResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).IdentityClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteIdentityUser(d *schema.ResourceData, m interface{}) error {
	sync := &IdentityUserResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).IdentityClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type IdentityUserResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_identity.IdentityClient
	Configuration          map[string]string
	Res                    *oci_identity.User
	DisableNotFoundRetries bool
}

func (s *IdentityUserResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *IdentityUserResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_identity.UserLifecycleStateCreating),
	}
}

func (s *IdentityUserResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_identity.UserLifecycleStateActive),
	}
}

func (s *IdentityUserResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_identity.UserLifecycleStateDeleting),
	}
}

func (s *IdentityUserResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_identity.UserLifecycleStateDeleted),
	}
}

func (s *IdentityUserResourceCrud) Create() error {
	request := oci_identity.CreateUserRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	} else { // @next-break: remove
		// Prevent potentially inferring wrong TenancyOCID from InstancePrincipal
		if auth := s.Configuration["auth"]; strings.ToLower(auth) == strings.ToLower(globalvar.AuthInstancePrincipalSetting) {
			return fmt.Errorf("compartment_id must be specified for this resource")
		}
		// Maintain legacy contract of compartment_id defaulting to tenancy ocid if not specified
		configProvider := s.Client.ConfigurationProvider()
		if configProvider == nil {
			return fmt.Errorf("cannot access tenancy OCID. No configuration provider could be found for identity client")
		}

		c := *configProvider
		tenancy, err := c.TenancyOCID()
		if err != nil {
			return err
		}
		request.CompartmentId = &tenancy
	}

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
		convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return err
		}
		request.DefinedTags = convertedDefinedTags
	}

	if description, ok := s.D.GetOkExists("description"); ok {
		tmp := description.(string)
		request.Description = &tmp
	}

	if email, ok := s.D.GetOkExists("email"); ok {
		tmp := email.(string)
		request.Email = &tmp
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if name, ok := s.D.GetOkExists("name"); ok {
		tmp := name.(string)
		request.Name = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "identity")

	response, err := s.Client.CreateUser(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.User
	return nil
}

func (s *IdentityUserResourceCrud) Get() error {
	request := oci_identity.GetUserRequest{}

	tmp := s.D.Id()
	request.UserId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "identity")

	response, err := s.Client.GetUser(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.User
	return nil
}

func (s *IdentityUserResourceCrud) Update() error {
	request := oci_identity.UpdateUserRequest{}

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
		convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return err
		}
		request.DefinedTags = convertedDefinedTags
	}

	if description, ok := s.D.GetOkExists("description"); ok {
		tmp := description.(string)
		request.Description = &tmp
	}

	if email, ok := s.D.GetOkExists("email"); ok {
		tmp := email.(string)
		request.Email = &tmp
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	tmp := s.D.Id()
	request.UserId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "identity")

	response, err := s.Client.UpdateUser(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.User
	return nil
}

func (s *IdentityUserResourceCrud) Delete() error {
	request := oci_identity.DeleteUserRequest{}

	tmp := s.D.Id()
	request.UserId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "identity")

	_, err := s.Client.DeleteUser(context.Background(), request)
	return err
}

func (s *IdentityUserResourceCrud) SetData() error {
	if s.Res.Capabilities != nil {
		s.D.Set("capabilities", []interface{}{UserCapabilitiesToMap(s.Res.Capabilities)})
	} else {
		s.D.Set("capabilities", nil)
	}

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DbUserName != nil {
		s.D.Set("db_user_name", *s.Res.DbUserName)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.Description != nil {
		s.D.Set("description", *s.Res.Description)
	}

	if s.Res.Email != nil {
		s.D.Set("email", *s.Res.Email)
	}

	if s.Res.EmailVerified != nil {
		s.D.Set("email_verified", *s.Res.EmailVerified)
	}

	if s.Res.ExternalIdentifier != nil {
		s.D.Set("external_identifier", *s.Res.ExternalIdentifier)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.IdentityProviderId != nil {
		s.D.Set("identity_provider_id", *s.Res.IdentityProviderId)
	}

	if s.Res.InactiveStatus != nil {
		s.D.Set("inactive_state", strconv.FormatInt(*s.Res.InactiveStatus, 10))
	}

	if s.Res.LastSuccessfulLoginTime != nil {
		s.D.Set("last_successful_login_time", s.Res.LastSuccessfulLoginTime.String())
	}

	if s.Res.Name != nil {
		s.D.Set("name", *s.Res.Name)
	}

	if s.Res.PreviousSuccessfulLoginTime != nil {
		s.D.Set("previous_successful_login_time", s.Res.PreviousSuccessfulLoginTime.String())
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	return nil
}

func UserCapabilitiesToMap(obj *oci_identity.UserCapabilities) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.CanUseApiKeys != nil {
		result["can_use_api_keys"] = bool(*obj.CanUseApiKeys)
	}

	if obj.CanUseAuthTokens != nil {
		result["can_use_auth_tokens"] = bool(*obj.CanUseAuthTokens)
	}

	if obj.CanUseConsolePassword != nil {
		result["can_use_console_password"] = bool(*obj.CanUseConsolePassword)
	}

	if obj.CanUseCustomerSecretKeys != nil {
		result["can_use_customer_secret_keys"] = bool(*obj.CanUseCustomerSecretKeys)
	}

	if obj.CanUseDbCredentials != nil {
		result["can_use_db_credentials"] = bool(*obj.CanUseDbCredentials)
	}

	if obj.CanUseOAuth2ClientCredentials != nil {
		result["can_use_oauth2client_credentials"] = bool(*obj.CanUseOAuth2ClientCredentials)
	}

	if obj.CanUseSmtpCredentials != nil {
		result["can_use_smtp_credentials"] = bool(*obj.CanUseSmtpCredentials)
	}

	return result
}
