package setup

import (
	"errors"
	"net/http"

	"helloworld/model/twofa"
	svc_twofa "helloworld/service/twofa"

	"github.com/forbearing/golib/service"
	"github.com/forbearing/golib/types"
	"github.com/pquerna/otp/totp"
	"github.com/skip2/go-qrcode"
)

type Lister struct {
	service.Base[*twofa.Setup, *twofa.Setup, *twofa.SetupRsp]
}

func (s *Lister) List(ctx *types.ServiceContext, req *twofa.Setup) (rsp *twofa.SetupRsp, err error) {
	if err = generate2FA(ctx); err != nil {
		return
	}
	return
}

func (s *Lister) ListBefore(ctx *types.ServiceContext, setups *[]*twofa.Setup) error {
	log := s.WithServiceContext(ctx, ctx.GetPhase())
	log.Info("setup list before")
	return nil
}

func (s *Lister) ListAfter(ctx *types.ServiceContext, setups *[]*twofa.Setup) error {
	log := s.WithServiceContext(ctx, ctx.GetPhase())
	log.Info("setup list after")
	return nil
}

func generate2FA(ctx *types.ServiceContext) error {
	c := ctx.GinContext
	username := c.Query("user")
	if username == "" {
		return errors.New("missing user")
	}
	key, err := totp.Generate(totp.GenerateOpts{Issuer: "MyGinApp", AccountName: username})
	if err != nil {
		return err
	}
	svc_twofa.UserSecrets[username] = key.Secret()
	pngData, _ := qrcode.Encode(key.URL(), qrcode.Medium, 256)
	c.Data(http.StatusOK, "image/png", pngData)
	return nil
}
