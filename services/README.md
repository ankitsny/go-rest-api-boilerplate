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

# Test

## Steps to write test for the service

> For assertion assert `"github.com/stretchr/testify/assert"` package has been used.

- Mock `[resource]DAO` by creating a `struct` with list of records. </br>`e.g.`
    ```
    // for user resource
    type mockUserDAO struct {
	    records []models.User
    }
    ```
- Create a function, with return type of `[resource]DAO`, this function should return instance of `[resource]DAO`. </br>`e.g.`
    ```
    // Data store
    func newMockUserDAO() userDAO {
        return &mockUserDAO{
            records: []models.User{
                {ID: "5a947f3a14032d3b384b0829", FirstName: "aaa", Email: "anks@anks.com"},
                {ID: "5a947f3a14032d3b384b0829", FirstName: "bbb", Email: "anso@ankso.com"},
                {ID: "5a947f3a14032d3b384b0829", FirstName: "ccc", Email: "yeah@yess.com"},
            },
        }
    }
    ```

- Implement all `DAO` interface methods. </br>`e.g.`
    ```
    func (m *mockUserDAO) Get(rs app.RequestScope, email string) (*models.User, error) {
        for _, record := range m.records {
            if record.Email == email {
                return &record, nil
            }
        }
        return nil, errors.New("not found")
    }
    // Implment rest of the methods
    ```

- Then test the service `newUserService`.</br>`e.g.`
    ```
    func TestNewUserService(t *testing.T) {
        dao := newMockUserDAO()
        s := NewUserService(dao)
        assert.Equal(t, dao, s.dao)
    }
    ```

- Then write test cases for the all service functions `TestNewUserService_(Get|Create|Count|Delete|...)` </br> `e.g.`
    ```
    func TestNewUserService_Get(t *testing.T) {
        s := NewUserService(newMockUserDAO())

        // Valid User
        user, err := s.Get(nil, "anks@anks.com")
        if assert.Nil(t, err) && assert.NotNil(t, user) {
            assert.Equal(t, user.FirstName, "aaa")
        }

        user, err = s.Get(nil, "anks1@anks.com")
        assert.NotNil(t, err)
        // Similarly we can write test for all func
    }
    ```

----