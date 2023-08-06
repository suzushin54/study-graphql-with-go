package auth

import "context"

func ExtractUserName(ctx context.Context) (string, bool) {
	switch v := ctx.Value(userNameKey{}).(type) {
	case string:
		return v, true
	default:
		return "", false
	}
}
