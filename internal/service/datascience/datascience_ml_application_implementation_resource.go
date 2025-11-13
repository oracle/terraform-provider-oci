// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package datascience

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/retry"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	oci_common "github.com/oracle/oci-go-sdk/v65/common"
	oci_datascience "github.com/oracle/oci-go-sdk/v65/datascience"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DatascienceMlApplicationImplementationResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts:      tfresource.DefaultTimeout,
		Create:        createDatascienceMlApplicationImplementation,
		Read:          readDatascienceMlApplicationImplementation,
		Update:        updateDatascienceMlApplicationImplementation,
		Delete:        deleteDatascienceMlApplicationImplementation,
		CustomizeDiff: tfresource.SetPackagePath,
		Schema: map[string]*schema.Schema{
			// Required
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"ml_application_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			// Optional
			"ml_application_package": {
				Type:     schema.TypeMap,
				Optional: true,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				ValidateFunc:     tfresource.ValidateMLApplicationPackage,
				DiffSuppressFunc: tfresource.SupressPackageUpload,
				Description:      "Specifies the ML application package as a map of key-value pairs. Valid keys include 'source_type', 'path', and 'uri'. Use 'file://' for local paths or 'https://' for object storage URIs.",
			},

			"opc_ml_app_package_args": {
				Type:     schema.TypeMap,
				Optional: true,
				Elem:     schema.TypeString,
			},
			"allowed_migration_destinations": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
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
			"logging": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional
						"aggregated_instance_view_log": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							MaxItems: 1,
							MinItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional
									"enable_logging": {
										Type:     schema.TypeBool,
										Optional: true,
										Computed: true,
									},
									"log_group_id": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"log_id": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},

									// Computed
								},
							},
						},
						"implementation_log": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							MaxItems: 1,
							MinItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional
									"enable_logging": {
										Type:     schema.TypeBool,
										Optional: true,
										Computed: true,
									},
									"log_group_id": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"log_id": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},

									// Computed
								},
							},
						},
						"trigger_log": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							MaxItems: 1,
							MinItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional
									"enable_logging": {
										Type:     schema.TypeBool,
										Optional: true,
										Computed: true,
									},
									"log_group_id": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"log_id": {
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

			// Computed
			"application_components": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"application_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"component_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"job_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"model_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"pipeline_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"resource_type": {
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
			"configuration_schema": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"default_value": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"description": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"is_mandatory": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"key_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"sample_value": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"validation_regexp": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"value_type": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"description": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"lifecycle_details": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"ml_application_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"ml_application_package_arguments": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"arguments": {
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
									"is_mandatory": {
										Type:     schema.TypeBool,
										Computed: true,
									},
									"name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"type": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"value": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
					},
				},
			},
			"package_version": {
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

func createDatascienceMlApplicationImplementation(d *schema.ResourceData, m interface{}) error {
	sync := &DatascienceMlApplicationImplementationResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataScienceClient()

	var packageContent io.ReadCloser
	var contentLen int64
	var err error
	var sourceType string
	var sourceTypeExists bool

	raw, ok := d.GetOkExists("ml_application_package")
	if ok {
		packageDetails, okMapConversion := raw.(map[string]interface{})
		if !okMapConversion {
			return fmt.Errorf("ml_application_package must be a map")
		}
		sourceType, sourceTypeExists = packageDetails["source_type"].(string)
		if sourceType == "object_storage_download" {
			uri, _ := packageDetails["uri"].(string)
			if !strings.HasPrefix(uri, "https://") {
				return fmt.Errorf("invalid URI for object storage download source: must start with 'https://'")
			}

			configProvider := sync.Client.ConfigurationProvider()
			packageContent, contentLen, err = tfresource.GetPackage(configProvider, uri)
			if err != nil {
				return fmt.Errorf("the specified MLApplication package is not available: %q", err)
			}
		}
	}
	if e := tfresource.CreateResource(d, sync); e != nil {
		return e
	}
	if sourceTypeExists && sourceType != "" {
		if e := sync.putMlApplicationPackage(packageContent, contentLen); e != nil {
			return e
		}
	}
	return tfresource.ReadResource(sync)
}

func readDatascienceMlApplicationImplementation(d *schema.ResourceData, m interface{}) error {
	sync := &DatascienceMlApplicationImplementationResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataScienceClient()

	return tfresource.ReadResource(sync)
}

func updateDatascienceMlApplicationImplementation(d *schema.ResourceData, m interface{}) error {
	sync := &DatascienceMlApplicationImplementationResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataScienceClient()

	var packageContent io.ReadCloser
	var contentLen int64
	var err error
	var sourceType string
	var sourceTypeExists bool

	raw, ok := d.GetOkExists("ml_application_package")
	if ok {
		d.Partial(true)
		packageDetails, okMapConversion := raw.(map[string]interface{})
		if !okMapConversion {
			return fmt.Errorf("ml_application_package must be a map")
		}
		sourceType, sourceTypeExists = packageDetails["source_type"].(string)
		if sourceType == "object_storage_download" {
			uri, _ := packageDetails["uri"].(string)
			if !strings.HasPrefix(uri, "https://") {
				return fmt.Errorf("invalid URI for object storage download source: must start with 'https://'")
			}

			configProvider := sync.Client.ConfigurationProvider()
			packageContent, contentLen, err = tfresource.GetPackage(configProvider, uri)
			if err != nil {
				return fmt.Errorf("the specified MLApplication package is not available: %q", err)
			}
		}
		d.Partial(false)
	}
	if e := tfresource.UpdateResource(d, sync); e != nil {
		return e
	}
	if sourceTypeExists && sourceType != "" {
		if e := sync.putMlApplicationPackage(packageContent, contentLen); e != nil {
			return e
		}
	}

	return tfresource.ReadResource(sync)
}

func deleteDatascienceMlApplicationImplementation(d *schema.ResourceData, m interface{}) error {
	sync := &DatascienceMlApplicationImplementationResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataScienceClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type DatascienceMlApplicationImplementationResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_datascience.DataScienceClient
	Res                    *oci_datascience.MlApplicationImplementation
	DisableNotFoundRetries bool
}

func (s *DatascienceMlApplicationImplementationResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *DatascienceMlApplicationImplementationResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_datascience.MlApplicationImplementationLifecycleStateUpdating),
	}
}

func (s *DatascienceMlApplicationImplementationResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_datascience.MlApplicationImplementationLifecycleStateCreating),
		string(oci_datascience.MlApplicationImplementationLifecycleStateActive),
		string(oci_datascience.MlApplicationImplementationLifecycleStateNeedsAttention),
	}
}

func (s *DatascienceMlApplicationImplementationResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_datascience.MlApplicationImplementationLifecycleStateDeleting),
	}
}

func (s *DatascienceMlApplicationImplementationResourceCrud) DeletedTarget() []string {
	return []string{}
}

func (s *DatascienceMlApplicationImplementationResourceCrud) Create() error {
	request := oci_datascience.CreateMlApplicationImplementationRequest{}

	if allowedMigrationDestinations, ok := s.D.GetOkExists("allowed_migration_destinations"); ok {
		interfaces := allowedMigrationDestinations.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("allowed_migration_destinations") {
			request.AllowedMigrationDestinations = tmp
		}
	}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
		convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return err
		}
		request.DefinedTags = convertedDefinedTags
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if logging, ok := s.D.GetOkExists("logging"); ok {
		if tmpList := logging.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "logging", 0)
			tmp, err := s.mapToImplementationLogging(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.Logging = &tmp
		}
	}

	if mlApplicationId, ok := s.D.GetOkExists("ml_application_id"); ok {
		tmp := mlApplicationId.(string)
		request.MlApplicationId = &tmp
	}

	if name, ok := s.D.GetOkExists("name"); ok {
		tmp := name.(string)
		request.Name = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "datascience")

	response, err := s.Client.CreateMlApplicationImplementation(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.MlApplicationImplementation
	return nil
}

func (s *DatascienceMlApplicationImplementationResourceCrud) getMlApplicationImplementationFromWorkRequest(workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_datascience.WorkRequestResourceActionTypeEnum, timeout time.Duration) error {

	// Wait until it finishes
	mlApplicationImplementationId, err := mlApplicationImplementationWaitForWorkRequest(workId, "mlapplicationimplementation",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.Client)

	if err != nil {
		// Try to cancel the work request
		log.Printf("[DEBUG] creation failed, attempting to cancel the workrequest: %v for identifier: %v\n", workId, mlApplicationImplementationId)
		_, cancelErr := s.Client.CancelWorkRequest(context.Background(),
			oci_datascience.CancelWorkRequestRequest{
				WorkRequestId: workId,
				RequestMetadata: oci_common.RequestMetadata{
					RetryPolicy: retryPolicy,
				},
			})
		if cancelErr != nil {
			log.Printf("[DEBUG] cleanup cancelWorkRequest failed with the error: %v\n", cancelErr)
		}
		return err
	}
	s.D.SetId(*mlApplicationImplementationId)

	return s.Get()
}

func mlApplicationImplementationWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
	startTime := time.Now()
	stopTime := startTime.Add(timeout)
	return func(response oci_common.OCIOperationResponse) bool {

		// Stop after timeout has elapsed
		if time.Now().After(stopTime) {
			return false
		}

		// Make sure we stop on default rules
		if tfresource.ShouldRetry(response, false, "datascience", startTime) {
			return true
		}

		// Only stop if the time Finished is set
		if workRequestResponse, ok := response.Response.(oci_datascience.GetWorkRequestResponse); ok {
			return workRequestResponse.TimeFinished == nil
		}
		return false
	}
}

func mlApplicationImplementationWaitForWorkRequest(wId *string, entityType string, action oci_datascience.WorkRequestResourceActionTypeEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_datascience.DataScienceClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "datascience")
	retryPolicy.ShouldRetryOperation = mlApplicationImplementationWorkRequestShouldRetryFunc(timeout)

	response := oci_datascience.GetWorkRequestResponse{}
	stateConf := &retry.StateChangeConf{
		Pending: []string{
			string(oci_datascience.WorkRequestStatusInProgress),
			string(oci_datascience.WorkRequestStatusAccepted),
			string(oci_datascience.WorkRequestStatusCanceling),
		},
		Target: []string{
			string(oci_datascience.WorkRequestStatusSucceeded),
			string(oci_datascience.WorkRequestStatusFailed),
			string(oci_datascience.WorkRequestStatusCanceled),
		},
		Refresh: func() (interface{}, string, error) {
			var err error
			response, err = client.GetWorkRequest(context.Background(),
				oci_datascience.GetWorkRequestRequest{
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
			if res.ActionType == action {
				identifier = res.Identifier
				break
			}
		}
	}

	// The workrequest may have failed, check for errors if identifier is not found or work failed or got cancelled
	if identifier == nil || response.Status == oci_datascience.WorkRequestStatusFailed || response.Status == oci_datascience.WorkRequestStatusCanceled {
		return nil, getErrorFromDatascienceMlApplicationImplementationWorkRequest(client, wId, retryPolicy, entityType, action)
	}

	return identifier, nil
}

func getErrorFromDatascienceMlApplicationImplementationWorkRequest(client *oci_datascience.DataScienceClient, workId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_datascience.WorkRequestResourceActionTypeEnum) error {
	response, err := client.ListWorkRequestErrors(context.Background(),
		oci_datascience.ListWorkRequestErrorsRequest{
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

func (s *DatascienceMlApplicationImplementationResourceCrud) Get() error {
	request := oci_datascience.GetMlApplicationImplementationRequest{}

	tmp := s.D.Id()
	request.MlApplicationImplementationId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "datascience")

	response, err := s.Client.GetMlApplicationImplementation(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.MlApplicationImplementation
	return nil
}

func (s *DatascienceMlApplicationImplementationResourceCrud) Update() error {
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(compartment)
			if err != nil {
				return err
			}
		}
	}
	request := oci_datascience.UpdateMlApplicationImplementationRequest{}

	if allowedMigrationDestinations, ok := s.D.GetOkExists("allowed_migration_destinations"); ok {
		interfaces := allowedMigrationDestinations.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("allowed_migration_destinations") {
			request.AllowedMigrationDestinations = tmp
		}
	}

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
		convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return err
		}
		request.DefinedTags = convertedDefinedTags
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if logging, ok := s.D.GetOkExists("logging"); ok {
		if tmpList := logging.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "logging", 0)
			tmp, err := s.mapToImplementationLogging(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.Logging = &tmp
		}
	}

	tmp := s.D.Id()
	request.MlApplicationImplementationId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "datascience")

	response, err := s.Client.UpdateMlApplicationImplementation(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getMlApplicationImplementationFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "datascience"), oci_datascience.WorkRequestResourceActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}

func (s *DatascienceMlApplicationImplementationResourceCrud) Delete() error {
	request := oci_datascience.DeleteMlApplicationImplementationRequest{}

	tmp := s.D.Id()
	request.MlApplicationImplementationId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "datascience")

	response, err := s.Client.DeleteMlApplicationImplementation(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	// Wait until it finishes
	_, delWorkRequestErr := mlApplicationImplementationWaitForWorkRequest(workId, "mlapplicationimplementation",
		oci_datascience.WorkRequestResourceActionTypeDeleted, s.D.Timeout(schema.TimeoutDelete), s.DisableNotFoundRetries, s.Client)
	return delWorkRequestErr
}

func (s *DatascienceMlApplicationImplementationResourceCrud) SetData() error {
	s.D.Set("allowed_migration_destinations", s.Res.AllowedMigrationDestinations)

	applicationComponents := []interface{}{}
	for _, item := range s.Res.ApplicationComponents {
		applicationComponents = append(applicationComponents, ApplicationComponentToMap(item))
	}
	s.D.Set("application_components", applicationComponents)

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	configurationSchema := []interface{}{}
	for _, item := range s.Res.ConfigurationSchema {
		configurationSchema = append(configurationSchema, ConfigurationPropertySchemaToMap(item))
	}
	s.D.Set("configuration_schema", configurationSchema)

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.Description != nil {
		s.D.Set("description", *s.Res.Description)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	if s.Res.Logging != nil {
		s.D.Set("logging", []interface{}{ImplementationLoggingToMap(s.Res.Logging)})
	} else {
		s.D.Set("logging", nil)
	}

	if s.Res.MlApplicationId != nil {
		s.D.Set("ml_application_id", *s.Res.MlApplicationId)
	}

	if s.Res.MlApplicationName != nil {
		s.D.Set("ml_application_name", *s.Res.MlApplicationName)
	}

	if s.Res.MlApplicationPackageArguments != nil {
		s.D.Set("ml_application_package_arguments", []interface{}{MlApplicationPackageArgumentsToMap(s.Res.MlApplicationPackageArguments)})
	} else {
		s.D.Set("ml_application_package_arguments", nil)
	}

	if s.Res.Name != nil {
		s.D.Set("name", *s.Res.Name)
	}

	if s.Res.PackageVersion != nil {
		s.D.Set("package_version", *s.Res.PackageVersion)
	}

	packageDetails, ok := s.D.Get("ml_application_package").(map[string]interface{})
	if ok && packageDetails != nil {
		sourceType, sourceTypeExists := packageDetails["source_type"].(string)
		if !sourceTypeExists || sourceType == "" {
			if err := s.D.Set("ml_application_package", make(map[string]interface{})); err != nil {
				return fmt.Errorf("failed to initialise ml_application_package with empty map: %s", err)
			}
		}

		pathVersion, pathExists := packageDetails["path"].(string)
		if pathExists && pathVersion != "" && s.Res.PackageVersion != nil {
			split := strings.SplitN(pathVersion, " ", 2)
			split = strings.SplitN(split[0], "::", 2)
			if len(split) > 1 {
				path := split[0]
				version := split[1]
				if version != *s.Res.PackageVersion {
					updatedPathVersion := fmt.Sprintf(
						"%s::%s (remote cloud packageVersion changes detected: %s)",
						path, version, *s.Res.PackageVersion,
					)
					updatedPackageDetails := map[string]interface{}{
						"path":        updatedPathVersion,
						"source_type": packageDetails["source_type"], // Preserve existing source_type if present
					}
					if packageDetails["uri"] != nil {
						updatedPackageDetails["uri"] = packageDetails["uri"]
					}

					if err := s.D.Set("ml_application_package", updatedPackageDetails); err != nil {
						return fmt.Errorf("failed to update ml_application_package: %s", err)
					}
				}
			}
		}
	} else {
		if err := s.D.Set("ml_application_package", make(map[string]interface{})); err != nil {
			return fmt.Errorf("failed to initialise ml_application_package with empty map: %s", err)
		}
	}

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

func ApplicationComponentToMap(obj oci_datascience.ApplicationComponent) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (obj).(type) {
	case oci_datascience.DataFlowApplicationApplicationComponent:
		result["type"] = "DATA_FLOW_APPLICATION"

		if v.ApplicationId != nil {
			result["application_id"] = string(*v.ApplicationId)
		}

		if v.ComponentName != nil {
			result["component_name"] = string(*v.ComponentName)
		}

		if v.Name != nil {
			result["name"] = string(*v.Name)
		}
	case oci_datascience.DataScienceJobApplicationComponent:
		result["type"] = "DATA_SCIENCE_JOB"

		if v.JobId != nil {
			result["job_id"] = string(*v.JobId)
		}

		if v.ComponentName != nil {
			result["component_name"] = string(*v.ComponentName)
		}

		if v.Name != nil {
			result["name"] = string(*v.Name)
		}
	case oci_datascience.DataScienceModelApplicationComponent:
		result["type"] = "DATA_SCIENCE_MODEL"

		if v.ModelId != nil {
			result["model_id"] = string(*v.ModelId)
		}

		if v.ComponentName != nil {
			result["component_name"] = string(*v.ComponentName)
		}

		if v.Name != nil {
			result["name"] = string(*v.Name)
		}
	case oci_datascience.DataSciencePipelineApplicationComponent:
		result["type"] = "DATA_SCIENCE_PIPELINE"

		if v.PipelineId != nil {
			result["pipeline_id"] = string(*v.PipelineId)
		}

		if v.ComponentName != nil {
			result["component_name"] = string(*v.ComponentName)
		}

		if v.Name != nil {
			result["name"] = string(*v.Name)
		}
	case oci_datascience.GenericOciResourceApplicationComponent:
		result["type"] = "GENERIC_OCI_RESOURCE"

		if v.Id != nil {
			result["id"] = string(*v.Id)
		}

		if v.ResourceType != nil {
			result["resource_type"] = string(*v.ResourceType)
		}

		if v.ComponentName != nil {
			result["component_name"] = string(*v.ComponentName)
		}

		if v.Name != nil {
			result["name"] = string(*v.Name)
		}
	default:
		log.Printf("[WARN] Received 'type' of unknown type %v", obj)
		return nil
	}

	return result
}

func ConfigurationPropertySchemaToMap(obj oci_datascience.ConfigurationPropertySchema) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.DefaultValue != nil {
		result["default_value"] = string(*obj.DefaultValue)
	}

	if obj.Description != nil {
		result["description"] = string(*obj.Description)
	}

	if obj.IsMandatory != nil {
		result["is_mandatory"] = bool(*obj.IsMandatory)
	}

	if obj.KeyName != nil {
		result["key_name"] = string(*obj.KeyName)
	}

	if obj.SampleValue != nil {
		result["sample_value"] = string(*obj.SampleValue)
	}

	if obj.ValidationRegexp != nil {
		result["validation_regexp"] = string(*obj.ValidationRegexp)
	}

	result["value_type"] = string(obj.ValueType)

	return result
}

func (s *DatascienceMlApplicationImplementationResourceCrud) mapToImplementationLogDetails(fieldKeyFormat string) (oci_datascience.ImplementationLogDetails, error) {
	result := oci_datascience.ImplementationLogDetails{}

	if enableLogging, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "enable_logging")); ok {
		tmp := enableLogging.(bool)
		result.EnableLogging = &tmp
	}

	if logGroupId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "log_group_id")); ok {
		tmp := logGroupId.(string)
		result.LogGroupId = &tmp
	}

	if logId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "log_id")); ok {
		tmp := logId.(string)
		result.LogId = &tmp
	}

	return result, nil
}

func ImplementationLogDetailsToMap(obj *oci_datascience.ImplementationLogDetails) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.EnableLogging != nil {
		result["enable_logging"] = bool(*obj.EnableLogging)
	}

	if obj.LogGroupId != nil {
		result["log_group_id"] = string(*obj.LogGroupId)
	}

	if obj.LogId != nil {
		result["log_id"] = string(*obj.LogId)
	}

	return result
}

func (s *DatascienceMlApplicationImplementationResourceCrud) mapToImplementationLogging(fieldKeyFormat string) (oci_datascience.ImplementationLogging, error) {
	result := oci_datascience.ImplementationLogging{}

	if aggregatedInstanceViewLog, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "aggregated_instance_view_log")); ok {
		if tmpList := aggregatedInstanceViewLog.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "aggregated_instance_view_log"), 0)
			tmp, err := s.mapToImplementationLogDetails(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert aggregated_instance_view_log, encountered error: %v", err)
			}
			result.AggregatedInstanceViewLog = &tmp
		}
	}

	if implementationLog, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "implementation_log")); ok {
		if tmpList := implementationLog.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "implementation_log"), 0)
			tmp, err := s.mapToImplementationLogDetails(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert implementation_log, encountered error: %v", err)
			}
			result.ImplementationLog = &tmp
		}
	}

	if triggerLog, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "trigger_log")); ok {
		if tmpList := triggerLog.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "trigger_log"), 0)
			tmp, err := s.mapToImplementationLogDetails(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert trigger_log, encountered error: %v", err)
			}
			result.TriggerLog = &tmp
		}
	}

	return result, nil
}

func ImplementationLoggingToMap(obj *oci_datascience.ImplementationLogging) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.AggregatedInstanceViewLog != nil {
		result["aggregated_instance_view_log"] = []interface{}{ImplementationLogDetailsToMap(obj.AggregatedInstanceViewLog)}
	}

	if obj.ImplementationLog != nil {
		result["implementation_log"] = []interface{}{ImplementationLogDetailsToMap(obj.ImplementationLog)}
	}

	if obj.TriggerLog != nil {
		result["trigger_log"] = []interface{}{ImplementationLogDetailsToMap(obj.TriggerLog)}
	}

	return result
}

func MlApplicationImplementationSummaryToMap(obj oci_datascience.MlApplicationImplementationSummary) map[string]interface{} {
	result := map[string]interface{}{}

	result["allowed_migration_destinations"] = obj.AllowedMigrationDestinations

	if obj.CompartmentId != nil {
		result["compartment_id"] = string(*obj.CompartmentId)
	}

	configurationSchema := []interface{}{}
	for _, item := range obj.ConfigurationSchema {
		configurationSchema = append(configurationSchema, ConfigurationPropertySchemaToMap(item))
	}
	result["configuration_schema"] = configurationSchema

	if obj.DefinedTags != nil {
		result["defined_tags"] = tfresource.DefinedTagsToMap(obj.DefinedTags)
	}

	if obj.Description != nil {
		result["description"] = string(*obj.Description)
	}

	result["freeform_tags"] = obj.FreeformTags

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	if obj.MlApplicationId != nil {
		result["ml_application_id"] = string(*obj.MlApplicationId)
	}

	if obj.MlApplicationName != nil {
		result["ml_application_name"] = string(*obj.MlApplicationName)
	}

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	if obj.PackageVersion != nil {
		result["package_version"] = string(*obj.PackageVersion)
	}

	result["state"] = string(obj.LifecycleState)

	if obj.SystemTags != nil {
		result["system_tags"] = tfresource.SystemTagsToMap(obj.SystemTags)
	}

	if obj.TimeCreated != nil {
		result["time_created"] = obj.TimeCreated.String()
	}

	return result
}

func MlApplicationPackageArgumentDetailsToMap(obj oci_datascience.MlApplicationPackageArgumentDetails) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Description != nil {
		result["description"] = string(*obj.Description)
	}

	if obj.IsMandatory != nil {
		result["is_mandatory"] = bool(*obj.IsMandatory)
	}

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	result["type"] = string(obj.Type)

	if obj.Value != nil {
		result["value"] = string(*obj.Value)
	}

	return result
}

func MlApplicationPackageArgumentsToMap(obj *oci_datascience.MlApplicationPackageArguments) map[string]interface{} {
	result := map[string]interface{}{}

	arguments := []interface{}{}
	for _, item := range obj.Arguments {
		arguments = append(arguments, MlApplicationPackageArgumentDetailsToMap(item))
	}
	result["arguments"] = arguments

	return result
}

func (s *DatascienceMlApplicationImplementationResourceCrud) updateCompartment(compartment interface{}) error {
	changeCompartmentRequest := oci_datascience.ChangeMlApplicationImplementationCompartmentRequest{}

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	idTmp := s.D.Id()
	changeCompartmentRequest.MlApplicationImplementationId = &idTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "datascience")

	response, err := s.Client.ChangeMlApplicationImplementationCompartment(context.Background(), changeCompartmentRequest)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getMlApplicationImplementationFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "datascience"), oci_datascience.WorkRequestResourceActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}

func (s *DatascienceMlApplicationImplementationResourceCrud) putMlApplicationPackage(packageContent io.ReadCloser, contentLen int64) error {
	request := oci_datascience.PutMlApplicationPackageRequest{}

	request.MlApplicationImplementationId = s.Res.Id

	if mlApplicationPackage, ok := s.D.GetOkExists("ml_application_package"); ok {
		packageDetails := mlApplicationPackage.(map[string]interface{})
		sourceType, _ := packageDetails["source_type"].(string)

		switch sourceType {
		case "local":
			pathVersion, _ := packageDetails["path"].(string)
			if !strings.HasPrefix(pathVersion, "file://") {
				return fmt.Errorf("invalid path for local source: must start with 'file://'")
			}
			pathSplit := strings.SplitN(pathVersion, "::", 2)
			packagePath := strings.TrimPrefix(pathSplit[0], "file://")
			artifactReader, err := os.Open(packagePath)
			if err != nil {
				return fmt.Errorf("the specified MLApplication package file is not available: %q", err)
			}
			request.PutMlApplicationPackage = ioutil.NopCloser(artifactReader)

			pathWithoutSpace := strings.SplitN(pathVersion, " ", 2)
			packageDetails["path"] = pathWithoutSpace[0]
			s.D.Set("ml_application_package", packageDetails)

		case "object_storage_download":
			request.PutMlApplicationPackage = packageContent
			request.ContentLength = &contentLen

		case "object_storage":
			return fmt.Errorf("source_type 'object_storage' is not yet supported")

		default:
			return fmt.Errorf("unsupported source_type: %s", sourceType)
		}
	} else {
		return fmt.Errorf("ml_application_package is required")
	}

	if opcMlAppPackageArgs, ok := s.D.GetOkExists("opc_ml_app_package_args"); ok {
		jsonData, err := json.Marshal(opcMlAppPackageArgs.(map[string]interface{}))
		if err != nil {
			return fmt.Errorf("error converting map to string: %q", err)
		}
		jsonString := string(jsonData)
		request.OpcMlAppPackageArgs = &jsonString
	}

	if contentDisposition, ok := s.D.GetOkExists("content_disposition"); ok {
		tmp := contentDisposition.(string)
		request.ContentDisposition = &tmp
	}

	if contentLength, ok := s.D.GetOkExists("content_length"); ok {
		tmp := contentLength.(string)
		tmpInt64, err := strconv.ParseInt(tmp, 10, 64)
		if err != nil {
			return fmt.Errorf("unable to convert content-length string: %s to an int64 and encountered error: %v", tmp, err)
		}
		request.ContentLength = &tmpInt64
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "datascience")

	response, err := s.Client.PutMlApplicationPackage(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	workRequestResponse := oci_datascience.GetWorkRequestResponse{}
	workRequestResponse, err = s.Client.GetWorkRequest(context.Background(),
		oci_datascience.GetWorkRequestRequest{
			WorkRequestId: workId,
			RequestMetadata: oci_common.RequestMetadata{
				RetryPolicy: tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "datascience"),
			},
		})
	if err == nil {
		// The work request response contains an array of objects
		for _, res := range workRequestResponse.Resources {
			if res.EntityType != nil && strings.Contains(strings.ToLower(*res.EntityType), "mlapplicationimplementation") && res.Identifier != nil {
				s.D.SetId(*res.Identifier)
				break
			}
		}
	}
	return s.getMlApplicationImplementationMlApplicationPackageFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "datascience"), oci_datascience.WorkRequestResourceActionTypeUpdated, s.D.Timeout(schema.TimeoutCreate))
}

func (s *DatascienceMlApplicationImplementationResourceCrud) getMlApplicationImplementationMlApplicationPackageFromWorkRequest(workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_datascience.WorkRequestResourceActionTypeEnum, timeout time.Duration) error {

	// Wait until it finishes
	mlApplicationImplementationMlApplicationPackageId, err := mlApplicationImplementationMlApplicationPackageWaitForWorkRequest(workId, "mlapplicationimplementation",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.Client)

	if err != nil {
		// Try to cancel the work request
		log.Printf("[DEBUG] creation failed, attempting to cancel the workrequest: %v for identifier: %v\n", workId, mlApplicationImplementationMlApplicationPackageId)
		_, cancelErr := s.Client.CancelWorkRequest(context.Background(),
			oci_datascience.CancelWorkRequestRequest{
				WorkRequestId: workId,
				RequestMetadata: oci_common.RequestMetadata{
					RetryPolicy: retryPolicy,
				},
			})
		if cancelErr != nil {
			log.Printf("[DEBUG] cleanup cancelWorkRequest failed with the error: %v\n", cancelErr)
		}
		return err
	}
	s.D.SetId(*mlApplicationImplementationMlApplicationPackageId)

	return s.Get()
}

func mlApplicationImplementationMlApplicationPackageWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
	startTime := time.Now()
	stopTime := startTime.Add(timeout)
	return func(response oci_common.OCIOperationResponse) bool {

		// Stop after timeout has elapsed
		if time.Now().After(stopTime) {
			return false
		}

		// Make sure we stop on default rules
		if tfresource.ShouldRetry(response, false, "datascience", startTime) {
			return true
		}

		// Only stop if the time Finished is set
		if workRequestResponse, ok := response.Response.(oci_datascience.GetWorkRequestResponse); ok {
			return workRequestResponse.TimeFinished == nil
		}
		return false
	}
}

func mlApplicationImplementationMlApplicationPackageWaitForWorkRequest(wId *string, entityType string, action oci_datascience.WorkRequestResourceActionTypeEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_datascience.DataScienceClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "datascience")
	retryPolicy.ShouldRetryOperation = mlApplicationImplementationMlApplicationPackageWorkRequestShouldRetryFunc(timeout)

	response := oci_datascience.GetWorkRequestResponse{}
	stateConf := &retry.StateChangeConf{
		Pending: []string{
			string(oci_datascience.WorkRequestStatusInProgress),
			string(oci_datascience.WorkRequestStatusAccepted),
			string(oci_datascience.WorkRequestStatusCanceling),
		},
		Target: []string{
			string(oci_datascience.WorkRequestStatusSucceeded),
			string(oci_datascience.WorkRequestStatusFailed),
			string(oci_datascience.WorkRequestStatusCanceled),
		},
		Refresh: func() (interface{}, string, error) {
			var err error
			response, err = client.GetWorkRequest(context.Background(),
				oci_datascience.GetWorkRequestRequest{
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
	if identifier == nil || response.Status == oci_datascience.WorkRequestStatusFailed || response.Status == oci_datascience.WorkRequestStatusCanceled {
		return nil, getErrorFromDatascienceMlApplicationImplementationMlApplicationPackageWorkRequest(client, wId, retryPolicy, entityType, action)
	}

	return identifier, nil
}

func getErrorFromDatascienceMlApplicationImplementationMlApplicationPackageWorkRequest(client *oci_datascience.DataScienceClient, workId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_datascience.WorkRequestResourceActionTypeEnum) error {
	response, err := client.ListWorkRequestErrors(context.Background(),
		oci_datascience.ListWorkRequestErrorsRequest{
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
