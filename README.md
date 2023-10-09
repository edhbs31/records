# Records
<p align="center">
  <h3 align="center">Records</h3>
  <p align="center">API Golang Records</p>
</p>
## Requirement to Running apps
- Go Version go.17.6
### Installation 
- go run .
#### env run
$env:DSN="jdbc://postgres:admin%3A@localhost:5432/postgres"

#####TABLE 
CREATE TABLE cas_records (
	id serial4 primary key,
	name VARCHAR(255),
  	marks jsonb,
  	created TIMESTAMPTZ
);
##### POSTMAN
GET : /data?startDate=2022-01-01&endDate=2024-01-01&minCount=100&maxCount=13121
POST: /data
payload"{
"name": "x",
"marks": [
	10,11,22,12121,12
]
}
https://api.postman.com/collections/23989743-b2d18ff0-8d7f-4d1c-b0c9-01391ccf369d?access_key=PMAT-01HCB1AVXP4MBZ9NVWPNAPBM92

#####DOCKER
