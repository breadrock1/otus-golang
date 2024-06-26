package hw09structvalidator

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

func fillIntRules(field, value string, vKind validateKind) (fieldRules, error) {
	rules := &intRules{
		field: field,
		vKind: vKind,
	}
	strs := strings.Split(value, "|")
	for _, str := range strs {
		pair := strings.Split(str, ":")
		if len(pair) != 2 {
			return nil, &ErrIncorrectUse{reason: IncorrectCondition, field: field, rule: str}
		}

		ruleName := pair[0]
		ruleValue := pair[1]
		var rule intRule
		var err error
		switch ruleName {
		case "min":
			rule, err = newIntMin(ruleValue)
		case "max":
			rule, err = newIntMax(ruleValue)
		case "in":
			rule, err = newIntIn(ruleValue)
		default:
			return nil, &ErrIncorrectUse{reason: UnknownRule, field: field, rule: ruleName}
		}
		if err != nil {
			return nil, &ErrIncorrectUse{reason: IncorrectCondition, field: field, rule: ruleName, err: err}
		}
		rules.rules = append(rules.rules, rule)
	}
	if len(rules.rules) == 0 {
		return nil, nil
	}
	return rules, nil
}

type intRule interface {
	validate(value int64) error
}

type intRules struct {
	field string
	vKind validateKind
	rules []intRule
}

func (r *intRules) fieldName() string {
	return r.field
}

func (r *intRules) validate(errs ValidationErrors, value reflect.Value) ValidationErrors {
	switch r.vKind { //nolint:exhaustive
	case validateRegular:
		return r.validateRegular(errs, value)
	default:
		return r.validateSlice(errs, value)
	}
}

func (r *intRules) validateRegular(errs ValidationErrors, value reflect.Value) ValidationErrors {
	val := value.Int()
	for _, rule := range r.rules {
		err := rule.validate(val)
		if err != nil {
			errs = append(errs, ValidationError{Field: r.field, Err: err})
		}
	}
	return errs
}

func (r *intRules) validateSlice(errs ValidationErrors, value reflect.Value) ValidationErrors {
	for i := 0; i < value.Len(); i++ {
		errs = r.validateRegular(errs, value.Index(i))
	}
	return errs
}

type intMin struct {
	cond int64
}

func newIntMin(value string) (*intMin, error) {
	val, err := strconv.ParseInt(value, 10, 0)
	if err != nil {
		return nil, err
	}
	return &intMin{cond: val}, nil
}

func (s intMin) validate(value int64) error {
	if value >= s.cond {
		return nil
	}
	return &ErrIntMin{value, s.cond}
}

type ErrIntMin struct {
	Value int64
	Cond  int64
}

func (e *ErrIntMin) Error() string {
	return fmt.Sprintf("number %d is less than specified %d", e.Value, e.Cond)
}

type intMax struct {
	cond int64
}

func newIntMax(value string) (*intMax, error) {
	val, err := strconv.ParseInt(value, 10, 0)
	if err != nil {
		return nil, err
	}
	return &intMax{cond: val}, nil
}

func (s intMax) validate(value int64) error {
	if value <= s.cond {
		return nil
	}
	return &ErrIntMax{value, s.cond}
}

type ErrIntMax struct {
	Value int64
	Cond  int64
}

func (e *ErrIntMax) Error() string {
	return fmt.Sprintf("number %d is greater than specified %d", e.Value, e.Cond)
}

type intIn struct {
	cond []int64
}

func newIntIn(value string) (*intIn, error) {
	val, err := strsToInts64(strings.Split(value, ","))
	if err != nil {
		return nil, err
	}
	return &intIn{cond: val}, nil
}

func (s intIn) validate(value int64) error {
	if intContains(s.cond, value) {
		return nil
	}
	return &ErrIntIn{value, s.cond}
}

type ErrIntIn struct {
	Value int64
	Cond  []int64
}

func (e *ErrIntIn) Error() string {
	return fmt.Sprintf("number %d is not included in the specified set %v", e.Value, e.Cond)
}
