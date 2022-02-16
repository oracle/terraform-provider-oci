// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package database_migration

import (
	"context"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_database_migration "github.com/oracle/oci-go-sdk/v58/databasemigration"
)

func DatabaseMigrationJobAdvisorReportDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularDatabaseMigrationJobAdvisorReport,
		Schema: map[string]*schema.Schema{
			"job_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			// Computed
			"number_of_fatal": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"number_of_fatal_blockers": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"number_of_informational_results": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"number_of_warnings": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"report_location_details": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"location_in_source": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"object_storage_details": {
							Type:     schema.TypeList,
							Computed: true,
							MaxItems: 1,
							MinItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"bucket": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"namespace": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"object": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
					},
				},
			},
			"result": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func readSingularDatabaseMigrationJobAdvisorReport(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseMigrationJobAdvisorReportDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseMigrationClient()

	return tfresource.ReadResource(sync)
}

type DatabaseMigrationJobAdvisorReportDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_database_migration.DatabaseMigrationClient
	Res    *oci_database_migration.GetAdvisorReportResponse
}

func (s *DatabaseMigrationJobAdvisorReportDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DatabaseMigrationJobAdvisorReportDataSourceCrud) Get() error {
	request := oci_database_migration.GetAdvisorReportRequest{}

	if jobId, ok := s.D.GetOkExists("job_id"); ok {
		tmp := jobId.(string)
		request.JobId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "database_migration")

	response, err := s.Client.GetAdvisorReport(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *DatabaseMigrationJobAdvisorReportDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DatabaseMigrationJobAdvisorReportDataSource-", DatabaseMigrationJobAdvisorReportDataSource(), s.D))

	if s.Res.NumberOfFatal != nil {
		s.D.Set("number_of_fatal", *s.Res.NumberOfFatal)
	}

	if s.Res.NumberOfFatalBlockers != nil {
		s.D.Set("number_of_fatal_blockers", *s.Res.NumberOfFatalBlockers)
	}

	if s.Res.NumberOfInformationalResults != nil {
		s.D.Set("number_of_informational_results", *s.Res.NumberOfInformationalResults)
	}

	if s.Res.NumberOfWarnings != nil {
		s.D.Set("number_of_warnings", *s.Res.NumberOfWarnings)
	}

	if s.Res.ReportLocationDetails != nil {
		s.D.Set("report_location_details", []interface{}{AdvisorReportLocationDetailsToMap(s.Res.ReportLocationDetails)})
	} else {
		s.D.Set("report_location_details", nil)
	}

	s.D.Set("result", s.Res.Result)

	return nil
}

func AdvisorReportBucketDetailsToMap(obj *oci_database_migration.AdvisorReportBucketDetails) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.BucketName != nil {
		result["bucket"] = string(*obj.BucketName)
	}

	if obj.Namespace != nil {
		result["namespace"] = string(*obj.Namespace)
	}

	if obj.ObjectName != nil {
		result["object"] = string(*obj.ObjectName)
	}

	return result
}

func AdvisorReportLocationDetailsToMap(obj *oci_database_migration.AdvisorReportLocationDetails) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.LocationInSource != nil {
		result["location_in_source"] = string(*obj.LocationInSource)
	}

	if obj.ObjectStorageDetails != nil {
		result["object_storage_details"] = []interface{}{AdvisorReportBucketDetailsToMap(obj.ObjectStorageDetails)}
	}

	return result
}
