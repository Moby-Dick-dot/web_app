package routes

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"reflect"
	"strings"
	g "web_app/global"
	"web_app/request"
	"web_app/response"
)

// ApiParamsErrCallback 参数绑定失败后的回调函数
func ApiParamsErrCallback(ctx *gin.Context, err error) {
	response.ResponseErrorWithMsg(ctx, g.CodeInvalidParams, err.Error())
}

type ParamsErrCallback func(ctx *gin.Context, err error)

func GenHandlerFunc(controllerFunc interface{}, paramsErrCallback ParamsErrCallback) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		controllerFuncType := reflect.TypeOf(controllerFunc)
		if controllerFuncType.Kind() != reflect.Func {
			panic(ctx.Request.URL.Path + " controllerFunc not a func")
		}

		args := make([]reflect.Value, 0)
		args = append(args, reflect.ValueOf(ctx))
		for i := 0; i < controllerFuncType.NumIn(); i++ {
			inParam := controllerFuncType.In(i)
			if !inParam.Implements(reflect.TypeOf((*request.Validator)(nil)).Elem()) {
				continue
			}
			inParamVale := reflect.New(inParam.Elem())
			req, ok := inParamVale.Interface().(request.Validator)
			if !ok {
				panic(ctx.Request.URL.Path + " req  convert requestutil.ValidateAble err")
			}
			err := bindRequestParams(ctx, req, false)
			if err != nil {
				if paramsErrCallback != nil {
					paramsErrCallback(ctx, err)
				}
				return
			}
			args = append(args, reflect.ValueOf(req))
		}
		reflect.ValueOf(controllerFunc).Call(args)
	}
}

// 解析请求参数，仅支持 get post delete put
func bindRequestParams(context *gin.Context, req request.Validator, aes bool) error {
	if req == nil {
		return errors.New("req is nil")
	}
	jsonAllow := context.Request.Method == http.MethodPost || context.Request.Method == http.MethodPut || context.Request.Method == http.MethodDelete
	if jsonAllow && strings.Contains(context.GetHeader("Content-Type"), "application/json") {
		data, err := ioutil.ReadAll(context.Request.Body)
		if err != nil {
			return err
		}
		defer context.Request.Body.Close()
		err = json.Unmarshal(data, req)
		if err != nil {
			return err
		}
		context.Request.Body = ioutil.NopCloser(bytes.NewBuffer(data))
	} else {
		err := context.ShouldBindQuery(req)
		if err != nil {
			return err
		}
	}
	err := req.Validate()
	if err != nil {
		return err
	}

	return nil
}
