# Gator
----------------------------------

A simple cli tool that aggregates rss feeds and displays the title and description.

You'll need to have [go](https://go.dev/doc/install) and [postgres](https://www.postgresql.org/download/) installed. 

To set up the config file for database and current user, create a json file named 
.gatorconfig.json to your root directory. I used neovim, but feel free to use your text
editor of choice.

`nvim ~/.gatorconfig.json`

Below is the json needed for the config file.  Replace \<username\> with your local username.  If you're not sure what your username is, `whoami` in the terminal will give you your logged in 
username. Leave current_user_name as an empty string.  This will be replaced once you 
log into gator.

`{
 "db_url": "postgres://<username>:@localhost:5432/gator?sslmode=disable",
 "current_user_name": ""
}`

Install gator:

`go install github.com/KarlHavoc/aggregatorCLI@latest`

# Gator Commands
---------------------------------
#### Command list

`command    <arguments>  - Description`

`register   <username>  - Registers a new user with the username`

`login  <username> - Logs in with the username`

`users     - Prints all users to the console`

`addfeed <feedname> <feedURL> - Adds a new feed with a feed name and url`

`feeds  - Prints all feeds and the user that created them to the console`

`follow <url_to_follow>  - Follows a feed for the current user`

`unfollow <url_to_unfollow>  - Unfollows a feed for the current user`

`following - Displays the feeds the current user is following`

`browse <number_of_feeds_to_display> - Prints the title and description of the latest feeds to the console.  If left blank, the number of feeds to display argument defaults to 2.`

`agg <time.Duration>   - Aggregates the feeds the current user is following every time.Duration, i.e. the command - agg 1m - would result in fetching the latest feeds once every minute. %s - seconds, %m - minutes, %h - hours`

