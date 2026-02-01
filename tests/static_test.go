package fail_test

import (
	"github.com/MintzyG/fail"
	"testing"
)

func TestStaticErrorBuilders(t *testing.T) {
	staticID := fail.ID(0, "STAT", 0, true, "StatStaticError")
	dynamicID := fail.ID(0, "STAT", 1, false, "StatDynamicError")

	fail.Register(fail.ErrorDefinition{
		ID:             staticID,
		DefaultMessage: "static message",
	})
	fail.Register(fail.ErrorDefinition{
		ID:             dynamicID,
		DefaultMessage: "dynamic message",
	})

	fail.AllowInternalLogs(true)

	t.Run("Static error should log warning on Msg", func(t *testing.T) {
		err := fail.New(staticID)
		// We can't easily capture stdout here without redirecting it,
		// but we can at least ensure it doesn't panic and we can manually verify if needed.
		err.Msg("new message")
		if err.Message != "static message" {
			t.Errorf("expected message to remain 'static message', got '%s'", err.Message)
		}
	})

	t.Run("Dynamic error should allow Msg", func(t *testing.T) {
		err := fail.New(dynamicID)
		err.Msg("new message")
		if err.Message != "new message" {
			t.Errorf("expected message to be 'new message', got '%s'", err.Message)
		}
	})
}
