// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package database

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	oci_database "github.com/oracle/oci-go-sdk/v65/database"
	oci_work_requests "github.com/oracle/oci-go-sdk/v65/workrequests"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DatabaseDatabaseDataPatchResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createDatabaseDataPatchResource,
		Read:     readDatabaseDataPatchResource,
		Delete:   deleteDatabaseDataPatchResource,
		Schema: map[string]*schema.Schema{
			"database_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"action": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"data_patch_options": {
				Type:     schema.TypeList,
				Optional: true,
				ForceNew: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"should_skip_closed_pdbs": {
							Type:     schema.TypeBool,
							Optional: true,
						},
					},
				},
			},
			"pluggable_databases": {
				Type:     schema.TypeList,
				Optional: true,
				ForceNew: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
		},
	}
}

func createDatabaseDataPatchResource(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseDatabaseDataPatchResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()
	sync.WorkRequestClient = m.(*client.OracleClients).WorkRequestClient

	return tfresource.CreateResource(d, sync)
}

func readDatabaseDataPatchResource(d *schema.ResourceData, m interface{}) error {
	return nil
}

func deleteDatabaseDataPatchResource(d *schema.ResourceData, m interface{}) error {
	return nil
}

type DatabaseDatabaseDataPatchResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_database.DatabaseClient
	Res                    *oci_database.Database
	DisableNotFoundRetries bool
	WorkRequestClient      *oci_work_requests.WorkRequestClient
}

func (s *DatabaseDatabaseDataPatchResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *DatabaseDatabaseDataPatchResourceCrud) SetData() error {
	return nil
}

func (s *DatabaseDatabaseDataPatchResourceCrud) Create() error {
	request := oci_database.RunDataPatchRequest{}
	if databaseId, ok := s.D.GetOkExists("database_id"); ok {
		tmp := databaseId.(string)
		request.DatabaseId = &tmp
	}

	action, ok := s.D.GetOkExists("action")
	if !ok {
		return fmt.Errorf("action is required")
	}
	request.Action = oci_database.RunDataPatchDetailsActionEnum(action.(string))

	if dataPatchOptions, ok := s.D.GetOkExists("data_patch_options"); ok {
		if dataPatchOptionsList, ok := dataPatchOptions.([]interface{}); ok && len(dataPatchOptionsList) > 0 {
			dataPatchOptionsMap := dataPatchOptionsList[0].(map[string]interface{})
			dataPatchOptionsObject := oci_database.DataPatchOptions{}
			if shouldSkipClosedPdbs, ok := dataPatchOptionsMap["should_skip_closed_pdbs"].(bool); ok {
				dataPatchOptionsObject.ShouldSkipClosedPdbs = &shouldSkipClosedPdbs
			}
			request.DataPatchOptions = &dataPatchOptionsObject
		}
	}

	if pluggableDatabases, ok := s.D.GetOkExists("pluggable_databases"); ok {
		interfaces := pluggableDatabases.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		request.PluggableDatabases = tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database")

	response, err := s.Client.RunDataPatch(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	s.Res = &response.Database

	if workId != nil {
		var identifier *string
		var err error
		identifier, err = tfresource.WaitForWorkRequestWithErrorHandling(s.WorkRequestClient, workId, "database", oci_work_requests.WorkRequestResourceActionTypeUpdated, s.D.Timeout(schema.TimeoutCreate), s.DisableNotFoundRetries)
		if err != nil {
			return err
		}
		// Prefer identifier from work request if present to set resource ID
		if identifier != nil {
			s.D.SetId(*identifier)
		}
	}

	return nil

}
