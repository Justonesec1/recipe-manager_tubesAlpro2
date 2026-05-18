package main

import "fmt"

type recipe struct {
	name, mainIngredient, difficulty, region string
	ingredients, steps                       [999]string
	duration, ingretientCount                int
	rating                                   float64
	favorite                                 bool
}

type tabMenu [999]recipe

// Menampilkan Interface Main Menu
func mainMenu(menu *tabMenu, index *int) {
	var input int

	fmt.Printf("\nMenu Manager Application\n")
	fmt.Printf("1. View Menu\n2. Add Recipe\n3. Edit Recipe\n4. Delete Recipe\n0. Exit\n")

	fmt.Print("Choose Option: ")
	fmt.Scanln(&input)

	switch input {
	case 1:
		viewMenu(menu, index)
	case 2:
		addRecipe(menu, index)
	case 3:
		editRecipe(menu, index)
	case 4:
		fmt.Println("4")
	case 0:
		fmt.Println("5")
	default:
		fmt.Println("Please input 1-5")
		mainMenu(menu, index)
	}
}

// Menampilkan Interface View Menu
func viewMenu(menu *tabMenu, index *int) {
	var input int

	fmt.Printf("\nView Menu\n")
	fmt.Printf("1. View All Menu\n2. Sort Menu\n3. Search by Main Ingredient\n0. Back\n")

	fmt.Print("Choose Option: ")
	fmt.Scanln(&input)

	switch input {
	case 1:
		viewAll(menu, index)
		viewMenu(menu, index)
	case 2:
		sortMenu(menu, index)
	case 3:
		fmt.Println("3")
	case 0:
		mainMenu(menu, index)
	default:
		fmt.Println("Please input 1-3")
		viewMenu(menu, index)
	}
}

func sortMenu(menu *tabMenu, index *int) {
	var input int

	fmt.Printf("\nSort Menu\n")
	fmt.Printf("1. Sort by Name \n2. Sort by Duration \n3. Sort by Rating \n4. Sort by Difficulty \n5. Sort by Category \n0. Back\n")

	fmt.Print("Choose Option: ")
	fmt.Scanln(&input)

	switch input {
	case 1:
		sortByName(menu)
		viewAll(menu, index)
		sortMenu(menu, index)
	case 2:
		sortByDuration(menu)
		viewAll(menu, index)
		sortMenu(menu, index)
	case 3:
		sortByRating(menu)
		viewAll(menu, index)
		sortMenu(menu, index)
	case 4:
		sortByDifficulty(menu)
		viewAll(menu, index)
		sortMenu(menu, index)
	case 5:
		fmt.Println("5")
	case 0:
		viewMenu(menu, index)
	default:
		fmt.Println("Please input 1-5")
		sortMenu(menu, index)
	}
}

func sortByName(menu *tabMenu) {
	var sortIndex, minIndex int
	var temp recipe

	for sortIndex = 0; menu[sortIndex].name != ""; sortIndex++ {
		minIndex = sortIndex
		for i := sortIndex; menu[i].name != ""; i++ {
			if menu[i].name < menu[minIndex].name {
				minIndex = i
			}
		}

		temp = menu[sortIndex]
		menu[sortIndex] = menu[minIndex]
		menu[minIndex] = temp
	}

}

func sortByDuration(menu *tabMenu) {
	var sortIndex, minIndex int
	var temp recipe

	for sortIndex = 0; menu[sortIndex].duration != 0; sortIndex++ {
		minIndex = sortIndex
		for i := sortIndex; menu[i].duration != 0; i++ {
			if menu[i].duration < menu[minIndex].duration {
				minIndex = i
			}
		}

		temp = menu[sortIndex]
		menu[sortIndex] = menu[minIndex]
		menu[minIndex] = temp
	}
}

func sortByRating(menu *tabMenu) {
	var sortIndex, minIndex int
	var temp recipe

	for sortIndex = 0; menu[sortIndex].rating != 0; sortIndex++ {
		minIndex = sortIndex
		for i := sortIndex; menu[i].rating != 0; i++ {
			if menu[i].rating < menu[minIndex].rating {
				minIndex = i
			}
		}

		temp = menu[sortIndex]
		menu[sortIndex] = menu[minIndex]
		menu[minIndex] = temp
	}
}

func sortByDifficulty(menu *tabMenu) {

	var tempRecipe recipe
	var sortIndex, minIndex, tempInt int
	var difficultyArr [999]int

	for i := 0; menu[i].difficulty != ""; i++ {
		switch menu[i].difficulty {
		case "Easy":
			difficultyArr[i] = 1
		case "Medium":
			difficultyArr[i] = 2
		case "Hard":
			difficultyArr[i] = 3
		}
	}

	for sortIndex = 0; menu[sortIndex].difficulty != ""; sortIndex++ {
		minIndex = sortIndex
		for i := sortIndex; menu[i].difficulty != ""; i++ {
			if difficultyArr[i] < difficultyArr[minIndex] {
				minIndex = i
			}
		}

		tempRecipe = menu[sortIndex]
		menu[sortIndex] = menu[minIndex]
		menu[minIndex] = tempRecipe

		tempInt = difficultyArr[sortIndex]
		difficultyArr[sortIndex] = difficultyArr[minIndex]
		difficultyArr[minIndex] = tempInt
	}

}

// Meminta user menambah resep
func addRecipe(menu *tabMenu, index *int) {

	fmt.Printf("\nRecipe's Name: ")
	fmt.Scanln(&menu[*index].name)

	fmt.Print("Recipe's region : ")
	fmt.Scanln(&menu[*index].region)

	fmt.Print("Recipe's Main Ingredient: ")
	fmt.Scanln(&menu[*index].mainIngredient)

	fmt.Print("Recipe's Ingredients : ")
	fmt.Scanln(&menu[*index].ingredients[0]) //!!!!!! harus scan array

	/*fmt.Print("Recipe's Ingredients Count : ")
	fmt.Scanln(&menu[*index].ingredients)*/ //bisa otomatis harusnya

	fmt.Print("Recipe's Steps : ")
	fmt.Scanln(&menu[*index].steps[0]) //!!!!!! harus scan array

	fmt.Print("Recipe's Cook Duration : ")
	fmt.Scanln(&menu[*index].duration)

	fmt.Print("Recipe's Difficluty : ")
	fmt.Scanln(&menu[*index].difficulty)

	fmt.Print("Recipe's Rating : ")
	fmt.Scanln(&menu[*index].rating)

	fmt.Print("Favorite? : ")
	fmt.Scanln(&menu[*index].favorite)

	*index++

	mainMenu(menu, index)
}

func editRecipe(menu *tabMenu, index *int) {
	var edit int
	viewAll(menu, index)

	fmt.Print("Choose edited recipe")

	fmt.Scanln(&edit)

	mainMenu(menu, index)
}

func deleteRecipe(menu *tabMenu, index *int) {
	var deleteIndex int

	fmt.Scanln(&deleteIndex)

	menu[deleteIndex] = menu[deleteIndex+1]

	mainMenu(menu, index)
}

func viewAll(menu *tabMenu, index *int) {
	fmt.Printf("\n| No | Name                   | Main Ingredient | Region     | Ingredients     | Ingredients Count | Steps                         | Duration | Difficulty     | Rating |\n")
	for i := 0; i < *index-1; i++ {
		fmt.Printf("  %-4d %-24s %-17s %-12s %-17s %-19d %-31s %-10d %-16s %-.2f \n", i+1, menu[i].name, menu[i].mainIngredient, menu[i].region, menu[i].ingredients[0], menu[i].ingretientCount, menu[i].steps[0], menu[i].duration, menu[i].difficulty, menu[i].rating)
		printIngredientsAndSteps(menu, i)
	}
}

func addTemplateRecipe(menu *tabMenu) {
	menu[0] = recipe{
		name:            "Nasi Goreng",
		mainIngredient:  "Rice",
		difficulty:      "Easy",
		region:          "Indonesia",
		duration:        20,
		ingretientCount: 5,
		rating:          4.8,
		favorite:        true,
		ingredients: [999]string{
			"Rice",
			"Egg",
			"Garlic",
			"Sweet Soy Sauce",
			"Chicken",
		},
		steps: [999]string{
			"Cook garlic",
			"Add chicken and egg",
			"Add rice",
			"Pour sweet soy sauce",
			"Serve hot",
		},
	}

	menu[1] = recipe{
		name:            "Sushi",
		mainIngredient:  "Fish",
		difficulty:      "Medium",
		region:          "Japan",
		duration:        45,
		ingretientCount: 5,
		rating:          4.9,
		favorite:        true,
		ingredients: [999]string{
			"Rice",
			"Salmon",
			"Nori",
			"Vinegar",
			"Soy Sauce",
		},
		steps: [999]string{
			"Prepare sushi rice",
			"Slice salmon",
			"Place nori on mat",
			"Roll sushi tightly",
			"Cut and serve",
		},
	}

	menu[2] = recipe{
		name:            "Spaghetti Carbonara",
		mainIngredient:  "Pasta",
		difficulty:      "Medium",
		region:          "Italy",
		duration:        30,
		ingretientCount: 5,
		rating:          4.7,
		favorite:        false,
		ingredients: [999]string{
			"Spaghetti",
			"Egg",
			"Cheese",
			"Bacon",
			"Pepper",
		},
		steps: [999]string{
			"Boil spaghetti",
			"Cook bacon",
			"Mix egg and cheese",
			"Combine all ingredients",
			"Add pepper and serve",
		},
	}

	menu[3] = recipe{
		name:            "Tom Yum",
		mainIngredient:  "Shrimp",
		difficulty:      "Hard",
		region:          "Thailand",
		duration:        40,
		ingretientCount: 6,
		rating:          4.6,
		favorite:        true,
		ingredients: [999]string{
			"Shrimp",
			"Mushroom",
			"Lemongrass",
			"Chili",
			"Lime",
			"Fish Sauce",
		},
		steps: [999]string{
			"Boil water",
			"Add herbs",
			"Add shrimp and mushroom",
			"Season soup",
			"Serve warm",
		},
	}

	menu[4] = recipe{
		name:            "Tacos",
		mainIngredient:  "Beef",
		difficulty:      "Easy",
		region:          "Mexico",
		duration:        25,
		ingretientCount: 5,
		rating:          4.5,
		favorite:        false,
		ingredients: [999]string{
			"Tortilla",
			"Beef",
			"Lettuce",
			"Cheese",
			"Tomato",
		},
		steps: [999]string{
			"Cook beef",
			"Prepare vegetables",
			"Heat tortilla",
			"Fill tortilla",
			"Serve tacos",
		},
	}

	menu[5] = recipe{
		name:            "Burrito",
		mainIngredient:  "Meat",
		difficulty:      "Medium",
		region:          "Mexico",
		duration:        25,
		ingretientCount: 5,
		rating:          4.5,
		favorite:        false,
		ingredients: [999]string{
			"Tortilla",
			"Beef",
			"Lettuce",
			"Cheese",
			"Tomato",
		},
		steps: [999]string{
			"Cook beef",
			"Prepare vegetables",
			"Heat tortilla",
			"Put Ingredient",
			"Wrap with Tortila",
		},
	}
}

func printIngredientsAndSteps(menu *tabMenu, index int) {
	for i := 1; menu[index].steps[i] != "" || menu[index].ingredients[i] != ""; i++ {
		fmt.Printf(" %-61s %-17s %-19s %-31s\n", "", menu[index].ingredients[i], "", menu[index].steps[i])
	}
	fmt.Println("")
}

func checkInteger(menu tabMenu) int {
	var i int
	for i = 0; menu[i].name != ""; i++ {
	}
	return i + 1
}

func main() {
	var menu tabMenu
	var index int

	addTemplateRecipe(&menu)

	index = checkInteger(menu)

	mainMenu(&menu, &index)
}
