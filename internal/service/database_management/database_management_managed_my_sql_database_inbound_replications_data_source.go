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

func DatabaseManagementManagedMySqlDatabaseInboundReplicationsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readDatabaseManagementManagedMySqlDatabaseInboundReplications,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"managed_my_sql_database_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"managed_my_sql_database_inbound_replication_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"inbound_replications_count": {
							Type:     schema.TypeInt,
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
									"applier_filters": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional

												// Computed
												"filter_name": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"filter_rule": {
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},
									"apply_delay": {
										Type:     schema.TypeFloat,
										Computed: true,
									},
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
									"apply_status": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"busy_workers": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									"channel_name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"desired_delay_seconds": {
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
									"fetch_status": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"gtid_assignment": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"relay_log_storage_space_used": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"remaining_delay_seconds": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"retrieved_gtid_set": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"seconds_behind_source": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"source_host": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"source_port": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									"source_server_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"source_uuid": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"transactions_received": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"parallel_workers": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"preserve_commit_order": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"replica_server_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"replica_uuid": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func readDatabaseManagementManagedMySqlDatabaseInboundReplications(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseManagementManagedMySqlDatabaseInboundReplicationsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ManagedMySqlDatabasesClient()

	return tfresource.ReadResource(sync)
}

type DatabaseManagementManagedMySqlDatabaseInboundReplicationsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_database_management.ManagedMySqlDatabasesClient
	Res    *oci_database_management.ListInboundReplicationsResponse
}

func (s *DatabaseManagementManagedMySqlDatabaseInboundReplicationsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DatabaseManagementManagedMySqlDatabaseInboundReplicationsDataSourceCrud) Get() error {
	request := oci_database_management.ListInboundReplicationsRequest{}

	if managedMySqlDatabaseId, ok := s.D.GetOkExists("managed_my_sql_database_id"); ok {
		tmp := managedMySqlDatabaseId.(string)
		request.ManagedMySqlDatabaseId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "database_management")

	response, err := s.Client.ListInboundReplications(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListInboundReplications(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *DatabaseManagementManagedMySqlDatabaseInboundReplicationsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DatabaseManagementManagedMySqlDatabaseInboundReplicationsDataSource-", DatabaseManagementManagedMySqlDatabaseInboundReplicationsDataSource(), s.D))
	resources := []map[string]interface{}{}
	managedMySqlDatabaseInboundReplication := map[string]interface{}{}

	if s.Res.InboundReplicationsCount != nil {
		managedMySqlDatabaseInboundReplication["inbound_replications_count"] = *s.Res.InboundReplicationsCount
	}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, ManagedMySqlDatabaseInboundReplicationSummaryToMap(item))
	}
	managedMySqlDatabaseInboundReplication["items"] = items

	if s.Res.ParallelWorkers != nil {
		managedMySqlDatabaseInboundReplication["parallel_workers"] = *s.Res.ParallelWorkers
	}

	if s.Res.PreserveCommitOrder != nil {
		managedMySqlDatabaseInboundReplication["preserve_commit_order"] = *s.Res.PreserveCommitOrder
	}

	if s.Res.ReplicaServerId != nil {
		managedMySqlDatabaseInboundReplication["replica_server_id"] = strconv.FormatInt(*s.Res.ReplicaServerId, 10)
	}

	if s.Res.ReplicaUuid != nil {
		managedMySqlDatabaseInboundReplication["replica_uuid"] = *s.Res.ReplicaUuid
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, DatabaseManagementManagedMySqlDatabaseInboundReplicationsDataSource().Schema["managed_my_sql_database_inbound_replication_collection"].Elem.(*schema.Resource).Schema)
		managedMySqlDatabaseInboundReplication["items"] = items
	}

	resources = append(resources, managedMySqlDatabaseInboundReplication)
	if err := s.D.Set("managed_my_sql_database_inbound_replication_collection", resources); err != nil {
		return err
	}

	return nil
}

func ManagedMySqlDatabaseInboundReplicationSummaryToMap(obj oci_database_management.ManagedMySqlDatabaseInboundReplicationSummary) map[string]interface{} {
	result := map[string]interface{}{}

	applierFilters := []interface{}{}
	for _, item := range obj.ApplierFilters {
		applierFilters = append(applierFilters, MySqlReplicationApplierFilterToMap(item))
	}
	result["applier_filters"] = applierFilters

	if obj.ApplyDelay != nil {
		result["apply_delay"] = float64(*obj.ApplyDelay)
	}

	if obj.ApplyError != nil {
		result["apply_error"] = []interface{}{DatabaseManagementMySqlApplyErrorToMap(obj.ApplyError)}
	}

	if obj.ApplyStatus != nil {
		result["apply_status"] = string(*obj.ApplyStatus)
	}

	if obj.BusyWorkers != nil {
		result["busy_workers"] = int(*obj.BusyWorkers)
	}

	if obj.ChannelName != nil {
		result["channel_name"] = string(*obj.ChannelName)
	}

	if obj.DesiredDelaySeconds != nil {
		result["desired_delay_seconds"] = strconv.FormatInt(*obj.DesiredDelaySeconds, 10)
	}

	if obj.FetchError != nil {
		result["fetch_error"] = []interface{}{DatabaseManagementMySqlFetchErrorToMap(obj.FetchError)}
	}

	if obj.FetchStatus != nil {
		result["fetch_status"] = string(*obj.FetchStatus)
	}

	if obj.GtidAssignment != nil {
		result["gtid_assignment"] = string(*obj.GtidAssignment)
	}

	if obj.RelayLogStorageSpaceUsed != nil {
		result["relay_log_storage_space_used"] = strconv.FormatInt(*obj.RelayLogStorageSpaceUsed, 10)
	}

	if obj.RemainingDelaySeconds != nil {
		result["remaining_delay_seconds"] = strconv.FormatInt(*obj.RemainingDelaySeconds, 10)
	}

	if obj.RetrievedGtidSet != nil {
		result["retrieved_gtid_set"] = string(*obj.RetrievedGtidSet)
	}

	if obj.SecondsBehindSource != nil {
		result["seconds_behind_source"] = strconv.FormatInt(*obj.SecondsBehindSource, 10)
	}

	if obj.SourceHost != nil {
		result["source_host"] = string(*obj.SourceHost)
	}

	if obj.SourcePort != nil {
		result["source_port"] = int(*obj.SourcePort)
	}

	if obj.SourceServerId != nil {
		result["source_server_id"] = strconv.FormatInt(*obj.SourceServerId, 10)
	}

	if obj.SourceUuid != nil {
		result["source_uuid"] = string(*obj.SourceUuid)
	}

	if obj.TransactionsReceived != nil {
		result["transactions_received"] = strconv.FormatInt(*obj.TransactionsReceived, 10)
	}

	return result
}

func DatabaseManagementMySqlApplyErrorToMap(obj *oci_database_management.MySqlApplyError) map[string]interface{} {
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
		workerErrors = append(workerErrors, DatabaseManagementMySqlApplyErrorWorkerToMap(item))
	}
	result["worker_errors"] = workerErrors

	return result
}

func DatabaseManagementMySqlApplyErrorWorkerToMap(obj oci_database_management.MySqlApplyErrorWorker) map[string]interface{} {
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

func DatabaseManagementMySqlFetchErrorToMap(obj *oci_database_management.MySqlFetchError) map[string]interface{} {
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

func MySqlReplicationApplierFilterToMap(obj oci_database_management.MySqlReplicationApplierFilter) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.FilterName != nil {
		result["filter_name"] = string(*obj.FilterName)
	}

	if obj.FilterRule != nil {
		result["filter_rule"] = string(*obj.FilterRule)
	}

	return result
}
