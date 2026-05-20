// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package database_tools

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_database_tools "github.com/oracle/oci-go-sdk/v65/databasetools"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DatabaseToolsDatabaseToolsMcpServerDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["database_tools_mcp_server_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchemaWithContext(DatabaseToolsDatabaseToolsMcpServerResource(), fieldMap, readSingularDatabaseToolsDatabaseToolsMcpServerWithContext)
}

func readSingularDatabaseToolsDatabaseToolsMcpServerWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &DatabaseToolsDatabaseToolsMcpServerDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseToolsClient()

	return tfresource.HandleDiagError(m, tfresource.ReadResourceWithContext(ctx, sync))
}

type DatabaseToolsDatabaseToolsMcpServerDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_database_tools.DatabaseToolsClient
	Res    *oci_database_tools.GetDatabaseToolsMcpServerResponse
}

func (s *DatabaseToolsDatabaseToolsMcpServerDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DatabaseToolsDatabaseToolsMcpServerDataSourceCrud) GetWithContext(ctx context.Context) error {
	request := oci_database_tools.GetDatabaseToolsMcpServerRequest{}

	if databaseToolsMcpServerId, ok := s.D.GetOkExists("database_tools_mcp_server_id"); ok {
		tmp := databaseToolsMcpServerId.(string)
		request.DatabaseToolsMcpServerId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "database_tools")

	response, err := s.Client.GetDatabaseToolsMcpServer(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *DatabaseToolsDatabaseToolsMcpServerDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.GetId())
	switch v := (s.Res.DatabaseToolsMcpServer).(type) {
	case oci_database_tools.DatabaseToolsMcpServerDefault:
		s.D.Set("type", "DEFAULT")

		if v.DomainAppId != nil {
			s.D.Set("domain_app_id", *v.DomainAppId)
		}

		if v.DomainId != nil {
			s.D.Set("domain_id", *v.DomainId)
		}

		if v.Storage != nil {
			storageArray := []interface{}{}
			if storageMap := DatabaseToolsMcpServerStorageToMap(&v.Storage); storageMap != nil {
				storageArray = append(storageArray, storageMap)
			}
			s.D.Set("storage", storageArray)
		} else {
			s.D.Set("storage", nil)
		}

		if v.AccessTokenExpiryInSeconds != nil {
			s.D.Set("access_token_expiry_in_seconds", *v.AccessTokenExpiryInSeconds)
		}

		builtInRoles := []interface{}{}
		for _, item := range v.BuiltInRoles {
			builtInRoles = append(builtInRoles, DatabaseToolsMcpServerBuiltInRoleToMap(item))
		}
		s.D.Set("built_in_roles", builtInRoles)

		if v.CompartmentId != nil {
			s.D.Set("compartment_id", *v.CompartmentId)
		}

		customRoles := []interface{}{}
		for _, item := range v.CustomRoles {
			customRoles = append(customRoles, DatabaseToolsMcpServerCustomRoleToMap(item))
		}
		s.D.Set("custom_roles", customRoles)

		if v.DatabaseToolsConnectionId != nil {
			s.D.Set("database_tools_connection_id", *v.DatabaseToolsConnectionId)
		}

		if v.DefinedTags != nil {
			s.D.Set("defined_tags", tfresource.DefinedTagsToMap(v.DefinedTags))
		}

		if v.Description != nil {
			s.D.Set("description", *v.Description)
		}

		if v.DisplayName != nil {
			s.D.Set("display_name", *v.DisplayName)
		}

		endpoints := []interface{}{}
		for _, item := range v.Endpoints {
			endpoints = append(endpoints, DatabaseToolsMcpServerEndpointToMap(item))
		}
		s.D.Set("endpoints", endpoints)

		s.D.Set("freeform_tags", v.FreeformTags)

		if v.LifecycleDetails != nil {
			s.D.Set("lifecycle_details", *v.LifecycleDetails)
		}

		locks := []interface{}{}
		for _, item := range v.Locks {
			locks = append(locks, DbtoolsMcpServerResourceLockToMap(item))
		}
		s.D.Set("locks", locks)

		if v.RefreshTokenExpiryInSeconds != nil {
			s.D.Set("refresh_token_expiry_in_seconds", *v.RefreshTokenExpiryInSeconds)
		}

		if v.RelatedResource != nil {
			s.D.Set("related_resource", []interface{}{DatabaseToolsMcpServerRelatedResourceToMap(v.RelatedResource)})
		} else {
			s.D.Set("related_resource", nil)
		}

		s.D.Set("runtime_identity", v.RuntimeIdentity)

		s.D.Set("state", v.LifecycleState)

		if v.SystemTags != nil {
			s.D.Set("system_tags", tfresource.SystemTagsToMap(v.SystemTags))
		}

		if v.TimeCreated != nil {
			s.D.Set("time_created", v.TimeCreated.String())
		}

		if v.TimeUpdated != nil {
			s.D.Set("time_updated", v.TimeUpdated.String())
		}
	default:
		log.Printf("[WARN] Received 'type' of unknown type %v", s.Res.DatabaseToolsMcpServer)
		return nil
	}

	return nil
}
