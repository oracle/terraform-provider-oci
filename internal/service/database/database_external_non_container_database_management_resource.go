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

func DatabaseExternalNonContainerDatabaseManagementResource() *schema.Resource {
	return &schema.Resource{
		Timeouts: tfresource.DefaultTimeout,
		Create:   createDatabaseExternalNonContainerDatabaseManagement,
		Update:   updateDatabaseExternalNonContainerDatabaseManagement,
		Read:     readDatabaseExternalNonContainerDatabaseManagement,
		Delete:   deleteDatabaseExternalNonContainerDatabaseManagement,
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

func createDatabaseExternalNonContainerDatabaseManagement(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseExternalNonContainerDatabaseManagementResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()
	sync.WorkRequestClient = m.(*client.OracleClients).WorkRequestClient
	sync.Res = &DatabaseExternalNonContainerDatabaseManagementResponse{}

	return tfresource.CreateResource(d, sync)
}

func updateDatabaseExternalNonContainerDatabaseManagement(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseExternalNonContainerDatabaseManagementResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()
	sync.WorkRequestClient = m.(*client.OracleClients).WorkRequestClient
	sync.Res = &DatabaseExternalNonContainerDatabaseManagementResponse{}
	return tfresource.UpdateResource(d, sync)
}

func readDatabaseExternalNonContainerDatabaseManagement(d *schema.ResourceData, m interface{}) error {
	return nil
}

func deleteDatabaseExternalNonContainerDatabaseManagement(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseExternalNonContainerDatabaseManagementResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()
	sync.WorkRequestClient = m.(*client.OracleClients).WorkRequestClient
	sync.Res = &DatabaseExternalNonContainerDatabaseManagementResponse{}
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type DatabaseExternalNonContainerDatabaseManagementResponse struct {
	enableResponse  *oci_database.EnableExternalNonContainerDatabaseDatabaseManagementResponse
	disableResponse *oci_database.DisableExternalNonContainerDatabaseDatabaseManagementResponse
}

type DatabaseExternalNonContainerDatabaseManagementResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_database.DatabaseClient
	Res                    *DatabaseExternalNonContainerDatabaseManagementResponse
	DisableNotFoundRetries bool
	WorkRequestClient      *oci_work_requests.WorkRequestClient
}

func (s *DatabaseExternalNonContainerDatabaseManagementResourceCrud) ID() string {
	return tfresource.GenerateDataSourceHashID("DatabaseExternalNonContainerDatabaseManagementResource-", DatabaseExternalNonContainerDatabaseManagementResource(), s.D)
}

func (s *DatabaseExternalNonContainerDatabaseManagementResourceCrud) Create() error {

	var operation bool
	if enableManagement, ok := s.D.GetOkExists("enable_management"); ok {
		operation = enableManagement.(bool)
	}

	if operation {
		// enable operation
		request := oci_database.EnableExternalNonContainerDatabaseDatabaseManagementRequest{}

		if externalNonContainerDatabaseId, ok := s.D.GetOkExists("external_non_container_database_id"); ok {
			tmp := externalNonContainerDatabaseId.(string)
			request.ExternalNonContainerDatabaseId = &tmp
		}

		if externalDatabaseConnectorId, ok := s.D.GetOkExists("external_database_connector_id"); ok {
			tmp := externalDatabaseConnectorId.(string)
			request.EnableExternalNonContainerDatabaseDatabaseManagementDetails.ExternalDatabaseConnectorId = &tmp
		}

		if licenseModel, ok := s.D.GetOkExists("license_model"); ok {
			request.EnableExternalNonContainerDatabaseDatabaseManagementDetails.LicenseModel = oci_database.EnableExternalNonContainerDatabaseDatabaseManagementDetailsLicenseModelEnum(licenseModel.(string))
		}
		request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database")

		response, err := s.Client.EnableExternalNonContainerDatabaseDatabaseManagement(context.Background(), request)
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
	// disable
	request := oci_database.DisableExternalNonContainerDatabaseDatabaseManagementRequest{}

	if externalNonContainerDatabaseId, ok := s.D.GetOkExists("external_non_container_database_id"); ok {
		tmp := externalNonContainerDatabaseId.(string)
		request.ExternalNonContainerDatabaseId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database")

	response, err := s.Client.DisableExternalNonContainerDatabaseDatabaseManagement(context.Background(), request)
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

func (s *DatabaseExternalNonContainerDatabaseManagementResourceCrud) Update() error {
	var operation bool
	if enableManagement, ok := s.D.GetOkExists("enable_management"); ok {
		operation = enableManagement.(bool)
	}

	if operation {
		// enable database management
		request := oci_database.EnableExternalNonContainerDatabaseDatabaseManagementRequest{}

		if externalNonContainerDatabaseId, ok := s.D.GetOkExists("external_non_container_database_id"); ok {
			tmp := externalNonContainerDatabaseId.(string)
			request.ExternalNonContainerDatabaseId = &tmp
		}

		if externalDatabaseConnectorId, ok := s.D.GetOkExists("external_database_connector_id"); ok {
			tmp := externalDatabaseConnectorId.(string)
			request.EnableExternalNonContainerDatabaseDatabaseManagementDetails.ExternalDatabaseConnectorId = &tmp
		}

		if licenseModel, ok := s.D.GetOkExists("license_model"); ok {
			request.EnableExternalNonContainerDatabaseDatabaseManagementDetails.LicenseModel = oci_database.EnableExternalNonContainerDatabaseDatabaseManagementDetailsLicenseModelEnum(licenseModel.(string))
		}

		request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database")

		response, err := s.Client.EnableExternalNonContainerDatabaseDatabaseManagement(context.Background(), request)
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
	// disable database management
	request := oci_database.DisableExternalNonContainerDatabaseDatabaseManagementRequest{}

	if externalNonContainerDatabaseId, ok := s.D.GetOkExists("external_non_container_database_id"); ok {
		tmp := externalNonContainerDatabaseId.(string)
		request.ExternalNonContainerDatabaseId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database")

	response, err := s.Client.DisableExternalNonContainerDatabaseDatabaseManagement(context.Background(), request)
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

func (s *DatabaseExternalNonContainerDatabaseManagementResourceCrud) Delete() error {
	var operation bool
	if enableManagement, ok := s.D.GetOkExists("enable_management"); ok {
		operation = enableManagement.(bool)
	}

	if !operation {
		return nil
	}
	// disable database management
	request := oci_database.DisableExternalNonContainerDatabaseDatabaseManagementRequest{}

	if externalNonContainerDatabaseId, ok := s.D.GetOkExists("external_non_container_database_id"); ok {
		tmp := externalNonContainerDatabaseId.(string)
		request.ExternalNonContainerDatabaseId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database")

	response, err := s.Client.DisableExternalNonContainerDatabaseDatabaseManagement(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	if workId != nil {
		_, err := tfresource.WaitForWorkRequestWithErrorHandling(s.WorkRequestClient, workId, "externalNonContainerDatabase", oci_work_requests.WorkRequestResourceActionTypeUpdated, s.D.Timeout(schema.TimeoutCreate), s.DisableNotFoundRetries)
		if err != nil {
			return err
		}
	}
	s.Res.disableResponse = &response
	return nil
}

func (s *DatabaseExternalNonContainerDatabaseManagementResourceCrud) SetData() error {
	return nil
}
