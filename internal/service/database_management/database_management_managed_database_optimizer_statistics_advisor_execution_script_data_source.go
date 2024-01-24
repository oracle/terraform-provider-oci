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

func DatabaseManagementManagedDatabaseOptimizerStatisticsAdvisorExecutionScriptDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularDatabaseManagementManagedDatabaseOptimizerStatisticsAdvisorExecutionScript,
		Schema: map[string]*schema.Schema{
			"execution_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"managed_database_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"task_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			// Computed
			"script": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func readSingularDatabaseManagementManagedDatabaseOptimizerStatisticsAdvisorExecutionScript(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseManagementManagedDatabaseOptimizerStatisticsAdvisorExecutionScriptDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DbManagementClient()

	return tfresource.ReadResource(sync)
}

type DatabaseManagementManagedDatabaseOptimizerStatisticsAdvisorExecutionScriptDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_database_management.DbManagementClient
	Res    *oci_database_management.GetOptimizerStatisticsAdvisorExecutionScriptResponse
}

func (s *DatabaseManagementManagedDatabaseOptimizerStatisticsAdvisorExecutionScriptDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DatabaseManagementManagedDatabaseOptimizerStatisticsAdvisorExecutionScriptDataSourceCrud) Get() error {
	request := oci_database_management.GetOptimizerStatisticsAdvisorExecutionScriptRequest{}

	if executionName, ok := s.D.GetOkExists("execution_name"); ok {
		tmp := executionName.(string)
		request.ExecutionName = &tmp
	}

	if managedDatabaseId, ok := s.D.GetOkExists("managed_database_id"); ok {
		tmp := managedDatabaseId.(string)
		request.ManagedDatabaseId = &tmp
	}

	if taskName, ok := s.D.GetOkExists("task_name"); ok {
		tmp := taskName.(string)
		request.TaskName = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "database_management")

	response, err := s.Client.GetOptimizerStatisticsAdvisorExecutionScript(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *DatabaseManagementManagedDatabaseOptimizerStatisticsAdvisorExecutionScriptDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DatabaseManagementManagedDatabaseOptimizerStatisticsAdvisorExecutionScriptDataSource-", DatabaseManagementManagedDatabaseOptimizerStatisticsAdvisorExecutionScriptDataSource(), s.D))

	if s.Res.Script != nil {
		s.D.Set("script", *s.Res.Script)
	}

	return nil
}
