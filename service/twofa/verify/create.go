package verify

import (
	"errors"

	"helloworld/model/twofa"
	svc_twofa "helloworld/service/twofa"

	"github.com/forbearing/golib/service"
	"github.com/forbearing/golib/types"
	"github.com/pquerna/otp/totp"
)

type Creator struct {
	service.Base[*twofa.Verify, *twofa.VerifyReq, *twofa.Verify]
}

func (v *Creator) Create(ctx *types.ServiceContext, req *twofa.VerifyReq) (rsp *twofa.Verify, err error) {
	if err = verify2FA(req); err != nil {
		return
	}
	return rsp, nil
}

func (v *Creator) CreateBefore(ctx *types.ServiceContext, verify *twofa.Verify) error {
	log := v.WithServiceContext(ctx, ctx.GetPhase())
	log.Info("verify create before")
	return nil
}

func (v *Creator) CreateAfter(ctx *types.ServiceContext, verify *twofa.Verify) error {
	log := v.WithServiceContext(ctx, ctx.GetPhase())
	log.Info("verify create after")
	return nil
}

func verify2FA(req *twofa.VerifyReq) error {
	secret, ok := svc_twofa.UserSecrets[req.User]
	if !ok {
		return errors.New("user not found")
	}
	valid := totp.Validate(req.Code, secret)
	if !valid {
		return errors.New("invalid code")
	}
	return nil
}
