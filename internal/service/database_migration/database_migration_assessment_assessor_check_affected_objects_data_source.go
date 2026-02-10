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

func DatabaseMigrationAssessmentAssessorCheckAffectedObjectsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readDatabaseMigrationAssessmentAssessorCheckAffectedObjects,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"assessment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"assessor_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"check_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"affected_objects_collection": {
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

func readDatabaseMigrationAssessmentAssessorCheckAffectedObjects(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseMigrationAssessmentAssessorCheckAffectedObjectsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseMigrationClient()

	return tfresource.ReadResource(sync)
}

type DatabaseMigrationAssessmentAssessorCheckAffectedObjectsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_database_migration.DatabaseMigrationClient
	Res    *oci_database_migration.ListAffectedObjectsResponse
}

func (s *DatabaseMigrationAssessmentAssessorCheckAffectedObjectsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DatabaseMigrationAssessmentAssessorCheckAffectedObjectsDataSourceCrud) Get() error {
	request := oci_database_migration.ListAffectedObjectsRequest{}

	if assessmentId, ok := s.D.GetOkExists("assessment_id"); ok {
		tmp := assessmentId.(string)
		request.AssessmentId = &tmp
	}

	if assessorName, ok := s.D.GetOkExists("assessor_name"); ok {
		tmp := assessorName.(string)
		request.AssessorName = &tmp
	}

	if checkName, ok := s.D.GetOkExists("check_name"); ok {
		tmp := checkName.(string)
		request.CheckName = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "database_migration")

	response, err := s.Client.ListAffectedObjects(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListAffectedObjects(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *DatabaseMigrationAssessmentAssessorCheckAffectedObjectsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DatabaseMigrationAssessmentAssessorCheckAffectedObjectsDataSource-", DatabaseMigrationAssessmentAssessorCheckAffectedObjectsDataSource(), s.D))
	resources := []map[string]interface{}{}
	assessmentAssessorCheckAffectedObject := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, AdvisorReportCheckObjectSummaryToMap(item))
	}
	assessmentAssessorCheckAffectedObject["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, DatabaseMigrationAssessmentAssessorCheckAffectedObjectsDataSource().Schema["affected_objects_collection"].Elem.(*schema.Resource).Schema)
		assessmentAssessorCheckAffectedObject["items"] = items
	}

	resources = append(resources, assessmentAssessorCheckAffectedObject)
	if err := s.D.Set("affected_objects_collection", resources); err != nil {
		return err
	}

	return nil
}
