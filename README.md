# Gator

A simple cli tool that aggregates rss feeds and displays the title and description.

You'll need to have [go](https://go.dev/doc/install) and [postgres](https://www.postgresql.org/download/) installed. 

To set up the config file for database and current user, create a json file named 
.gatorconfig.json to your root directory. I used neovim, but feel free to use your text
editor of choice.

`nvim ~/.gatorconfig.json`

Put this in the config file, but replace <username> with your local username.  If you're
not sure what your's is, a simple `whoami` in the terminal will give you your logged in 
username.

`{
 "db_url": "postgres://<username>:@localhost:5432/gator?sslmode=disable",
 "current_user_name": ""
}`


