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

func DatabaseMigrationJobAdvisorReportCheckObjectsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readDatabaseMigrationJobAdvisorReportCheckObjects,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"advisor_report_check_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"job_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"advisor_report_check_objects_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"defined_tags": {
										Type:     schema.TypeMap,
										Computed: true,
										Elem:     schema.TypeString,
									},
									"fields": {
										Type:     schema.TypeMap,
										Computed: true,
										Elem:     schema.TypeString,
									},
									"freeform_tags": {
										Type:     schema.TypeMap,
										Computed: true,
										Elem:     schema.TypeString,
									},
									"is_excluded": {
										Type:     schema.TypeBool,
										Computed: true,
									},
									"system_tags": {
										Type:     schema.TypeMap,
										Computed: true,
										Elem:     schema.TypeString,
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

func readDatabaseMigrationJobAdvisorReportCheckObjects(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseMigrationJobAdvisorReportCheckObjectsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseMigrationClient()

	return tfresource.ReadResource(sync)
}

type DatabaseMigrationJobAdvisorReportCheckObjectsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_database_migration.DatabaseMigrationClient
	Res    *oci_database_migration.ListAdvisorReportCheckObjectsResponse
}

func (s *DatabaseMigrationJobAdvisorReportCheckObjectsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DatabaseMigrationJobAdvisorReportCheckObjectsDataSourceCrud) Get() error {
	request := oci_database_migration.ListAdvisorReportCheckObjectsRequest{}

	if advisorReportCheckId, ok := s.D.GetOkExists("advisor_report_check_id"); ok {
		tmp := advisorReportCheckId.(string)
		request.AdvisorReportCheckId = &tmp
	}

	if jobId, ok := s.D.GetOkExists("job_id"); ok {
		tmp := jobId.(string)
		request.JobId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "database_migration")

	response, err := s.Client.ListAdvisorReportCheckObjects(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListAdvisorReportCheckObjects(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *DatabaseMigrationJobAdvisorReportCheckObjectsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DatabaseMigrationJobAdvisorReportCheckObjectsDataSource-", DatabaseMigrationJobAdvisorReportCheckObjectsDataSource(), s.D))
	resources := []map[string]interface{}{}
	jobAdvisorReportCheckObject := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, AdvisorReportCheckObjectSummaryToMap(item))
	}
	jobAdvisorReportCheckObject["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, DatabaseMigrationJobAdvisorReportCheckObjectsDataSource().Schema["advisor_report_check_objects_collection"].Elem.(*schema.Resource).Schema)
		jobAdvisorReportCheckObject["items"] = items
	}

	resources = append(resources, jobAdvisorReportCheckObject)
	if err := s.D.Set("advisor_report_check_objects_collection", resources); err != nil {
		return err
	}

	return nil
}

func AdvisorReportCheckObjectSummaryToMap(obj oci_database_migration.AdvisorReportCheckObjectSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.DefinedTags != nil {
		result["defined_tags"] = tfresource.DefinedTagsToMap(obj.DefinedTags)
	}

	result["fields"] = obj.Fields

	result["freeform_tags"] = obj.FreeformTags

	if obj.IsExcluded != nil {
		result["is_excluded"] = bool(*obj.IsExcluded)
	}

	if obj.SystemTags != nil {
		result["system_tags"] = tfresource.SystemTagsToMap(obj.SystemTags)
	}

	return result
}
