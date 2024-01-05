// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package identity

import (
	"context"
	"fmt"
	"log"
	"net/url"
	"regexp"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"

	oci_identity "github.com/oracle/oci-go-sdk/v65/identity"
)

func IdentityAuthenticationPolicyResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createIdentityAuthenticationPolicy,
		Read:     readIdentityAuthenticationPolicy,
		Update:   updateIdentityAuthenticationPolicy,
		Delete:   deleteIdentityAuthenticationPolicy,
		Schema: map[string]*schema.Schema{
			// Required
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			// Optional
			"network_policy": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional
						"network_source_ids": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},

						// Computed
					},
				},
			},
			"password_policy": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				MaxItems: 1,
				MinItems: 1,
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

			// Computed
		},
	}
}

func createIdentityAuthenticationPolicy(d *schema.ResourceData, m interface{}) error {
	sync := &IdentityAuthenticationPolicyResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).IdentityClient()

	return tfresource.CreateResource(d, sync)
}

func readIdentityAuthenticationPolicy(d *schema.ResourceData, m interface{}) error {
	sync := &IdentityAuthenticationPolicyResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).IdentityClient()

	return tfresource.ReadResource(sync)
}

func updateIdentityAuthenticationPolicy(d *schema.ResourceData, m interface{}) error {
	sync := &IdentityAuthenticationPolicyResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).IdentityClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteIdentityAuthenticationPolicy(d *schema.ResourceData, m interface{}) error {
	return nil
}

type IdentityAuthenticationPolicyResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_identity.IdentityClient
	Res                    *oci_identity.AuthenticationPolicy
	DisableNotFoundRetries bool
}

func (s *IdentityAuthenticationPolicyResourceCrud) ID() string {
	return GetAuthenticationPolicyCompositeId(s.D.Get("compartment_id").(string))
}

func (s *IdentityAuthenticationPolicyResourceCrud) Create() error {
	request := oci_identity.UpdateAuthenticationPolicyRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if networkPolicy, ok := s.D.GetOkExists("network_policy"); ok {
		if tmpList := networkPolicy.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "network_policy", 0)
			tmp, err := s.mapToNetworkPolicy(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.NetworkPolicy = &tmp
		}
	}

	if passwordPolicy, ok := s.D.GetOkExists("password_policy"); ok {
		if tmpList := passwordPolicy.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "password_policy", 0)
			tmp, err := s.mapToPasswordPolicy(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.PasswordPolicy = &tmp
		}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "identity")

	response, err := s.Client.UpdateAuthenticationPolicy(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.AuthenticationPolicy
	return nil
}

func (s *IdentityAuthenticationPolicyResourceCrud) Get() error {
	request := oci_identity.GetAuthenticationPolicyRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	compartmentId, err := parseAuthenticationPolicyCompositeId(s.D.Id())
	if err == nil {
		request.CompartmentId = &compartmentId
	} else {
		log.Printf("[WARN] Get() unable to parse current ID: %s", s.D.Id())
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "identity")

	response, err := s.Client.GetAuthenticationPolicy(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.AuthenticationPolicy
	return nil
}

func (s *IdentityAuthenticationPolicyResourceCrud) Update() error {
	request := oci_identity.UpdateAuthenticationPolicyRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if networkPolicy, ok := s.D.GetOkExists("network_policy"); ok {
		if tmpList := networkPolicy.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "network_policy", 0)
			tmp, err := s.mapToNetworkPolicy(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.NetworkPolicy = &tmp
		}
	}

	if passwordPolicy, ok := s.D.GetOkExists("password_policy"); ok {
		if tmpList := passwordPolicy.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "password_policy", 0)
			tmp, err := s.mapToPasswordPolicy(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.PasswordPolicy = &tmp
		}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "identity")

	response, err := s.Client.UpdateAuthenticationPolicy(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.AuthenticationPolicy
	return nil
}

func (s *IdentityAuthenticationPolicyResourceCrud) SetData() error {

	compartmentId, err := parseAuthenticationPolicyCompositeId(s.D.Id())
	if err == nil {
		s.D.Set("compartment_id", &compartmentId)
	} else {
		log.Printf("[WARN] SetData() unable to parse current ID: %s", s.D.Id())
	}

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.NetworkPolicy != nil {
		s.D.Set("network_policy", []interface{}{NetworkPolicyToMap(s.Res.NetworkPolicy)})
	} else {
		s.D.Set("network_policy", nil)
	}

	if s.Res.PasswordPolicy != nil {
		s.D.Set("password_policy", []interface{}{PasswordPolicyToMap(s.Res.PasswordPolicy)})
	} else {
		s.D.Set("password_policy", nil)
	}

	return nil
}

func GetAuthenticationPolicyCompositeId(compartmentId string) string {
	compartmentId = url.PathEscape(compartmentId)
	compositeId := "authenticationPolicies/" + compartmentId
	return compositeId
}

func parseAuthenticationPolicyCompositeId(compositeId string) (compartmentId string, err error) {
	parts := strings.Split(compositeId, "/")
	match, _ := regexp.MatchString("authenticationPolicies/.*", compositeId)
	if !match || len(parts) != 2 {
		err = fmt.Errorf("illegal compositeId %s encountered", compositeId)
		return
	}
	compartmentId, _ = url.PathUnescape(parts[1])

	return
}

func (s *IdentityAuthenticationPolicyResourceCrud) mapToNetworkPolicy(fieldKeyFormat string) (oci_identity.NetworkPolicy, error) {
	result := oci_identity.NetworkPolicy{}

	if networkSourceIds, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "network_source_ids")); ok {
		interfaces := networkSourceIds.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "network_source_ids")) {
			result.NetworkSourceIds = tmp
		}
	}

	return result, nil
}

func NetworkPolicyToMap(obj *oci_identity.NetworkPolicy) map[string]interface{} {
	result := map[string]interface{}{}

	result["network_source_ids"] = obj.NetworkSourceIds

	return result
}

func (s *IdentityAuthenticationPolicyResourceCrud) mapToPasswordPolicy(fieldKeyFormat string) (oci_identity.PasswordPolicy, error) {
	result := oci_identity.PasswordPolicy{}

	if isLowercaseCharactersRequired, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_lowercase_characters_required")); ok {
		tmp := isLowercaseCharactersRequired.(bool)
		result.IsLowercaseCharactersRequired = &tmp
	}

	if isNumericCharactersRequired, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_numeric_characters_required")); ok {
		tmp := isNumericCharactersRequired.(bool)
		result.IsNumericCharactersRequired = &tmp
	}

	if isSpecialCharactersRequired, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_special_characters_required")); ok {
		tmp := isSpecialCharactersRequired.(bool)
		result.IsSpecialCharactersRequired = &tmp
	}

	if isUppercaseCharactersRequired, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_uppercase_characters_required")); ok {
		tmp := isUppercaseCharactersRequired.(bool)
		result.IsUppercaseCharactersRequired = &tmp
	}

	if isUsernameContainmentAllowed, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_username_containment_allowed")); ok {
		tmp := isUsernameContainmentAllowed.(bool)
		result.IsUsernameContainmentAllowed = &tmp
	}

	if minimumPasswordLength, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "minimum_password_length")); ok {
		tmp := minimumPasswordLength.(int)
		result.MinimumPasswordLength = &tmp
	}

	return result, nil
}

func PasswordPolicyToMap(obj *oci_identity.PasswordPolicy) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.IsLowercaseCharactersRequired != nil {
		result["is_lowercase_characters_required"] = bool(*obj.IsLowercaseCharactersRequired)
	}

	if obj.IsNumericCharactersRequired != nil {
		result["is_numeric_characters_required"] = bool(*obj.IsNumericCharactersRequired)
	}

	if obj.IsSpecialCharactersRequired != nil {
		result["is_special_characters_required"] = bool(*obj.IsSpecialCharactersRequired)
	}

	if obj.IsUppercaseCharactersRequired != nil {
		result["is_uppercase_characters_required"] = bool(*obj.IsUppercaseCharactersRequired)
	}

	if obj.IsUsernameContainmentAllowed != nil {
		result["is_username_containment_allowed"] = bool(*obj.IsUsernameContainmentAllowed)
	}

	if obj.MinimumPasswordLength != nil {
		result["minimum_password_length"] = int(*obj.MinimumPasswordLength)
	}

	return result
}
