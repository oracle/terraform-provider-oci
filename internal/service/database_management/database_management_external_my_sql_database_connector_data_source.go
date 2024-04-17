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

func DatabaseManagementExternalMySqlDatabaseConnectorDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["external_my_sql_database_connector_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(DatabaseManagementExternalMySqlDatabaseConnectorResource(), fieldMap, readSingularDatabaseManagementExternalMySqlDatabaseConnector)
}

func readSingularDatabaseManagementExternalMySqlDatabaseConnector(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseManagementExternalMySqlDatabaseConnectorDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DbManagementClient()

	return tfresource.ReadResource(sync)
}

type DatabaseManagementExternalMySqlDatabaseConnectorDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_database_management.DbManagementClient
	Res    *oci_database_management.GetExternalMySqlDatabaseConnectorResponse
}

func (s *DatabaseManagementExternalMySqlDatabaseConnectorDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DatabaseManagementExternalMySqlDatabaseConnectorDataSourceCrud) Get() error {
	request := oci_database_management.GetExternalMySqlDatabaseConnectorRequest{}

	if externalMySqlDatabaseConnectorId, ok := s.D.GetOkExists("external_my_sql_database_connector_id"); ok {
		tmp := externalMySqlDatabaseConnectorId.(string)
		request.ExternalMySqlDatabaseConnectorId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "database_management")

	response, err := s.Client.GetExternalMySqlDatabaseConnector(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *DatabaseManagementExternalMySqlDatabaseConnectorDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.AssociatedServices != nil {
		s.D.Set("associated_services", *s.Res.AssociatedServices)
	}

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.ConnectionStatus != nil {
		s.D.Set("connection_status", *s.Res.ConnectionStatus)
	}

	s.D.Set("connector_type", s.Res.ConnectorType)

	s.D.Set("credential_type", s.Res.CredentialType)

	if s.Res.ExternalDatabaseId != nil {
		s.D.Set("external_database_id", *s.Res.ExternalDatabaseId)
	}

	if s.Res.HostName != nil {
		s.D.Set("host_name", *s.Res.HostName)
	}

	if s.Res.MacsAgentId != nil {
		s.D.Set("macs_agent_id", *s.Res.MacsAgentId)
	}

	if s.Res.Name != nil {
		s.D.Set("name", *s.Res.Name)
	}

	s.D.Set("network_protocol", s.Res.NetworkProtocol)

	if s.Res.Port != nil {
		s.D.Set("port", *s.Res.Port)
	}

	if s.Res.SourceDatabase != nil {
		s.D.Set("source_database", *s.Res.SourceDatabase)
	}

	s.D.Set("source_database_type", s.Res.SourceDatabaseType)

	if s.Res.SslSecretId != nil {
		s.D.Set("ssl_secret_id", *s.Res.SslSecretId)
	}

	if s.Res.SslSecretName != nil {
		s.D.Set("ssl_secret_name", *s.Res.SslSecretName)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeConnectionStatusUpdated != nil {
		s.D.Set("time_connection_status_updated", s.Res.TimeConnectionStatusUpdated.String())
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	return nil
}
