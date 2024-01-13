// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package data_safe

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
	oci_data_safe "github.com/oracle/oci-go-sdk/v65/datasafe"
)

func DataSafeTargetDatabaseResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createDataSafeTargetDatabase,
		Read:     readDataSafeTargetDatabase,
		Update:   updateDataSafeTargetDatabase,
		Delete:   deleteDataSafeTargetDatabase,
		Schema: map[string]*schema.Schema{
			// Required
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"database_details": {
				Type:     schema.TypeList,
				Required: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"database_type": {
							Type:             schema.TypeString,
							Required:         true,
							DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
							ValidateFunc: validation.StringInSlice([]string{
								"AUTONOMOUS_DATABASE",
								"DATABASE_CLOUD_SERVICE",
								"INSTALLED_DATABASE",
							}, true),
						},
						"infrastructure_type": {
							Type:     schema.TypeString,
							Required: true,
						},

						// Optional
						"autonomous_database_id": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"db_system_id": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"instance_id": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"ip_addresses": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"listener_port": {
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"service_name": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"vm_cluster_id": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},

						// Computed
					},
				},
			},

			// Optional
			"connection_option": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"connection_type": {
							Type:             schema.TypeString,
							Required:         true,
							DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
							ValidateFunc: validation.StringInSlice([]string{
								"ONPREM_CONNECTOR",
								"PRIVATE_ENDPOINT",
							}, true),
						},

						// Optional
						"datasafe_private_endpoint_id": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"on_prem_connector_id": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},

						// Computed
					},
				},
			},
			"credentials": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"password": {
							Type:      schema.TypeString,
							Required:  true,
							Sensitive: true,
						},
						"user_name": {
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
			"peer_target_database_details": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				ForceNew: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"database_details": {
							Type:     schema.TypeList,
							Required: true,
							ForceNew: true,
							MaxItems: 1,
							MinItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required
									"database_type": {
										Type:             schema.TypeString,
										Required:         true,
										ForceNew:         true,
										DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
										ValidateFunc: validation.StringInSlice([]string{
											"AUTONOMOUS_DATABASE",
											"DATABASE_CLOUD_SERVICE",
											"INSTALLED_DATABASE",
										}, true),
									},
									"infrastructure_type": {
										Type:     schema.TypeString,
										Required: true,
										ForceNew: true,
									},

									// Optional
									"autonomous_database_id": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
										ForceNew: true,
									},
									"db_system_id": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
										ForceNew: true,
									},
									"instance_id": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
										ForceNew: true,
									},
									"ip_addresses": {
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										ForceNew: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"listener_port": {
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
										ForceNew: true,
									},
									"service_name": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
										ForceNew: true,
									},
									"vm_cluster_id": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
										ForceNew: true,
									},

									// Computed
								},
							},
						},

						// Optional
						"dataguard_association_id": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},
						"description": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},
						"display_name": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},
						"tls_config": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							ForceNew: true,
							MaxItems: 1,
							MinItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required
									"status": {
										Type:     schema.TypeString,
										Required: true,
										ForceNew: true,
									},

									// Optional
									"certificate_store_type": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
										ForceNew: true,
									},
									"key_store_content": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
										ForceNew: true,
									},
									"store_password": {
										Type:      schema.TypeString,
										Optional:  true,
										Computed:  true,
										ForceNew:  true,
										Sensitive: true,
									},
									"trust_store_content": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
										ForceNew: true,
									},

									// Computed
								},
							},
						},

						// Computed
					},
				},
			},
			"tls_config": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"status": {
							Type:     schema.TypeString,
							Required: true,
						},

						// Optional
						"certificate_store_type": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"key_store_content": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"store_password": {
							Type:      schema.TypeString,
							Optional:  true,
							Computed:  true,
							Sensitive: true,
						},
						"trust_store_content": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},

						// Computed
					},
				},
			},

			// Computed
			"associated_resource_ids": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"lifecycle_details": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"peer_target_databases": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"database_details": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"autonomous_database_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"database_type": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"db_system_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"infrastructure_type": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"instance_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"ip_addresses": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"listener_port": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									"service_name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"vm_cluster_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"database_unique_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"dataguard_association_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"description": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"display_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"key": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"lifecycle_details": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"role": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"state": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"time_created": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"tls_config": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"certificate_store_type": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"key_store_content": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"status": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"store_password": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"trust_store_content": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
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

func createDataSafeTargetDatabase(d *schema.ResourceData, m interface{}) error {
	sync := &DataSafeTargetDatabaseResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataSafeClient()

	return tfresource.CreateResource(d, sync)
}

func readDataSafeTargetDatabase(d *schema.ResourceData, m interface{}) error {
	sync := &DataSafeTargetDatabaseResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataSafeClient()

	return tfresource.ReadResource(sync)
}

func updateDataSafeTargetDatabase(d *schema.ResourceData, m interface{}) error {
	sync := &DataSafeTargetDatabaseResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataSafeClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteDataSafeTargetDatabase(d *schema.ResourceData, m interface{}) error {
	sync := &DataSafeTargetDatabaseResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataSafeClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type DataSafeTargetDatabaseResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_data_safe.DataSafeClient
	Res                    *oci_data_safe.TargetDatabase
	DisableNotFoundRetries bool
}

func (s *DataSafeTargetDatabaseResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *DataSafeTargetDatabaseResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_data_safe.TargetDatabaseLifecycleStateCreating),
	}
}

func (s *DataSafeTargetDatabaseResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_data_safe.TargetDatabaseLifecycleStateActive),
	}
}

func (s *DataSafeTargetDatabaseResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_data_safe.TargetDatabaseLifecycleStateDeleting),
	}
}

func (s *DataSafeTargetDatabaseResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_data_safe.TargetDatabaseLifecycleStateDeleted),
	}
}

func (s *DataSafeTargetDatabaseResourceCrud) Create() error {
	request := oci_data_safe.CreateTargetDatabaseRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if connectionOption, ok := s.D.GetOkExists("connection_option"); ok {
		if tmpList := connectionOption.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "connection_option", 0)
			tmp, err := s.mapToConnectionOption(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.ConnectionOption = tmp
		}
	}

	if credentials, ok := s.D.GetOkExists("credentials"); ok {
		if tmpList := credentials.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "credentials", 0)
			tmp, err := s.mapToCredentials(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.Credentials = &tmp
		}
	}

	if databaseDetails, ok := s.D.GetOkExists("database_details"); ok {
		if tmpList := databaseDetails.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "database_details", 0)
			tmp, err := s.mapToDatabaseDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.DatabaseDetails = tmp
		}
	}

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

	if peerTargetDatabaseDetails, ok := s.D.GetOkExists("peer_target_database_details"); ok {
		interfaces := peerTargetDatabaseDetails.([]interface{})
		tmp := make([]oci_data_safe.CreatePeerTargetDatabaseDetails, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "peer_target_database_details", stateDataIndex)
			converted, err := s.mapToCreatePeerTargetDatabaseDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("peer_target_database_details") {
			request.PeerTargetDatabaseDetails = tmp
		}
	}

	if tlsConfig, ok := s.D.GetOkExists("tls_config"); ok {
		if tmpList := tlsConfig.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "tls_config", 0)
			tmp, err := s.mapToTlsConfig(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.TlsConfig = &tmp
		}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "data_safe")

	response, err := s.Client.CreateTargetDatabase(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	var identifier *string
	identifier = response.Id
	if identifier != nil {
		s.D.SetId(*identifier)
	}
	return s.getTargetDatabaseFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "data_safe"), oci_data_safe.WorkRequestResourceActionTypeCreated, s.D.Timeout(schema.TimeoutCreate))
}

func (s *DataSafeTargetDatabaseResourceCrud) getTargetDatabaseFromWorkRequest(workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_data_safe.WorkRequestResourceActionTypeEnum, timeout time.Duration) error {

	// Wait until it finishes
	targetDatabaseId, err := targetDatabaseWaitForWorkRequest(workId, "target-database",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.Client)

	if err != nil {
		return err
	}
	s.D.SetId(*targetDatabaseId)

	return s.Get()
}

func targetDatabaseWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
	startTime := time.Now()
	stopTime := startTime.Add(timeout)
	return func(response oci_common.OCIOperationResponse) bool {

		// Stop after timeout has elapsed
		if time.Now().After(stopTime) {
			return false
		}

		// Make sure we stop on default rules
		if tfresource.ShouldRetry(response, false, "data_safe", startTime) {
			return true
		}

		// Only stop if the time Finished is set
		if workRequestResponse, ok := response.Response.(oci_data_safe.GetWorkRequestResponse); ok {
			return workRequestResponse.TimeFinished == nil
		}
		return false
	}
}

func targetDatabaseWaitForWorkRequest(wId *string, entityType string, action oci_data_safe.WorkRequestResourceActionTypeEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_data_safe.DataSafeClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "data_safe")
	retryPolicy.ShouldRetryOperation = targetDatabaseWorkRequestShouldRetryFunc(timeout)

	response := oci_data_safe.GetWorkRequestResponse{}
	stateConf := &resource.StateChangeConf{
		Pending: []string{
			string(oci_data_safe.WorkRequestStatusInProgress),
			string(oci_data_safe.WorkRequestStatusAccepted),
		},
		Target: []string{
			string(oci_data_safe.WorkRequestStatusSucceeded),
			string(oci_data_safe.WorkRequestStatusFailed),
		},
		Refresh: func() (interface{}, string, error) {
			var err error
			response, err = client.GetWorkRequest(context.Background(),
				oci_data_safe.GetWorkRequestRequest{
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
	if identifier == nil || response.Status == oci_data_safe.WorkRequestStatusFailed {
		return nil, getErrorFromDataSafeTargetDatabaseWorkRequest(client, wId, retryPolicy, entityType, action)
	}

	return identifier, nil
}

func getErrorFromDataSafeTargetDatabaseWorkRequest(client *oci_data_safe.DataSafeClient, workId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_data_safe.WorkRequestResourceActionTypeEnum) error {
	response, err := client.ListWorkRequestErrors(context.Background(),
		oci_data_safe.ListWorkRequestErrorsRequest{
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

func (s *DataSafeTargetDatabaseResourceCrud) Get() error {
	request := oci_data_safe.GetTargetDatabaseRequest{}

	tmp := s.D.Id()
	request.TargetDatabaseId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "data_safe")

	response, err := s.Client.GetTargetDatabase(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.TargetDatabase
	return nil
}

func (s *DataSafeTargetDatabaseResourceCrud) Update() error {
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(compartment)
			if err != nil {
				return err
			}
		}
	}
	request := oci_data_safe.UpdateTargetDatabaseRequest{}

	if connectionOption, ok := s.D.GetOkExists("connection_option"); ok {
		if tmpList := connectionOption.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "connection_option", 0)
			tmp, err := s.mapToConnectionOption(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.ConnectionOption = tmp
		}
	}

	if credentials, ok := s.D.GetOkExists("credentials"); ok {
		if tmpList := credentials.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "credentials", 0)
			tmp, err := s.mapToCredentials(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.Credentials = &tmp
		}
	}

	if databaseDetails, ok := s.D.GetOkExists("database_details"); ok {
		if tmpList := databaseDetails.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "database_details", 0)
			tmp, err := s.mapToDatabaseDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.DatabaseDetails = tmp
		}
	}

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

	tmp := s.D.Id()
	request.TargetDatabaseId = &tmp

	if tlsConfig, ok := s.D.GetOkExists("tls_config"); ok {
		if tmpList := tlsConfig.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "tls_config", 0)
			tmp, err := s.mapToTlsConfig(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.TlsConfig = &tmp
		}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "data_safe")

	response, err := s.Client.UpdateTargetDatabase(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getTargetDatabaseFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "data_safe"), oci_data_safe.WorkRequestResourceActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}

func (s *DataSafeTargetDatabaseResourceCrud) Delete() error {
	request := oci_data_safe.DeleteTargetDatabaseRequest{}

	tmp := s.D.Id()
	request.TargetDatabaseId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "data_safe")

	response, err := s.Client.DeleteTargetDatabase(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	// Wait until it finishes
	_, delWorkRequestErr := targetDatabaseWaitForWorkRequest(workId, "target-database",
		oci_data_safe.WorkRequestResourceActionTypeDeleted, s.D.Timeout(schema.TimeoutDelete), s.DisableNotFoundRetries, s.Client)
	return delWorkRequestErr
}

func (s *DataSafeTargetDatabaseResourceCrud) SetData() error {
	s.D.Set("associated_resource_ids", s.Res.AssociatedResourceIds)

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.ConnectionOption != nil {
		connectionOptionArray := []interface{}{}
		if connectionOptionMap := ConnectionOptionToMap(&s.Res.ConnectionOption); connectionOptionMap != nil {
			connectionOptionArray = append(connectionOptionArray, connectionOptionMap)
		}
		s.D.Set("connection_option", connectionOptionArray)
	} else {
		s.D.Set("connection_option", nil)
	}

	if s.Res.DatabaseDetails != nil {
		databaseDetailsArray := []interface{}{}
		if databaseDetailsMap := DatabaseDetailsToMap(&s.Res.DatabaseDetails); databaseDetailsMap != nil {
			databaseDetailsArray = append(databaseDetailsArray, databaseDetailsMap)
		}
		s.D.Set("database_details", databaseDetailsArray)
	} else {
		s.D.Set("database_details", nil)
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

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	peerTargetDatabases := []interface{}{}
	for _, item := range s.Res.PeerTargetDatabases {
		peerTargetDatabases = append(peerTargetDatabases, PeerTargetDatabaseToMap(item))
	}
	s.D.Set("peer_target_databases", peerTargetDatabases)

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.SystemTags != nil {
		s.D.Set("system_tags", tfresource.SystemTagsToMap(s.Res.SystemTags))
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	return nil
}

func (s *DataSafeTargetDatabaseResourceCrud) mapToConnectionOption(fieldKeyFormat string) (oci_data_safe.ConnectionOption, error) {
	var baseObject oci_data_safe.ConnectionOption
	//discriminator
	connectionTypeRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "connection_type"))
	var connectionType string
	if ok {
		connectionType = connectionTypeRaw.(string)
	} else {
		connectionType = "" // default value
	}
	switch strings.ToLower(connectionType) {
	case strings.ToLower("ONPREM_CONNECTOR"):
		details := oci_data_safe.OnPremiseConnector{}
		if onPremConnectorId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "on_prem_connector_id")); ok {
			tmp := onPremConnectorId.(string)
			details.OnPremConnectorId = &tmp
		}
		baseObject = details
	case strings.ToLower("PRIVATE_ENDPOINT"):
		details := oci_data_safe.PrivateEndpoint{}
		if datasafePrivateEndpointId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "datasafe_private_endpoint_id")); ok {
			tmp := datasafePrivateEndpointId.(string)
			details.DatasafePrivateEndpointId = &tmp
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown connection_type '%v' was specified", connectionType)
	}
	return baseObject, nil
}

func ConnectionOptionToMap(obj *oci_data_safe.ConnectionOption) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (*obj).(type) {
	case oci_data_safe.OnPremiseConnector:
		result["connection_type"] = "ONPREM_CONNECTOR"

		if v.OnPremConnectorId != nil {
			result["on_prem_connector_id"] = string(*v.OnPremConnectorId)
		}
	case oci_data_safe.PrivateEndpoint:
		result["connection_type"] = "PRIVATE_ENDPOINT"

		if v.DatasafePrivateEndpointId != nil {
			result["datasafe_private_endpoint_id"] = string(*v.DatasafePrivateEndpointId)
		}
	default:
		log.Printf("[WARN] Received 'connection_type' of unknown type %v", *obj)
		return nil
	}

	return result
}

func (s *DataSafeTargetDatabaseResourceCrud) mapToCreatePeerTargetDatabaseDetails(fieldKeyFormat string) (oci_data_safe.CreatePeerTargetDatabaseDetails, error) {
	result := oci_data_safe.CreatePeerTargetDatabaseDetails{}

	if databaseDetails, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "database_details")); ok {
		if tmpList := databaseDetails.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "database_details"), 0)
			tmp, err := s.mapToDatabaseDetails(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert database_details, encountered error: %v", err)
			}
			result.DatabaseDetails = tmp
		}
	}

	if dataguardAssociationId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "dataguard_association_id")); ok {
		tmp := dataguardAssociationId.(string)
		result.DataguardAssociationId = &tmp
	}

	if description, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "description")); ok {
		tmp := description.(string)
		result.Description = &tmp
	}

	if displayName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "display_name")); ok {
		tmp := displayName.(string)
		result.DisplayName = &tmp
	}

	if tlsConfig, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "tls_config")); ok {
		if tmpList := tlsConfig.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "tls_config"), 0)
			tmp, err := s.mapToTlsConfig(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert tls_config, encountered error: %v", err)
			}
			result.TlsConfig = &tmp
		}
	}

	return result, nil
}

func CreatePeerTargetDatabaseDetailsToMap(obj oci_data_safe.CreatePeerTargetDatabaseDetails) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.DatabaseDetails != nil {
		databaseDetailsArray := []interface{}{}
		if databaseDetailsMap := DatabaseDetailsToMap(&obj.DatabaseDetails); databaseDetailsMap != nil {
			databaseDetailsArray = append(databaseDetailsArray, databaseDetailsMap)
		}
		result["database_details"] = databaseDetailsArray
	}

	if obj.DataguardAssociationId != nil {
		result["dataguard_association_id"] = string(*obj.DataguardAssociationId)
	}

	if obj.Description != nil {
		result["description"] = string(*obj.Description)
	}

	if obj.DisplayName != nil {
		result["display_name"] = string(*obj.DisplayName)
	}

	if obj.TlsConfig != nil {
		result["tls_config"] = []interface{}{TlsConfigToMap(obj.TlsConfig)}
	}

	return result
}

func (s *DataSafeTargetDatabaseResourceCrud) mapToCredentials(fieldKeyFormat string) (oci_data_safe.Credentials, error) {
	result := oci_data_safe.Credentials{}

	if password, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "password")); ok {
		tmp := password.(string)
		result.Password = &tmp
	}

	if userName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "user_name")); ok {
		tmp := userName.(string)
		result.UserName = &tmp
	}

	return result, nil
}

func CredentialsToMap(obj *oci_data_safe.Credentials) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Password != nil {
		result["password"] = string(*obj.Password)
	}

	if obj.UserName != nil {
		result["user_name"] = string(*obj.UserName)
	}

	return result
}

func (s *DataSafeTargetDatabaseResourceCrud) mapToDatabaseDetails(fieldKeyFormat string) (oci_data_safe.DatabaseDetails, error) {
	var baseObject oci_data_safe.DatabaseDetails
	//discriminator
	databaseTypeRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "database_type"))
	var databaseType string
	if ok {
		databaseType = databaseTypeRaw.(string)
	} else {
		databaseType = "" // default value
	}
	switch strings.ToLower(databaseType) {
	case strings.ToLower("AUTONOMOUS_DATABASE"):
		details := oci_data_safe.AutonomousDatabaseDetails{}
		if autonomousDatabaseId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "autonomous_database_id")); ok {
			tmp := autonomousDatabaseId.(string)
			details.AutonomousDatabaseId = &tmp
		}
		if infrastructureType, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "infrastructure_type")); ok {
			details.InfrastructureType = oci_data_safe.InfrastructureTypeEnum(infrastructureType.(string))
		}
		baseObject = details
	case strings.ToLower("DATABASE_CLOUD_SERVICE"):
		details := oci_data_safe.DatabaseCloudServiceDetails{}
		if dbSystemId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "db_system_id")); ok {
			tmp := dbSystemId.(string)
			details.DbSystemId = &tmp
		}
		if vmClusterId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "vm_cluster_id")); ok {
			tmp := vmClusterId.(string)
			details.VmClusterId = &tmp
		}
		if listenerPort, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "listener_port")); ok {
			tmp := listenerPort.(int)
			details.ListenerPort = &tmp
		}
		if serviceName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "service_name")); ok {
			tmp := serviceName.(string)
			details.ServiceName = &tmp
		}
		if infrastructureType, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "infrastructure_type")); ok {
			details.InfrastructureType = oci_data_safe.InfrastructureTypeEnum(infrastructureType.(string))
		}
		baseObject = details
	case strings.ToLower("INSTALLED_DATABASE"):
		details := oci_data_safe.InstalledDatabaseDetails{}
		if instanceId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "instance_id")); ok {
			tmp := instanceId.(string)
			details.InstanceId = &tmp
		}
		if ipAddresses, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "ip_addresses")); ok {
			interfaces := ipAddresses.([]interface{})
			tmp := make([]string, len(interfaces))
			for i := range interfaces {
				if interfaces[i] != nil {
					tmp[i] = interfaces[i].(string)
				}
			}
			if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "ip_addresses")) {
				details.IpAddresses = tmp
			}
		}
		if listenerPort, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "listener_port")); ok {
			tmp := listenerPort.(int)
			details.ListenerPort = &tmp
		}
		if serviceName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "service_name")); ok {
			tmp := serviceName.(string)
			details.ServiceName = &tmp
		}
		if infrastructureType, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "infrastructure_type")); ok {
			details.InfrastructureType = oci_data_safe.InfrastructureTypeEnum(infrastructureType.(string))
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown database_type '%v' was specified", databaseType)
	}
	return baseObject, nil
}

func DatabaseDetailsToMap(obj *oci_data_safe.DatabaseDetails) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (*obj).(type) {
	case oci_data_safe.AutonomousDatabaseDetails:
		result["database_type"] = "AUTONOMOUS_DATABASE"

		if v.AutonomousDatabaseId != nil {
			result["autonomous_database_id"] = string(*v.AutonomousDatabaseId)
		}

		result["infrastructure_type"] = string(v.InfrastructureType)
	case oci_data_safe.DatabaseCloudServiceDetails:
		result["database_type"] = "DATABASE_CLOUD_SERVICE"

		if v.DbSystemId != nil {
			result["db_system_id"] = string(*v.DbSystemId)
		}

		if v.VmClusterId != nil {
			result["vm_cluster_id"] = string(*v.VmClusterId)
		}

		if v.ListenerPort != nil {
			result["listener_port"] = int(*v.ListenerPort)
		}

		if v.ServiceName != nil {
			result["service_name"] = string(*v.ServiceName)
		}

		result["infrastructure_type"] = string(v.InfrastructureType)
	case oci_data_safe.InstalledDatabaseDetails:
		result["database_type"] = "INSTALLED_DATABASE"

		if v.InstanceId != nil {
			result["instance_id"] = string(*v.InstanceId)
		}

		result["ip_addresses"] = v.IpAddresses

		if v.ListenerPort != nil {
			result["listener_port"] = int(*v.ListenerPort)
		}

		if v.ServiceName != nil {
			result["service_name"] = string(*v.ServiceName)
		}

		result["infrastructure_type"] = string(v.InfrastructureType)
	default:
		log.Printf("[WARN] Received 'database_type' of unknown type %v", *obj)
		return nil
	}

	return result
}

func PeerTargetDatabaseToMap(obj oci_data_safe.PeerTargetDatabase) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.DatabaseDetails != nil {
		databaseDetailsArray := []interface{}{}
		if databaseDetailsMap := DatabaseDetailsToMap(&obj.DatabaseDetails); databaseDetailsMap != nil {
			databaseDetailsArray = append(databaseDetailsArray, databaseDetailsMap)
		}
		result["database_details"] = databaseDetailsArray
	}

	if obj.DatabaseUniqueName != nil {
		result["database_unique_name"] = string(*obj.DatabaseUniqueName)
	}

	if obj.DataguardAssociationId != nil {
		result["dataguard_association_id"] = string(*obj.DataguardAssociationId)
	}

	if obj.Description != nil {
		result["description"] = string(*obj.Description)
	}

	if obj.DisplayName != nil {
		result["display_name"] = string(*obj.DisplayName)
	}

	if obj.Key != nil {
		result["key"] = int(*obj.Key)
	}

	if obj.LifecycleDetails != nil {
		result["lifecycle_details"] = string(*obj.LifecycleDetails)
	}

	if obj.Role != nil {
		result["role"] = string(*obj.Role)
	}

	result["state"] = string(obj.LifecycleState)

	if obj.TimeCreated != nil {
		result["time_created"] = obj.TimeCreated.String()
	}

	if obj.TlsConfig != nil {
		result["tls_config"] = []interface{}{TlsConfigToMap(obj.TlsConfig)}
	}

	return result
}

func (s *DataSafeTargetDatabaseResourceCrud) mapToTlsConfig(fieldKeyFormat string) (oci_data_safe.TlsConfig, error) {
	result := oci_data_safe.TlsConfig{}

	if certificateStoreType, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "certificate_store_type")); ok {
		result.CertificateStoreType = oci_data_safe.TlsConfigCertificateStoreTypeEnum(certificateStoreType.(string))
	}

	if keyStoreContent, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "key_store_content")); ok {
		tmp := keyStoreContent.(string)
		result.KeyStoreContent = &tmp
	}

	if status, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "status")); ok {
		result.Status = oci_data_safe.TlsConfigStatusEnum(status.(string))
	}

	if storePassword, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "store_password")); ok {
		tmp := storePassword.(string)
		result.StorePassword = &tmp
	}

	if trustStoreContent, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "trust_store_content")); ok {
		tmp := trustStoreContent.(string)
		result.TrustStoreContent = &tmp
	}

	return result, nil
}

func TlsConfigToMap(obj *oci_data_safe.TlsConfig) map[string]interface{} {
	result := map[string]interface{}{}

	result["certificate_store_type"] = string(obj.CertificateStoreType)

	if obj.KeyStoreContent != nil {
		result["key_store_content"] = string(*obj.KeyStoreContent)
	}

	result["status"] = string(obj.Status)

	if obj.StorePassword != nil {
		result["store_password"] = string(*obj.StorePassword)
	}

	if obj.TrustStoreContent != nil {
		result["trust_store_content"] = string(*obj.TrustStoreContent)
	}

	return result
}

func (s *DataSafeTargetDatabaseResourceCrud) updateCompartment(compartment interface{}) error {
	changeCompartmentRequest := oci_data_safe.ChangeTargetDatabaseCompartmentRequest{}

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	idTmp := s.D.Id()
	changeCompartmentRequest.TargetDatabaseId = &idTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "data_safe")

	_, err := s.Client.ChangeTargetDatabaseCompartment(context.Background(), changeCompartmentRequest)
	if err != nil {
		return err
	}

	if waitErr := tfresource.WaitForUpdatedState(s.D, s); waitErr != nil {
		return waitErr
	}

	return nil
}
