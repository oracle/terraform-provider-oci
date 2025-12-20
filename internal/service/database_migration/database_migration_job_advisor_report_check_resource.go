// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package database_migration

import (
	"context"
	"fmt"
	"log"
	"net/url"
	"regexp"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	oci_database_migration "github.com/oracle/oci-go-sdk/v65/databasemigration"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DatabaseMigrationJobAdvisorReportCheckResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createDatabaseMigrationJobAdvisorReportCheck,
		Read:     readDatabaseMigrationJobAdvisorReportCheck,
		Update:   updateDatabaseMigrationJobAdvisorReportCheck,
		Delete:   deleteDatabaseMigrationJobAdvisorReportCheck,
		Schema: map[string]*schema.Schema{
			// Required
			"advisor_report_check_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"is_reviewed": {
				Type:     schema.TypeBool,
				Required: true,
			},
			"job_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

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
						"action": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"columns": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"display_name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"key": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"display_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"fixup_script_location": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"impact": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"is_exclusion_allowed": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"is_reviewed": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"issue": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"key": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"metadata": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"object_name_column": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"object_type_column": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"object_type_fixed": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"schema_owner_column": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"object_count": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"result_type": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
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
	}
}

func createDatabaseMigrationJobAdvisorReportCheck(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseMigrationJobAdvisorReportCheckResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseMigrationClient()

	return tfresource.CreateResource(d, sync)
}

func readDatabaseMigrationJobAdvisorReportCheck(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseMigrationJobAdvisorReportCheckResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseMigrationClient()

	return tfresource.ReadResource(sync)
}

func updateDatabaseMigrationJobAdvisorReportCheck(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseMigrationJobAdvisorReportCheckResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseMigrationClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteDatabaseMigrationJobAdvisorReportCheck(d *schema.ResourceData, m interface{}) error {
	return nil
}

type DatabaseMigrationJobAdvisorReportCheckResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_database_migration.DatabaseMigrationClient
	Res                    *oci_database_migration.AdvisorReportCheckCollection
	DisableNotFoundRetries bool
}

func (s *DatabaseMigrationJobAdvisorReportCheckResourceCrud) ID() string {
	return s.D.Id()
}

func (s *DatabaseMigrationJobAdvisorReportCheckResourceCrud) Create() error {
	request := oci_database_migration.UpdateAdvisorReportCheckRequest{}

	if advisorReportCheckId, ok := s.D.GetOkExists("advisor_report_check_id"); ok {
		tmp := advisorReportCheckId.(string)
		request.AdvisorReportCheckId = &tmp
	}

	if jobId, ok := s.D.GetOkExists("job_id"); ok {
		tmp := jobId.(string)
		request.JobId = &tmp
	}

	if isReviewed, ok := s.D.GetOkExists("is_reviewed"); ok {
		val := isReviewed.(bool)
		request.UpdateAdvisorReportCheck = oci_database_migration.UpdateAdvisorReportCheckDetails{
			IsReviewed: &val,
		}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database_migration")

	_, err := s.Client.UpdateAdvisorReportCheck(context.Background(), request)
	if err != nil {
		return err
	}

	if request.AdvisorReportCheckId != nil && request.JobId != nil {
		s.D.SetId(GetJobAdvisorReportCheckCompositeId(*request.AdvisorReportCheckId, *request.JobId))
	}

	return s.Get()
}

func (s *DatabaseMigrationJobAdvisorReportCheckResourceCrud) Get() error {
	request := oci_database_migration.ListAdvisorReportChecksRequest{}

	if jobId, ok := s.D.GetOkExists("job_id"); ok {
		tmp := jobId.(string)
		request.JobId = &tmp
	}

	advisorReportCheckId, jobId, err := parseJobAdvisorReportCheckCompositeId(s.D.Id())
	if err == nil {
		request.JobId = &jobId
	} else {
		log.Printf("[WARN] Get() unable to parse current ID: %s", s.D.Id())
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database_migration")

	response, err := s.Client.ListAdvisorReportChecks(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.AdvisorReportCheckCollection

	// Filter to the requested advisorReportCheckId if present in composite ID
	if advisorReportCheckId != "" {
		filtered := make([]oci_database_migration.AdvisorReportCheckSummary, 0, len(s.Res.Items))
		for _, it := range s.Res.Items {
			if it.Key != nil && *it.Key == advisorReportCheckId {
				filtered = append(filtered, it)
			}
		}
		s.Res.Items = filtered
	}

	return nil
}

func (s *DatabaseMigrationJobAdvisorReportCheckResourceCrud) Update() error {
	request := oci_database_migration.UpdateAdvisorReportCheckRequest{}

	if advisorReportCheckId, ok := s.D.GetOkExists("advisor_report_check_id"); ok {
		tmp := advisorReportCheckId.(string)
		request.AdvisorReportCheckId = &tmp
	}

	if jobId, ok := s.D.GetOkExists("job_id"); ok {
		tmp := jobId.(string)
		request.JobId = &tmp
	}

	if isReviewed, ok := s.D.GetOkExists("is_reviewed"); ok {
		val := isReviewed.(bool)
		request.UpdateAdvisorReportCheck = oci_database_migration.UpdateAdvisorReportCheckDetails{
			IsReviewed: &val,
		}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database_migration")

	_, err := s.Client.UpdateAdvisorReportCheck(context.Background(), request)
	if err != nil {
		return err
	}

	return s.Get()
}

func (s *DatabaseMigrationJobAdvisorReportCheckResourceCrud) SetData() error {

	_, jobId, err := parseJobAdvisorReportCheckCompositeId(s.D.Id())
	if err == nil {
		s.D.Set("job_id", jobId)
	} else {
		log.Printf("[WARN] SetData() unable to parse current ID: %s", s.D.Id())
	}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, AdvisorReportCheckSummaryToMap(item))
	}
	s.D.Set("items", items)

	if s.Res.Summary != nil {
		s.D.Set("summary", []interface{}{AdvisorCheckSummaryToMap(s.Res.Summary)})
	} else {
		s.D.Set("summary", nil)
	}

	return nil
}

func GetJobAdvisorReportCheckCompositeId(advisorReportCheckId string, jobId string) string {
	advisorReportCheckId = url.PathEscape(advisorReportCheckId)
	jobId = url.PathEscape(jobId)
	compositeId := "jobs/" + jobId + "/advisorReportChecks/" + advisorReportCheckId
	return compositeId
}

func parseJobAdvisorReportCheckCompositeId(compositeId string) (advisorReportCheckId string, jobId string, err error) {
	parts := strings.Split(compositeId, "/")
	match, _ := regexp.MatchString("jobs/.*/advisorReportChecks/.*", compositeId)
	if !match || len(parts) != 4 {
		err = fmt.Errorf("illegal compositeId %s encountered", compositeId)
		return
	}
	jobId, _ = url.PathUnescape(parts[1])
	advisorReportCheckId, _ = url.PathUnescape(parts[3])

	return
}

func AdvisorCheckSummaryToMap(obj *oci_database_migration.AdvisorCheckSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.BlockerResultsTotalCount != nil {
		result["blocker_results_total_count"] = int(*obj.BlockerResultsTotalCount)
	}

	if obj.FatalResultsTotalCount != nil {
		result["fatal_results_total_count"] = int(*obj.FatalResultsTotalCount)
	}

	if obj.InformationalResultsTotalCount != nil {
		result["informational_results_total_count"] = int(*obj.InformationalResultsTotalCount)
	}

	if obj.PassResultsTotalCount != nil {
		result["pass_results_total_count"] = int(*obj.PassResultsTotalCount)
	}

	if obj.WarningResultsTotalCount != nil {
		result["warning_results_total_count"] = int(*obj.WarningResultsTotalCount)
	}

	return result
}

func AdvisorReportCheckColumnToMap(obj oci_database_migration.AdvisorReportCheckColumn) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.DisplayName != nil {
		result["display_name"] = string(*obj.DisplayName)
	}

	if obj.Key != nil {
		result["key"] = string(*obj.Key)
	}

	return result
}

func AdvisorReportCheckSummaryToMap(obj oci_database_migration.AdvisorReportCheckSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Action != nil {
		result["action"] = string(*obj.Action)
	}

	columns := []interface{}{}
	for _, item := range obj.Columns {
		columns = append(columns, AdvisorReportCheckColumnToMap(item))
	}
	result["columns"] = columns

	if obj.DisplayName != nil {
		result["display_name"] = string(*obj.DisplayName)
	}

	if obj.FixupScriptLocation != nil {
		result["fixup_script_location"] = string(*obj.FixupScriptLocation)
	}

	if obj.Impact != nil {
		result["impact"] = string(*obj.Impact)
	}

	if obj.IsExclusionAllowed != nil {
		result["is_exclusion_allowed"] = bool(*obj.IsExclusionAllowed)
	}

	if obj.IsReviewed != nil {
		result["is_reviewed"] = bool(*obj.IsReviewed)
	}

	if obj.Issue != nil {
		result["issue"] = string(*obj.Issue)
	}

	if obj.Key != nil {
		result["key"] = string(*obj.Key)
	}

	if obj.Metadata != nil {
		result["metadata"] = []interface{}{ObjectMetadataToMap(obj.Metadata)}
	}

	if obj.ObjectCount != nil {
		result["object_count"] = int(*obj.ObjectCount)
	}

	result["result_type"] = string(obj.ResultType)

	return result
}

func ObjectMetadataToMap(obj *oci_database_migration.ObjectMetadata) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.ObjectNameColumn != nil {
		result["object_name_column"] = string(*obj.ObjectNameColumn)
	}

	if obj.ObjectTypeColumn != nil {
		result["object_type_column"] = string(*obj.ObjectTypeColumn)
	}

	if obj.ObjectTypeFixed != nil {
		result["object_type_fixed"] = string(*obj.ObjectTypeFixed)
	}

	if obj.SchemaOwnerColumn != nil {
		result["schema_owner_column"] = string(*obj.SchemaOwnerColumn)
	}

	return result
}
