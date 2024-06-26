package hw10programoptimization

import (
	json "encoding/json"

	easyjson "github.com/mailru/easyjson"
	jlexer "github.com/mailru/easyjson/jlexer"
	jwriter "github.com/mailru/easyjson/jwriter"
)

var (
	_ *json.RawMessage
	_ *jlexer.Lexer
	_ *jwriter.Writer
	_ easyjson.Marshaler
)

func easyjsonE3ab7953DecodeGithubComAnfilatOtusGoHw10ProgramOptimization(in *jlexer.Lexer, out *User) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "ID":
			out.ID = in.Int()
		case "Name":
			out.Name = in.String()
		case "Username":
			out.Username = in.String()
		case "Email":
			out.Email = in.String()
		case "Phone":
			out.Phone = in.String()
		case "Password":
			out.Password = in.String()
		case "Address":
			out.Address = in.String()
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}

func easyjsonE3ab7953EncodeGithubComAnfilatOtusGoHw10ProgramOptimization(out *jwriter.Writer, in User) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"ID\":"
		out.RawString(prefix[1:])
		out.Int(in.ID)
	}
	{
		const prefix string = ",\"Name\":"
		out.RawString(prefix)
		out.String(in.Name)
	}
	{
		const prefix string = ",\"Username\":"
		out.RawString(prefix)
		out.String(in.Username)
	}
	{
		const prefix string = ",\"Email\":"
		out.RawString(prefix)
		out.String(in.Email)
	}
	{
		const prefix string = ",\"Phone\":"
		out.RawString(prefix)
		out.String(in.Phone)
	}
	{
		const prefix string = ",\"Password\":"
		out.RawString(prefix)
		out.String(in.Password)
	}
	{
		const prefix string = ",\"Address\":"
		out.RawString(prefix)
		out.String(in.Address)
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface.
func (v User) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonE3ab7953EncodeGithubComAnfilatOtusGoHw10ProgramOptimization(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface.
func (v User) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonE3ab7953EncodeGithubComAnfilatOtusGoHw10ProgramOptimization(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface.
func (v *User) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonE3ab7953DecodeGithubComAnfilatOtusGoHw10ProgramOptimization(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface.
func (v *User) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonE3ab7953DecodeGithubComAnfilatOtusGoHw10ProgramOptimization(l, v)
}
