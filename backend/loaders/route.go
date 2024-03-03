package loaders

import (
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	//"github.com/gofiber/fiber/v2/middleware/filesystem"
	"github.com/wuhoops/silenda/backend/internal/handlers"
	"github.com/wuhoops/silenda/backend/internal/repositories"
	"github.com/wuhoops/silenda/backend/internal/services"
	//"io/fs"
	//"net/http"
	"time"
)

func SetupRoutes() {

	serverAddr := fmt.Sprintf("%s:%d", Env.ServerHost, Env.ServerPort)

	// Repositories
	var secretRepository = repositories.NewSecretRepository(DB)
	var workspaceRepository = repositories.NewWorkspaceRepository(DB)

	// Services
	var secretService = services.NewSecretService(secretRepository)
	var workspaceService = services.NewWorkspaceService(workspaceRepository)

	// Handlers
	var secretHandlers = handlers.NewSecretHandler(secretService)
	var workspaceHandlers = handlers.NewWorkspaceHandler(workspaceService)

	// Fiber App
	app := NewFiberApp()

	app.Use(cors.New(cors.Config{
		AllowOrigins:     Env.ServerOrigins,
		AllowCredentials: true,
	}))

	//f, err := fs.Sub(dist, "web/dist")
	//
	//if err != nil {
	//	panic(err)
	//}

	// Routes
	v1 := app.Group("/api/v1")

	v1.Post("secret.findSecretsByEnvMode", secretHandlers.FindSecretsByEnvMode)
	v1.Post("secret.createSecret", secretHandlers.CreateSecret)
	v1.Post("secret.deleteSecretById", secretHandlers.DeleteSecretById)

	v1.Post("workspace.createWorkspace", workspaceHandlers.CreateWorkspace)
	v1.Post("workspace.findAllWorkspaces", workspaceHandlers.FindAllWorkspaces)

	//app.Use("*", filesystem.New(filesystem.Config{
	//	Root:         http.FS(f),
	//	Index:        "index.html",
	//	Browse:       true,
	//	NotFoundFile: "index.html",
	//}))

	ListenAndServe(app, serverAddr)
}

func NewFiberApp() *fiber.App {
	fiberConfig := fiber.Config{
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			// Status code defaults to 500
			code := fiber.StatusInternalServerError
			Message := "Something went wrong"

			// Retrieve the custom status code if it's a *fiber.Error
			var e *fiber.Error
			if errors.As(err, &e) {
				code = e.Code
				Message = e.Message
			}

			if err != nil {
				// In case the SendFile fails
				return ctx.Status(code).JSON(handlers.ErrResponse{
					Code:      code,
					Message:   Message,
					Timestamp: time.Now(),
				})
			}

			// Return from handler
			return nil
		},
	}

	app := fiber.New(fiberConfig)
	return app
}

func ListenAndServe(app *fiber.App, serverAddr string) {
	err := app.Listen(serverAddr)
	if err != nil {
		panic(err)
	}
}
