// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package database_tools

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"
	"github.com/terraform-providers/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"

	oci_common "github.com/oracle/oci-go-sdk/v58/common"
	oci_database_tools "github.com/oracle/oci-go-sdk/v58/databasetools"
)

func DatabaseToolsDatabaseToolsConnectionResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createDatabaseToolsDatabaseToolsConnection,
		Read:     readDatabaseToolsDatabaseToolsConnection,
		Update:   updateDatabaseToolsDatabaseToolsConnection,
		Delete:   deleteDatabaseToolsDatabaseToolsConnection,
		Schema: map[string]*schema.Schema{
			// Required
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
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
					"ORACLE_DATABASE",
				}, true),
			},

			// Optional
			"advanced_properties": {
				Type:     schema.TypeMap,
				Optional: true,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"connection_string": {
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
			"freeform_tags": {
				Type:     schema.TypeMap,
				Optional: true,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"key_stores": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional
						"key_store_content": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							MaxItems: 1,
							MinItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required
									"value_type": {
										Type:             schema.TypeString,
										Required:         true,
										DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
										ValidateFunc: validation.StringInSlice([]string{
											"SECRETID",
										}, true),
									},

									// Optional
									"secret_id": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},

									// Computed
								},
							},
						},
						"key_store_password": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							MaxItems: 1,
							MinItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required
									"value_type": {
										Type:             schema.TypeString,
										Required:         true,
										DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
										ValidateFunc: validation.StringInSlice([]string{
											"SECRETID",
										}, true),
									},

									// Optional
									"secret_id": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},

									// Computed
								},
							},
						},
						"key_store_type": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},

						// Computed
					},
				},
			},
			"private_endpoint_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"related_resource": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"entity_type": {
							Type:     schema.TypeString,
							Required: true,
						},
						"identifier": {
							Type:     schema.TypeString,
							Required: true,
						},

						// Optional

						// Computed
					},
				},
			},
			"user_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"user_password": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"value_type": {
							Type:             schema.TypeString,
							Required:         true,
							DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
							ValidateFunc: validation.StringInSlice([]string{
								"SECRETID",
							}, true),
						},

						// Optional
						"secret_id": {
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

func createDatabaseToolsDatabaseToolsConnection(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseToolsDatabaseToolsConnectionResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseToolsClient()

	return tfresource.CreateResource(d, sync)
}

func readDatabaseToolsDatabaseToolsConnection(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseToolsDatabaseToolsConnectionResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseToolsClient()

	return tfresource.ReadResource(sync)
}

func updateDatabaseToolsDatabaseToolsConnection(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseToolsDatabaseToolsConnectionResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseToolsClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteDatabaseToolsDatabaseToolsConnection(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseToolsDatabaseToolsConnectionResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseToolsClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type DatabaseToolsDatabaseToolsConnectionResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_database_tools.DatabaseToolsClient
	Res                    *oci_database_tools.DatabaseToolsConnection
	DisableNotFoundRetries bool
}

func (s *DatabaseToolsDatabaseToolsConnectionResourceCrud) ID() string {
	databaseToolsConnection := *s.Res
	return *databaseToolsConnection.GetId()
}

func (s *DatabaseToolsDatabaseToolsConnectionResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_database_tools.LifecycleStateCreating),
	}
}

func (s *DatabaseToolsDatabaseToolsConnectionResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_database_tools.LifecycleStateActive),
	}
}

func (s *DatabaseToolsDatabaseToolsConnectionResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_database_tools.LifecycleStateDeleting),
	}
}

func (s *DatabaseToolsDatabaseToolsConnectionResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_database_tools.LifecycleStateDeleted),
	}
}

func (s *DatabaseToolsDatabaseToolsConnectionResourceCrud) Create() error {
	request := oci_database_tools.CreateDatabaseToolsConnectionRequest{}
	err := s.populateTopLevelPolymorphicCreateDatabaseToolsConnectionRequest(&request)
	if err != nil {
		return err
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database_tools")

	response, err := s.Client.CreateDatabaseToolsConnection(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getDatabaseToolsConnectionFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database_tools"), oci_database_tools.ActionTypeCreated, s.D.Timeout(schema.TimeoutCreate))
}

func (s *DatabaseToolsDatabaseToolsConnectionResourceCrud) getDatabaseToolsConnectionFromWorkRequest(workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_database_tools.ActionTypeEnum, timeout time.Duration) error {

	// Wait until it finishes
	databaseToolsConnectionId, err := databaseToolsConnectionWaitForWorkRequest(workId, "databasetoolsconnection", // ""database_tools", // Rashik Bhasin indicates that it should match the entityType in the Work request Response
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.Client)

	if err != nil {
		return err
	}
	s.D.SetId(*databaseToolsConnectionId)

	return s.Get()
}

func databaseToolsConnectionWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
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

func databaseToolsConnectionWaitForWorkRequest(wId *string, entityType string, action oci_database_tools.ActionTypeEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_database_tools.DatabaseToolsClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "database_tools")
	retryPolicy.ShouldRetryOperation = databaseToolsConnectionWorkRequestShouldRetryFunc(timeout)

	response := oci_database_tools.GetWorkRequestResponse{}
	stateConf := &resource.StateChangeConf{
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
			response, err = client.GetWorkRequest(context.Background(),
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
		return nil, getErrorFromDatabaseToolsDatabaseToolsConnectionWorkRequest(client, wId, retryPolicy, entityType, action)
	}

	return identifier, nil
}

func getErrorFromDatabaseToolsDatabaseToolsConnectionWorkRequest(client *oci_database_tools.DatabaseToolsClient, workId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_database_tools.ActionTypeEnum) error {
	response, err := client.ListWorkRequestErrors(context.Background(),
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

func (s *DatabaseToolsDatabaseToolsConnectionResourceCrud) Get() error {
	request := oci_database_tools.GetDatabaseToolsConnectionRequest{}

	tmp := s.D.Id()
	request.DatabaseToolsConnectionId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database_tools")

	response, err := s.Client.GetDatabaseToolsConnection(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.DatabaseToolsConnection
	return nil
}

func (s *DatabaseToolsDatabaseToolsConnectionResourceCrud) Update() error {
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(compartment)
			if err != nil {
				return err
			}
		}
	}
	request := oci_database_tools.UpdateDatabaseToolsConnectionRequest{}
	err := s.populateTopLevelPolymorphicUpdateDatabaseToolsConnectionRequest(&request)
	if err != nil {
		return err
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database_tools")

	response, err := s.Client.UpdateDatabaseToolsConnection(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getDatabaseToolsConnectionFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database_tools"), oci_database_tools.ActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}

func (s *DatabaseToolsDatabaseToolsConnectionResourceCrud) Delete() error {
	request := oci_database_tools.DeleteDatabaseToolsConnectionRequest{}

	tmp := s.D.Id()
	request.DatabaseToolsConnectionId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database_tools")

	response, err := s.Client.DeleteDatabaseToolsConnection(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	// Wait until it finishes
	_, delWorkRequestErr := databaseToolsConnectionWaitForWorkRequest(workId, "databasetoolsconnection",
		oci_database_tools.ActionTypeDeleted, s.D.Timeout(schema.TimeoutDelete), s.DisableNotFoundRetries, s.Client)
	return delWorkRequestErr
}

func (s *DatabaseToolsDatabaseToolsConnectionResourceCrud) SetData() error {
	switch v := (*s.Res).(type) {
	case oci_database_tools.DatabaseToolsConnectionOracleDatabase:
		s.D.Set("type", "ORACLE_DATABASE")

		s.D.Set("advanced_properties", v.AdvancedProperties)

		if v.ConnectionString != nil {
			s.D.Set("connection_string", *v.ConnectionString)
		}

		keyStores := []interface{}{}
		for _, item := range v.KeyStores {
			keyStores = append(keyStores, DatabaseToolsKeyStoreToMap(item))
		}
		s.D.Set("key_stores", keyStores)

		if v.PrivateEndpointId != nil {
			s.D.Set("private_endpoint_id", *v.PrivateEndpointId)
		}

		if v.RelatedResource != nil {
			s.D.Set("related_resource", []interface{}{DatabaseToolsRelatedResourceToMap(v.RelatedResource)})
		} else {
			s.D.Set("related_resource", nil)
		}

		if v.UserName != nil {
			s.D.Set("user_name", *v.UserName)
		}

		if v.UserPassword != nil {
			userPasswordArray := []interface{}{}
			if userPasswordMap := DatabaseToolsUserPasswordToMap(&v.UserPassword); userPasswordMap != nil {
				userPasswordArray = append(userPasswordArray, userPasswordMap)
			}
			s.D.Set("user_password", userPasswordArray)
		} else {
			s.D.Set("user_password", nil)
		}

		if v.CompartmentId != nil {
			s.D.Set("compartment_id", *v.CompartmentId)
		}

		if v.DefinedTags != nil {
			s.D.Set("defined_tags", tfresource.DefinedTagsToMap(v.DefinedTags))
		}

		if v.DisplayName != nil {
			s.D.Set("display_name", *v.DisplayName)
		}

		s.D.Set("freeform_tags", v.FreeformTags)

		if v.Id != nil {
			s.D.Set("id", *v.Id)
		}

		if v.LifecycleDetails != nil {
			s.D.Set("lifecycle_details", *v.LifecycleDetails)
		}

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

func (s *DatabaseToolsDatabaseToolsConnectionResourceCrud) mapToCreateDatabaseToolsRelatedResourceDetails(fieldKeyFormat string) (oci_database_tools.CreateDatabaseToolsRelatedResourceDetails, error) {
	result := oci_database_tools.CreateDatabaseToolsRelatedResourceDetails{}

	if entityType, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "entity_type")); ok {
		result.EntityType = oci_database_tools.RelatedResourceEntityTypeEnum(entityType.(string))
	}

	if identifier, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "identifier")); ok {
		tmp := identifier.(string)
		result.Identifier = &tmp
	}

	return result, nil
}

func CreateDatabaseToolsRelatedResourceDetailsToMap(obj *oci_database_tools.CreateDatabaseToolsRelatedResourceDetails) map[string]interface{} {
	result := map[string]interface{}{}

	result["entity_type"] = string(obj.EntityType)

	if obj.Identifier != nil {
		result["identifier"] = string(*obj.Identifier)
	}

	return result
}

func DatabaseToolsConnectionSummaryToMap(obj oci_database_tools.DatabaseToolsConnectionSummary) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (obj).(type) {
	case oci_database_tools.DatabaseToolsConnectionOracleDatabaseSummary:
		result["type"] = "ORACLE_DATABASE"

		result["advanced_properties"] = v.AdvancedProperties

		if v.ConnectionString != nil {
			result["connection_string"] = string(*v.ConnectionString)
		}

		keyStores := []interface{}{}
		for _, item := range v.KeyStores {
			keyStores = append(keyStores, DatabaseToolsKeyStoreSummaryToMap(item))
		}
		result["key_stores"] = keyStores

		if v.PrivateEndpointId != nil {
			result["private_endpoint_id"] = string(*v.PrivateEndpointId)
		}

		if v.RelatedResource != nil {
			result["related_resource"] = []interface{}{DatabaseToolsRelatedResourceToMap(v.RelatedResource)}
		}

		if v.UserName != nil {
			result["user_name"] = string(*v.UserName)
		}

		if v.UserPassword != nil {
			userPasswordArray := []interface{}{}
			if userPasswordMap := DatabaseToolsUserPasswordSummaryToMap(&v.UserPassword); userPasswordMap != nil {
				userPasswordArray = append(userPasswordArray, userPasswordMap)
			}
			result["user_password"] = userPasswordArray
		}

		// frobert 2021-08-06
		// These missing fields correspond to the fields defined in parent DatabaseToolsConnectionSummary as we
		// only get the field defined in child DatabaseToolsConnectionOracleDatabaseSummary.

		if v.Id != nil {
			result["id"] = string(*v.Id)
		}

		if v.DisplayName != nil {
			result["display_name"] = string(*v.DisplayName)
		}

		if v.CompartmentId != nil {
			result["compartment_id"] = string(*v.CompartmentId)
		}

		if v.TimeCreated != nil {
			result["time_created"] = v.TimeCreated.String()
		}

		if v.TimeUpdated != nil {
			result["time_updated"] = v.TimeUpdated.String()
		}

		if v.DefinedTags != nil {
			result["defined_tags"] = tfresource.DefinedTagsToMap(v.DefinedTags)
		}

		if v.SystemTags != nil {
			result["system_tags"] = tfresource.SystemTagsToMap(v.SystemTags)
		}

		if v.FreeformTags != nil {
			result["freeform_tags"] = v.FreeformTags
		}

		if v.PrivateEndpointId != nil {
			result["private_endpoint_id"] = string(*v.PrivateEndpointId)
		}

		result["state"] = string(v.LifecycleState)

		if v.LifecycleDetails != nil {
			result["lifecycle_details"] = string(*v.LifecycleDetails)
		}

	default:
		log.Printf("[WARN] Received 'type' of unknown type %v", obj)
		return nil
	}

	return result
}

func (s *DatabaseToolsDatabaseToolsConnectionResourceCrud) mapToDatabaseToolsKeyStore(fieldKeyFormat string) (oci_database_tools.DatabaseToolsKeyStore, error) {
	result := oci_database_tools.DatabaseToolsKeyStore{}

	if keyStoreContent, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "key_store_content")); ok {
		if tmpList := keyStoreContent.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "key_store_content"), 0)
			tmp, err := s.mapToDatabaseToolsKeyStoreContent(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert key_store_content, encountered error: %v", err)
			}
			result.KeyStoreContent = tmp
		}
	}

	if keyStorePassword, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "key_store_password")); ok {
		if tmpList := keyStorePassword.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "key_store_password"), 0)
			tmp, err := s.mapToDatabaseToolsKeyStorePassword(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert key_store_password, encountered error: %v", err)
			}
			result.KeyStorePassword = tmp
		}
	}

	if keyStoreType, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "key_store_type")); ok {
		result.KeyStoreType = oci_database_tools.KeyStoreTypeEnum(keyStoreType.(string))
	}

	return result, nil
}

func DatabaseToolsKeyStoreToMap(obj oci_database_tools.DatabaseToolsKeyStore) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.KeyStoreContent != nil {
		keyStoreContentArray := []interface{}{}
		if keyStoreContentMap := DatabaseToolsKeyStoreContentToMap(&obj.KeyStoreContent); keyStoreContentMap != nil {
			keyStoreContentArray = append(keyStoreContentArray, keyStoreContentMap)
		}
		result["key_store_content"] = keyStoreContentArray
	}

	if obj.KeyStorePassword != nil {
		keyStorePasswordArray := []interface{}{}
		if keyStorePasswordMap := DatabaseToolsKeyStorePasswordToMap(&obj.KeyStorePassword); keyStorePasswordMap != nil {
			keyStorePasswordArray = append(keyStorePasswordArray, keyStorePasswordMap)
		}
		result["key_store_password"] = keyStorePasswordArray
	}

	result["key_store_type"] = string(obj.KeyStoreType)

	return result
}

func (s *DatabaseToolsDatabaseToolsConnectionResourceCrud) mapToDatabaseToolsKeyStoreContent(fieldKeyFormat string) (oci_database_tools.DatabaseToolsKeyStoreContent, error) {
	var baseObject oci_database_tools.DatabaseToolsKeyStoreContent
	//discriminator
	valueTypeRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "value_type"))
	var valueType string
	if ok {
		valueType = valueTypeRaw.(string)
	} else {
		valueType = "" // default value
	}
	switch strings.ToLower(valueType) {
	case strings.ToLower("SECRETID"):
		details := oci_database_tools.DatabaseToolsKeyStoreContentSecretId{}
		if secretId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "secret_id")); ok {
			tmp := secretId.(string)
			details.SecretId = &tmp
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown value_type '%v' was specified", valueType)
	}
	return baseObject, nil
}

func DatabaseToolsKeyStoreContentToMap(obj *oci_database_tools.DatabaseToolsKeyStoreContent) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (*obj).(type) {
	case oci_database_tools.DatabaseToolsKeyStoreContentSecretId:
		result["value_type"] = "SECRETID"

		if v.SecretId != nil {
			result["secret_id"] = string(*v.SecretId)
		}
	default:
		log.Printf("[WARN] Received 'value_type' of unknown type %v", *obj)
		return nil
	}

	return result
}

func (s *DatabaseToolsDatabaseToolsConnectionResourceCrud) mapToDatabaseToolsKeyStoreContentDetails(fieldKeyFormat string) (oci_database_tools.DatabaseToolsKeyStoreContentDetails, error) {
	var baseObject oci_database_tools.DatabaseToolsKeyStoreContentDetails
	//discriminator
	valueTypeRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "value_type"))
	var valueType string
	if ok {
		valueType = valueTypeRaw.(string)
	} else {
		valueType = "" // default value
	}
	switch strings.ToLower(valueType) {
	case strings.ToLower("SECRETID"):
		details := oci_database_tools.DatabaseToolsKeyStoreContentSecretIdDetails{}
		if secretId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "secret_id")); ok {
			tmp := secretId.(string)
			details.SecretId = &tmp
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown value_type '%v' was specified", valueType)
	}
	return baseObject, nil
}

func DatabaseToolsKeyStoreContentDetailsToMap(obj *oci_database_tools.DatabaseToolsKeyStoreContentDetails) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (*obj).(type) {
	case oci_database_tools.DatabaseToolsKeyStoreContentSecretIdDetails:
		result["value_type"] = "SECRETID"

		if v.SecretId != nil {
			result["secret_id"] = string(*v.SecretId)
		}
	default:
		log.Printf("[WARN] Received 'value_type' of unknown type %v", *obj)
		return nil
	}

	return result
}

func (s *DatabaseToolsDatabaseToolsConnectionResourceCrud) mapToDatabaseToolsKeyStoreContentSummary(fieldKeyFormat string) (oci_database_tools.DatabaseToolsKeyStoreContentSummary, error) {
	var baseObject oci_database_tools.DatabaseToolsKeyStoreContentSummary
	//discriminator
	valueTypeRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "value_type"))
	var valueType string
	if ok {
		valueType = valueTypeRaw.(string)
	} else {
		valueType = "" // default value
	}
	switch strings.ToLower(valueType) {
	case strings.ToLower("SECRETID"):
		details := oci_database_tools.DatabaseToolsKeyStoreContentSecretIdSummary{}
		if secretId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "secret_id")); ok {
			tmp := secretId.(string)
			details.SecretId = &tmp
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown value_type '%v' was specified", valueType)
	}
	return baseObject, nil
}

func DatabaseToolsKeyStoreContentSummaryToMap(obj *oci_database_tools.DatabaseToolsKeyStoreContentSummary) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (*obj).(type) {
	case oci_database_tools.DatabaseToolsKeyStoreContentSecretIdSummary:
		result["value_type"] = "SECRETID"

		if v.SecretId != nil {
			result["secret_id"] = string(*v.SecretId)
		}
	default:
		log.Printf("[WARN] Received 'value_type' of unknown type %v", *obj)
		return nil
	}

	return result
}

func (s *DatabaseToolsDatabaseToolsConnectionResourceCrud) mapToDatabaseToolsKeyStoreDetails(fieldKeyFormat string) (oci_database_tools.DatabaseToolsKeyStoreDetails, error) {
	result := oci_database_tools.DatabaseToolsKeyStoreDetails{}

	if keyStoreContent, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "key_store_content")); ok {
		if tmpList := keyStoreContent.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "key_store_content"), 0)
			tmp, err := s.mapToDatabaseToolsKeyStoreContentDetails(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert key_store_content, encountered error: %v", err)
			}
			result.KeyStoreContent = tmp
		}
	}

	if keyStorePassword, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "key_store_password")); ok {
		if tmpList := keyStorePassword.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "key_store_password"), 0)
			tmp, err := s.mapToDatabaseToolsKeyStorePasswordDetails(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert key_store_password, encountered error: %v", err)
			}
			result.KeyStorePassword = tmp
		}
	}

	if keyStoreType, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "key_store_type")); ok {
		result.KeyStoreType = oci_database_tools.KeyStoreTypeEnum(keyStoreType.(string))
	}

	return result, nil
}

func DatabaseToolsKeyStoreDetailsToMap(obj oci_database_tools.DatabaseToolsKeyStoreDetails) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.KeyStoreContent != nil {
		keyStoreContentArray := []interface{}{}
		if keyStoreContentMap := DatabaseToolsKeyStoreContentDetailsToMap(&obj.KeyStoreContent); keyStoreContentMap != nil {
			keyStoreContentArray = append(keyStoreContentArray, keyStoreContentMap)
		}
		result["key_store_content"] = keyStoreContentArray
	}

	if obj.KeyStorePassword != nil {
		keyStorePasswordArray := []interface{}{}
		if keyStorePasswordMap := DatabaseToolsKeyStorePasswordDetailsToMap(&obj.KeyStorePassword); keyStorePasswordMap != nil {
			keyStorePasswordArray = append(keyStorePasswordArray, keyStorePasswordMap)
		}
		result["key_store_password"] = keyStorePasswordArray
	}

	result["key_store_type"] = string(obj.KeyStoreType)

	return result
}

func (s *DatabaseToolsDatabaseToolsConnectionResourceCrud) mapToDatabaseToolsKeyStorePassword(fieldKeyFormat string) (oci_database_tools.DatabaseToolsKeyStorePassword, error) {
	var baseObject oci_database_tools.DatabaseToolsKeyStorePassword
	//discriminator
	valueTypeRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "value_type"))
	var valueType string
	if ok {
		valueType = valueTypeRaw.(string)
	} else {
		valueType = "" // default value
	}
	switch strings.ToLower(valueType) {
	case strings.ToLower("SECRETID"):
		details := oci_database_tools.DatabaseToolsKeyStorePasswordSecretId{}
		if secretId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "secret_id")); ok {
			tmp := secretId.(string)
			details.SecretId = &tmp
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown value_type '%v' was specified", valueType)
	}
	return baseObject, nil
}

func DatabaseToolsKeyStorePasswordToMap(obj *oci_database_tools.DatabaseToolsKeyStorePassword) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (*obj).(type) {
	case oci_database_tools.DatabaseToolsKeyStorePasswordSecretId:
		result["value_type"] = "SECRETID"

		if v.SecretId != nil {
			result["secret_id"] = string(*v.SecretId)
		}
	default:
		log.Printf("[WARN] Received 'value_type' of unknown type %v", *obj)
		return nil
	}

	return result
}

func (s *DatabaseToolsDatabaseToolsConnectionResourceCrud) mapToDatabaseToolsKeyStorePasswordDetails(fieldKeyFormat string) (oci_database_tools.DatabaseToolsKeyStorePasswordDetails, error) {
	var baseObject oci_database_tools.DatabaseToolsKeyStorePasswordDetails
	//discriminator
	valueTypeRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "value_type"))
	var valueType string
	if ok {
		valueType = valueTypeRaw.(string)
	} else {
		valueType = "" // default value
	}
	switch strings.ToLower(valueType) {
	case strings.ToLower("SECRETID"):
		details := oci_database_tools.DatabaseToolsKeyStorePasswordSecretIdDetails{}
		if secretId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "secret_id")); ok {
			tmp := secretId.(string)
			details.SecretId = &tmp
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown value_type '%v' was specified", valueType)
	}
	return baseObject, nil
}

func DatabaseToolsKeyStorePasswordDetailsToMap(obj *oci_database_tools.DatabaseToolsKeyStorePasswordDetails) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (*obj).(type) {
	case oci_database_tools.DatabaseToolsKeyStorePasswordSecretIdDetails:
		result["value_type"] = "SECRETID"

		if v.SecretId != nil {
			result["secret_id"] = string(*v.SecretId)
		}
	default:
		log.Printf("[WARN] Received 'value_type' of unknown type %v", *obj)
		return nil
	}

	return result
}

func (s *DatabaseToolsDatabaseToolsConnectionResourceCrud) mapToDatabaseToolsKeyStorePasswordSummary(fieldKeyFormat string) (oci_database_tools.DatabaseToolsKeyStorePasswordSummary, error) {
	var baseObject oci_database_tools.DatabaseToolsKeyStorePasswordSummary
	//discriminator
	valueTypeRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "value_type"))
	var valueType string
	if ok {
		valueType = valueTypeRaw.(string)
	} else {
		valueType = "" // default value
	}
	switch strings.ToLower(valueType) {
	case strings.ToLower("SECRETID"):
		details := oci_database_tools.DatabaseToolsKeyStorePasswordSecretIdSummary{}
		if secretId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "secret_id")); ok {
			tmp := secretId.(string)
			details.SecretId = &tmp
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown value_type '%v' was specified", valueType)
	}
	return baseObject, nil
}

func DatabaseToolsKeyStorePasswordSummaryToMap(obj *oci_database_tools.DatabaseToolsKeyStorePasswordSummary) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (*obj).(type) {
	case oci_database_tools.DatabaseToolsKeyStorePasswordSecretIdSummary:
		result["value_type"] = "SECRETID"

		if v.SecretId != nil {
			result["secret_id"] = string(*v.SecretId)
		}
	default:
		log.Printf("[WARN] Received 'value_type' of unknown type %v", *obj)
		return nil
	}

	return result
}

func (s *DatabaseToolsDatabaseToolsConnectionResourceCrud) mapToDatabaseToolsKeyStoreSummary(fieldKeyFormat string) (oci_database_tools.DatabaseToolsKeyStoreSummary, error) {
	result := oci_database_tools.DatabaseToolsKeyStoreSummary{}

	if keyStoreContent, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "key_store_content")); ok {
		if tmpList := keyStoreContent.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "key_store_content"), 0)
			tmp, err := s.mapToDatabaseToolsKeyStoreContentSummary(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert key_store_content, encountered error: %v", err)
			}
			result.KeyStoreContent = tmp
		}
	}

	if keyStorePassword, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "key_store_password")); ok {
		if tmpList := keyStorePassword.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "key_store_password"), 0)
			tmp, err := s.mapToDatabaseToolsKeyStorePasswordSummary(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert key_store_password, encountered error: %v", err)
			}
			result.KeyStorePassword = tmp
		}
	}

	if keyStoreType, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "key_store_type")); ok {
		result.KeyStoreType = oci_database_tools.KeyStoreTypeEnum(keyStoreType.(string))
	}

	return result, nil
}

func DatabaseToolsKeyStoreSummaryToMap(obj oci_database_tools.DatabaseToolsKeyStoreSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.KeyStoreContent != nil {
		keyStoreContentArray := []interface{}{}
		if keyStoreContentMap := DatabaseToolsKeyStoreContentSummaryToMap(&obj.KeyStoreContent); keyStoreContentMap != nil {
			keyStoreContentArray = append(keyStoreContentArray, keyStoreContentMap)
		}
		result["key_store_content"] = keyStoreContentArray
	}

	if obj.KeyStorePassword != nil {
		keyStorePasswordArray := []interface{}{}
		if keyStorePasswordMap := DatabaseToolsKeyStorePasswordSummaryToMap(&obj.KeyStorePassword); keyStorePasswordMap != nil {
			keyStorePasswordArray = append(keyStorePasswordArray, keyStorePasswordMap)
		}
		result["key_store_password"] = keyStorePasswordArray
	}

	result["key_store_type"] = string(obj.KeyStoreType)

	return result
}

func (s *DatabaseToolsDatabaseToolsConnectionResourceCrud) mapToDatabaseToolsRelatedResource(fieldKeyFormat string) (oci_database_tools.DatabaseToolsRelatedResource, error) {
	result := oci_database_tools.DatabaseToolsRelatedResource{}

	if entityType, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "entity_type")); ok {
		result.EntityType = oci_database_tools.RelatedResourceEntityTypeEnum(entityType.(string))
	}

	if identifier, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "identifier")); ok {
		tmp := identifier.(string)
		result.Identifier = &tmp
	}

	return result, nil
}

func DatabaseToolsRelatedResourceToMap(obj *oci_database_tools.DatabaseToolsRelatedResource) map[string]interface{} {
	result := map[string]interface{}{}

	result["entity_type"] = string(obj.EntityType)

	if obj.Identifier != nil {
		result["identifier"] = string(*obj.Identifier)
	}

	return result
}

func (s *DatabaseToolsDatabaseToolsConnectionResourceCrud) mapToDatabaseToolsUserPassword(fieldKeyFormat string) (oci_database_tools.DatabaseToolsUserPassword, error) {
	var baseObject oci_database_tools.DatabaseToolsUserPassword
	//discriminator
	valueTypeRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "value_type"))
	var valueType string
	if ok {
		valueType = valueTypeRaw.(string)
	} else {
		valueType = "" // default value
	}
	switch strings.ToLower(valueType) {
	case strings.ToLower("SECRETID"):
		details := oci_database_tools.DatabaseToolsUserPasswordSecretId{}
		if secretId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "secret_id")); ok {
			tmp := secretId.(string)
			details.SecretId = &tmp
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown value_type '%v' was specified", valueType)
	}
	return baseObject, nil
}

func DatabaseToolsUserPasswordToMap(obj *oci_database_tools.DatabaseToolsUserPassword) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (*obj).(type) {
	case oci_database_tools.DatabaseToolsUserPasswordSecretId:
		result["value_type"] = "SECRETID"

		if v.SecretId != nil {
			result["secret_id"] = string(*v.SecretId)
		}
	default:
		log.Printf("[WARN] Received 'value_type' of unknown type %v", *obj)
		return nil
	}

	return result
}

func (s *DatabaseToolsDatabaseToolsConnectionResourceCrud) mapToDatabaseToolsUserPasswordDetails(fieldKeyFormat string) (oci_database_tools.DatabaseToolsUserPasswordDetails, error) {
	var baseObject oci_database_tools.DatabaseToolsUserPasswordDetails
	//discriminator
	valueTypeRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "value_type"))
	var valueType string
	if ok {
		valueType = valueTypeRaw.(string)
	} else {
		valueType = "" // default value
	}
	switch strings.ToLower(valueType) {
	case strings.ToLower("SECRETID"):
		details := oci_database_tools.DatabaseToolsUserPasswordSecretIdDetails{}
		if secretId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "secret_id")); ok {
			tmp := secretId.(string)
			details.SecretId = &tmp
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown value_type '%v' was specified", valueType)
	}
	return baseObject, nil
}

func DatabaseToolsUserPasswordDetailsToMap(obj *oci_database_tools.DatabaseToolsUserPasswordDetails) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (*obj).(type) {
	case oci_database_tools.DatabaseToolsUserPasswordSecretIdDetails:
		result["value_type"] = "SECRETID"

		if v.SecretId != nil {
			result["secret_id"] = string(*v.SecretId)
		}
	default:
		log.Printf("[WARN] Received 'value_type' of unknown type %v", *obj)
		return nil
	}

	return result
}

func (s *DatabaseToolsDatabaseToolsConnectionResourceCrud) mapToDatabaseToolsUserPasswordSummary(fieldKeyFormat string) (oci_database_tools.DatabaseToolsUserPasswordSummary, error) {
	var baseObject oci_database_tools.DatabaseToolsUserPasswordSummary
	//discriminator
	valueTypeRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "value_type"))
	var valueType string
	if ok {
		valueType = valueTypeRaw.(string)
	} else {
		valueType = "" // default value
	}
	switch strings.ToLower(valueType) {
	case strings.ToLower("SECRETID"):
		details := oci_database_tools.DatabaseToolsUserPasswordSecretIdSummary{}
		if secretId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "secret_id")); ok {
			tmp := secretId.(string)
			details.SecretId = &tmp
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown value_type '%v' was specified", valueType)
	}
	return baseObject, nil
}

func DatabaseToolsUserPasswordSummaryToMap(obj *oci_database_tools.DatabaseToolsUserPasswordSummary) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (*obj).(type) {
	case oci_database_tools.DatabaseToolsUserPasswordSecretIdSummary:
		result["value_type"] = "SECRETID"

		if v.SecretId != nil {
			result["secret_id"] = string(*v.SecretId)
		}
	default:
		log.Printf("[WARN] Received 'value_type' of unknown type %v", *obj)
		return nil
	}

	return result
}

func (s *DatabaseToolsDatabaseToolsConnectionResourceCrud) mapToUpdateDatabaseToolsRelatedResourceDetails(fieldKeyFormat string) (oci_database_tools.UpdateDatabaseToolsRelatedResourceDetails, error) {
	result := oci_database_tools.UpdateDatabaseToolsRelatedResourceDetails{}

	if entityType, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "entity_type")); ok {
		result.EntityType = oci_database_tools.RelatedResourceEntityTypeEnum(entityType.(string))
	}

	if identifier, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "identifier")); ok {
		tmp := identifier.(string)
		result.Identifier = &tmp
	}

	return result, nil
}

func UpdateDatabaseToolsRelatedResourceDetailsToMap(obj *oci_database_tools.UpdateDatabaseToolsRelatedResourceDetails) map[string]interface{} {
	result := map[string]interface{}{}

	result["entity_type"] = string(obj.EntityType)

	if obj.Identifier != nil {
		result["identifier"] = string(*obj.Identifier)
	}

	return result
}

func (s *DatabaseToolsDatabaseToolsConnectionResourceCrud) populateTopLevelPolymorphicCreateDatabaseToolsConnectionRequest(request *oci_database_tools.CreateDatabaseToolsConnectionRequest) error {
	//discriminator
	typeRaw, ok := s.D.GetOkExists("type")
	var type_ string
	if ok {
		type_ = typeRaw.(string)
	} else {
		type_ = "" // default value
	}
	switch strings.ToLower(type_) {
	case strings.ToLower("ORACLE_DATABASE"):
		details := oci_database_tools.CreateDatabaseToolsConnectionOracleDatabaseDetails{}
		if advancedProperties, ok := s.D.GetOkExists("advanced_properties"); ok {
			details.AdvancedProperties = utils.ObjectMapToStringMap(advancedProperties.(map[string]interface{}))
		}
		if connectionString, ok := s.D.GetOkExists("connection_string"); ok {
			tmp := connectionString.(string)
			details.ConnectionString = &tmp
		}
		if keyStores, ok := s.D.GetOkExists("key_stores"); ok {
			interfaces := keyStores.([]interface{})
			tmp := make([]oci_database_tools.DatabaseToolsKeyStoreDetails, len(interfaces))
			for i := range interfaces {
				stateDataIndex := i
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "key_stores", stateDataIndex)
				converted, err := s.mapToDatabaseToolsKeyStoreDetails(fieldKeyFormat)
				if err != nil {
					return err
				}
				tmp[i] = converted
			}
			if len(tmp) != 0 || s.D.HasChange("key_stores") {
				details.KeyStores = tmp
			}
		}
		if privateEndpointId, ok := s.D.GetOkExists("private_endpoint_id"); ok {
			tmp := privateEndpointId.(string)
			details.PrivateEndpointId = &tmp
		}
		if relatedResource, ok := s.D.GetOkExists("related_resource"); ok {
			if tmpList := relatedResource.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "related_resource", 0)
				tmp, err := s.mapToCreateDatabaseToolsRelatedResourceDetails(fieldKeyFormat)
				if err != nil {
					return err
				}
				details.RelatedResource = &tmp
			}
		}
		if userName, ok := s.D.GetOkExists("user_name"); ok {
			tmp := userName.(string)
			details.UserName = &tmp
		}
		if userPassword, ok := s.D.GetOkExists("user_password"); ok {
			if tmpList := userPassword.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "user_password", 0)
				tmp, err := s.mapToDatabaseToolsUserPasswordDetails(fieldKeyFormat)
				if err != nil {
					return err
				}
				details.UserPassword = tmp
			}
		}
		if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
			tmp := compartmentId.(string)
			details.CompartmentId = &tmp
		}
		if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
			convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
			if err != nil {
				return err
			}
			details.DefinedTags = convertedDefinedTags
		}
		if displayName, ok := s.D.GetOkExists("display_name"); ok {
			tmp := displayName.(string)
			details.DisplayName = &tmp
		}
		if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
			details.FreeformTags = utils.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
		}
		request.CreateDatabaseToolsConnectionDetails = details
	default:
		return fmt.Errorf("unknown type '%v' was specified", type_)
	}
	return nil
}

func (s *DatabaseToolsDatabaseToolsConnectionResourceCrud) populateTopLevelPolymorphicUpdateDatabaseToolsConnectionRequest(request *oci_database_tools.UpdateDatabaseToolsConnectionRequest) error {
	//discriminator
	typeRaw, ok := s.D.GetOkExists("type")
	var type_ string
	if ok {
		type_ = typeRaw.(string)
	} else {
		type_ = "" // default value
	}
	switch strings.ToLower(type_) {
	case strings.ToLower("ORACLE_DATABASE"):
		details := oci_database_tools.UpdateDatabaseToolsConnectionOracleDatabaseDetails{}
		if advancedProperties, ok := s.D.GetOkExists("advanced_properties"); ok {
			details.AdvancedProperties = utils.ObjectMapToStringMap(advancedProperties.(map[string]interface{}))
		}
		if connectionString, ok := s.D.GetOkExists("connection_string"); ok {
			tmp := connectionString.(string)
			details.ConnectionString = &tmp
		}
		if keyStores, ok := s.D.GetOkExists("key_stores"); ok {
			interfaces := keyStores.([]interface{})
			tmp := make([]oci_database_tools.DatabaseToolsKeyStoreDetails, len(interfaces))
			for i := range interfaces {
				stateDataIndex := i
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "key_stores", stateDataIndex)
				converted, err := s.mapToDatabaseToolsKeyStoreDetails(fieldKeyFormat)
				if err != nil {
					return err
				}
				tmp[i] = converted
			}
			if len(tmp) != 0 || s.D.HasChange("key_stores") {
				details.KeyStores = tmp
			}
		}
		if privateEndpointId, ok := s.D.GetOkExists("private_endpoint_id"); ok {
			tmp := privateEndpointId.(string)
			details.PrivateEndpointId = &tmp
		}
		if relatedResource, ok := s.D.GetOkExists("related_resource"); ok {
			if tmpList := relatedResource.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "related_resource", 0)
				tmp, err := s.mapToUpdateDatabaseToolsRelatedResourceDetails(fieldKeyFormat)
				if err != nil {
					return err
				}
				details.RelatedResource = &tmp
			}
		}
		if userName, ok := s.D.GetOkExists("user_name"); ok {
			tmp := userName.(string)
			details.UserName = &tmp
		}
		if userPassword, ok := s.D.GetOkExists("user_password"); ok {
			if tmpList := userPassword.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "user_password", 0)
				tmp, err := s.mapToDatabaseToolsUserPasswordDetails(fieldKeyFormat)
				if err != nil {
					return err
				}
				details.UserPassword = tmp
			}
		}
		tmp := s.D.Id()
		request.DatabaseToolsConnectionId = &tmp
		if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
			convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
			if err != nil {
				return err
			}
			details.DefinedTags = convertedDefinedTags
		}
		if displayName, ok := s.D.GetOkExists("display_name"); ok {
			tmp := displayName.(string)
			details.DisplayName = &tmp
		}
		if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
			details.FreeformTags = utils.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
		}
		request.UpdateDatabaseToolsConnectionDetails = details
	default:
		return fmt.Errorf("unknown type '%v' was specified", type_)
	}
	return nil
}

func (s *DatabaseToolsDatabaseToolsConnectionResourceCrud) updateCompartment(compartment interface{}) error {
	changeCompartmentRequest := oci_database_tools.ChangeDatabaseToolsConnectionCompartmentRequest{}

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	idTmp := s.D.Id()
	changeCompartmentRequest.DatabaseToolsConnectionId = &idTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database_tools")

	response, err := s.Client.ChangeDatabaseToolsConnectionCompartment(context.Background(), changeCompartmentRequest)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getDatabaseToolsConnectionFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database_tools"), oci_database_tools.ActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}
