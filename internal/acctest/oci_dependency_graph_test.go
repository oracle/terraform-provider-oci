package acctest

import (
	"testing"
)

func TestUnit_DependencyGraph(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "Validate OCI dependency graph is getting set",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Logf("Running %s", tt.name)
			InitDependencyGraph()
			if DependencyGraph["instance"] == nil {
				t.Errorf("Dependency Graph not getting set")
			}
			if DependencyGraph["loadBalancer"] == nil {
				t.Errorf("Dependency Graph not getting set")
			}
		})
	}
}
