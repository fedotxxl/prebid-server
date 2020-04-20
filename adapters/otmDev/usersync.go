package otmDev

import (
	"net/url"

	"github.com/prebid/prebid-server/adapters"
	"github.com/prebid/prebid-server/config"
	"github.com/prebid/prebid-server/usersync"
)

func NewOtmSyncer(cfg *config.Configuration) usersync.Usersyncer {
	redirectURI := url.QueryEscape(cfg.ExternalURL) + "%2Fsetuid%3Fbidder%3DotmDev%26uid%3D%24%7BUSER_ID%7D"
	usersyncURL := "//pix.ssp.otm-r.com/match?callback_url="
	return adapters.NewSyncer("otmDev", 48, adapters.ResolveMacros(usersyncURL+redirectURI), adapters.SyncTypeRedirect)
}
