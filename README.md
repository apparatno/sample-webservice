# Webservice eksempel

Dette er et enkelt (men også fortenkt) eksempel på en
webservice skrevet i Go.

Det består i hovedsak av tre lag:

1. Transportlag ([main.go](main.go)):
  Dette laget gjør alt som har med transport
  (i dette tilfellet HTTP) å gjøre.
  Det vil si dekode JSON, sende til servicen,
  håndtere feil og mappe tilbake til HTTP og JSON.
2. Servicelag ([pets/service.go](pets/service.go)):
  Her ligger forretningslogikken.
  Servicen har et veldefinert grensesnitt
  (både i form av funksjoner og typer/data)
  og implementerer forretningslogikken som vi prøver
  å løse.
3. Persistenslag ([repository/database.go](repository/database.go)):
  I dette laget foregår all kommunikasjon med hva enn vi
  bruker til å lagre data med.
  I dette eksempelet bruker vi bare en map i minne
  men det kunne også være en database-klient
  eller lignende.

Siden dette er et veldig fortenkt eksempel kan det virke
som om det er for mye mapping mellom typer som ellers er
make til hverandre.
Man kan selv vurdere hvor mye man ønsker å følge dette
paradigmet.
På lenger sikt kan det være veldig nyttig
å holde ting adskilt,
hvis man senere trenger å legge om til gRPC for eksempel.     

Det er kun brukt komponenter fra standardbiblioteket her
men det kan gjøre livet lettere å se på en annen
implementasjon av Mux, for eksempel
[Gorilla Mux](https://www.gorillatoolkit.org/pkg/mux).

## Se også

Bibliotek som kan være til hjelp:

* [Gorilla web toolkit](https://www.gorillatoolkit.org)
* [go-kit](https://gokit.io)
  _har også links til mange andre interessante alternativer_
* [gRPC](https://github.com/grpc/grpc-go)

## Kontakt

Ricco Førgaard <ricco@apparat.no>
eller `@ricco` på GDG Bergen Slack.
