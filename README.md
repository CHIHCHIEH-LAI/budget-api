# Budget APIs
Create APIs to manage budgets for various categories

## Functional Requirements
1.	Add a new category with a budget and initial used budget (optional, default to 0).
2.	Update the budget or budget used for an existing category.
3.	View all categories with their budgets and usage.
4.	Delete a category and its associated budget.

## API Design
### Add a New Category
- Endpoint: POST /categories
- Description: Adds a new category with a specified budget and optional initial budget used.
- Request Body:
```
{
    "name": "Groceries",
    "budget": 500,
    "used": 50
}
```
- Response:
    - 201 Created:
    ```
    {
        "id": 1,
        "name": "Groceries",
        "budget": 500,
        "used": 50
    }
    ```
    - 400 Bad Request (if name or budget is invalid).

### Update Budget or Budget Used
- Endpoint: PUT /categories/{category_id}
- Description: Updates the budget or budget used for a category.
- Request Body:
```
{
    "budget": 600,
    "used": 100
}
```
- Response Body:
    - 200 OK
    ```
    {
        "id": 1,
        "name": "Groceries",
        "budget": 600,
        "used": 100
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


## Improvement
1. Use database transaction to ensure strong consistency.
