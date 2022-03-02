package main

import "charmtest/internal/types"

var entries = []*types.MenuEntry{
	{Name: "First", Value: "This is the first menu entry"},
	{Name: "Second", Value: "This is the second menu entry. Its comes after the first"},
	{Name: "Third", Value: "The third menu entry comes between the second and fourth."},
	{Name: "Fourth", Value: "The fourth is the second to last menu entry."},
	{Name: "Fifth", Value: "This is the last, there is no more."},
}
