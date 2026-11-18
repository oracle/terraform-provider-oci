package core

import (
	"context"
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	oci_common "github.com/oracle/oci-go-sdk/v65/common"
	oci_core "github.com/oracle/oci-go-sdk/v65/core"
)

func TestCoreSecurityListResource_outOfBandRuleKeepsUnchangedRulesKnown(t *testing.T) {
	r := CoreSecurityListResource()
	ssh := oci_core.IngressSecurityRule{Protocol: oci_common.String("6"), Source: oci_common.String("0.0.0.0/0"), SourceType: "CIDR_BLOCK",
		TcpOptions: &oci_core.TcpOptions{DestinationPortRange: &oci_core.PortRange{Min: oci_common.Int(22), Max: oci_common.Int(22)}}}
	stateless := oci_core.IngressSecurityRule{Protocol: oci_common.String("6"), Source: oci_common.String("10.250.0.0/16"), SourceType: "CIDR_BLOCK", IsStateless: oci_common.Bool(true),
		TcpOptions: &oci_core.TcpOptions{DestinationPortRange: &oci_core.PortRange{Min: oci_common.Int(6443), Max: oci_common.Int(6443)}}}
	described := oci_core.IngressSecurityRule{Protocol: oci_common.String("all"), Source: oci_common.String("10.0.0.0/8"), SourceType: "CIDR_BLOCK", Description: oci_common.String("internal")}
	added := oci_core.IngressSecurityRule{Protocol: oci_common.String("6"), Source: oci_common.String("10.0.0.0/8"), SourceType: "CIDR_BLOCK",
		TcpOptions: &oci_core.TcpOptions{DestinationPortRange: &oci_core.PortRange{Min: oci_common.Int(9090), Max: oci_common.Int(9090)}}}

	d := r.Data(nil)
	d.SetId("ocid1.securitylist.oc1..unique")
	d.Set("compartment_id", "c")
	d.Set("vcn_id", "v")
	maps := []interface{}{}
	for _, rule := range []oci_core.IngressSecurityRule{ssh, stateless, described, added} {
		maps = append(maps, IngressSecurityRuleToMap(rule))
	}
	if err := d.Set("ingress_security_rules", schema.NewSet(ingressSecurityRulesHashCodeForSets, maps)); err != nil {
		t.Fatalf("d.Set(ingress_security_rules) error = %v", err)
	}
	state := d.State()

	// Proposed-new-state shaped config (see routeTableProposedConfig); "added" was created out of band.
	proposed := func(m map[string]interface{}) map[string]interface{} {
		full := map[string]interface{}{"description": "", "source_type": "CIDR_BLOCK", "stateless": false}
		for k, v := range m {
			full[k] = v
		}
		return full
	}
	cfg := terraform.NewResourceConfigRaw(map[string]interface{}{
		"compartment_id": "c",
		"vcn_id":         "v",
		"ingress_security_rules": []interface{}{
			proposed(map[string]interface{}{"protocol": "6", "source": "0.0.0.0/0", "tcp_options": []interface{}{map[string]interface{}{"min": 22, "max": 22}}}),
			proposed(map[string]interface{}{"protocol": "6", "source": "10.250.0.0/16", "stateless": true, "tcp_options": []interface{}{map[string]interface{}{"min": 6443, "max": 6443}}}),
			proposed(map[string]interface{}{"protocol": "all", "source": "10.0.0.0/8", "description": "internal"}),
		},
	})

	diff, err := r.SimpleDiff(context.Background(), state, cfg, nil)
	if err != nil {
		t.Fatalf("SimpleDiff() error = %v", err)
	}
	planned, err := diff.Apply(state.Attributes, r.CoreConfigSchema())
	if err != nil {
		t.Fatalf("diff.Apply() error = %v", err)
	}

	if got := planned["ingress_security_rules.#"]; got != "3" {
		t.Errorf("ingress_security_rules.# = %q, want 3", got)
	}
	attrs := func(src map[string]string, rule oci_core.IngressSecurityRule) map[string]string {
		prefix := fmt.Sprintf("ingress_security_rules.%d.", ingressSecurityRulesHashCodeForSets(IngressSecurityRuleToMap(rule)))
		out := map[string]string{}
		for k, v := range src {
			if strings.HasPrefix(k, prefix) {
				out[strings.TrimPrefix(k, prefix)] = v
			}
		}
		return out
	}
	for _, rule := range []oci_core.IngressSecurityRule{ssh, stateless, described} {
		before, after := attrs(state.Attributes, rule), attrs(planned, rule)
		for k, v := range before {
			if after[k] == sdkUnknownValue {
				t.Errorf("unchanged rule %s/%s: %s planned as unknown", *rule.Protocol, *rule.Source, k)
			} else if after[k] != v {
				t.Errorf("unchanged rule %s/%s: %s planned as %q, want %q", *rule.Protocol, *rule.Source, k, after[k], v)
			}
		}
	}
	if left := attrs(planned, added); len(left) != 0 {
		t.Errorf("out-of-band rule should be removed from the plan, got %v", left)
	}
}
