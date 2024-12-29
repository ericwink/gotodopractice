package utils

import (
	"fmt"
	"todoServer/types"
)

// capitalizing the first letter will export the function
func ValidateTodoFields(payload map[string]interface{}) error {
	requiredFields := []string{"ID", "Title", "Body", "IsCompleted"} // List of required fields

	for _, field := range requiredFields {
		if _, exists := payload[field]; !exists { // declare vars value (unused) and bool exists in if, then set the if condition (if it doesn't exist)
			return fmt.Errorf("missing required field: %s", field)
		}
	}
	return nil
}

// returns a pointer to the actual todo in the array, not a copy
func FindTodoById(id float64, todos []types.Todo) (*types.Todo, int) {
	for index := range todos {
		if todos[index].ID == id {
			return &todos[index], index
		}
	}
	return nil, -1
}

func ValidateAndParseJSON(payloadData map[string]interface{}) (types.Todo, error) {
	err := ValidateTodoFields(payloadData)
	if err != nil {
		return types.Todo{}, err
	} // return empty todo, and the error

	parsedTodo := types.Todo{
		ID:          payloadData["ID"].(float64),       // Assert the type
		Title:       payloadData["Title"].(string),     // Assert the type
		Body:        payloadData["Body"].(string),      // Assert the type
		IsCompleted: payloadData["IsCompleted"].(bool), // Assert the type
	}

	return parsedTodo, nil
}

func RemoveByID(todos []types.Todo, id float64) ([]types.Todo, error) {
	todo, index := FindTodoById(id, todos)

	if todo == nil || index == -1 {
		return todos, fmt.Errorf("no todo found with id %v", id)
	}

	return append(todos[:index], todos[index+1:]...), nil
}
