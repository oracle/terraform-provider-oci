// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package dataintegration

import (
	"context"
	"fmt"
	"log"
	"net/url"
	"regexp"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	oci_dataintegration "github.com/oracle/oci-go-sdk/v65/dataintegration"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DataintegrationWorkspaceApplicationPatchResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createDataintegrationWorkspaceApplicationPatch,
		Read:     readDataintegrationWorkspaceApplicationPatch,
		Delete:   deleteDataintegrationWorkspaceApplicationPatch,
		Schema: map[string]*schema.Schema{
			// Required
			"application_key": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"identifier": {
				Type:     schema.TypeString,
				Optional: true,
				// identifier is computed based on name.
				Computed: true,
				ForceNew: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"object_keys": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				ForceNew: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"patch_type": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"workspace_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			// Optional
			"description": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"key": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"model_version": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"object_status": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"registry_metadata": {
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
						"aggregator_key": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},
						"is_favorite": {
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},
						"key": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},
						"labels": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							ForceNew: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"registry_version": {
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},

						// Computed
					},
				},
			},

			// Computed
			"application_version": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"dependent_object_metadata": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"action": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"identifier": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"key": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"name_path": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"object_version": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"type": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"error_messages": {
				Type:     schema.TypeMap,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"key_map": {
				Type:     schema.TypeMap,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"metadata": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"aggregator": {
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
									"identifier": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"key": {
										Type:     schema.TypeString,
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
								},
							},
						},
						"aggregator_key": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"count_statistics": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"object_type_count_list": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional

												// Computed
												"object_count": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"object_type": {
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},
								},
							},
						},
						"created_by": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"created_by_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"identifier_path": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"info_fields": {
							Type:     schema.TypeMap,
							Computed: true,
							Elem:     schema.TypeString,
						},
						"is_favorite": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"labels": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"registry_version": {
							Type:     schema.TypeInt,
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
						"updated_by": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"updated_by_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"model_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"object_version": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"parent_ref": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"parent": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"root_doc_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"patch_object_metadata": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"action": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"identifier": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"key": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"name_path": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"object_version": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"type": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"patch_status": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_patched": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func createDataintegrationWorkspaceApplicationPatch(d *schema.ResourceData, m interface{}) error {
	sync := &DataintegrationWorkspaceApplicationPatchResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataIntegrationClient()

	return tfresource.CreateResource(d, sync)
}

func readDataintegrationWorkspaceApplicationPatch(d *schema.ResourceData, m interface{}) error {
	sync := &DataintegrationWorkspaceApplicationPatchResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataIntegrationClient()

	return tfresource.ReadResource(sync)
}

func deleteDataintegrationWorkspaceApplicationPatch(d *schema.ResourceData, m interface{}) error {
	sync := &DataintegrationWorkspaceApplicationPatchResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataIntegrationClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type DataintegrationWorkspaceApplicationPatchResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_dataintegration.DataIntegrationClient
	Res                    *oci_dataintegration.Patch
	DisableNotFoundRetries bool
}

func (s *DataintegrationWorkspaceApplicationPatchResourceCrud) ID() string {
	return GetWorkspaceApplicationPatchCompositeId(s.D.Get("application_key").(string), *s.Res.Key, s.D.Get("workspace_id").(string))
}

func (s *DataintegrationWorkspaceApplicationPatchResourceCrud) Create() error {
	request := oci_dataintegration.CreatePatchRequest{}

	if applicationKey, ok := s.D.GetOkExists("application_key"); ok {
		tmp := applicationKey.(string)
		request.ApplicationKey = &tmp
	}

	if description, ok := s.D.GetOkExists("description"); ok {
		tmp := description.(string)
		request.Description = &tmp
	}

	if identifier, ok := s.D.GetOkExists("identifier"); ok {
		tmp := identifier.(string)
		request.Identifier = &tmp
	}

	if key, ok := s.D.GetOkExists("key"); ok {
		tmp := key.(string)
		request.Key = &tmp
	}

	if modelVersion, ok := s.D.GetOkExists("model_version"); ok {
		tmp := modelVersion.(string)
		request.ModelVersion = &tmp
	}

	if name, ok := s.D.GetOkExists("name"); ok {
		tmp := name.(string)
		request.Name = &tmp
	}

	if objectKeys, ok := s.D.GetOkExists("object_keys"); ok {
		interfaces := objectKeys.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("object_keys") {
			request.ObjectKeys = tmp
		}
	}

	if objectStatus, ok := s.D.GetOkExists("object_status"); ok {
		tmp := objectStatus.(int)
		request.ObjectStatus = &tmp
	}

	if patchType, ok := s.D.GetOkExists("patch_type"); ok {
		request.PatchType = oci_dataintegration.CreatePatchDetailsPatchTypeEnum(patchType.(string))
	}

	if registryMetadata, ok := s.D.GetOkExists("registry_metadata"); ok {
		if tmpList := registryMetadata.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "registry_metadata", 0)
			tmp, err := s.mapToRegistryMetadata(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.RegistryMetadata = &tmp
		}
	}

	if workspaceId, ok := s.D.GetOkExists("workspace_id"); ok {
		tmp := workspaceId.(string)
		request.WorkspaceId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "dataintegration")

	response, err := s.Client.CreatePatch(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Patch

	s.D.SetId(GetWorkspaceApplicationPatchCompositeId(*request.ApplicationKey, *response.Key, *request.WorkspaceId))

	retentionPolicyFunc := func() bool {
		return s.Res.PatchStatus == oci_dataintegration.PatchPatchStatusFailed || s.Res.PatchStatus == oci_dataintegration.PatchPatchStatusSuccessful
	}

	if err := tfresource.WaitForResourceCondition(s, retentionPolicyFunc, s.D.Timeout(schema.TimeoutUpdate)); err != nil {
		return err
	}

	return nil
}

func (s *DataintegrationWorkspaceApplicationPatchResourceCrud) Get() error {
	request := oci_dataintegration.GetPatchRequest{}

	if applicationKey, ok := s.D.GetOkExists("application_key"); ok {
		tmp := applicationKey.(string)
		request.ApplicationKey = &tmp
	}

	if patchKey, ok := s.D.GetOkExists("key"); ok {
		tmp := patchKey.(string)
		request.PatchKey = &tmp
	}

	if workspaceId, ok := s.D.GetOkExists("workspace_id"); ok {
		tmp := workspaceId.(string)
		request.WorkspaceId = &tmp
	}

	applicationKey, patchKey, workspaceId, err := parseWorkspaceApplicationPatchCompositeId(s.D.Id())
	if err == nil {
		request.ApplicationKey = &applicationKey
		request.PatchKey = &patchKey
		request.WorkspaceId = &workspaceId
	} else {
		log.Printf("[WARN] Get() unable to parse current ID: %s", s.D.Id())
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "dataintegration")

	response, err := s.Client.GetPatch(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Patch
	return nil
}

func (s *DataintegrationWorkspaceApplicationPatchResourceCrud) Delete() error {
	request := oci_dataintegration.DeletePatchRequest{}

	if applicationKey, ok := s.D.GetOkExists("application_key"); ok {
		tmp := applicationKey.(string)
		request.ApplicationKey = &tmp
	}

	if patchKey, ok := s.D.GetOkExists("key"); ok {
		tmp := patchKey.(string)
		request.PatchKey = &tmp
	}

	if workspaceId, ok := s.D.GetOkExists("workspace_id"); ok {
		tmp := workspaceId.(string)
		request.WorkspaceId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "dataintegration")

	_, err := s.Client.DeletePatch(context.Background(), request)

	if err != nil && strings.Contains(err.Error(), "Cannot delete a patch in 'SUCCESSFUL' state") {
		fmt.Printf("[ERROR] Fail to delete a patch in 'SUCCESSFUL' state: %s", err)
		return nil
	}

	return err
}

func (s *DataintegrationWorkspaceApplicationPatchResourceCrud) SetData() error {

	applicationKey, patchKey, workspaceId, err := parseWorkspaceApplicationPatchCompositeId(s.D.Id())
	if err == nil {
		s.D.Set("application_key", &applicationKey)
		s.D.Set("key", &patchKey)
		s.D.Set("workspace_id", &workspaceId)
	} else {
		log.Printf("[WARN] SetData() unable to parse current ID: %s", s.D.Id())
	}

	if s.Res.ApplicationVersion != nil {
		s.D.Set("application_version", *s.Res.ApplicationVersion)
	}

	dependentObjectMetadata := []interface{}{}
	for _, item := range s.Res.DependentObjectMetadata {
		dependentObjectMetadata = append(dependentObjectMetadata, PatchObjectMetadataToMap(item))
	}
	s.D.Set("dependent_object_metadata", dependentObjectMetadata)

	if s.Res.Description != nil {
		s.D.Set("description", *s.Res.Description)
	}

	s.D.Set("error_messages", s.Res.ErrorMessages)
	s.D.Set("error_messages", s.Res.ErrorMessages)

	if s.Res.Identifier != nil {
		s.D.Set("identifier", *s.Res.Identifier)
	}

	if s.Res.Key != nil {
		s.D.Set("key", *s.Res.Key)
	}

	s.D.Set("key_map", s.Res.KeyMap)
	s.D.Set("key_map", s.Res.KeyMap)

	if s.Res.Metadata != nil {
		s.D.Set("metadata", []interface{}{ObjectMetadataToMap(s.Res.Metadata)})
	} else {
		s.D.Set("metadata", nil)
	}

	if s.Res.ModelType != nil {
		s.D.Set("model_type", *s.Res.ModelType)
	}

	if s.Res.ModelVersion != nil {
		s.D.Set("model_version", *s.Res.ModelVersion)
	}

	if s.Res.Name != nil {
		s.D.Set("name", *s.Res.Name)
	}

	if s.Res.ObjectStatus != nil {
		s.D.Set("object_status", *s.Res.ObjectStatus)
	}

	if s.Res.ObjectVersion != nil {
		s.D.Set("object_version", *s.Res.ObjectVersion)
	}

	if s.Res.ParentRef != nil {
		s.D.Set("parent_ref", []interface{}{ParentReferenceToMap(s.Res.ParentRef)})
	} else {
		s.D.Set("parent_ref", nil)
	}

	patchObjectMetadata := []interface{}{}
	for _, item := range s.Res.PatchObjectMetadata {
		patchObjectMetadata = append(patchObjectMetadata, PatchObjectMetadataToMap(item))
	}
	s.D.Set("patch_object_metadata", patchObjectMetadata)

	s.D.Set("patch_status", s.Res.PatchStatus)

	s.D.Set("patch_type", s.Res.PatchType)

	if s.Res.TimePatched != nil {
		s.D.Set("time_patched", s.Res.TimePatched.String())
	}

	return nil
}

func GetWorkspaceApplicationPatchCompositeId(applicationKey string, patchKey string, workspaceId string) string {
	applicationKey = url.PathEscape(applicationKey)
	patchKey = url.PathEscape(patchKey)
	workspaceId = url.PathEscape(workspaceId)
	compositeId := "workspaces/" + workspaceId + "/applications/" + applicationKey + "/patches/" + patchKey
	return compositeId
}

func parseWorkspaceApplicationPatchCompositeId(compositeId string) (applicationKey string, patchKey string, workspaceId string, err error) {
	parts := strings.Split(compositeId, "/")
	match, _ := regexp.MatchString("workspaces/.*/applications/.*/patches/.*", compositeId)
	if !match || len(parts) != 6 {
		err = fmt.Errorf("illegal compositeId %s encountered", compositeId)
		return
	}
	workspaceId, _ = url.PathUnescape(parts[1])
	applicationKey, _ = url.PathUnescape(parts[3])
	patchKey, _ = url.PathUnescape(parts[5])

	return
}

func PatchSummaryToMap(obj oci_dataintegration.PatchSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.ApplicationVersion != nil {
		result["application_version"] = int(*obj.ApplicationVersion)
	}

	dependentObjectMetadata := []interface{}{}
	for _, item := range obj.DependentObjectMetadata {
		dependentObjectMetadata = append(dependentObjectMetadata, PatchObjectMetadataToMap(item))
	}
	result["dependent_object_metadata"] = dependentObjectMetadata

	if obj.Description != nil {
		result["description"] = string(*obj.Description)
	}

	result["error_messages"] = obj.ErrorMessages
	result["error_messages"] = obj.ErrorMessages

	if obj.Identifier != nil {
		result["identifier"] = string(*obj.Identifier)
	}

	if obj.Key != nil {
		result["key"] = string(*obj.Key)
	}

	result["key_map"] = obj.KeyMap
	result["key_map"] = obj.KeyMap

	if obj.Metadata != nil {
		result["metadata"] = []interface{}{ObjectMetadataToMap(obj.Metadata)}
	}

	if obj.ModelType != nil {
		result["model_type"] = string(*obj.ModelType)
	}

	if obj.ModelVersion != nil {
		result["model_version"] = string(*obj.ModelVersion)
	}

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	if obj.ObjectStatus != nil {
		result["object_status"] = int(*obj.ObjectStatus)
	}

	if obj.ObjectVersion != nil {
		result["object_version"] = int(*obj.ObjectVersion)
	}

	if obj.ParentRef != nil {
		result["parent_ref"] = []interface{}{ParentReferenceToMap(obj.ParentRef)}
	}

	patchObjectMetadata := []interface{}{}
	for _, item := range obj.PatchObjectMetadata {
		patchObjectMetadata = append(patchObjectMetadata, PatchObjectMetadataToMap(item))
	}
	result["patch_object_metadata"] = patchObjectMetadata

	result["patch_status"] = string(obj.PatchStatus)

	result["patch_type"] = string(obj.PatchType)

	if obj.TimePatched != nil {
		result["time_patched"] = obj.TimePatched.String()
	}

	return result
}

func (s *DataintegrationWorkspaceApplicationPatchResourceCrud) mapToRegistryMetadata(fieldKeyFormat string) (oci_dataintegration.RegistryMetadata, error) {
	result := oci_dataintegration.RegistryMetadata{}

	if aggregatorKey, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "aggregator_key")); ok {
		tmp := aggregatorKey.(string)
		result.AggregatorKey = &tmp
	}

	if isFavorite, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_favorite")); ok {
		tmp := isFavorite.(bool)
		result.IsFavorite = &tmp
	}

	if key, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "key")); ok {
		tmp := key.(string)
		result.Key = &tmp
	}

	if labels, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "labels")); ok {
		interfaces := labels.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "labels")) {
			result.Labels = tmp
		}
	}

	if registryVersion, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "registry_version")); ok {
		tmp := registryVersion.(int)
		result.RegistryVersion = &tmp
	}

	return result, nil
}
