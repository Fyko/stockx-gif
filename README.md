<div align="center">
  <br />
  <p>
    <a href="https://stockx-gif.com"><img src=".github/synth.gif" width="640" alt="stockx-gif icon" /></a>
  </p>
  <br />
  <p>
    <a href="https://github.com/Fyko/stockx-gif/blob/main/go.mod"><img src="https://img.shields.io/github/go-mod/go-version/fyko/stockx-gif" alt="Go version" /></a>
    <a href="https://david-dm.org/fyko/stockx-gif?path=web"> <img src="https://img.shields.io/david/fyko/stockx-gif?path=web" alt="Dependencies" /></a>
    <a href="https://github.com/sponsors/fyko"><img src="https://img.shields.io/github/sponsors/Fyko" alt="Sponsors" /></a>
  </p>
</div>

# About
StockX GIF is a web service for generating GIFs of StockX products which have [360 degree images](https://stockx.com/news/360-degree-images-are-live/).  
The API is written in [Go](https://golang.org) and uses [Go Fiber](https://gofiber.io) for routing.
The front-end is written in [React](https://reactjs.org/) and uses [Next.js](https://nextjs.org/).

# Deployment 

## Deploying backend
Deploying the backend is as simple as cloning the project and running `docker-compose up -d`.  
I would **not** reccomend hosting on Heroku as performance suffers.
## Deploying frontend
You can deploy the frontend to Vercel by clicking the button below:  
[![Deploy with Vercel](https://vercel.com/button)](https://vercel.com/new/git/external?repository-url=https%3A%2F%2Fgithub.com%2FFyko%2Fstockx-gif-next%2Ftree%2Fmain%2Fweb&env=NODE_ENV,API_URL)

Ensure you set the following environment variables:  
* `NODE_ENV` should be `production`
* `API_URL` should be the URL to where the API is running
