package functional

import (
	"net/http"
	"testing"

	"github.com/google/uuid"
	"github.com/steinfletcher/apitest"
	jsonpath "github.com/steinfletcher/apitest-jsonpath"
)

func TestProductCreation(t *testing.T) {
	t.Run("successfull", func(t *testing.T) {
		var res struct {
			ID uuid.UUID `json:"id"`
		}

		apitest.New().
			Handler(handler).
			Post("/public/api/v1/products").
			JSON(map[string]any{
				"description": "Milk",
				"quantity":    10,
			}).
			Expect(t).
			Assert(jsonpath.Present(`$.id`)).
			Status(http.StatusOK).
			End().JSON(&res)

		apitest.New().
			Handler(handler).
			Get("/public/api/v1/products/" + res.ID.String()).
			Expect(t).
			Assert(
				jsonpath.Chain().
					Equal("$.id", res.ID.String()).
					Equal("$.description", "Milk").
					Equal("$.quantity", float64(10)).
					End(),
			).
			Status(http.StatusOK).
			End()
	})
	t.Run("0 quantity", func(t *testing.T) {
		var res struct {
			ID uuid.UUID `json:"id"`
		}

		apitest.New().
			Handler(handler).
			Post("/public/api/v1/products").
			JSON(map[string]any{
				"description": "Milk",
				"quantity":    0,
			}).
			Expect(t).
			Assert(jsonpath.Present(`$.id`)).
			Status(http.StatusOK).
			End().JSON(&res)

		apitest.New().
			Handler(handler).
			Get("/public/api/v1/products/" + res.ID.String()).
			Expect(t).
			Assert(
				jsonpath.Chain().
					Equal("$.id", res.ID.String()).
					Equal("$.description", "Milk").
					Equal("$.quantity", float64(0)).
					End(),
			).
			Status(http.StatusOK).
			End()
	})
	t.Run("missed description", func(t *testing.T) {
		apitest.New().
			Handler(handler).
			Post("/public/api/v1/products").
			JSON(map[string]any{
				"quantity": 0,
			}).
			Expect(t).
			Status(http.StatusBadRequest).
			Assert(
				jsonpath.Chain().
					Equal("$.code", "MISSED_REQUIRED_FIELD").
					Equal("$.details", "description is required").
					End(),
			).
			End()
	})
}
