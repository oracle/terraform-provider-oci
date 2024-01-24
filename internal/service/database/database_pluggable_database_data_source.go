// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package database

import (
	"context"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_database "github.com/oracle/oci-go-sdk/v65/database"
)

func DatabasePluggableDatabaseDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["pluggable_database_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(DatabasePluggableDatabaseResource(), fieldMap, readSingularDatabasePluggableDatabase)
}

func readSingularDatabasePluggableDatabase(d *schema.ResourceData, m interface{}) error {
	sync := &DatabasePluggableDatabaseDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()

	return tfresource.ReadResource(sync)
}

type DatabasePluggableDatabaseDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_database.DatabaseClient
	Res    *oci_database.GetPluggableDatabaseResponse
}

func (s *DatabasePluggableDatabaseDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DatabasePluggableDatabaseDataSourceCrud) Get() error {
	request := oci_database.GetPluggableDatabaseRequest{}

	if pluggableDatabaseId, ok := s.D.GetOkExists("pluggable_database_id"); ok {
		tmp := pluggableDatabaseId.(string)
		request.PluggableDatabaseId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "database")

	response, err := s.Client.GetPluggableDatabase(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *DatabasePluggableDatabaseDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.ConnectionStrings != nil {
		s.D.Set("connection_strings", []interface{}{PluggableDatabaseConnectionStringsToMap(s.Res.ConnectionStrings)})
	} else {
		s.D.Set("connection_strings", nil)
	}

	if s.Res.ContainerDatabaseId != nil {
		s.D.Set("container_database_id", *s.Res.ContainerDatabaseId)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.IsRestricted != nil {
		s.D.Set("is_restricted", *s.Res.IsRestricted)
	}

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	s.D.Set("open_mode", s.Res.OpenMode)

	if s.Res.PdbName != nil {
		s.D.Set("pdb_name", *s.Res.PdbName)
	}

	pdbNodeLevelDetails := []interface{}{}
	for _, item := range s.Res.PdbNodeLevelDetails {
		pdbNodeLevelDetails = append(pdbNodeLevelDetails, PluggableDatabaseNodeLevelDetailsToMap(item))
	}
	s.D.Set("pdb_node_level_details", pdbNodeLevelDetails)

	if s.Res.PluggableDatabaseManagementConfig != nil {
		s.D.Set("pluggable_database_management_config", []interface{}{PluggableDatabaseManagementConfigToMap(s.Res.PluggableDatabaseManagementConfig)})
	} else {
		s.D.Set("pluggable_database_management_config", nil)
	}

	if s.Res.RefreshableCloneConfig != nil {
		s.D.Set("refreshable_clone_config", []interface{}{PluggableDatabaseRefreshableCloneConfigToMap(s.Res.RefreshableCloneConfig)})
	} else {
		s.D.Set("refreshable_clone_config", nil)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	return nil
}
