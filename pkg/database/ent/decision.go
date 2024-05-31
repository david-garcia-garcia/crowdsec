// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/crowdsecurity/crowdsec/pkg/database/ent/alert"
	"github.com/crowdsecurity/crowdsec/pkg/database/ent/decision"
)

// Decision is the model entity for the Decision schema.
type Decision struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Until holds the value of the "until" field.
	Until time.Time `json:"until,omitempty"`
	// Scenario holds the value of the "scenario" field.
	Scenario string `json:"scenario,omitempty"`
	// Type holds the value of the "type" field.
	Type string `json:"type,omitempty"`
	// StartIP holds the value of the "start_ip" field.
	StartIP int64 `json:"start_ip,omitempty"`
	// EndIP holds the value of the "end_ip" field.
	EndIP int64 `json:"end_ip,omitempty"`
	// StartSuffix holds the value of the "start_suffix" field.
	StartSuffix int64 `json:"start_suffix,omitempty"`
	// EndSuffix holds the value of the "end_suffix" field.
	EndSuffix int64 `json:"end_suffix,omitempty"`
	// IPSize holds the value of the "ip_size" field.
	IPSize int64 `json:"ip_size,omitempty"`
	// Scope holds the value of the "scope" field.
	Scope string `json:"scope,omitempty"`
	// Value holds the value of the "value" field.
	Value string `json:"value,omitempty"`
	// Origin holds the value of the "origin" field.
	Origin string `json:"origin,omitempty"`
	// Simulated holds the value of the "simulated" field.
	Simulated bool `json:"simulated,omitempty"`
	// UUID holds the value of the "uuid" field.
	UUID string `json:"uuid,omitempty"`
	// AlertDecisions holds the value of the "alert_decisions" field.
	AlertDecisions int `json:"alert_decisions,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the DecisionQuery when eager-loading is set.
	Edges        DecisionEdges `json:"edges"`
	selectValues sql.SelectValues
}

// DecisionEdges holds the relations/edges for other nodes in the graph.
type DecisionEdges struct {
	// Owner holds the value of the owner edge.
	Owner *Alert `json:"owner,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// OwnerOrErr returns the Owner value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e DecisionEdges) OwnerOrErr() (*Alert, error) {
	if e.Owner != nil {
		return e.Owner, nil
	} else if e.loadedTypes[0] {
		return nil, &NotFoundError{label: alert.Label}
	}
	return nil, &NotLoadedError{edge: "owner"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Decision) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case decision.FieldSimulated:
			values[i] = new(sql.NullBool)
		case decision.FieldID, decision.FieldStartIP, decision.FieldEndIP, decision.FieldStartSuffix, decision.FieldEndSuffix, decision.FieldIPSize, decision.FieldAlertDecisions:
			values[i] = new(sql.NullInt64)
		case decision.FieldScenario, decision.FieldType, decision.FieldScope, decision.FieldValue, decision.FieldOrigin, decision.FieldUUID:
			values[i] = new(sql.NullString)
		case decision.FieldCreatedAt, decision.FieldUpdatedAt, decision.FieldUntil:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Decision fields.
func (d *Decision) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case decision.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			d.ID = int(value.Int64)
		case decision.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				d.CreatedAt = value.Time
			}
		case decision.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				d.UpdatedAt = value.Time
			}
		case decision.FieldUntil:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field until", values[i])
			} else if value.Valid {
				d.Until = value.Time
			}
		case decision.FieldScenario:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field scenario", values[i])
			} else if value.Valid {
				d.Scenario = value.String
			}
		case decision.FieldType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field type", values[i])
			} else if value.Valid {
				d.Type = value.String
			}
		case decision.FieldStartIP:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field start_ip", values[i])
			} else if value.Valid {
				d.StartIP = value.Int64
			}
		case decision.FieldEndIP:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field end_ip", values[i])
			} else if value.Valid {
				d.EndIP = value.Int64
			}
		case decision.FieldStartSuffix:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field start_suffix", values[i])
			} else if value.Valid {
				d.StartSuffix = value.Int64
			}
		case decision.FieldEndSuffix:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field end_suffix", values[i])
			} else if value.Valid {
				d.EndSuffix = value.Int64
			}
		case decision.FieldIPSize:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field ip_size", values[i])
			} else if value.Valid {
				d.IPSize = value.Int64
			}
		case decision.FieldScope:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field scope", values[i])
			} else if value.Valid {
				d.Scope = value.String
			}
		case decision.FieldValue:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field value", values[i])
			} else if value.Valid {
				d.Value = value.String
			}
		case decision.FieldOrigin:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field origin", values[i])
			} else if value.Valid {
				d.Origin = value.String
			}
		case decision.FieldSimulated:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field simulated", values[i])
			} else if value.Valid {
				d.Simulated = value.Bool
			}
		case decision.FieldUUID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field uuid", values[i])
			} else if value.Valid {
				d.UUID = value.String
			}
		case decision.FieldAlertDecisions:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field alert_decisions", values[i])
			} else if value.Valid {
				d.AlertDecisions = int(value.Int64)
			}
		default:
			d.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// GetValue returns the ent.Value that was dynamically selected and assigned to the Decision.
// This includes values selected through modifiers, order, etc.
func (d *Decision) GetValue(name string) (ent.Value, error) {
	return d.selectValues.Get(name)
}

// QueryOwner queries the "owner" edge of the Decision entity.
func (d *Decision) QueryOwner() *AlertQuery {
	return NewDecisionClient(d.config).QueryOwner(d)
}

// Update returns a builder for updating this Decision.
// Note that you need to call Decision.Unwrap() before calling this method if this Decision
// was returned from a transaction, and the transaction was committed or rolled back.
func (d *Decision) Update() *DecisionUpdateOne {
	return NewDecisionClient(d.config).UpdateOne(d)
}

// Unwrap unwraps the Decision entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (d *Decision) Unwrap() *Decision {
	_tx, ok := d.config.driver.(*txDriver)
	if !ok {
		panic("ent: Decision is not a transactional entity")
	}
	d.config.driver = _tx.drv
	return d
}

// String implements the fmt.Stringer.
func (d *Decision) String() string {
	var builder strings.Builder
	builder.WriteString("Decision(")
	builder.WriteString(fmt.Sprintf("id=%v, ", d.ID))
	builder.WriteString("created_at=")
	builder.WriteString(d.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(d.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("until=")
	builder.WriteString(d.Until.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("scenario=")
	builder.WriteString(d.Scenario)
	builder.WriteString(", ")
	builder.WriteString("type=")
	builder.WriteString(d.Type)
	builder.WriteString(", ")
	builder.WriteString("start_ip=")
	builder.WriteString(fmt.Sprintf("%v", d.StartIP))
	builder.WriteString(", ")
	builder.WriteString("end_ip=")
	builder.WriteString(fmt.Sprintf("%v", d.EndIP))
	builder.WriteString(", ")
	builder.WriteString("start_suffix=")
	builder.WriteString(fmt.Sprintf("%v", d.StartSuffix))
	builder.WriteString(", ")
	builder.WriteString("end_suffix=")
	builder.WriteString(fmt.Sprintf("%v", d.EndSuffix))
	builder.WriteString(", ")
	builder.WriteString("ip_size=")
	builder.WriteString(fmt.Sprintf("%v", d.IPSize))
	builder.WriteString(", ")
	builder.WriteString("scope=")
	builder.WriteString(d.Scope)
	builder.WriteString(", ")
	builder.WriteString("value=")
	builder.WriteString(d.Value)
	builder.WriteString(", ")
	builder.WriteString("origin=")
	builder.WriteString(d.Origin)
	builder.WriteString(", ")
	builder.WriteString("simulated=")
	builder.WriteString(fmt.Sprintf("%v", d.Simulated))
	builder.WriteString(", ")
	builder.WriteString("uuid=")
	builder.WriteString(d.UUID)
	builder.WriteString(", ")
	builder.WriteString("alert_decisions=")
	builder.WriteString(fmt.Sprintf("%v", d.AlertDecisions))
	builder.WriteByte(')')
	return builder.String()
}

// Decisions is a parsable slice of Decision.
type Decisions []*Decision
