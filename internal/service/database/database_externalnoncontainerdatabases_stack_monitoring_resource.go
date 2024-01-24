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

func DatabaseExternalnoncontainerdatabasesStackMonitoringResource() *schema.Resource {
	return &schema.Resource{
		Timeouts: tfresource.DefaultTimeout,
		Create:   createDatabaseExternalnoncontainerdatabasesStackMonitoring,
		Update:   updateDatabaseExternalnoncontainerdatabasesStackMonitoring,
		Read:     readDatabaseExternalnoncontainerdatabasesStackMonitoring,
		Delete:   deleteDatabaseExternalnoncontainerdatabasesStackMonitoring,
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

			"enable_stack_monitoring": {
				Type:     schema.TypeBool,
				Required: true,
			},

			// Optional

			// Computed
		},
	}
}

func createDatabaseExternalnoncontainerdatabasesStackMonitoring(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseExternalnoncontainerdatabasesStackMonitoringResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()
	sync.WorkRequestClient = m.(*client.OracleClients).WorkRequestClient
	sync.Res = &DatabaseExternalnoncontainerDatabaseStackMonitoringResponse{}
	return tfresource.CreateResource(d, sync)
}

func updateDatabaseExternalnoncontainerdatabasesStackMonitoring(d *schema.ResourceData, m interface{}) error {
	return nil
}

func readDatabaseExternalnoncontainerdatabasesStackMonitoring(d *schema.ResourceData, m interface{}) error {
	return nil
}

func deleteDatabaseExternalnoncontainerdatabasesStackMonitoring(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseExternalnoncontainerdatabasesStackMonitoringResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()
	sync.WorkRequestClient = m.(*client.OracleClients).WorkRequestClient
	sync.Res = &DatabaseExternalnoncontainerDatabaseStackMonitoringResponse{}
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type DatabaseExternalnoncontainerDatabaseStackMonitoringResponse struct {
	enableResponse  *oci_database.EnableExternalNonContainerDatabaseStackMonitoringResponse
	disableResponse *oci_database.DisableExternalNonContainerDatabaseStackMonitoringResponse
}

type DatabaseExternalnoncontainerdatabasesStackMonitoringResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_database.DatabaseClient
	Res                    *DatabaseExternalnoncontainerDatabaseStackMonitoringResponse
	DisableNotFoundRetries bool
	WorkRequestClient      *oci_work_requests.WorkRequestClient
}

func (s *DatabaseExternalnoncontainerdatabasesStackMonitoringResourceCrud) ID() string {
	return tfresource.GenerateDataSourceHashID("DatabaseExternalnoncontainerdatabasesStackMonitoringResource-", DatabaseExternalnoncontainerdatabasesStackMonitoringResource(), s.D)
}

func (s *DatabaseExternalnoncontainerdatabasesStackMonitoringResourceCrud) Create() error {

	var operation_stack_monitoring bool
	if enableStackMonitoring, ok := s.D.GetOkExists("enable_stack_monitoring"); ok {
		operation_stack_monitoring = enableStackMonitoring.(bool)
	}

	if operation_stack_monitoring {
		// enable stack monitoring
		request := oci_database.EnableExternalNonContainerDatabaseStackMonitoringRequest{}

		if externalDatabaseConnectorId, ok := s.D.GetOkExists("external_database_connector_id"); ok {
			tmp := externalDatabaseConnectorId.(string)
			request.ExternalDatabaseConnectorId = &tmp
		}

		if externalNonContainerDatabaseId, ok := s.D.GetOkExists("external_non_container_database_id"); ok {
			tmp := externalNonContainerDatabaseId.(string)
			request.ExternalNonContainerDatabaseId = &tmp
		}

		request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database")

		response, err := s.Client.EnableExternalNonContainerDatabaseStackMonitoring(context.Background(), request)
		if err != nil {
			return err
		}
		workId := response.OpcWorkRequestId
		if workId != nil {
			identifier, err := tfresource.WaitForWorkRequestWithErrorHandling(s.WorkRequestClient, workId, "externalNonContainerDatabase", oci_work_requests.WorkRequestResourceActionTypeUpdated, s.D.Timeout(schema.TimeoutCreate), s.DisableNotFoundRetries)
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
	request := oci_database.DisableExternalNonContainerDatabaseStackMonitoringRequest{}

	if externalNonContainerDatabaseId, ok := s.D.GetOkExists("external_non_container_database_id"); ok {
		tmp := externalNonContainerDatabaseId.(string)
		request.ExternalNonContainerDatabaseId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database")

	response, err := s.Client.DisableExternalNonContainerDatabaseStackMonitoring(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	if workId != nil {
		_, err = tfresource.WaitForWorkRequestWithErrorHandling(s.WorkRequestClient, workId, "externalNonContainerDatabase", oci_work_requests.WorkRequestResourceActionTypeCreated, s.D.Timeout(schema.TimeoutCreate), s.DisableNotFoundRetries)
		if err != nil {
			return err
		}
	}
	s.Res.disableResponse = &response
	return nil
}

func (s *DatabaseExternalnoncontainerdatabasesStackMonitoringResourceCrud) Delete() error {
	var operation_stack_monitoring bool
	if enableOpsi, ok := s.D.GetOkExists("enable_stack_monitoring"); ok {
		operation_stack_monitoring = enableOpsi.(bool)
	}

	if !operation_stack_monitoring {
		return nil
	}
	// Disable Operations Insights
	request := oci_database.DisableExternalNonContainerDatabaseStackMonitoringRequest{}

	if externalNonContainerDatabaseId, ok := s.D.GetOkExists("external_non_container_database_id"); ok {
		tmp := externalNonContainerDatabaseId.(string)
		request.ExternalNonContainerDatabaseId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database")

	response, err := s.Client.DisableExternalNonContainerDatabaseStackMonitoring(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	if workId != nil {
		// verification required for entity type name "externalPluggableDatabase"
		_, err = tfresource.WaitForWorkRequestWithErrorHandling(s.WorkRequestClient, workId, "externalNonContainerDatabase", oci_work_requests.WorkRequestResourceActionTypeUpdated, s.D.Timeout(schema.TimeoutCreate), s.DisableNotFoundRetries)
		if err != nil {
			return err
		}
	}
	s.Res.disableResponse = &response
	return nil
}

func (s *DatabaseExternalnoncontainerdatabasesStackMonitoringResourceCrud) SetData() error {
	return nil
}
