package main

import (
	"io"
	"live-chat/config"
	"live-chat/users"
	"text/template"

	"github.com/labstack/echo/v4"
	"gopkg.in/olahol/melody.v1"
)

func main() {
	r := echo.New()
	m := melody.New()

	// Initializes database
	db := config.ConnectPGLocal()

	userRepository := users.NewRepository(db)
	userService := users.NewService(userRepository)
	userHandler := users.NewUserController(userService)

	r.Renderer = NewRenderer("views/*.html", true)
	r.Any("/login", userHandler.LoginRegisterPage)
	r.Any("/home", userHandler.HomePage)
	r.Any("/chat", userHandler.ChatPage)
	r.Any("/register", userHandler.LoginRegisterPage)
	r.POST("/process/register", userHandler.Register)
	r.POST("/process/login", userHandler.Login)
	r.POST("/process/logout", userHandler.Logout)

	r.Static("/chat-page", "views/live-chat")
	r.GET("/ws", func(c echo.Context) error {
		return m.HandleRequest(c.Response().Writer, c.Request())
	})
	m.HandleMessage(func(s *melody.Session, msg []byte) {
		m.Broadcast(msg)
	})

	r.Start(":8080")
}

type Renderer struct {
	template *template.Template
	debug    bool
	location string
}

func NewRenderer(location string, debug bool) *Renderer {
	tpl := new(Renderer)
	tpl.location = location
	tpl.debug = debug

	tpl.ReloadTemplates()

	return tpl
}

func (t *Renderer) ReloadTemplates() {
	t.template = template.Must(template.ParseGlob(t.location))
}

func (t *Renderer) Render(
	w io.Writer,
	name string,
	data interface{},
	c echo.Context,
) error {
	if t.debug {
		t.ReloadTemplates()
	}

	return t.template.ExecuteTemplate(w, name, data)
}
