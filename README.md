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
    - 400 Bad Request (if name, budget or expense is invalid).
    - 409 Conflict (if category name already exists)

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
    - 400 Bad Request (if budget is invalid).
    - 404 Not Found (if the category ID does not exist).

### Update Category Expense
- Endpoint: PUT /categories/{category_id}/expense
- Description: Updates the expense for a category.
- Request Body:
```
{
    "expense": 100
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
    - 400 Bad Request (if expense is invalid).
    - 404 Not Found (if the category ID does not exist).

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

## Deploy
### Prerequisites
1. Install Docker Desktop and ensure the Kubernetes option is enabled.
2. Install kubectl 
```
brew install kubectl
```
3. Verify Kubernetes setup 
```
kubectl cluster-info
```
4. Create a local Docker image 
```
docker build -t budget-api:latest .
```

### Apply Kubernetes Configuration
1. Prepare the YAML files:
    - `deployment.yaml`: Defines how the app should be deployed.
	- `service.yaml`: Exposes the app on localhost.
2. Apply the configurations:
```
kubectl apply -f k8s/deployment.yaml
kubectl apply -f k8s/service.yaml
```

### Verify the Deployment
1. Check the status of the pods:
```
kubectl get pods
```
2. Check the status of the service:
```
kubectl get services
```

### Access the Application
If you used the NodePort type in service.yaml, access the application via the node’s port:
1. Find the assigned NodePort:
```
kubectl get service budget-api-service
```
Example output:
```
NAME                  TYPE       CLUSTER-IP      EXTERNAL-IP   PORT(S)          AGE
budget-api-service    NodePort   10.96.123.45    <none>        80:31234/TCP     5m
```
In this case, the NodePort is 31234.

2. Access the application in your browser or via curl:
```
curl http://localhost:31234
```
