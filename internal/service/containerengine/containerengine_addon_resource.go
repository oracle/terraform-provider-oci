// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package containerengine

import (
	"context"
	"fmt"
	"log"
	"net/url"
	"regexp"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	oci_common "github.com/oracle/oci-go-sdk/v65/common"
	oci_containerengine "github.com/oracle/oci-go-sdk/v65/containerengine"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func ContainerengineAddonResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createContainerengineAddon,
		Read:     readContainerengineAddon,
		Update:   updateContainerengineAddon,
		Delete:   deleteContainerengineAddon,
		Schema: map[string]*schema.Schema{
			// Required
			"addon_name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"cluster_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"remove_addon_resources_on_delete": {
				Type:     schema.TypeBool,
				Required: true,
			},

			// Optional
			"configurations": {
				Type:             schema.TypeList,
				Optional:         true,
				Computed:         true,
				DiffSuppressFunc: tfresource.ListOfMapEqualIgnoreOrderSuppressDiff,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional
						"key": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"value": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},

						// Computed
					},
				},
			},
			"version": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			// Computed
			"addon_error": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"code": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"message": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"status": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"current_installed_version": {
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
		},
	}
}

func createContainerengineAddon(d *schema.ResourceData, m interface{}) error {
	sync := &ContainerengineAddonResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ContainerEngineClient()

	return tfresource.CreateResource(d, sync)
}

func readContainerengineAddon(d *schema.ResourceData, m interface{}) error {
	sync := &ContainerengineAddonResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ContainerEngineClient()

	return tfresource.ReadResource(sync)
}

func updateContainerengineAddon(d *schema.ResourceData, m interface{}) error {
	sync := &ContainerengineAddonResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ContainerEngineClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteContainerengineAddon(d *schema.ResourceData, m interface{}) error {
	sync := &ContainerengineAddonResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ContainerEngineClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type ContainerengineAddonResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_containerengine.ContainerEngineClient
	Res                    *oci_containerengine.Addon
	DisableNotFoundRetries bool
}

func (s *ContainerengineAddonResourceCrud) ID() string {
	return GetAddonCompositeId(s.D.Get("addon_name").(string), s.D.Get("cluster_id").(string))
}

func (s *ContainerengineAddonResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_containerengine.AddonLifecycleStateCreating),
	}
}

func (s *ContainerengineAddonResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_containerengine.AddonLifecycleStateActive),
		string(oci_containerengine.AddonLifecycleStateNeedsAttention),
	}
}

func (s *ContainerengineAddonResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_containerengine.AddonLifecycleStateDeleting),
	}
}

func (s *ContainerengineAddonResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_containerengine.AddonLifecycleStateDeleted),
	}
}

func (s *ContainerengineAddonResourceCrud) Create() error {
	request := oci_containerengine.InstallAddonRequest{}

	if name, ok := s.D.GetOkExists("addon_name"); ok {
		tmp := name.(string)
		request.AddonName = &tmp
	}

	if clusterId, ok := s.D.GetOkExists("cluster_id"); ok {
		tmp := clusterId.(string)
		request.ClusterId = &tmp
	}

	if configurations, ok := s.D.GetOkExists("configurations"); ok {
		interfaces := configurations.([]interface{})
		tmp := make([]oci_containerengine.AddonConfiguration, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "configurations", stateDataIndex)
			converted, err := s.mapToAddonConfiguration(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("configurations") {
			request.Configurations = tmp
		}
	}

	if version, ok := s.D.GetOkExists("version"); ok {
		tmp := version.(string)
		request.Version = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "containerengine")

	response, err := s.Client.InstallAddon(context.Background(), request)
	if err != nil {
		return err
	}

	s.D.SetId(s.ID())
	workId := response.OpcWorkRequestId
	err = s.getAddonFromWorkRequest(workId, s.D.Timeout(schema.TimeoutCreate))

	if err != nil {
		// Try to delete the addon
		log.Printf("[DEBUG] creation failed, attempting to delete the addon: %v\n", request.AddonName)
		disableAddonRequest := oci_containerengine.DisableAddonRequest{
			AddonName: request.AddonName,
			ClusterId: request.ClusterId,
			RequestMetadata: oci_common.RequestMetadata{
				RetryPolicy: tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "containerengine"),
			},
		}

		if isRemoveExistingAddon, ok := s.D.GetOkExists("remove_addon_resources_on_delete"); ok {
			tmp := isRemoveExistingAddon.(bool)
			disableAddonRequest.IsRemoveExistingAddOn = &tmp
		}

		_, deleteErr := s.Client.DisableAddon(context.Background(), disableAddonRequest)
		if deleteErr != nil {
			return deleteErr
		}
	}

	return nil
}

func (s *ContainerengineAddonResourceCrud) getAddonFromWorkRequest(workId *string, timeout time.Duration) error {

	// Wait until it finishes
	err := addonWaitForWorkRequest(workId, "addon", timeout, s.DisableNotFoundRetries, s.Client)

	if err != nil {
		return err
	}

	return s.Get()
}

func addonWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
	startTime := time.Now()
	stopTime := startTime.Add(timeout)
	return func(response oci_common.OCIOperationResponse) bool {

		// Stop after timeout has elapsed
		if time.Now().After(stopTime) {
			return false
		}

		// Make sure we stop on default rules
		if tfresource.ShouldRetry(response, false, "containerengine", startTime) {
			return true
		}

		// Only stop if the time Finished is set
		if workRequestResponse, ok := response.Response.(oci_containerengine.GetWorkRequestResponse); ok {
			return workRequestResponse.TimeFinished == nil
		}
		return false
	}
}

func addonWaitForWorkRequest(wId *string, entityType string,
	timeout time.Duration, disableFoundRetries bool, client *oci_containerengine.ContainerEngineClient) error {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "containerengine")
	retryPolicy.ShouldRetryOperation = addonWorkRequestShouldRetryFunc(timeout)

	response := oci_containerengine.GetWorkRequestResponse{}
	stateConf := &resource.StateChangeConf{
		Pending: []string{
			string(oci_containerengine.WorkRequestStatusInProgress),
			string(oci_containerengine.WorkRequestStatusAccepted),
			string(oci_containerengine.WorkRequestStatusCanceling),
		},
		Target: []string{
			string(oci_containerengine.WorkRequestStatusSucceeded),
			string(oci_containerengine.WorkRequestStatusFailed),
			string(oci_containerengine.WorkRequestStatusCanceled),
		},
		Refresh: func() (interface{}, string, error) {
			var err error
			response, err = client.GetWorkRequest(context.Background(),
				oci_containerengine.GetWorkRequestRequest{
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
		return e
	}

	// The workrequest may have failed, check for errors if identifier is not found or work failed or got cancelled
	if response.Status == oci_containerengine.WorkRequestStatusFailed || response.Status == oci_containerengine.WorkRequestStatusCanceled {
		return getErrorFromContainerengineAddonWorkRequest(client, wId, response.CompartmentId, retryPolicy, entityType)
	}

	return nil
}

func getErrorFromContainerengineAddonWorkRequest(client *oci_containerengine.ContainerEngineClient, workId *string, compartmentId *string, retryPolicy *oci_common.RetryPolicy, entityType string) error {
	response, err := client.ListWorkRequestErrors(context.Background(),
		oci_containerengine.ListWorkRequestErrorsRequest{
			WorkRequestId: workId,
			CompartmentId: compartmentId,
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

	workRequestErr := fmt.Errorf("work request did not succeed, workId: %s, entity: %s. Message: %s", *workId, entityType, errorMessage)

	return workRequestErr
}

func (s *ContainerengineAddonResourceCrud) Get() error {
	request := oci_containerengine.GetAddonRequest{}

	if addonName, ok := s.D.GetOkExists("addon_name"); ok {
		tmp := addonName.(string)
		request.AddonName = &tmp
	}

	if clusterId, ok := s.D.GetOkExists("cluster_id"); ok {
		tmp := clusterId.(string)
		request.ClusterId = &tmp
	}

	addonName, clusterId, err := parseAddonCompositeId(s.D.Id())
	if err == nil {
		request.AddonName = &addonName
		request.ClusterId = &clusterId
	} else {
		log.Printf("[WARN] Get() unable to parse current ID: %s", s.D.Id())
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "containerengine")

	response, err := s.Client.GetAddon(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Addon
	return nil
}

func (s *ContainerengineAddonResourceCrud) Update() error {
	request := oci_containerengine.UpdateAddonRequest{}

	if addonName, ok := s.D.GetOkExists("addon_name"); ok {
		tmp := addonName.(string)
		request.AddonName = &tmp
	}

	if clusterId, ok := s.D.GetOkExists("cluster_id"); ok {
		tmp := clusterId.(string)
		request.ClusterId = &tmp
	}

	if configurations, ok := s.D.GetOkExists("configurations"); ok {
		interfaces := configurations.([]interface{})
		tmp := make([]oci_containerengine.AddonConfiguration, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "configurations", stateDataIndex)
			converted, err := s.mapToAddonConfiguration(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("configurations") {
			request.Configurations = tmp
		}
	}

	if version, ok := s.D.GetOkExists("version"); ok {
		tmp := version.(string)
		request.Version = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "containerengine")

	response, err := s.Client.UpdateAddon(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getAddonFromWorkRequest(workId, s.D.Timeout(schema.TimeoutUpdate))
}

func (s *ContainerengineAddonResourceCrud) Delete() error {
	request := oci_containerengine.DisableAddonRequest{}

	if addonName, ok := s.D.GetOkExists("addon_name"); ok {
		tmp := addonName.(string)
		request.AddonName = &tmp
	}

	if clusterId, ok := s.D.GetOkExists("cluster_id"); ok {
		tmp := clusterId.(string)
		request.ClusterId = &tmp
	}

	if isRemoveExistingAddOn, ok := s.D.GetOkExists("remove_addon_resources_on_delete"); ok {
		tmp := isRemoveExistingAddOn.(bool)
		request.IsRemoveExistingAddOn = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "containerengine")

	response, err := s.Client.DisableAddon(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	// Wait until it finishes
	delWorkRequestErr := addonWaitForWorkRequest(workId, "addon", s.D.Timeout(schema.TimeoutDelete), s.DisableNotFoundRetries, s.Client)
	return delWorkRequestErr
}

func (s *ContainerengineAddonResourceCrud) SetData() error {

	addonName, clusterId, err := parseAddonCompositeId(s.D.Id())
	if err == nil {
		s.D.Set("addon_name", &addonName)
		s.D.Set("cluster_id", &clusterId)
	} else {
		log.Printf("[WARN] SetData() unable to parse current ID: %s", s.D.Id())
	}

	if s.Res.AddonError != nil {
		s.D.Set("addon_error", []interface{}{AddonErrorToMap(s.Res.AddonError)})
	} else {
		s.D.Set("addon_error", nil)
	}

	configurations := []interface{}{}
	for _, item := range s.Res.Configurations {
		configurations = append(configurations, AddonConfigurationToMap(item))
	}
	s.D.Set("configurations", configurations)

	if s.Res.CurrentInstalledVersion != nil {
		s.D.Set("current_installed_version", *s.Res.CurrentInstalledVersion)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.Version != nil {
		s.D.Set("version", *s.Res.Version)
	}

	return nil
}

func GetAddonCompositeId(addonName string, clusterId string) string {
	addonName = url.PathEscape(addonName)
	clusterId = url.PathEscape(clusterId)
	compositeId := "clusters/" + clusterId + "/addons/" + addonName
	return compositeId
}

func parseAddonCompositeId(compositeId string) (addonName string, clusterId string, err error) {
	parts := strings.Split(compositeId, "/")
	match, _ := regexp.MatchString("clusters/.*/addons/.*", compositeId)
	if !match || len(parts) != 4 {
		err = fmt.Errorf("illegal compositeId %s encountered", compositeId)
		return
	}
	clusterId, _ = url.PathUnescape(parts[1])
	addonName, _ = url.PathUnescape(parts[3])

	return
}

func (s *ContainerengineAddonResourceCrud) mapToAddonConfiguration(fieldKeyFormat string) (oci_containerengine.AddonConfiguration, error) {
	result := oci_containerengine.AddonConfiguration{}

	if key, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "key")); ok {
		tmp := key.(string)
		result.Key = &tmp
	}

	if value, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "value")); ok {
		tmp := value.(string)
		result.Value = &tmp
	}

	return result, nil
}

func AddonConfigurationToMap(obj oci_containerengine.AddonConfiguration) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Key != nil {
		result["key"] = string(*obj.Key)
	}

	if obj.Value != nil {
		result["value"] = string(*obj.Value)
	}

	return result
}

func AddonErrorToMap(obj *oci_containerengine.AddonError) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Code != nil {
		result["code"] = string(*obj.Code)
	}

	if obj.Message != nil {
		result["message"] = string(*obj.Message)
	}

	if obj.Status != nil {
		result["status"] = string(*obj.Status)
	}

	return result
}
