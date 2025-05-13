// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package database_management

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_database_management "github.com/oracle/oci-go-sdk/v65/databasemanagement"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DatabaseManagementManagedMySqlDatabaseDigestErrorsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readDatabaseManagementManagedMySqlDatabaseDigestErrors,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"digest": {
				Type:     schema.TypeString,
				Required: true,
			},
			"managed_my_sql_database_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"my_sql_digest_errors_collection": {
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
									"error": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional

												// Computed
												"code": {
													Type:     schema.TypeInt,
													Computed: true,
												},
												"level": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"message_text": {
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},
									"occurrence_count": {
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

func readDatabaseManagementManagedMySqlDatabaseDigestErrors(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseManagementManagedMySqlDatabaseDigestErrorsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ManagedMySqlDatabasesClient()

	return tfresource.ReadResource(sync)
}

type DatabaseManagementManagedMySqlDatabaseDigestErrorsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_database_management.ManagedMySqlDatabasesClient
	Res    *oci_database_management.ListMySqlDigestErrorsResponse
}

func (s *DatabaseManagementManagedMySqlDatabaseDigestErrorsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DatabaseManagementManagedMySqlDatabaseDigestErrorsDataSourceCrud) Get() error {
	request := oci_database_management.ListMySqlDigestErrorsRequest{}

	if digest, ok := s.D.GetOkExists("digest"); ok {
		tmp := digest.(string)
		request.Digest = &tmp
	}

	if managedMySqlDatabaseId, ok := s.D.GetOkExists("managed_my_sql_database_id"); ok {
		tmp := managedMySqlDatabaseId.(string)
		request.ManagedMySqlDatabaseId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "database_management")

	response, err := s.Client.ListMySqlDigestErrors(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListMySqlDigestErrors(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *DatabaseManagementManagedMySqlDatabaseDigestErrorsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DatabaseManagementManagedMySqlDatabaseDigestErrorsDataSource-", DatabaseManagementManagedMySqlDatabaseDigestErrorsDataSource(), s.D))
	resources := []map[string]interface{}{}
	managedMySqlDatabaseDigestError := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, MySqlDigestErrorSummaryToMap(item))
	}
	managedMySqlDatabaseDigestError["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, DatabaseManagementManagedMySqlDatabaseDigestErrorsDataSource().Schema["my_sql_digest_errors_collection"].Elem.(*schema.Resource).Schema)
		managedMySqlDatabaseDigestError["items"] = items
	}

	resources = append(resources, managedMySqlDatabaseDigestError)
	if err := s.D.Set("my_sql_digest_errors_collection", resources); err != nil {
		return err
	}

	return nil
}

func MySqlDigestErrorSummaryToMap(obj oci_database_management.MySqlDigestErrorSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Error != nil {
		result["error"] = []interface{}{MySqlQueryMessageToMap(obj.Error)}
	}

	if obj.OccurrenceCount != nil {
		result["occurrence_count"] = int(*obj.OccurrenceCount)
	}

	return result
}

func MySqlQueryMessageToMap(obj *oci_database_management.MySqlQueryMessage) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Code != nil {
		result["code"] = int(*obj.Code)
	}

	result["level"] = string(obj.Level)

	if obj.MessageText != nil {
		result["message_text"] = string(*obj.MessageText)
	}

	return result
}
