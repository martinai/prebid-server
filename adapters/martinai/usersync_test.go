package martinai

import (
	"testing"
	"text/template"

	"github.com/stretchr/testify/assert"
)

func TestMartinaiSyncer(t *testing.T) {
	temp := template.Must(template.New("sync-template").Parse("https://bidder.martin.ai/sync?redirect_url=localhost%2Fsetuid%3Fbidder%3Dmartinai%26gdpr%3D{{.GDPR}}%26gdpr_consent%3D{{.GDPRConsent}}%26uid%3D%24%7BMARTIN_AI_ID%7D"))
	syncer := NewMartinaiSyncer(temp)
	syncInfo, err := syncer.GetUsersyncInfo("", "")
	assert.NoError(t, err)
	assert.Equal(t, "https://bidder.martin.ai/sync?redirect_url=localhost%2Fsetuid%3Fbidder%3Dmartinai%26gdpr%3D%26gdpr_consent%3D%26uid%3D%24%7BMARTIN_AI_ID%7D", syncInfo.URL)
	assert.Equal(t, "redirect", syncInfo.Type)
	assert.EqualValues(t, 0, syncer.GDPRVendorID())
	assert.Equal(t, false, syncInfo.SupportCORS)
}
