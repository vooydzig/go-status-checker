Usage
=
1. clone
2. compile 
3. run `./checker config.ini`

Sample output
=
```
Service status check
	Database...error
		pq: database "dbname" does not exist
	Google...OK

```

Config structure
=
Structure for HTTP services
-
```
[Service verbose name]
type=http
url=http://test.service.com
```

Structure for Databases
-
Supported drivers are `postres` and `mysql`
```
[Service verbose name]
type=db
driver=postgres
host=127.0.0.1
port=5432
username=postgres
password=postgres
database=dbname
```