### Curl commands for testing the endpoints
---
Make sure to edit the fields inside brackets []

---
- Create a new checkout basket
```bash
curl -d '' http://localhost:3000/Basket
```
- Add a product to a basket
```bash
curl -H "Content-Type: application/json" -d '{"Code":"[PEN|MUG|TSHIRT]","Quantity":"[0-N]"}' http://localhost:3000/Baskets/[id]/items
```
- Get the total amount in a basket
```bash
curl http://localhost:3000/Baskets/[id]
```
- Remove the basket
```bash
curl -X DELETE http://localhost:3000/Baskets/[id]
```
