package timek

import (
	"encoding/json"

	"github.com/dromara/carbon/v2"
)

// MarshalJSON implements json.Marshaler interface
func (t Time) MarshalJSON() ([]byte, error) {
	if t.IsZero() {
		return json.Marshal(nil)
	}
	// Default JSON format is usually standard ISO8601/RFC3339
	return json.Marshal(t.ToDateTimeString())
}

// UnmarshalJSON implements json.Unmarshaler interface
func (t *Time) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}

	if s == "" {
		*t = Time{Carbon: *carbon.NewCarbon()}
		return nil
	}

	*t = Parse(s)
	return nil
}
