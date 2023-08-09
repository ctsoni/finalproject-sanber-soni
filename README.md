# Final Project Sanber Go Batch 42 Soni
Simple Hike and Camp Stuff Rental API made with Gin and Postgresql
## Resource
* Railway : https://finalproject-sanber--golang-soni-production.up.railway.app/
* Postman : https://documenter.getpostman.com/view/25656509/2s93CNMtFD
* Slides  : https://docs.google.com/presentation/d/1kTdb4B7ZtmlzIEEz91bV6CD8hW7-bZhhHiBZ-kk4jN8/edit?usp=sharing

## Key Functionalities
1. User Account Management:
* Implemented user registration, enabling users to create accounts with ease.
* Developed functionality for users to delete their accounts and manage their account details.
* Ensured data privacy and security through JWT (JSON Web Tokens) authentication.
2. Role-based Access Control:
* Employed role-based access control to differentiate between admin and normal users.
* Admin users were granted special privileges, allowing them to edit specific tables, such as the inventory table.
3. Inventory Management:
* Designed and maintained an "inventory" table to manage the availability of hiking and camping equipment.
* Admin users could efficiently update the inventory in real-time, ensuring accurate and up-to-date equipment availability.
4. Transaction Handling:
* Implemented a seamless transaction process, enabling users to rent hiking and camping equipment effortlessly.
* Users could initiate transactions and receive detailed receipts with relevant information.
5. Review System:
* Introduced a review system, enabling users to provide feedback and ratings after completing transactions.
* Enhanced user engagement and fostered a sense of community through user-generated reviews.

## ERD
* dbdiagram : https://dbdiagram.io/d/63f49b28296d97641d828f6a
![image](/img/ERDFinalProject.png)

## API
The full documentation of the API can be viewed on [postman documentation](https://documenter.getpostman.com/view/25656509/2s93CNMtFD) above.
### Endpoint Users 
| Method | Path                      | Description                                                                                 | Auth        |
|--------|---------------------------|---------------------------------------------------------------------------------------------|-------------|
| POST   | *`…/users/register`*     | Registering user and token generated. Request body consist of full_name, email and password | token       |
| POST   | *`…/users/login`*         | User login and token generated. Request body consist email and password                     | token       |
| PUT    | *`…/users/edit`*          | User edit profile validated with token                                                      | token       |
| DELETE | *`…/users/delete`*        | User delete profile validated with token                                                    | token       |
| GET    | *`…/users/get-all-users`* | Admin get all users data except password hash                                               | token admin |

### Endpoint Category
| Method | Path                                | Description                                                                                                    | Auth        |
|--------|-------------------------------------|----------------------------------------------------------------------------------------------------------------|-------------|
| POST   | *`…/category/add_`*                 | Admin add new category with request body consist name                                                          | token       |
| PUT    | *`…/category/edit/{category_id}`*   | Admin can edit category name                                                                                   | token       |
| DELETE | *`…/category/delete/{category_id}`* | Admin delete category by its id. The item on inventory table that has deleted category id will be deleted too  | token       |
| GET    | *`…/category/{category_id}/items`*  | All users can get inventories by its category id                                                               |             |

### Endpoint Inventory
| Method | Path                                  | Description                                                                                                                              | Auth        |
|--------|---------------------------------------|------------------------------------------------------------------------------------------------------------------------------------------|-------------|
| POST   | *`…/inventory/add`*                   | Admin can add item into inventory. Request body consist of category id, name, description, availability, stock unit, and price per unit. | token admin |
| PUT    | *`…/inventory/edit/{inventory id}`*   | Admin can edit existed item on inventory by its inventory id                                                                             | token admin |
| DELETE | *`…/inventory/delete/{inventory id}`* | Admin can delete existed item on inventory by its inventory id. The stock on invenotory_stocks table will be deleted too.                | token admin |
| GET    | *`…/inventory/get-all`*               | All user can get all inventories                                                                                                         |             |
| GET    | *`…/inventory/get/{category id}`*     | All user can get all inventories by its category id                                                                                      |             |

### Endpoint Transaction
*`.../users/transaction`*
| Method | Path                                                    | Description                                                                                                                                                                                                       | Auth        |
|--------|---------------------------------------------------------|-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|-------------|
| POST   | *`.../create`*                                          | User can create transaction, every transaction made has expiration, 5 minutes after its made. The transaction will cut the item stock automatically.                                                              | token       |
| PUT    | *`…/{transaction id}?action={pay/cancel}`*              | User can make updates to their transaction status. If user decides to cancel the transaction, the item stock will back up automatically. If the transaction already expire, the stock also back up automatically. | token       |
| PUT    | *`…/admin/{transaction id}`*                            | Admin can update transaction with paid status to retrieve the item stock                                                                                                                                          | token admin |
| GET    | *`…/get-all`*                                           | User can get all their transactions history                                                                                                                                                                       | token       |
| GET    | *`…/get?status={Failed, Unpaid, Paid, Canceled}`*       | User can get all their transactions history filtered by its status                                                                                                                                                | token       |
| GET    | *`…/admin/get-all`*                                     | Admin can get all users transactions history                                                                                                                                                                      | token admin |
| GET    | *`…/admin/get?status={Failed, Unpaid, Paid, Canceled}`* | Admin can get all users transactions history filtered by its status                                                                                                                                               | token admin |

### Endpoint Review
| Method | Path                              | Description                                                                                                                 | Auth  |
|--------|-----------------------------------|-----------------------------------------------------------------------------------------------------------------------------|-------|
| POST   | *`…/add/{transaction id}`*        | User can make review by its transaction id after making transaction.  Request body consist of review text and rating number | token |
| PUT    | *`…/edit/{review id}`*            | User can edit review their made before                                                                                      | token |
| DELETE | *`…/delete/{review id}`*          | User can delete their review                                                                                                | token |
| GET    | *`…/get-all`*                     | All user can get all reviews                                                                                                |       |
| GET    | *`…/get-by-user/{user id}`*       | All user can get all reviews specify by its user id                                                                         |       |
| GET    | *`…/get-by-inven/{inventory id}`* | All user can get all reviews specify by its inventory id                                                                    |       |
