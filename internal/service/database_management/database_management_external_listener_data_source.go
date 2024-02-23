// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package database_management

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_database_management "github.com/oracle/oci-go-sdk/v65/databasemanagement"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DatabaseManagementExternalListenerDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["external_listener_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(DatabaseManagementExternalListenerResource(), fieldMap, readSingularDatabaseManagementExternalListener)
}

func readSingularDatabaseManagementExternalListener(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseManagementExternalListenerDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DbManagementClient()

	return tfresource.ReadResource(sync)
}

type DatabaseManagementExternalListenerDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_database_management.DbManagementClient
	Res    *oci_database_management.GetExternalListenerResponse
}

func (s *DatabaseManagementExternalListenerDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DatabaseManagementExternalListenerDataSourceCrud) Get() error {
	request := oci_database_management.GetExternalListenerRequest{}

	if externalListenerId, ok := s.D.GetOkExists("external_listener_id"); ok {
		tmp := externalListenerId.(string)
		request.ExternalListenerId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "database_management")

	response, err := s.Client.GetExternalListener(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *DatabaseManagementExternalListenerDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	s.D.Set("additional_details", s.Res.AdditionalDetails)
	s.D.Set("additional_details", s.Res.AdditionalDetails)

	if s.Res.AdrHomeDirectory != nil {
		s.D.Set("adr_home_directory", *s.Res.AdrHomeDirectory)
	}

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.ComponentName != nil {
		s.D.Set("component_name", *s.Res.ComponentName)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	endpoints := []interface{}{}
	for _, item := range s.Res.Endpoints {
		endpoints = append(endpoints, ExternalListenerEndpointToMap(item))
	}
	s.D.Set("endpoints", endpoints)

	if s.Res.ExternalConnectorId != nil {
		s.D.Set("external_connector_id", *s.Res.ExternalConnectorId)
	}

	if s.Res.ExternalDbHomeId != nil {
		s.D.Set("external_db_home_id", *s.Res.ExternalDbHomeId)
	}

	if s.Res.ExternalDbNodeId != nil {
		s.D.Set("external_db_node_id", *s.Res.ExternalDbNodeId)
	}

	if s.Res.ExternalDbSystemId != nil {
		s.D.Set("external_db_system_id", *s.Res.ExternalDbSystemId)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.HostName != nil {
		s.D.Set("host_name", *s.Res.HostName)
	}

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	if s.Res.ListenerAlias != nil {
		s.D.Set("listener_alias", *s.Res.ListenerAlias)
	}

	if s.Res.ListenerOraLocation != nil {
		s.D.Set("listener_ora_location", *s.Res.ListenerOraLocation)
	}

	s.D.Set("listener_type", s.Res.ListenerType)

	if s.Res.LogDirectory != nil {
		s.D.Set("log_directory", *s.Res.LogDirectory)
	}

	if s.Res.OracleHome != nil {
		s.D.Set("oracle_home", *s.Res.OracleHome)
	}

	servicedAsms := []interface{}{}
	for _, item := range s.Res.ServicedAsms {
		servicedAsms = append(servicedAsms, ExternalServicedAsmToMap(item))
	}
	s.D.Set("serviced_asms", servicedAsms)

	servicedDatabases := []interface{}{}
	for _, item := range s.Res.ServicedDatabases {
		servicedDatabases = append(servicedDatabases, ExternalListenerServicedDatabaseToMap(item))
	}
	s.D.Set("serviced_databases", servicedDatabases)

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	if s.Res.TraceDirectory != nil {
		s.D.Set("trace_directory", *s.Res.TraceDirectory)
	}

	if s.Res.Version != nil {
		s.D.Set("version", *s.Res.Version)
	}

	return nil
}
