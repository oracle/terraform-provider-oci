// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package database_management

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_database_management "github.com/oracle/oci-go-sdk/v65/databasemanagement"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DatabaseManagementExternalExadataInfrastructureDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["external_exadata_infrastructure_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(DatabaseManagementExternalExadataInfrastructureResource(), fieldMap, readSingularDatabaseManagementExternalExadataInfrastructure)
}

func readSingularDatabaseManagementExternalExadataInfrastructure(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseManagementExternalExadataInfrastructureDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DbManagementClient()

	return tfresource.ReadResource(sync)
}

type DatabaseManagementExternalExadataInfrastructureDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_database_management.DbManagementClient
	Res    *oci_database_management.GetExternalExadataInfrastructureResponse
}

func (s *DatabaseManagementExternalExadataInfrastructureDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DatabaseManagementExternalExadataInfrastructureDataSourceCrud) Get() error {
	request := oci_database_management.GetExternalExadataInfrastructureRequest{}

	if externalExadataInfrastructureId, ok := s.D.GetOkExists("external_exadata_infrastructure_id"); ok {
		tmp := externalExadataInfrastructureId.(string)
		request.ExternalExadataInfrastructureId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "database_management")

	response, err := s.Client.GetExternalExadataInfrastructure(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *DatabaseManagementExternalExadataInfrastructureDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	s.D.Set("additional_details", s.Res.AdditionalDetails)

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	s.D.Set("database_compartments", s.Res.DatabaseCompartments)

	databaseSystems := []interface{}{}
	for _, item := range s.Res.DatabaseSystems {
		databaseSystems = append(databaseSystems, ExternalExadataDatabaseSystemSummaryToMap(item))
	}
	s.D.Set("database_systems", databaseSystems)

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	if s.Res.InternalId != nil {
		s.D.Set("internal_id", *s.Res.InternalId)
	}

	s.D.Set("license_model", s.Res.LicenseModel)

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	s.D.Set("rack_size", s.Res.RackSize)

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.Status != nil {
		s.D.Set("status", *s.Res.Status)
	}

	if s.Res.StorageGrid != nil {
		s.D.Set("storage_grid", []interface{}{ExternalExadataStorageGridSummaryToMap(s.Res.StorageGrid)})
	} else {
		s.D.Set("storage_grid", nil)
	}

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
