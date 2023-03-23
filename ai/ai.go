package ai

import (
	"github.com/Lama06/Dame-Go.git/dame"
)

func FindBestMove(brett dame.Brett, spieler dame.Spieler, maxDepth int) dame.Brett {
	type Node struct {
		depth          int
		brett          dame.Brett
		amZug          dame.Spieler
		bewertung      int
		bestChildIndex int
		children       []int
	}

	nodes := []Node{
		{
			depth:          0,
			brett:          brett,
			amZug:          spieler,
			bewertung:      0,
			bestChildIndex: -1,
			children:       nil,
		},
	}

	// Mögliche Züge generieren
	for depth := 1; depth <= maxDepth; depth++ {
		for parentIndex, parent := range nodes {
			if parent.depth != depth-1 {
				continue
			}

			for possibleMove := range parent.brett.PossibleMoves(parent.amZug) {
				child := Node{
					depth:          depth,
					amZug:          !parent.amZug,
					bewertung:      0,
					brett:          possibleMove,
					bestChildIndex: -1,
					children:       nil,
				}
				nodes[parentIndex].children = append(nodes[parentIndex].children, len(nodes))
				nodes = append(nodes, child)
			}
		}
	}

	// Bewertungen der untersten Zeile berechnen
	for nodeIndex, node := range nodes {
		if node.depth != maxDepth {
			continue
		}

		nodes[nodeIndex].bewertung = statistikFromBrett(node.brett).bewertung(spieler)
	}

	// Bewertungen in den Zeilen darüber berechnen
	for depth := maxDepth - 1; depth >= 0; depth-- {
		for nodeIndex, node := range nodes {
			if node.depth != depth {
				continue
			}

			var (
				bestChildIndex     int
				bestChildBewertung int
			)

			for childIndexIndex, childIndex := range node.children {
				child := nodes[childIndex]

				if childIndexIndex == 0 {
					bestChildIndex = childIndex
					bestChildBewertung = child.bewertung
					continue
				}

				if node.amZug == spieler && child.bewertung > bestChildBewertung {
					bestChildIndex = childIndex
					bestChildBewertung = child.bewertung
				} else if node.amZug != spieler && child.bewertung < bestChildBewertung {
					bestChildIndex = childIndex
					bestChildBewertung = child.bewertung
				}
			}

			nodes[nodeIndex].bewertung = bestChildBewertung
			nodes[nodeIndex].bestChildIndex = bestChildIndex
		}
	}

	return nodes[nodes[0].bestChildIndex].brett
}
