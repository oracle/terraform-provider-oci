// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package managed_kafka

import (
	"context"
	"fmt"
	"log"
	"net/url"
	"regexp"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/retry"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	oci_common "github.com/oracle/oci-go-sdk/v65/common"
	oci_managed_kafka "github.com/oracle/oci-go-sdk/v65/managedkafka"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func ManagedKafkaKafkaClusterAddonResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts:      tfresource.DefaultTimeout,
		CreateContext: createManagedKafkaKafkaClusterAddonWithContext,
		ReadContext:   readManagedKafkaKafkaClusterAddonWithContext,
		UpdateContext: updateManagedKafkaKafkaClusterAddonWithContext,
		DeleteContext: deleteManagedKafkaKafkaClusterAddonWithContext,
		Schema: map[string]*schema.Schema{
			// Required
			"addon_type": {
				Type:             schema.TypeString,
				Required:         true,
				DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
				ValidateFunc: validation.StringInSlice([]string{
					"PUBLICCONNECTIVITY",
				}, true),
			},
			"authentication_mechanism": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"kafka_cluster_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"network_cidrs": {
				Type:     schema.TypeList,
				Required: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},

			// Optional
			"description": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			// Computed
			"bootstrap_url": {
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
	}
}

func createManagedKafkaKafkaClusterAddonWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &ManagedKafkaKafkaClusterAddonResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).KafkaClusterClient()

	return tfresource.HandleDiagError(m, tfresource.CreateResourceWithContext(ctx, d, sync))
}

func readManagedKafkaKafkaClusterAddonWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &ManagedKafkaKafkaClusterAddonResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).KafkaClusterClient()

	return tfresource.HandleDiagError(m, tfresource.ReadResourceWithContext(ctx, sync))
}

func updateManagedKafkaKafkaClusterAddonWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &ManagedKafkaKafkaClusterAddonResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).KafkaClusterClient()

	return tfresource.HandleDiagError(m, tfresource.UpdateResourceWithContext(ctx, d, sync))
}

func deleteManagedKafkaKafkaClusterAddonWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &ManagedKafkaKafkaClusterAddonResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).KafkaClusterClient()
	sync.DisableNotFoundRetries = true

	return tfresource.HandleDiagError(m, tfresource.DeleteResourceWithContext(ctx, d, sync))
}

type ManagedKafkaKafkaClusterAddonResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_managed_kafka.KafkaClusterClient
	Res                    *oci_managed_kafka.KafkaClusterAddon
	DisableNotFoundRetries bool
}

func (s *ManagedKafkaKafkaClusterAddonResourceCrud) ID() string {
	addonName, _ := s.D.Get("name").(string)
	kafkaClusterId, _ := s.D.Get("kafka_cluster_id").(string)
	return GetKafkaClusterAddonCompositeId(addonName, kafkaClusterId)
}

func (s *ManagedKafkaKafkaClusterAddonResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_managed_kafka.KafkaClusterAddonLifecycleStateCreating),
	}
}

func (s *ManagedKafkaKafkaClusterAddonResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_managed_kafka.KafkaClusterAddonLifecycleStateActive),
	}
}

func (s *ManagedKafkaKafkaClusterAddonResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_managed_kafka.KafkaClusterAddonLifecycleStateDeleting),
	}
}

func (s *ManagedKafkaKafkaClusterAddonResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_managed_kafka.KafkaClusterAddonLifecycleStateDeleted),
	}
}

func (s *ManagedKafkaKafkaClusterAddonResourceCrud) CreateWithContext(ctx context.Context) error {
	request := oci_managed_kafka.InstallAddonRequest{}
	err := s.populateTopLevelPolymorphicInstallAddonRequest(&request)
	if err != nil {
		return err
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "managed_kafka")

	response, err := s.Client.InstallAddon(ctx, request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getKafkaClusterAddonFromWorkRequest(ctx, workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "managed_kafka"), oci_managed_kafka.ActionTypeCreated, s.D.Timeout(schema.TimeoutCreate))
}

func (s *ManagedKafkaKafkaClusterAddonResourceCrud) getKafkaClusterAddonFromWorkRequest(ctx context.Context, workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_managed_kafka.ActionTypeEnum, timeout time.Duration) error {

	// Wait until it finishes
	kafkaClusterAddonId, err := kafkaClusterAddonWaitForWorkRequest(ctx, workId, "kafkacluster",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.Client)

	if err != nil {
		// Try to cancel the work request
		log.Printf("[DEBUG] creation failed, attempting to cancel the workrequest: %v for identifier: %v\n", workId, kafkaClusterAddonId)
		_, cancelErr := s.Client.CancelWorkRequest(ctx,
			oci_managed_kafka.CancelWorkRequestRequest{
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
	addonName, _ := s.D.Get("name").(string)

	kafkaClusterId, _ := s.D.Get("kafka_cluster_id").(string)
	s.D.SetId(GetKafkaClusterAddonCompositeId(addonName, kafkaClusterId))

	return s.GetWithContext(ctx)
}

func kafkaClusterAddonWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
	startTime := time.Now()
	stopTime := startTime.Add(timeout)
	return func(response oci_common.OCIOperationResponse) bool {

		// Stop after timeout has elapsed
		if time.Now().After(stopTime) {
			return false
		}

		// Make sure we stop on default rules
		if tfresource.ShouldRetry(response, false, "managed_kafka", startTime) {
			return true
		}

		// Only stop if the time Finished is set
		if workRequestResponse, ok := response.Response.(oci_managed_kafka.GetWorkRequestResponse); ok {
			return workRequestResponse.TimeFinished == nil
		}
		return false
	}
}

func kafkaClusterAddonWaitForWorkRequest(ctx context.Context, wId *string, entityType string, action oci_managed_kafka.ActionTypeEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_managed_kafka.KafkaClusterClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "managed_kafka")
	retryPolicy.ShouldRetryOperation = kafkaClusterAddonWorkRequestShouldRetryFunc(timeout)

	response := oci_managed_kafka.GetWorkRequestResponse{}
	stateConf := &retry.StateChangeConf{
		Pending: []string{
			string(oci_managed_kafka.OperationStatusInProgress),
			string(oci_managed_kafka.OperationStatusAccepted),
			string(oci_managed_kafka.OperationStatusCanceling),
		},
		Target: []string{
			string(oci_managed_kafka.OperationStatusSucceeded),
			string(oci_managed_kafka.OperationStatusFailed),
			string(oci_managed_kafka.OperationStatusCanceled),
		},
		Refresh: func() (interface{}, string, error) {
			var err error
			response, err = client.GetWorkRequest(ctx,
				oci_managed_kafka.GetWorkRequestRequest{
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
	if identifier == nil || response.Status == oci_managed_kafka.OperationStatusFailed || response.Status == oci_managed_kafka.OperationStatusCanceled {
		return nil, getErrorFromManagedKafkaKafkaClusterAddonWorkRequest(ctx, client, wId, retryPolicy, entityType, action)
	}

	return identifier, nil
}

func getErrorFromManagedKafkaKafkaClusterAddonWorkRequest(ctx context.Context, client *oci_managed_kafka.KafkaClusterClient, workId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_managed_kafka.ActionTypeEnum) error {
	response, err := client.ListWorkRequestErrors(ctx,
		oci_managed_kafka.ListWorkRequestErrorsRequest{
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

func (s *ManagedKafkaKafkaClusterAddonResourceCrud) GetWithContext(ctx context.Context) error {
	request := oci_managed_kafka.GetAddonRequest{}

	if addonName, ok := s.D.GetOkExists("name"); ok {
		tmp := addonName.(string)
		request.AddonName = &tmp
	}

	if kafkaClusterId, ok := s.D.GetOkExists("kafka_cluster_id"); ok {
		tmp := kafkaClusterId.(string)
		request.KafkaClusterId = &tmp
	}

	addonName, kafkaClusterId, err := parseKafkaClusterAddonCompositeId(s.D.Id())
	if err == nil {
		request.AddonName = &addonName
		request.KafkaClusterId = &kafkaClusterId
	} else {
		log.Printf("[WARN] Get() unable to parse current ID: %s", s.D.Id())
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "managed_kafka")

	response, err := s.Client.GetAddon(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response.KafkaClusterAddon
	return nil
}

func (s *ManagedKafkaKafkaClusterAddonResourceCrud) UpdateWithContext(ctx context.Context) error {
	request := oci_managed_kafka.UpdateAddonRequest{}
	err := s.populateTopLevelPolymorphicUpdateAddonRequest(&request)
	if err != nil {
		return err
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "managed_kafka")

	response, err := s.Client.UpdateAddon(ctx, request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getKafkaClusterAddonFromWorkRequest(ctx, workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "managed_kafka"), oci_managed_kafka.ActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}

func (s *ManagedKafkaKafkaClusterAddonResourceCrud) DeleteWithContext(ctx context.Context) error {
	request := oci_managed_kafka.UninstallAddonRequest{}

	if addonName, ok := s.D.GetOkExists("name"); ok {
		tmp := addonName.(string)
		request.AddonName = &tmp
	}

	if kafkaClusterId, ok := s.D.GetOkExists("kafka_cluster_id"); ok {
		tmp := kafkaClusterId.(string)
		request.KafkaClusterId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "managed_kafka")

	_, err := s.Client.UninstallAddon(ctx, request)
	return err
}

func (s *ManagedKafkaKafkaClusterAddonResourceCrud) SetData() error {
	addonName, kafkaClusterId, err := parseKafkaClusterAddonCompositeId(s.D.Id())
	if err == nil {
		s.D.Set("name", &addonName)
		s.D.Set("kafka_cluster_id", &kafkaClusterId)
	} else {
		log.Printf("[WARN] SetData() unable to parse current ID: %s", s.D.Id())
	}

	switch v := (*s.Res).(type) {
	case oci_managed_kafka.PublicConnectivityAddon:
		s.D.Set("addon_type", "PUBLICCONNECTIVITY")

		s.D.Set("authentication_mechanism", v.AuthenticationMechanism)

		if v.BootstrapUrl != nil {
			s.D.Set("bootstrap_url", *v.BootstrapUrl)
		}

		s.D.Set("network_cidrs", v.NetworkCidrs)

		if v.Description != nil {
			s.D.Set("description", *v.Description)
		}

		if v.Name != nil {
			s.D.Set("name", *v.Name)
		}

		s.D.Set("state", v.LifecycleState)

		if v.TimeCreated != nil {
			s.D.Set("time_created", v.TimeCreated.String())
		}

		if v.TimeUpdated != nil {
			s.D.Set("time_updated", v.TimeUpdated.String())
		}
	default:
		log.Printf("[WARN] Received 'addon_type' of unknown type %v", *s.Res)
		return nil
	}
	return nil
}

func GetKafkaClusterAddonCompositeId(addonName string, kafkaClusterId string) string {
	addonName = url.PathEscape(addonName)
	kafkaClusterId = url.PathEscape(kafkaClusterId)
	compositeId := "kafkaClusters/" + kafkaClusterId + "/addons/" + addonName
	return compositeId
}

func parseKafkaClusterAddonCompositeId(compositeId string) (addonName string, kafkaClusterId string, err error) {
	parts := strings.Split(compositeId, "/")
	match, _ := regexp.MatchString("kafkaClusters/.*/addons/.*", compositeId)
	if !match || len(parts) != 4 {
		err = fmt.Errorf("illegal compositeId %s encountered", compositeId)
		return
	}
	kafkaClusterId, _ = url.PathUnescape(parts[1])
	addonName, _ = url.PathUnescape(parts[3])
	return
}

func AddonSummaryToMap(obj oci_managed_kafka.AddonSummary) map[string]interface{} {
	result := map[string]interface{}{}

	result["addon_type"] = string(obj.AddonType)

	if obj.LifecycleDetails != nil {
		result["lifecycle_details"] = string(*obj.LifecycleDetails)
	}

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
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

func (s *ManagedKafkaKafkaClusterAddonResourceCrud) populateTopLevelPolymorphicInstallAddonRequest(request *oci_managed_kafka.InstallAddonRequest) error {
	//discriminator
	addonTypeRaw, ok := s.D.GetOkExists("addon_type")
	var addonType string
	if ok {
		addonType = addonTypeRaw.(string)
	} else {
		addonType = "" // default value
	}
	switch strings.ToLower(addonType) {
	case strings.ToLower("PUBLICCONNECTIVITY"):
		details := oci_managed_kafka.InstallPublicConnectivityAddonDetails{}
		if authenticationMechanism, ok := s.D.GetOkExists("authentication_mechanism"); ok {
			details.AuthenticationMechanism = oci_managed_kafka.AuthenticationMechanismEnum(authenticationMechanism.(string))
		}
		if networkCidrs, ok := s.D.GetOkExists("network_cidrs"); ok {
			interfaces := networkCidrs.([]interface{})
			tmp := make([]string, len(interfaces))
			for i := range interfaces {
				if interfaces[i] != nil {
					tmp[i] = interfaces[i].(string)
				}
			}
			if len(tmp) != 0 || s.D.HasChange("network_cidrs") {
				details.NetworkCidrs = tmp
			}
		}
		if description, ok := s.D.GetOkExists("description"); ok {
			tmp := description.(string)
			details.Description = &tmp
		}
		if name, ok := s.D.GetOkExists("name"); ok {
			tmp := name.(string)
			details.Name = &tmp
		}
		if kafkaClusterId, ok := s.D.GetOkExists("kafka_cluster_id"); ok {
			tmp := kafkaClusterId.(string)
			request.KafkaClusterId = &tmp
		}
		request.InstallAddonDetails = details
	default:
		return fmt.Errorf("unknown addon_type '%v' was specified", addonType)
	}
	return nil
}

func (s *ManagedKafkaKafkaClusterAddonResourceCrud) populateTopLevelPolymorphicUpdateAddonRequest(request *oci_managed_kafka.UpdateAddonRequest) error {
	//discriminator
	addonTypeRaw, ok := s.D.GetOkExists("addon_type")
	var addonType string
	if ok {
		addonType = addonTypeRaw.(string)
	} else {
		addonType = "" // default value
	}
	switch strings.ToLower(addonType) {
	case strings.ToLower("PUBLICCONNECTIVITY"):
		details := oci_managed_kafka.UpdatePublicConnectivityAddonDetails{}
		if networkCidrs, ok := s.D.GetOkExists("network_cidrs"); ok {
			interfaces := networkCidrs.([]interface{})
			tmp := make([]string, len(interfaces))
			for i := range interfaces {
				if interfaces[i] != nil {
					tmp[i] = interfaces[i].(string)
				}
			}
			if len(tmp) != 0 || s.D.HasChange("network_cidrs") {
				details.NetworkCidrs = tmp
			}
		}
		if addonName, ok := s.D.GetOkExists("name"); ok {
			tmp := addonName.(string)
			request.AddonName = &tmp
		}
		if description, ok := s.D.GetOkExists("description"); ok {
			tmp := description.(string)
			details.Description = &tmp
		}
		if kafkaClusterId, ok := s.D.GetOkExists("kafka_cluster_id"); ok {
			tmp := kafkaClusterId.(string)
			request.KafkaClusterId = &tmp
		}
		request.UpdateAddonDetails = details
	default:
		return fmt.Errorf("unknown addon_type '%v' was specified", addonType)
	}
	return nil
}
