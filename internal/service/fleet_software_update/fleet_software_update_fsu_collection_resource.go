// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package fleet_software_update

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"

	oci_common "github.com/oracle/oci-go-sdk/v65/common"
	oci_fleet_software_update "github.com/oracle/oci-go-sdk/v65/fleetsoftwareupdate"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func FleetSoftwareUpdateFsuCollectionResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createFleetSoftwareUpdateFsuCollection,
		Read:     readFleetSoftwareUpdateFsuCollection,
		Update:   updateFleetSoftwareUpdateFsuCollection,
		Delete:   deleteFleetSoftwareUpdateFsuCollection,
		Schema: map[string]*schema.Schema{
			"id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			// Required
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"service_type": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"source_major_version": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"type": {
				Type:             schema.TypeString,
				Required:         true,
				ForceNew:         true,
				DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
				ValidateFunc: validation.StringInSlice([]string{
					"DB",
					"GI",
				}, true),
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
			"fleet_discovery": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				ForceNew: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"strategy": {
							Type:             schema.TypeString,
							Required:         true,
							ForceNew:         true,
							DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
							ValidateFunc: validation.StringInSlice([]string{
								"DISCOVERY_RESULTS",
								"FILTERS",
								"SEARCH_QUERY",
								"TARGET_LIST",
							}, true),
						},

						// Optional
						"filters": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							ForceNew: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required
									"type": {
										Type:             schema.TypeString,
										Required:         true,
										ForceNew:         true,
										DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
										ValidateFunc: validation.StringInSlice([]string{
											"COMPARTMENT_ID",
											"DB_HOME_NAME",
											"DB_NAME",
											"DB_UNIQUE_NAME",
											"DEFINED_TAG",
											"FREEFORM_TAG",
											"RESOURCE_ID",
											"VERSION",
										}, true),
									},

									// Optional
									"entity_type": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
										ForceNew: true,
									},
									"identifiers": {
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										ForceNew: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"mode": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
										ForceNew: true,
									},
									"names": {
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										ForceNew: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"operator": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
										ForceNew: true,
									},
									"tags": {
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										ForceNew: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required
												"key": {
													Type:     schema.TypeString,
													Required: true,
													ForceNew: true,
												},
												"value": {
													Type:     schema.TypeString,
													Required: true,
													ForceNew: true,
												},

												// Optional
												"namespace": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
													ForceNew: true,
												},

												// Computed
											},
										},
									},
									"versions": {
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										ForceNew: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},

									// Computed
								},
							},
						},
						"fsu_discovery_id": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},
						"query": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},
						"targets": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							ForceNew: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},

						// Computed
					},
				},
			},
			"freeform_tags": {
				Type:     schema.TypeMap,
				Optional: true,
				Computed: true,
				Elem:     schema.TypeString,
			},

			// Computed
			"active_fsu_cycle": {
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
					},
				},
			},
			"lifecycle_details": {
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
			"target_count": {
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
		},
	}
}

func createFleetSoftwareUpdateFsuCollection(d *schema.ResourceData, m interface{}) error {
	sync := &FleetSoftwareUpdateFsuCollectionResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).FleetSoftwareUpdateClient()

	return tfresource.CreateResource(d, sync)
}

func readFleetSoftwareUpdateFsuCollection(d *schema.ResourceData, m interface{}) error {
	sync := &FleetSoftwareUpdateFsuCollectionResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).FleetSoftwareUpdateClient()

	return tfresource.ReadResource(sync)
}

func updateFleetSoftwareUpdateFsuCollection(d *schema.ResourceData, m interface{}) error {
	sync := &FleetSoftwareUpdateFsuCollectionResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).FleetSoftwareUpdateClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteFleetSoftwareUpdateFsuCollection(d *schema.ResourceData, m interface{}) error {
	sync := &FleetSoftwareUpdateFsuCollectionResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).FleetSoftwareUpdateClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type FleetSoftwareUpdateFsuCollectionResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_fleet_software_update.FleetSoftwareUpdateClient
	Res                    *oci_fleet_software_update.FsuCollection
	DisableNotFoundRetries bool
}

func (s *FleetSoftwareUpdateFsuCollectionResourceCrud) ID() string {
	fsuCollection := *s.Res
	return *fsuCollection.GetId()
}

func (s *FleetSoftwareUpdateFsuCollectionResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_fleet_software_update.CollectionLifecycleStatesCreating),
	}
}

func (s *FleetSoftwareUpdateFsuCollectionResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_fleet_software_update.CollectionLifecycleStatesActive),
		string(oci_fleet_software_update.CollectionLifecycleStatesNeedsAttention),
	}
}

func (s *FleetSoftwareUpdateFsuCollectionResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_fleet_software_update.CollectionLifecycleStatesDeleting),
	}
}

func (s *FleetSoftwareUpdateFsuCollectionResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_fleet_software_update.CollectionLifecycleStatesDeleted),
	}
}

func (s *FleetSoftwareUpdateFsuCollectionResourceCrud) Create() error {
	request := oci_fleet_software_update.CreateFsuCollectionRequest{}
	err := s.populateTopLevelPolymorphicCreateFsuCollectionRequest(&request)
	if err != nil {
		return err
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "fleet_software_update")

	response, err := s.Client.CreateFsuCollection(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	var identifier *string
	identifier = response.GetId()
	if identifier != nil {
		s.D.SetId(*identifier)
	}
	return s.getFsuCollectionFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "fleet_software_update"), oci_fleet_software_update.ActionTypeCreated, s.D.Timeout(schema.TimeoutCreate))
}

func (s *FleetSoftwareUpdateFsuCollectionResourceCrud) getFsuCollectionFromWorkRequest(workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_fleet_software_update.ActionTypeEnum, timeout time.Duration) error {

	// Wait until it finishes
	fsuCollectionId, err := fsuCollectionWaitForWorkRequest(workId, "collection",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.Client)

	if err != nil {
		return err
	}
	s.D.SetId(*fsuCollectionId)

	return s.Get()
}

func fsuCollectionWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
	startTime := time.Now()
	stopTime := startTime.Add(timeout)
	return func(response oci_common.OCIOperationResponse) bool {

		// Stop after timeout has elapsed
		if time.Now().After(stopTime) {
			return false
		}

		// Make sure we stop on default rules
		if tfresource.ShouldRetry(response, false, "fleet_software_update", startTime) {
			return true
		}

		// Only stop if the time Finished is set
		if workRequestResponse, ok := response.Response.(oci_fleet_software_update.GetWorkRequestResponse); ok {
			return workRequestResponse.TimeFinished == nil
		}
		return false
	}
}

func fsuCollectionWaitForWorkRequest(wId *string, entityType string, action oci_fleet_software_update.ActionTypeEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_fleet_software_update.FleetSoftwareUpdateClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "fleet_software_update")
	retryPolicy.ShouldRetryOperation = fsuCollectionWorkRequestShouldRetryFunc(timeout)

	response := oci_fleet_software_update.GetWorkRequestResponse{}
	stateConf := &resource.StateChangeConf{
		Pending: []string{
			string(oci_fleet_software_update.OperationStatusInProgress),
			string(oci_fleet_software_update.OperationStatusAccepted),
			string(oci_fleet_software_update.OperationStatusCanceling),
		},
		Target: []string{
			string(oci_fleet_software_update.OperationStatusSucceeded),
			string(oci_fleet_software_update.OperationStatusFailed),
			string(oci_fleet_software_update.OperationStatusCanceled),
		},
		Refresh: func() (interface{}, string, error) {
			var err error
			response, err = client.GetWorkRequest(context.Background(),
				oci_fleet_software_update.GetWorkRequestRequest{
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
	if identifier == nil || response.Status == oci_fleet_software_update.OperationStatusFailed || response.Status == oci_fleet_software_update.OperationStatusCanceled {
		return nil, getErrorFromFleetSoftwareUpdateFsuCollectionWorkRequest(client, wId, retryPolicy, entityType, action)
	}

	return identifier, nil
}

func getErrorFromFleetSoftwareUpdateFsuCollectionWorkRequest(client *oci_fleet_software_update.FleetSoftwareUpdateClient, workId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_fleet_software_update.ActionTypeEnum) error {
	response, err := client.ListWorkRequestErrors(context.Background(),
		oci_fleet_software_update.ListWorkRequestErrorsRequest{
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

func (s *FleetSoftwareUpdateFsuCollectionResourceCrud) Get() error {
	request := oci_fleet_software_update.GetFsuCollectionRequest{}

	tmp := s.D.Id()
	request.FsuCollectionId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "fleet_software_update")

	response, err := s.Client.GetFsuCollection(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.FsuCollection
	return nil
}

func (s *FleetSoftwareUpdateFsuCollectionResourceCrud) Update() error {
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(compartment)
			if err != nil {
				return err
			}
		}
	}
	request := oci_fleet_software_update.UpdateFsuCollectionRequest{}

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

	tmp := s.D.Id()
	request.FsuCollectionId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "fleet_software_update")

	response, err := s.Client.UpdateFsuCollection(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getFsuCollectionFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "fleet_software_update"), oci_fleet_software_update.ActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}

func (s *FleetSoftwareUpdateFsuCollectionResourceCrud) Delete() error {
	request := oci_fleet_software_update.DeleteFsuCollectionRequest{}

	tmp := s.D.Id()
	request.FsuCollectionId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "fleet_software_update")

	response, err := s.Client.DeleteFsuCollection(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	// Wait until it finishes
	_, delWorkRequestErr := fsuCollectionWaitForWorkRequest(workId, "collection",
		oci_fleet_software_update.ActionTypeRelated, s.D.Timeout(schema.TimeoutDelete), s.DisableNotFoundRetries, s.Client)
	return delWorkRequestErr
}

func (s *FleetSoftwareUpdateFsuCollectionResourceCrud) SetData() error {
	s.D.SetId(s.D.Id())
	switch v := (*s.Res).(type) {
	case oci_fleet_software_update.DbCollection:
		s.D.Set("type", "DB")

		if v.FleetDiscovery != nil {
			fleetDiscoveryArray := []interface{}{}
			if fleetDiscoveryMap := DbFleetDiscoveryDetailsToMap(&v.FleetDiscovery); fleetDiscoveryMap != nil {
				fleetDiscoveryArray = append(fleetDiscoveryArray, fleetDiscoveryMap)
			}
			s.D.Set("fleet_discovery", fleetDiscoveryArray)
		} else {
			s.D.Set("fleet_discovery", nil)
		}

		s.D.Set("source_major_version", v.SourceMajorVersion)

		if v.ActiveFsuCycle != nil {
			s.D.Set("active_fsu_cycle", []interface{}{ActiveCycleDetailsToMap(v.ActiveFsuCycle)})
		} else {
			s.D.Set("active_fsu_cycle", nil)
		}

		if v.CompartmentId != nil {
			s.D.Set("compartment_id", *v.CompartmentId)
		}

		if v.DefinedTags != nil {
			s.D.Set("defined_tags", tfresource.DefinedTagsToMap(v.DefinedTags))
		}

		if v.DisplayName != nil {
			s.D.Set("display_name", *v.DisplayName)
		}

		s.D.Set("freeform_tags", v.FreeformTags)

		if v.Id != nil {
			s.D.Set("id", *v.Id)
		}

		if v.LifecycleDetails != nil {
			s.D.Set("lifecycle_details", *v.LifecycleDetails)
		}

		s.D.Set("service_type", v.ServiceType)

		s.D.Set("state", v.LifecycleState)

		if v.SystemTags != nil {
			s.D.Set("system_tags", tfresource.SystemTagsToMap(v.SystemTags))
		}

		if v.TargetCount != nil {
			s.D.Set("target_count", *v.TargetCount)
		}

		if v.TimeCreated != nil {
			s.D.Set("time_created", v.TimeCreated.String())
		}

		if v.TimeUpdated != nil {
			s.D.Set("time_updated", v.TimeUpdated.String())
		}
	case oci_fleet_software_update.GiCollection:
		s.D.Set("type", "GI")

		if v.FleetDiscovery != nil {
			fleetDiscoveryArray := []interface{}{}
			if fleetDiscoveryMap := GiFleetDiscoveryDetailsToMap(&v.FleetDiscovery); fleetDiscoveryMap != nil {
				fleetDiscoveryArray = append(fleetDiscoveryArray, fleetDiscoveryMap)
			}
			s.D.Set("fleet_discovery", fleetDiscoveryArray)
		} else {
			s.D.Set("fleet_discovery", nil)
		}

		s.D.Set("source_major_version", v.SourceMajorVersion)

		if v.ActiveFsuCycle != nil {
			s.D.Set("active_fsu_cycle", []interface{}{ActiveCycleDetailsToMap(v.ActiveFsuCycle)})
		} else {
			s.D.Set("active_fsu_cycle", nil)
		}

		if v.CompartmentId != nil {
			s.D.Set("compartment_id", *v.CompartmentId)
		}

		if v.DefinedTags != nil {
			s.D.Set("defined_tags", tfresource.DefinedTagsToMap(v.DefinedTags))
		}

		if v.DisplayName != nil {
			s.D.Set("display_name", *v.DisplayName)
		}

		s.D.Set("freeform_tags", v.FreeformTags)

		if v.Id != nil {
			s.D.Set("id", *v.Id)
		}

		if v.LifecycleDetails != nil {
			s.D.Set("lifecycle_details", *v.LifecycleDetails)
		}

		s.D.Set("service_type", v.ServiceType)

		s.D.Set("state", v.LifecycleState)

		if v.SystemTags != nil {
			s.D.Set("system_tags", tfresource.SystemTagsToMap(v.SystemTags))
		}

		if v.TargetCount != nil {
			s.D.Set("target_count", *v.TargetCount)
		}

		if v.TimeCreated != nil {
			s.D.Set("time_created", v.TimeCreated.String())
		}

		if v.TimeUpdated != nil {
			s.D.Set("time_updated", v.TimeUpdated.String())
		}
	default:
		log.Printf("[WARN] Received 'type' of unknown type %v", *s.Res)
		return nil
	}
	return nil
}

func ActiveCycleDetailsToMap(obj *oci_fleet_software_update.ActiveCycleDetails) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.DisplayName != nil {
		result["display_name"] = string(*obj.DisplayName)
	}

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	return result
}

func (s *FleetSoftwareUpdateFsuCollectionResourceCrud) mapToDbFleetDiscoveryDetails(fieldKeyFormat string) (oci_fleet_software_update.DbFleetDiscoveryDetails, error) {
	var baseObject oci_fleet_software_update.DbFleetDiscoveryDetails
	//discriminator
	strategyRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "strategy"))
	var strategy string
	if ok {
		strategy = strategyRaw.(string)
	} else {
		strategy = "" // default value
	}
	switch strings.ToLower(strategy) {
	case strings.ToLower("DISCOVERY_RESULTS"):
		details := oci_fleet_software_update.DbDiscoveryResults{}
		if fsuDiscoveryId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "fsu_discovery_id")); ok {
			tmp := fsuDiscoveryId.(string)
			details.FsuDiscoveryId = &tmp
		}
		baseObject = details
	case strings.ToLower("FILTERS"):
		details := oci_fleet_software_update.DbFiltersDiscovery{}
		if filters, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "filters")); ok {
			interfaces := filters.([]interface{})
			tmp := make([]oci_fleet_software_update.DbFleetDiscoveryFilter, len(interfaces))
			for i := range interfaces {
				stateDataIndex := i
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "filters"), stateDataIndex)
				converted, err := s.mapToDbFleetDiscoveryFilter(fieldKeyFormatNextLevel)
				if err != nil {
					return details, err
				}
				tmp[i] = converted
			}
			if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "filters")) {
				details.Filters = tmp
			}
		}
		baseObject = details
	case strings.ToLower("SEARCH_QUERY"):
		details := oci_fleet_software_update.DbSearchQueryDiscovery{}
		if query, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "query")); ok {
			tmp := query.(string)
			details.Query = &tmp
		}
		baseObject = details
	case strings.ToLower("TARGET_LIST"):
		details := oci_fleet_software_update.DbTargetListDiscovery{}
		if targets, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "targets")); ok {
			interfaces := targets.([]interface{})
			tmp := make([]string, len(interfaces))
			for i := range interfaces {
				if interfaces[i] != nil {
					tmp[i] = interfaces[i].(string)
				}
			}
			if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "targets")) {
				details.Targets = tmp
			}
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown strategy '%v' was specified", strategy)
	}
	return baseObject, nil
}

func DbFleetDiscoveryDetailsToMap(obj *oci_fleet_software_update.DbFleetDiscoveryDetails) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (*obj).(type) {
	case oci_fleet_software_update.DbDiscoveryResults:
		result["strategy"] = "DISCOVERY_RESULTS"

		if v.FsuDiscoveryId != nil {
			result["fsu_discovery_id"] = string(*v.FsuDiscoveryId)
		}
	case oci_fleet_software_update.DbFiltersDiscovery:
		result["strategy"] = "FILTERS"

		filters := []interface{}{}
		for _, item := range v.Filters {
			filters = append(filters, DbFleetDiscoveryFilterToMap(item))
		}
		result["filters"] = filters
	case oci_fleet_software_update.DbSearchQueryDiscovery:
		result["strategy"] = "SEARCH_QUERY"

		if v.Query != nil {
			result["query"] = string(*v.Query)
		}
	case oci_fleet_software_update.DbTargetListDiscovery:
		result["strategy"] = "TARGET_LIST"

		result["targets"] = v.Targets
	default:
		log.Printf("[WARN] Received 'strategy' of unknown type %v", *obj)
		return nil
	}

	return result
}

func (s *FleetSoftwareUpdateFsuCollectionResourceCrud) mapToDbFleetDiscoveryFilter(fieldKeyFormat string) (oci_fleet_software_update.DbFleetDiscoveryFilter, error) {
	var baseObject oci_fleet_software_update.DbFleetDiscoveryFilter
	//discriminator
	typeRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "type"))
	var type_ string
	if ok {
		type_ = typeRaw.(string)
	} else {
		type_ = "" // default value
	}
	switch strings.ToLower(type_) {
	case strings.ToLower("COMPARTMENT_ID"):
		details := oci_fleet_software_update.DbCompartmentIdFilter{}
		if identifiers, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "identifiers")); ok {
			interfaces := identifiers.([]interface{})
			tmp := make([]string, len(interfaces))
			for i := range interfaces {
				if interfaces[i] != nil {
					tmp[i] = interfaces[i].(string)
				}
			}
			if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "identifiers")) {
				details.Identifiers = tmp
			}
		}
		if mode, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "mode")); ok {
			details.Mode = oci_fleet_software_update.DbFleetDiscoveryFilterModeEnum(mode.(string))
		}
		baseObject = details
	case strings.ToLower("DB_HOME_NAME"):
		details := oci_fleet_software_update.DbHomeNameFilter{}
		if names, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "names")); ok {
			interfaces := names.([]interface{})
			tmp := make([]string, len(interfaces))
			for i := range interfaces {
				if interfaces[i] != nil {
					tmp[i] = interfaces[i].(string)
				}
			}
			if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "names")) {
				details.Names = tmp
			}
		}
		if mode, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "mode")); ok {
			details.Mode = oci_fleet_software_update.DbFleetDiscoveryFilterModeEnum(mode.(string))
		}
		baseObject = details
	case strings.ToLower("DB_NAME"):
		details := oci_fleet_software_update.DbNameFilter{}
		if names, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "names")); ok {
			interfaces := names.([]interface{})
			tmp := make([]string, len(interfaces))
			for i := range interfaces {
				if interfaces[i] != nil {
					tmp[i] = interfaces[i].(string)
				}
			}
			if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "names")) {
				details.Names = tmp
			}
		}
		if mode, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "mode")); ok {
			details.Mode = oci_fleet_software_update.DbFleetDiscoveryFilterModeEnum(mode.(string))
		}
		baseObject = details
	case strings.ToLower("DB_UNIQUE_NAME"):
		details := oci_fleet_software_update.DbUniqueNameFilter{}
		if names, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "names")); ok {
			interfaces := names.([]interface{})
			tmp := make([]string, len(interfaces))
			for i := range interfaces {
				if interfaces[i] != nil {
					tmp[i] = interfaces[i].(string)
				}
			}
			if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "names")) {
				details.Names = tmp
			}
		}
		if mode, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "mode")); ok {
			details.Mode = oci_fleet_software_update.DbFleetDiscoveryFilterModeEnum(mode.(string))
		}
		baseObject = details
	case strings.ToLower("DEFINED_TAG"):
		details := oci_fleet_software_update.DbDefinedTagsFilter{}
		if operator, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "operator")); ok {
			details.Operator = oci_fleet_software_update.FleetDiscoveryOperatorsEnum(operator.(string))
		}
		if tags, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "tags")); ok {
			interfaces := tags.([]interface{})
			tmp := make([]oci_fleet_software_update.DefinedTagFilterEntry, len(interfaces))
			for i := range interfaces {
				stateDataIndex := i
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "tags"), stateDataIndex)
				converted, err := s.mapToDefinedTagFilterEntry(fieldKeyFormatNextLevel)
				if err != nil {
					return details, err
				}
				tmp[i] = converted
			}
			if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "tags")) {
				details.Tags = tmp
			}
		}
		if mode, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "mode")); ok {
			details.Mode = oci_fleet_software_update.DbFleetDiscoveryFilterModeEnum(mode.(string))
		}
		baseObject = details
	case strings.ToLower("FREEFORM_TAG"):
		details := oci_fleet_software_update.DbFreeformTagsFilter{}
		if operator, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "operator")); ok {
			details.Operator = oci_fleet_software_update.FleetDiscoveryOperatorsEnum(operator.(string))
		}
		if tags, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "tags")); ok {
			interfaces := tags.([]interface{})
			tmp := make([]oci_fleet_software_update.FreeformTagFilterEntry, len(interfaces))
			for i := range interfaces {
				stateDataIndex := i
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "tags"), stateDataIndex)
				converted, err := s.mapToFreeformTagFilterEntry(fieldKeyFormatNextLevel)
				if err != nil {
					return details, err
				}
				tmp[i] = converted
			}
			if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "tags")) {
				details.Tags = tmp
			}
		}
		if mode, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "mode")); ok {
			details.Mode = oci_fleet_software_update.DbFleetDiscoveryFilterModeEnum(mode.(string))
		}
		baseObject = details
	case strings.ToLower("RESOURCE_ID"):
		details := oci_fleet_software_update.DbResourceIdFilter{}
		if entityType, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "entity_type")); ok {
			details.EntityType = oci_fleet_software_update.DbResourceIdFilterEntityTypeEnum(entityType.(string))
		}
		if identifiers, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "identifiers")); ok {
			interfaces := identifiers.([]interface{})
			tmp := make([]string, len(interfaces))
			for i := range interfaces {
				if interfaces[i] != nil {
					tmp[i] = interfaces[i].(string)
				}
			}
			if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "identifiers")) {
				details.Identifiers = tmp
			}
		}
		if operator, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "operator")); ok {
			details.Operator = oci_fleet_software_update.FleetDiscoveryOperatorsEnum(operator.(string))
		}
		if mode, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "mode")); ok {
			details.Mode = oci_fleet_software_update.DbFleetDiscoveryFilterModeEnum(mode.(string))
		}
		baseObject = details
	case strings.ToLower("VERSION"):
		details := oci_fleet_software_update.DbVersionFilter{}
		if versions, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "versions")); ok {
			interfaces := versions.([]interface{})
			tmp := make([]string, len(interfaces))
			for i := range interfaces {
				if interfaces[i] != nil {
					tmp[i] = interfaces[i].(string)
				}
			}
			if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "versions")) {
				details.Versions = tmp
			}
		}
		if mode, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "mode")); ok {
			details.Mode = oci_fleet_software_update.DbFleetDiscoveryFilterModeEnum(mode.(string))
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown type '%v' was specified", type_)
	}
	return baseObject, nil
}

func DbFleetDiscoveryFilterToMap(obj oci_fleet_software_update.DbFleetDiscoveryFilter) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (obj).(type) {
	case oci_fleet_software_update.DbCompartmentIdFilter:
		result["type"] = "COMPARTMENT_ID"

		result["identifiers"] = v.Identifiers
	case oci_fleet_software_update.DbHomeNameFilter:
		result["type"] = "DB_HOME_NAME"

		result["names"] = v.Names
	case oci_fleet_software_update.DbNameFilter:
		result["type"] = "DB_NAME"

		result["names"] = v.Names
	case oci_fleet_software_update.DbUniqueNameFilter:
		result["type"] = "DB_UNIQUE_NAME"

		result["names"] = v.Names
	case oci_fleet_software_update.DbDefinedTagsFilter:
		result["type"] = "DEFINED_TAG"

		result["operator"] = string(v.Operator)

		tags := []interface{}{}
		for _, item := range v.Tags {
			tags = append(tags, DefinedTagFilterEntryToMap(item))
		}
		result["tags"] = tags
	case oci_fleet_software_update.DbFreeformTagsFilter:
		result["type"] = "FREEFORM_TAG"

		result["operator"] = string(v.Operator)

		tags := []interface{}{}
		for _, item := range v.Tags {
			tags = append(tags, FreeformTagFilterEntryToMap(item))
		}
		result["tags"] = tags
	case oci_fleet_software_update.DbResourceIdFilter:
		result["type"] = "RESOURCE_ID"

		result["entity_type"] = string(v.EntityType)

		result["identifiers"] = v.Identifiers

		result["operator"] = string(v.Operator)
	case oci_fleet_software_update.DbVersionFilter:
		result["type"] = "VERSION"

		result["versions"] = v.Versions
	default:
		log.Printf("[WARN] Received 'type' of unknown type %v", obj)
		return nil
	}

	return result
}

func (s *FleetSoftwareUpdateFsuCollectionResourceCrud) mapToDefinedTagFilterEntry(fieldKeyFormat string) (oci_fleet_software_update.DefinedTagFilterEntry, error) {
	result := oci_fleet_software_update.DefinedTagFilterEntry{}

	if key, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "key")); ok {
		tmp := key.(string)
		result.Key = &tmp
	}

	if namespace, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "namespace")); ok {
		tmp := namespace.(string)
		result.Namespace = &tmp
	}

	if value, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "value")); ok {
		tmp := value.(string)
		result.Value = &tmp
	}

	return result, nil
}

func DefinedTagFilterEntryToMap(obj oci_fleet_software_update.DefinedTagFilterEntry) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Key != nil {
		result["key"] = string(*obj.Key)
	}

	if obj.Namespace != nil {
		result["namespace"] = string(*obj.Namespace)
	}

	if obj.Value != nil {
		result["value"] = string(*obj.Value)
	}

	return result
}

func (s *FleetSoftwareUpdateFsuCollectionResourceCrud) mapToFreeformTagFilterEntry(fieldKeyFormat string) (oci_fleet_software_update.FreeformTagFilterEntry, error) {
	result := oci_fleet_software_update.FreeformTagFilterEntry{}

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

func FreeformTagFilterEntryToMap(obj oci_fleet_software_update.FreeformTagFilterEntry) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Key != nil {
		result["key"] = string(*obj.Key)
	}

	if obj.Value != nil {
		result["value"] = string(*obj.Value)
	}

	return result
}

func FsuCollectionSummaryToMap(obj oci_fleet_software_update.FsuCollectionSummary) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (obj).(type) {
	case oci_fleet_software_update.DbFsuCollectionSummary:
		result["type"] = "DB"

		result["source_major_version"] = string(v.SourceMajorVersion)
		result["id"] = string(*v.Id)

		if v.ActiveFsuCycle != nil {
			result["active_fsu_cycle"] = []interface{}{ActiveCycleDetailsToMap(v.ActiveFsuCycle)}
		} else {
			result["active_fsu_cycle"] = nil
		}

		result["compartment_id"] = string(*v.CompartmentId)
		if v.DefinedTags != nil {
			result["defined_tags"] = tfresource.DefinedTagsToMap(v.DefinedTags)
		}

		if v.DisplayName != nil {
			result["display_name"] = string(*v.DisplayName)
		}

		result["freeform_tags"] = v.FreeformTags

		if v.LifecycleDetails != nil {
			result["lifecycle_details"] = string(*v.LifecycleDetails)
		}
		result["service_type"] = string(v.ServiceType)
		result["state"] = string(v.LifecycleState)
		if v.TargetCount != nil {
			result["target_count"] = *v.TargetCount
		}
		if v.TimeCreated != nil {
			result["time_created"] = v.TimeCreated.String()
		}

		if v.TimeUpdated != nil {
			result["time_updated"] = v.TimeUpdated.String()
		}

	case oci_fleet_software_update.GiFsuCollectionSummary:
		result["type"] = "GI"

		result["source_major_version"] = string(v.SourceMajorVersion)
		result["id"] = string(*v.Id)

		if v.ActiveFsuCycle != nil {
			result["active_fsu_cycle"] = []interface{}{ActiveCycleDetailsToMap(v.ActiveFsuCycle)}
		} else {
			result["active_fsu_cycle"] = nil
		}

		result["compartment_id"] = string(*v.CompartmentId)
		if v.DefinedTags != nil {
			result["defined_tags"] = tfresource.DefinedTagsToMap(v.DefinedTags)
		}

		if v.DisplayName != nil {
			result["display_name"] = string(*v.DisplayName)
		}

		result["freeform_tags"] = v.FreeformTags

		if v.LifecycleDetails != nil {
			result["lifecycle_details"] = string(*v.LifecycleDetails)
		}
		result["service_type"] = string(v.ServiceType)
		result["state"] = string(v.LifecycleState)
		if v.TargetCount != nil {
			result["target_count"] = *v.TargetCount
		}
		if v.TimeCreated != nil {
			result["time_created"] = v.TimeCreated.String()
		}

		if v.TimeUpdated != nil {
			result["time_updated"] = v.TimeUpdated.String()
		}

	default:
		log.Printf("[WARN] Received 'type' of unknown type %v", obj)
		return nil
	}

	return result
}

func (s *FleetSoftwareUpdateFsuCollectionResourceCrud) mapToGiFleetDiscoveryDetails(fieldKeyFormat string) (oci_fleet_software_update.GiFleetDiscoveryDetails, error) {
	var baseObject oci_fleet_software_update.GiFleetDiscoveryDetails
	//discriminator
	strategyRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "strategy"))
	var strategy string
	if ok {
		strategy = strategyRaw.(string)
	} else {
		strategy = "" // default value
	}
	switch strings.ToLower(strategy) {
	case strings.ToLower("DISCOVERY_RESULTS"):
		details := oci_fleet_software_update.GiDiscoveryResults{}
		if fsuDiscoveryId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "fsu_discovery_id")); ok {
			tmp := fsuDiscoveryId.(string)
			details.FsuDiscoveryId = &tmp
		}
		baseObject = details
	case strings.ToLower("FILTERS"):
		details := oci_fleet_software_update.GiFiltersDiscovery{}
		if filters, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "filters")); ok {
			interfaces := filters.([]interface{})
			tmp := make([]oci_fleet_software_update.GiFleetDiscoveryFilter, len(interfaces))
			for i := range interfaces {
				stateDataIndex := i
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "filters"), stateDataIndex)
				converted, err := s.mapToGiFleetDiscoveryFilter(fieldKeyFormatNextLevel)
				if err != nil {
					return details, err
				}
				tmp[i] = converted
			}
			if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "filters")) {
				details.Filters = tmp
			}
		}
		baseObject = details
	case strings.ToLower("SEARCH_QUERY"):
		details := oci_fleet_software_update.GiSearchQueryDiscovery{}
		if query, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "query")); ok {
			tmp := query.(string)
			details.Query = &tmp
		}
		baseObject = details
	case strings.ToLower("TARGET_LIST"):
		details := oci_fleet_software_update.GiTargetListDiscovery{}
		if targets, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "targets")); ok {
			interfaces := targets.([]interface{})
			tmp := make([]string, len(interfaces))
			for i := range interfaces {
				if interfaces[i] != nil {
					tmp[i] = interfaces[i].(string)
				}
			}
			if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "targets")) {
				details.Targets = tmp
			}
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown strategy '%v' was specified", strategy)
	}
	return baseObject, nil
}

func GiFleetDiscoveryDetailsToMap(obj *oci_fleet_software_update.GiFleetDiscoveryDetails) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (*obj).(type) {
	case oci_fleet_software_update.GiDiscoveryResults:
		result["strategy"] = "DISCOVERY_RESULTS"

		if v.FsuDiscoveryId != nil {
			result["fsu_discovery_id"] = string(*v.FsuDiscoveryId)
		}
	case oci_fleet_software_update.GiFiltersDiscovery:
		result["strategy"] = "FILTERS"

		filters := []interface{}{}
		for _, item := range v.Filters {
			filters = append(filters, GiFleetDiscoveryFilterToMap(item))
		}
		result["filters"] = filters
	case oci_fleet_software_update.GiSearchQueryDiscovery:
		result["strategy"] = "SEARCH_QUERY"

		if v.Query != nil {
			result["query"] = string(*v.Query)
		}
	case oci_fleet_software_update.GiTargetListDiscovery:
		result["strategy"] = "TARGET_LIST"

		result["targets"] = v.Targets
	default:
		log.Printf("[WARN] Received 'strategy' of unknown type %v", *obj)
		return nil
	}

	return result
}

func (s *FleetSoftwareUpdateFsuCollectionResourceCrud) mapToGiFleetDiscoveryFilter(fieldKeyFormat string) (oci_fleet_software_update.GiFleetDiscoveryFilter, error) {
	var baseObject oci_fleet_software_update.GiFleetDiscoveryFilter
	//discriminator
	typeRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "type"))
	var type_ string
	if ok {
		type_ = typeRaw.(string)
	} else {
		type_ = "" // default value
	}
	switch strings.ToLower(type_) {
	case strings.ToLower("COMPARTMENT_ID"):
		details := oci_fleet_software_update.GiCompartmentIdFilter{}
		if identifiers, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "identifiers")); ok {
			interfaces := identifiers.([]interface{})
			tmp := make([]string, len(interfaces))
			for i := range interfaces {
				if interfaces[i] != nil {
					tmp[i] = interfaces[i].(string)
				}
			}
			if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "identifiers")) {
				details.Identifiers = tmp
			}
		}
		if mode, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "mode")); ok {
			details.Mode = oci_fleet_software_update.GiFleetDiscoveryFilterModeEnum(mode.(string))
		}
		baseObject = details
	case strings.ToLower("DEFINED_TAG"):
		details := oci_fleet_software_update.GiDefinedTagsFilter{}
		if operator, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "operator")); ok {
			details.Operator = oci_fleet_software_update.FleetDiscoveryOperatorsEnum(operator.(string))
		}
		if tags, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "tags")); ok {
			interfaces := tags.([]interface{})
			tmp := make([]oci_fleet_software_update.DefinedTagFilterEntry, len(interfaces))
			for i := range interfaces {
				stateDataIndex := i
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "tags"), stateDataIndex)
				converted, err := s.mapToDefinedTagFilterEntry(fieldKeyFormatNextLevel)
				if err != nil {
					return details, err
				}
				tmp[i] = converted
			}
			if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "tags")) {
				details.Tags = tmp
			}
		}
		if mode, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "mode")); ok {
			details.Mode = oci_fleet_software_update.GiFleetDiscoveryFilterModeEnum(mode.(string))
		}
		baseObject = details
	case strings.ToLower("FREEFORM_TAG"):
		details := oci_fleet_software_update.GiFreeformTagsFilter{}
		if operator, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "operator")); ok {
			details.Operator = oci_fleet_software_update.FleetDiscoveryOperatorsEnum(operator.(string))
		}
		if tags, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "tags")); ok {
			interfaces := tags.([]interface{})
			tmp := make([]oci_fleet_software_update.FreeformTagFilterEntry, len(interfaces))
			for i := range interfaces {
				stateDataIndex := i
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "tags"), stateDataIndex)
				converted, err := s.mapToFreeformTagFilterEntry(fieldKeyFormatNextLevel)
				if err != nil {
					return details, err
				}
				tmp[i] = converted
			}
			if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "tags")) {
				details.Tags = tmp
			}
		}
		if mode, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "mode")); ok {
			details.Mode = oci_fleet_software_update.GiFleetDiscoveryFilterModeEnum(mode.(string))
		}
		baseObject = details
	case strings.ToLower("RESOURCE_ID"):
		details := oci_fleet_software_update.GiResourceIdFilter{}
		if entityType, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "entity_type")); ok {
			details.EntityType = oci_fleet_software_update.GiResourceIdFilterEntityTypeEnum(entityType.(string))
		}
		if identifiers, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "identifiers")); ok {
			interfaces := identifiers.([]interface{})
			tmp := make([]string, len(interfaces))
			for i := range interfaces {
				if interfaces[i] != nil {
					tmp[i] = interfaces[i].(string)
				}
			}
			if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "identifiers")) {
				details.Identifiers = tmp
			}
		}
		if operator, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "operator")); ok {
			details.Operator = oci_fleet_software_update.FleetDiscoveryOperatorsEnum(operator.(string))
		}
		if mode, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "mode")); ok {
			details.Mode = oci_fleet_software_update.GiFleetDiscoveryFilterModeEnum(mode.(string))
		}
		baseObject = details
	case strings.ToLower("VERSION"):
		details := oci_fleet_software_update.GiVersionFilter{}
		if versions, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "versions")); ok {
			interfaces := versions.([]interface{})
			tmp := make([]string, len(interfaces))
			for i := range interfaces {
				if interfaces[i] != nil {
					tmp[i] = interfaces[i].(string)
				}
			}
			if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "versions")) {
				details.Versions = tmp
			}
		}
		if mode, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "mode")); ok {
			details.Mode = oci_fleet_software_update.GiFleetDiscoveryFilterModeEnum(mode.(string))
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown type '%v' was specified", type_)
	}
	return baseObject, nil
}

func GiFleetDiscoveryFilterToMap(obj oci_fleet_software_update.GiFleetDiscoveryFilter) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (obj).(type) {
	case oci_fleet_software_update.GiCompartmentIdFilter:
		result["type"] = "COMPARTMENT_ID"

		result["identifiers"] = v.Identifiers
	case oci_fleet_software_update.GiDefinedTagsFilter:
		result["type"] = "DEFINED_TAG"

		result["operator"] = string(v.Operator)

		tags := []interface{}{}
		for _, item := range v.Tags {
			tags = append(tags, DefinedTagFilterEntryToMap(item))
		}
		result["tags"] = tags
	case oci_fleet_software_update.GiFreeformTagsFilter:
		result["type"] = "FREEFORM_TAG"

		result["operator"] = string(v.Operator)

		tags := []interface{}{}
		for _, item := range v.Tags {
			tags = append(tags, FreeformTagFilterEntryToMap(item))
		}
		result["tags"] = tags
	case oci_fleet_software_update.GiResourceIdFilter:
		result["type"] = "RESOURCE_ID"

		result["entity_type"] = string(v.EntityType)

		result["identifiers"] = v.Identifiers

		result["operator"] = string(v.Operator)
	case oci_fleet_software_update.GiVersionFilter:
		result["type"] = "VERSION"

		result["versions"] = v.Versions
	default:
		log.Printf("[WARN] Received 'type' of unknown type %v", obj)
		return nil
	}

	return result
}

func (s *FleetSoftwareUpdateFsuCollectionResourceCrud) populateTopLevelPolymorphicCreateFsuCollectionRequest(request *oci_fleet_software_update.CreateFsuCollectionRequest) error {
	//discriminator
	typeRaw, ok := s.D.GetOkExists("type")
	var type_ string
	if ok {
		type_ = typeRaw.(string)
	} else {
		type_ = "" // default value
	}
	switch strings.ToLower(type_) {
	case strings.ToLower("DB"):
		details := oci_fleet_software_update.CreateDbFsuCollectionDetails{}
		if fleetDiscovery, ok := s.D.GetOkExists("fleet_discovery"); ok {
			if tmpList := fleetDiscovery.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "fleet_discovery", 0)
				tmp, err := s.mapToDbFleetDiscoveryDetails(fieldKeyFormat)
				if err != nil {
					return err
				}
				details.FleetDiscovery = tmp
			}
		}
		if sourceMajorVersion, ok := s.D.GetOkExists("source_major_version"); ok {
			details.SourceMajorVersion = oci_fleet_software_update.DbSourceMajorVersionsEnum(sourceMajorVersion.(string))
		}
		if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
			tmp := compartmentId.(string)
			details.CompartmentId = &tmp
		}
		if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
			convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
			if err != nil {
				return err
			}
			details.DefinedTags = convertedDefinedTags
		}
		if displayName, ok := s.D.GetOkExists("display_name"); ok {
			tmp := displayName.(string)
			details.DisplayName = &tmp
		}
		if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
			details.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
		}
		if serviceType, ok := s.D.GetOkExists("service_type"); ok {
			details.ServiceType = oci_fleet_software_update.CollectionServiceTypesEnum(serviceType.(string))
		}
		request.CreateFsuCollectionDetails = details
	case strings.ToLower("GI"):
		details := oci_fleet_software_update.CreateGiFsuCollectionDetails{}
		if fleetDiscovery, ok := s.D.GetOkExists("fleet_discovery"); ok {
			if tmpList := fleetDiscovery.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "fleet_discovery", 0)
				tmp, err := s.mapToGiFleetDiscoveryDetails(fieldKeyFormat)
				if err != nil {
					return err
				}
				details.FleetDiscovery = tmp
			}
		}
		if sourceMajorVersion, ok := s.D.GetOkExists("source_major_version"); ok {
			details.SourceMajorVersion = oci_fleet_software_update.GiSourceMajorVersionsEnum(sourceMajorVersion.(string))
		}
		if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
			tmp := compartmentId.(string)
			details.CompartmentId = &tmp
		}
		if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
			convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
			if err != nil {
				return err
			}
			details.DefinedTags = convertedDefinedTags
		}
		if displayName, ok := s.D.GetOkExists("display_name"); ok {
			tmp := displayName.(string)
			details.DisplayName = &tmp
		}
		if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
			details.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
		}
		if serviceType, ok := s.D.GetOkExists("service_type"); ok {
			details.ServiceType = oci_fleet_software_update.CollectionServiceTypesEnum(serviceType.(string))
		}
		request.CreateFsuCollectionDetails = details
	default:
		return fmt.Errorf("unknown type '%v' was specified", type_)
	}
	return nil
}

func (s *FleetSoftwareUpdateFsuCollectionResourceCrud) updateCompartment(compartment interface{}) error {
	changeCompartmentRequest := oci_fleet_software_update.ChangeFsuCollectionCompartmentRequest{}

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	idTmp := s.D.Id()
	changeCompartmentRequest.FsuCollectionId = &idTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "fleet_software_update")

	response, err := s.Client.ChangeFsuCollectionCompartment(context.Background(), changeCompartmentRequest)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getFsuCollectionFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "fleet_software_update"), oci_fleet_software_update.ActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}
