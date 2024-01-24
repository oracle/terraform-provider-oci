// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package data_safe

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_data_safe "github.com/oracle/oci-go-sdk/v65/datasafe"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DataSafeSqlFirewallViolationsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readDataSafeSqlFirewallViolations,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"access_level": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"compartment_id_in_subtree": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"scim_query": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"sql_firewall_violations_collection": {
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
									"client_ip": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"client_os_user_name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"client_program": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"compartment_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"current_db_user_name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"db_user_name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"operation": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"operation_time": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"sql_accessed_objects": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"sql_level": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"sql_text": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"target_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"target_name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"time_collected": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"violation_action": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"violation_cause": {
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

func readDataSafeSqlFirewallViolations(d *schema.ResourceData, m interface{}) error {
	sync := &DataSafeSqlFirewallViolationsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataSafeClient()

	return tfresource.ReadResource(sync)
}

type DataSafeSqlFirewallViolationsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_data_safe.DataSafeClient
	Res    *oci_data_safe.ListSqlFirewallViolationsResponse
}

func (s *DataSafeSqlFirewallViolationsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DataSafeSqlFirewallViolationsDataSourceCrud) Get() error {
	request := oci_data_safe.ListSqlFirewallViolationsRequest{}

	if accessLevel, ok := s.D.GetOkExists("access_level"); ok {
		request.AccessLevel = oci_data_safe.ListSqlFirewallViolationsAccessLevelEnum(accessLevel.(string))
	}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if compartmentIdInSubtree, ok := s.D.GetOkExists("compartment_id_in_subtree"); ok {
		tmp := compartmentIdInSubtree.(bool)
		request.CompartmentIdInSubtree = &tmp
	}

	if scimQuery, ok := s.D.GetOkExists("scim_query"); ok {
		tmp := scimQuery.(string)
		request.ScimQuery = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "data_safe")

	response, err := s.Client.ListSqlFirewallViolations(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListSqlFirewallViolations(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *DataSafeSqlFirewallViolationsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DataSafeSqlFirewallViolationsDataSource-", DataSafeSqlFirewallViolationsDataSource(), s.D))
	resources := []map[string]interface{}{}
	sqlFirewallViolation := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, SqlFirewallViolationSummaryToMap(item))
	}
	sqlFirewallViolation["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, DataSafeSqlFirewallViolationsDataSource().Schema["sql_firewall_violations_collection"].Elem.(*schema.Resource).Schema)
		sqlFirewallViolation["items"] = items
	}

	resources = append(resources, sqlFirewallViolation)
	if err := s.D.Set("sql_firewall_violations_collection", resources); err != nil {
		return err
	}

	return nil
}

func SqlFirewallViolationSummaryToMap(obj oci_data_safe.SqlFirewallViolationSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.ClientIp != nil {
		result["client_ip"] = string(*obj.ClientIp)
	}

	if obj.ClientOsUserName != nil {
		result["client_os_user_name"] = string(*obj.ClientOsUserName)
	}

	if obj.ClientProgram != nil {
		result["client_program"] = string(*obj.ClientProgram)
	}

	if obj.CompartmentId != nil {
		result["compartment_id"] = string(*obj.CompartmentId)
	}

	if obj.CurrentDbUserName != nil {
		result["current_db_user_name"] = string(*obj.CurrentDbUserName)
	}

	if obj.DbUserName != nil {
		result["db_user_name"] = string(*obj.DbUserName)
	}

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	if obj.Operation != nil {
		result["operation"] = string(*obj.Operation)
	}

	if obj.OperationTime != nil {
		result["operation_time"] = obj.OperationTime.String()
	}

	if obj.SqlAccessedObjects != nil {
		result["sql_accessed_objects"] = string(*obj.SqlAccessedObjects)
	}

	result["sql_level"] = string(obj.SqlLevel)

	if obj.SqlText != nil {
		result["sql_text"] = string(*obj.SqlText)
	}

	if obj.TargetId != nil {
		result["target_id"] = string(*obj.TargetId)
	}

	if obj.TargetName != nil {
		result["target_name"] = string(*obj.TargetName)
	}

	if obj.TimeCollected != nil {
		result["time_collected"] = obj.TimeCollected.String()
	}

	result["violation_action"] = string(obj.ViolationAction)

	if obj.ViolationCause != nil {
		result["violation_cause"] = string(*obj.ViolationCause)
	}

	return result
}
