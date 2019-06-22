// +build !quick

package tests

import (
	"testing"

	pokeapi "github.com/mtslzr/pokeapi-go"
	"github.com/stretchr/testify/assert"
)

func TestResource(t *testing.T) {
	tests := []string{
		"berry", "berry-firmness", "berry-flavor", "contest-type",
		"contest-effect", "super-contest-effect", "encounter-method",
		"encounter-condition", "encounter-condition-value", "evolution-chain",
		"evolution-trigger", "generation", "pokedex", "version", "version-group",
		"item", "item-attribute", "item-category", "item-fling-effect",
		"item-pocket", "location", "location-area", "pal-park-area", "region",
		"machine", "move", "move-ailment", "move-battle-style", "move-category",
		"move-damage-class", "move-learn-method", "move-target", "ability",
		"egg-group", "gender", "growth-rate", "nature", "pokeathlon-stat",
		"pokemon", "pokemon-color", "pokemon-form", "pokemon-habitat",
		"pokemon-species", "stat", "type", "language",
	}
	for _, test := range tests {
		result, _ := pokeapi.Resource(test)
		assert.NotEqual(t, 0, len(result.Results),
			"Expect to receive more than zero results.")
		assert.NotEqual(t, nil, result.Results[0].URL,
			"Expect to receive results with URLs.")
	}
}
