# True Size

Prerequisite
=============
* Add a file named _.env_

  With Following Properties:
  ```DBHOST```(for e.g.: localhost if running locally)
  ```DBPORT``` Port that _postgres_ db is running on. Default for _postgres_ is 5432
  ```DBUSER```User that has access to the db schema.
  ```PASS``` Password for db user.
  ```DBNAME```Name of the db your trying to connect to. ⋅⋅

End Points
===========

* Get all sneakers ```/api/sneakers``` HTTP_METHOD: GET
* Get true size by brand and model ```/api/sneaker/true-size``` HTTP_METHOD: GET
* Add a sneaker ```/api/sneaker/new``` HTTP_METHOD: POST


