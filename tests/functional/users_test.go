package functional

import (
	"net/http"
	"testing"

	"github.com/google/uuid"
	"github.com/steinfletcher/apitest"
	jsonpath "github.com/steinfletcher/apitest-jsonpath"
)

func TestUserRegistration(t *testing.T) {
	t.Run("successfull", func(t *testing.T) {
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

		apitest.New().
			Handler(handler).
			Get("/public/api/v1/users/" + res.ID.String()).
			Expect(t).
			Assert(
				jsonpath.Chain().
					Equal("$.id", res.ID.String()).
					Equal("$.firstname", "test").
					End(),
			).
			Status(http.StatusOK).
			End()
	})
	t.Run("unacceptable age", func(t *testing.T) {
		apitest.New().
			Handler(handler).
			Post("/public/api/v1/users").
			JSON(map[string]any{
				"firstname":  "test",
				"lastname":   "test",
				"full_name":  "",
				"age":        12,
				"is_married": true,
				"password":   "qwerty1212",
			}).
			Expect(t).
			Status(http.StatusBadRequest).
			Assert(
				jsonpath.Chain().
					Equal("$.code", "INSUFFICIENT_CONDITIONS").
					Equal("$.details", "insufficient age").
					End(),
			).
			End()
	})
	t.Run("short password", func(t *testing.T) {
		apitest.New().
			Handler(handler).
			Post("/public/api/v1/users").
			JSON(map[string]any{
				"firstname":  "test",
				"lastname":   "test",
				"full_name":  "",
				"age":        18,
				"is_married": true,
				"password":   "qwerty",
			}).
			Expect(t).
			Status(http.StatusBadRequest).
			Assert(
				jsonpath.Chain().
					Equal("$.code", "INSUFFICIENT_CONDITIONS").
					Equal("$.details", "password too short").
					End(),
			).
			End()
	})
}
