package pokecache

import (
	"fmt"
	"time"
	"testing"
)

func TestAddGet(t *testing.T) {
	const interval = 5 * time.Second
	cases := []struct {
		key string
		val []byte
	}{
		{
			key: "https://pokeapi.co/api/v2/location-area",
			val: []byte("canalave-city-area\neterna-city-area\npastoria-city-area\nsunyshore-city-area\nsinnoh-pokemon-league-area\noreburgh-mine-1f\noreburgh-mine-b1f\nvalley-windworks-area\neterna-forest-area\nfuego-ironworks-area\nmt-coronet-1f-route-207\nmt-coronet-2f\nmt-coronet-3f\nmt-coronet-exterior-snowfall\nmt-coronet-exterior-blizzard\nmt-coronet-4f\nmt-coronet-4f-small-room\nmt-coronet-5f\nmt-coronet-6f\nmt-coronet-1f-from-exterior"),
		},
		{
			key: "https://pokeapi.co/api/v2/location-area?offset=20&limit=20",
			val: []byte("mt-coronet-1f-route-216\nmt-coronet-1f-route-211\nmt-coronet-b1f\ngreat-marsh-area-1\ngreat-marsh-area-2\ngreat-marsh-area-3\ngreat-marsh-area-4\ngreat-marsh-area-5\ngreat-marsh-area-6\nsolaceon-ruins-2f\nsolaceon-ruins-1f\nsolaceon-ruins-b1f-a\nsolaceon-ruins-b1f-b\nsolaceon-ruins-b1f-c\nsolaceon-ruins-b2f-a\nsolaceon-ruins-b2f-b\nsolaceon-ruins-b2f-c\nsolaceon-ruins-b3f-a\nsolaceon-ruins-b3f-b\nsolaceon-ruins-b3f-c"),
		},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("Test case %v", i), func(t *testing.T) {
			cache := NewCache(interval)
			cache.Add(c.key, c.val)
			val, ok := cache.Get(c.key)
			if !ok {
				t.Errorf("expected to find key")
				return
			}
			if string(val) != string(c.val) {
				t.Errorf("expected to find value")
				return
			}
		})
	}
}

func TestReapLoop(t *testing.T) {
	const baseTime = 5 * time.Millisecond
	const waitTime = baseTime + 5*time.Millisecond
	cache := NewCache(baseTime)
	cache.Add("https://pokeapi.co/api/v2/location-area", []byte("canalave-city-area\neterna-city-area\npastoria-city-area\nsunyshore-city-area\nsinnoh-pokemon-league-area\noreburgh-mine-1f\noreburgh-mine-b1f\nvalley-windworks-area\neterna-forest-area\nfuego-ironworks-area\nmt-coronet-1f-route-207\nmt-coronet-2f\nmt-coronet-3f\nmt-coronet-exterior-snowfall\nmt-coronet-exterior-blizzard\nmt-coronet-4f\nmt-coronet-4f-small-room\nmt-coronet-5f\nmt-coronet-6f\nmt-coronet-1f-from-exterior"))

	_, ok := cache.Get("https://pokeapi.co/api/v2/location-area")
	if !ok {
		t.Errorf("expected to find key")
		return
	}

	time.Sleep(waitTime)

	_, ok = cache.Get("https://pokeapi.co/api/v2/location-area")
	if ok {
		t.Errorf("expected to not find key")
		return
	}
}