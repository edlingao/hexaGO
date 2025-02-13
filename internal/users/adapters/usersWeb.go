package adapters

import (
	authCore "github.com/edlingao/go-auth/auth/core"
	"github.com/edlingao/hexago/internal/users/core"
	"github.com/edlingao/hexago/internal/users/ports/driven"
	"github.com/edlingao/hexago/lib"
	"github.com/edlingao/hexago/web/views/auth"
	"github.com/edlingao/hexago/web/views/users"
	"github.com/labstack/echo/v4"
)

type UsersWebService struct {
  URL         string
  http        *echo.Group
  sessionService authCore.SessionService
  usersService core.UserService
  dbService   driven.StoringUsers[core.User]
}

func NewUsersWebService(
  url string,
  httpService *echo.Group,
  sessionService authCore.SessionService,
  dbService driven.StoringUsers[core.User],
  usersService core.UserService,
) *UsersWebService {

  usersWebService := &UsersWebService{
    URL:         url,
    http:        httpService,
    sessionService: sessionService,
    dbService:   dbService,
    usersService: usersService,
  }

  usersWebService.http.GET("/login", usersWebService.Login)
  usersWebService.http.GET("/register", usersWebService.SignUp)

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

func (uws *UsersWebService) LoginEndpoint(c echo.Context) error {
  username := c.FormValue("username")
  password := c.FormValue("password")

  _, err := uws.usersService.SignIn(username, password)

  if err != nil {
    return lib.Render(
      c,
      auth.SignIn(auth.SignInVM{
        Error: err,
      }),
      400,
    )
  }

  return nil
}

func (uws *UsersWebService) SignUp(c echo.Context) error {
  return lib.Render(
    c,
    auth.Register(auth.RegisterVM{}),
    200,
  )
}

func (uws *UsersWebService) Home(c echo.Context) error {
  return lib.Render(
    c,
    users.Home(users.HomeVM{}),
    200,
  )
}



