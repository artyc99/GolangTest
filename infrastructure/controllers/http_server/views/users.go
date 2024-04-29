package views

import (
	"EchoAPI/core/modules"
	"EchoAPI/core/modules/users"
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"io"
	"net/http"
	"strconv"
)

func MountUsers(router *echo.Group, viewContext ViewContext) (err error) {
	router.GET("/users", func(eCtx echo.Context) error {
		var res users.GetListRes
		var req users.GetListReq

		rawReq := eCtx.QueryParams()
		if err != nil {
			return echo.NewHTTPError(http.StatusOK, err)
		}
		_, ok := rawReq["to_age"]
		if ok {
			toAge, err := strconv.Atoi(rawReq["to_age"][0])
			if err != nil {
				return echo.NewHTTPError(http.StatusOK, err)
			}
			req.ToAge = &toAge
		}

		_, ok = rawReq["from_age"]
		if ok {
			fromAge, err := strconv.Atoi(rawReq["from_age"][0])
			if err != nil {
				return echo.NewHTTPError(http.StatusOK, err)
			}
			req.FromAge = &fromAge
		}

		_, ok = rawReq["from_date"]
		if ok {
			fromDate, err := strconv.ParseInt(rawReq["from_date"][0], 10, 64)
			if err != nil {
				return echo.NewHTTPError(http.StatusOK, err)
			}
			req.FromDate = &fromDate
		}

		_, ok = rawReq["to_date"]
		if ok {
			toDate, err := strconv.ParseInt(rawReq["to_date"][0], 10, 64)
			if err != nil {
				return echo.NewHTTPError(http.StatusOK, err)
			}
			req.ToDate = &toDate
		}

		validate := validator.New()
		err = validate.Struct(req)
		if err != nil {
			return echo.NewHTTPError(http.StatusOK, err)
		}

		err = viewContext.Transaction(func(modules *modules.Modules) (err error) {
			res, err = modules.Users.GetList(eCtx.Request().Context(), req)
			if err != nil {
				return err
			}

			return nil
		})
		if err != nil {
			return err
		}

		return eCtx.JSON(http.StatusOK, res)
	})

	router.POST("/users", func(eCtx echo.Context) error {
		var res users.CreateRes
		var req users.CreateReq

		rawReq, err := io.ReadAll(eCtx.Request().Body)
		if err != nil {
			return nil
		}
		err = json.Unmarshal(rawReq, &req)
		if err != nil {
			return err
		}

		validate := validator.New()
		err = validate.Struct(req)
		if err != nil {
			return err
		}

		err = viewContext.Transaction(func(modules *modules.Modules) (err error) {
			res, err = modules.Users.Create(eCtx.Request().Context(), req)
			if err != nil {
				return err
			}

			return nil
		})
		if err != nil {
			return err
		}

		return eCtx.JSON(http.StatusOK, res)
	})

	return nil
}
