// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package database_management

import (
	"context"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_database_management "github.com/oracle/oci-go-sdk/v65/databasemanagement"
	"github.com/oracle/terraform-provider-oci/internal/client"

	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DatabaseManagementManagedMySqlDatabaseGeneralReplicationInformationDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularDatabaseManagementManagedMySqlDatabaseGeneralReplicationInformation,
		Schema: map[string]*schema.Schema{
			"managed_my_sql_database_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			// Computed
			"apply_status_summary": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"binary_log_format": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"binary_logging": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"executed_gtid_set": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"fetch_status_summary": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"gtid_mode": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"high_availability_member_state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"host_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"inbound_replications_count": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"instance_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"is_high_availability_enabled": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"outbound_replications_count": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"port": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"read_only": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"seconds_behind_source_max": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"server_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"server_uuid": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func readSingularDatabaseManagementManagedMySqlDatabaseGeneralReplicationInformation(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseManagementManagedMySqlDatabaseGeneralReplicationInformationDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ManagedMySqlDatabasesClient()

	return tfresource.ReadResource(sync)
}

type DatabaseManagementManagedMySqlDatabaseGeneralReplicationInformationDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_database_management.ManagedMySqlDatabasesClient
	Res    *oci_database_management.GetGeneralReplicationInformationResponse
}

func (s *DatabaseManagementManagedMySqlDatabaseGeneralReplicationInformationDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DatabaseManagementManagedMySqlDatabaseGeneralReplicationInformationDataSourceCrud) Get() error {
	request := oci_database_management.GetGeneralReplicationInformationRequest{}

	if managedMySqlDatabaseId, ok := s.D.GetOkExists("managed_my_sql_database_id"); ok {
		tmp := managedMySqlDatabaseId.(string)
		request.ManagedMySqlDatabaseId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "database_management")

	response, err := s.Client.GetGeneralReplicationInformation(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *DatabaseManagementManagedMySqlDatabaseGeneralReplicationInformationDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DatabaseManagementManagedMySqlDatabaseGeneralReplicationInformationDataSource-", DatabaseManagementManagedMySqlDatabaseGeneralReplicationInformationDataSource(), s.D))

	if s.Res.ApplyStatusSummary != nil {
		s.D.Set("apply_status_summary", *s.Res.ApplyStatusSummary)
	}

	if s.Res.BinaryLogFormat != nil {
		s.D.Set("binary_log_format", *s.Res.BinaryLogFormat)
	}

	if s.Res.BinaryLogging != nil {
		s.D.Set("binary_logging", *s.Res.BinaryLogging)
	}

	if s.Res.ExecutedGtidSet != nil {
		s.D.Set("executed_gtid_set", *s.Res.ExecutedGtidSet)
	}

	if s.Res.FetchStatusSummary != nil {
		s.D.Set("fetch_status_summary", *s.Res.FetchStatusSummary)
	}

	if s.Res.GtidMode != nil {
		s.D.Set("gtid_mode", *s.Res.GtidMode)
	}

	if s.Res.HighAvailabilityMemberState != nil {
		s.D.Set("high_availability_member_state", *s.Res.HighAvailabilityMemberState)
	}

	if s.Res.HostName != nil {
		s.D.Set("host_name", *s.Res.HostName)
	}

	if s.Res.InboundReplicationsCount != nil {
		s.D.Set("inbound_replications_count", *s.Res.InboundReplicationsCount)
	}

	if s.Res.InstanceType != nil {
		s.D.Set("instance_type", *s.Res.InstanceType)
	}

	if s.Res.IsHighAvailabilityEnabled != nil {
		s.D.Set("is_high_availability_enabled", *s.Res.IsHighAvailabilityEnabled)
	}

	if s.Res.OutboundReplicationsCount != nil {
		s.D.Set("outbound_replications_count", *s.Res.OutboundReplicationsCount)
	}

	if s.Res.Port != nil {
		s.D.Set("port", *s.Res.Port)
	}

	s.D.Set("read_only", s.Res.ReadOnly)

	if s.Res.SecondsBehindSourceMax != nil {
		s.D.Set("seconds_behind_source_max", strconv.FormatInt(*s.Res.SecondsBehindSourceMax, 10))
	}

	if s.Res.ServerId != nil {
		s.D.Set("server_id", strconv.FormatInt(*s.Res.ServerId, 10))
	}

	if s.Res.ServerUuid != nil {
		s.D.Set("server_uuid", *s.Res.ServerUuid)
	}

	return nil
}
