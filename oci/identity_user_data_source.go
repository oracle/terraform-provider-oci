// Copyright (c) 2017, 2019, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"context"
	"strconv"

	"github.com/hashicorp/terraform/helper/schema"
	oci_identity "github.com/oracle/oci-go-sdk/identity"
)

func IdentityUserDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularIdentityUser,
		Schema: map[string]*schema.Schema{
			"user_id": {
				Type:     schema.TypeString,
				Required: true,
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
						"can_use_smtp_credentials": {
							Type:     schema.TypeBool,
							Computed: true,
						},
					},
				},
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"defined_tags": {
				Type:     schema.TypeMap,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"description": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"email": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"external_identifier": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"freeform_tags": {
				Type:     schema.TypeMap,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"identity_provider_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"inactive_state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"name": {
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

func readSingularIdentityUser(d *schema.ResourceData, m interface{}) error {
	sync := &IdentityUserDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).identityClient

	return ReadResource(sync)
}

type IdentityUserDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_identity.IdentityClient
	Res    *oci_identity.GetUserResponse
}

func (s *IdentityUserDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *IdentityUserDataSourceCrud) Get() error {
	request := oci_identity.GetUserRequest{}

	if userId, ok := s.D.GetOkExists("user_id"); ok {
		tmp := userId.(string)
		request.UserId = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(false, "identity")

	response, err := s.Client.GetUser(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *IdentityUserDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.Capabilities != nil {
		s.D.Set("capabilities", []interface{}{UserCapabilitiesToMap(s.Res.Capabilities)})
	} else {
		s.D.Set("capabilities", nil)
	}

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", definedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.Description != nil {
		s.D.Set("description", *s.Res.Description)
	}

	if s.Res.Email != nil {
		s.D.Set("email", *s.Res.Email)
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

	if s.Res.Name != nil {
		s.D.Set("name", *s.Res.Name)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	return nil
}
