// Copyright (c) 2017, 2020, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	oci_database "github.com/oracle/oci-go-sdk/v38/database"
	oci_work_requests "github.com/oracle/oci-go-sdk/v38/workrequests"
)

func init() {
	RegisterResource("oci_database_external_container_database_management", DatabaseExternalContainerDatabaseManagementResource())
}

func DatabaseExternalContainerDatabaseManagementResource() *schema.Resource {
	return &schema.Resource{
		Timeouts: DefaultTimeout,
		Create:   createDatabaseExternalContainerDatabaseManagement,
		Update:   updateDatabaseExternalContainerDatabaseManagement,
		Read:     readDatabaseExternalContainerDatabaseManagement,
		Delete:   deleteDatabaseExternalContainerDatabaseManagement,
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
			"enable_management": {
				Type:     schema.TypeBool,
				Required: true,
			},

			// Optional
			"license_model": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			// Computed
		},
	}
}

func createDatabaseExternalContainerDatabaseManagement(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseExternalContainerDatabaseManagementResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).databaseClient()
	sync.workRequestClient = m.(*OracleClients).workRequestClient
	sync.Res = &DatabaseExternalContainerDatabaseManagementResponse{}
	return CreateResource(d, sync)
}

func updateDatabaseExternalContainerDatabaseManagement(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseExternalContainerDatabaseManagementResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).databaseClient()
	sync.workRequestClient = m.(*OracleClients).workRequestClient
	sync.Res = &DatabaseExternalContainerDatabaseManagementResponse{}
	return UpdateResource(d, sync)
}

func readDatabaseExternalContainerDatabaseManagement(d *schema.ResourceData, m interface{}) error {
	return nil
}

func deleteDatabaseExternalContainerDatabaseManagement(d *schema.ResourceData, m interface{}) error {
	return nil
}

type DatabaseExternalContainerDatabaseManagementResponse struct {
	enableResponse  *oci_database.EnableExternalContainerDatabaseDatabaseManagementResponse
	disableResponse *oci_database.DisableExternalContainerDatabaseDatabaseManagementResponse
}

type DatabaseExternalContainerDatabaseManagementResourceCrud struct {
	BaseCrud
	Client                 *oci_database.DatabaseClient
	workRequestClient      *oci_work_requests.WorkRequestClient
	Res                    *DatabaseExternalContainerDatabaseManagementResponse
	DisableNotFoundRetries bool
}

func (s *DatabaseExternalContainerDatabaseManagementResourceCrud) ID() string {
	return GenerateDataSourceHashID("DatabaseExternalContainerDatabaseManagementResource-", DatabaseExternalContainerDatabaseManagementResource(), s.D)
}

func (s *DatabaseExternalContainerDatabaseManagementResourceCrud) Create() error {

	var operation bool
	if enableManagement, ok := s.D.GetOkExists("enable_management"); ok {
		operation = enableManagement.(bool)
	}

	if operation {
		// enable operation
		request := oci_database.EnableExternalContainerDatabaseDatabaseManagementRequest{}

		if externalContainerDatabaseId, ok := s.D.GetOkExists("external_container_database_id"); ok {
			tmp := externalContainerDatabaseId.(string)
			request.ExternalContainerDatabaseId = &tmp
		}

		if externalDatabaseConnectorId, ok := s.D.GetOkExists("external_database_connector_id"); ok {
			tmp := externalDatabaseConnectorId.(string)
			request.EnableExternalContainerDatabaseDatabaseManagementDetails.ExternalDatabaseConnectorId = &tmp
		}

		if licenseModel, ok := s.D.GetOkExists("license_model"); ok {
			request.EnableExternalContainerDatabaseDatabaseManagementDetails.LicenseModel = oci_database.EnableExternalContainerDatabaseDatabaseManagementDetailsLicenseModelEnum(licenseModel.(string))
		}
		request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "database")

		response, err := s.Client.EnableExternalContainerDatabaseDatabaseManagement(context.Background(), request)
		if err != nil {
			return err
		}

		workId := response.OpcWorkRequestId
		if workId != nil {
			_, err = WaitForWorkRequestWithErrorHandling(s.workRequestClient, workId, "externalContainerDatabase", oci_work_requests.WorkRequestResourceActionTypeUpdated, s.D.Timeout(schema.TimeoutCreate), s.DisableNotFoundRetries)
			if err != nil {
				return err
			}
		}
		s.Res.enableResponse = &response
		return nil
	}
	// disable
	request := oci_database.DisableExternalContainerDatabaseDatabaseManagementRequest{}

	if externalContainerDatabaseId, ok := s.D.GetOkExists("external_container_database_id"); ok {
		tmp := externalContainerDatabaseId.(string)
		request.ExternalContainerDatabaseId = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "database")

	response, err := s.Client.DisableExternalContainerDatabaseDatabaseManagement(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	if workId != nil {
		_, err = WaitForWorkRequestWithErrorHandling(s.workRequestClient, workId, "externalContainerDatabase", oci_work_requests.WorkRequestResourceActionTypeUpdated, s.D.Timeout(schema.TimeoutCreate), s.DisableNotFoundRetries)
		if err != nil {
			return err
		}
	}
	s.Res.disableResponse = &response
	return nil
}

func (s *DatabaseExternalContainerDatabaseManagementResourceCrud) Update() error {
	var operation bool
	if enableManagement, ok := s.D.GetOkExists("enable_management"); ok {
		operation = enableManagement.(bool)
	}
	operation = false
	if operation {
		// enable operation
		request := oci_database.EnableExternalContainerDatabaseDatabaseManagementRequest{}

		if externalContainerDatabaseId, ok := s.D.GetOkExists("external_container_database_id"); ok {
			tmp := externalContainerDatabaseId.(string)
			request.ExternalContainerDatabaseId = &tmp
		}

		if externalDatabaseConnectorId, ok := s.D.GetOkExists("external_database_connector_id"); ok {
			tmp := externalDatabaseConnectorId.(string)
			request.EnableExternalContainerDatabaseDatabaseManagementDetails.ExternalDatabaseConnectorId = &tmp
		}

		if licenseModel, ok := s.D.GetOkExists("license_model"); ok {
			request.EnableExternalContainerDatabaseDatabaseManagementDetails.LicenseModel = oci_database.EnableExternalContainerDatabaseDatabaseManagementDetailsLicenseModelEnum(licenseModel.(string))
		}

		request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "database")

		response, err := s.Client.EnableExternalContainerDatabaseDatabaseManagement(context.Background(), request)
		if err != nil {
			return err
		}

		workId := response.OpcWorkRequestId
		if workId != nil {
			_, err = WaitForWorkRequestWithErrorHandling(s.workRequestClient, workId, "externalContainerDatabase", oci_work_requests.WorkRequestResourceActionTypeUpdated, s.D.Timeout(schema.TimeoutCreate), s.DisableNotFoundRetries)
			if err != nil {
				return err
			}
		}
		s.Res.enableResponse = &response
		return nil
	}
	// disable
	request := oci_database.DisableExternalContainerDatabaseDatabaseManagementRequest{}

	if externalContainerDatabaseId, ok := s.D.GetOkExists("external_container_database_id"); ok {
		tmp := externalContainerDatabaseId.(string)
		request.ExternalContainerDatabaseId = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "database")

	response, err := s.Client.DisableExternalContainerDatabaseDatabaseManagement(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	if workId != nil {
		_, err = WaitForWorkRequestWithErrorHandling(s.workRequestClient, workId, "externalContainerDatabase", oci_work_requests.WorkRequestResourceActionTypeUpdated, s.D.Timeout(schema.TimeoutCreate), s.DisableNotFoundRetries)
		if err != nil {
			return err
		}
	}
	s.Res.disableResponse = &response
	return nil
}

func (s *DatabaseExternalContainerDatabaseManagementResourceCrud) SetData() error {
	return nil
}
