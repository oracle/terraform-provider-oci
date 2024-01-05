// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package dataintegration

import (
	"context"
	"fmt"
	"log"
	"net/url"
	"regexp"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	oci_dataintegration "github.com/oracle/oci-go-sdk/v65/dataintegration"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DataintegrationWorkspaceApplicationResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createDataintegrationWorkspaceApplication,
		Read:     readDataintegrationWorkspaceApplication,
		Update:   updateDataintegrationWorkspaceApplication,
		Delete:   deleteDataintegrationWorkspaceApplication,
		Schema: map[string]*schema.Schema{
			// Required
			"identifier": {
				Type:     schema.TypeString,
				Required: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"workspace_id": {
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
			"key": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"model_type": {
				Type:     schema.TypeString,
				Required: true,
			},
			"model_version": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"object_status": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
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
			"source_application_info": {
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
						"application_key": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},
						"copy_type": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},
						"workspace_id": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},

						// Computed
						"application_version": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"last_patch_key": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			// Computed
			"application_version": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"compartment_id": {
				Type:     schema.TypeString,
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
			"published_object_metadata": {
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
			"time_created": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_patched": {
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

func createDataintegrationWorkspaceApplication(d *schema.ResourceData, m interface{}) error {
	sync := &DataintegrationWorkspaceApplicationResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataIntegrationClient()

	return tfresource.CreateResource(d, sync)
}

func readDataintegrationWorkspaceApplication(d *schema.ResourceData, m interface{}) error {
	sync := &DataintegrationWorkspaceApplicationResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataIntegrationClient()

	return tfresource.ReadResource(sync)
}

func updateDataintegrationWorkspaceApplication(d *schema.ResourceData, m interface{}) error {
	sync := &DataintegrationWorkspaceApplicationResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataIntegrationClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteDataintegrationWorkspaceApplication(d *schema.ResourceData, m interface{}) error {
	sync := &DataintegrationWorkspaceApplicationResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataIntegrationClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type DataintegrationWorkspaceApplicationResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_dataintegration.DataIntegrationClient
	Res                    *oci_dataintegration.Application
	DisableNotFoundRetries bool
}

func (s *DataintegrationWorkspaceApplicationResourceCrud) ID() string {
	return GetWorkspaceApplicationCompositeId(*s.Res.Key, s.D.Get("workspace_id").(string))
}

func (s *DataintegrationWorkspaceApplicationResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_dataintegration.ApplicationLifecycleStateCreating),
	}
}

func (s *DataintegrationWorkspaceApplicationResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_dataintegration.ApplicationLifecycleStateActive),
	}
}

func (s *DataintegrationWorkspaceApplicationResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_dataintegration.ApplicationLifecycleStateDeleting),
	}
}

func (s *DataintegrationWorkspaceApplicationResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_dataintegration.ApplicationLifecycleStateDeleted),
	}
}

func (s *DataintegrationWorkspaceApplicationResourceCrud) Create() error {
	request := oci_dataintegration.CreateApplicationRequest{}

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

	if identifier, ok := s.D.GetOkExists("identifier"); ok {
		tmp := identifier.(string)
		request.Identifier = &tmp
	}

	if key, ok := s.D.GetOkExists("key"); ok {
		tmp := key.(string)
		request.Key = &tmp
	}

	if modelType, ok := s.D.GetOkExists("model_type"); ok {
		request.ModelType = oci_dataintegration.CreateApplicationDetailsModelTypeEnum(modelType.(string))
	}

	if modelVersion, ok := s.D.GetOkExists("model_version"); ok {
		tmp := modelVersion.(string)
		request.ModelVersion = &tmp
	}

	if name, ok := s.D.GetOkExists("name"); ok {
		tmp := name.(string)
		request.Name = &tmp
	}

	if objectStatus, ok := s.D.GetOkExists("object_status"); ok {
		tmp := objectStatus.(int)
		request.ObjectStatus = &tmp
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

	if sourceApplicationInfo, ok := s.D.GetOkExists("source_application_info"); ok {
		if tmpList := sourceApplicationInfo.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "source_application_info", 0)
			tmp, err := s.mapToCreateSourceApplicationInfo(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.SourceApplicationInfo = &tmp
		}
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_dataintegration.CreateApplicationDetailsLifecycleStateEnum(state.(string))
	}

	if workspaceId, ok := s.D.GetOkExists("workspace_id"); ok {
		tmp := workspaceId.(string)
		request.WorkspaceId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "dataintegration")

	response, err := s.Client.CreateApplication(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Application
	return nil
}

func (s *DataintegrationWorkspaceApplicationResourceCrud) Get() error {
	request := oci_dataintegration.GetApplicationRequest{}

	if applicationKey, ok := s.D.GetOkExists("key"); ok {
		tmp := applicationKey.(string)
		request.ApplicationKey = &tmp
	}

	if workspaceId, ok := s.D.GetOkExists("workspace_id"); ok {
		tmp := workspaceId.(string)
		request.WorkspaceId = &tmp
	}

	applicationKey, workspaceId, err := parseWorkspaceApplicationCompositeId(s.D.Id())
	if err == nil {
		request.ApplicationKey = &applicationKey
		request.WorkspaceId = &workspaceId
	} else {
		log.Printf("[WARN] Get() unable to parse current ID: %s", s.D.Id())
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "dataintegration")

	response, err := s.Client.GetApplication(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Application
	return nil
}

func (s *DataintegrationWorkspaceApplicationResourceCrud) Update() error {
	request := oci_dataintegration.UpdateApplicationRequest{}

	if applicationKey, ok := s.D.GetOkExists("key"); ok {
		tmp := applicationKey.(string)
		request.ApplicationKey = &tmp
	}

	if applicationVersion, ok := s.D.GetOkExists("application_version"); ok {
		tmp := applicationVersion.(int)
		request.ApplicationVersion = &tmp
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

	if identifier, ok := s.D.GetOkExists("identifier"); ok {
		tmp := identifier.(string)
		request.Identifier = &tmp
	}

	if key, ok := s.D.GetOkExists("key"); ok {
		tmp := key.(string)
		request.Key = &tmp
	}

	if metadata, ok := s.D.GetOkExists("metadata"); ok {
		if tmpList := metadata.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "metadata", 0)
			tmp, err := s.mapToObjectMetadata(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.Metadata = &tmp
		}
	}

	if modelType, ok := s.D.GetOkExists("model_type"); ok {
		tmp := modelType.(string)
		request.ModelType = &tmp
	}

	if modelVersion, ok := s.D.GetOkExists("model_version"); ok {
		tmp := modelVersion.(string)
		request.ModelVersion = &tmp
	}

	if name, ok := s.D.GetOkExists("name"); ok {
		tmp := name.(string)
		request.Name = &tmp
	}

	if objectStatus, ok := s.D.GetOkExists("object_status"); ok {
		tmp := objectStatus.(int)
		request.ObjectStatus = &tmp
	}

	if objectVersion, ok := s.D.GetOkExists("object_version"); ok {
		tmp := objectVersion.(int)
		request.ObjectVersion = &tmp
	}

	if parentRef, ok := s.D.GetOkExists("parent_ref"); ok {
		if tmpList := parentRef.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "parent_ref", 0)
			tmp, err := s.mapToParentReference(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.ParentRef = &tmp
		}
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_dataintegration.UpdateApplicationDetailsLifecycleStateEnum(state.(string))
	}

	if workspaceId, ok := s.D.GetOkExists("workspace_id"); ok {
		tmp := workspaceId.(string)
		request.WorkspaceId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "dataintegration")

	response, err := s.Client.UpdateApplication(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Application
	return nil
}

func (s *DataintegrationWorkspaceApplicationResourceCrud) Delete() error {
	request := oci_dataintegration.DeleteApplicationRequest{}

	if applicationKey, ok := s.D.GetOkExists("key"); ok {
		tmp := applicationKey.(string)
		request.ApplicationKey = &tmp
	}

	if workspaceId, ok := s.D.GetOkExists("workspace_id"); ok {
		tmp := workspaceId.(string)
		request.WorkspaceId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "dataintegration")

	_, err := s.Client.DeleteApplication(context.Background(), request)
	return err
}

func (s *DataintegrationWorkspaceApplicationResourceCrud) SetData() error {

	applicationKey, workspaceId, err := parseWorkspaceApplicationCompositeId(s.D.Id())
	if err == nil {
		s.D.Set("key", &applicationKey)
		s.D.Set("workspace_id", &workspaceId)
	} else {
		log.Printf("[WARN] SetData() unable to parse current ID: %s", s.D.Id())
	}

	if s.Res.ApplicationVersion != nil {
		s.D.Set("application_version", *s.Res.ApplicationVersion)
	}

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	dependentObjectMetadata := []interface{}{}
	for _, item := range s.Res.DependentObjectMetadata {
		dependentObjectMetadata = append(dependentObjectMetadata, PatchObjectMetadataToMap(item))
	}
	s.D.Set("dependent_object_metadata", dependentObjectMetadata)

	if s.Res.Description != nil {
		s.D.Set("description", *s.Res.Description)
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.Identifier != nil {
		s.D.Set("identifier", *s.Res.Identifier)
	}

	if s.Res.Key != nil {
		s.D.Set("key", *s.Res.Key)
	}

	s.D.Set("key_map", s.Res.KeyMap)

	if s.Res.Metadata != nil {
		s.D.Set("metadata", []interface{}{ApplicationObjectMetadataToMap(s.Res.Metadata)})
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
		s.D.Set("parent_ref", []interface{}{ApplicationParentReferenceToMap(s.Res.ParentRef)})
	} else {
		s.D.Set("parent_ref", nil)
	}

	publishedObjectMetadataList := []interface{}{}
	for _, value := range s.Res.PublishedObjectMetadata {
		publishedObjectMetadataList = append(publishedObjectMetadataList, PatchObjectMetadataToMap(value))
	}
	s.D.Set("published_object_metadata", publishedObjectMetadataList)

	if s.Res.SourceApplicationInfo != nil {
		s.D.Set("source_application_info", []interface{}{SourceApplicationInfoToMap(s.Res.SourceApplicationInfo)})
	} else {
		s.D.Set("source_application_info", nil)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimePatched != nil {
		s.D.Set("time_patched", s.Res.TimePatched.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	return nil
}

func GetWorkspaceApplicationCompositeId(applicationKey string, workspaceId string) string {
	applicationKey = url.PathEscape(applicationKey)
	workspaceId = url.PathEscape(workspaceId)
	compositeId := "workspaces/" + workspaceId + "/applications/" + applicationKey
	return compositeId
}

func parseWorkspaceApplicationCompositeId(compositeId string) (applicationKey string, workspaceId string, err error) {
	parts := strings.Split(compositeId, "/")
	match, _ := regexp.MatchString("workspaces/.*/applications/.*", compositeId)
	if !match || len(parts) != 4 {
		err = fmt.Errorf("illegal compositeId %s encountered", compositeId)
		return
	}
	workspaceId, _ = url.PathUnescape(parts[1])
	applicationKey, _ = url.PathUnescape(parts[3])

	return
}

func (s *DataintegrationWorkspaceApplicationResourceCrud) mapToAggregatorSummary(fieldKeyFormat string) (oci_dataintegration.AggregatorSummary, error) {
	result := oci_dataintegration.AggregatorSummary{}

	if description, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "description")); ok {
		tmp := description.(string)
		result.Description = &tmp
	}

	if identifier, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "identifier")); ok {
		tmp := identifier.(string)
		result.Identifier = &tmp
	}

	if key, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "key")); ok {
		tmp := key.(string)
		result.Key = &tmp
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

func ApplicationAggregatorSummaryToMap(obj *oci_dataintegration.AggregatorSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Description != nil {
		result["description"] = string(*obj.Description)
	}

	if obj.Identifier != nil {
		result["identifier"] = string(*obj.Identifier)
	}

	if obj.Key != nil {
		result["key"] = string(*obj.Key)
	}

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	if obj.Type != nil {
		result["type"] = string(*obj.Type)
	}

	return result
}

func ApplicationSummaryToMap(obj oci_dataintegration.ApplicationSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.ApplicationVersion != nil {
		result["application_version"] = int(*obj.ApplicationVersion)
	}

	if obj.CompartmentId != nil {
		result["compartment_id"] = string(*obj.CompartmentId)
	}

	if obj.DefinedTags != nil {
		result["defined_tags"] = tfresource.DefinedTagsToMap(obj.DefinedTags)
	}

	dependentObjectMetadata := []interface{}{}
	for _, item := range obj.DependentObjectMetadata {
		dependentObjectMetadata = append(dependentObjectMetadata, PatchObjectMetadataToMap(item))
	}
	result["dependent_object_metadata"] = dependentObjectMetadata

	if obj.Description != nil {
		result["description"] = string(*obj.Description)
	}

	if obj.DisplayName != nil {
		result["display_name"] = string(*obj.DisplayName)
	}

	result["freeform_tags"] = obj.FreeformTags

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	if obj.Identifier != nil {
		result["identifier"] = string(*obj.Identifier)
	}

	if obj.Key != nil {
		result["key"] = string(*obj.Key)
	}

	result["key_map"] = obj.KeyMap

	if obj.Metadata != nil {
		result["metadata"] = []interface{}{ApplicationObjectMetadataToMap(obj.Metadata)}
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
		result["parent_ref"] = []interface{}{ApplicationParentReferenceToMap(obj.ParentRef)}
	}

	publishedObjectMetadataList := []interface{}{}
	for _, value := range obj.PublishedObjectMetadata {
		publishedObjectMetadataList = append(publishedObjectMetadataList, PatchObjectMetadataToMap(value))
	}
	result["published_object_metadata"] = publishedObjectMetadataList

	if obj.SourceApplicationInfo != nil {
		result["source_application_info"] = []interface{}{SourceApplicationInfoToMap(obj.SourceApplicationInfo)}
	}

	result["state"] = string(obj.LifecycleState)

	if obj.TimeCreated != nil {
		result["time_created"] = obj.TimeCreated.String()
	}

	if obj.TimePatched != nil {
		result["time_patched"] = obj.TimePatched.String()
	}

	if obj.TimeUpdated != nil {
		result["time_updated"] = obj.TimeUpdated.String()
	}

	return result
}

func (s *DataintegrationWorkspaceApplicationResourceCrud) mapToCountStatistic(fieldKeyFormat string) (oci_dataintegration.CountStatistic, error) {
	result := oci_dataintegration.CountStatistic{}

	if objectTypeCountList, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "object_type_count_list")); ok {
		interfaces := objectTypeCountList.([]interface{})
		tmp := make([]oci_dataintegration.CountStatisticSummary, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "object_type_count_list"), stateDataIndex)
			converted, err := s.mapToCountStatisticSummary(fieldKeyFormatNextLevel)
			if err != nil {
				return result, err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "object_type_count_list")) {
			result.ObjectTypeCountList = tmp
		}
	}

	return result, nil
}

func ApplicationCountStatisticToMap(obj *oci_dataintegration.CountStatistic) map[string]interface{} {
	result := map[string]interface{}{}

	objectTypeCountList := []interface{}{}
	for _, item := range obj.ObjectTypeCountList {
		objectTypeCountList = append(objectTypeCountList, ApplicationCountStatisticSummaryToMap(item))
	}
	result["object_type_count_list"] = objectTypeCountList

	return result
}

func (s *DataintegrationWorkspaceApplicationResourceCrud) mapToCountStatisticSummary(fieldKeyFormat string) (oci_dataintegration.CountStatisticSummary, error) {
	result := oci_dataintegration.CountStatisticSummary{}

	if objectCount, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "object_count")); ok {
		tmp := objectCount.(string)
		tmpInt64, err := strconv.ParseInt(tmp, 10, 64)
		if err != nil {
			return result, fmt.Errorf("unable to convert objectCount string: %s to an int64 and encountered error: %v", tmp, err)
		}
		result.ObjectCount = &tmpInt64
	}

	if objectType, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "object_type")); ok {
		result.ObjectType = oci_dataintegration.CountStatisticSummaryObjectTypeEnum(objectType.(string))
	}

	return result, nil
}

func ApplicationCountStatisticSummaryToMap(obj oci_dataintegration.CountStatisticSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.ObjectCount != nil {
		result["object_count"] = strconv.FormatInt(*obj.ObjectCount, 10)
	}

	result["object_type"] = string(obj.ObjectType)

	return result
}

func (s *DataintegrationWorkspaceApplicationResourceCrud) mapToCreateSourceApplicationInfo(fieldKeyFormat string) (oci_dataintegration.CreateSourceApplicationInfo, error) {
	result := oci_dataintegration.CreateSourceApplicationInfo{}

	if applicationKey, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "application_key")); ok {
		tmp := applicationKey.(string)
		result.ApplicationKey = &tmp
	}

	if copyType, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "copy_type")); ok {
		result.CopyType = oci_dataintegration.CreateSourceApplicationInfoCopyTypeEnum(copyType.(string))
	}

	if workspaceId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "workspace_id")); ok {
		tmp := workspaceId.(string)
		result.WorkspaceId = &tmp
	}

	return result, nil
}

func SourceApplicationInfoToMap(obj *oci_dataintegration.SourceApplicationInfo) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.ApplicationKey != nil {
		result["application_key"] = string(*obj.ApplicationKey)
	}

	if obj.ApplicationVersion != nil {
		result["application_version"] = string(*obj.ApplicationVersion)
	}

	if obj.LastPatchKey != nil {
		result["last_patch_key"] = string(*obj.LastPatchKey)
	}

	if obj.WorkspaceId != nil {
		result["workspace_id"] = string(*obj.WorkspaceId)
	}

	return result
}

func (s *DataintegrationWorkspaceApplicationResourceCrud) mapToObjectMetadata(fieldKeyFormat string) (oci_dataintegration.ObjectMetadata, error) {
	result := oci_dataintegration.ObjectMetadata{}

	if aggregator, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "aggregator")); ok {
		if tmpList := aggregator.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "aggregator"), 0)
			tmp, err := s.mapToAggregatorSummary(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert aggregator, encountered error: %v", err)
			}
			result.Aggregator = &tmp
		}
	}

	if aggregatorKey, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "aggregator_key")); ok {
		tmp := aggregatorKey.(string)
		result.AggregatorKey = &tmp
	}

	if countStatistics, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "count_statistics")); ok {
		if tmpList := countStatistics.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "count_statistics"), 0)
			tmp, err := s.mapToCountStatistic(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert count_statistics, encountered error: %v", err)
			}
			result.CountStatistics = &tmp
		}
	}

	if createdBy, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "created_by")); ok {
		tmp := createdBy.(string)
		result.CreatedBy = &tmp
	}

	if createdByName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "created_by_name")); ok {
		tmp := createdByName.(string)
		result.CreatedByName = &tmp
	}

	if identifierPath, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "identifier_path")); ok {
		tmp := identifierPath.(string)
		result.IdentifierPath = &tmp
	}

	if infoFields, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "info_fields")); ok {
		result.InfoFields = tfresource.ObjectMapToStringMap(infoFields.(map[string]interface{}))
	}

	if isFavorite, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_favorite")); ok {
		tmp := isFavorite.(bool)
		result.IsFavorite = &tmp
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

	if updatedBy, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "updated_by")); ok {
		tmp := updatedBy.(string)
		result.UpdatedBy = &tmp
	}

	if updatedByName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "updated_by_name")); ok {
		tmp := updatedByName.(string)
		result.UpdatedByName = &tmp
	}

	return result, nil
}

func ApplicationObjectMetadataToMap(obj *oci_dataintegration.ObjectMetadata) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Aggregator != nil {
		result["aggregator"] = []interface{}{ApplicationAggregatorSummaryToMap(obj.Aggregator)}
	}

	if obj.AggregatorKey != nil {
		result["aggregator_key"] = string(*obj.AggregatorKey)
	}

	if obj.CountStatistics != nil {
		result["count_statistics"] = []interface{}{ApplicationCountStatisticToMap(obj.CountStatistics)}
	}

	if obj.CreatedBy != nil {
		result["created_by"] = string(*obj.CreatedBy)
	}

	if obj.CreatedByName != nil {
		result["created_by_name"] = string(*obj.CreatedByName)
	}

	if obj.IdentifierPath != nil {
		result["identifier_path"] = string(*obj.IdentifierPath)
	}

	result["info_fields"] = obj.InfoFields

	if obj.IsFavorite != nil {
		result["is_favorite"] = bool(*obj.IsFavorite)
	}

	result["labels"] = obj.Labels

	if obj.RegistryVersion != nil {
		result["registry_version"] = int(*obj.RegistryVersion)
	}

	if obj.TimeCreated != nil {
		result["time_created"] = obj.TimeCreated.String()
	}

	if obj.TimeUpdated != nil {
		result["time_updated"] = obj.TimeUpdated.String()
	}

	if obj.UpdatedBy != nil {
		result["updated_by"] = string(*obj.UpdatedBy)
	}

	if obj.UpdatedByName != nil {
		result["updated_by_name"] = string(*obj.UpdatedByName)
	}

	return result
}

func (s *DataintegrationWorkspaceApplicationResourceCrud) mapToParentReference(fieldKeyFormat string) (oci_dataintegration.ParentReference, error) {
	result := oci_dataintegration.ParentReference{}

	if parent, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "parent")); ok {
		tmp := parent.(string)
		result.Parent = &tmp
	}

	if rootDocId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "root_doc_id")); ok {
		tmp := rootDocId.(string)
		result.RootDocId = &tmp
	}

	return result, nil
}

func ApplicationParentReferenceToMap(obj *oci_dataintegration.ParentReference) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Parent != nil {
		result["parent"] = string(*obj.Parent)
	}

	if obj.RootDocId != nil {
		result["root_doc_id"] = string(*obj.RootDocId)
	}

	return result
}

func PatchObjectMetadataToMap(obj oci_dataintegration.PatchObjectMetadata) map[string]interface{} {
	result := map[string]interface{}{}

	result["action"] = string(obj.Action)

	if obj.Identifier != nil {
		result["identifier"] = string(*obj.Identifier)
	}

	if obj.Key != nil {
		result["key"] = string(*obj.Key)
	}

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	if obj.NamePath != nil {
		result["name_path"] = string(*obj.NamePath)
	}

	if obj.ObjectVersion != nil {
		result["object_version"] = int(*obj.ObjectVersion)
	}

	result["type"] = string(obj.Type)

	return result
}

func (s *DataintegrationWorkspaceApplicationResourceCrud) mapToRegistryMetadata(fieldKeyFormat string) (oci_dataintegration.RegistryMetadata, error) {
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

func RegistryMetadataToMap(obj *oci_dataintegration.RegistryMetadata) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.AggregatorKey != nil {
		result["aggregator_key"] = string(*obj.AggregatorKey)
	}

	if obj.IsFavorite != nil {
		result["is_favorite"] = bool(*obj.IsFavorite)
	}

	if obj.Key != nil {
		result["key"] = string(*obj.Key)
	}

	result["labels"] = obj.Labels

	if obj.RegistryVersion != nil {
		result["registry_version"] = int(*obj.RegistryVersion)
	}

	return result
}
