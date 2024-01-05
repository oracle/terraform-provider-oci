// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package database_tools

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"

	oci_common "github.com/oracle/oci-go-sdk/v65/common"
	oci_database_tools "github.com/oracle/oci-go-sdk/v65/databasetools"
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
				ForceNew:         true,
				DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
				ValidateFunc: validation.StringInSlice([]string{
					"GENERIC_JDBC",
					"MYSQL",
					"ORACLE_DATABASE",
					"POSTGRESQL",
				}, true),
			},
			"user_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"user_password": {
				Type:     schema.TypeList,
				Required: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"secret_id": {
							Type:     schema.TypeString,
							Required: true,
						},
						"value_type": {
							Type:             schema.TypeString,
							Required:         true,
							DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
							ValidateFunc: validation.StringInSlice([]string{
								"SECRETID",
							}, true),
						},

						// Optional

						// Computed
					},
				},
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
			"private_endpoint_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"proxy_client": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"proxy_authentication_type": {
							Type:             schema.TypeString,
							Required:         true,
							DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
							ValidateFunc: validation.StringInSlice([]string{
								"NO_PROXY",
								"USER_NAME",
							}, true),
						},

						// Optional
						"roles": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
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
									"secret_id": {
										Type:     schema.TypeString,
										Required: true,
									},
									"value_type": {
										Type:             schema.TypeString,
										Required:         true,
										DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
										ValidateFunc: validation.StringInSlice([]string{
											"SECRETID",
										}, true),
									},

									// Optional

									// Computed
								},
							},
						},

						// Computed
					},
				},
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

						// Optional
						"entity_type": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"identifier": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},

						// Computed
					},
				},
			},
			"runtime_support": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"url": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
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
	var identifier *string
	identifier = response.GetId()
	if identifier != nil {
		s.D.SetId(*identifier)
	}
	return s.getDatabaseToolsConnectionFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database_tools"), oci_database_tools.ActionTypeCreated, s.D.Timeout(schema.TimeoutCreate))
}

func (s *DatabaseToolsDatabaseToolsConnectionResourceCrud) getDatabaseToolsConnectionFromWorkRequest(workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_database_tools.ActionTypeEnum, timeout time.Duration) error {

	// Wait until it finishes
	databaseToolsConnectionId, err := databaseToolsConnectionWaitForWorkRequest(workId, "databasetoolsconnection",
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

	if isLockOverride, ok := s.D.GetOkExists("is_lock_override"); ok {
		tmp := isLockOverride.(bool)
		request.IsLockOverride = &tmp
	}

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
	case oci_database_tools.DatabaseToolsConnectionGenericJdbc:
		s.D.Set("type", "GENERIC_JDBC")
		s.D.Set("advanced_properties", v.AdvancedProperties)

		keyStores := []interface{}{}
		for _, item := range v.KeyStores {
			keyStores = append(keyStores, DatabaseToolsKeyStoreGenericJdbcToMap(item))
		}
		s.D.Set("key_stores", keyStores)

		if v.Url != nil {
			s.D.Set("url", *v.Url)
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

		if v.LifecycleDetails != nil {
			s.D.Set("lifecycle_details", *v.LifecycleDetails)
		}

		locks := []interface{}{}
		for _, item := range v.Locks {
			locks = append(locks, ConnectionResourceLockToMap(item))
		}
		s.D.Set("locks", locks)

		s.D.Set("runtime_support", v.RuntimeSupport)

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
	case oci_database_tools.DatabaseToolsConnectionMySql:
		s.D.Set("type", "MYSQL")

		s.D.Set("advanced_properties", v.AdvancedProperties)

		if v.ConnectionString != nil {
			s.D.Set("connection_string", *v.ConnectionString)
		}

		keyStores := []interface{}{}
		for _, item := range v.KeyStores {
			keyStores = append(keyStores, DatabaseToolsKeyStoreMySqlToMap(item))
		}
		s.D.Set("key_stores", keyStores)

		if v.PrivateEndpointId != nil {
			s.D.Set("private_endpoint_id", *v.PrivateEndpointId)
		}

		if v.RelatedResource != nil {
			s.D.Set("related_resource", []interface{}{DatabaseToolsRelatedResourceMySqlToMap(v.RelatedResource)})
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

		if v.LifecycleDetails != nil {
			s.D.Set("lifecycle_details", *v.LifecycleDetails)
		}

		locks := []interface{}{}
		for _, item := range v.Locks {
			locks = append(locks, ConnectionResourceLockToMap(item))
		}
		s.D.Set("locks", locks)

		s.D.Set("runtime_support", v.RuntimeSupport)

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

		if v.ProxyClient != nil {
			proxyClientArray := []interface{}{}
			if proxyClientMap := DatabaseToolsConnectionOracleDatabaseProxyClientToMap(&v.ProxyClient); proxyClientMap != nil {
				proxyClientArray = append(proxyClientArray, proxyClientMap)
			}
			s.D.Set("proxy_client", proxyClientArray)
		} else {
			s.D.Set("proxy_client", nil)
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

		if v.LifecycleDetails != nil {
			s.D.Set("lifecycle_details", *v.LifecycleDetails)
		}

		locks := []interface{}{}
		for _, item := range v.Locks {
			locks = append(locks, ConnectionResourceLockToMap(item))
		}
		s.D.Set("locks", locks)

		s.D.Set("runtime_support", v.RuntimeSupport)

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
	case oci_database_tools.DatabaseToolsConnectionPostgresql:
		s.D.Set("type", "POSTGRESQL")
		s.D.Set("advanced_properties", v.AdvancedProperties)

		if v.ConnectionString != nil {
			s.D.Set("connection_string", *v.ConnectionString)
		}

		keyStores := []interface{}{}
		for _, item := range v.KeyStores {
			keyStores = append(keyStores, DatabaseToolsKeyStorePostgresqlToMap(item))
		}
		s.D.Set("key_stores", keyStores)

		if v.PrivateEndpointId != nil {
			s.D.Set("private_endpoint_id", *v.PrivateEndpointId)
		}

		if v.RelatedResource != nil {
			s.D.Set("related_resource", []interface{}{DatabaseToolsRelatedResourcePostgresqlToMap(v.RelatedResource)})
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

		if v.LifecycleDetails != nil {
			s.D.Set("lifecycle_details", *v.LifecycleDetails)
		}

		locks := []interface{}{}
		for _, item := range v.Locks {
			locks = append(locks, ConnectionResourceLockToMap(item))
		}
		s.D.Set("locks", locks)

		s.D.Set("runtime_support", v.RuntimeSupport)

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

// For type: MYSQL
func (s *DatabaseToolsDatabaseToolsConnectionResourceCrud) mapToCreateDatabaseToolsRelatedResourceMySqlDetails(fieldKeyFormat string) (oci_database_tools.CreateDatabaseToolsRelatedResourceMySqlDetails, error) {
	result := oci_database_tools.CreateDatabaseToolsRelatedResourceMySqlDetails{}

	if entityType, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "entity_type")); ok {
		result.EntityType = oci_database_tools.RelatedResourceEntityTypeMySqlEnum(entityType.(string))
	}

	if identifier, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "identifier")); ok {
		tmp := identifier.(string)
		result.Identifier = &tmp
	}

	return result, nil
}

// For type: MYSQL
func CreateDatabaseToolsRelatedResourceMySqlDetailsToMap(obj *oci_database_tools.CreateDatabaseToolsRelatedResourceMySqlDetails) map[string]interface{} {
	result := map[string]interface{}{}

	result["entity_type"] = string(obj.EntityType)

	if obj.Identifier != nil {
		result["identifier"] = string(*obj.Identifier)
	}

	return result
}

// For type: POSTGRESQL
func (s *DatabaseToolsDatabaseToolsConnectionResourceCrud) mapToCreateDatabaseToolsRelatedResourcePostgresqlDetails(fieldKeyFormat string) (oci_database_tools.CreateDatabaseToolsRelatedResourcePostgresqlDetails, error) {
	result := oci_database_tools.CreateDatabaseToolsRelatedResourcePostgresqlDetails{}

	if entityType, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "entity_type")); ok {
		result.EntityType = oci_database_tools.RelatedResourceEntityTypePostgresqlEnum(entityType.(string))
	}

	if identifier, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "identifier")); ok {
		tmp := identifier.(string)
		result.Identifier = &tmp
	}

	return result, nil
}

// For type: POSTGRESQL
func CreateDatabaseToolsRelatedResourcePostgresqlDetailsToMap(obj *oci_database_tools.CreateDatabaseToolsRelatedResourcePostgresqlDetails) map[string]interface{} {
	result := map[string]interface{}{}

	result["entity_type"] = string(obj.EntityType)

	if obj.Identifier != nil {
		result["identifier"] = string(*obj.Identifier)
	}

	return result
}

func (s *DatabaseToolsDatabaseToolsConnectionResourceCrud) mapToDatabaseToolsConnectionOracleDatabaseProxyClient(fieldKeyFormat string) (oci_database_tools.DatabaseToolsConnectionOracleDatabaseProxyClient, error) {
	var baseObject oci_database_tools.DatabaseToolsConnectionOracleDatabaseProxyClient
	//discriminator
	proxyAuthenticationTypeRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "proxy_authentication_type"))
	var proxyAuthenticationType string
	if ok {
		proxyAuthenticationType = proxyAuthenticationTypeRaw.(string)
	} else {
		proxyAuthenticationType = "NO_PROXY" // default value
	}
	switch strings.ToLower(proxyAuthenticationType) {
	case strings.ToLower("NO_PROXY"):
		details := oci_database_tools.DatabaseToolsConnectionOracleDatabaseProxyClientNoProxy{}
		baseObject = details
	case strings.ToLower("USER_NAME"):
		details := oci_database_tools.DatabaseToolsConnectionOracleDatabaseProxyClientUserName{}
		if roles, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "roles")); ok {
			interfaces := roles.([]interface{})
			tmp := make([]string, len(interfaces))
			for i := range interfaces {
				if interfaces[i] != nil {
					tmp[i] = interfaces[i].(string)
				}
			}
			if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "roles")) {
				details.Roles = tmp
			}
		}
		if userName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "user_name")); ok {
			tmp := userName.(string)
			details.UserName = &tmp
		}
		if userPassword, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "user_password")); ok {
			if tmpList := userPassword.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "user_password"), 0)
				tmp, err := s.mapToDatabaseToolsUserPassword(fieldKeyFormatNextLevel)
				if err != nil {
					return details, fmt.Errorf("unable to convert user_password, encountered error: %v", err)
				}
				details.UserPassword = tmp
			}
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown proxy_authentication_type '%v' was specified", proxyAuthenticationType)
	}
	return baseObject, nil
}

func DatabaseToolsConnectionOracleDatabaseProxyClientToMap(obj *oci_database_tools.DatabaseToolsConnectionOracleDatabaseProxyClient) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (*obj).(type) {
	case oci_database_tools.DatabaseToolsConnectionOracleDatabaseProxyClientNoProxy:
		result["proxy_authentication_type"] = "NO_PROXY"
	case oci_database_tools.DatabaseToolsConnectionOracleDatabaseProxyClientUserName:
		result["proxy_authentication_type"] = "USER_NAME"

		result["roles"] = v.Roles

		if v.UserName != nil {
			result["user_name"] = string(*v.UserName)
		}

		if v.UserPassword != nil {
			userPasswordArray := []interface{}{}
			if userPasswordMap := DatabaseToolsUserPasswordToMap(&v.UserPassword); userPasswordMap != nil {
				userPasswordArray = append(userPasswordArray, userPasswordMap)
			}
			result["user_password"] = userPasswordArray
		}
	default:
		log.Printf("[WARN] Received 'proxy_authentication_type' of unknown type %v", *obj)
		return nil
	}

	return result
}

func (s *DatabaseToolsDatabaseToolsConnectionResourceCrud) mapToDatabaseToolsConnectionOracleDatabaseProxyClientDetails(fieldKeyFormat string) (oci_database_tools.DatabaseToolsConnectionOracleDatabaseProxyClientDetails, error) {
	var baseObject oci_database_tools.DatabaseToolsConnectionOracleDatabaseProxyClientDetails
	//discriminator
	proxyAuthenticationTypeRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "proxy_authentication_type"))
	var proxyAuthenticationType string
	if ok {
		proxyAuthenticationType = proxyAuthenticationTypeRaw.(string)
	} else {
		proxyAuthenticationType = "NO_PROXY" // default value
	}
	switch strings.ToLower(proxyAuthenticationType) {
	case strings.ToLower("NO_PROXY"):
		details := oci_database_tools.DatabaseToolsConnectionOracleDatabaseProxyClientNoProxyDetails{}
		baseObject = details
	case strings.ToLower("USER_NAME"):
		details := oci_database_tools.DatabaseToolsConnectionOracleDatabaseProxyClientUserNameDetails{}
		if roles, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "roles")); ok {
			interfaces := roles.([]interface{})
			tmp := make([]string, len(interfaces))
			for i := range interfaces {
				if interfaces[i] != nil {
					tmp[i] = interfaces[i].(string)
				}
			}
			if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "roles")) {
				details.Roles = tmp
			}
		}
		if userName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "user_name")); ok {
			tmp := userName.(string)
			details.UserName = &tmp
		}
		if userPassword, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "user_password")); ok {
			if tmpList := userPassword.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "user_password"), 0)
				tmp, err := s.mapToDatabaseToolsUserPasswordDetails(fieldKeyFormatNextLevel)
				if err != nil {
					return details, fmt.Errorf("unable to convert user_password, encountered error: %v", err)
				}
				details.UserPassword = tmp
			}
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown proxy_authentication_type '%v' was specified", proxyAuthenticationType)
	}
	return baseObject, nil
}

func DatabaseToolsConnectionOracleDatabaseProxyClientDetailsToMap(obj *oci_database_tools.DatabaseToolsConnectionOracleDatabaseProxyClientDetails) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (*obj).(type) {
	case oci_database_tools.DatabaseToolsConnectionOracleDatabaseProxyClientNoProxyDetails:
		result["proxy_authentication_type"] = "NO_PROXY"
	case oci_database_tools.DatabaseToolsConnectionOracleDatabaseProxyClientUserNameDetails:
		result["proxy_authentication_type"] = "USER_NAME"

		result["roles"] = v.Roles
		result["roles"] = v.Roles

		if v.UserName != nil {
			result["user_name"] = string(*v.UserName)
		}

		if v.UserPassword != nil {
			userPasswordArray := []interface{}{}
			if userPasswordMap := DatabaseToolsUserPasswordDetailsToMap(&v.UserPassword); userPasswordMap != nil {
				userPasswordArray = append(userPasswordArray, userPasswordMap)
			}
			result["user_password"] = userPasswordArray
		}
	default:
		log.Printf("[WARN] Received 'proxy_authentication_type' of unknown type %v", *obj)
		return nil
	}

	return result
}

func (s *DatabaseToolsDatabaseToolsConnectionResourceCrud) mapToDatabaseToolsConnectionOracleDatabaseProxyClientSummary(fieldKeyFormat string) (oci_database_tools.DatabaseToolsConnectionOracleDatabaseProxyClientSummary, error) {
	var baseObject oci_database_tools.DatabaseToolsConnectionOracleDatabaseProxyClientSummary
	//discriminator
	proxyAuthenticationTypeRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "proxy_authentication_type"))
	var proxyAuthenticationType string
	if ok {
		proxyAuthenticationType = proxyAuthenticationTypeRaw.(string)
	} else {
		proxyAuthenticationType = "NO_PROXY" // default value
	}
	switch strings.ToLower(proxyAuthenticationType) {
	case strings.ToLower("NO_PROXY"):
		details := oci_database_tools.DatabaseToolsConnectionOracleDatabaseProxyClientNoProxySummary{}
		baseObject = details
	case strings.ToLower("USER_NAME"):
		details := oci_database_tools.DatabaseToolsConnectionOracleDatabaseProxyClientUserNameSummary{}
		if roles, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "roles")); ok {
			interfaces := roles.([]interface{})
			tmp := make([]string, len(interfaces))
			for i := range interfaces {
				if interfaces[i] != nil {
					tmp[i] = interfaces[i].(string)
				}
			}
			if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "roles")) {
				details.Roles = tmp
			}
		}
		if userName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "user_name")); ok {
			tmp := userName.(string)
			details.UserName = &tmp
		}
		if userPassword, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "user_password")); ok {
			if tmpList := userPassword.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "user_password"), 0)
				tmp, err := s.mapToDatabaseToolsUserPasswordSummary(fieldKeyFormatNextLevel)
				if err != nil {
					return details, fmt.Errorf("unable to convert user_password, encountered error: %v", err)
				}
				details.UserPassword = tmp
			}
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown proxy_authentication_type '%v' was specified", proxyAuthenticationType)
	}
	return baseObject, nil
}

func DatabaseToolsConnectionOracleDatabaseProxyClientSummaryToMap(obj *oci_database_tools.DatabaseToolsConnectionOracleDatabaseProxyClientSummary) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (*obj).(type) {
	case oci_database_tools.DatabaseToolsConnectionOracleDatabaseProxyClientNoProxySummary:
		result["proxy_authentication_type"] = "NO_PROXY"
	case oci_database_tools.DatabaseToolsConnectionOracleDatabaseProxyClientUserNameSummary:
		result["proxy_authentication_type"] = "USER_NAME"
		result["roles"] = v.Roles
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
	default:
		log.Printf("[WARN] Received 'proxy_authentication_type' of unknown type %v", *obj)
		return nil
	}

	return result
}

func DatabaseToolsConnectionSummaryToMap(obj oci_database_tools.DatabaseToolsConnectionSummary) map[string]interface{} {
	result := map[string]interface{}{}

	switch v := (obj).(type) {
	case oci_database_tools.DatabaseToolsConnectionGenericJdbcSummary:
		result["type"] = "GENERIC_JDBC"
		result["advanced_properties"] = v.AdvancedProperties

		keyStores := []interface{}{}
		for _, item := range v.KeyStores {
			keyStores = append(keyStores, DatabaseToolsKeyStoreGenericJdbcSummaryToMap(item))
		}
		result["key_stores"] = keyStores

		if v.Url != nil {
			result["url"] = string(*v.Url)
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

		// These missing fields correspond to the fields defined in parent DatabaseToolsConnectionSummary
		result["runtime_support"] = v.RuntimeSupport

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

		result["state"] = string(v.LifecycleState)

		if v.LifecycleDetails != nil {
			result["lifecycle_details"] = string(*v.LifecycleDetails)
		}
	case oci_database_tools.DatabaseToolsConnectionMySqlSummary:
		result["type"] = "MYSQL"
		result["advanced_properties"] = v.AdvancedProperties

		if v.ConnectionString != nil {
			result["connection_string"] = string(*v.ConnectionString)
		}

		keyStores := []interface{}{}
		for _, item := range v.KeyStores {
			keyStores = append(keyStores, DatabaseToolsKeyStoreMySqlSummaryToMap(item))
		}
		result["key_stores"] = keyStores

		if v.PrivateEndpointId != nil {
			result["private_endpoint_id"] = string(*v.PrivateEndpointId)
		}

		if v.RelatedResource != nil {
			result["related_resource"] = []interface{}{DatabaseToolsRelatedResourceMySqlToMap(v.RelatedResource)}
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

		// These missing fields correspond to the fields defined in parent DatabaseToolsConnectionSummary
		result["runtime_support"] = v.RuntimeSupport

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

		if v.ProxyClient != nil {
			proxyClientArray := []interface{}{}
			if proxyClientMap := DatabaseToolsConnectionOracleDatabaseProxyClientSummaryToMap(&v.ProxyClient); proxyClientMap != nil {
				proxyClientArray = append(proxyClientArray, proxyClientMap)
			}
			result["proxy_client"] = proxyClientArray
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

		// These missing fields correspond to the fields defined in parent DatabaseToolsConnectionSummary
		result["runtime_support"] = v.RuntimeSupport

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
	case oci_database_tools.DatabaseToolsConnectionPostgresqlSummary:
		result["type"] = "POSTGRESQL"

		result["advanced_properties"] = v.AdvancedProperties

		if v.ConnectionString != nil {
			result["connection_string"] = string(*v.ConnectionString)
		}

		keyStores := []interface{}{}
		for _, item := range v.KeyStores {
			keyStores = append(keyStores, DatabaseToolsKeyStorePostgresqlSummaryToMap(item))
		}
		result["key_stores"] = keyStores

		if v.PrivateEndpointId != nil {
			result["private_endpoint_id"] = string(*v.PrivateEndpointId)
		}

		if v.RelatedResource != nil {
			result["related_resource"] = []interface{}{DatabaseToolsRelatedResourcePostgresqlToMap(v.RelatedResource)}
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

		// These missing fields correspond to the fields defined in parent DatabaseToolsConnectionSummary
		result["runtime_support"] = v.RuntimeSupport

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

func (s *DatabaseToolsDatabaseToolsConnectionResourceCrud) mapToDatabaseToolsKeyStoreContentGenericJdbc(fieldKeyFormat string) (oci_database_tools.DatabaseToolsKeyStoreContentGenericJdbc, error) {
	var baseObject oci_database_tools.DatabaseToolsKeyStoreContentGenericJdbc
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
		details := oci_database_tools.DatabaseToolsKeyStoreContentSecretIdGenericJdbc{}
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

func DatabaseToolsKeyStoreContentGenericJdbcToMap(obj *oci_database_tools.DatabaseToolsKeyStoreContentGenericJdbc) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (*obj).(type) {
	case oci_database_tools.DatabaseToolsKeyStoreContentSecretIdGenericJdbc:
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

func (s *DatabaseToolsDatabaseToolsConnectionResourceCrud) mapToDatabaseToolsKeyStoreContentGenericJdbcDetails(fieldKeyFormat string) (oci_database_tools.DatabaseToolsKeyStoreContentGenericJdbcDetails, error) {
	var baseObject oci_database_tools.DatabaseToolsKeyStoreContentGenericJdbcDetails
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
		details := oci_database_tools.DatabaseToolsKeyStoreContentSecretIdGenericJdbcDetails{}
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

func DatabaseToolsKeyStoreContentGenericJdbcDetailsToMap(obj *oci_database_tools.DatabaseToolsKeyStoreContentGenericJdbcDetails) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (*obj).(type) {
	case oci_database_tools.DatabaseToolsKeyStoreContentSecretIdGenericJdbcDetails:
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

func (s *DatabaseToolsDatabaseToolsConnectionResourceCrud) mapToDatabaseToolsKeyStoreContentGenericJdbcSummary(fieldKeyFormat string) (oci_database_tools.DatabaseToolsKeyStoreContentGenericJdbcSummary, error) {
	var baseObject oci_database_tools.DatabaseToolsKeyStoreContentGenericJdbcSummary
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
		details := oci_database_tools.DatabaseToolsKeyStoreContentSecretIdGenericJdbcSummary{}
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

func DatabaseToolsKeyStoreContentGenericJdbcSummaryToMap(obj *oci_database_tools.DatabaseToolsKeyStoreContentGenericJdbcSummary) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (*obj).(type) {
	case oci_database_tools.DatabaseToolsKeyStoreContentSecretIdGenericJdbcSummary:
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

func (s *DatabaseToolsDatabaseToolsConnectionResourceCrud) mapToDatabaseToolsKeyStoreContentMySql(fieldKeyFormat string) (oci_database_tools.DatabaseToolsKeyStoreContentMySql, error) {
	var baseObject oci_database_tools.DatabaseToolsKeyStoreContentMySql
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
		details := oci_database_tools.DatabaseToolsKeyStoreContentSecretIdMySql{}
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

// For type: ORACLE_DATABASE
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

func DatabaseToolsKeyStoreContentMySqlToMap(obj *oci_database_tools.DatabaseToolsKeyStoreContentMySql) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (*obj).(type) {
	case oci_database_tools.DatabaseToolsKeyStoreContentSecretIdMySql:
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

func (s *DatabaseToolsDatabaseToolsConnectionResourceCrud) mapToDatabaseToolsKeyStoreContentMySqlDetails(fieldKeyFormat string) (oci_database_tools.DatabaseToolsKeyStoreContentMySqlDetails, error) {
	var baseObject oci_database_tools.DatabaseToolsKeyStoreContentMySqlDetails
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
		details := oci_database_tools.DatabaseToolsKeyStoreContentSecretIdMySqlDetails{}
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

func DatabaseToolsKeyStoreContentMySqlDetailsToMap(obj *oci_database_tools.DatabaseToolsKeyStoreContentMySqlDetails) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (*obj).(type) {
	case oci_database_tools.DatabaseToolsKeyStoreContentSecretIdMySqlDetails:
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

func (s *DatabaseToolsDatabaseToolsConnectionResourceCrud) mapToDatabaseToolsKeyStoreContentMySqlSummary(fieldKeyFormat string) (oci_database_tools.DatabaseToolsKeyStoreContentMySqlSummary, error) {
	var baseObject oci_database_tools.DatabaseToolsKeyStoreContentMySqlSummary
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
		details := oci_database_tools.DatabaseToolsKeyStoreContentSecretIdMySqlSummary{}
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

func DatabaseToolsKeyStoreContentMySqlSummaryToMap(obj *oci_database_tools.DatabaseToolsKeyStoreContentMySqlSummary) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (*obj).(type) {
	case oci_database_tools.DatabaseToolsKeyStoreContentSecretIdMySqlSummary:
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

func (s *DatabaseToolsDatabaseToolsConnectionResourceCrud) mapToDatabaseToolsKeyStoreContentPostgresql(fieldKeyFormat string) (oci_database_tools.DatabaseToolsKeyStoreContentPostgresql, error) {
	var baseObject oci_database_tools.DatabaseToolsKeyStoreContentPostgresql
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
		details := oci_database_tools.DatabaseToolsKeyStoreContentSecretIdPostgresql{}
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

func DatabaseToolsKeyStoreContentPostgresqlToMap(obj *oci_database_tools.DatabaseToolsKeyStoreContentPostgresql) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (*obj).(type) {
	case oci_database_tools.DatabaseToolsKeyStoreContentSecretIdPostgresql:
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

func (s *DatabaseToolsDatabaseToolsConnectionResourceCrud) mapToDatabaseToolsKeyStoreContentPostgresqlDetails(fieldKeyFormat string) (oci_database_tools.DatabaseToolsKeyStoreContentPostgresqlDetails, error) {
	var baseObject oci_database_tools.DatabaseToolsKeyStoreContentPostgresqlDetails
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
		details := oci_database_tools.DatabaseToolsKeyStoreContentSecretIdPostgresqlDetails{}
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

func DatabaseToolsKeyStoreContentPostgresqlDetailsToMap(obj *oci_database_tools.DatabaseToolsKeyStoreContentPostgresqlDetails) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (*obj).(type) {
	case oci_database_tools.DatabaseToolsKeyStoreContentSecretIdPostgresqlDetails:
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

func (s *DatabaseToolsDatabaseToolsConnectionResourceCrud) mapToDatabaseToolsKeyStoreContentPostgresqlSummary(fieldKeyFormat string) (oci_database_tools.DatabaseToolsKeyStoreContentPostgresqlSummary, error) {
	var baseObject oci_database_tools.DatabaseToolsKeyStoreContentPostgresqlSummary
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
		details := oci_database_tools.DatabaseToolsKeyStoreContentSecretIdPostgresqlSummary{}
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

func DatabaseToolsKeyStoreContentPostgresqlSummaryToMap(obj *oci_database_tools.DatabaseToolsKeyStoreContentPostgresqlSummary) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (*obj).(type) {
	case oci_database_tools.DatabaseToolsKeyStoreContentSecretIdPostgresqlSummary:
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

// For type: ORACLE_DATABASE
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

func (s *DatabaseToolsDatabaseToolsConnectionResourceCrud) mapToDatabaseToolsKeyStoreGenericJdbc(fieldKeyFormat string) (oci_database_tools.DatabaseToolsKeyStoreGenericJdbc, error) {
	result := oci_database_tools.DatabaseToolsKeyStoreGenericJdbc{}

	if keyStoreContent, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "key_store_content")); ok {
		if tmpList := keyStoreContent.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "key_store_content"), 0)
			tmp, err := s.mapToDatabaseToolsKeyStoreContentGenericJdbc(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert key_store_content, encountered error: %v", err)
			}
			result.KeyStoreContent = tmp
		}
	}

	if keyStorePassword, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "key_store_password")); ok {
		if tmpList := keyStorePassword.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "key_store_password"), 0)
			tmp, err := s.mapToDatabaseToolsKeyStorePasswordGenericJdbc(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert key_store_password, encountered error: %v", err)
			}
			result.KeyStorePassword = tmp
		}
	}

	if keyStoreType, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "key_store_type")); ok {
		result.KeyStoreType = oci_database_tools.KeyStoreTypeGenericJdbcEnum(keyStoreType.(string))
	}

	return result, nil
}

func DatabaseToolsKeyStoreGenericJdbcToMap(obj oci_database_tools.DatabaseToolsKeyStoreGenericJdbc) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.KeyStoreContent != nil {
		keyStoreContentArray := []interface{}{}
		if keyStoreContentMap := DatabaseToolsKeyStoreContentGenericJdbcToMap(&obj.KeyStoreContent); keyStoreContentMap != nil {
			keyStoreContentArray = append(keyStoreContentArray, keyStoreContentMap)
		}
		result["key_store_content"] = keyStoreContentArray
	}

	if obj.KeyStorePassword != nil {
		keyStorePasswordArray := []interface{}{}
		if keyStorePasswordMap := DatabaseToolsKeyStorePasswordGenericJdbcToMap(&obj.KeyStorePassword); keyStorePasswordMap != nil {
			keyStorePasswordArray = append(keyStorePasswordArray, keyStorePasswordMap)
		}
		result["key_store_password"] = keyStorePasswordArray
	}

	result["key_store_type"] = string(obj.KeyStoreType)

	return result
}

func (s *DatabaseToolsDatabaseToolsConnectionResourceCrud) mapToDatabaseToolsKeyStoreGenericJdbcDetails(fieldKeyFormat string) (oci_database_tools.DatabaseToolsKeyStoreGenericJdbcDetails, error) {
	result := oci_database_tools.DatabaseToolsKeyStoreGenericJdbcDetails{}

	if keyStoreContent, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "key_store_content")); ok {
		if tmpList := keyStoreContent.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "key_store_content"), 0)
			tmp, err := s.mapToDatabaseToolsKeyStoreContentGenericJdbcDetails(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert key_store_content, encountered error: %v", err)
			}
			result.KeyStoreContent = tmp
		}
	}

	if keyStorePassword, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "key_store_password")); ok {
		if tmpList := keyStorePassword.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "key_store_password"), 0)
			tmp, err := s.mapToDatabaseToolsKeyStorePasswordGenericJdbcDetails(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert key_store_password, encountered error: %v", err)
			}
			result.KeyStorePassword = tmp
		}
	}

	if keyStoreType, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "key_store_type")); ok {
		result.KeyStoreType = oci_database_tools.KeyStoreTypeGenericJdbcEnum(keyStoreType.(string))
	}

	return result, nil
}

func DatabaseToolsKeyStoreGenericJdbcDetailsToMap(obj oci_database_tools.DatabaseToolsKeyStoreGenericJdbcDetails) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.KeyStoreContent != nil {
		keyStoreContentArray := []interface{}{}
		if keyStoreContentMap := DatabaseToolsKeyStoreContentGenericJdbcDetailsToMap(&obj.KeyStoreContent); keyStoreContentMap != nil {
			keyStoreContentArray = append(keyStoreContentArray, keyStoreContentMap)
		}
		result["key_store_content"] = keyStoreContentArray
	}

	if obj.KeyStorePassword != nil {
		keyStorePasswordArray := []interface{}{}
		if keyStorePasswordMap := DatabaseToolsKeyStorePasswordGenericJdbcDetailsToMap(&obj.KeyStorePassword); keyStorePasswordMap != nil {
			keyStorePasswordArray = append(keyStorePasswordArray, keyStorePasswordMap)
		}
		result["key_store_password"] = keyStorePasswordArray
	}

	result["key_store_type"] = string(obj.KeyStoreType)

	return result
}

func (s *DatabaseToolsDatabaseToolsConnectionResourceCrud) mapToDatabaseToolsKeyStoreGenericJdbcSummary(fieldKeyFormat string) (oci_database_tools.DatabaseToolsKeyStoreGenericJdbcSummary, error) {
	result := oci_database_tools.DatabaseToolsKeyStoreGenericJdbcSummary{}

	if keyStoreContent, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "key_store_content")); ok {
		if tmpList := keyStoreContent.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "key_store_content"), 0)
			tmp, err := s.mapToDatabaseToolsKeyStoreContentGenericJdbcSummary(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert key_store_content, encountered error: %v", err)
			}
			result.KeyStoreContent = tmp
		}
	}

	if keyStorePassword, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "key_store_password")); ok {
		if tmpList := keyStorePassword.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "key_store_password"), 0)
			tmp, err := s.mapToDatabaseToolsKeyStorePasswordGenericJdbcSummary(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert key_store_password, encountered error: %v", err)
			}
			result.KeyStorePassword = tmp
		}
	}

	if keyStoreType, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "key_store_type")); ok {
		result.KeyStoreType = oci_database_tools.KeyStoreTypeGenericJdbcEnum(keyStoreType.(string))
	}

	return result, nil
}

func DatabaseToolsKeyStoreGenericJdbcSummaryToMap(obj oci_database_tools.DatabaseToolsKeyStoreGenericJdbcSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.KeyStoreContent != nil {
		keyStoreContentArray := []interface{}{}
		if keyStoreContentMap := DatabaseToolsKeyStoreContentGenericJdbcSummaryToMap(&obj.KeyStoreContent); keyStoreContentMap != nil {
			keyStoreContentArray = append(keyStoreContentArray, keyStoreContentMap)
		}
		result["key_store_content"] = keyStoreContentArray
	}

	if obj.KeyStorePassword != nil {
		keyStorePasswordArray := []interface{}{}
		if keyStorePasswordMap := DatabaseToolsKeyStorePasswordGenericJdbcSummaryToMap(&obj.KeyStorePassword); keyStorePasswordMap != nil {
			keyStorePasswordArray = append(keyStorePasswordArray, keyStorePasswordMap)
		}
		result["key_store_password"] = keyStorePasswordArray
	}

	result["key_store_type"] = string(obj.KeyStoreType)

	return result
}

// For type: MYSQL
func (s *DatabaseToolsDatabaseToolsConnectionResourceCrud) mapToDatabaseToolsKeyStoreMySql(fieldKeyFormat string) (oci_database_tools.DatabaseToolsKeyStoreMySql, error) {
	result := oci_database_tools.DatabaseToolsKeyStoreMySql{}

	if keyStoreContent, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "key_store_content")); ok {
		if tmpList := keyStoreContent.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "key_store_content"), 0)
			tmp, err := s.mapToDatabaseToolsKeyStoreContentMySql(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert key_store_content, encountered error: %v", err)
			}
			result.KeyStoreContent = tmp
		}
	}

	if keyStorePassword, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "key_store_password")); ok {
		if tmpList := keyStorePassword.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "key_store_password"), 0)
			tmp, err := s.mapToDatabaseToolsKeyStorePasswordMySql(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert key_store_password, encountered error: %v", err)
			}
			result.KeyStorePassword = tmp
		}
	}

	if keyStoreType, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "key_store_type")); ok {
		result.KeyStoreType = oci_database_tools.KeyStoreTypeMySqlEnum(keyStoreType.(string))
	}

	return result, nil
}

// For type: MYSQL
func DatabaseToolsKeyStoreMySqlToMap(obj oci_database_tools.DatabaseToolsKeyStoreMySql) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.KeyStoreContent != nil {
		keyStoreContentArray := []interface{}{}
		if keyStoreContentMap := DatabaseToolsKeyStoreContentMySqlToMap(&obj.KeyStoreContent); keyStoreContentMap != nil {
			keyStoreContentArray = append(keyStoreContentArray, keyStoreContentMap)
		}
		result["key_store_content"] = keyStoreContentArray
	}

	if obj.KeyStorePassword != nil {
		keyStorePasswordArray := []interface{}{}
		if keyStorePasswordMap := DatabaseToolsKeyStorePasswordMySqlToMap(&obj.KeyStorePassword); keyStorePasswordMap != nil {
			keyStorePasswordArray = append(keyStorePasswordArray, keyStorePasswordMap)
		}
		result["key_store_password"] = keyStorePasswordArray
	}

	result["key_store_type"] = string(obj.KeyStoreType)

	return result
}

// For type: MYSQL
func (s *DatabaseToolsDatabaseToolsConnectionResourceCrud) mapToDatabaseToolsKeyStoreMySqlDetails(fieldKeyFormat string) (oci_database_tools.DatabaseToolsKeyStoreMySqlDetails, error) {
	result := oci_database_tools.DatabaseToolsKeyStoreMySqlDetails{}

	if keyStoreContent, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "key_store_content")); ok {
		if tmpList := keyStoreContent.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "key_store_content"), 0)
			tmp, err := s.mapToDatabaseToolsKeyStoreContentMySqlDetails(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert key_store_content, encountered error: %v", err)
			}
			result.KeyStoreContent = tmp
		}
	}

	if keyStorePassword, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "key_store_password")); ok {
		if tmpList := keyStorePassword.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "key_store_password"), 0)
			tmp, err := s.mapToDatabaseToolsKeyStorePasswordMySqlDetails(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert key_store_password, encountered error: %v", err)
			}
			result.KeyStorePassword = tmp
		}
	}

	if keyStoreType, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "key_store_type")); ok {
		result.KeyStoreType = oci_database_tools.KeyStoreTypeMySqlEnum(keyStoreType.(string))
	}

	return result, nil
}

// For type: MYSQL
func DatabaseToolsKeyStoreMySqlDetailsToMap(obj oci_database_tools.DatabaseToolsKeyStoreMySqlDetails) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.KeyStoreContent != nil {
		keyStoreContentArray := []interface{}{}
		if keyStoreContentMap := DatabaseToolsKeyStoreContentMySqlDetailsToMap(&obj.KeyStoreContent); keyStoreContentMap != nil {
			keyStoreContentArray = append(keyStoreContentArray, keyStoreContentMap)
		}
		result["key_store_content"] = keyStoreContentArray
	}

	if obj.KeyStorePassword != nil {
		keyStorePasswordArray := []interface{}{}
		if keyStorePasswordMap := DatabaseToolsKeyStorePasswordMySqlDetailsToMap(&obj.KeyStorePassword); keyStorePasswordMap != nil {
			keyStorePasswordArray = append(keyStorePasswordArray, keyStorePasswordMap)
		}
		result["key_store_password"] = keyStorePasswordArray
	}

	result["key_store_type"] = string(obj.KeyStoreType)

	return result
}

// For type: MYSQL
func (s *DatabaseToolsDatabaseToolsConnectionResourceCrud) mapToDatabaseToolsKeyStoreMySqlSummary(fieldKeyFormat string) (oci_database_tools.DatabaseToolsKeyStoreMySqlSummary, error) {
	result := oci_database_tools.DatabaseToolsKeyStoreMySqlSummary{}

	if keyStoreContent, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "key_store_content")); ok {
		if tmpList := keyStoreContent.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "key_store_content"), 0)
			tmp, err := s.mapToDatabaseToolsKeyStoreContentMySqlSummary(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert key_store_content, encountered error: %v", err)
			}
			result.KeyStoreContent = tmp
		}
	}

	if keyStorePassword, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "key_store_password")); ok {
		if tmpList := keyStorePassword.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "key_store_password"), 0)
			tmp, err := s.mapToDatabaseToolsKeyStorePasswordMySqlSummary(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert key_store_password, encountered error: %v", err)
			}
			result.KeyStorePassword = tmp
		}
	}

	if keyStoreType, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "key_store_type")); ok {
		result.KeyStoreType = oci_database_tools.KeyStoreTypeMySqlEnum(keyStoreType.(string))
	}

	return result, nil
}

// For type: MYSQL
func DatabaseToolsKeyStoreMySqlSummaryToMap(obj oci_database_tools.DatabaseToolsKeyStoreMySqlSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.KeyStoreContent != nil {
		keyStoreContentArray := []interface{}{}
		if keyStoreContentMap := DatabaseToolsKeyStoreContentMySqlSummaryToMap(&obj.KeyStoreContent); keyStoreContentMap != nil {
			keyStoreContentArray = append(keyStoreContentArray, keyStoreContentMap)
		}
		result["key_store_content"] = keyStoreContentArray
	}

	if obj.KeyStorePassword != nil {
		keyStorePasswordArray := []interface{}{}
		if keyStorePasswordMap := DatabaseToolsKeyStorePasswordMySqlSummaryToMap(&obj.KeyStorePassword); keyStorePasswordMap != nil {
			keyStorePasswordArray = append(keyStorePasswordArray, keyStorePasswordMap)
		}
		result["key_store_password"] = keyStorePasswordArray
	}

	result["key_store_type"] = string(obj.KeyStoreType)

	return result
}

// For type: ORACLE_DATABASE
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

// For type: ORACLE_DATABASE
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

// For type: ORACLE_DATABASE
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

// For type: ORACLE_DATABASE
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

func (s *DatabaseToolsDatabaseToolsConnectionResourceCrud) mapToDatabaseToolsKeyStorePasswordGenericJdbc(fieldKeyFormat string) (oci_database_tools.DatabaseToolsKeyStorePasswordGenericJdbc, error) {
	var baseObject oci_database_tools.DatabaseToolsKeyStorePasswordGenericJdbc
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
		details := oci_database_tools.DatabaseToolsKeyStorePasswordSecretIdGenericJdbc{}
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

func DatabaseToolsKeyStorePasswordGenericJdbcToMap(obj *oci_database_tools.DatabaseToolsKeyStorePasswordGenericJdbc) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (*obj).(type) {
	case oci_database_tools.DatabaseToolsKeyStorePasswordSecretIdGenericJdbc:
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

func (s *DatabaseToolsDatabaseToolsConnectionResourceCrud) mapToDatabaseToolsKeyStorePasswordGenericJdbcDetails(fieldKeyFormat string) (oci_database_tools.DatabaseToolsKeyStorePasswordGenericJdbcDetails, error) {
	var baseObject oci_database_tools.DatabaseToolsKeyStorePasswordGenericJdbcDetails
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
		details := oci_database_tools.DatabaseToolsKeyStorePasswordSecretIdGenericJdbcDetails{}
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

func DatabaseToolsKeyStorePasswordGenericJdbcDetailsToMap(obj *oci_database_tools.DatabaseToolsKeyStorePasswordGenericJdbcDetails) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (*obj).(type) {
	case oci_database_tools.DatabaseToolsKeyStorePasswordSecretIdGenericJdbcDetails:
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

func (s *DatabaseToolsDatabaseToolsConnectionResourceCrud) mapToDatabaseToolsKeyStorePasswordGenericJdbcSummary(fieldKeyFormat string) (oci_database_tools.DatabaseToolsKeyStorePasswordGenericJdbcSummary, error) {
	var baseObject oci_database_tools.DatabaseToolsKeyStorePasswordGenericJdbcSummary
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
		details := oci_database_tools.DatabaseToolsKeyStorePasswordSecretIdGenericJdbcSummary{}
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

func DatabaseToolsKeyStorePasswordGenericJdbcSummaryToMap(obj *oci_database_tools.DatabaseToolsKeyStorePasswordGenericJdbcSummary) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (*obj).(type) {
	case oci_database_tools.DatabaseToolsKeyStorePasswordSecretIdGenericJdbcSummary:
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

func (s *DatabaseToolsDatabaseToolsConnectionResourceCrud) mapToDatabaseToolsKeyStorePasswordMySql(fieldKeyFormat string) (oci_database_tools.DatabaseToolsKeyStorePasswordMySql, error) {
	var baseObject oci_database_tools.DatabaseToolsKeyStorePasswordMySql
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
		details := oci_database_tools.DatabaseToolsKeyStorePasswordSecretIdMySql{}
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

// For type: MYSQL
func DatabaseToolsKeyStorePasswordMySqlToMap(obj *oci_database_tools.DatabaseToolsKeyStorePasswordMySql) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (*obj).(type) {
	case oci_database_tools.DatabaseToolsKeyStorePasswordSecretIdMySql:
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

// For type: MYSQL
func (s *DatabaseToolsDatabaseToolsConnectionResourceCrud) mapToDatabaseToolsKeyStorePasswordMySqlDetails(fieldKeyFormat string) (oci_database_tools.DatabaseToolsKeyStorePasswordMySqlDetails, error) {
	var baseObject oci_database_tools.DatabaseToolsKeyStorePasswordMySqlDetails
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
		details := oci_database_tools.DatabaseToolsKeyStorePasswordSecretIdMySqlDetails{}
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

func DatabaseToolsKeyStorePasswordMySqlDetailsToMap(obj *oci_database_tools.DatabaseToolsKeyStorePasswordMySqlDetails) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (*obj).(type) {
	case oci_database_tools.DatabaseToolsKeyStorePasswordSecretIdMySqlDetails:
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

func (s *DatabaseToolsDatabaseToolsConnectionResourceCrud) mapToDatabaseToolsKeyStorePasswordMySqlSummary(fieldKeyFormat string) (oci_database_tools.DatabaseToolsKeyStorePasswordMySqlSummary, error) {
	var baseObject oci_database_tools.DatabaseToolsKeyStorePasswordMySqlSummary
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
		details := oci_database_tools.DatabaseToolsKeyStorePasswordSecretIdMySqlSummary{}
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

func DatabaseToolsKeyStorePasswordMySqlSummaryToMap(obj *oci_database_tools.DatabaseToolsKeyStorePasswordMySqlSummary) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (*obj).(type) {
	case oci_database_tools.DatabaseToolsKeyStorePasswordSecretIdMySqlSummary:
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

func (s *DatabaseToolsDatabaseToolsConnectionResourceCrud) mapToDatabaseToolsKeyStorePasswordPostgresql(fieldKeyFormat string) (oci_database_tools.DatabaseToolsKeyStorePasswordPostgresql, error) {
	var baseObject oci_database_tools.DatabaseToolsKeyStorePasswordPostgresql
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
		details := oci_database_tools.DatabaseToolsKeyStorePasswordSecretIdPostgresql{}
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

func DatabaseToolsKeyStorePasswordPostgresqlToMap(obj *oci_database_tools.DatabaseToolsKeyStorePasswordPostgresql) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (*obj).(type) {
	case oci_database_tools.DatabaseToolsKeyStorePasswordSecretIdPostgresql:
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

func (s *DatabaseToolsDatabaseToolsConnectionResourceCrud) mapToDatabaseToolsKeyStorePasswordPostgresqlDetails(fieldKeyFormat string) (oci_database_tools.DatabaseToolsKeyStorePasswordPostgresqlDetails, error) {
	var baseObject oci_database_tools.DatabaseToolsKeyStorePasswordPostgresqlDetails
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
		details := oci_database_tools.DatabaseToolsKeyStorePasswordSecretIdPostgresqlDetails{}
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

func DatabaseToolsKeyStorePasswordPostgresqlDetailsToMap(obj *oci_database_tools.DatabaseToolsKeyStorePasswordPostgresqlDetails) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (*obj).(type) {
	case oci_database_tools.DatabaseToolsKeyStorePasswordSecretIdPostgresqlDetails:
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

func (s *DatabaseToolsDatabaseToolsConnectionResourceCrud) mapToDatabaseToolsKeyStorePasswordPostgresqlSummary(fieldKeyFormat string) (oci_database_tools.DatabaseToolsKeyStorePasswordPostgresqlSummary, error) {
	var baseObject oci_database_tools.DatabaseToolsKeyStorePasswordPostgresqlSummary
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
		details := oci_database_tools.DatabaseToolsKeyStorePasswordSecretIdPostgresqlSummary{}
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

func DatabaseToolsKeyStorePasswordPostgresqlSummaryToMap(obj *oci_database_tools.DatabaseToolsKeyStorePasswordPostgresqlSummary) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (*obj).(type) {
	case oci_database_tools.DatabaseToolsKeyStorePasswordSecretIdPostgresqlSummary:
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

// For type: ORACLE_DATABASE
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

// For type: ORACLE_DATABASE
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

func (s *DatabaseToolsDatabaseToolsConnectionResourceCrud) mapToDatabaseToolsKeyStorePostgresql(fieldKeyFormat string) (oci_database_tools.DatabaseToolsKeyStorePostgresql, error) {
	result := oci_database_tools.DatabaseToolsKeyStorePostgresql{}

	if keyStoreContent, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "key_store_content")); ok {
		if tmpList := keyStoreContent.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "key_store_content"), 0)
			tmp, err := s.mapToDatabaseToolsKeyStoreContentPostgresql(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert key_store_content, encountered error: %v", err)
			}
			result.KeyStoreContent = tmp
		}
	}

	if keyStorePassword, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "key_store_password")); ok {
		if tmpList := keyStorePassword.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "key_store_password"), 0)
			tmp, err := s.mapToDatabaseToolsKeyStorePasswordPostgresql(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert key_store_password, encountered error: %v", err)
			}
			result.KeyStorePassword = tmp
		}
	}

	if keyStoreType, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "key_store_type")); ok {
		result.KeyStoreType = oci_database_tools.KeyStoreTypePostgresqlEnum(keyStoreType.(string))
	}

	return result, nil
}

func DatabaseToolsKeyStorePostgresqlToMap(obj oci_database_tools.DatabaseToolsKeyStorePostgresql) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.KeyStoreContent != nil {
		keyStoreContentArray := []interface{}{}
		if keyStoreContentMap := DatabaseToolsKeyStoreContentPostgresqlToMap(&obj.KeyStoreContent); keyStoreContentMap != nil {
			keyStoreContentArray = append(keyStoreContentArray, keyStoreContentMap)
		}
		result["key_store_content"] = keyStoreContentArray
	}

	if obj.KeyStorePassword != nil {
		keyStorePasswordArray := []interface{}{}
		if keyStorePasswordMap := DatabaseToolsKeyStorePasswordPostgresqlToMap(&obj.KeyStorePassword); keyStorePasswordMap != nil {
			keyStorePasswordArray = append(keyStorePasswordArray, keyStorePasswordMap)
		}
		result["key_store_password"] = keyStorePasswordArray
	}

	result["key_store_type"] = string(obj.KeyStoreType)

	return result
}

func (s *DatabaseToolsDatabaseToolsConnectionResourceCrud) mapToDatabaseToolsKeyStorePostgresqlDetails(fieldKeyFormat string) (oci_database_tools.DatabaseToolsKeyStorePostgresqlDetails, error) {
	result := oci_database_tools.DatabaseToolsKeyStorePostgresqlDetails{}

	if keyStoreContent, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "key_store_content")); ok {
		if tmpList := keyStoreContent.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "key_store_content"), 0)
			tmp, err := s.mapToDatabaseToolsKeyStoreContentPostgresqlDetails(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert key_store_content, encountered error: %v", err)
			}
			result.KeyStoreContent = tmp
		}
	}

	if keyStorePassword, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "key_store_password")); ok {
		if tmpList := keyStorePassword.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "key_store_password"), 0)
			tmp, err := s.mapToDatabaseToolsKeyStorePasswordPostgresqlDetails(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert key_store_password, encountered error: %v", err)
			}
			result.KeyStorePassword = tmp
		}
	}

	if keyStoreType, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "key_store_type")); ok {
		result.KeyStoreType = oci_database_tools.KeyStoreTypePostgresqlEnum(keyStoreType.(string))
	}

	return result, nil
}

func DatabaseToolsKeyStorePostgresqlDetailsToMap(obj oci_database_tools.DatabaseToolsKeyStorePostgresqlDetails) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.KeyStoreContent != nil {
		keyStoreContentArray := []interface{}{}
		if keyStoreContentMap := DatabaseToolsKeyStoreContentPostgresqlDetailsToMap(&obj.KeyStoreContent); keyStoreContentMap != nil {
			keyStoreContentArray = append(keyStoreContentArray, keyStoreContentMap)
		}
		result["key_store_content"] = keyStoreContentArray
	}

	if obj.KeyStorePassword != nil {
		keyStorePasswordArray := []interface{}{}
		if keyStorePasswordMap := DatabaseToolsKeyStorePasswordPostgresqlDetailsToMap(&obj.KeyStorePassword); keyStorePasswordMap != nil {
			keyStorePasswordArray = append(keyStorePasswordArray, keyStorePasswordMap)
		}
		result["key_store_password"] = keyStorePasswordArray
	}

	result["key_store_type"] = string(obj.KeyStoreType)

	return result
}

func (s *DatabaseToolsDatabaseToolsConnectionResourceCrud) mapToDatabaseToolsKeyStorePostgresqlSummary(fieldKeyFormat string) (oci_database_tools.DatabaseToolsKeyStorePostgresqlSummary, error) {
	result := oci_database_tools.DatabaseToolsKeyStorePostgresqlSummary{}

	if keyStoreContent, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "key_store_content")); ok {
		if tmpList := keyStoreContent.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "key_store_content"), 0)
			tmp, err := s.mapToDatabaseToolsKeyStoreContentPostgresqlSummary(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert key_store_content, encountered error: %v", err)
			}
			result.KeyStoreContent = tmp
		}
	}

	if keyStorePassword, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "key_store_password")); ok {
		if tmpList := keyStorePassword.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "key_store_password"), 0)
			tmp, err := s.mapToDatabaseToolsKeyStorePasswordPostgresqlSummary(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert key_store_password, encountered error: %v", err)
			}
			result.KeyStorePassword = tmp
		}
	}

	if keyStoreType, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "key_store_type")); ok {
		result.KeyStoreType = oci_database_tools.KeyStoreTypePostgresqlEnum(keyStoreType.(string))
	}

	return result, nil
}

func DatabaseToolsKeyStorePostgresqlSummaryToMap(obj oci_database_tools.DatabaseToolsKeyStorePostgresqlSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.KeyStoreContent != nil {
		keyStoreContentArray := []interface{}{}
		if keyStoreContentMap := DatabaseToolsKeyStoreContentPostgresqlSummaryToMap(&obj.KeyStoreContent); keyStoreContentMap != nil {
			keyStoreContentArray = append(keyStoreContentArray, keyStoreContentMap)
		}
		result["key_store_content"] = keyStoreContentArray
	}

	if obj.KeyStorePassword != nil {
		keyStorePasswordArray := []interface{}{}
		if keyStorePasswordMap := DatabaseToolsKeyStorePasswordPostgresqlSummaryToMap(&obj.KeyStorePassword); keyStorePasswordMap != nil {
			keyStorePasswordArray = append(keyStorePasswordArray, keyStorePasswordMap)
		}
		result["key_store_password"] = keyStorePasswordArray
	}

	result["key_store_type"] = string(obj.KeyStoreType)

	return result
}

// For type: ORACLE_DATABASE
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

// For type: ORACLE_DATABASE
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

// For type: ORACLE_DATABASE
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

// For type: ORACLE_DATABASE
func DatabaseToolsRelatedResourceToMap(obj *oci_database_tools.DatabaseToolsRelatedResource) map[string]interface{} {
	result := map[string]interface{}{}

	result["entity_type"] = string(obj.EntityType)

	if obj.Identifier != nil {
		result["identifier"] = string(*obj.Identifier)
	}

	return result
}

// For type: MYSQL
func (s *DatabaseToolsDatabaseToolsConnectionResourceCrud) mapToDatabaseToolsRelatedResourceMySql(fieldKeyFormat string) (oci_database_tools.DatabaseToolsRelatedResourceMySql, error) {
	result := oci_database_tools.DatabaseToolsRelatedResourceMySql{}

	if entityType, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "entity_type")); ok {
		result.EntityType = oci_database_tools.RelatedResourceEntityTypeMySqlEnum(entityType.(string))
	}

	if identifier, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "identifier")); ok {
		tmp := identifier.(string)
		result.Identifier = &tmp
	}

	return result, nil
}

// For type: MYSQL
func DatabaseToolsRelatedResourceMySqlToMap(obj *oci_database_tools.DatabaseToolsRelatedResourceMySql) map[string]interface{} {
	result := map[string]interface{}{}

	result["entity_type"] = string(obj.EntityType)

	if obj.Identifier != nil {
		result["identifier"] = string(*obj.Identifier)
	}

	return result
}

func (s *DatabaseToolsDatabaseToolsConnectionResourceCrud) mapToDatabaseToolsRelatedResourcePostgresql(fieldKeyFormat string) (oci_database_tools.DatabaseToolsRelatedResourcePostgresql, error) {
	result := oci_database_tools.DatabaseToolsRelatedResourcePostgresql{}

	if entityType, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "entity_type")); ok {
		result.EntityType = oci_database_tools.RelatedResourceEntityTypePostgresqlEnum(entityType.(string))
	}

	if identifier, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "identifier")); ok {
		tmp := identifier.(string)
		result.Identifier = &tmp
	}

	return result, nil
}

func DatabaseToolsRelatedResourcePostgresqlToMap(obj *oci_database_tools.DatabaseToolsRelatedResourcePostgresql) map[string]interface{} {
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

func (s *DatabaseToolsDatabaseToolsConnectionResourceCrud) mapToResourceLock(fieldKeyFormat string) (oci_database_tools.ResourceLock, error) {
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

func ConnectionResourceLockToMap(obj oci_database_tools.ResourceLock) map[string]interface{} {
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

// For type: ORACLE_DATABASE
func UpdateDatabaseToolsRelatedResourceDetailsToMap(obj *oci_database_tools.UpdateDatabaseToolsRelatedResourceDetails) map[string]interface{} {
	result := map[string]interface{}{}

	result["entity_type"] = string(obj.EntityType)

	if obj.Identifier != nil {
		result["identifier"] = string(*obj.Identifier)
	}

	return result
}

// For type: MYSQL
func (s *DatabaseToolsDatabaseToolsConnectionResourceCrud) mapToUpdateDatabaseToolsRelatedResourceMySqlDetails(fieldKeyFormat string) (oci_database_tools.UpdateDatabaseToolsRelatedResourceMySqlDetails, error) {
	result := oci_database_tools.UpdateDatabaseToolsRelatedResourceMySqlDetails{}

	if entityType, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "entity_type")); ok {
		result.EntityType = oci_database_tools.RelatedResourceEntityTypeMySqlEnum(entityType.(string))
	}

	if identifier, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "identifier")); ok {
		tmp := identifier.(string)
		result.Identifier = &tmp
	}

	return result, nil
}

// For type: MYSQL
func UpdateDatabaseToolsRelatedResourceMySqlDetailsToMap(obj *oci_database_tools.UpdateDatabaseToolsRelatedResourceMySqlDetails) map[string]interface{} {
	result := map[string]interface{}{}

	result["entity_type"] = string(obj.EntityType)

	if obj.Identifier != nil {
		result["identifier"] = string(*obj.Identifier)
	}

	return result
}

func (s *DatabaseToolsDatabaseToolsConnectionResourceCrud) mapToUpdateDatabaseToolsRelatedResourcePostgresqlDetails(fieldKeyFormat string) (oci_database_tools.UpdateDatabaseToolsRelatedResourcePostgresqlDetails, error) {
	result := oci_database_tools.UpdateDatabaseToolsRelatedResourcePostgresqlDetails{}

	if entityType, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "entity_type")); ok {
		result.EntityType = oci_database_tools.RelatedResourceEntityTypePostgresqlEnum(entityType.(string))
	}

	if identifier, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "identifier")); ok {
		tmp := identifier.(string)
		result.Identifier = &tmp
	}

	return result, nil
}

func UpdateDatabaseToolsRelatedResourcePostgresqlDetailsToMap(obj *oci_database_tools.UpdateDatabaseToolsRelatedResourcePostgresqlDetails) map[string]interface{} {
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
	case strings.ToLower("GENERIC_JDBC"):
		details := oci_database_tools.CreateDatabaseToolsConnectionGenericJdbcDetails{}
		if advancedProperties, ok := s.D.GetOkExists("advanced_properties"); ok {
			details.AdvancedProperties = tfresource.ObjectMapToStringMap(advancedProperties.(map[string]interface{}))
		}
		if keyStores, ok := s.D.GetOkExists("key_stores"); ok {
			interfaces := keyStores.([]interface{})
			tmp := make([]oci_database_tools.DatabaseToolsKeyStoreGenericJdbcDetails, len(interfaces))
			for i := range interfaces {
				stateDataIndex := i
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "key_stores", stateDataIndex)
				converted, err := s.mapToDatabaseToolsKeyStoreGenericJdbcDetails(fieldKeyFormat)
				if err != nil {
					return err
				}
				tmp[i] = converted
			}
			if len(tmp) != 0 || s.D.HasChange("key_stores") {
				details.KeyStores = tmp
			}
		}
		if url, ok := s.D.GetOkExists("url"); ok {
			tmp := url.(string)
			details.Url = &tmp
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
		if runtimeSupport, ok := s.D.GetOkExists("runtime_support"); ok {
			details.RuntimeSupport = oci_database_tools.RuntimeSupportEnum(runtimeSupport.(string))
		}
		request.CreateDatabaseToolsConnectionDetails = details
	case strings.ToLower("MYSQL"):
		details := oci_database_tools.CreateDatabaseToolsConnectionMySqlDetails{}
		if advancedProperties, ok := s.D.GetOkExists("advanced_properties"); ok {
			details.AdvancedProperties = tfresource.ObjectMapToStringMap(advancedProperties.(map[string]interface{}))
		}
		if connectionString, ok := s.D.GetOkExists("connection_string"); ok {
			tmp := connectionString.(string)
			details.ConnectionString = &tmp
		}
		if keyStores, ok := s.D.GetOkExists("key_stores"); ok {
			interfaces := keyStores.([]interface{})
			tmp := make([]oci_database_tools.DatabaseToolsKeyStoreMySqlDetails, len(interfaces))
			for i := range interfaces {
				stateDataIndex := i
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "key_stores", stateDataIndex)
				converted, err := s.mapToDatabaseToolsKeyStoreMySqlDetails(fieldKeyFormat)
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
				tmp, err := s.mapToCreateDatabaseToolsRelatedResourceMySqlDetails(fieldKeyFormat)
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
		if runtimeSupport, ok := s.D.GetOkExists("runtime_support"); ok {
			details.RuntimeSupport = oci_database_tools.RuntimeSupportEnum(runtimeSupport.(string))
		}
		request.CreateDatabaseToolsConnectionDetails = details
	case strings.ToLower("ORACLE_DATABASE"):
		details := oci_database_tools.CreateDatabaseToolsConnectionOracleDatabaseDetails{}
		if advancedProperties, ok := s.D.GetOkExists("advanced_properties"); ok {
			details.AdvancedProperties = tfresource.ObjectMapToStringMap(advancedProperties.(map[string]interface{}))
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
		if proxyClient, ok := s.D.GetOkExists("proxy_client"); ok {
			if tmpList := proxyClient.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "proxy_client", 0)
				tmp, err := s.mapToDatabaseToolsConnectionOracleDatabaseProxyClientDetails(fieldKeyFormat)
				if err != nil {
					return err
				}
				details.ProxyClient = tmp
			}
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
		if runtimeSupport, ok := s.D.GetOkExists("runtime_support"); ok {
			details.RuntimeSupport = oci_database_tools.RuntimeSupportEnum(runtimeSupport.(string))
		}
		request.CreateDatabaseToolsConnectionDetails = details
	case strings.ToLower("POSTGRESQL"):
		details := oci_database_tools.CreateDatabaseToolsConnectionPostgresqlDetails{}
		if advancedProperties, ok := s.D.GetOkExists("advanced_properties"); ok {
			details.AdvancedProperties = tfresource.ObjectMapToStringMap(advancedProperties.(map[string]interface{}))
		}
		if connectionString, ok := s.D.GetOkExists("connection_string"); ok {
			tmp := connectionString.(string)
			details.ConnectionString = &tmp
		}
		if keyStores, ok := s.D.GetOkExists("key_stores"); ok {
			interfaces := keyStores.([]interface{})
			tmp := make([]oci_database_tools.DatabaseToolsKeyStorePostgresqlDetails, len(interfaces))
			for i := range interfaces {
				stateDataIndex := i
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "key_stores", stateDataIndex)
				converted, err := s.mapToDatabaseToolsKeyStorePostgresqlDetails(fieldKeyFormat)
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
				tmp, err := s.mapToCreateDatabaseToolsRelatedResourcePostgresqlDetails(fieldKeyFormat)
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
		if runtimeSupport, ok := s.D.GetOkExists("runtime_support"); ok {
			details.RuntimeSupport = oci_database_tools.RuntimeSupportEnum(runtimeSupport.(string))
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
	case strings.ToLower("GENERIC_JDBC"):
		details := oci_database_tools.UpdateDatabaseToolsConnectionGenericJdbcDetails{}
		if advancedProperties, ok := s.D.GetOkExists("advanced_properties"); ok {
			details.AdvancedProperties = tfresource.ObjectMapToStringMap(advancedProperties.(map[string]interface{}))
		}
		if keyStores, ok := s.D.GetOkExists("key_stores"); ok {
			interfaces := keyStores.([]interface{})
			tmp := make([]oci_database_tools.DatabaseToolsKeyStoreGenericJdbcDetails, len(interfaces))
			for i := range interfaces {
				stateDataIndex := i
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "key_stores", stateDataIndex)
				converted, err := s.mapToDatabaseToolsKeyStoreGenericJdbcDetails(fieldKeyFormat)
				if err != nil {
					return err
				}
				tmp[i] = converted
			}
			if len(tmp) != 0 || s.D.HasChange("key_stores") {
				details.KeyStores = tmp
			}
		}
		if url, ok := s.D.GetOkExists("url"); ok {
			tmp := url.(string)
			details.Url = &tmp
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
			details.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
		}
		if isLockOverride, ok := s.D.GetOkExists("is_lock_override"); ok {
			tmp := isLockOverride.(bool)
			request.IsLockOverride = &tmp
		}
		request.UpdateDatabaseToolsConnectionDetails = details
	case strings.ToLower("MYSQL"):
		details := oci_database_tools.UpdateDatabaseToolsConnectionMySqlDetails{}
		if advancedProperties, ok := s.D.GetOkExists("advanced_properties"); ok {
			details.AdvancedProperties = tfresource.ObjectMapToStringMap(advancedProperties.(map[string]interface{}))
		}
		if connectionString, ok := s.D.GetOkExists("connection_string"); ok {
			tmp := connectionString.(string)
			details.ConnectionString = &tmp
		}
		if keyStores, ok := s.D.GetOkExists("key_stores"); ok {
			interfaces := keyStores.([]interface{})
			tmp := make([]oci_database_tools.DatabaseToolsKeyStoreMySqlDetails, len(interfaces))
			for i := range interfaces {
				stateDataIndex := i
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "key_stores", stateDataIndex)
				converted, err := s.mapToDatabaseToolsKeyStoreMySqlDetails(fieldKeyFormat)
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
				tmp, err := s.mapToUpdateDatabaseToolsRelatedResourceMySqlDetails(fieldKeyFormat)
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
			details.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
		}
		if isLockOverride, ok := s.D.GetOkExists("is_lock_override"); ok {
			tmp := isLockOverride.(bool)
			request.IsLockOverride = &tmp
		}
		request.UpdateDatabaseToolsConnectionDetails = details
	case strings.ToLower("ORACLE_DATABASE"):
		details := oci_database_tools.UpdateDatabaseToolsConnectionOracleDatabaseDetails{}
		if advancedProperties, ok := s.D.GetOkExists("advanced_properties"); ok {
			details.AdvancedProperties = tfresource.ObjectMapToStringMap(advancedProperties.(map[string]interface{}))
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
		if proxyClient, ok := s.D.GetOkExists("proxy_client"); ok {
			if tmpList := proxyClient.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "proxy_client", 0)
				tmp, err := s.mapToDatabaseToolsConnectionOracleDatabaseProxyClientDetails(fieldKeyFormat)
				if err != nil {
					return err
				}
				details.ProxyClient = tmp
			}
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
			details.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
		}
		if isLockOverride, ok := s.D.GetOkExists("is_lock_override"); ok {
			tmp := isLockOverride.(bool)
			request.IsLockOverride = &tmp
		}
		request.UpdateDatabaseToolsConnectionDetails = details
	case strings.ToLower("POSTGRESQL"):
		details := oci_database_tools.UpdateDatabaseToolsConnectionPostgresqlDetails{}
		if advancedProperties, ok := s.D.GetOkExists("advanced_properties"); ok {
			details.AdvancedProperties = tfresource.ObjectMapToStringMap(advancedProperties.(map[string]interface{}))
		}
		if connectionString, ok := s.D.GetOkExists("connection_string"); ok {
			tmp := connectionString.(string)
			details.ConnectionString = &tmp
		}
		if keyStores, ok := s.D.GetOkExists("key_stores"); ok {
			interfaces := keyStores.([]interface{})
			tmp := make([]oci_database_tools.DatabaseToolsKeyStorePostgresqlDetails, len(interfaces))
			for i := range interfaces {
				stateDataIndex := i
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "key_stores", stateDataIndex)
				converted, err := s.mapToDatabaseToolsKeyStorePostgresqlDetails(fieldKeyFormat)
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
				tmp, err := s.mapToUpdateDatabaseToolsRelatedResourcePostgresqlDetails(fieldKeyFormat)
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
			details.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
		}
		if isLockOverride, ok := s.D.GetOkExists("is_lock_override"); ok {
			tmp := isLockOverride.(bool)
			request.IsLockOverride = &tmp
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

	if isLockOverride, ok := s.D.GetOkExists("is_lock_override"); ok {
		tmp := isLockOverride.(bool)
		changeCompartmentRequest.IsLockOverride = &tmp
	}

	changeCompartmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database_tools")

	response, err := s.Client.ChangeDatabaseToolsConnectionCompartment(context.Background(), changeCompartmentRequest)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getDatabaseToolsConnectionFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database_tools"), oci_database_tools.ActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}
