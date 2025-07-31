// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package datascience

import (
	"context"
	"fmt"
	"strings"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"

	oci_datascience "github.com/oracle/oci-go-sdk/v65/datascience"
)

func DatascienceNotebookSessionResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createDatascienceNotebookSession,
		Read:     readDatascienceNotebookSession,
		Update:   updateDatascienceNotebookSession,
		Delete:   deleteDatascienceNotebookSession,
		Schema: map[string]*schema.Schema{
			// Required
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"project_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			// Optional
			"defined_tags": {
				Type:             schema.TypeMap,
				Optional:         true,
				Computed:         true,
				DiffSuppressFunc: tfresource.DefinedTagsDiffSuppressFunction,
				Elem:             schema.TypeString,
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
			"notebook_session_config_details": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				ForceNew: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"shape": {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},

						// Optional
						"block_storage_size_in_gbs": {
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},
						"notebook_session_shape_config_details": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							ForceNew: true,
							MaxItems: 1,
							MinItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional
									"cpu_baseline": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
										ForceNew: true,
									},
									"memory_in_gbs": {
										Type:     schema.TypeFloat,
										Optional: true,
										Computed: true,
										ForceNew: true,
									},
									"ocpus": {
										Type:     schema.TypeFloat,
										Optional: true,
										Computed: true,
										ForceNew: true,
									},

									// Computed
								},
							},
						},
						"private_endpoint_id": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},
						"subnet_id": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},

						// Computed
					},
				},
			},
			"notebook_session_configuration_details": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"shape": {
							Type:     schema.TypeString,
							Required: true,
						},
						"subnet_id": {
							Type:     schema.TypeString,
							Required: true,
						},

						// Optional
						"block_storage_size_in_gbs": {
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"notebook_session_shape_config_details": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							MaxItems: 1,
							MinItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional
									"cpu_baseline": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"memory_in_gbs": {
										Type:     schema.TypeFloat,
										Optional: true,
										Computed: true,
									},
									"ocpus": {
										Type:     schema.TypeFloat,
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

						// Computed
					},
				},
			},
			"notebook_session_runtime_config_details": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional
						"custom_environment_variables": {
							Type:     schema.TypeMap,
							Optional: true,
							Computed: true,
							Elem:     schema.TypeString,
						},
						"notebook_session_git_config_details": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							MaxItems: 1,
							MinItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional
									"notebook_session_git_repo_config_collection": {
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required
												"url": {
													Type:     schema.TypeString,
													Required: true,
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

						// Computed
					},
				},
			},
			"notebook_session_storage_mount_configuration_details_list": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"destination_directory_name": {
							Type:     schema.TypeString,
							Required: true,
						},
						"storage_type": {
							Type:             schema.TypeString,
							Required:         true,
							DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
							ValidateFunc: validation.StringInSlice([]string{
								"FILE_STORAGE",
								"OBJECT_STORAGE",
							}, true),
						},

						// Optional
						"bucket": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"destination_path": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"export_id": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"mount_target_id": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"namespace": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"prefix": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},

						// Computed
					},
				},
			},
			"state": {
				Type:             schema.TypeString,
				Optional:         true,
				Computed:         true,
				DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
				ValidateFunc: validation.StringInSlice([]string{
					string(oci_datascience.NotebookSessionLifecycleStateActive),
					string(oci_datascience.NotebookSessionLifecycleStateInactive),
				}, true),
			},

			// Computed
			"created_by": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"lifecycle_details": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"notebook_session_url": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_created": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func createDatascienceNotebookSession(d *schema.ResourceData, m interface{}) error {
	sync := &DatascienceNotebookSessionResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataScienceClient()

	var deactivateNotebookSession = false
	if state, ok := sync.D.GetOkExists("state"); ok {
		desiredState := oci_datascience.NotebookSessionLifecycleStateEnum(strings.ToUpper(state.(string)))
		if desiredState == oci_datascience.NotebookSessionLifecycleStateInactive {
			deactivateNotebookSession = true
		}
	}

	if e := tfresource.CreateResource(d, sync); e != nil {
		return e
	}
	if deactivateNotebookSession {
		if e := sync.DeactivateNotebookSession(); e != nil {
			return e
		}
		sync.D.Set("state", oci_datascience.NotebookSessionLifecycleStateInactive)
	}
	return nil
}

func readDatascienceNotebookSession(d *schema.ResourceData, m interface{}) error {
	sync := &DatascienceNotebookSessionResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataScienceClient()

	return tfresource.ReadResource(sync)
}

func updateDatascienceNotebookSession(d *schema.ResourceData, m interface{}) error {
	sync := &DatascienceNotebookSessionResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataScienceClient()

	// Activate/Deactivate NotebookSession
	activate, deactivate := false, false

	if sync.D.HasChange("state") {
		desiredState := strings.ToUpper(sync.D.Get("state").(string))
		if oci_datascience.NotebookSessionLifecycleStateActive == oci_datascience.NotebookSessionLifecycleStateEnum(desiredState) {
			activate = true
		} else if oci_datascience.NotebookSessionLifecycleStateInactive == oci_datascience.NotebookSessionLifecycleStateEnum(desiredState) {
			deactivate = true
		}
	} else {
		currentState := strings.ToUpper(sync.D.Get("state").(string))
		if oci_datascience.NotebookSessionLifecycleStateActive == oci_datascience.NotebookSessionLifecycleStateEnum(currentState) {
			activate = true
			deactivate = true
		}
	}

	if deactivate {
		if err := sync.DeactivateNotebookSession(); err != nil {
			return err
		}
		sync.D.Set("state", oci_datascience.NotebookSessionLifecycleStateInactive)
	}
	if err := tfresource.UpdateResource(d, sync); err != nil {
		return err
	}

	if activate {
		if err := sync.ActivateNotebookSession(); err != nil {
			return err
		}
		sync.D.Set("state", oci_datascience.NotebookSessionLifecycleStateActive)
	}
	return nil
}

func deleteDatascienceNotebookSession(d *schema.ResourceData, m interface{}) error {
	sync := &DatascienceNotebookSessionResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataScienceClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type DatascienceNotebookSessionResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_datascience.DataScienceClient
	Res                    *oci_datascience.NotebookSession
	DisableNotFoundRetries bool
}

func (s *DatascienceNotebookSessionResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *DatascienceNotebookSessionResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_datascience.NotebookSessionLifecycleStateCreating),
	}
}

func (s *DatascienceNotebookSessionResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_datascience.NotebookSessionLifecycleStateActive),
	}
}

func (s *DatascienceNotebookSessionResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_datascience.NotebookSessionLifecycleStateDeleting),
	}
}

func (s *DatascienceNotebookSessionResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_datascience.NotebookSessionLifecycleStateDeleted),
	}
}

func (s *DatascienceNotebookSessionResourceCrud) Create() error {
	request := oci_datascience.CreateNotebookSessionRequest{}

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

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if notebookSessionConfigDetails, ok := s.D.GetOkExists("notebook_session_config_details"); ok {
		if tmpList := notebookSessionConfigDetails.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "notebook_session_config_details", 0)
			tmp, err := s.mapToNotebookSessionConfigDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.NotebookSessionConfigDetails = &tmp
		}
	}

	if notebookSessionConfigurationDetails, ok := s.D.GetOkExists("notebook_session_configuration_details"); ok {
		if tmpList := notebookSessionConfigurationDetails.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "notebook_session_configuration_details", 0)
			tmp, err := s.mapToNotebookSessionConfigurationDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.NotebookSessionConfigurationDetails = &tmp
		}
	}

	if notebookSessionRuntimeConfigDetails, ok := s.D.GetOkExists("notebook_session_runtime_config_details"); ok {
		if tmpList := notebookSessionRuntimeConfigDetails.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "notebook_session_runtime_config_details", 0)
			tmp, err := s.mapToNotebookSessionRuntimeConfigDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.NotebookSessionRuntimeConfigDetails = &tmp
		}
	}

	if notebookSessionStorageMountConfigurationDetailsList, ok := s.D.GetOkExists("notebook_session_storage_mount_configuration_details_list"); ok {
		interfaces := notebookSessionStorageMountConfigurationDetailsList.([]interface{})
		tmp := make([]oci_datascience.StorageMountConfigurationDetails, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "notebook_session_storage_mount_configuration_details_list", stateDataIndex)
			converted, err := s.mapToStorageMountConfigurationDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("notebook_session_storage_mount_configuration_details_list") {
			request.NotebookSessionStorageMountConfigurationDetailsList = tmp
		}
	}

	if projectId, ok := s.D.GetOkExists("project_id"); ok {
		tmp := projectId.(string)
		request.ProjectId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "datascience")

	response, err := s.Client.CreateNotebookSession(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.NotebookSession
	return nil
}

func (s *DatascienceNotebookSessionResourceCrud) Get() error {
	request := oci_datascience.GetNotebookSessionRequest{}

	tmp := s.D.Id()
	request.NotebookSessionId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "datascience")

	response, err := s.Client.GetNotebookSession(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.NotebookSession
	return nil
}

func (s *DatascienceNotebookSessionResourceCrud) Update() error {
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(compartment)
			if err != nil {
				return err
			}
		}
	}
	request := oci_datascience.UpdateNotebookSessionRequest{}

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
		convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return err
		}
		request.DefinedTags = convertedDefinedTags
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if notebookSessionConfigurationDetails, ok := s.D.GetOkExists("notebook_session_configuration_details"); ok {
		if tmpList := notebookSessionConfigurationDetails.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "notebook_session_configuration_details", 0)
			tmp, err := s.mapToNotebookSessionConfigurationDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.NotebookSessionConfigurationDetails = &tmp
		}
	}

	tmp := s.D.Id()
	request.NotebookSessionId = &tmp

	if notebookSessionRuntimeConfigDetails, ok := s.D.GetOkExists("notebook_session_runtime_config_details"); ok {
		if tmpList := notebookSessionRuntimeConfigDetails.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "notebook_session_runtime_config_details", 0)
			tmp, err := s.mapToNotebookSessionRuntimeConfigDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.NotebookSessionRuntimeConfigDetails = &tmp
		}
	}

	if notebookSessionStorageMountConfigurationDetailsList, ok := s.D.GetOkExists("notebook_session_storage_mount_configuration_details_list"); ok {
		interfaces := notebookSessionStorageMountConfigurationDetailsList.([]interface{})
		tmp := make([]oci_datascience.StorageMountConfigurationDetails, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "notebook_session_storage_mount_configuration_details_list", stateDataIndex)
			converted, err := s.mapToStorageMountConfigurationDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("notebook_session_storage_mount_configuration_details_list") {
			request.NotebookSessionStorageMountConfigurationDetailsList = tmp
		}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "datascience")

	response, err := s.Client.UpdateNotebookSession(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.NotebookSession
	return nil
}

func (s *DatascienceNotebookSessionResourceCrud) Delete() error {
	request := oci_datascience.DeleteNotebookSessionRequest{}

	tmp := s.D.Id()
	request.NotebookSessionId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "datascience")

	_, err := s.Client.DeleteNotebookSession(context.Background(), request)
	return err
}

func (s *DatascienceNotebookSessionResourceCrud) SetData() error {
	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.CreatedBy != nil {
		s.D.Set("created_by", *s.Res.CreatedBy)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	if s.Res.NotebookSessionConfigDetails != nil {
		s.D.Set("notebook_session_config_details", []interface{}{NotebookSessionConfigDetailsToMap(s.Res.NotebookSessionConfigDetails)})
	} else {
		s.D.Set("notebook_session_config_details", nil)
	}

	if s.Res.NotebookSessionConfigurationDetails != nil {
		s.D.Set("notebook_session_configuration_details", []interface{}{NotebookSessionConfigurationDetailsToMap(s.Res.NotebookSessionConfigurationDetails)})
	} else {
		s.D.Set("notebook_session_configuration_details", nil)
	}

	if s.Res.NotebookSessionRuntimeConfigDetails != nil {
		s.D.Set("notebook_session_runtime_config_details", []interface{}{NotebookSessionRuntimeConfigDetailsToMap(s.Res.NotebookSessionRuntimeConfigDetails)})
	} else {
		s.D.Set("notebook_session_runtime_config_details", nil)
	}

	notebookSessionStorageMountConfigurationDetailsList := []interface{}{}
	for _, item := range s.Res.NotebookSessionStorageMountConfigurationDetailsList {
		notebookSessionStorageMountConfigurationDetailsList = append(notebookSessionStorageMountConfigurationDetailsList, StorageMountConfigurationDetailsToMap(item))
	}
	s.D.Set("notebook_session_storage_mount_configuration_details_list", notebookSessionStorageMountConfigurationDetailsList)

	if s.Res.NotebookSessionUrl != nil {
		s.D.Set("notebook_session_url", *s.Res.NotebookSessionUrl)
	}

	if s.Res.ProjectId != nil {
		s.D.Set("project_id", *s.Res.ProjectId)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	return nil
}

func (s *DatascienceNotebookSessionResourceCrud) mapToNotebookSessionConfigDetails(fieldKeyFormat string) (oci_datascience.NotebookSessionConfigDetails, error) {
	result := oci_datascience.NotebookSessionConfigDetails{}

	if blockStorageSizeInGBs, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "block_storage_size_in_gbs")); ok {
		tmp := blockStorageSizeInGBs.(int)
		result.BlockStorageSizeInGBs = &tmp
	}

	if notebookSessionShapeConfigDetails, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "notebook_session_shape_config_details")); ok {
		if tmpList := notebookSessionShapeConfigDetails.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "notebook_session_shape_config_details"), 0)
			tmp, err := s.mapToNotebookSessionShapeConfigDetails(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert notebook_session_shape_config_details, encountered error: %v", err)
			}
			result.NotebookSessionShapeConfigDetails = &tmp
		}
	}

	if privateEndpointId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "private_endpoint_id")); ok {
		tmp := privateEndpointId.(string)
		result.PrivateEndpointId = &tmp
	}

	if shape, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "shape")); ok {
		tmp := shape.(string)
		result.Shape = &tmp
	}

	if subnetId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "subnet_id")); ok {
		tmp := subnetId.(string)
		result.SubnetId = &tmp
	}

	return result, nil
}

func NotebookSessionConfigDetailsToMap(obj *oci_datascience.NotebookSessionConfigDetails) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.BlockStorageSizeInGBs != nil {
		result["block_storage_size_in_gbs"] = int(*obj.BlockStorageSizeInGBs)
	}

	if obj.NotebookSessionShapeConfigDetails != nil {
		result["notebook_session_shape_config_details"] = []interface{}{NotebookSessionShapeConfigDetailsToMap(obj.NotebookSessionShapeConfigDetails)}
	}

	if obj.PrivateEndpointId != nil {
		result["private_endpoint_id"] = string(*obj.PrivateEndpointId)
	}

	if obj.Shape != nil {
		result["shape"] = string(*obj.Shape)
	}

	if obj.SubnetId != nil {
		result["subnet_id"] = string(*obj.SubnetId)
	}

	return result
}

func (s *DatascienceNotebookSessionResourceCrud) mapToNotebookSessionConfigurationDetails(fieldKeyFormat string) (oci_datascience.NotebookSessionConfigurationDetails, error) {
	result := oci_datascience.NotebookSessionConfigurationDetails{}

	if blockStorageSizeInGBs, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "block_storage_size_in_gbs")); ok {
		tmp := blockStorageSizeInGBs.(int)
		result.BlockStorageSizeInGBs = &tmp
	}

	if notebookSessionShapeConfigDetails, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "notebook_session_shape_config_details")); ok {
		if tmpList := notebookSessionShapeConfigDetails.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "notebook_session_shape_config_details"), 0)
			tmp, err := s.mapToNotebookSessionShapeConfigDetails(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert notebook_session_shape_config_details, encountered error: %v", err)
			}
			result.NotebookSessionShapeConfigDetails = &tmp
		}
	}

	if privateEndpointId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "private_endpoint_id")); ok {
		tmp := privateEndpointId.(string)
		result.PrivateEndpointId = &tmp
	}

	if shape, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "shape")); ok {
		tmp := shape.(string)
		result.Shape = &tmp
	}

	if subnetId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "subnet_id")); ok {
		tmp := subnetId.(string)
		result.SubnetId = &tmp
	}

	return result, nil
}

func NotebookSessionConfigurationDetailsToMap(obj *oci_datascience.NotebookSessionConfigurationDetails) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.BlockStorageSizeInGBs != nil {
		result["block_storage_size_in_gbs"] = int(*obj.BlockStorageSizeInGBs)
	}

	if obj.NotebookSessionShapeConfigDetails != nil {
		result["notebook_session_shape_config_details"] = []interface{}{NotebookSessionShapeConfigDetailsToMap(obj.NotebookSessionShapeConfigDetails)}
	}

	if obj.PrivateEndpointId != nil {
		result["private_endpoint_id"] = string(*obj.PrivateEndpointId)
	}

	if obj.Shape != nil {
		result["shape"] = string(*obj.Shape)
	}

	if obj.SubnetId != nil {
		result["subnet_id"] = string(*obj.SubnetId)
	}

	return result
}

func (s *DatascienceNotebookSessionResourceCrud) mapToNotebookSessionGitConfigDetails(fieldKeyFormat string) (oci_datascience.NotebookSessionGitConfigDetails, error) {
	result := oci_datascience.NotebookSessionGitConfigDetails{}

	if notebookSessionGitRepoConfigCollection, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "notebook_session_git_repo_config_collection")); ok {
		interfaces := notebookSessionGitRepoConfigCollection.([]interface{})
		gitRepoConfigDetails := make([]oci_datascience.NotebookSessionGitRepoConfigDetails, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "notebook_session_git_repo_config_collection"), stateDataIndex)
			converted, err := s.mapToNotebookSessionGitRepoConfigDetails(fieldKeyFormatNextLevel)
			if err != nil {
				return result, err
			}
			gitRepoConfigDetails[i] = converted
		}
		if len(gitRepoConfigDetails) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "notebook_session_git_repo_config_collection")) {
			result.NotebookSessionGitRepoConfigCollection = gitRepoConfigDetails
		}
	}

	return result, nil
}

func NotebookSessionGitConfigDetailsToMap(obj *oci_datascience.NotebookSessionGitConfigDetails) map[string]interface{} {
	result := map[string]interface{}{}

	notebookSessionGitRepoConfigCollection := []interface{}{}
	for _, item := range obj.NotebookSessionGitRepoConfigCollection {
		notebookSessionGitRepoConfigCollection = append(notebookSessionGitRepoConfigCollection, NotebookSessionGitRepoConfigDetailsToMap(item))
	}
	result["notebook_session_git_repo_config_collection"] = notebookSessionGitRepoConfigCollection

	return result
}

func (s *DatascienceNotebookSessionResourceCrud) mapToNotebookSessionGitRepoConfigDetails(fieldKeyFormat string) (oci_datascience.NotebookSessionGitRepoConfigDetails, error) {
	result := oci_datascience.NotebookSessionGitRepoConfigDetails{}

	if url, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "url")); ok {
		tmp := url.(string)
		result.Url = &tmp
	}

	return result, nil
}

func NotebookSessionGitRepoConfigDetailsToMap(obj oci_datascience.NotebookSessionGitRepoConfigDetails) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Url != nil {
		result["url"] = string(*obj.Url)
	}

	return result
}

func (s *DatascienceNotebookSessionResourceCrud) mapToNotebookSessionRuntimeConfigDetails(fieldKeyFormat string) (oci_datascience.NotebookSessionRuntimeConfigDetails, error) {
	result := oci_datascience.NotebookSessionRuntimeConfigDetails{}

	if customEnvironmentVariables, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "custom_environment_variables")); ok {
		result.CustomEnvironmentVariables = tfresource.ObjectMapToStringMap(customEnvironmentVariables.(map[string]interface{}))
	}

	if notebookSessionGitConfigDetails, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "notebook_session_git_config_details")); ok {
		if tmpList := notebookSessionGitConfigDetails.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "notebook_session_git_config_details"), 0)
			tmp, err := s.mapToNotebookSessionGitConfigDetails(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert notebook_session_git_config_details, encountered error: %v", err)
			}
			result.NotebookSessionGitConfigDetails = &tmp
		}
	}

	return result, nil
}

func NotebookSessionRuntimeConfigDetailsToMap(obj *oci_datascience.NotebookSessionRuntimeConfigDetails) map[string]interface{} {
	result := map[string]interface{}{}

	result["custom_environment_variables"] = obj.CustomEnvironmentVariables

	if obj.NotebookSessionGitConfigDetails != nil {
		result["notebook_session_git_config_details"] = []interface{}{NotebookSessionGitConfigDetailsToMap(obj.NotebookSessionGitConfigDetails)}
	}

	return result
}

func (s *DatascienceNotebookSessionResourceCrud) mapToNotebookSessionShapeConfigDetails(fieldKeyFormat string) (oci_datascience.NotebookSessionShapeConfigDetails, error) {
	result := oci_datascience.NotebookSessionShapeConfigDetails{}

	if cpuBaseline, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "cpu_baseline")); ok {
		result.CpuBaseline = oci_datascience.NotebookSessionShapeConfigDetailsCpuBaselineEnum(cpuBaseline.(string))
	}

	if memoryInGBs, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "memory_in_gbs")); ok {
		tmp := float32(memoryInGBs.(float64))
		result.MemoryInGBs = &tmp
	}

	if ocpus, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "ocpus")); ok {
		tmp := float32(ocpus.(float64))
		result.Ocpus = &tmp
	}

	return result, nil
}

func NotebookSessionShapeConfigDetailsToMap(obj *oci_datascience.NotebookSessionShapeConfigDetails) map[string]interface{} {
	result := map[string]interface{}{}

	result["cpu_baseline"] = string(obj.CpuBaseline)

	if obj.MemoryInGBs != nil {
		result["memory_in_gbs"] = float32(*obj.MemoryInGBs)
	}

	if obj.Ocpus != nil {
		result["ocpus"] = float32(*obj.Ocpus)
	}

	return result
}

func (s *DatascienceNotebookSessionResourceCrud) mapToStorageMountConfigurationDetails(fieldKeyFormat string) (oci_datascience.StorageMountConfigurationDetails, error) {
	var baseObject oci_datascience.StorageMountConfigurationDetails
	//discriminator
	storageTypeRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "storage_type"))
	var storageType string
	if ok {
		storageType = storageTypeRaw.(string)
	} else {
		storageType = "" // default value
	}
	switch strings.ToLower(storageType) {
	case strings.ToLower("FILE_STORAGE"):
		details := oci_datascience.FileStorageMountConfigurationDetails{}
		if exportId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "export_id")); ok {
			tmp := exportId.(string)
			details.ExportId = &tmp
		}
		if mountTargetId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "mount_target_id")); ok {
			tmp := mountTargetId.(string)
			details.MountTargetId = &tmp
		}
		if destinationDirectoryName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "destination_directory_name")); ok {
			tmp := destinationDirectoryName.(string)
			details.DestinationDirectoryName = &tmp
		}
		if destinationPath, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "destination_path")); ok {
			tmp := destinationPath.(string)
			details.DestinationPath = &tmp
		}
		baseObject = details
	case strings.ToLower("OBJECT_STORAGE"):
		details := oci_datascience.ObjectStorageMountConfigurationDetails{}
		if bucket, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "bucket")); ok {
			tmp := bucket.(string)
			details.Bucket = &tmp
		}
		if namespace, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "namespace")); ok {
			tmp := namespace.(string)
			details.Namespace = &tmp
		}
		if prefix, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "prefix")); ok {
			tmp := prefix.(string)
			details.Prefix = &tmp
		}
		if destinationDirectoryName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "destination_directory_name")); ok {
			tmp := destinationDirectoryName.(string)
			details.DestinationDirectoryName = &tmp
		}
		if destinationPath, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "destination_path")); ok {
			tmp := destinationPath.(string)
			details.DestinationPath = &tmp
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown storage_type '%v' was specified", storageType)
	}
	return baseObject, nil
}

func (s *DatascienceNotebookSessionResourceCrud) updateCompartment(compartment interface{}) error {
	changeCompartmentRequest := oci_datascience.ChangeNotebookSessionCompartmentRequest{}

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	idTmp := s.D.Id()
	changeCompartmentRequest.NotebookSessionId = &idTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "datascience")

	_, err := s.Client.ChangeNotebookSessionCompartment(context.Background(), changeCompartmentRequest)
	if err != nil {
		return err
	}

	if waitErr := tfresource.WaitForUpdatedState(s.D, s); waitErr != nil {
		return waitErr
	}

	return nil
}

func (s *DatascienceNotebookSessionResourceCrud) ActivateNotebookSession() error {
	request := oci_datascience.ActivateNotebookSessionRequest{}

	tmp := s.D.Id()
	request.NotebookSessionId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "datascience")

	_, err := s.Client.ActivateNotebookSession(context.Background(), request)
	if err != nil {
		return err
	}

	retentionPolicyFunc := func() bool { return s.Res.LifecycleState == oci_datascience.NotebookSessionLifecycleStateActive }
	return tfresource.WaitForResourceCondition(s, retentionPolicyFunc, s.D.Timeout(schema.TimeoutUpdate))
}

func (s *DatascienceNotebookSessionResourceCrud) DeactivateNotebookSession() error {
	request := oci_datascience.DeactivateNotebookSessionRequest{}

	tmp := s.D.Id()
	request.NotebookSessionId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "datascience")

	_, err := s.Client.DeactivateNotebookSession(context.Background(), request)
	if err != nil {
		return err
	}

	retentionPolicyFunc := func() bool { return s.Res.LifecycleState == oci_datascience.NotebookSessionLifecycleStateInactive }
	return tfresource.WaitForResourceCondition(s, retentionPolicyFunc, s.D.Timeout(schema.TimeoutUpdate))
}
