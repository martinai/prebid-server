package martinai

import (
	"text/template"

	"github.com/prebid/prebid-server/adapters"
	"github.com/prebid/prebid-server/usersync"
)

func NewMartinaiSyncer(temp *template.Template) usersync.Usersyncer {
	return adapters.NewSyncer("martinia", 69, temp, adapters.SyncTypeRedirect)
}
