package hw09structvalidator

import (
	"fmt"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

func fillStringRules(field, value string, vKind validateKind) (fieldRules, error) {
	rules := &stringRules{
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
		var rule stringRule
		var err error
		switch ruleName {
		case "len":
			rule, err = newStrLen(ruleValue)
		case "regexp":
			rule, err = newStrRegexp(ruleValue)
		case "in":
			rule = newStrIn(ruleValue)
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

type stringRule interface {
	validate(value string) error
}

type stringRules struct {
	field string
	vKind validateKind
	rules []stringRule
}

func (r *stringRules) fieldName() string {
	return r.field
}

func (r *stringRules) validate(errs ValidationErrors, value reflect.Value) ValidationErrors {
	switch r.vKind { //nolint:exhaustive
	case validateRegular:
		return r.validateRegular(errs, value)
	default:
		return r.validateSlice(errs, value)
	}
}

func (r *stringRules) validateRegular(errs ValidationErrors, value reflect.Value) ValidationErrors {
	val := value.String()
	for _, rule := range r.rules {
		err := rule.validate(val)
		if err != nil {
			errs = append(errs, ValidationError{Field: r.field, Err: err})
		}
	}
	return errs
}

func (r *stringRules) validateSlice(errs ValidationErrors, value reflect.Value) ValidationErrors {
	for i := 0; i < value.Len(); i++ {
		errs = r.validateRegular(errs, value.Index(i))
	}
	return errs
}

type strLen struct {
	cond int
}

func newStrLen(value string) (*strLen, error) {
	val, err := strconv.ParseInt(value, 10, 0)
	if err != nil {
		return nil, err
	}
	return &strLen{cond: int(val)}, nil
}

func (s strLen) validate(value string) error {
	if len(value) == s.cond {
		return nil
	}
	return &ErrStrLen{len(value), s.cond}
}

type ErrStrLen struct {
	Value int
	Cond  int
}

func (e *ErrStrLen) Error() string {
	return fmt.Sprintf("string length %d is not equal to required %d", e.Value, e.Cond)
}

type strRegexp struct {
	cond *regexp.Regexp
}

func newStrRegexp(value string) (*strRegexp, error) {
	rg, err := regexp.Compile(value)
	if err != nil {
		return nil, err
	}
	return &strRegexp{cond: rg}, nil
}

func (s strRegexp) validate(value string) error {
	if s.cond.MatchString(value) {
		return nil
	}
	return &ErrStrRegexp{value, s.cond}
}

type ErrStrRegexp struct {
	Value string
	Cond  *regexp.Regexp
}

func (e *ErrStrRegexp) Error() string {
	return fmt.Sprintf("string `%s` does not match the regexp `%v`", e.Value, e.Cond)
}

type strIn struct {
	cond []string
}

func newStrIn(value string) *strIn {
	val := strings.Split(value, ",")
	return &strIn{cond: val}
}

func (s strIn) validate(value string) error {
	if stringContains(s.cond, value) {
		return nil
	}
	return &ErrStrIn{value, s.cond}
}

type ErrStrIn struct {
	Value string
	Cond  []string
}

func (e *ErrStrIn) Error() string {
	return fmt.Sprintf("string `%s` is not included in the specified set %v", e.Value, e.Cond)
}
