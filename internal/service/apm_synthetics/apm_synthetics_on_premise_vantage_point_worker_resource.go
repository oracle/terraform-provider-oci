// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package apm_synthetics

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/url"
	"regexp"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	oci_apm_synthetics "github.com/oracle/oci-go-sdk/v65/apmsynthetics"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func ApmSyntheticsOnPremiseVantagePointWorkerResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createApmSyntheticsOnPremiseVantagePointWorker,
		Read:     readApmSyntheticsOnPremiseVantagePointWorker,
		Update:   updateApmSyntheticsOnPremiseVantagePointWorker,
		Delete:   deleteApmSyntheticsOnPremiseVantagePointWorker,
		Schema: map[string]*schema.Schema{
			// Required
			"apm_domain_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"on_premise_vantage_point_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"resource_principal_token_public_key": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"version": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			// Optional
			"configuration_details": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
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
			"priority": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"status": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"worker_type": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},

			// Computed
			"display_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"geo_info": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"identity_info": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"apm_short_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"collector_end_point": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"region_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"monitor_list": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"display_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"is_run_now": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"monitor_type": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"time_assigned": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"opvp_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"opvp_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"runtime_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_created": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_last_sync_up": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_updated": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"version_details": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"latest_version": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"min_supported_version": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"version": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func createApmSyntheticsOnPremiseVantagePointWorker(d *schema.ResourceData, m interface{}) error {
	sync := &ApmSyntheticsOnPremiseVantagePointWorkerResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ApmSyntheticClient()

	return tfresource.CreateResource(d, sync)
}

func readApmSyntheticsOnPremiseVantagePointWorker(d *schema.ResourceData, m interface{}) error {
	sync := &ApmSyntheticsOnPremiseVantagePointWorkerResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ApmSyntheticClient()

	return tfresource.ReadResource(sync)
}

func updateApmSyntheticsOnPremiseVantagePointWorker(d *schema.ResourceData, m interface{}) error {
	sync := &ApmSyntheticsOnPremiseVantagePointWorkerResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ApmSyntheticClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteApmSyntheticsOnPremiseVantagePointWorker(d *schema.ResourceData, m interface{}) error {
	sync := &ApmSyntheticsOnPremiseVantagePointWorkerResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ApmSyntheticClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type ApmSyntheticsOnPremiseVantagePointWorkerResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_apm_synthetics.ApmSyntheticClient
	Res                    *oci_apm_synthetics.Worker
	DisableNotFoundRetries bool
}

func (s *ApmSyntheticsOnPremiseVantagePointWorkerResourceCrud) ID() string {
	return GetOnPremiseVantagePointWorkerCompositeId(s.D.Get("on_premise_vantage_point_id").(string), *s.Res.Id, s.D.Get("apm_domain_id").(string))
}

func (s *ApmSyntheticsOnPremiseVantagePointWorkerResourceCrud) Create() error {
	request := oci_apm_synthetics.CreateWorkerRequest{}

	if apmDomainId, ok := s.D.GetOkExists("apm_domain_id"); ok {
		tmp := apmDomainId.(string)
		request.ApmDomainId = &tmp
	}

	if configurationDetails, ok := s.D.GetOkExists("configuration_details"); ok {
		tmp, err := stringToConfigurationDetailsJsonObject(configurationDetails.(string))
		if err != nil {
			return err
		}
		request.ConfigurationDetails = tmp
		/*if tmpList := configurationDetails.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "configuration_details", 0)
			tmp, err := s.mapToobject(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.ConfigurationDetails = &tmp
		}*/
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

	if name, ok := s.D.GetOkExists("name"); ok {
		tmp := name.(string)
		request.Name = &tmp
	}

	if onPremiseVantagePointCompositeId, ok := s.D.GetOkExists("on_premise_vantage_point_id"); ok {
		tmp := onPremiseVantagePointCompositeId.(string)
		onPremiseVantagePointId, _, err := parseOnPremiseVantagePointCompositeId(tmp)
		if err == nil {
			request.OnPremiseVantagePointId = &onPremiseVantagePointId
		} else {
			log.Printf("[WARN] Create() unable to parse onPremiseVantagePointCompositeId: %s", onPremiseVantagePointCompositeId)
		}
	}

	if priority, ok := s.D.GetOkExists("priority"); ok {
		tmp := priority.(int)
		request.Priority = &tmp
	}

	if resourcePrincipalTokenPublicKey, ok := s.D.GetOkExists("resource_principal_token_public_key"); ok {
		tmp := resourcePrincipalTokenPublicKey.(string)
		request.ResourcePrincipalTokenPublicKey = &tmp
	}

	if status, ok := s.D.GetOkExists("status"); ok {
		request.Status = oci_apm_synthetics.OnPremiseVantagePointWorkerStatusEnum(status.(string))
	}

	if version, ok := s.D.GetOkExists("version"); ok {
		tmp := version.(string)
		request.Version = &tmp
	}

	if workerType, ok := s.D.GetOkExists("worker_type"); ok {
		request.WorkerType = oci_apm_synthetics.OnPremiseVantagePointWorkerTypeEnum(workerType.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "apm_synthetics")

	response, err := s.Client.CreateWorker(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Worker
	return nil
}

func (s *ApmSyntheticsOnPremiseVantagePointWorkerResourceCrud) Get() error {
	request := oci_apm_synthetics.GetWorkerRequest{}

	onPremiseVantagePointCompositeId, workerId, apmDomainId, err := parseOnPremiseVantagePointWorkerCompositeId(s.D.Id())
	if err == nil {
		request.WorkerId = &workerId
		request.ApmDomainId = &apmDomainId
		onPremiseVantagePointId, _, err1 := parseOnPremiseVantagePointCompositeId(onPremiseVantagePointCompositeId)
		if err1 == nil {
			request.OnPremiseVantagePointId = &onPremiseVantagePointId
		} else {
			log.Printf("[WARN] Get() unable to parse onPremiseVantagePointCompositeId: %s", onPremiseVantagePointCompositeId)
		}
	} else {
		log.Printf("[WARN] Get() unable to parse current ID: %s", s.D.Id())
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "apm_synthetics")

	response, err := s.Client.GetWorker(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Worker
	return nil
}

func (s *ApmSyntheticsOnPremiseVantagePointWorkerResourceCrud) Update() error {
	request := oci_apm_synthetics.UpdateWorkerRequest{}

	if apmDomainId, ok := s.D.GetOkExists("apm_domain_id"); ok {
		tmp := apmDomainId.(string)
		request.ApmDomainId = &tmp
	}

	if configurationDetails, ok := s.D.GetOkExists("configuration_details"); ok {
		tmp, err := stringToConfigurationDetailsJsonObject(configurationDetails.(string))
		if err != nil {
			return err
		}
		request.ConfigurationDetails = tmp
		/*if tmpList := configurationDetails.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "configuration_details", 0)
			tmp, err := s.mapToobject(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.ConfigurationDetails = &tmp
		}*/
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

	if onPremiseVantagePointId, ok := s.D.GetOkExists("on_premise_vantage_point_id"); ok {
		tmp := onPremiseVantagePointId.(string)
		request.OnPremiseVantagePointId = &tmp
	}

	if priority, ok := s.D.GetOkExists("priority"); ok {
		tmp := priority.(int)
		request.Priority = &tmp
	}

	if status, ok := s.D.GetOkExists("status"); ok {
		request.Status = oci_apm_synthetics.OnPremiseVantagePointWorkerStatusEnum(status.(string))
	}

	onPremiseVantagePointCompositeId, workerId, apmDomainId, err := parseOnPremiseVantagePointWorkerCompositeId(s.D.Id())
	if err == nil {
		request.WorkerId = &workerId
		request.ApmDomainId = &apmDomainId
		onPremiseVantagePointId, _, err1 := parseOnPremiseVantagePointCompositeId(onPremiseVantagePointCompositeId)
		if err1 == nil {
			request.OnPremiseVantagePointId = &onPremiseVantagePointId
		} else {
			log.Printf("[WARN] Update() unable to parse onPremiseVantagePointCompositeId: %s", onPremiseVantagePointCompositeId)
		}
	} else {
		log.Printf("[WARN] Update() unable to parse current ID: %s", s.D.Id())
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "apm_synthetics")

	response, err := s.Client.UpdateWorker(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Worker
	return nil
}

func (s *ApmSyntheticsOnPremiseVantagePointWorkerResourceCrud) Delete() error {
	request := oci_apm_synthetics.DeleteWorkerRequest{}

	if apmDomainId, ok := s.D.GetOkExists("apm_domain_id"); ok {
		tmp := apmDomainId.(string)
		request.ApmDomainId = &tmp
	}

	if onPremiseVantagePointId, ok := s.D.GetOkExists("on_premise_vantage_point_id"); ok {
		tmp := onPremiseVantagePointId.(string)
		request.OnPremiseVantagePointId = &tmp
	}

	onPremiseVantagePointCompositeId, workerId, apmDomainId, err1 := parseOnPremiseVantagePointWorkerCompositeId(s.D.Id())
	if err1 == nil {
		request.WorkerId = &workerId
		request.ApmDomainId = &apmDomainId
		onPremiseVantagePointId, _, err2 := parseOnPremiseVantagePointCompositeId(onPremiseVantagePointCompositeId)
		if err2 == nil {
			request.OnPremiseVantagePointId = &onPremiseVantagePointId
		} else {
			log.Printf("[WARN] Delete() unable to parse current ID: %s", s.D.Id())
		}
	} else {
		log.Printf("[WARN] Delete() unable to parse current ID: %s", s.D.Id())
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "apm_synthetics")

	_, err := s.Client.DeleteWorker(context.Background(), request)
	return err
}

func (s *ApmSyntheticsOnPremiseVantagePointWorkerResourceCrud) SetData() error {

	onPremiseVantagePointId, _, apmDomainId, err := parseOnPremiseVantagePointWorkerCompositeId(s.D.Id())
	if err == nil {
		s.D.Set("on_premise_vantage_point_id", onPremiseVantagePointId)
		s.D.Set("apm_domain_id", apmDomainId)
	} else {
		log.Printf("[WARN] SetData() unable to parse current ID: %s", s.D.Id())
	}

	if s.Res.ConfigurationDetails != nil {
		s.D.Set("configuration_details", configurationDetailsJsonObjectToString(s.Res.ConfigurationDetails))
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.GeoInfo != nil {
		s.D.Set("geo_info", *s.Res.GeoInfo)
	}

	if s.Res.IdentityInfo != nil {
		s.D.Set("identity_info", []interface{}{IdentityInfoDetailsToMap(s.Res.IdentityInfo)})
	} else {
		s.D.Set("identity_info", nil)
	}

	monitorList := []interface{}{}
	for _, item := range s.Res.MonitorList {
		monitorList = append(monitorList, WorkerMonitorListToMap(item))
	}
	s.D.Set("monitor_list", monitorList)

	if s.Res.Name != nil {
		s.D.Set("name", *s.Res.Name)
	}

	if s.Res.OpvpId != nil {
		s.D.Set("opvp_id", *s.Res.OpvpId)
	}

	if s.Res.OpvpName != nil {
		s.D.Set("opvp_name", *s.Res.OpvpName)
	}

	if s.Res.Priority != nil {
		s.D.Set("priority", *s.Res.Priority)
	}

	if s.Res.RuntimeId != nil {
		s.D.Set("runtime_id", *s.Res.RuntimeId)
	}

	s.D.Set("status", s.Res.Status)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeLastSyncUp != nil {
		s.D.Set("time_last_sync_up", s.Res.TimeLastSyncUp.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	if s.Res.VersionDetails != nil {
		s.D.Set("version_details", []interface{}{OnPremiseVpWorkerVersionDetailsToMap(s.Res.VersionDetails)})
	} else {
		s.D.Set("version_details", nil)
	}

	s.D.Set("worker_type", s.Res.WorkerType)

	return nil
}

func GetOnPremiseVantagePointWorkerCompositeId(onPremiseVantagePointId string, workerId string, apmDomainId string) string {
	apmDomainId = url.PathEscape(apmDomainId)
	onPremiseVantagePointId = url.PathEscape(onPremiseVantagePointId)
	workerId = url.PathEscape(workerId)
	compositeId := "onPremiseVantagePoints/" + onPremiseVantagePointId + "/workers/" + workerId + "/apmDomainId/" + apmDomainId
	return compositeId
}

func parseOnPremiseVantagePointWorkerCompositeId(compositeId string) (onPremiseVantagePointId string, workerId string, apmDomainId string, err error) {
	parts := strings.Split(compositeId, "/")
	match, _ := regexp.MatchString("onPremiseVantagePoints/.*/workers/.*/apmDomainId/.*", compositeId)
	if !match || len(parts) != 6 {
		err = fmt.Errorf("illegal compositeId %s encountered", compositeId)
		return
	}
	onPremiseVantagePointId, _ = url.PathUnescape(parts[1])
	workerId, _ = url.PathUnescape(parts[3])
	apmDomainId, _ = url.PathUnescape(parts[5])

	return
}

func IdentityInfoDetailsToMap(obj *oci_apm_synthetics.IdentityInfoDetails) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.ApmShortId != nil {
		result["apm_short_id"] = string(*obj.ApmShortId)
	}

	if obj.CollectorEndPoint != nil {
		result["collector_end_point"] = string(*obj.CollectorEndPoint)
	}

	if obj.RegionName != nil {
		result["region_name"] = string(*obj.RegionName)
	}

	return result
}

func OnPremiseVpWorkerVersionDetailsToMap(obj *oci_apm_synthetics.OnPremiseVpWorkerVersionDetails) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.LatestVersion != nil {
		result["latest_version"] = string(*obj.LatestVersion)
	}

	if obj.MinSupportedVersion != nil {
		result["min_supported_version"] = string(*obj.MinSupportedVersion)
	}

	if obj.Version != nil {
		result["version"] = string(*obj.Version)
	}

	return result
}

func WorkerMonitorListToMap(obj oci_apm_synthetics.WorkerMonitorList) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.DisplayName != nil {
		result["display_name"] = string(*obj.DisplayName)
	}

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	if obj.IsRunNow != nil {
		result["is_run_now"] = bool(*obj.IsRunNow)
	}

	result["monitor_type"] = string(obj.MonitorType)

	if obj.TimeAssigned != nil {
		result["time_assigned"] = obj.TimeAssigned.String()
	}

	return result
}

func WorkerSummaryToMap(obj oci_apm_synthetics.WorkerSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.ConfigurationDetails != nil {
		result["configuration_details"] = configurationDetailsJsonObjectToString(obj.ConfigurationDetails)
	}

	if obj.DefinedTags != nil {
		result["defined_tags"] = tfresource.DefinedTagsToMap(obj.DefinedTags)
	}

	if obj.DisplayName != nil {
		result["display_name"] = string(*obj.DisplayName)
	}

	result["freeform_tags"] = obj.FreeformTags

	if obj.GeoInfo != nil {
		result["geo_info"] = string(*obj.GeoInfo)
	}

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	monitorList := []interface{}{}
	for _, item := range obj.MonitorList {
		monitorList = append(monitorList, WorkerMonitorListToMap(item))
	}
	result["monitor_list"] = monitorList

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	if obj.Priority != nil {
		result["priority"] = int(*obj.Priority)
	}

	result["status"] = string(obj.Status)

	if obj.TimeCreated != nil {
		result["time_created"] = obj.TimeCreated.String()
	}

	/*if obj.TimeLastSyncup != nil {
		result["time_last_sync_up"] = obj.TimeLastSyncup.String()
	}*/

	if obj.TimeUpdated != nil {
		result["time_updated"] = obj.TimeUpdated.String()
	}

	if obj.VersionDetails != nil {
		result["version_details"] = []interface{}{OnPremiseVpWorkerVersionDetailsToMap(obj.VersionDetails)}
	}

	result["worker_type"] = string(obj.WorkerType)

	return result
}

func stringToConfigurationDetailsJsonObject(options string) (*interface{}, error) {
	var result interface{}
	var err error

	var obj interface{}
	err = json.Unmarshal([]byte(options), &obj)
	result = &obj

	return &result, err
}

func configurationDetailsJsonObjectToString(obj *interface{}) string {
	var result string

	if obj != nil {
		var bytes, _ = json.Marshal(obj)
		result = string(bytes)
	}

	return result
}
