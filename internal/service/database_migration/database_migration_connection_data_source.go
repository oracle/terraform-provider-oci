// // Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// // Licensed under the Mozilla Public License v2.0
package database_migration

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_database_migration "github.com/oracle/oci-go-sdk/v65/databasemigration"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DatabaseMigrationConnectionDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["connection_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(DatabaseMigrationConnectionResource(), fieldMap, readSingularDatabaseMigrationConnection)
}

func readSingularDatabaseMigrationConnection(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseMigrationConnectionDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseMigrationClient()

	return tfresource.ReadResource(sync)
}

type DatabaseMigrationConnectionDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_database_migration.DatabaseMigrationClient
	Res    *oci_database_migration.GetConnectionResponse
}

func (s *DatabaseMigrationConnectionDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DatabaseMigrationConnectionDataSourceCrud) Get() error {
	request := oci_database_migration.GetConnectionRequest{}

	if connectionId, ok := s.D.GetOkExists("connection_id"); ok {
		tmp := connectionId.(string)
		request.ConnectionId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "database_migration")

	response, err := s.Client.GetConnection(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *DatabaseMigrationConnectionDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.GetId())
	switch v := (s.Res.Connection).(type) {
	case oci_database_migration.MysqlConnection:
		s.D.Set("connection_type", "MYSQL")

		additionalAttributes := []interface{}{}
		for _, item := range v.AdditionalAttributes {
			additionalAttributes = append(additionalAttributes, NameValuePairToMap(item))
		}
		s.D.Set("additional_attributes", additionalAttributes)

		if v.DatabaseName != nil {
			s.D.Set("database_name", *v.DatabaseName)
		}

		if v.DbSystemId != nil {
			s.D.Set("db_system_id", *v.DbSystemId)
		}

		if v.Host != nil {
			s.D.Set("host", *v.Host)
		}

		if v.Port != nil {
			s.D.Set("port", *v.Port)
		}

		s.D.Set("security_protocol", v.SecurityProtocol)

		s.D.Set("ssl_mode", v.SslMode)

		s.D.Set("technology_type", v.TechnologyType)

		if v.CompartmentId != nil {
			s.D.Set("compartment_id", *v.CompartmentId)
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

		s.D.Set("freeform_tags", v.FreeformTags)

		ingressIps := []interface{}{}
		for _, item := range v.IngressIps {
			ingressIps = append(ingressIps, IngressIpDetailsToMap(item))
		}
		s.D.Set("ingress_ips", ingressIps)

		if v.KeyId != nil {
			s.D.Set("key_id", *v.KeyId)
		}

		if v.LifecycleDetails != nil {
			s.D.Set("lifecycle_details", *v.LifecycleDetails)
		}

		nsgIds := []interface{}{}
		for _, item := range v.NsgIds {
			nsgIds = append(nsgIds, item)
		}
		s.D.Set("nsg_ids", v.NsgIds)

		if v.Password != nil {
			s.D.Set("password", *v.Password)
		}

		if v.PrivateEndpointId != nil {
			s.D.Set("private_endpoint_id", *v.PrivateEndpointId)
		}

		if v.ReplicationPassword != nil {
			s.D.Set("replication_password", *v.ReplicationPassword)
		}

		if v.ReplicationUsername != nil {
			s.D.Set("replication_username", *v.ReplicationUsername)
		}

		if v.SecretId != nil {
			s.D.Set("secret_id", *v.SecretId)
		}

		s.D.Set("state", v.LifecycleState)

		if v.SubnetId != nil {
			s.D.Set("subnet_id", *v.SubnetId)
		}

		if v.SystemTags != nil {
			s.D.Set("system_tags", tfresource.SystemTagsToMap(v.SystemTags))
		}

		s.D.Set("technology_type", v.TechnologyType)

		if v.TimeCreated != nil {
			s.D.Set("time_created", v.TimeCreated.String())
		}

		if v.TimeUpdated != nil {
			s.D.Set("time_updated", v.TimeUpdated.String())
		}

		if v.Username != nil {
			s.D.Set("username", *v.Username)
		}

		if v.VaultId != nil {
			s.D.Set("vault_id", *v.VaultId)
		}
	case oci_database_migration.OracleConnection:
		s.D.Set("connection_type", "ORACLE")

		if v.ConnectionString != nil {
			s.D.Set("connection_string", *v.ConnectionString)
		}

		if v.DatabaseId != nil {
			s.D.Set("database_id", *v.DatabaseId)
		}

		if v.SshHost != nil {
			s.D.Set("ssh_host", *v.SshHost)
		}

		if v.SshKey != nil {
			s.D.Set("ssh_key", *v.SshKey)
		}

		if v.SshSudoLocation != nil {
			s.D.Set("ssh_sudo_location", *v.SshSudoLocation)
		}

		if v.SshUser != nil {
			s.D.Set("ssh_user", *v.SshUser)
		}

		s.D.Set("technology_type", v.TechnologyType)

		if v.CompartmentId != nil {
			s.D.Set("compartment_id", *v.CompartmentId)
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

		s.D.Set("freeform_tags", v.FreeformTags)

		ingressIps := []interface{}{}
		for _, item := range v.IngressIps {
			ingressIps = append(ingressIps, IngressIpDetailsToMap(item))
		}
		s.D.Set("ingress_ips", ingressIps)

		if v.KeyId != nil {
			s.D.Set("key_id", *v.KeyId)
		}

		if v.LifecycleDetails != nil {
			s.D.Set("lifecycle_details", *v.LifecycleDetails)
		}

		nsgIds := []interface{}{}
		for _, item := range v.NsgIds {
			nsgIds = append(nsgIds, item)
		}
		s.D.Set("nsg_ids", v.NsgIds)

		if v.Password != nil {
			s.D.Set("password", *v.Password)
		}

		if v.PrivateEndpointId != nil {
			s.D.Set("private_endpoint_id", *v.PrivateEndpointId)
		}

		if v.ReplicationPassword != nil {
			s.D.Set("replication_password", *v.ReplicationPassword)
		}

		if v.ReplicationUsername != nil {
			s.D.Set("replication_username", *v.ReplicationUsername)
		}

		if v.SecretId != nil {
			s.D.Set("secret_id", *v.SecretId)
		}

		s.D.Set("state", v.LifecycleState)

		if v.SubnetId != nil {
			s.D.Set("subnet_id", *v.SubnetId)
		}

		if v.SystemTags != nil {
			s.D.Set("system_tags", tfresource.SystemTagsToMap(v.SystemTags))
		}

		s.D.Set("technology_type", v.TechnologyType)

		if v.TimeCreated != nil {
			s.D.Set("time_created", v.TimeCreated.String())
		}

		if v.TimeUpdated != nil {
			s.D.Set("time_updated", v.TimeUpdated.String())
		}

		if v.Username != nil {
			s.D.Set("username", *v.Username)
		}

		if v.VaultId != nil {
			s.D.Set("vault_id", *v.VaultId)
		}
	default:
		log.Printf("[WARN] Received 'connection_type' of unknown type %v", s.Res.Connection)
		return nil
	}

	return nil
}
