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

func IdentityDomainsCloudGateMappingResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createIdentityDomainsCloudGateMapping,
		Read:     readIdentityDomainsCloudGateMapping,
		Update:   updateIdentityDomainsCloudGateMapping,
		Delete:   deleteIdentityDomainsCloudGateMapping,
		Schema: map[string]*schema.Schema{
			// Required
			"cloud_gate": {
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
						"ref": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"gateway_app": {
				Type:     schema.TypeList,
				Required: true,
				MaxItems: 1,
				MinItems: 1,
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
			"policy_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"resource_prefix": {
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
			"server": {
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
						"ref": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
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
			"nginx_settings": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ocid": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"proxy_pass": {
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
			"upstream_server_group": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional
						"ssl": {
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
						},
						"value": {
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
			"is_opc_service": {
				Type:     schema.TypeBool,
				Computed: true,
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

func createIdentityDomainsCloudGateMapping(d *schema.ResourceData, m interface{}) error {
	sync := &IdentityDomainsCloudGateMappingResourceCrud{}
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

func readIdentityDomainsCloudGateMapping(d *schema.ResourceData, m interface{}) error {
	sync := &IdentityDomainsCloudGateMappingResourceCrud{}
	sync.D = d
	idcsEndpoint, err := getIdcsEndpointForRead(d, "cloudGateMappings")
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

func updateIdentityDomainsCloudGateMapping(d *schema.ResourceData, m interface{}) error {
	sync := &IdentityDomainsCloudGateMappingResourceCrud{}
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

func deleteIdentityDomainsCloudGateMapping(d *schema.ResourceData, m interface{}) error {
	sync := &IdentityDomainsCloudGateMappingResourceCrud{}
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

type IdentityDomainsCloudGateMappingResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_identity_domains.IdentityDomainsClient
	Res                    *oci_identity_domains.CloudGateMapping
	DisableNotFoundRetries bool
}

func (s *IdentityDomainsCloudGateMappingResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *IdentityDomainsCloudGateMappingResourceCrud) Create() error {
	request := oci_identity_domains.CreateCloudGateMappingRequest{}

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

	if cloudGate, ok := s.D.GetOkExists("cloud_gate"); ok {
		if tmpList := cloudGate.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "cloud_gate", 0)
			tmp, err := s.mapToCloudGateMappingCloudGate(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.CloudGate = &tmp
		}
	}

	if description, ok := s.D.GetOkExists("description"); ok {
		tmp := description.(string)
		request.Description = &tmp
	}

	if gatewayApp, ok := s.D.GetOkExists("gateway_app"); ok {
		if tmpList := gatewayApp.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "gateway_app", 0)
			tmp, err := s.mapToCloudGateMappingGatewayApp(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.GatewayApp = &tmp
		}
	}

	if id, ok := s.D.GetOkExists("id"); ok {
		tmp := id.(string)
		request.Id = &tmp
	}

	if isOPCService, ok := s.D.GetOkExists("is_opc_service"); ok {
		tmp := isOPCService.(bool)
		request.IsOPCService = &tmp
	}

	if nginxSettings, ok := s.D.GetOkExists("nginx_settings"); ok {
		tmp := nginxSettings.(string)
		request.NginxSettings = &tmp
	}

	if ocid, ok := s.D.GetOkExists("ocid"); ok {
		tmp := ocid.(string)
		request.Ocid = &tmp
	}

	if policyName, ok := s.D.GetOkExists("policy_name"); ok {
		tmp := policyName.(string)
		request.PolicyName = &tmp
	}

	if proxyPass, ok := s.D.GetOkExists("proxy_pass"); ok {
		tmp := proxyPass.(string)
		request.ProxyPass = &tmp
	}

	if resourcePrefix, ok := s.D.GetOkExists("resource_prefix"); ok {
		tmp := resourcePrefix.(string)
		request.ResourcePrefix = &tmp
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

	if server, ok := s.D.GetOkExists("server"); ok {
		if tmpList := server.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "server", 0)
			tmp, err := s.mapToCloudGateMappingServer(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.Server = &tmp
		}
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

	if upstreamServerGroup, ok := s.D.GetOkExists("upstream_server_group"); ok {
		if tmpList := upstreamServerGroup.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "upstream_server_group", 0)
			tmp, err := s.mapToCloudGateMappingUpstreamServerGroup(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.UpstreamServerGroup = &tmp
		}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "identity_domains")

	response, err := s.Client.CreateCloudGateMapping(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.CloudGateMapping
	return nil
}

func (s *IdentityDomainsCloudGateMappingResourceCrud) Get() error {
	request := oci_identity_domains.GetCloudGateMappingRequest{}

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
	request.CloudGateMappingId = &tmp

	if resourceTypeSchemaVersion, ok := s.D.GetOkExists("resource_type_schema_version"); ok {
		tmp := resourceTypeSchemaVersion.(string)
		request.ResourceTypeSchemaVersion = &tmp
	}

	cloudGateMappingId, err := parseCloudGateMappingCompositeId(s.D.Id())
	if err == nil {
		request.CloudGateMappingId = &cloudGateMappingId
	} else {
		log.Printf("[WARN] Get() unable to parse current ID: %s", s.D.Id())
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "identity_domains")

	response, err := s.Client.GetCloudGateMapping(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.CloudGateMapping
	return nil
}

func (s *IdentityDomainsCloudGateMappingResourceCrud) Update() error {
	request := oci_identity_domains.PutCloudGateMappingRequest{}

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

	if cloudGate, ok := s.D.GetOkExists("cloud_gate"); ok {
		if tmpList := cloudGate.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "cloud_gate", 0)
			tmp, err := s.mapToCloudGateMappingCloudGate(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.CloudGate = &tmp
		}
	}

	tmp := s.D.Id()
	request.CloudGateMappingId = &tmp

	if description, ok := s.D.GetOkExists("description"); ok {
		tmp := description.(string)
		request.Description = &tmp
	}

	if gatewayApp, ok := s.D.GetOkExists("gateway_app"); ok {
		if tmpList := gatewayApp.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "gateway_app", 0)
			tmp, err := s.mapToCloudGateMappingGatewayApp(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.GatewayApp = &tmp
		}
	}

	tmp = s.D.Id()
	request.Id = &tmp

	if isOPCService, ok := s.D.GetOkExists("is_opc_service"); ok {
		tmp := isOPCService.(bool)
		request.IsOPCService = &tmp
	}

	if nginxSettings, ok := s.D.GetOkExists("nginx_settings"); ok {
		tmp := nginxSettings.(string)
		request.NginxSettings = &tmp
	}

	if ocid, ok := s.D.GetOkExists("ocid"); ok {
		tmp := ocid.(string)
		request.Ocid = &tmp
	}

	if policyName, ok := s.D.GetOkExists("policy_name"); ok {
		tmp := policyName.(string)
		request.PolicyName = &tmp
	}

	if proxyPass, ok := s.D.GetOkExists("proxy_pass"); ok {
		tmp := proxyPass.(string)
		request.ProxyPass = &tmp
	}

	if resourcePrefix, ok := s.D.GetOkExists("resource_prefix"); ok {
		tmp := resourcePrefix.(string)
		request.ResourcePrefix = &tmp
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

	if server, ok := s.D.GetOkExists("server"); ok {
		if tmpList := server.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "server", 0)
			tmp, err := s.mapToCloudGateMappingServer(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.Server = &tmp
		}
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

	if upstreamServerGroup, ok := s.D.GetOkExists("upstream_server_group"); ok {
		if tmpList := upstreamServerGroup.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "upstream_server_group", 0)
			tmp, err := s.mapToCloudGateMappingUpstreamServerGroup(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.UpstreamServerGroup = &tmp
		}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "identity_domains")

	response, err := s.Client.PutCloudGateMapping(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.CloudGateMapping
	return nil
}

func (s *IdentityDomainsCloudGateMappingResourceCrud) Delete() error {
	request := oci_identity_domains.DeleteCloudGateMappingRequest{}

	if authorization, ok := s.D.GetOkExists("authorization"); ok {
		tmp := authorization.(string)
		request.Authorization = &tmp
	}

	tmp := s.D.Id()
	request.CloudGateMappingId = &tmp

	if forceDelete, ok := s.D.GetOkExists("force_delete"); ok {
		tmp := forceDelete.(bool)
		request.ForceDelete = &tmp
	}

	if resourceTypeSchemaVersion, ok := s.D.GetOkExists("resource_type_schema_version"); ok {
		tmp := resourceTypeSchemaVersion.(string)
		request.ResourceTypeSchemaVersion = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "identity_domains")

	_, err := s.Client.DeleteCloudGateMapping(context.Background(), request)
	return err
}

func (s *IdentityDomainsCloudGateMappingResourceCrud) SetData() error {

	cloudGateMappingId, err := parseCloudGateMappingCompositeId(s.D.Id())
	if err == nil {
		s.D.SetId(cloudGateMappingId)
	} else {
		log.Printf("[WARN] SetData() unable to parse current ID: %s", s.D.Id())
	}

	if s.Res.CloudGate != nil {
		s.D.Set("cloud_gate", []interface{}{CloudGateMappingCloudGateToMap(s.Res.CloudGate)})
	} else {
		s.D.Set("cloud_gate", nil)
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

	if s.Res.GatewayApp != nil {
		s.D.Set("gateway_app", []interface{}{CloudGateMappingGatewayAppToMap(s.Res.GatewayApp)})
	} else {
		s.D.Set("gateway_app", nil)
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

	if s.Res.IsOPCService != nil {
		s.D.Set("is_opc_service", *s.Res.IsOPCService)
	}

	if s.Res.Meta != nil {
		s.D.Set("meta", []interface{}{metaToMap(s.Res.Meta)})
	} else {
		s.D.Set("meta", nil)
	}

	if s.Res.NginxSettings != nil {
		s.D.Set("nginx_settings", *s.Res.NginxSettings)
	}

	if s.Res.Ocid != nil {
		s.D.Set("ocid", *s.Res.Ocid)
	}

	if s.Res.PolicyName != nil {
		s.D.Set("policy_name", *s.Res.PolicyName)
	}

	if s.Res.ProxyPass != nil {
		s.D.Set("proxy_pass", *s.Res.ProxyPass)
	}

	if s.Res.ResourcePrefix != nil {
		s.D.Set("resource_prefix", *s.Res.ResourcePrefix)
	}

	s.D.Set("schemas", s.Res.Schemas)

	if s.Res.Server != nil {
		s.D.Set("server", []interface{}{CloudGateMappingServerToMap(s.Res.Server)})
	} else {
		s.D.Set("server", nil)
	}

	tags := []interface{}{}
	for _, item := range s.Res.Tags {
		tags = append(tags, tagsToMap(item))
	}
	s.D.Set("tags", tags)

	if s.Res.TenancyOcid != nil {
		s.D.Set("tenancy_ocid", *s.Res.TenancyOcid)
	}

	if s.Res.UpstreamServerGroup != nil {
		s.D.Set("upstream_server_group", []interface{}{CloudGateMappingUpstreamServerGroupToMap(s.Res.UpstreamServerGroup)})
	} else {
		s.D.Set("upstream_server_group", nil)
	}

	return nil
}

func parseCloudGateMappingCompositeId(compositeId string) (cloudGateMappingId string, err error) {
	parts := strings.Split(compositeId, "/")
	match, _ := regexp.MatchString("idcsEndpoint/.*/cloudGateMappings/.*", compositeId)
	if !match || len(parts) != 4 {
		err = fmt.Errorf("illegal compositeId %s encountered", compositeId)
		return
	}
	cloudGateMappingId, _ = url.PathUnescape(parts[3])

	return
}

func CloudGateMappingToMap(obj oci_identity_domains.CloudGateMapping) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.CloudGate != nil {
		result["cloud_gate"] = []interface{}{CloudGateMappingCloudGateToMap(obj.CloudGate)}
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

	if obj.GatewayApp != nil {
		result["gateway_app"] = []interface{}{CloudGateMappingGatewayAppToMap(obj.GatewayApp)}
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

	if obj.IsOPCService != nil {
		result["is_opc_service"] = bool(*obj.IsOPCService)
	}

	if obj.Meta != nil {
		result["meta"] = []interface{}{metaToMap(obj.Meta)}
	}

	if obj.NginxSettings != nil {
		result["nginx_settings"] = string(*obj.NginxSettings)
	}

	if obj.Ocid != nil {
		result["ocid"] = string(*obj.Ocid)
	}

	if obj.PolicyName != nil {
		result["policy_name"] = string(*obj.PolicyName)
	}

	if obj.ProxyPass != nil {
		result["proxy_pass"] = string(*obj.ProxyPass)
	}

	if obj.ResourcePrefix != nil {
		result["resource_prefix"] = string(*obj.ResourcePrefix)
	}

	result["schemas"] = obj.Schemas

	if obj.Server != nil {
		result["server"] = []interface{}{CloudGateMappingServerToMap(obj.Server)}
	}

	tags := []interface{}{}
	for _, item := range obj.Tags {
		tags = append(tags, tagsToMap(item))
	}
	result["tags"] = tags

	if obj.TenancyOcid != nil {
		result["tenancy_ocid"] = string(*obj.TenancyOcid)
	}

	if obj.UpstreamServerGroup != nil {
		result["upstream_server_group"] = []interface{}{CloudGateMappingUpstreamServerGroupToMap(obj.UpstreamServerGroup)}
	}

	return result
}

func (s *IdentityDomainsCloudGateMappingResourceCrud) mapToCloudGateMappingCloudGate(fieldKeyFormat string) (oci_identity_domains.CloudGateMappingCloudGate, error) {
	result := oci_identity_domains.CloudGateMappingCloudGate{}

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

func CloudGateMappingCloudGateToMap(obj *oci_identity_domains.CloudGateMappingCloudGate) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Ref != nil {
		result["ref"] = string(*obj.Ref)
	}

	if obj.Value != nil {
		result["value"] = string(*obj.Value)
	}

	return result
}

func (s *IdentityDomainsCloudGateMappingResourceCrud) mapToCloudGateMappingGatewayApp(fieldKeyFormat string) (oci_identity_domains.CloudGateMappingGatewayApp, error) {
	result := oci_identity_domains.CloudGateMappingGatewayApp{}

	if ref, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "ref")); ok {
		tmp := ref.(string)
		result.Ref = &tmp
	}

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

func CloudGateMappingGatewayAppToMap(obj *oci_identity_domains.CloudGateMappingGatewayApp) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Ref != nil {
		result["ref"] = string(*obj.Ref)
	}

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	if obj.Value != nil {
		result["value"] = string(*obj.Value)
	}

	return result
}

func (s *IdentityDomainsCloudGateMappingResourceCrud) mapToCloudGateMappingServer(fieldKeyFormat string) (oci_identity_domains.CloudGateMappingServer, error) {
	result := oci_identity_domains.CloudGateMappingServer{}

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

func CloudGateMappingServerToMap(obj *oci_identity_domains.CloudGateMappingServer) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Ref != nil {
		result["ref"] = string(*obj.Ref)
	}

	if obj.Value != nil {
		result["value"] = string(*obj.Value)
	}

	return result
}

func (s *IdentityDomainsCloudGateMappingResourceCrud) mapToCloudGateMappingUpstreamServerGroup(fieldKeyFormat string) (oci_identity_domains.CloudGateMappingUpstreamServerGroup, error) {
	result := oci_identity_domains.CloudGateMappingUpstreamServerGroup{}

	if ref, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "ref")); ok {
		tmp := ref.(string)
		result.Ref = &tmp
	}

	if ssl, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "ssl")); ok {
		tmp := ssl.(bool)
		result.Ssl = &tmp
	}

	if value, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "value")); ok {
		tmp := value.(string)
		result.Value = &tmp
	}

	return result, nil
}

func CloudGateMappingUpstreamServerGroupToMap(obj *oci_identity_domains.CloudGateMappingUpstreamServerGroup) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Ref != nil {
		result["ref"] = string(*obj.Ref)
	}

	if obj.Ssl != nil {
		result["ssl"] = bool(*obj.Ssl)
	}

	if obj.Value != nil {
		result["value"] = string(*obj.Value)
	}

	return result
}

func (s *IdentityDomainsCloudGateMappingResourceCrud) mapTotags(fieldKeyFormat string) (oci_identity_domains.Tags, error) {
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
