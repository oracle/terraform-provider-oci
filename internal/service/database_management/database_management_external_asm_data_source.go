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

func DatabaseManagementExternalAsmDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["external_asm_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(DatabaseManagementExternalAsmResource(), fieldMap, readSingularDatabaseManagementExternalAsm)
}

func readSingularDatabaseManagementExternalAsm(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseManagementExternalAsmDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DbManagementClient()

	return tfresource.ReadResource(sync)
}

type DatabaseManagementExternalAsmDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_database_management.DbManagementClient
	Res    *oci_database_management.GetExternalAsmResponse
}

func (s *DatabaseManagementExternalAsmDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DatabaseManagementExternalAsmDataSourceCrud) Get() error {
	request := oci_database_management.GetExternalAsmRequest{}

	if externalAsmId, ok := s.D.GetOkExists("external_asm_id"); ok {
		tmp := externalAsmId.(string)
		request.ExternalAsmId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "database_management")

	response, err := s.Client.GetExternalAsm(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *DatabaseManagementExternalAsmDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	s.D.Set("additional_details", s.Res.AdditionalDetails)

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.ComponentName != nil {
		s.D.Set("component_name", *s.Res.ComponentName)
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	if s.Res.ExternalConnectorId != nil {
		s.D.Set("external_connector_id", *s.Res.ExternalConnectorId)
	}

	if s.Res.ExternalDbSystemId != nil {
		s.D.Set("external_db_system_id", *s.Res.ExternalDbSystemId)
	}

	if s.Res.GridHome != nil {
		s.D.Set("grid_home", *s.Res.GridHome)
	}

	if s.Res.IsCluster != nil {
		s.D.Set("is_cluster", *s.Res.IsCluster)
	}

	if s.Res.IsFlexEnabled != nil {
		s.D.Set("is_flex_enabled", *s.Res.IsFlexEnabled)
	}

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	servicedDatabases := []interface{}{}
	for _, item := range s.Res.ServicedDatabases {
		servicedDatabases = append(servicedDatabases, ExternalAsmServicedDatabaseToMap(item))
	}
	s.D.Set("serviced_databases", servicedDatabases)

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	if s.Res.Version != nil {
		s.D.Set("version", *s.Res.Version)
	}

	return nil
}
