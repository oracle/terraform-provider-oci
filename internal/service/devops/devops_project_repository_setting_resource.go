// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package devops

import (
	"context"
	"fmt"
	"log"
	"net/url"
	"regexp"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	oci_devops "github.com/oracle/oci-go-sdk/v65/devops"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DevopsProjectRepositorySettingResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createDevopsProjectRepositorySetting,
		Read:     readDevopsProjectRepositorySetting,
		Update:   updateDevopsProjectRepositorySetting,
		Delete:   deleteDevopsProjectRepositorySetting,
		Schema: map[string]*schema.Schema{
			// Required
			"project_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			// Optional
			"approval_rules": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"items": {
							Type:     schema.TypeList,
							Required: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required
									"min_approvals_count": {
										Type:     schema.TypeInt,
										Required: true,
									},
									"name": {
										Type:     schema.TypeString,
										Required: true,
									},

									// Optional
									"destination_branch": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"reviewers": {
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required
												"principal_id": {
													Type:     schema.TypeString,
													Required: true,
												},

												// Optional

												// Computed
												"principal_name": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"principal_state": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"principal_type": {
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},

									// Computed
								},
							},
						},

						// Optional

						// Computed
					},
				},
			},
			"merge_settings": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"allowed_merge_strategies": {
							Type:     schema.TypeList,
							Required: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"default_merge_strategy": {
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
	}
}

func createDevopsProjectRepositorySetting(d *schema.ResourceData, m interface{}) error {
	sync := &DevopsProjectRepositorySettingResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DevopsClient()

	return tfresource.CreateResource(d, sync)
}

func readDevopsProjectRepositorySetting(d *schema.ResourceData, m interface{}) error {
	sync := &DevopsProjectRepositorySettingResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DevopsClient()

	return tfresource.ReadResource(sync)
}

func updateDevopsProjectRepositorySetting(d *schema.ResourceData, m interface{}) error {
	sync := &DevopsProjectRepositorySettingResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DevopsClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteDevopsProjectRepositorySetting(d *schema.ResourceData, m interface{}) error {
	sync := &DevopsProjectRepositorySettingResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DevopsClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type DevopsProjectRepositorySettingResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_devops.DevopsClient
	Res                    *oci_devops.ProjectRepositorySettings
	DisableNotFoundRetries bool
}

func (s *DevopsProjectRepositorySettingResourceCrud) ID() string {
	return GetProjectRepositorySettingCompositeId(s.D.Get("project_id").(string))
}

func (s *DevopsProjectRepositorySettingResourceCrud) Create() error {
	request := oci_devops.UpdateProjectRepositorySettingsRequest{}

	if approvalRules, ok := s.D.GetOkExists("approval_rules"); ok {
		if tmpList := approvalRules.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "approval_rules", 0)
			tmp, err := s.mapToUpdateApprovalRuleDetailsCollection(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.ApprovalRules = &tmp
		}
	}

	if mergeSettings, ok := s.D.GetOkExists("merge_settings"); ok {
		if tmpList := mergeSettings.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "merge_settings", 0)
			tmp, err := s.mapToMergeSettings(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.MergeSettings = &tmp
		}
	}

	if projectId, ok := s.D.GetOkExists("project_id"); ok {
		tmp := projectId.(string)
		request.ProjectId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "devops")

	response, err := s.Client.UpdateProjectRepositorySettings(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.ProjectRepositorySettings
	return nil
}

func (s *DevopsProjectRepositorySettingResourceCrud) Get() error {
	request := oci_devops.GetProjectRepositorySettingsRequest{}

	if projectId, ok := s.D.GetOkExists("project_id"); ok {
		tmp := projectId.(string)
		request.ProjectId = &tmp
	}

	projectId, err := parseProjectRepositorySettingCompositeId(s.D.Id())
	if err == nil {
		request.ProjectId = &projectId
	} else {
		log.Printf("[WARN] Get() unable to parse current ID: %s", s.D.Id())
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "devops")

	response, err := s.Client.GetProjectRepositorySettings(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.ProjectRepositorySettings
	return nil
}

func (s *DevopsProjectRepositorySettingResourceCrud) Update() error {
	request := oci_devops.UpdateProjectRepositorySettingsRequest{}

	if approvalRules, ok := s.D.GetOkExists("approval_rules"); ok {
		if tmpList := approvalRules.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "approval_rules", 0)
			tmp, err := s.mapToUpdateApprovalRuleDetailsCollection(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.ApprovalRules = &tmp
		}
	}

	if mergeSettings, ok := s.D.GetOkExists("merge_settings"); ok {
		if tmpList := mergeSettings.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "merge_settings", 0)
			tmp, err := s.mapToMergeSettings(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.MergeSettings = &tmp
		}
	}

	if projectId, ok := s.D.GetOkExists("project_id"); ok {
		tmp := projectId.(string)
		request.ProjectId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "devops")

	response, err := s.Client.UpdateProjectRepositorySettings(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.ProjectRepositorySettings
	return nil
}

func (s *DevopsProjectRepositorySettingResourceCrud) Delete() error {
	request := oci_devops.DeleteProjectRepositorySettingsRequest{}

	if projectId, ok := s.D.GetOkExists("project_id"); ok {
		tmp := projectId.(string)
		request.ProjectId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "devops")

	_, err := s.Client.DeleteProjectRepositorySettings(context.Background(), request)
	return err
}

func (s *DevopsProjectRepositorySettingResourceCrud) SetData() error {

	projectId, err := parseProjectRepositorySettingCompositeId(s.D.Id())
	if err == nil {
		s.D.Set("project_id", &projectId)
	} else {
		log.Printf("[WARN] SetData() unable to parse current ID: %s", s.D.Id())
	}

	if s.Res.ApprovalRules != nil {
		s.D.Set("approval_rules", []interface{}{ApprovalRuleCollectionToMap(s.Res.ApprovalRules)})
	} else {
		s.D.Set("approval_rules", nil)
	}

	if s.Res.MergeSettings != nil {
		s.D.Set("merge_settings", []interface{}{MergeSettingsToMap(s.Res.MergeSettings)})
	} else {
		s.D.Set("merge_settings", nil)
	}

	return nil
}

func GetProjectRepositorySettingCompositeId(projectId string) string {
	projectId = url.PathEscape(projectId)
	compositeId := "projects/" + projectId + "/repositorySettings"
	return compositeId
}

func parseProjectRepositorySettingCompositeId(compositeId string) (projectId string, err error) {
	parts := strings.Split(compositeId, "/")
	match, _ := regexp.MatchString("projects/.*/repositorySettings", compositeId)
	if !match || len(parts) != 3 {
		err = fmt.Errorf("illegal compositeId %s encountered", compositeId)
		return
	}
	projectId, _ = url.PathUnescape(parts[1])

	return
}

func (s *DevopsProjectRepositorySettingResourceCrud) mapToMergeSettings(fieldKeyFormat string) (oci_devops.MergeSettings, error) {
	result := oci_devops.MergeSettings{}

	if allowedMergeStrategies, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "allowed_merge_strategies")); ok {
		interfaces := allowedMergeStrategies.([]interface{})
		tmp := make([]oci_devops.MergeStrategyEnum, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = oci_devops.MergeStrategyEnum(interfaces[i].(string))
			}
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "allowed_merge_strategies")) {
			result.AllowedMergeStrategies = tmp
		}
	}

	if defaultMergeStrategy, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "default_merge_strategy")); ok {
		result.DefaultMergeStrategy = oci_devops.MergeStrategyEnum(defaultMergeStrategy.(string))
	}

	return result, nil
}

func (s *DevopsProjectRepositorySettingResourceCrud) mapToUpdateApprovalRuleDetails(fieldKeyFormat string) (oci_devops.UpdateApprovalRuleDetails, error) {
	result := oci_devops.UpdateApprovalRuleDetails{}

	if destinationBranch, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "destination_branch")); ok {
		tmp := destinationBranch.(string)
		result.DestinationBranch = &tmp
	}

	if minApprovalsCount, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "min_approvals_count")); ok {
		tmp := minApprovalsCount.(int)
		result.MinApprovalsCount = &tmp
	}

	if name, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "name")); ok {
		tmp := name.(string)
		result.Name = &tmp
	}

	if reviewers, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "reviewers")); ok {
		interfaces := reviewers.([]interface{})
		tmp := make([]oci_devops.UpdateReviewerDetails, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "reviewers"), stateDataIndex)
			converted, err := s.mapToUpdateReviewerDetails(fieldKeyFormatNextLevel)
			if err != nil {
				return result, err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "reviewers")) {
			result.Reviewers = tmp
		}
	}

	return result, nil
}

/* Same function is defined in devops_repository_setting_resource.go, so commenting out it in this file.
func ApprovalRuleToMap(obj oci_devops.ApprovalRule) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.DestinationBranch != nil {
		result["destination_branch"] = string(*obj.DestinationBranch)
	}

	if obj.MinApprovalsCount != nil {
		result["min_approvals_count"] = int(*obj.MinApprovalsCount)
	}

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	reviewers := []interface{}{}
	for _, item := range obj.Reviewers {
		reviewers = append(reviewers, PrincipalDetailsToMap(item))
	}
	result["reviewers"] = reviewers

	return result
}
*/

func (s *DevopsProjectRepositorySettingResourceCrud) mapToUpdateApprovalRuleDetailsCollection(fieldKeyFormat string) (oci_devops.UpdateApprovalRuleDetailsCollection, error) {
	result := oci_devops.UpdateApprovalRuleDetailsCollection{}

	if items, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "items")); ok {
		interfaces := items.([]interface{})
		tmp := make([]oci_devops.UpdateApprovalRuleDetails, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "items"), stateDataIndex)
			converted, err := s.mapToUpdateApprovalRuleDetails(fieldKeyFormatNextLevel)
			if err != nil {
				return result, err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "items")) {
			result.Items = tmp
		}
	}

	return result, nil
}

/* Same function is defined in devops_repository_setting_resource.go, so commenting out it in this file.
func ApprovalRuleCollectionToMap(obj *oci_devops.ApprovalRuleCollection) map[string]interface{} {
	result := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range obj.Items {
		items = append(items, ApprovalRuleToMap(item))
	}
	result["items"] = items

	return result
}
*/

func (s *DevopsProjectRepositorySettingResourceCrud) mapToUpdateReviewerDetails(fieldKeyFormat string) (oci_devops.UpdateReviewerDetails, error) {
	result := oci_devops.UpdateReviewerDetails{}

	if principalId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "principal_id")); ok {
		tmp := principalId.(string)
		result.PrincipalId = &tmp
	}

	return result, nil
}

/* Same function is defined in devops_repository_setting_resource.go, so commenting out it in this file.
func PrincipalDetailsToMap(obj oci_devops.PrincipalDetails) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.PrincipalId != nil {
		result["principal_id"] = string(*obj.PrincipalId)
	}

	if obj.PrincipalName != nil {
		result["principal_name"] = string(*obj.PrincipalName)
	}

	result["principal_state"] = string(obj.PrincipalState)

	result["principal_type"] = string(obj.PrincipalType)

	return result
}
*/
