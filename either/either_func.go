package either

import base "soundcream/funcgemon"

func Bind[R, R2, L any](e base.Either[R, L], fun func(right R) base.Either[R2, L]) base.Either[R2, L] {
	if e.IsLeft() {
		return base.Either[R2, L]{Left: e.Left}
	}
	return fun(*e.Right)
}

func Map[R, R2, L any](e base.Either[R, L], fun func(left *L, right *R) base.Either[R2, L]) base.Either[R2, L] {
	if e.IsLeft() {
		return base.Either[R2, L]{Left: e.Left}
	}
	return fun(e.Left, e.Right)
}

type Either[R, L any] struct {
	Right *R
	Left  *L
}

func (e Either[R, L]) IsLeft() bool {
	return e.Left != nil
}

func (e Either[R, L]) Then(fn func(R) Either[R, L]) Either[R, L] {
	if e.IsLeft() {
		return Either[R, L]{Left: e.Left}
	}
	return fn(*e.Right)
}
