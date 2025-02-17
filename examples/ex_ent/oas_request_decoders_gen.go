// Code generated by ogen, DO NOT EDIT.

package api

import (
	"io"
	"mime"
	"net/http"

	"github.com/go-faster/errors"
	"github.com/go-faster/jx"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/multierr"

	"github.com/ogen-go/ogen/validate"
)

func (s *Server) decodeCreatePetRequest(r *http.Request, span trace.Span) (
	req CreatePetReq,
	close func() error,
	rerr error,
) {
	var closers []func() error
	close = func() error {
		var merr error
		// Close in reverse order, to match defer behavior.
		for i := len(closers) - 1; i >= 0; i-- {
			c := closers[i]
			merr = multierr.Append(merr, c())
		}
		return merr
	}
	defer func() {
		if rerr != nil {
			rerr = multierr.Append(rerr, close())
		}
	}()

	ct, _, err := mime.ParseMediaType(r.Header.Get("Content-Type"))
	if err != nil {
		return req, close, errors.Wrap(err, "parse media type")
	}
	switch ct {
	case "application/json":
		if r.ContentLength == 0 {
			return req, close, validate.ErrBodyRequired
		}

		var request CreatePetReq
		buf, err := io.ReadAll(r.Body)
		if err != nil {
			return req, close, err
		}

		if len(buf) == 0 {
			return req, close, validate.ErrBodyRequired
		}

		d := jx.DecodeBytes(buf)
		if err := func() error {
			if err := request.Decode(d); err != nil {
				return err
			}
			return nil
		}(); err != nil {
			return req, close, errors.Wrap(err, "decode \"application/json\"")
		}
		return request, close, nil
	default:
		return req, close, validate.InvalidContentType(ct)
	}
}

func (s *Server) decodeCreatePetCategoriesRequest(r *http.Request, span trace.Span) (
	req CreatePetCategoriesReq,
	close func() error,
	rerr error,
) {
	var closers []func() error
	close = func() error {
		var merr error
		// Close in reverse order, to match defer behavior.
		for i := len(closers) - 1; i >= 0; i-- {
			c := closers[i]
			merr = multierr.Append(merr, c())
		}
		return merr
	}
	defer func() {
		if rerr != nil {
			rerr = multierr.Append(rerr, close())
		}
	}()

	ct, _, err := mime.ParseMediaType(r.Header.Get("Content-Type"))
	if err != nil {
		return req, close, errors.Wrap(err, "parse media type")
	}
	switch ct {
	case "application/json":
		if r.ContentLength == 0 {
			return req, close, validate.ErrBodyRequired
		}

		var request CreatePetCategoriesReq
		buf, err := io.ReadAll(r.Body)
		if err != nil {
			return req, close, err
		}

		if len(buf) == 0 {
			return req, close, validate.ErrBodyRequired
		}

		d := jx.DecodeBytes(buf)
		if err := func() error {
			if err := request.Decode(d); err != nil {
				return err
			}
			return nil
		}(); err != nil {
			return req, close, errors.Wrap(err, "decode \"application/json\"")
		}
		return request, close, nil
	default:
		return req, close, validate.InvalidContentType(ct)
	}
}

func (s *Server) decodeCreatePetFriendsRequest(r *http.Request, span trace.Span) (
	req CreatePetFriendsReq,
	close func() error,
	rerr error,
) {
	var closers []func() error
	close = func() error {
		var merr error
		// Close in reverse order, to match defer behavior.
		for i := len(closers) - 1; i >= 0; i-- {
			c := closers[i]
			merr = multierr.Append(merr, c())
		}
		return merr
	}
	defer func() {
		if rerr != nil {
			rerr = multierr.Append(rerr, close())
		}
	}()

	ct, _, err := mime.ParseMediaType(r.Header.Get("Content-Type"))
	if err != nil {
		return req, close, errors.Wrap(err, "parse media type")
	}
	switch ct {
	case "application/json":
		if r.ContentLength == 0 {
			return req, close, validate.ErrBodyRequired
		}

		var request CreatePetFriendsReq
		buf, err := io.ReadAll(r.Body)
		if err != nil {
			return req, close, err
		}

		if len(buf) == 0 {
			return req, close, validate.ErrBodyRequired
		}

		d := jx.DecodeBytes(buf)
		if err := func() error {
			if err := request.Decode(d); err != nil {
				return err
			}
			return nil
		}(); err != nil {
			return req, close, errors.Wrap(err, "decode \"application/json\"")
		}
		return request, close, nil
	default:
		return req, close, validate.InvalidContentType(ct)
	}
}

func (s *Server) decodeCreatePetOwnerRequest(r *http.Request, span trace.Span) (
	req CreatePetOwnerReq,
	close func() error,
	rerr error,
) {
	var closers []func() error
	close = func() error {
		var merr error
		// Close in reverse order, to match defer behavior.
		for i := len(closers) - 1; i >= 0; i-- {
			c := closers[i]
			merr = multierr.Append(merr, c())
		}
		return merr
	}
	defer func() {
		if rerr != nil {
			rerr = multierr.Append(rerr, close())
		}
	}()

	ct, _, err := mime.ParseMediaType(r.Header.Get("Content-Type"))
	if err != nil {
		return req, close, errors.Wrap(err, "parse media type")
	}
	switch ct {
	case "application/json":
		if r.ContentLength == 0 {
			return req, close, validate.ErrBodyRequired
		}

		var request CreatePetOwnerReq
		buf, err := io.ReadAll(r.Body)
		if err != nil {
			return req, close, err
		}

		if len(buf) == 0 {
			return req, close, validate.ErrBodyRequired
		}

		d := jx.DecodeBytes(buf)
		if err := func() error {
			if err := request.Decode(d); err != nil {
				return err
			}
			return nil
		}(); err != nil {
			return req, close, errors.Wrap(err, "decode \"application/json\"")
		}
		return request, close, nil
	default:
		return req, close, validate.InvalidContentType(ct)
	}
}

func (s *Server) decodeUpdatePetRequest(r *http.Request, span trace.Span) (
	req UpdatePetReq,
	close func() error,
	rerr error,
) {
	var closers []func() error
	close = func() error {
		var merr error
		// Close in reverse order, to match defer behavior.
		for i := len(closers) - 1; i >= 0; i-- {
			c := closers[i]
			merr = multierr.Append(merr, c())
		}
		return merr
	}
	defer func() {
		if rerr != nil {
			rerr = multierr.Append(rerr, close())
		}
	}()

	ct, _, err := mime.ParseMediaType(r.Header.Get("Content-Type"))
	if err != nil {
		return req, close, errors.Wrap(err, "parse media type")
	}
	switch ct {
	case "application/json":
		if r.ContentLength == 0 {
			return req, close, validate.ErrBodyRequired
		}

		var request UpdatePetReq
		buf, err := io.ReadAll(r.Body)
		if err != nil {
			return req, close, err
		}

		if len(buf) == 0 {
			return req, close, validate.ErrBodyRequired
		}

		d := jx.DecodeBytes(buf)
		if err := func() error {
			if err := request.Decode(d); err != nil {
				return err
			}
			return nil
		}(); err != nil {
			return req, close, errors.Wrap(err, "decode \"application/json\"")
		}
		return request, close, nil
	default:
		return req, close, validate.InvalidContentType(ct)
	}
}
