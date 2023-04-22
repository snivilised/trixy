package i18n

// TODO: Should be updated to use url of the implementing project,
// so should not be left as arcadia.
const trixySourceID = "github.com/snivilised/trixy"

type trixyTemplData struct{}

func (td trixyTemplData) SourceID() string {
	return trixySourceID
}
