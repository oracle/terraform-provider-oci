// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package redis

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_redis "github.com/oracle/oci-go-sdk/v65/redis"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func RedisOciCacheBackupDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["oci_cache_backup_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchemaWithContext(RedisOciCacheBackupResource(), fieldMap, readSingularRedisOciCacheBackupWithContext)
}

func readSingularRedisOciCacheBackupWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &RedisOciCacheBackupDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).OciCacheBackupClient()

	return tfresource.HandleDiagError(m, tfresource.ReadResourceWithContext(ctx, sync))
}

type RedisOciCacheBackupDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_redis.OciCacheBackupClient
	Res    *oci_redis.GetOciCacheBackupResponse
}

func (s *RedisOciCacheBackupDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *RedisOciCacheBackupDataSourceCrud) GetWithContext(ctx context.Context) error {
	request := oci_redis.GetOciCacheBackupRequest{}

	if ociCacheBackupId, ok := s.D.GetOkExists("oci_cache_backup_id"); ok {
		tmp := ociCacheBackupId.(string)
		request.OciCacheBackupId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "redis")

	response, err := s.Client.GetOciCacheBackup(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *RedisOciCacheBackupDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.BackupSizeInGBs != nil {
		s.D.Set("backup_size_in_gbs", *s.Res.BackupSizeInGBs)
	}

	s.D.Set("backup_source", s.Res.BackupSource)

	s.D.Set("backup_type", s.Res.BackupType)

	if s.Res.ClusterMemoryInGBs != nil {
		s.D.Set("cluster_memory_in_gbs", *s.Res.ClusterMemoryInGBs)
	}

	s.D.Set("cluster_mode", s.Res.ClusterMode)

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.Description != nil {
		s.D.Set("description", *s.Res.Description)
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.RetentionPeriodInDays != nil {
		s.D.Set("retention_period_in_days", *s.Res.RetentionPeriodInDays)
	}

	if s.Res.ShardCount != nil {
		s.D.Set("shard_count", *s.Res.ShardCount)
	}

	s.D.Set("software_version", s.Res.SoftwareVersion)

	if s.Res.SourceClusterId != nil {
		s.D.Set("source_cluster_id", *s.Res.SourceClusterId)
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
