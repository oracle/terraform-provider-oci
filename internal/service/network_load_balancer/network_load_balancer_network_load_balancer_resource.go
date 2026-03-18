// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package network_load_balancer

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/customdiff"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/retry"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	oci_common "github.com/oracle/oci-go-sdk/v65/common"
	oci_network_load_balancer "github.com/oracle/oci-go-sdk/v65/networkloadbalancer"
)

func NetworkLoadBalancerNetworkLoadBalancerResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createNetworkLoadBalancerNetworkLoadBalancer,
		Read:     readNetworkLoadBalancerNetworkLoadBalancer,
		Update:   updateNetworkLoadBalancerNetworkLoadBalancer,
		Delete:   deleteNetworkLoadBalancerNetworkLoadBalancer,
		Schema: map[string]*schema.Schema{
			// Required
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"subnet_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			// Optional
			"assigned_ipv6": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: false,
				ForceNew: true,
			},
			"assigned_private_ipv4": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: false,
				ForceNew: true,
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
			"is_preserve_source_destination": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"is_private": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"is_symmetric_hash_enabled": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"network_security_group_ids": {
				Type:     schema.TypeSet,
				Optional: true,
				Set:      tfresource.LiteralTypeHashCodeForSets,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"nlb_ip_version": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"reserved_ips": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: false,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional
						"id": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},

						// Computed
					},
				},
			},
			"security_attributes": {
				Type:     schema.TypeMap,
				Optional: true,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"subnet_ipv6cidr": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: false,
				ForceNew: true,
			},

			// Computed
			"ip_addresses": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"ip_address": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"ip_version": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"is_public": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"reserved_ip": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"id": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
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
			"time_created": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_updated": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
		CustomizeDiff: customdiff.All(
			// force change for certain reserved ips combination
			func(ctx context.Context, d *schema.ResourceDiff, meta interface{}) error {
				return forceNewReservedIps(ctx, d, meta)
			},
		),
	}
}

// forceNewReservedIps decides whether a change to reserved_ips should force recreation
// of the NLB resource or can be handled by an in-place Update.
//
// Rules:
// - Force NEW (destroy + create):
//  1. Any change in reserved IPv4 (both private and public) for single-stack IPv4 and dual-stack
//  2. Any change in reserved IPv6 for single-stack IPv6 and dual-stack
//  3. Mixed case: single-stack IPv4 → dual-stack AND IPv4 changes → force new (IPv4 change wins)
//
// - ALLOW UPDATE (no recreate):
//  1. NLB is updated from single-stack IPv4 to dual-stack using a reserved IPv6
//     (IPv4 unchanged, only IPv6 added/changed)
//  2. NLB is updated from dual-stack (created using reserved IPv6) to single-stack IPv4
//     (IPv4 unchanged, effectively just dropping IPv6)
func forceNewReservedIps(ctx context.Context, d *schema.ResourceDiff, meta interface{}) error {
	// Name of the Terraform attribute that holds reserved IPs
	const attrReservedIps = "reserved_ips"
	// Name of the Terraform attribute that holds the NLB IP version ("IPV4", "IPV4_AND_IPV6", "IPV6")
	const attrNlbIpVersion = "nlb_ip_version"

	// If reserved_ips didn't change in this plan, we don't need to do anything
	if !d.HasChange(attrReservedIps) {
		return nil
	}

	// Get old and new values of reserved_ips from the diff
	oldRaw, newRaw := d.GetChange(attrReservedIps)

	// Cast old value to []interface{} (Terraform internal type for list)
	oldList, _ := oldRaw.([]interface{})
	// Cast new value to []interface{}
	newList, _ := newRaw.([]interface{})

	// DEBUG: log the raw reserved_ips diff shape
	log.Printf("[DEBUG] forceNewReservedIps: old reserved_ips=%#v", oldList)
	log.Printf("[DEBUG] forceNewReservedIps: new reserved_ips=%#v", newList)

	// Helper: classify OCID as IPv6 if it starts with "ocid1.ipv6."
	// Otherwise treat it as IPv4 (covers both privateip and publicip).
	isIPv6ID := func(id string) bool {
		return strings.HasPrefix(id, "ocid1.ipv6.")
	}

	// Tiny enum type to tag an ID as none/IPv4/IPv6
	type ipType int

	const (
		ipNone ipType = iota // not used in logic, just a base value
		ipv4                 // represents an IPv4 OCID
		ipv6                 // represents an IPv6 OCID
	)

	// Convert Terraform list of reserved_ips into map[id] -> ipType
	toMap := func(list []interface{}) map[string]ipType {
		// Initialize an empty map
		m := make(map[string]ipType)
		// Iterate over each element in the list
		for _, v := range list {
			// Skip nil entries
			if v == nil {
				continue
			}
			// Each element should be map[string]interface{} (object with "id")
			mv, ok := v.(map[string]interface{})
			if !ok {
				continue
			}
			// Extract "id" as string
			id, ok := mv["id"].(string)
			if !ok || id == "" {
				continue
			}
			// Classify the ID
			if isIPv6ID(id) {
				m[id] = ipv6
			} else {
				m[id] = ipv4
			}
		}
		// Return the populated map
		return m
	}

	// Map of old reserved IPs: id -> type (ipv4/ipv6)
	oldMap := toMap(oldList)
	// Map of new reserved IPs: id -> type (ipv4/ipv6)
	newMap := toMap(newList)

	// Flags to indicate whether any IPv4 or IPv6 reserved IP changed
	var ipv4Changed, ipv6Changed bool
	var ipv4Removed bool

	// Find IDs that were present before but are missing now (removed IDs)
	for id, t := range oldMap {
		if _, exists := newMap[id]; !exists {
			// This ID was removed; mark IPv4 or IPv6 as changed
			if t == ipv6 {
				ipv6Changed = true
			} else {
				ipv4Changed = true
				ipv4Removed = true // <-- explicitly track removals
			}
		}
	}

	// Find IDs that are newly added (present now, absent before)
	for id, t := range newMap {
		if _, exists := oldMap[id]; !exists {
			// This ID was added; mark IPv4 or IPv6 as changed
			if t == ipv6 {
				ipv6Changed = true
			} else {
				ipv4Changed = true
			}
		}
	}

	// Now determine old and new values of nlb_ip_version
	var oldVer, newVer string

	// If nlb_ip_version changed in this plan
	if d.HasChange(attrNlbIpVersion) {
		// Get old and new values from diff
		ov, nv := d.GetChange(attrNlbIpVersion)
		// Cast old value to string if possible
		if s, ok := ov.(string); ok {
			oldVer = s
		}
		// Cast new value to string if possible
		if s, ok := nv.(string); ok {
			newVer = s
		}
	} else {
		// If nlb_ip_version didn't change, take current value (if set) as both old and new
		if v, ok := d.GetOkExists(attrNlbIpVersion); ok {
			if s, ok := v.(string); ok {
				oldVer, newVer = s, s
			}
		}
	}

	// If only newVer is set, copy it into oldVer so they match
	if oldVer == "" && newVer != "" {
		oldVer = newVer
	}
	// If only oldVer is set, copy it into newVer so they match
	if newVer == "" && oldVer != "" {
		newVer = oldVer
	}

	// Helper to test specific version values
	isIPv4Only := func(v string) bool { return v == "IPV4" }
	isIPv6Only := func(v string) bool { return v == "IPV6" }
	isDual := func(v string) bool { return v == "IPV4_AND_IPV6" }

	// DEBUG: log the versions and change flags
	log.Printf("[DEBUG] forceNewReservedIps: oldVer=%q newVer=%q ipv4Changed=%v ipv6Changed=%v",
		oldVer, newVer, ipv4Changed, ipv6Changed)

	// --- Allowed UPDATE scenario #1 ---
	// single-stack IPv4 -> dual-stack using reserved IPv6
	// - oldVer must be IPV4
	// - newVer must be IPV4_AND_IPV6
	// - IPv4 must NOT change
	// - IPv6 MUST change
	if isIPv4Only(oldVer) && isDual(newVer) && !ipv4Changed && ipv6Changed {
		log.Printf("[DEBUG] forceNewReservedIps: allowing update (IPV4 -> IPV4_AND_IPV6, only IPv6 changed)")
		// No ForceNew: Terraform will use the Update path
		return nil
	}

	// --- Allowed UPDATE scenario #2 ---
	// dual-stack (created using reserved IPv6) -> single-stack IPv4
	// - oldVer must be IPV4_AND_IPV6
	// - newVer must be IPV4
	// - IPv4 unchanged (we are effectively dropping IPv6)
	if isDual(oldVer) && isIPv4Only(newVer) && !ipv4Changed {
		log.Printf("[DEBUG] forceNewReservedIps: allowing update (IPV4_AND_IPV6 -> IPV4, dropping IPv6)")
		// No ForceNew: Terraform will use the Update path
		return nil
	}

	// Special: if an IPv4 entry was *removed*, force new on parent "reserved_ips"
	if ipv4Removed {
		log.Printf("[DEBUG] forceNewReservedIps: IPv4 reserved IP removed, forcing new on %q", attrReservedIps)
		return d.ForceNew(attrReservedIps)
	}

	// --- All other scenarios fall under "force new" logic ---

	// Helper: for any index where reserved_ips.<i>.id changed, mark that path ForceNew.
	forceNewIds := func() error {
		maxLen := len(oldList)
		if len(newList) > maxLen {
			maxLen = len(newList)
		}
		for i := 0; i < maxLen; i++ {
			key := fmt.Sprintf("reserved_ips.%d.id", i)

			// Only bother if Terraform thinks this specific key changed
			if d.HasChange(key) {
				log.Printf("[DEBUG] forceNewReservedIps: forcing new on %q", key)
				if err := d.ForceNew(key); err != nil {
					return err
				}
			}
		}
		return nil
	}

	// Rule: Any IPv4 change takes precedence and forces recreation.
	if ipv4Changed {
		log.Printf("[DEBUG] forceNewReservedIps: forcing new (IPv4 reserved IP changed)")
		// Mark reserved_ips as requiring a new resource
		// return d.ForceNew("reserved_ips")
		return forceNewIds()
	}

	// Rule: Any IPv6 change for single-stack IPv6 or dual-stack forces recreation,
	// unless already handled by allowed-update cases above.
	if ipv6Changed && (isIPv6Only(newVer) || isDual(newVer)) {
		log.Printf("[DEBUG] forceNewReservedIps: forcing new (IPv6 reserved IP changed for IPv6/dual-stack)")
		return forceNewIds()
	}

	log.Printf("[DEBUG] forceNewReservedIps: no ForceNew required for this change")
	// If none of the above conditions matched, do not force new.
	// Terraform will proceed with an in-place Update.
	return nil
}

func createNetworkLoadBalancerNetworkLoadBalancer(d *schema.ResourceData, m interface{}) error {
	sync := &NetworkLoadBalancerNetworkLoadBalancerResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).NetworkLoadBalancerClient()

	return tfresource.CreateResource(d, sync)
}

func readNetworkLoadBalancerNetworkLoadBalancer(d *schema.ResourceData, m interface{}) error {
	sync := &NetworkLoadBalancerNetworkLoadBalancerResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).NetworkLoadBalancerClient()

	return tfresource.ReadResource(sync)
}

func updateNetworkLoadBalancerNetworkLoadBalancer(d *schema.ResourceData, m interface{}) error {
	sync := &NetworkLoadBalancerNetworkLoadBalancerResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).NetworkLoadBalancerClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteNetworkLoadBalancerNetworkLoadBalancer(d *schema.ResourceData, m interface{}) error {
	sync := &NetworkLoadBalancerNetworkLoadBalancerResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).NetworkLoadBalancerClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type NetworkLoadBalancerNetworkLoadBalancerResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_network_load_balancer.NetworkLoadBalancerClient
	Res                    *oci_network_load_balancer.NetworkLoadBalancer
	DisableNotFoundRetries bool
}

func (s *NetworkLoadBalancerNetworkLoadBalancerResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *NetworkLoadBalancerNetworkLoadBalancerResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_network_load_balancer.LifecycleStateCreating),
	}
}

func (s *NetworkLoadBalancerNetworkLoadBalancerResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_network_load_balancer.LifecycleStateActive),
		string(oci_network_load_balancer.LifecycleStateFailed),
	}
}

func (s *NetworkLoadBalancerNetworkLoadBalancerResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_network_load_balancer.LifecycleStateDeleting),
	}
}

func (s *NetworkLoadBalancerNetworkLoadBalancerResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_network_load_balancer.LifecycleStateDeleted),
	}
}

func (s *NetworkLoadBalancerNetworkLoadBalancerResourceCrud) Create() error {
	request := oci_network_load_balancer.CreateNetworkLoadBalancerRequest{}

	if assignedIpv6, ok := s.D.GetOkExists("assigned_ipv6"); ok {
		tmp := assignedIpv6.(string)
		request.AssignedIpv6 = &tmp
	}

	if assignedPrivateIpv4, ok := s.D.GetOkExists("assigned_private_ipv4"); ok {
		tmp := assignedPrivateIpv4.(string)
		request.AssignedPrivateIpv4 = &tmp
	}

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

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if isPreserveSourceDestination, ok := s.D.GetOkExists("is_preserve_source_destination"); ok {
		tmp := isPreserveSourceDestination.(bool)
		request.IsPreserveSourceDestination = &tmp
	}

	if isPrivate, ok := s.D.GetOkExists("is_private"); ok {
		tmp := isPrivate.(bool)
		request.IsPrivate = &tmp
	}

	if isSymmetricHashEnabled, ok := s.D.GetOkExists("is_symmetric_hash_enabled"); ok {
		tmp := isSymmetricHashEnabled.(bool)
		request.IsSymmetricHashEnabled = &tmp
	}

	if networkSecurityGroupIds, ok := s.D.GetOkExists("network_security_group_ids"); ok {
		set := networkSecurityGroupIds.(*schema.Set)
		interfaces := set.List()
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("network_security_group_ids") {
			request.NetworkSecurityGroupIds = tmp
		}
	}
	if nlbIpVersion, ok := s.D.GetOkExists("nlb_ip_version"); ok {
		request.NlbIpVersion = oci_network_load_balancer.NlbIpVersionEnum(nlbIpVersion.(string))
	}
	if reservedIps, ok := s.D.GetOkExists("reserved_ips"); ok {
		interfaces := reservedIps.([]interface{})
		tmp := make([]oci_network_load_balancer.ReservedIp, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "reserved_ips", stateDataIndex)
			converted, err := s.mapToNetworkLoadBalancerReservedIp(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("reserved_ips") {
			request.ReservedIps = tmp
		}
	}

	if securityAttributes, ok := s.D.GetOkExists("security_attributes"); ok {
		convertedAttributes := tfresource.MapToSecurityAttributes(securityAttributes.(map[string]interface{}))
		request.SecurityAttributes = convertedAttributes
	}

	if subnetId, ok := s.D.GetOkExists("subnet_id"); ok {
		tmp := subnetId.(string)
		request.SubnetId = &tmp
	}

	if subnetIpv6Cidr, ok := s.D.GetOkExists("subnet_ipv6cidr"); ok {
		tmp := subnetIpv6Cidr.(string)
		request.SubnetIpv6Cidr = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "network_load_balancer")

	response, err := s.Client.CreateNetworkLoadBalancer(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	var identifier *string
	identifier = response.Id
	if identifier != nil {
		s.D.SetId(*identifier)
	}
	return s.getNetworkLoadBalancerFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "network_load_balancer"), oci_network_load_balancer.ActionTypeCreated, s.D.Timeout(schema.TimeoutCreate))
}

func (s *NetworkLoadBalancerNetworkLoadBalancerResourceCrud) getNetworkLoadBalancerFromWorkRequest(workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_network_load_balancer.ActionTypeEnum, timeout time.Duration) error {

	// Wait until it finishes
	networkLoadBalancerId, err := networkLoadBalancerWaitForWorkRequest(workId,
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.Client)

	if err != nil {
		return err
	}
	s.D.SetId(*networkLoadBalancerId)

	return s.Get()
}

func networkLoadBalancerWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
	startTime := time.Now()
	stopTime := startTime.Add(timeout)
	return func(response oci_common.OCIOperationResponse) bool {

		// Stop after timeout has elapsed
		if time.Now().After(stopTime) {
			return false
		}

		// Make sure we stop on default rules
		if tfresource.ShouldRetry(response, false, "network_load_balancer", startTime) {
			return true
		}

		// Only stop if the time Finished is set
		if workRequestResponse, ok := response.Response.(oci_network_load_balancer.GetWorkRequestResponse); ok {
			return workRequestResponse.TimeFinished == nil
		}
		return false
	}
}

func networkLoadBalancerWaitForWorkRequest(wId *string, action oci_network_load_balancer.ActionTypeEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_network_load_balancer.NetworkLoadBalancerClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "network_load_balancer")
	retryPolicy.ShouldRetryOperation = networkLoadBalancerWorkRequestShouldRetryFunc(timeout)

	response := oci_network_load_balancer.GetWorkRequestResponse{}
	stateConf := &retry.StateChangeConf{
		Pending: []string{
			string(oci_network_load_balancer.OperationStatusInProgress),
			string(oci_network_load_balancer.OperationStatusAccepted),
			string(oci_network_load_balancer.OperationStatusCanceling),
		},
		Target: []string{
			string(oci_network_load_balancer.OperationStatusSucceeded),
			string(oci_network_load_balancer.OperationStatusFailed),
			string(oci_network_load_balancer.OperationStatusCanceled),
		},
		Refresh: func() (interface{}, string, error) {
			var err error
			response, err = client.GetWorkRequest(context.Background(),
				oci_network_load_balancer.GetWorkRequestRequest{
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
		if strings.Contains(strings.ToLower(*res.EntityType), "networkloadbalancer") {
			if res.ActionType == action {
				identifier = res.Identifier
				break
			}
		}
	}

	var workRequestErr error
	if response.WorkRequest.Status == oci_network_load_balancer.OperationStatusFailed {
		errorMessage := getErrorFromNetworkLoadBalancerWorkRequest(response.WorkRequest, client, retryPolicy)
		workRequestErr = fmt.Errorf("work request did not succeed, workId: %s, action: %s. Message: %s", *wId, action, errorMessage)
	}

	return identifier, workRequestErr
}

func getErrorFromNetworkLoadBalancerWorkRequest(wr oci_network_load_balancer.WorkRequest,
	client *oci_network_load_balancer.NetworkLoadBalancerClient, retryPolicy *oci_common.RetryPolicy) string {
	// Fetch the list of work request errors
	response, err := client.ListWorkRequestErrors(context.Background(),
		oci_network_load_balancer.ListWorkRequestErrorsRequest{
			WorkRequestId: wr.Id,
			CompartmentId: wr.CompartmentId,
			RequestMetadata: oci_common.RequestMetadata{
				RetryPolicy: retryPolicy,
			},
		})

	if err != nil {
		return "Unknown failure reason"
	}

	allErrs := make([]string, 0)
	for _, wrkErr := range response.WorkRequestErrorCollection.Items {
		allErrs = append(allErrs, *wrkErr.Message)
	}
	errorMessage := strings.Join(allErrs, "\n")
	return errorMessage
}

func (s *NetworkLoadBalancerNetworkLoadBalancerResourceCrud) Get() error {
	request := oci_network_load_balancer.GetNetworkLoadBalancerRequest{}

	tmp := s.D.Id()
	request.NetworkLoadBalancerId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "network_load_balancer")

	response, err := s.Client.GetNetworkLoadBalancer(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.NetworkLoadBalancer
	return nil
}

func (s *NetworkLoadBalancerNetworkLoadBalancerResourceCrud) Update() error {
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(compartment)
			if err != nil {
				return err
			}
		}
	}

	if s.D.HasChange("network_security_group_ids") {
		err := s.updateNetworkSecurityGroups()
		if err != nil {
			return fmt.Errorf("unable to update 'network_security_group_ids', error: %v", err)
		}
	}

	request := oci_network_load_balancer.UpdateNetworkLoadBalancerRequest{}

	if assignedIpv6, ok := s.D.GetOkExists("assigned_ipv6"); ok &&
		s.D.HasChange("assigned_ipv6") {
		tmp := assignedIpv6.(string)
		request.AssignedIpv6 = &tmp
	}

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

	if isPreserveSourceDestination, ok := s.D.GetOkExists("is_preserve_source_destination"); ok {
		tmp := isPreserveSourceDestination.(bool)
		request.IsPreserveSourceDestination = &tmp
	}

	if isSymmetricHashEnabled, ok := s.D.GetOkExists("is_symmetric_hash_enabled"); ok {
		tmp := isSymmetricHashEnabled.(bool)
		request.IsSymmetricHashEnabled = &tmp
	}

	tmp := s.D.Id()
	request.NetworkLoadBalancerId = &tmp

	// 1. Guard: disallow dual-stack -> IPV4 if reserved IPv6 exists
	if s.D.HasChange("nlb_ip_version") {
		log.Printf("[DEBUG] nlb ip version has changed")
		oldRaw, newRaw := s.D.GetChange("nlb_ip_version")
		oldVer, _ := oldRaw.(string)
		newVer, _ := newRaw.(string)

		if oldVer == "IPV4_AND_IPV6" && newVer == "IPV4" {
			if reservedIps, ok := s.D.GetOkExists("reserved_ips"); ok {
				interfaces := reservedIps.([]interface{})
				for _, v := range interfaces {
					if v == nil {
						continue
					}
					mv, ok := v.(map[string]interface{})
					if !ok {
						continue
					}
					id, ok := mv["id"].(string)
					if !ok || id == "" {
						continue
					}
					if strings.HasPrefix(id, "ocid1.ipv6.") {
						log.Printf("[DEBUG] reserved ips list contains ipv6 ocid")
						return fmt.Errorf(
							"cannot update nlb_ip_version from IPV4_AND_IPV6 to IPV4 while an IPv6 reserved IP (%s) "+
								"is still configured in reserved_ips; remove the IPv6 entry first", id)
					}
				}
			}
		}
	}

	if nlbIpVersion, ok := s.D.GetOkExists("nlb_ip_version"); ok {
		request.NlbIpVersion = oci_network_load_balancer.NlbIpVersionEnum(nlbIpVersion.(string))
	}

	if reservedIps, ok := s.D.GetOkExists("reserved_ips"); ok && s.D.HasChange("reserved_ips") {
		interfaces := reservedIps.([]interface{})

		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "reserved_ips", stateDataIndex)
			converted, err := s.mapToNetworkLoadBalancerReservedIp(fieldKeyFormat)
			if err != nil {
				return err
			}

			if converted.Id != nil && strings.Contains(strings.ToLower(*converted.Id), "ipv6") {
				// Only one ipv6 is expected: set it and stop scanning
				request.ReservedIpv6Id = converted.Id
				break
			}
		}
	}

	if securityAttributes, ok := s.D.GetOkExists("security_attributes"); ok {
		convertedAttributes := tfresource.MapToSecurityAttributes(securityAttributes.(map[string]interface{}))
		request.SecurityAttributes = convertedAttributes
	}

	if subnetIpv6Cidr, ok := s.D.GetOkExists("subnet_ipv6cidr"); ok &&
		s.D.HasChange("subnet_ipv6cidr") {
		tmp := subnetIpv6Cidr.(string)
		request.SubnetIpv6Cidr = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "network_load_balancer")

	response, err := s.Client.UpdateNetworkLoadBalancer(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getNetworkLoadBalancerFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "network_load_balancer"), oci_network_load_balancer.ActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}

func (s *NetworkLoadBalancerNetworkLoadBalancerResourceCrud) Delete() error {
	request := oci_network_load_balancer.DeleteNetworkLoadBalancerRequest{}

	tmp := s.D.Id()
	request.NetworkLoadBalancerId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "network_load_balancer")

	response, err := s.Client.DeleteNetworkLoadBalancer(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	// Wait until it finishes
	_, delWorkRequestErr := networkLoadBalancerWaitForWorkRequest(workId,
		oci_network_load_balancer.ActionTypeDeleted, s.D.Timeout(schema.TimeoutDelete), s.DisableNotFoundRetries, s.Client)
	return delWorkRequestErr
}

func (s *NetworkLoadBalancerNetworkLoadBalancerResourceCrud) SetData() error {
	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	ipAddresses := []interface{}{}
	for _, item := range s.Res.IpAddresses {
		ipAddresses = append(ipAddresses, NetworkLoadBalancerIpAddressToMap(item))
	}
	s.D.Set("ip_addresses", ipAddresses)

	if s.Res.IsPreserveSourceDestination != nil {
		s.D.Set("is_preserve_source_destination", *s.Res.IsPreserveSourceDestination)
	}

	if s.Res.IsPrivate != nil {
		s.D.Set("is_private", *s.Res.IsPrivate)
	}

	if s.Res.IsSymmetricHashEnabled != nil {
		s.D.Set("is_symmetric_hash_enabled", *s.Res.IsSymmetricHashEnabled)
	}

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	networkSecurityGroupIds := []interface{}{}
	for _, item := range s.Res.NetworkSecurityGroupIds {
		networkSecurityGroupIds = append(networkSecurityGroupIds, item)
	}
	s.D.Set("network_security_group_ids", schema.NewSet(tfresource.LiteralTypeHashCodeForSets, networkSecurityGroupIds))
	s.D.Set("nlb_ip_version", s.Res.NlbIpVersion)

	//s.D.Set("security_attributes", s.Res.SecurityAttributes)
	s.D.Set("security_attributes", tfresource.SecurityAttributesToMap(s.Res.SecurityAttributes))

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.SubnetId != nil {
		s.D.Set("subnet_id", *s.Res.SubnetId)
	}

	if s.Res.SystemTags != nil {
		s.D.Set("system_tags", tfresource.SystemTagsToMap(s.Res.SystemTags))
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	return nil
}

func NetworkLoadBalancerIpAddressToMap(obj oci_network_load_balancer.IpAddress) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.IpAddress != nil {
		result["ip_address"] = string(*obj.IpAddress)
	}
	result["ip_version"] = string(obj.IpVersion)
	if obj.IsPublic != nil {
		result["is_public"] = bool(*obj.IsPublic)
	}

	if obj.ReservedIp != nil {
		result["reserved_ip"] = []interface{}{NetworkLoadBalancerReservedIpToMap(*obj.ReservedIp)}
	}

	return result
}

func NetworkLoadBalancerSummaryToMap(obj oci_network_load_balancer.NetworkLoadBalancerSummary, datasource bool) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.CompartmentId != nil {
		result["compartment_id"] = string(*obj.CompartmentId)
	}

	if obj.DefinedTags != nil {
		result["defined_tags"] = tfresource.DefinedTagsToMap(obj.DefinedTags)
	}

	if obj.DisplayName != nil {
		result["display_name"] = string(*obj.DisplayName)
	}

	result["freeform_tags"] = obj.FreeformTags

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	ipAddresses := []interface{}{}
	for _, item := range obj.IpAddresses {
		ipAddresses = append(ipAddresses, NetworkLoadBalancerIpAddressToMap(item))
	}
	result["ip_addresses"] = ipAddresses

	if obj.IsPreserveSourceDestination != nil {
		result["is_preserve_source_destination"] = bool(*obj.IsPreserveSourceDestination)
	}

	if obj.IsPrivate != nil {
		result["is_private"] = bool(*obj.IsPrivate)
	}

	if obj.IsSymmetricHashEnabled != nil {
		result["is_symmetric_hash_enabled"] = bool(*obj.IsSymmetricHashEnabled)
	}

	if obj.LifecycleDetails != nil {
		result["lifecycle_details"] = string(*obj.LifecycleDetails)
	}

	networkSecurityGroupIds := []interface{}{}
	for _, item := range obj.NetworkSecurityGroupIds {
		networkSecurityGroupIds = append(networkSecurityGroupIds, item)
	}
	if datasource {
		result["network_security_group_ids"] = networkSecurityGroupIds
	} else {
		result["network_security_group_ids"] = schema.NewSet(tfresource.LiteralTypeHashCodeForSets, networkSecurityGroupIds)
	}
	result["nlb_ip_version"] = string(obj.NlbIpVersion)

	if obj.SecurityAttributes != nil {
		result["security_attributes"] = tfresource.SecurityAttributesToMap(obj.SecurityAttributes)
	}

	result["state"] = string(obj.LifecycleState)

	if obj.SubnetId != nil {
		result["subnet_id"] = string(*obj.SubnetId)
	}

	if obj.SystemTags != nil {
		result["system_tags"] = tfresource.SystemTagsToMap(obj.SystemTags)
	}

	if obj.TimeCreated != nil {
		result["time_created"] = obj.TimeCreated.String()
	}

	if obj.TimeUpdated != nil {
		result["time_updated"] = obj.TimeUpdated.String()
	}

	return result
}

func (s *NetworkLoadBalancerNetworkLoadBalancerResourceCrud) mapToNetworkLoadBalancerReservedIp(fieldKeyFormat string) (oci_network_load_balancer.ReservedIp, error) {
	result := oci_network_load_balancer.ReservedIp{}

	if id, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "id")); ok {
		tmp := id.(string)
		result.Id = &tmp
	}

	return result, nil
}

func NetworkLoadBalancerReservedIpToMap(obj oci_network_load_balancer.ReservedIp) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	return result
}

func (s *NetworkLoadBalancerNetworkLoadBalancerResourceCrud) updateNetworkSecurityGroups() error {
	updateNsgIdsRequest := oci_network_load_balancer.UpdateNetworkSecurityGroupsRequest{}

	//@Codegen: Unless explicitly specified by the user, network_security_group_ids will not be supplied as the feature may or may not be supported
	if networkSecurityGroupIds, ok := s.D.GetOkExists("network_security_group_ids"); ok {
		set := networkSecurityGroupIds.(*schema.Set)
		interfaces := set.List()
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		updateNsgIdsRequest.NetworkSecurityGroupIds = tmp
	}

	tmp := s.D.Id()
	updateNsgIdsRequest.NetworkLoadBalancerId = &tmp

	updateNsgIdsRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "network_load_balancer")

	response, err := s.Client.UpdateNetworkSecurityGroups(context.Background(), updateNsgIdsRequest)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getNetworkLoadBalancerFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "network_load_balancer"), oci_network_load_balancer.ActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}

func (s *NetworkLoadBalancerNetworkLoadBalancerResourceCrud) updateCompartment(compartment interface{}) error {
	changeCompartmentRequest := oci_network_load_balancer.ChangeNetworkLoadBalancerCompartmentRequest{}

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	idTmp := s.D.Id()
	changeCompartmentRequest.NetworkLoadBalancerId = &idTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "network_load_balancer")

	response, err := s.Client.ChangeNetworkLoadBalancerCompartment(context.Background(), changeCompartmentRequest)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getNetworkLoadBalancerFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "network_load_balancer"), oci_network_load_balancer.ActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}
