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

func DatabaseManagementManagedMySqlDatabaseHighAvailabilityMembersDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readDatabaseManagementManagedMySqlDatabaseHighAvailabilityMembers,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"managed_my_sql_database_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"managed_my_sql_database_high_availability_member_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"flow_control": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"group_auto_increment": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"group_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"member_host": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"member_port": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									"member_role": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"member_state": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"member_uuid": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"member_role": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"member_state": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"single_primary_mode": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"status_summary": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"channel_apply_errors": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional

												// Computed
												"apply_error": {
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															// Required

															// Optional

															// Computed
															"last_error_message": {
																Type:     schema.TypeString,
																Computed: true,
															},
															"last_error_number": {
																Type:     schema.TypeInt,
																Computed: true,
															},
															"time_last_error": {
																Type:     schema.TypeString,
																Computed: true,
															},
															"worker_errors": {
																Type:     schema.TypeList,
																Computed: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		// Required

																		// Optional

																		// Computed
																		"last_error_message": {
																			Type:     schema.TypeString,
																			Computed: true,
																		},
																		"last_error_number": {
																			Type:     schema.TypeInt,
																			Computed: true,
																		},
																		"time_last_error": {
																			Type:     schema.TypeString,
																			Computed: true,
																		},
																	},
																},
															},
														},
													},
												},
												"channel_name": {
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},
									"channel_fetch_errors": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional

												// Computed
												"channel_name": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"fetch_error": {
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															// Required

															// Optional

															// Computed
															"last_error_message": {
																Type:     schema.TypeString,
																Computed: true,
															},
															"last_error_number": {
																Type:     schema.TypeInt,
																Computed: true,
															},
															"time_last_error": {
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
							},
						},
						"transactions_in_gtid_executed": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"view_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func readDatabaseManagementManagedMySqlDatabaseHighAvailabilityMembers(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseManagementManagedMySqlDatabaseHighAvailabilityMembersDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ManagedMySqlDatabasesClient()

	return tfresource.ReadResource(sync)
}

type DatabaseManagementManagedMySqlDatabaseHighAvailabilityMembersDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_database_management.ManagedMySqlDatabasesClient
	Res    *oci_database_management.ListHighAvailabilityMembersResponse
}

func (s *DatabaseManagementManagedMySqlDatabaseHighAvailabilityMembersDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DatabaseManagementManagedMySqlDatabaseHighAvailabilityMembersDataSourceCrud) Get() error {
	request := oci_database_management.ListHighAvailabilityMembersRequest{}

	if managedMySqlDatabaseId, ok := s.D.GetOkExists("managed_my_sql_database_id"); ok {
		tmp := managedMySqlDatabaseId.(string)
		request.ManagedMySqlDatabaseId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "database_management")

	response, err := s.Client.ListHighAvailabilityMembers(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListHighAvailabilityMembers(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *DatabaseManagementManagedMySqlDatabaseHighAvailabilityMembersDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DatabaseManagementManagedMySqlDatabaseHighAvailabilityMembersDataSource-", DatabaseManagementManagedMySqlDatabaseHighAvailabilityMembersDataSource(), s.D))
	resources := []map[string]interface{}{}
	managedMySqlDatabaseHighAvailabilityMember := map[string]interface{}{}

	if s.Res.FlowControl != nil {
		managedMySqlDatabaseHighAvailabilityMember["flow_control"] = *s.Res.FlowControl
	}

	if s.Res.GroupAutoIncrement != nil {
		managedMySqlDatabaseHighAvailabilityMember["group_auto_increment"] = *s.Res.GroupAutoIncrement
	}

	if s.Res.GroupName != nil {
		managedMySqlDatabaseHighAvailabilityMember["group_name"] = *s.Res.GroupName
	}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, ManagedMySqlDatabaseHighAvailabilityMemberSummaryToMap(item))
	}
	managedMySqlDatabaseHighAvailabilityMember["items"] = items

	if s.Res.MemberRole != nil {
		managedMySqlDatabaseHighAvailabilityMember["member_role"] = *s.Res.MemberRole
	}

	if s.Res.MemberState != nil {
		managedMySqlDatabaseHighAvailabilityMember["member_state"] = *s.Res.MemberState
	}

	if s.Res.SinglePrimaryMode != nil {
		managedMySqlDatabaseHighAvailabilityMember["single_primary_mode"] = *s.Res.SinglePrimaryMode
	}

	if s.Res.StatusSummary != nil {
		managedMySqlDatabaseHighAvailabilityMember["status_summary"] = []interface{}{MySqlHighAvailabilityStatusSummaryToMap(s.Res.StatusSummary)}
	} else {
		managedMySqlDatabaseHighAvailabilityMember["status_summary"] = nil
	}

	if s.Res.TransactionsInGtidExecuted != nil {
		managedMySqlDatabaseHighAvailabilityMember["transactions_in_gtid_executed"] = strconv.FormatInt(*s.Res.TransactionsInGtidExecuted, 10)
	}

	if s.Res.ViewId != nil {
		managedMySqlDatabaseHighAvailabilityMember["view_id"] = *s.Res.ViewId
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, DatabaseManagementManagedMySqlDatabaseHighAvailabilityMembersDataSource().Schema["managed_my_sql_database_high_availability_member_collection"].Elem.(*schema.Resource).Schema)
		managedMySqlDatabaseHighAvailabilityMember["items"] = items
	}

	resources = append(resources, managedMySqlDatabaseHighAvailabilityMember)
	if err := s.D.Set("managed_my_sql_database_high_availability_member_collection", resources); err != nil {
		return err
	}

	return nil
}

func ManagedMySqlDatabaseHighAvailabilityMemberSummaryToMap(obj oci_database_management.ManagedMySqlDatabaseHighAvailabilityMemberSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.MemberHost != nil {
		result["member_host"] = string(*obj.MemberHost)
	}

	if obj.MemberPort != nil {
		result["member_port"] = int(*obj.MemberPort)
	}

	if obj.MemberRole != nil {
		result["member_role"] = string(*obj.MemberRole)
	}

	if obj.MemberState != nil {
		result["member_state"] = string(*obj.MemberState)
	}

	if obj.MemberUuid != nil {
		result["member_uuid"] = string(*obj.MemberUuid)
	}

	return result
}

func MySqlApplyErrorToMap(obj *oci_database_management.MySqlApplyError) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.LastErrorMessage != nil {
		result["last_error_message"] = string(*obj.LastErrorMessage)
	}

	if obj.LastErrorNumber != nil {
		result["last_error_number"] = int(*obj.LastErrorNumber)
	}

	if obj.TimeLastError != nil {
		result["time_last_error"] = obj.TimeLastError.String()
	}

	workerErrors := []interface{}{}
	for _, item := range obj.WorkerErrors {
		workerErrors = append(workerErrors, MySqlApplyErrorWorkerToMap(item))
	}
	result["worker_errors"] = workerErrors

	return result
}

func MySqlApplyErrorWorkerToMap(obj oci_database_management.MySqlApplyErrorWorker) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.LastErrorMessage != nil {
		result["last_error_message"] = string(*obj.LastErrorMessage)
	}

	if obj.LastErrorNumber != nil {
		result["last_error_number"] = int(*obj.LastErrorNumber)
	}

	if obj.TimeLastError != nil {
		result["time_last_error"] = obj.TimeLastError.String()
	}

	return result
}

func MySqlChannelApplyErrorToMap(obj oci_database_management.MySqlChannelApplyError) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.ApplyError != nil {
		result["apply_error"] = []interface{}{MySqlApplyErrorToMap(obj.ApplyError)}
	}

	if obj.ChannelName != nil {
		result["channel_name"] = string(*obj.ChannelName)
	}

	return result
}

func MySqlChannelFetchErrorToMap(obj oci_database_management.MySqlChannelFetchError) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.ChannelName != nil {
		result["channel_name"] = string(*obj.ChannelName)
	}

	if obj.FetchError != nil {
		result["fetch_error"] = []interface{}{MySqlFetchErrorToMap(obj.FetchError)}
	}

	return result
}

func MySqlFetchErrorToMap(obj *oci_database_management.MySqlFetchError) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.LastErrorMessage != nil {
		result["last_error_message"] = string(*obj.LastErrorMessage)
	}

	if obj.LastErrorNumber != nil {
		result["last_error_number"] = int(*obj.LastErrorNumber)
	}

	if obj.TimeLastError != nil {
		result["time_last_error"] = obj.TimeLastError.String()
	}

	return result
}

func MySqlHighAvailabilityStatusSummaryToMap(obj *oci_database_management.MySqlHighAvailabilityStatusSummary) map[string]interface{} {
	result := map[string]interface{}{}

	channelApplyErrors := []interface{}{}
	for _, item := range obj.ChannelApplyErrors {
		channelApplyErrors = append(channelApplyErrors, MySqlChannelApplyErrorToMap(item))
	}
	result["channel_apply_errors"] = channelApplyErrors

	channelFetchErrors := []interface{}{}
	for _, item := range obj.ChannelFetchErrors {
		channelFetchErrors = append(channelFetchErrors, MySqlChannelFetchErrorToMap(item))
	}
	result["channel_fetch_errors"] = channelFetchErrors

	return result
}
