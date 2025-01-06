# Budget APIs
Create APIs to manage budgets for various categories

## Functional Requirements
1.	Add a new category with a budget and initial expense (optional, default to 0).
2.	Update the budget for an existing category.
3.	Update the expense for an existing category.
4.	View all categories with their budgets and expenses.
5.	Delete a category and its associated budget.

## Non-Functional Requirements
1. Strong consistency

## API Design
### Add a New Category
- Endpoint: POST /categories
- Description: Adds a new category with a specified budget and optional initial expense.
- Request Body:
```
{
    "name": "Groceries",
    "budget": 500,
    "expense": 50
}
```
- Response:
    - 201 Created:
    ```
    {
        "id": 1,
        "name": "Groceries",
        "budget": 500,
        "expense": 50
    }
    ```
    - 400 Bad Request (if name or budget is invalid).

### Update Category Budget
- Endpoint: PUT /categories/{category_id}/budget
- Description: Updates the budget for a category.
- Request Body:
```
{
    "budget": 600
}
```
- Response Body:
    - 200 OK
    ```
    {
        "id": 1,
        "name": "Groceries",
        "budget": 600,
        "expense": 100
    }
    ```

### Update Category Expense
- Endpoint: PUT /categories/{category_id}/expense
- Description: Updates the expense for a category.
- Request Body:
```
{
    "expense": 600
}
```
- Response Body:
    - 200 OK
    ```
    {
        "id": 1,
        "name": "Groceries",
        "budget": 600,
        "expense": 100
    }
    ```

### View All Categories
- Endpoint: GET /categories
- Description: Retrieves all categories with their budgets and usages.
- Response Body:
```
[
  {
    "id": 1,
    "name": "Groceries",
    "budget": 600,
    "used": 100
  },
  {
    "id": 2,
    "name": "Entertainment",
    "budget": 200,
    "used": 50
  }
]
```

### Delete a Category
- Endpoint: DELETE /categories/{category_id}
- Description: Deletes a category and its associated budget.
- Response:
	- 204 No Content (success).
	- 404 Not Found (if the category ID does not exist).
