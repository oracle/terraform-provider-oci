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

func DatabaseManagementCloudDbHomeDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["cloud_db_home_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(DatabaseManagementCloudDbHomeResource(), fieldMap, readSingularDatabaseManagementCloudDbHome)
}

func readSingularDatabaseManagementCloudDbHome(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseManagementCloudDbHomeDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DbManagementClient()

	return tfresource.ReadResource(sync)
}

type DatabaseManagementCloudDbHomeDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_database_management.DbManagementClient
	Res    *oci_database_management.GetCloudDbHomeResponse
}

func (s *DatabaseManagementCloudDbHomeDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DatabaseManagementCloudDbHomeDataSourceCrud) Get() error {
	request := oci_database_management.GetCloudDbHomeRequest{}

	if cloudDbHomeId, ok := s.D.GetOkExists("cloud_db_home_id"); ok {
		tmp := cloudDbHomeId.(string)
		request.CloudDbHomeId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "database_management")

	response, err := s.Client.GetCloudDbHome(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *DatabaseManagementCloudDbHomeDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	s.D.Set("additional_details", s.Res.AdditionalDetails)

	if s.Res.CloudDbSystemId != nil {
		s.D.Set("cloud_db_system_id", *s.Res.CloudDbSystemId)
	}

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.ComponentName != nil {
		s.D.Set("component_name", *s.Res.ComponentName)
	}

	if s.Res.DbaasId != nil {
		s.D.Set("dbaas_id", *s.Res.DbaasId)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.HomeDirectory != nil {
		s.D.Set("home_directory", *s.Res.HomeDirectory)
	}

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
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

	return nil
}
