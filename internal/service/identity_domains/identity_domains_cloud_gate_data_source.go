// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package identity_domains

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_identity_domains "github.com/oracle/oci-go-sdk/v65/identitydomains"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func IdentityDomainsCloudGateDataSource() *schema.Resource {
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
	fieldMap["cloud_gate_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	fieldMap["idcs_endpoint"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	fieldMap["resource_type_schema_version"] = &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(IdentityDomainsCloudGateResource(), fieldMap, readSingularIdentityDomainsCloudGate)
}

func readSingularIdentityDomainsCloudGate(d *schema.ResourceData, m interface{}) error {
	sync := &IdentityDomainsCloudGateDataSourceCrud{}
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

type IdentityDomainsCloudGateDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_identity_domains.IdentityDomainsClient
	Res    *oci_identity_domains.GetCloudGateResponse
}

func (s *IdentityDomainsCloudGateDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *IdentityDomainsCloudGateDataSourceCrud) Get() error {
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

	if cloudGateId, ok := s.D.GetOkExists("cloud_gate_id"); ok {
		tmp := cloudGateId.(string)
		request.CloudGateId = &tmp
	}

	if resourceTypeSchemaVersion, ok := s.D.GetOkExists("resource_type_schema_version"); ok {
		tmp := resourceTypeSchemaVersion.(string)
		request.ResourceTypeSchemaVersion = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "identity_domains")

	response, err := s.Client.GetCloudGate(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *IdentityDomainsCloudGateDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

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
