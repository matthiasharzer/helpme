package konsole

import "github.com/matthiasharzer/helpme/knowledge/resource"

var Resource = resource.NestedResource{
	Items: []resource.Resource{
		resource.DescriptiveResource{
			Title: "Konsole – KDE terminal emulator",
			Items: []resource.Item{
				resource.ActionPairItem{
					Action:      "Ctrl + Shift + ,",
					Description: "Show Konsole settings",
				},
				resource.ActionPairItem{
					Action:      "Ctrl + Shift + m",
					Description: "Toggle toolbar",
				},
			},
		},
	},
}
