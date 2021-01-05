# Go API Basic with struct only

## Deploy to Heroku

Before you deploy to heroku you must change port from `port := "80"` to `port := os.Getenv("PORT")`

`git init` initialization git repository

`heroku create -b https://github.com/heroku/heroku-buildpack-go.git` create heroku and add buildpack go lang

`git add .` add all file to stage

`git commit -m "your massage"` make tag or massage to your file

`git push heroku master` add push to heroku server