# proxyservice

*ProxyService* è un microservizio scritto in [Go](https://golang.org) atto a comunicare con ASL Regionali per i quali sono stete imposte restrizioni relative all'indirizzo IP del chiamante.

Le chiamate HTTP di questo microservizio non sono autenticate ma l'autenticazione è necessaria poiché viene trasmessa al server remoto.


- `POST /proxy`
Imposta una particolare chiamata nel quale sono richiesti l'url, il metodo e il body da passare al server remoto.


## API Documentation

Il progetto prevede un sistema automatico di generazione della documentazione rispettando lo standard OpenAPI 2.0 (Swagger), mediante il tool [Swaggo](https://github.com/swaggo/swag).

La cartella `docs` contiene  la documentazione generata. Non dovrebbe mai essere modificata manualmente, ma sempre generata tramite i seguenti comandi:

```
swag fmt
swag init -ot go,json
```

Una volta online, il server esporrà l'interfaccia Swagger UI al path: `{baseUrl}/docs/`

Di seguito i collegamenti per tutti gli ambienti:

- [LOCAL](http://localhost:8000/docs/index.html)
- [DEVELOPMENT](https://dev.gayadeed.it/proxyservice/docs/index.html)
- [STAGING](https://staging.gayadeed.it/proxyservice/docs/index.html)
- [DEMO](https://demo.gayadeed.it/proxyservice/docs/index.html)
- [PRODUCTION](https://app.gayadeed.it/proxyservice/docs/index.html)

### Postman

È possibile accedere agli endpoints tramite Postman:

[![Run in Postman](https://run.pstmn.io/button.svg)](https://app.getpostman.com/run-collection/3273428-d269888f-7633-401f-9341-45d5abb9fb9a?action=collection%2Ffork&collection-url=entityId%3D3273428-d269888f-7633-401f-9341-45d5abb9fb9a%26entityType%3Dcollection%26workspaceId%3D6bd1ee9d-b7f9-436c-9ffa-d6452576c9d2)

## Build

Per ottenere l'eseguibile del server sulla tua macchina basta usare il seguente comando:
```
go build -o proxyservice
```

### Docker

È disponibile un Dockerfile per costruire l'immagine Docker del progetto. Per quanto riguarda la configurazione delle
variabili d'ambiente, basta creare un nuovo file sotto la directory `env` chiamato `.env.local` contenente
i valori di produzione, prima di effettuare la build effettiva mediante i seguenti comandi:
```
docker build -t proxyservice/server .
docker run -p 5000:5000 --name proxyservice proxyservice/server
```
Questo presuppone un'istanza di MongoDB già attiva e pronta all'uso.

In alternativa, è possibile adoperare `docker-compose` per costruire il container per il server e per MongoDB. I servizi
disponibili sono `proxyservice` e `mongo`. Ovviamente **vanno configurati i path correttamente nel file `docker-compose.yml`**.
```
docker-compose up -d --build proxyservice mongo
```
