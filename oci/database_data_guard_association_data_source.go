// Copyright (c) 2017, 2019, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"context"

	"github.com/hashicorp/terraform/helper/schema"
	oci_database "github.com/oracle/oci-go-sdk/database"
)

func DatabaseDataGuardAssociationDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularDatabaseDataGuardAssociation,
		Schema: map[string]*schema.Schema{
			"data_guard_association_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"database_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			// Computed
			"apply_lag": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"apply_rate": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"lifecycle_details": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"peer_data_guard_association_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"peer_database_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"peer_db_home_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"peer_db_system_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"peer_role": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"protection_mode": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"role": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_created": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"transport_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func readSingularDatabaseDataGuardAssociation(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseDataGuardAssociationDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).databaseClient

	return ReadResource(sync)
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

	request.RequestMetadata.RetryPolicy = getRetryPolicy(false, "database")

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
