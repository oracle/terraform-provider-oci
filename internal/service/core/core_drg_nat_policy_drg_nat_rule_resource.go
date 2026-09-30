package core

import (
	"context"
	"fmt"
	"log"
	"net/url"
	"regexp"
	"strings"
	"sync"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/oracle/oci-go-sdk/v65/common"
	oci_core "github.com/oracle/oci-go-sdk/v65/core"
	oci_work_requests "github.com/oracle/oci-go-sdk/v65/workrequests"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

var drgNatPolicyRuleMutexes drgNatPolicyRuleSafeMutexMap
var drgNatPolicyRuleReadMutexes drgNatPolicyRuleSafeMutexMap
var drgNatPolicyRuleCaches drgNatPolicyRuleCacheMap

type drgNatPolicyRuleSafeMutexMap struct {
	mutexes map[string]*sync.Mutex
	m       sync.Mutex
}

// drgNatPolicyRuleCacheMap stores successful full ListDrgNatRules results per
// DRG NAT policy for the lifetime of the current provider process.
type drgNatPolicyRuleCacheMap struct {
	rulesByPolicy map[string][]oci_core.DrgNatRule
	m             sync.Mutex
}

// GetOrCreatePolicyMutex returns a mutex for mutations on a single DRG NAT policy.
func (safeMap *drgNatPolicyRuleSafeMutexMap) GetOrCreatePolicyMutex(drgNatPolicyId string) *sync.Mutex {
	if drgNatPolicyId == "" {
		return nil
	}

	safeMap.m.Lock()
	defer safeMap.m.Unlock()

	if safeMap.mutexes == nil {
		safeMap.mutexes = map[string]*sync.Mutex{}
	}

	m, exists := safeMap.mutexes[drgNatPolicyId]
	if !exists {
		m = &sync.Mutex{}
		safeMap.mutexes[drgNatPolicyId] = m
	}

	return m
}

// Get returns a defensive copy of the cached rules for a DRG NAT policy.
func (cacheMap *drgNatPolicyRuleCacheMap) Get(drgNatPolicyId string) ([]oci_core.DrgNatRule, bool) {
	cacheMap.m.Lock()
	defer cacheMap.m.Unlock()

	if cacheMap.rulesByPolicy == nil {
		return nil, false
	}

	rules, ok := cacheMap.rulesByPolicy[drgNatPolicyId]
	if !ok {
		return nil, false
	}

	copiedRules := append([]oci_core.DrgNatRule(nil), rules...)
	return copiedRules, true
}

// Set replaces the cached full rule list for a DRG NAT policy.
func (cacheMap *drgNatPolicyRuleCacheMap) Set(drgNatPolicyId string, rules []oci_core.DrgNatRule) {
	cacheMap.m.Lock()
	defer cacheMap.m.Unlock()

	if cacheMap.rulesByPolicy == nil {
		cacheMap.rulesByPolicy = map[string][]oci_core.DrgNatRule{}
	}

	cacheMap.rulesByPolicy[drgNatPolicyId] = append([]oci_core.DrgNatRule(nil), rules...)
}

// Invalidate clears any cached rule list for the given DRG NAT policy.
func (cacheMap *drgNatPolicyRuleCacheMap) Invalidate(drgNatPolicyId string) {
	if drgNatPolicyId == "" {
		return
	}

	cacheMap.m.Lock()
	defer cacheMap.m.Unlock()

	if cacheMap.rulesByPolicy == nil {
		return
	}

	delete(cacheMap.rulesByPolicy, drgNatPolicyId)
}

// CoreDrgNatPolicyDrgNatRuleResource defines the Terraform resource schema and
// CRUD entry points for "oci_core_drg_nat_policy_drg_nat_rule".
// This resource models a single DRG NAT rule that belongs to a DRG NAT policy.
func CoreDrgNatPolicyDrgNatRuleResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts:      tfresource.DefaultTimeout,
		CreateContext: CreateCoreDrgNatPolicyDrgNatRuleWithContext,
		ReadContext:   ReadCoreDrgNatPolicyDrgNatRuleWithContext,
		UpdateContext: UpdateCoreDrgNatPolicyDrgNatRuleWithContext,
		DeleteContext: DeleteCoreDrgNatPolicyDrgNatRuleWithContext,
		Schema: map[string]*schema.Schema{
			// Required
			"drg_nat_policy_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			"drg_nat_rule_priority": {
				Type:     schema.TypeInt,
				Required: true,
			},

			// Optional
			"original_source": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"translated_source": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"original_destination": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"translated_destination": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func CreateCoreDrgNatPolicyDrgNatRuleWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &CoreDrgNatPolicyDrgNatRuleResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).VirtualNetworkClient()
	sync.WorkRequestClient = m.(*client.OracleClients).WorkRequestClient

	if e := tfresource.CreateResourceWithContext(ctx, d, sync); e != nil {
		return tfresource.HandleDiagError(m, e)
	}

	return nil
}

func ReadCoreDrgNatPolicyDrgNatRuleWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &CoreDrgNatPolicyDrgNatRuleResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).VirtualNetworkClient()

	return tfresource.HandleDiagError(m, tfresource.ReadResourceWithContext(ctx, sync))
}

func UpdateCoreDrgNatPolicyDrgNatRuleWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &CoreDrgNatPolicyDrgNatRuleResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).VirtualNetworkClient()
	sync.WorkRequestClient = m.(*client.OracleClients).WorkRequestClient

	if err := tfresource.UpdateResourceWithContext(ctx, d, sync); err != nil {
		return tfresource.HandleDiagError(m, err)
	}

	return nil

}

func DeleteCoreDrgNatPolicyDrgNatRuleWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &CoreDrgNatPolicyDrgNatRuleResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).VirtualNetworkClient()
	sync.DisableNotFoundRetries = true
	sync.WorkRequestClient = m.(*client.OracleClients).WorkRequestClient

	return tfresource.HandleDiagError(m, tfresource.DeleteResourceWithContext(ctx, d, sync))
}

type CoreDrgNatPolicyDrgNatRuleResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_core.VirtualNetworkClient
	Res                    *oci_core.DrgNatRule
	DisableNotFoundRetries bool
	WorkRequestClient      *oci_work_requests.WorkRequestClient
}

// GetMutex serializes rule mutations for a single DRG NAT policy so Terraform
// does not issue concurrent add/update/remove operations against a parent
// policy that only allows one active mutation at a time.
func (s *CoreDrgNatPolicyDrgNatRuleResourceCrud) GetMutex() *sync.Mutex {
	return drgNatPolicyRuleMutexes.GetOrCreatePolicyMutex(s.D.Get("drg_nat_policy_id").(string))
}

// ID returns the Terraform resource ID from the schema. This ID is a composite
// of DRG NAT policy OCID and rule OCID (see GetDrgNatPolicyDrgNatRuleCompositeId).
func (s *CoreDrgNatPolicyDrgNatRuleResourceCrud) ID() string {
	return s.D.Id()
}

// getDrgNatRuleMutationRetryPolicy retries only the transient 409 cases that
// occur when the parent DRG NAT policy is still applying a previous mutation.
func getDrgNatRuleMutationRetryPolicy(timeoutKey string, d *schema.ResourceData) *common.RetryPolicy {
	return tfresource.GetRetryPolicyWithAdditionalRetryCondition(d.Timeout(timeoutKey), func(response common.OCIOperationResponse) bool {
		failure, isServiceError := common.IsServiceError(response.Error)
		if !isServiceError || failure.GetHTTPStatusCode() != 409 {
			return false
		}

		message := strings.ToLower(failure.GetMessage())
		return strings.Contains(message, "invalid lifecycle state") || strings.Contains(message, "updating")
	}, "core")
}

// listDrgNatRules returns the full rule list for a DRG NAT policy, serving it
// from the in-process cache when available to avoid repeated identical list calls.
func (s *CoreDrgNatPolicyDrgNatRuleResourceCrud) listDrgNatRules(ctx context.Context, drgNatPolicyId string) ([]oci_core.DrgNatRule, error) {
	if cachedRules, ok := drgNatPolicyRuleCaches.Get(drgNatPolicyId); ok {
		return cachedRules, nil
	}

	// Serialize cache misses per policy so concurrent refreshes do not all issue
	// the same initial ListDrgNatRules request before the cache is populated.
	readMutex := drgNatPolicyRuleReadMutexes.GetOrCreatePolicyMutex(drgNatPolicyId)
	readMutex.Lock()
	defer readMutex.Unlock()

	if cachedRules, ok := drgNatPolicyRuleCaches.Get(drgNatPolicyId); ok {
		return cachedRules, nil
	}

	req := oci_core.ListDrgNatRulesRequest{
		DrgNatPolicyId: &drgNatPolicyId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "core"),
		},
	}

	var rules []oci_core.DrgNatRule
	for {
		resp, err := s.Client.ListDrgNatRules(ctx, req)
		if err != nil {
			return nil, err
		}
		rules = append(rules, resp.Items...)
		if resp.OpcNextPage == nil {
			break
		}
		req.Page = resp.OpcNextPage
	}

	drgNatPolicyRuleCaches.Set(drgNatPolicyId, rules)
	return append([]oci_core.DrgNatRule(nil), rules...), nil
}

// CreateWithContext implements creation of a DRG NAT rule.
// It sends an AddDrgNatRules request, waits for the policy to return to ACTIVE
// (via CheckDrgNatPolicyState), and then looks up the created rule by policy
// and priority to populate the composite ID and Res.
//
// If the rule cannot be found after the policy returns to ACTIVE, the resource
// ID is left empty and the resource is treated as absent.
func (s *CoreDrgNatPolicyDrgNatRuleResourceCrud) CreateWithContext(ctx context.Context) error {
	log.Printf("[DEBUG] CreateWithContext: start, Id=%q, drg_nat_policy_id=%v, priority=%v",
		s.D.Id(), s.D.Get("drg_nat_policy_id"), s.D.Get("drg_nat_rule_priority"))
	request := oci_core.AddDrgNatRulesRequest{}

	if drgNatPolicyId, ok := s.D.GetOkExists("drg_nat_policy_id"); ok {
		tmp := drgNatPolicyId.(string)
		request.DrgNatPolicyId = &tmp
	}

	addDrgNatRuleDetails := oci_core.AddDrgNatRuleDetails{}

	if drgNatRulePriority, ok := s.D.GetOkExists("drg_nat_rule_priority"); ok {
		tmp := int64(drgNatRulePriority.(int))
		addDrgNatRuleDetails.DrgNatRulePriority = &tmp
	}

	if originalSource, ok := s.D.GetOkExists("original_source"); ok {
		tmp := originalSource.(string)
		addDrgNatRuleDetails.OriginalSource = &tmp
	}

	if translatedSource, ok := s.D.GetOkExists("translated_source"); ok {
		tmp := translatedSource.(string)
		addDrgNatRuleDetails.TranslatedSource = &tmp
	}

	if originalDestination, ok := s.D.GetOkExists("original_destination"); ok {
		tmp := originalDestination.(string)
		addDrgNatRuleDetails.OriginalDestination = &tmp
	}

	if translatedDestination, ok := s.D.GetOkExists("translated_destination"); ok {
		tmp := translatedDestination.(string)
		addDrgNatRuleDetails.TranslatedDestination = &tmp
	}

	tmp := []oci_core.AddDrgNatRuleDetails{addDrgNatRuleDetails}
	request.Rules = tmp

	request.RequestMetadata.RetryPolicy = getDrgNatRuleMutationRetryPolicy(schema.TimeoutCreate, s.D)

	log.Printf("[DEBUG] CreateWithContext: AddDrgNatRules request: %+v", request)
	_, err := s.Client.AddDrgNatRules(ctx, request)
	if err != nil {
		log.Printf("[DEBUG] CreateWithContext: AddDrgNatRules error: %v", err)
		return err
	}

	// The service may already have applied the mutation, so do not reuse the
	// pre-mutation cached rule list on retries if the lifecycle wait fails.
	drgNatPolicyRuleCaches.Invalidate(*request.DrgNatPolicyId)

	// First ensure the NAT policy returns to ACTIVE. The rule itself has no
	// standalone lifecycle, so we use the policy state as the async completion signal.
	if err := s.CheckDrgNatPolicyState(ctx); err != nil {
		return err
	}

	// Immediately fetch the created rule and populate Id/Res. This is required
	// because the AddDrgNatRules API does not directly return the rule OCID.
	log.Printf("[DEBUG] CreateWithContext: calling GetWithContext to populate Id/Res after create")
	if err := s.GetWithContext(ctx); err != nil {
		return err
	}

	log.Printf("[DEBUG] CreateWithContext: done, Id=%q", s.D.Id())
	return nil
}

// SetData updates Terraform state fields from the last-read rule (Res).
// If Res is nil, the ID is cleared, which tells Terraform the resource no
// longer exists and should be dropped from state.
func (s *CoreDrgNatPolicyDrgNatRuleResourceCrud) SetData() error {
	log.Printf("[DEBUG] SetData: entry, Id=%q, ResNil=%v", s.D.Id(), s.Res == nil)

	if s.Res == nil {
		// When Res is nil, we explicitly clear the ID so Terraform treats this
		// resource as destroyed/absent.
		log.Printf("[DEBUG] SetData: Res is nil, clearing Id")
		s.D.SetId("")
		return nil
	}

	drgNatPolicyId, drgNatRuleId, err := ParseDrgNatPolicyDrgNatRuleCompositeId(s.D.Id())
	if err != nil {
		log.Printf("[WARN] SetData: unable to parse current ID: %s (err=%v)", s.D.Id(), err)
		return err
	}

	log.Printf("[DEBUG] SetData: parsed Id, policy=%q rule=%q", drgNatPolicyId, drgNatRuleId)

	s.D.Set("drg_nat_policy_id", drgNatPolicyId)
	s.D.SetId(GetDrgNatPolicyDrgNatRuleCompositeId(drgNatPolicyId, drgNatRuleId))

	if s.Res.DrgNatRulePriority != nil {
		s.D.Set("drg_nat_rule_priority", int(*s.Res.DrgNatRulePriority))
	} else {
		s.D.Set("drg_nat_rule_priority", nil)
	}

	s.D.Set("original_source", valueOrNilString(s.Res.OriginalSource))
	s.D.Set("original_destination", valueOrNilString(s.Res.OriginalDestination))
	s.D.Set("translated_source", valueOrNilString(s.Res.TranslatedSource))
	s.D.Set("translated_destination", valueOrNilString(s.Res.TranslatedDestination))

	return nil
}

func valueOrNilString(p *string) interface{} {
	if p == nil {
		return nil
	}
	return *p
}

// drgNatPolicyWaiter is a lightweight waiter that repeatedly calls
// GetDrgNatPolicy and caches the last lifecycleState and error.
// It is used by CheckDrgNatPolicyState with WaitForResourceConditionWithContext.
type drgNatPolicyWaiter struct {
	client         *oci_core.VirtualNetworkClient
	policyId       string
	disableRetries bool

	lifecycleState oci_core.DrgNatPolicyLifecycleStateEnum
	fetchErr       error
}

// GetWithContext implements the polling step for drgNatPolicyWaiter.
// It calls GetDrgNatPolicy and updates lifecycleState and fetchErr.
// On error, it records the error but returns nil so that the condition
// function can decide whether to continue waiting or fail.
func (w *drgNatPolicyWaiter) GetWithContext(ctx context.Context) error {
	req := oci_core.GetDrgNatPolicyRequest{
		DrgNatPolicyId: &w.policyId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: tfresource.GetRetryPolicy(w.disableRetries, "core"),
		},
	}

	resp, err := w.client.GetDrgNatPolicy(ctx, req)
	if err != nil {
		w.fetchErr = err
		w.lifecycleState = ""
		// Returning nil here ensures WaitForResourceConditionWithContext continues
		// to evaluate the condition function, which inspects fetchErr.
		return nil
	}

	w.fetchErr = nil
	w.lifecycleState = resp.DrgNatPolicy.LifecycleState
	return nil
}

// CheckDrgNatPolicyState waits until the associated DRG NAT policy returns
// to the ACTIVE lifecycle state. The async DRG NAT rule operations do not
// expose a separate rule lifecycle, so the policy lifecycle acts as the
// completion indicator.
func (s *CoreDrgNatPolicyDrgNatRuleResourceCrud) CheckDrgNatPolicyState(ctx context.Context) error {
	policyVal, policyOk := s.D.GetOk("drg_nat_policy_id")
	if !policyOk {
		return fmt.Errorf("cannot determine drg_nat_policy_id for state check")
	}
	drgNatPolicyId := policyVal.(string)

	// Use the local waiter that only calls GetDrgNatPolicy.
	waiter := &drgNatPolicyWaiter{
		client:         s.Client,
		policyId:       drgNatPolicyId,
		disableRetries: s.DisableNotFoundRetries,
	}

	// Condition function: WaitForResourceConditionWithContext will call
	// waiter.GetWithContext(ctx) before each evaluation, and we stop waiting
	// when the policy reaches ACTIVE and there is no fetch error.
	cond := func() bool {
		if waiter.fetchErr != nil {
			return false
		}
		return waiter.lifecycleState == oci_core.DrgNatPolicyLifecycleStateActive
	}

	// IMPORTANT: pass `waiter`, not `s`, to avoid recursion into the rule's GetWithContext.
	return tfresource.WaitForResourceConditionWithContext(
		ctx,
		waiter,
		cond,
		s.D.Timeout(schema.TimeoutUpdate),
	)
}

// GetWithContext implements read semantics for the DRG NAT rule.
//
// There are two cases:
//  1. If the resource ID is empty (typically right after Create), it looks
//     up the rule by policy ID and drg_nat_rule_priority, paginating through
//     all ListDrgNatRules pages. If found, it sets a composite ID and Res.
//     If not found across all pages, it clears ID and treats the rule as gone.
//  2. If the resource has a composite ID, it parses policy/rule IDs, lists
//     all rules for that policy (with pagination), and finds the rule by its
//     OCID. If not found, the ID is cleared and the resource is removed
//     from state.
func (s *CoreDrgNatPolicyDrgNatRuleResourceCrud) GetWithContext(ctx context.Context) error {

	log.Printf("[DEBUG] GetWithContext: start, Id=%q, drg_nat_policy_id=%v, priority=%v",
		s.D.Id(), s.D.Get("drg_nat_policy_id"), s.D.Get("drg_nat_rule_priority"))

	// Case 1: ID is empty – this happens right after Create
	if s.D.Id() == "" {
		log.Printf("[DEBUG] GetWithContext: empty Id, trying lookup by policy/priority")

		policyVal, policyOk := s.D.GetOk("drg_nat_policy_id")
		prioVal, prioOk := s.D.GetOk("drg_nat_rule_priority")
		if !policyOk || !prioOk {
			return fmt.Errorf("resource ID is not set and cannot determine policy/prio")
		}
		drgNatPolicyId := policyVal.(string)
		critPriority := prioVal.(int)

		rules, err := s.listDrgNatRules(ctx, drgNatPolicyId)
		if err != nil {
			return err
		}

		for _, r := range rules {
			if r.DrgNatRulePriority != nil && int(*r.DrgNatRulePriority) == critPriority {
				// Found our rule – set a proper composite ID and Res
				s.D.SetId(GetDrgNatPolicyDrgNatRuleCompositeId(drgNatPolicyId, *r.Id))
				s.Res = &r
				return nil
			}
		}

		log.Printf("[DEBUG] GetWithContext: no rule found by policy/prio across all pages, treating as gone")
		// Not found – treat as gone
		s.D.SetId("")
		s.Res = nil
		return nil
	}

	// Case 2: we already have a composite ID – parse and list by policy
	drgNatPolicyId, drgNatRuleId, err := ParseDrgNatPolicyDrgNatRuleCompositeId(s.D.Id())
	if err != nil {
		log.Printf("[WARN] GetWithContext() unable to parse current ID: %s", s.D.Id())
		s.D.SetId("")
		s.Res = nil
		return nil
	}

	log.Printf("[DEBUG] GetWithContext: parsed Id, policy=%q rule=%q", drgNatPolicyId, drgNatRuleId)
	rules, err := s.listDrgNatRules(ctx, drgNatPolicyId)
	if err != nil {
		return err
	}

	var found *oci_core.DrgNatRule
	for i := range rules {
		r := &rules[i]
		if r.Id != nil && *r.Id == drgNatRuleId {
			found = r
			break
		}
	}

	if found == nil {
		// Not found; resource is gone
		log.Printf("[DEBUG] GetWithContext: rule %q not found; clearing Id", drgNatRuleId)
		s.D.SetId("")
		s.Res = nil
		return nil
	}

	log.Printf("[DEBUG] GetWithContext: found rule %q; Res set", drgNatRuleId)

	s.Res = found
	return nil

}

func (s *CoreDrgNatPolicyDrgNatRuleResourceCrud) DeleteWithContext(ctx context.Context) error {
	request := oci_core.RemoveDrgNatRulesRequest{}

	drgNatPolicyId, drgNatRuleId, err := ParseDrgNatPolicyDrgNatRuleCompositeId(s.D.Id())
	if err != nil {
		return err
	}

	request.DrgNatPolicyId = &drgNatPolicyId
	request.RuleIds = []string{drgNatRuleId}
	request.RequestMetadata.RetryPolicy = getDrgNatRuleMutationRetryPolicy(schema.TimeoutDelete, s.D)

	log.Printf("[DEBUG] DeleteWithContext: RemoveDrgNatRules request: %+v", request)
	_, err = s.Client.RemoveDrgNatRules(ctx, request)
	if err != nil {
		return fmt.Errorf("failed to remove nat rules, error: %v", err)
	}

	// The service may already have applied the mutation, so do not reuse the
	// pre-mutation cached rule list on retries if the lifecycle wait fails.
	drgNatPolicyRuleCaches.Invalidate(drgNatPolicyId)

	// Wait for the policy to return to ACTIVE after rule removal
	log.Printf("[DEBUG] DeleteWithContext: waiting for DRG NAT policy to become ACTIVE")
	if err := s.CheckDrgNatPolicyState(ctx); err != nil {
		return err
	}

	return nil
}

func (s *CoreDrgNatPolicyDrgNatRuleResourceCrud) UpdateWithContext(ctx context.Context) error {
	request := oci_core.UpdateDrgNatRulesRequest{}

	drgNatPolicyId, drgNatRuleId, err := ParseDrgNatPolicyDrgNatRuleCompositeId(s.D.Id())
	if err == nil {
		request.DrgNatPolicyId = &drgNatPolicyId
		s.D.Set("drg_nat_policy_id", drgNatPolicyId)
	} else {
		return err
	}

	updateDrgNatRuleDetails := oci_core.UpdateDrgNatRuleDetails{}
	updateDrgNatRuleDetails.Id = &drgNatRuleId

	if drgNatRulePriority, ok := s.D.GetOkExists("drg_nat_rule_priority"); ok && s.D.HasChange("drg_nat_rule_priority") {
		tmp := int64(drgNatRulePriority.(int))
		updateDrgNatRuleDetails.DrgNatRulePriority = &tmp
	}

	if originalSource, ok := s.D.GetOkExists("original_source"); ok && s.D.HasChange("original_source") {
		tmp := originalSource.(string)
		updateDrgNatRuleDetails.OriginalSource = &tmp
	}

	if originalDestination, ok := s.D.GetOkExists("original_destination"); ok && s.D.HasChange("original_destination") {
		tmp := originalDestination.(string)
		updateDrgNatRuleDetails.OriginalDestination = &tmp
	}

	if translatedSource, ok := s.D.GetOkExists("translated_source"); ok && s.D.HasChange("translated_source") {
		tmp := translatedSource.(string)
		updateDrgNatRuleDetails.TranslatedSource = &tmp
	}

	if translatedDestination, ok := s.D.GetOkExists("translated_destination"); ok && s.D.HasChange("translated_destination") {
		tmp := translatedDestination.(string)
		updateDrgNatRuleDetails.TranslatedDestination = &tmp
	}

	tmp := []oci_core.UpdateDrgNatRuleDetails{updateDrgNatRuleDetails}
	request.Rules = tmp

	request.RequestMetadata.RetryPolicy = getDrgNatRuleMutationRetryPolicy(schema.TimeoutUpdate, s.D)
	_, err = s.Client.UpdateDrgNatRules(ctx, request)
	if err != nil {
		return fmt.Errorf("failed to Update nat rules, error: %v", err)
	}

	// The service may already have applied the mutation, so do not reuse the
	// pre-mutation cached rule list on retries if the lifecycle wait fails.
	drgNatPolicyRuleCaches.Invalidate(drgNatPolicyId)

	// Wait for the policy to return to ACTIVE
	log.Printf("[DEBUG] UpdateWithContext: waiting for DRG NAT policy to become ACTIVE")
	if err := s.CheckDrgNatPolicyState(ctx); err != nil {
		return err
	}

	// Refresh the resource so SetData sees the latest values
	log.Printf("[DEBUG] UpdateWithContext: calling GetWithContext to refresh after update")
	if err := s.GetWithContext(ctx); err != nil {
		return err
	}

	log.Printf("[DEBUG] UpdateWithContext: done, Id=%q", s.D.Id())
	return nil
}

func GetDrgNatPolicyDrgNatRuleCompositeId(drgNatPolicyId string, drgNatRuleId string) string {
	drgNatPolicyId = url.PathEscape(drgNatPolicyId)
	drgNatRuleId = url.PathEscape(drgNatRuleId)
	compositeId := "drgNatPolicies/" + drgNatPolicyId + "/drgNatRules/" + drgNatRuleId
	return compositeId
}

func ParseDrgNatPolicyDrgNatRuleCompositeId(compositeId string) (drgNatPolicyId string, drgNatRuleId string, err error) {
	parts := strings.Split(compositeId, "/")
	match, _ := regexp.MatchString("drgNatPolicies/.*/drgNatRules/.*", compositeId)
	if !match || len(parts) != 4 {
		err = fmt.Errorf("illegal compositeId %s encountered", compositeId)
		return
	}
	drgNatPolicyId, _ = url.PathUnescape(parts[1])
	drgNatRuleId, _ = url.PathUnescape(parts[3])

	return
}
