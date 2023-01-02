package router

import (
	"api/controllers/appkey"
	"api/controllers/otp"
	"api/controllers/session"
	"api/controllers/user"

	"github.com/gin-gonic/gin"
)

func userRouter(r *gin.Engine) {
	u := r.Group("/user")
	u.GET("/", user.GetCurrentUser)
	u.POST("/", user.CreateUser)
	u.DELETE("/", user.DeleteUser)
	u.PUT("/", user.UpdateUser)

	ue := u.Group("/email")
	ue.GET("/", user.VerifyEmailAddress)
	ue.PATCH("/", user.ResendEmailVerification)

	up := u.Group("/password")
	up.PUT("/", user.UpdatePassword)
	up.POST("/", user.ResetPassword)
	up.PATCH("/", user.ForgotPassword)

	us := u.Group("/sessions")
	us.GET("/", session.ListSessions)
	us.POST("/", session.SignIn)
	us.PATCH("/", session.RefreshToken)
	us.DELETE("/", session.SignOut)
	us.DELETE("/:id", session.DeleteSession)

	uo := u.Group("/otp")
	uo.GET("/", otp.SetupOtp)
	uo.PATCH("/", otp.EnableOtp)
	uo.DELETE("/", otp.DisableOtp)

	ua := u.Group("/appkeys")
	ua.GET("/", appkey.ListAppkeys)
	ua.POST("/", appkey.CreateAppkey)
	ua.DELETE("/:id", appkey.DeleteAppkey)
}
