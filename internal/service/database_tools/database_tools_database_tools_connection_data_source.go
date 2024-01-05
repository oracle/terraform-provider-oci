// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package database_tools

import (
	"context"
	"log"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_database_tools "github.com/oracle/oci-go-sdk/v65/databasetools"
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
	Res    *oci_database_tools.GetDatabaseToolsConnectionResponse //Res *oci_database_tools.GetDatabaseToolsConnectionResponse
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

	s.Res = &response //s.Res = &response
	return nil
}

func (s *DatabaseToolsDatabaseToolsConnectionDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}
	s.D.SetId(tfresource.GenerateDataSourceHashID("DatabaseToolsDatabaseToolsConnectionDataSource-", DatabaseToolsDatabaseToolsConnectionDataSource(), s.D))

	switch v := (s.Res.DatabaseToolsConnection).(type) {
	case oci_database_tools.DatabaseToolsConnectionGenericJdbc:
		s.D.Set("type", "GENERIC_JDBC")
		s.D.Set("advanced_properties", v.AdvancedProperties)

		keyStores := []interface{}{}
		for _, item := range v.KeyStores {
			keyStores = append(keyStores, DatabaseToolsKeyStoreGenericJdbcToMap(item))
		}
		s.D.Set("key_stores", keyStores)

		if v.Url != nil {
			s.D.Set("url", *v.Url)
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

		if v.LifecycleDetails != nil {
			s.D.Set("lifecycle_details", *v.LifecycleDetails)
		}

		locks := []interface{}{}
		for _, item := range v.Locks {
			locks = append(locks, ConnectionResourceLockToMap(item))
		}
		s.D.Set("locks", locks)

		s.D.Set("runtime_support", v.RuntimeSupport)

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
	case oci_database_tools.DatabaseToolsConnectionMySql:
		s.D.Set("type", "MYSQL")

		s.D.Set("advanced_properties", v.AdvancedProperties)

		if v.ConnectionString != nil {
			s.D.Set("connection_string", *v.ConnectionString)
		}

		keyStores := []interface{}{}
		for _, item := range v.KeyStores {
			keyStores = append(keyStores, DatabaseToolsKeyStoreMySqlToMap(item))
		}
		s.D.Set("key_stores", keyStores)

		if v.PrivateEndpointId != nil {
			s.D.Set("private_endpoint_id", *v.PrivateEndpointId)
		}

		if v.RelatedResource != nil {
			s.D.Set("related_resource", []interface{}{DatabaseToolsRelatedResourceMySqlToMap(v.RelatedResource)})
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

		if v.LifecycleDetails != nil {
			s.D.Set("lifecycle_details", *v.LifecycleDetails)
		}

		locks := []interface{}{}
		for _, item := range v.Locks {
			locks = append(locks, ConnectionResourceLockToMap(item))
		}
		s.D.Set("locks", locks)

		s.D.Set("runtime_support", v.RuntimeSupport)

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

		if v.ProxyClient != nil {
			proxyClientArray := []interface{}{}
			if proxyClientMap := DatabaseToolsConnectionOracleDatabaseProxyClientToMap(&v.ProxyClient); proxyClientMap != nil {
				proxyClientArray = append(proxyClientArray, proxyClientMap)
			}
			s.D.Set("proxy_client", proxyClientArray)
		} else {
			s.D.Set("proxy_client", nil)
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

		if v.LifecycleDetails != nil {
			s.D.Set("lifecycle_details", *v.LifecycleDetails)
		}

		locks := []interface{}{}
		for _, item := range v.Locks {
			locks = append(locks, ConnectionResourceLockToMap(item))
		}
		s.D.Set("locks", locks)

		s.D.Set("runtime_support", v.RuntimeSupport)

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
	case oci_database_tools.DatabaseToolsConnectionPostgresql:
		s.D.Set("type", "POSTGRESQL")

		s.D.Set("advanced_properties", v.AdvancedProperties)

		if v.ConnectionString != nil {
			s.D.Set("connection_string", *v.ConnectionString)
		}

		keyStores := []interface{}{}
		for _, item := range v.KeyStores {
			keyStores = append(keyStores, DatabaseToolsKeyStorePostgresqlToMap(item))
		}
		s.D.Set("key_stores", keyStores)

		if v.PrivateEndpointId != nil {
			s.D.Set("private_endpoint_id", *v.PrivateEndpointId)
		}

		if v.RelatedResource != nil {
			s.D.Set("related_resource", []interface{}{DatabaseToolsRelatedResourcePostgresqlToMap(v.RelatedResource)})
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

		if v.LifecycleDetails != nil {
			s.D.Set("lifecycle_details", *v.LifecycleDetails)
		}

		locks := []interface{}{}
		for _, item := range v.Locks {
			locks = append(locks, ConnectionResourceLockToMap(item))
		}
		s.D.Set("locks", locks)

		s.D.Set("runtime_support", v.RuntimeSupport)

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
		log.Printf("[WARN] Received 'type' of unknown type %v", s.Res.DatabaseToolsConnection)
		return nil
	}

	return nil
}
