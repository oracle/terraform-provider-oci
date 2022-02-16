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

func DatabaseAutonomousContainerDatabaseDataguardAssociationDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["autonomous_container_database_dataguard_association_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	fieldMap["autonomous_container_database_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(DatabaseAutonomousContainerDatabaseDataguardAssociationResource(), fieldMap, readSingularDatabaseAutonomousContainerDatabaseDataguardAssociation)
}

func readSingularDatabaseAutonomousContainerDatabaseDataguardAssociation(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseAutonomousContainerDatabaseDataguardAssociationDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()

	return tfresource.ReadResource(sync)
}

type DatabaseAutonomousContainerDatabaseDataguardAssociationDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_database.DatabaseClient
	Res    *oci_database.GetAutonomousContainerDatabaseDataguardAssociationResponse
}

func (s *DatabaseAutonomousContainerDatabaseDataguardAssociationDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DatabaseAutonomousContainerDatabaseDataguardAssociationDataSourceCrud) Get() error {
	request := oci_database.GetAutonomousContainerDatabaseDataguardAssociationRequest{}

	if autonomousContainerDatabaseDataguardAssociationId, ok := s.D.GetOkExists("autonomous_container_database_dataguard_association_id"); ok {
		tmp := autonomousContainerDatabaseDataguardAssociationId.(string)
		request.AutonomousContainerDatabaseDataguardAssociationId = &tmp
	}

	if autonomousContainerDatabaseId, ok := s.D.GetOkExists("autonomous_container_database_id"); ok {
		tmp := autonomousContainerDatabaseId.(string)
		request.AutonomousContainerDatabaseId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "database")

	response, err := s.Client.GetAutonomousContainerDatabaseDataguardAssociation(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *DatabaseAutonomousContainerDatabaseDataguardAssociationDataSourceCrud) SetData() error {
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

	if s.Res.IsAutomaticFailoverEnabled != nil {
		s.D.Set("is_automatic_failover_enabled", *s.Res.IsAutomaticFailoverEnabled)
	}

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	if s.Res.PeerAutonomousContainerDatabaseDataguardAssociationId != nil {
		s.D.Set("peer_autonomous_container_database_dataguard_association_id", *s.Res.PeerAutonomousContainerDatabaseDataguardAssociationId)
	}

	if s.Res.PeerAutonomousContainerDatabaseId != nil {
		s.D.Set("peer_autonomous_container_database_id", *s.Res.PeerAutonomousContainerDatabaseId)
	}

	s.D.Set("peer_lifecycle_state", s.Res.PeerLifecycleState)

	s.D.Set("peer_role", s.Res.PeerRole)

	s.D.Set("protection_mode", s.Res.ProtectionMode)

	s.D.Set("role", s.Res.Role)

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeLastRoleChanged != nil {
		s.D.Set("time_last_role_changed", s.Res.TimeLastRoleChanged.String())
	}

	if s.Res.TimeLastSynced != nil {
		s.D.Set("time_last_synced", s.Res.TimeLastSynced.String())
	}

	if s.Res.TransportLag != nil {
		s.D.Set("transport_lag", *s.Res.TransportLag)
	}

	return nil
}
