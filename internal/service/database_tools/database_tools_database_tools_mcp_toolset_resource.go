// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package database_tools

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/retry"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	oci_common "github.com/oracle/oci-go-sdk/v65/common"
	oci_database_tools "github.com/oracle/oci-go-sdk/v65/databasetools"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DatabaseToolsDatabaseToolsMcpToolsetResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts:      tfresource.DefaultTimeout,
		CreateContext: createDatabaseToolsDatabaseToolsMcpToolsetWithContext,
		ReadContext:   readDatabaseToolsDatabaseToolsMcpToolsetWithContext,
		UpdateContext: updateDatabaseToolsDatabaseToolsMcpToolsetWithContext,
		DeleteContext: deleteDatabaseToolsDatabaseToolsMcpToolsetWithContext,
		Schema: map[string]*schema.Schema{
			// Required
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"database_tools_mcp_server_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"type": {
				Type:             schema.TypeString,
				Required:         true,
				DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
				ValidateFunc: validation.StringInSlice([]string{
					"BUILT_IN_SQL_TOOLS",
					"CUSTOMIZABLE_REPORTING_TOOLS",
					"CUSTOM_SQL_TOOL",
					"GENAI_SQL_ASSISTANT",
				}, true),
			},
			"version": {
				Type:     schema.TypeInt,
				Required: true,
			},

			// Optional
			"allowed_roles": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"default_execution_type": {
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
			"freeform_tags": {
				Type:     schema.TypeMap,
				Optional: true,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"generative_ai_semantic_store_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"locks": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				ForceNew: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"type": {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},

						// Optional
						"message": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},
						"related_resource_id": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},
						"time_created": {
							Type:             schema.TypeString,
							Optional:         true,
							Computed:         true,
							ForceNew:         true,
							DiffSuppressFunc: tfresource.TimeDiffSuppressFunction,
						},

						// Computed
					},
				},
			},
			"reports": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional
						"allowed_roles": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"database_tools_sql_report_id": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},

						// Computed
					},
				},
			},
			"source": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional
						"type": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"value": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},

						// Computed
					},
				},
			},
			"tool_description": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"tool_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"tools": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional
						"allowed_roles": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"name": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"status": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},

						// Computed
						"display_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"variables": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional
						"description": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"name": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"type": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},

						// Computed
					},
				},
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
			"time_updated": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func createDatabaseToolsDatabaseToolsMcpToolsetWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &DatabaseToolsDatabaseToolsMcpToolsetResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseToolsClient()

	return tfresource.HandleDiagError(m, tfresource.CreateResourceWithContext(ctx, d, sync))
}

func readDatabaseToolsDatabaseToolsMcpToolsetWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &DatabaseToolsDatabaseToolsMcpToolsetResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseToolsClient()

	return tfresource.HandleDiagError(m, tfresource.ReadResourceWithContext(ctx, sync))
}

func updateDatabaseToolsDatabaseToolsMcpToolsetWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &DatabaseToolsDatabaseToolsMcpToolsetResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseToolsClient()

	return tfresource.HandleDiagError(m, tfresource.UpdateResourceWithContext(ctx, d, sync))
}

func deleteDatabaseToolsDatabaseToolsMcpToolsetWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &DatabaseToolsDatabaseToolsMcpToolsetResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseToolsClient()
	sync.DisableNotFoundRetries = true

	return tfresource.HandleDiagError(m, tfresource.DeleteResourceWithContext(ctx, d, sync))
}

type DatabaseToolsDatabaseToolsMcpToolsetResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_database_tools.DatabaseToolsClient
	Res                    *oci_database_tools.DatabaseToolsMcpToolset
	DisableNotFoundRetries bool
}

func (s *DatabaseToolsDatabaseToolsMcpToolsetResourceCrud) ID() string {
	databaseToolsMcpToolset := *s.Res
	return *databaseToolsMcpToolset.GetId()
}

func (s *DatabaseToolsDatabaseToolsMcpToolsetResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_database_tools.DatabaseToolsMcpToolsetLifecycleStateCreating),
	}
}

func (s *DatabaseToolsDatabaseToolsMcpToolsetResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_database_tools.DatabaseToolsMcpToolsetLifecycleStateActive),
	}
}

func (s *DatabaseToolsDatabaseToolsMcpToolsetResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_database_tools.DatabaseToolsMcpToolsetLifecycleStateDeleting),
	}
}

func (s *DatabaseToolsDatabaseToolsMcpToolsetResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_database_tools.DatabaseToolsMcpToolsetLifecycleStateDeleted),
	}
}

func (s *DatabaseToolsDatabaseToolsMcpToolsetResourceCrud) CreateWithContext(ctx context.Context) error {
	request := oci_database_tools.CreateDatabaseToolsMcpToolsetRequest{}
	err := s.populateTopLevelPolymorphicCreateDatabaseToolsMcpToolsetRequest(&request)
	if err != nil {
		return err
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database_tools")

	response, err := s.Client.CreateDatabaseToolsMcpToolset(ctx, request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	var identifier *string
	identifier = response.GetId()
	if identifier != nil {
		s.D.SetId(*identifier)
	}
	return s.getDatabaseToolsMcpToolsetFromWorkRequest(ctx, workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database_tools"), oci_database_tools.ActionTypeCreated, s.D.Timeout(schema.TimeoutCreate))
}

func (s *DatabaseToolsDatabaseToolsMcpToolsetResourceCrud) getDatabaseToolsMcpToolsetFromWorkRequest(ctx context.Context, workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_database_tools.ActionTypeEnum, timeout time.Duration) error {

	// Wait until it finishes
	databaseToolsMcpToolsetId, err := databaseToolsMcpToolsetWaitForWorkRequest(ctx, workId, "databasetoolsmcptoolset",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.Client)

	if err != nil {
		return err
	}
	s.D.SetId(*databaseToolsMcpToolsetId)

	return s.GetWithContext(ctx)
}

func databaseToolsMcpToolsetWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
	startTime := time.Now()
	stopTime := startTime.Add(timeout)
	return func(response oci_common.OCIOperationResponse) bool {

		// Stop after timeout has elapsed
		if time.Now().After(stopTime) {
			return false
		}

		// Make sure we stop on default rules
		if tfresource.ShouldRetry(response, false, "database_tools", startTime) {
			return true
		}

		// Only stop if the time Finished is set
		if workRequestResponse, ok := response.Response.(oci_database_tools.GetWorkRequestResponse); ok {
			return workRequestResponse.TimeFinished == nil
		}
		return false
	}
}

func databaseToolsMcpToolsetWaitForWorkRequest(ctx context.Context, wId *string, entityType string, action oci_database_tools.ActionTypeEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_database_tools.DatabaseToolsClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "database_tools")
	retryPolicy.ShouldRetryOperation = databaseToolsMcpToolsetWorkRequestShouldRetryFunc(timeout)

	response := oci_database_tools.GetWorkRequestResponse{}
	stateConf := &retry.StateChangeConf{
		Pending: []string{
			string(oci_database_tools.OperationStatusInProgress),
			string(oci_database_tools.OperationStatusAccepted),
			string(oci_database_tools.OperationStatusCanceling),
		},
		Target: []string{
			string(oci_database_tools.OperationStatusSucceeded),
			string(oci_database_tools.OperationStatusFailed),
			string(oci_database_tools.OperationStatusCanceled),
		},
		Refresh: func() (interface{}, string, error) {
			var err error
			response, err = client.GetWorkRequest(ctx,
				oci_database_tools.GetWorkRequestRequest{
					WorkRequestId: wId,
					RequestMetadata: oci_common.RequestMetadata{
						RetryPolicy: retryPolicy,
					},
				})
			wr := &response.WorkRequest
			return wr, string(wr.Status), err
		},
		Timeout: timeout,
	}
	if _, e := stateConf.WaitForState(); e != nil {
		return nil, e
	}

	var identifier *string
	// The work request response contains an array of objects that finished the operation
	for _, res := range response.Resources {
		if strings.Contains(strings.ToLower(*res.EntityType), entityType) {
			if res.ActionType == action {
				identifier = res.Identifier
				break
			}
		}
	}

	// The workrequest may have failed, check for errors if identifier is not found or work failed or got cancelled
	if identifier == nil || response.Status == oci_database_tools.OperationStatusFailed || response.Status == oci_database_tools.OperationStatusCanceled {
		return nil, getErrorFromDatabaseToolsDatabaseToolsMcpToolsetWorkRequest(ctx, client, wId, retryPolicy, entityType, action)
	}

	return identifier, nil
}

func getErrorFromDatabaseToolsDatabaseToolsMcpToolsetWorkRequest(ctx context.Context, client *oci_database_tools.DatabaseToolsClient, workId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_database_tools.ActionTypeEnum) error {
	response, err := client.ListWorkRequestErrors(ctx,
		oci_database_tools.ListWorkRequestErrorsRequest{
			WorkRequestId: workId,
			RequestMetadata: oci_common.RequestMetadata{
				RetryPolicy: retryPolicy,
			},
		})
	if err != nil {
		return err
	}

	allErrs := make([]string, 0)
	for _, wrkErr := range response.Items {
		allErrs = append(allErrs, *wrkErr.Message)
	}
	errorMessage := strings.Join(allErrs, "\n")

	workRequestErr := fmt.Errorf("work request did not succeed, workId: %s, entity: %s, action: %s. Message: %s", *workId, entityType, action, errorMessage)

	return workRequestErr
}

func (s *DatabaseToolsDatabaseToolsMcpToolsetResourceCrud) GetWithContext(ctx context.Context) error {
	request := oci_database_tools.GetDatabaseToolsMcpToolsetRequest{}

	tmp := s.D.Id()
	request.DatabaseToolsMcpToolsetId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database_tools")

	response, err := s.Client.GetDatabaseToolsMcpToolset(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response.DatabaseToolsMcpToolset
	return nil
}

func (s *DatabaseToolsDatabaseToolsMcpToolsetResourceCrud) UpdateWithContext(ctx context.Context) error {
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(ctx, compartment)
			if err != nil {
				return err
			}
		}
	}
	request := oci_database_tools.UpdateDatabaseToolsMcpToolsetRequest{}
	err := s.populateTopLevelPolymorphicUpdateDatabaseToolsMcpToolsetRequest(&request)
	if err != nil {
		return err
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database_tools")

	response, err := s.Client.UpdateDatabaseToolsMcpToolset(ctx, request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getDatabaseToolsMcpToolsetFromWorkRequest(ctx, workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database_tools"), oci_database_tools.ActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}

func (s *DatabaseToolsDatabaseToolsMcpToolsetResourceCrud) DeleteWithContext(ctx context.Context) error {
	request := oci_database_tools.DeleteDatabaseToolsMcpToolsetRequest{}

	tmp := s.D.Id()
	request.DatabaseToolsMcpToolsetId = &tmp

	if isLockOverride, ok := s.D.GetOkExists("is_lock_override"); ok {
		tmp := isLockOverride.(bool)
		request.IsLockOverride = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database_tools")

	response, err := s.Client.DeleteDatabaseToolsMcpToolset(ctx, request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	// Wait until it finishes
	_, delWorkRequestErr := databaseToolsMcpToolsetWaitForWorkRequest(ctx, workId, "databasetoolsmcptoolset",
		oci_database_tools.ActionTypeDeleted, s.D.Timeout(schema.TimeoutDelete), s.DisableNotFoundRetries, s.Client)
	return delWorkRequestErr
}

func (s *DatabaseToolsDatabaseToolsMcpToolsetResourceCrud) SetData() error {
	switch v := (*s.Res).(type) {
	case oci_database_tools.DatabaseToolsMcpToolsetBuiltInSqlTools:
		s.D.Set("type", "BUILT_IN_SQL_TOOLS")

		s.D.Set("allowed_roles", nil)

		s.D.Set("default_execution_type", v.DefaultExecutionType)

		tools := []interface{}{}
		for _, item := range v.Tools {
			tools = append(tools, DatabaseToolsMcpToolsetToolDetailsToMap(item))
		}
		s.D.Set("tools", tools)

		if v.CompartmentId != nil {
			s.D.Set("compartment_id", *v.CompartmentId)
		}

		if v.DatabaseToolsMcpServerId != nil {
			s.D.Set("database_tools_mcp_server_id", *v.DatabaseToolsMcpServerId)
		}

		if v.DefinedTags != nil {
			s.D.Set("defined_tags", tfresource.DefinedTagsToMap(v.DefinedTags))
		}

		if v.Description != nil {
			s.D.Set("description", *v.Description)
		}

		if v.DisplayName != nil {
			s.D.Set("display_name", *v.DisplayName)
		}

		s.D.Set("freeform_tags", v.FreeformTags)

		if v.Id != nil {
			s.D.SetId(*v.Id)
		}

		if v.LifecycleDetails != nil {
			s.D.Set("lifecycle_details", *v.LifecycleDetails)
		}

		locks := []interface{}{}
		for _, item := range v.Locks {
			locks = append(locks, DbtoolsMcpToolsetResourceLockToMap(item))
		}
		s.D.Set("locks", locks)

		s.D.Set("state", v.LifecycleState)

		if v.SystemTags != nil {
			s.D.Set("system_tags", tfresource.SystemTagsToMap(v.SystemTags))
		}

		if v.TimeCreated != nil {
			s.D.Set("time_created", v.TimeCreated.String())
		}

		if v.TimeUpdated != nil {
			s.D.Set("time_updated", v.TimeUpdated.String())
		}

		if v.Version != nil {
			s.D.Set("version", *v.Version)
		}
	case oci_database_tools.DatabaseToolsMcpToolsetCustomizableReportingTools:
		s.D.Set("type", "CUSTOMIZABLE_REPORTING_TOOLS")

		s.D.Set("allowed_roles", nil)

		s.D.Set("default_execution_type", v.DefaultExecutionType)

		reports := []interface{}{}
		for _, item := range v.Reports {
			reports = append(reports, DatabaseToolsMcpToolsetCustomizableReportingToolsReportToMap(item))
		}
		s.D.Set("reports", reports)

		tools := []interface{}{}
		for _, item := range v.Tools {
			tools = append(tools, DatabaseToolsMcpToolsetToolDetailsToMap(item))
		}
		s.D.Set("tools", tools)

		if v.CompartmentId != nil {
			s.D.Set("compartment_id", *v.CompartmentId)
		}

		if v.DatabaseToolsMcpServerId != nil {
			s.D.Set("database_tools_mcp_server_id", *v.DatabaseToolsMcpServerId)
		}

		if v.DefinedTags != nil {
			s.D.Set("defined_tags", tfresource.DefinedTagsToMap(v.DefinedTags))
		}

		if v.Description != nil {
			s.D.Set("description", *v.Description)
		}

		if v.DisplayName != nil {
			s.D.Set("display_name", *v.DisplayName)
		}

		s.D.Set("freeform_tags", v.FreeformTags)

		if v.Id != nil {
			s.D.SetId(*v.Id)
		}

		if v.LifecycleDetails != nil {
			s.D.Set("lifecycle_details", *v.LifecycleDetails)
		}

		locks := []interface{}{}
		for _, item := range v.Locks {
			locks = append(locks, DbtoolsMcpToolsetResourceLockToMap(item))
		}
		s.D.Set("locks", locks)

		s.D.Set("state", v.LifecycleState)

		if v.SystemTags != nil {
			s.D.Set("system_tags", tfresource.SystemTagsToMap(v.SystemTags))
		}

		if v.TimeCreated != nil {
			s.D.Set("time_created", v.TimeCreated.String())
		}

		if v.TimeUpdated != nil {
			s.D.Set("time_updated", v.TimeUpdated.String())
		}

		if v.Version != nil {
			s.D.Set("version", *v.Version)
		}
	case oci_database_tools.DatabaseToolsMcpToolsetCustomSqlTool:
		s.D.Set("type", "CUSTOM_SQL_TOOL")

		s.D.Set("allowed_roles", v.AllowedRoles)

		s.D.Set("default_execution_type", v.DefaultExecutionType)

		if v.Source != nil {
			s.D.Set("source", []interface{}{DatabaseToolsCustomSqlToolToolsetSourceToMap(v.Source)})
		} else {
			s.D.Set("source", nil)
		}

		if v.ToolDescription != nil {
			s.D.Set("tool_description", *v.ToolDescription)
		}

		if v.ToolName != nil {
			s.D.Set("tool_name", *v.ToolName)
		}

		variables := []interface{}{}
		for _, item := range v.Variables {
			variables = append(variables, DatabaseToolsMcpToolsetCustomSqlToolVariableToMap(item))
		}
		s.D.Set("variables", variables)

		if v.CompartmentId != nil {
			s.D.Set("compartment_id", *v.CompartmentId)
		}

		if v.DatabaseToolsMcpServerId != nil {
			s.D.Set("database_tools_mcp_server_id", *v.DatabaseToolsMcpServerId)
		}

		if v.DefinedTags != nil {
			s.D.Set("defined_tags", tfresource.DefinedTagsToMap(v.DefinedTags))
		}

		if v.Description != nil {
			s.D.Set("description", *v.Description)
		}

		if v.DisplayName != nil {
			s.D.Set("display_name", *v.DisplayName)
		}

		s.D.Set("freeform_tags", v.FreeformTags)

		if v.Id != nil {
			s.D.SetId(*v.Id)
		}

		if v.LifecycleDetails != nil {
			s.D.Set("lifecycle_details", *v.LifecycleDetails)
		}

		locks := []interface{}{}
		for _, item := range v.Locks {
			locks = append(locks, DbtoolsMcpToolsetResourceLockToMap(item))
		}
		s.D.Set("locks", locks)

		s.D.Set("state", v.LifecycleState)

		if v.SystemTags != nil {
			s.D.Set("system_tags", tfresource.SystemTagsToMap(v.SystemTags))
		}

		if v.TimeCreated != nil {
			s.D.Set("time_created", v.TimeCreated.String())
		}

		if v.TimeUpdated != nil {
			s.D.Set("time_updated", v.TimeUpdated.String())
		}

		if v.Version != nil {
			s.D.Set("version", *v.Version)
		}
	case oci_database_tools.DatabaseToolsMcpToolsetGenAiSqlAssistant:
		s.D.Set("type", "GENAI_SQL_ASSISTANT")

		s.D.Set("allowed_roles", nil)

		s.D.Set("default_execution_type", v.DefaultExecutionType)

		if v.GenerativeAiSemanticStoreId != nil {
			s.D.Set("generative_ai_semantic_store_id", *v.GenerativeAiSemanticStoreId)
		}

		tools := []interface{}{}
		for _, item := range v.Tools {
			tools = append(tools, DatabaseToolsMcpToolsetToolDetailsToMap(item))
		}
		s.D.Set("tools", tools)

		if v.CompartmentId != nil {
			s.D.Set("compartment_id", *v.CompartmentId)
		}

		if v.DatabaseToolsMcpServerId != nil {
			s.D.Set("database_tools_mcp_server_id", *v.DatabaseToolsMcpServerId)
		}

		if v.DefinedTags != nil {
			s.D.Set("defined_tags", tfresource.DefinedTagsToMap(v.DefinedTags))
		}

		if v.Description != nil {
			s.D.Set("description", *v.Description)
		}

		if v.DisplayName != nil {
			s.D.Set("display_name", *v.DisplayName)
		}

		s.D.Set("freeform_tags", v.FreeformTags)

		if v.Id != nil {
			s.D.SetId(*v.Id)
		}

		if v.LifecycleDetails != nil {
			s.D.Set("lifecycle_details", *v.LifecycleDetails)
		}

		locks := []interface{}{}
		for _, item := range v.Locks {
			locks = append(locks, DbtoolsMcpToolsetResourceLockToMap(item))
		}
		s.D.Set("locks", locks)

		s.D.Set("state", v.LifecycleState)

		if v.SystemTags != nil {
			s.D.Set("system_tags", tfresource.SystemTagsToMap(v.SystemTags))
		}

		if v.TimeCreated != nil {
			s.D.Set("time_created", v.TimeCreated.String())
		}

		if v.TimeUpdated != nil {
			s.D.Set("time_updated", v.TimeUpdated.String())
		}

		if v.Version != nil {
			s.D.Set("version", *v.Version)
		}
	default:
		log.Printf("[WARN] Received 'type' of unknown type %v", *s.Res)
		return nil
	}
	return nil
}

func (s *DatabaseToolsDatabaseToolsMcpToolsetResourceCrud) mapToCreateDatabaseToolsMcpToolsetToolDetails(fieldKeyFormat string) (oci_database_tools.CreateDatabaseToolsMcpToolsetToolDetails, error) {
	result := oci_database_tools.CreateDatabaseToolsMcpToolsetToolDetails{}

	if allowedRoles, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "allowed_roles")); ok {
		interfaces := allowedRoles.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "allowed_roles")) {
			result.AllowedRoles = tmp
		}
	}

	if name, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "name")); ok {
		tmp := name.(string)
		result.Name = &tmp
	}

	if status, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "status")); ok {
		result.Status = oci_database_tools.DatabaseToolsMcpToolsetToolStatusEnum(status.(string))
	}

	return result, nil
}

func CreateDatabaseToolsMcpToolsetToolDetailsToMap(obj oci_database_tools.CreateDatabaseToolsMcpToolsetToolDetails) map[string]interface{} {
	result := map[string]interface{}{}

	result["allowed_roles"] = obj.AllowedRoles

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	result["status"] = string(obj.Status)

	return result
}

func (s *DatabaseToolsDatabaseToolsMcpToolsetResourceCrud) mapToDatabaseToolsCustomSqlToolToolsetSource(fieldKeyFormat string) (oci_database_tools.DatabaseToolsCustomSqlToolToolsetSource, error) {
	result := oci_database_tools.DatabaseToolsCustomSqlToolToolsetSource{}

	if type_, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "type")); ok {
		result.Type = oci_database_tools.DatabaseToolsCustomSqlToolToolsetSourceTypeEnum(type_.(string))
	}

	if value, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "value")); ok {
		tmp := value.(string)
		result.Value = &tmp
	}

	return result, nil
}

func DatabaseToolsCustomSqlToolToolsetSourceToMap(obj *oci_database_tools.DatabaseToolsCustomSqlToolToolsetSource) map[string]interface{} {
	result := map[string]interface{}{}

	result["type"] = string(obj.Type)

	if obj.Value != nil {
		result["value"] = string(*obj.Value)
	}

	return result
}

func (s *DatabaseToolsDatabaseToolsMcpToolsetResourceCrud) mapToDatabaseToolsMcpToolsetCustomSqlToolVariable(fieldKeyFormat string) (oci_database_tools.DatabaseToolsMcpToolsetCustomSqlToolVariable, error) {
	result := oci_database_tools.DatabaseToolsMcpToolsetCustomSqlToolVariable{}

	if description, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "description")); ok {
		tmp := description.(string)
		result.Description = &tmp
	}

	if name, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "name")); ok {
		tmp := name.(string)
		result.Name = &tmp
	}

	if type_, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "type")); ok {
		tmp := type_.(string)
		result.Type = &tmp
	}

	return result, nil
}

func DatabaseToolsMcpToolsetCustomSqlToolVariableToMap(obj oci_database_tools.DatabaseToolsMcpToolsetCustomSqlToolVariable) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Description != nil {
		result["description"] = string(*obj.Description)
	}

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	if obj.Type != nil {
		result["type"] = string(*obj.Type)
	}

	return result
}

func (s *DatabaseToolsDatabaseToolsMcpToolsetResourceCrud) mapToDatabaseToolsMcpToolsetCustomizableReportingToolsReport(fieldKeyFormat string) (oci_database_tools.DatabaseToolsMcpToolsetCustomizableReportingToolsReport, error) {
	result := oci_database_tools.DatabaseToolsMcpToolsetCustomizableReportingToolsReport{}

	if allowedRoles, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "allowed_roles")); ok {
		interfaces := allowedRoles.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "allowed_roles")) {
			result.AllowedRoles = tmp
		}
	}

	if databaseToolsSqlReportId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "database_tools_sql_report_id")); ok {
		tmp := databaseToolsSqlReportId.(string)
		result.DatabaseToolsSqlReportId = &tmp
	}

	return result, nil
}

func DatabaseToolsMcpToolsetCustomizableReportingToolsReportToMap(obj oci_database_tools.DatabaseToolsMcpToolsetCustomizableReportingToolsReport) map[string]interface{} {
	result := map[string]interface{}{}

	result["allowed_roles"] = obj.AllowedRoles

	if obj.DatabaseToolsSqlReportId != nil {
		result["database_tools_sql_report_id"] = string(*obj.DatabaseToolsSqlReportId)
	}

	return result
}

func DatabaseToolsMcpToolsetSummaryToMap(obj oci_database_tools.DatabaseToolsMcpToolsetSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.GetId() != nil {
		result["id"] = *obj.GetId()
	}

	if obj.GetCompartmentId() != nil {
		result["compartment_id"] = *obj.GetCompartmentId()
	}

	if obj.GetDisplayName() != nil {
		result["display_name"] = *obj.GetDisplayName()
	}

	result["state"] = string(obj.GetLifecycleState())

	if obj.GetTimeCreated() != nil {
		result["time_created"] = obj.GetTimeCreated().String()
	}

	switch v := (obj).(type) {
	case oci_database_tools.DatabaseToolsMcpToolsetBuiltInSqlToolsSummary:
		result["type"] = "BUILT_IN_SQL_TOOLS"

		result["default_execution_type"] = string(v.DefaultExecutionType)

		tools := []interface{}{}
		for _, item := range v.Tools {
			tools = append(tools, DatabaseToolsMcpToolsetToolDetailsToMap(item))
		}
		result["tools"] = tools
	case oci_database_tools.DatabaseToolsMcpToolsetCustomizableReportingToolsSummary:
		result["type"] = "CUSTOMIZABLE_REPORTING_TOOLS"

		result["default_execution_type"] = string(v.DefaultExecutionType)

		reports := []interface{}{}
		for _, item := range v.Reports {
			reports = append(reports, DatabaseToolsMcpToolsetCustomizableReportingToolsReportToMap(item))
		}
		result["reports"] = reports

		tools := []interface{}{}
		for _, item := range v.Tools {
			tools = append(tools, DatabaseToolsMcpToolsetToolDetailsToMap(item))
		}
		result["tools"] = tools
	case oci_database_tools.DatabaseToolsMcpToolsetCustomSqlToolSummary:
		result["type"] = "CUSTOM_SQL_TOOL"

		result["allowed_roles"] = v.AllowedRoles

		result["default_execution_type"] = string(v.DefaultExecutionType)

		if v.ToolDescription != nil {
			result["tool_description"] = string(*v.ToolDescription)
		}

		if v.ToolName != nil {
			result["tool_name"] = string(*v.ToolName)
		}
	case oci_database_tools.DatabaseToolsMcpToolsetGenAiSqlAssistantSummary:
		result["type"] = "GENAI_SQL_ASSISTANT"

		result["default_execution_type"] = string(v.DefaultExecutionType)

		if v.GenerativeAiSemanticStoreId != nil {
			result["generative_ai_semantic_store_id"] = string(*v.GenerativeAiSemanticStoreId)
		}

		tools := []interface{}{}
		for _, item := range v.Tools {
			tools = append(tools, DatabaseToolsMcpToolsetToolDetailsToMap(item))
		}
		result["tools"] = tools
	default:
		log.Printf("[WARN] Received 'type' of unknown type %v", obj)
		return nil
	}

	return result
}

func (s *DatabaseToolsDatabaseToolsMcpToolsetResourceCrud) mapToDatabaseToolsMcpToolsetToolDetails(fieldKeyFormat string) (oci_database_tools.DatabaseToolsMcpToolsetToolDetails, error) {
	result := oci_database_tools.DatabaseToolsMcpToolsetToolDetails{}

	if allowedRoles, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "allowed_roles")); ok {
		interfaces := allowedRoles.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "allowed_roles")) {
			result.AllowedRoles = tmp
		}
	}

	if displayName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "display_name")); ok {
		tmp := displayName.(string)
		result.DisplayName = &tmp
	}

	if name, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "name")); ok {
		tmp := name.(string)
		result.Name = &tmp
	}

	if status, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "status")); ok {
		result.Status = oci_database_tools.DatabaseToolsMcpToolsetToolStatusEnum(status.(string))
	}

	return result, nil
}

func DatabaseToolsMcpToolsetToolDetailsToMap(obj oci_database_tools.DatabaseToolsMcpToolsetToolDetails) map[string]interface{} {
	result := map[string]interface{}{}

	result["allowed_roles"] = obj.AllowedRoles

	if obj.DisplayName != nil {
		result["display_name"] = string(*obj.DisplayName)
	}

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	result["status"] = string(obj.Status)

	return result
}

func (s *DatabaseToolsDatabaseToolsMcpToolsetResourceCrud) mapToResourceLock(fieldKeyFormat string) (oci_database_tools.ResourceLock, error) {
	result := oci_database_tools.ResourceLock{}

	if message, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "message")); ok {
		tmp := message.(string)
		result.Message = &tmp
	}

	if relatedResourceId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "related_resource_id")); ok {
		tmp := relatedResourceId.(string)
		result.RelatedResourceId = &tmp
	}

	if timeCreated, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "time_created")); ok {
		tmp, err := time.Parse(time.RFC3339, timeCreated.(string))
		if err != nil {
			return result, err
		}
		result.TimeCreated = &oci_common.SDKTime{Time: tmp}
	}

	if type_, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "type")); ok {
		result.Type = oci_database_tools.ResourceLockTypeEnum(type_.(string))
	}

	return result, nil
}

func DbtoolsMcpToolsetResourceLockToMap(obj oci_database_tools.ResourceLock) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Message != nil {
		result["message"] = string(*obj.Message)
	}

	if obj.RelatedResourceId != nil {
		result["related_resource_id"] = string(*obj.RelatedResourceId)
	}

	if obj.TimeCreated != nil {
		result["time_created"] = obj.TimeCreated.Format(time.RFC3339Nano)
	}

	result["type"] = string(obj.Type)

	return result
}

func (s *DatabaseToolsDatabaseToolsMcpToolsetResourceCrud) mapToUpdateDatabaseToolsMcpToolsetToolDetails(fieldKeyFormat string) (oci_database_tools.UpdateDatabaseToolsMcpToolsetToolDetails, error) {
	result := oci_database_tools.UpdateDatabaseToolsMcpToolsetToolDetails{}

	if allowedRoles, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "allowed_roles")); ok {
		interfaces := allowedRoles.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "allowed_roles")) {
			result.AllowedRoles = tmp
		}
	}

	if name, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "name")); ok {
		tmp := name.(string)
		result.Name = &tmp
	}

	if status, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "status")); ok {
		result.Status = oci_database_tools.DatabaseToolsMcpToolsetToolStatusEnum(status.(string))
	}

	return result, nil
}

func UpdateDatabaseToolsMcpToolsetToolDetailsToMap(obj oci_database_tools.UpdateDatabaseToolsMcpToolsetToolDetails) map[string]interface{} {
	result := map[string]interface{}{}

	result["allowed_roles"] = obj.AllowedRoles

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	result["status"] = string(obj.Status)

	return result
}

func (s *DatabaseToolsDatabaseToolsMcpToolsetResourceCrud) populateTopLevelPolymorphicCreateDatabaseToolsMcpToolsetRequest(request *oci_database_tools.CreateDatabaseToolsMcpToolsetRequest) error {
	//discriminator
	typeRaw, ok := s.D.GetOkExists("type")
	var type_ string
	if ok {
		type_ = typeRaw.(string)
	} else {
		type_ = "" // default value
	}
	switch strings.ToLower(type_) {
	case strings.ToLower("BUILT_IN_SQL_TOOLS"):
		details := oci_database_tools.CreateDatabaseToolsMcpToolsetBuiltInSqlToolsDetails{}
		if defaultExecutionType, ok := s.D.GetOkExists("default_execution_type"); ok {
			details.DefaultExecutionType = oci_database_tools.DatabaseToolsMcpToolsetDefaultExecutionTypeEnum(defaultExecutionType.(string))
		}
		if tools, ok := s.D.GetOkExists("tools"); ok {
			interfaces := tools.([]interface{})
			tmp := make([]oci_database_tools.CreateDatabaseToolsMcpToolsetToolDetails, len(interfaces))
			for i := range interfaces {
				stateDataIndex := i
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "tools", stateDataIndex)
				converted, err := s.mapToCreateDatabaseToolsMcpToolsetToolDetails(fieldKeyFormat)
				if err != nil {
					return err
				}
				tmp[i] = converted
			}
			if len(tmp) != 0 || s.D.HasChange("tools") {
				details.Tools = tmp
			}
		}
		if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
			tmp := compartmentId.(string)
			details.CompartmentId = &tmp
		}
		if databaseToolsMcpServerId, ok := s.D.GetOkExists("database_tools_mcp_server_id"); ok {
			tmp := databaseToolsMcpServerId.(string)
			details.DatabaseToolsMcpServerId = &tmp
		}
		if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
			convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
			if err != nil {
				return err
			}
			details.DefinedTags = convertedDefinedTags
		}
		if description, ok := s.D.GetOkExists("description"); ok {
			tmp := description.(string)
			details.Description = &tmp
		}
		if displayName, ok := s.D.GetOkExists("display_name"); ok {
			tmp := displayName.(string)
			details.DisplayName = &tmp
		}
		if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
			details.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
		}
		if locks, ok := s.D.GetOkExists("locks"); ok {
			interfaces := locks.([]interface{})
			tmp := make([]oci_database_tools.ResourceLock, len(interfaces))
			for i := range interfaces {
				stateDataIndex := i
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "locks", stateDataIndex)
				converted, err := s.mapToResourceLock(fieldKeyFormat)
				if err != nil {
					return err
				}
				tmp[i] = converted
			}
			if len(tmp) != 0 || s.D.HasChange("locks") {
				details.Locks = tmp
			}
		}
		if version, ok := s.D.GetOkExists("version"); ok {
			tmp := version.(int)
			details.Version = &tmp
		}
		request.CreateDatabaseToolsMcpToolsetDetails = details
	case strings.ToLower("CUSTOMIZABLE_REPORTING_TOOLS"):
		details := oci_database_tools.CreateDatabaseToolsMcpToolsetCustomizableReportingToolsDetails{}
		if defaultExecutionType, ok := s.D.GetOkExists("default_execution_type"); ok {
			details.DefaultExecutionType = oci_database_tools.DatabaseToolsMcpToolsetDefaultExecutionTypeEnum(defaultExecutionType.(string))
		}
		if reports, ok := s.D.GetOkExists("reports"); ok {
			interfaces := reports.([]interface{})
			tmp := make([]oci_database_tools.DatabaseToolsMcpToolsetCustomizableReportingToolsReport, len(interfaces))
			for i := range interfaces {
				stateDataIndex := i
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "reports", stateDataIndex)
				converted, err := s.mapToDatabaseToolsMcpToolsetCustomizableReportingToolsReport(fieldKeyFormat)
				if err != nil {
					return err
				}
				tmp[i] = converted
			}
			if len(tmp) != 0 || s.D.HasChange("reports") {
				details.Reports = tmp
			}
		}
		if tools, ok := s.D.GetOkExists("tools"); ok {
			interfaces := tools.([]interface{})
			tmp := make([]oci_database_tools.CreateDatabaseToolsMcpToolsetToolDetails, len(interfaces))
			for i := range interfaces {
				stateDataIndex := i
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "tools", stateDataIndex)
				converted, err := s.mapToCreateDatabaseToolsMcpToolsetToolDetails(fieldKeyFormat)
				if err != nil {
					return err
				}
				tmp[i] = converted
			}
			if len(tmp) != 0 || s.D.HasChange("tools") {
				details.Tools = tmp
			}
		}
		if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
			tmp := compartmentId.(string)
			details.CompartmentId = &tmp
		}
		if databaseToolsMcpServerId, ok := s.D.GetOkExists("database_tools_mcp_server_id"); ok {
			tmp := databaseToolsMcpServerId.(string)
			details.DatabaseToolsMcpServerId = &tmp
		}
		if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
			convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
			if err != nil {
				return err
			}
			details.DefinedTags = convertedDefinedTags
		}
		if description, ok := s.D.GetOkExists("description"); ok {
			tmp := description.(string)
			details.Description = &tmp
		}
		if displayName, ok := s.D.GetOkExists("display_name"); ok {
			tmp := displayName.(string)
			details.DisplayName = &tmp
		}
		if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
			details.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
		}
		if locks, ok := s.D.GetOkExists("locks"); ok {
			interfaces := locks.([]interface{})
			tmp := make([]oci_database_tools.ResourceLock, len(interfaces))
			for i := range interfaces {
				stateDataIndex := i
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "locks", stateDataIndex)
				converted, err := s.mapToResourceLock(fieldKeyFormat)
				if err != nil {
					return err
				}
				tmp[i] = converted
			}
			if len(tmp) != 0 || s.D.HasChange("locks") {
				details.Locks = tmp
			}
		}
		if version, ok := s.D.GetOkExists("version"); ok {
			tmp := version.(int)
			details.Version = &tmp
		}
		request.CreateDatabaseToolsMcpToolsetDetails = details
	case strings.ToLower("CUSTOM_SQL_TOOL"):
		details := oci_database_tools.CreateDatabaseToolsMcpToolsetCustomSqlToolDetails{}
		if allowedRoles, ok := s.D.GetOkExists("allowed_roles"); ok {
			interfaces := allowedRoles.([]interface{})
			tmp := make([]string, len(interfaces))
			for i := range interfaces {
				if interfaces[i] != nil {
					tmp[i] = interfaces[i].(string)
				}
			}
			if len(tmp) != 0 || s.D.HasChange("allowed_roles") {
				details.AllowedRoles = tmp
			}
		}
		if defaultExecutionType, ok := s.D.GetOkExists("default_execution_type"); ok {
			details.DefaultExecutionType = oci_database_tools.DatabaseToolsMcpToolsetDefaultExecutionTypeEnum(defaultExecutionType.(string))
		}
		if source, ok := s.D.GetOkExists("source"); ok {
			if tmpList := source.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "source", 0)
				tmp, err := s.mapToDatabaseToolsCustomSqlToolToolsetSource(fieldKeyFormat)
				if err != nil {
					return err
				}
				details.Source = &tmp
			}
		}
		if toolDescription, ok := s.D.GetOkExists("tool_description"); ok {
			tmp := toolDescription.(string)
			details.ToolDescription = &tmp
		}
		if toolName, ok := s.D.GetOkExists("tool_name"); ok {
			tmp := toolName.(string)
			details.ToolName = &tmp
		}
		if variables, ok := s.D.GetOkExists("variables"); ok {
			interfaces := variables.([]interface{})
			tmp := make([]oci_database_tools.DatabaseToolsMcpToolsetCustomSqlToolVariable, len(interfaces))
			for i := range interfaces {
				stateDataIndex := i
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "variables", stateDataIndex)
				converted, err := s.mapToDatabaseToolsMcpToolsetCustomSqlToolVariable(fieldKeyFormat)
				if err != nil {
					return err
				}
				tmp[i] = converted
			}
			if len(tmp) != 0 || s.D.HasChange("variables") {
				details.Variables = tmp
			}
		}
		if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
			tmp := compartmentId.(string)
			details.CompartmentId = &tmp
		}
		if databaseToolsMcpServerId, ok := s.D.GetOkExists("database_tools_mcp_server_id"); ok {
			tmp := databaseToolsMcpServerId.(string)
			details.DatabaseToolsMcpServerId = &tmp
		}
		if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
			convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
			if err != nil {
				return err
			}
			details.DefinedTags = convertedDefinedTags
		}
		if description, ok := s.D.GetOkExists("description"); ok {
			tmp := description.(string)
			details.Description = &tmp
		}
		if displayName, ok := s.D.GetOkExists("display_name"); ok {
			tmp := displayName.(string)
			details.DisplayName = &tmp
		}
		if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
			details.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
		}
		if locks, ok := s.D.GetOkExists("locks"); ok {
			interfaces := locks.([]interface{})
			tmp := make([]oci_database_tools.ResourceLock, len(interfaces))
			for i := range interfaces {
				stateDataIndex := i
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "locks", stateDataIndex)
				converted, err := s.mapToResourceLock(fieldKeyFormat)
				if err != nil {
					return err
				}
				tmp[i] = converted
			}
			if len(tmp) != 0 || s.D.HasChange("locks") {
				details.Locks = tmp
			}
		}
		if version, ok := s.D.GetOkExists("version"); ok {
			tmp := version.(int)
			details.Version = &tmp
		}
		request.CreateDatabaseToolsMcpToolsetDetails = details
	case strings.ToLower("GENAI_SQL_ASSISTANT"):
		details := oci_database_tools.CreateDatabaseToolsMcpToolsetGenAiSqlAssistantDetails{}
		if defaultExecutionType, ok := s.D.GetOkExists("default_execution_type"); ok {
			details.DefaultExecutionType = oci_database_tools.DatabaseToolsMcpToolsetDefaultExecutionTypeEnum(defaultExecutionType.(string))
		}
		if generativeAiSemanticStoreId, ok := s.D.GetOkExists("generative_ai_semantic_store_id"); ok {
			tmp := generativeAiSemanticStoreId.(string)
			details.GenerativeAiSemanticStoreId = &tmp
		}
		if tools, ok := s.D.GetOkExists("tools"); ok {
			interfaces := tools.([]interface{})
			tmp := make([]oci_database_tools.CreateDatabaseToolsMcpToolsetToolDetails, len(interfaces))
			for i := range interfaces {
				stateDataIndex := i
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "tools", stateDataIndex)
				converted, err := s.mapToCreateDatabaseToolsMcpToolsetToolDetails(fieldKeyFormat)
				if err != nil {
					return err
				}
				tmp[i] = converted
			}
			if len(tmp) != 0 || s.D.HasChange("tools") {
				details.Tools = tmp
			}
		}
		if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
			tmp := compartmentId.(string)
			details.CompartmentId = &tmp
		}
		if databaseToolsMcpServerId, ok := s.D.GetOkExists("database_tools_mcp_server_id"); ok {
			tmp := databaseToolsMcpServerId.(string)
			details.DatabaseToolsMcpServerId = &tmp
		}
		if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
			convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
			if err != nil {
				return err
			}
			details.DefinedTags = convertedDefinedTags
		}
		if description, ok := s.D.GetOkExists("description"); ok {
			tmp := description.(string)
			details.Description = &tmp
		}
		if displayName, ok := s.D.GetOkExists("display_name"); ok {
			tmp := displayName.(string)
			details.DisplayName = &tmp
		}
		if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
			details.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
		}
		if locks, ok := s.D.GetOkExists("locks"); ok {
			interfaces := locks.([]interface{})
			tmp := make([]oci_database_tools.ResourceLock, len(interfaces))
			for i := range interfaces {
				stateDataIndex := i
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "locks", stateDataIndex)
				converted, err := s.mapToResourceLock(fieldKeyFormat)
				if err != nil {
					return err
				}
				tmp[i] = converted
			}
			if len(tmp) != 0 || s.D.HasChange("locks") {
				details.Locks = tmp
			}
		}
		if version, ok := s.D.GetOkExists("version"); ok {
			tmp := version.(int)
			details.Version = &tmp
		}
		request.CreateDatabaseToolsMcpToolsetDetails = details
	default:
		return fmt.Errorf("unknown type '%v' was specified", type_)
	}
	return nil
}

func (s *DatabaseToolsDatabaseToolsMcpToolsetResourceCrud) populateTopLevelPolymorphicUpdateDatabaseToolsMcpToolsetRequest(request *oci_database_tools.UpdateDatabaseToolsMcpToolsetRequest) error {
	//discriminator
	typeRaw, ok := s.D.GetOkExists("type")
	var type_ string
	if ok {
		type_ = typeRaw.(string)
	} else {
		type_ = "" // default value
	}
	switch strings.ToLower(type_) {
	case strings.ToLower("BUILT_IN_SQL_TOOLS"):
		details := oci_database_tools.UpdateDatabaseToolsMcpToolsetBuiltInSqlToolsDetails{}
		if defaultExecutionType, ok := s.D.GetOkExists("default_execution_type"); ok {
			details.DefaultExecutionType = oci_database_tools.DatabaseToolsMcpToolsetDefaultExecutionTypeEnum(defaultExecutionType.(string))
		}
		if tools, ok := s.D.GetOkExists("tools"); ok {
			interfaces := tools.([]interface{})
			tmp := make([]oci_database_tools.UpdateDatabaseToolsMcpToolsetToolDetails, len(interfaces))
			for i := range interfaces {
				stateDataIndex := i
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "tools", stateDataIndex)
				converted, err := s.mapToUpdateDatabaseToolsMcpToolsetToolDetails(fieldKeyFormat)
				if err != nil {
					return err
				}
				tmp[i] = converted
			}
			if len(tmp) != 0 || s.D.HasChange("tools") {
				details.Tools = tmp
			}
		}
		tmp := s.D.Id()
		request.DatabaseToolsMcpToolsetId = &tmp
		if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
			convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
			if err != nil {
				return err
			}
			details.DefinedTags = convertedDefinedTags
		}
		if description, ok := s.D.GetOkExists("description"); ok {
			tmp := description.(string)
			details.Description = &tmp
		}
		if displayName, ok := s.D.GetOkExists("display_name"); ok {
			tmp := displayName.(string)
			details.DisplayName = &tmp
		}
		if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
			details.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
		}
		if version, ok := s.D.GetOkExists("version"); ok {
			tmp := version.(int)
			details.Version = &tmp
		}
		request.UpdateDatabaseToolsMcpToolsetDetails = details
	case strings.ToLower("CUSTOMIZABLE_REPORTING_TOOLS"):
		details := oci_database_tools.UpdateDatabaseToolsMcpToolsetCustomizableReportingToolsDetails{}
		if defaultExecutionType, ok := s.D.GetOkExists("default_execution_type"); ok {
			details.DefaultExecutionType = oci_database_tools.DatabaseToolsMcpToolsetDefaultExecutionTypeEnum(defaultExecutionType.(string))
		}
		if reports, ok := s.D.GetOkExists("reports"); ok {
			interfaces := reports.([]interface{})
			tmp := make([]oci_database_tools.DatabaseToolsMcpToolsetCustomizableReportingToolsReport, len(interfaces))
			for i := range interfaces {
				stateDataIndex := i
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "reports", stateDataIndex)
				converted, err := s.mapToDatabaseToolsMcpToolsetCustomizableReportingToolsReport(fieldKeyFormat)
				if err != nil {
					return err
				}
				tmp[i] = converted
			}
			if len(tmp) != 0 || s.D.HasChange("reports") {
				details.Reports = tmp
			}
		}
		if tools, ok := s.D.GetOkExists("tools"); ok {
			interfaces := tools.([]interface{})
			tmp := make([]oci_database_tools.UpdateDatabaseToolsMcpToolsetToolDetails, len(interfaces))
			for i := range interfaces {
				stateDataIndex := i
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "tools", stateDataIndex)
				converted, err := s.mapToUpdateDatabaseToolsMcpToolsetToolDetails(fieldKeyFormat)
				if err != nil {
					return err
				}
				tmp[i] = converted
			}
			if len(tmp) != 0 || s.D.HasChange("tools") {
				details.Tools = tmp
			}
		}
		tmp := s.D.Id()
		request.DatabaseToolsMcpToolsetId = &tmp
		if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
			convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
			if err != nil {
				return err
			}
			details.DefinedTags = convertedDefinedTags
		}
		if description, ok := s.D.GetOkExists("description"); ok {
			tmp := description.(string)
			details.Description = &tmp
		}
		if displayName, ok := s.D.GetOkExists("display_name"); ok {
			tmp := displayName.(string)
			details.DisplayName = &tmp
		}
		if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
			details.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
		}
		if version, ok := s.D.GetOkExists("version"); ok {
			tmp := version.(int)
			details.Version = &tmp
		}
		request.UpdateDatabaseToolsMcpToolsetDetails = details
	case strings.ToLower("CUSTOM_SQL_TOOL"):
		details := oci_database_tools.UpdateDatabaseToolsMcpToolsetCustomSqlToolDetails{}
		if allowedRoles, ok := s.D.GetOkExists("allowed_roles"); ok {
			interfaces := allowedRoles.([]interface{})
			tmp := make([]string, len(interfaces))
			for i := range interfaces {
				if interfaces[i] != nil {
					tmp[i] = interfaces[i].(string)
				}
			}
			if len(tmp) != 0 || s.D.HasChange("allowed_roles") {
				details.AllowedRoles = tmp
			}
		}
		if defaultExecutionType, ok := s.D.GetOkExists("default_execution_type"); ok {
			details.DefaultExecutionType = oci_database_tools.DatabaseToolsMcpToolsetDefaultExecutionTypeEnum(defaultExecutionType.(string))
		}
		if source, ok := s.D.GetOkExists("source"); ok {
			if tmpList := source.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "source", 0)
				tmp, err := s.mapToDatabaseToolsCustomSqlToolToolsetSource(fieldKeyFormat)
				if err != nil {
					return err
				}
				details.Source = &tmp
			}
		}
		if toolDescription, ok := s.D.GetOkExists("tool_description"); ok {
			tmp := toolDescription.(string)
			details.ToolDescription = &tmp
		}
		if toolName, ok := s.D.GetOkExists("tool_name"); ok {
			tmp := toolName.(string)
			details.ToolName = &tmp
		}
		if variables, ok := s.D.GetOkExists("variables"); ok {
			interfaces := variables.([]interface{})
			tmp := make([]oci_database_tools.DatabaseToolsMcpToolsetCustomSqlToolVariable, len(interfaces))
			for i := range interfaces {
				stateDataIndex := i
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "variables", stateDataIndex)
				converted, err := s.mapToDatabaseToolsMcpToolsetCustomSqlToolVariable(fieldKeyFormat)
				if err != nil {
					return err
				}
				tmp[i] = converted
			}
			if len(tmp) != 0 || s.D.HasChange("variables") {
				details.Variables = tmp
			}
		}
		tmp := s.D.Id()
		request.DatabaseToolsMcpToolsetId = &tmp
		if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
			convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
			if err != nil {
				return err
			}
			details.DefinedTags = convertedDefinedTags
		}
		if description, ok := s.D.GetOkExists("description"); ok {
			tmp := description.(string)
			details.Description = &tmp
		}
		if displayName, ok := s.D.GetOkExists("display_name"); ok {
			tmp := displayName.(string)
			details.DisplayName = &tmp
		}
		if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
			details.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
		}
		if version, ok := s.D.GetOkExists("version"); ok {
			tmp := version.(int)
			details.Version = &tmp
		}
		request.UpdateDatabaseToolsMcpToolsetDetails = details
	case strings.ToLower("GENAI_SQL_ASSISTANT"):
		details := oci_database_tools.UpdateDatabaseToolsMcpToolsetGenAiSqlAssistantDetails{}
		if defaultExecutionType, ok := s.D.GetOkExists("default_execution_type"); ok {
			details.DefaultExecutionType = oci_database_tools.DatabaseToolsMcpToolsetDefaultExecutionTypeEnum(defaultExecutionType.(string))
		}
		if generativeAiSemanticStoreId, ok := s.D.GetOkExists("generative_ai_semantic_store_id"); ok {
			tmp := generativeAiSemanticStoreId.(string)
			details.GenerativeAiSemanticStoreId = &tmp
		}
		if tools, ok := s.D.GetOkExists("tools"); ok {
			interfaces := tools.([]interface{})
			tmp := make([]oci_database_tools.UpdateDatabaseToolsMcpToolsetToolDetails, len(interfaces))
			for i := range interfaces {
				stateDataIndex := i
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "tools", stateDataIndex)
				converted, err := s.mapToUpdateDatabaseToolsMcpToolsetToolDetails(fieldKeyFormat)
				if err != nil {
					return err
				}
				tmp[i] = converted
			}
			if len(tmp) != 0 || s.D.HasChange("tools") {
				details.Tools = tmp
			}
		}
		tmp := s.D.Id()
		request.DatabaseToolsMcpToolsetId = &tmp
		if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
			convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
			if err != nil {
				return err
			}
			details.DefinedTags = convertedDefinedTags
		}
		if description, ok := s.D.GetOkExists("description"); ok {
			tmp := description.(string)
			details.Description = &tmp
		}
		if displayName, ok := s.D.GetOkExists("display_name"); ok {
			tmp := displayName.(string)
			details.DisplayName = &tmp
		}
		if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
			details.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
		}
		if version, ok := s.D.GetOkExists("version"); ok {
			tmp := version.(int)
			details.Version = &tmp
		}
		request.UpdateDatabaseToolsMcpToolsetDetails = details
	default:
		return fmt.Errorf("unknown type '%v' was specified", type_)
	}
	return nil
}

func (s *DatabaseToolsDatabaseToolsMcpToolsetResourceCrud) updateCompartment(ctx context.Context, compartment interface{}) error {
	changeCompartmentRequest := oci_database_tools.ChangeDatabaseToolsMcpToolsetCompartmentRequest{}

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	idTmp := s.D.Id()
	changeCompartmentRequest.DatabaseToolsMcpToolsetId = &idTmp

	if isLockOverride, ok := s.D.GetOkExists("is_lock_override"); ok {
		tmp := isLockOverride.(bool)
		changeCompartmentRequest.IsLockOverride = &tmp
	}

	changeCompartmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database_tools")

	_, err := s.Client.ChangeDatabaseToolsMcpToolsetCompartment(ctx, changeCompartmentRequest)
	if err != nil {
		return err
	}

	if waitErr := tfresource.WaitForUpdatedStateWithContext(ctx, s.D, s); waitErr != nil {
		return waitErr
	}

	return nil
}
