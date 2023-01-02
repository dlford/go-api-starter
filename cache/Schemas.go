package cache

import "net/url"

var AccessTokenTag = "access-token"

func MakeAccessTokenKey(userId string, sessionId string, token string) string {
	return userId + ":" + sessionId + ":" + AccessTokenTag + ":" + token
}

var EmailVerificationTokenTag = "email-verification"

func MakeEmailVerificationTokenKey(userId string, email string) string {
	return userId + ":" + url.QueryEscape(email) + ":" + EmailVerificationTokenTag
}

var PasswordResetTokenTag = "password-reset"

func MakePasswordResetTokenKey(userId string) string {
	return userId + ":" + PasswordResetTokenTag
}

var AppkeyTag = "appkey"

func MakeAppkeyKey(prefix string) string {
	return AppkeyTag + ":" + prefix
}
