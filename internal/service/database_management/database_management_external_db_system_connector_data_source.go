// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package database_management

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_database_management "github.com/oracle/oci-go-sdk/v65/databasemanagement"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DatabaseManagementExternalDbSystemConnectorDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["external_db_system_connector_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(DatabaseManagementExternalDbSystemConnectorResource(), fieldMap, readSingularDatabaseManagementExternalDbSystemConnector)
}

func readSingularDatabaseManagementExternalDbSystemConnector(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseManagementExternalDbSystemConnectorDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DbManagementClient()

	return tfresource.ReadResource(sync)
}

type DatabaseManagementExternalDbSystemConnectorDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_database_management.DbManagementClient
	Res    *oci_database_management.GetExternalDbSystemConnectorResponse
}

func (s *DatabaseManagementExternalDbSystemConnectorDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DatabaseManagementExternalDbSystemConnectorDataSourceCrud) Get() error {
	request := oci_database_management.GetExternalDbSystemConnectorRequest{}

	if externalDbSystemConnectorId, ok := s.D.GetOkExists("external_db_system_connector_id"); ok {
		tmp := externalDbSystemConnectorId.(string)
		request.ExternalDbSystemConnectorId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "database_management")

	response, err := s.Client.GetExternalDbSystemConnector(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *DatabaseManagementExternalDbSystemConnectorDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.GetId())
	switch v := (s.Res.ExternalDbSystemConnector).(type) {
	case oci_database_management.ExternalDbSystemMacsConnector:
		s.D.Set("connector_type", "MACS")

		if v.AgentId != nil {
			s.D.Set("agent_id", *v.AgentId)
		}

		if v.ConnectionInfo != nil {
			connectionInfoArray := []interface{}{}
			if connectionInfoMap := ExternalDbSystemConnectionInfoToMap(&v.ConnectionInfo); connectionInfoMap != nil {
				connectionInfoArray = append(connectionInfoArray, connectionInfoMap)
			}
			s.D.Set("connection_info", connectionInfoArray)
		} else {
			s.D.Set("connection_info", nil)
		}

		if v.DefinedTags != nil {
			s.D.Set("defined_tags", tfresource.DefinedTagsToMap(v.DefinedTags))
		}

		s.D.Set("freeform_tags", v.FreeformTags)

		if v.SystemTags != nil {
			s.D.Set("system_tags", tfresource.SystemTagsToMap(v.SystemTags))
		}

		if v.CompartmentId != nil {
			s.D.Set("compartment_id", *v.CompartmentId)
		}

		if v.ConnectionFailureMessage != nil {
			s.D.Set("connection_failure_message", *v.ConnectionFailureMessage)
		}

		if v.ConnectionStatus != nil {
			s.D.Set("connection_status", *v.ConnectionStatus)
		}

		if v.DisplayName != nil {
			s.D.Set("display_name", *v.DisplayName)
		}

		if v.ExternalDbSystemId != nil {
			s.D.Set("external_db_system_id", *v.ExternalDbSystemId)
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

		if v.TimeUpdated != nil {
			s.D.Set("time_updated", v.TimeUpdated.String())
		}
	default:
		log.Printf("[WARN] Received 'connector_type' of unknown type %v", s.Res.ExternalDbSystemConnector)
		return nil
	}

	return nil
}
