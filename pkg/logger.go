package pkg

import (
	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
)

func LogError(err error) string {
	refCode := uuid.New().String()
	log.Error().Str("ref_code", refCode).Err(err).Msg("an error occurred")
	return refCode
}
