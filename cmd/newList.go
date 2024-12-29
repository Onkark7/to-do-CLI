package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

type Item struct {
	Name     string `json:"name"`
	Quantity int    `json:"quantity"`
}

var lists = make(map[string][]Item)

// newListCmd represents the newList command
var newListCmd = &cobra.Command{
	Use:   "newList",
	Short: "Add a new item to a list",
	Long: `This command allows you to create or update a named list
by adding items with their respective quantities.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 3 {
			fmt.Println("Usage: newList <listName> <itemName> <quantity>")
			return
		}

		listName := args[0]
		itemName := args[1]
		quantity, err := strconv.Atoi(args[2])
		if err != nil {
			fmt.Println("Quantity must be a valid number.")
			return
		}

		// Add the new item to the specified list
		newItem := Item{Name: itemName, Quantity: quantity}
		lists[listName] = append(lists[listName], newItem)

		fmt.Printf("Item '%s' with quantity %d added to list '%s'.\n", itemName, quantity, listName)
	},
}

// showListsCmd represents the showLists command
var showListsCmd = &cobra.Command{
	Use:   "showLists",
	Short: "Show all lists and their items",
	Long:  `This command displays all the lists and their respective items and quantities.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(lists) == 0 {
			fmt.Println("No lists found.")
			return
		}
		for listName, items := range lists {
			fmt.Printf("List: %s\n", listName)
			for _, item := range items {
				// fmt.Printf("  Item: %s, Quantity: %d\n", item.Name, item.Quantity)
				fmt.Println(item)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(newListCmd)
	rootCmd.AddCommand(showListsCmd)
}
