## All Services will go here
----
### Steps to create a service
- First we need to create `[resource]DAO` interface will all methods defined in the `[resource]DAO` file</br>
    `e.g.` 
    ```
    type [resource]Dao interface {
        // Get returns the user with the specified user email.
        Get(rs app.RequestScope, email string) (*models.User, error)
        // ...
        // put all function signatures which we havw written in the [resource]Dao file
    }
    ```    
- Create struct `[Resource]Service` type with member variable `dao`, so that we can inject tha `dao` to the `[Resource]Service` </br>
    `e.g.` 
    ```
    type [Resource]Service struct {
        dao [resource]DAO 
        // basically the interface which we have created in step 1
    }
    ```
- Then create a function which will return ad instance of `[Resource]Service`, with given `[resource]DAO` </br>
    `e.g.`
    ```
    // New[Resource]Service creates a new [Resource]Service with the given resource DAO.
    func New[Resource]Service(dao [resource]DAO) *[Resource]Service {
	    return &[Resource]Service{dao}
    }
    ```
- Now implement all methods which are present in the </br>`[resource]DAO interface`

- Assume we have </br>`Get(rs app.RequestScope, email string) (*models.User, error)`

- So the implementation will be like this
    ```
    // Get returns the user with the specified the user email.
    func (s *[Resource]Service) Get(rs app.RequestScope, email string) (*models.User, error) {
	    return s.dao.Get(rs, email)
    }
    ```
- `[Resource]Service` struct **must** implement the `[resource]DAO` interface

-----