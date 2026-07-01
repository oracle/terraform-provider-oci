// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package database

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_database "github.com/oracle/oci-go-sdk/v65/database"
	oci_work_requests "github.com/oracle/oci-go-sdk/v65/workrequests"
	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DatabaseDatabaseSoftwareScheduleManagementResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts:      tfresource.DefaultTimeout,
		CreateContext: createDatabaseDatabaseSoftwareScheduleManagementWithContext,
		ReadContext:   readDatabaseDatabaseSoftwareScheduleManagementWithContext,
		DeleteContext: deleteDatabaseDatabaseSoftwareScheduleManagementWithContext,
		Schema: map[string]*schema.Schema{
			// Required
			"database_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func createDatabaseDatabaseSoftwareScheduleManagementWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &DatabaseDatabaseSoftwareScheduleManagementResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()
	sync.WorkRequestClient = m.(*client.OracleClients).WorkRequestClient

	return tfresource.HandleDiagError(m, tfresource.CreateResourceWithContext(ctx, d, sync))
}

func readDatabaseDatabaseSoftwareScheduleManagementWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	return nil
}

func deleteDatabaseDatabaseSoftwareScheduleManagementWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	return nil
}

type DatabaseDatabaseSoftwareScheduleManagementResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_database.DatabaseClient
	Res                    *oci_database.Database
	DisableNotFoundRetries bool
	WorkRequestClient      *oci_work_requests.WorkRequestClient
}

func (s *DatabaseDatabaseSoftwareScheduleManagementResourceCrud) CreateWithContext(ctx context.Context) error {
	request := oci_database.RescheduleManagedDbSoftwareUpdateRequest{}

	if databaseId, ok := s.D.GetOkExists("database_id"); ok {
		tmp := databaseId.(string)
		request.DatabaseId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database")

	response, err := s.Client.RescheduleManagedDbSoftwareUpdate(ctx, request)
	if err != nil {
		return fmt.Errorf("failed to call RescheduleManagedDbSoftwareUpdate: %w", err)
	}

	workId := response.OpcWorkRequestId
	s.Res = &response.Database

	if workId != nil {
		var identifier *string
		var err error
		identifier = response.Id
		if identifier != nil {
			s.D.SetId(*identifier)
		}
		identifier, err = tfresource.WaitForWorkRequestWithErrorHandlingAndContext(ctx, s.WorkRequestClient, workId, "database", oci_work_requests.WorkRequestResourceActionTypeUpdated, s.D.Timeout(schema.TimeoutCreate), s.DisableNotFoundRetries)
		if identifier != nil {
			s.D.SetId(*identifier)
		}
		if err != nil {
			return err
		}
	}

	return nil
}
func (s *DatabaseDatabaseSoftwareScheduleManagementResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *DatabaseDatabaseSoftwareScheduleManagementResourceCrud) SetData() error {
	if s.Res.Id != nil {
		s.D.Set("database_id", *s.Res.Id)
	}
	return nil
}
