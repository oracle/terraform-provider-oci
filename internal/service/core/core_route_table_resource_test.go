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

// sdkUnknownValue is the flatmap placeholder SDKv2 uses for "(known after apply)".
const sdkUnknownValue = "74D93920-ED26-11E3-AC10-0800200C9A66"

func routeTableStateFromRules(t *testing.T, rules []oci_core.RouteRule) *terraform.InstanceState {
	t.Helper()
	d := CoreRouteTableResource().Data(nil)
	d.SetId("ocid1.routetable.oc1..unique")
	d.Set("compartment_id", "ocid1.compartment.oc1..unique")
	d.Set("vcn_id", "ocid1.vcn.oc1..unique")
	maps := []interface{}{}
	for _, rule := range rules {
		maps = append(maps, RouteRuleToMap(rule))
	}
	if err := d.Set("route_rules", schema.NewSet(routeRulesHashCodeForSets, maps)); err != nil {
		t.Fatalf("d.Set(route_rules) error = %v", err)
	}
	return d.State()
}

// routeTableProposedConfig mimics core's proposed new state, where unset attributes carry "" instead of being absent.
func routeTableProposedConfig(rules ...map[string]interface{}) *terraform.ResourceConfig {
	list := []interface{}{}
	for _, rule := range rules {
		full := map[string]interface{}{
			"cidr_block":       "",
			"description":      "",
			"destination_type": "CIDR_BLOCK",
			"route_type":       "STATIC",
		}
		for k, v := range rule {
			full[k] = v
		}
		list = append(list, full)
	}
	return terraform.NewResourceConfigRaw(map[string]interface{}{
		"compartment_id": "ocid1.compartment.oc1..unique",
		"vcn_id":         "ocid1.vcn.oc1..unique",
		"route_rules":    list,
	})
}

// planRouteTable returns the planned flatmap attributes, built the same way the SDK does.
func planRouteTable(t *testing.T, state *terraform.InstanceState, cfg *terraform.ResourceConfig) map[string]string {
	t.Helper()
	r := CoreRouteTableResource()
	diff, err := r.SimpleDiff(context.Background(), state, cfg, nil)
	if err != nil {
		t.Fatalf("SimpleDiff() error = %v", err)
	}
	planned, err := diff.Apply(state.Attributes, r.CoreConfigSchema())
	if err != nil {
		t.Fatalf("diff.Apply() error = %v", err)
	}
	return planned
}

func routeRuleAttrs(planned map[string]string, code int) map[string]string {
	prefix := fmt.Sprintf("route_rules.%d.", code)
	out := map[string]string{}
	for k, v := range planned {
		if strings.HasPrefix(k, prefix) {
			out[strings.TrimPrefix(k, prefix)] = v
		}
	}
	return out
}

func routeRuleCode(rule oci_core.RouteRule) int {
	return routeRulesHashCodeForSets(RouteRuleToMap(rule))
}

func TestCoreRouteTableResource_outOfBandRuleKeepsUnchangedRulesKnown(t *testing.T) {
	igw := oci_core.RouteRule{Destination: oci_common.String("0.0.0.0/0"), DestinationType: "CIDR_BLOCK", NetworkEntityId: oci_common.String("igw"), RouteType: "STATIC"}
	peer := oci_core.RouteRule{Destination: oci_common.String("192.168.0.0/16"), DestinationType: "CIDR_BLOCK", NetworkEntityId: oci_common.String("drg"), RouteType: "STATIC", Description: oci_common.String("peer")}
	added := oci_core.RouteRule{Destination: oci_common.String("10.0.0.0/8"), DestinationType: "CIDR_BLOCK", NetworkEntityId: oci_common.String("drg"), RouteType: "STATIC"}
	state := routeTableStateFromRules(t, []oci_core.RouteRule{igw, peer, added})

	// "added" was created out of band, so it is not in the config.
	planned := planRouteTable(t, state, routeTableProposedConfig(
		map[string]interface{}{"destination": "0.0.0.0/0", "network_entity_id": "igw"},
		map[string]interface{}{"destination": "192.168.0.0/16", "network_entity_id": "drg", "description": "peer"},
	))

	if got := planned["route_rules.#"]; got != "2" {
		t.Errorf("route_rules.# = %q, want 2", got)
	}
	for _, rule := range []oci_core.RouteRule{igw, peer} {
		code := routeRuleCode(rule)
		before := routeRuleAttrs(state.Attributes, code)
		after := routeRuleAttrs(planned, code)
		for k, v := range before {
			if after[k] == sdkUnknownValue {
				t.Errorf("unchanged rule %s: %s planned as unknown", *rule.Destination, k)
			} else if after[k] != v {
				t.Errorf("unchanged rule %s: %s planned as %q, want %q", *rule.Destination, k, after[k], v)
			}
		}
	}
	if left := routeRuleAttrs(planned, routeRuleCode(added)); len(left) != 0 {
		t.Errorf("out-of-band rule should be removed from the plan, got %v", left)
	}
}

func TestCoreRouteTableResource_echoedCidrBlockDoesNotChurn(t *testing.T) {
	// The service echoes cidr_block for rules created with it, even after the config moves to destination.
	sticky := oci_core.RouteRule{CidrBlock: oci_common.String("10.112.0.0/16"), Destination: oci_common.String("10.112.0.0/16"), DestinationType: "CIDR_BLOCK", NetworkEntityId: oci_common.String("lpg"), RouteType: "STATIC"}
	igw := oci_core.RouteRule{Destination: oci_common.String("0.0.0.0/0"), DestinationType: "CIDR_BLOCK", NetworkEntityId: oci_common.String("igw"), RouteType: "STATIC"}
	added := oci_core.RouteRule{Destination: oci_common.String("10.120.0.0/16"), DestinationType: "CIDR_BLOCK", NetworkEntityId: oci_common.String("lpg"), RouteType: "STATIC"}
	state := routeTableStateFromRules(t, []oci_core.RouteRule{sticky, igw, added})

	planned := planRouteTable(t, state, routeTableProposedConfig(
		map[string]interface{}{"destination": "10.112.0.0/16", "network_entity_id": "lpg"},
		map[string]interface{}{"destination": "0.0.0.0/0", "network_entity_id": "igw"},
	))

	after := routeRuleAttrs(planned, routeRuleCode(sticky))
	if got := after["cidr_block"]; got != "10.112.0.0/16" {
		t.Errorf("echoed cidr_block planned as %q, want it kept as 10.112.0.0/16", got)
	}
	if left := routeRuleAttrs(planned, routeRuleCode(added)); len(left) != 0 {
		t.Errorf("out-of-band rule should be removed from the plan, got %v", left)
	}
}

func TestCoreRouteTableResource_removedRuleWithEchoedCidrBlockIsFullyRemoved(t *testing.T) {
	keep := oci_core.RouteRule{Destination: oci_common.String("0.0.0.0/0"), DestinationType: "CIDR_BLOCK", NetworkEntityId: oci_common.String("igw"), RouteType: "STATIC"}
	removed := oci_core.RouteRule{CidrBlock: oci_common.String("10.112.0.0/16"), Destination: oci_common.String("10.112.0.0/16"), DestinationType: "CIDR_BLOCK", NetworkEntityId: oci_common.String("lpg"), RouteType: "STATIC"}
	state := routeTableStateFromRules(t, []oci_core.RouteRule{keep, removed})

	planned := planRouteTable(t, state, routeTableProposedConfig(
		map[string]interface{}{"destination": "0.0.0.0/0", "network_entity_id": "igw"},
	))

	if got := planned["route_rules.#"]; got != "1" {
		t.Errorf("route_rules.# = %q, want 1", got)
	}
	if left := routeRuleAttrs(planned, routeRuleCode(removed)); len(left) != 0 {
		t.Errorf("removed rule must not leave a partial element in the plan, got %v", left)
	}
}
