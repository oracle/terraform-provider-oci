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

func DataSafeSecurityPolicyReportDatabaseTableAccessEntriesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readDataSafeSecurityPolicyReportDatabaseTableAccessEntries,
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
			"database_table_access_entry_collection": {
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
									"access_through_object": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"access_type": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"are_all_tables_accessible": {
										Type:     schema.TypeBool,
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
									"is_access_constrained_by_label_security": {
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
									"is_access_constrained_by_view": {
										Type:     schema.TypeBool,
										Computed: true,
									},
									"is_access_constrained_by_virtual_private_database": {
										Type:     schema.TypeBool,
										Computed: true,
									},
									"is_sensitive": {
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
								},
							},
						},
					},
				},
			},
		},
	}
}

func readDataSafeSecurityPolicyReportDatabaseTableAccessEntries(d *schema.ResourceData, m interface{}) error {
	sync := &DataSafeSecurityPolicyReportDatabaseTableAccessEntriesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataSafeClient()

	return tfresource.ReadResource(sync)
}

type DataSafeSecurityPolicyReportDatabaseTableAccessEntriesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_data_safe.DataSafeClient
	Res    *oci_data_safe.ListDatabaseTableAccessEntriesResponse
}

func (s *DataSafeSecurityPolicyReportDatabaseTableAccessEntriesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DataSafeSecurityPolicyReportDatabaseTableAccessEntriesDataSourceCrud) Get() error {
	request := oci_data_safe.ListDatabaseTableAccessEntriesRequest{}

	if scimQuery, ok := s.D.GetOkExists("scim_query"); ok {
		tmp := scimQuery.(string)
		request.ScimQuery = &tmp
	}

	if securityPolicyReportId, ok := s.D.GetOkExists("security_policy_report_id"); ok {
		tmp := securityPolicyReportId.(string)
		request.SecurityPolicyReportId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "data_safe")

	response, err := s.Client.ListDatabaseTableAccessEntries(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListDatabaseTableAccessEntries(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *DataSafeSecurityPolicyReportDatabaseTableAccessEntriesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DataSafeSecurityPolicyReportDatabaseTableAccessEntriesDataSource-", DataSafeSecurityPolicyReportDatabaseTableAccessEntriesDataSource(), s.D))
	resources := []map[string]interface{}{}
	securityPolicyReportDatabaseTableAccessEntry := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, DatabaseTableAccessEntrySummaryToMap(item))
	}
	securityPolicyReportDatabaseTableAccessEntry["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, DataSafeSecurityPolicyReportDatabaseTableAccessEntriesDataSource().Schema["database_table_access_entry_collection"].Elem.(*schema.Resource).Schema)
		securityPolicyReportDatabaseTableAccessEntry["items"] = items
	}

	resources = append(resources, securityPolicyReportDatabaseTableAccessEntry)
	if err := s.D.Set("database_table_access_entry_collection", resources); err != nil {
		return err
	}

	return nil
}

func DatabaseTableAccessEntrySummaryToMap(obj oci_data_safe.DatabaseTableAccessEntrySummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.AccessThroughObject != nil {
		result["access_through_object"] = string(*obj.AccessThroughObject)
	}

	result["access_type"] = string(obj.AccessType)

	if obj.AreAllTablesAccessible != nil {
		result["are_all_tables_accessible"] = bool(*obj.AreAllTablesAccessible)
	}

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

	if obj.IsAccessConstrainedByLabelSecurity != nil {
		result["is_access_constrained_by_label_security"] = bool(*obj.IsAccessConstrainedByLabelSecurity)
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

	if obj.IsAccessConstrainedByView != nil {
		result["is_access_constrained_by_view"] = bool(*obj.IsAccessConstrainedByView)
	}

	if obj.IsAccessConstrainedByVirtualPrivateDatabase != nil {
		result["is_access_constrained_by_virtual_private_database"] = bool(*obj.IsAccessConstrainedByVirtualPrivateDatabase)
	}

	if obj.IsSensitive != nil {
		result["is_sensitive"] = bool(*obj.IsSensitive)
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

	return result
}
