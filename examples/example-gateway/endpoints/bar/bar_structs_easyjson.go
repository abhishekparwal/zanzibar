// TEMPORARY AUTOGENERATED FILE: easyjson stub code to make the package
// compilable during generation.

package  bar

import (
  "github.com/mailru/easyjson/jwriter"
  "github.com/mailru/easyjson/jlexer"
)

func ( ArgNotStructHTTPRequest ) MarshalJSON() ([]byte, error) { return nil, nil }
func (* ArgNotStructHTTPRequest ) UnmarshalJSON([]byte) error { return nil }
func ( ArgNotStructHTTPRequest ) MarshalEasyJSON(w *jwriter.Writer) {}
func (* ArgNotStructHTTPRequest ) UnmarshalEasyJSON(l *jlexer.Lexer) {}

type EasyJSON_exporter_ArgNotStructHTTPRequest *ArgNotStructHTTPRequest

func ( NormalHTTPRequest ) MarshalJSON() ([]byte, error) { return nil, nil }
func (* NormalHTTPRequest ) UnmarshalJSON([]byte) error { return nil }
func ( NormalHTTPRequest ) MarshalEasyJSON(w *jwriter.Writer) {}
func (* NormalHTTPRequest ) UnmarshalEasyJSON(l *jlexer.Lexer) {}

type EasyJSON_exporter_NormalHTTPRequest *NormalHTTPRequest

func ( TooManyArgsHTTPRequest ) MarshalJSON() ([]byte, error) { return nil, nil }
func (* TooManyArgsHTTPRequest ) UnmarshalJSON([]byte) error { return nil }
func ( TooManyArgsHTTPRequest ) MarshalEasyJSON(w *jwriter.Writer) {}
func (* TooManyArgsHTTPRequest ) UnmarshalEasyJSON(l *jlexer.Lexer) {}

type EasyJSON_exporter_TooManyArgsHTTPRequest *TooManyArgsHTTPRequest
