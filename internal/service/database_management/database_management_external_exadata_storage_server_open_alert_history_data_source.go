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

func DatabaseManagementExternalExadataStorageServerOpenAlertHistoryDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularDatabaseManagementExternalExadataStorageServerOpenAlertHistory,
		Schema: map[string]*schema.Schema{
			"external_exadata_storage_server_id": {
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

func readSingularDatabaseManagementExternalExadataStorageServerOpenAlertHistory(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseManagementExternalExadataStorageServerOpenAlertHistoryDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DbManagementClient()

	return tfresource.ReadResource(sync)
}

type DatabaseManagementExternalExadataStorageServerOpenAlertHistoryDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_database_management.DbManagementClient
	Res    *oci_database_management.GetOpenAlertHistoryResponse
}

func (s *DatabaseManagementExternalExadataStorageServerOpenAlertHistoryDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DatabaseManagementExternalExadataStorageServerOpenAlertHistoryDataSourceCrud) Get() error {
	request := oci_database_management.GetOpenAlertHistoryRequest{}

	if externalExadataStorageServerId, ok := s.D.GetOkExists("external_exadata_storage_server_id"); ok {
		tmp := externalExadataStorageServerId.(string)
		request.ExternalExadataStorageServerId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "database_management")

	response, err := s.Client.GetOpenAlertHistory(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *DatabaseManagementExternalExadataStorageServerOpenAlertHistoryDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DatabaseManagementExternalExadataStorageServerOpenAlertHistoryDataSource-", DatabaseManagementExternalExadataStorageServerOpenAlertHistoryDataSource(), s.D))

	alerts := []interface{}{}
	for _, item := range s.Res.Alerts {
		alerts = append(alerts, OpenAlertSummaryToMap(item))
	}
	s.D.Set("alerts", alerts)

	return nil
}

func OpenAlertSummaryToMap(obj oci_database_management.OpenAlertSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Message != nil {
		result["message"] = string(*obj.Message)
	}

	result["severity"] = string(obj.Severity)

	if obj.TimeStartAt != nil {
		result["time_start_at"] = obj.TimeStartAt.String()
	}

	result["type"] = string(obj.Type)

	return result
}
