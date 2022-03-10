package test

import (
	"context"
	"os"
	"os/signal"
	"time"

	"github.com/Deny7676yar/URL_shortener/URL_shortener1/internal/entities/linkentity"
	"github.com/Deny7676yar/URL_shortener/URL_shortener1/internal/usecase/app/repo"
	"github.com/google/uuid"
	gc "gopkg.in/check.v1"
)

type SuitBase struct {
	l repo.LinkeStore
}

func (s *SuitBase) Setlink(link repo.LinkeStore) {
	s.l = link
}

func (s *SuitBase) TestCreateLink(c *gc.C) {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	// Создание новой ссылки
	origin := &linkentity.Link{
		OriginLink: "https://example.com",
		LinkAt:     time.Now().Add(-10 * time.Hour),
	}
	_, err := s.l.Create(ctx, *origin)
	c.Assert(err, gc.IsNil)
	c.Assert(origin.LinkID, gc.Not(gc.Equals), uuid.Nil, gc.Commentf(""))

	accessedAt := time.Now().Truncate(time.Second).UTC()
	existing := &linkentity.Link{
		LinkID:     origin.LinkID,
		OriginLink: "https://example.com",
		LinkAt:     accessedAt,
	}
	_, err = s.l.Create(ctx, *existing)
	c.Assert(err, gc.IsNil)
	c.Assert(existing.LinkID, gc.Equals, origin.LinkID, gc.Commentf("link ID changed while upserting"))
	cancel()
}

// TestGetLongURL проверяет логику поиска ссылок по короткому URL
func (s *SuitBase) TestGetLongURL(c *gc.C) {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	// Создание новой ссылки
	origin := &linkentity.Link{
		OriginLink: "https://example.com",
		LinkAt:     time.Now().Add(-10 * time.Hour),
	}
	_, err := s.l.Create(ctx, *origin)
	c.Assert(err, gc.IsNil)
	c.Assert(origin.LinkID, gc.Not(gc.Equals), uuid.Nil, gc.Commentf(""))

	// Поиск ссылки по короткой ссылке
	other, err := s.l.GetLongURL(ctx, origin.ResultLink)
	c.Assert(err, gc.IsNil)
	c.Assert(other, gc.DeepEquals, origin, gc.Commentf("lookup by short link returned the wrong link"))

	// Поиск ссылки по неизвестному идентификатору
	_, err = s.l.GetLongURL(ctx, "")
	c.Assert(err, gc.Equals, true)
	cancel()
}