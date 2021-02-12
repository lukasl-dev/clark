package lexing

import "errors"

var (
  // ErrNoPrefix occurs when no prefix was found.
  ErrNoPrefix = errors.New("no prefix found")

  // ErrNoLabel occurs when no label was found.
  ErrNoLabel = errors.New("no label found")
)
