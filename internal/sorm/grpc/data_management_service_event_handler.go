package grpc

import (
	"context"
	"github.com/google/uuid"
	"gitlab.gid.team/research-development/libraries/go-libs/pkg/log"
	"gitlab.gid.team/sso/auth/sorm/internal/sorm/entity"
	"gitlab.gid.team/sso/auth/sorm/internal/sorm/kafka"
	"gitlab.gid.team/sso/auth/sorm/internal/sorm/security"
	"google.golang.org/grpc/codes"
	"time"
)

const (
	InvalidArgument = "INVALID_ARGUMENT"
	Ok              = "OK"
	Aborted         = "ABORTED"
	Unauthenticated = "UNAUTHENTICATED"
)

type DataManagementServiceEventHandler struct {
	ctx          *context.Context
	target       entity.IEntity
	validator    entity.IValidator
	producer     kafka.IProducer
	security     security.ISecurity
	traceId      uuid.UUID
	unitTestMode bool
}

func NewDataManagementServiceEventHandler(
	ctx *context.Context, target entity.IEntity, validator entity.IValidator, producer kafka.IProducer, security security.ISecurity, traceId uuid.UUID, unitTestMode bool,
) *DataManagementServiceEventHandler {
	return &DataManagementServiceEventHandler{
		ctx:          ctx,
		target:       target,
		validator:    validator,
		producer:     producer,
		security:     security,
		traceId:      traceId,
		unitTestMode: unitTestMode,
	}
}

func (h *DataManagementServiceEventHandler) Handle() (codes.Code, string, string) {
	if ok, err := h.security.Auth(h.ctx); !ok {
		return h.Result(codes.Unauthenticated, Unauthenticated, err.Error())
	}
	if ok, info := h.validator.Validate(h.target); !ok {
		json := info.ToJson()
		return h.Result(codes.InvalidArgument, InvalidArgument, json)
	}
	if ok, err := h.producer.Send(entity.ToBytes(h.target)); !ok {
		return h.Result(codes.Aborted, Aborted, err.Error())
	}
	return h.Result(codes.OK, Ok, "")
}

func (h *DataManagementServiceEventHandler) Result(code codes.Code, message string, details string) (codes.Code, string, string) {
	if !h.unitTestMode {
		if code == codes.OK {
			log.Info(context.Background(), message, "id", uuid.New(), "correlation_id", h.target.GetCorrelationId(), "timestamp", time.Now(), "method", h.target.GetName(), "trace_id", h.traceId, "status", code)
		} else {
			log.Error(context.Background(), message, "id", uuid.New(), "correlation_id", h.target.GetCorrelationId(), "timestamp", time.Now(), "method", h.target.GetName(), "trace_id", h.traceId, "status", code, "error", details)
		}
	}
	return code, message, details
}
