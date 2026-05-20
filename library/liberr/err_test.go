package liberr

import (
	"context"
	"errors"
	"testing"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
)

func TestErrIsNil(t *testing.T) {
	ctx := context.Background()


	t.Run("nil error", func(t *testing.T) {
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("ErrIsNil should not panic with nil error, but got: %v", r)
			}
		}()
		ErrIsNil(ctx, nil)
	})


	t.Run("non-nil error without message", func(t *testing.T) {
		testErr := errors.New("test error")
		defer func() {
			if r := recover(); r == nil {
				t.Error("ErrIsNil should panic with non-nil error")
			} else if r != testErr {
				t.Errorf("Expected panic with error %v, but got: %v", testErr, r)
			}
		}()
		ErrIsNil(ctx, testErr)
	})


	t.Run("non-nil error with custom message", func(t *testing.T) {
		testErr := errors.New("test error")
		customMsg := "custom error message"
		defer func() {
			if r := recover(); r == nil {
				t.Error("ErrIsNil should panic with non-nil error and custom message")
			} else if r != customMsg {
				t.Errorf("Expected panic with message %q, but got: %v", customMsg, r)
			}
		}()
		ErrIsNil(ctx, testErr, customMsg)
	})


	t.Run("non-nil error with multiple messages", func(t *testing.T) {
		testErr := errors.New("test error")
		firstMsg := "first message"
		secondMsg := "second message"
		defer func() {
			if r := recover(); r == nil {
				t.Error("ErrIsNil should panic with non-nil error and messages")
			} else if r != firstMsg {
				t.Errorf("Expected panic with first message %q, but got: %v", firstMsg, r)
			}
		}()
		ErrIsNil(ctx, testErr, firstMsg, secondMsg)
	})


	t.Run("gerror with CodeValidationFailed", func(t *testing.T) {
		testErr := gerror.NewCode(gcode.CodeValidationFailed, "验证码已过期或不存在")
		defer func() {
			if r := recover(); r == nil {
				t.Error("ErrIsNil should panic with gerror")
			} else if r != testErr {
				t.Errorf("Expected panic with gerror %v, but got: %v", testErr, r)
			}
		}()
		ErrIsNil(ctx, testErr)
	})


	t.Run("gerror with custom message", func(t *testing.T) {
		testErr := gerror.NewCode(gcode.CodeValidationFailed, "验证码已过期或不存在")
		customMsg := "验证失败"
		defer func() {
			if r := recover(); r == nil {
				t.Error("ErrIsNil should panic with gerror and custom message")
			} else if r != customMsg {
				t.Errorf("Expected panic with message %q, but got: %v", customMsg, r)
			}
		}()
		ErrIsNil(ctx, testErr, customMsg)
	})


	t.Run("gerror with CodeInternalError", func(t *testing.T) {
		testErr := gerror.NewCode(gcode.CodeInternalError, "内部服务错误")
		defer func() {
			if r := recover(); r == nil {
				t.Error("ErrIsNil should panic with CodeInternalError")
			} else if r != testErr {
				t.Errorf("Expected panic with gerror %v, but got: %v", testErr, r)
			}
		}()
		ErrIsNil(ctx, testErr)
	})


	t.Run("multiple validation scenarios", func(t *testing.T) {
		scenarios := []struct {
			name string
			err  error
			msg  string
		}{
			{
				name: "验证码过期",
				err:  gerror.NewCode(gcode.CodeValidationFailed, "验证码已过期或不存在"),
				msg:  "验证码验证失败",
			},
			{
				name: "参数验证失败",
				err:  gerror.NewCode(gcode.CodeValidationFailed, "参数格式不正确"),
				msg:  "请求参数错误",
			},
			{
				name: "业务验证失败",
				err:  gerror.NewCode(gcode.CodeValidationFailed, "用户状态异常"),
				msg:  "",
			},
		}

		for _, scenario := range scenarios {
			t.Run(scenario.name, func(t *testing.T) {
				defer func() {
					if r := recover(); r == nil {
						t.Error("ErrIsNil should panic")
					} else {
						if scenario.msg != "" {
							if r != scenario.msg {
								t.Errorf("Expected panic with message %q, but got: %v", scenario.msg, r)
							}
						} else {
							if r != scenario.err {
								t.Errorf("Expected panic with error %v, but got: %v", scenario.err, r)
							}
						}
					}
				}()

				if scenario.msg != "" {
					ErrIsNil(ctx, scenario.err, scenario.msg)
				} else {
					ErrIsNil(ctx, scenario.err)
				}
			})
		}
	})
}
