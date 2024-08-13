package main

import (
    "fmt"
    "pro.cligolang/pkg"
)

func main() {
    shoppingList := pkg.ShoppingList{}

    pkg.LoadFromFile(&shoppingList)

    for {
        fmt.Println("\nChoose an option:")
        fmt.Println("1. Add item")
        fmt.Println("2. Show items")
        fmt.Println("3. Mark item")
        fmt.Println("4. Delete item")
        fmt.Println("5. Exit and Save")

        var choice int
        fmt.Print("Enter your choice: ")
        fmt.Scanln(&choice)

        switch choice {
        case 1:
            fmt.Print("Enter item to add: ")
            var item string
            fmt.Scanln(&item)
            shoppingList.AddItem(item)
            fmt.Println("Item added successfully!")

        case 2:
            fmt.Println("Current Shopping List:")
            shoppingList.ShowItems()

        case 3:
            fmt.Print("Enter the item number to mark: ")
            var index int
            fmt.Scanln(&index)
            if index > 0 && index <= len(shoppingList.Items) {
                shoppingList.MarkItem(index - 1)
                fmt.Println("Item marked successfully!")
            } else {
                fmt.Println("Invalid item number.")
            }

        case 4:
            fmt.Print("Enter the item number to delete: ")
            var index int
            fmt.Scanln(&index)
            if index > 0 && index <= len(shoppingList.Items) {
                shoppingList.DeleteItem(index - 1)
                fmt.Println("Item deleted successfully!")
            } else {
                fmt.Println("Invalid item number.")
            }

        case 5:
            pkg.SaveToFile(&shoppingList)
            fmt.Println("Shopping list saved. Exiting...")
            return

        default:
            fmt.Println("Invalid choice. Please try again.")
        }
    }
}
