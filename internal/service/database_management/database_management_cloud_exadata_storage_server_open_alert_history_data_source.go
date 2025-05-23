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

func DatabaseManagementCloudExadataStorageServerOpenAlertHistoryDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularDatabaseManagementCloudExadataStorageServerOpenAlertHistory,
		Schema: map[string]*schema.Schema{
			"cloud_exadata_storage_server_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			// Computed
			"alerts": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"message": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"severity": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"time_start_at": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"type": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func readSingularDatabaseManagementCloudExadataStorageServerOpenAlertHistory(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseManagementCloudExadataStorageServerOpenAlertHistoryDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DbManagementClient()

	return tfresource.ReadResource(sync)
}

type DatabaseManagementCloudExadataStorageServerOpenAlertHistoryDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_database_management.DbManagementClient
	Res    *oci_database_management.GetCloudOpenAlertHistoryResponse
}

func (s *DatabaseManagementCloudExadataStorageServerOpenAlertHistoryDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DatabaseManagementCloudExadataStorageServerOpenAlertHistoryDataSourceCrud) Get() error {
	request := oci_database_management.GetCloudOpenAlertHistoryRequest{}

	if cloudExadataStorageServerId, ok := s.D.GetOkExists("cloud_exadata_storage_server_id"); ok {
		tmp := cloudExadataStorageServerId.(string)
		request.CloudExadataStorageServerId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "database_management")

	response, err := s.Client.GetCloudOpenAlertHistory(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *DatabaseManagementCloudExadataStorageServerOpenAlertHistoryDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DatabaseManagementCloudExadataStorageServerOpenAlertHistoryDataSource-", DatabaseManagementCloudExadataStorageServerOpenAlertHistoryDataSource(), s.D))

	alerts := []interface{}{}
	for _, item := range s.Res.Alerts {
		alerts = append(alerts, OpenAlertSummaryToMap(item))
	}
	s.D.Set("alerts", alerts)

	return nil
}

//func OpenAlertSummaryToMap(obj oci_database_management.OpenAlertSummary) map[string]interface{} {
//	result := map[string]interface{}{}
//
//	if obj.Message != nil {
//		result["message"] = string(*obj.Message)
//	}
//
//	result["severity"] = string(obj.Severity)
//
//	if obj.TimeStartAt != nil {
//		result["time_start_at"] = obj.TimeStartAt.String()
//	}
//
//	result["type"] = string(obj.Type)
//
//	return result
//}
