// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package database_migration

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_database_migration "github.com/oracle/oci-go-sdk/v65/databasemigration"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DatabaseMigrationJobAdvisorReportChecksDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readDatabaseMigrationJobAdvisorReportChecks,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"job_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"advisor_report_check_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem:     DatabaseMigrationJobAdvisorReportCheckResource(),
						},

						"summary": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"blocker_results_total_count": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									"fatal_results_total_count": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									"informational_results_total_count": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									"pass_results_total_count": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									"warning_results_total_count": {
										Type:     schema.TypeInt,
										Computed: true,
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func readDatabaseMigrationJobAdvisorReportChecks(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseMigrationJobAdvisorReportChecksDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseMigrationClient()

	return tfresource.ReadResource(sync)
}

type DatabaseMigrationJobAdvisorReportChecksDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_database_migration.DatabaseMigrationClient
	Res    *oci_database_migration.ListAdvisorReportChecksResponse
}

func (s *DatabaseMigrationJobAdvisorReportChecksDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DatabaseMigrationJobAdvisorReportChecksDataSourceCrud) Get() error {
	request := oci_database_migration.ListAdvisorReportChecksRequest{}

	if jobId, ok := s.D.GetOkExists("job_id"); ok {
		tmp := jobId.(string)
		request.JobId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "database_migration")

	response, err := s.Client.ListAdvisorReportChecks(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response

	return nil
}

func (s *DatabaseMigrationJobAdvisorReportChecksDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DatabaseMigrationJobAdvisorReportChecksDataSource-", DatabaseMigrationJobAdvisorReportChecksDataSource(), s.D))
	resources := []map[string]interface{}{}
	jobAdvisorReportCheck := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, AdvisorReportCheckSummaryToMap(item))
	}
	jobAdvisorReportCheck["items"] = items

	if s.Res.Summary != nil {
		jobAdvisorReportCheck["summary"] = []interface{}{AdvisorCheckSummaryToMap(s.Res.Summary)}
	} else {
		jobAdvisorReportCheck["summary"] = nil
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, DatabaseMigrationJobAdvisorReportChecksDataSource().Schema["advisor_report_check_collection"].Elem.(*schema.Resource).Schema)
		jobAdvisorReportCheck["items"] = items
	}

	resources = append(resources, jobAdvisorReportCheck)
	if err := s.D.Set("advisor_report_check_collection", resources); err != nil {
		return err
	}

	return nil
}
