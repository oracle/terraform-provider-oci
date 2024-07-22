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

func DevopsRepositorySettingResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createDevopsRepositorySetting,
		Read:     readDevopsRepositorySetting,
		Update:   updateDevopsRepositorySetting,
		Delete:   deleteDevopsRepositorySetting,
		Schema: map[string]*schema.Schema{
			// Required
			"repository_id": {
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
			"merge_checks": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"last_build_succeeded": {
							Type:     schema.TypeString,
							Required: true,
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

func createDevopsRepositorySetting(d *schema.ResourceData, m interface{}) error {
	sync := &DevopsRepositorySettingResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DevopsClient()

	return tfresource.CreateResource(d, sync)
}

func readDevopsRepositorySetting(d *schema.ResourceData, m interface{}) error {
	sync := &DevopsRepositorySettingResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DevopsClient()

	return tfresource.ReadResource(sync)
}

func updateDevopsRepositorySetting(d *schema.ResourceData, m interface{}) error {
	sync := &DevopsRepositorySettingResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DevopsClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteDevopsRepositorySetting(d *schema.ResourceData, m interface{}) error {
	sync := &DevopsRepositorySettingResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DevopsClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type DevopsRepositorySettingResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_devops.DevopsClient
	Res                    *oci_devops.RepositorySettings
	DisableNotFoundRetries bool
}

func (s *DevopsRepositorySettingResourceCrud) ID() string {
	return GetRepositorySettingCompositeId(s.D.Get("repository_id").(string))
}

func (s *DevopsRepositorySettingResourceCrud) Create() error {
	request := oci_devops.UpdateRepositorySettingsRequest{}

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

	if mergeChecks, ok := s.D.GetOkExists("merge_checks"); ok {
		if tmpList := mergeChecks.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "merge_checks", 0)
			tmp, err := s.mapToMergeChecks(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.MergeChecks = &tmp
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

	if repositoryId, ok := s.D.GetOkExists("repository_id"); ok {
		tmp := repositoryId.(string)
		request.RepositoryId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "devops")

	response, err := s.Client.UpdateRepositorySettings(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.RepositorySettings
	return nil
}

func (s *DevopsRepositorySettingResourceCrud) Get() error {
	request := oci_devops.GetRepositorySettingsRequest{}

	if repositoryId, ok := s.D.GetOkExists("repository_id"); ok {
		tmp := repositoryId.(string)
		request.RepositoryId = &tmp
	}

	repositoryId, err := parseRepositorySettingCompositeId(s.D.Id())
	if err == nil {
		request.RepositoryId = &repositoryId
	} else {
		log.Printf("[WARN] Get() unable to parse current ID: %s", s.D.Id())
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "devops")

	response, err := s.Client.GetRepositorySettings(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.RepositorySettings
	return nil
}

func (s *DevopsRepositorySettingResourceCrud) Update() error {
	request := oci_devops.UpdateRepositorySettingsRequest{}

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

	if mergeChecks, ok := s.D.GetOkExists("merge_checks"); ok {
		if tmpList := mergeChecks.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "merge_checks", 0)
			tmp, err := s.mapToMergeChecks(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.MergeChecks = &tmp
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

	if repositoryId, ok := s.D.GetOkExists("repository_id"); ok {
		tmp := repositoryId.(string)
		request.RepositoryId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "devops")

	response, err := s.Client.UpdateRepositorySettings(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.RepositorySettings
	return nil
}

func (s *DevopsRepositorySettingResourceCrud) Delete() error {
	request := oci_devops.DeleteRepositorySettingsRequest{}

	if repositoryId, ok := s.D.GetOkExists("repository_id"); ok {
		tmp := repositoryId.(string)
		request.RepositoryId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "devops")

	_, err := s.Client.DeleteRepositorySettings(context.Background(), request)
	return err
}

func (s *DevopsRepositorySettingResourceCrud) SetData() error {

	repositoryId, err := parseRepositorySettingCompositeId(s.D.Id())
	if err == nil {
		s.D.Set("repository_id", &repositoryId)
	} else {
		log.Printf("[WARN] SetData() unable to parse current ID: %s", s.D.Id())
	}

	if s.Res.ApprovalRules != nil {
		s.D.Set("approval_rules", []interface{}{ApprovalRuleCollectionToMap(s.Res.ApprovalRules)})
	} else {
		s.D.Set("approval_rules", nil)
	}

	if s.Res.MergeChecks != nil {
		s.D.Set("merge_checks", []interface{}{MergeChecksToMap(s.Res.MergeChecks)})
	} else {
		s.D.Set("merge_checks", nil)
	}

	if s.Res.MergeSettings != nil {
		s.D.Set("merge_settings", []interface{}{MergeSettingsToMap(s.Res.MergeSettings)})
	} else {
		s.D.Set("merge_settings", nil)
	}

	return nil
}

func GetRepositorySettingCompositeId(repositoryId string) string {
	repositoryId = url.PathEscape(repositoryId)
	compositeId := "repositories/" + repositoryId + "/repositorySettings"
	return compositeId
}

func parseRepositorySettingCompositeId(compositeId string) (repositoryId string, err error) {
	parts := strings.Split(compositeId, "/")
	match, _ := regexp.MatchString("repositories/.*/repositorySettings", compositeId)
	if !match || len(parts) != 3 {
		err = fmt.Errorf("illegal compositeId %s encountered", compositeId)
		return
	}
	repositoryId, _ = url.PathUnescape(parts[1])

	return
}

func (s *DevopsRepositorySettingResourceCrud) mapToMergeChecks(fieldKeyFormat string) (oci_devops.MergeChecks, error) {
	result := oci_devops.MergeChecks{}

	if lastBuildSucceeded, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "last_build_succeeded")); ok {
		result.LastBuildSucceeded = oci_devops.MergeCheckSettingsValueEnum(lastBuildSucceeded.(string))
	}

	return result, nil
}

func MergeChecksToMap(obj *oci_devops.MergeChecks) map[string]interface{} {
	result := map[string]interface{}{}

	result["last_build_succeeded"] = string(obj.LastBuildSucceeded)

	return result
}

func (s *DevopsRepositorySettingResourceCrud) mapToMergeSettings(fieldKeyFormat string) (oci_devops.MergeSettings, error) {
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

func MergeSettingsToMap(obj *oci_devops.MergeSettings) map[string]interface{} {
	result := map[string]interface{}{}

	result["allowed_merge_strategies"] = obj.AllowedMergeStrategies

	result["default_merge_strategy"] = string(obj.DefaultMergeStrategy)

	return result
}

func (s *DevopsRepositorySettingResourceCrud) mapToUpdateApprovalRuleDetails(fieldKeyFormat string) (oci_devops.UpdateApprovalRuleDetails, error) {
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

func (s *DevopsRepositorySettingResourceCrud) mapToUpdateApprovalRuleDetailsCollection(fieldKeyFormat string) (oci_devops.UpdateApprovalRuleDetailsCollection, error) {
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

func ApprovalRuleCollectionToMap(obj *oci_devops.ApprovalRuleCollection) map[string]interface{} {
	result := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range obj.Items {
		items = append(items, ApprovalRuleToMap(item))
	}
	result["items"] = items

	return result
}

func (s *DevopsRepositorySettingResourceCrud) mapToUpdateReviewerDetails(fieldKeyFormat string) (oci_devops.UpdateReviewerDetails, error) {
	result := oci_devops.UpdateReviewerDetails{}

	if principalId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "principal_id")); ok {
		tmp := principalId.(string)
		result.PrincipalId = &tmp
	}

	return result, nil
}

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
