package pkg

import (
    "fmt"
)

type ShoppingList struct {
    Items  []string
    Marked []bool
}

func (s *ShoppingList) AddItem(item string) {
    s.Items = append(s.Items, item)
    s.Marked = append(s.Marked, false)
}

func (s *ShoppingList) ShowItems() {
    for i, item := range s.Items {
        mark := ""
        if s.Marked[i] {
            mark = " (marked)"
        }
        fmt.Printf("%d. %s%s\n", i+1, item, mark)
    }
}

func (s *ShoppingList) MarkItem(index int) {
    if index >= 0 && index < len(s.Items) {
        s.Marked[index] = true
    }
}

func (s *ShoppingList) DeleteItem(index int) {
    if index >= 0 && index < len(s.Items) {
        s.Items = append(s.Items[:index], s.Items[index+1:]...)
        s.Marked = append(s.Marked[:index], s.Marked[index+1:]...)
    }
}

func (s *ShoppingList) ReorderList(fromIndex, toIndex int) {
    if fromIndex < len(s.Items) && toIndex < len(s.Items) && fromIndex != toIndex {
        item := s.Items[fromIndex]
        s.Items = append(s.Items[:fromIndex], s.Items[fromIndex+1:]...)
        s.Items = append(s.Items[:toIndex], append([]string{item}, s.Items[toIndex:]...)...)
    }
}




func (s *ShoppingList) AddItem(item string) {

}

func (s *ShoppingList) ShowItems() {

}

func (s *ShoppingList) MarkItem(index int) {

}

func (s *ShoppingList) RemoveItem(index int) {

}
