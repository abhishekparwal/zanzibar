// Copyright (c) 2018 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package zanzibar

import (
	"context"
	"time"

	"github.com/pborman/uuid"
	"github.com/uber-go/tally"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type contextFieldKey string

// ContextScopeTagsExtractor defines func where extracts tags from context
type ContextScopeTagsExtractor func(context.Context) map[string]string

// ContextLogFieldsExtractor defines func where extracts log fields from context
type ContextLogFieldsExtractor func(context.Context) []zap.Field

const (
	endpointKey           = contextFieldKey("endpoint")
	requestUUIDKey        = contextFieldKey("requestUUID")
	routingDelegateKey    = contextFieldKey("rd")
	endpointRequestHeader = contextFieldKey("endpointRequestHeader")
	requestLogFields      = contextFieldKey("requestLogFields")
	scopeTags             = contextFieldKey("scopeTags")
)

const (
	logFieldRequestMethod       = "method"
	logFieldRequestURL          = "url"
	logFieldRequestStartTime    = "timestamp-started"
	logFieldRequestFinishedTime = "timestamp-finished"
	logFieldRequestHeaderPrefix = "Request-Header"
	logFieldResponseStatusCode  = "statusCode"
	logFieldRequestUUID         = "requestUUID"
	logFieldEndpointID          = "endpointID"
	logFieldHandlerID           = "handlerID"
)

const (
	scopeTagClientMethod    = "clientmethod"
	scopeTagEndpointMethod  = "endpointmethod"
	scopeTagClient          = "clientid"
	scopeTagEndpoint        = "endpointid"
	scopeTagHandler         = "handlerid"
	scopeTagError           = "error"
	scopeTagStatus          = "status"
	scopeTagProtocol        = "protocol"
	scopeTagHTTP            = "HTTP"
	scopeTagTChannel        = "TChannel"
	scopeTagsTargetService  = "targetservice"
	scopeTagsTargetEndpoint = "targetendpoint"
)

// WithEndpointField adds the endpoint information in the
// request context.
func WithEndpointField(ctx context.Context, endpoint string) context.Context {
	return context.WithValue(ctx, endpointKey, endpoint)
}

// GetRequestEndpointFromCtx returns the endpoint, if it exists on context
func GetRequestEndpointFromCtx(ctx context.Context) string {
	if val := ctx.Value(endpointKey); val != nil {
		endpoint, _ := val.(string)
		return endpoint
	}
	return ""
}

// WithEndpointRequestHeadersField adds the endpoint request header information in the
// request context.
func WithEndpointRequestHeadersField(ctx context.Context, requestHeaders map[string]string) context.Context {
	headers := GetEndpointRequestHeadersFromCtx(ctx)
	for k, v := range requestHeaders {
		headers[k] = v
	}

	return context.WithValue(ctx, endpointRequestHeader, headers)
}

// GetEndpointRequestHeadersFromCtx returns the endpoint request headers, if it exists on context
func GetEndpointRequestHeadersFromCtx(ctx context.Context) map[string]string {
	requestHeaders := make(map[string]string)
	if val := ctx.Value(endpointRequestHeader); val != nil {
		headers, _ := val.(map[string]string)
		for k, v := range headers {
			requestHeaders[k] = v
		}
	}

	return requestHeaders
}

// withRequestUUID returns a context with request uuid.
func withRequestUUID(ctx context.Context, reqUUID uuid.UUID) context.Context {
	return context.WithValue(ctx, requestUUIDKey, reqUUID)
}

// GetRequestUUIDFromCtx returns the RequestUUID, if it exists on context
// TODO: in future, we can extend this to have request object
func GetRequestUUIDFromCtx(ctx context.Context) uuid.UUID {
	if val := ctx.Value(requestUUIDKey); val != nil {
		uuid, _ := val.(uuid.UUID)
		return uuid
	}
	return nil
}

// WithRoutingDelegate adds the tchannel routing delegate information in the
// request context.
func WithRoutingDelegate(ctx context.Context, rd string) context.Context {
	return context.WithValue(ctx, routingDelegateKey, rd)
}

// GetRoutingDelegateFromCtx returns the tchannel routing delegate info
// extracted from context.
func GetRoutingDelegateFromCtx(ctx context.Context) string {
	if val := ctx.Value(routingDelegateKey); val != nil {
		rd, _ := val.(string)
		return rd
	}
	return ""
}

// WithLogFields returns a new context with the given log fields attached to context.Context
func WithLogFields(ctx context.Context, newFields ...zap.Field) context.Context {
	return context.WithValue(ctx, requestLogFields, accumulateLogFields(ctx, newFields))
}

func getLogFieldsFromCtx(ctx context.Context) []zap.Field {
	var fields []zap.Field
	v := ctx.Value(requestLogFields)
	if v != nil {
		fields = v.([]zap.Field)
	}
	return fields
}

// WithScopeTags returns a new context with the given scope tags attached to context.Context
func WithScopeTags(ctx context.Context, newFields map[string]string) context.Context {
	fields := GetScopeTagsFromCtx(ctx)
	for k, v := range newFields {
		fields[k] = v
	}

	return context.WithValue(ctx, scopeTags, fields)
}

// GetScopeTagsFromCtx returns the tag info extracted from context.
func GetScopeTagsFromCtx(ctx context.Context) map[string]string {
	fields := make(map[string]string)
	if val := ctx.Value(scopeTags); val != nil {
		headers, _ := val.(map[string]string)
		for k, v := range headers {
			fields[k] = v
		}
	}

	return fields
}

func accumulateLogFields(ctx context.Context, newFields []zap.Field) []zap.Field {
	previousFields := getLogFieldsFromCtx(ctx)
	return append(previousFields, newFields...)
}

// ContextExtractor is a extractor that extracts some log fields from the context
type ContextExtractor interface {
	ExtractScopeTags(ctx context.Context) map[string]string
	ExtractLogFields(ctx context.Context) []zap.Field
}

// AddContextScopeTagsExtractor added a scope tags extractor to contextExtractor.
func (c *ContextExtractors) AddContextScopeTagsExtractor(extractors ...ContextScopeTagsExtractor) {
	c.contextScopeExtractors = extractors
}

// AddContextLogFieldsExtractor added a log fields extractor to contextExtractor.
func (c *ContextExtractors) AddContextLogFieldsExtractor(extractors ...ContextLogFieldsExtractor) {
	c.contextLogFieldsExtractors = extractors
}

// MakeContextExtractor returns a extractor that extracts log fields a context.
func (c *ContextExtractors) MakeContextExtractor() ContextExtractor {
	return &ContextExtractors{
		contextScopeExtractors:     c.contextScopeExtractors,
		contextLogFieldsExtractors: c.contextLogFieldsExtractors,
	}
}

// ContextExtractors warps extractors for context
type ContextExtractors struct {
	contextScopeExtractors     []ContextScopeTagsExtractor
	contextLogFieldsExtractors []ContextLogFieldsExtractor
}

// ExtractScopeTags extracts scope fields from a context into a tag.
func (c *ContextExtractors) ExtractScopeTags(ctx context.Context) map[string]string {
	tags := make(map[string]string)
	for _, fn := range c.contextScopeExtractors {
		sc := fn(ctx)
		for k, v := range sc {
			tags[k] = v
		}
	}

	return tags
}

// ExtractLogFields extracts log fields from a context into a field.
func (c *ContextExtractors) ExtractLogFields(ctx context.Context) []zap.Field {
	var fields []zap.Field
	for _, fn := range c.contextLogFieldsExtractors {
		logFields := fn(ctx)
		fields = append(fields, logFields...)

	}

	return fields
}

// ContextLogger is a logger that extracts some log fields from the context before passing through to underlying zap logger.
type ContextLogger interface {
	Debug(ctx context.Context, msg string, fields ...zap.Field)
	Error(ctx context.Context, msg string, fields ...zap.Field)
	Info(ctx context.Context, msg string, fields ...zap.Field)
	Panic(ctx context.Context, msg string, fields ...zap.Field)
	Warn(ctx context.Context, msg string, fields ...zap.Field)

	// Other utility methods on the logger
	Check(lvl zapcore.Level, msg string) *zapcore.CheckedEntry
}

// NewContextLogger returns a logger that extracts log fields a context before passing through to underlying zap logger.
func NewContextLogger(log *zap.Logger) ContextLogger {
	return &contextLogger{
		log: log,
	}
}

type contextLogger struct {
	log *zap.Logger
}

func (c *contextLogger) Debug(ctx context.Context, msg string, userFields ...zap.Field) {
	c.log.Debug(msg, accumulateLogFields(ctx, userFields)...)
}

func (c *contextLogger) Error(ctx context.Context, msg string, userFields ...zap.Field) {
	c.log.Error(msg, accumulateLogFields(ctx, userFields)...)
}

func (c *contextLogger) Info(ctx context.Context, msg string, userFields ...zap.Field) {
	c.log.Info(msg, accumulateLogFields(ctx, userFields)...)
}

func (c *contextLogger) Panic(ctx context.Context, msg string, userFields ...zap.Field) {
	c.log.Panic(msg, accumulateLogFields(ctx, userFields)...)
}

func (c *contextLogger) Warn(ctx context.Context, msg string, userFields ...zap.Field) {
	c.log.Warn(msg, accumulateLogFields(ctx, userFields)...)
}

func (c *contextLogger) Check(lvl zapcore.Level, msg string) *zapcore.CheckedEntry {
	return c.log.Check(lvl, msg)
}

// Logger is a generic logger interface that zap.Logger implements.
type Logger interface {
	Debug(msg string, fields ...zap.Field)
	Error(msg string, fields ...zap.Field)
	Info(msg string, fields ...zap.Field)
	Panic(msg string, fields ...zap.Field)
	Warn(msg string, fields ...zap.Field)
	Check(lvl zapcore.Level, msg string) *zapcore.CheckedEntry
}

// newLoggerWithFields creates a lightweight logger that implements Logger interface
func newLoggerWithFields(logger *zap.Logger, fields []zap.Field) Logger {
	return &loggerWithFields{
		logger: logger,
		fields: fields,
	}
}

// loggerWithFields is a logger that logs entries with default fields, it is not
// a child logger and does not clone, therefore suitable for per request creation.
type loggerWithFields struct {
	logger *zap.Logger
	fields []zap.Field
}

func (l *loggerWithFields) Debug(msg string, userFields ...zap.Field) {
	l.logger.Debug(msg, append(l.fields, userFields...)...)
}

func (l *loggerWithFields) Error(msg string, userFields ...zap.Field) {
	l.logger.Error(msg, append(l.fields, userFields...)...)
}

func (l *loggerWithFields) Info(msg string, userFields ...zap.Field) {
	l.logger.Info(msg, append(l.fields, userFields...)...)
}

func (l *loggerWithFields) Panic(msg string, userFields ...zap.Field) {
	l.logger.Panic(msg, append(l.fields, userFields...)...)
}

func (l *loggerWithFields) Warn(msg string, userFields ...zap.Field) {
	l.logger.Warn(msg, append(l.fields, userFields...)...)
}

func (l *loggerWithFields) Check(lvl zapcore.Level, msg string) *zapcore.CheckedEntry {
	return l.logger.Check(lvl, msg)
}

// ContextMetrics emit metrics with tags extracted from context.
type ContextMetrics interface {
	IncCounter(ctx context.Context, name string, value int64)
	RecordTimer(ctx context.Context, name string, d time.Duration)
}

type contextMetrics struct {
	scope tally.Scope
}

// NewContextMetrics create ContextMetrics to emit metrics with tags extracted from context.
func NewContextMetrics(scope tally.Scope) ContextMetrics {
	return &contextMetrics{
		scope: scope,
	}
}

// IncCounter increments the counter with current tags from context
func (c *contextMetrics) IncCounter(ctx context.Context, name string, value int64) {
	tags := GetScopeTagsFromCtx(ctx)
	c.scope.Tagged(tags).Counter(name).Inc(value)
}

// RecordTimer records the duration with current tags from context
func (c *contextMetrics) RecordTimer(ctx context.Context, name string, d time.Duration) {
	tags := GetScopeTagsFromCtx(ctx)
	c.scope.Tagged(tags).Timer(name).Record(d)
}
