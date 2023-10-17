package eraspacelog

import (
	"errors"
	"strings"

	otel "github.com/erajayatech/go-opentelemetry"
	"github.com/sirupsen/logrus"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
)

// Ref to https://github.com/open-telemetry/opentelemetry-specification/blob/main/specification/logs/overview.md#json-formats
const (
	traceIDKey       = "trace_id"
	spanIDKey        = "span_id"
	traceFlagsKey    = "trace_flags"
	logEventKey      = "log"
	requestHeaderKey = "request_header"
	authHeaderKey    = "auth_header"
)

var (
	logSeverityTextKey = attribute.Key("otel.logrus.severity.text")
	logMessageKey      = attribute.Key("otel.logrus.message")
)

type TraceHookConfig struct {
	RecordStackTraceInSpan bool
	EnableLevels           []logrus.Level
	ErrorSpanLevel         logrus.Level
}

type OtelTraceHook struct {
	cfg *TraceHookConfig
}

func NewOtelTraceHook(cfg *TraceHookConfig) *OtelTraceHook {
	return &OtelTraceHook{cfg: cfg}
}

func (h *OtelTraceHook) Levels() []logrus.Level {
	return h.cfg.EnableLevels
}

func (h *OtelTraceHook) Fire(entry *logrus.Entry) error {
	if entry.Context == nil {
		return nil
	}

	entry.Data[requestHeaderKey] = GetRequestHeaderInfoFromContext(entry.Context)
	entry.Data[authHeaderKey] = GetAuthHeaderInfoFromContext(entry.Context)

	span := otel.SpanFromContext(entry.Context)
	if !span.IsRecording() {
		return nil
	}

	// attach span context to log entry data fields
	entry.Data[traceIDKey] = span.SpanContext().TraceID()
	entry.Data[spanIDKey] = span.SpanContext().SpanID()
	entry.Data[traceFlagsKey] = span.SpanContext().TraceFlags()

	// attach log to span event attributes
	attrs := []attribute.KeyValue{
		logMessageKey.String(entry.Message),
		logSeverityTextKey.String(OtelSeverityText(entry.Level)),
	}
	span.AddEvent(logEventKey, trace.WithAttributes(attrs...))

	// set span status
	if entry.Level <= h.cfg.ErrorSpanLevel {
		span.SetStatus(codes.Error, entry.Message)
		span.RecordError(errors.New(entry.Message), trace.WithStackTrace(h.cfg.RecordStackTraceInSpan))
	}

	return nil
}

// OtelSeverityText convert logrus level to otel severityText
// ref to https://github.com/open-telemetry/opentelemetry-specification/blob/main/specification/logs/data-model.md#severity-fields
func OtelSeverityText(lv logrus.Level) string {
	s := lv.String()
	if s == "warning" {
		s = "warn"
	}
	return strings.ToUpper(s)
}
