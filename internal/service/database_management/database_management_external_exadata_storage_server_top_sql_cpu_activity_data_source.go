// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package database_management

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_database_management "github.com/oracle/oci-go-sdk/v65/databasemanagement"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DatabaseManagementExternalExadataStorageServerTopSqlCpuActivityDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularDatabaseManagementExternalExadataStorageServerTopSqlCpuActivity,
		Schema: map[string]*schema.Schema{
			"external_exadata_storage_server_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			// Computed
			"activity": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"cpu_activity": {
							Type:     schema.TypeFloat,
							Computed: true,
						},
						"database_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"sql_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func readSingularDatabaseManagementExternalExadataStorageServerTopSqlCpuActivity(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseManagementExternalExadataStorageServerTopSqlCpuActivityDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DbManagementClient()

	return tfresource.ReadResource(sync)
}

type DatabaseManagementExternalExadataStorageServerTopSqlCpuActivityDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_database_management.DbManagementClient
	Res    *oci_database_management.GetTopSqlCpuActivityResponse
}

func (s *DatabaseManagementExternalExadataStorageServerTopSqlCpuActivityDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DatabaseManagementExternalExadataStorageServerTopSqlCpuActivityDataSourceCrud) Get() error {
	request := oci_database_management.GetTopSqlCpuActivityRequest{}

	if externalExadataStorageServerId, ok := s.D.GetOkExists("external_exadata_storage_server_id"); ok {
		tmp := externalExadataStorageServerId.(string)
		request.ExternalExadataStorageServerId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "database_management")

	response, err := s.Client.GetTopSqlCpuActivity(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *DatabaseManagementExternalExadataStorageServerTopSqlCpuActivityDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DatabaseManagementExternalExadataStorageServerTopSqlCpuActivityDataSource-", DatabaseManagementExternalExadataStorageServerTopSqlCpuActivityDataSource(), s.D))

	activity := []interface{}{}
	for _, item := range s.Res.Activity {
		activity = append(activity, SqlCpuActivityToMap(item))
	}
	s.D.Set("activity", activity)

	return nil
}

func SqlCpuActivityToMap(obj oci_database_management.SqlCpuActivity) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.CpuActivity != nil {
		result["cpu_activity"] = float32(*obj.CpuActivity)
	}

	if obj.DatabaseName != nil {
		result["database_name"] = string(*obj.DatabaseName)
	}

	if obj.SqlId != nil {
		result["sql_id"] = string(*obj.SqlId)
	}

	return result
}
