// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package identity_domains

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_identity_domains "github.com/oracle/oci-go-sdk/v65/identitydomains"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func IdentityDomainsPasswordPolicyDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["attribute_sets"] = &schema.Schema{
		Type:     schema.TypeList,
		Optional: true,
		Elem: &schema.Schema{
			Type: schema.TypeString,
		},
	}
	fieldMap["attributes"] = &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
	}
	fieldMap["authorization"] = &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
	}
	fieldMap["idcs_endpoint"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	fieldMap["password_policy_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	fieldMap["resource_type_schema_version"] = &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(IdentityDomainsPasswordPolicyResource(), fieldMap, readSingularIdentityDomainsPasswordPolicy)
}

func readSingularIdentityDomainsPasswordPolicy(d *schema.ResourceData, m interface{}) error {
	sync := &IdentityDomainsPasswordPolicyDataSourceCrud{}
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

	return tfresource.ReadResource(sync)
}

type IdentityDomainsPasswordPolicyDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_identity_domains.IdentityDomainsClient
	Res    *oci_identity_domains.GetPasswordPolicyResponse
}

func (s *IdentityDomainsPasswordPolicyDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *IdentityDomainsPasswordPolicyDataSourceCrud) Get() error {
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

	if passwordPolicyId, ok := s.D.GetOkExists("password_policy_id"); ok {
		tmp := passwordPolicyId.(string)
		request.PasswordPolicyId = &tmp
	}

	if resourceTypeSchemaVersion, ok := s.D.GetOkExists("resource_type_schema_version"); ok {
		tmp := resourceTypeSchemaVersion.(string)
		request.ResourceTypeSchemaVersion = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "identity_domains")

	response, err := s.Client.GetPasswordPolicy(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *IdentityDomainsPasswordPolicyDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

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
