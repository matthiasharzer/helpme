package knowledge

import (
	"github.com/matthiasharzer/helpme/knowledge/builtin/konsole"
	"github.com/matthiasharzer/helpme/knowledge/builtin/tmux"
	"github.com/matthiasharzer/helpme/knowledge/resource"
)

func LoadResources() map[string]resource.Resource {
	return map[string]resource.Resource{
		"tmux":    tmux.Resource,
		"konsole": konsole.Resource,
	}
}
