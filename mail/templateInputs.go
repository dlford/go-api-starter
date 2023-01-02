package mail

type emailVerificationInput struct {
	FirstName string
	VerifyUrl string
}

type resetPasswordInput struct {
	FirstName string
	ResetUrl  string
}
