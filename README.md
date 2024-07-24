# Puumat Stat API Server

A project to practice building a RESTful API server with Go.

The stats are from [this website](https://finland.wbsc.org/en/events/2024-superbaseball/stats/general/all). I scraped the data from the website (check the [`py-scrape` folder](./py-scraper/)).

## Plan

- [ ] Add endpoints to get stats for individual players
- [ ] Add endpoints for updating and deleting stats (although it doesn't make sense but for the sake of practicing)
- [ ] Package the API server with Docker
- [ ] Package the scraper with Docker
- [ ] Use a production-level database (PostgreSQL, MySQL, etc.)
- [ ] Use Docker Compose to run the database, the API server, and the scaper
  - [ ] The scraper should run periodically