package commonexport

import (
	"sync"
	"testing"
)

func TestUnitCheckDuplicateResourceName(t *testing.T) {
	ResourceNameCount = make(map[string]int)
	t.Run("Concurrency", func(t *testing.T) {
		names := []string{"resource", "resource_1", "resource_2", "resource", "resource", "resource_1"}
		var wg sync.WaitGroup
		wg.Add(len(names))

		results := make(chan string, len(names))

		for _, name := range names {
			go func(name string) {
				defer wg.Done()
				results <- CheckDuplicateResourceName(name)
			}(name)
		}

		wg.Wait()
		close(results)

		// Map to track unique results
		uniqueResults := make(map[string]bool)

		for result := range results {
			t.Logf("result: %v\n", result)
			if _, exists := uniqueResults[result]; exists {
				t.Errorf("duplicate name found: %v", result)
			} else {
				uniqueResults[result] = true
			}
		}

	})
}
