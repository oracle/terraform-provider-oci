// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package identity_domains

import (
	"context"
	"fmt"
	"log"
	"net/url"
	"regexp"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	oci_identity_domains "github.com/oracle/oci-go-sdk/v65/identitydomains"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func IdentityDomainsPasswordPolicyResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createIdentityDomainsPasswordPolicy,
		Read:     readIdentityDomainsPasswordPolicy,
		Update:   updateIdentityDomainsPasswordPolicy,
		Delete:   deleteIdentityDomainsPasswordPolicy,
		Schema: map[string]*schema.Schema{
			// Required
			"idcs_endpoint": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"schemas": {
				Type:     schema.TypeList,
				Required: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},

			// Optional
			"allowed_chars": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"attribute_sets": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"attributes": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"authorization": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dictionary_delimiter": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dictionary_location": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dictionary_word_disallowed": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"disallowed_chars": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"disallowed_substrings": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"disallowed_user_attribute_values": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"distinct_characters": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"external_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"first_name_disallowed": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"force_password_reset": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"groups": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"value": {
							Type:     schema.TypeString,
							Required: true,
						},

						// Optional

						// Computed
						"display": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"ref": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"last_name_disallowed": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"lockout_duration": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"max_incorrect_attempts": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"max_length": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"max_repeated_chars": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"max_special_chars": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"min_alpha_numerals": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"min_alphas": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"min_length": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"min_lower_case": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"min_numerals": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"min_password_age": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"min_special_chars": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"min_unique_chars": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"min_upper_case": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"num_passwords_in_history": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"ocid": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"password_expire_warning": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"password_expires_after": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"password_strength": {
				Type:      schema.TypeString,
				Optional:  true,
				Computed:  true,
				Sensitive: true,
			},
			"priority": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"required_chars": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"resource_type_schema_version": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"starts_with_alphabet": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"tags": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"key": {
							Type:     schema.TypeString,
							Required: true,
						},
						"value": {
							Type:     schema.TypeString,
							Required: true,
						},

						// Optional

						// Computed
					},
				},
			},
			"user_name_disallowed": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},

			// Computed
			"compartment_ocid": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"configured_password_policy_rules": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"key": {
							Type:     schema.TypeString,
							Required: true,
						},
						"value": {
							Type:     schema.TypeString,
							Required: true,
						},

						// Optional

						// Computed
					},
				},
			},
			"delete_in_progress": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"domain_ocid": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"idcs_created_by": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"value": {
							Type:     schema.TypeString,
							Required: true,
						},

						// Optional
						"display": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"ocid": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"type": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},

						// Computed
						"ref": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"idcs_last_modified_by": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"value": {
							Type:     schema.TypeString,
							Required: true,
						},

						// Optional
						"display": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"ocid": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"type": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},

						// Computed
						"ref": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"idcs_last_upgraded_in_release": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"idcs_prevented_operations": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"meta": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional
						"created": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"last_modified": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"location": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"resource_type": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"version": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},

						// Computed
					},
				},
			},
			"tenancy_ocid": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func createIdentityDomainsPasswordPolicy(d *schema.ResourceData, m interface{}) error {
	sync := &IdentityDomainsPasswordPolicyResourceCrud{}
	sync.D = d
	idcsEndpoint, err := getIdcsEndpoint(d)
	if err != nil {
		return err
	}
	client, err := m.(*client.OracleClients).IdentityDomainsClientWithEndpoint(idcsEndpoint)
	if err != nil {
		return err
	}
	sync.Client = client
	return tfresource.CreateResource(d, sync)
}

func readIdentityDomainsPasswordPolicy(d *schema.ResourceData, m interface{}) error {
	sync := &IdentityDomainsPasswordPolicyResourceCrud{}
	sync.D = d
	idcsEndpoint, err := getIdcsEndpointForRead(d, "passwordPolicies")
	if err != nil {
		return err
	}
	client, err := m.(*client.OracleClients).IdentityDomainsClientWithEndpoint(idcsEndpoint)
	if err != nil {
		return err
	}
	sync.Client = client
	return tfresource.ReadResource(sync)
}

func updateIdentityDomainsPasswordPolicy(d *schema.ResourceData, m interface{}) error {
	sync := &IdentityDomainsPasswordPolicyResourceCrud{}
	sync.D = d
	idcsEndpoint, err := getIdcsEndpoint(d)
	if err != nil {
		return err
	}
	client, err := m.(*client.OracleClients).IdentityDomainsClientWithEndpoint(idcsEndpoint)
	if err != nil {
		return err
	}
	sync.Client = client
	return tfresource.UpdateResource(d, sync)
}

func deleteIdentityDomainsPasswordPolicy(d *schema.ResourceData, m interface{}) error {
	sync := &IdentityDomainsPasswordPolicyResourceCrud{}
	sync.D = d
	idcsEndpoint, err := getIdcsEndpoint(d)
	if err != nil {
		return err
	}
	client, err := m.(*client.OracleClients).IdentityDomainsClientWithEndpoint(idcsEndpoint)
	if err != nil {
		return err
	}
	sync.Client = client
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type IdentityDomainsPasswordPolicyResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_identity_domains.IdentityDomainsClient
	Res                    *oci_identity_domains.PasswordPolicy
	DisableNotFoundRetries bool
}

func (s *IdentityDomainsPasswordPolicyResourceCrud) ID() string {
	return *s.Res.Id
	//return GetPasswordPolicyCompositeId(s.D.Get("id").(string))
}

func (s *IdentityDomainsPasswordPolicyResourceCrud) Create() error {
	request := oci_identity_domains.CreatePasswordPolicyRequest{}

	if allowedChars, ok := s.D.GetOkExists("allowed_chars"); ok {
		tmp := allowedChars.(string)
		request.AllowedChars = &tmp
	}

	if attributeSets, ok := s.D.GetOkExists("attribute_sets"); ok {
		interfaces := attributeSets.([]interface{})
		tmp := make([]oci_identity_domains.AttributeSetsEnum, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = oci_identity_domains.AttributeSetsEnum(interfaces[i].(string))
			}
		}
		if len(tmp) != 0 || s.D.HasChange("attribute_sets") {
			request.AttributeSets = tmp
		}
	}

	if attributes, ok := s.D.GetOkExists("attributes"); ok {
		tmp := attributes.(string)
		request.Attributes = &tmp
	}

	if authorization, ok := s.D.GetOkExists("authorization"); ok {
		tmp := authorization.(string)
		request.Authorization = &tmp
	}

	if description, ok := s.D.GetOkExists("description"); ok {
		tmp := description.(string)
		request.Description = &tmp
	}

	if dictionaryDelimiter, ok := s.D.GetOkExists("dictionary_delimiter"); ok {
		tmp := dictionaryDelimiter.(string)
		request.DictionaryDelimiter = &tmp
	}

	if dictionaryLocation, ok := s.D.GetOkExists("dictionary_location"); ok {
		tmp := dictionaryLocation.(string)
		request.DictionaryLocation = &tmp
	}

	if dictionaryWordDisallowed, ok := s.D.GetOkExists("dictionary_word_disallowed"); ok {
		tmp := dictionaryWordDisallowed.(bool)
		request.DictionaryWordDisallowed = &tmp
	}

	if disallowedChars, ok := s.D.GetOkExists("disallowed_chars"); ok {
		tmp := disallowedChars.(string)
		request.DisallowedChars = &tmp
	}

	if disallowedSubstrings, ok := s.D.GetOkExists("disallowed_substrings"); ok {
		interfaces := disallowedSubstrings.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("disallowed_substrings") {
			request.DisallowedSubstrings = tmp
		}
	}

	if disallowedUserAttributeValues, ok := s.D.GetOkExists("disallowed_user_attribute_values"); ok {
		interfaces := disallowedUserAttributeValues.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("disallowed_user_attribute_values") {
			request.DisallowedUserAttributeValues = tmp
		}
	}

	if distinctCharacters, ok := s.D.GetOkExists("distinct_characters"); ok {
		tmp := distinctCharacters.(int)
		request.DistinctCharacters = &tmp
	}

	if externalId, ok := s.D.GetOkExists("external_id"); ok {
		tmp := externalId.(string)
		request.ExternalId = &tmp
	}

	if firstNameDisallowed, ok := s.D.GetOkExists("first_name_disallowed"); ok {
		tmp := firstNameDisallowed.(bool)
		request.FirstNameDisallowed = &tmp
	}

	if forcePasswordReset, ok := s.D.GetOkExists("force_password_reset"); ok {
		tmp := forcePasswordReset.(bool)
		request.ForcePasswordReset = &tmp
	}

	if groups, ok := s.D.GetOkExists("groups"); ok {
		interfaces := groups.([]interface{})
		tmp := make([]oci_identity_domains.PasswordPolicyGroups, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "groups", stateDataIndex)
			converted, err := s.mapToPasswordPolicyGroups(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("groups") {
			request.Groups = tmp
		}
	}

	if id, ok := s.D.GetOkExists("id"); ok {
		tmp := id.(string)
		request.Id = &tmp
	}

	if lastNameDisallowed, ok := s.D.GetOkExists("last_name_disallowed"); ok {
		tmp := lastNameDisallowed.(bool)
		request.LastNameDisallowed = &tmp
	}

	if lockoutDuration, ok := s.D.GetOkExists("lockout_duration"); ok {
		tmp := lockoutDuration.(int)
		request.LockoutDuration = &tmp
	}

	if maxIncorrectAttempts, ok := s.D.GetOkExists("max_incorrect_attempts"); ok {
		tmp := maxIncorrectAttempts.(int)
		request.MaxIncorrectAttempts = &tmp
	}

	if maxLength, ok := s.D.GetOkExists("max_length"); ok {
		tmp := maxLength.(int)
		request.MaxLength = &tmp
	}

	if maxRepeatedChars, ok := s.D.GetOkExists("max_repeated_chars"); ok {
		tmp := maxRepeatedChars.(int)
		request.MaxRepeatedChars = &tmp
	}

	if maxSpecialChars, ok := s.D.GetOkExists("max_special_chars"); ok {
		tmp := maxSpecialChars.(int)
		request.MaxSpecialChars = &tmp
	}

	if minAlphaNumerals, ok := s.D.GetOkExists("min_alpha_numerals"); ok {
		tmp := minAlphaNumerals.(int)
		request.MinAlphaNumerals = &tmp
	}

	if minAlphas, ok := s.D.GetOkExists("min_alphas"); ok {
		tmp := minAlphas.(int)
		request.MinAlphas = &tmp
	}

	if minLength, ok := s.D.GetOkExists("min_length"); ok {
		tmp := minLength.(int)
		request.MinLength = &tmp
	}

	if minLowerCase, ok := s.D.GetOkExists("min_lower_case"); ok {
		tmp := minLowerCase.(int)
		request.MinLowerCase = &tmp
	}

	if minNumerals, ok := s.D.GetOkExists("min_numerals"); ok {
		tmp := minNumerals.(int)
		request.MinNumerals = &tmp
	}

	if minPasswordAge, ok := s.D.GetOkExists("min_password_age"); ok {
		tmp := minPasswordAge.(int)
		request.MinPasswordAge = &tmp
	}

	if minSpecialChars, ok := s.D.GetOkExists("min_special_chars"); ok {
		tmp := minSpecialChars.(int)
		request.MinSpecialChars = &tmp
	}

	if minUniqueChars, ok := s.D.GetOkExists("min_unique_chars"); ok {
		tmp := minUniqueChars.(int)
		request.MinUniqueChars = &tmp
	}

	if minUpperCase, ok := s.D.GetOkExists("min_upper_case"); ok {
		tmp := minUpperCase.(int)
		request.MinUpperCase = &tmp
	}

	if name, ok := s.D.GetOkExists("name"); ok {
		tmp := name.(string)
		request.Name = &tmp
	}

	if numPasswordsInHistory, ok := s.D.GetOkExists("num_passwords_in_history"); ok {
		tmp := numPasswordsInHistory.(int)
		request.NumPasswordsInHistory = &tmp
	}

	if ocid, ok := s.D.GetOkExists("ocid"); ok {
		tmp := ocid.(string)
		request.Ocid = &tmp
	}

	if passwordExpireWarning, ok := s.D.GetOkExists("password_expire_warning"); ok {
		tmp := passwordExpireWarning.(int)
		request.PasswordExpireWarning = &tmp
	}

	if passwordExpiresAfter, ok := s.D.GetOkExists("password_expires_after"); ok {
		tmp := passwordExpiresAfter.(int)
		request.PasswordExpiresAfter = &tmp
	}

	if passwordStrength, ok := s.D.GetOkExists("password_strength"); ok {
		request.PasswordStrength = oci_identity_domains.PasswordPolicyPasswordStrengthEnum(passwordStrength.(string))
	}

	if priority, ok := s.D.GetOkExists("priority"); ok {
		tmp := priority.(int)
		request.Priority = &tmp
	}

	if requiredChars, ok := s.D.GetOkExists("required_chars"); ok {
		tmp := requiredChars.(string)
		request.RequiredChars = &tmp
	}

	if resourceTypeSchemaVersion, ok := s.D.GetOkExists("resource_type_schema_version"); ok {
		tmp := resourceTypeSchemaVersion.(string)
		request.ResourceTypeSchemaVersion = &tmp
	}

	if schemas, ok := s.D.GetOkExists("schemas"); ok {
		interfaces := schemas.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("schemas") {
			request.Schemas = tmp
		}
	}

	if startsWithAlphabet, ok := s.D.GetOkExists("starts_with_alphabet"); ok {
		tmp := startsWithAlphabet.(bool)
		request.StartsWithAlphabet = &tmp
	}

	if tags, ok := s.D.GetOkExists("tags"); ok {
		interfaces := tags.([]interface{})
		tmp := make([]oci_identity_domains.Tags, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "tags", stateDataIndex)
			converted, err := s.mapTotags(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("tags") {
			request.Tags = tmp
		}
	}

	if userNameDisallowed, ok := s.D.GetOkExists("user_name_disallowed"); ok {
		tmp := userNameDisallowed.(bool)
		request.UserNameDisallowed = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "identity_domains")

	response, err := s.Client.CreatePasswordPolicy(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.PasswordPolicy
	return nil
}

func (s *IdentityDomainsPasswordPolicyResourceCrud) Get() error {
	request := oci_identity_domains.GetPasswordPolicyRequest{}

	if attributeSets, ok := s.D.GetOkExists("attribute_sets"); ok {
		interfaces := attributeSets.([]interface{})
		tmp := make([]oci_identity_domains.AttributeSetsEnum, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = oci_identity_domains.AttributeSetsEnum(interfaces[i].(string))
			}
		}
		if len(tmp) != 0 || s.D.HasChange("attribute_sets") {
			request.AttributeSets = tmp
		}
	}

	if attributes, ok := s.D.GetOkExists("attributes"); ok {
		tmp := attributes.(string)
		request.Attributes = &tmp
	}

	if authorization, ok := s.D.GetOkExists("authorization"); ok {
		tmp := authorization.(string)
		request.Authorization = &tmp
	}

	tmp := s.D.Id()
	request.PasswordPolicyId = &tmp

	if resourceTypeSchemaVersion, ok := s.D.GetOkExists("resource_type_schema_version"); ok {
		tmp := resourceTypeSchemaVersion.(string)
		request.ResourceTypeSchemaVersion = &tmp
	}

	passwordPolicyId, err := parsePasswordPolicyCompositeId(s.D.Id())
	if err == nil {
		request.PasswordPolicyId = &passwordPolicyId
	} else {
		log.Printf("[WARN] Get() unable to parse current ID: %s", s.D.Id())
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "identity_domains")

	response, err := s.Client.GetPasswordPolicy(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.PasswordPolicy
	return nil
}

func (s *IdentityDomainsPasswordPolicyResourceCrud) Update() error {
	request := oci_identity_domains.PutPasswordPolicyRequest{}

	if allowedChars, ok := s.D.GetOkExists("allowed_chars"); ok {
		tmp := allowedChars.(string)
		request.AllowedChars = &tmp
	}

	if attributeSets, ok := s.D.GetOkExists("attribute_sets"); ok {
		interfaces := attributeSets.([]interface{})
		tmp := make([]oci_identity_domains.AttributeSetsEnum, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = oci_identity_domains.AttributeSetsEnum(interfaces[i].(string))
			}
		}
		if len(tmp) != 0 || s.D.HasChange("attribute_sets") {
			request.AttributeSets = tmp
		}
	}

	if attributes, ok := s.D.GetOkExists("attributes"); ok {
		tmp := attributes.(string)
		request.Attributes = &tmp
	}

	if authorization, ok := s.D.GetOkExists("authorization"); ok {
		tmp := authorization.(string)
		request.Authorization = &tmp
	}

	if description, ok := s.D.GetOkExists("description"); ok {
		tmp := description.(string)
		request.Description = &tmp
	}

	if dictionaryDelimiter, ok := s.D.GetOkExists("dictionary_delimiter"); ok {
		tmp := dictionaryDelimiter.(string)
		request.DictionaryDelimiter = &tmp
	}

	if dictionaryLocation, ok := s.D.GetOkExists("dictionary_location"); ok {
		tmp := dictionaryLocation.(string)
		request.DictionaryLocation = &tmp
	}

	if dictionaryWordDisallowed, ok := s.D.GetOkExists("dictionary_word_disallowed"); ok {
		tmp := dictionaryWordDisallowed.(bool)
		request.DictionaryWordDisallowed = &tmp
	}

	if disallowedChars, ok := s.D.GetOkExists("disallowed_chars"); ok {
		tmp := disallowedChars.(string)
		request.DisallowedChars = &tmp
	}

	if disallowedSubstrings, ok := s.D.GetOkExists("disallowed_substrings"); ok {
		interfaces := disallowedSubstrings.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("disallowed_substrings") {
			request.DisallowedSubstrings = tmp
		}
	}

	if disallowedUserAttributeValues, ok := s.D.GetOkExists("disallowed_user_attribute_values"); ok {
		interfaces := disallowedUserAttributeValues.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("disallowed_user_attribute_values") {
			request.DisallowedUserAttributeValues = tmp
		}
	}

	if distinctCharacters, ok := s.D.GetOkExists("distinct_characters"); ok {
		tmp := distinctCharacters.(int)
		request.DistinctCharacters = &tmp
	}

	if externalId, ok := s.D.GetOkExists("external_id"); ok {
		tmp := externalId.(string)
		request.ExternalId = &tmp
	}

	if firstNameDisallowed, ok := s.D.GetOkExists("first_name_disallowed"); ok {
		tmp := firstNameDisallowed.(bool)
		request.FirstNameDisallowed = &tmp
	}

	if forcePasswordReset, ok := s.D.GetOkExists("force_password_reset"); ok {
		tmp := forcePasswordReset.(bool)
		request.ForcePasswordReset = &tmp
	}

	if groups, ok := s.D.GetOkExists("groups"); ok {
		interfaces := groups.([]interface{})
		tmp := make([]oci_identity_domains.PasswordPolicyGroups, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "groups", stateDataIndex)
			converted, err := s.mapToPasswordPolicyGroups(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("groups") {
			request.Groups = tmp
		}
	}

	tmp := s.D.Id()
	request.Id = &tmp

	if lastNameDisallowed, ok := s.D.GetOkExists("last_name_disallowed"); ok {
		tmp := lastNameDisallowed.(bool)
		request.LastNameDisallowed = &tmp
	}

	if lockoutDuration, ok := s.D.GetOkExists("lockout_duration"); ok {
		tmp := lockoutDuration.(int)
		request.LockoutDuration = &tmp
	}

	if maxIncorrectAttempts, ok := s.D.GetOkExists("max_incorrect_attempts"); ok {
		tmp := maxIncorrectAttempts.(int)
		request.MaxIncorrectAttempts = &tmp
	}

	if maxLength, ok := s.D.GetOkExists("max_length"); ok {
		tmp := maxLength.(int)
		request.MaxLength = &tmp
	}

	if maxRepeatedChars, ok := s.D.GetOkExists("max_repeated_chars"); ok {
		tmp := maxRepeatedChars.(int)
		request.MaxRepeatedChars = &tmp
	}

	if maxSpecialChars, ok := s.D.GetOkExists("max_special_chars"); ok {
		tmp := maxSpecialChars.(int)
		request.MaxSpecialChars = &tmp
	}

	if minAlphaNumerals, ok := s.D.GetOkExists("min_alpha_numerals"); ok {
		tmp := minAlphaNumerals.(int)
		request.MinAlphaNumerals = &tmp
	}

	if minAlphas, ok := s.D.GetOkExists("min_alphas"); ok {
		tmp := minAlphas.(int)
		request.MinAlphas = &tmp
	}

	if minLength, ok := s.D.GetOkExists("min_length"); ok {
		tmp := minLength.(int)
		request.MinLength = &tmp
	}

	if minLowerCase, ok := s.D.GetOkExists("min_lower_case"); ok {
		tmp := minLowerCase.(int)
		request.MinLowerCase = &tmp
	}

	if minNumerals, ok := s.D.GetOkExists("min_numerals"); ok {
		tmp := minNumerals.(int)
		request.MinNumerals = &tmp
	}

	if minPasswordAge, ok := s.D.GetOkExists("min_password_age"); ok {
		tmp := minPasswordAge.(int)
		request.MinPasswordAge = &tmp
	}

	if minSpecialChars, ok := s.D.GetOkExists("min_special_chars"); ok {
		tmp := minSpecialChars.(int)
		request.MinSpecialChars = &tmp
	}

	if minUniqueChars, ok := s.D.GetOkExists("min_unique_chars"); ok {
		tmp := minUniqueChars.(int)
		request.MinUniqueChars = &tmp
	}

	if minUpperCase, ok := s.D.GetOkExists("min_upper_case"); ok {
		tmp := minUpperCase.(int)
		request.MinUpperCase = &tmp
	}

	if name, ok := s.D.GetOkExists("name"); ok {
		tmp := name.(string)
		request.Name = &tmp
	}

	if numPasswordsInHistory, ok := s.D.GetOkExists("num_passwords_in_history"); ok {
		tmp := numPasswordsInHistory.(int)
		request.NumPasswordsInHistory = &tmp
	}

	if ocid, ok := s.D.GetOkExists("ocid"); ok {
		tmp := ocid.(string)
		request.Ocid = &tmp
	}

	if passwordExpireWarning, ok := s.D.GetOkExists("password_expire_warning"); ok {
		tmp := passwordExpireWarning.(int)
		request.PasswordExpireWarning = &tmp
	}

	if passwordExpiresAfter, ok := s.D.GetOkExists("password_expires_after"); ok {
		tmp := passwordExpiresAfter.(int)
		request.PasswordExpiresAfter = &tmp
	}

	tmp = s.D.Id()
	request.PasswordPolicyId = &tmp

	if passwordStrength, ok := s.D.GetOkExists("password_strength"); ok {
		request.PasswordStrength = oci_identity_domains.PasswordPolicyPasswordStrengthEnum(passwordStrength.(string))
	}

	if priority, ok := s.D.GetOkExists("priority"); ok {
		tmp := priority.(int)
		request.Priority = &tmp
	}

	if requiredChars, ok := s.D.GetOkExists("required_chars"); ok {
		tmp := requiredChars.(string)
		request.RequiredChars = &tmp
	}

	if resourceTypeSchemaVersion, ok := s.D.GetOkExists("resource_type_schema_version"); ok {
		tmp := resourceTypeSchemaVersion.(string)
		request.ResourceTypeSchemaVersion = &tmp
	}

	if schemas, ok := s.D.GetOkExists("schemas"); ok {
		interfaces := schemas.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("schemas") {
			request.Schemas = tmp
		}
	}

	if startsWithAlphabet, ok := s.D.GetOkExists("starts_with_alphabet"); ok {
		tmp := startsWithAlphabet.(bool)
		request.StartsWithAlphabet = &tmp
	}

	if tags, ok := s.D.GetOkExists("tags"); ok {
		interfaces := tags.([]interface{})
		tmp := make([]oci_identity_domains.Tags, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "tags", stateDataIndex)
			converted, err := s.mapTotags(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("tags") {
			request.Tags = tmp
		}
	}

	if userNameDisallowed, ok := s.D.GetOkExists("user_name_disallowed"); ok {
		tmp := userNameDisallowed.(bool)
		request.UserNameDisallowed = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "identity_domains")

	response, err := s.Client.PutPasswordPolicy(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.PasswordPolicy
	return nil
}

func (s *IdentityDomainsPasswordPolicyResourceCrud) Delete() error {
	request := oci_identity_domains.DeletePasswordPolicyRequest{}

	if authorization, ok := s.D.GetOkExists("authorization"); ok {
		tmp := authorization.(string)
		request.Authorization = &tmp
	}

	if forceDelete, ok := s.D.GetOkExists("force_delete"); ok {
		tmp := forceDelete.(bool)
		request.ForceDelete = &tmp
	}

	tmp := s.D.Id()
	request.PasswordPolicyId = &tmp

	if resourceTypeSchemaVersion, ok := s.D.GetOkExists("resource_type_schema_version"); ok {
		tmp := resourceTypeSchemaVersion.(string)
		request.ResourceTypeSchemaVersion = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "identity_domains")

	_, err := s.Client.DeletePasswordPolicy(context.Background(), request)
	return err
}

func (s *IdentityDomainsPasswordPolicyResourceCrud) SetData() error {

	passwordPolicyId, err := parsePasswordPolicyCompositeId(s.D.Id())
	if err == nil {
		s.D.SetId(passwordPolicyId)
	} else {
		log.Printf("[WARN] SetData() unable to parse current ID: %s", s.D.Id())
	}

	if s.Res.AllowedChars != nil {
		s.D.Set("allowed_chars", *s.Res.AllowedChars)
	}

	if s.Res.CompartmentOcid != nil {
		s.D.Set("compartment_ocid", *s.Res.CompartmentOcid)
	}

	configuredPasswordPolicyRules := []interface{}{}
	for _, item := range s.Res.ConfiguredPasswordPolicyRules {
		configuredPasswordPolicyRules = append(configuredPasswordPolicyRules, PasswordPolicyConfiguredPasswordPolicyRulesToMap(item))
	}
	s.D.Set("configured_password_policy_rules", configuredPasswordPolicyRules)

	if s.Res.DeleteInProgress != nil {
		s.D.Set("delete_in_progress", *s.Res.DeleteInProgress)
	}

	if s.Res.Description != nil {
		s.D.Set("description", *s.Res.Description)
	}

	if s.Res.DictionaryDelimiter != nil {
		s.D.Set("dictionary_delimiter", *s.Res.DictionaryDelimiter)
	}

	if s.Res.DictionaryLocation != nil {
		s.D.Set("dictionary_location", *s.Res.DictionaryLocation)
	}

	if s.Res.DictionaryWordDisallowed != nil {
		s.D.Set("dictionary_word_disallowed", *s.Res.DictionaryWordDisallowed)
	}

	if s.Res.DisallowedChars != nil {
		s.D.Set("disallowed_chars", *s.Res.DisallowedChars)
	}

	s.D.Set("disallowed_substrings", s.Res.DisallowedSubstrings)

	s.D.Set("disallowed_user_attribute_values", s.Res.DisallowedUserAttributeValues)

	if s.Res.DistinctCharacters != nil {
		s.D.Set("distinct_characters", *s.Res.DistinctCharacters)
	}

	if s.Res.DomainOcid != nil {
		s.D.Set("domain_ocid", *s.Res.DomainOcid)
	}

	if s.Res.ExternalId != nil {
		s.D.Set("external_id", *s.Res.ExternalId)
	}

	if s.Res.FirstNameDisallowed != nil {
		s.D.Set("first_name_disallowed", *s.Res.FirstNameDisallowed)
	}

	if s.Res.ForcePasswordReset != nil {
		s.D.Set("force_password_reset", *s.Res.ForcePasswordReset)
	}

	groups := []interface{}{}
	for _, item := range s.Res.Groups {
		groups = append(groups, PasswordPolicyGroupsToMap(item))
	}
	s.D.Set("groups", groups)

	if s.Res.IdcsCreatedBy != nil {
		s.D.Set("idcs_created_by", []interface{}{idcsCreatedByToMap(s.Res.IdcsCreatedBy)})
	} else {
		s.D.Set("idcs_created_by", nil)
	}

	if s.Res.IdcsLastModifiedBy != nil {
		s.D.Set("idcs_last_modified_by", []interface{}{idcsLastModifiedByToMap(s.Res.IdcsLastModifiedBy)})
	} else {
		s.D.Set("idcs_last_modified_by", nil)
	}

	if s.Res.IdcsLastUpgradedInRelease != nil {
		s.D.Set("idcs_last_upgraded_in_release", *s.Res.IdcsLastUpgradedInRelease)
	}

	s.D.Set("idcs_prevented_operations", s.Res.IdcsPreventedOperations)

	if s.Res.LastNameDisallowed != nil {
		s.D.Set("last_name_disallowed", *s.Res.LastNameDisallowed)
	}

	if s.Res.LockoutDuration != nil {
		s.D.Set("lockout_duration", *s.Res.LockoutDuration)
	}

	if s.Res.MaxIncorrectAttempts != nil {
		s.D.Set("max_incorrect_attempts", *s.Res.MaxIncorrectAttempts)
	}

	if s.Res.MaxLength != nil {
		s.D.Set("max_length", *s.Res.MaxLength)
	}

	if s.Res.MaxRepeatedChars != nil {
		s.D.Set("max_repeated_chars", *s.Res.MaxRepeatedChars)
	}

	if s.Res.MaxSpecialChars != nil {
		s.D.Set("max_special_chars", *s.Res.MaxSpecialChars)
	}

	if s.Res.Meta != nil {
		s.D.Set("meta", []interface{}{metaToMap(s.Res.Meta)})
	} else {
		s.D.Set("meta", nil)
	}

	if s.Res.MinAlphaNumerals != nil {
		s.D.Set("min_alpha_numerals", *s.Res.MinAlphaNumerals)
	}

	if s.Res.MinAlphas != nil {
		s.D.Set("min_alphas", *s.Res.MinAlphas)
	}

	if s.Res.MinLength != nil {
		s.D.Set("min_length", *s.Res.MinLength)
	}

	if s.Res.MinLowerCase != nil {
		s.D.Set("min_lower_case", *s.Res.MinLowerCase)
	}

	if s.Res.MinNumerals != nil {
		s.D.Set("min_numerals", *s.Res.MinNumerals)
	}

	if s.Res.MinPasswordAge != nil {
		s.D.Set("min_password_age", *s.Res.MinPasswordAge)
	}

	if s.Res.MinSpecialChars != nil {
		s.D.Set("min_special_chars", *s.Res.MinSpecialChars)
	}

	if s.Res.MinUniqueChars != nil {
		s.D.Set("min_unique_chars", *s.Res.MinUniqueChars)
	}

	if s.Res.MinUpperCase != nil {
		s.D.Set("min_upper_case", *s.Res.MinUpperCase)
	}

	if s.Res.Name != nil {
		s.D.Set("name", *s.Res.Name)
	}

	if s.Res.NumPasswordsInHistory != nil {
		s.D.Set("num_passwords_in_history", *s.Res.NumPasswordsInHistory)
	}

	if s.Res.Ocid != nil {
		s.D.Set("ocid", *s.Res.Ocid)
	}

	if s.Res.PasswordExpireWarning != nil {
		s.D.Set("password_expire_warning", *s.Res.PasswordExpireWarning)
	}

	if s.Res.PasswordExpiresAfter != nil {
		s.D.Set("password_expires_after", *s.Res.PasswordExpiresAfter)
	}

	s.D.Set("password_strength", s.Res.PasswordStrength)

	if s.Res.Priority != nil {
		s.D.Set("priority", *s.Res.Priority)
	}

	if s.Res.RequiredChars != nil {
		s.D.Set("required_chars", *s.Res.RequiredChars)
	}

	s.D.Set("schemas", s.Res.Schemas)

	if s.Res.StartsWithAlphabet != nil {
		s.D.Set("starts_with_alphabet", *s.Res.StartsWithAlphabet)
	}

	tags := []interface{}{}
	for _, item := range s.Res.Tags {
		tags = append(tags, tagsToMap(item))
	}
	s.D.Set("tags", tags)

	if s.Res.TenancyOcid != nil {
		s.D.Set("tenancy_ocid", *s.Res.TenancyOcid)
	}

	if s.Res.UserNameDisallowed != nil {
		s.D.Set("user_name_disallowed", *s.Res.UserNameDisallowed)
	}

	return nil
}

//func GetPasswordPolicyCompositeId(passwordPolicyId string) string {
//	id = url.PathEscape(id)
//	idcsEndpoint = url.PathEscape(idcsEndpoint)
//	passwordPolicyId = url.PathEscape(passwordPolicyId)
//	compositeId := "idcsEndpoint/" + idcsEndpoint + "/passwordPolicies/" + passwordPolicyId
//	return compositeId
//}

func parsePasswordPolicyCompositeId(compositeId string) (passwordPolicyId string, err error) {
	parts := strings.Split(compositeId, "/")
	match, _ := regexp.MatchString("idcsEndpoint/.*/passwordPolicies/.*", compositeId)
	if !match || len(parts) != 4 {
		err = fmt.Errorf("illegal compositeId %s encountered", compositeId)
		return
	}
	//idcsEndpoint, _ = url.PathUnescape(parts[1])
	passwordPolicyId, _ = url.PathUnescape(parts[3])

	return
}

func PasswordPolicyToMap(obj oci_identity_domains.PasswordPolicy) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.AllowedChars != nil {
		result["allowed_chars"] = string(*obj.AllowedChars)
	}

	if obj.CompartmentOcid != nil {
		result["compartment_ocid"] = string(*obj.CompartmentOcid)
	}

	configuredPasswordPolicyRules := []interface{}{}
	for _, item := range obj.ConfiguredPasswordPolicyRules {
		configuredPasswordPolicyRules = append(configuredPasswordPolicyRules, PasswordPolicyConfiguredPasswordPolicyRulesToMap(item))
	}
	result["configured_password_policy_rules"] = configuredPasswordPolicyRules

	if obj.DeleteInProgress != nil {
		result["delete_in_progress"] = bool(*obj.DeleteInProgress)
	}

	if obj.Description != nil {
		result["description"] = string(*obj.Description)
	}

	if obj.DictionaryDelimiter != nil {
		result["dictionary_delimiter"] = string(*obj.DictionaryDelimiter)
	}

	if obj.DictionaryLocation != nil {
		result["dictionary_location"] = string(*obj.DictionaryLocation)
	}

	if obj.DictionaryWordDisallowed != nil {
		result["dictionary_word_disallowed"] = bool(*obj.DictionaryWordDisallowed)
	}

	if obj.DisallowedChars != nil {
		result["disallowed_chars"] = string(*obj.DisallowedChars)
	}

	result["disallowed_substrings"] = obj.DisallowedSubstrings

	result["disallowed_user_attribute_values"] = obj.DisallowedUserAttributeValues

	if obj.DistinctCharacters != nil {
		result["distinct_characters"] = int(*obj.DistinctCharacters)
	}

	if obj.DomainOcid != nil {
		result["domain_ocid"] = string(*obj.DomainOcid)
	}

	if obj.ExternalId != nil {
		result["external_id"] = string(*obj.ExternalId)
	}

	if obj.FirstNameDisallowed != nil {
		result["first_name_disallowed"] = bool(*obj.FirstNameDisallowed)
	}

	if obj.ForcePasswordReset != nil {
		result["force_password_reset"] = bool(*obj.ForcePasswordReset)
	}

	groups := []interface{}{}
	for _, item := range obj.Groups {
		groups = append(groups, PasswordPolicyGroupsToMap(item))
	}
	result["groups"] = groups

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	if obj.IdcsCreatedBy != nil {
		result["idcs_created_by"] = []interface{}{idcsCreatedByToMap(obj.IdcsCreatedBy)}
	}

	if obj.IdcsLastModifiedBy != nil {
		result["idcs_last_modified_by"] = []interface{}{idcsLastModifiedByToMap(obj.IdcsLastModifiedBy)}
	}

	if obj.IdcsLastUpgradedInRelease != nil {
		result["idcs_last_upgraded_in_release"] = string(*obj.IdcsLastUpgradedInRelease)
	}

	result["idcs_prevented_operations"] = obj.IdcsPreventedOperations

	if obj.LastNameDisallowed != nil {
		result["last_name_disallowed"] = bool(*obj.LastNameDisallowed)
	}

	if obj.LockoutDuration != nil {
		result["lockout_duration"] = int(*obj.LockoutDuration)
	}

	if obj.MaxIncorrectAttempts != nil {
		result["max_incorrect_attempts"] = int(*obj.MaxIncorrectAttempts)
	}

	if obj.MaxLength != nil {
		result["max_length"] = int(*obj.MaxLength)
	}

	if obj.MaxRepeatedChars != nil {
		result["max_repeated_chars"] = int(*obj.MaxRepeatedChars)
	}

	if obj.MaxSpecialChars != nil {
		result["max_special_chars"] = int(*obj.MaxSpecialChars)
	}

	if obj.Meta != nil {
		result["meta"] = []interface{}{metaToMap(obj.Meta)}
	}

	if obj.MinAlphaNumerals != nil {
		result["min_alpha_numerals"] = int(*obj.MinAlphaNumerals)
	}

	if obj.MinAlphas != nil {
		result["min_alphas"] = int(*obj.MinAlphas)
	}

	if obj.MinLength != nil {
		result["min_length"] = int(*obj.MinLength)
	}

	if obj.MinLowerCase != nil {
		result["min_lower_case"] = int(*obj.MinLowerCase)
	}

	if obj.MinNumerals != nil {
		result["min_numerals"] = int(*obj.MinNumerals)
	}

	if obj.MinPasswordAge != nil {
		result["min_password_age"] = int(*obj.MinPasswordAge)
	}

	if obj.MinSpecialChars != nil {
		result["min_special_chars"] = int(*obj.MinSpecialChars)
	}

	if obj.MinUniqueChars != nil {
		result["min_unique_chars"] = int(*obj.MinUniqueChars)
	}

	if obj.MinUpperCase != nil {
		result["min_upper_case"] = int(*obj.MinUpperCase)
	}

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	if obj.NumPasswordsInHistory != nil {
		result["num_passwords_in_history"] = int(*obj.NumPasswordsInHistory)
	}

	if obj.Ocid != nil {
		result["ocid"] = string(*obj.Ocid)
	}

	if obj.PasswordExpireWarning != nil {
		result["password_expire_warning"] = int(*obj.PasswordExpireWarning)
	}

	if obj.PasswordExpiresAfter != nil {
		result["password_expires_after"] = int(*obj.PasswordExpiresAfter)
	}

	result["password_strength"] = string(obj.PasswordStrength)

	if obj.Priority != nil {
		result["priority"] = int(*obj.Priority)
	}

	if obj.RequiredChars != nil {
		result["required_chars"] = string(*obj.RequiredChars)
	}

	result["schemas"] = obj.Schemas

	if obj.StartsWithAlphabet != nil {
		result["starts_with_alphabet"] = bool(*obj.StartsWithAlphabet)
	}

	tags := []interface{}{}
	for _, item := range obj.Tags {
		tags = append(tags, tagsToMap(item))
	}
	result["tags"] = tags

	if obj.TenancyOcid != nil {
		result["tenancy_ocid"] = string(*obj.TenancyOcid)
	}

	if obj.UserNameDisallowed != nil {
		result["user_name_disallowed"] = bool(*obj.UserNameDisallowed)
	}

	return result
}

func PasswordPolicyConfiguredPasswordPolicyRulesToMap(obj oci_identity_domains.PasswordPolicyConfiguredPasswordPolicyRules) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Key != nil {
		result["key"] = string(*obj.Key)
	}

	if obj.Value != nil {
		result["value"] = string(*obj.Value)
	}

	return result
}

func (s *IdentityDomainsPasswordPolicyResourceCrud) mapToPasswordPolicyGroups(fieldKeyFormat string) (oci_identity_domains.PasswordPolicyGroups, error) {
	result := oci_identity_domains.PasswordPolicyGroups{}

	if display, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "display")); ok {
		tmp := display.(string)
		result.Display = &tmp
	}

	if ref, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "ref")); ok {
		tmp := ref.(string)
		result.Ref = &tmp
	}

	if value, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "value")); ok {
		tmp := value.(string)
		result.Value = &tmp
	}

	return result, nil
}

func PasswordPolicyGroupsToMap(obj oci_identity_domains.PasswordPolicyGroups) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Display != nil {
		result["display"] = string(*obj.Display)
	}

	if obj.Ref != nil {
		result["ref"] = string(*obj.Ref)
	}

	if obj.Value != nil {
		result["value"] = string(*obj.Value)
	}

	return result
}

func (s *IdentityDomainsPasswordPolicyResourceCrud) mapTotags(fieldKeyFormat string) (oci_identity_domains.Tags, error) {
	result := oci_identity_domains.Tags{}

	if key, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "key")); ok {
		tmp := key.(string)
		result.Key = &tmp
	}

	if value, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "value")); ok {
		tmp := value.(string)
		result.Value = &tmp
	}

	return result, nil
}
