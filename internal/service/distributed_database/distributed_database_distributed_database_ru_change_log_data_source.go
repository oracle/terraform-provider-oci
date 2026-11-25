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

func DistributedDatabaseDistributedDatabaseRuChangeLogDataSource() *schema.Resource {
	return &schema.Resource{
		ReadContext: readSingularDistributedDatabaseDistributedDatabaseRuChangeLogWithContext,
		Schema: map[string]*schema.Schema{
			"distributed_database_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"source_shard_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			// Computed
			"ru_change_log": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"ru_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"source_shard_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"status": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"target_shard_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"time_created": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"time_updated": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"triggered_by": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func readSingularDistributedDatabaseDistributedDatabaseRuChangeLogWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &DistributedDatabaseDistributedDatabaseRuChangeLogDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DistributedDbServiceClient()

	return tfresource.HandleDiagError(m, tfresource.ReadResourceWithContext(ctx, sync))
}

type DistributedDatabaseDistributedDatabaseRuChangeLogDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_distributed_database.DistributedDbServiceClient
	Res    *oci_distributed_database.GetDistributedDatabaseRuChangeLogResponse
}

func (s *DistributedDatabaseDistributedDatabaseRuChangeLogDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DistributedDatabaseDistributedDatabaseRuChangeLogDataSourceCrud) GetWithContext(ctx context.Context) error {
	request := oci_distributed_database.GetDistributedDatabaseRuChangeLogRequest{}

	if distributedDatabaseId, ok := s.D.GetOkExists("distributed_database_id"); ok {
		tmp := distributedDatabaseId.(string)
		request.DistributedDatabaseId = &tmp
	}

	if sourceShardName, ok := s.D.GetOkExists("source_shard_name"); ok {
		tmp := sourceShardName.(string)
		request.SourceShardName = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "distributed_database")

	response, err := s.Client.GetDistributedDatabaseRuChangeLog(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *DistributedDatabaseDistributedDatabaseRuChangeLogDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DistributedDatabaseDistributedDatabaseRuChangeLogDataSource-", DistributedDatabaseDistributedDatabaseRuChangeLogDataSource(), s.D))

	ruChangeLog := []interface{}{}
	for _, item := range s.Res.RuChangeLog {
		ruChangeLog = append(ruChangeLog, DistributedDatabaseRuChangeLogSummaryToMap(item))
	}
	s.D.Set("ru_change_log", ruChangeLog)

	return nil
}

func DistributedDatabaseRuChangeLogSummaryToMap(obj oci_distributed_database.DistributedDatabaseRuChangeLogSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.RuId != nil {
		result["ru_id"] = string(*obj.RuId)
	}

	if obj.SourceShardName != nil {
		result["source_shard_name"] = string(*obj.SourceShardName)
	}

	result["status"] = string(obj.Status)

	if obj.TargetShardName != nil {
		result["target_shard_name"] = string(*obj.TargetShardName)
	}

	if obj.TimeCreated != nil {
		result["time_created"] = obj.TimeCreated.String()
	}

	if obj.TimeUpdated != nil {
		result["time_updated"] = obj.TimeUpdated.String()
	}

	if obj.TriggeredBy != nil {
		result["triggered_by"] = string(*obj.TriggeredBy)
	}

	return result
}
