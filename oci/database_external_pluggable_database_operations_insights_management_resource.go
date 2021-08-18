// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	oci_database "github.com/oracle/oci-go-sdk/v46/database"
	oci_work_requests "github.com/oracle/oci-go-sdk/v46/workrequests"
)

func init() {
	RegisterResource("oci_database_external_pluggable_database_operations_insights_management", DatabaseExternalPluggableDatabaseOperationsInsightsManagementResource())
}

func DatabaseExternalPluggableDatabaseOperationsInsightsManagementResource() *schema.Resource {
	return &schema.Resource{
		Timeouts: DefaultTimeout,
		Create:   createDatabaseExternalPluggableDatabaseOperationsInsightsManagement,
		Update:   updateDatabaseExternalPluggableDatabaseOperationsInsightsManagement,
		Read:     readDatabaseExternalPluggableDatabaseOperationsInsightsManagement,
		Delete:   deleteDatabaseExternalPluggableDatabaseOperationsInsightsManagement,
		Schema: map[string]*schema.Schema{
			// Required
			"external_database_connector_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"external_pluggable_database_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			"enable_operations_insights": {
				Type:     schema.TypeBool,
				Required: true,
			},

			// Optional

			// Computed
		},
	}
}

func createDatabaseExternalPluggableDatabaseOperationsInsightsManagement(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseExternalPluggableDatabaseOperationsInsightsManagementResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).databaseClient()
	sync.WorkRequestClient = m.(*OracleClients).workRequestClient
	sync.Res = &DatabaseExternalPluggableDatabaseOperationsInsightsResponse{}

	return CreateResource(d, sync)
}

func updateDatabaseExternalPluggableDatabaseOperationsInsightsManagement(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseExternalPluggableDatabaseOperationsInsightsManagementResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).databaseClient()
	sync.WorkRequestClient = m.(*OracleClients).workRequestClient
	sync.Res = &DatabaseExternalPluggableDatabaseOperationsInsightsResponse{}
	return UpdateResource(d, sync)
}

func readDatabaseExternalPluggableDatabaseOperationsInsightsManagement(d *schema.ResourceData, m interface{}) error {
	return nil
}

func deleteDatabaseExternalPluggableDatabaseOperationsInsightsManagement(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseExternalPluggableDatabaseOperationsInsightsManagementResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).databaseClient()
	sync.WorkRequestClient = m.(*OracleClients).workRequestClient
	sync.Res = &DatabaseExternalPluggableDatabaseOperationsInsightsResponse{}
	sync.DisableNotFoundRetries = true

	return DeleteResource(d, sync)
}

type DatabaseExternalPluggableDatabaseOperationsInsightsResponse struct {
	enableResponse  *oci_database.EnableExternalPluggableDatabaseOperationsInsightsResponse
	disableResponse *oci_database.DisableExternalPluggableDatabaseOperationsInsightsResponse
}

type DatabaseExternalPluggableDatabaseOperationsInsightsManagementResourceCrud struct {
	BaseCrud
	Client                 *oci_database.DatabaseClient
	Res                    *DatabaseExternalPluggableDatabaseOperationsInsightsResponse
	DisableNotFoundRetries bool
	WorkRequestClient      *oci_work_requests.WorkRequestClient
}

func (s *DatabaseExternalPluggableDatabaseOperationsInsightsManagementResourceCrud) ID() string {
	return GenerateDataSourceHashID("DatabaseExternalPluggableDatabaseOperationsInsightsManagementResource-", DatabaseExternalPluggableDatabaseOperationsInsightsManagementResource(), s.D)
}

func (s *DatabaseExternalPluggableDatabaseOperationsInsightsManagementResourceCrud) Create() error {
	var operation bool
	if enableOperationsInsights, ok := s.D.GetOkExists("enable_operations_insights"); ok {
		operation = enableOperationsInsights.(bool)
	}

	if operation {
		// enable operations insights
		request := oci_database.EnableExternalPluggableDatabaseOperationsInsightsRequest{}

		if externalDatabaseConnectorId, ok := s.D.GetOkExists("external_database_connector_id"); ok {
			tmp := externalDatabaseConnectorId.(string)
			request.ExternalDatabaseConnectorId = &tmp
		}

		if externalPluggableDatabaseId, ok := s.D.GetOkExists("external_pluggable_database_id"); ok {
			tmp := externalPluggableDatabaseId.(string)
			request.ExternalPluggableDatabaseId = &tmp
		}

		request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "database")

		response, err := s.Client.EnableExternalPluggableDatabaseOperationsInsights(context.Background(), request)
		if err != nil {
			return err
		}

		workId := response.OpcWorkRequestId
		if workId != nil {
			_, err = WaitForWorkRequestWithErrorHandling(s.WorkRequestClient, workId, "externalPluggableDatabase", oci_work_requests.WorkRequestResourceActionTypeUpdated, s.D.Timeout(schema.TimeoutCreate), s.DisableNotFoundRetries)
			if err != nil {
				return err
			}
		}
		s.Res.enableResponse = &response
		return nil
	}

	// disable operations insights
	request := oci_database.DisableExternalPluggableDatabaseOperationsInsightsRequest{}

	if externalPluggableDatabaseId, ok := s.D.GetOkExists("external_pluggable_database_id"); ok {
		tmp := externalPluggableDatabaseId.(string)
		request.ExternalPluggableDatabaseId = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "database")

	response, err := s.Client.DisableExternalPluggableDatabaseOperationsInsights(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	if workId != nil {
		_, err = WaitForWorkRequestWithErrorHandling(s.WorkRequestClient, workId, "externalPluggableDatabase", oci_work_requests.WorkRequestResourceActionTypeUpdated, s.D.Timeout(schema.TimeoutCreate), s.DisableNotFoundRetries)
		if err != nil {
			return err
		}
	}
	s.Res.disableResponse = &response
	return nil
}

func (s *DatabaseExternalPluggableDatabaseOperationsInsightsManagementResourceCrud) Update() error {
	var operation bool
	if enableOperationsInsights, ok := s.D.GetOkExists("enable_operations_insights"); ok {
		operation = enableOperationsInsights.(bool)
	}

	if operation {
		// enable operations insights
		request := oci_database.EnableExternalPluggableDatabaseOperationsInsightsRequest{}

		if externalDatabaseConnectorId, ok := s.D.GetOkExists("external_database_connector_id"); ok {
			tmp := externalDatabaseConnectorId.(string)
			request.ExternalDatabaseConnectorId = &tmp
		}

		if externalPluggableDatabaseId, ok := s.D.GetOkExists("external_pluggable_database_id"); ok {
			tmp := externalPluggableDatabaseId.(string)
			request.ExternalPluggableDatabaseId = &tmp
		}

		request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "database")

		response, err := s.Client.EnableExternalPluggableDatabaseOperationsInsights(context.Background(), request)
		if err != nil {
			return err
		}

		workId := response.OpcWorkRequestId
		if workId != nil {
			_, err = WaitForWorkRequestWithErrorHandling(s.WorkRequestClient, workId, "externalPluggableDatabase", oci_work_requests.WorkRequestResourceActionTypeUpdated, s.D.Timeout(schema.TimeoutCreate), s.DisableNotFoundRetries)
			if err != nil {
				return err
			}
		}
		s.Res.enableResponse = &response
		return nil
	}

	// disable operations insights
	request := oci_database.DisableExternalPluggableDatabaseOperationsInsightsRequest{}

	if externalPluggableDatabaseId, ok := s.D.GetOkExists("external_pluggable_database_id"); ok {
		tmp := externalPluggableDatabaseId.(string)
		request.ExternalPluggableDatabaseId = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "database")

	response, err := s.Client.DisableExternalPluggableDatabaseOperationsInsights(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	if workId != nil {
		_, err = WaitForWorkRequestWithErrorHandling(s.WorkRequestClient, workId, "externalPluggableDatabase", oci_work_requests.WorkRequestResourceActionTypeUpdated, s.D.Timeout(schema.TimeoutCreate), s.DisableNotFoundRetries)
		if err != nil {
			return err
		}
	}
	s.Res.disableResponse = &response
	return nil
}

func (s *DatabaseExternalPluggableDatabaseOperationsInsightsManagementResourceCrud) Delete() error {
	var operation bool
	if enableOpsi, ok := s.D.GetOkExists("enable_operations_insights"); ok {
		operation = enableOpsi.(bool)
	}

	if !operation {
		return nil
	}
	// Disable Operations Insights
	request := oci_database.DisableExternalPluggableDatabaseOperationsInsightsRequest{}

	if externalPluggableDatabaseId, ok := s.D.GetOkExists("external_pluggable_database_id"); ok {
		tmp := externalPluggableDatabaseId.(string)
		request.ExternalPluggableDatabaseId = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "database")

	response, err := s.Client.DisableExternalPluggableDatabaseOperationsInsights(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	if workId != nil {
		_, err = WaitForWorkRequestWithErrorHandling(s.WorkRequestClient, workId, "externalPluggableDatabase", oci_work_requests.WorkRequestResourceActionTypeUpdated, s.D.Timeout(schema.TimeoutCreate), s.DisableNotFoundRetries)
		if err != nil {
			return err
		}
	}
	s.Res.disableResponse = &response
	return nil
}

func (s *DatabaseExternalPluggableDatabaseOperationsInsightsManagementResourceCrud) SetData() error {
	return nil
}
