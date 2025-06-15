package adapters

import (
	"errors"
	"net/http"
	"os"
	"time"

	authCore "github.com/edlingao/go-auth/auth/core"
	calculatorCore "github.com/edlingao/hexago/internal/calculator/core"
	"github.com/edlingao/hexago/internal/users/core"
	"github.com/edlingao/hexago/internal/users/ports/driven"
	"github.com/edlingao/hexago/lib"
	"github.com/edlingao/hexago/web/views/auth"
	"github.com/edlingao/hexago/web/views/users"
	"github.com/labstack/echo/v4"
)

type UsersWebService struct {
	URL               string
	CalculatorService *calculatorCore.Calculator
	http              *echo.Group
	sessionService    authCore.SessionService
	usersService      core.UserService
	dbService         driven.StoringUsers[core.User]
}

func NewUsersWebService(
	url string,
	httpService *echo.Group,
	sessionService authCore.SessionService,
	dbService driven.StoringUsers[core.User],
	usersService core.UserService,
	calculatorService *calculatorCore.Calculator,
) *UsersWebService {

	usersWebService := &UsersWebService{
		URL:               url,
		CalculatorService: calculatorService,
		http:              httpService,
		sessionService:    sessionService,
		dbService:         dbService,
		usersService:      usersService,
	}
	// Public routes
	usersWebService.http.GET("/login", usersWebService.Login)
	usersWebService.http.GET("/signup", usersWebService.SignUp)
	usersWebService.http.POST("/login", usersWebService.LoginEndpoint)
	usersWebService.http.POST("/register", usersWebService.SignUpEndpoint)

	// Protected routes
	protectedAPI := usersWebService.http.Group("", sessionService.APIAuth)
	protectedAPI.GET("/all", usersWebService.GetAllUsers)

	protectedWeb := usersWebService.http.Group("", sessionService.WebAuth)
	protectedWeb.GET("/home", usersWebService.Home)

	return usersWebService
}

func (uws *UsersWebService) GetAllUsers(c echo.Context) error {
	return nil
}

func (uws *UsersWebService) Login(c echo.Context) error {
	return lib.Render(
		c,
		auth.SignIn(auth.SignInVM{}),
		200,
	)
}

func (uws *UsersWebService) SignUp(c echo.Context) error {
	return lib.Render(
		c,
		auth.Register(auth.RegisterVM{}),
		200,
	)
}

func (uws *UsersWebService) LoginEndpoint(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")

	user, err := uws.usersService.SignIn(username, password)

	if err != nil {
		return lib.Render(
			c,
			auth.SignIn(auth.SignInVM{
				Error: err,
			}),
			400,
		)
	}

	token, err := uws.sessionService.Create(user.ID, user.Username, os.Getenv("JWT_SECRET"))

	if err != nil {
		return lib.Render(
			c,
			auth.SignIn(auth.SignInVM{
				Error: err,
			}),
			500,
		)
	}
	cookie := uws.SetCookie(token.Token, c)
	c.SetCookie(cookie)
	c.Response().Header().Set("HX-Location", "/home")

	return lib.Render(
		c,
		auth.SignIn(auth.SignInVM{}),
		200,
	)
}

func (uws *UsersWebService) SignUpEndpoint(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")
	if username == "" || password == "" {
		return lib.Render(
			c,
			auth.SignIn(auth.SignInVM{
				Error: errors.New("Username and password are required"),
			}),
			400,
		)
	}

	err := uws.usersService.Register(username, password)
	if err != nil {
		return lib.Render(
			c,
			auth.SignIn(auth.SignInVM{
				Error: err,
			}),
			400,
		)
	}

	user, err := uws.usersService.GetByUsername(username)
	if err != nil {
		return lib.Render(
			c,
			auth.SignIn(auth.SignInVM{
				Error: err,
			}),
			400,
		)
	}

	token, err := uws.sessionService.Create(user.ID, user.Username, os.Getenv("JWT_SECRET"))
	if err != nil {
		return lib.Render(
			c,
			auth.SignIn(auth.SignInVM{
				Error: err,
			}),
			500,
		)
	}

	cookie := uws.SetCookie(token.Token, c)
	c.SetCookie(cookie)

	c.Response().Header().Set("HX-Location", "/home")

	return lib.Render(
		c,
		auth.SignIn(auth.SignInVM{}),
		200,
	)

}

func (uws *UsersWebService) Home(c echo.Context) error {
	history := uws.CalculatorService.GetAllCalculations()

	return lib.Render(
		c,
		users.Home(users.HomeVM{
			History: history,
		}),
		200,
	)
}

func (uh UsersWebService) SetCookie(key string, c echo.Context) *http.Cookie {
	secure := os.Getenv("ENVIRONMENT") == "prod"
	cookie := &http.Cookie{
		Name:     "Auth",
		Value:    key,
		Expires:  time.Now().Add(24 * time.Hour),
		HttpOnly: true,
		Secure:   secure, // Set to true in production with HTTPS
		SameSite: http.SameSiteLaxMode,
		Path:     "/",
	}

	return cookie
}
