// Copyright (c) 2017, 2019, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"context"

	"github.com/hashicorp/terraform/helper/schema"
	oci_identity "github.com/oracle/oci-go-sdk/identity"
)

func IdentityAuthenticationPolicyDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularIdentityAuthenticationPolicy,
		Schema: map[string]*schema.Schema{
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			// Computed
			"password_policy": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional
						"is_lowercase_characters_required": {
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
						},
						"is_numeric_characters_required": {
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
						},
						"is_special_characters_required": {
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
						},
						"is_uppercase_characters_required": {
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
						},
						"is_username_containment_allowed": {
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
						},
						"minimum_password_length": {
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},

						// Computed
					},
				},
			},
		},
	}
}

func readSingularIdentityAuthenticationPolicy(d *schema.ResourceData, m interface{}) error {
	sync := &IdentityAuthenticationPolicyDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).identityClient

	return ReadResource(sync)
}

type IdentityAuthenticationPolicyDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_identity.IdentityClient
	Res    *oci_identity.GetAuthenticationPolicyResponse
}

func (s *IdentityAuthenticationPolicyDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *IdentityAuthenticationPolicyDataSourceCrud) Get() error {
	request := oci_identity.GetAuthenticationPolicyRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(false, "identity")

	response, err := s.Client.GetAuthenticationPolicy(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *IdentityAuthenticationPolicyDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(GenerateDataSourceID())

	if s.Res.PasswordPolicy != nil {
		s.D.Set("password_policy", []interface{}{PasswordPolicyToMap(s.Res.PasswordPolicy)})
	} else {
		s.D.Set("password_policy", nil)
	}

	return nil
}
