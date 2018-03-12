# apis
-----

## Steps to create apis or handler functions

- Create one `interface` `[resource]Service` with all methods which had been implemented in [resource] service file. </br>`e.g.`
    ```
    type userService interface{
        Get(rs app.RequestScope, email string) (*models.User, error)
        Count(rs app.RequestScope) (int, error)
        Create(rs app.RequestScope, model *models.User) error
        Update(rs app.RequestScope, email string, model *models.User) (*models.User, error)
        Delete(rs app.RequestScope, email string) error
        // these all methods **must** be defined in the service file
    }
    ```

- Create a Resourse, resourse is basically a entity or we can call it data store, it may be user collection or any collection. </br>`e.g.`
    ```
    type userResource struct {
        service userService 
        // inject userService in the [user] resource
    }    
    ```

- Create a contructor or method to instantiate the service, basically this function will bind all **layers** and defines the **end-points** </br>`e.g.`
    ```
    // ServeUserResource sets up the routing of user endpoints and the corresponding handlers.
    func ServeUserResource(rg *mux.Router, service userService) {
        // Create instance of userResource
        r := &userResource{service}
        rg.HandleFunc("/users/{email}", r.get).Methods("GET")
        rg.HandleFunc("/users", r.create).Methods("POST")
        rg.HandleFunc("/users/{email}", r.update).Methods("PUT")
        rg.HandleFunc("/users/{email}", r.delete).Methods("DELETE")
    }
    ```
    NOTE:
        *this function expects a router object as first param, and service instance in second param*</br>
    This function **should be called in main** </br>`e.g.`
    ```
    // Create a dao, to instantiate the service
    userDao := daos.NewUserDao()
	apis.ServeUserResource(r, services.NewUserService(userDao))
    ```

- Now Create the handler ~~functions~~  Methods, here we will not create a function rather we'll create methods, which will be specific to a resource.
    Each method will have **resource** as the receiver
    > *Basically it will be a handler function with a receiver*

    Method signature:

        `func (rs userResourse) get(w http.ResponseWriter, r *http.Request)` 
    
    &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;`e.g.`
    ```
    func (ur *userResource) create(w http.ResponseWriter, r *http.Request) {
        // Create requestScope obj
        rs := app.GetRequestScope(r)
        var x *models.User
        json.Unmarshal(rs.GetBody(), &x)
        
        // Call any service method using userResource
        err := ur.service.Create(rs, x)

        if err != nil {
            fmt.Fprintf(w, err.Error())
            return
        }

        fmt.Fprintf(w, "Done")
    }
    
    // Similarly we can create varoius handler with the resource as the receiver
    ```    



