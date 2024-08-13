package tests

import (
    "testing"
    "os"
    "bytes"
    "encoding/json"
	"pro.cligolang/pkg"
	"io"
)

func TestAddItem(t *testing.T) {
    list := pkg.ShoppingList{}
    list.AddItem("Milk")

    if len(list.Items) != 1 {
        t.Errorf("Expected 1 item, got %d", len(list.Items))
    }

    if list.Items[0] != "Milk" {
        t.Errorf("Expected 'Milk', got '%s'", list.Items[0])
    }
}

func TestShowItems(t *testing.T) {
    list := pkg.ShoppingList{}
    list.AddItem("Milk")
    list.AddItem("Eggs")

    expected := "1. Milk\n2. Eggs\n"
    
    old := os.Stdout
    r, w, _ := os.Pipe()
    os.Stdout = w

    list.ShowItems()

    w.Close()
    os.Stdout = old

    var buf bytes.Buffer
    io.Copy(&buf, r)

    if buf.String() != expected {
        t.Errorf("Expected output:\n%s\nGot:\n%s", expected, buf.String())
    }
}

func TestMarkItem(t *testing.T) {
    list := pkg.ShoppingList{}
    list.AddItem("Milk")
    list.MarkItem(0)

    if !list.Marked[0] {
        t.Errorf("Expected item 0 to be marked")
    }
}

func TestDeleteItem(t *testing.T) {
    list := pkg.ShoppingList{}
    list.AddItem("Milk")
    list.AddItem("Eggs")
    list.DeleteItem(0)

    if len(list.Items) != 1 {
        t.Errorf("Expected 1 item, got %d", len(list.Items))
    }

    if list.Items[0] != "Eggs" {
        t.Errorf("Expected 'Eggs', got '%s'", list.Items[0])
    }
}

func TestReorderList(t *testing.T) {
    list := pkg.ShoppingList{}
    list.AddItem("Milk")
    list.AddItem("Eggs")
    list.AddItem("Bread")

    list.ReorderList(2, 0)

    if list.Items[0] != "Bread" || list.Items[1] != "Milk" {
        t.Errorf("Expected order 'Bread', 'Milk', 'Eggs', got '%s', '%s'", list.Items[0], list.Items[1])
    }
}

func TestSaveToFile(t *testing.T) {
    list := pkg.ShoppingList{}
    list.AddItem("Milk")
    list.AddItem("Eggs")

    pkg.SaveToFile(&list)

    defer os.Remove("shopping_list.json")

    file, err := os.Open("shopping_list.json")
    if err != nil {
        t.Fatalf("Expected no error opening file, got %v", err)
    }
    defer file.Close()

    var loadedList pkg.ShoppingList
    json.NewDecoder(file).Decode(&loadedList)

    if len(loadedList.Items) != 2 || loadedList.Items[0] != "Milk" || loadedList.Items[1] != "Eggs" {
        t.Errorf("Expected 'Milk', 'Eggs', got '%v'", loadedList.Items)
    }
}


func TestLoadFromFile(t *testing.T) {
    list := pkg.ShoppingList{}
    list.AddItem("Milk")
    list.AddItem("Eggs")
    pkg.SaveToFile(&list)
    defer os.Remove("shopping_list.json")

    var loadedList pkg.ShoppingList
    pkg.LoadFromFile(&loadedList)

    if len(loadedList.Items) != 2 || loadedList.Items[0] != "Milk" || loadedList.Items[1] != "Eggs" {
        t.Errorf("Expected 'Milk', 'Eggs', got '%v'", loadedList.Items)
    }
}
