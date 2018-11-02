# go-api-targomo

Go: API that handles requests for isochrones from Targomo.com API

__*This is a work in progress*__

This is a Golang API that takes care of simple drive-time (isochrone) requests.  

- I will be using this with other projects.  
- I have no intention on reproducing all of the options available for a Targomo REST API request so the funcitonality will be narrow in scope.
- The API returns JSON in text format of only the geojson portion of the json that Targomo returns.
- API requests are limited to North America.

__*Deployment:*__ *http://zotact1.ddns.net:8001/v1/targomo-isochrone/{lng}/{lat}/{time}/{key}*

- __*lng*__ => longitude (decimal degrees)
- __*lat*__ => latitude (decimal degrees)
- __*time*__ => drive time polygon in seconds
- __*key*__ => targomo key