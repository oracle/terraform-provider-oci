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

func DatabaseManagementExternalClusterInstanceDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["external_cluster_instance_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(DatabaseManagementExternalClusterInstanceResource(), fieldMap, readSingularDatabaseManagementExternalClusterInstance)
}

func readSingularDatabaseManagementExternalClusterInstance(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseManagementExternalClusterInstanceDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DbManagementClient()

	return tfresource.ReadResource(sync)
}

type DatabaseManagementExternalClusterInstanceDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_database_management.DbManagementClient
	Res    *oci_database_management.GetExternalClusterInstanceResponse
}

func (s *DatabaseManagementExternalClusterInstanceDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DatabaseManagementExternalClusterInstanceDataSourceCrud) Get() error {
	request := oci_database_management.GetExternalClusterInstanceRequest{}

	if externalClusterInstanceId, ok := s.D.GetOkExists("external_cluster_instance_id"); ok {
		tmp := externalClusterInstanceId.(string)
		request.ExternalClusterInstanceId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "database_management")

	response, err := s.Client.GetExternalClusterInstance(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *DatabaseManagementExternalClusterInstanceDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.AdrHomeDirectory != nil {
		s.D.Set("adr_home_directory", *s.Res.AdrHomeDirectory)
	}

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.ComponentName != nil {
		s.D.Set("component_name", *s.Res.ComponentName)
	}

	if s.Res.CrsBaseDirectory != nil {
		s.D.Set("crs_base_directory", *s.Res.CrsBaseDirectory)
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	if s.Res.ExternalClusterId != nil {
		s.D.Set("external_cluster_id", *s.Res.ExternalClusterId)
	}

	if s.Res.ExternalConnectorId != nil {
		s.D.Set("external_connector_id", *s.Res.ExternalConnectorId)
	}

	if s.Res.ExternalDbNodeId != nil {
		s.D.Set("external_db_node_id", *s.Res.ExternalDbNodeId)
	}

	if s.Res.ExternalDbSystemId != nil {
		s.D.Set("external_db_system_id", *s.Res.ExternalDbSystemId)
	}

	if s.Res.HostName != nil {
		s.D.Set("host_name", *s.Res.HostName)
	}

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	s.D.Set("node_role", s.Res.NodeRole)

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	return nil
}
