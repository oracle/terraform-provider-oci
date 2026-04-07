// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package distributed_database

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_distributed_database "github.com/oracle/oci-go-sdk/v65/distributeddatabase"

	// TF_CODE_GEN: TERSI-4920-TOP-17 generated raft metric data sources reference client.OracleClients but omit the internal/client import.
	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DistributedDatabaseDistributedAutonomousDatabaseRaftMetricDataSource() *schema.Resource {
	return &schema.Resource{
		ReadContext: readSingularDistributedDatabaseDistributedAutonomousDatabaseRaftMetricWithContext,
		Schema: map[string]*schema.Schema{
			"distributed_autonomous_database_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			// Computed
			"config_tasks": {
				Type:     schema.TypeMap,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"raft_metrics": {
				Type:     schema.TypeMap,
				Computed: true,
				Elem:     schema.TypeString,
			},
		},
	}
}

func readSingularDistributedDatabaseDistributedAutonomousDatabaseRaftMetricWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &DistributedDatabaseDistributedAutonomousDatabaseRaftMetricDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DistributedAutonomousDbServiceClient()

	return tfresource.HandleDiagError(m, tfresource.ReadResourceWithContext(ctx, sync))
}

type DistributedDatabaseDistributedAutonomousDatabaseRaftMetricDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_distributed_database.DistributedAutonomousDbServiceClient
	Res    *oci_distributed_database.GetDistributedAutonomousDatabaseRaftMetricResponse
}

func (s *DistributedDatabaseDistributedAutonomousDatabaseRaftMetricDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DistributedDatabaseDistributedAutonomousDatabaseRaftMetricDataSourceCrud) GetWithContext(ctx context.Context) error {
	request := oci_distributed_database.GetDistributedAutonomousDatabaseRaftMetricRequest{}

	if distributedAutonomousDatabaseId, ok := s.D.GetOkExists("distributed_autonomous_database_id"); ok {
		tmp := distributedAutonomousDatabaseId.(string)
		request.DistributedAutonomousDatabaseId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "distributed_database")

	response, err := s.Client.GetDistributedAutonomousDatabaseRaftMetric(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *DistributedDatabaseDistributedAutonomousDatabaseRaftMetricDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DistributedDatabaseDistributedAutonomousDatabaseRaftMetricDataSource-", DistributedDatabaseDistributedAutonomousDatabaseRaftMetricDataSource(), s.D))

	s.D.Set("config_tasks", s.Res.ConfigTasks)

	s.D.Set("raft_metrics", s.Res.RaftMetrics)

	return nil
}
