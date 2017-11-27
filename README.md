# Gold Ecstasy

Client et serveur minimalistes en go, communiquant à travers deux conteneurs Docker.
__________________
### Pré-requis
- docker
- docker-compose
- go
__________________
### Installation 
Clonage du dépôt 
```bash
git clone git@github.com:YrenWu/gold-ecstasy.git
```
Test 
```bash
cd gold-ecstasy
docker-compose up -d --build
```
Lire les logs des échanges entre conteneurs 
```bash
docker-compose logs
```
Le serveur est aussi accessible dans le navigateur à:
``` http://localhost:8080 ```
__________________
### Sources
https://golang.org  
https://golang.io/
