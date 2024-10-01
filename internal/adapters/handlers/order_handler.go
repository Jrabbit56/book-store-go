package handlers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/jrabbit56/book-store/internal/core/domain"
	"github.com/jrabbit56/book-store/internal/core/ports"
	//"github.com/jrabbit56/library-management/internal/core/services"
)

type OrderHandler struct {
	service ports.OrderService
}

func NewOrderHandler(service ports.OrderService) *OrderHandler {
	return &OrderHandler{service: service}
}

type SwaggerOrderFiberMap map[string]interface{}

// CreateOrder godoc
//
//	@Summary		Create a new order
//	@Description	Creates a new order with the provided information
//	@Tags			orders
//	@Accept			json
//	@Produce		json
//	@Param			order	body		domain.SwaggerOrder	true	"Order information"
//	@Success		201		{object}	handlers.SwaggerOrderFiberMap
//	@Failure		400		{object}	handlers.SwaggerOrderFiberMap
//	@Failure		500		{object}	handlers.SwaggerOrderFiberMap
//	@Router			/orders [post]
func (h *OrderHandler) CreateOrder(c *fiber.Ctx) error {

	order := new(domain.Order)
	if err := c.BodyParser(order); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid request"})
	}

	if err := h.service.CreateOrder(order); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	// return c.Status(fiber.StatusCreated).JSON(order)
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Order created successfully",
		"order":   order,
	})

}

// GetAllOrder godoc
//
//	@Summary		Get all orders
//	@Description	Retrieve a list of all orders
//	@Tags			orders
//	@Accept			json
//	@Produce		json
//	@Success		200	{array}		domain.SwaggerOrder
//	@Failure		500	{object}	handlers.SwaggerOrderFiberMap
//	@Router			/orders [get]
func (h *OrderHandler) GetAllOrder(c *fiber.Ctx) error {
	orders, err := h.service.GetAllOrder()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(orders)
}

// / GetOrder godoc
//
//	@Summary		Get a specific order
//	@Description	Retrieve a specific order by its ID, including its items
//	@Tags			orders
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int	true	"Order ID"
//	@Success		200	{array}		domain.SwaggerOrder
//	@Failure		400	{object}	handlers.SwaggerOrderFiberMap
//	@Failure		500	{object}	handlers.SwaggerOrderFiberMap
//	@Router			/orders/{id} [get]
func (h *OrderHandler) GetOrder(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid order ID"})
	}

	order, err := h.service.GetOrderWithItems(uint(id))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to retrieve order"})
	}

	return c.JSON(order)
}
