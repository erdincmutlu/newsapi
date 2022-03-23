# News API

Usage
## To start
To start DB
```
cd resources/docker
docker compose up
```
to start the API
```
go run . -emailpass=<EmailPassword> -dbpass=<DBPassword>
```

## To get list of news

```
curl "http://localhost:8080/news?provider=bbc&category=tech"
```
GET /news

provider: `bbc` or `sky`

category: `uk` or `tech`


## To share a news

```
curl --location --request POST 'http://localhost:8080/share' \
--header 'Content-Type: application/json' \
--data-raw '{
    "id": "60836962",
    "action": "email",
    "recipient": "test@example.com"
}'
```
POST /share

with data
```
{
   "id": "60836962",
   "action": "email",
   "recipient": "test@example.com"
}
```
