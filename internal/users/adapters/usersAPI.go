package adapters

import (
	"log"
	"os"

	auth "github.com/edlingao/go-auth/auth/core"
	"github.com/edlingao/hexago/internal/users/core"
	"github.com/edlingao/hexago/internal/users/ports/driven"
	"github.com/edlingao/hexago/internal/users/ports/driving"
	"github.com/labstack/echo/v4"
)

type UsersAPIService struct {
	dbService      driven.StoringUsers[core.User]
	httpService    *echo.Group
	sessionService auth.SessionService
	usersService   core.UserService
  secret         string
}

type SignInResponse struct {
  User core.User `json:"user"`
  Token string `json:"token"`
}

func NewUsersAPIService(
	dbService driven.StoringUsers[core.User],
	httpService *echo.Group,
	sessionService auth.SessionService,
	usersService core.UserService,
) *UsersAPIService {

  secret := os.Getenv("JWT_SECRET")

	uApiService := &UsersAPIService{
		dbService:      dbService,
		httpService:    httpService,
		sessionService: sessionService,
		usersService:   usersService,
    secret:         secret,
	}

  uApiService.httpService.POST("/signin", uApiService.SignIn)
  uApiService.httpService.POST("/signup", uApiService.SignUp)

  // Protected routes
  protected := uApiService.httpService.Group("", sessionService.APIAuth)
  protected.GET("/all", uApiService.GetAllUsers)

	return uApiService
}

func (uas *UsersAPIService) GetAllUsers(c echo.Context) error {
  users := uas.dbService.GetAll("users")

  return c.JSON(200, driving.Response[[]core.User]{
    Status: 200,
    Message: "Success",
    Data: users,
  })
}

func (uas *UsersAPIService) SignIn(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")

	if username == "" || password == "" {
		return c.JSON(400, driving.Response[interface{}]{
      Status: 400,
      Message: "Username and password are required",
    })
	}

	user, err := uas.usersService.SignIn(username, password)

	if err != nil {
		return c.JSON(500, driving.Response[interface{}]{
      Status: 500,
      Message: err.Error(),
    })
	}

  token, err := uas.sessionService.Create(user.ID, user.Username, uas.secret)

  if err != nil {
    log.Println(user.ID, user.Username, uas.secret, err)
    return c.JSON(500, driving.Response[interface{}]{
      Status: 500,
      Message: err.Error(),
    })
  }

	return c.JSON(200, driving.Response[SignInResponse]{
    Status: 200,
    Message: "User signed in",
    Data: SignInResponse{
      User: user,
      Token: token.Token,
    },
  })
}

func (uas *UsersAPIService) SignUp(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")
	err := uas.usersService.Register(username, password)

	if err != nil {
		return c.JSON(500, driving.Response[interface{}]{
      Status: 500,
      Message: err.Error(),
    })
	}

  user, err := uas.usersService.GetByUsername(username)
  
  if err != nil {
    return c.JSON(500, driving.Response[interface{}]{
      Status: 500,
      Message: err.Error(),
    })
  }

  token, err := uas.sessionService.Create(user.ID, user.Username, uas.secret)

  if err != nil {
    return c.JSON(500, driving.Response[interface{}]{
      Status: 500,
      Message: err.Error(),
    })
  }

	return c.JSON(200, driving.Response[SignInResponse]{
    Status: 200,
    Message: "User registered",
    Data: SignInResponse {
      User: user,
      Token: token.Token,
    },
  })
}
