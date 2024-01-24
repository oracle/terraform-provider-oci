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

func DatabaseExternalpluggabledatabasesStackMonitoringResource() *schema.Resource {
	return &schema.Resource{
		Timeouts: tfresource.DefaultTimeout,
		Create:   createDatabaseExternalpluggabledatabasesStackMonitoring,
		Update:   updateDatabaseExternalpluggabledatabasesStackMonitoring,
		Read:     readDatabaseExternalpluggabledatabasesStackMonitoring,
		Delete:   deleteDatabaseExternalpluggabledatabasesStackMonitoring,
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
			"enable_stack_monitoring": {
				Type:     schema.TypeBool,
				Required: true,
			},

			// Optional

			// Computed
		},
	}
}

func createDatabaseExternalpluggabledatabasesStackMonitoring(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseExternalpluggabledatabasesStackMonitoringResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()
	sync.WorkRequestClient = m.(*client.OracleClients).WorkRequestClient
	sync.Res = &DatabaseExternalpluggabledatabaseStackMonitoringResponse{}
	return tfresource.CreateResource(d, sync)
}

func updateDatabaseExternalpluggabledatabasesStackMonitoring(d *schema.ResourceData, m interface{}) error {
	return nil
}

func readDatabaseExternalpluggabledatabasesStackMonitoring(d *schema.ResourceData, m interface{}) error {
	return nil
}

func deleteDatabaseExternalpluggabledatabasesStackMonitoring(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseExternalpluggabledatabasesStackMonitoringResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()
	sync.WorkRequestClient = m.(*client.OracleClients).WorkRequestClient
	sync.Res = &DatabaseExternalpluggabledatabaseStackMonitoringResponse{}
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type DatabaseExternalpluggabledatabaseStackMonitoringResponse struct {
	enableResponse  *oci_database.EnableExternalPluggableDatabaseStackMonitoringResponse
	disableResponse *oci_database.DisableExternalPluggableDatabaseStackMonitoringResponse
}

type DatabaseExternalpluggabledatabasesStackMonitoringResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_database.DatabaseClient
	Res                    *DatabaseExternalpluggabledatabaseStackMonitoringResponse
	DisableNotFoundRetries bool
	WorkRequestClient      *oci_work_requests.WorkRequestClient
}

func (s *DatabaseExternalpluggabledatabasesStackMonitoringResourceCrud) ID() string {
	return tfresource.GenerateDataSourceHashID("DatabaseExternalpluggabledatabasesStackMonitoringResource-", DatabaseExternalpluggabledatabasesStackMonitoringResource(), s.D)
}

func (s *DatabaseExternalpluggabledatabasesStackMonitoringResourceCrud) Create() error {
	var operation_stack_monitoring bool
	if enableStackMonitoring, ok := s.D.GetOkExists("enable_stack_monitoring"); ok {
		operation_stack_monitoring = enableStackMonitoring.(bool)
	}

	if operation_stack_monitoring {
		// enable stack monitoring
		request := oci_database.EnableExternalPluggableDatabaseStackMonitoringRequest{}

		if externalDatabaseConnectorId, ok := s.D.GetOkExists("external_database_connector_id"); ok {
			tmp := externalDatabaseConnectorId.(string)
			request.ExternalDatabaseConnectorId = &tmp
		}

		if externalPluggableDatabaseId, ok := s.D.GetOkExists("external_pluggable_database_id"); ok {
			tmp := externalPluggableDatabaseId.(string)
			request.ExternalPluggableDatabaseId = &tmp
		}

		request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database")

		response, err := s.Client.EnableExternalPluggableDatabaseStackMonitoring(context.Background(), request)
		if err != nil {
			return err
		}
		workId := response.OpcWorkRequestId
		if workId != nil {
			identifier, err := tfresource.WaitForWorkRequestWithErrorHandling(s.WorkRequestClient, workId, "externalPluggableDatabase", oci_work_requests.WorkRequestResourceActionTypeUpdated, s.D.Timeout(schema.TimeoutCreate), s.DisableNotFoundRetries)
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
	request := oci_database.DisableExternalPluggableDatabaseStackMonitoringRequest{}

	if externalPluggableDatabaseId, ok := s.D.GetOkExists("external_pluggable_database_id"); ok {
		tmp := externalPluggableDatabaseId.(string)
		request.ExternalPluggableDatabaseId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database")

	response, err := s.Client.DisableExternalPluggableDatabaseStackMonitoring(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	if workId != nil {
		_, err = tfresource.WaitForWorkRequestWithErrorHandling(s.WorkRequestClient, workId, "externalPluggableDatabase", oci_work_requests.WorkRequestResourceActionTypeUpdated, s.D.Timeout(schema.TimeoutCreate), s.DisableNotFoundRetries)
		if err != nil {
			return err
		}
	}
	s.Res.disableResponse = &response
	return nil
}

func (s *DatabaseExternalpluggabledatabasesStackMonitoringResourceCrud) Delete() error {
	var operation_stack_monitoring bool
	if enableOpsi, ok := s.D.GetOkExists("enable_stack_monitoring"); ok {
		operation_stack_monitoring = enableOpsi.(bool)
	}

	if !operation_stack_monitoring {
		return nil
	}
	// Disable Operations Insights
	request := oci_database.DisableExternalPluggableDatabaseStackMonitoringRequest{}

	if externalPluggableDatabaseId, ok := s.D.GetOkExists("external_pluggable_database_id"); ok {
		tmp := externalPluggableDatabaseId.(string)
		request.ExternalPluggableDatabaseId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database")

	response, err := s.Client.DisableExternalPluggableDatabaseStackMonitoring(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	if workId != nil {
		// verification required for entity type name "externalPluggableDatabase"
		_, err = tfresource.WaitForWorkRequestWithErrorHandling(s.WorkRequestClient, workId, "externalPluggableDatabase", oci_work_requests.WorkRequestResourceActionTypeUpdated, s.D.Timeout(schema.TimeoutCreate), s.DisableNotFoundRetries)
		if err != nil {
			return err
		}
	}
	s.Res.disableResponse = &response
	return nil
}

func (s *DatabaseExternalpluggabledatabasesStackMonitoringResourceCrud) SetData() error {
	return nil
}
