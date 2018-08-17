# True Size

Prerequisite
=============
* Add a file named _.env_

  With Following Properties:  
  ```DBHOST```(for e.g.: localhost if running locally)  
  ```DBPORT``` Port that _postgres_ db is running on. Default for _postgres_ is 5432  
  ```DBUSER```User that has access to the db schema.  
  ```PASS``` Password for db user.  
  ```DBNAME```Name of the db your trying to connect to.  

End Points
===========

* Get all sneakers ```/api/sneakers``` HTTP_METHOD: GET  
  + Sample _curl_ Request:  
    ```curl "http://localhost:8000/api/sneakers"```
* Get true size by brand and model ```/api/sneaker/true-size``` HTTP_METHOD: GET  
  + Sample _curl_ Request:  
    ```curl -v -L "http://localhost:8000/api/sneaker/true-size?brand=adidas&sneaker_model=yeezy"```
* Add a sneaker ```/api/sneaker/new``` HTTP_METHOD: POST  
  + Sample _curl_ Request:  
    ```curl -d '{"brand":"adidas", "sneaker_model":"yeezy", "true_size": 1.00}' -H "Content-Type: application/json" -X POST http://localhost:8000/api/sneaker/new```


