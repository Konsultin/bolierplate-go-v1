package timek

import (
	"database/sql/driver"

	"github.com/dromara/carbon/v2"
)

// Value implements driver.Valuer interface for database serialization
func (t Time) Value() (driver.Value, error) {
	if t.IsZero() {
		return nil, nil
	}
	return t.ToTime(), nil
}

// Scan implements sql.Scanner interface for database deserialization
func (t *Time) Scan(value interface{}) error {
	if value == nil {
		*t = Time{Carbon: *carbon.NewCarbon()}
		return nil
	}

	c := carbon.NewCarbon()
	if err := c.Scan(value); err != nil {
		return err
	}
	if c.Error != nil {
		return c.Error
	}
	*t = Time{Carbon: *c.SetLanguage(carbon.NewLanguage().SetLocale(LangID))}
	return nil
}
