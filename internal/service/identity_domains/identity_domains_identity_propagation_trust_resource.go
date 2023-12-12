// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package identity_domains

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"net/url"
	"regexp"
	"strings"

	"github.com/oracle/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	oci_identity_domains "github.com/oracle/oci-go-sdk/v65/identitydomains"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func IdentityDomainsIdentityPropagationTrustResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createIdentityDomainsIdentityPropagationTrust,
		Read:     readIdentityDomainsIdentityPropagationTrust,
		Update:   updateIdentityDomainsIdentityPropagationTrust,
		Delete:   deleteIdentityDomainsIdentityPropagationTrust,
		Schema: map[string]*schema.Schema{
			// Required
			"idcs_endpoint": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"issuer": {
				Type:     schema.TypeString,
				Required: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"schemas": {
				Type:     schema.TypeList,
				Required: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"type": {
				Type:     schema.TypeString,
				Required: true,
			},

			// Optional
			"account_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"active": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"allow_impersonation": {
				Type:     schema.TypeBool,
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
			"client_claim_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"client_claim_values": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Set:      tfresource.LiteralTypeHashCodeForSets,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"clock_skew_seconds": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"impersonation_service_users": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Set:      impersonationServiceUsersHashCodeForSets,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"rule": {
							Type:     schema.TypeString,
							Required: true,
						},
						"value": {
							Type:     schema.TypeString,
							Required: true,
						},

						// Optional
						"ocid": {
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
			"keytab": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"secret_ocid": {
							Type:     schema.TypeString,
							Required: true,
						},

						// Optional
						"secret_version": {
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},

						// Computed
					},
				},
			},
			"oauth_clients": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Set:      tfresource.LiteralTypeHashCodeForSets,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"ocid": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"public_certificate": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"public_key_endpoint": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"resource_type_schema_version": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"subject_claim_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"subject_mapping_attribute": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"subject_type": {
				Type:     schema.TypeString,
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

func createIdentityDomainsIdentityPropagationTrust(d *schema.ResourceData, m interface{}) error {
	sync := &IdentityDomainsIdentityPropagationTrustResourceCrud{}
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

func readIdentityDomainsIdentityPropagationTrust(d *schema.ResourceData, m interface{}) error {
	sync := &IdentityDomainsIdentityPropagationTrustResourceCrud{}
	sync.D = d
	idcsEndpoint, err := getIdcsEndpointForRead(d, "identityPropagationTrusts")
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

func updateIdentityDomainsIdentityPropagationTrust(d *schema.ResourceData, m interface{}) error {
	sync := &IdentityDomainsIdentityPropagationTrustResourceCrud{}
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

func deleteIdentityDomainsIdentityPropagationTrust(d *schema.ResourceData, m interface{}) error {
	sync := &IdentityDomainsIdentityPropagationTrustResourceCrud{}
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

type IdentityDomainsIdentityPropagationTrustResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_identity_domains.IdentityDomainsClient
	Res                    *oci_identity_domains.IdentityPropagationTrust
	DisableNotFoundRetries bool
}

func (s *IdentityDomainsIdentityPropagationTrustResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *IdentityDomainsIdentityPropagationTrustResourceCrud) Create() error {
	request := oci_identity_domains.CreateIdentityPropagationTrustRequest{}

	if accountId, ok := s.D.GetOkExists("account_id"); ok {
		tmp := accountId.(string)
		request.AccountId = &tmp
	}

	if active, ok := s.D.GetOkExists("active"); ok {
		tmp := active.(bool)
		request.Active = &tmp
	}

	if allowImpersonation, ok := s.D.GetOkExists("allow_impersonation"); ok {
		tmp := allowImpersonation.(bool)
		request.AllowImpersonation = &tmp
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

	if clientClaimName, ok := s.D.GetOkExists("client_claim_name"); ok {
		tmp := clientClaimName.(string)
		request.ClientClaimName = &tmp
	}

	if clientClaimValues, ok := s.D.GetOkExists("client_claim_values"); ok {
		set := clientClaimValues.(*schema.Set)
		interfaces := set.List()
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("client_claim_values") {
			request.ClientClaimValues = tmp
		}
	}

	if clockSkewSeconds, ok := s.D.GetOkExists("clock_skew_seconds"); ok {
		tmp := clockSkewSeconds.(int)
		request.ClockSkewSeconds = &tmp
	}

	if description, ok := s.D.GetOkExists("description"); ok {
		tmp := description.(string)
		request.Description = &tmp
	}

	if impersonationServiceUsers, ok := s.D.GetOkExists("impersonation_service_users"); ok {
		set := impersonationServiceUsers.(*schema.Set)
		interfaces := set.List()
		tmp := make([]oci_identity_domains.IdentityPropagationTrustImpersonationServiceUsers, len(interfaces))
		for i := range interfaces {
			stateDataIndex := impersonationServiceUsersHashCodeForSets(interfaces[i])
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "impersonation_service_users", stateDataIndex)
			converted, err := s.mapToIdentityPropagationTrustImpersonationServiceUsers(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("impersonation_service_users") {
			request.ImpersonationServiceUsers = tmp
		}
	}

	if issuer, ok := s.D.GetOkExists("issuer"); ok {
		tmp := issuer.(string)
		request.Issuer = &tmp
	}

	if keytab, ok := s.D.GetOkExists("keytab"); ok {
		if tmpList := keytab.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "keytab", 0)
			tmp, err := s.mapToIdentityPropagationTrustKeytab(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.Keytab = &tmp
		}
	}

	if name, ok := s.D.GetOkExists("name"); ok {
		tmp := name.(string)
		request.Name = &tmp
	}

	if oauthClients, ok := s.D.GetOkExists("oauth_clients"); ok {
		set := oauthClients.(*schema.Set)
		interfaces := set.List()
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("oauth_clients") {
			request.OauthClients = tmp
		}
	}

	if ocid, ok := s.D.GetOkExists("ocid"); ok {
		tmp := ocid.(string)
		request.Ocid = &tmp
	}

	if publicCertificate, ok := s.D.GetOkExists("public_certificate"); ok {
		tmp := publicCertificate.(string)
		request.PublicCertificate = &tmp
	}

	if publicKeyEndpoint, ok := s.D.GetOkExists("public_key_endpoint"); ok {
		tmp := publicKeyEndpoint.(string)
		request.PublicKeyEndpoint = &tmp
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

	if subjectClaimName, ok := s.D.GetOkExists("subject_claim_name"); ok {
		tmp := subjectClaimName.(string)
		request.SubjectClaimName = &tmp
	}

	if subjectMappingAttribute, ok := s.D.GetOkExists("subject_mapping_attribute"); ok {
		tmp := subjectMappingAttribute.(string)
		request.SubjectMappingAttribute = &tmp
	}

	if subjectType, ok := s.D.GetOkExists("subject_type"); ok {
		request.SubjectType = oci_identity_domains.IdentityPropagationTrustSubjectTypeEnum(subjectType.(string))
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

	if type_, ok := s.D.GetOkExists("type"); ok {
		request.Type = oci_identity_domains.IdentityPropagationTrustTypeEnum(type_.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "identity_domains")

	response, err := s.Client.CreateIdentityPropagationTrust(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.IdentityPropagationTrust
	return nil
}

func (s *IdentityDomainsIdentityPropagationTrustResourceCrud) Get() error {
	request := oci_identity_domains.GetIdentityPropagationTrustRequest{}

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
	request.IdentityPropagationTrustId = &tmp

	if resourceTypeSchemaVersion, ok := s.D.GetOkExists("resource_type_schema_version"); ok {
		tmp := resourceTypeSchemaVersion.(string)
		request.ResourceTypeSchemaVersion = &tmp
	}

	identityPropagationTrustId, err := parseIdentityPropagationTrustCompositeId(s.D.Id())
	if err == nil {
		request.IdentityPropagationTrustId = &identityPropagationTrustId
	} else {
		log.Printf("[WARN] Get() unable to parse current ID: %s", s.D.Id())
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "identity_domains")

	response, err := s.Client.GetIdentityPropagationTrust(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.IdentityPropagationTrust
	return nil
}

func (s *IdentityDomainsIdentityPropagationTrustResourceCrud) Update() error {
	request := oci_identity_domains.PutIdentityPropagationTrustRequest{}

	if accountId, ok := s.D.GetOkExists("account_id"); ok {
		tmp := accountId.(string)
		request.AccountId = &tmp
	}

	if active, ok := s.D.GetOkExists("active"); ok {
		tmp := active.(bool)
		request.Active = &tmp
	}

	if allowImpersonation, ok := s.D.GetOkExists("allow_impersonation"); ok {
		tmp := allowImpersonation.(bool)
		request.AllowImpersonation = &tmp
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

	if clientClaimName, ok := s.D.GetOkExists("client_claim_name"); ok {
		tmp := clientClaimName.(string)
		request.ClientClaimName = &tmp
	}

	if clientClaimValues, ok := s.D.GetOkExists("client_claim_values"); ok {
		set := clientClaimValues.(*schema.Set)
		interfaces := set.List()
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("client_claim_values") {
			request.ClientClaimValues = tmp
		}
	}

	if clockSkewSeconds, ok := s.D.GetOkExists("clock_skew_seconds"); ok {
		tmp := clockSkewSeconds.(int)
		request.ClockSkewSeconds = &tmp
	}

	if description, ok := s.D.GetOkExists("description"); ok {
		tmp := description.(string)
		request.Description = &tmp
	}

	tmp := s.D.Id()
	request.IdentityPropagationTrustId = &tmp

	if impersonationServiceUsers, ok := s.D.GetOkExists("impersonation_service_users"); ok {
		set := impersonationServiceUsers.(*schema.Set)
		interfaces := set.List()
		tmp := make([]oci_identity_domains.IdentityPropagationTrustImpersonationServiceUsers, len(interfaces))
		for i := range interfaces {
			stateDataIndex := impersonationServiceUsersHashCodeForSets(interfaces[i])
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "impersonation_service_users", stateDataIndex)
			converted, err := s.mapToIdentityPropagationTrustImpersonationServiceUsers(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("impersonation_service_users") {
			request.ImpersonationServiceUsers = tmp
		}
	}

	if issuer, ok := s.D.GetOkExists("issuer"); ok {
		tmp := issuer.(string)
		request.Issuer = &tmp
	}

	if keytab, ok := s.D.GetOkExists("keytab"); ok {
		if tmpList := keytab.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "keytab", 0)
			tmp, err := s.mapToIdentityPropagationTrustKeytab(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.Keytab = &tmp
		}
	}

	// Update() still needs to handle "name" since it's a required field
	if name, ok := s.D.GetOkExists("name"); ok {
		tmp := name.(string)
		request.Name = &tmp
	}

	if oauthClients, ok := s.D.GetOkExists("oauth_clients"); ok {
		set := oauthClients.(*schema.Set)
		interfaces := set.List()
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("oauth_clients") {
			request.OauthClients = tmp
		}
	}

	if ocid, ok := s.D.GetOkExists("ocid"); ok {
		tmp := ocid.(string)
		request.Ocid = &tmp
	}

	if publicCertificate, ok := s.D.GetOkExists("public_certificate"); ok {
		tmp := publicCertificate.(string)
		request.PublicCertificate = &tmp
	}

	if publicKeyEndpoint, ok := s.D.GetOkExists("public_key_endpoint"); ok {
		tmp := publicKeyEndpoint.(string)
		request.PublicKeyEndpoint = &tmp
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

	if subjectClaimName, ok := s.D.GetOkExists("subject_claim_name"); ok {
		tmp := subjectClaimName.(string)
		request.SubjectClaimName = &tmp
	}

	if subjectMappingAttribute, ok := s.D.GetOkExists("subject_mapping_attribute"); ok {
		tmp := subjectMappingAttribute.(string)
		request.SubjectMappingAttribute = &tmp
	}

	if subjectType, ok := s.D.GetOkExists("subject_type"); ok {
		request.SubjectType = oci_identity_domains.IdentityPropagationTrustSubjectTypeEnum(subjectType.(string))
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

	if type_, ok := s.D.GetOkExists("type"); ok {
		request.Type = oci_identity_domains.IdentityPropagationTrustTypeEnum(type_.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "identity_domains")

	response, err := s.Client.PutIdentityPropagationTrust(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.IdentityPropagationTrust
	return nil
}

func (s *IdentityDomainsIdentityPropagationTrustResourceCrud) Delete() error {
	request := oci_identity_domains.DeleteIdentityPropagationTrustRequest{}

	if authorization, ok := s.D.GetOkExists("authorization"); ok {
		tmp := authorization.(string)
		request.Authorization = &tmp
	}

	if forceDelete, ok := s.D.GetOkExists("force_delete"); ok {
		tmp := forceDelete.(bool)
		request.ForceDelete = &tmp
	}

	tmp := s.D.Id()
	request.IdentityPropagationTrustId = &tmp

	if resourceTypeSchemaVersion, ok := s.D.GetOkExists("resource_type_schema_version"); ok {
		tmp := resourceTypeSchemaVersion.(string)
		request.ResourceTypeSchemaVersion = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "identity_domains")

	_, err := s.Client.DeleteIdentityPropagationTrust(context.Background(), request)
	return err
}

func (s *IdentityDomainsIdentityPropagationTrustResourceCrud) SetData() error {

	identityPropagationTrustId, err := parseIdentityPropagationTrustCompositeId(s.D.Id())
	if err == nil {
		s.D.SetId(identityPropagationTrustId)
	} else {
		log.Printf("[WARN] SetData() unable to parse current ID: %s", s.D.Id())
	}

	if s.Res.AccountId != nil {
		s.D.Set("account_id", *s.Res.AccountId)
	}

	if s.Res.Active != nil {
		s.D.Set("active", *s.Res.Active)
	}

	if s.Res.AllowImpersonation != nil {
		s.D.Set("allow_impersonation", *s.Res.AllowImpersonation)
	}

	if s.Res.ClientClaimName != nil {
		s.D.Set("client_claim_name", *s.Res.ClientClaimName)
	}

	clientClaimValues := []interface{}{}
	for _, item := range s.Res.ClientClaimValues {
		clientClaimValues = append(clientClaimValues, item)
	}
	s.D.Set("client_claim_values", schema.NewSet(tfresource.LiteralTypeHashCodeForSets, clientClaimValues))

	if s.Res.ClockSkewSeconds != nil {
		s.D.Set("clock_skew_seconds", *s.Res.ClockSkewSeconds)
	}

	if s.Res.CompartmentOcid != nil {
		s.D.Set("compartment_ocid", *s.Res.CompartmentOcid)
	}

	if s.Res.DeleteInProgress != nil {
		s.D.Set("delete_in_progress", *s.Res.DeleteInProgress)
	}

	if s.Res.Description != nil {
		s.D.Set("description", *s.Res.Description)
	}

	if s.Res.DomainOcid != nil {
		s.D.Set("domain_ocid", *s.Res.DomainOcid)
	}

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

	impersonationServiceUsers := []interface{}{}
	for _, item := range s.Res.ImpersonationServiceUsers {
		impersonationServiceUsers = append(impersonationServiceUsers, IdentityPropagationTrustImpersonationServiceUsersToMap(item))
	}
	s.D.Set("impersonation_service_users", schema.NewSet(impersonationServiceUsersHashCodeForSets, impersonationServiceUsers))

	if s.Res.Issuer != nil {
		s.D.Set("issuer", *s.Res.Issuer)
	}

	if s.Res.Keytab != nil {
		s.D.Set("keytab", []interface{}{IdentityPropagationTrustKeytabToMap(s.Res.Keytab)})
	} else {
		s.D.Set("keytab", nil)
	}

	if s.Res.Meta != nil {
		s.D.Set("meta", []interface{}{metaToMap(s.Res.Meta)})
	} else {
		s.D.Set("meta", nil)
	}

	if s.Res.Name != nil {
		s.D.Set("name", *s.Res.Name)
	}

	oauthClients := []interface{}{}
	for _, item := range s.Res.OauthClients {
		oauthClients = append(oauthClients, item)
	}
	s.D.Set("oauth_clients", schema.NewSet(tfresource.LiteralTypeHashCodeForSets, oauthClients))

	if s.Res.Ocid != nil {
		s.D.Set("ocid", *s.Res.Ocid)
	}

	if s.Res.PublicCertificate != nil {
		s.D.Set("public_certificate", *s.Res.PublicCertificate)
	}

	if s.Res.PublicKeyEndpoint != nil {
		s.D.Set("public_key_endpoint", *s.Res.PublicKeyEndpoint)
	}

	s.D.Set("schemas", s.Res.Schemas)

	if s.Res.SubjectClaimName != nil {
		s.D.Set("subject_claim_name", *s.Res.SubjectClaimName)
	}

	if s.Res.SubjectMappingAttribute != nil {
		s.D.Set("subject_mapping_attribute", *s.Res.SubjectMappingAttribute)
	}

	s.D.Set("subject_type", s.Res.SubjectType)

	tags := []interface{}{}
	for _, item := range s.Res.Tags {
		tags = append(tags, tagsToMap(item))
	}
	s.D.Set("tags", tags)

	if s.Res.TenancyOcid != nil {
		s.D.Set("tenancy_ocid", *s.Res.TenancyOcid)
	}

	s.D.Set("type", s.Res.Type)

	return nil
}

func parseIdentityPropagationTrustCompositeId(compositeId string) (identityPropagationTrustId string, err error) {
	parts := strings.Split(compositeId, "/")
	match, _ := regexp.MatchString("idcsEndpoint/.*/identityPropagationTrusts/.*", compositeId)
	if !match || len(parts) != 4 {
		err = fmt.Errorf("illegal compositeId %s encountered", compositeId)
		return
	}
	identityPropagationTrustId, _ = url.PathUnescape(parts[3])

	return
}

func IdentityPropagationTrustToMap(obj oci_identity_domains.IdentityPropagationTrust, datasource bool) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.AccountId != nil {
		result["account_id"] = string(*obj.AccountId)
	}

	if obj.Active != nil {
		result["active"] = bool(*obj.Active)
	}

	if obj.AllowImpersonation != nil {
		result["allow_impersonation"] = bool(*obj.AllowImpersonation)
	}

	if obj.ClientClaimName != nil {
		result["client_claim_name"] = string(*obj.ClientClaimName)
	}

	clientClaimValues := []interface{}{}
	for _, item := range obj.ClientClaimValues {
		clientClaimValues = append(clientClaimValues, item)
	}
	if datasource {
		result["client_claim_values"] = clientClaimValues
	} else {
		result["client_claim_values"] = schema.NewSet(tfresource.LiteralTypeHashCodeForSets, clientClaimValues)
	}

	if obj.ClockSkewSeconds != nil {
		result["clock_skew_seconds"] = int(*obj.ClockSkewSeconds)
	}

	if obj.CompartmentOcid != nil {
		result["compartment_ocid"] = string(*obj.CompartmentOcid)
	}

	if obj.DeleteInProgress != nil {
		result["delete_in_progress"] = bool(*obj.DeleteInProgress)
	}

	if obj.Description != nil {
		result["description"] = string(*obj.Description)
	}

	if obj.DomainOcid != nil {
		result["domain_ocid"] = string(*obj.DomainOcid)
	}

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

	impersonationServiceUsers := []interface{}{}
	for _, item := range obj.ImpersonationServiceUsers {
		impersonationServiceUsers = append(impersonationServiceUsers, IdentityPropagationTrustImpersonationServiceUsersToMap(item))
	}
	if datasource {
		result["impersonation_service_users"] = impersonationServiceUsers
	} else {
		result["impersonation_service_users"] = schema.NewSet(impersonationServiceUsersHashCodeForSets, impersonationServiceUsers)
	}

	if obj.Issuer != nil {
		result["issuer"] = string(*obj.Issuer)
	}

	if obj.Keytab != nil {
		result["keytab"] = []interface{}{IdentityPropagationTrustKeytabToMap(obj.Keytab)}
	}

	if obj.Meta != nil {
		result["meta"] = []interface{}{metaToMap(obj.Meta)}
	}

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	oauthClients := []interface{}{}
	for _, item := range obj.OauthClients {
		oauthClients = append(oauthClients, item)
	}
	if datasource {
		result["oauth_clients"] = oauthClients
	} else {
		result["oauth_clients"] = schema.NewSet(tfresource.LiteralTypeHashCodeForSets, oauthClients)
	}

	if obj.Ocid != nil {
		result["ocid"] = string(*obj.Ocid)
	}

	if obj.PublicCertificate != nil {
		result["public_certificate"] = string(*obj.PublicCertificate)
	}

	if obj.PublicKeyEndpoint != nil {
		result["public_key_endpoint"] = string(*obj.PublicKeyEndpoint)
	}

	result["schemas"] = obj.Schemas

	if obj.SubjectClaimName != nil {
		result["subject_claim_name"] = string(*obj.SubjectClaimName)
	}

	if obj.SubjectMappingAttribute != nil {
		result["subject_mapping_attribute"] = string(*obj.SubjectMappingAttribute)
	}

	result["subject_type"] = string(obj.SubjectType)

	tags := []interface{}{}
	for _, item := range obj.Tags {
		tags = append(tags, tagsToMap(item))
	}
	result["tags"] = tags

	if obj.TenancyOcid != nil {
		result["tenancy_ocid"] = string(*obj.TenancyOcid)
	}

	result["type"] = string(obj.Type)

	return result
}

func (s *IdentityDomainsIdentityPropagationTrustResourceCrud) mapToIdentityPropagationTrustImpersonationServiceUsers(fieldKeyFormat string) (oci_identity_domains.IdentityPropagationTrustImpersonationServiceUsers, error) {
	result := oci_identity_domains.IdentityPropagationTrustImpersonationServiceUsers{}

	if ref, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "ref")); ok {
		tmp := ref.(string)
		result.Ref = &tmp
	}

	if ocid, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "ocid")); ok {
		tmp := ocid.(string)
		result.Ocid = &tmp
	}

	if rule, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "rule")); ok {
		tmp := rule.(string)
		result.Rule = &tmp
	}

	if value, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "value")); ok {
		tmp := value.(string)
		result.Value = &tmp
	}

	return result, nil
}

func IdentityPropagationTrustImpersonationServiceUsersToMap(obj oci_identity_domains.IdentityPropagationTrustImpersonationServiceUsers) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Ref != nil {
		result["ref"] = string(*obj.Ref)
	}

	if obj.Ocid != nil {
		result["ocid"] = string(*obj.Ocid)
	}

	if obj.Rule != nil {
		result["rule"] = string(*obj.Rule)
	}

	if obj.Value != nil {
		result["value"] = string(*obj.Value)
	}

	return result
}

func (s *IdentityDomainsIdentityPropagationTrustResourceCrud) mapToIdentityPropagationTrustKeytab(fieldKeyFormat string) (oci_identity_domains.IdentityPropagationTrustKeytab, error) {
	result := oci_identity_domains.IdentityPropagationTrustKeytab{}

	if secretOcid, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "secret_ocid")); ok {
		tmp := secretOcid.(string)
		result.SecretOcid = &tmp
	}

	if secretVersion, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "secret_version")); ok {
		tmp := secretVersion.(int)
		result.SecretVersion = &tmp
	}

	return result, nil
}

func IdentityPropagationTrustKeytabToMap(obj *oci_identity_domains.IdentityPropagationTrustKeytab) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.SecretOcid != nil {
		result["secret_ocid"] = string(*obj.SecretOcid)
	}

	if obj.SecretVersion != nil {
		result["secret_version"] = int(*obj.SecretVersion)
	}

	return result
}

func (s *IdentityDomainsIdentityPropagationTrustResourceCrud) mapTotags(fieldKeyFormat string) (oci_identity_domains.Tags, error) {
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

func impersonationServiceUsersHashCodeForSets(v interface{}) int {
	var buf bytes.Buffer
	m := v.(map[string]interface{})
	if ocid, ok := m["ocid"]; ok && ocid != "" {
		buf.WriteString(fmt.Sprintf("%v-", ocid))
	}
	if rule, ok := m["rule"]; ok && rule != "" {
		buf.WriteString(fmt.Sprintf("%v-", rule))
	}
	if value, ok := m["value"]; ok && value != "" {
		buf.WriteString(fmt.Sprintf("%v-", value))
	}
	return utils.GetStringHashcode(buf.String())
}
