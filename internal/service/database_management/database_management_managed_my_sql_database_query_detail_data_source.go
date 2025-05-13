// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package database_management

import (
	"context"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_database_management "github.com/oracle/oci-go-sdk/v65/databasemanagement"
	"github.com/oracle/terraform-provider-oci/internal/client"

	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DatabaseManagementManagedMySqlDatabaseQueryDetailDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularDatabaseManagementManagedMySqlDatabaseQueryDetail,
		Schema: map[string]*schema.Schema{
			"digest": {
				Type:     schema.TypeString,
				Required: true,
			},
			"managed_my_sql_database_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			// Computed
			"query_explain_plan": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"json_explain": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"json_explain_version": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"query_messages": {
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
			"query_sample_details": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"execution_time": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"host": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"mysql_instance": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"query_sample_text": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"thread_id": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"time_query_sample_seen": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"user": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func readSingularDatabaseManagementManagedMySqlDatabaseQueryDetail(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseManagementManagedMySqlDatabaseQueryDetailDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ManagedMySqlDatabasesClient()

	return tfresource.ReadResource(sync)
}

type DatabaseManagementManagedMySqlDatabaseQueryDetailDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_database_management.ManagedMySqlDatabasesClient
	Res    *oci_database_management.GetMySqlQueryDetailsResponse
}

func (s *DatabaseManagementManagedMySqlDatabaseQueryDetailDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DatabaseManagementManagedMySqlDatabaseQueryDetailDataSourceCrud) Get() error {
	request := oci_database_management.GetMySqlQueryDetailsRequest{}

	if digest, ok := s.D.GetOkExists("digest"); ok {
		tmp := digest.(string)
		request.Digest = &tmp
	}

	if managedMySqlDatabaseId, ok := s.D.GetOkExists("managed_my_sql_database_id"); ok {
		tmp := managedMySqlDatabaseId.(string)
		request.ManagedMySqlDatabaseId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "database_management")

	response, err := s.Client.GetMySqlQueryDetails(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *DatabaseManagementManagedMySqlDatabaseQueryDetailDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DatabaseManagementManagedMySqlDatabaseQueryDetailDataSource-", DatabaseManagementManagedMySqlDatabaseQueryDetailDataSource(), s.D))

	if s.Res.QueryExplainPlan != nil {
		s.D.Set("query_explain_plan", []interface{}{MySqlQueryExplainPlanToMap(s.Res.QueryExplainPlan)})
	} else {
		s.D.Set("query_explain_plan", nil)
	}

	queryMessages := []interface{}{}
	for _, item := range s.Res.QueryMessages {
		queryMessages = append(queryMessages, DatabasemanagementMySqlQueryMessageToMap(item))
	}
	s.D.Set("query_messages", queryMessages)

	if s.Res.QuerySampleDetails != nil {
		s.D.Set("query_sample_details", []interface{}{MySqlQuerySampleDetailsToMap(s.Res.QuerySampleDetails)})
	} else {
		s.D.Set("query_sample_details", nil)
	}

	return nil
}

func MySqlQueryExplainPlanToMap(obj *oci_database_management.MySqlQueryExplainPlan) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.JsonExplain != nil {
		result["json_explain"] = string(*obj.JsonExplain)
	}

	result["json_explain_version"] = string(obj.JsonExplainVersion)

	return result
}

func DatabasemanagementMySqlQueryMessageToMap(obj oci_database_management.MySqlQueryMessage) map[string]interface{} {
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

func MySqlQuerySampleDetailsToMap(obj *oci_database_management.MySqlQuerySampleDetails) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.ExecutionTime != nil {
		result["execution_time"] = strconv.FormatInt(*obj.ExecutionTime, 10)
	}

	if obj.Host != nil {
		result["host"] = string(*obj.Host)
	}

	if obj.MysqlInstance != nil {
		result["mysql_instance"] = string(*obj.MysqlInstance)
	}

	if obj.QuerySampleText != nil {
		result["query_sample_text"] = string(*obj.QuerySampleText)
	}

	if obj.ThreadId != nil {
		result["thread_id"] = int(*obj.ThreadId)
	}

	if obj.TimeQuerySampleSeen != nil {
		result["time_query_sample_seen"] = obj.TimeQuerySampleSeen.String()
	}

	if obj.User != nil {
		result["user"] = string(*obj.User)
	}

	return result
}
