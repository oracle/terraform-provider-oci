// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
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

func IdentityDomainsIdentityProofingProviderTemplateResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createIdentityDomainsIdentityProofingProviderTemplate,
		Read:     readIdentityDomainsIdentityProofingProviderTemplate,
		Update:   updateIdentityDomainsIdentityProofingProviderTemplate,
		Delete:   deleteIdentityDomainsIdentityProofingProviderTemplate,
		Schema: map[string]*schema.Schema{
			// Required
			"identity_proofing_provider_template_provider": {
				Type:     schema.TypeString,
				Required: true,
			},
			"idcs_endpoint": {
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
			"service_type": {
				Type:     schema.TypeSet,
				Required: true,
				Set:      tfresource.LiteralTypeHashCodeForSets,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"verification_url": {
				Type:     schema.TypeString,
				Required: true,
			},

			// Optional
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
			"configuration": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Set:      configurationHashCodeForSets,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"name": {
							Type:     schema.TypeString,
							Required: true,
						},
						"sensitivity": {
							Type:     schema.TypeBool,
							Required: true,
						},
						"type": {
							Type:     schema.TypeString,
							Required: true,
						},

						// Optional

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
			"idcs_locked_by": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

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
						"value": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},

						// Computed
						"_ref": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"idcs_locked_on": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"idcs_locked_operations": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
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

func createIdentityDomainsIdentityProofingProviderTemplate(d *schema.ResourceData, m interface{}) error {
	sync := &IdentityDomainsIdentityProofingProviderTemplateResourceCrud{}
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

func readIdentityDomainsIdentityProofingProviderTemplate(d *schema.ResourceData, m interface{}) error {
	sync := &IdentityDomainsIdentityProofingProviderTemplateResourceCrud{}
	sync.D = d
	idcsEndpoint, err := getIdcsEndpointForRead(d, "identityProofingProviderTemplates")
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

func updateIdentityDomainsIdentityProofingProviderTemplate(d *schema.ResourceData, m interface{}) error {

	sync := &IdentityDomainsIdentityProofingProviderTemplateResourceCrud{}
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

func deleteIdentityDomainsIdentityProofingProviderTemplate(d *schema.ResourceData, m interface{}) error {
	sync := &IdentityDomainsIdentityProofingProviderTemplateResourceCrud{}
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

type IdentityDomainsIdentityProofingProviderTemplateResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_identity_domains.IdentityDomainsClient
	Res                    *oci_identity_domains.IdentityProofingProviderTemplate
	DisableNotFoundRetries bool
}

func (s *IdentityDomainsIdentityProofingProviderTemplateResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *IdentityDomainsIdentityProofingProviderTemplateResourceCrud) Create() error {
	request := oci_identity_domains.CreateIdentityProofingProviderTemplateRequest{}

	if identityProofingProviderTemplateProvider, ok := s.D.GetOkExists("identity_proofing_provider_template_provider"); ok {
		tmp := identityProofingProviderTemplateProvider.(string)
		request.Provider = &tmp
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

	if configuration, ok := s.D.GetOkExists("configuration"); ok {
		set := configuration.(*schema.Set)
		interfaces := set.List()
		tmp := make([]oci_identity_domains.IdentityProofingProviderTemplateConfiguration, len(interfaces))
		for i := range interfaces {
			stateDataIndex := configurationHashCodeForSets(interfaces[i])
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "configuration", stateDataIndex)
			converted, err := s.mapToIdentityProofingProviderTemplateConfiguration(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("configuration") {
			request.Configuration = tmp
		}
	}

	if ocid, ok := s.D.GetOkExists("ocid"); ok {
		tmp := ocid.(string)
		request.Ocid = &tmp
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

	if serviceType, ok := s.D.GetOkExists("service_type"); ok {
		set := serviceType.(*schema.Set)
		interfaces := set.List()
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("service_type") {
			request.ServiceType = tmp
		}
	}

	if verificationUrl, ok := s.D.GetOkExists("verification_url"); ok {
		tmp := verificationUrl.(string)
		request.VerificationUrl = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "identity_domains")

	response, err := s.Client.CreateIdentityProofingProviderTemplate(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.IdentityProofingProviderTemplate
	return nil
}

func (s *IdentityDomainsIdentityProofingProviderTemplateResourceCrud) Get() error {
	request := oci_identity_domains.GetIdentityProofingProviderTemplateRequest{}

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
	request.IdentityProofingProviderTemplateId = &tmp

	if resourceTypeSchemaVersion, ok := s.D.GetOkExists("resource_type_schema_version"); ok {
		tmp := resourceTypeSchemaVersion.(string)
		request.ResourceTypeSchemaVersion = &tmp
	}

	identityProofingProviderTemplateId, err := parseIdentityProofingProviderTemplateCompositeId(s.D.Id())
	if err == nil {
		request.IdentityProofingProviderTemplateId = &identityProofingProviderTemplateId
	} else {
		log.Printf("[WARN] Get() unable to parse current ID: %s", s.D.Id())
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "identity_domains")

	response, err := s.Client.GetIdentityProofingProviderTemplate(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.IdentityProofingProviderTemplate
	return nil
}

func (s *IdentityDomainsIdentityProofingProviderTemplateResourceCrud) Update() error {
	request := oci_identity_domains.PutIdentityProofingProviderTemplateRequest{}

	if identityProofingProviderTemplateProvider, ok := s.D.GetOkExists("identity_proofing_provider_template_provider"); ok {
		tmp := identityProofingProviderTemplateProvider.(string)
		request.Provider = &tmp
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

	if configuration, ok := s.D.GetOkExists("configuration"); ok {
		set := configuration.(*schema.Set)
		interfaces := set.List()
		tmp := make([]oci_identity_domains.IdentityProofingProviderTemplateConfiguration, len(interfaces))
		for i := range interfaces {
			stateDataIndex := configurationHashCodeForSets(interfaces[i])
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "configuration", stateDataIndex)
			converted, err := s.mapToIdentityProofingProviderTemplateConfiguration(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("configuration") {
			request.Configuration = tmp
		}
	}

	tmp := s.D.Id()
	request.IdentityProofingProviderTemplateId = &tmp

	if ocid, ok := s.D.GetOkExists("ocid"); ok {
		tmp := ocid.(string)
		request.Ocid = &tmp
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

	if serviceType, ok := s.D.GetOkExists("service_type"); ok {
		set := serviceType.(*schema.Set)
		interfaces := set.List()
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("service_type") {
			request.ServiceType = tmp
		}
	}

	if verificationUrl, ok := s.D.GetOkExists("verification_url"); ok {
		tmp := verificationUrl.(string)
		request.VerificationUrl = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "identity_domains")

	response, err := s.Client.PutIdentityProofingProviderTemplate(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.IdentityProofingProviderTemplate
	return nil
}

func (s *IdentityDomainsIdentityProofingProviderTemplateResourceCrud) Delete() error {
	request := oci_identity_domains.DeleteIdentityProofingProviderTemplateRequest{}

	if authorization, ok := s.D.GetOkExists("authorization"); ok {
		tmp := authorization.(string)
		request.Authorization = &tmp
	}

	if forceDelete, ok := s.D.GetOkExists("force_delete"); ok {
		tmp := forceDelete.(bool)
		request.ForceDelete = &tmp
	}

	tmp := s.D.Id()
	request.IdentityProofingProviderTemplateId = &tmp

	if resourceTypeSchemaVersion, ok := s.D.GetOkExists("resource_type_schema_version"); ok {
		tmp := resourceTypeSchemaVersion.(string)
		request.ResourceTypeSchemaVersion = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "identity_domains")

	_, err := s.Client.DeleteIdentityProofingProviderTemplate(context.Background(), request)
	return err
}

func (s *IdentityDomainsIdentityProofingProviderTemplateResourceCrud) SetData() error {

	_, err := parseIdentityProofingProviderTemplateCompositeId(s.D.Id())
	if err == nil {
	} else {
		log.Printf("[WARN] SetData() unable to parse current ID: %s", s.D.Id())
	}

	if s.Res.Provider != nil {
		s.D.Set("identity_proofing_provider_template_provider", *s.Res.Provider)
	}

	if s.Res.CompartmentOcid != nil {
		s.D.Set("compartment_ocid", *s.Res.CompartmentOcid)
	}

	configuration := []interface{}{}
	for _, item := range s.Res.Configuration {
		configuration = append(configuration, IdentityProofingProviderTemplateConfigurationToMap(item))
	}
	s.D.Set("configuration", schema.NewSet(configurationHashCodeForSets, configuration))

	if s.Res.DeleteInProgress != nil {
		s.D.Set("delete_in_progress", *s.Res.DeleteInProgress)
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

	if s.Res.IdcsLockedBy != nil {
		s.D.Set("idcs_locked_by", []interface{}{IdentityProofingProviderTemplateIdcsLockedByToMap(s.Res.IdcsLockedBy)})
	} else {
		s.D.Set("idcs_locked_by", nil)
	}

	if s.Res.IdcsLockedOn != nil {
		s.D.Set("idcs_locked_on", *s.Res.IdcsLockedOn)
	}

	s.D.Set("idcs_locked_operations", s.Res.IdcsLockedOperations)

	s.D.Set("idcs_prevented_operations", s.Res.IdcsPreventedOperations)

	if s.Res.Meta != nil {
		s.D.Set("meta", []interface{}{metaToMap(s.Res.Meta)})
	} else {
		s.D.Set("meta", nil)
	}

	if s.Res.Ocid != nil {
		s.D.Set("ocid", *s.Res.Ocid)
	}

	s.D.Set("schemas", s.Res.Schemas)

	serviceType := []interface{}{}
	for _, item := range s.Res.ServiceType {
		serviceType = append(serviceType, item)
	}
	s.D.Set("service_type", schema.NewSet(tfresource.LiteralTypeHashCodeForSets, serviceType))

	if s.Res.TenancyOcid != nil {
		s.D.Set("tenancy_ocid", *s.Res.TenancyOcid)
	}

	if s.Res.VerificationUrl != nil {
		s.D.Set("verification_url", *s.Res.VerificationUrl)
	}

	return nil
}

func parseIdentityProofingProviderTemplateCompositeId(compositeId string) (identityProofingProviderTemplateId string, err error) {
	parts := strings.Split(compositeId, "/")
	match, _ := regexp.MatchString("idcsEndpoint/.*/identityProofingProviderTemplates/.*", compositeId)
	if !match || len(parts) != 4 {
		err = fmt.Errorf("illegal compositeId %s encountered", compositeId)
		return
	}
	//idcsEndpoint, _ = url.PathUnescape(parts[1])
	identityProofingProviderTemplateId, _ = url.PathUnescape(parts[3])

	return
}

func IdentityProofingProviderTemplateToMap(obj oci_identity_domains.IdentityProofingProviderTemplate, datasource bool) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Provider != nil {
		result["identity_proofing_provider_template_provider"] = string(*obj.Provider)
	}

	if obj.CompartmentOcid != nil {
		result["compartment_ocid"] = string(*obj.CompartmentOcid)
	}

	configuration := []interface{}{}
	for _, item := range obj.Configuration {
		configuration = append(configuration, IdentityProofingProviderTemplateConfigurationToMap(item))
	}
	if datasource {
		result["configuration"] = configuration
	} else {
		result["configuration"] = schema.NewSet(configurationHashCodeForSets, configuration)
	}

	if obj.DeleteInProgress != nil {
		result["delete_in_progress"] = bool(*obj.DeleteInProgress)
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

	if obj.IdcsLockedBy != nil {
		result["idcs_locked_by"] = []interface{}{IdentityProofingProviderTemplateIdcsLockedByToMap(obj.IdcsLockedBy)}
	}

	if obj.IdcsLockedOn != nil {
		result["idcs_locked_on"] = string(*obj.IdcsLockedOn)
	}

	result["idcs_locked_operations"] = obj.IdcsLockedOperations

	result["idcs_prevented_operations"] = obj.IdcsPreventedOperations

	if obj.Meta != nil {
		result["meta"] = []interface{}{metaToMap(obj.Meta)}
	}

	if obj.Ocid != nil {
		result["ocid"] = string(*obj.Ocid)
	}

	result["schemas"] = obj.Schemas

	serviceType := []interface{}{}
	for _, item := range obj.ServiceType {
		serviceType = append(serviceType, item)
	}
	if datasource {
		result["service_type"] = serviceType
	} else {
		result["service_type"] = schema.NewSet(tfresource.LiteralTypeHashCodeForSets, serviceType)
	}

	if obj.TenancyOcid != nil {
		result["tenancy_ocid"] = string(*obj.TenancyOcid)
	}

	if obj.VerificationUrl != nil {
		result["verification_url"] = string(*obj.VerificationUrl)
	}

	return result
}

func (s *IdentityDomainsIdentityProofingProviderTemplateResourceCrud) mapToIdentityProofingProviderTemplateConfiguration(fieldKeyFormat string) (oci_identity_domains.IdentityProofingProviderTemplateConfiguration, error) {
	result := oci_identity_domains.IdentityProofingProviderTemplateConfiguration{}

	if name, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "name")); ok {
		tmp := name.(string)
		result.Name = &tmp
	}

	if sensitivity, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "sensitivity")); ok {
		tmp := sensitivity.(bool)
		result.Sensitivity = &tmp
	}

	if type_, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "type")); ok {
		tmp := type_.(string)
		result.Type = &tmp
	}

	return result, nil
}

func IdentityProofingProviderTemplateConfigurationToMap(obj oci_identity_domains.IdentityProofingProviderTemplateConfiguration) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	if obj.Sensitivity != nil {
		result["sensitivity"] = bool(*obj.Sensitivity)
	}

	if obj.Type != nil {
		result["type"] = string(*obj.Type)
	}

	return result
}

func IdentityProofingProviderTemplateIdcsLockedByToMap(obj *oci_identity_domains.IdentityProofingProviderTemplateIdcsLockedBy) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Ref != nil {
		result["_ref"] = string(*obj.Ref)
	}

	if obj.Display != nil {
		result["display"] = string(*obj.Display)
	}

	if obj.Ocid != nil {
		result["ocid"] = string(*obj.Ocid)
	}

	result["type"] = string(obj.Type)

	if obj.Value != nil {
		result["value"] = string(*obj.Value)
	}

	return result
}
