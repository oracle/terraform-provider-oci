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

func IdentityDomainsCloudGateResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createIdentityDomainsCloudGate,
		Read:     readIdentityDomainsCloudGate,
		Update:   updateIdentityDomainsCloudGate,
		Delete:   deleteIdentityDomainsCloudGate,
		Schema: map[string]*schema.Schema{
			// Required
			"display_name": {
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

			// Optional
			"active": {
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
			"description": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"last_modified_time": {
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
			"type": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
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
			"mappings": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"resource_prefix": {
							Type:     schema.TypeString,
							Required: true,
						},
						"server_id": {
							Type:     schema.TypeString,
							Required: true,
						},

						// Optional
						"app_id": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"mapping_id": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"name": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"nginx_settings": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"policy_name": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"proxy_pass": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"upstream_server_group_id": {
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
						"ref": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
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
			"oauth_client": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional
						"client_id": {
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
						"ref": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"servers": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"host_name": {
							Type:     schema.TypeString,
							Required: true,
						},
						"port": {
							Type:     schema.TypeInt,
							Required: true,
						},
						"ssl": {
							Type:     schema.TypeBool,
							Required: true,
						},

						// Optional
						"nginx_settings": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"server_id": {
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
			"upstream_server_groups": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"display_name": {
							Type:     schema.TypeString,
							Required: true,
						},
						"value": {
							Type:     schema.TypeString,
							Required: true,
						},

						// Optional
						"nginx_settings": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"ssl": {
							Type:     schema.TypeBool,
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
			"upstream_servers": {
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
						"display_name": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"group_id": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"host_name": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"nginx_settings": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"port": {
							Type:     schema.TypeInt,
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
		},
	}
}

func createIdentityDomainsCloudGate(d *schema.ResourceData, m interface{}) error {
	sync := &IdentityDomainsCloudGateResourceCrud{}
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

func readIdentityDomainsCloudGate(d *schema.ResourceData, m interface{}) error {
	sync := &IdentityDomainsCloudGateResourceCrud{}
	sync.D = d
	idcsEndpoint, err := getIdcsEndpointForRead(d, "cloudGates")
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

func updateIdentityDomainsCloudGate(d *schema.ResourceData, m interface{}) error {
	sync := &IdentityDomainsCloudGateResourceCrud{}
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

func deleteIdentityDomainsCloudGate(d *schema.ResourceData, m interface{}) error {
	sync := &IdentityDomainsCloudGateResourceCrud{}
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

type IdentityDomainsCloudGateResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_identity_domains.IdentityDomainsClient
	Res                    *oci_identity_domains.CloudGate
	DisableNotFoundRetries bool
}

func (s *IdentityDomainsCloudGateResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *IdentityDomainsCloudGateResourceCrud) Create() error {
	request := oci_identity_domains.CreateCloudGateRequest{}

	if active, ok := s.D.GetOkExists("active"); ok {
		tmp := active.(bool)
		request.Active = &tmp
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

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if lastModifiedTime, ok := s.D.GetOkExists("last_modified_time"); ok {
		tmp := lastModifiedTime.(string)
		request.LastModifiedTime = &tmp
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
		request.Type = oci_identity_domains.CloudGateTypeEnum(type_.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "identity_domains")

	response, err := s.Client.CreateCloudGate(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.CloudGate
	return nil
}

func (s *IdentityDomainsCloudGateResourceCrud) Get() error {
	request := oci_identity_domains.GetCloudGateRequest{}

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
	request.CloudGateId = &tmp

	if resourceTypeSchemaVersion, ok := s.D.GetOkExists("resource_type_schema_version"); ok {
		tmp := resourceTypeSchemaVersion.(string)
		request.ResourceTypeSchemaVersion = &tmp
	}

	cloudGateId, err := parseCloudGateCompositeId(s.D.Id())
	if err == nil {
		request.CloudGateId = &cloudGateId
	} else {
		log.Printf("[WARN] Get() unable to parse current ID: %s", s.D.Id())
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "identity_domains")

	response, err := s.Client.GetCloudGate(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.CloudGate
	return nil
}

func (s *IdentityDomainsCloudGateResourceCrud) Update() error {
	request := oci_identity_domains.PutCloudGateRequest{}

	if active, ok := s.D.GetOkExists("active"); ok {
		tmp := active.(bool)
		request.Active = &tmp
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

	tmp := s.D.Id()
	request.CloudGateId = &tmp

	if description, ok := s.D.GetOkExists("description"); ok {
		tmp := description.(string)
		request.Description = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if lastModifiedTime, ok := s.D.GetOkExists("last_modified_time"); ok {
		tmp := lastModifiedTime.(string)
		request.LastModifiedTime = &tmp
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
		request.Type = oci_identity_domains.CloudGateTypeEnum(type_.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "identity_domains")

	response, err := s.Client.PutCloudGate(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.CloudGate
	return nil
}

func (s *IdentityDomainsCloudGateResourceCrud) Delete() error {
	request := oci_identity_domains.DeleteCloudGateRequest{}

	if authorization, ok := s.D.GetOkExists("authorization"); ok {
		tmp := authorization.(string)
		request.Authorization = &tmp
	}

	tmp := s.D.Id()
	request.CloudGateId = &tmp

	if forceDelete, ok := s.D.GetOkExists("force_delete"); ok {
		tmp := forceDelete.(bool)
		request.ForceDelete = &tmp
	}

	if resourceTypeSchemaVersion, ok := s.D.GetOkExists("resource_type_schema_version"); ok {
		tmp := resourceTypeSchemaVersion.(string)
		request.ResourceTypeSchemaVersion = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "identity_domains")

	_, err := s.Client.DeleteCloudGate(context.Background(), request)
	return err
}

func (s *IdentityDomainsCloudGateResourceCrud) SetData() error {

	cloudGateId, err := parseCloudGateCompositeId(s.D.Id())
	if err == nil {
		s.D.SetId(cloudGateId)
	} else {
		log.Printf("[WARN] SetData() unable to parse current ID: %s", s.D.Id())
	}

	if s.Res.Active != nil {
		s.D.Set("active", *s.Res.Active)
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

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
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

	if s.Res.IsOPCService != nil {
		s.D.Set("is_opc_service", *s.Res.IsOPCService)
	}

	if s.Res.LastModifiedTime != nil {
		s.D.Set("last_modified_time", *s.Res.LastModifiedTime)
	}

	mappings := []interface{}{}
	for _, item := range s.Res.Mappings {
		mappings = append(mappings, CloudGateExtMappingsToMap(item))
	}
	s.D.Set("mappings", mappings)

	if s.Res.Meta != nil {
		s.D.Set("meta", []interface{}{metaToMap(s.Res.Meta)})
	} else {
		s.D.Set("meta", nil)
	}

	if s.Res.OauthClient != nil {
		s.D.Set("oauth_client", []interface{}{CloudGateOauthClientToMap(s.Res.OauthClient)})
	} else {
		s.D.Set("oauth_client", nil)
	}

	if s.Res.Ocid != nil {
		s.D.Set("ocid", *s.Res.Ocid)
	}

	s.D.Set("schemas", s.Res.Schemas)

	servers := []interface{}{}
	for _, item := range s.Res.Servers {
		servers = append(servers, CloudGateExtServersToMap(item))
	}
	s.D.Set("servers", servers)

	tags := []interface{}{}
	for _, item := range s.Res.Tags {
		tags = append(tags, tagsToMap(item))
	}
	s.D.Set("tags", tags)

	if s.Res.TenancyOcid != nil {
		s.D.Set("tenancy_ocid", *s.Res.TenancyOcid)
	}

	s.D.Set("type", s.Res.Type)

	upstreamServerGroups := []interface{}{}
	for _, item := range s.Res.UpstreamServerGroups {
		upstreamServerGroups = append(upstreamServerGroups, CloudGateExtUpstreamServerGroupsToMap(item))
	}
	s.D.Set("upstream_server_groups", upstreamServerGroups)

	upstreamServers := []interface{}{}
	for _, item := range s.Res.UpstreamServers {
		upstreamServers = append(upstreamServers, CloudGateExtUpstreamServersToMap(item))
	}
	s.D.Set("upstream_servers", upstreamServers)

	return nil
}

func parseCloudGateCompositeId(compositeId string) (cloudGateId string, err error) {
	parts := strings.Split(compositeId, "/")
	match, _ := regexp.MatchString("idcsEndpoint/.*/cloudGates/.*", compositeId)
	if !match || len(parts) != 4 {
		err = fmt.Errorf("illegal compositeId %s encountered", compositeId)
		return
	}
	cloudGateId, _ = url.PathUnescape(parts[3])

	return
}

func CloudGateToMap(obj oci_identity_domains.CloudGate) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Active != nil {
		result["active"] = bool(*obj.Active)
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

	if obj.DisplayName != nil {
		result["display_name"] = string(*obj.DisplayName)
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

	if obj.IsOPCService != nil {
		result["is_opc_service"] = bool(*obj.IsOPCService)
	}

	if obj.LastModifiedTime != nil {
		result["last_modified_time"] = string(*obj.LastModifiedTime)
	}

	mappings := []interface{}{}
	for _, item := range obj.Mappings {
		mappings = append(mappings, CloudGateExtMappingsToMap(item))
	}
	result["mappings"] = mappings

	if obj.Meta != nil {
		result["meta"] = []interface{}{metaToMap(obj.Meta)}
	}

	if obj.OauthClient != nil {
		result["oauth_client"] = []interface{}{CloudGateOauthClientToMap(obj.OauthClient)}
	}

	if obj.Ocid != nil {
		result["ocid"] = string(*obj.Ocid)
	}

	result["schemas"] = obj.Schemas

	servers := []interface{}{}
	for _, item := range obj.Servers {
		servers = append(servers, CloudGateExtServersToMap(item))
	}
	result["servers"] = servers

	tags := []interface{}{}
	for _, item := range obj.Tags {
		tags = append(tags, tagsToMap(item))
	}
	result["tags"] = tags

	if obj.TenancyOcid != nil {
		result["tenancy_ocid"] = string(*obj.TenancyOcid)
	}

	result["type"] = string(obj.Type)

	upstreamServerGroups := []interface{}{}
	for _, item := range obj.UpstreamServerGroups {
		upstreamServerGroups = append(upstreamServerGroups, CloudGateExtUpstreamServerGroupsToMap(item))
	}
	result["upstream_server_groups"] = upstreamServerGroups

	upstreamServers := []interface{}{}
	for _, item := range obj.UpstreamServers {
		upstreamServers = append(upstreamServers, CloudGateExtUpstreamServersToMap(item))
	}
	result["upstream_servers"] = upstreamServers

	return result
}

func CloudGateExtMappingsToMap(obj oci_identity_domains.CloudGateExtMappings) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Ref != nil {
		result["ref"] = string(*obj.Ref)
	}

	if obj.AppId != nil {
		result["app_id"] = string(*obj.AppId)
	}

	if obj.MappingId != nil {
		result["mapping_id"] = string(*obj.MappingId)
	}

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	if obj.NginxSettings != nil {
		result["nginx_settings"] = string(*obj.NginxSettings)
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

	if obj.ServerId != nil {
		result["server_id"] = string(*obj.ServerId)
	}

	if obj.UpstreamServerGroupId != nil {
		result["upstream_server_group_id"] = string(*obj.UpstreamServerGroupId)
	}

	if obj.Value != nil {
		result["value"] = string(*obj.Value)
	}

	return result
}

func CloudGateExtServersToMap(obj oci_identity_domains.CloudGateExtServers) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.HostName != nil {
		result["host_name"] = string(*obj.HostName)
	}

	if obj.NginxSettings != nil {
		result["nginx_settings"] = string(*obj.NginxSettings)
	}

	if obj.Port != nil {
		result["port"] = int(*obj.Port)
	}

	if obj.ServerId != nil {
		result["server_id"] = string(*obj.ServerId)
	}

	if obj.Ssl != nil {
		result["ssl"] = bool(*obj.Ssl)
	}

	return result
}

func CloudGateExtUpstreamServerGroupsToMap(obj oci_identity_domains.CloudGateExtUpstreamServerGroups) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Ref != nil {
		result["ref"] = string(*obj.Ref)
	}

	if obj.DisplayName != nil {
		result["display_name"] = string(*obj.DisplayName)
	}

	if obj.NginxSettings != nil {
		result["nginx_settings"] = string(*obj.NginxSettings)
	}

	if obj.Ssl != nil {
		result["ssl"] = bool(*obj.Ssl)
	}

	if obj.Value != nil {
		result["value"] = string(*obj.Value)
	}

	return result
}

func CloudGateExtUpstreamServersToMap(obj oci_identity_domains.CloudGateExtUpstreamServers) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Ref != nil {
		result["ref"] = string(*obj.Ref)
	}

	if obj.DisplayName != nil {
		result["display_name"] = string(*obj.DisplayName)
	}

	if obj.GroupId != nil {
		result["group_id"] = string(*obj.GroupId)
	}

	if obj.HostName != nil {
		result["host_name"] = string(*obj.HostName)
	}

	if obj.NginxSettings != nil {
		result["nginx_settings"] = string(*obj.NginxSettings)
	}

	if obj.Port != nil {
		result["port"] = int(*obj.Port)
	}

	if obj.Value != nil {
		result["value"] = string(*obj.Value)
	}

	return result
}

func CloudGateOauthClientToMap(obj *oci_identity_domains.CloudGateOauthClient) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Ref != nil {
		result["ref"] = string(*obj.Ref)
	}

	if obj.ClientId != nil {
		result["client_id"] = string(*obj.ClientId)
	}

	if obj.Value != nil {
		result["value"] = string(*obj.Value)
	}

	return result
}

func (s *IdentityDomainsCloudGateResourceCrud) mapTotags(fieldKeyFormat string) (oci_identity_domains.Tags, error) {
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
