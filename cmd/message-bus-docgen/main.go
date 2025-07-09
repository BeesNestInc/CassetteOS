package main

import (
	"github.com/BeesNestInc/CassetteOS-Common/external"
	"github.com/BeesNestInc/CassetteOS/codegen/message_bus"
	"github.com/BeesNestInc/CassetteOS/common"
	"github.com/samber/lo"
)

func main() {
	eventTypes := lo.Map(common.EventTypes, func(item message_bus.EventType, index int) external.EventType {
		return external.EventType{
			Name:     item.Name,
			SourceID: item.SourceID,
			PropertyTypeList: lo.Map(
				item.PropertyTypeList, func(item message_bus.PropertyType, index int) external.PropertyType {
					return external.PropertyType{
						Name:        item.Name,
						Description: item.Description,
						Example:     item.Example,
					}
				},
			),
		}
	})

	external.PrintEventTypesAsMarkdown(common.SERVICENAME, common.VERSION, eventTypes)
}
