// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package database_tools

import (
	"context"
	"log"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_database_tools "github.com/oracle/oci-go-sdk/v58/databasetools"
)

func DatabaseToolsDatabaseToolsConnectionDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["database_tools_connection_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(DatabaseToolsDatabaseToolsConnectionResource(), fieldMap, readSingularDatabaseToolsDatabaseToolsConnection)
}

func readSingularDatabaseToolsDatabaseToolsConnection(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseToolsDatabaseToolsConnectionDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseToolsClient()

	return tfresource.ReadResource(sync)
}

type DatabaseToolsDatabaseToolsConnectionDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_database_tools.DatabaseToolsClient
	Res    *oci_database_tools.DatabaseToolsConnection //Res *oci_database_tools.GetDatabaseToolsConnectionResponse
}

func (s *DatabaseToolsDatabaseToolsConnectionDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DatabaseToolsDatabaseToolsConnectionDataSourceCrud) Get() error {
	request := oci_database_tools.GetDatabaseToolsConnectionRequest{}

	if databaseToolsConnectionId, ok := s.D.GetOkExists("database_tools_connection_id"); ok {
		tmp := databaseToolsConnectionId.(string)
		request.DatabaseToolsConnectionId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "database_tools")

	response, err := s.Client.GetDatabaseToolsConnection(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.DatabaseToolsConnection //s.Res = &response
	return nil
}

func (s *DatabaseToolsDatabaseToolsConnectionDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}
	s.D.SetId(tfresource.GenerateDataSourceHashID("DatabaseToolsDatabaseToolsConnectionDataSource-", DatabaseToolsDatabaseToolsConnectionDataSource(), s.D))

	switch v := (*s.Res).(type) {
	case oci_database_tools.DatabaseToolsConnectionOracleDatabase:
		s.D.Set("type", "ORACLE_DATABASE")

		s.D.Set("advanced_properties", v.AdvancedProperties)

		if v.ConnectionString != nil {
			s.D.Set("connection_string", *v.ConnectionString)
		}

		keyStores := []interface{}{}
		for _, item := range v.KeyStores {
			keyStores = append(keyStores, DatabaseToolsKeyStoreToMap(item))
		}
		s.D.Set("key_stores", keyStores)

		if v.PrivateEndpointId != nil {
			s.D.Set("private_endpoint_id", *v.PrivateEndpointId)
		}

		if v.RelatedResource != nil {
			s.D.Set("related_resource", []interface{}{DatabaseToolsRelatedResourceToMap(v.RelatedResource)})
		} else {
			s.D.Set("related_resource", nil)
		}

		if v.UserName != nil {
			s.D.Set("user_name", *v.UserName)
		}

		if v.UserPassword != nil {
			userPasswordArray := []interface{}{}
			if userPasswordMap := DatabaseToolsUserPasswordToMap(&v.UserPassword); userPasswordMap != nil {
				userPasswordArray = append(userPasswordArray, userPasswordMap)
			}
			s.D.Set("user_password", userPasswordArray)
		} else {
			s.D.Set("user_password", nil)
		}

		if v.CompartmentId != nil {
			s.D.Set("compartment_id", *v.CompartmentId)
		}

		if v.DefinedTags != nil {
			s.D.Set("defined_tags", tfresource.DefinedTagsToMap(v.DefinedTags))
		}

		if v.DisplayName != nil {
			s.D.Set("display_name", *v.DisplayName)
		}

		s.D.Set("freeform_tags", v.FreeformTags)

		if v.Id != nil {
			s.D.Set("id", *v.Id)
		}

		if v.LifecycleDetails != nil {
			s.D.Set("lifecycle_details", *v.LifecycleDetails)
		}

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
		log.Printf("[WARN] Received 'type' of unknown type %v", *s.Res)
		return nil
	}
	return nil
}
