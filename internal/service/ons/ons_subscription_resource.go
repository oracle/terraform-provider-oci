// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package ons

import (
	"context"
	"encoding/json"
	"reflect"
	"strconv"

	"github.com/terraform-providers/terraform-provider-oci/internal/utils"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	oci_ons "github.com/oracle/oci-go-sdk/v56/ons"
)

func OnsSubscriptionResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createOnsSubscription,
		Read:     readOnsSubscription,
		Update:   updateOnsSubscription,
		Delete:   deleteOnsSubscription,
		Schema: map[string]*schema.Schema{
			// Required
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"endpoint": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"protocol": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"topic_id": {
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
			"delivery_policy": {
				Type:             schema.TypeString,
				Optional:         true,
				Computed:         true,
				DiffSuppressFunc: jsonStringDiffSuppresionFunction,
			},
			"freeform_tags": {
				Type:     schema.TypeMap,
				Optional: true,
				Computed: true,
				Elem:     schema.TypeString,
			},

			// Computed
			"created_time": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"etag": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func createOnsSubscription(d *schema.ResourceData, m interface{}) error {
	sync := &OnsSubscriptionResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).NotificationDataPlaneClient()

	return tfresource.CreateResource(d, sync)
}

func readOnsSubscription(d *schema.ResourceData, m interface{}) error {
	sync := &OnsSubscriptionResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).NotificationDataPlaneClient()

	return tfresource.ReadResource(sync)
}

func updateOnsSubscription(d *schema.ResourceData, m interface{}) error {
	sync := &OnsSubscriptionResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).NotificationDataPlaneClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteOnsSubscription(d *schema.ResourceData, m interface{}) error {
	sync := &OnsSubscriptionResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).NotificationDataPlaneClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type OnsSubscriptionResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_ons.NotificationDataPlaneClient
	Res                    *oci_ons.Subscription
	DisableNotFoundRetries bool
}

func (s *OnsSubscriptionResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *OnsSubscriptionResourceCrud) CreatedPending() []string {
	return []string{}
}

func (s *OnsSubscriptionResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_ons.SubscriptionLifecycleStatePending),
		string(oci_ons.SubscriptionLifecycleStateActive),
	}
}

func (s *OnsSubscriptionResourceCrud) DeletedPending() []string {
	return []string{}
}

func (s *OnsSubscriptionResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_ons.SubscriptionLifecycleStateDeleted),
	}
}

func (s *OnsSubscriptionResourceCrud) Create() error {
	request := oci_ons.CreateSubscriptionRequest{}

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

	if endpoint, ok := s.D.GetOkExists("endpoint"); ok {
		tmp := endpoint.(string)
		request.Endpoint = &tmp
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = utils.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if protocol, ok := s.D.GetOkExists("protocol"); ok {
		tmp := protocol.(string)
		request.Protocol = &tmp
	}

	if topicId, ok := s.D.GetOkExists("topic_id"); ok {
		tmp := topicId.(string)
		request.TopicId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "ons")

	response, err := s.Client.CreateSubscription(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Subscription
	return nil
}

func (s *OnsSubscriptionResourceCrud) Get() error {
	request := oci_ons.GetSubscriptionRequest{}

	tmp := s.D.Id()
	request.SubscriptionId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "ons")

	response, err := s.Client.GetSubscription(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Subscription
	return nil
}

func (s *OnsSubscriptionResourceCrud) Update() error {
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(compartment)
			if err != nil {
				return err
			}
		}
	}
	request := oci_ons.UpdateSubscriptionRequest{}

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
		convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return err
		}
		request.DefinedTags = convertedDefinedTags
	}

	if deliveryPolicy, ok := s.D.GetOkExists("delivery_policy"); ok {
		res := oci_ons.DeliveryPolicy{}
		// due to deliveryPolicy's difference between Update and Get Api
		if err := json.Unmarshal([]byte(deliveryPolicy.(string)), &res); err != nil {
			return err
		}
		request.DeliveryPolicy = &res
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = utils.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	tmp := s.D.Id()
	request.SubscriptionId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "ons")

	_, err := s.Client.UpdateSubscription(context.Background(), request)
	if err == nil {
		return s.Get()
	}

	return err
}

func (s *OnsSubscriptionResourceCrud) Delete() error {
	request := oci_ons.DeleteSubscriptionRequest{}

	tmp := s.D.Id()
	request.SubscriptionId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "ons")

	_, err := s.Client.DeleteSubscription(context.Background(), request)
	return err
}

func (s *OnsSubscriptionResourceCrud) SetData() error {
	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.CreatedTime != nil {
		s.D.Set("created_time", strconv.FormatInt(*s.Res.CreatedTime, 10))
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DeliverPolicy != nil {
		s.D.Set("delivery_policy", *s.Res.DeliverPolicy)
	}

	if s.Res.Endpoint != nil {
		s.D.Set("endpoint", *s.Res.Endpoint)
	}

	if s.Res.Etag != nil {
		s.D.Set("etag", *s.Res.Etag)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.Protocol != nil {
		s.D.Set("protocol", *s.Res.Protocol)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TopicId != nil {
		s.D.Set("topic_id", *s.Res.TopicId)
	}

	return nil
}

func BackoffRetryPolicyToMap(obj *oci_ons.BackoffRetryPolicy) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.MaxRetryDuration != nil {
		result["max_retry_duration"] = int(*obj.MaxRetryDuration)
	}

	result["policy_type"] = string(obj.PolicyType)

	return result
}

func DeliveryPolicyToMap(obj *oci_ons.DeliveryPolicy) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.BackoffRetryPolicy != nil {
		result["backoff_retry_policy"] = []interface{}{BackoffRetryPolicyToMap(obj.BackoffRetryPolicy)}
	}

	return result
}

// check if the json data are the same
func jsonStringDiffSuppresionFunction(key string, old string, new string, d *schema.ResourceData) bool {
	if old == "" && new == "" {
		return true
	} else if old == "" || new == "" {
		return false
	} else {
		if old == new {
			return true
		}
		oldValue := oci_ons.DeliveryPolicy{}
		newValue := oci_ons.DeliveryPolicy{}
		if err := json.Unmarshal([]byte(old), &oldValue); err != nil {
			return false
		}
		if err := json.Unmarshal([]byte(new), &newValue); err != nil {
			return false
		}
		return reflect.DeepEqual(oldValue, newValue)
	}
}

func (s *OnsSubscriptionResourceCrud) updateCompartment(compartment interface{}) error {
	changeCompartmentRequest := oci_ons.ChangeSubscriptionCompartmentRequest{}

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.ChangeSubscriptionCompartmentDetails.CompartmentId = &compartmentTmp

	idTmp := s.D.Id()
	changeCompartmentRequest.SubscriptionId = &idTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "ons")

	_, err := s.Client.ChangeSubscriptionCompartment(context.Background(), changeCompartmentRequest)
	if err != nil {
		return err
	}

	if waitErr := tfresource.WaitForUpdatedState(s.D, s); waitErr != nil {
		return waitErr
	}

	return nil
}
