// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package data_safe

import (
	"context"
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_common "github.com/oracle/oci-go-sdk/v65/common"
	oci_data_safe "github.com/oracle/oci-go-sdk/v65/datasafe"
	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DataSafeDatabaseSecurityConfigManagementResource() *schema.Resource {
	return &schema.Resource{
		Timeouts: tfresource.DefaultTimeout,
		Create:   createDataSafeDatabaseSecurityConfigManagement,
		Read:     readDataSafeDatabaseSecurityConfigManagement,
		Update:   updateDataSafeDatabaseSecurityConfigManagement,
		Delete:   deleteDataSafeDatabaseSecurityConfigManagement,
		Schema: map[string]*schema.Schema{
			// Required

			// Optional
			"compartment_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"defined_tags": {
				Type:             schema.TypeMap,
				Optional:         true,
				Computed:         true,
				DiffSuppressFunc: tfresource.DefinedTagsDiffSuppressFunction,
				Elem:             schema.TypeString,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"freeform_tags": {
				Type:     schema.TypeMap,
				Optional: true,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"target_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"sql_firewall_config": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional
						"exclude_job": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"status": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"violation_log_auto_purge": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},

						// Computed
						"time_status_updated": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"refresh_trigger": {
				Type:     schema.TypeBool,
				Optional: true,
			},

			// Computed
			"lifecycle_details": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"system_tags": {
				Type:     schema.TypeMap,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"time_created": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_last_refreshed": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_updated": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func createDataSafeDatabaseSecurityConfigManagement(d *schema.ResourceData, m interface{}) error {
	sync := &DataSafeDatabaseSecurityConfigManagementResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataSafeClient()

	err := sync.GetDatabaseSecurityConfigWorkReq()
	if err != nil {
		return err
	}

	err = sync.Get()
	if err != nil {
		return err
	}

	return updateDataSafeDatabaseSecurityConfigManagement(d, m)
}

func readDataSafeDatabaseSecurityConfigManagement(d *schema.ResourceData, m interface{}) error {
	sync := &DataSafeDatabaseSecurityConfigResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataSafeClient()

	return tfresource.ReadResource(sync)
}

func updateDataSafeDatabaseSecurityConfigManagement(d *schema.ResourceData, m interface{}) error {
	sync := &DataSafeDatabaseSecurityConfigManagementResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataSafeClient()

	if refreshTrigger, ok := sync.D.GetOkExists("refresh_trigger"); ok {
		if refreshTrigger == true {
			err := sync.RefreshDatabaseSecurityConfiguration()

			if err != nil {
				return err
			}
		}
	}

	return tfresource.UpdateResource(d, sync)
}

func deleteDataSafeDatabaseSecurityConfigManagement(d *schema.ResourceData, m interface{}) error {
	return nil
}

type DataSafeDatabaseSecurityConfigManagementResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_data_safe.DataSafeClient
	Res                    *oci_data_safe.DatabaseSecurityConfig
	DisableNotFoundRetries bool
}

func (s *DataSafeDatabaseSecurityConfigManagementResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *DataSafeDatabaseSecurityConfigManagementResourceCrud) RefreshDatabaseSecurityConfiguration() error {
	request := oci_data_safe.RefreshDatabaseSecurityConfigurationRequest{}

	idTmp := s.D.Id()
	request.DatabaseSecurityConfigId = &idTmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "data_safe")

	response, err := s.Client.RefreshDatabaseSecurityConfiguration(context.Background(), request)
	if err != nil {
		return err
	}

	if waitErr := tfresource.WaitForUpdatedState(s.D, s); waitErr != nil {
		return waitErr
	}

	val := s.D.Get("refresh_trigger")
	s.D.Set("refresh_trigger", val)

	workId := response.OpcWorkRequestId
	return s.getDatabaseSecurityConfigFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "data_safe"), oci_data_safe.WorkRequestResourceActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}

func (s *DataSafeDatabaseSecurityConfigManagementResourceCrud) GetDatabaseSecurityConfigWorkReq() error {
	listWorkRequestsRequest := oci_data_safe.ListWorkRequestsRequest{SortBy: oci_data_safe.ListWorkRequestsSortByEnum("ACCEPTEDTIME"), SortOrder: oci_data_safe.ListWorkRequestsSortOrderEnum("DESC")}
	var workId *string
	tmp := "CREATE_DB_SECURITY_CONFIG"
	listWorkRequestsRequest.OperationType = &tmp

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		listWorkRequestsRequest.CompartmentId = &tmp
	}

	if targetId, ok := s.D.GetOkExists("target_id"); ok {
		tmp := targetId.(string)
		listWorkRequestsRequest.TargetDatabaseId = &tmp
	}

	listWorkRequestsResponse, err := s.Client.ListWorkRequests(context.Background(), listWorkRequestsRequest)
	if listWorkRequestsResponse.Items != nil && len(listWorkRequestsResponse.Items) > 0 {
		var items = &listWorkRequestsResponse.Items[0]
		workId = items.Id
	}

	if err != nil {
		return err
	}

	if workId != nil {
		return s.getDatabaseSecurityConfigFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "data_safe"), oci_data_safe.WorkRequestResourceActionTypeCreated, s.D.Timeout(schema.TimeoutUpdate))
	} else {
		return s.GetDatabaseSecurityConfigList()
	}
}

func (s *DataSafeDatabaseSecurityConfigManagementResourceCrud) GetDatabaseSecurityConfigList() error {
	request := oci_data_safe.ListDatabaseSecurityConfigsRequest{}
	var databaseSecurityConfig = new(oci_data_safe.DatabaseSecurityConfig)

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if targetId, ok := s.D.GetOkExists("target_id"); ok {
		tmp := targetId.(string)
		request.TargetId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "data_safe")

	response, err := s.Client.ListDatabaseSecurityConfigs(context.Background(), request)
	if err != nil {
		return err
	}

	if response.DatabaseSecurityConfigCollection.Items != nil && len(response.DatabaseSecurityConfigCollection.Items) > 0 {
		tmp := &response.DatabaseSecurityConfigCollection.Items[0]
		databaseSecurityConfig.Id = tmp.Id
	}

	if databaseSecurityConfig.Id == nil {
		return nil
	}

	s.D.SetId(*databaseSecurityConfig.Id)
	return nil
}

func (s *DataSafeDatabaseSecurityConfigManagementResourceCrud) Get() error {
	request := oci_data_safe.GetDatabaseSecurityConfigRequest{}

	tmp := s.D.Id()
	request.DatabaseSecurityConfigId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "data_safe")

	response, err := s.Client.GetDatabaseSecurityConfig(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.DatabaseSecurityConfig
	return nil
}

func (s *DataSafeDatabaseSecurityConfigManagementResourceCrud) Update() error {
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(compartment)
			if err != nil {
				return err
			}
		}
	}
	request := oci_data_safe.UpdateDatabaseSecurityConfigRequest{}

	tmp := s.D.Id()
	request.DatabaseSecurityConfigId = &tmp

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
		convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return err
		}
		request.DefinedTags = convertedDefinedTags
	}

	if description, ok := s.D.GetOkExists("description"); ok {
		tmp := description.(string)
		request.Description = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if sqlFirewallConfig, ok := s.D.GetOkExists("sql_firewall_config"); ok {
		if tmpList := sqlFirewallConfig.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "sql_firewall_config", 0)
			tmp, err := s.mapToUpdateSqlFirewallConfigDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.SqlFirewallConfig = &tmp
		}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "data_safe")

	response, err := s.Client.UpdateDatabaseSecurityConfig(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getDatabaseSecurityConfigFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "data_safe"), oci_data_safe.WorkRequestResourceActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}

func (s *DataSafeDatabaseSecurityConfigManagementResourceCrud) SetData() error {
	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.Description != nil {
		s.D.Set("description", *s.Res.Description)
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)
	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	if s.Res.SqlFirewallConfig != nil {
		s.D.Set("sql_firewall_config", []interface{}{SqlFirewallConfigToMap(s.Res.SqlFirewallConfig)})
	} else {
		s.D.Set("sql_firewall_config", nil)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.SystemTags != nil {
		s.D.Set("system_tags", tfresource.SystemTagsToMap(s.Res.SystemTags))
	}

	if s.Res.TargetId != nil {
		s.D.Set("target_id", *s.Res.TargetId)
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeLastRefreshed != nil {
		s.D.Set("time_last_refreshed", s.Res.TimeLastRefreshed.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	return nil
}

func (s *DataSafeDatabaseSecurityConfigManagementResourceCrud) mapToUpdateSqlFirewallConfigDetails(fieldKeyFormat string) (oci_data_safe.UpdateSqlFirewallConfigDetails, error) {
	result := oci_data_safe.UpdateSqlFirewallConfigDetails{}

	if excludeJob, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "exclude_job")); ok {
		result.ExcludeJob = oci_data_safe.UpdateSqlFirewallConfigDetailsExcludeJobEnum(excludeJob.(string))
	}

	if status, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "status")); ok {
		result.Status = oci_data_safe.UpdateSqlFirewallConfigDetailsStatusEnum(status.(string))
	}

	if violationLogAutoPurge, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "violation_log_auto_purge")); ok {
		result.ViolationLogAutoPurge = oci_data_safe.UpdateSqlFirewallConfigDetailsViolationLogAutoPurgeEnum(violationLogAutoPurge.(string))
	}

	return result, nil
}

func (s *DataSafeDatabaseSecurityConfigManagementResourceCrud) getDatabaseSecurityConfigFromWorkRequest(workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_data_safe.WorkRequestResourceActionTypeEnum, timeout time.Duration) error {

	// Wait until it finishes
	databaseSecurityConfigId, err := databaseSecurityConfigWaitForWorkRequest(workId, "databasesecurityconfig",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.Client)

	if err != nil {
		return err
	}

	s.D.SetId(*databaseSecurityConfigId)

	return s.Get()
}

func (s *DataSafeDatabaseSecurityConfigManagementResourceCrud) updateCompartment(compartment interface{}) error {
	changeCompartmentRequest := oci_data_safe.ChangeDatabaseSecurityConfigCompartmentRequest{}

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	idTmp := s.D.Id()
	changeCompartmentRequest.DatabaseSecurityConfigId = &idTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "data_safe")

	response, err := s.Client.ChangeDatabaseSecurityConfigCompartment(context.Background(), changeCompartmentRequest)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getDatabaseSecurityConfigFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "data_safe"), oci_data_safe.WorkRequestResourceActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}
