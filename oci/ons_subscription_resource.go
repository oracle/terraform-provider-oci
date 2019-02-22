// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"context"
	"encoding/json"
	"reflect"

	"github.com/hashicorp/terraform/helper/schema"

	oci_ons "github.com/oracle/oci-go-sdk/ons"
)

func OnsSubscriptionResource() *schema.Resource {
	return &schema.Resource{
		Timeouts: DefaultTimeout,
		Create:   createSubscription,
		Read:     readSubscription,
		Update:   updateSubscription,
		Delete:   deleteSubscription,
		Schema: map[string]*schema.Schema{
			// Required
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
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
				DiffSuppressFunc: definedTagsDiffSuppressFunction,
				Elem:             schema.TypeString,
			},
			"freeform_tags": {
				Type:     schema.TypeMap,
				Optional: true,
				Computed: true,
				Elem:     schema.TypeString,
			},

			// Computed
			"delivery_policy": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				//json string, convert to json to see if they are the same
				DiffSuppressFunc: jsonStringDiffSuppresionFunction,
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

func createSubscription(d *schema.ResourceData, m interface{}) error {
	sync := &SubscriptionResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).notificationDataPlaneClient

	return CreateResource(d, sync)
}

func readSubscription(d *schema.ResourceData, m interface{}) error {
	sync := &SubscriptionResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).notificationDataPlaneClient

	return ReadResource(sync)
}

func updateSubscription(d *schema.ResourceData, m interface{}) error {
	sync := &SubscriptionResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).notificationDataPlaneClient

	return UpdateResource(d, sync)
}

func deleteSubscription(d *schema.ResourceData, m interface{}) error {
	sync := &SubscriptionResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).notificationDataPlaneClient
	sync.DisableNotFoundRetries = true

	return DeleteResource(d, sync)
}

type SubscriptionResourceCrud struct {
	BaseCrud
	Client                 *oci_ons.NotificationDataPlaneClient
	Res                    *oci_ons.Subscription
	DisableNotFoundRetries bool
}

func (s *SubscriptionResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *SubscriptionResourceCrud) CreatedPending() []string {
	return []string{}
}

func (s *SubscriptionResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_ons.SubscriptionLifecycleStatePending),
		string(oci_ons.SubscriptionLifecycleStateActive),
	}
}

func (s *SubscriptionResourceCrud) DeletedPending() []string {
	return []string{}
}

func (s *SubscriptionResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_ons.SubscriptionLifecycleStateDeleted),
	}
}

func (s *SubscriptionResourceCrud) Create() error {
	request := oci_ons.CreateSubscriptionRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
		convertedDefinedTags, err := mapToDefinedTags(definedTags.(map[string]interface{}))
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
		request.FreeformTags = objectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if protocol, ok := s.D.GetOkExists("protocol"); ok {
		tmp := protocol.(string)
		request.Protocol = &tmp
	}

	if topicId, ok := s.D.GetOkExists("topic_id"); ok {
		tmp := topicId.(string)
		request.TopicId = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "ons")

	response, err := s.Client.CreateSubscription(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Subscription
	return nil
}

func (s *SubscriptionResourceCrud) Get() error {
	request := oci_ons.GetSubscriptionRequest{}

	tmp := s.D.Id()
	request.SubscriptionId = &tmp

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "ons")

	response, err := s.Client.GetSubscription(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Subscription
	return nil
}

func (s *SubscriptionResourceCrud) Update() error {
	request := oci_ons.UpdateSubscriptionRequest{}

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
		convertedDefinedTags, err := mapToDefinedTags(definedTags.(map[string]interface{}))
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
		request.FreeformTags = objectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	tmp := s.D.Id()
	request.SubscriptionId = &tmp

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "ons")

	_, err := s.Client.UpdateSubscription(context.Background(), request)
	if err == nil {
		return s.Get()
	}

	return err
}

func (s *SubscriptionResourceCrud) Delete() error {
	request := oci_ons.DeleteSubscriptionRequest{}

	tmp := s.D.Id()
	request.SubscriptionId = &tmp

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "ons")

	_, err := s.Client.DeleteSubscription(context.Background(), request)
	return err
}

func (s *SubscriptionResourceCrud) SetData() error {
	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", definedTagsToMap(s.Res.DefinedTags))
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
