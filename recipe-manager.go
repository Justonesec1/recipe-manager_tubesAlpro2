package main

import "fmt"

type recipe struct {
	name, mainIngredient, difficulty, region string
	ingredients, steps                       [999]string
	duration, ingredientsCount               int
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
	fmt.Scan(&input)

	switch input {
	case 1:
		viewMenu(menu, index)
	case 2:
		addRecipe(menu, index)
	case 3:
		editRecipe(menu, index)
	case 4:
		deleteRecipe(menu, index)
	case 0:
		fmt.Println("-")
	default:
		fmt.Printf("\nPlease input 1-4 or 0!")
		mainMenu(menu, index)
	}
}

// Menampilkan Interface View Menu
func viewMenu(menu *tabMenu, index *int) {
	var input int

	fmt.Printf("\nView Menu\n")
	fmt.Printf("1. View All Menu\n2. Sort Menu\n3. Search by Main Ingredient\n0. Back\n")

	fmt.Print("Choose Option: ")
	fmt.Scan(&input)

	switch input {
	case 1:
		viewAll(menu, index)
		viewMenu(menu, index)
	case 2:
		sortMenu(menu, index)
	case 3:
		searchByMainIngredient(menu, index)
		viewMenu(menu, index)
	case 0:
		mainMenu(menu, index)
	default:
		fmt.Println("Please input 1-3")
		viewMenu(menu, index)
	}
}

// Menampilkan Interface Opsi Sort Menu
func sortMenu(menu *tabMenu, index *int) {
	var input int

	fmt.Printf("\nSort Menu\n")
	fmt.Printf("1. Sort by Name \n2. Sort by Duration \n3. Sort by Rating \n4. Sort by Difficulty \n5. Sort by Region \n6. Sort by Main Ingredient \n0. Back\n")

	fmt.Print("Choose Option: ")
	fmt.Scan(&input)

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
		sortByRegion(menu)
		viewAll(menu, index)
		sortMenu(menu, index)
	case 6:
		sortByMainIngredient(menu)
		viewAll(menu, index)
		sortMenu(menu, index)
	case 0:
		viewMenu(menu, index)
	default:
		fmt.Println("Please input 1-6")
		sortMenu(menu, index)
	}
}

// Menu Sorting Funcions
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

func sortByRegion(menu *tabMenu) {
	var sortIndex, minIndex int
	var temp recipe

	for sortIndex = 0; menu[sortIndex].region != ""; sortIndex++ {

		minIndex = sortIndex

		for i := sortIndex; menu[i].region != ""; i++ {

			// Ascending alphabetical order
			if menu[i].region < menu[minIndex].region {
				minIndex = i
			}
		}

		temp = menu[sortIndex]
		menu[sortIndex] = menu[minIndex]
		menu[minIndex] = temp
	}
}

func sortByMainIngredient(menu *tabMenu) {
	var sortIndex, minIndex int
	var temp recipe

	for sortIndex = 0; menu[sortIndex].mainIngredient != ""; sortIndex++ {

		minIndex = sortIndex

		for i := sortIndex; menu[i].mainIngredient != ""; i++ {

			if menu[i].mainIngredient < menu[minIndex].mainIngredient || (menu[i].mainIngredient == menu[minIndex].mainIngredient && menu[i].name < menu[minIndex].name) {
				minIndex = i
			}
		}

		temp = menu[sortIndex]
		menu[sortIndex] = menu[minIndex]
		menu[minIndex] = temp
	}
}

// Add Recipe Functions
// Meminta user menambah resep
func addRecipe(menu *tabMenu, index *int) {

	fmt.Println("Input string. Type ( . ) to end")
	fmt.Printf("\nRecipe's Name: ")
	menu[*index].name = inputString()

	fmt.Println("Input string. Press {enter} to end")
	fmt.Print("Recipe's region : ")
	fmt.Scan(&menu[*index].region)

	fmt.Println("Input string. Press {enter} to end")
	fmt.Print("Recipe's Main Ingredient: ")
	fmt.Scan(&menu[*index].mainIngredient)

	addIngredients(menu, index)

	countIngredients(menu, index)

	addSteps(menu, index)

	for menu[*index].duration <= 0 {
		fmt.Print("Recipe's Cook Duration (minutes) : ")
		fmt.Scan(&menu[*index].duration)

		if menu[*index].duration <= 0 {
			fmt.Println("Duration must be greater than 0.")
			fmt.Println()
		}
	}

	fmt.Print("Recipe's Difficluty : ")
	fmt.Scan(&menu[*index].difficulty)

	if menu[*index].difficulty != "Easy" && menu[*index].difficulty != "Medium" && menu[*index].difficulty != "Hard" {

		fmt.Println("Please enter Easy, Medium, or Hard.")
		fmt.Println()
	}

	for menu[*index].rating <= 0 && menu[*index].rating >= 5 {

		fmt.Print("Recipe's Rating (0-5): ")
		fmt.Scan(&menu[*index].rating)

		if menu[*index].rating <= 0 && menu[*index].rating >= 5 {
			fmt.Println("Rating must be between 0 and 5.")
		}
	}

	fmt.Print("Favorite? (true/false): ")
	fmt.Scan(&menu[*index].favorite)

	(*index)++

	mainMenu(menu, index)
}

func inputString() string {
	var input, result string

	for input != "." {
		fmt.Scan(&input)
		if input != "." {
			result += input + " "
		}
	}

	return result

}

func chooseDifficulty() string {
	var input int

	fmt.Println("1. Easy")
	fmt.Println("2. Medium")
	fmt.Println("3. Hard")
	fmt.Print("Choose: ")
	fmt.Scan(&input)

	switch input {
	case 1:
		return "Easy"
	case 2:
		return "Medium"
	case 3:
		return "Hard"
	default:
		return chooseDifficulty()
	}

}

//Input Array Ingredients
func addIngredients(menu *tabMenu, index *int) {
	var input string
	var ingredientsIndex int

	fmt.Printf("\n\nPlease Input Recipe's Ingredients\nEnter ( , ) to add next step\nEnter ( . ) to end \n\n")

	for input != "." {
		fmt.Scan(&input)

		if input != "," && input != "." {
			menu[*index].ingredients[ingredientsIndex] += input + " "
		} else {
			ingredientsIndex++
		}

	}
}

// Input Array Steps
func addSteps(menu *tabMenu, index *int) {
	var input string
	var stepsIndex int

	fmt.Printf("\n\nPlease Input Recipe's Steps\nEnter ( , ) to add next step\nEnter ( . ) to end \n\n")

	for input != "." {
		fmt.Scan(&input)

		if input != "," && input != "." {
			menu[*index].steps[stepsIndex] += input + " "
		} else {
			stepsIndex++
		}

	}
}

func countIngredients(menu *tabMenu, index *int) {
	for i := 0; menu[*index].ingredients[i] != ""; i++ {
		menu[*index].ingredientsCount++
	}
}

func searchByMainIngredient(menu *tabMenu, index *int) {
	var ingredient string
	var found bool

	showAvailableIngredients(menu, index)

	fmt.Print("\nInput Main Ingredient: ")
	fmt.Scan(&ingredient)

	fmt.Printf("\nSearch Result for \"%s\"\n", ingredient)
	fmt.Printf("| No | Name                   | Main Ingredient | Region     | Ingredients     | Ingredients Count | Steps                         | Duration | Difficulty     | Rating |\n")

	for i := 0; i < *index; i++ {
		if menu[i].mainIngredient == ingredient {
			found = true

			fmt.Printf("  %-4d %-24s %-17s %-12s %-17s %-19d %-31s %-10d %-16s %-.2f \n",
				i+1,
				menu[i].name,
				menu[i].mainIngredient,
				menu[i].region,
				menu[i].ingredients[0],
				menu[i].ingredientsCount,
				menu[i].steps[0],
				menu[i].duration,
				menu[i].difficulty,
				menu[i].rating)

			printIngredientsAndSteps(menu, i)
		}
	}

	if !found {
		fmt.Println("No recipes found with that main ingredient.")
	}
}

func showAvailableIngredients(menu *tabMenu, index *int) {
	var found bool
	var count int

	fmt.Println("\nAvailable Main Ingredients:")

	for i := 0; i < *index; i++ {

		found = false

		for j := 0; j < i; j++ {
			if menu[i].mainIngredient == menu[j].mainIngredient {
				found = true
			}
		}

		if !found {
			count++
			fmt.Printf("%d. %s\n", count, menu[i].mainIngredient)
		}
	}
}

func editRecipe(menu *tabMenu, index *int) {
	var edit, option int

	viewAll(menu, index)

	fmt.Print("Choose edited recipe: ")
	fmt.Scan(&edit)

	for edit > *index || edit == 0 {
		fmt.Println("No Recipe Selected. Please Try Again")
		fmt.Print("Choose edited recipe: ")
		fmt.Scan(&edit)
	}

	edit -= 1

	fmt.Println("Choose what to edit")
	fmt.Printf("1. Name\n2. Main Ingredient\n3. Region\n4. Ingredients\n5. Steps\n6. Duration\n7. Difficulty\n8. Rating\n0. Back\n")
	fmt.Print("Choose option: ")
	fmt.Scan(&option)

	switch option {
	case 1:
		fmt.Print("New Name: ")
		menu[edit].name = inputString()
	case 2:
		fmt.Print("New Main Ingredient: ")
		fmt.Scan(&menu[edit].mainIngredient)
	case 3:
		fmt.Print("New Region: ")
		fmt.Scan(&menu[edit].region)
	case 4:
		menu[edit].ingredients = [999]string{}
		menu[edit].ingredientsCount = 0
		addIngredients(menu, &edit)
		countIngredients(menu, &edit)
	case 5:
		menu[edit].steps = [999]string{}
		addSteps(menu, &edit)
	case 6:
		fmt.Print("New Duration: ")
		fmt.Scan(&menu[edit].duration)
	case 7:
		fmt.Print("New Difficulty: ")
		fmt.Scan(&menu[edit].difficulty)
	case 8:
		fmt.Print("New Rating: ")
		fmt.Scan(&menu[edit].rating)
	case 0:
	default:
		fmt.Print("Please input 0-8")
		editRecipe(menu, index)
	}

	mainMenu(menu, index)
}

func deleteRecipe(menu *tabMenu, index *int) {
	var deleteIndex int

	viewAll(menu, index)

	if *index == 0 {
		fmt.Println("\nNo Recipe Stored")
		mainMenu(menu, index)
		return
	}

	fmt.Print("Choose Recipe to Delete: ")
	fmt.Scan(&deleteIndex)

	for deleteIndex > *index || deleteIndex == 0 {
		fmt.Println("No Recipe Selected. Please Try Again")
		fmt.Print("Choose Recipe to Delete: ")
		fmt.Scan(&deleteIndex)
	}

	deleteIndex -= 1

	for i := deleteIndex; i < *index-1; i++ {
		menu[i] = menu[i+1]
	}

	menu[*index-1] = recipe{}

	(*index)--

	fmt.Println("Recipe deleted successfully!")

	mainMenu(menu, index)
}

func viewAll(menu *tabMenu, index *int) {
	fmt.Printf("\n| No | Name                   | Main Ingredient | Region     | Ingredients     | Ingredients Count | Steps                         | Duration | Difficulty     | Rating |\n")
	for i := 0; i < *index; i++ {
		fmt.Printf("  %-4d %-24s %-17s %-12s %-17s %-19d %-31s %-10d %-16s %-.2f \n", i+1, menu[i].name, menu[i].mainIngredient, menu[i].region, menu[i].ingredients[0], menu[i].ingredientsCount, menu[i].steps[0], menu[i].duration, menu[i].difficulty, menu[i].rating)
		printIngredientsAndSteps(menu, i)
	}
}

func addTemplateRecipe(menu *tabMenu) {
	menu[0] = recipe{
		name:             "Nasi Goreng",
		mainIngredient:   "Rice",
		difficulty:       "Easy",
		region:           "Indonesia",
		duration:         20,
		ingredientsCount: 5,
		rating:           4.8,
		favorite:         true,
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
		name:             "Sushi",
		mainIngredient:   "Fish",
		difficulty:       "Medium",
		region:           "Japan",
		duration:         45,
		ingredientsCount: 5,
		rating:           4.9,
		favorite:         true,
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
		name:             "Spaghetti Carbonara",
		mainIngredient:   "Pasta",
		difficulty:       "Medium",
		region:           "Italy",
		duration:         30,
		ingredientsCount: 5,
		rating:           4.7,
		favorite:         false,
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
		name:             "Tom Yum",
		mainIngredient:   "Shrimp",
		difficulty:       "Hard",
		region:           "Thailand",
		duration:         40,
		ingredientsCount: 6,
		rating:           4.6,
		favorite:         true,
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
		name:             "Tacos",
		mainIngredient:   "Beef",
		difficulty:       "Easy",
		region:           "Mexico",
		duration:         25,
		ingredientsCount: 5,
		rating:           4.5,
		favorite:         false,
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
		name:             "Burrito",
		mainIngredient:   "Beef",
		difficulty:       "Medium",
		region:           "Mexico",
		duration:         25,
		ingredientsCount: 5,
		rating:           4.5,
		favorite:         false,
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
	menu[6] = recipe{
		name:             "Sushi",
		mainIngredient:   "Fish",
		difficulty:       "Hard",
		region:           "Japan",
		duration:         40,
		ingredientsCount: 5,
		rating:           4.8,
		favorite:         true,
		ingredients: [999]string{
			"Rice",
			"Salmon",
			"Nori",
			"Vinegar",
			"Cucumber",
		},
		steps: [999]string{
			"Cook rice",
			"Season rice with vinegar",
			"Prepare salmon",
			"Place nori on mat",
			"Roll and slice sushi",
		},
	}

	menu[7] = recipe{
		name:             "Spaghetti Carbonara",
		mainIngredient:   "Pasta",
		difficulty:       "Easy",
		region:           "Italy",
		duration:         20,
		ingredientsCount: 5,
		rating:           4.6,
		favorite:         true,
		ingredients: [999]string{
			"Spaghetti",
			"Egg",
			"Parmesan",
			"Bacon",
			"Black Pepper",
		},
		steps: [999]string{
			"Boil spaghetti",
			"Cook bacon",
			"Mix egg and cheese",
			"Combine with pasta",
			"Add pepper and serve",
		},
	}

	menu[8] = recipe{
		name:             "Nasi Goreng",
		mainIngredient:   "Rice",
		difficulty:       "Easy",
		region:           "Indonesia",
		duration:         15,
		ingredientsCount: 5,
		rating:           4.7,
		favorite:         true,
		ingredients: [999]string{
			"Rice",
			"Egg",
			"Garlic",
			"Soy Sauce",
			"Chicken",
		},
		steps: [999]string{
			"Heat oil",
			"Saute garlic",
			"Add chicken and egg",
			"Add rice and soy sauce",
			"Stir fry and serve",
		},
	}

	menu[9] = recipe{
		name:             "Chicken Curry",
		mainIngredient:   "Chicken",
		difficulty:       "Medium",
		region:           "India",
		duration:         35,
		ingredientsCount: 5,
		rating:           4.4,
		favorite:         false,
		ingredients: [999]string{
			"Chicken",
			"Onion",
			"Curry Powder",
			"Tomato",
			"Coconut Milk",
		},
		steps: [999]string{
			"Cook onions",
			"Add chicken",
			"Mix in curry powder",
			"Add tomato and coconut milk",
			"Simmer until cooked",
		},
	}

	menu[10] = recipe{
		name:             "Cheeseburger",
		mainIngredient:   "Beef",
		difficulty:       "Easy",
		region:           "US",
		duration:         20,
		ingredientsCount: 5,
		rating:           4.3,
		favorite:         false,
		ingredients: [999]string{
			"Burger Bun",
			"Beef Patty",
			"Cheese",
			"Lettuce",
			"Tomato",
		},
		steps: [999]string{
			"Cook beef patty",
			"Toast buns",
			"Melt cheese on patty",
			"Assemble ingredients",
			"Serve burger",
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
	return i
}

func main() {
	var menu tabMenu
	var index int

	addTemplateRecipe(&menu)

	index = checkInteger(menu)

	mainMenu(&menu, &index)
}
