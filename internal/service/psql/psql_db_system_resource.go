// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package psql

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"

	oci_common "github.com/oracle/oci-go-sdk/v65/common"
	oci_psql "github.com/oracle/oci-go-sdk/v65/psql"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func PsqlDbSystemResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createPsqlDbSystem,
		Read:     readPsqlDbSystem,
		Update:   updatePsqlDbSystem,
		Delete:   deletePsqlDbSystem,
		Schema: map[string]*schema.Schema{
			// Required
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"db_version": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"network_details": {
				Type:     schema.TypeList,
				Required: true,
				ForceNew: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"subnet_id": {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},

						// Optional
						"nsg_ids": {
							Type:     schema.TypeSet,
							Optional: true,
							Computed: true,
							ForceNew: true,
							Set:      tfresource.LiteralTypeHashCodeForSets,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"primary_db_endpoint_private_ip": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},

						// Computed
					},
				},
			},
			"shape": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"storage_details": {
				Type:     schema.TypeList,
				Required: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"is_regionally_durable": {
							Type:     schema.TypeBool,
							Required: true,
							ForceNew: true,
						},
						"system_type": {
							Type:             schema.TypeString,
							Required:         true,
							ForceNew:         true,
							DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
							ValidateFunc: validation.StringInSlice([]string{
								"OCI_OPTIMIZED_STORAGE",
							}, true),
						},

						// Optional
						"availability_domain": {
							Type:             schema.TypeString,
							Optional:         true,
							ForceNew:         true,
							DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
						},
						"iops": {
							Type:             schema.TypeString,
							Optional:         true,
							Computed:         true,
							ValidateFunc:     tfresource.ValidateInt64TypeString,
							DiffSuppressFunc: tfresource.Int64StringDiffSuppressFunction,
						},

						// Computed
					},
				},
			},

			// Optional
			"config_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"apply_config": {
				Type:     schema.TypeString,
				Optional: true,
				ValidateFunc: validation.StringInSlice([]string{
					"RESTART",
					"RELOAD",
				}, true),
			},
			"credentials": {
				Type:     schema.TypeList,
				Optional: true,
				ForceNew: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"password_details": {
							Type:     schema.TypeList,
							Required: true,
							MaxItems: 1,
							MinItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required
									"password_type": {
										Type:             schema.TypeString,
										Required:         true,
										Sensitive:        true,
										DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
										ValidateFunc: validation.StringInSlice([]string{
											"PLAIN_TEXT",
											"VAULT_SECRET",
										}, true),
									},

									// Optional
									"password": {
										Type:      schema.TypeString,
										Optional:  true,
										Computed:  true,
										Sensitive: true,
									},
									"secret_id": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"secret_version": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									// Computed
								},
							},
						},
						"username": {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
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
			"instance_count": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"instance_memory_size_in_gbs": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"instance_ocpu_count": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"instances_details": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				ForceNew: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional
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
						"private_ip": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},

						// Computed
					},
				},
			},
			"management_policy": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional
						"backup_policy": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							MaxItems: 1,
							MinItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional
									"backup_start": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"days_of_the_month": {
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										MaxItems: 28,
										MinItems: 1,
										Elem: &schema.Schema{
											Type: schema.TypeInt,
										},
									},
									"days_of_the_week": {
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"kind": {
										Type:             schema.TypeString,
										Optional:         true,
										Computed:         true,
										DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
										ValidateFunc: validation.StringInSlice([]string{
											"DAILY",
											"MONTHLY",
											"NONE",
											"WEEKLY",
										}, true),
									},
									"retention_days": {
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
									},

									// Computed
								},
							},
						},
						"maintenance_window_start": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},

						// Computed
					},
				},
			},
			"patch_operations": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"operation": {
							Type:             schema.TypeString,
							Required:         true,
							DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
							ValidateFunc: validation.StringInSlice([]string{
								"INSERT",
								"REMOVE",
							}, true),
						},
						"selection": {
							Type:     schema.TypeString,
							Required: true,
						},

						// Optional
						"value": {
							Type:             schema.TypeMap,
							Optional:         true,
							DiffSuppressFunc: tfresource.JsonStringDiffSuppressFunction,
						},
						"from": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"position": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"selected_item": {
							Type:     schema.TypeString,
							Optional: true,
						},

						// Computed
					},
				},
			},
			"source": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				ForceNew: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"source_type": {
							Type:             schema.TypeString,
							Required:         true,
							ForceNew:         true,
							DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
							ValidateFunc: validation.StringInSlice([]string{
								"BACKUP",
								"NONE",
							}, true),
						},

						// Optional
						"backup_id": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},
						"is_having_restore_config_overrides": {
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},

						// Computed
					},
				},
			},
			"system_type": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},

			// Computed
			"admin_username": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"instances": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"availability_domain": {
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
						"id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"lifecycle_details": {
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
						"time_updated": {
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

func createPsqlDbSystem(d *schema.ResourceData, m interface{}) error {
	sync := &PsqlDbSystemResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).PostgresqlClient()

	return tfresource.CreateResource(d, sync)
}

func readPsqlDbSystem(d *schema.ResourceData, m interface{}) error {
	sync := &PsqlDbSystemResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).PostgresqlClient()

	return tfresource.ReadResource(sync)
}

func updatePsqlDbSystem(d *schema.ResourceData, m interface{}) error {
	sync := &PsqlDbSystemResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).PostgresqlClient()

	return tfresource.UpdateResource(d, sync)
}

func deletePsqlDbSystem(d *schema.ResourceData, m interface{}) error {
	sync := &PsqlDbSystemResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).PostgresqlClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type PsqlDbSystemResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_psql.PostgresqlClient
	Res                    *oci_psql.DbSystem
	PatchResponse          *oci_psql.DbSystem
	DisableNotFoundRetries bool
}

func (s *PsqlDbSystemResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *PsqlDbSystemResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_psql.DbSystemLifecycleStateCreating),
	}
}

func (s *PsqlDbSystemResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_psql.DbSystemLifecycleStateActive),
		string(oci_psql.DbSystemLifecycleStateNeedsAttention),
	}
}

func (s *PsqlDbSystemResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_psql.DbSystemLifecycleStateDeleting),
	}
}

func (s *PsqlDbSystemResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_psql.DbSystemLifecycleStateDeleted),
	}
}

func (s *PsqlDbSystemResourceCrud) Create() error {
	request := oci_psql.CreateDbSystemRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if configId, ok := s.D.GetOkExists("config_id"); ok {
		tmp := configId.(string)
		request.ConfigId = &tmp
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

	if dbVersion, ok := s.D.GetOkExists("db_version"); ok {
		tmp := dbVersion.(string)
		request.DbVersion = &tmp
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

	if instanceCount, ok := s.D.GetOkExists("instance_count"); ok {
		tmp := instanceCount.(int)
		request.InstanceCount = &tmp
	}

	if instanceMemorySizeInGBs, ok := s.D.GetOkExists("instance_memory_size_in_gbs"); ok {
		tmp := instanceMemorySizeInGBs.(int)
		request.InstanceMemorySizeInGBs = &tmp
	}

	if instanceOcpuCount, ok := s.D.GetOkExists("instance_ocpu_count"); ok {
		tmp := instanceOcpuCount.(int)
		request.InstanceOcpuCount = &tmp
	}

	if instancesDetails, ok := s.D.GetOkExists("instances_details"); ok {
		interfaces := instancesDetails.([]interface{})
		tmp := make([]oci_psql.CreateDbInstanceDetails, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "instances_details", stateDataIndex)
			converted, err := s.mapToCreateDbInstanceDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("instances_details") {
			request.InstancesDetails = tmp
		}
	}

	if managementPolicy, ok := s.D.GetOkExists("management_policy"); ok {
		if tmpList := managementPolicy.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "management_policy", 0)
			tmp, err := s.mapToManagementPolicyDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.ManagementPolicy = &tmp
		}
	}

	if networkDetails, ok := s.D.GetOkExists("network_details"); ok {
		if tmpList := networkDetails.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "network_details", 0)
			tmp, err := s.mapToNetworkDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.NetworkDetails = &tmp
		}
	}

	if shape, ok := s.D.GetOkExists("shape"); ok {
		tmp := shape.(string)
		request.Shape = &tmp
	}

	if source, ok := s.D.GetOkExists("source"); ok {
		if tmpList := source.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "source", 0)
			tmp, err := s.mapToSourceDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.Source = tmp
		}
	}

	if storageDetails, ok := s.D.GetOkExists("storage_details"); ok {
		if tmpList := storageDetails.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "storage_details", 0)
			tmp, err := s.mapToStorageDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.StorageDetails = tmp
		}
	}

	if systemType, ok := s.D.GetOkExists("system_type"); ok {
		request.SystemType = oci_psql.DbSystemSystemTypeEnum(systemType.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "psql")

	response, err := s.Client.CreateDbSystem(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	var identifier *string
	identifier = response.Id
	if identifier != nil {
		s.D.SetId(*identifier)
	}
	return s.getDbSystemFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "psql"), oci_psql.ActionTypeCreated, s.D.Timeout(schema.TimeoutCreate))
}

func (s *PsqlDbSystemResourceCrud) Patch() error {
	request := oci_psql.PatchDbSystemRequest{}

	tmp := s.D.Id()
	request.DbSystemId = &tmp

	if patchOperations, ok := s.D.GetOkExists("patch_operations"); ok {

		interfaces := patchOperations.([]interface{})
		tmp := make([]oci_psql.PatchInstruction, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "patch_operations", stateDataIndex)
			converted, err := s.mapToPatchInstruction(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 {
			request.Items = tmp
		} else {
			return nil
		}
	}

	//Check if instance count has changed when applying patch operation
	if (len(request.Items) != 0 && !s.D.HasChange("instance_count")) || (len(request.Items) == 0 && s.D.HasChange("instance_count")) {
		return fmt.Errorf("please update the instance count of the dbSystem while doing a patch operation to add/remove replica")
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "psql")
	response, err := s.Client.PatchDbSystem(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getDbSystemFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "psql"), oci_psql.ActionTypeCreated, s.D.Timeout(schema.TimeoutCreate))
}

func (s *PsqlDbSystemResourceCrud) getDbSystemFromWorkRequest(workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_psql.ActionTypeEnum, timeout time.Duration) error {

	// Wait until it finishes
	dbSystemId, err := dbSystemWaitForWorkRequest(workId, "dbsystem",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.Client)

	if err != nil {
		return err
	}
	s.D.SetId(*dbSystemId)

	return s.Get()
}

func dbSystemWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
	startTime := time.Now()
	stopTime := startTime.Add(timeout)
	return func(response oci_common.OCIOperationResponse) bool {

		// Stop after timeout has elapsed
		if time.Now().After(stopTime) {
			return false
		}

		// Make sure we stop on default rules
		if tfresource.ShouldRetry(response, false, "psql", startTime) {
			return true
		}

		// Only stop if the time Finished is set
		if workRequestResponse, ok := response.Response.(oci_psql.GetWorkRequestResponse); ok {
			return workRequestResponse.TimeFinished == nil
		}
		return false
	}
}

func dbSystemWaitForWorkRequest(wId *string, entityType string, action oci_psql.ActionTypeEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_psql.PostgresqlClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "psql")
	retryPolicy.ShouldRetryOperation = dbSystemWorkRequestShouldRetryFunc(timeout)

	response := oci_psql.GetWorkRequestResponse{}
	stateConf := &resource.StateChangeConf{
		Pending: []string{
			string(oci_psql.OperationStatusInProgress),
			string(oci_psql.OperationStatusAccepted),
			string(oci_psql.OperationStatusCanceling),
		},
		Target: []string{
			string(oci_psql.OperationStatusSucceeded),
			string(oci_psql.OperationStatusFailed),
			string(oci_psql.OperationStatusCanceled),
		},
		Refresh: func() (interface{}, string, error) {
			var err error
			response, err = client.GetWorkRequest(context.Background(),
				oci_psql.GetWorkRequestRequest{
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
			identifier = res.Identifier
			break
		}
	}

	// The workrequest may have failed, check for errors if identifier is not found or work failed or got cancelled
	if identifier == nil || response.Status == oci_psql.OperationStatusFailed || response.Status == oci_psql.OperationStatusCanceled {
		return nil, getErrorFromPsqlDbSystemWorkRequest(client, wId, retryPolicy, entityType, action)
	}

	return identifier, nil
}

func getErrorFromPsqlDbSystemWorkRequest(client *oci_psql.PostgresqlClient, workId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_psql.ActionTypeEnum) error {
	response, err := client.ListWorkRequestErrors(context.Background(),
		oci_psql.ListWorkRequestErrorsRequest{
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

func (s *PsqlDbSystemResourceCrud) Get() error {
	request := oci_psql.GetDbSystemRequest{}

	tmp := s.D.Id()
	request.DbSystemId = &tmp

	if excludedFields, ok := s.D.GetOkExists("excluded_fields"); ok {
		interfaces := excludedFields.([]interface{})
		tmp := make([]oci_psql.GetDbSystemExcludedFieldsEnum, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(oci_psql.GetDbSystemExcludedFieldsEnum)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("excluded_fields") {
			request.ExcludedFields = tmp
		}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "psql")

	response, err := s.Client.GetDbSystem(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.DbSystem
	return nil
}

func (s *PsqlDbSystemResourceCrud) Update() error {

	if _, ok := s.D.GetOkExists("passwordDetails"); ok && s.D.HasChange("passwordDetails") {
		err := s.ResetMasterUserPassword()
		if err != nil {
			return err
		}
	}
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(compartment)
			if err != nil {
				return err
			}
		}
	}
	request := oci_psql.UpdateDbSystemRequest{}

	if configId, ok := s.D.GetOkExists("config_id"); ok && s.D.HasChange("config_id") {
		configRequest := oci_psql.UpdateDbConfigParams{}

		tmp := configId.(string)
		configRequest.ConfigId = &tmp

		if applyConfig, ok := s.D.GetOkExists("apply_config"); ok {
			configRequest.ApplyConfig = oci_psql.UpdateDbConfigParamsApplyConfigEnum(applyConfig.(string))
		} else {
			configRequest.ApplyConfig = oci_psql.UpdateDbConfigParamsApplyConfigReload
		}

		request.DbConfigurationParams = &configRequest
	}

	tmp := s.D.Id()
	request.DbSystemId = &tmp

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

	if managementPolicy, ok := s.D.GetOkExists("management_policy"); ok {
		if tmpList := managementPolicy.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "management_policy", 0)
			tmp, err := s.mapToManagementPolicyDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.ManagementPolicy = &tmp
		}
	}

	if storageDetails, ok := s.D.GetOkExists("storage_details"); ok {
		if tmpList := storageDetails.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "storage_details", 0)
			tmp, err := s.mapToUpdateStorageDetailsParams(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.StorageDetails = &tmp
		}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "psql")

	response, err := s.Client.UpdateDbSystem(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	err = s.getDbSystemFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "psql"), oci_psql.ActionTypeInProgress, s.D.Timeout(schema.TimeoutUpdate))
	if err != nil {
		return err
	}
	err = s.Patch()
	if err != nil {
		log.Printf("[ERROR] Failed to execute Patch operation: %v", err)
		return err
	}
	return nil
}

func (s *PsqlDbSystemResourceCrud) Delete() error {
	request := oci_psql.DeleteDbSystemRequest{}

	tmp := s.D.Id()
	request.DbSystemId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "psql")

	response, err := s.Client.DeleteDbSystem(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	// Wait until it finishes
	_, delWorkRequestErr := dbSystemWaitForWorkRequest(workId, "dbsystem",
		oci_psql.ActionTypeDeleted, s.D.Timeout(schema.TimeoutDelete), s.DisableNotFoundRetries, s.Client)
	return delWorkRequestErr
}

func (s *PsqlDbSystemResourceCrud) SetData() error {
	if s.Res.AdminUsername != nil {
		s.D.Set("admin_username", *s.Res.AdminUsername)
	}

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.ConfigId != nil {
		s.D.Set("config_id", *s.Res.ConfigId)
	}

	if s.Res.DbVersion != nil {
		parts := strings.Split(*s.Res.DbVersion, ".")
		s.D.Set("db_version", parts[0])
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

	if s.Res.InstanceCount != nil {
		s.D.Set("instance_count", *s.Res.InstanceCount)
	}

	if s.Res.InstanceMemorySizeInGBs != nil {
		s.D.Set("instance_memory_size_in_gbs", *s.Res.InstanceMemorySizeInGBs)
	}

	if s.Res.InstanceOcpuCount != nil {
		s.D.Set("instance_ocpu_count", *s.Res.InstanceOcpuCount)
	}

	instances := []interface{}{}
	for _, item := range s.Res.Instances {
		instances = append(instances, DbInstanceToMap(item))
	}
	s.D.Set("instances", instances)

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	if s.Res.ManagementPolicy != nil {
		s.D.Set("management_policy", []interface{}{ManagementPolicyToMap(s.Res.ManagementPolicy)})
	} else {
		s.D.Set("management_policy", nil)
	}

	if s.Res.NetworkDetails != nil {
		s.D.Set("network_details", []interface{}{NetworkDetailsToMap(s.Res.NetworkDetails, false)})
	} else {
		s.D.Set("network_details", nil)
	}

	if s.Res.Shape != nil {
		s.D.Set("shape", "PostgreSQL."+*s.Res.Shape+"."+strconv.Itoa(*s.Res.InstanceOcpuCount)+"."+strconv.Itoa(*s.Res.InstanceMemorySizeInGBs)+"GB")
	}

	if s.Res.Source != nil {
		sourceArray := []interface{}{}
		if sourceMap := SourceDetailsToMap(&s.Res.Source); sourceMap != nil {
			sourceArray = append(sourceArray, sourceMap)
		}
		s.D.Set("source", sourceArray)
	} else {
		s.D.Set("source", nil)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.StorageDetails != nil {
		storageDetailsArray := []interface{}{}
		if storageDetailsMap := StorageDetailsToMap(&s.Res.StorageDetails); storageDetailsMap != nil {
			storageDetailsArray = append(storageDetailsArray, storageDetailsMap)
		}
		s.D.Set("storage_details", storageDetailsArray)
	} else {
		s.D.Set("storage_details", nil)
	}

	if s.Res.SystemTags != nil {
		s.D.Set("system_tags", tfresource.SystemTagsToMap(s.Res.SystemTags))
	}

	s.D.Set("system_type", s.Res.SystemType)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	return nil
}

func (s *PsqlDbSystemResourceCrud) ResetMasterUserPassword() error {
	request := oci_psql.ResetMasterUserPasswordRequest{}

	idTmp := s.D.Id()
	request.DbSystemId = &idTmp

	if passwordDetails, ok := s.D.GetOkExists("password_details"); ok {
		if tmpList := passwordDetails.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "password_details", 0)
			tmp, err := s.mapToPasswordDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.PasswordDetails = tmp
		}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "psql")

	_, err := s.Client.ResetMasterUserPassword(context.Background(), request)
	if err != nil {
		return err
	}

	if waitErr := tfresource.WaitForUpdatedState(s.D, s); waitErr != nil {
		return waitErr
	}

	return nil
}

func (s *PsqlDbSystemResourceCrud) mapToBackupPolicy(fieldKeyFormat string) (oci_psql.BackupPolicy, error) {
	var baseObject oci_psql.BackupPolicy
	//discriminator
	kindRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "kind"))
	var kind string
	if ok {
		kind = kindRaw.(string)
	} else {
		kind = "" // default value
	}
	switch strings.ToLower(kind) {
	case strings.ToLower("DAILY"):
		details := oci_psql.DailyBackupPolicy{}
		if backupStart, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "backup_start")); ok {
			tmp := strings.TrimSuffix(backupStart.(string), " UTC")
			details.BackupStart = &tmp
		}
		if retentionDays, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "retention_days")); ok {
			tmp := retentionDays.(int)
			details.RetentionDays = &tmp
		}
		baseObject = details
	case strings.ToLower("MONTHLY"):
		details := oci_psql.MonthlyBackupPolicy{}
		if backupStart, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "backup_start")); ok {
			tmp := strings.TrimSuffix(backupStart.(string), " UTC")
			details.BackupStart = &tmp
		}
		if daysOfTheMonth, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "days_of_the_month")); ok {
			interfaces := daysOfTheMonth.([]interface{})
			tmp := make([]int, len(interfaces))
			for i := range interfaces {
				if interfaces[i] != nil {
					tmp[i] = interfaces[i].(int)
				}
			}
			if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "days_of_the_month")) {
				details.DaysOfTheMonth = tmp
			}
		}
		if retentionDays, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "retention_days")); ok {
			tmp := retentionDays.(int)
			details.RetentionDays = &tmp
		}
		baseObject = details
	case strings.ToLower("NONE"):
		details := oci_psql.NoneBackupPolicy{}
		if retentionDays, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "retention_days")); ok {
			tmp := retentionDays.(int)
			details.RetentionDays = &tmp
		}
		baseObject = details
	case strings.ToLower("WEEKLY"):
		details := oci_psql.WeeklyBackupPolicy{}
		if backupStart, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "backup_start")); ok {
			tmp := strings.TrimSuffix(backupStart.(string), " UTC")
			details.BackupStart = &tmp
		}
		if daysOfTheWeek, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "days_of_the_week")); ok {
			interfaces := daysOfTheWeek.([]interface{})
			tmp := make([]oci_psql.WeeklyBackupPolicyDaysOfTheWeekEnum, len(interfaces))
			for i := range interfaces {
				if interfaces[i] != nil {
					str := interfaces[i].(string)

					tmp[i], _ = oci_psql.GetMappingWeeklyBackupPolicyDaysOfTheWeekEnum(str)
				}
			}
			if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "days_of_the_week")) {
				details.DaysOfTheWeek = tmp
			}
		}
		if retentionDays, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "retention_days")); ok {
			tmp := retentionDays.(int)
			details.RetentionDays = &tmp
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown kind '%v' was specified", kind)
	}
	return baseObject, nil
}

func BackupPolicyToMap(obj *oci_psql.BackupPolicy) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (*obj).(type) {
	case oci_psql.DailyBackupPolicy:
		result["kind"] = "DAILY"

		if v.BackupStart != nil {
			result["backup_start"] = strings.TrimSuffix(string(*v.BackupStart), " UTC")
		}

		if v.RetentionDays != nil {
			result["retention_days"] = int(*v.RetentionDays)
		}
	case oci_psql.MonthlyBackupPolicy:
		result["kind"] = "MONTHLY"

		if v.BackupStart != nil {
			result["backup_start"] = strings.TrimSuffix(string(*v.BackupStart), " UTC")
		}

		result["days_of_the_month"] = v.DaysOfTheMonth
		result["days_of_the_month"] = v.DaysOfTheMonth

		if v.RetentionDays != nil {
			result["retention_days"] = int(*v.RetentionDays)
		}
	case oci_psql.NoneBackupPolicy:
		result["kind"] = "NONE"

		if v.RetentionDays != nil {
			result["retention_days"] = int(*v.RetentionDays)
		}
	case oci_psql.WeeklyBackupPolicy:
		result["kind"] = "WEEKLY"

		if v.BackupStart != nil {
			result["backup_start"] = strings.TrimSuffix(string(*v.BackupStart), " UTC")
		}

		result["days_of_the_week"] = v.DaysOfTheWeek
		result["days_of_the_week"] = v.DaysOfTheWeek

		if v.RetentionDays != nil {
			result["retention_days"] = int(*v.RetentionDays)
		}
	default:
		log.Printf("[WARN] Received 'kind' of unknown type %v", *obj)
		return nil
	}

	return result
}

func (s *PsqlDbSystemResourceCrud) mapToCreateDbInstanceDetails(fieldKeyFormat string) (oci_psql.CreateDbInstanceDetails, error) {
	result := oci_psql.CreateDbInstanceDetails{}

	if description, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "description")); ok {
		tmp := description.(string)
		result.Description = &tmp
	}

	if displayName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "display_name")); ok {
		tmp := displayName.(string)
		result.DisplayName = &tmp
	}

	if privateIp, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "private_ip")); ok {
		tmp := privateIp.(string)
		result.PrivateIp = &tmp
	}

	return result, nil
}

func CreateDbInstanceDetailsToMap(obj oci_psql.CreateDbInstanceDetails) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Description != nil {
		result["description"] = string(*obj.Description)
	}

	if obj.DisplayName != nil {
		result["display_name"] = string(*obj.DisplayName)
	}

	if obj.PrivateIp != nil {
		result["private_ip"] = string(*obj.PrivateIp)
	}

	return result
}

func (s *PsqlDbSystemResourceCrud) mapToCredentials(fieldKeyFormat string) (oci_psql.Credentials, error) {
	result := oci_psql.Credentials{}

	if passwordDetails, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "password_details")); ok {
		if tmpList := passwordDetails.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "password_details"), 0)
			tmp, err := s.mapToPasswordDetails(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert password_details, encountered error: %v", err)
			}
			result.PasswordDetails = tmp
		}
	}

	if username, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "username")); ok {
		tmp := username.(string)
		result.Username = &tmp
	}

	return result, nil
}

func CredentialsToMap(obj *oci_psql.Credentials) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.PasswordDetails != nil {
		passwordDetailsArray := []interface{}{}
		if passwordDetailsMap := PasswordDetailsToMap(&obj.PasswordDetails); passwordDetailsMap != nil {
			passwordDetailsArray = append(passwordDetailsArray, passwordDetailsMap)
		}
		result["password_details"] = passwordDetailsArray
	}

	if obj.Username != nil {
		result["username"] = string(*obj.Username)
	}

	return result
}

func DbInstanceToMap(obj oci_psql.DbInstance) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.AvailabilityDomain != nil {
		result["availability_domain"] = string(*obj.AvailabilityDomain)
	}

	if obj.Description != nil {
		result["description"] = string(*obj.Description)
	}

	if obj.DisplayName != nil {
		result["display_name"] = string(*obj.DisplayName)
	}

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	if obj.LifecycleDetails != nil {
		result["lifecycle_details"] = string(*obj.LifecycleDetails)
	}

	result["state"] = string(obj.LifecycleState)

	if obj.TimeCreated != nil {
		result["time_created"] = obj.TimeCreated.String()
	}

	if obj.TimeUpdated != nil {
		result["time_updated"] = obj.TimeUpdated.String()
	}

	return result
}

func DbSystemSummaryToMap(obj oci_psql.DbSystemSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.CompartmentId != nil {
		result["compartment_id"] = string(*obj.CompartmentId)
	}

	if obj.ConfigId != nil {
		result["config_id"] = string(*obj.ConfigId)
	}

	if obj.DbVersion != nil {
		result["db_version"] = string(*obj.DbVersion)
	}

	if obj.DefinedTags != nil {
		result["defined_tags"] = tfresource.DefinedTagsToMap(obj.DefinedTags)
	}

	if obj.DisplayName != nil {
		result["display_name"] = string(*obj.DisplayName)
	}

	result["freeform_tags"] = obj.FreeformTags
	result["freeform_tags"] = obj.FreeformTags

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	if obj.InstanceCount != nil {
		result["instance_count"] = int(*obj.InstanceCount)
	}

	if obj.InstanceMemorySizeInGBs != nil {
		result["instance_memory_size_in_gbs"] = int(*obj.InstanceMemorySizeInGBs)
	}

	if obj.InstanceOcpuCount != nil {
		result["instance_ocpu_count"] = int(*obj.InstanceOcpuCount)
	}

	if obj.LifecycleDetails != nil {
		result["lifecycle_details"] = string(*obj.LifecycleDetails)
	}

	if obj.Shape != nil {
		result["shape"] = string(*obj.Shape)
	}

	result["state"] = string(obj.LifecycleState)

	if obj.SystemTags != nil {
		result["system_tags"] = tfresource.SystemTagsToMap(obj.SystemTags)
	}

	result["system_type"] = string(obj.SystemType)

	if obj.TimeCreated != nil {
		result["time_created"] = obj.TimeCreated.String()
	}

	if obj.TimeUpdated != nil {
		result["time_updated"] = obj.TimeUpdated.String()
	}

	return result
}

func (s *PsqlDbSystemResourceCrud) mapToManagementPolicyDetails(fieldKeyFormat string) (oci_psql.ManagementPolicyDetails, error) {
	result := oci_psql.ManagementPolicyDetails{}

	if backupPolicy, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "backup_policy")); ok {
		if tmpList := backupPolicy.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "backup_policy"), 0)
			tmp, err := s.mapToBackupPolicy(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert backup_policy, encountered error: %v", err)
			}
			result.BackupPolicy = tmp
		}
	}

	if maintenanceWindowStart, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "maintenance_window_start")); ok {
		tmp := maintenanceWindowStart.(string)
		result.MaintenanceWindowStart = &tmp
	}

	return result, nil
}

func ManagementPolicyToMap(obj *oci_psql.ManagementPolicy) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.BackupPolicy != nil {
		backupPolicyArray := []interface{}{}
		if backupPolicyMap := BackupPolicyToMap(&obj.BackupPolicy); backupPolicyMap != nil {
			backupPolicyArray = append(backupPolicyArray, backupPolicyMap)
		}
		result["backup_policy"] = backupPolicyArray
	}

	if obj.MaintenanceWindowStart != nil {
		result["maintenance_window_start"] = string(*obj.MaintenanceWindowStart)
	}

	return result
}

func (s *PsqlDbSystemResourceCrud) mapToNetworkDetails(fieldKeyFormat string) (oci_psql.NetworkDetails, error) {
	result := oci_psql.NetworkDetails{}

	if nsgIds, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "nsg_ids")); ok {
		set := nsgIds.(*schema.Set)
		interfaces := set.List()
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "nsg_ids")) {
			result.NsgIds = tmp
		}
	}

	if primaryDbEndpointPrivateIp, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "primary_db_endpoint_private_ip")); ok {
		tmp := primaryDbEndpointPrivateIp.(string)
		result.PrimaryDbEndpointPrivateIp = &tmp
	}

	if subnetId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "subnet_id")); ok {
		tmp := subnetId.(string)
		result.SubnetId = &tmp
	}

	return result, nil
}

func NetworkDetailsToMap(obj *oci_psql.NetworkDetails, datasource bool) map[string]interface{} {
	result := map[string]interface{}{}

	nsgIds := []interface{}{}
	for _, item := range obj.NsgIds {
		nsgIds = append(nsgIds, item)
	}
	if datasource {
		result["nsg_ids"] = nsgIds
	} else {
		result["nsg_ids"] = schema.NewSet(tfresource.LiteralTypeHashCodeForSets, nsgIds)
	}

	if obj.PrimaryDbEndpointPrivateIp != nil {
		result["primary_db_endpoint_private_ip"] = string(*obj.PrimaryDbEndpointPrivateIp)
	}

	if obj.SubnetId != nil {
		result["subnet_id"] = string(*obj.SubnetId)
	}

	return result
}

func (s *PsqlDbSystemResourceCrud) mapToPasswordDetails(fieldKeyFormat string) (oci_psql.PasswordDetails, error) {
	var baseObject oci_psql.PasswordDetails
	//discriminator
	passwordTypeRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "password_type"))
	var passwordType string
	if ok {
		passwordType = passwordTypeRaw.(string)
	} else {
		passwordType = "" // default value
	}
	switch strings.ToLower(passwordType) {
	case strings.ToLower("PLAIN_TEXT"):
		details := oci_psql.PlainTextPasswordDetails{}
		if password, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "password")); ok {
			tmp := password.(string)
			details.Password = &tmp
		}
		baseObject = details
	case strings.ToLower("VAULT_SECRET"):
		details := oci_psql.VaultSecretPasswordDetails{}
		if secretId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "secret_id")); ok {
			tmp := secretId.(string)
			details.SecretId = &tmp
		}
		if secretVersion, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "secret_version")); ok {
			tmp := secretVersion.(string)
			details.SecretVersion = &tmp
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown password_type '%v' was specified", passwordType)
	}
	return baseObject, nil
}

func PasswordDetailsToMap(obj *oci_psql.PasswordDetails) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (*obj).(type) {
	case oci_psql.PlainTextPasswordDetails:
		result["password_type"] = "PLAIN_TEXT"

		if v.Password != nil {
			result["password"] = string(*v.Password)
		}
	case oci_psql.VaultSecretPasswordDetails:
		result["password_type"] = "VAULT_SECRET"

		if v.SecretId != nil {
			result["secret_id"] = string(*v.SecretId)
		}

		if v.SecretVersion != nil {
			result["secret_version"] = string(*v.SecretVersion)
		}
	default:
		log.Printf("[WARN] Received 'password_type' of unknown type %v", *obj)
		return nil
	}

	return result
}

func (s *PsqlDbSystemResourceCrud) mapToPatchInstruction(fieldKeyFormat string) (oci_psql.PatchInstruction, error) {
	var baseObject oci_psql.PatchInstruction
	//discriminator
	operationRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "operation"))
	var operation string
	if ok {
		operation = operationRaw.(string)
	} else {
		operation = "" // default value
	}
	switch strings.ToLower(operation) {
	case strings.ToLower("INSERT"):
		details := oci_psql.PatchInsertInstruction{}
		if value, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "value")); ok {
			details.Value = &value
		}
		if selection, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "selection")); ok {
			tmp := selection.(string)
			details.Selection = &tmp
		}
		baseObject = details
	case strings.ToLower("REMOVE"):
		details := oci_psql.PatchRemoveInstruction{}
		if selection, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "selection")); ok {
			tmp := selection.(string)
			details.Selection = &tmp
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown operation '%v' was specified", operation)
	}
	return baseObject, nil
}

func (s *PsqlDbSystemResourceCrud) mapToSourceDetails(fieldKeyFormat string) (oci_psql.SourceDetails, error) {
	var baseObject oci_psql.SourceDetails
	//discriminator
	sourceTypeRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "source_type"))
	var sourceType string
	if ok {
		sourceType = sourceTypeRaw.(string)
	} else {
		sourceType = "" // default value
	}
	switch strings.ToLower(sourceType) {
	case strings.ToLower("BACKUP"):
		details := oci_psql.BackupSourceDetails{}
		if backupId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "backup_id")); ok {
			tmp := backupId.(string)
			details.BackupId = &tmp
		}
		if isHavingRestoreConfigOverrides, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_having_restore_config_overrides")); ok {
			tmp := isHavingRestoreConfigOverrides.(bool)
			details.IsHavingRestoreConfigOverrides = &tmp
		}
		baseObject = details
	case strings.ToLower("NONE"):
		details := oci_psql.NoneSourceDetails{}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown source_type '%v' was specified", sourceType)
	}
	return baseObject, nil
}

func SourceDetailsToMap(obj *oci_psql.SourceDetails) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (*obj).(type) {
	case oci_psql.BackupSourceDetails:
		result["source_type"] = "BACKUP"

		if v.BackupId != nil {
			result["backup_id"] = string(*v.BackupId)
		}

		if v.IsHavingRestoreConfigOverrides != nil {
			result["is_having_restore_config_overrides"] = bool(*v.IsHavingRestoreConfigOverrides)
		}
	case oci_psql.NoneSourceDetails:
		result["source_type"] = "NONE"
	default:
		log.Printf("[WARN] Received 'source_type' of unknown type %v", *obj)
		return nil
	}

	return result
}

func (s *PsqlDbSystemResourceCrud) mapToStorageDetails(fieldKeyFormat string) (oci_psql.StorageDetails, error) {
	var baseObject oci_psql.StorageDetails
	//discriminator
	systemTypeRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "system_type"))
	var systemType string
	if ok {
		systemType = systemTypeRaw.(string)
	} else {
		systemType = "" // default value
	}
	switch strings.ToLower(systemType) {
	case strings.ToLower("OCI_OPTIMIZED_STORAGE"):
		details := oci_psql.OciOptimizedStorageDetails{}
		if iops, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "iops")); ok {
			tmp := iops.(string)
			tmpInt64, err := strconv.ParseInt(tmp, 10, 64)
			if err != nil {
				return details, fmt.Errorf("unable to convert iops string: %s to an int64 and encountered error: %v", tmp, err)
			}
			details.Iops = &tmpInt64
		}
		if availabilityDomain, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "availability_domain")); ok {
			tmp := availabilityDomain.(string)
			details.AvailabilityDomain = &tmp
		}
		if isRegionallyDurable, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_regionally_durable")); ok {
			tmp := isRegionallyDurable.(bool)
			details.IsRegionallyDurable = &tmp
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown system_type '%v' was specified", systemType)
	}
	return baseObject, nil
}

func (s *PsqlDbSystemResourceCrud) mapToUpdateStorageDetailsParams(fieldKeyFormat string) (oci_psql.UpdateStorageDetailsParams, error) {
	var baseObject oci_psql.UpdateStorageDetailsParams
	//discriminator
	systemTypeRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "system_type"))
	var systemType string
	if ok {
		systemType = systemTypeRaw.(string)
	} else {
		systemType = "" // default value
	}
	switch strings.ToLower(systemType) {
	case strings.ToLower("OCI_OPTIMIZED_STORAGE"):
		details := oci_psql.OciOptimizedStorageDetails{}
		if iops, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "iops")); ok {
			tmp := iops.(string)
			tmpInt64, err := strconv.ParseInt(tmp, 10, 64)
			if err != nil {
				return baseObject, fmt.Errorf("unable to convert iops string: %s to an int64 and encountered error: %v", tmp, err)
			}
			details.Iops = &tmpInt64
		}
		baseObject = oci_psql.UpdateStorageDetailsParams{Iops: details.Iops}
	default:
		return baseObject, fmt.Errorf("unknown system_type '%v' was specified", systemType)
	}
	return baseObject, nil
}

func StorageDetailsToMap(obj *oci_psql.StorageDetails) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (*obj).(type) {
	case oci_psql.OciOptimizedStorageDetails:
		result["system_type"] = "OCI_OPTIMIZED_STORAGE"

		if v.Iops != nil {
			result["iops"] = strconv.FormatInt(*v.Iops, 10)
		}

		if v.AvailabilityDomain != nil {
			result["availability_domain"] = string(*v.AvailabilityDomain)
		}

		if v.IsRegionallyDurable != nil {
			result["is_regionally_durable"] = bool(*v.IsRegionallyDurable)
		}
	default:
		log.Printf("[WARN] Received 'system_type' of unknown type %v", *obj)
		return nil
	}

	return result
}

func (s *PsqlDbSystemResourceCrud) mapToUpdateDbConfigParams(fieldKeyFormat string) (oci_psql.UpdateDbConfigParams, error) {
	result := oci_psql.UpdateDbConfigParams{}

	if applyConfig, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "apply_config")); ok {
		result.ApplyConfig = oci_psql.UpdateDbConfigParamsApplyConfigEnum(applyConfig.(string))
	}

	if configId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "config_id")); ok {
		tmp := configId.(string)
		result.ConfigId = &tmp
	}

	return result, nil
}

func UpdateDbConfigParamsToMap(obj *oci_psql.UpdateDbConfigParams) map[string]interface{} {
	result := map[string]interface{}{}

	result["apply_config"] = string(obj.ApplyConfig)

	if obj.ConfigId != nil {
		result["config_id"] = string(*obj.ConfigId)
	}

	return result
}

func (s *PsqlDbSystemResourceCrud) updateCompartment(compartment interface{}) error {
	changeCompartmentRequest := oci_psql.ChangeDbSystemCompartmentRequest{}

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	idTmp := s.D.Id()
	changeCompartmentRequest.DbSystemId = &idTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "psql")

	_, err := s.Client.ChangeDbSystemCompartment(context.Background(), changeCompartmentRequest)
	if err != nil {
		return err
	}

	return nil
}
