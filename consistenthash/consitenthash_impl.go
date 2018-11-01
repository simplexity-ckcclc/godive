package consistenthash

import (
    "errors"
)

type ring struct {
    totalSlot int
    slot2node map[int]string
    node2slot map[string]int
}

func NewRing(slots int) *ring {
    return &ring {
        totalSlot: slots,
        slot2node: make(map[int]string, slots),
        node2slot: make(map[string]int),
    }
}

func GetNodeBySlot(ring *ring, hashSlot int) (string, error) {
    if hashSlot < 0 || hashSlot >= ring.totalSlot {
        return "", errors.New("slot is invalid")
    }

    min := ring.totalSlot
    next := ring.totalSlot
    for slot := range ring.slot2node {
        min = minInt(min, slot)
        if hashSlot < slot && slot < next {
            next = slot
        }
    }

    if next == ring.totalSlot {
        return ring.slot2node[min], nil
    }
    return ring.slot2node[next], nil
}

func AddNode(ring *ring, slot int, node string) {
    ring.slot2node[slot] = node
    ring.node2slot[node] = slot
}

func RemoveNode(ring *ring, node string) {
    slot, ok := ring.node2slot[node]
    if ok {
        delete(ring.node2slot, node)
        delete(ring.slot2node, slot)
    }
}

func minInt(a int, b int) int {
    if a < b {
        return a
    }
    return b
}