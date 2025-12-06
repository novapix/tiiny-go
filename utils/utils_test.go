package utils

import "testing"

func TestGenerateKeyLength(t *testing.T) {
	key := GenerateKey()

	if len(key) != 8 {
		t.Fatalf("expected key length 8, got %d", len(key))
	}
}

func TestGenerateKeyUniqueness(t *testing.T) {
	keys := make(map[string]bool)

	for i := 0; i < 1000; i++ {
		k := GenerateKey()
		if keys[k] {
			t.Fatalf("duplicate key generated: %s", k)
		}
		keys[k] = true
	}
}

func TestGenerateDomainName(t *testing.T) {
	domain := GenerateDomainName("localhost", 8080)

	if domain == "" {
		t.Fatal("domain should not be empty")
	}
}
