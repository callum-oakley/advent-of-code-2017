package main

import (
	"bufio"
	"fmt"
	"os"
)

type component struct {
	back, front int
}

type bridge struct {
	length, weight int
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var components []component
	for scanner.Scan() {
		var c component
		fmt.Sscanf(scanner.Text(), "%d/%d", &c.back, &c.front)
		components = append(components, c)
	}
	fmt.Printf("part 1: %v\n", strongest(0, components).weight)
	fmt.Printf("part 2: %v\n", longest(0, components).weight)
}

func longest(front int, components []component) bridge {
	return bestBridge(func(a, b bridge) bridge {
		if a.length > b.length {
			return a
		} else if a.length == b.length && a.weight > b.weight {
			return a
		}
		return b
	}, front, components)
}

func strongest(front int, components []component) bridge {
	return bestBridge(func(a, b bridge) bridge {
		if a.weight > b.weight {
			return a
		}
		return b
	}, front, components)
}

func bestBridge(
	max func(bridge, bridge) bridge,
	front int,
	components []component,
) bridge {
	var best bridge
	for i, c := range components {
		if c.back == front {
			b := bestBridge(max, c.front, remove(components, i))
			best = max(bridge{b.length + 1, b.weight + c.back + c.front}, best)
		} else if c.front == front {
			b := bestBridge(max, c.back, remove(components, i))
			best = max(bridge{b.length + 1, b.weight + c.front + c.back}, best)
		}
	}
	return best
}

func remove(components []component, i int) []component {
	result := make([]component, len(components)-1)
	for j := range components {
		if j < i {
			result[j] = components[j]
		} else if j > i {
			result[j-1] = components[j]
		}
	}
	return result
}
