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

func DataintegrationWorkspaceProjectResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createDataintegrationWorkspaceProject,
		Read:     readDataintegrationWorkspaceProject,
		Update:   updateDataintegrationWorkspaceProject,
		Delete:   deleteDataintegrationWorkspaceProject,
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
			"project_key": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"key": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
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
						},
						"is_favorite": {
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
						},
						"key": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"labels": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"registry_version": {
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},

						// Computed
					},
				},
			},

			// Computed
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
		},
	}
}

func createDataintegrationWorkspaceProject(d *schema.ResourceData, m interface{}) error {
	sync := &DataintegrationWorkspaceProjectResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataIntegrationClient()

	return tfresource.CreateResource(d, sync)
}

func readDataintegrationWorkspaceProject(d *schema.ResourceData, m interface{}) error {
	sync := &DataintegrationWorkspaceProjectResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataIntegrationClient()

	return tfresource.ReadResource(sync)
}

func updateDataintegrationWorkspaceProject(d *schema.ResourceData, m interface{}) error {
	sync := &DataintegrationWorkspaceProjectResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataIntegrationClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteDataintegrationWorkspaceProject(d *schema.ResourceData, m interface{}) error {
	sync := &DataintegrationWorkspaceProjectResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataIntegrationClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type DataintegrationWorkspaceProjectResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_dataintegration.DataIntegrationClient
	Res                    *oci_dataintegration.Project
	DisableNotFoundRetries bool
}

func (s *DataintegrationWorkspaceProjectResourceCrud) ID() string {
	return GetWorkspaceProjectCompositeId(*s.Res.Key, s.D.Get("workspace_id").(string))
}

func (s *DataintegrationWorkspaceProjectResourceCrud) Create() error {
	request := oci_dataintegration.CreateProjectRequest{}

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

	if workspaceId, ok := s.D.GetOkExists("workspace_id"); ok {
		tmp := workspaceId.(string)
		request.WorkspaceId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "dataintegration")

	response, err := s.Client.CreateProject(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Project
	return nil
}

func (s *DataintegrationWorkspaceProjectResourceCrud) Get() error {
	request := oci_dataintegration.GetProjectRequest{}

	if projectKey, ok := s.D.GetOkExists("project_key"); ok {
		tmp := projectKey.(string)
		request.ProjectKey = &tmp
	}

	if workspaceId, ok := s.D.GetOkExists("workspace_id"); ok {
		tmp := workspaceId.(string)
		request.WorkspaceId = &tmp
	}

	projectKey, workspaceId, err := parseWorkspaceProjectCompositeId(s.D.Id())
	if err == nil {
		request.ProjectKey = &projectKey
		request.WorkspaceId = &workspaceId
	} else {
		log.Printf("[WARN] Get() unable to parse current ID: %s", s.D.Id())
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "dataintegration")

	response, err := s.Client.GetProject(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Project
	return nil
}

func (s *DataintegrationWorkspaceProjectResourceCrud) Update() error {
	request := oci_dataintegration.UpdateProjectRequest{}

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

	if projectKey, ok := s.D.GetOkExists("project_key"); ok {
		tmp := projectKey.(string)
		request.ProjectKey = &tmp
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

	response, err := s.Client.UpdateProject(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Project
	return nil
}

func (s *DataintegrationWorkspaceProjectResourceCrud) Delete() error {
	request := oci_dataintegration.DeleteProjectRequest{}

	if projectKey, ok := s.D.GetOkExists("project_key"); ok {
		tmp := projectKey.(string)
		request.ProjectKey = &tmp
	}

	if workspaceId, ok := s.D.GetOkExists("workspace_id"); ok {
		tmp := workspaceId.(string)
		request.WorkspaceId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "dataintegration")

	_, err := s.Client.DeleteProject(context.Background(), request)
	return err
}

func (s *DataintegrationWorkspaceProjectResourceCrud) SetData() error {

	projectKey, workspaceId, err := parseWorkspaceProjectCompositeId(s.D.Id())
	if err == nil {
		s.D.Set("project_key", &projectKey)
		s.D.Set("workspace_id", &workspaceId)
	} else {
		log.Printf("[WARN] SetData() unable to parse current ID: %s", s.D.Id())
	}

	if s.Res.Description != nil {
		s.D.Set("description", *s.Res.Description)
	}

	if s.Res.Identifier != nil {
		s.D.Set("identifier", *s.Res.Identifier)
	}

	if s.Res.Key != nil {
		s.D.Set("key", *s.Res.Key)
	}

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

	return nil
}

func GetWorkspaceProjectCompositeId(projectKey string, workspaceId string) string {
	projectKey = url.PathEscape(projectKey)
	workspaceId = url.PathEscape(workspaceId)
	compositeId := "workspaces/" + workspaceId + "/projects/" + projectKey
	return compositeId
}

func parseWorkspaceProjectCompositeId(compositeId string) (projectKey string, workspaceId string, err error) {
	parts := strings.Split(compositeId, "/")
	match, _ := regexp.MatchString("workspaces/.*/projects/.*", compositeId)
	if !match || len(parts) != 4 {
		err = fmt.Errorf("illegal compositeId %s encountered", compositeId)
		return
	}
	workspaceId, _ = url.PathUnescape(parts[1])
	projectKey, _ = url.PathUnescape(parts[3])
	return
}

func AggregatorSummaryToMap(obj *oci_dataintegration.AggregatorSummary) map[string]interface{} {
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

func CountStatisticToMap(obj *oci_dataintegration.CountStatistic) map[string]interface{} {
	result := map[string]interface{}{}

	objectTypeCountList := []interface{}{}
	for _, item := range obj.ObjectTypeCountList {
		objectTypeCountList = append(objectTypeCountList, CountStatisticSummaryToMap(item))
	}
	result["object_type_count_list"] = objectTypeCountList

	return result
}

func CountStatisticSummaryToMap(obj oci_dataintegration.CountStatisticSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.ObjectCount != nil {
		result["object_count"] = strconv.FormatInt(*obj.ObjectCount, 10)
	}

	result["object_type"] = string(obj.ObjectType)

	return result
}

func ObjectMetadataToMap(obj *oci_dataintegration.ObjectMetadata) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Aggregator != nil {
		result["aggregator"] = []interface{}{AggregatorSummaryToMap(obj.Aggregator)}
	}

	if obj.AggregatorKey != nil {
		result["aggregator_key"] = string(*obj.AggregatorKey)
	}

	if obj.CountStatistics != nil {
		result["count_statistics"] = []interface{}{CountStatisticToMap(obj.CountStatistics)}
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

func (s *DataintegrationWorkspaceProjectResourceCrud) mapToParentReference(fieldKeyFormat string) (oci_dataintegration.ParentReference, error) {
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

func ParentReferenceToMap(obj *oci_dataintegration.ParentReference) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Parent != nil {
		result["parent"] = string(*obj.Parent)
	}

	if obj.RootDocId != nil {
		result["root_doc_id"] = string(*obj.RootDocId)
	}

	return result
}

func DataintegrationProjectProjectSummaryToMap(obj oci_dataintegration.ProjectSummary) map[string]interface{} {
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

	return result
}

func (s *DataintegrationWorkspaceProjectResourceCrud) mapToRegistryMetadata(fieldKeyFormat string) (oci_dataintegration.RegistryMetadata, error) {
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

func WorkspaceProjectRegistryMetadataToMap(obj *oci_dataintegration.RegistryMetadata) map[string]interface{} {
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
