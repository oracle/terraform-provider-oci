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

func DatabaseManagementManagedDatabaseSqlPlanBaselineDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularDatabaseManagementManagedDatabaseSqlPlanBaseline,
		Schema: map[string]*schema.Schema{
			"managed_database_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"opc_named_credential_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"plan_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			// Computed
			"accepted": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"action": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"adaptive": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"auto_purge": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"enabled": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"execution_plan": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"fixed": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"module": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"origin": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"reproduced": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"sql_handle": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"sql_text": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_created": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_last_executed": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_last_modified": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func readSingularDatabaseManagementManagedDatabaseSqlPlanBaseline(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseManagementManagedDatabaseSqlPlanBaselineDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DbManagementClient()

	return tfresource.ReadResource(sync)
}

type DatabaseManagementManagedDatabaseSqlPlanBaselineDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_database_management.DbManagementClient
	Res    *oci_database_management.GetSqlPlanBaselineResponse
}

func (s *DatabaseManagementManagedDatabaseSqlPlanBaselineDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DatabaseManagementManagedDatabaseSqlPlanBaselineDataSourceCrud) Get() error {
	request := oci_database_management.GetSqlPlanBaselineRequest{}

	if managedDatabaseId, ok := s.D.GetOkExists("managed_database_id"); ok {
		tmp := managedDatabaseId.(string)
		request.ManagedDatabaseId = &tmp
	}

	if opcNamedCredentialId, ok := s.D.GetOkExists("opc_named_credential_id"); ok {
		tmp := opcNamedCredentialId.(string)
		request.OpcNamedCredentialId = &tmp
	}

	if planName, ok := s.D.GetOkExists("plan_name"); ok {
		tmp := planName.(string)
		request.PlanName = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "database_management")

	response, err := s.Client.GetSqlPlanBaseline(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *DatabaseManagementManagedDatabaseSqlPlanBaselineDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DatabaseManagementManagedDatabaseSqlPlanBaselineDataSource-", DatabaseManagementManagedDatabaseSqlPlanBaselineDataSource(), s.D))

	s.D.Set("accepted", s.Res.Accepted)

	if s.Res.Action != nil {
		s.D.Set("action", *s.Res.Action)
	}

	s.D.Set("adaptive", s.Res.Adaptive)

	s.D.Set("auto_purge", s.Res.AutoPurge)

	s.D.Set("enabled", s.Res.Enabled)

	if s.Res.ExecutionPlan != nil {
		s.D.Set("execution_plan", *s.Res.ExecutionPlan)
	}

	s.D.Set("fixed", s.Res.Fixed)

	if s.Res.Module != nil {
		s.D.Set("module", *s.Res.Module)
	}

	s.D.Set("origin", s.Res.Origin)

	s.D.Set("reproduced", s.Res.Reproduced)

	if s.Res.SqlHandle != nil {
		s.D.Set("sql_handle", *s.Res.SqlHandle)
	}

	if s.Res.SqlText != nil {
		s.D.Set("sql_text", *s.Res.SqlText)
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeLastExecuted != nil {
		s.D.Set("time_last_executed", s.Res.TimeLastExecuted.String())
	}

	if s.Res.TimeLastModified != nil {
		s.D.Set("time_last_modified", s.Res.TimeLastModified.String())
	}

	return nil
}
