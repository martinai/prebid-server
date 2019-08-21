package martinai

import (
	"text/template"

	"github.com/prebid/prebid-server/adapters"
	"github.com/prebid/prebid-server/usersync"
)

func NewMartinaiSyncer(temp *template.Template) usersync.Usersyncer {
	// Use 0 until we have our own vendor id
	return adapters.NewSyncer("martinai", 0, temp, adapters.SyncTypeRedirect)
}
