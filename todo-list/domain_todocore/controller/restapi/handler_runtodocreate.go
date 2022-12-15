package restapi

import (
	"context"
	"net/http"
	"time"
	"todo-list/domain_todocore/model/entity"
	"todo-list/domain_todocore/usecase/runtodocreate"
	"todo-list/shared/gogen"
	"todo-list/shared/infrastructure/logger"
	"todo-list/shared/model/payload"
	"todo-list/shared/util"

	"github.com/gin-gonic/gin"
)

func (r *ginController) runTodoCreateHandler() gin.HandlerFunc {

	type InportRequest = runtodocreate.InportRequest
	type InportResponse = runtodocreate.InportResponse

	inport := gogen.GetInport[InportRequest, InportResponse](r.GetUsecase(InportRequest{}))

	type request struct {
		InportRequest
	}

	type response struct {
		Todo *entity.Todo `json:"todo"`
	}

	return func(c *gin.Context) {

		traceID := util.GenerateID(16)

		ctx := logger.SetTraceID(context.Background(), traceID)

		var jsonReq request
		if err := c.BindJSON(&jsonReq); err != nil {
			r.log.Error(ctx, err.Error())
			c.JSON(http.StatusBadRequest, payload.NewErrorResponse(err, traceID))
			return
		}

		var req InportRequest
		req.Message = jsonReq.Message
		req.Now = time.Now()
		req.RandomString = util.GenerateID(5)

		r.log.Info(ctx, util.MustJSON(req))

		res, err := inport.Execute(ctx, req)
		if err != nil {
			r.log.Error(ctx, err.Error())
			c.JSON(http.StatusBadRequest, payload.NewErrorResponse(err, traceID))
			return
		}

		var jsonRes response
		jsonRes.Todo = res.Todo

		r.log.Info(ctx, util.MustJSON(jsonRes))
		c.JSON(http.StatusOK, payload.NewSuccessResponse(jsonRes, traceID))

	}
}
