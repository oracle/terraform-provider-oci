// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package database_tools

import (
	"context"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_database_tools "github.com/oracle/oci-go-sdk/v56/databasetools"
)

func DatabaseToolsDatabaseToolsPrivateEndpointDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["database_tools_private_endpoint_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(DatabaseToolsDatabaseToolsPrivateEndpointResource(), fieldMap, readSingularDatabaseToolsDatabaseToolsPrivateEndpoint)
}

func readSingularDatabaseToolsDatabaseToolsPrivateEndpoint(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseToolsDatabaseToolsPrivateEndpointDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseToolsClient()

	return tfresource.ReadResource(sync)
}

type DatabaseToolsDatabaseToolsPrivateEndpointDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_database_tools.DatabaseToolsClient
	Res    *oci_database_tools.GetDatabaseToolsPrivateEndpointResponse
}

func (s *DatabaseToolsDatabaseToolsPrivateEndpointDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DatabaseToolsDatabaseToolsPrivateEndpointDataSourceCrud) Get() error {
	request := oci_database_tools.GetDatabaseToolsPrivateEndpointRequest{}

	if databaseToolsPrivateEndpointId, ok := s.D.GetOkExists("database_tools_private_endpoint_id"); ok {
		tmp := databaseToolsPrivateEndpointId.(string)
		request.DatabaseToolsPrivateEndpointId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "database_tools")

	response, err := s.Client.GetDatabaseToolsPrivateEndpoint(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *DatabaseToolsDatabaseToolsPrivateEndpointDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	s.D.Set("additional_fqdns", s.Res.AdditionalFqdns)

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

	if s.Res.EndpointFqdn != nil {
		s.D.Set("endpoint_fqdn", *s.Res.EndpointFqdn)
	}

	if s.Res.EndpointServiceId != nil {
		s.D.Set("endpoint_service_id", *s.Res.EndpointServiceId)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	s.D.Set("nsg_ids", s.Res.NsgIds)

	if s.Res.PrivateEndpointIp != nil {
		s.D.Set("private_endpoint_ip", *s.Res.PrivateEndpointIp)
	}

	if s.Res.PrivateEndpointVnicId != nil {
		s.D.Set("private_endpoint_vnic_id", *s.Res.PrivateEndpointVnicId)
	}

	if s.Res.ReverseConnectionConfiguration != nil {
		s.D.Set("reverse_connection_configuration", []interface{}{DatabaseToolsPrivateEndpointReverseConnectionConfigurationToMap(s.Res.ReverseConnectionConfiguration)})
	} else {
		s.D.Set("reverse_connection_configuration", nil)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.SubnetId != nil {
		s.D.Set("subnet_id", *s.Res.SubnetId)
	}

	if s.Res.SystemTags != nil {
		s.D.Set("system_tags", tfresource.SystemTagsToMap(s.Res.SystemTags))
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	if s.Res.VcnId != nil {
		s.D.Set("vcn_id", *s.Res.VcnId)
	}

	return nil
}
