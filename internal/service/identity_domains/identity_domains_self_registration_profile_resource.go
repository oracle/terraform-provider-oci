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

func IdentityDomainsSelfRegistrationProfileResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createIdentityDomainsSelfRegistrationProfile,
		Read:     readIdentityDomainsSelfRegistrationProfile,
		Update:   updateIdentityDomainsSelfRegistrationProfile,
		Delete:   deleteIdentityDomainsSelfRegistrationProfile,
		Schema: map[string]*schema.Schema{
			// Required
			"activation_email_required": {
				Type:     schema.TypeBool,
				Required: true,
			},
			"consent_text_present": {
				Type:     schema.TypeBool,
				Required: true,
			},
			"display_name": {
				Type:     schema.TypeList,
				Required: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"locale": {
							Type:     schema.TypeString,
							Required: true,
						},
						"value": {
							Type:     schema.TypeString,
							Required: true,
						},

						// Optional
						"default": {
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
						},

						// Computed
					},
				},
			},
			"email_template": {
				Type:     schema.TypeList,
				Required: true,
				MaxItems: 1,
				MinItems: 1,
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
			"idcs_endpoint": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"number_of_days_redirect_url_is_valid": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"redirect_url": {
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
			"show_on_login_page": {
				Type:     schema.TypeBool,
				Required: true,
			},

			// Optional
			"active": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"after_submit_text": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"locale": {
							Type:     schema.TypeString,
							Required: true,
						},
						"value": {
							Type:     schema.TypeString,
							Required: true,
						},

						// Optional
						"default": {
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
						},

						// Computed
					},
				},
			},
			"allowed_email_domains": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
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
			"consent_text": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"locale": {
							Type:     schema.TypeString,
							Required: true,
						},
						"value": {
							Type:     schema.TypeString,
							Required: true,
						},

						// Optional
						"default": {
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
						},

						// Computed
					},
				},
			},
			"default_groups": {
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
			"disallowed_email_domains": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"external_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"footer_logo": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"footer_text": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"locale": {
							Type:     schema.TypeString,
							Required: true,
						},
						"value": {
							Type:     schema.TypeString,
							Required: true,
						},

						// Optional
						"default": {
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
						},

						// Computed
					},
				},
			},
			"header_logo": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"header_text": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"locale": {
							Type:     schema.TypeString,
							Required: true,
						},
						"value": {
							Type:     schema.TypeString,
							Required: true,
						},

						// Optional
						"default": {
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
						},

						// Computed
					},
				},
			},
			"ocid": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"resource_type_schema_version": {
				Type:     schema.TypeString,
				Optional: true,
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
			"user_attributes": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"seq_number": {
							Type:     schema.TypeInt,
							Required: true,
						},
						"value": {
							Type:     schema.TypeString,
							Required: true,
						},

						// Optional
						"fully_qualified_attribute_name": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},

						// Computed
						"deletable": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"metadata": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},

			// Computed
			"compartment_ocid": {
				Type:     schema.TypeString,
				Computed: true,
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

func createIdentityDomainsSelfRegistrationProfile(d *schema.ResourceData, m interface{}) error {
	sync := &IdentityDomainsSelfRegistrationProfileResourceCrud{}
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

func readIdentityDomainsSelfRegistrationProfile(d *schema.ResourceData, m interface{}) error {
	sync := &IdentityDomainsSelfRegistrationProfileResourceCrud{}
	sync.D = d
	idcsEndpoint, err := getIdcsEndpointForRead(d, "selfRegistrationProfiles")
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

func updateIdentityDomainsSelfRegistrationProfile(d *schema.ResourceData, m interface{}) error {
	sync := &IdentityDomainsSelfRegistrationProfileResourceCrud{}
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

func deleteIdentityDomainsSelfRegistrationProfile(d *schema.ResourceData, m interface{}) error {
	sync := &IdentityDomainsSelfRegistrationProfileResourceCrud{}
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
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type IdentityDomainsSelfRegistrationProfileResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_identity_domains.IdentityDomainsClient
	Res                    *oci_identity_domains.SelfRegistrationProfile
	DisableNotFoundRetries bool
}

func (s *IdentityDomainsSelfRegistrationProfileResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *IdentityDomainsSelfRegistrationProfileResourceCrud) Create() error {
	request := oci_identity_domains.CreateSelfRegistrationProfileRequest{}

	if activationEmailRequired, ok := s.D.GetOkExists("activation_email_required"); ok {
		tmp := activationEmailRequired.(bool)
		request.ActivationEmailRequired = &tmp
	}

	if active, ok := s.D.GetOkExists("active"); ok {
		tmp := active.(bool)
		request.Active = &tmp
	}

	if afterSubmitText, ok := s.D.GetOkExists("after_submit_text"); ok {
		interfaces := afterSubmitText.([]interface{})
		tmp := make([]oci_identity_domains.SelfRegistrationProfileAfterSubmitText, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "after_submit_text", stateDataIndex)
			converted, err := s.mapToSelfRegistrationProfileAfterSubmitText(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("after_submit_text") {
			request.AfterSubmitText = tmp
		}
	}

	if allowedEmailDomains, ok := s.D.GetOkExists("allowed_email_domains"); ok {
		interfaces := allowedEmailDomains.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("allowed_email_domains") {
			request.AllowedEmailDomains = tmp
		}
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

	if consentText, ok := s.D.GetOkExists("consent_text"); ok {
		interfaces := consentText.([]interface{})
		tmp := make([]oci_identity_domains.SelfRegistrationProfileConsentText, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "consent_text", stateDataIndex)
			converted, err := s.mapToSelfRegistrationProfileConsentText(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("consent_text") {
			request.ConsentText = tmp
		}
	}

	if consentTextPresent, ok := s.D.GetOkExists("consent_text_present"); ok {
		tmp := consentTextPresent.(bool)
		request.ConsentTextPresent = &tmp
	}

	if defaultGroups, ok := s.D.GetOkExists("default_groups"); ok {
		interfaces := defaultGroups.([]interface{})
		tmp := make([]oci_identity_domains.SelfRegistrationProfileDefaultGroups, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "default_groups", stateDataIndex)
			converted, err := s.mapToSelfRegistrationProfileDefaultGroups(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("default_groups") {
			request.DefaultGroups = tmp
		}
	}

	if disallowedEmailDomains, ok := s.D.GetOkExists("disallowed_email_domains"); ok {
		interfaces := disallowedEmailDomains.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("disallowed_email_domains") {
			request.DisallowedEmailDomains = tmp
		}
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		interfaces := displayName.([]interface{})
		tmp := make([]oci_identity_domains.SelfRegistrationProfileDisplayName, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "display_name", stateDataIndex)
			converted, err := s.mapToSelfRegistrationProfileDisplayName(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("display_name") {
			request.DisplayName = tmp
		}
	}

	if emailTemplate, ok := s.D.GetOkExists("email_template"); ok {
		if tmpList := emailTemplate.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "email_template", 0)
			tmp, err := s.mapToSelfRegistrationProfileEmailTemplate(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.EmailTemplate = &tmp
		}
	}

	if externalId, ok := s.D.GetOkExists("external_id"); ok {
		tmp := externalId.(string)
		request.ExternalId = &tmp
	}

	if footerLogo, ok := s.D.GetOkExists("footer_logo"); ok {
		tmp := footerLogo.(string)
		request.FooterLogo = &tmp
	}

	if footerText, ok := s.D.GetOkExists("footer_text"); ok {
		interfaces := footerText.([]interface{})
		tmp := make([]oci_identity_domains.SelfRegistrationProfileFooterText, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "footer_text", stateDataIndex)
			converted, err := s.mapToSelfRegistrationProfileFooterText(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("footer_text") {
			request.FooterText = tmp
		}
	}

	if headerLogo, ok := s.D.GetOkExists("header_logo"); ok {
		tmp := headerLogo.(string)
		request.HeaderLogo = &tmp
	}

	if headerText, ok := s.D.GetOkExists("header_text"); ok {
		interfaces := headerText.([]interface{})
		tmp := make([]oci_identity_domains.SelfRegistrationProfileHeaderText, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "header_text", stateDataIndex)
			converted, err := s.mapToSelfRegistrationProfileHeaderText(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("header_text") {
			request.HeaderText = tmp
		}
	}

	if id, ok := s.D.GetOkExists("id"); ok {
		tmp := id.(string)
		request.Id = &tmp
	}

	if name, ok := s.D.GetOkExists("name"); ok {
		tmp := name.(string)
		request.Name = &tmp
	}

	if numberOfDaysRedirectUrlIsValid, ok := s.D.GetOkExists("number_of_days_redirect_url_is_valid"); ok {
		tmp := numberOfDaysRedirectUrlIsValid.(int)
		request.NumberOfDaysRedirectUrlIsValid = &tmp
	}

	if ocid, ok := s.D.GetOkExists("ocid"); ok {
		tmp := ocid.(string)
		request.Ocid = &tmp
	}

	if redirectUrl, ok := s.D.GetOkExists("redirect_url"); ok {
		tmp := redirectUrl.(string)
		request.RedirectUrl = &tmp
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

	if showOnLoginPage, ok := s.D.GetOkExists("show_on_login_page"); ok {
		tmp := showOnLoginPage.(bool)
		request.ShowOnLoginPage = &tmp
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

	if userAttributes, ok := s.D.GetOkExists("user_attributes"); ok {
		interfaces := userAttributes.([]interface{})
		tmp := make([]oci_identity_domains.SelfRegistrationProfileUserAttributes, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "user_attributes", stateDataIndex)
			converted, err := s.mapToSelfRegistrationProfileUserAttributes(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("user_attributes") {
			request.UserAttributes = tmp
		}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "identity_domains")

	response, err := s.Client.CreateSelfRegistrationProfile(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.SelfRegistrationProfile
	return nil
}

func (s *IdentityDomainsSelfRegistrationProfileResourceCrud) Get() error {
	request := oci_identity_domains.GetSelfRegistrationProfileRequest{}

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

	if resourceTypeSchemaVersion, ok := s.D.GetOkExists("resource_type_schema_version"); ok {
		tmp := resourceTypeSchemaVersion.(string)
		request.ResourceTypeSchemaVersion = &tmp
	}

	tmp := s.D.Id()
	request.SelfRegistrationProfileId = &tmp

	selfRegistrationProfileId, err := parseSelfRegistrationProfileCompositeId(s.D.Id())
	if err == nil {
		request.SelfRegistrationProfileId = &selfRegistrationProfileId
	} else {
		log.Printf("[WARN] Get() unable to parse current ID: %s", s.D.Id())
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "identity_domains")

	response, err := s.Client.GetSelfRegistrationProfile(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.SelfRegistrationProfile
	return nil
}

func (s *IdentityDomainsSelfRegistrationProfileResourceCrud) Update() error {
	request := oci_identity_domains.PutSelfRegistrationProfileRequest{}

	if activationEmailRequired, ok := s.D.GetOkExists("activation_email_required"); ok {
		tmp := activationEmailRequired.(bool)
		request.ActivationEmailRequired = &tmp
	}

	if active, ok := s.D.GetOkExists("active"); ok {
		tmp := active.(bool)
		request.Active = &tmp
	}

	if afterSubmitText, ok := s.D.GetOkExists("after_submit_text"); ok {
		interfaces := afterSubmitText.([]interface{})
		tmp := make([]oci_identity_domains.SelfRegistrationProfileAfterSubmitText, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "after_submit_text", stateDataIndex)
			converted, err := s.mapToSelfRegistrationProfileAfterSubmitText(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("after_submit_text") {
			request.AfterSubmitText = tmp
		}
	}

	if allowedEmailDomains, ok := s.D.GetOkExists("allowed_email_domains"); ok {
		interfaces := allowedEmailDomains.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("allowed_email_domains") {
			request.AllowedEmailDomains = tmp
		}
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

	if consentText, ok := s.D.GetOkExists("consent_text"); ok {
		interfaces := consentText.([]interface{})
		tmp := make([]oci_identity_domains.SelfRegistrationProfileConsentText, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "consent_text", stateDataIndex)
			converted, err := s.mapToSelfRegistrationProfileConsentText(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("consent_text") {
			request.ConsentText = tmp
		}
	}

	if consentTextPresent, ok := s.D.GetOkExists("consent_text_present"); ok {
		tmp := consentTextPresent.(bool)
		request.ConsentTextPresent = &tmp
	}

	if defaultGroups, ok := s.D.GetOkExists("default_groups"); ok {
		interfaces := defaultGroups.([]interface{})
		tmp := make([]oci_identity_domains.SelfRegistrationProfileDefaultGroups, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "default_groups", stateDataIndex)
			converted, err := s.mapToSelfRegistrationProfileDefaultGroups(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("default_groups") {
			request.DefaultGroups = tmp
		}
	}

	if disallowedEmailDomains, ok := s.D.GetOkExists("disallowed_email_domains"); ok {
		interfaces := disallowedEmailDomains.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("disallowed_email_domains") {
			request.DisallowedEmailDomains = tmp
		}
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		interfaces := displayName.([]interface{})
		tmp := make([]oci_identity_domains.SelfRegistrationProfileDisplayName, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "display_name", stateDataIndex)
			converted, err := s.mapToSelfRegistrationProfileDisplayName(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("display_name") {
			request.DisplayName = tmp
		}
	}

	if emailTemplate, ok := s.D.GetOkExists("email_template"); ok {
		if tmpList := emailTemplate.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "email_template", 0)
			tmp, err := s.mapToSelfRegistrationProfileEmailTemplate(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.EmailTemplate = &tmp
		}
	}

	if externalId, ok := s.D.GetOkExists("external_id"); ok {
		tmp := externalId.(string)
		request.ExternalId = &tmp
	}

	if footerLogo, ok := s.D.GetOkExists("footer_logo"); ok {
		tmp := footerLogo.(string)
		request.FooterLogo = &tmp
	}

	if footerText, ok := s.D.GetOkExists("footer_text"); ok {
		interfaces := footerText.([]interface{})
		tmp := make([]oci_identity_domains.SelfRegistrationProfileFooterText, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "footer_text", stateDataIndex)
			converted, err := s.mapToSelfRegistrationProfileFooterText(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("footer_text") {
			request.FooterText = tmp
		}
	}

	if headerLogo, ok := s.D.GetOkExists("header_logo"); ok {
		tmp := headerLogo.(string)
		request.HeaderLogo = &tmp
	}

	if headerText, ok := s.D.GetOkExists("header_text"); ok {
		interfaces := headerText.([]interface{})
		tmp := make([]oci_identity_domains.SelfRegistrationProfileHeaderText, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "header_text", stateDataIndex)
			converted, err := s.mapToSelfRegistrationProfileHeaderText(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("header_text") {
			request.HeaderText = tmp
		}
	}

	tmp := s.D.Id()
	request.Id = &tmp

	if name, ok := s.D.GetOkExists("name"); ok {
		tmp := name.(string)
		request.Name = &tmp
	}

	if numberOfDaysRedirectUrlIsValid, ok := s.D.GetOkExists("number_of_days_redirect_url_is_valid"); ok {
		tmp := numberOfDaysRedirectUrlIsValid.(int)
		request.NumberOfDaysRedirectUrlIsValid = &tmp
	}

	if ocid, ok := s.D.GetOkExists("ocid"); ok {
		tmp := ocid.(string)
		request.Ocid = &tmp
	}

	if redirectUrl, ok := s.D.GetOkExists("redirect_url"); ok {
		tmp := redirectUrl.(string)
		request.RedirectUrl = &tmp
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

	tmp = s.D.Id()
	request.SelfRegistrationProfileId = &tmp

	if showOnLoginPage, ok := s.D.GetOkExists("show_on_login_page"); ok {
		tmp := showOnLoginPage.(bool)
		request.ShowOnLoginPage = &tmp
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

	if userAttributes, ok := s.D.GetOkExists("user_attributes"); ok {
		interfaces := userAttributes.([]interface{})
		tmp := make([]oci_identity_domains.SelfRegistrationProfileUserAttributes, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "user_attributes", stateDataIndex)
			converted, err := s.mapToSelfRegistrationProfileUserAttributes(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("user_attributes") {
			request.UserAttributes = tmp
		}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "identity_domains")

	response, err := s.Client.PutSelfRegistrationProfile(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.SelfRegistrationProfile
	return nil
}

func (s *IdentityDomainsSelfRegistrationProfileResourceCrud) Delete() error {
	request := oci_identity_domains.DeleteSelfRegistrationProfileRequest{}

	if authorization, ok := s.D.GetOkExists("authorization"); ok {
		tmp := authorization.(string)
		request.Authorization = &tmp
	}

	if forceDelete, ok := s.D.GetOkExists("force_delete"); ok {
		tmp := forceDelete.(bool)
		request.ForceDelete = &tmp
	}

	if resourceTypeSchemaVersion, ok := s.D.GetOkExists("resource_type_schema_version"); ok {
		tmp := resourceTypeSchemaVersion.(string)
		request.ResourceTypeSchemaVersion = &tmp
	}

	tmp := s.D.Id()
	request.SelfRegistrationProfileId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "identity_domains")

	_, err := s.Client.DeleteSelfRegistrationProfile(context.Background(), request)
	return err
}

func (s *IdentityDomainsSelfRegistrationProfileResourceCrud) SetData() error {

	selfRegistrationProfileId, err := parseSelfRegistrationProfileCompositeId(s.D.Id())
	if err == nil {
		s.D.SetId(selfRegistrationProfileId)
	} else {
		log.Printf("[WARN] SetData() unable to parse current ID: %s", s.D.Id())
	}

	if s.Res.ActivationEmailRequired != nil {
		s.D.Set("activation_email_required", *s.Res.ActivationEmailRequired)
	}

	if s.Res.Active != nil {
		s.D.Set("active", *s.Res.Active)
	}

	afterSubmitText := []interface{}{}
	for _, item := range s.Res.AfterSubmitText {
		afterSubmitText = append(afterSubmitText, SelfRegistrationProfileAfterSubmitTextToMap(item))
	}
	s.D.Set("after_submit_text", afterSubmitText)

	s.D.Set("allowed_email_domains", s.Res.AllowedEmailDomains)

	if s.Res.CompartmentOcid != nil {
		s.D.Set("compartment_ocid", *s.Res.CompartmentOcid)
	}

	consentText := []interface{}{}
	for _, item := range s.Res.ConsentText {
		consentText = append(consentText, SelfRegistrationProfileConsentTextToMap(item))
	}
	s.D.Set("consent_text", consentText)

	if s.Res.ConsentTextPresent != nil {
		s.D.Set("consent_text_present", *s.Res.ConsentTextPresent)
	}

	defaultGroups := []interface{}{}
	for _, item := range s.Res.DefaultGroups {
		defaultGroups = append(defaultGroups, SelfRegistrationProfileDefaultGroupsToMap(item))
	}
	s.D.Set("default_groups", defaultGroups)

	if s.Res.DeleteInProgress != nil {
		s.D.Set("delete_in_progress", *s.Res.DeleteInProgress)
	}

	s.D.Set("disallowed_email_domains", s.Res.DisallowedEmailDomains)

	displayName := []interface{}{}
	for _, item := range s.Res.DisplayName {
		displayName = append(displayName, SelfRegistrationProfileDisplayNameToMap(item))
	}
	s.D.Set("display_name", displayName)

	if s.Res.DomainOcid != nil {
		s.D.Set("domain_ocid", *s.Res.DomainOcid)
	}

	if s.Res.EmailTemplate != nil {
		s.D.Set("email_template", []interface{}{SelfRegistrationProfileEmailTemplateToMap(s.Res.EmailTemplate)})
	} else {
		s.D.Set("email_template", nil)
	}

	if s.Res.ExternalId != nil {
		s.D.Set("external_id", *s.Res.ExternalId)
	}

	if s.Res.FooterLogo != nil {
		s.D.Set("footer_logo", *s.Res.FooterLogo)
	}

	footerText := []interface{}{}
	for _, item := range s.Res.FooterText {
		footerText = append(footerText, SelfRegistrationProfileFooterTextToMap(item))
	}
	s.D.Set("footer_text", footerText)

	if s.Res.HeaderLogo != nil {
		s.D.Set("header_logo", *s.Res.HeaderLogo)
	}

	headerText := []interface{}{}
	for _, item := range s.Res.HeaderText {
		headerText = append(headerText, SelfRegistrationProfileHeaderTextToMap(item))
	}
	s.D.Set("header_text", headerText)

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

	if s.Res.Meta != nil {
		s.D.Set("meta", []interface{}{metaToMap(s.Res.Meta)})
	} else {
		s.D.Set("meta", nil)
	}

	if s.Res.Name != nil {
		s.D.Set("name", *s.Res.Name)
	}

	if s.Res.NumberOfDaysRedirectUrlIsValid != nil {
		s.D.Set("number_of_days_redirect_url_is_valid", *s.Res.NumberOfDaysRedirectUrlIsValid)
	}

	if s.Res.Ocid != nil {
		s.D.Set("ocid", *s.Res.Ocid)
	}

	if s.Res.RedirectUrl != nil {
		s.D.Set("redirect_url", *s.Res.RedirectUrl)
	}

	s.D.Set("schemas", s.Res.Schemas)

	if s.Res.ShowOnLoginPage != nil {
		s.D.Set("show_on_login_page", *s.Res.ShowOnLoginPage)
	}

	tags := []interface{}{}
	for _, item := range s.Res.Tags {
		tags = append(tags, tagsToMap(item))
	}
	s.D.Set("tags", tags)

	if s.Res.TenancyOcid != nil {
		s.D.Set("tenancy_ocid", *s.Res.TenancyOcid)
	}

	userAttributes := []interface{}{}
	for _, item := range s.Res.UserAttributes {
		userAttributes = append(userAttributes, SelfRegistrationProfileUserAttributesToMap(item))
	}
	s.D.Set("user_attributes", userAttributes)

	return nil
}

func parseSelfRegistrationProfileCompositeId(compositeId string) (selfRegistrationProfileId string, err error) {
	parts := strings.Split(compositeId, "/")
	match, _ := regexp.MatchString("idcsEndpoint/.*/selfRegistrationProfiles/.*", compositeId)
	if !match || len(parts) != 4 {
		err = fmt.Errorf("illegal compositeId %s encountered", compositeId)
		return
	}
	selfRegistrationProfileId, _ = url.PathUnescape(parts[3])

	return
}

func SelfRegistrationProfileToMap(obj oci_identity_domains.SelfRegistrationProfile) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.ActivationEmailRequired != nil {
		result["activation_email_required"] = bool(*obj.ActivationEmailRequired)
	}

	if obj.Active != nil {
		result["active"] = bool(*obj.Active)
	}

	afterSubmitText := []interface{}{}
	for _, item := range obj.AfterSubmitText {
		afterSubmitText = append(afterSubmitText, SelfRegistrationProfileAfterSubmitTextToMap(item))
	}
	result["after_submit_text"] = afterSubmitText

	result["allowed_email_domains"] = obj.AllowedEmailDomains

	if obj.CompartmentOcid != nil {
		result["compartment_ocid"] = string(*obj.CompartmentOcid)
	}

	consentText := []interface{}{}
	for _, item := range obj.ConsentText {
		consentText = append(consentText, SelfRegistrationProfileConsentTextToMap(item))
	}
	result["consent_text"] = consentText

	if obj.ConsentTextPresent != nil {
		result["consent_text_present"] = bool(*obj.ConsentTextPresent)
	}

	defaultGroups := []interface{}{}
	for _, item := range obj.DefaultGroups {
		defaultGroups = append(defaultGroups, SelfRegistrationProfileDefaultGroupsToMap(item))
	}
	result["default_groups"] = defaultGroups

	if obj.DeleteInProgress != nil {
		result["delete_in_progress"] = bool(*obj.DeleteInProgress)
	}

	result["disallowed_email_domains"] = obj.DisallowedEmailDomains

	displayName := []interface{}{}
	for _, item := range obj.DisplayName {
		displayName = append(displayName, SelfRegistrationProfileDisplayNameToMap(item))
	}
	result["display_name"] = displayName

	if obj.DomainOcid != nil {
		result["domain_ocid"] = string(*obj.DomainOcid)
	}

	if obj.EmailTemplate != nil {
		result["email_template"] = []interface{}{SelfRegistrationProfileEmailTemplateToMap(obj.EmailTemplate)}
	}

	if obj.ExternalId != nil {
		result["external_id"] = string(*obj.ExternalId)
	}

	if obj.FooterLogo != nil {
		result["footer_logo"] = string(*obj.FooterLogo)
	}

	footerText := []interface{}{}
	for _, item := range obj.FooterText {
		footerText = append(footerText, SelfRegistrationProfileFooterTextToMap(item))
	}
	result["footer_text"] = footerText

	if obj.HeaderLogo != nil {
		result["header_logo"] = string(*obj.HeaderLogo)
	}

	headerText := []interface{}{}
	for _, item := range obj.HeaderText {
		headerText = append(headerText, SelfRegistrationProfileHeaderTextToMap(item))
	}
	result["header_text"] = headerText

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

	if obj.Meta != nil {
		result["meta"] = []interface{}{metaToMap(obj.Meta)}
	}

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	if obj.NumberOfDaysRedirectUrlIsValid != nil {
		result["number_of_days_redirect_url_is_valid"] = int(*obj.NumberOfDaysRedirectUrlIsValid)
	}

	if obj.Ocid != nil {
		result["ocid"] = string(*obj.Ocid)
	}

	if obj.RedirectUrl != nil {
		result["redirect_url"] = string(*obj.RedirectUrl)
	}

	result["schemas"] = obj.Schemas

	if obj.ShowOnLoginPage != nil {
		result["show_on_login_page"] = bool(*obj.ShowOnLoginPage)
	}

	tags := []interface{}{}
	for _, item := range obj.Tags {
		tags = append(tags, tagsToMap(item))
	}
	result["tags"] = tags

	if obj.TenancyOcid != nil {
		result["tenancy_ocid"] = string(*obj.TenancyOcid)
	}

	userAttributes := []interface{}{}
	for _, item := range obj.UserAttributes {
		userAttributes = append(userAttributes, SelfRegistrationProfileUserAttributesToMap(item))
	}
	result["user_attributes"] = userAttributes

	return result
}

func (s *IdentityDomainsSelfRegistrationProfileResourceCrud) mapToSelfRegistrationProfileAfterSubmitText(fieldKeyFormat string) (oci_identity_domains.SelfRegistrationProfileAfterSubmitText, error) {
	result := oci_identity_domains.SelfRegistrationProfileAfterSubmitText{}

	if default_, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "default")); ok {
		tmp := default_.(bool)
		result.IsDefault = &tmp
	}

	if locale, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "locale")); ok {
		tmp := locale.(string)
		result.Locale = &tmp
	}

	if value, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "value")); ok {
		tmp := value.(string)
		result.Value = &tmp
	}

	return result, nil
}

func SelfRegistrationProfileAfterSubmitTextToMap(obj oci_identity_domains.SelfRegistrationProfileAfterSubmitText) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.IsDefault != nil {
		result["default"] = bool(*obj.IsDefault)
	}

	if obj.Locale != nil {
		result["locale"] = string(*obj.Locale)
	}

	if obj.Value != nil {
		result["value"] = string(*obj.Value)
	}

	return result
}

func (s *IdentityDomainsSelfRegistrationProfileResourceCrud) mapToSelfRegistrationProfileConsentText(fieldKeyFormat string) (oci_identity_domains.SelfRegistrationProfileConsentText, error) {
	result := oci_identity_domains.SelfRegistrationProfileConsentText{}

	if default_, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "default")); ok {
		tmp := default_.(bool)
		result.IsDefault = &tmp
	}

	if locale, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "locale")); ok {
		tmp := locale.(string)
		result.Locale = &tmp
	}

	if value, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "value")); ok {
		tmp := value.(string)
		result.Value = &tmp
	}

	return result, nil
}

func SelfRegistrationProfileConsentTextToMap(obj oci_identity_domains.SelfRegistrationProfileConsentText) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.IsDefault != nil {
		result["default"] = bool(*obj.IsDefault)
	}

	if obj.Locale != nil {
		result["locale"] = string(*obj.Locale)
	}

	if obj.Value != nil {
		result["value"] = string(*obj.Value)
	}

	return result
}

func (s *IdentityDomainsSelfRegistrationProfileResourceCrud) mapToSelfRegistrationProfileDefaultGroups(fieldKeyFormat string) (oci_identity_domains.SelfRegistrationProfileDefaultGroups, error) {
	result := oci_identity_domains.SelfRegistrationProfileDefaultGroups{}

	if ref, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "ref")); ok {
		tmp := ref.(string)
		result.Ref = &tmp
	}

	if display, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "display")); ok {
		tmp := display.(string)
		result.Display = &tmp
	}

	if value, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "value")); ok {
		tmp := value.(string)
		result.Value = &tmp
	}

	return result, nil
}

func SelfRegistrationProfileDefaultGroupsToMap(obj oci_identity_domains.SelfRegistrationProfileDefaultGroups) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Ref != nil {
		result["ref"] = string(*obj.Ref)
	}

	if obj.Display != nil {
		result["display"] = string(*obj.Display)
	}

	if obj.Value != nil {
		result["value"] = string(*obj.Value)
	}

	return result
}

func (s *IdentityDomainsSelfRegistrationProfileResourceCrud) mapToSelfRegistrationProfileDisplayName(fieldKeyFormat string) (oci_identity_domains.SelfRegistrationProfileDisplayName, error) {
	result := oci_identity_domains.SelfRegistrationProfileDisplayName{}

	if default_, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "default")); ok {
		tmp := default_.(bool)
		result.IsDefault = &tmp
	}

	if locale, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "locale")); ok {
		tmp := locale.(string)
		result.Locale = &tmp
	}

	if value, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "value")); ok {
		tmp := value.(string)
		result.Value = &tmp
	}

	return result, nil
}

func SelfRegistrationProfileDisplayNameToMap(obj oci_identity_domains.SelfRegistrationProfileDisplayName) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.IsDefault != nil {
		result["default"] = bool(*obj.IsDefault)
	}

	if obj.Locale != nil {
		result["locale"] = string(*obj.Locale)
	}

	if obj.Value != nil {
		result["value"] = string(*obj.Value)
	}

	return result
}

func (s *IdentityDomainsSelfRegistrationProfileResourceCrud) mapToSelfRegistrationProfileEmailTemplate(fieldKeyFormat string) (oci_identity_domains.SelfRegistrationProfileEmailTemplate, error) {
	result := oci_identity_domains.SelfRegistrationProfileEmailTemplate{}

	if ref, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "ref")); ok {
		tmp := ref.(string)
		result.Ref = &tmp
	}

	if display, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "display")); ok {
		tmp := display.(string)
		result.Display = &tmp
	}

	if value, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "value")); ok {
		tmp := value.(string)
		result.Value = &tmp
	}

	return result, nil
}

func SelfRegistrationProfileEmailTemplateToMap(obj *oci_identity_domains.SelfRegistrationProfileEmailTemplate) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Ref != nil {
		result["ref"] = string(*obj.Ref)
	}

	if obj.Display != nil {
		result["display"] = string(*obj.Display)
	}

	if obj.Value != nil {
		result["value"] = string(*obj.Value)
	}

	return result
}

func (s *IdentityDomainsSelfRegistrationProfileResourceCrud) mapToSelfRegistrationProfileFooterText(fieldKeyFormat string) (oci_identity_domains.SelfRegistrationProfileFooterText, error) {
	result := oci_identity_domains.SelfRegistrationProfileFooterText{}

	if default_, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "default")); ok {
		tmp := default_.(bool)
		result.IsDefault = &tmp
	}

	if locale, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "locale")); ok {
		tmp := locale.(string)
		result.Locale = &tmp
	}

	if value, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "value")); ok {
		tmp := value.(string)
		result.Value = &tmp
	}

	return result, nil
}

func SelfRegistrationProfileFooterTextToMap(obj oci_identity_domains.SelfRegistrationProfileFooterText) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.IsDefault != nil {
		result["default"] = bool(*obj.IsDefault)
	}

	if obj.Locale != nil {
		result["locale"] = string(*obj.Locale)
	}

	if obj.Value != nil {
		result["value"] = string(*obj.Value)
	}

	return result
}

func (s *IdentityDomainsSelfRegistrationProfileResourceCrud) mapToSelfRegistrationProfileHeaderText(fieldKeyFormat string) (oci_identity_domains.SelfRegistrationProfileHeaderText, error) {
	result := oci_identity_domains.SelfRegistrationProfileHeaderText{}

	if default_, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "default")); ok {
		tmp := default_.(bool)
		result.IsDefault = &tmp
	}

	if locale, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "locale")); ok {
		tmp := locale.(string)
		result.Locale = &tmp
	}

	if value, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "value")); ok {
		tmp := value.(string)
		result.Value = &tmp
	}

	return result, nil
}

func SelfRegistrationProfileHeaderTextToMap(obj oci_identity_domains.SelfRegistrationProfileHeaderText) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.IsDefault != nil {
		result["default"] = bool(*obj.IsDefault)
	}

	if obj.Locale != nil {
		result["locale"] = string(*obj.Locale)
	}

	if obj.Value != nil {
		result["value"] = string(*obj.Value)
	}

	return result
}

func (s *IdentityDomainsSelfRegistrationProfileResourceCrud) mapToSelfRegistrationProfileUserAttributes(fieldKeyFormat string) (oci_identity_domains.SelfRegistrationProfileUserAttributes, error) {
	result := oci_identity_domains.SelfRegistrationProfileUserAttributes{}

	if deletable, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "deletable")); ok {
		tmp := deletable.(bool)
		result.Deletable = &tmp
	}

	if fullyQualifiedAttributeName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "fully_qualified_attribute_name")); ok {
		tmp := fullyQualifiedAttributeName.(string)
		result.FullyQualifiedAttributeName = &tmp
	}

	if metadata, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "metadata")); ok {
		tmp := metadata.(string)
		result.Metadata = &tmp
	}

	if seqNumber, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "seq_number")); ok {
		tmp := seqNumber.(int)
		result.SeqNumber = &tmp
	}

	if value, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "value")); ok {
		tmp := value.(string)
		result.Value = &tmp
	}

	return result, nil
}

func SelfRegistrationProfileUserAttributesToMap(obj oci_identity_domains.SelfRegistrationProfileUserAttributes) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Deletable != nil {
		result["deletable"] = bool(*obj.Deletable)
	}

	if obj.FullyQualifiedAttributeName != nil {
		result["fully_qualified_attribute_name"] = string(*obj.FullyQualifiedAttributeName)
	}

	if obj.Metadata != nil {
		result["metadata"] = string(*obj.Metadata)
	}

	if obj.SeqNumber != nil {
		result["seq_number"] = int(*obj.SeqNumber)
	}

	if obj.Value != nil {
		result["value"] = string(*obj.Value)
	}

	return result
}

func (s *IdentityDomainsSelfRegistrationProfileResourceCrud) mapTotags(fieldKeyFormat string) (oci_identity_domains.Tags, error) {
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
