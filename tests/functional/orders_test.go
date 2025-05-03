package functional

import (
	"net/http"
	"testing"

	"github.com/google/uuid"
	"github.com/steinfletcher/apitest"
	jsonpath "github.com/steinfletcher/apitest-jsonpath"
)

func TestOrderCreation(t *testing.T) {
	t.Run("successfull", func(t *testing.T) {
		// Register user
		var res struct {
			ID uuid.UUID `json:"id"`
		}

		apitest.New().
			Handler(handler).
			Post("/public/api/v1/users").
			JSON(map[string]any{
				"firstname":  "test",
				"lastname":   "test",
				"full_name":  "",
				"age":        18,
				"is_married": true,
				"password":   "qwerty1212",
			}).
			Expect(t).
			Assert(jsonpath.Present(`$.id`)).
			Status(http.StatusOK).
			End().JSON(&res)

		userID := res.ID

		// Create product
		apitest.New().
			Handler(handler).
			Post("/public/api/v1/products").
			JSON(map[string]any{
				"description": "Milk",
				"quantity":    12,
			}).
			Expect(t).
			Assert(jsonpath.Present(`$.id`)).
			Status(http.StatusOK).
			End().JSON(&res)

		productID := res.ID

		// Create order
		apitest.New().
			Handler(handler).
			Post("/public/api/v1/orders").
			JSON(map[string]any{
				"user_id": userID,
				"items": []map[string]any{
					map[string]any{
						"product_id":  productID,
						"description": "Milk",
						"price":       500,
						"quantity":    1,
					},
				},
			}).
			Expect(t).
			Assert(jsonpath.Present(`$.id`)).
			Status(http.StatusOK).
			End().JSON(&res)
	})
	t.Run("not enough products", func(t *testing.T) {
		// Register user
		var res struct {
			ID uuid.UUID `json:"id"`
		}

		apitest.New().
			Handler(handler).
			Post("/public/api/v1/users").
			JSON(map[string]any{
				"firstname":  "test",
				"lastname":   "test",
				"full_name":  "",
				"age":        18,
				"is_married": true,
				"password":   "qwerty1212",
			}).
			Expect(t).
			Assert(jsonpath.Present(`$.id`)).
			Status(http.StatusOK).
			End().JSON(&res)

		userID := res.ID

		// Create product
		apitest.New().
			Handler(handler).
			Post("/public/api/v1/products").
			JSON(map[string]any{
				"description": "Milk",
				"quantity":    5,
			}).
			Expect(t).
			Assert(jsonpath.Present(`$.id`)).
			Status(http.StatusOK).
			End().JSON(&res)

		productID := res.ID

		// Create order
		apitest.New().
			Handler(handler).
			Post("/public/api/v1/orders").
			JSON(map[string]any{
				"user_id": userID,
				"items": []map[string]any{
					map[string]any{
						"product_id":  productID,
						"description": "Milk",
						"price":       500,
						"quantity":    7,
					},
				},
			}).
			Expect(t).
			Status(http.StatusBadRequest).
			Assert(
				jsonpath.Chain().
					Equal("$.code", "INSUFFICIENT_RESOURCES").
					Equal("$.details", "not enough products").
					End(),
			).
			End()
	})
}
