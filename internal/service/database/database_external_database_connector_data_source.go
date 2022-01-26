// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package database

import (
	"context"
	"log"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_database "github.com/oracle/oci-go-sdk/v56/database"
)

func DatabaseExternalDatabaseConnectorDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["external_database_connector_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(DatabaseExternalDatabaseConnectorResource(), fieldMap, readSingularDatabaseExternalDatabaseConnector)
}

func readSingularDatabaseExternalDatabaseConnector(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseExternalDatabaseConnectorDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()

	return tfresource.ReadResource(sync)
}

type DatabaseExternalDatabaseConnectorDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_database.DatabaseClient
	Res    *oci_database.GetExternalDatabaseConnectorResponse
}

func (s *DatabaseExternalDatabaseConnectorDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DatabaseExternalDatabaseConnectorDataSourceCrud) Get() error {
	request := oci_database.GetExternalDatabaseConnectorRequest{}

	if externalDatabaseConnectorId, ok := s.D.GetOkExists("external_database_connector_id"); ok {
		tmp := externalDatabaseConnectorId.(string)
		request.ExternalDatabaseConnectorId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "database")

	response, err := s.Client.GetExternalDatabaseConnector(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *DatabaseExternalDatabaseConnectorDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.GetId())
	switch v := (s.Res.ExternalDatabaseConnector).(type) {
	case oci_database.ExternalMacsConnector:
		s.D.Set("connector_type", "MACS")

		if v.ConnectionCredentials != nil {
			connectionCredentialsArray := []interface{}{}
			if connectionCredentialsMap := s.DatabaseConnectionCredentialsToMap(&v.ConnectionCredentials); connectionCredentialsMap != nil {
				connectionCredentialsArray = append(connectionCredentialsArray, connectionCredentialsMap)
			}
			s.D.Set("connection_credentials", connectionCredentialsArray)
		} else {
			s.D.Set("connection_credentials", nil)
		}

		if v.ConnectionString != nil {
			s.D.Set("connection_string", []interface{}{DatabaseConnectionStringToMap(v.ConnectionString)})
		} else {
			s.D.Set("connection_string", nil)
		}

		if v.ConnectorAgentId != nil {
			s.D.Set("connector_agent_id", *v.ConnectorAgentId)
		}

		if v.CompartmentId != nil {
			s.D.Set("compartment_id", *v.CompartmentId)
		}

		if v.ConnectionStatus != nil {
			s.D.Set("connection_status", *v.ConnectionStatus)
		}

		if v.DefinedTags != nil {
			s.D.Set("defined_tags", tfresource.DefinedTagsToMap(v.DefinedTags))
		}

		if v.DisplayName != nil {
			s.D.Set("display_name", *v.DisplayName)
		}

		if v.ExternalDatabaseId != nil {
			s.D.Set("external_database_id", *v.ExternalDatabaseId)
		}

		s.D.Set("freeform_tags", v.FreeformTags)

		if v.Id != nil {
			s.D.Set("id", *v.Id)
		}

		if v.LifecycleDetails != nil {
			s.D.Set("lifecycle_details", *v.LifecycleDetails)
		}

		s.D.Set("state", v.LifecycleState)

		if v.TimeConnectionStatusLastUpdated != nil {
			s.D.Set("time_connection_status_last_updated", v.TimeConnectionStatusLastUpdated.String())
		}

		if v.TimeCreated != nil {
			s.D.Set("time_created", v.TimeCreated.String())
		}
	default:
		log.Printf("[WARN] Received 'connector_type' of unknown type %v", *s.Res)
		return nil
	}

	return nil
}

func (s *DatabaseExternalDatabaseConnectorDataSourceCrud) DatabaseConnectionCredentialsToMap(obj *oci_database.DatabaseConnectionCredentials) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (*obj).(type) {
	case oci_database.DatabaseConnectionCredentialsByDetails:
		result["credential_type"] = "DETAILS"

		if v.CredentialName != nil {
			result["credential_name"] = string(*v.CredentialName)
		}

		if password, ok := s.D.GetOkExists("connection_credentials.0.password"); ok && password != nil {
			result["password"] = password.(string)
		}

		result["role"] = string(v.Role)

		if username, ok := s.D.GetOkExists("connection_credentials.0.username"); ok && username != nil {
			result["username"] = username.(string)
		}
	case oci_database.DatabaseConnectionCredentailsByName:
		result["credential_type"] = "NAME_REFERENCE"

		if v.CredentialName != nil {
			result["credential_name"] = string(*v.CredentialName)
		}
	default:
		log.Printf("[WARN] Received 'credential_type' of unknown type %v", *obj)
		return nil
	}

	return result
}
