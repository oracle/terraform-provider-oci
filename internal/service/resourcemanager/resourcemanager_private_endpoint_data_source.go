// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package resourcemanager

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_resourcemanager "github.com/oracle/oci-go-sdk/v65/resourcemanager"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func ResourcemanagerPrivateEndpointDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["private_endpoint_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(ResourcemanagerPrivateEndpointResource(), fieldMap, readSingularResourcemanagerPrivateEndpoint)
}

func readSingularResourcemanagerPrivateEndpoint(d *schema.ResourceData, m interface{}) error {
	sync := &ResourcemanagerPrivateEndpointDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ResourceManagerClient()

	return tfresource.ReadResource(sync)
}

type ResourcemanagerPrivateEndpointDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_resourcemanager.ResourceManagerClient
	Res    *oci_resourcemanager.GetPrivateEndpointResponse
}

func (s *ResourcemanagerPrivateEndpointDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *ResourcemanagerPrivateEndpointDataSourceCrud) Get() error {
	request := oci_resourcemanager.GetPrivateEndpointRequest{}

	if privateEndpointId, ok := s.D.GetOkExists("private_endpoint_id"); ok {
		tmp := privateEndpointId.(string)
		request.PrivateEndpointId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "resourcemanager")

	response, err := s.Client.GetPrivateEndpoint(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *ResourcemanagerPrivateEndpointDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.Description != nil {
		s.D.Set("description", *s.Res.Description)
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	s.D.Set("dns_zones", s.Res.DnsZones)

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.IsUsedWithConfigurationSourceProvider != nil {
		s.D.Set("is_used_with_configuration_source_provider", *s.Res.IsUsedWithConfigurationSourceProvider)
	}

	s.D.Set("nsg_id_list", s.Res.NsgIdList)

	s.D.Set("source_ips", s.Res.SourceIps)

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.SubnetId != nil {
		s.D.Set("subnet_id", *s.Res.SubnetId)
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.VcnId != nil {
		s.D.Set("vcn_id", *s.Res.VcnId)
	}

	return nil
}
