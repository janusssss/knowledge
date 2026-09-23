package gomonkey

import (
	"errors"
	"github.com/agiledragon/gomonkey/v2"
	"github.com/labstack/echo/v4"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

type PlanService interface {
	GetPlan(planId string) (string, error)
}
type PlanServiceImpl struct{}

func (p *PlanServiceImpl) GetPlan(planId string) (string, error) {
	return planId, nil
}

func GetPlan(ctx echo.Context) error {
	planId := ctx.Param("plan_id")
	var planService PlanService
	planService = &PlanServiceImpl{}
	var s, err = planService.GetPlan(planId)
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, s)
}

func TestHelloConvey(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/hello", nil)
	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)
	p := &PlanServiceImpl{}

	Convey("TestGetPlans", t, func() {
		Convey("err plan list", func() {
			errA := errors.New("test err A")
			patchA := gomonkey.ApplyMethod(reflect.TypeOf(p), "GetPlan",
				func(_ *PlanServiceImpl, planId string) (string, error) {
					return "", errA
				})
			assert.Equal(t, GetPlan(ctx), errA)
			patchA.Reset()
		})
		Convey("norm", func() {
			patchA := gomonkey.ApplyMethod(reflect.TypeOf(p), "GetPlan",
				func(_ *PlanServiceImpl, planId string) (string, error) {
					return "", nil
				})
			assert.Equal(t, GetPlan(ctx), ctx.String(http.StatusOK, ""))
			patchA.Reset()
		})
	})
}
