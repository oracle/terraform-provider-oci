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

func RedisOciCacheBackupsDataSource() *schema.Resource {
	return &schema.Resource{
		ReadContext: readRedisOciCacheBackupsWithContext,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"oci_cache_backup_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"source_cluster_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"oci_cache_backup_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem:     tfresource.GetDataSourceItemSchema(RedisOciCacheBackupResource()),
						},
					},
				},
			},
		},
	}
}

func readRedisOciCacheBackupsWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &RedisOciCacheBackupsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).OciCacheBackupClient()

	return tfresource.HandleDiagError(m, tfresource.ReadResourceWithContext(ctx, sync))
}

type RedisOciCacheBackupsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_redis.OciCacheBackupClient
	Res    *oci_redis.ListOciCacheBackupsResponse
}

func (s *RedisOciCacheBackupsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *RedisOciCacheBackupsDataSourceCrud) GetWithContext(ctx context.Context) error {
	request := oci_redis.ListOciCacheBackupsRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if ociCacheBackupId, ok := s.D.GetOkExists("id"); ok {
		tmp := ociCacheBackupId.(string)
		request.OciCacheBackupId = &tmp
	}

	if sourceClusterId, ok := s.D.GetOkExists("source_cluster_id"); ok {
		tmp := sourceClusterId.(string)
		request.SourceClusterId = &tmp
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_redis.OciCacheBackupLifecycleStateEnum(state.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "redis")

	response, err := s.Client.ListOciCacheBackups(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListOciCacheBackups(ctx, request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *RedisOciCacheBackupsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("RedisOciCacheBackupsDataSource-", RedisOciCacheBackupsDataSource(), s.D))
	resources := []map[string]interface{}{}
	ociCacheBackup := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, OciCacheBackupSummaryToMap(item))
	}
	ociCacheBackup["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, RedisOciCacheBackupsDataSource().Schema["oci_cache_backup_collection"].Elem.(*schema.Resource).Schema)
		ociCacheBackup["items"] = items
	}

	resources = append(resources, ociCacheBackup)
	if err := s.D.Set("oci_cache_backup_collection", resources); err != nil {
		return err
	}

	return nil
}
