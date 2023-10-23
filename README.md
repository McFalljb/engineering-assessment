# Engineering Challenge

This is a simple API that allows you to create, update, delete, and search for food trucks in San Francisco. It also allows you to get a random food truck from the database. The API is written in Go and uses a Postgres database. The API is containerized using Docker and docker-compose.

## Table of Contents

- [Engineering Challenge](#engineering-challenge)
  - [Table of Contents](#table-of-contents)
  - [Requirements](#requirements)
  - [Getting Started](#getting-started)
  - [Usage](#usage)
    - [Create Food Truck](#create-food-truck)
    - [Update Food Truck](#update-food-truck)
    - [Delete Food Truck](#delete-food-truck)
    - [Get Food Truck by ID](#get-food-truck-by-id)
    - [Search Food Truck by Food Type](#search-food-truck-by-food-type)
    - [Get Random Food Truck](#get-random-food-truck)
    - [KillSwitch](#killswitch)

## Requirements

- Docker
- Docker-compose

## Getting Started

```bash
cd docker
docker-compose up -d --build
```

## Usage

### Create Food Truck

```bash
 curl -X POST http://127.0.0.1:8080/api/v1/foodTruck/create \
  -H "Content-Type: application/json" \
  -d '{"locationid":1569145,"applicant":"Casita Vegana","facilitytype":"Truck","cnn":7553000,"locationdescription":"JOHN MUIR DR: LAKE MERCED BLVD to SKYLINE BLVD (200 - 699)","address":"Assessors Block 7283/Lot004","blocklot":"7283004","block":"7283","lot":"004","permit":"21MFF-00105","status":"APPROVED","fooditems":"Coffee: Vegan Pastries: Vegan Hot Dogs: Vegan Tamales: Te: Vegan Shakes","x":5985417.15,"y":2091453.145,"latitude":37.7218897087084,"longitude":-122.492521244995,"schedule":"http://bsm.sfdpw.org/PermitsTracker/reports/report.aspx?title=schedule&report=rptSchedule&params=permit=21MFF-00105&ExportPDF=1&Filename=21MFF-00105_schedule.pdf","approved":"11/05/2021 12:00:00 AM","received":20211105,"priorpermit":false,"expirationdate":"","location":"(37.72188970870838, -122.4925212449949)","fire_prevention_districts":1,"police_districts":8,"supervisor_districts":4,"zip_codes":64,"neighborhoods_old":14}'
```

### Update Food Truck

```bash
curl -X PUT http://127.0.0.1:8080/api/v1/foodTruck/update \
  -H "Content-Type: application/json" \
  -d '{
      "id": 1,
      "fooditems": ""
  }'
```

### Delete Food Truck

```bash
curl -X POST http://127.0.0.1:8080/api/v1/foodTruck/delete \
  -H "Content-Type: application/json" \
  -d '{"id": 1}'
```

### Get Food Truck by ID

```bash
curl -X POST http://127.0.0.1:8080/api/v1/foodTruck/getByID \
  -H "Content-Type: application/json" \
  -d '{"id": 1}'
```

### Search Food Truck by Food Type

```bash
curl -X POST http://127.0.0.1:8080/api/v1/foodTruck/searchFoodItem \
  -H "Content-Type: application/json" \
  -d '{"fooditems": "sandwiches"}'
```

### Get Random Food Truck

```bash
curl -X POST http://127.0.0.1:8080/api/v1/foodTruck/getRandom
```

### KillSwitch

```bash
curl -X POST http://127.0.0.1:8080/killSwitch
```
