# Final Project Sanber Go Batch 42 Soni
Simple Hike and Camp Stuff Rental API made with Gin and Postgresql
## Resource
* Railway : https://finalproject-sanber--golang-soni-production.up.railway.app/
* Postman : https://documenter.getpostman.com/view/25656509/2s93CNMtFD
* Slides  : https://docs.google.com/presentation/d/1kTdb4B7ZtmlzIEEz91bV6CD8hW7-bZhhHiBZ-kk4jN8/edit?usp=sharing

## ERD
* dbdiagram : https://dbdiagram.io/d/63f49b28296d97641d828f6a
{{ url img erd di sini}

## API
The full documentation of the API can be viewed on [postman documentation](https://documenter.getpostman.com/view/25656509/2s93CNMtFD) above.
### Endpoint Users 
| Method | Path                      | Description                                                                                 | Auth        |
|--------|---------------------------|---------------------------------------------------------------------------------------------|-------------|
| POST   | `_…/users/register_ `     | Registering user and token generated. Request body consist of full_name, email and password | token       |
| POST   | `_…/users/login_`         | User login and token generated. Request body consist email and password                     | token       |
| PUT    | `_…/users/edit_`          | User edit profile validated with token                                                      | token       |
| DELETE | `_…/users/delete_`        | User delete profile validated with token                                                    | token       |
| GET    | `_…/users/get-all-users_` | Admin get all users data except password hash                                               | token admin |

### Endpoint Category
| Method | Path                                | Description                                                                                                    | Auth        |
|--------|-------------------------------------|----------------------------------------------------------------------------------------------------------------|-------------|
| POST   | `_…/category/add_ `                 | Admin add new category with request body consist name                                                          | token       |
| PUT    | `_…/category/edit/{category_id}_`   | Admin can edit category name                                                                                   | token       |
| DELETE | `_…/category/delete/{category_id}_` | Admin delete category by its id. The item on inventory table that has deleted category id will be deleted too  | token       |
| GET    | `_…/category/{category_id}/items_`  | All users can get inventories by its category id                                                               |             |

### Endpoint Inventory
| Method | Path                                | Description                                                                                                                              | Auth        |
|--------|-------------------------------------|------------------------------------------------------------------------------------------------------------------------------------------|-------------|
| POST   | _…/inventory/add_                   | Admin can add item into inventory. Request body consist of category id, name, description, availability, stock unit, and price per unit. | token admin |
| PUT    | _…/inventory/edit/{inventory id}_   | Admin can edit existed item on inventory by its inventory id                                                                             | token admin |
| DELETE | _…/inventory/delete/{inventory id}_ | Admin can delete existed item on inventory by its inventory id. The stock on invenotory_stocks table will be deleted too.                | token admin |
| GET    | _…/inventory/get-all_               | All user can get all inventories                                                                                                         |             |
| GET    | _…/inventory/get/{category id}_     | All user can get all inventories by its category id                                                                                      |             |

### Endpoint Transaction
`_.../users/transaction_`
| Method | Path                                                    | Description                                                                                                                                                                                                       | Auth        |
|--------|---------------------------------------------------------|-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|-------------|
| POST   | `_.../create_`                                          | User can create transaction, every transaction made has expiration, 5 minutes after its made. The transaction will cut the item stock automatically.                                                              | token       |
| PUT    | `_…/{transaction id}?action={pay/cancel}_`              | User can make updates to their transaction status. If user decides to cancel the transaction, the item stock will back up automatically. If the transaction already expire, the stock also back up automatically. | token       |
| PUT    | `_…/admin/{transaction id}_`                            | Admin can update transaction with paid status to retrieve the item stock                                                                                                                                          | token admin |
| GET    | `_…/get-all_`                                           | User can get all their transactions history                                                                                                                                                                       | token       |
| GET    | `_…/get?status={Failed, Unpaid, Paid, Canceled}_`       | User can get all their transactions history filtered by its status                                                                                                                                                | token       |
| GET    | `_…/admin/get-all_`                                     | Admin can get all users transactions history                                                                                                                                                                      | token admin |
| GET    | `_…/admin/get?status={Failed, Unpaid, Paid, Canceled}_` | Admin can get all users transactions history filtered by its status                                                                                                                                               | token admin |

### Endpoint Review
| Method | Path                              | Description                                                                                                                 | Auth  |
|--------|-----------------------------------|-----------------------------------------------------------------------------------------------------------------------------|-------|
| POST   | `_…/add/{transaction id}_`        | User can make review by its transaction id after making transaction.  Request body consist of review text and rating number | token |
| PUT    | `_…/edit/{review id}_`            | User can edit review their made before                                                                                      | token |
| DELETE | `_…/delete/{review id}_`          | User can delete their review                                                                                                | token |
| GET    | `_…/get-all_`                     | All user can get all reviews                                                                                                |       |
| GET    | `_…/get-by-user/{user id}_`       | All user can get all reviews specify by its user id                                                                         |       |
| GET    | `_…/get-by-inven/{inventory id}_` | All user can get all reviews specify by its inventory id                                                                    |       |
