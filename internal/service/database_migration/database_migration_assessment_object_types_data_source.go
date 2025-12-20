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

func DatabaseMigrationAssessmentObjectTypesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readDatabaseMigrationAssessmentObjectTypes,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"connection_type": {
				Type:     schema.TypeString,
				Required: true,
			},
			"assessment_object_type_summary_collection": {
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
									"name": {
										Type:     schema.TypeString,
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

func readDatabaseMigrationAssessmentObjectTypes(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseMigrationAssessmentObjectTypesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseMigrationClient()

	return tfresource.ReadResource(sync)
}

type DatabaseMigrationAssessmentObjectTypesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_database_migration.DatabaseMigrationClient
	Res    *oci_database_migration.ListAssessmentObjectTypesResponse
}

func (s *DatabaseMigrationAssessmentObjectTypesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DatabaseMigrationAssessmentObjectTypesDataSourceCrud) Get() error {
	request := oci_database_migration.ListAssessmentObjectTypesRequest{}

	if connectionType, ok := s.D.GetOkExists("connection_type"); ok {
		request.ConnectionType = oci_database_migration.ListAssessmentObjectTypesConnectionTypeEnum(connectionType.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "database_migration")

	response, err := s.Client.ListAssessmentObjectTypes(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListAssessmentObjectTypes(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *DatabaseMigrationAssessmentObjectTypesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DatabaseMigrationAssessmentObjectTypesDataSource-", DatabaseMigrationAssessmentObjectTypesDataSource(), s.D))
	resources := []map[string]interface{}{}
	assessmentObjectType := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, AssessmentObjectTypeSummaryToMap(item))
	}
	assessmentObjectType["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, DatabaseMigrationAssessmentObjectTypesDataSource().Schema["assessment_object_type_summary_collection"].Elem.(*schema.Resource).Schema)
		assessmentObjectType["items"] = items
	}

	resources = append(resources, assessmentObjectType)
	if err := s.D.Set("assessment_object_type_summary_collection", resources); err != nil {
		return err
	}

	return nil
}

func AssessmentObjectTypeSummaryToMap(obj oci_database_migration.AssessmentObjectTypeSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	return result
}
