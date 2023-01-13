
# MarieJean

Laravel Kit Starter (Actions, Tasks, Repositories, Criteria, Models, Transformers.)

![Alt Text](https://media.giphy.com/media/6wM4Zhs4h4PGo/giphy.gif)

# inspiration

This package was written after 4 days of usage of macOS (Hackintosh).

Last 7 years I used only Linux.

# package

Current solution provides a fast start for developing module structured monolith.
Tested on real-world solutions.

# install

Native (Linux)

1. git clone https://github.com/andrewozarko/mariejean
2. cd mariejean
3. make install

Docker

1. In progress

# commands

1. ```mj version``` - show mj version - done
2. ```mj create <TEMPLATE_NAME> <APP_NAME> --version <APP_VERSION> --force``` - create new php app by template - done
3. ```mj update mj``` - update mj

```mj create ozarko/mariejean app_name```

# templates

1. ```laravel/laravel``` - to create an empty Laravel project
2. ```ozarko/mariejean``` - to create a RestAPI - oriented skeleton, that supports modern programming patterns and development ways.


# laravel post-install

1. ```php artisan sail:install``` - to add docker integration (optional) <br>
   ```./vendor/bin/sail up``` - up docker (or docker compose up) (optional) <br>
   ```./vendor/bin/sail shell``` - use container console for next command execution (optional)
2. ```php artisan migrate``` - run migrations
3. ```php artisan passport:install``` - setup passport
4. Setup generated credentials to .env (Authentication module it require) <br>
   ```CLIENT_WEB_ADMIN_ID=2``` <br>
   ```CLIENT_WEB_ADMIN_SECRET=```
5. Done


**Commands** <br>
1. ```php artisan module:make <module_name>``` - generate new module
2. ```php artisan module:make-controller <controller_name> <module_name>``` - create new controller
3. ```php artisan module:make-model <entity_name> <module_name>``` - create new entity
4. ```php artisan module:make-repository <repository_name> <module_name>``` - create new repository
5. ```php artisan module:make-request <request_name>  <module_name>``` - create new request
5. ```php artisan module:make-action <action_name>  <module_name>``` - create new action
6. ```php artisan module:make-task <tusk_name>  <module_name>``` - create new task
7. ```php artisan module:make-transformer <transformer_name>  <module_name>``` - create new transformer

--------------------------------------------------------------
*  Get validated data from request, paste it to Actions, get back, transform to json and return.
* Call Actions From Controller. Call Tasks from Actions.
* Controller -> []Action -> []Tasks (Controller has a lot of Actions. Action has a lot of tasks).
* Code hasn't limit. It just recommendation. (You can call action from task, but don't do it.)
--------------------------------------------------------------
* You could validate get parameters in Request. Put your param to $urlParameters, and add validation schema.
* Use task extra params and criteria for filtering.
---------------------------------------------------------------
* Don't think about response. Just generate new Transformer. Put your entity, collection to "transform".
* Use ValidationException for returning errors.
* Use transformer includes for dynamic entity relations loading.

# modules

**1. Authentication** <br>

<pre>
curl --location -g --request POST 'http://0.0.0.0:80/api/login' \
--header 'Content-Type: application/json' \
--header 'Accept: application/json, text/plain, */*' \
--data-raw '{
    "email": "andrew@gmail.com",
    "password": "12345678"
}'
</pre>

**2. User** <br>

<pre>
curl --location -g --request POST 'http://0.0.0.0:80/api/registration' \
--header 'Content-Type: application/json' \
--header 'Accept: application/json, text/plain, */*' \
--data-raw '{
    "email": "andrew@gmail.com",
    "password": "12345678",
    "password_repeat": "12345678"
}'
</pre>

# additional packages

<a href="https://github.com/nWidart/laravel-modules">nwidart/laravel-modules</a> - nwidart/laravel-modules is a Laravel package which created to manage your large Laravel app using modules.<br>
<a href="https://laravel.com/docs/9.x/passport">laravel/passport</a> - Laravel Passport provides a full OAuth2 server implementation for your Laravel application in a matter of minutes.<br>
<a href="https://fractal.thephpleague.com/">league/fractal</a> - Fractal provides a presentation and transformation layer for complex data output, the like found in RESTful APIs, and works really well with JSON.<br>
<a href="https://packagist.org/packages/prettus/l5-repository">prettus/l5-repository</a> - Repositories to the database layer. <br>

# additional sources (credits)

<a href="http://apiato.io/docs/9.x/getting-started/software-architectural-patterns/">Porto</a> <br>
<a href="http://apiato.io/">Apiato</a>

# comment

My implementation based on popular packages, all sources code related to my implementation will generate
to folder with laravel. You can control everything. Tested on real-world.

# author
<pre>Andrew O</pre>
