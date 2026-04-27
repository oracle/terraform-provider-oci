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

func DatabaseToolsDatabaseToolsMcpServerResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts:      tfresource.DefaultTimeout,
		CreateContext: createDatabaseToolsDatabaseToolsMcpServerWithContext,
		ReadContext:   readDatabaseToolsDatabaseToolsMcpServerWithContext,
		UpdateContext: updateDatabaseToolsDatabaseToolsMcpServerWithContext,
		DeleteContext: deleteDatabaseToolsDatabaseToolsMcpServerWithContext,
		Schema: map[string]*schema.Schema{
			// Required
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"database_tools_connection_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"domain_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"storage": {
				Type:     schema.TypeList,
				Required: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"type": {
							Type:             schema.TypeString,
							Required:         true,
							DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
							ValidateFunc: validation.StringInSlice([]string{
								"NONE",
								"OBJECT_STORAGE",
							}, true),
						},

						// Optional
						"bucket": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							MaxItems: 1,
							MinItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional
									"bucket": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"namespace": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},

									// Computed
								},
							},
						},

						// Computed
					},
				},
			},
			"type": {
				Type:             schema.TypeString,
				Required:         true,
				DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
				ValidateFunc: validation.StringInSlice([]string{
					"DEFAULT",
				}, true),
			},

			// Optional
			"access_token_expiry_in_seconds": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"custom_roles": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"description": {
							Type:     schema.TypeString,
							Required: true,
						},
						"display_name": {
							Type:     schema.TypeString,
							Required: true,
						},

						// Optional

						// Computed
					},
				},
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
			"refresh_token_expiry_in_seconds": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"runtime_identity": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},

			// Computed
			"built_in_roles": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"description": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"display_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"domain_app_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"endpoints": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"endpoint": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"type": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"lifecycle_details": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"related_resource": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"entity_type": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"identifier": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
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

func createDatabaseToolsDatabaseToolsMcpServerWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &DatabaseToolsDatabaseToolsMcpServerResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseToolsClient()

	return tfresource.HandleDiagError(m, tfresource.CreateResourceWithContext(ctx, d, sync))
}

func readDatabaseToolsDatabaseToolsMcpServerWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &DatabaseToolsDatabaseToolsMcpServerResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseToolsClient()

	return tfresource.HandleDiagError(m, tfresource.ReadResourceWithContext(ctx, sync))
}

func updateDatabaseToolsDatabaseToolsMcpServerWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &DatabaseToolsDatabaseToolsMcpServerResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseToolsClient()

	return tfresource.HandleDiagError(m, tfresource.UpdateResourceWithContext(ctx, d, sync))
}

func deleteDatabaseToolsDatabaseToolsMcpServerWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &DatabaseToolsDatabaseToolsMcpServerResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseToolsClient()
	sync.DisableNotFoundRetries = true

	return tfresource.HandleDiagError(m, tfresource.DeleteResourceWithContext(ctx, d, sync))
}

type DatabaseToolsDatabaseToolsMcpServerResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_database_tools.DatabaseToolsClient
	Res                    *oci_database_tools.DatabaseToolsMcpServer
	DisableNotFoundRetries bool
}

func (s *DatabaseToolsDatabaseToolsMcpServerResourceCrud) ID() string {
	databaseToolsMcpServer := *s.Res
	return *databaseToolsMcpServer.GetId()
}

func (s *DatabaseToolsDatabaseToolsMcpServerResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_database_tools.DatabaseToolsMcpServerLifecycleStateCreating),
	}
}

func (s *DatabaseToolsDatabaseToolsMcpServerResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_database_tools.DatabaseToolsMcpServerLifecycleStateActive),
		string(oci_database_tools.DatabaseToolsMcpServerLifecycleStateNeedsAttention),
	}
}

func (s *DatabaseToolsDatabaseToolsMcpServerResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_database_tools.DatabaseToolsMcpServerLifecycleStateDeleting),
	}
}

func (s *DatabaseToolsDatabaseToolsMcpServerResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_database_tools.DatabaseToolsMcpServerLifecycleStateDeleted),
	}
}

func (s *DatabaseToolsDatabaseToolsMcpServerResourceCrud) CreateWithContext(ctx context.Context) error {
	request := oci_database_tools.CreateDatabaseToolsMcpServerRequest{}
	err := s.populateTopLevelPolymorphicCreateDatabaseToolsMcpServerRequest(&request)
	if err != nil {
		return err
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database_tools")

	response, err := s.Client.CreateDatabaseToolsMcpServer(ctx, request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	var identifier *string
	identifier = response.GetId()
	if identifier != nil {
		s.D.SetId(*identifier)
	}
	return s.getDatabaseToolsMcpServerFromWorkRequest(ctx, workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database_tools"), oci_database_tools.ActionTypeCreated, s.D.Timeout(schema.TimeoutCreate))
}

func (s *DatabaseToolsDatabaseToolsMcpServerResourceCrud) getDatabaseToolsMcpServerFromWorkRequest(ctx context.Context, workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_database_tools.ActionTypeEnum, timeout time.Duration) error {

	// Wait until it finishes
	databaseToolsMcpServerId, err := databaseToolsMcpServerWaitForWorkRequest(ctx, workId, "databasetoolsmcpserver",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.Client)

	if err != nil {
		return err
	}
	s.D.SetId(*databaseToolsMcpServerId)

	return s.GetWithContext(ctx)
}

func databaseToolsMcpServerWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
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

func databaseToolsMcpServerWaitForWorkRequest(ctx context.Context, wId *string, entityType string, action oci_database_tools.ActionTypeEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_database_tools.DatabaseToolsClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "database_tools")
	retryPolicy.ShouldRetryOperation = databaseToolsMcpServerWorkRequestShouldRetryFunc(timeout)

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
		return nil, getErrorFromDatabaseToolsDatabaseToolsMcpServerWorkRequest(ctx, client, wId, retryPolicy, entityType, action)
	}

	return identifier, nil
}

func getErrorFromDatabaseToolsDatabaseToolsMcpServerWorkRequest(ctx context.Context, client *oci_database_tools.DatabaseToolsClient, workId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_database_tools.ActionTypeEnum) error {
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

func (s *DatabaseToolsDatabaseToolsMcpServerResourceCrud) GetWithContext(ctx context.Context) error {
	request := oci_database_tools.GetDatabaseToolsMcpServerRequest{}

	tmp := s.D.Id()
	request.DatabaseToolsMcpServerId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database_tools")

	response, err := s.Client.GetDatabaseToolsMcpServer(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response.DatabaseToolsMcpServer
	return nil
}

func (s *DatabaseToolsDatabaseToolsMcpServerResourceCrud) UpdateWithContext(ctx context.Context) error {
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(ctx, compartment)
			if err != nil {
				return err
			}
		}
	}
	request := oci_database_tools.UpdateDatabaseToolsMcpServerRequest{}
	err := s.populateTopLevelPolymorphicUpdateDatabaseToolsMcpServerRequest(&request)
	if err != nil {
		return err
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database_tools")

	response, err := s.Client.UpdateDatabaseToolsMcpServer(ctx, request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getDatabaseToolsMcpServerFromWorkRequest(ctx, workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database_tools"), oci_database_tools.ActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}

func (s *DatabaseToolsDatabaseToolsMcpServerResourceCrud) DeleteWithContext(ctx context.Context) error {
	request := oci_database_tools.DeleteDatabaseToolsMcpServerRequest{}

	tmp := s.D.Id()
	request.DatabaseToolsMcpServerId = &tmp

	if isLockOverride, ok := s.D.GetOkExists("is_lock_override"); ok {
		tmp := isLockOverride.(bool)
		request.IsLockOverride = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database_tools")

	response, err := s.Client.DeleteDatabaseToolsMcpServer(ctx, request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	// Wait until it finishes
	_, delWorkRequestErr := databaseToolsMcpServerWaitForWorkRequest(ctx, workId, "databasetoolsmcpserver",
		oci_database_tools.ActionTypeDeleted, s.D.Timeout(schema.TimeoutDelete), s.DisableNotFoundRetries, s.Client)
	return delWorkRequestErr
}

func (s *DatabaseToolsDatabaseToolsMcpServerResourceCrud) SetData() error {
	switch v := (*s.Res).(type) {
	case oci_database_tools.DatabaseToolsMcpServerDefault:
		s.D.Set("type", "DEFAULT")

		if v.DomainAppId != nil {
			s.D.Set("domain_app_id", *v.DomainAppId)
		}

		if v.DomainId != nil {
			s.D.Set("domain_id", *v.DomainId)
		}

		if v.Storage != nil {
			storageArray := []interface{}{}
			if storageMap := DatabaseToolsMcpServerStorageToMap(&v.Storage); storageMap != nil {
				storageArray = append(storageArray, storageMap)
			}
			s.D.Set("storage", storageArray)
		} else {
			s.D.Set("storage", nil)
		}

		if v.AccessTokenExpiryInSeconds != nil {
			s.D.Set("access_token_expiry_in_seconds", *v.AccessTokenExpiryInSeconds)
		}

		builtInRoles := []interface{}{}
		for _, item := range v.BuiltInRoles {
			builtInRoles = append(builtInRoles, DatabaseToolsMcpServerBuiltInRoleToMap(item))
		}
		s.D.Set("built_in_roles", builtInRoles)

		if v.CompartmentId != nil {
			s.D.Set("compartment_id", *v.CompartmentId)
		}

		customRoles := []interface{}{}
		for _, item := range v.CustomRoles {
			customRoles = append(customRoles, DatabaseToolsMcpServerCustomRoleToMap(item))
		}
		s.D.Set("custom_roles", customRoles)

		if v.DatabaseToolsConnectionId != nil {
			s.D.Set("database_tools_connection_id", *v.DatabaseToolsConnectionId)
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

		endpoints := []interface{}{}
		for _, item := range v.Endpoints {
			endpoints = append(endpoints, DatabaseToolsMcpServerEndpointToMap(item))
		}
		s.D.Set("endpoints", endpoints)

		s.D.Set("freeform_tags", v.FreeformTags)

		if v.Id != nil {
			s.D.SetId(*v.Id)
		}

		if v.LifecycleDetails != nil {
			s.D.Set("lifecycle_details", *v.LifecycleDetails)
		}

		locks := []interface{}{}
		for _, item := range v.Locks {
			locks = append(locks, DbtoolsMcpServerResourceLockToMap(item))
		}
		s.D.Set("locks", locks)

		if v.RefreshTokenExpiryInSeconds != nil {
			s.D.Set("refresh_token_expiry_in_seconds", *v.RefreshTokenExpiryInSeconds)
		}

		if v.RelatedResource != nil {
			s.D.Set("related_resource", []interface{}{DatabaseToolsMcpServerRelatedResourceToMap(v.RelatedResource)})
		} else {
			s.D.Set("related_resource", nil)
		}

		s.D.Set("runtime_identity", v.RuntimeIdentity)

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
	default:
		log.Printf("[WARN] Received 'type' of unknown type %v", *s.Res)
		return nil
	}
	return nil
}

func DatabaseToolsMcpServerBuiltInRoleToMap(obj oci_database_tools.DatabaseToolsMcpServerBuiltInRole) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Description != nil {
		result["description"] = string(*obj.Description)
	}

	if obj.DisplayName != nil {
		result["display_name"] = string(*obj.DisplayName)
	}

	return result
}

func (s *DatabaseToolsDatabaseToolsMcpServerResourceCrud) mapToDatabaseToolsMcpServerCustomRole(fieldKeyFormat string) (oci_database_tools.DatabaseToolsMcpServerCustomRole, error) {
	result := oci_database_tools.DatabaseToolsMcpServerCustomRole{}

	if description, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "description")); ok {
		tmp := description.(string)
		result.Description = &tmp
	}

	if displayName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "display_name")); ok {
		tmp := displayName.(string)
		result.DisplayName = &tmp
	}

	return result, nil
}

func DatabaseToolsMcpServerCustomRoleToMap(obj oci_database_tools.DatabaseToolsMcpServerCustomRole) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Description != nil {
		result["description"] = string(*obj.Description)
	}

	if obj.DisplayName != nil {
		result["display_name"] = string(*obj.DisplayName)
	}

	return result
}

func DatabaseToolsMcpServerEndpointToMap(obj oci_database_tools.DatabaseToolsMcpServerEndpoint) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Endpoint != nil {
		result["endpoint"] = string(*obj.Endpoint)
	}

	result["type"] = string(obj.Type)

	return result
}

func DatabaseToolsMcpServerRelatedResourceToMap(obj *oci_database_tools.DatabaseToolsMcpServerRelatedResource) map[string]interface{} {
	result := map[string]interface{}{}

	result["entity_type"] = string(obj.EntityType)

	if obj.Identifier != nil {
		result["identifier"] = string(*obj.Identifier)
	}

	return result
}

func (s *DatabaseToolsDatabaseToolsMcpServerResourceCrud) mapToDatabaseToolsMcpServerStorage(fieldKeyFormat string) (oci_database_tools.DatabaseToolsMcpServerStorage, error) {
	var baseObject oci_database_tools.DatabaseToolsMcpServerStorage
	//discriminator
	typeRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "type"))
	var type_ string
	if ok {
		type_ = typeRaw.(string)
	} else {
		type_ = "" // default value
	}
	switch strings.ToLower(type_) {
	case strings.ToLower("NONE"):
		details := oci_database_tools.DatabaseToolsMcpServerStorageNone{}
		baseObject = details
	case strings.ToLower("OBJECT_STORAGE"):
		details := oci_database_tools.DatabaseToolsMcpServerStorageObjectStorage{}
		if bucket, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "bucket")); ok {
			if tmpList := bucket.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "bucket"), 0)
				tmp, err := s.mapToDatabaseToolsMcpServerStorageObjectStorageBucket(fieldKeyFormatNextLevel)
				if err != nil {
					return details, fmt.Errorf("unable to convert bucket, encountered error: %v", err)
				}
				details.Bucket = &tmp
			}
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown type '%v' was specified", type_)
	}
	return baseObject, nil
}

func DatabaseToolsMcpServerStorageToMap(obj *oci_database_tools.DatabaseToolsMcpServerStorage) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (*obj).(type) {
	case oci_database_tools.DatabaseToolsMcpServerStorageNone:
		result["type"] = "NONE"
	case oci_database_tools.DatabaseToolsMcpServerStorageObjectStorage:
		result["type"] = "OBJECT_STORAGE"

		if v.Bucket != nil {
			result["bucket"] = []interface{}{DatabaseToolsMcpServerStorageObjectStorageBucketToMap(v.Bucket)}
		}
	default:
		log.Printf("[WARN] Received 'type' of unknown type %v", *obj)
		return nil
	}

	return result
}

func (s *DatabaseToolsDatabaseToolsMcpServerResourceCrud) mapToDatabaseToolsMcpServerStorageObjectStorageBucket(fieldKeyFormat string) (oci_database_tools.DatabaseToolsMcpServerStorageObjectStorageBucket, error) {
	result := oci_database_tools.DatabaseToolsMcpServerStorageObjectStorageBucket{}

	if bucket, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "bucket")); ok {
		tmp := bucket.(string)
		result.BucketName = &tmp
	}

	if namespace, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "namespace")); ok {
		tmp := namespace.(string)
		result.Namespace = &tmp
	}

	return result, nil
}

func DatabaseToolsMcpServerStorageObjectStorageBucketToMap(obj *oci_database_tools.DatabaseToolsMcpServerStorageObjectStorageBucket) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.BucketName != nil {
		result["bucket"] = string(*obj.BucketName)
	}

	if obj.Namespace != nil {
		result["namespace"] = string(*obj.Namespace)
	}

	return result
}

func DatabaseToolsMcpServerSummaryToMap(obj oci_database_tools.DatabaseToolsMcpServerSummary) map[string]interface{} {
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
	case oci_database_tools.DatabaseToolsMcpServerSummaryDefault:
		result["type"] = "DEFAULT"

		if v.DomainAppId != nil {
			result["domain_app_id"] = string(*v.DomainAppId)
		}

		if v.DomainId != nil {
			result["domain_id"] = string(*v.DomainId)
		}

		if v.Storage != nil {
			storageArray := []interface{}{}
			if storageMap := DatabaseToolsMcpServerStorageToMap(&v.Storage); storageMap != nil {
				storageArray = append(storageArray, storageMap)
			}
			result["storage"] = storageArray
		}
	default:
		log.Printf("[WARN] Received 'type' of unknown type %v", obj)
		return nil
	}

	return result
}

func (s *DatabaseToolsDatabaseToolsMcpServerResourceCrud) mapToResourceLock(fieldKeyFormat string) (oci_database_tools.ResourceLock, error) {
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

func DbtoolsMcpServerResourceLockToMap(obj oci_database_tools.ResourceLock) map[string]interface{} {
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

func (s *DatabaseToolsDatabaseToolsMcpServerResourceCrud) populateTopLevelPolymorphicCreateDatabaseToolsMcpServerRequest(request *oci_database_tools.CreateDatabaseToolsMcpServerRequest) error {
	//discriminator
	typeRaw, ok := s.D.GetOkExists("type")
	var type_ string
	if ok {
		type_ = typeRaw.(string)
	} else {
		type_ = "" // default value
	}
	switch strings.ToLower(type_) {
	case strings.ToLower("DEFAULT"):
		details := oci_database_tools.CreateDatabaseToolsMcpServerDefaultDetails{}
		if domainId, ok := s.D.GetOkExists("domain_id"); ok {
			tmp := domainId.(string)
			details.DomainId = &tmp
		}
		if storage, ok := s.D.GetOkExists("storage"); ok {
			if tmpList := storage.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "storage", 0)
				tmp, err := s.mapToDatabaseToolsMcpServerStorage(fieldKeyFormat)
				if err != nil {
					return err
				}
				details.Storage = tmp
			}
		}
		if accessTokenExpiryInSeconds, ok := s.D.GetOkExists("access_token_expiry_in_seconds"); ok {
			tmp := accessTokenExpiryInSeconds.(int)
			details.AccessTokenExpiryInSeconds = &tmp
		}
		if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
			tmp := compartmentId.(string)
			details.CompartmentId = &tmp
		}
		if customRoles, ok := s.D.GetOkExists("custom_roles"); ok {
			interfaces := customRoles.([]interface{})
			tmp := make([]oci_database_tools.DatabaseToolsMcpServerCustomRole, len(interfaces))
			for i := range interfaces {
				stateDataIndex := i
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "custom_roles", stateDataIndex)
				converted, err := s.mapToDatabaseToolsMcpServerCustomRole(fieldKeyFormat)
				if err != nil {
					return err
				}
				tmp[i] = converted
			}
			if len(tmp) != 0 || s.D.HasChange("custom_roles") {
				details.CustomRoles = tmp
			}
		}
		if databaseToolsConnectionId, ok := s.D.GetOkExists("database_tools_connection_id"); ok {
			tmp := databaseToolsConnectionId.(string)
			details.DatabaseToolsConnectionId = &tmp
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
		if refreshTokenExpiryInSeconds, ok := s.D.GetOkExists("refresh_token_expiry_in_seconds"); ok {
			tmp := refreshTokenExpiryInSeconds.(int)
			details.RefreshTokenExpiryInSeconds = &tmp
		}
		if runtimeIdentity, ok := s.D.GetOkExists("runtime_identity"); ok {
			details.RuntimeIdentity = oci_database_tools.DatabaseToolsMcpServerRuntimeIdentityEnum(runtimeIdentity.(string))
		}
		request.CreateDatabaseToolsMcpServerDetails = details
	default:
		return fmt.Errorf("unknown type '%v' was specified", type_)
	}
	return nil
}

func (s *DatabaseToolsDatabaseToolsMcpServerResourceCrud) populateTopLevelPolymorphicUpdateDatabaseToolsMcpServerRequest(request *oci_database_tools.UpdateDatabaseToolsMcpServerRequest) error {
	//discriminator
	typeRaw, ok := s.D.GetOkExists("type")
	var type_ string
	if ok {
		type_ = typeRaw.(string)
	} else {
		type_ = "" // default value
	}
	switch strings.ToLower(type_) {
	case strings.ToLower("DEFAULT"):
		details := oci_database_tools.UpdateDatabaseToolsMcpServerDetailsDefault{}
		if storage, ok := s.D.GetOkExists("storage"); ok {
			if tmpList := storage.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "storage", 0)
				tmp, err := s.mapToDatabaseToolsMcpServerStorage(fieldKeyFormat)
				if err != nil {
					return err
				}
				details.Storage = tmp
			}
		}
		if accessTokenExpiryInSeconds, ok := s.D.GetOkExists("access_token_expiry_in_seconds"); ok {
			tmp := accessTokenExpiryInSeconds.(int)
			details.AccessTokenExpiryInSeconds = &tmp
		}
		if customRoles, ok := s.D.GetOkExists("custom_roles"); ok {
			interfaces := customRoles.([]interface{})
			tmp := make([]oci_database_tools.DatabaseToolsMcpServerCustomRole, len(interfaces))
			for i := range interfaces {
				stateDataIndex := i
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "custom_roles", stateDataIndex)
				converted, err := s.mapToDatabaseToolsMcpServerCustomRole(fieldKeyFormat)
				if err != nil {
					return err
				}
				tmp[i] = converted
			}
			if len(tmp) != 0 || s.D.HasChange("custom_roles") {
				details.CustomRoles = tmp
			}
		}
		tmp := s.D.Id()
		request.DatabaseToolsMcpServerId = &tmp
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
		if refreshTokenExpiryInSeconds, ok := s.D.GetOkExists("refresh_token_expiry_in_seconds"); ok {
			tmp := refreshTokenExpiryInSeconds.(int)
			details.RefreshTokenExpiryInSeconds = &tmp
		}
		request.UpdateDatabaseToolsMcpServerDetails = details
	default:
		return fmt.Errorf("unknown type '%v' was specified", type_)
	}
	return nil
}

func (s *DatabaseToolsDatabaseToolsMcpServerResourceCrud) updateCompartment(ctx context.Context, compartment interface{}) error {
	changeCompartmentRequest := oci_database_tools.ChangeDatabaseToolsMcpServerCompartmentRequest{}

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	idTmp := s.D.Id()
	changeCompartmentRequest.DatabaseToolsMcpServerId = &idTmp

	if isLockOverride, ok := s.D.GetOkExists("is_lock_override"); ok {
		tmp := isLockOverride.(bool)
		changeCompartmentRequest.IsLockOverride = &tmp
	}

	changeCompartmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database_tools")

	response, err := s.Client.ChangeDatabaseToolsMcpServerCompartment(ctx, changeCompartmentRequest)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getDatabaseToolsMcpServerFromWorkRequest(ctx, workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database_tools"), oci_database_tools.ActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}
