// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package usage_proxy

import (
	"context"
	"fmt"
	"log"
	"net/url"
	"regexp"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	oci_usage_proxy "github.com/oracle/oci-go-sdk/v56/usage"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"
)

func UsageProxySubscriptionRedeemableUserResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createUsageProxySubscriptionRedeemableUser,
		Read:     readUsageProxySubscriptionRedeemableUser,
		Delete:   deleteUsageProxySubscriptionRedeemableUser,
		Schema: map[string]*schema.Schema{
			// Required
			"subscription_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"tenancy_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			// Optional
			"items": {
				Type:     schema.TypeList,
				Required: true,
				ForceNew: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"email_id": {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},

						// Optional

						// Computed
					},
				},
			},
			"user_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},

			// Computed
		},
	}
}

func createUsageProxySubscriptionRedeemableUser(d *schema.ResourceData, m interface{}) error {
	sync := &UsageProxySubscriptionRedeemableUserResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).RewardsClient()

	return tfresource.CreateResource(d, sync)
}

func readUsageProxySubscriptionRedeemableUser(d *schema.ResourceData, m interface{}) error {
	sync := &UsageProxySubscriptionRedeemableUserResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).RewardsClient()

	return tfresource.ReadResource(sync)
}

func deleteUsageProxySubscriptionRedeemableUser(d *schema.ResourceData, m interface{}) error {
	sync := &UsageProxySubscriptionRedeemableUserResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).RewardsClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type UsageProxySubscriptionRedeemableUserResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_usage_proxy.RewardsClient
	Res                    *oci_usage_proxy.RedeemableUserCollection
	DisableNotFoundRetries bool
}

func (s *UsageProxySubscriptionRedeemableUserResourceCrud) ID() string {
	return GetSubscriptionRedeemableUserCompositeId(s.D.Get("subscription_id").(string), s.D.Get("tenancy_id").(string))
}

func (s *UsageProxySubscriptionRedeemableUserResourceCrud) Create() error {
	request := oci_usage_proxy.CreateRedeemableUserRequest{}

	if items, ok := s.D.GetOkExists("items"); ok {
		interfaces := items.([]interface{})
		tmp := make([]oci_usage_proxy.RedeemableUser, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "items", stateDataIndex)
			converted, err := s.mapToRedeemableUser(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("items") {
			request.Items = tmp
		}
	}

	if subscriptionId, ok := s.D.GetOkExists("subscription_id"); ok {
		tmp := subscriptionId.(string)
		request.SubscriptionId = &tmp
	}

	if tenancyId, ok := s.D.GetOkExists("tenancy_id"); ok {
		tmp := tenancyId.(string)
		request.TenancyId = &tmp
	}

	if userId, ok := s.D.GetOkExists("user_id"); ok {
		tmp := userId.(string)
		request.UserId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "usage_proxy")

	response, err := s.Client.CreateRedeemableUser(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.RedeemableUserCollection
	return nil
}

func (s *UsageProxySubscriptionRedeemableUserResourceCrud) Get() error {
	request := oci_usage_proxy.ListRedeemableUsersRequest{}

	if subscriptionId, ok := s.D.GetOkExists("subscription_id"); ok {
		tmp := subscriptionId.(string)
		request.SubscriptionId = &tmp
	}

	if tenancyId, ok := s.D.GetOkExists("tenancy_id"); ok {
		tmp := tenancyId.(string)
		request.TenancyId = &tmp
	}

	subscriptionId, tenancyId, err := parseSubscriptionRedeemableUserCompositeId(s.D.Id())
	if err == nil {
		request.SubscriptionId = &subscriptionId
		request.TenancyId = &tenancyId
	} else {
		log.Printf("[WARN] Get() unable to parse current ID: %s", s.D.Id())
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "usage_proxy")

	response, err := s.Client.ListRedeemableUsers(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.RedeemableUserCollection
	return nil
}

func (s *UsageProxySubscriptionRedeemableUserResourceCrud) Delete() error {
	request := oci_usage_proxy.DeleteRedeemableUserRequest{}
	if emailId, ok := s.D.GetOkExists("items.0.email_id"); ok {
		tmp := emailId.(string)
		request.EmailId = &tmp
	}

	if subscriptionId, ok := s.D.GetOkExists("subscription_id"); ok {
		tmp := subscriptionId.(string)
		request.SubscriptionId = &tmp
	}

	if tenancyId, ok := s.D.GetOkExists("tenancy_id"); ok {
		tmp := tenancyId.(string)
		request.TenancyId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "usage_proxy")

	_, err := s.Client.DeleteRedeemableUser(context.Background(), request)
	return err
}

func (s *UsageProxySubscriptionRedeemableUserResourceCrud) SetData() error {

	subscriptionId, tenancyId, err := parseSubscriptionRedeemableUserCompositeId(s.D.Id())
	if err == nil {
		s.D.Set("subscription_id", &subscriptionId)
		s.D.Set("tenancy_id", &tenancyId)
	} else {
		log.Printf("[WARN] SetData() unable to parse current ID: %s", s.D.Id())
	}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, RedeemableUserSummaryToMap(item))
	}
	s.D.Set("items", items)

	return nil
}

func GetSubscriptionRedeemableUserCompositeId(subscriptionId string, tenancyId string) string {
	subscriptionId = url.PathEscape(subscriptionId)
	tenancyId = url.PathEscape(tenancyId)
	compositeId := "subscriptions/" + subscriptionId + "/redeemableUsers/tenancyId/" + tenancyId + ""
	return compositeId
}

func parseSubscriptionRedeemableUserCompositeId(compositeId string) (subscriptionId string, tenancyId string, err error) {

	parts := strings.Split(compositeId, "/")
	match, _ := regexp.MatchString("subscriptions/.*/redeemableUsers/tenancyId/.*", compositeId)
	if !match || len(parts) != 5 {
		err = fmt.Errorf("illegal compositeId %s encountered", compositeId)
		return
	}
	subscriptionId, _ = url.PathUnescape(parts[1])
	tenancyId, _ = url.PathUnescape(parts[4])

	return
}

func (s *UsageProxySubscriptionRedeemableUserResourceCrud) mapToRedeemableUser(fieldKeyFormat string) (oci_usage_proxy.RedeemableUser, error) {
	result := oci_usage_proxy.RedeemableUser{}

	if emailId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "email_id")); ok {
		tmp := emailId.(string)
		result.EmailId = &tmp
	}

	return result, nil
}

func RedeemableUserSummaryToMap(obj oci_usage_proxy.RedeemableUserSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.EmailId != nil {
		result["email_id"] = string(*obj.EmailId)
	}

	return result
}
