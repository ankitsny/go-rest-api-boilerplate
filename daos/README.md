# daos
-----

## Steps to create a dao

- Create an empty struct of type `[Resource]DAO`. </br> `e.g.`
    ```
    type [Resource]DAO struct{}
    ```
- Create a getter for this `dao` so that we can create instance from other packages.</br>`e.g.`
    ```
    func New[Resource]Dao() *[Resource]DAO {
	    return &[Resource]DAO{}
    }
    // we need to create an instance of dao, because we will inject the same dao in the service to create a service 
    ```
- Declare and define all functions to interact with the DB. </br>`e.g.`
    ```
    // Get reads the user with the specified email from the database.
    // receiver : empty struct which we have created in step #1
    func (dao *[Resource]DAO) Get(rs app.RequestScope, email string) (*models.User, error) {
        var user models.User
        err := rs.DB().C("users").Find(bson.M{"email": email}).One(&user)
        return &user, err
    }
    ```
----


