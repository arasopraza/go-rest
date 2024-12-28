package user

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type handlerProduct struct {
	usecase    Usecase
	repository Repository
}

func UserHandler(userRoute *echo.Group, usecase Usecase, repository Repository) {
	handler := &handlerProduct{
		usecase:    usecase,
		repository: repository,
	}

	userRoute.POST("/user", handler.CreateUser)
	// productRoute.GET("/product", handler.GetProducts)
	// productRoute.PUT("/product/:id", handler.UpdateProduct)
	// productRoute.DELETE("/product/:id", handler.DeleteProduct)
}

type User struct {
	ID    string `json:"id"`
	Name  string `json:"name" validate:"required"`
	Email string `json:"email" validate:"required,email"`
}

func (h *handlerProduct) CreateUser(c echo.Context) error {

	var req User

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
	}

	if err := c.Validate(req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
	}

	id, createdAt, err := h.usecase.CreateUser(c.Request().Context(), req)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{"id": id, "createdAt": createdAt})
}
