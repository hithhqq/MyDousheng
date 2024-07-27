package verify

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/zeromicro/go-zero/rest/token"

	"MyDouSheng/app/identity/cmd/api/internal/svc"
	"MyDouSheng/app/identity/cmd/api/internal/types"
	"MyDouSheng/app/identity/cmd/rpc/identity"
	"MyDouSheng/common/ctxdata"
	"MyDouSheng/common/xerr"

	"github.com/golang-jwt/jwt/v4"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type VerifyTokenLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

var ValidateTokenError = xerr.NewErrCode(xerr.TOKEN_EXPIRE_ERROR)

// 验证请求token
func NewVerifyTokenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *VerifyTokenLogic {
	return &VerifyTokenLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *VerifyTokenLogic) VerifyToken(r *http.Request) (resp *types.VerifyTokenResp, err error) {
	// todo: add your logic here and delete this line
	fmt.Printf("VerifyToken\n")
	realRequestPath := r.Header.Get("X-Original-Uri")
	authorization := r.Header.Get("Authorization")
	if strings.Contains(realRequestPath, "?") {
		realRequestPath = strings.Split(realRequestPath, "?")[0]
	}
	var resultUserId int64
	if !l.urlNoAuth(realRequestPath) {
		userId, err := l.isPass(r)
		if err != nil {
			logx.WithContext(l.ctx).Errorf("authorization:%s, realRequestPath:%s", authorization, realRequestPath)
			return nil, err
		}
		if userId == 0 {
			return nil, errors.Wrapf(ValidateTokenError, "urlIsAuth.true isPass userId  is 0 , authorization: %s ,realRequestPath:%s", authorization, realRequestPath)
		}
		fmt.Printf("user is %v\n", userId)
		resultUserId = userId
	}
	return &types.VerifyTokenResp{
		StatusCode: resultUserId,
		StatusMsg:  "token有效",
	}, nil
}
func (l *VerifyTokenLogic) urlNoAuth(path string) bool {
	for _, val := range l.svcCtx.Config.NoAuthUrls {
		if val == path {
			return true
		}
	}
	return false
}

func (l *VerifyTokenLogic) isPass(r *http.Request) (int64, error) {
	parser := token.NewTokenParser()
	tok, err := parser.ParseToken(r, l.svcCtx.Config.JwtAuth.AccessSecret, "")

	if err != nil {
		return 0, errors.Wrapf(ValidateTokenError, "JwtAuthLogic isPass  ParseToken err : %v", err)
	}
	if tok.Valid {
		claims, ok := tok.Claims.(jwt.MapClaims) // 解析token中对内容
		if ok {
			userId, _ := claims[ctxdata.CtxKeyJwtUserId].(json.Number).Int64() // 获取userId 并且到后端redis校验是否过期
			if userId <= 0 {
				return 0, errors.Wrapf(ValidateTokenError, "JwtAuthLogic.isPass invalid userId  tokRaw:%s , tokValid :%v ,userId:%d ", tok.Raw, tok.Valid, userId)
			}
			resp, err := l.svcCtx.IdentityRpc.ValidateToken(l.ctx, &identity.ValidateTokenReq{
				UserId: userId,
				Token:  tok.Raw,
			})
			if err != nil || !resp.Ok {
				return 0, errors.Wrapf(ValidateTokenError, "JwtAuthLogic.isPass IdentityRpc . ValidateToken err:%v ,resp:%+v , tokRaw:%s , tokValid : %v,userId:%d ", err, resp, tok.Raw, tok.Valid, userId)
			}
			return userId, nil
		} else {
			return 0, errors.Wrapf(ValidateTokenError, "tok.Claims is not ok ,tok.Claims ：%+v , claims : %+v , ok:%v ", tok.Claims, claims, ok)
		}
	}
	return 0, nil
}
