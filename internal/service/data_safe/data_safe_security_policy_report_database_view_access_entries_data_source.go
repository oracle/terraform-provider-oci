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

func DataSafeSecurityPolicyReportDatabaseViewAccessEntriesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readDataSafeSecurityPolicyReportDatabaseViewAccessEntries,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"scim_query": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"security_policy_report_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"target_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"database_view_access_entry_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"access_type": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"column_name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"grant_from_role": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"grantee": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"grantor": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"is_access_constrained_by_database_vault": {
										Type:     schema.TypeBool,
										Computed: true,
									},
									"is_access_constrained_by_real_application_security": {
										Type:     schema.TypeBool,
										Computed: true,
									},
									"is_access_constrained_by_redaction": {
										Type:     schema.TypeBool,
										Computed: true,
									},
									"is_access_constrained_by_sql_firewall": {
										Type:     schema.TypeBool,
										Computed: true,
									},
									"is_access_constrained_by_virtual_private_database": {
										Type:     schema.TypeBool,
										Computed: true,
									},
									"key": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"privilege": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"privilege_grantable": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"privilege_type": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"table_name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"table_schema": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"target_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"view_name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"view_schema": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"view_text": {
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

func readDataSafeSecurityPolicyReportDatabaseViewAccessEntries(d *schema.ResourceData, m interface{}) error {
	sync := &DataSafeSecurityPolicyReportDatabaseViewAccessEntriesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataSafeClient()

	return tfresource.ReadResource(sync)
}

type DataSafeSecurityPolicyReportDatabaseViewAccessEntriesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_data_safe.DataSafeClient
	Res    *oci_data_safe.ListDatabaseViewAccessEntriesResponse
}

func (s *DataSafeSecurityPolicyReportDatabaseViewAccessEntriesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DataSafeSecurityPolicyReportDatabaseViewAccessEntriesDataSourceCrud) Get() error {
	request := oci_data_safe.ListDatabaseViewAccessEntriesRequest{}

	if scimQuery, ok := s.D.GetOkExists("scim_query"); ok {
		tmp := scimQuery.(string)
		request.ScimQuery = &tmp
	}

	if securityPolicyReportId, ok := s.D.GetOkExists("security_policy_report_id"); ok {
		tmp := securityPolicyReportId.(string)
		request.SecurityPolicyReportId = &tmp
	}

	if targetId, ok := s.D.GetOkExists("target_id"); ok {
		tmp := targetId.(string)
		request.TargetId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "data_safe")

	response, err := s.Client.ListDatabaseViewAccessEntries(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListDatabaseViewAccessEntries(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *DataSafeSecurityPolicyReportDatabaseViewAccessEntriesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DataSafeSecurityPolicyReportDatabaseViewAccessEntriesDataSource-", DataSafeSecurityPolicyReportDatabaseViewAccessEntriesDataSource(), s.D))
	resources := []map[string]interface{}{}
	securityPolicyReportDatabaseViewAccessEntry := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, DatabaseViewAccessEntrySummaryToMap(item))
	}
	securityPolicyReportDatabaseViewAccessEntry["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, DataSafeSecurityPolicyReportDatabaseViewAccessEntriesDataSource().Schema["database_view_access_entry_collection"].Elem.(*schema.Resource).Schema)
		securityPolicyReportDatabaseViewAccessEntry["items"] = items
	}

	resources = append(resources, securityPolicyReportDatabaseViewAccessEntry)
	if err := s.D.Set("database_view_access_entry_collection", resources); err != nil {
		return err
	}

	return nil
}

func DatabaseViewAccessEntrySummaryToMap(obj oci_data_safe.DatabaseViewAccessEntrySummary) map[string]interface{} {
	result := map[string]interface{}{}

	result["access_type"] = string(obj.AccessType)

	if obj.ColumnName != nil {
		result["column_name"] = string(*obj.ColumnName)
	}

	if obj.GrantFromRole != nil {
		result["grant_from_role"] = string(*obj.GrantFromRole)
	}

	if obj.Grantee != nil {
		result["grantee"] = string(*obj.Grantee)
	}

	if obj.Grantor != nil {
		result["grantor"] = string(*obj.Grantor)
	}

	if obj.IsAccessConstrainedByDatabaseVault != nil {
		result["is_access_constrained_by_database_vault"] = bool(*obj.IsAccessConstrainedByDatabaseVault)
	}

	if obj.IsAccessConstrainedByRealApplicationSecurity != nil {
		result["is_access_constrained_by_real_application_security"] = bool(*obj.IsAccessConstrainedByRealApplicationSecurity)
	}

	if obj.IsAccessConstrainedByRedaction != nil {
		result["is_access_constrained_by_redaction"] = bool(*obj.IsAccessConstrainedByRedaction)
	}

	if obj.IsAccessConstrainedBySqlFirewall != nil {
		result["is_access_constrained_by_sql_firewall"] = bool(*obj.IsAccessConstrainedBySqlFirewall)
	}

	if obj.IsAccessConstrainedByVirtualPrivateDatabase != nil {
		result["is_access_constrained_by_virtual_private_database"] = bool(*obj.IsAccessConstrainedByVirtualPrivateDatabase)
	}

	if obj.Key != nil {
		result["key"] = string(*obj.Key)
	}

	result["privilege"] = string(obj.Privilege)

	result["privilege_grantable"] = string(obj.PrivilegeGrantable)

	if obj.PrivilegeType != nil {
		result["privilege_type"] = string(*obj.PrivilegeType)
	}

	if obj.TableName != nil {
		result["table_name"] = string(*obj.TableName)
	}

	if obj.TableSchema != nil {
		result["table_schema"] = string(*obj.TableSchema)
	}

	if obj.TargetId != nil {
		result["target_id"] = string(*obj.TargetId)
	}

	if obj.ViewName != nil {
		result["view_name"] = string(*obj.ViewName)
	}

	if obj.ViewSchema != nil {
		result["view_schema"] = string(*obj.ViewSchema)
	}

	if obj.ViewText != nil {
		result["view_text"] = string(*obj.ViewText)
	}

	return result
}
