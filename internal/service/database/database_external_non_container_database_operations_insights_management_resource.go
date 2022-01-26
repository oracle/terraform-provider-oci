// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package database

import (
	"context"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	oci_database "github.com/oracle/oci-go-sdk/v56/database"
	oci_work_requests "github.com/oracle/oci-go-sdk/v56/workrequests"
)

func DatabaseExternalNonContainerDatabaseOperationsInsightsManagementResource() *schema.Resource {
	return &schema.Resource{
		Timeouts: tfresource.DefaultTimeout,
		Create:   createDatabaseExternalNonContainerDatabaseOperationsInsightsManagement,
		Update:   updateDatabaseExternalNonContainerDatabaseOperationsInsightsManagement,
		Read:     readDatabaseExternalNonContainerDatabaseOperationsInsightsManagement,
		Delete:   deleteDatabaseExternalNonContainerDatabaseOperationsInsightsManagement,
		Schema: map[string]*schema.Schema{
			// Required
			"external_database_connector_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"external_non_container_database_id": {
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

func createDatabaseExternalNonContainerDatabaseOperationsInsightsManagement(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseExternalNonContainerDatabaseOperationsInsightsManagementResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()
	sync.WorkRequestClient = m.(*client.OracleClients).WorkRequestClient
	sync.Res = &DatabaseExternalNonContainerDatabaseOperationsInsightsResponse{}

	return tfresource.CreateResource(d, sync)
}

func updateDatabaseExternalNonContainerDatabaseOperationsInsightsManagement(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseExternalNonContainerDatabaseOperationsInsightsManagementResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()
	sync.WorkRequestClient = m.(*client.OracleClients).WorkRequestClient
	sync.Res = &DatabaseExternalNonContainerDatabaseOperationsInsightsResponse{}
	return tfresource.UpdateResource(d, sync)
}

func readDatabaseExternalNonContainerDatabaseOperationsInsightsManagement(d *schema.ResourceData, m interface{}) error {
	return nil
}

func deleteDatabaseExternalNonContainerDatabaseOperationsInsightsManagement(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseExternalNonContainerDatabaseOperationsInsightsManagementResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()
	sync.WorkRequestClient = m.(*client.OracleClients).WorkRequestClient
	sync.Res = &DatabaseExternalNonContainerDatabaseOperationsInsightsResponse{}
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type DatabaseExternalNonContainerDatabaseOperationsInsightsResponse struct {
	enableResponse  *oci_database.EnableExternalNonContainerDatabaseOperationsInsightsResponse
	disableResponse *oci_database.DisableExternalNonContainerDatabaseOperationsInsightsResponse
}

type DatabaseExternalNonContainerDatabaseOperationsInsightsManagementResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_database.DatabaseClient
	Res                    *DatabaseExternalNonContainerDatabaseOperationsInsightsResponse
	DisableNotFoundRetries bool
	WorkRequestClient      *oci_work_requests.WorkRequestClient
}

func (s *DatabaseExternalNonContainerDatabaseOperationsInsightsManagementResourceCrud) ID() string {
	return tfresource.GenerateDataSourceHashID("DatabaseExternalNonContainerDatabaseOperationsInsightsManagementResource-", DatabaseExternalNonContainerDatabaseOperationsInsightsManagementResource(), s.D)
}

func (s *DatabaseExternalNonContainerDatabaseOperationsInsightsManagementResourceCrud) Create() error {
	var operation bool
	if enableOperationsInsights, ok := s.D.GetOkExists("enable_operations_insights"); ok {
		operation = enableOperationsInsights.(bool)
	}

	if operation {
		// enable operations insights
		request := oci_database.EnableExternalNonContainerDatabaseOperationsInsightsRequest{}

		if externalDatabaseConnectorId, ok := s.D.GetOkExists("external_database_connector_id"); ok {
			tmp := externalDatabaseConnectorId.(string)
			request.ExternalDatabaseConnectorId = &tmp
		}

		if externalNonContainerDatabaseId, ok := s.D.GetOkExists("external_non_container_database_id"); ok {
			tmp := externalNonContainerDatabaseId.(string)
			request.ExternalNonContainerDatabaseId = &tmp
		}

		request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database")

		response, err := s.Client.EnableExternalNonContainerDatabaseOperationsInsights(context.Background(), request)
		if err != nil {
			return err
		}

		workId := response.OpcWorkRequestId
		if workId != nil {
			_, err = tfresource.WaitForWorkRequestWithErrorHandling(s.WorkRequestClient, workId, "externalNonContainerDatabase", oci_work_requests.WorkRequestResourceActionTypeUpdated, s.D.Timeout(schema.TimeoutCreate), s.DisableNotFoundRetries)
			if err != nil {
				return err
			}
		}
		s.Res.enableResponse = &response
		return nil
	}

	// disable operations insights
	request := oci_database.DisableExternalNonContainerDatabaseOperationsInsightsRequest{}

	if externalNonContainerDatabaseId, ok := s.D.GetOkExists("external_non_container_database_id"); ok {
		tmp := externalNonContainerDatabaseId.(string)
		request.ExternalNonContainerDatabaseId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database")

	response, err := s.Client.DisableExternalNonContainerDatabaseOperationsInsights(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	if workId != nil {
		_, err := tfresource.WaitForWorkRequestWithErrorHandling(s.WorkRequestClient, workId, "externalNonContainerDatabase", oci_work_requests.WorkRequestResourceActionTypeCreated, s.D.Timeout(schema.TimeoutCreate), s.DisableNotFoundRetries)
		if err != nil {
			return err
		}
	}
	s.Res.disableResponse = &response
	return nil
}

func (s *DatabaseExternalNonContainerDatabaseOperationsInsightsManagementResourceCrud) Update() error {

	var operation bool
	if enableOperationsInsights, ok := s.D.GetOkExists("enable_operations_insights"); ok {
		operation = enableOperationsInsights.(bool)
	}

	if operation {
		// enable operations insights
		request := oci_database.EnableExternalNonContainerDatabaseOperationsInsightsRequest{}

		if externalDatabaseConnectorId, ok := s.D.GetOkExists("external_database_connector_id"); ok {
			tmp := externalDatabaseConnectorId.(string)
			request.ExternalDatabaseConnectorId = &tmp
		}

		if externalNonContainerDatabaseId, ok := s.D.GetOkExists("external_non_container_database_id"); ok {
			tmp := externalNonContainerDatabaseId.(string)
			request.ExternalNonContainerDatabaseId = &tmp
		}

		request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database")

		response, err := s.Client.EnableExternalNonContainerDatabaseOperationsInsights(context.Background(), request)
		if err != nil {
			return err
		}

		workId := response.OpcWorkRequestId
		if workId != nil {
			_, err = tfresource.WaitForWorkRequestWithErrorHandling(s.WorkRequestClient, workId, "externalNonContainerDatabase", oci_work_requests.WorkRequestResourceActionTypeUpdated, s.D.Timeout(schema.TimeoutCreate), s.DisableNotFoundRetries)
			if err != nil {
				return err
			}
		}
		s.Res.enableResponse = &response
		return nil
	}

	// disable operations insights
	request := oci_database.DisableExternalNonContainerDatabaseOperationsInsightsRequest{}

	if externalNonContainerDatabaseId, ok := s.D.GetOkExists("external_non_container_database_id"); ok {
		tmp := externalNonContainerDatabaseId.(string)
		request.ExternalNonContainerDatabaseId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database")

	response, err := s.Client.DisableExternalNonContainerDatabaseOperationsInsights(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	if workId != nil {
		_, err = tfresource.WaitForWorkRequestWithErrorHandling(s.WorkRequestClient, workId, "externalNonContainerDatabase", oci_work_requests.WorkRequestResourceActionTypeUpdated, s.D.Timeout(schema.TimeoutCreate), s.DisableNotFoundRetries)
		if err != nil {
			return err
		}
	}
	s.Res.disableResponse = &response
	return nil
}

func (s *DatabaseExternalNonContainerDatabaseOperationsInsightsManagementResourceCrud) Delete() error {
	var operation bool
	if enableOpsi, ok := s.D.GetOkExists("enable_operations_insights"); ok {
		operation = enableOpsi.(bool)
	}

	if !operation {
		return nil
	}
	// disable database operations insights
	request := oci_database.DisableExternalNonContainerDatabaseOperationsInsightsRequest{}

	if externalNonContainerDatabaseId, ok := s.D.GetOkExists("external_non_container_database_id"); ok {
		tmp := externalNonContainerDatabaseId.(string)
		request.ExternalNonContainerDatabaseId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database")

	response, err := s.Client.DisableExternalNonContainerDatabaseOperationsInsights(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	if workId != nil {
		_, err = tfresource.WaitForWorkRequestWithErrorHandling(s.WorkRequestClient, workId, "externalNonContainerDatabase", oci_work_requests.WorkRequestResourceActionTypeUpdated, s.D.Timeout(schema.TimeoutCreate), s.DisableNotFoundRetries)
		if err != nil {
			return err
		}
	}
	s.Res.disableResponse = &response
	return nil
}

func (s *DatabaseExternalNonContainerDatabaseOperationsInsightsManagementResourceCrud) SetData() error {
	return nil
}
