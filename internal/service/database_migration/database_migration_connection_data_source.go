// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package database_migration

import (
	"context"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_database_migration "github.com/oracle/oci-go-sdk/v65/databasemigration"
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

	s.D.SetId(*s.Res.Id)

	if s.Res.AdminCredentials != nil {
		s.D.Set("admin_credentials", []interface{}{AdminCredentialsToMap(s.Res.AdminCredentials)})
	} else {
		s.D.Set("admin_credentials", nil)
	}

	if s.Res.CertificateTdn != nil {
		s.D.Set("certificate_tdn", *s.Res.CertificateTdn)
	}

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.ConnectDescriptor != nil {
		s.D.Set("connect_descriptor", []interface{}{ConnectDescriptorToMap(s.Res.ConnectDescriptor)})
	} else {
		s.D.Set("connect_descriptor", nil)
	}

	if s.Res.CredentialsSecretId != nil {
		s.D.Set("credentials_secret_id", *s.Res.CredentialsSecretId)
	}

	if s.Res.DatabaseId != nil {
		s.D.Set("database_id", *s.Res.DatabaseId)
	}

	s.D.Set("database_type", s.Res.DatabaseType)

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	s.D.Set("manual_database_sub_type", s.Res.ManualDatabaseSubType)

	if s.Res.PrivateEndpoint != nil {
		s.D.Set("private_endpoint", []interface{}{PrivateEndpointDetailsToMap(s.Res.PrivateEndpoint)})
	} else {
		s.D.Set("private_endpoint", nil)
	}

	if s.Res.ReplicationCredentials != nil {
		s.D.Set("replication_credentials", []interface{}{AdminCredentialsToMapPassword(s.Res.ReplicationCredentials, s.D)})
	} else {
		s.D.Set("replication_credentials", nil)
	}

	if s.Res.SshDetails != nil {
		s.D.Set("ssh_details", []interface{}{SshDetailsToMap(s.Res.SshDetails)})
	} else {
		s.D.Set("ssh_details", nil)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.SystemTags != nil {
		s.D.Set("system_tags", tfresource.SystemTagsToMap(s.Res.SystemTags))
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	if s.Res.VaultDetails != nil {
		s.D.Set("vault_details", []interface{}{VaultDetailsToMap(s.Res.VaultDetails)})
	} else {
		s.D.Set("vault_details", nil)
	}

	return nil
}
