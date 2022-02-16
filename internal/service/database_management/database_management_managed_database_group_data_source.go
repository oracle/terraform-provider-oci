// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package database_management

import (
	"context"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_database_management "github.com/oracle/oci-go-sdk/v58/databasemanagement"
)

func DatabaseManagementManagedDatabaseGroupDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["managed_database_group_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(DatabaseManagementManagedDatabaseGroupResource(), fieldMap, readSingularDatabaseManagementManagedDatabaseGroup)
}

func readSingularDatabaseManagementManagedDatabaseGroup(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseManagementManagedDatabaseGroupDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DbManagementClient()

	return tfresource.ReadResource(sync)
}

type DatabaseManagementManagedDatabaseGroupDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_database_management.DbManagementClient
	Res    *oci_database_management.GetManagedDatabaseGroupResponse
}

func (s *DatabaseManagementManagedDatabaseGroupDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DatabaseManagementManagedDatabaseGroupDataSourceCrud) Get() error {
	request := oci_database_management.GetManagedDatabaseGroupRequest{}

	if managedDatabaseGroupId, ok := s.D.GetOkExists("managed_database_group_id"); ok {
		tmp := managedDatabaseGroupId.(string)
		request.ManagedDatabaseGroupId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "database_management")

	response, err := s.Client.GetManagedDatabaseGroup(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *DatabaseManagementManagedDatabaseGroupDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.Description != nil {
		s.D.Set("description", *s.Res.Description)
	}

	managedDatabases := []interface{}{}
	for _, item := range s.Res.ManagedDatabases {
		managedDatabases = append(managedDatabases, ChildDatabaseToMap(item))
	}
	s.D.Set("managed_databases", managedDatabases)

	if s.Res.Name != nil {
		s.D.Set("name", *s.Res.Name)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	return nil
}
