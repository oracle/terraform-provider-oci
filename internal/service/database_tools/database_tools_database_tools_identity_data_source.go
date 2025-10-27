// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package database_tools

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_database_tools "github.com/oracle/oci-go-sdk/v65/databasetools"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DatabaseToolsDatabaseToolsIdentityDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["database_tools_identity_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(DatabaseToolsDatabaseToolsIdentityResource(), fieldMap, readSingularDatabaseToolsDatabaseToolsIdentity)
}

func readSingularDatabaseToolsDatabaseToolsIdentity(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseToolsDatabaseToolsIdentityDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseToolsClient()

	return tfresource.ReadResource(sync)
}

type DatabaseToolsDatabaseToolsIdentityDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_database_tools.DatabaseToolsClient
	Res    *oci_database_tools.GetDatabaseToolsIdentityResponse
}

func (s *DatabaseToolsDatabaseToolsIdentityDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DatabaseToolsDatabaseToolsIdentityDataSourceCrud) Get() error {
	request := oci_database_tools.GetDatabaseToolsIdentityRequest{}

	if databaseToolsIdentityId, ok := s.D.GetOkExists("database_tools_identity_id"); ok {
		tmp := databaseToolsIdentityId.(string)
		request.DatabaseToolsIdentityId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "database_tools")

	response, err := s.Client.GetDatabaseToolsIdentity(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *DatabaseToolsDatabaseToolsIdentityDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.GetId())
	switch v := (s.Res.DatabaseToolsIdentity).(type) {
	case oci_database_tools.DatabaseToolsIdentityOracleDatabaseResourcePrincipal:
		s.D.Set("type", "ORACLE_DATABASE_RESOURCE_PRINCIPAL")

		if v.CredentialKey != nil {
			s.D.Set("credential_key", *v.CredentialKey)
		}

		if v.CompartmentId != nil {
			s.D.Set("compartment_id", *v.CompartmentId)
		}

		if v.DatabaseToolsConnectionId != nil {
			s.D.Set("database_tools_connection_id", *v.DatabaseToolsConnectionId)
		}

		if v.DefinedTags != nil {
			s.D.Set("defined_tags", tfresource.DefinedTagsToMap(v.DefinedTags))
		}

		if v.DisplayName != nil {
			s.D.Set("display_name", *v.DisplayName)
		}

		s.D.Set("freeform_tags", v.FreeformTags)

		if v.LifecycleDetails != nil {
			s.D.Set("lifecycle_details", *v.LifecycleDetails)
		}

		locks := []interface{}{}
		for _, item := range v.Locks {
			locks = append(locks, ResourceLockToMap(item))
		}
		s.D.Set("locks", locks)

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
		log.Printf("[WARN] Received 'type' of unknown type %v", s.Res.DatabaseToolsIdentity)
		return nil
	}

	return nil
}
