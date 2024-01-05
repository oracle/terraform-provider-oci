// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package database

import (
	"context"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	oci_database "github.com/oracle/oci-go-sdk/v65/database"
	oci_work_requests "github.com/oracle/oci-go-sdk/v65/workrequests"
)

func DatabaseExternalcontainerdatabasesStackMonitoringResource() *schema.Resource {
	return &schema.Resource{
		Timeouts: tfresource.DefaultTimeout,
		Create:   createDatabaseExternalcontainerdatabasesStackMonitoring,
		Update:   updateDatabaseExternalcontainerdatabasesStackMonitoring,
		Read:     readDatabaseExternalcontainerdatabasesStackMonitoring,
		Delete:   deleteDatabaseExternalcontainerdatabasesStackMonitoring,
		Schema: map[string]*schema.Schema{
			// Required
			"external_container_database_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"external_database_connector_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"enable_stack_monitoring": {
				Type:     schema.TypeBool,
				Required: true,
			},

			// Optional

			// Computed
		},
	}
}

func createDatabaseExternalcontainerdatabasesStackMonitoring(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseExternalcontainerdatabasesStackMonitoringResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()
	sync.WorkRequestClient = m.(*client.OracleClients).WorkRequestClient
	sync.Res = &DatabaseExternalcontainerdatabaseStackMonitoringResponse{}
	return tfresource.CreateResource(d, sync)
}

func updateDatabaseExternalcontainerdatabasesStackMonitoring(d *schema.ResourceData, m interface{}) error {
	return nil
}

func readDatabaseExternalcontainerdatabasesStackMonitoring(d *schema.ResourceData, m interface{}) error {
	return nil
}

func deleteDatabaseExternalcontainerdatabasesStackMonitoring(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseExternalcontainerdatabasesStackMonitoringResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()
	sync.WorkRequestClient = m.(*client.OracleClients).WorkRequestClient
	sync.Res = &DatabaseExternalcontainerdatabaseStackMonitoringResponse{}
	sync.DisableNotFoundRetries = true
	return tfresource.DeleteResource(d, sync)
}

type DatabaseExternalcontainerdatabaseStackMonitoringResponse struct {
	enableResponse  *oci_database.EnableExternalContainerDatabaseStackMonitoringResponse
	disableResponse *oci_database.DisableExternalContainerDatabaseStackMonitoringResponse
}

type DatabaseExternalcontainerdatabasesStackMonitoringResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_database.DatabaseClient
	Res                    *DatabaseExternalcontainerdatabaseStackMonitoringResponse
	DisableNotFoundRetries bool
	WorkRequestClient      *oci_work_requests.WorkRequestClient
}

func (s *DatabaseExternalcontainerdatabasesStackMonitoringResourceCrud) ID() string {
	return tfresource.GenerateDataSourceHashID("DatabaseExternalcontainerdatabasesStackMonitoringResource-", DatabaseExternalcontainerdatabasesStackMonitoringResource(), s.D)
}

func (s *DatabaseExternalcontainerdatabasesStackMonitoringResourceCrud) Create() error {
	var operation_stack_monitoring bool
	if enableStackMonitoring, ok := s.D.GetOkExists("enable_stack_monitoring"); ok {
		operation_stack_monitoring = enableStackMonitoring.(bool)
	}

	if operation_stack_monitoring {
		// enable stack monitoring
		request := oci_database.EnableExternalContainerDatabaseStackMonitoringRequest{}

		if externalDatabaseConnectorId, ok := s.D.GetOkExists("external_database_connector_id"); ok {
			tmp := externalDatabaseConnectorId.(string)
			request.ExternalDatabaseConnectorId = &tmp
		}

		if externalContainerDatabaseId, ok := s.D.GetOkExists("external_container_database_id"); ok {
			tmp := externalContainerDatabaseId.(string)
			request.ExternalContainerDatabaseId = &tmp
		}

		request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database")

		response, err := s.Client.EnableExternalContainerDatabaseStackMonitoring(context.Background(), request)
		if err != nil {
			return err
		}
		workId := response.OpcWorkRequestId
		if workId != nil {
			identifier, err := tfresource.WaitForWorkRequestWithErrorHandling(s.WorkRequestClient, workId, "externalContainerDatabase", oci_work_requests.WorkRequestResourceActionTypeUpdated, s.D.Timeout(schema.TimeoutCreate), s.DisableNotFoundRetries)
			if identifier != nil {
				s.D.SetId(*identifier)
			}
			if err != nil {
				return err
			}
		}
		s.Res.enableResponse = &response
		return nil
	}
	// disable stack monitoring
	request := oci_database.DisableExternalContainerDatabaseStackMonitoringRequest{}

	if externalContainerDatabaseId, ok := s.D.GetOkExists("external_container_database_id"); ok {
		tmp := externalContainerDatabaseId.(string)
		request.ExternalContainerDatabaseId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database")

	response, err := s.Client.DisableExternalContainerDatabaseStackMonitoring(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	if workId != nil {
		_, err = tfresource.WaitForWorkRequestWithErrorHandling(s.WorkRequestClient, workId, "externalContainerDatabase", oci_work_requests.WorkRequestResourceActionTypeUpdated, s.D.Timeout(schema.TimeoutCreate), s.DisableNotFoundRetries)
		if err != nil {
			return err
		}
	}
	s.Res.disableResponse = &response
	return nil
}

func (s *DatabaseExternalcontainerdatabasesStackMonitoringResourceCrud) Delete() error {
	var operation_stack_monitoring bool
	if enableOpsi, ok := s.D.GetOkExists("enable_stack_monitoring"); ok {
		operation_stack_monitoring = enableOpsi.(bool)
	}

	if !operation_stack_monitoring {
		return nil
	}
	// Disable Operations Insights
	request := oci_database.DisableExternalContainerDatabaseStackMonitoringRequest{}

	if externalContainerDatabaseId, ok := s.D.GetOkExists("external_container_database_id"); ok {
		tmp := externalContainerDatabaseId.(string)
		request.ExternalContainerDatabaseId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database")

	response, err := s.Client.DisableExternalContainerDatabaseStackMonitoring(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	if workId != nil {
		// verification required for entity type name "externalPluggableDatabase"
		_, err = tfresource.WaitForWorkRequestWithErrorHandling(s.WorkRequestClient, workId, "externalContainerDatabase", oci_work_requests.WorkRequestResourceActionTypeUpdated, s.D.Timeout(schema.TimeoutCreate), s.DisableNotFoundRetries)
		if err != nil {
			return err
		}
	}
	s.Res.disableResponse = &response
	return nil
}

func (s *DatabaseExternalcontainerdatabasesStackMonitoringResourceCrud) SetData() error {
	return nil
}
