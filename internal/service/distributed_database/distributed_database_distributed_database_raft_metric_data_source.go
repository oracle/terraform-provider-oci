// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package distributed_database

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_distributed_database "github.com/oracle/oci-go-sdk/v65/distributeddatabase"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DistributedDatabaseDistributedDatabaseRaftMetricDataSource() *schema.Resource {
	return &schema.Resource{
		ReadContext: readSingularDistributedDatabaseDistributedDatabaseRaftMetricWithContext,
		Schema: map[string]*schema.Schema{
			"distributed_database_id": {
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

func readSingularDistributedDatabaseDistributedDatabaseRaftMetricWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &DistributedDatabaseDistributedDatabaseRaftMetricDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DistributedDbServiceClient()

	return tfresource.HandleDiagError(m, tfresource.ReadResourceWithContext(ctx, sync))
}

type DistributedDatabaseDistributedDatabaseRaftMetricDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_distributed_database.DistributedDbServiceClient
	Res    *oci_distributed_database.GetDistributedDatabaseRaftMetricResponse
}

func (s *DistributedDatabaseDistributedDatabaseRaftMetricDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DistributedDatabaseDistributedDatabaseRaftMetricDataSourceCrud) GetWithContext(ctx context.Context) error {
	request := oci_distributed_database.GetDistributedDatabaseRaftMetricRequest{}

	if distributedDatabaseId, ok := s.D.GetOkExists("distributed_database_id"); ok {
		tmp := distributedDatabaseId.(string)
		request.DistributedDatabaseId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "distributed_database")

	response, err := s.Client.GetDistributedDatabaseRaftMetric(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *DistributedDatabaseDistributedDatabaseRaftMetricDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DistributedDatabaseDistributedDatabaseRaftMetricDataSource-", DistributedDatabaseDistributedDatabaseRaftMetricDataSource(), s.D))

	s.D.Set("config_tasks", s.Res.ConfigTasks)

	s.D.Set("raft_metrics", s.Res.RaftMetrics)

	return nil
}
