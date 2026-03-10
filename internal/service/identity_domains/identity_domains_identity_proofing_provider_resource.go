// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
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

func IdentityDomainsIdentityProofingProviderResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createIdentityDomainsIdentityProofingProvider,
		Read:     readIdentityDomainsIdentityProofingProvider,
		Update:   updateIdentityDomainsIdentityProofingProvider,
		Delete:   deleteIdentityDomainsIdentityProofingProvider,
		Schema: map[string]*schema.Schema{
			// Required
			"identity_proofing_provider_provider": {
				Type:     schema.TypeString,
				Required: true,
			},
			"claim_mapping": {
				Type:     schema.TypeSet,
				Required: true,
				Set:      claimMappingHashCodeForSets,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"attr_match": {
							Type:     schema.TypeString,
							Required: true,
						},
						"verifiable_claim": {
							Type:     schema.TypeString,
							Required: true,
						},

						// Optional

						// Computed
					},
				},
			},
			"configuration": {
				Type:     schema.TypeSet,
				Required: true,
				Set:      configurationHashCodeForSets,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"name": {
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
			"ocid": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"resource_type_schema_version": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"runtime_data": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Set:      runtimeDataHashCodeForSets,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"attr_name": {
							Type:     schema.TypeString,
							Required: true,
						},
						"attr_value": {
							Type:     schema.TypeString,
							Required: true,
						},

						// Optional

						// Computed
					},
				},
			},
			"status": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
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

func createIdentityDomainsIdentityProofingProvider(d *schema.ResourceData, m interface{}) error {
	sync := &IdentityDomainsIdentityProofingProviderResourceCrud{}
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

func readIdentityDomainsIdentityProofingProvider(d *schema.ResourceData, m interface{}) error {
	sync := &IdentityDomainsIdentityProofingProviderResourceCrud{}
	sync.D = d
	idcsEndpoint, err := getIdcsEndpointForRead(d, "identityProofingProviders")
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

func updateIdentityDomainsIdentityProofingProvider(d *schema.ResourceData, m interface{}) error {
	sync := &IdentityDomainsIdentityProofingProviderResourceCrud{}
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

func deleteIdentityDomainsIdentityProofingProvider(d *schema.ResourceData, m interface{}) error {
	sync := &IdentityDomainsIdentityProofingProviderResourceCrud{}
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

type IdentityDomainsIdentityProofingProviderResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_identity_domains.IdentityDomainsClient
	Res                    *oci_identity_domains.IdentityProofingProvider
	DisableNotFoundRetries bool
}

func (s *IdentityDomainsIdentityProofingProviderResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *IdentityDomainsIdentityProofingProviderResourceCrud) Create() error {
	request := oci_identity_domains.CreateIdentityProofingProviderRequest{}

	if identityProofingProviderProvider, ok := s.D.GetOkExists("identity_proofing_provider_provider"); ok {
		tmp := identityProofingProviderProvider.(string)
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

	if claimMapping, ok := s.D.GetOkExists("claim_mapping"); ok {
		set := claimMapping.(*schema.Set)
		interfaces := set.List()
		tmp := make([]oci_identity_domains.IdentityProofingProviderClaimMapping, len(interfaces))
		for i := range interfaces {
			stateDataIndex := claimMappingHashCodeForSets(interfaces[i])
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "claim_mapping", stateDataIndex)
			converted, err := s.mapToIdentityProofingProviderClaimMapping(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("claim_mapping") {
			request.ClaimMapping = tmp
		}
	}

	if configuration, ok := s.D.GetOkExists("configuration"); ok {
		set := configuration.(*schema.Set)
		interfaces := set.List()
		tmp := make([]oci_identity_domains.IdentityProofingProviderConfiguration, len(interfaces))
		for i := range interfaces {
			stateDataIndex := configurationHashCodeForSets(interfaces[i])
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "configuration", stateDataIndex)
			converted, err := s.mapToIdentityProofingProviderConfiguration(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("configuration") {
			request.Configuration = tmp
		}
	}

	if description, ok := s.D.GetOkExists("description"); ok {
		tmp := description.(string)
		request.Description = &tmp
	}

	if name, ok := s.D.GetOkExists("name"); ok {
		tmp := name.(string)
		request.Name = &tmp
	}

	if ocid, ok := s.D.GetOkExists("ocid"); ok {
		tmp := ocid.(string)
		request.Ocid = &tmp
	}

	if resourceTypeSchemaVersion, ok := s.D.GetOkExists("resource_type_schema_version"); ok {
		tmp := resourceTypeSchemaVersion.(string)
		request.ResourceTypeSchemaVersion = &tmp
	}

	if runtimeData, ok := s.D.GetOkExists("runtime_data"); ok {
		set := runtimeData.(*schema.Set)
		interfaces := set.List()
		tmp := make([]oci_identity_domains.IdentityProofingProviderRuntimeData, len(interfaces))
		for i := range interfaces {
			stateDataIndex := runtimeDataHashCodeForSets(interfaces[i])
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "runtime_data", stateDataIndex)
			converted, err := s.mapToIdentityProofingProviderRuntimeData(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("runtime_data") {
			request.RuntimeData = tmp
		}
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

	if status, ok := s.D.GetOkExists("status"); ok {
		request.Status = oci_identity_domains.IdentityProofingProviderStatusEnum(status.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "identity_domains")

	response, err := s.Client.CreateIdentityProofingProvider(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.IdentityProofingProvider
	return nil
}

func (s *IdentityDomainsIdentityProofingProviderResourceCrud) Get() error {
	request := oci_identity_domains.GetIdentityProofingProviderRequest{}

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
	request.IdentityProofingProviderId = &tmp

	if resourceTypeSchemaVersion, ok := s.D.GetOkExists("resource_type_schema_version"); ok {
		tmp := resourceTypeSchemaVersion.(string)
		request.ResourceTypeSchemaVersion = &tmp
	}

	identityProofingProviderId, err := parseIdentityProofingProviderCompositeId(s.D.Id())
	if err == nil {
		request.IdentityProofingProviderId = &identityProofingProviderId
	} else {
		log.Printf("[WARN] Get() unable to parse current ID: %s", s.D.Id())
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "identity_domains")

	response, err := s.Client.GetIdentityProofingProvider(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.IdentityProofingProvider
	return nil
}

func (s *IdentityDomainsIdentityProofingProviderResourceCrud) Update() error {
	request := oci_identity_domains.PutIdentityProofingProviderRequest{}

	if identityProofingProviderProvider, ok := s.D.GetOkExists("identity_proofing_provider_provider"); ok {
		tmp := identityProofingProviderProvider.(string)
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

	if claimMapping, ok := s.D.GetOkExists("claim_mapping"); ok {
		set := claimMapping.(*schema.Set)
		interfaces := set.List()
		tmp := make([]oci_identity_domains.IdentityProofingProviderClaimMapping, len(interfaces))
		for i := range interfaces {
			stateDataIndex := claimMappingHashCodeForSets(interfaces[i])
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "claim_mapping", stateDataIndex)
			converted, err := s.mapToIdentityProofingProviderClaimMapping(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("claim_mapping") {
			request.ClaimMapping = tmp
		}
	}

	if configuration, ok := s.D.GetOkExists("configuration"); ok {
		set := configuration.(*schema.Set)
		interfaces := set.List()
		tmp := make([]oci_identity_domains.IdentityProofingProviderConfiguration, len(interfaces))
		for i := range interfaces {
			stateDataIndex := configurationHashCodeForSets(interfaces[i])
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "configuration", stateDataIndex)
			converted, err := s.mapToIdentityProofingProviderConfiguration(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("configuration") {
			request.Configuration = tmp
		}
	}

	if description, ok := s.D.GetOkExists("description"); ok {
		tmp := description.(string)
		request.Description = &tmp
	}

	tmp := s.D.Id()
	request.IdentityProofingProviderId = &tmp

	if name, ok := s.D.GetOkExists("name"); ok {
		tmp := name.(string)
		request.Name = &tmp
	}

	if ocid, ok := s.D.GetOkExists("ocid"); ok {
		tmp := ocid.(string)
		request.Ocid = &tmp
	}

	if resourceTypeSchemaVersion, ok := s.D.GetOkExists("resource_type_schema_version"); ok {
		tmp := resourceTypeSchemaVersion.(string)
		request.ResourceTypeSchemaVersion = &tmp
	}

	if runtimeData, ok := s.D.GetOkExists("runtime_data"); ok {
		set := runtimeData.(*schema.Set)
		interfaces := set.List()
		tmp := make([]oci_identity_domains.IdentityProofingProviderRuntimeData, len(interfaces))
		for i := range interfaces {
			stateDataIndex := runtimeDataHashCodeForSets(interfaces[i])
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "runtime_data", stateDataIndex)
			converted, err := s.mapToIdentityProofingProviderRuntimeData(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("runtime_data") {
			request.RuntimeData = tmp
		}
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

	if status, ok := s.D.GetOkExists("status"); ok {
		request.Status = oci_identity_domains.IdentityProofingProviderStatusEnum(status.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "identity_domains")

	response, err := s.Client.PutIdentityProofingProvider(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.IdentityProofingProvider
	return nil
}

func (s *IdentityDomainsIdentityProofingProviderResourceCrud) Delete() error {
	request := oci_identity_domains.DeleteIdentityProofingProviderRequest{}

	if authorization, ok := s.D.GetOkExists("authorization"); ok {
		tmp := authorization.(string)
		request.Authorization = &tmp
	}

	if forceDelete, ok := s.D.GetOkExists("force_delete"); ok {
		tmp := forceDelete.(bool)
		request.ForceDelete = &tmp
	}

	tmp := s.D.Id()
	request.IdentityProofingProviderId = &tmp

	if resourceTypeSchemaVersion, ok := s.D.GetOkExists("resource_type_schema_version"); ok {
		tmp := resourceTypeSchemaVersion.(string)
		request.ResourceTypeSchemaVersion = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "identity_domains")

	_, err := s.Client.DeleteIdentityProofingProvider(context.Background(), request)
	return err
}

func (s *IdentityDomainsIdentityProofingProviderResourceCrud) SetData() error {

	_, err := parseIdentityProofingProviderCompositeId(s.D.Id())
	if err == nil {
	} else {
		log.Printf("[WARN] SetData() unable to parse current ID: %s", s.D.Id())
	}

	if s.Res.Provider != nil {
		s.D.Set("identity_proofing_provider_provider", *s.Res.Provider)
	}

	claimMapping := []interface{}{}
	for _, item := range s.Res.ClaimMapping {
		claimMapping = append(claimMapping, IdentityProofingProviderClaimMappingToMap(item))
	}
	s.D.Set("claim_mapping", schema.NewSet(claimMappingHashCodeForSets, claimMapping))

	if s.Res.CompartmentOcid != nil {
		s.D.Set("compartment_ocid", *s.Res.CompartmentOcid)
	}

	configuration := []interface{}{}
	for _, item := range s.Res.Configuration {
		configuration = append(configuration, IdentityProofingProviderConfigurationToMap(item))
	}
	s.D.Set("configuration", schema.NewSet(configurationHashCodeForSets, configuration))

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

	if s.Res.IdcsLockedBy != nil {
		s.D.Set("idcs_locked_by", []interface{}{IdentityProofingProviderIdcsLockedByToMap(s.Res.IdcsLockedBy)})
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

	if s.Res.Name != nil {
		s.D.Set("name", *s.Res.Name)
	}

	if s.Res.Ocid != nil {
		s.D.Set("ocid", *s.Res.Ocid)
	}

	runtimeData := []interface{}{}
	for _, item := range s.Res.RuntimeData {
		runtimeData = append(runtimeData, IdentityProofingProviderRuntimeDataToMap(item))
	}
	s.D.Set("runtime_data", schema.NewSet(runtimeDataHashCodeForSets, runtimeData))

	s.D.Set("schemas", s.Res.Schemas)

	s.D.Set("status", s.Res.Status)

	if s.Res.TenancyOcid != nil {
		s.D.Set("tenancy_ocid", *s.Res.TenancyOcid)
	}

	return nil
}

func parseIdentityProofingProviderCompositeId(compositeId string) (identityProofingProviderId string, err error) {
	parts := strings.Split(compositeId, "/")
	match, _ := regexp.MatchString("idcsEndpoint/.*/identityProofingProviders/.*", compositeId)
	if !match || len(parts) != 4 {
		err = fmt.Errorf("illegal compositeId %s encountered", compositeId)
		return
	}
	//idcsEndpoint, _ = url.PathUnescape(parts[1])
	identityProofingProviderId, _ = url.PathUnescape(parts[3])

	return
}

func IdentityProofingProviderToMap(obj oci_identity_domains.IdentityProofingProvider, datasource bool) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Provider != nil {
		result["identity_proofing_provider_provider"] = string(*obj.Provider)
	}

	claimMapping := []interface{}{}
	for _, item := range obj.ClaimMapping {
		claimMapping = append(claimMapping, IdentityProofingProviderClaimMappingToMap(item))
	}
	if datasource {
		result["claim_mapping"] = claimMapping
	} else {
		result["claim_mapping"] = schema.NewSet(claimMappingHashCodeForSets, claimMapping)
	}

	if obj.CompartmentOcid != nil {
		result["compartment_ocid"] = string(*obj.CompartmentOcid)
	}

	configuration := []interface{}{}
	for _, item := range obj.Configuration {
		configuration = append(configuration, IdentityProofingProviderConfigurationToMap(item))
	}
	if datasource {
		result["configuration"] = configuration
	} else {
		result["configuration"] = schema.NewSet(configurationHashCodeForSets, configuration)
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

	if obj.IdcsLockedBy != nil {
		result["idcs_locked_by"] = []interface{}{IdentityProofingProviderIdcsLockedByToMap(obj.IdcsLockedBy)}
	}

	if obj.IdcsLockedOn != nil {
		result["idcs_locked_on"] = string(*obj.IdcsLockedOn)
	}

	result["idcs_locked_operations"] = obj.IdcsLockedOperations

	result["idcs_prevented_operations"] = obj.IdcsPreventedOperations

	if obj.Meta != nil {
		result["meta"] = []interface{}{metaToMap(obj.Meta)}
	}

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	if obj.Ocid != nil {
		result["ocid"] = string(*obj.Ocid)
	}

	runtimeData := []interface{}{}
	for _, item := range obj.RuntimeData {
		runtimeData = append(runtimeData, IdentityProofingProviderRuntimeDataToMap(item))
	}
	if datasource {
		result["runtime_data"] = runtimeData
	} else {
		result["runtime_data"] = schema.NewSet(runtimeDataHashCodeForSets, runtimeData)
	}

	result["schemas"] = obj.Schemas

	result["status"] = string(obj.Status)

	if obj.TenancyOcid != nil {
		result["tenancy_ocid"] = string(*obj.TenancyOcid)
	}

	return result
}

func (s *IdentityDomainsIdentityProofingProviderResourceCrud) mapToIdentityProofingProviderClaimMapping(fieldKeyFormat string) (oci_identity_domains.IdentityProofingProviderClaimMapping, error) {
	result := oci_identity_domains.IdentityProofingProviderClaimMapping{}

	if attrMatch, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "attr_match")); ok {
		tmp := attrMatch.(string)
		result.AttrMatch = &tmp
	}

	if verifiableClaim, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "verifiable_claim")); ok {
		tmp := verifiableClaim.(string)
		result.VerifiableClaim = &tmp
	}

	return result, nil
}

func IdentityProofingProviderClaimMappingToMap(obj oci_identity_domains.IdentityProofingProviderClaimMapping) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.AttrMatch != nil {
		result["attr_match"] = string(*obj.AttrMatch)
	}

	if obj.VerifiableClaim != nil {
		result["verifiable_claim"] = string(*obj.VerifiableClaim)
	}

	return result
}

func (s *IdentityDomainsIdentityProofingProviderResourceCrud) mapToIdentityProofingProviderConfiguration(fieldKeyFormat string) (oci_identity_domains.IdentityProofingProviderConfiguration, error) {
	result := oci_identity_domains.IdentityProofingProviderConfiguration{}

	if name, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "name")); ok {
		tmp := name.(string)
		result.Name = &tmp
	}

	if value, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "value")); ok {
		tmp := value.(string)
		result.Value = &tmp
	}

	return result, nil
}

func IdentityProofingProviderConfigurationToMap(obj oci_identity_domains.IdentityProofingProviderConfiguration) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	if obj.Value != nil {
		result["value"] = string(*obj.Value)
	}

	return result
}

func IdentityProofingProviderIdcsLockedByToMap(obj *oci_identity_domains.IdentityProofingProviderIdcsLockedBy) map[string]interface{} {
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

func (s *IdentityDomainsIdentityProofingProviderResourceCrud) mapToIdentityProofingProviderRuntimeData(fieldKeyFormat string) (oci_identity_domains.IdentityProofingProviderRuntimeData, error) {
	result := oci_identity_domains.IdentityProofingProviderRuntimeData{}

	if attrName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "attr_name")); ok {
		tmp := attrName.(string)
		result.AttrName = &tmp
	}

	if attrValue, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "attr_value")); ok {
		tmp := attrValue.(string)
		result.AttrValue = &tmp
	}

	return result, nil
}

func IdentityProofingProviderRuntimeDataToMap(obj oci_identity_domains.IdentityProofingProviderRuntimeData) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.AttrName != nil {
		result["attr_name"] = string(*obj.AttrName)
	}

	if obj.AttrValue != nil {
		result["attr_value"] = string(*obj.AttrValue)
	}

	return result
}

func claimMappingHashCodeForSets(v interface{}) int {
	var buf bytes.Buffer
	m := v.(map[string]interface{})
	if attrMatch, ok := m["attr_match"]; ok && attrMatch != "" {
		buf.WriteString(fmt.Sprintf("%v-", attrMatch))
	}
	if verifiableClaim, ok := m["verifiable_claim"]; ok && verifiableClaim != "" {
		buf.WriteString(fmt.Sprintf("%v-", verifiableClaim))
	}
	return utils.GetStringHashcode(buf.String())
}

func configurationHashCodeForSets(v interface{}) int {
	var buf bytes.Buffer
	m := v.(map[string]interface{})
	if name, ok := m["name"]; ok && name != "" {
		buf.WriteString(fmt.Sprintf("%v-", name))
	}
	if value, ok := m["value"]; ok && value != "" {
		buf.WriteString(fmt.Sprintf("%v-", value))
	}
	return utils.GetStringHashcode(buf.String())
}

func runtimeDataHashCodeForSets(v interface{}) int {
	var buf bytes.Buffer
	m := v.(map[string]interface{})
	if attrName, ok := m["attr_name"]; ok && attrName != "" {
		buf.WriteString(fmt.Sprintf("%v-", attrName))
	}
	if attrValue, ok := m["attr_value"]; ok && attrValue != "" {
		buf.WriteString(fmt.Sprintf("%v-", attrValue))
	}
	return utils.GetStringHashcode(buf.String())
}
