package iam

import (
	"context"

	"github.com/caos/zitadel/internal/eventstore"

	"github.com/caos/zitadel/internal/eventstore/repository"
	"github.com/caos/zitadel/internal/repository/idpconfig"
)

const (
	IDPJWTConfigAddedEventType   eventstore.EventType = "iam.idp." + idpconfig.JWTConfigAddedEventType
	IDPJWTConfigChangedEventType eventstore.EventType = "iam.idp." + idpconfig.JWTConfigChangedEventType
)

type IDPJWTConfigAddedEvent struct {
	idpconfig.JWTConfigAddedEvent
}

func NewIDPJWTConfigAddedEvent(
	ctx context.Context,
	aggregate *eventstore.Aggregate,
	idpConfigID,
	jwtEndpoint,
	issuer,
	keysEndpoint string,
) *IDPJWTConfigAddedEvent {
	return &IDPJWTConfigAddedEvent{
		JWTConfigAddedEvent: *idpconfig.NewJWTConfigAddedEvent(
			eventstore.NewBaseEventForPush(
				ctx,
				aggregate,
				IDPJWTConfigAddedEventType,
			),
			idpConfigID,
			jwtEndpoint,
			issuer,
			keysEndpoint,
		),
	}
}

func IDPJWTConfigAddedEventMapper(event *repository.Event) (eventstore.EventReader, error) {
	e, err := idpconfig.JWTConfigAddedEventMapper(event)
	if err != nil {
		return nil, err
	}

	return &IDPJWTConfigAddedEvent{JWTConfigAddedEvent: *e.(*idpconfig.JWTConfigAddedEvent)}, nil
}

type IDPJWTConfigChangedEvent struct {
	idpconfig.JWTConfigChangedEvent
}

func NewIDPJWTConfigChangedEvent(
	ctx context.Context,
	aggregate *eventstore.Aggregate,
	idpConfigID string,
	changes []idpconfig.JWTConfigChanges,
) (*IDPJWTConfigChangedEvent, error) {
	changeEvent, err := idpconfig.NewJWTConfigChangedEvent(
		eventstore.NewBaseEventForPush(
			ctx,
			aggregate,
			IDPJWTConfigChangedEventType),
		idpConfigID,
		changes,
	)
	if err != nil {
		return nil, err
	}
	return &IDPJWTConfigChangedEvent{JWTConfigChangedEvent: *changeEvent}, nil
}

func IDPJWTConfigChangedEventMapper(event *repository.Event) (eventstore.EventReader, error) {
	e, err := idpconfig.JWTConfigChangedEventMapper(event)
	if err != nil {
		return nil, err
	}

	return &IDPJWTConfigChangedEvent{JWTConfigChangedEvent: *e.(*idpconfig.JWTConfigChangedEvent)}, nil
}