// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package database

import (
	"context"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_database "github.com/oracle/oci-go-sdk/v58/database"
)

func DatabaseDataGuardAssociationDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["data_guard_association_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	fieldMap["database_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(DatabaseDataGuardAssociationResource(), fieldMap, readSingularDatabaseDataGuardAssociation)
}

func readSingularDatabaseDataGuardAssociation(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseDataGuardAssociationDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()

	return tfresource.ReadResource(sync)
}

type DatabaseDataGuardAssociationDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_database.DatabaseClient
	Res    *oci_database.GetDataGuardAssociationResponse
}

func (s *DatabaseDataGuardAssociationDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DatabaseDataGuardAssociationDataSourceCrud) Get() error {
	request := oci_database.GetDataGuardAssociationRequest{}

	if dataGuardAssociationId, ok := s.D.GetOkExists("data_guard_association_id"); ok {
		tmp := dataGuardAssociationId.(string)
		request.DataGuardAssociationId = &tmp
	}

	if databaseId, ok := s.D.GetOkExists("database_id"); ok {
		tmp := databaseId.(string)
		request.DatabaseId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "database")

	response, err := s.Client.GetDataGuardAssociation(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *DatabaseDataGuardAssociationDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.ApplyLag != nil {
		s.D.Set("apply_lag", *s.Res.ApplyLag)
	}

	if s.Res.ApplyRate != nil {
		s.D.Set("apply_rate", *s.Res.ApplyRate)
	}

	if s.Res.IsActiveDataGuardEnabled != nil {
		s.D.Set("is_active_data_guard_enabled", *s.Res.IsActiveDataGuardEnabled)
	}

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	if s.Res.PeerDataGuardAssociationId != nil {
		s.D.Set("peer_data_guard_association_id", *s.Res.PeerDataGuardAssociationId)
	}

	if s.Res.PeerDatabaseId != nil {
		s.D.Set("peer_database_id", *s.Res.PeerDatabaseId)
	}

	if s.Res.PeerDbHomeId != nil {
		s.D.Set("peer_db_home_id", *s.Res.PeerDbHomeId)
	}

	if s.Res.PeerDbSystemId != nil {
		s.D.Set("peer_db_system_id", *s.Res.PeerDbSystemId)
	}

	s.D.Set("peer_role", s.Res.PeerRole)

	s.D.Set("protection_mode", s.Res.ProtectionMode)

	s.D.Set("role", s.Res.Role)

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	s.D.Set("transport_type", s.Res.TransportType)

	return nil
}
