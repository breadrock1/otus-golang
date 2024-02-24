package hw09structvalidator

import (
	"encoding/json"
	"fmt"
	"testing"
)

type UserRole string

// Test the function on different structures and other types.
type (
	User struct {
		ID     string `json:"id" validate:"len:36"`
		Name   string
		Age    int             `validate:"min:18|max:50"`
		Email  string          `validate:"regexp:^\\w+@\\w+\\.\\w+$"`
		Role   UserRole        `validate:"in:admin,stuff"`
		Phones []string        `validate:"len:11"`
		meta   json.RawMessage //nolint:unused
	}

	App struct {
		Version string `validate:"len:5"`
	}

	Token struct {
		Header    []byte
		Payload   []byte
		Signature []byte
	}

	Response struct {
		Code int    `validate:"in:200,404,500"`
		Body string `json:"omitempty"`
	}

	PrivateField struct {
		Login    string
		password string `validate:"in:123456,pass"`
	}

	Ints struct {
		IntField   int   `validate:"min:1|max:100"`
		Int8Field  int8  `validate:"min:1|max:100"`
		Int16Field int16 `validate:"min:1|max:100"`
		Int32Field int32 `validate:"min:1|max:100"`
		Int64Field int64 `validate:"min:1|max:100"`
	}
)

var user = User{
	ID:     "012345678901234567890123456789123456",
	Name:   "Somebody",
	Age:    20,
	Email:  "test@mail.ru",
	Role:   "admin",
	Phones: []string{"79270000000"},
	meta:   []byte("{}"),
}

var app = App{
	Version: "12345",
}

var token = Token{
	Header:    []byte("12345"),
	Payload:   []byte("12345"),
	Signature: []byte("12345"),
}

var response = Response{
	Code: 200,
	Body: "content",
}

var privateField = PrivateField{
	Login:    "somebody",
	password: "dv740Z_I!hrU&aW11dWYbrQ$t$QHez1*r@x%`WBU",
}

var ints = Ints{
	IntField:   42,
	Int8Field:  42,
	Int16Field: 42,
	Int32Field: 42,
	Int64Field: 42,
}

var emptyUser = User{}

var wrongUser = User{
	ID:     "012345678",
	Age:    51,
	Email:  "test.mail.ru",
	Role:   "hacker",
	Phones: []string{"03"},
}

var wrongResponse = Response{
	Code: 418,
}

var wrongInts = Ints{
	IntField:   0,
	Int8Field:  0,
	Int16Field: 0,
	Int32Field: 0,
	Int64Field: 100500,
}

func TestValidate(t *testing.T) {
	tests := []struct {
		name string
		in   interface{}
	}{
		{
			name: "nil",
			in:   nil,
		},
		{
			name: "User",
			in:   user,
		},
		{
			name: "&User",
			in:   &user,
		},
		{
			name: "App",
			in:   app,
		},
		{
			name: "Token",
			in:   token,
		},
		{
			name: "Response",
			in:   response,
		},
		{
			name: "PrivateField",
			in:   privateField,
		},
		{
			name: "Ints",
			in:   ints,
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("case %d", i), func(t *testing.T) {
			tt := tt
			t.Parallel()

			// Place your code here.
			_ = tt
		})
	}
}

func TestValidateFail(t *testing.T) {
}

func TestErrIncorrectUse(t *testing.T) {
}

func BenchmarkValidateSuccess(b *testing.B) {
}

func BenchmarkValidateFail(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_ = Validate(emptyUser)
			_ = Validate(wrongUser)
			_ = Validate(wrongResponse)
			_ = Validate(wrongInts)
		}
	})
}
