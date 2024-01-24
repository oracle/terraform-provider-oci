// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package data_safe

import (
	"context"
	"fmt"
	"log"
	"net/url"
	"regexp"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"

	oci_data_safe "github.com/oracle/oci-go-sdk/v65/datasafe"
)

func DataSafeDiscoveryJobsResultResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createDataSafeDiscoveryJobsResult,
		Read:     readDataSafeDiscoveryJobsResult,
		Delete:   deleteDataSafeDiscoveryJobsResult,
		Schema: map[string]*schema.Schema{
			// Required
			"discovery_job_id": {
				Type:       schema.TypeString,
				Required:   true,
				ForceNew:   true,
				Deprecated: tfresource.ResourceDeprecated("oci_data_safe_discovery_jobs_result"),
			},
			// Optional

			// Computed
			"app_defined_child_column_keys": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"app_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"column_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"data_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"db_defined_child_column_keys": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"discovery_type": {
				Type:       schema.TypeString,
				Computed:   true,
				Deprecated: tfresource.ResourceDeprecated("oci_data_safe_discovery_jobs_result"),
			},
			"estimated_data_value_count": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"is_result_applied": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"key": {
				Type:       schema.TypeString,
				Computed:   true,
				Deprecated: tfresource.ResourceDeprecated("oci_data_safe_discovery_jobs_result"),
			},
			"modified_attributes": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"app_defined_child_column_keys": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"db_defined_child_column_keys": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
					},
				},
			},
			"object": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"object_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"parent_column_keys": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"planned_action": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"relation_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"sample_data_values": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"schema_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"sensitive_columnkey": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"sensitive_type_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func createDataSafeDiscoveryJobsResult(d *schema.ResourceData, m interface{}) error {
	sync := &DataSafeDiscoveryJobsResultResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataSafeClient()

	return tfresource.CreateResource(d, sync)
}

func readDataSafeDiscoveryJobsResult(d *schema.ResourceData, m interface{}) error {
	sync := &DataSafeDiscoveryJobsResultResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataSafeClient()

	return tfresource.ReadResource(sync)
}

func deleteDataSafeDiscoveryJobsResult(d *schema.ResourceData, m interface{}) error {
	sync := &DataSafeDiscoveryJobsResultResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataSafeClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type DataSafeDiscoveryJobsResultResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_data_safe.DataSafeClient
	Res                    *oci_data_safe.DiscoveryJobResult
	DisableNotFoundRetries bool
}

func (s *DataSafeDiscoveryJobsResultResourceCrud) ID() string {
	return s.D.Id()
}

func (s *DataSafeDiscoveryJobsResultResourceCrud) Create() error {
	request := oci_data_safe.ListDiscoveryJobResultsRequest{}

	if discoveryJobId, ok := s.D.GetOkExists("discovery_job_id"); ok {
		tmp := discoveryJobId.(string)
		request.DiscoveryJobId = &tmp
	}
	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "data_safe")

	response, err := s.Client.ListDiscoveryJobResults(context.Background(), request)
	if err != nil {
		return err
	}

	listResponse := response.Items
	resultKey := listResponse[0].Key
	s.D.Set("key", &resultKey)

	url := "discoveryJobs/" + *request.DiscoveryJobId + "/results/" + *resultKey
	s.D.SetId(url)

	return s.Get()
}

func (s *DataSafeDiscoveryJobsResultResourceCrud) Get() error {
	request := oci_data_safe.GetDiscoveryJobResultRequest{}

	if discoveryJobId, ok := s.D.GetOkExists("discovery_job_id"); ok {
		tmp := discoveryJobId.(string)
		request.DiscoveryJobId = &tmp
	}

	if resultKey, ok := s.D.GetOkExists("key"); ok {
		tmp := resultKey.(string)
		request.ResultKey = &tmp
	}

	discoveryJobId, resultKey, err := parseDiscoveryJobsResultCompositeId(s.D.Id())
	if err == nil {
		request.DiscoveryJobId = &discoveryJobId
		request.ResultKey = &resultKey
	} else {
		log.Printf("[WARN] Get() unable to parse current ID: %s", s.D.Id())
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "data_safe")

	response, err := s.Client.GetDiscoveryJobResult(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.DiscoveryJobResult
	return nil
}

func (s *DataSafeDiscoveryJobsResultResourceCrud) Delete() error {
	request := oci_data_safe.DeleteDiscoveryJobResultRequest{}

	if discoveryJobId, ok := s.D.GetOkExists("discovery_job_id"); ok {
		tmp := discoveryJobId.(string)
		request.DiscoveryJobId = &tmp
	}

	if resultKey, ok := s.D.GetOkExists("key"); ok {
		tmp := resultKey.(string)
		request.ResultKey = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "data_safe")

	_, err := s.Client.DeleteDiscoveryJobResult(context.Background(), request)
	return err
}

func (s *DataSafeDiscoveryJobsResultResourceCrud) SetData() error {

	discoveryJobId, resultKey, err := parseDiscoveryJobsResultCompositeId(s.D.Id())
	if err == nil {
		s.D.Set("discovery_job_id", &discoveryJobId)
		s.D.Set("result_key", &resultKey)
	} else {
		log.Printf("[WARN] SetData() unable to parse current ID: %s", s.D.Id())
	}

	s.D.Set("app_defined_child_column_keys", s.Res.AppDefinedChildColumnKeys)

	if s.Res.AppName != nil {
		s.D.Set("app_name", *s.Res.AppName)
	}

	if s.Res.ColumnName != nil {
		s.D.Set("column_name", *s.Res.ColumnName)
	}

	if s.Res.DataType != nil {
		s.D.Set("data_type", *s.Res.DataType)
	}

	s.D.Set("db_defined_child_column_keys", s.Res.DbDefinedChildColumnKeys)

	if s.Res.DiscoveryJobId != nil {
		s.D.Set("discovery_job_id", *s.Res.DiscoveryJobId)
	}

	s.D.Set("discovery_type", s.Res.DiscoveryType)

	if s.Res.EstimatedDataValueCount != nil {
		s.D.Set("estimated_data_value_count", strconv.FormatInt(*s.Res.EstimatedDataValueCount, 10))
	}

	if s.Res.IsResultApplied != nil {
		s.D.Set("is_result_applied", *s.Res.IsResultApplied)
	}

	if s.Res.Key != nil {
		s.D.Set("key", *s.Res.Key)
	}

	if s.Res.ModifiedAttributes != nil {
		s.D.Set("modified_attributes", []interface{}{ModifiedAttributesToMap(s.Res.ModifiedAttributes)})
	} else {
		s.D.Set("modified_attributes", nil)
	}

	if s.Res.ObjectName != nil {
		s.D.Set("object", *s.Res.ObjectName)
	}

	s.D.Set("object_type", s.Res.ObjectType)

	s.D.Set("parent_column_keys", s.Res.ParentColumnKeys)

	s.D.Set("planned_action", s.Res.PlannedAction)

	s.D.Set("relation_type", s.Res.RelationType)

	s.D.Set("sample_data_values", s.Res.SampleDataValues)

	if s.Res.SchemaName != nil {
		s.D.Set("schema_name", *s.Res.SchemaName)
	}

	if s.Res.SensitiveColumnkey != nil {
		s.D.Set("sensitive_columnkey", *s.Res.SensitiveColumnkey)
	}

	if s.Res.SensitiveTypeId != nil {
		s.D.Set("sensitive_type_id", *s.Res.SensitiveTypeId)
	}

	return nil
}

func GetDiscoveryJobsResultCompositeId(discoveryJobId string, resultKey string) string {
	discoveryJobId = url.PathEscape(discoveryJobId)
	resultKey = url.PathEscape(resultKey)
	compositeId := "discoveryJobs/" + discoveryJobId + "/results/" + resultKey
	return compositeId
}

func parseDiscoveryJobsResultCompositeId(compositeId string) (discoveryJobId string, resultKey string, err error) {
	parts := strings.Split(compositeId, "/")
	match, _ := regexp.MatchString("discoveryJobs/.*/results/.*", compositeId)
	if !match || len(parts) != 4 {
		err = fmt.Errorf("illegal compositeId %s encountered", compositeId)
		return
	}
	discoveryJobId, _ = url.PathUnescape(parts[1])
	resultKey, _ = url.PathUnescape(parts[3])

	return
}

func DiscoveryJobResultSummaryToMap(obj oci_data_safe.DiscoveryJobResultSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.ColumnName != nil {
		result["column_name"] = string(*obj.ColumnName)
	}

	if obj.DataType != nil {
		result["data_type"] = string(*obj.DataType)
	}

	if obj.DiscoveryJobId != nil {
		result["discovery_job_id"] = string(*obj.DiscoveryJobId)
	}

	result["discovery_type"] = string(obj.DiscoveryType)

	if obj.EstimatedDataValueCount != nil {
		result["estimated_data_value_count"] = strconv.FormatInt(*obj.EstimatedDataValueCount, 10)
	}

	if obj.IsResultApplied != nil {
		result["is_result_applied"] = bool(*obj.IsResultApplied)
	}

	if obj.Key != nil {
		result["key"] = string(*obj.Key)
	}

	if obj.ObjectName != nil {
		result["object"] = string(*obj.ObjectName)
	}

	result["object_type"] = string(obj.ObjectType)

	result["parent_column_keys"] = obj.ParentColumnKeys

	result["planned_action"] = string(obj.PlannedAction)

	result["relation_type"] = string(obj.RelationType)

	result["sample_data_values"] = obj.SampleDataValues

	if obj.SchemaName != nil {
		result["schema_name"] = string(*obj.SchemaName)
	}

	if obj.SensitiveColumnkey != nil {
		result["sensitive_columnkey"] = string(*obj.SensitiveColumnkey)
	}

	if obj.SensitiveTypeId != nil {
		result["sensitive_type_id"] = string(*obj.SensitiveTypeId)
	}

	return result
}

func ModifiedAttributesToMap(obj *oci_data_safe.ModifiedAttributes) map[string]interface{} {
	result := map[string]interface{}{}

	result["app_defined_child_column_keys"] = obj.AppDefinedChildColumnKeys

	result["db_defined_child_column_keys"] = obj.DbDefinedChildColumnKeys

	return result
}
