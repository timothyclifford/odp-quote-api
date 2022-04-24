# Quote API

## Endpoints

```
Dashboard:          /
Create quote:       /quote
View/edit quote:    /quote/{id}
Create location:    /quote/{id}/location
View/edit location: /quote/{id}/location/{id}
Create item:        /quote/{id}/item
View/edit item:     /quote/{id}/item/{id}
Quote summary:      /quote/{id}/summary
Quote review:       /quote/{id}/review
```

## Debug

```
curl -X POST -H 'Content-Type:application/json' -d '{"Id":"id", "SalesPersonId":"sales", "Created":0, "FirstName":"Mr", "LastName":"Sells-a-lot", "Email":"sales@quotes.com", "Phone":"000", "Street":"123 Blabla St","Suburb":"Blaville", "Postcode":"3000"}' http://localhost:3000/quotes
```